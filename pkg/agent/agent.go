package agent

import (
	"go-emas/pkg/common_types"
	"go-emas/pkg/tag_calculator"

	"strconv"
)

type IAgent interface {
	Solution() common_types.Solution
	ActionTag() common_types.ActionTag
	Energy() common_types.Energy
	ModifyEnergy(energyDelta common_types.Energy)
	Execute()
	String() string
}

// Agent struct
type Agent struct {
	id       int
	solution int
}

// NewAgent creates new Agent object
func NewAgent(id int, solution int) *Agent {
	a := Agent{id, solution}
	return &a
}

func (a Agent) Solution() int {
	return a.solution
}

func (a Agent) String() string {
	return "Agent: " + strconv.Itoa(a.id) + " solution: " + strconv.Itoa(a.solution)
}
