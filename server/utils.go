package converter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AllMeasurements interface {
	HeightFeet | HeightMetric | BodyWeightUnits
}

type PostInputs interface {
	RecipeInput | []RecipeInput | AllMeasurements
}

type PostOutputs interface {
	RecipeOutput | []RecipeOutput | AllMeasurements
}

type TestOutputs interface {
	RecipeOutput | AllMeasurements
}

func HandleGetRequestMarshal(w http.ResponseWriter, r *http.Request, data []RecipeInput, inputFn func(data []RecipeInput) ([]RecipeOutput, error)) {
	w.Header().Set("Content-Type", "application/json")

	result, err := inputFn(data)
	if err != nil {
		log.Printf("Conversion error: %v", err)
		http.Error(w, "Conversion error", http.StatusBadRequest)
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

func HandleGetRequestEncode(w http.ResponseWriter, r *http.Request, data []RecipeInput, inputFn func(data []RecipeInput) ([]RecipeOutput, error)) {
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

func HandlePostRequest[I PostInputs, O PostOutputs](w http.ResponseWriter, r *http.Request, inputFn func(data I) (O, error)) {
	w.Header().Set("Content-Type", "application/json")
	var input I

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Printf("Unable to decode JSON: %v", err)
		http.Error(w, "Unable to decode JSON", http.StatusBadRequest)
		return
	}

	result, err := inputFn(input)
	if err != nil {
		errMsg := fmt.Sprintf("Conversion error: %v", err)
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
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

func CheckGotWant[TO TestOutputs](t testing.TB, got, want TO) {
	t.Helper()
	if got != want {
		t.Errorf(`got %v want %v`, got, want)
	}
}

func CheckError(t testing.TB, err error, expectedErrMsg string) {
	t.Helper()
	if err == nil {
		t.Errorf("got %v want nil", err)
	}
	assert.ErrorContains(t, err, expectedErrMsg)
}

// Returns a float64 rounded to the specified amount of decimal places.
func RoundToCustom(num float64, decimalPlaces int) float64 {
	multiplier := math.Pow(10, float64(decimalPlaces))
	return math.Round((num)*multiplier) / multiplier
}
