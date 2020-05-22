package test

import (
	"errors"
	"go-emas/pkg/common_types"
	"go-emas/pkg/environment"
	"go-emas/pkg/i_agent"
	"testing"
)

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

func (m *mockAgent) SetId(id int64) {
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

func TestEnvironmentInit(t *testing.T) {

	populationFactory := &mockPopulationFactory{}

	t.Run("0 or negative populationSize", func(t *testing.T) {

		populationGeneratorMock = func(populationSize int) (map[int64]i_agent.IAgent, error) {
			return nil, errors.New("0 or negative populationSize")
		}

		testParams := []int{-5, 0}

		for _, param := range testParams {
			_, err := environment.NewEnvironment(param, populationFactory)
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
			obj, err := environment.NewEnvironment(param, populationFactory)

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

		env, err := environment.NewEnvironment(populationSize, populationFactory)

		var testID int64 = 3
		err = env.DeleteFromPopulation(testID)

		if err != nil {
			t.Errorf("Got unexpected err: %s", err)
		}

	})

	t.Run("Normal DeleteFromPopulation", func(t *testing.T) {

		env, err := environment.NewEnvironment(populationSize, populationFactory)

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

	t.Run("Add element to list", func(t *testing.T) {

		env, err := environment.NewEnvironment(populationSize, populationFactory)
		newAgent := &mockAgent{10}

		if newAgent.ID() != 10 {
			t.Errorf("Error in agent creating, expected id: %d got %d.", 10, newAgent.ID())
		}

		err = env.AddToPopulation(newAgent)

		if err != nil {
			t.Errorf("Got unexpected err: %s", err)
		}
	})
}
