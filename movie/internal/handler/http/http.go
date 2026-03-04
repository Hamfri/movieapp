package http

import (
	"encoding/json"
	"errors"
	"log"
	"movieapp/movie/internal/service/movie"
	"net/http"
)

type Handler struct {
	service *movie.Service
}

func New(service *movie.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetMovieDetails(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	details, err := h.service.Get(r.Context(), id)
	if err != nil && errors.Is(err, movie.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Response encode error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
