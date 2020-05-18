package population_factory

import (
	"go-emas/pkg/agent_factory"
)

// IPopulationFactory interface for population factories
type IPopulationFactory interface {
	CreatePopulation(populationSize int) error
}

// BasicPopulationFactroy is a basic variant of IPopulationFactory
type BasicPopulationFactroy struct{
	agentFactory *agent_factory.AgentFactory
}

// NewBasicPopulationFactroy creates new BasicPopulationFactroy object
func NewBasicPopulationFactroy(agentFactory *agent_factory.AgentFactory) *BasicPopulationFactroy {
	b := BasicPopulationFactroy{agentFactory}
	return &b
}

// CreatePopulation is used to creating initial population
func (b* BasicPopulationFactroy) CreatePopulation(populationSize int) map[int]int, error {
	
	for i := 0; i < size; i++ {
		population[i] = agent.NewAgent(i)
	}
}