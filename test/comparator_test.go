package test

import (
	"go-emas/pkg/common_types"
	"go-emas/pkg/comparator"
	"go-emas/pkg/i_agent"
	"testing"
)

type MockAgent struct {
	solution common_types.Solution
}

func (m *MockAgent) Id() common_types.AgentId {
	var result common_types.AgentId
	return result
}

func (m *MockAgent) Solution() common_types.Solution {
	return m.solution
}

func (m *MockAgent) ActionTag() common_types.ActionTag {
	var result common_types.ActionTag
	return result
}

func (m *MockAgent) Energy() common_types.Energy {
	var result common_types.Energy
	return result
}

func (m *MockAgent) ModifyEnergy(energyDelta common_types.Energy) {
}

func (m *MockAgent) Tag() {
}

func (m *MockAgent) Execute() {
}

func (m *MockAgent) String() string {
	return ""
}
func TestLinearAgentComparator(t *testing.T) {
	sut := comparator.NewLinearAgentComparator()

	t.Run("Test base cases", func(t *testing.T) {
		testParams := []struct {
			firstAgent  i_agent.IAgent
			secondAgent i_agent.IAgent
			result      bool
		}{
			{&MockAgent{1}, &MockAgent{2}, false},
			{&MockAgent{4}, &MockAgent{3}, true},
		}

		for _, param := range testParams {
			result := sut.Compare(param.firstAgent, param.secondAgent)
			if result != param.result {
				t.Errorf("Error in agent comparison, for agents: %s and %s.", param.firstAgent, param.secondAgent)
			}
		}
	})

}
