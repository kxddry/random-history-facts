package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/kxddry/random-history-facts/internal/config"
	"github.com/kxddry/random-history-facts/internal/http-server/handlers/get"
	"github.com/kxddry/random-history-facts/internal/http-server/handlers/post"
	mwLogger "github.com/kxddry/random-history-facts/internal/http-server/middleware/logger"
	"github.com/kxddry/random-history-facts/internal/lib/factmatcher"
	"github.com/kxddry/random-history-facts/internal/lib/logger"
	"github.com/kxddry/random-history-facts/internal/lib/logger/sl"
	"github.com/kxddry/random-history-facts/internal/storage/postgres"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init config
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)
	log.Debug("debug messages are enabled")

	// init storage
	store, err := postgres.New(cfg)
	if err != nil {
		log.Error("error connecting to db", sl.Err(err))
		os.Exit(1)
	}

	fm := factmatcher.Fact_Matcher{}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(cors.Handler(newCorsOptions()))

	router.Get("/api/", get.New(log, store))
	router.Post("/api/", post.New(log, store, fm))

	log.Info("Starting HTTP server", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	go func() {
		err = srv.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			log.Error("Failed to start HTTP server", sl.Err(err))
			os.Exit(1)
		}
	}()

	// graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Info("Shutting down HTTP server")
	_ = srv.Shutdown(context.Background())
	log.Info("Shutting down Redis server")
	log.Info("Shutting down SQL connection")
	_ = store.Close()
	log.Info("Application stopped")
}

func newCorsOptions() cors.Options {
	return cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Frontend URL
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}
}
