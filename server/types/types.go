package types

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

type HeightFromMetric interface {
	GetHeightInFeet() (HeightFeet, error)
}

type HeightFeet struct {
	Feet   float32 `json:"feet"`
	Inches float32 `json:"inches"`
}

type HeightCentimetres struct {
	Centimetres float32 `json:"centimetres"`
}

type HeightMetres struct {
	Metres float32 `json:"metres"`
}

func (h HeightCentimetres) GetHeightInFeet() (HeightFeet, error) {
	var output HeightFeet

	if h.Centimetres < 0 {
		return output, errors.New("invalid input")
	}
	totalInches := float64(h.Centimetres) / 2.54
	feet := math.Floor(totalInches / 12)
	remainingInches := math.Mod(totalInches, 12)
	output.Feet = float32(feet)
	output.Inches = float32(remainingInches)

	return output, nil
}

func (h HeightMetres) GetHeightInFeet() (HeightFeet, error) {
	var output HeightFeet

	if h.Metres < 0 {
		return output, errors.New("invalid input")
	}
	totalInches := float64(h.Metres / 0.0254)
	feet := math.Floor(totalInches / 12)
	remainingInches := math.Mod(totalInches, 12)
	output.Feet = float32(feet)
	output.Inches = float32(remainingInches)

	return output, nil
}
