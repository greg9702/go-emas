package agent

import (
	"go-emas/pkg/common_types"
	"go-emas/pkg/tag_calculator"

	"strconv"
)

type IAgent interface {
	Id() common_types.AgentId
	Solution() common_types.Solution
	ActionTag() common_types.ActionTag
	Energy() common_types.Energy
	ModifyEnergy(energyDelta common_types.Energy)
	Tag()
	Execute()
	String() string
}

type Agent struct {
	id            common_types.AgentId
	solution      common_types.Solution
	actionTag     common_types.ActionTag
	energy        common_types.Energy
	tagCalculator tag_calculator.ITagCalulator
}

func NewAgent(id common_types.AgentId,
	solution common_types.Solution,
	actionTag common_types.ActionTag,
	energy common_types.Energy,
	tagCalculator tag_calculator.ITagCalulator) *Agent {
	a := Agent{id, solution, actionTag, energy, tagCalculator}
	return &a
}

func (a Agent) Solution() common_types.Solution {
	return a.solution
}

func (a Agent) ActionTag() common_types.ActionTag {
	return a.actionTag
}

func (a Agent) Energy() common_types.Energy {
	return a.energy
}

// TODO remove cast if possible
func (a Agent) String() string {
	return "Agent: " + strconv.Itoa(int(a.id)) + " solution: " + strconv.Itoa(int(a.solution))
}

func (a *Agent) ModifyEnergy(energyDelta common_types.Energy) {
	if a.energy+energyDelta < 0 {
		return
	}
	a.energy += energyDelta
}

func (a *Agent) Tag() {
	a.actionTag = a.tagCalculator.Calculate(a.energy)
}

func (a *Agent) Fight() {

}

func (a *Agent) Reproduce() {

}

func (a *Agent) Die() {

}
