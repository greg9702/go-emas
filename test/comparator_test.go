package test

import (
	"go-emas/pkg/comparator"
	"testing"
)

func TestBasicAgentComparator(t *testing.T) {

	t.Run("Test agent comparator with linear fitness calculator", func(t *testing.T) {
		sut := comparator.NewBasicAgentComparator()

		testParams := []struct {
			agentsFitness int
			rivalsFitness int
			result        bool
		}{
			{1, 2, false},
			{4, 3, true},
		}

		for _, param := range testParams {
			agent := new(MockAgent)
			rival := new(MockAgent)
			agent.On("Fitness").Return(param.agentsFitness)
			rival.On("Fitness").Return(param.rivalsFitness)

			result := sut.Compare(agent, rival)
			if result != param.result {
				t.Errorf("Error in agent comparison, for solutions: %d and %d.", param.agentsFitness, param.rivalsFitness)
			}
		}
	})

}
