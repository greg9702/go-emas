package tag_calculator

import "go-emas/pkg/common_types"

type ActionTag string

const REPRODUCTION_THRESHOLD = 80

type TagCalculator struct {
}

func (tc TagCalculator) Calculate(energy int) common_types.ActionTag {
	if energy == 0 {
		return common_types.Death
	} else if energy >= REPRODUCTION_THRESHOLD {
		return common_types.Reproduction
	}
	return common_types.Fight
}
