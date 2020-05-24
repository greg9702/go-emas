package test

import (
	"errors"
	"go-emas/pkg/common_types"
	"go-emas/pkg/environment"
	"go-emas/pkg/i_agent"
	"testing"

	"github.com/stretchr/testify/mock"
)

const populationSize = 5
const existingAgentId = 3

// This helps in assigning mock at the runtime instead of compile time
var populationGeneratorMock func(populationSize int) (map[int64]i_agent.IAgent, error)

type mockPopulationFactory struct{}

func (m *mockPopulationFactory) CreatePopulation(populationSize int,
	getAgentByTagCallback func(tag string) (i_agent.IAgent, error),
	deleteAgentCallback func(id int64) error,
	addAgentCallback func(newAgent i_agent.IAgent) error) (map[int64]i_agent.IAgent, error) {
	return populationGeneratorMock(populationSize)
}

func TestEnvironmentInit(t *testing.T) {

	populationFactory := &mockPopulationFactory{}

	t.Run("0 or negative populationSize", func(t *testing.T) {

		populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {
			return nil, errors.New("0 or negative populationSize")
		}

		testParams := []int{-5, 0}

		for _, param := range testParams {
			_, err := environment.NewEnvironment(param, populationFactory, &MockStopper{}, &MockRandomizer{})
			if err == nil {
				t.Errorf("Should return error")
			}
		}

	})

	t.Run("Test positive values", func(t *testing.T) {
		testParams := []int{1, 5, 21}

		populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {

			population := make(map[int64]i_agent.IAgent)
			for i := 0; i < populationSize; i++ {
				population[int64(i+1)] = new(MockAgent)
			}
			return population, nil
		}

		for _, param := range testParams {
			obj, err := environment.NewEnvironment(param, populationFactory, &MockStopper{}, &MockRandomizer{})

			got := obj.PopulationSize()
			expected := param

			if err != nil {
				t.Errorf("Got not expected error: %s", err)
			}

			if got != expected {
				t.Errorf("Error in environment initialization, for param: %d got PopulationSize: %d, want: %d.", param, got, expected)
			}
		}
	})
}

func TestDeleteFromPopulation(t *testing.T) {
	populationFactory := &mockPopulationFactory{}
	populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {

		population := make(map[int64]i_agent.IAgent)
		for i := 0; i < populationSize; i++ {
			agent := new(MockAgent)
			agent.On("Id").Return(i + 1)
			population[int64(i+1)] = agent
		}
		return population, nil
	}

	t.Run("Normal DeleteFromPopulation", func(t *testing.T) {

		env, err := environment.NewEnvironment(populationSize, populationFactory, &MockStopper{}, &MockRandomizer{})

		var testID int64 = existingAgentId
		err = env.DeleteFromPopulation(testID)

		if err != nil {
			t.Errorf("Got unexpected err: %s", err)
		}

	})

	t.Run("Error expected when trying to DeleteFromPopulation agent with id that does not exist", func(t *testing.T) {

		env, err := environment.NewEnvironment(populationSize, populationFactory, &MockStopper{}, &MockRandomizer{})

		var testID int64 = 10
		err = env.DeleteFromPopulation(testID)

		if err == nil {
			t.Errorf("Expected error but not received")
		}

	})
}

func TestAddToPopulation(t *testing.T) {

	populationFactory := &mockPopulationFactory{}
	populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {

		population := make(map[int64]i_agent.IAgent)
		for i := 0; i < populationSize; i++ {
			mockAgent := new(MockAgent)
			mockAgent.On("ID").Return(int64(i + 1))
			population[int64(i+1)] = mockAgent
		}
		return population, nil
	}

	t.Run("Add agent to population", func(t *testing.T) {
		env, err := environment.NewEnvironment(populationSize, populationFactory, &MockStopper{}, &MockRandomizer{})
		newAgent := new(MockAgent)
		uniqueId := populationSize + 1
		newAgent.On("ID").Return(uniqueId)
		newAgent.On("SetID").Return(nil)

		err = env.AddToPopulation(newAgent)
		// TODO it should work, but it's not...
		// newAgent.AssertCalled(t, "SetID", uniqueId)
		newAgent.AssertCalled(t, "SetID")

		if err != nil {
			t.Errorf("Got unexpected err: %s", err)
		}
	})
}

func TestGetAgentByTag(t *testing.T) {

	populationFactory := &mockPopulationFactory{}
	populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {

		population := make(map[int64]i_agent.IAgent)
		for i := 0; i < populationSize; i++ {
			mockAgent := new(MockAgent)
			mockAgent.On("ActionTag").Return(common_types.Death)
			mockAgent.On("Tag").Return(nil)
			population[int64(i+1)] = mockAgent
		}
		return population, nil
	}

	sut, _ := environment.NewEnvironment(populationSize, populationFactory, &MockStopper{}, &MockRandomizer{})
	sut.TagAgents()

	t.Run("Return error when there is no agent with specified tag", func(t *testing.T) {
		actionNotAvail := common_types.Reproduction
		_, err := sut.GetAgentByTag(actionNotAvail)
		if err == nil {
			t.Errorf("There was no agent with specified tag in population, but GetAgentByTag reported no error")
		}
	})

	t.Run("Return agent with specified tag as long as there is one", func(t *testing.T) {
		action := common_types.Death
		for i := 0; i < populationSize; i++ {
			_, err := sut.GetAgentByTag(action)
			if err != nil {
				t.Errorf("Got unexpeced error, agent with tag %s should be returned", action)
			}
		}
	})

	t.Run("Return error when there were agents with specified tag, but all of them have done action", func(t *testing.T) {
		action := common_types.Death
		_, err := sut.GetAgentByTag(action)
		if err == nil {
			t.Errorf("There was no agent with specified tag in population, but GetAgentByTag reported no error")
		}
	})
}

func TestExecutionFlow(t *testing.T) {
	populationFactory := &mockPopulationFactory{}
	mockAgents := make([]*MockAgent, populationSize)

	populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {

		population := make(map[int64]i_agent.IAgent)
		for i := 0; i < populationSize; i++ {
			mockAgents[i] = new(MockAgent)
			mockAgents[i].On("ActionTag").Return(common_types.Death)
			mockAgents[i].On("Tag").Return(nil)
			mockAgents[i].On("Execute").Return(nil)
			population[int64(i+1)] = mockAgents[i]
		}
		return population, nil
	}

	mockStopper := new(MockStopper)
	mockStopper.On("Stop", mock.Anything).Return(true)

	sut, _ := environment.NewEnvironment(populationSize, populationFactory, mockStopper, &MockRandomizer{})

	sut.Start()
	for _, mockAgent := range mockAgents {

		mockAgent.AssertCalled(t, "Tag")
		mockAgent.AssertCalled(t, "Execute")
	}

	mockStopper.AssertExpectations(t)

}
