package test

import (
	"go-emas/pkg/tag_calculator"
	"testing"
)

func TestTagCalculationDummy(t *testing.T) {
	sut := tag_calculator.TagCalculator{}

	t.Run("Test base cases", func(t *testing.T) {
		expected := []struct {
			energy int
			tag    tag_calculator.AgentTag
		}{
			{0, tag_calculator.Death},
			{20, tag_calculator.Fight},
			{80, tag_calculator.Fight},
			{81, tag_calculator.Reproduction},
		}

		for _, exp := range expected {
			tag := sut.DummyCalculate(exp.energy)
			if tag != exp.tag {
				t.Errorf("Error in tag calculations, for energy: %d got tag: %s, want: %s.", exp.energy, tag, exp.tag)
			}
		}
	})
}
