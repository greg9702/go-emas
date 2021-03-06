package agent

import (
	"go-emas/pkg/common"
	"go-emas/pkg/comparator"
	"go-emas/pkg/fitness_calculator"
	"go-emas/pkg/i_agent"
	"go-emas/pkg/logger"
	"go-emas/pkg/randomizer"
	"go-emas/pkg/solution"
	"go-emas/pkg/tag_calculator"

	"strconv"
)

// percent of current parent energy passed to a child as inital energy value
const energyPercentageToChild float32 = 0.5

// Agent struct represent an Agent
type Agent struct {
	id                    int64
	solution              solution.ISolution
	fitness               int
	actionTag             string
	energy                int
	tagCalculator         tag_calculator.ITagCalulator
	agentComparator       comparator.IAgentComparator
	fitnessCalculator     fitness_calculator.IFitnessCalculator
	randomizer            randomizer.IRandomizer
	getAgentByTagCallback func(tag string) (i_agent.IAgent, error)
	deleteAgentCallback   func(id int64) error
	addAgentCallback      func(newAgent i_agent.IAgent) error
}

// NewAgent creates new Agent object
func NewAgent(
	id int64,
	solution solution.ISolution,
	actionTag string, energy int,
	tagCalculator tag_calculator.ITagCalulator,
	agentComparator comparator.IAgentComparator,
	randomizer randomizer.IRandomizer,
	getAgentByTagCallback func(tag string) (i_agent.IAgent, error),
	deleteAgentCallback func(id int64) error,
	addAgentCallback func(newAgent i_agent.IAgent) error,
	fitnessCalculator fitness_calculator.IFitnessCalculator) i_agent.IAgent {
	fitness := fitnessCalculator.CalculateFitness(solution)
	a := Agent{id, solution, fitness, actionTag, energy, tagCalculator, agentComparator, fitnessCalculator, randomizer, getAgentByTagCallback, deleteAgentCallback, addAgentCallback}
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
func (a *Agent) Solution() solution.ISolution {
	return a.solution
}

func (a *Agent) Fitness() int {
	return a.fitness
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
	return "Agent [" + strconv.Itoa(int(a.id)) + "] " + a.solution.String() + " fitness: " + strconv.Itoa(a.Fitness()) + " energy: " + strconv.Itoa(a.energy)
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
	rival, err := a.getAgentByTagCallback(common.Fight)

	if err != nil {
		logger.BaseLog().Debug("[Figth] Agent [" + strconv.Itoa(int(a.id)) + "] there is no rival for him")
		return
	}

	var won bool = a.agentComparator.Compare(a, rival)
	var result string

	if won {
		result = "1:0"
		a.ModifyEnergy(common.LossPenalty)
		rival.ModifyEnergy(-common.LossPenalty)
	} else {
		result = "0:1"
		a.ModifyEnergy(-common.LossPenalty)
		rival.ModifyEnergy(common.LossPenalty)
	}
	logger.BaseLog().Debug("[Figth] Agent [" + strconv.Itoa(int(a.id)) + "] vs Agent [" + strconv.Itoa(int(rival.ID())) + "] result: " + result)
}

// Reproduce is used to perform fight action
func (a *Agent) reproduce() {
	var newAgentID int64

	newAgentSolution := a.solution.Mutate()

	var newAgentEnergy int = int(float32(a.energy) * energyPercentageToChild) // TODO this must be int!

	child := NewAgent(newAgentID,
		newAgentSolution,
		common.Fight,
		newAgentEnergy,
		a.tagCalculator,
		a.agentComparator,
		a.randomizer,
		a.getAgentByTagCallback,
		a.deleteAgentCallback,
		a.addAgentCallback,
		a.fitnessCalculator)

	a.addAgentCallback(child)
	a.ModifyEnergy(-newAgentEnergy)
	logger.BaseLog().Debug("[Reproduce] Agent [" + strconv.Itoa(int(a.id)) + "] reproduced, spawned Agent [" + strconv.Itoa(int(child.ID())) + "]")
}

func (a *Agent) die() {
	logger.BaseLog().Debug("[Die] Agent [" + strconv.Itoa(int(a.id)) + "] died")
	a.deleteAgentCallback(a.id)
}
