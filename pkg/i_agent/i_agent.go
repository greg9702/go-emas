package i_agent

import "go-emas/pkg/common_types"

// IAgent is an interface for agents
type IAgent interface {
	ID() int64
	Solution() common_types.Solution
	ActionTag() string
	Energy() int
	ModifyEnergy(energyDelta int)
	Tag()
	Execute()
	SetID(int64)
	String() string
}
