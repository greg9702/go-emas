package i_agent

import "go-emas/pkg/common_types"

type IAgent interface {
	Id() common_types.AgentId
	Solution() common_types.Solution
	ActionTag() common_types.ActionTag
	Energy() common_types.Energy
	ModifyEnergy(energyDelta common_types.Energy)
	Tag()
	Execute()
	String() string
}
