package test

import (
	"go-emas/pkg/comparator"
	"go-emas/pkg/fitness_calculator"
	"go-emas/pkg/solution"
	"testing"

	"github.com/willf/bitset"
)

type MockAgentWithBitSetSolution struct {
	*MockAgent
	solution solution.ISolution
}

func (m MockAgentWithBitSetSolution) Solution() solution.ISolution {
	return m.solution
}

func TestBasicAgentComparator(t *testing.T) {

	t.Run("Test agent comparator with linear fitness calculator", func(t *testing.T) {
		sut := comparator.NewBasicAgentComparator(fitness_calculator.NewLinearFitnessCalculator())

		testParams := []struct {
			agentsSolution int
			rivalsSolution int
			result         bool
		}{
			{1, 2, false},
			{4, 3, true},
		}

		for _, param := range testParams {
			agent := new(MockAgent)
			rival := new(MockAgent)
			agent.On("Solution").Return(param.agentsSolution)
			rival.On("Solution").Return(param.rivalsSolution)

			result := sut.Compare(agent, rival)
			if result != param.result {
				t.Errorf("Error in agent comparison, for solutions: %d and %d.", param.agentsSolution, param.rivalsSolution)
			}
		}
	})

	t.Run("Test agent comparator with bitset fitness calculator", func(t *testing.T) {
		sut := comparator.NewBasicAgentComparator(fitness_calculator.NewBitSetFitnessCalculator())

		testParams := []struct {
			agentsSolution bitset.BitSet
			rivalsSolution bitset.BitSet
			result         bool
		}{
			{*bitset.From([]uint64{0, 1}), *bitset.From([]uint64{0, 1, 2}), false},
			{*bitset.From([]uint64{5, 6, 7, 8}), *bitset.From([]uint64{1, 2, 3}), true},
		}

		for _, param := range testParams {
			agent := MockAgentWithBitSetSolution{MockAgent: new(MockAgent), solution: solution.NewBitSetSolution(param.agentsSolution)}
			rival := MockAgentWithBitSetSolution{MockAgent: new(MockAgent), solution: solution.NewBitSetSolution(param.rivalsSolution)}

			result := sut.Compare(agent, rival)
			if result != param.result {
				t.Errorf("Error in agent comparison, for solutions: %d and %d.", param.agentsSolution, param.rivalsSolution)
			}
		}
	})
}
