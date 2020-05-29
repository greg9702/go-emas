package test

import (
	"go-emas/pkg/comparator"
	"testing"
)

func TestLinearAgentComparator(t *testing.T) {
	sut := comparator.NewLinearAgentComparator()

	t.Run("Test base cases", func(t *testing.T) {
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

}
