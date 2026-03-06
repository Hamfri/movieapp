package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	httphandler "movieapp/metadata/internal/handler/http" // aliased to avoid name collision with stdlib net/http
	"movieapp/metadata/internal/repository/memory"
	"movieapp/metadata/internal/service/metadata"
	"movieapp/pkg/discovery"
	"movieapp/pkg/discovery/consul"
	"net/http"
	"time"
)

const serviceName = "metadata"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("Starting metadata service on port %d", port)

	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	repo := memory.New()
	service := metadata.New(repo)
	h := httphandler.New(service)
	http.HandleFunc("/metadata", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetMetadata(w, r)
		case http.MethodPut:
			h.PutMetadata(w, r)
		default:
			http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		}
	})
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
