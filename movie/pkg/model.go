package model

import model "movieapp/metadata/pkg"

type MovieDetails struct {
	Rating   *float64 `json:"rating omitEmpty"`
	Metadata model.Metadata
}
