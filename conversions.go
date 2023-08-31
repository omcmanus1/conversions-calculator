package main

import (
	"errors"
	"math"

	"github.com/omcmanus1/converter/types"
)

func MlConversion(inp types.Setup) float64 {
	floz := float64(inp.Amount) / 28.413
	return floz
}

func GramConversion(inp types.Setup) (float64, error) {
	var output float64

	if inp.OutputUnit == "oz" {
		output = float64(inp.Amount) / 28.35
	} else if inp.OutputUnit == "cups" {
		if inp.Amount < 40 {
			return 0.0, errors.New("too small for cups")
		} else {
			output = float64(inp.Amount) / 250
		}
	} else {
		return 0.0, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	return math.Round(output/0.05) * 0.05, nil
}
