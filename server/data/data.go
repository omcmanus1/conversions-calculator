package data

import "github.com/omcmanus1/converter/types"

var Input = []types.Input{
	{
		Ingredient:   "haribos",
		InputSystem:  "US",
		InputUnit:    "cups",
		OutputSystem: "metric",
		OutputUnit:   "grams",
		Type:         "weight",
		Amount:       2,
	},
	{
		Ingredient:   "basmati rice",
		InputSystem:  "US",
		InputUnit:    "lbs",
		OutputSystem: "metric",
		OutputUnit:   "grams",
		Type:         "weight",
		Amount:       1.5,
	},
	{
		Ingredient:   "chicken thighs",
		InputSystem:  "US",
		InputUnit:    "oz",
		OutputSystem: "metric",
		OutputUnit:   "grams",
		Type:         "weight",
		Amount:       12,
	},
	{
		Ingredient:   "golden syrup",
		InputSystem:  "metric",
		InputUnit:    "millilitres",
		OutputSystem: "US",
		OutputUnit:   "fluid oz",
		Type:         "volume",
		Amount:       12,
	},
}

var SingleInput = []types.Input{
	{
		Ingredient:   "tomatoes",
		InputSystem:  "US",
		InputUnit:    "cups",
		OutputSystem: "metric",
		OutputUnit:   "grams",
		Type:         "weight",
		Amount:       1.5,
	},
}

var EmptyInput = types.Input{
	Ingredient:   "",
	InputSystem:  "",
	InputUnit:    "",
	OutputSystem: "",
	OutputUnit:   "",
	Type:         "",
	Amount:       0.0,
}

var VMBadInput = types.Input{
	Ingredient:   "haribos",
	InputSystem:  "dollars",
	InputUnit:    "litres",
	OutputSystem: "US",
	OutputUnit:   "fluid oz",
	Type:         "volume",
	Amount:       24,
}
