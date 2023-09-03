package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/omcmanus1/converter/types"
)

func Flow(inp types.Setup) {
	var result float64
	var err error

	if inp.Amount <= 0 {
		fmt.Println(errors.New("please supply a valid amount"))
	}

	if inp.Type == "volume" {
		if inp.InputSystem == "US" {
			result, err = VolumeUS(inp)
		} else {
			result, err = VolumeMetric(inp)
		}
	} else if inp.Type == "weight" {
		if inp.InputSystem == "US" {
			result, err = WeightUS(inp)
		} else {
			result, err = WeightMetric(inp)
		}
	} else {
		err = errors.New("invalid measurement type: " + inp.Type)
	}

	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("%v %v --> %v %v of %v\n", inp.Amount, inp.InputUnit, result, inp.OutputUnit, inp.Ingredient)
	}
}
