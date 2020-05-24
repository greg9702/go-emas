package test

import (
	"go-emas/pkg/common_types"

	"github.com/stretchr/testify/mock"
)

type MockAgent struct {
	mock.Mock
}

func (m *MockAgent) ID() int64 {
	args := m.Called()
	// TODO fix these casts are everywhere
	return int64(args.Int(0))
}

func (m *MockAgent) SetID(id int64) {
	m.Called()
}

func (m *MockAgent) Solution() common_types.Solution {
	args := m.Called()
	// TODO this mock can be not sufficient if solution ever changes
	return common_types.Solution(args.Int(0))
}

func (m *MockAgent) ActionTag() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockAgent) Energy() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockAgent) ModifyEnergy(energyDelta int) {
}

func (m *MockAgent) Execute() {
	m.Called()
}

func (m *MockAgent) String() string {
	return ""
}

func (m *MockAgent) Tag() {
	m.Called()
}
