package test

import "github.com/stretchr/testify/mock"

type MockRandomizer struct {
	mock.Mock
}

func (m *MockRandomizer) RandInt(min int, max int) (int, error) {
	args := m.Called()
	// TODO fix these casts are everywhere
	return args.Int(0), nil
}
