package main

import (
	"log"
	httphandler "movieapp/rating/internal/handler/http"
	"movieapp/rating/internal/repository/memory"
	"movieapp/rating/internal/service/rating"
	"net/http"
)

func main() {
	log.Println("starting the rating service")
	repo := memory.New()
	service := rating.New(repo)
	h := httphandler.New(service)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
