package test

import (
	"errors"
	"go-emas/pkg/common_types"
	"go-emas/pkg/environment"
	"go-emas/pkg/i_agent"
	"testing"

	"github.com/stretchr/testify/mock"
)

const defaultRandom = 2

// This helps in assigning mock at the runtime instead of compile time
var populationGeneratorMock func(populationSize int) (map[int64]i_agent.IAgent, error)

type mockPopulationFactory struct{}

func (m *mockPopulationFactory) CreatePopulation(populationSize int,
	getAgentByTagCallback func(tag string) (i_agent.IAgent, error),
	deleteAgentCallback func(id int64) error,
	addAgentCallback func(newAgent i_agent.IAgent) error) (map[int64]i_agent.IAgent, error) {
	return populationGeneratorMock(populationSize)
}

type mockAgent struct {
	id int64
}

func (m *mockAgent) ID() int64 {
	return m.id
}

func (m *mockAgent) SetID(id int64) {
	m.id = id
}

func (m *mockAgent) Solution() common_types.Solution {
	var result common_types.Solution
	return result
}

func (m *mockAgent) ActionTag() string {
	var result string
	return result
}

func (m *mockAgent) Energy() int {
	var result int
	return result
}

func (m *mockAgent) ModifyEnergy(energyDelta int) {
}

func (m *mockAgent) Execute() {
}

func (m *mockAgent) String() string {
	return ""
}

func (m *mockAgent) Tag() {
}

type MockStopper struct {
	mock.Mock
}

func (m *MockStopper) Stop(iteration int) bool {
	args := m.Called()
	return args.Bool(0)
}

func TestEnvironmentInit(t *testing.T) {

	populationFactory := &mockPopulationFactory{}

	t.Run("0 or negative populationSize", func(t *testing.T) {

		populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {
			return nil, errors.New("0 or negative populationSize")
		}

		testParams := []int{-5, 0}

		for _, param := range testParams {
			_, err := environment.NewEnvironment(param, populationFactory, &MockStopper{}, MockRandomizer{defaultRandom})
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
				population[int64(i+1)] = &mockAgent{int64(i + 1)}
			}
			return population, nil
		}

		for _, param := range testParams {
			obj, err := environment.NewEnvironment(param, populationFactory, &MockStopper{}, MockRandomizer{defaultRandom})

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

	populationSize := 5

	populationFactory := &mockPopulationFactory{}
	populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {

		population := make(map[int64]i_agent.IAgent)
		for i := 0; i < populationSize; i++ {
			population[int64(i+1)] = &mockAgent{int64(i + 1)}
		}
		return population, nil
	}

	t.Run("Normal DeleteFromPopulation", func(t *testing.T) {

		env, err := environment.NewEnvironment(populationSize, populationFactory, &MockStopper{}, MockRandomizer{defaultRandom})

		var testID int64 = 3
		err = env.DeleteFromPopulation(testID)

		if err != nil {
			t.Errorf("Got unexpected err: %s", err)
		}

	})

	t.Run("Error expected when trying to DeleteFromPopulation agent with id that does not exist", func(t *testing.T) {

		env, err := environment.NewEnvironment(populationSize, populationFactory, &MockStopper{}, MockRandomizer{defaultRandom})

		var testID int64 = 10
		err = env.DeleteFromPopulation(testID)

		if err == nil {
			t.Errorf("Expected error but not received")
		}

	})
}

func TestAddToPopulation(t *testing.T) {

	populationSize := 5

	populationFactory := &mockPopulationFactory{}
	populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {

		population := make(map[int64]i_agent.IAgent)
		for i := 0; i < populationSize; i++ {
			population[int64(i+1)] = &mockAgent{int64(i + 1)}
		}
		return population, nil
	}

	t.Run("Add agent to population", func(t *testing.T) {

		env, err := environment.NewEnvironment(populationSize, populationFactory, &MockStopper{}, MockRandomizer{defaultRandom})
		newAgent := &mockAgent{10}

		err = env.AddToPopulation(newAgent)

		// TODO consider changing way of assigning unique ids
		uniqueId := int64(populationSize + 1)
		if newAgent.ID() != uniqueId {
			t.Errorf("Error in agent creating, expected id: %d got %d.", uniqueId, newAgent.ID())
		}

		if err != nil {
			t.Errorf("Got unexpected err: %s", err)
		}
	})
}

type MockAgentWithTag struct {
	mockAgent
	actionTag string
}

func (m *MockAgentWithTag) ActionTag() string {
	return m.actionTag
}

func TestGetAgentByTag(t *testing.T) {

	populationSize := 5

	populationFactory := &mockPopulationFactory{}
	populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {

		population := make(map[int64]i_agent.IAgent)
		for i := 0; i < populationSize; i++ {
			population[int64(i+1)] = &MockAgentWithTag{mockAgent{int64(i + 1)}, common_types.Death}
		}
		return population, nil
	}

	sut, _ := environment.NewEnvironment(populationSize, populationFactory, &MockStopper{}, MockRandomizer{defaultRandom})
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
	populationSize := 5

	populationFactory := &mockPopulationFactory{}
	populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {

		population := make(map[int64]i_agent.IAgent)
		for i := 0; i < populationSize; i++ {
			population[int64(i+1)] = &MockAgentWithTag{mockAgent{int64(i + 1)}, common_types.Death}
		}
		return population, nil
	}

	mockStopper := new(MockStopper)
	mockStopper.On("Stop", mock.Anything).Return(true)

	sut, _ := environment.NewEnvironment(populationSize, populationFactory, mockStopper, MockRandomizer{defaultRandom})

	t.Run("Return error when there were agents with specified tag, but all of them have done action", func(t *testing.T) {

		sut.Start()

	})
	mockStopper.AssertExpectations(t)

}
