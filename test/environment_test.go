package test

import (
	"errors"
	"go-emas/pkg/agent"
	"go-emas/pkg/common_types"
	"go-emas/pkg/environment"
	"testing"
)

// This helps in assigning mock at the runtime instead of compile time
var populationGeneratorMock func(populationSize int) (map[int]agent.IAgent, error)

type mockPopulationFactory struct{}

func (b *mockPopulationFactory) CreatePopulation(populationSize int) (map[int]agent.IAgent, error) {
	return populationGeneratorMock(populationSize)
}

type mockAgent struct {
	id int
}

func (m *mockAgent) ID() int {
	return m.id
}

func (m *mockAgent) Solution() common_types.Solution {
	var result common_types.Solution
	return result
}

func (m *mockAgent) ActionTag() common_types.ActionTag {
	var result common_types.ActionTag
	return result
}

func (m *mockAgent) Energy() common_types.Energy {
	var result common_types.Energy
	return result
}

func (m *mockAgent) ModifyEnergy(energyDelta common_types.Energy) {
}

func (m *mockAgent) Execute() {
}

func (m *mockAgent) String() string {
	return ""
}

func TestEnvironmentInit(t *testing.T) {

	populationFactory := &mockPopulationFactory{}

	t.Run("0 or negative populationSize", func(t *testing.T) {

		populationGeneratorMock = func(populationSize int) (map[int]agent.IAgent, error) {
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

		populationGeneratorMock = func(populationSize int) (map[int]agent.IAgent, error) {

			population := make(map[int]agent.IAgent)
			for i := 0; i < populationSize; i++ {
				population[i+1] = &mockAgent{i + 1}
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
	populationGeneratorMock = func(populationSize int) (map[int]agent.IAgent, error) {

		population := make(map[int]agent.IAgent)
		for i := 0; i < populationSize; i++ {
			population[i+1] = &mockAgent{i + 1}
		}
		return population, nil
	}

	t.Run("Normal DeleteFromPopulation", func(t *testing.T) {

		env, err := environment.NewEnvironment(populationSize, populationFactory)

		testID := 3
		err = env.DeleteFromPopulation(testID)

		if err != nil {
			t.Errorf("Got unexpected err: %s", err)
		}

	})

	t.Run("Normal DeleteFromPopulation", func(t *testing.T) {

		env, err := environment.NewEnvironment(populationSize, populationFactory)

		testID := 10
		err = env.DeleteFromPopulation(testID)

		if err == nil {
			t.Errorf("Expected error but not received")
		}

	})
}

func TestAddToPopulation(t *testing.T) {

	populationSize := 5

	populationFactory := &mockPopulationFactory{}
	populationGeneratorMock = func(populationSize int) (map[int]agent.IAgent, error) {

		population := make(map[int]agent.IAgent)
		for i := 0; i < populationSize; i++ {
			population[i+1] = &mockAgent{i + 1}
		}
		return population, nil
	}

	t.Run("Add element to list which do not exists", func(t *testing.T) {

		env, err := environment.NewEnvironment(populationSize, populationFactory)

		newAgent := &mockAgent{10}

		err = env.AddToPopulation(newAgent)

		if err != nil {
			t.Errorf("Got unexpected err: %s", err)
		}
	})

	t.Run("Add element to list which already exists", func(t *testing.T) {

		env, err := environment.NewEnvironment(populationSize, populationFactory)

		newAgent := &mockAgent{3}

		err = env.AddToPopulation(newAgent)

		if err == nil {
			t.Errorf("Expected error but not received")
		}
	})
}
