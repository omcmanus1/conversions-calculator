package types

type Input struct {
	Ingredient   string
	InputSystem  string
	InputUnit    string
	OutputSystem string
	OutputUnit   string
	Type         string
	Amount       float32
}

type Output struct {
	Ingredient   string  `json:"ingredient"`
	OutputSystem string  `json:"system"`
	OutputUnit   string  `json:"unit"`
	Type         string  `json:"type"`
	Amount       float64 `json:"amount"`
}
