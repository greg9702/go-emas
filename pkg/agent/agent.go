package agent

import (
	"go-emas/pkg/common_types"
	"go-emas/pkg/comparator"
	"go-emas/pkg/i_agent"
	"go-emas/pkg/tag_calculator"

	"strconv"
)

const LOSS_PENALTY common_types.Energy = 20

type IAgent interface {
	Solution() common_types.Solution
	ActionTag() common_types.ActionTag
	Energy() common_types.Energy
	ModifyEnergy(energyDelta common_types.Energy)
	Execute()
	String() string
	ID() int
}

type Agent struct {
	id                    common_types.AgentId
	solution              common_types.Solution
	actionTag             common_types.ActionTag
	energy                common_types.Energy
	tagCalculator         tag_calculator.ITagCalulator
	agentComparator       comparator.IAgentComparator
	getAgentByTagCallback func(tag common_types.ActionTag) i_agent.IAgent
	deleteAgentCallback   func(id common_types.AgentId)
}

func NewAgent(id common_types.AgentId,
	solution common_types.Solution,
	actionTag common_types.ActionTag,
	energy common_types.Energy,
	tagCalculator tag_calculator.ITagCalulator,
	agentComparator comparator.IAgentComparator,
	getAgentByTagCallback func(tag common_types.ActionTag) i_agent.IAgent,
	deleteAgentCallback func(id common_types.AgentId)) i_agent.IAgent {
	a := Agent{id, solution, actionTag, energy, tagCalculator, agentComparator, getAgentByTagCallback, deleteAgentCallback}
	return &a
}

func (a Agent) Id() common_types.AgentId {
	return a.id
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
	var rival i_agent.IAgent = a.getAgentByTagCallback(common_types.Fight)
	var won bool = a.agentComparator.Compare(a, rival)
	if won {
		a.ModifyEnergy(LOSS_PENALTY)
		rival.ModifyEnergy(-LOSS_PENALTY)
	} else {
		a.ModifyEnergy(-LOSS_PENALTY)
		rival.ModifyEnergy(LOSS_PENALTY)
	}
}

func (a *Agent) Reproduce() {

}

func (a *Agent) Die() {
	a.deleteAgentCallback(a.id)
}

func (a *Agent) Execute() {
	switch at := a.actionTag; at {
	case "Death":
		a.Die()
	case "Reproduction":
		a.Reproduce()
	case "Fight":
		a.Fight()
	}
}
