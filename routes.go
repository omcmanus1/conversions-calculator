package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/omcmanus1/converter/data"
	"github.com/omcmanus1/converter/types"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, welcome to the recipe converter..."))
}

func GetHandlerMarshal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := json.MarshalIndent(Flow(data.Input), "", " ")
	if err != nil {
		log.Println("unable to encode JSON")
	}
	w.Write([]byte(result))
}

func GetHandlerEncode(w http.ResponseWriter, r *http.Request) {
	EncodeJSON(w, r, data.Input, Flow)
}

func PostConversions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input []types.Input
	
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("unable to decode JSON")
	}
	result, err := json.MarshalIndent(Flow(input), "", "  ")
	if err != nil {
		log.Println("unable to encode JSON")
	}
	w.Write([]byte(result))
}
