package main

import (
	"log"
	httphandler "movieapp/metadata/internal/handler/http" // aliased to avoid name collision with stdlib net/http
	"movieapp/metadata/internal/repository/memory"
	"movieapp/metadata/internal/service/metadata"
	"net/http"
)

func main() {
	log.Println("Starting movie metadata service")
	repo := memory.New()
	service := metadata.New(repo)
	h := httphandler.New(service)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
