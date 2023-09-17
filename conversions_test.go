package main

import (
	"testing"

	"github.com/omcmanus1/converter/data"
)

func checkGotWant(t testing.TB, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf(`got %v want %v`, got, want)
	}
}

func checkError(t testing.TB, err interface{}) {
	t.Helper()
	if err == nil {
		t.Errorf("got %v want nil", err)
	}
}

func TestVMEmpty(t *testing.T) {
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

func TestVMBadInput(t *testing.T) {
	badInp := data.VMBadInput
	got, err := VolumeMetric(badInp)
	want := 0.0
	if got != want {
		t.Errorf(`got %v want %v`, got, want)
	}
	if err == nil {
		t.Errorf("got %v want nil", err)
	}
}

func TestVolumeMetric(t *testing.T) {
	t.Run("empty inputs result in error", func(t *testing.T) {
		inp := data.EmptyInput
		got, err := VolumeMetric(inp)
		want := 0.0
		checkGotWant(t, got, want)
		checkError(t, err)
	})
	t.Run("invalid input system result in error", func(t *testing.T) {
		inp := data.VMBadInput
		got, err := VolumeMetric(inp)
		want := 0.0
		checkGotWant(t, got, want)
		checkError(t, err)
	})
}
