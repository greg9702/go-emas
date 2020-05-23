package agent

import (
	"fmt"
	"go-emas/pkg/common_types"
	"go-emas/pkg/comparator"
	"go-emas/pkg/i_agent"
	"go-emas/pkg/randomizer"
	"go-emas/pkg/tag_calculator"

	"strconv"
)

const lossPenalty int = 20
const mutationRate float32 = 0.5

// percent of current parent energy passed to a child as inital energy value
const energyPercentageToChild float32 = 0.5

// Agent struct represent an Agent
type Agent struct {
	id                    int64
	solution              common_types.Solution
	actionTag             string
	energy                int
	tagCalculator         tag_calculator.ITagCalulator
	agentComparator       comparator.IAgentComparator
	randomizer            randomizer.IRandomizer
	getAgentByTagCallback func(tag string) (i_agent.IAgent, error)
	deleteAgentCallback   func(id int64) error
	addAgentCallback      func(newAgent i_agent.IAgent) error
}

// NewAgent creates new Agent object
func NewAgent(
	id int64,
	solution common_types.Solution,
	actionTag string, energy int,
	tagCalculator tag_calculator.ITagCalulator,
	agentComparator comparator.IAgentComparator,
	randomizer randomizer.IRandomizer,
	getAgentByTagCallback func(tag string) (i_agent.IAgent, error),
	deleteAgentCallback func(id int64) error,
	addAgentCallback func(newAgent i_agent.IAgent) error) i_agent.IAgent {
	a := Agent{id, solution, actionTag, energy, tagCalculator, agentComparator, randomizer, getAgentByTagCallback, deleteAgentCallback, addAgentCallback}
	return &a
}

// ID returns id
func (a *Agent) ID() int64 {
	return a.id
}

// SetID is used to set agents id
func (a *Agent) SetID(id int64) {
	a.id = id
}

// Solution returns agent solution
func (a *Agent) Solution() common_types.Solution {
	return a.solution
}

// ActionTag returns agent tag
func (a *Agent) ActionTag() string {
	return a.actionTag
}

// Energy returns agent energy
func (a *Agent) Energy() int {
	return a.energy
}

// String used to display agent struct using fmt
func (a *Agent) String() string {
	return "Agent: " + strconv.Itoa(int(a.id)) + " solution: " + strconv.Itoa(int(a.solution)) + " energy: " + strconv.Itoa(a.energy) + "\n"
}

// ModifyEnergy is used to modify agent energy
func (a *Agent) ModifyEnergy(energyDelta int) {
	if a.energy+energyDelta < 0 {
		return
	}
	a.energy += energyDelta
}

// Tag returns tag of an agent
func (a *Agent) Tag() {
	a.actionTag = a.tagCalculator.Calculate(a.energy)
}

// Execute used to execute action on an agent
func (a *Agent) Execute() {
	switch at := a.actionTag; at {
	case "Death":
		a.die()
	case "Reproduction":
		a.reproduce()
	case "Fight":
		a.fight()
	}
}

// Fight is used to perform fight action
func (a *Agent) fight() {
	rival, err := a.getAgentByTagCallback(common_types.Fight)

	if err != nil {
		// TODO what about this cast?
		fmt.Println("[Agent] Agent with id " + strconv.Itoa(int(a.id)) + " could not perform fight - there is no rival for him")
		return
	}

	var won bool = a.agentComparator.Compare(a, rival)

	if won {
		a.ModifyEnergy(lossPenalty)
		rival.ModifyEnergy(-lossPenalty)
	} else {
		a.ModifyEnergy(-lossPenalty)
		rival.ModifyEnergy(lossPenalty)
	}
}

// Reproduce is used to perform fight action
func (a *Agent) reproduce() {
	// TODO get unique id - from environment?
	// TODO environment.addAgent should generate it
	var newAgentID int64 = a.id + 50

	solutionDelta, _ := a.randomizer.RandInt(-int(float32(a.solution)*mutationRate), int(float32(a.solution)*mutationRate))

	var newAgentSolution common_types.Solution = a.solution + common_types.Solution(solutionDelta)
	var newAgentEnergy int = int(float32(a.energy) * energyPercentageToChild) // TODO this must be int!

	child := NewAgent(newAgentID,
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

func (a *Agent) die() {
	a.deleteAgentCallback(a.id)
}
