package fitness_calculator

import "go-emas/pkg/common_types"

// IFitnessCalculator is an interface for fitness calculators
type IFitnessCalculator interface {
	CalculateFitness(solution common_types.Solution) int
}

// LinearFitnessCalculator represents linear function
type LinearFitnessCalculator struct {
}

// NewLinearFitnessCalculator creates new LinearFitnessCalculator object
func NewLinearFitnessCalculator() LinearFitnessCalculator {
	return LinearFitnessCalculator{}
}

// CalculateFitness calculate fitness value for passed soultion argument
func (flc LinearFitnessCalculator) CalculateFitness(solution common_types.Solution) int {
	// TODO this cast cannot be used like this here
	return int(solution)
}
