package main

import (
	"fmt"
	"log"
	"net/http"

	lambda "github.com/omcmanus1/converter/aws/lambdas"
)

func main() {
	router := SetupRoutes()
	port := ":8080"
	fmt.Printf("Listening on port https://localhost%v...\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("failed to start server: %w", err)
	}
	lambda.AwsLambdaFunction()
}
