package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/omcmanus1/converter/types"
)

func Flow(inp []types.Input) ([]types.Output, error) {
	var output []string
	var arrOutput []types.Output
	var err error

	for _, entry := range inp {
		var result float64

		if entry.Amount <= 0 {
			return nil, errors.New("please supply a valid amount")
		}

		if entry.InputSystem == "US" {
			if entry.Type == "volume" {
				result, err = VolumeUS(entry)
			} else if entry.Type == "weight" {
				result, err = WeightUS(entry)
			} else {
				return nil, errors.New("invalid measurement type: " + entry.Type)
			}
		} else if entry.InputSystem == "metric" {
			if entry.Type == "weight" {
				result, err = WeightMetric(entry)
			} else if entry.Type == "volume" {
				result, err = VolumeMetric(entry)
			} else {
				return nil, errors.New("invalid measurement type: " + entry.Type)
			}
		} else {
			return nil, errors.New("invalid system: " + entry.InputSystem)
		}

		if err != nil {
			log.Fatal("Error: ", err)
		}

		structOutput := types.Output{Ingredient: entry.Ingredient, OutputUnit: entry.OutputUnit, Amount: result}
		arrOutput = append(arrOutput, structOutput)
		formattedAnswer := fmt.Sprintf("%v %v ----> %v %v of %v\n", entry.Amount, entry.InputUnit, result, entry.OutputUnit, entry.Ingredient)
		output = append(output, formattedAnswer)

	}
	fmt.Println(output)
	return arrOutput, err
}
