package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/omcmanus1/converter"
)

func main() {
	router := converter.SetupRoutes()
	port := "8090"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
		fmt.Printf("Listening on port https://localhost%v...\n", port)
		err := http.ListenAndServe(port, router)
		if err != nil {
			log.Fatal("failed to start server: %w", err)
		}
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
