package main

import (
	"errors"
	"math"

	"github.com/omcmanus1/converter/types"
)

func VolumeMetric(inp types.Setup) (float64, error) {
	var input float64
	var output float64

	if inp.InputUnit == "litres" {
		input = float64(inp.Amount) * 1000
	} else if inp.InputUnit == "millilitres" {
		input = float64(inp.Amount)
	} else {
		return 0.0, errors.New("invalid input unit: " + inp.InputUnit)
	}

	if inp.OutputUnit == "fluid oz" {
		output = float64(input / 29.574)
	} else if inp.OutputUnit == "pints" {
		output = float64(input / 473.2)
	} else if inp.OutputUnit == "quarts" {
		output = float64(input / 946.4)
	} else if inp.OutputUnit == "gallons" {
		output = float64(input / 3785)
	} else if inp.OutputUnit == "cups" {
		if input < 40 {
			return 0.0, errors.New("too small for cups")
		}
		output = float64(input) / 236.6
	} else {
		return 0.0, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	return float64(int(output*10)) / 10, nil
}

func VolumeUS(inp types.Setup) (float64, error) {
	var output float64

	if inp.InputUnit == "cups" {
		output = float64(inp.Amount) * 236.6
	} else if inp.InputUnit == "gallons" {
		output = float64(inp.Amount) * 3785
	} else if inp.InputUnit == "quarts" {
		output = float64(inp.Amount) * 946.4
	} else if inp.InputUnit == "pints" {
		output = float64(inp.Amount) * 473.2
	} else if inp.InputUnit == "fluid oz" {
		output = float64(inp.Amount) * 240
	} else {
		return 0.0, errors.New("invalid input unit" + inp.InputUnit)
	}

	if inp.OutputUnit == "millilitres" {
		return math.Round(output), nil
	} else if inp.OutputUnit == "litres" {
		output = output / 1000
		return float64(int(output*10)) / 10, nil
	} else {
		return 0.0, errors.New("invalid output unit: " + inp.OutputUnit)
	}
}

func WeightMetric(inp types.Setup) (float64, error) {
	var input float64
	var output float64

	if inp.InputUnit == "kg" {
		input = float64(inp.Amount) * 1000
	} else if inp.InputUnit == "grams" {
		input = float64(inp.Amount)
	} else {
		return 0.0, errors.New("invalid input unit: " + inp.InputUnit)
	}

	if inp.OutputUnit == "oz" {
		output = float64(input / 28.35)
	} else if inp.OutputUnit == "lbs" {
		output = float64(input / 453.6)
	} else if inp.OutputUnit == "cups" {
		if input < 40 {
			return 0.0, errors.New("too small for cups")
		}
		output = input / 250
	} else {
		return 0.0, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	return float64(int(output*10)) / 10, nil
}

func WeightUS(inp types.Setup) (float64, error) {
	var output float64

	if inp.InputUnit == "cups" {
		output = float64(inp.Amount) * 250
	} else if inp.InputUnit == "lbs" {
		output = float64(inp.Amount) * 453.6
	} else if inp.InputUnit == "oz" {
		output = float64(inp.Amount) * 28.35
	} else {
		return 0.0, errors.New("invalid input unit: " + inp.InputUnit)
	}

	if inp.OutputUnit == "grams" {
		output = math.Round(output)
	} else if inp.OutputUnit == "kg" {
		output = output / 1000
		output = float64(int(output*10)) / 10
	}

	return output, nil
}
