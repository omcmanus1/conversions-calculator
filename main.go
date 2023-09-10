package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/omcmanus1/converter/data"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	result, err := json.MarshalIndent(Flow(data.Input), "", " ")
	if err != nil {
		log.Println("unable to encode JSON")
	}
	w.Write([]byte(result))
}

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/convert", getHandler)
	http.ListenAndServe(":8080", r)
}
