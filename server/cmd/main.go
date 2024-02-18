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
	port := ":8080"
	env := os.Getenv("ENVIRONMENT")
	fmt.Println("env:", env)
	if env == "local" {
		fmt.Printf("Listening on port https://localhost%v...\n", port)
		err := http.ListenAndServe(port, router)
		if err != nil {
			log.Fatal("failed to start server: %w", err)
		}
	} else {
		fmt.Println("Running using cloud functions")
		if err := funcframework.Start(port); err != nil {
			log.Fatalf("funcframework.Start: %v\n", err)
		}
	}
}
