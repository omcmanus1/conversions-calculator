package converter

import (
	"errors"
	"math"
)

type HeightFeet struct {
	Feet   int     `json:"feet"`
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
		return output, errors.New("please input an imperial height")
	case inp.Feet < 0 || inp.Inches < 0:
		return output, errors.New("invalid input")
	}

	feet := 30.48 * float32(inp.Feet)
	inches := 2.54 * inp.Inches

	output.Centimetres = float32(math.Round(float64(feet + inches)))
	output.Metres = float32(math.Round(float64((feet+inches)/100*100))) / 100
	return output, nil
}

func FromMetric(inp HeightMetric) (HeightFeet, error) {
	var output HeightFeet
	var totalInches float64

	switch {
	case inp.Centimetres == 0 && inp.Metres == 0:
		return output, errors.New("please input a metric height")
	case inp.Centimetres > 0 && inp.Metres != 0 || inp.Metres > 0 && inp.Centimetres != 0:
		return output, errors.New("please only input one unit")
	case inp.Centimetres < 0 || inp.Metres < 0:
		return output, errors.New("invalid input")
	}

	switch {
	case inp.Centimetres > 0:
		totalInches = float64(inp.Centimetres) / 2.54
	case inp.Metres > 0:
		totalInches = float64(inp.Metres) / 0.0254
	}

	totalInches = math.Round(totalInches*10) / 10
	feet := math.Floor(totalInches / 12)
	remainingInches := math.Mod(totalInches, 12)

	output.Feet = int(feet)
	output.Inches = float32(remainingInches)
	return output, nil
}
