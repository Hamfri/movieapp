package http

import (
	"encoding/json"
	"errors"
	"log"
	"movieapp/metadata/internal/repository"
	"movieapp/metadata/internal/service/metadata"
	"net/http"
)

// Handler defines a movie metadata HTTP handler.
type Handler struct {
	service *metadata.Service
}

// New creates a new movie metadata HTTP handler.
func New(service *metadata.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	m, err := h.service.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if err != nil {
		log.Printf("Repository get error for movie %s: %v\n", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
