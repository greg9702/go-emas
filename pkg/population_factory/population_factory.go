package population_factory

import "go-emas/pkg/agent"

// IPopulationFactory interface for population factories
type IPopulationFactory interface {
	CreatePopulation(populationSize int) (map[int]agent.Agent, error)
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
func (b *BasicPopulationFactroy) CreatePopulation(populationSize int) (map[int]agent.Agent, error) {
	var population = make(map[int]agent.Agent)
	for i := 0; i < populationSize; i++ {
		population[i] = *agent.NewAgent(i)
	}
	return population, nil
}
