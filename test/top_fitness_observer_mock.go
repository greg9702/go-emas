package test

import (
	"go-emas/pkg/i_agent"

	"github.com/stretchr/testify/mock"
)

type MockTopFitnessObserver struct {
	mock.Mock
}

func (m *MockTopFitnessObserver) Update(agent i_agent.IAgent) {
}

func (m *MockTopFitnessObserver) TopFitness() int {
	args := m.Called()
	return args.Int(0)
}
