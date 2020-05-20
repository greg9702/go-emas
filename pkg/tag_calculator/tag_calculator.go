package tag_calculator

import "go-emas/pkg/common_types"

const reproductionTreshold = 80

// ITagCalulator is an interface for tag calculators
type ITagCalulator interface {
	Calculate(energy int) string
}

// TagCalculator is a tag calculator
type TagCalculator struct {
}

var instantiated *TagCalculator = nil

// NewTagCalculator creates new TagCalculator object
func NewTagCalculator() *TagCalculator {
	if instantiated == nil {
		instantiated = new(TagCalculator)
	}
	return instantiated
}

// Calculate is used to calculate action for passed energy value
func (tc TagCalculator) Calculate(energy int) string {
	if energy == 0 {
		return common_types.Death
	} else if energy >= reproductionTreshold {
		return common_types.Reproduction
	}
	return common_types.Fight
}
