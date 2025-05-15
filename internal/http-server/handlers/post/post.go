package post

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	resp "github.com/kxddry/random-history-facts/internal/lib/api/response"
	"github.com/kxddry/random-history-facts/internal/lib/logger/sl"
	val "github.com/kxddry/random-history-facts/internal/lib/validator"
	"github.com/kxddry/random-history-facts/internal/storage"
	"io"
	"log/slog"
	"net/http"
)

type Request struct {
	Fact string `json:"fact"`
}

type FactSaver interface {
	AddFact(matcher FactMatcher, fact string) (int64, error)
}

type Response struct {
	resp.Response
	FactId int64 `json:"id"`
}

type FactMatcher interface {
	IsDuplicateFact(fact1, fact2 string) bool
	Normalize(fact string) string
}

func New(log *slog.Logger, fs FactSaver, fm FactMatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.post.New"

		log = log.With(slog.String("op", op))

		var req Request

		if err := render.DecodeJSON(r.Body, &req); err != nil {
			if errors.Is(err, io.EOF) {
				log.Info("request body is empty")
				w.WriteHeader(http.StatusBadRequest)
				render.JSON(w, r, resp.Error(resp.BadRequest, "request body is empty"))
				return
			}

			log.Error("failed to decode request", sl.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, resp.Error(resp.BadRequest, "failed to decode request"))
			return
		}

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("invalid request", sl.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, resp.ValidationError(validateErr))
			return
		}

		fact := req.Fact

		if err := val.Validate(req.Fact); err != nil {
			log.Info("Fact didn't get validated", slog.String("fact", fact))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, resp.Error(resp.BadRequest, "only cyrillic characters are allowed"))
			return
		}

		id, err := fs.AddFact(fm, fact)

		if err != nil {
			if errors.Is(err, storage.ErrFactAlreadyExists) {
				log.Info("fact already exists", sl.Err(err))
				w.WriteHeader(http.StatusBadRequest)
				render.JSON(w, r, resp.Error(resp.BadRequest, "a similar fact already exists"))
				return
			}
			log.Error("failed to add fact", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error(resp.InternalServerError, "failed to add fact"))
			return
		}

		log.Info("fact added", slog.Int64("id", id), slog.String("fact", fact))
		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, Response{
			Response: resp.OK(),
			FactId:   id,
		})
	}
}
