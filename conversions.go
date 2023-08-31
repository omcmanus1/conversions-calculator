package main

import (
	"github.com/omcmanus1/converter/types"
)

func MlConversion(inp types.Setup) float64 {
	floz := float64(inp.Amount) / 28.413
	return floz
}

func GramConversion(inp types.Setup) float64 {
	oz := float64(inp.Amount) / 28.35
	return oz
}
