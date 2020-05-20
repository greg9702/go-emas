package population_factory

import (
	"go-emas/pkg/i_agent"
)

// IPopulationFactory interface for population factories
type IPopulationFactory interface {
	CreatePopulation(populationSize int) (map[int64]i_agent.IAgent, error)
}

// BasicPopulationFactroy is a basic variant of IPopulationFactory
type BasicPopulationFactroy struct {
}

// NewBasicPopulationFactroy creates new BasicPopulationFactroy object
func NewBasicPopulationFactroy() *BasicPopulationFactroy {
	b := BasicPopulationFactroy{}
	return &b
}

// CreatePopulation is used to creating initial population
func (b *BasicPopulationFactroy) CreatePopulation(populationSize int) (map[int64]i_agent.IAgent, error) {
	var population = make(map[int64]i_agent.IAgent)
	for i := 0; i < populationSize; i++ {
		// TODO
		// population[i] = *agent.NewAgent(i)
	}
	return population, nil
}
