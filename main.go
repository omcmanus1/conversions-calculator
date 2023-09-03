package main

import "github.com/omcmanus1/converter/types"

func main() {
	data := types.Setup{
		Ingredient:  "haribos",
		InputSystem: "US",
		InputUnit:   "cups",
		OutputUnit:  "grams",
		Type:        "vol",
		Amount:      25,
	}
	Flow(data)
}
