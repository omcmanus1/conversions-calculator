package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/omcmanus1/converter/data"
)

func main() {
	port := ":8080"
	r := chi.NewRouter()
	r.Get("/api", home)
	r.Get("/api/get-encode", getHandlerEncode)
	r.Get("/api/get-marshal", getHandlerMarshal)
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func getHandlerMarshal(w http.ResponseWriter, r *http.Request) {
	result, err := json.MarshalIndent(Flow(data.Input), "", " ")
	if err != nil {
		log.Println("unable to encode JSON")
	}
	w.Write([]byte(result))
}

func getHandlerEncode(w http.ResponseWriter, r *http.Request) {
	EncodeJSON(w, r, data.Input, Flow)
}
