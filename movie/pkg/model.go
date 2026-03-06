package model

import model "movieapp/metadata/pkg"

type MovieDetails struct {
	Rating   *float64       `json:"rating" validate:"omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
