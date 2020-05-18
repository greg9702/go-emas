package common_types

type AgentId int
type Solution int
type Fitness int
type Energy int
type ActionTag string

const (
	Death        ActionTag = "Death"
	Reproduction           = "Reproduction"
	Fight                  = "Fight"
)
