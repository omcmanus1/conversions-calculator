package main

import (
	"errors"
	"math"

	"github.com/omcmanus1/converter/types"
)

func VolumeMetric(inp types.Input) (types.Output, error) {
	var input float64
	var output float64
	var result = types.Output{
		Ingredient: inp.Ingredient,
		OutputUnit: inp.OutputUnit,
	}

	if inp.Ingredient == "" {
		return types.Output{}, errors.New("please submit an ingredient name")
	}

	if inp.Type != "volume" {
		return types.Output{}, errors.New("invalid conversion type" + inp.Type)
	}

	if inp.InputSystem != "metric" {
		return types.Output{}, errors.New("invalid input system: " + inp.InputSystem)
	}

	if inp.InputUnit == "litres" {
		input = float64(inp.Amount) * 1000
	} else if inp.InputUnit == "millilitres" {
		input = float64(inp.Amount)
	} else {
		return types.Output{}, errors.New("invalid input unit: " + inp.InputUnit)
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
			return types.Output{}, errors.New("too small for cups")
		}
		output = float64(input) / 236.6
	} else {
		return types.Output{}, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	result.Amount = float64(int(output*10)) / 10

	return result, nil
}

func VolumeUS(inp types.Input) (types.Output, error) {
	var output float64
	var result = types.Output{
		Ingredient: inp.Ingredient,
		OutputUnit: inp.OutputUnit,
	}

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
		return types.Output{}, errors.New("invalid input unit" + inp.InputUnit)
	}

	if inp.OutputUnit == "millilitres" {
		output = math.Round(output)
	} else if inp.OutputUnit == "litres" {
		output = output / 1000
		output = float64(int(output*10)) / 10
	} else {
		return types.Output{}, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	result.Amount = output

	return result, nil
}

func WeightMetric(inp types.Input) (types.Output, error) {
	var input float64
	var output float64
	var result = types.Output{
		Ingredient: inp.Ingredient,
		OutputUnit: inp.OutputUnit,
	}

	if inp.InputUnit == "kg" {
		input = float64(inp.Amount) * 1000
	} else if inp.InputUnit == "grams" {
		input = float64(inp.Amount)
	} else {
		return types.Output{}, errors.New("invalid input unit: " + inp.InputUnit)
	}

	if inp.OutputUnit == "oz" {
		output = float64(input / 28.35)
	} else if inp.OutputUnit == "lbs" {
		output = float64(input / 453.6)
	} else if inp.OutputUnit == "cups" {
		if input < 40 {
			return types.Output{}, errors.New("too small for cups")
		}
		output = input / 250
	} else {
		return types.Output{}, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	result.Amount = float64(int(output*10)) / 10

	return result, nil
}

func WeightUS(inp types.Input) (types.Output, error) {
	var output float64
	var result = types.Output{
		Ingredient: inp.Ingredient,
		OutputUnit: inp.OutputUnit,
	}

	if inp.Type != "weight" {
		return types.Output{}, errors.New("invalid type: " + inp.Type)
	}

	if inp.InputUnit == "cups" {
		output = float64(inp.Amount) * 240
	} else if inp.InputUnit == "lbs" {
		output = float64(inp.Amount) * 453.6
	} else if inp.InputUnit == "oz" {
		output = float64(inp.Amount) * 28.35
	} else {
		return types.Output{}, errors.New("invalid input unit: " + inp.InputUnit)
	}

	if inp.OutputUnit == "grams" {
		output = math.Round(output)
	} else if inp.OutputUnit == "kg" {
		output = output / 1000
		output = float64(int(output*10)) / 10
	} else {
		return types.Output{}, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	result.Amount = output

	return result, nil
}
