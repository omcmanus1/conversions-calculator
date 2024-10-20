package converter

import (
	"testing"
)

func TestFromBodyWeightMetric(t *testing.T) {
	input := BodyWeightUnits{
		Kilograms: 70,
	}
	t.Run("valid input returns correct output", func(t *testing.T) {
		want := BodyWeightUnits{
			Kilograms:  70,
			TotalLbs:   154.32,
			TotalStone: 11.02,
			Stone:      11,
			Lbs:        0.32,
		}
		got, _ := FromBodyWeightMetric(input)
		CheckGotWant(t, got, want)
	})
	t.Run("invalid input returns error", func(t *testing.T) {
		input = BodyWeightUnits{
			Kilograms: -10,
		}
		_, err := FromBodyWeightMetric(input)
		CheckError(t, err, "please input a positive value for kilograms")
	})
}

func TestFromBodyWeightStone(t *testing.T) {
	t.Run("valid conversion from stone and lbs to kilograms", func(t *testing.T) {
		input := BodyWeightUnits{
			Stone: 11,
			Lbs:   3.5,
		}
		want := BodyWeightUnits{
			Stone:      11,
			Lbs:        3.5,
			TotalLbs:   157.5,
			TotalStone: 11.25,
			Kilograms:  71.44,
		}
		got, _ := FromBodyWeightStone(input)
		CheckGotWant(t, got, want)
	})

	t.Run("invalid lbs input (lbs > 14)", func(t *testing.T) {
		input := BodyWeightUnits{
			Stone: 11,
			Lbs:   15,
		}
		_, err := FromBodyWeightStone(input)
		CheckError(t, err, "please input a valid 'lbs' value")
	})
}

func TestFromBodyWeightLbs(t *testing.T) {
	t.Run("valid conversion from lbs to kilograms", func(t *testing.T) {
		input := BodyWeightUnits{
			TotalLbs: 154.32,
		}
		want := BodyWeightUnits{
			TotalLbs:   154.32,
			TotalStone: 11.02,
			Stone:      11,
			Lbs:        0.32,
			Kilograms:  70.0,
		}
		got, _ := FromBodyWeightLbs(input)
		CheckGotWant(t, got, want)
	})

	t.Run("invalid input (negative lbs)", func(t *testing.T) {
		input := BodyWeightUnits{
			TotalLbs: -100,
		}
		_, err := FromBodyWeightLbs(input)
		CheckError(t, err, "please input a positive value for total lbs")
	})
}
