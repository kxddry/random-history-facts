package get

import (
    "errors"
    "github.com/go-chi/render"
    resp "github.com/kxddry/random-history-facts/internal/lib/api/response"
    "github.com/kxddry/random-history-facts/internal/lib/logger/sl"
    "github.com/kxddry/random-history-facts/internal/storage"
    "log/slog"
    "net/http"
)

type FactGetter interface {
    Fact() (string, error)
}

func New(log *slog.Logger, fg FactGetter) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "handlers.get.New"
        
        log = log.With(
            slog.String("op", op),
        )
        
        fact, err := fg.Fact()
        
        if errors.Is(err, storage.ErrNoFacts) {
            log.Info("No facts found. Create one.")
            w.WriteHeader(http.StatusNotFound)
            render.JSON(w, r, resp.Error(resp.NotFound, "no facts found"))
            return
        }
        
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            render.JSON(w, r, resp.Error(resp.InternalServerError, "internal server error"))
            log.Error("error getting a fact", sl.Err(err))
            return
        }
        
        log.Debug("generated fact", slog.String("fact", fact))
        w.WriteHeader(http.StatusOK)
        render.JSON(w, r, resp.Info(fact))
        return
    }
}
