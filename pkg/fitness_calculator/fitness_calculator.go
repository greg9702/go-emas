package fitness_calculator

import "go-emas/pkg/common_types"

// IFitnessCalculator is an interface for fitness calculators
type IFitnessCalculator interface {
	CalculateFitness(solution common_types.Solution) common_types.Fitness
}

type LinearFitnessCalculator struct {
}

func NewLinearFitnessCalculator() LinearFitnessCalculator {
	return LinearFitnessCalculator{}
}

func (flc LinearFitnessCalculator) CalculateFitness(solution common_types.Solution) common_types.Fitness {
	return common_types.Fitness(solution)
}
