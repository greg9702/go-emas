package test

import (
	"go-emas/pkg/agent"
	"go-emas/pkg/common_types"
	"testing"
)

const ID common_types.AgentId = 0
const SOLUTION common_types.Solution = 10
const ACTION_TAG common_types.ActionTag = common_types.Fight
const ENERGY common_types.Energy = 50

type MockTagCalculator struct {
	result common_types.ActionTag
}

func (m MockTagCalculator) Calculate(common_types.Energy) common_types.ActionTag {
	return m.result
}

func TestAgent(t *testing.T) {
	sut := agent.NewAgent(ID, SOLUTION, ACTION_TAG, ENERGY, MockTagCalculator{common_types.Fight})

	t.Run("Test modifying energy", func(t *testing.T) {
		testParams := []struct {
			energyDelta common_types.Energy
			finalEnergy common_types.Energy
		}{
			{0, 50},
			{20, 70},
			{-20, 50},
			{-100, 50},
		}

		for _, param := range testParams {
			sut.ModifyEnergy(param.energyDelta)
			energy := sut.Energy()
			if energy != param.finalEnergy {
				t.Errorf("Error in energy modification, expected: %d, got: %d.", param.finalEnergy, energy)
			}
		}
	})
}
