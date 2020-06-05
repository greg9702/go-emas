package test

import "github.com/stretchr/testify/mock"

type MockRandomizer struct {
	mock.Mock
}

func (m *MockRandomizer) RandInt(min int, max int) (int, error) {
	args := m.Called()
	return args.Int(0), nil
}

func (m *MockRandomizer) RandFloat64(min float64, max float64) (float64, error) {
	args := m.Called()
	// TODO fix these casts are everywhere
	return float64(args.Int(0)), nil
}
