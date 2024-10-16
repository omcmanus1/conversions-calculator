package converter

import (
	"errors"
)

type BodyWeightUnits struct {
	TotalLbs   float64 `json:"totalLbs"`
	TotalStone float64 `json:"totalStone"`
	Stone      int     `json:"stone"`
	Lbs        float64 `json:"lbs"`
	Kilograms  float64 `json:"kilograms"`
}

func FromBodyWeightMetric(inp BodyWeightUnits) (BodyWeightUnits, error) {
	if inp.Kilograms <= 0 {
		return inp, errors.New("please input a positive value for kilograms")
	}

	totalLbs := float64(inp.Kilograms) * 2.20462
	stone := (totalLbs) / 14
	remainingLbs := totalLbs - (float64(int(stone)) * 14)

	inp.TotalLbs = RoundToCustom((totalLbs), 2)
	inp.TotalStone = RoundToCustom(stone, 2)
	inp.Stone = int(stone)
	inp.Lbs = RoundToCustom(remainingLbs, 2)

	return inp, nil
}

func FromBodyWeightStone(inp BodyWeightUnits) (BodyWeightUnits, error) {
	if inp.TotalStone <= 0 && inp.Stone <= 0 && inp.Lbs <= 0 {
		return inp, errors.New("please input a positive value for total_stone or stone & lbs")
	}

	if inp.Stone > 0 {
		if inp.Lbs > 14 {
			return inp, errors.New("please input a valid 'lbs' value")
		}
		inp.TotalLbs = RoundToCustom((float64(inp.Stone)*14 + inp.Lbs), 2)
		inp.TotalStone = RoundToCustom((inp.TotalLbs * 0.0714286), 2)
		inp.Kilograms = RoundToCustom((inp.TotalLbs / 2.20462), 2)

		return inp, nil
	}

	if inp.TotalStone > 0 {
		inp.Stone = int(inp.TotalStone)
		inp.Lbs = RoundToCustom(((inp.TotalStone - float64(inp.Stone)) * 14), 2)
		inp.TotalLbs = RoundToCustom((inp.TotalStone * 14), 2)
		inp.Kilograms = RoundToCustom((inp.TotalStone * 6.35029), 2)

		return inp, nil
	}

	return inp, nil
}

func FromBodyWeightLbs(inp BodyWeightUnits) (BodyWeightUnits, error) {
	if inp.TotalLbs <= 0 {
		return inp, errors.New("please input a positive value for total lbs")
	}

	inp.TotalStone = RoundToCustom((inp.TotalLbs * 0.0714286), 2)
	inp.Stone = int(inp.TotalStone)
	inp.Lbs = RoundToCustom(((inp.TotalStone - float64(inp.Stone)) * 14), 2)
	inp.Kilograms = RoundToCustom((inp.TotalStone * 6.35029), 2)

	return inp, nil
}
