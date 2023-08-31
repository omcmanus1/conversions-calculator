package main

import "github.com/omcmanus1/converter/types"

func main() {
	data := types.Setup{
		InputSystem: "metric",
		InputUnit:   "millilitres",
		OutputUnit:  "fluid oz",
		Baking:      false,
		Amount:      245,
	}
	Flow(data)
}
