package fitness_calculator

import "go-emas/pkg/solution"

// IFitnessCalculator is an interface for fitness calculators
type IFitnessCalculator interface {
	CalculateFitness(solution solution.Solution) int
}

// LinearFitnessCalculator represents linear function
type LinearFitnessCalculator struct {
}

// NewLinearFitnessCalculator creates new LinearFitnessCalculator object
func NewLinearFitnessCalculator() *LinearFitnessCalculator {
	l := LinearFitnessCalculator{}
	return &l
}

// CalculateFitness calculate fitness value for passed soultion argument
func (flc *LinearFitnessCalculator) CalculateFitness(sol solution.Solution) int {
	// TODO this cast cannot be used like this here
	return int(sol.(solution.IntSolution).Solution())
}
