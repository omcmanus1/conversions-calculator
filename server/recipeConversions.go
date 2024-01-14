package main

import (
	"errors"
	"math"
)

type RecipeInput struct {
	Ingredient   string
	InputSystem  string
	InputUnit    string
	OutputSystem string
	OutputUnit   string
	Type         string
	Amount       float32
}

type RecipeOutput struct {
	Ingredient string  `json:"ingredient"`
	OutputUnit string  `json:"unit"`
	Amount     float64 `json:"amount"`
}

func VolumeMetric(inp RecipeInput) (RecipeOutput, error) {
	var input float64
	var output float64
	var result = RecipeOutput{
		Ingredient: inp.Ingredient,
		OutputUnit: inp.OutputUnit,
	}

	if inp.InputSystem != "metric" || inp.Type != "volume" {
		return RecipeOutput{}, errors.New("please provide a metric volume input")
	}

	switch inp.InputUnit {
	case "litres":
		input = float64(inp.Amount) * 1000
	case "millilitres":
		input = float64(inp.Amount)
	default:
		return RecipeOutput{}, errors.New("invalid input unit: " + inp.InputUnit)
	}

	switch inp.OutputUnit {
	case "fluid oz":
		output = float64(input / 29.574)
	case "pints":
		output = float64(input / 473.2)
	case "quarts":
		output = float64(input / 946.4)
	case "gallons":
		output = float64(input / 3785)
	case "cups":
		if input < 40 {
			return RecipeOutput{}, errors.New("too small for cups")
		} else {
			output = float64(input) / 236.6
		}
	default:
		return RecipeOutput{}, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	if output < 0.1 {
		result.Amount = float64(math.Round(output*100)) / 100
	} else {
		result.Amount = float64(math.Round(output*10)) / 10
	}

	return result, nil
}

func VolumeUS(inp RecipeInput) (RecipeOutput, error) {
	var output float64
	var result = RecipeOutput{
		Ingredient: inp.Ingredient,
		OutputUnit: inp.OutputUnit,
	}

	if inp.InputSystem != "usa" || inp.Type != "volume" {
		return RecipeOutput{}, errors.New("please provide a US volume input")
	}

	switch inp.InputUnit {
	case "cups":
		output = float64(inp.Amount) * 236.6
	case "gallons":
		output = float64(inp.Amount) * 3785
	case "quarts":
		output = float64(inp.Amount) * 946.4
	case "pints":
		output = float64(inp.Amount) * 473.2
	case "fluid oz":
		output = float64(inp.Amount) * 240
	default:
		return RecipeOutput{}, errors.New("invalid input unit" + inp.InputUnit)
	}

	switch inp.OutputUnit {
	case "millilitres":
		output = math.Round(output)
	case "litres":
		convertedOutput := output / 1000
		output = float64(math.Round(convertedOutput*10)) / 10
	default:
		return RecipeOutput{}, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	result.Amount = output
	return result, nil
}

func WeightMetric(inp RecipeInput) (RecipeOutput, error) {
	var input float64
	var output float64
	var result = RecipeOutput{
		Ingredient: inp.Ingredient,
		OutputUnit: inp.OutputUnit,
	}

	if inp.InputSystem != "metric" || inp.Type != "weight" {
		return RecipeOutput{}, errors.New("please provide a metric weight input")
	}

	switch inp.InputUnit {
	case "kg":
		input = float64(inp.Amount) * 1000
	case "grams":
		input = float64(inp.Amount)
	default:
		return RecipeOutput{}, errors.New("invalid input unit: " + inp.InputUnit)
	}

	switch inp.OutputUnit {
	case "oz":
		output = float64(input / 28.35)
	case "lbs":
		output = float64(input / 453.6)
	case "cups":
		if input < 40 {
			return RecipeOutput{}, errors.New("too small for cups")
		}
		output = input / 250
	default:
		return RecipeOutput{}, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	if output < 0.1 {
		result.Amount = float64(math.Round(output*100)) / 100
	} else {
		result.Amount = float64(math.Round(output*10)) / 10
	}

	return result, nil
}

func WeightUS(inp RecipeInput) (RecipeOutput, error) {
	var output float64
	var result = RecipeOutput{
		Ingredient: inp.Ingredient,
		OutputUnit: inp.OutputUnit,
	}

	if inp.InputSystem != "usa" || inp.Type != "weight" {
		return RecipeOutput{}, errors.New("please provide a US weight input")
	}

	switch inp.InputUnit {
	case "cups":
		output = float64(inp.Amount) * 240
	case "lbs":
		output = float64(inp.Amount) * 453.6
	case "oz":
		output = float64(inp.Amount) * 28.35
	default:
		return RecipeOutput{}, errors.New("invalid input unit: " + inp.InputUnit)
	}

	switch inp.OutputUnit {
	case "grams":
		output = math.Round(output)
	case "kg":
		output = output / 1000
		output = float64(math.Round(output*10)) / 10
	default:
		return RecipeOutput{}, errors.New("invalid output unit: " + inp.OutputUnit)
	}

	if output < 0.1 {
		result.Amount = float64(math.Round(output*100)) / 100
	} else {
		result.Amount = float64(math.Round(output*10)) / 10
	}

	return result, nil
}
