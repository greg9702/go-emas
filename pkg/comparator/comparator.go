package comparator

import (
	"go-emas/pkg/agent"
	"go-emas/pkg/fitness_calculator"
)

// IComparator is an interface for comparators
type IAgentComparator interface {
	Compare(firstAgent agent.IAgent, secondAgent agent.IAgent) bool
}

// Agent with higher solution wins
type LinearAgentComparator struct {
	fitnessCalculator fitness_calculator.IFitnessCalculator
}

func NewLinearAgentComparator() LinearAgentComparator {
	lac := LinearAgentComparator{fitness_calculator.NewLinearFitnessCalculator()}
	return lac
}

func (lac LinearAgentComparator) Compare(firstAgent agent.IAgent, secondAgent agent.IAgent) bool {
	return lac.fitnessCalculator.CalculateFitness(firstAgent.Solution()) > lac.fitnessCalculator.CalculateFitness(secondAgent.Solution())
}
