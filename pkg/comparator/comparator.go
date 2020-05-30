package comparator

import (
	"go-emas/pkg/fitness_calculator"
	"go-emas/pkg/i_agent"
)

// IAgentComparator is an interface for comparators
type IAgentComparator interface {
	Compare(firstAgent i_agent.IAgent, secondAgent i_agent.IAgent) bool
}

// BasicAgentComparator treat agent with higher solution wins
type BasicAgentComparator struct {
	fitnessCalculator fitness_calculator.IFitnessCalculator
}

// NewBasicAgentComparator creates new BasicAgentComparator object
func NewBasicAgentComparator(fitness_calculator fitness_calculator.IFitnessCalculator) *BasicAgentComparator {
	lac := BasicAgentComparator{fitness_calculator}
	return &lac
}

// Compare method used to compare two agents, returns true if the first passed agnet won, false otherwise
func (lac *BasicAgentComparator) Compare(firstAgent i_agent.IAgent, secondAgent i_agent.IAgent) bool {
	return lac.fitnessCalculator.CalculateFitness(firstAgent.Solution()) > lac.fitnessCalculator.CalculateFitness(secondAgent.Solution())
}
