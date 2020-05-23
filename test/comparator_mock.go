package test

import (
	"go-emas/pkg/i_agent"

	"github.com/stretchr/testify/mock"
)

type MockAgentComparator struct {
	mock.Mock
}

func (m *MockAgentComparator) Compare(firstAgent i_agent.IAgent, secondAgent i_agent.IAgent) bool {
	args := m.Called()
	return args.Bool(0)
}
