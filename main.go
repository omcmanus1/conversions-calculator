package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	port := ":8080"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/api", Home)
	r.Get("/api/get-encode", GetHandlerEncode)
	r.Get("/api/get-marshal", GetHandlerMarshal)
	fmt.Println("Listening on port " + port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("failed to start server: %w", err)
	}
}
