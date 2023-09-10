package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/omcmanus1/converter/types"
)

func EncodeJSON(w http.ResponseWriter, r *http.Request, data []types.Input, inputFn func(data []types.Input) []types.Output) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the data to a buffer
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(inputFn(data)); err != nil {
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
