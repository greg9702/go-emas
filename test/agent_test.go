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

type MockRandomizer struct {
	result int
}

func (mr MockRandomizer) RandInt(min int, max int) (int, error) {
	return mr.result, nil
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
	rival := agent.NewAgent(rivalId, rivalSolution, ACTION_TAG, ENERGY, MockTagCalculator{common_types.Fight}, MockAgentComparator{true}, MockRandomizer{2}, mockGetAgentByTagEmpty, mockDeleteAgent, mockAddAgent)
	return rival
}

// todo replace with mock.Called
var agentDeleted bool = false

func mockDeleteAgent(common_types.AgentId) {
	agentDeleted = true
}

// todo replace with mock.Called
var agentAdded bool = false

func mockAddAgent(newAgent i_agent.IAgent) {
	agentAdded = true
}

func expectFight(t *testing.T, sut i_agent.IAgent, expectedEnergyAfterFight common_types.Energy) {
	energyAfterFight := sut.Energy()

	if energyAfterFight != expectedEnergyAfterFight {
		t.Errorf("Error in agents fight, expected energy after fight: %d, got: %d.", expectedEnergyAfterFight, energyAfterFight)
	}
}

func TestAgent(t *testing.T) {
	sut := agent.NewAgent(ID, SOLUTION, ACTION_TAG, ENERGY, MockTagCalculator{common_types.Fight}, MockAgentComparator{false}, MockRandomizer{2}, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)

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

	t.Run("Test fight", func(t *testing.T) {
		sut.Execute()
		expectFight(t, sut, 30)

		sut := agent.NewAgent(ID, SOLUTION, ACTION_TAG, ENERGY, MockTagCalculator{common_types.Fight}, MockAgentComparator{true}, MockRandomizer{2}, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)
		sut.Execute()
		expectFight(t, sut, 70)
	})

}

func expectAgentDeath(t *testing.T, agent i_agent.IAgent) {
	if agentDeleted == false {
		t.Errorf("Error - agent with id: %d has not been deleted", agent.Id())
	}
	agentDeleted = false
}

func TestAgentGoingToDie(t *testing.T) {
	sut := agent.NewAgent(ID, SOLUTION, common_types.Death, ENERGY, MockTagCalculator{common_types.Death}, MockAgentComparator{false}, MockRandomizer{2}, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)
	t.Run("Test death", func(t *testing.T) {
		sut.Execute()
		expectAgentDeath(t, sut)
	})
}

func expectAgentMutation(t *testing.T, agent i_agent.IAgent, finalEnergy common_types.Energy) {
	if agentAdded == false {
		t.Errorf("Error - agent with id: %d has not been added", agent.Id())
	}
	agentAdded = false

	if agent.Energy() != finalEnergy {
		t.Errorf("Error in mutation, expected energy: %d, got: %d.", finalEnergy, agent.Energy())
	}
}

func TestAgentGoingToReproduce(t *testing.T) {
	var energy common_types.Energy = 80
	sut := agent.NewAgent(ID, SOLUTION, common_types.Reproduction, energy, MockTagCalculator{common_types.Reproduction}, MockAgentComparator{false}, MockRandomizer{2}, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)
	t.Run("Test mutation", func(t *testing.T) {
		sut.Execute()
		var expectedEnergyAfterMutation common_types.Energy = 40
		expectAgentMutation(t, sut, expectedEnergyAfterMutation)
	})
}
