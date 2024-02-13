package converter

import (
	"errors"
	"fmt"
	"log"
)

func RecipeList(inp []RecipeInput) ([]RecipeOutput, error) {
	var output []string
	var arrOutput []RecipeOutput
	var err error

	for _, entry := range inp {
		var result RecipeOutput

		if entry.Amount <= 0 {
			return nil, errors.New("please supply a valid amount")
		}

		if entry.InputSystem == "usa" {
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

		structOutput := result
		arrOutput = append(arrOutput, structOutput)
		formattedAnswer := fmt.Sprintf("%v %v ----> %v %v of %v\n", entry.Amount, entry.InputUnit, result.Amount, result.OutputUnit, entry.Ingredient)
		output = append(output, formattedAnswer)

	}
	fmt.Println(output)
	return arrOutput, err
}
