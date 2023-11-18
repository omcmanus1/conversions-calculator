package data

import "github.com/omcmanus1/converter/types"

var Input = []types.RecipeInput{
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

var SingleInput = []types.RecipeInput{
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
