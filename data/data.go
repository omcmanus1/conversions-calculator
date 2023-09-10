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
		Ingredient:   "peas",
		InputSystem:  "metric",
		InputUnit:    "kg",
		OutputSystem: "US",
		OutputUnit:   "lbs",
		Type:         "weight",
		Amount:       12,
	},
}
