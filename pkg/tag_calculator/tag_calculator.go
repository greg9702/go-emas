package tag_calculator

import "go-emas/pkg/common_types"

type ActionTag string

const REPRODUCTION_THRESHOLD = 80

type ITagCalulator interface {
	Calculate(common_types.Energy) common_types.ActionTag
}

type TagCalculator struct {
}

var instantiated *TagCalculator = nil

func NewTagCalculator() *TagCalculator {
	if instantiated == nil {
		instantiated = new(TagCalculator)
	}
	return instantiated
}

func (tc TagCalculator) Calculate(energy common_types.Energy) common_types.ActionTag {
	if energy == 0 {
		return common_types.Death
	} else if energy >= REPRODUCTION_THRESHOLD {
		return common_types.Reproduction
	}
	return common_types.Fight
}
