package main

import (
	"testing"

	"github.com/omcmanus1/converter/types"
)

func TestVolumeMetricEmpty(t *testing.T) {
	inp := types.Input{Ingredient: "", InputSystem: "", InputUnit: "", OutputSystem: "", OutputUnit: "", Type: "", Amount: 0.0}
	msg, err := VolumeMetric(inp)
	if msg != 0.0 || err == nil {
		t.Fatalf(`VolumeMetric() = %v, %v, want 0.0, error`, msg, err)
	}
}
