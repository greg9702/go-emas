package fitness_calculator

import (
	"go-emas/pkg/solution"
	"math"
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

// Dejong1FitnessCalculator represents function that counts set bits
type Dejong1FitnessCalculator struct {
}

// NewDejong1FitnessCalculator creates new Dejong1FitnessCalculator object
func NewDejong1FitnessCalculator() *Dejong1FitnessCalculator {
	l := Dejong1FitnessCalculator{}
	return &l
}

// CalculateFitness calculate fitness value for passed soultion argument - count bits that are set
func (flc *Dejong1FitnessCalculator) CalculateFitness(sol solution.ISolution) int {
	x1, x2 := sol.(*solution.PairSolution).Solution()
	fitness := 0 - (math.Pow(x1, 2) + math.Pow(x2, 2))

	return int(fitness)
}
