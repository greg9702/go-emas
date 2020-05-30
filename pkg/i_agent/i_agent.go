package i_agent

import "go-emas/pkg/solution"

// IAgent is an interface for agents
type IAgent interface {
	ID() int64
	Solution() solution.ISolution
	ActionTag() string
	Energy() int
	ModifyEnergy(energyDelta int)
	Tag()
	Execute()
	SetID(int64)
	String() string
}
