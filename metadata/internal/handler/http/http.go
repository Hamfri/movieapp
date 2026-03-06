package http

import (
	"encoding/json"
	"errors"
	"log"
	metadata "movieapp/metadata/internal/service/metadata"
	model "movieapp/metadata/pkg"
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

	if err != nil {
		switch {
		case errors.Is(err, metadata.ErrNotFound):
			w.WriteHeader(http.StatusNotFound)
		default:
			log.Printf("Repository get error for movie %s: %v\n", id, err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) PutMetadata(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	if description == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	director := r.FormValue("director")
	if director == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	metadata := model.Metadata{
		ID:          id,
		Title:       title,
		Description: description,
		Director:    director,
	}

	ctx := r.Context()
	err := h.service.Put(ctx, &metadata)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
