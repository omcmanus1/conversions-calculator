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

	result, err := Flow(data.Input)
	if err != nil {
		log.Printf("Conversion error (Flow func): %v", err)
		http.Error(w, "Conversion error (Flow func)", http.StatusBadRequest)
		return
	}
	jsonResult, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Printf("Unable to encode JSON: %v", err)
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResult)
}

func GetHandlerEncode(w http.ResponseWriter, r *http.Request) {
	EncodeJSON(w, r, data.SingleInput, Flow)
}

func PostConversions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input []types.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Printf("Unable to decode JSON: %v", err)
		http.Error(w, "Unable to decode JSON", http.StatusBadRequest)
		return
	}
	result, err := Flow(input)
	if err != nil {
		log.Printf("Conversion error: %v", err)
		http.Error(w, "Conversion error", http.StatusInternalServerError)
		return
	}
	jsonResult, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Printf("Unable to encode JSON: %v", err)
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResult)
}

func PostWeightUS(w http.ResponseWriter, r *http.Request) {
	PostSingleConversion(w, r, WeightUS)
}

func PostVolumeUS(w http.ResponseWriter, r *http.Request) {
	PostSingleConversion(w, r, VolumeUS)
}

func PostWeightMetric(w http.ResponseWriter, r *http.Request) {
	PostSingleConversion(w, r, WeightMetric)
}

func PostVolumeMetric(w http.ResponseWriter, r *http.Request) {
	PostSingleConversion(w, r, VolumeMetric)
}
