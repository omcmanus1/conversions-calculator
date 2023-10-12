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
			Amount:     8.4,
		}
		result, _ := VolumeMetric(inp)
		CheckGotWant(t, result, got)
	})
}
