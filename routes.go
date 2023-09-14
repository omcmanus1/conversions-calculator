package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/omcmanus1/converter/data"
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
