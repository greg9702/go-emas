package fitness_calculator

import (
	"go-emas/pkg/solution"
)

// IFitnessCalculator is an interface for fitness calculators
type IFitnessCalculator interface {
	CalculateFitness(solution solution.ISolution) int
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
func (flc *LinearFitnessCalculator) CalculateFitness(sol solution.ISolution) int {
	return int(sol.(*solution.IntSolution).Solution())
}

// BitSetFitnessCalculator represents function that counts set bits
type BitSetFitnessCalculator struct {
}

// NewBitSetFitnessCalculator creates new BitSetFitnessCalculator object
func NewBitSetFitnessCalculator() *BitSetFitnessCalculator {
	l := BitSetFitnessCalculator{}
	return &l
}

// CalculateFitness calculate fitness value for passed soultion argument - count bits that are set
func (flc *BitSetFitnessCalculator) CalculateFitness(sol solution.ISolution) int {
	return int(sol.(*solution.BitSetSolution).Solution().Count())
}
