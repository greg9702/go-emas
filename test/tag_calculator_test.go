package test

import (
	"go-emas/pkg/common"
	"go-emas/pkg/tag_calculator"
	"testing"
)

func TestTagCalculation(t *testing.T) {
	sut := tag_calculator.NewTagCalculator()

	t.Run("Test base cases", func(t *testing.T) {
		testParams := []struct {
			energy int
			tag    string
		}{
			{0, common.Death},
			{20, common.Fight},
			{80, common.Reproduction},
			{81, common.Reproduction},
		}

		for _, param := range testParams {
			tag := sut.Calculate(param.energy)
			if tag != param.tag {
				t.Errorf("Error in tag calculations, for energy: %d got tag: %s, want: %s.", param.energy, tag, param.tag)
			}
		}
	})
}
