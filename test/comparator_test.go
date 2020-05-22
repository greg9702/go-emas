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

func (m *MockAgent) ID() int64 {
	var result int64
	return result
}

func (m *MockAgent) Solution() common_types.Solution {
	return m.solution
}

func (m *MockAgent) ActionTag() string {
	var result string
	return result
}

func (m *MockAgent) Energy() int {
	var result int
	return result
}

func (m *MockAgent) ModifyEnergy(energyDelta int) {
}

func (m *MockAgent) Tag() {
}

func (m *MockAgent) SetId(id int64) {
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
