package test

import "github.com/stretchr/testify/mock"

type MockStopper struct {
	mock.Mock
}

func (m *MockStopper) Stop() bool {
	args := m.Called()
	return args.Bool(0)
}
