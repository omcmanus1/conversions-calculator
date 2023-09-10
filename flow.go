package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/omcmanus1/converter/types"
)

func Flow(inp []types.Input) []types.Output {
	var output []string
	var arrOutput []types.Output

	for _, entry := range inp {
		var result float64
		var err error

		if entry.Amount <= 0 {
			fmt.Println(errors.New("please supply a valid amount"))
		}

		if entry.Type == "volume" {
			if entry.InputSystem == "US" {
				result, err = VolumeUS(entry)
			} else {
				result, err = VolumeMetric(entry)
			}
		} else if entry.Type == "weight" {
			if entry.InputSystem == "US" {
				result, err = WeightUS(entry)
			} else {
				result, err = WeightMetric(entry)
			}
		} else {
			err = errors.New("invalid measurement type: " + entry.Type)
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
	return arrOutput
}
