package test

import (
	"go-emas/pkg/agent"
	"go-emas/pkg/common_types"
	"go-emas/pkg/comparator"
	"testing"
)

type mockAgent struct {
	solution common_types.Solution
}

func (m *mockAgent) Id() common_types.AgentId {
	var result common_types.AgentId
	return result
}

func (m *mockAgent) Solution() common_types.Solution {
	return m.solution
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

func (m *mockAgent) Tag() {
}

func (m *mockAgent) Execute() {
}

func (m *mockAgent) String() string {
	return ""
}
func TestLinearAgentComparator(t *testing.T) {
	sut := comparator.NewLinearAgentComparator()

	t.Run("Test base cases", func(t *testing.T) {
		testParams := []struct {
			firstAgent  agent.IAgent
			secondAgent agent.IAgent
			result      bool
		}{
			{&mockAgent{1}, &mockAgent{2}, false},
			{&mockAgent{4}, &mockAgent{3}, true},
		}

		for _, param := range testParams {
			result := sut.Compare(param.firstAgent, param.secondAgent)
			if result != param.result {
				t.Errorf("Error in agent comparison, for agents: %s and %s.", param.firstAgent, param.secondAgent)
			}
		}
	})

}
