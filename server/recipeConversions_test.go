package main

import (
	"testing"
)

func TestVolumeMetric(t *testing.T) {
	t.Run("invalid input system results in error", func(t *testing.T) {
		inp := RecipeInput{
			Ingredient:  "haribos",
			Type:        "volume",
			InputSystem: "dollars",
			InputUnit:   "litres",
			OutputUnit:  "fluid oz",
			Amount:      24,
		}
		got, err := VolumeMetric(inp)
		want := RecipeOutput{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "please provide a metric volume input")
	})
	t.Run("invalid conversion type results in error", func(t *testing.T) {
		inp := RecipeInput{
			Ingredient:  "sugar",
			Type:        "weight",
			InputSystem: "metric",
			InputUnit:   "millilitres",
			OutputUnit:  "fluid oz",
			Amount:      1000,
		}
		got, err := VolumeMetric(inp)
		want := RecipeOutput{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "please provide a metric volume input")
	})
	t.Run("valid millilitres to fluid oz conversion", func(t *testing.T) {
		inp := RecipeInput{
			Ingredient:  "water",
			Type:        "volume",
			InputSystem: "metric",
			InputUnit:   "millilitres",
			OutputUnit:  "fluid oz",
			Amount:      1000,
		}
		got := RecipeOutput{
			Ingredient: "water",
			OutputUnit: "fluid oz",
			Amount:     33.8,
		}
		result, _ := VolumeMetric(inp)
		CheckGotWant(t, result, got)
	})
	t.Run("valid litres to cups conversion", func(t *testing.T) {
		inp := RecipeInput{
			Ingredient:  "milk",
			Type:        "volume",
			InputSystem: "metric",
			InputUnit:   "litres",
			OutputUnit:  "cups",
			Amount:      2,
		}
		got := RecipeOutput{
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
		input := RecipeInput{
			Ingredient:  "coffee",
			Type:        "volume",
			InputSystem: "usa",
			InputUnit:   "cups",
			OutputUnit:  "millilitres",
			Amount:      2,
		}
		expected := RecipeOutput{
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
		input := RecipeInput{
			Ingredient:  "water",
			Type:        "volume",
			InputSystem: "usa",
			InputUnit:   "gallons",
			OutputUnit:  "litres",
			Amount:      5,
		}
		expected := RecipeOutput{
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
		input := RecipeInput{
			Ingredient: "soda",
			Type:       "volume",
			InputUnit:  "invalid",
			OutputUnit: "millilitres",
			Amount:     1,
		}
		_, err := VolumeUS(input)
		CheckError(t, err, "please provide a US volume input")
	})
	t.Run("InvalidInputType", func(t *testing.T) {
		input := RecipeInput{
			Ingredient: "soda",
			Type:       "blah",
			InputUnit:  "invalid",
			OutputUnit: "millilitres",
			Amount:     1,
		}
		_, err := VolumeUS(input)
		CheckError(t, err, "please provide a US volume input")
	})
	t.Run("InvalidOutputUnit", func(t *testing.T) {
		input := RecipeInput{
			Ingredient:  "juice",
			InputSystem: "usa",
			Type:        "volume",
			InputUnit:   "pints",
			OutputUnit:  "invalid",
			Amount:      1,
		}
		_, err := VolumeUS(input)
		CheckError(t, err, "invalid output unit")
	})
}

func TestWeightMetric(t *testing.T) {
	t.Run("ValidConversionFromKilogramsToPounds", func(t *testing.T) {
		input := RecipeInput{
			Ingredient:  "flour",
			Type:        "weight",
			InputSystem: "metric",
			InputUnit:   "kg",
			OutputUnit:  "lbs",
			Amount:      2,
		}
		expected := RecipeOutput{
			Ingredient: "flour",
			OutputUnit: "lbs",
			Amount:     4.4,
		}
		result, _ := WeightMetric(input)
		CheckGotWant(t, result, expected)
	})
	t.Run("ValidConversionFromGramsToOunces", func(t *testing.T) {
		input := RecipeInput{
			Ingredient:  "sugar",
			Type:        "weight",
			InputSystem: "metric",
			InputUnit:   "grams",
			OutputUnit:  "oz",
			Amount:      500,
		}
		expected := RecipeOutput{
			Ingredient: "sugar",
			OutputUnit: "oz",
			Amount:     17.6,
		}
		result, _ := WeightMetric(input)
		CheckGotWant(t, result, expected)
	})
	t.Run("InvalidInputSystem", func(t *testing.T) {
		input := RecipeInput{
			Ingredient:  "salt",
			Type:        "weight",
			InputSystem: "hectares",
			InputUnit:   "bricks",
			OutputUnit:  "oz",
			Amount:      100,
		}
		_, err := WeightMetric(input)
		CheckError(t, err, "please provide a metric weight input")
	})
	t.Run("InvalidInputType", func(t *testing.T) {
		input := RecipeInput{
			Ingredient:  "salt",
			Type:        "football fields",
			InputSystem: "metric",
			InputUnit:   "bricks",
			OutputUnit:  "oz",
			Amount:      100,
		}
		_, err := WeightMetric(input)
		CheckError(t, err, "please provide a metric weight input")
	})
	t.Run("InvalidInputUnit", func(t *testing.T) {
		input := RecipeInput{
			Ingredient:  "salt",
			Type:        "weight",
			InputSystem: "metric",
			InputUnit:   "bricks",
			OutputUnit:  "oz",
			Amount:      100,
		}
		_, err := WeightMetric(input)
		CheckError(t, err, "invalid input unit: ")
	})
	t.Run("InvalidOutputUnit", func(t *testing.T) {
		input := RecipeInput{
			Ingredient:  "rice",
			Type:        "weight",
			InputSystem: "metric",
			InputUnit:   "kg",
			OutputUnit:  "invalid",
			Amount:      1,
		}
		_, err := WeightMetric(input)
		CheckError(t, err, "invalid output unit: ")
	})
}

func TestWeightUS(t *testing.T) {
	t.Run("ValidConversionFromCupsToGrams", func(t *testing.T) {
		input := RecipeInput{
			Ingredient:  "sugar",
			Type:        "weight",
			InputSystem: "usa",
			InputUnit:   "cups",
			OutputUnit:  "grams",
			Amount:      2,
		}
		expected := RecipeOutput{
			Ingredient: "sugar",
			OutputUnit: "grams",
			Amount:     480,
		}
		result, _ := WeightUS(input)
		CheckGotWant(t, result, expected)
	})
	t.Run("ValidConversionFromPoundsToKilograms", func(t *testing.T) {
		input := RecipeInput{
			Ingredient:  "flour",
			Type:        "weight",
			InputSystem: "usa",
			InputUnit:   "lbs",
			OutputUnit:  "kg",
			Amount:      5,
		}
		expected := RecipeOutput{
			Ingredient: "flour",
			OutputUnit: "kg",
			Amount:     2.3,
		}
		result, _ := WeightUS(input)
		CheckGotWant(t, result, expected)
	})
	t.Run("InvalidSystem", func(t *testing.T) {
		input := RecipeInput{
			Ingredient:  "rice",
			Type:        "volume",
			InputSystem: "alaska",
			InputUnit:   "cups",
			OutputUnit:  "grams",
			Amount:      1,
		}
		_, err := WeightUS(input)
		CheckError(t, err, "please provide a US weight input")
	})
	t.Run("InvalidType", func(t *testing.T) {
		input := RecipeInput{
			Ingredient:  "rice",
			Type:        "volume",
			InputSystem: "usa",
			InputUnit:   "cups",
			OutputUnit:  "grams",
			Amount:      1,
		}
		_, err := WeightUS(input)
		CheckError(t, err, "please provide a US weight input")
	})
}
