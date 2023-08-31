package main

import (
	"errors"
	"fmt"

	"github.com/omcmanus1/converter/types"
)

func Flow(data types.Setup) {
	var result float64

	if data.Amount <= 0 {
		fmt.Println(errors.New("please supply a valid amount"))
		return
	}

	if data.InputSystem == "metric" && data.InputUnit == "millilitres" {
		result = MlConversion(data)
	}

	fmt.Printf("%v %v --> %v %v\n", data.Amount, data.InputUnit, result, data.OutputUnit)
}
