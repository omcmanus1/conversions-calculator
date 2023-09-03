package main

import (
	"errors"
	"math"

	"github.com/omcmanus1/converter/types"
)

func LiquidConversion(inp types.Setup) (float64, error) {
	var output float64

	if inp.OutputUnit == "fluid oz" {
		output = math.Round(float64(inp.Amount) / 28.413)
		output = math.Round(output)
	} else if inp.OutputUnit == "cups" {
		if inp.Amount < 40 {
			return 0.0, errors.New("too small for cups")
		}
		output = float64(inp.Amount) / 236.6
		output = float64(int(output*10)) / 10
	} else {
		return 0.0, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	return output, nil
}

func SolidConversion(inp types.Setup) (float64, error) {
	var output float64

	if inp.OutputUnit == "oz" {
		output = float64(inp.Amount) / 28.35
	} else if inp.OutputUnit == "cups" {
		if inp.Amount < 40 {
			return 0.0, errors.New("too small for cups")
		}
		output = float64(inp.Amount) / 250
		output = float64(int(output*10)) / 10
	} else {
		return 0.0, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	return output, nil
}
