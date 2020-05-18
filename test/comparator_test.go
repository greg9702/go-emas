package test

import (
	"go-emas/pkg/agent"
	"go-emas/pkg/comparator"
	"testing"
)

func TestLinearAgentComparator(t *testing.T) {
	sut := comparator.NewLinearAgentComparator()

	t.Run("Test base cases", func(t *testing.T) {
		testParams := []struct {
			firstAgent  agent.Agent
			secondAgent agent.Agent
			result      bool
		}{
			{*agent.NewAgent(0, 1), *agent.NewAgent(1, 2), false},
			{*agent.NewAgent(0, 2), *agent.NewAgent(1, 1), true},
		}

		for _, param := range testParams {
			result := sut.Compare(param.firstAgent, param.secondAgent)
			if result != param.result {
				t.Errorf("Error in agent comparison, for agents: %s and %s.", param.firstAgent, param.secondAgent)
			}
		}
	})

}
