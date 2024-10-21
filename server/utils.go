package converter

import (
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

// Reusable handler function for HTTP POST requests. Decodes the incoming JSON into a generic input type,
// executes a provided input function, and returns a JSON-encoded result or an error.
//
// Accepts generic input types using interface implementation.
//
// If JSON decoding or encoding fails, or if the input function returns
// an error, appropriate errors are logged and returned to the client.
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

// Testing util to compare result with expected output -
// throws error if not equal.
func CheckGotWant[TO TestOutputs](t testing.TB, got, want TO) {
	t.Helper()
	if got != want {
		t.Errorf(`got %v want %v`, got, want)
	}
}

// Testing util for checking expected errors as return type -
// throws new error if err == nil
func CheckError(t testing.TB, err error, expectedErrMsg string) {
	t.Helper()
	if err == nil {
		t.Errorf("got %v want nil", err)
	}
	assert.ErrorContains(t, err, expectedErrMsg)
}

// Returns the input number rounded to the specified amount of decimal places.
func RoundToCustom(num float64, decimalPlaces int) float64 {
	multiplier := math.Pow(10, float64(decimalPlaces))
	return math.Round((num)*multiplier) / multiplier
}
