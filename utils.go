package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/omcmanus1/converter/types"
)

func EncodeJSON(w http.ResponseWriter, r *http.Request, data []types.Input, inputFn func(data []types.Input) ([]types.Output, error)) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	result, err := inputFn(data)
	if err != nil {
		log.Printf("Input function error: %v", err)
		http.Error(w, "Input function error", http.StatusInternalServerError)
		return
	}
	// Encode the data to a buffer
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Indent the JSON data
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, buf.Bytes(), "", "    "); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write the indented JSON data to the response writer
	w.Write(prettyJSON.Bytes())
}
