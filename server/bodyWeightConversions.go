package converter

import (
	"errors"
)

type BodyWeightUnits struct {
	TotalLbs   float32 `json:"totalLbs"`
	TotalStone float32 `json:"totalStone"`
	Stone      int     `json:"stone"`
	Lbs        float32 `json:"lbs"`
	Kilograms  float32 `json:"kilograms"`
}

func FromBodyWeightMetric(inp BodyWeightUnits) (BodyWeightUnits, error) {
	if inp.Kilograms <= 0 {
		return inp, errors.New("please input a positive value for kilograms")
	}

	totalLbs := float32(inp.Kilograms) * 2.20462
	stone := (totalLbs) / 14
	remainingLbs := totalLbs - (float32(int(stone)) * 14)

	inp.TotalLbs = totalLbs
	inp.TotalStone = stone
	inp.Stone = int(stone)
	inp.Lbs = remainingLbs

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
		inp.TotalLbs = float32(inp.Stone)*14 + inp.Lbs
		inp.TotalStone = inp.TotalLbs * 0.0714286
		inp.Kilograms = inp.TotalLbs / 2.20462

		return inp, nil
	}

	if inp.TotalStone > 0 {
		inp.Stone = int(inp.TotalStone)
		inp.Lbs = (inp.TotalStone - float32(inp.Stone)) * 14
		inp.TotalLbs = inp.TotalStone * 14
		inp.Kilograms = inp.TotalStone * 6.35029

		return inp, nil
	}

	return inp, nil
}

func FromBodyWeightLbs(inp BodyWeightUnits) (BodyWeightUnits, error) {
	if inp.TotalLbs <= 0 {
		return inp, errors.New("please input a positive value for total lbs")
	}

	inp.TotalStone = inp.TotalLbs * 0.0714286
	inp.Stone = int(inp.TotalStone)
	inp.Lbs = (inp.TotalStone - float32(inp.Stone)) * 14
	inp.Kilograms = inp.TotalStone * 6.35029

	return inp, nil
}
