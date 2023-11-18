package main

import (
	"testing"

	"github.com/omcmanus1/converter/types"
)

func TestVolumeMetric(t *testing.T) {
	t.Run("empty inputs results in error", func(t *testing.T) {
		inp := types.Input{
			Ingredient:   "",
			InputSystem:  "",
			InputUnit:    "",
			OutputSystem: "",
			OutputUnit:   "",
			Type:         "",
			Amount:       0.0,
		}
		got, err := VolumeMetric(inp)
		want := types.Output{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "please submit an ingredient name")
	})
	t.Run("invalid input system results in error", func(t *testing.T) {
		inp := types.Input{
			Ingredient:   "haribos",
			InputSystem:  "dollars",
			InputUnit:    "litres",
			OutputSystem: "US",
			OutputUnit:   "fluid oz",
			Type:         "volume",
			Amount:       24,
		}
		got, err := VolumeMetric(inp)
		want := types.Output{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "invalid input system")
	})
	t.Run("invalid ingredient results in error", func(t *testing.T) {
		inp := types.Input{
			Ingredient:  "",
			Type:        "volume",
			InputSystem: "metric",
			InputUnit:   "millilitres",
			OutputUnit:  "fluid oz",
			Amount:      344,
		}
		got, err := VolumeMetric(inp)
		want := types.Output{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "please submit an ingredient name")
	})
	t.Run("invalid conversion type results in error", func(t *testing.T) {
		inp := types.Input{
			Ingredient:  "sugar",
			Type:        "weight",
			InputSystem: "metric",
			InputUnit:   "millilitres",
			OutputUnit:  "fluid oz",
			Amount:      1000,
		}
		got, err := VolumeMetric(inp)
		want := types.Output{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "invalid conversion type")
	})
	t.Run("valid millilitres to fluid oz conversion", func(t *testing.T) {
		inp := types.Input{
			Ingredient:  "water",
			Type:        "volume",
			InputSystem: "metric",
			InputUnit:   "millilitres",
			OutputUnit:  "fluid oz",
			Amount:      1000,
		}
		got := types.Output{
			Ingredient: "water",
			OutputUnit: "fluid oz",
			Amount:     33.8,
		}
		result, _ := VolumeMetric(inp)
		CheckGotWant(t, result, got)
	})
	t.Run("valid litres to cups conversion", func(t *testing.T) {
		inp := types.Input{
			Ingredient:  "milk",
			Type:        "volume",
			InputSystem: "metric",
			InputUnit:   "litres",
			OutputUnit:  "cups",
			Amount:      2,
		}
		got := types.Output{
			Ingredient: "milk",
			OutputUnit: "cups",
			Amount:     8.5,
		}
		result, _ := VolumeMetric(inp)
		CheckGotWant(t, result, got)
	})
}

func TestVolumeUS(t *testing.T) {
	t.Run("ValidConversionFromCupsToMillilitres", func(t *testing.T) {
		input := types.Input{
			Ingredient: "coffee",
			Type:       "volume",
			InputUnit:  "cups",
			OutputUnit: "millilitres",
			Amount:     2,
		}
		expected := types.Output{
			Ingredient: "coffee",
			OutputUnit: "millilitres",
			Amount:     473,
		}
		result, err := VolumeUS(input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		CheckGotWant(t, result, expected)
	})
	t.Run("ValidConversionFromGallonsToLitres", func(t *testing.T) {
		input := types.Input{
			Ingredient: "water",
			Type:       "volume",
			InputUnit:  "gallons",
			OutputUnit: "litres",
			Amount:     5,
		}
		expected := types.Output{
			Ingredient: "water",
			OutputUnit: "litres",
			Amount:     18.9,
		}
		result, err := VolumeUS(input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		CheckGotWant(t, result, expected)
	})
	t.Run("InvalidInputUnit", func(t *testing.T) {
		input := types.Input{
			Ingredient: "soda",
			Type:       "volume",
			InputUnit:  "invalid",
			OutputUnit: "millilitres",
			Amount:     1,
		}
		_, err := VolumeUS(input)
		CheckError(t, err, "invalid input unit")
	})
	t.Run("InvalidOutputUnit", func(t *testing.T) {
		input := types.Input{
			Ingredient: "juice",
			Type:       "volume",
			InputUnit:  "pints",
			OutputUnit: "invalid",
			Amount:     1,
		}
		_, err := VolumeUS(input)
		CheckError(t, err, "invalid output unit")
	})
}

func TestWeightMetric(t *testing.T) {
	t.Run("ValidConversionFromKilogramsToPounds", func(t *testing.T) {
		input := types.Input{
			Ingredient: "flour",
			Type:       "weight",
			InputUnit:  "kg",
			OutputUnit: "lbs",
			Amount:     2,
		}
		expected := types.Output{
			Ingredient: "flour",
			OutputUnit: "lbs",
			Amount:     4.4,
		}
		result, _ := WeightMetric(input)
		CheckGotWant(t, result, expected)
	})
	t.Run("ValidConversionFromGramsToOunces", func(t *testing.T) {
		input := types.Input{
			Ingredient: "sugar",
			Type:       "weight",
			InputUnit:  "grams",
			OutputUnit: "oz",
			Amount:     500,
		}
		expected := types.Output{
			Ingredient: "sugar",
			OutputUnit: "oz",
			Amount:     17.6,
		}
		result, _ := WeightMetric(input)
		CheckGotWant(t, result, expected)
	})

	t.Run("InvalidInputUnit", func(t *testing.T) {
		input := types.Input{
			Ingredient: "salt",
			Type:       "weight",
			InputUnit:  "invalid",
			OutputUnit: "oz",
			Amount:     100,
		}
		_, err := WeightMetric(input)
		CheckError(t, err, "invalid input unit")
	})
	t.Run("InvalidOutputUnit", func(t *testing.T) {
		input := types.Input{
			Ingredient: "rice",
			Type:       "weight",
			InputUnit:  "kg",
			OutputUnit: "invalid",
			Amount:     1,
		}
		_, err := WeightMetric(input)
		CheckError(t, err, "invalid output unit")
	})
}

func TestWeightUS(t *testing.T) {
	t.Run("ValidConversionFromCupsToGrams", func(t *testing.T) {
		input := types.Input{
			Ingredient: "sugar",
			Type:       "weight",
			InputUnit:  "cups",
			OutputUnit: "grams",
			Amount:     2,
		}
		expected := types.Output{
			Ingredient: "sugar",
			OutputUnit: "grams",
			Amount:     480,
		}
		result, _ := WeightUS(input)
		CheckGotWant(t, result, expected)
	})
	t.Run("ValidConversionFromPoundsToKilograms", func(t *testing.T) {
		input := types.Input{
			Ingredient: "flour",
			Type:       "weight",
			InputUnit:  "lbs",
			OutputUnit: "kg",
			Amount:     5,
		}
		expected := types.Output{
			Ingredient: "flour",
			OutputUnit: "kg",
			Amount:     2.3,
		}
		result, _ := WeightUS(input)
		CheckGotWant(t, result, expected)
	})
	t.Run("InvalidType", func(t *testing.T) {
		input := types.Input{
			Ingredient: "rice",
			Type:       "volume",
			InputUnit:  "cups",
			OutputUnit: "grams",
			Amount:     1,
		}
		_, err := WeightUS(input)
		CheckError(t, err, "invalid type")
	})
}
