package test

import (
	"go-emas/pkg/common_types"
	"go-emas/pkg/tag_calculator"
	"testing"
)

func TestTagCalculation(t *testing.T) {
	sut := tag_calculator.NewTagCalculator()

	t.Run("Test base cases", func(t *testing.T) {
		testParams := []struct {
			energy common_types.Energy
			tag    common_types.ActionTag
		}{
			{0, common_types.Death},
			{20, common_types.Fight},
			{80, common_types.Reproduction},
			{81, common_types.Reproduction},
		}

		for _, param := range testParams {
			tag := sut.Calculate(param.energy)
			if tag != param.tag {
				t.Errorf("Error in tag calculations, for energy: %d got tag: %s, want: %s.", param.energy, tag, param.tag)
			}
		}
	})
}
