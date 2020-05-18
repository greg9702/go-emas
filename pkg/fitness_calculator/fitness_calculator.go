package fitness_calculator

// IFitnessCalculator is an interface for fitness calculators
type IFitnessCalculator interface {
	CalculateFitness(solution int) int
}

type LinearFitnessCalculator struct {
}

func NewLinearFitnessCalculator() LinearFitnessCalculator {
	return LinearFitnessCalculator{}
}

func (flc LinearFitnessCalculator) CalculateFitness(solution int) int {
	return solution
}
