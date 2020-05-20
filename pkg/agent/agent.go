package agent

import (
	"go-emas/pkg/common_types"
	"go-emas/pkg/comparator"
	"go-emas/pkg/i_agent"
	"go-emas/pkg/randomizer"
	"go-emas/pkg/tag_calculator"

	"strconv"
)

const lossPenalty int = 20
const mnutationRate float32 = 0.5

// percent of current parent energy passed to a child as inital energy value
const energyPercentageToChild float32 = 0.5

type Agent struct {
	id                    common_types.AgentId
	solution              common_types.Solution
	actionTag             common_types.ActionTag
	energy                common_types.Energy
	tagCalculator         tag_calculator.ITagCalulator
	agentComparator       comparator.IAgentComparator
	randomizer            randomizer.IRandomizer
	getAgentByTagCallback func(tag common_types.ActionTag) i_agent.IAgent
	deleteAgentCallback   func(id common_types.AgentId)
	addAgentCallback      func(newAgent i_agent.IAgent)
}

func NewAgent(id common_types.AgentId,
	solution common_types.Solution,
	actionTag common_types.ActionTag,
	energy common_types.Energy,
	tagCalculator tag_calculator.ITagCalulator,
	agentComparator comparator.IAgentComparator,
	randomizer randomizer.IRandomizer,
	getAgentByTagCallback func(tag common_types.ActionTag) i_agent.IAgent,
	deleteAgentCallback func(id common_types.AgentId),
	addAgentCallback func(newAgent i_agent.IAgent)) i_agent.IAgent {
	a := Agent{id, solution, actionTag, energy, tagCalculator, agentComparator, randomizer, getAgentByTagCallback, deleteAgentCallback, addAgentCallback}
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
	// TODO get unique id - from environment?
	var newAgentId common_types.AgentId = a.id + 50
	solutionDelta, _ := a.randomizer.RandInt(-int(float32(a.solution)*MUTATION_RATE), int(float32(a.solution)*MUTATION_RATE))
	var newAgentSolution common_types.Solution = a.solution + common_types.Solution(solutionDelta)
	var newAgentEnergy common_types.Energy = common_types.Energy(float32(a.energy) * ENERGY_PERCENTAGE_TRANSFERRED_TO_CHILD)
	child := NewAgent(newAgentId,
		newAgentSolution,
		common_types.Fight,
		newAgentEnergy,
		a.tagCalculator,
		a.agentComparator,
		a.randomizer,
		a.getAgentByTagCallback,
		a.deleteAgentCallback,
		a.addAgentCallback)
	a.addAgentCallback(child)
	a.ModifyEnergy(-newAgentEnergy)
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
