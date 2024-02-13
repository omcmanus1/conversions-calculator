package converter

import "testing"

func TestFromFeet(t *testing.T) {
	t.Run("empty input returns empty error", func(t *testing.T) {
		inp := HeightFeet{}
		got, err := FromFeet(inp)
		want := HeightMetric{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "please input an imperial height")
	})
	t.Run("empty input returns empty error", func(t *testing.T) {
		inp := HeightFeet{
			Feet:   0,
			Inches: 0,
		}
		got, err := FromFeet(inp)
		want := HeightMetric{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "please input an imperial height")
	})
	t.Run("invalid feet returns invalid error", func(t *testing.T) {
		inp := HeightFeet{
			Feet:   -23,
			Inches: 2,
		}
		got, err := FromFeet(inp)
		want := HeightMetric{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "invalid input")
	})
	t.Run("invalid inches returns invalid error", func(t *testing.T) {
		inp := HeightFeet{
			Feet:   5,
			Inches: -2,
		}
		got, err := FromFeet(inp)
		want := HeightMetric{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "invalid input")
	})
	t.Run("invalid inches returns invalid error", func(t *testing.T) {
		inp := HeightFeet{
			Feet:   4,
			Inches: 23,
		}
		got, err := FromFeet(inp)
		want := HeightMetric{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "invalid input")
	})
	t.Run("valid input returns correct output", func(t *testing.T) {
		inp := HeightFeet{
			Feet:   5,
			Inches: 2,
		}
		want := HeightMetric{
			Centimetres: 157,
			Metres:      1.57,
		}
		got, _ := FromFeet(inp)
		CheckGotWant(t, got, want)
	})
}

func TestFromMetric(t *testing.T) {
	t.Run("empty input returns empty error", func(t *testing.T) {
		inp := HeightMetric{
			Centimetres: 0,
			Metres:      0,
		}
		got, err := FromMetric(inp)
		want := HeightFeet{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "please input a metric height")
	})
	t.Run("empty input returns empty error", func(t *testing.T) {
		inp := HeightMetric{}
		got, err := FromMetric(inp)
		want := HeightFeet{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "please input a metric height")
	})
	t.Run("conflicting input returns invalid error", func(t *testing.T) {
		inp := HeightMetric{
			Metres:      1.86,
			Centimetres: 160,
		}
		got, err := FromMetric(inp)
		want := HeightFeet{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "please only input one unit")
	})
	t.Run("invalid metres input returns invalid error", func(t *testing.T) {
		inp := HeightMetric{
			Metres:      -2,
			Centimetres: 0,
		}
		got, err := FromMetric(inp)
		want := HeightFeet{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "invalid input")
	})
	t.Run("invalid centimetres input returns invalid error", func(t *testing.T) {
		inp := HeightMetric{
			Metres:      0,
			Centimetres: -186,
		}
		got, err := FromMetric(inp)
		want := HeightFeet{}
		CheckGotWant(t, got, want)
		CheckError(t, err, "invalid input")
	})
	t.Run("valid centimetres input returns correct output", func(t *testing.T) {
		inp := HeightMetric{
			Metres:      0,
			Centimetres: 186,
		}
		want := HeightFeet{
			Feet:   6,
			Inches: 1.2,
		}
		got, _ := FromMetric(inp)
		CheckGotWant(t, got, want)
	})
	t.Run("valid metres input returns correct output", func(t *testing.T) {
		inp := HeightMetric{
			Metres:      1.86,
			Centimetres: 0,
		}
		want := HeightFeet{
			Feet:   6,
			Inches: 1.2,
		}
		got, _ := FromMetric(inp)
		CheckGotWant(t, got, want)
	})
}
