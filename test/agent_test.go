package test

import (
	"go-emas/pkg/agent"
	"go-emas/pkg/common_types"
	"go-emas/pkg/i_agent"
	"testing"

	"github.com/stretchr/testify/mock"
)

const ID int64 = 0
const solution common_types.Solution = 10
const actionTag string = common_types.Fight
const energy int = 50

type MockTagCalculator struct {
	result string
}

func (m MockTagCalculator) Calculate(energy int) string {
	return m.result
}

func mockGetAgentByTagEmpty(tag string) (i_agent.IAgent, error) {
	return nil, nil
}

func mockGetAgentByTag(tag string) (i_agent.IAgent, error) {
	rivalID := ID + 1
	rivalSolution := solution + 10
	rival := agent.NewAgent(rivalID, rivalSolution, actionTag, energy, MockTagCalculator{common_types.Fight},
		&MockAgentComparator{}, &MockRandomizer{}, mockGetAgentByTagEmpty, mockDeleteAgent, mockAddAgent)
	return rival, nil
}

var agentDeleted bool = false

func mockDeleteAgent(id int64) error {
	agentDeleted = true
	return nil
}

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
	sut := agent.NewAgent(ID, solution, actionTag, energy, MockTagCalculator{common_types.Fight}, &MockAgentComparator{},
		&MockRandomizer{}, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)

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
		agentComparator := new(MockAgentComparator)
		// TODO improve expectation - remove Once and specify Compare() arguments
		agentComparator.On("Compare").Return(false).Once()
		sut := agent.NewAgent(ID, solution, actionTag, energy, MockTagCalculator{common_types.Fight}, agentComparator,
			&MockRandomizer{}, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)

		sut.Execute()
		expectFight(t, sut, 30)

		agentComparator.On("Compare").Return(true).Once()
		sut.Execute()
		expectFight(t, sut, 50)
	})

}

func expectAgentDeath(t *testing.T, agent i_agent.IAgent) {
	if agentDeleted == false {
		t.Errorf("Error - agent with id: %d has not been deleted", agent.ID())
	}
	agentDeleted = false
}

func TestAgentGoingToDie(t *testing.T) {
	sut := agent.NewAgent(ID, solution, common_types.Death, energy, MockTagCalculator{common_types.Death}, &MockAgentComparator{},
		&MockRandomizer{}, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)
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
	randomizer := new(MockRandomizer)
	randomizer.On("RandInt", mock.Anything, mock.Anything).Return(2)
	sut := agent.NewAgent(ID, solution, common_types.Reproduction, energy, MockTagCalculator{common_types.Reproduction}, &MockAgentComparator{},
		randomizer, mockGetAgentByTag, mockDeleteAgent, mockAddAgent)
	t.Run("Test mutation", func(t *testing.T) {
		sut.Execute()
		var expectedEnergyAfterMutation int = 40
		expectAgentMutation(t, sut, expectedEnergyAfterMutation)
	})
}
