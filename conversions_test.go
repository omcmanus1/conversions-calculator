package main

import (
	"testing"

	"github.com/omcmanus1/converter/data"
)

func TestVolumeMetricEmpty(t *testing.T) {
	inp := data.EmptyInput
	got, err := VolumeMetric(inp)
	want := 0.0
	if got != want {
		t.Errorf(`got %v want %v`, got, want)
	}
	if err == nil {
		t.Errorf("got %v want nil", err)
	}
}
