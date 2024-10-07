package converter

import (
	"errors"
)

type BodyWeightMetric struct {
	Kilograms float32 `json:"kilograms"`
}

type BodyWeightImperial struct {
	TotalLbs   float32 `json:"total_lbs"`
	TotalStone float32 `json:"total_stone"`
	Stone      int     `json:"stone"`
	Lbs        float32 `json:"lbs"`
}

func FromBodyWeightMetric(inp BodyWeightMetric) (BodyWeightImperial, error) {
	var output BodyWeightImperial

	if inp.Kilograms <= 0 {
		return output, errors.New("please input a valid weight")
	}

	totalLbs := float32(inp.Kilograms) * 2.20462
	stone := (totalLbs) / 14
	remainingLbs := totalLbs - (float32(int(stone)) * 14)

	output.TotalLbs = totalLbs
	output.TotalStone = stone
	output.Stone = int(stone)
	output.Lbs = remainingLbs

	return output, nil
}
