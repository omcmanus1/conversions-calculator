package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/omcmanus1/converter/types"
)

func Flow(data types.Setup) {
	var result float64
	var err error

	if data.Amount <= 0 {
		fmt.Println(errors.New("please supply a valid amount"))
	}

	if data.InputUnit == "millilitres" {
		result, err = LiquidConversion(data)
	} else if data.InputUnit == "grams" {
		result, err = SolidConversion(data)
	} else {
		err = errors.New("unsupported input unit" + data.InputUnit)
	}

	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("%v %v --> %v %v\n", data.Amount, data.InputUnit, result, data.OutputUnit)
	}
}
