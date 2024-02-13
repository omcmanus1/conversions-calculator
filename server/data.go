package converter

var Input = []RecipeInput{
	{
		Ingredient:   "haribos",
		InputSystem:  "usa",
		InputUnit:    "cups",
		OutputSystem: "metric",
		OutputUnit:   "grams",
		Type:         "weight",
		Amount:       2,
	},
	{
		Ingredient:   "basmati rice",
		InputSystem:  "usa",
		InputUnit:    "lbs",
		OutputSystem: "metric",
		OutputUnit:   "grams",
		Type:         "weight",
		Amount:       1.5,
	},
	{
		Ingredient:   "chicken thighs",
		InputSystem:  "usa",
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
		OutputSystem: "usa",
		OutputUnit:   "fluid oz",
		Type:         "volume",
		Amount:       12,
	},
}

var SingleInput = []RecipeInput{
	{
		Ingredient:   "tomatoes",
		InputSystem:  "usa",
		InputUnit:    "cups",
		OutputSystem: "metric",
		OutputUnit:   "grams",
		Type:         "weight",
		Amount:       1.5,
	},
}
