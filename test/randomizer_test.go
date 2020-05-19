package test

import (
	"go-emas/pkg/randomizer"
	"testing"
)

func TestBaseRandomizer(t *testing.T) {

	t.Run("Test RadInt expected usage", func(t *testing.T) {
		testParams := []struct {
			min int
			max int
		}{
			{0, 20},
			{21, 30},
			{50, 51},
			{81, 88},
		}

		numberOfIteration := 100

		for _, param := range testParams {
			for i := 0; i < numberOfIteration; i++ {
				result, err := randomizer.BaseRand().RandInt(param.min, param.max)

				if err != nil {
					t.Errorf("Error got unexpected error: %s", err)
				}

				if param.min > result || result > param.max {
					t.Errorf("Error in RandInt, for min: %d and max: %d got: %d.", param.min, param.max, result)
				}
			}
		}
	})

	t.Run("Test RadInt equal min and max values", func(t *testing.T) {

		numberOfIteration := 100

		testParams := []int{20, 20}

		min := testParams[0]
		max := testParams[1]

		for i := 0; i < numberOfIteration; i++ {
			result, err := randomizer.BaseRand().RandInt(min, max)

			if err != nil {
				t.Errorf("Error got unexpected error: %s", err)
			}

			if min > result || result > max {
				t.Errorf("Error in RandInt, for min: %d and max: %d got: %d.", min, max, result)
			}
		}
	})

	t.Run("Test RadInt negative values and min > max", func(t *testing.T) {

		numberOfIteration := 100

		testParams := []struct {
			min int
			max int
		}{
			{0, -5},
			{-1, 0},
			{50, 49},
		}

		for _, param := range testParams {
			for i := 0; i < numberOfIteration; i++ {
				_, err := randomizer.BaseRand().RandInt(param.min, param.max)

				if err == nil {
					t.Errorf("Expecting error but not received")
				}
			}
		}
	})
}
