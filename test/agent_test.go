package test

import (
	"go-emas/pkg/agent"
	"go-emas/pkg/common_types"
	"go-emas/pkg/i_agent"
	"testing"
)

const ID int64 = 0
const solution common_types.Solution = 10
const actionTag string = common_types.Fight
const energy int = 50

type MockAgentComparator struct {
	result bool
}

func (m MockAgentComparator) Compare(firstAgent i_agent.IAgent, secondAgent i_agent.IAgent) bool {
	return m.result
}

type MockTagCalculator struct {
	result string
}

type MockRandomizer struct {
	result int
}

func (mr MockRandomizer) RandInt(min int, max int) (int, error) {
	return mr.result, nil
}

func (m MockTagCalculator) Calculate(energy int) string {
	return m.result
}

func mockGetAgentByTagEmpty(tag string) i_agent.IAgent {
	return nil
}

func mockGetAgentByTag(tag string) i_agent.IAgent {
	rivalID := ID + 1
	rivalSolution := solution + 10
	rival := agent.NewAgent(rivalID, rivalSolution, actionTag, energy, MockTagCalculator{common_types.Fight},
		MockAgentComparator{true}, MockRandomizer{2}, mockGetAgentByTagEmpty, mockDeleteAgent, mockAddAgent)
	return rival
}

// todo replace with mock.Called
var agentDeleted bool = false

func mockDeleteAgent(id int64) error {
	agentDeleted = true
	return nil
}

// todo replace with mock.Called
var agentAdded bool = false

func mockAddAgent(newAgent i_agent.IAgent) error {
	agentAdded = true
	return nil
}

func expectFight(t *testing.T, sut i_agent.IAgent, expectedEnergyAfterFight int) {
	energyAfterFight := sut.Energy()

	if energyAfterFight != expectedEnergyAfterFight {
		t.Errorf("Error in agents fight, expected energy after fight: %d, got: %d.", expectedEnergyAfterFight, energyAfterFight)
	}
}

func TestAgent(t *testing.T) {
	sut := agent.NewAgent(ID, solution, actionTag, energy, MockTagCalculator{common_types.Fight}, MockAgentComparator{false},
		MockRandomizer{2}, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)

	t.Run("Test modifying energy", func(t *testing.T) {
		testParams := []struct {
			energyDelta int
			finalEnergy int
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

	t.Run("Test fight", func(t *testing.T) {
		sut.Execute()
		expectFight(t, sut, 30)

		sut := agent.NewAgent(ID, solution, actionTag, energy, MockTagCalculator{common_types.Fight}, MockAgentComparator{true},
			MockRandomizer{2}, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)
		sut.Execute()
		expectFight(t, sut, 70)
	})

}

func expectAgentDeath(t *testing.T, agent i_agent.IAgent) {
	if agentDeleted == false {
		t.Errorf("Error - agent with id: %d has not been deleted", agent.ID())
	}
	agentDeleted = false
}

func TestAgentGoingToDie(t *testing.T) {
	sut := agent.NewAgent(ID, solution, common_types.Death, energy, MockTagCalculator{common_types.Death}, MockAgentComparator{false},
		MockRandomizer{2}, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)
	t.Run("Test death", func(t *testing.T) {
		sut.Execute()
		expectAgentDeath(t, sut)
	})
}

func expectAgentMutation(t *testing.T, agent i_agent.IAgent, finalEnergy int) {
	if agentAdded == false {
		t.Errorf("Error - agent with id: %d has not been added", agent.ID())
	}
	agentAdded = false

	if agent.Energy() != finalEnergy {
		t.Errorf("Error in mutation, expected energy: %d, got: %d.", finalEnergy, agent.Energy())
	}
}

func TestAgentGoingToReproduce(t *testing.T) {
	var energy int = 80
	sut := agent.NewAgent(ID, solution, common_types.Reproduction, energy, MockTagCalculator{common_types.Reproduction}, MockAgentComparator{false},
		MockRandomizer{2}, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)
	t.Run("Test mutation", func(t *testing.T) {
		sut.Execute()
		var expectedEnergyAfterMutation int = 40
		expectAgentMutation(t, sut, expectedEnergyAfterMutation)
	})
}
