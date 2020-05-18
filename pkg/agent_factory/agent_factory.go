package agent_factory

import (
	"go-emas/pkg/agent"
	"go-emas/pkg/randomizer"
)

// IAgnetFactory interface for agent factories
type IAgnetFactory interface {
	CreateSingleAgent() *agent.Agent
}

// BaseAgentFactory is a basic variant of IAgentFactory
type BaseAgentFactory struct {
	randomizer *randomizer.IRandomizer
	id         int
}

// NewBaseAgentFactory creates new BaseAgentFactory object
func NewBaseAgentFactory(randomizer *randomizer.IRandomizer) {
	b := BaseAgentFactory{randomizer, 1}
}

// CreateSingleAgent creates single agent.Agnet object
func (b *BaseAgentFactory) CreateSingleAgent() *agent.Agent {
	a := agent.NewAgent(id)
	id++
	return &a
}
