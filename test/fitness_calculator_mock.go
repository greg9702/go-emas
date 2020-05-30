package test

import (
	"go-emas/pkg/solution"

	"github.com/stretchr/testify/mock"
)

type MockFitnessCalculator struct {
	mock.Mock
}

func (m *MockFitnessCalculator) CalculateFitness(solution solution.ISolution) int {
	args := m.Called()
	return args.Int(0)
}
