package main

import (
	"log"
	metadatagateway "movieapp/movie/internal/gateway/metadata/http"
	ratinggateway "movieapp/movie/internal/gateway/rating/http"
	httphandler "movieapp/movie/internal/handler/http"
	"movieapp/movie/internal/service/movie"
	"net/http"
)

func main() {
	log.Println("starting movie service")

	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")

	service := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(service)

	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
