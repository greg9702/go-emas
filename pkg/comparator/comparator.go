package comparator

import (
	"go-emas/pkg/i_agent"
)

// IAgentComparator is an interface for comparators
type IAgentComparator interface {
	Compare(firstAgent i_agent.IAgent, secondAgent i_agent.IAgent) bool
}

// BasicAgentComparator treat agent with higher solution wins
type BasicAgentComparator struct {
}

// NewBasicAgentComparator creates new BasicAgentComparator object
func NewBasicAgentComparator() *BasicAgentComparator {
	lac := BasicAgentComparator{}
	return &lac
}

// Compare method used to compare two agents, returns true if the first passed agnet won, false otherwise
func (lac *BasicAgentComparator) Compare(firstAgent i_agent.IAgent, secondAgent i_agent.IAgent) bool {
	return firstAgent.Fitness() > secondAgent.Fitness()
}
