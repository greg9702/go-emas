package test

import (
	"go-emas/pkg/agent"
	"go-emas/pkg/common_types"
	"go-emas/pkg/i_agent"
	"testing"
)

const ID common_types.AgentId = 0
const SOLUTION common_types.Solution = 10
const ACTION_TAG common_types.ActionTag = common_types.Fight
const ENERGY common_types.Energy = 50

type MockAgentComparator struct {
	result bool
}

func (m MockAgentComparator) Compare(firstAgent i_agent.IAgent, secondAgent i_agent.IAgent) bool {
	return m.result
}

type MockTagCalculator struct {
	result common_types.ActionTag
}

func (m MockTagCalculator) Calculate(common_types.Energy) common_types.ActionTag {
	return m.result
}

func mockGetAgentByTagEmpty(tag common_types.ActionTag) i_agent.IAgent {
	return nil
}

func mockGetAgentByTag(tag common_types.ActionTag) i_agent.IAgent {
	rivalId := ID + 1
	rivalSolution := SOLUTION + 10
	rival := agent.NewAgent(rivalId, rivalSolution, ACTION_TAG, ENERGY, MockTagCalculator{common_types.Fight}, MockAgentComparator{true}, mockGetAgentByTagEmpty)
	return rival
}

func expectFight(t *testing.T, sut i_agent.IAgent, expectedEnergyAfterFight common_types.Energy) {
	energyAfterFight := sut.Energy()

	if energyAfterFight != expectedEnergyAfterFight {
		t.Errorf("Error in agents fight, expected energy after fight: %d, got: %d.", expectedEnergyAfterFight, energyAfterFight)
	}
}

func TestAgent(t *testing.T) {
	sut := agent.NewAgent(ID, SOLUTION, ACTION_TAG, ENERGY, MockTagCalculator{common_types.Fight}, MockAgentComparator{false}, mockGetAgentByTag)

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

	t.Run("Test won fight", func(t *testing.T) {
		sut.Execute()
		expectFight(t, sut, 30)

		sut := agent.NewAgent(ID, SOLUTION, ACTION_TAG, ENERGY, MockTagCalculator{common_types.Fight}, MockAgentComparator{true}, mockGetAgentByTag)
		sut.Execute()
		expectFight(t, sut, 70)
	})

}