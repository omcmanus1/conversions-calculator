package main

import (
	"errors"
	"math"
)

type HeightFeet struct {
	Feet   float32 `json:"feet"`
	Inches float32 `json:"inches"`
}

type HeightMetric struct {
	Centimetres float32 `json:"centimetres"`
	Metres      float32 `json:"metres"`
}

func FromFeet(inp HeightFeet) (HeightMetric, error) {
	var output HeightMetric

	switch {
	case inp.Feet == 0 && inp.Inches == 0:
		return output, errors.New("please input a height")
	case inp.Feet < 0 || inp.Inches < 0 || inp.Inches >= 12:
		return output, errors.New("invalid input")
	}

	feet := 30.48 * inp.Feet
	inches := 2.54 * inp.Inches

	output.Centimetres = feet + inches
	output.Metres = (feet + inches) / 100
	return output, nil
}

func FromMetric(inp HeightMetric) (HeightFeet, error) {
	var output HeightFeet
	var totalInches float64

	switch {
	case inp.Centimetres == 0 && inp.Metres == 0:
		return output, errors.New("please enter a height")
	case inp.Centimetres > 0 && inp.Metres != 0 || inp.Metres > 0 && inp.Centimetres != 0:
		return output, errors.New("please only enter one unit")
	}

	switch {
	case inp.Centimetres > 0:
		totalInches = float64(inp.Centimetres) / 2.54
	case inp.Metres > 0:
		totalInches = float64(inp.Metres) / 0.0254
	}

	feet := math.Floor(totalInches / 12)
	remainingInches := math.Mod(totalInches, 12)

	output.Feet = float32(feet)
	output.Inches = float32(remainingInches)
	return output, nil
}
