package types

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

type HeightFeet struct {
	Feet   float32 `json:"feet"`
	Inches float32 `json:"inches"`
}

type HeightMetric struct {
	Centimetres float32 `json:"centimetres"`
}
