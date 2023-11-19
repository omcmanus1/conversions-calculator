package main

import (
	"errors"

	"github.com/omcmanus1/converter/types"
)

func FromFeet(inp types.HeightFeet) (types.HeightCentimetres, error) {
	var output types.HeightCentimetres

	if inp.Feet < 0 || inp.Inches < 0 || inp.Inches >= 12 {
		return output, errors.New("invalid input")
	}
	feet := 30.48 * inp.Feet
	inches := 2.54 * inp.Inches

	output.Centimetres = feet + inches
	return output, nil
}

func FromMetric(inp types.HeightFromMetric) (types.HeightFeet, error) {
	output, err := inp.GetHeightInFeet()
	return output, err
}
