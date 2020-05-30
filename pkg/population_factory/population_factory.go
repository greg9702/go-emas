package population_factory

import (
	"go-emas/pkg/agent"
	"go-emas/pkg/common"
	"go-emas/pkg/comparator"
	"go-emas/pkg/fitness_calculator"
	"go-emas/pkg/i_agent"
	"go-emas/pkg/randomizer"
	"go-emas/pkg/solution"
	"go-emas/pkg/tag_calculator"
)

// IPopulationFactory interface for population factories
type IPopulationFactory interface {
	CreatePopulation(populationSize int,
		getAgentByTagCallback func(tag string) (i_agent.IAgent, error),
		deleteAgentCallback func(id int64) error,
		addAgentCallback func(newAgent i_agent.IAgent) error) (map[int64]i_agent.IAgent, error)
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
func (b *BasicPopulationFactroy) CreatePopulation(populationSize int,
	getAgentByTagCallback func(tag string) (i_agent.IAgent, error),
	deleteAgentCallback func(id int64) error,
	addAgentCallback func(newAgent i_agent.IAgent) error) (map[int64]i_agent.IAgent, error) {

	var population = make(map[int64]i_agent.IAgent)
	randomizer := randomizer.BaseRand()
	tagCalculator := tag_calculator.NewTagCalculator()
	agentComparator := comparator.NewBasicAgentComparator(fitness_calculator.NewBitSetFitnessCalculator())

	for i := 0; i < populationSize; i++ {

		agentSolution, err := solution.NewRandomBitSetSolution(2)
		if err != nil {
			return nil, err
		}
		energy := 40
		population[int64(i)] = agent.NewAgent(int64(i),
			solution.ISolution(agentSolution),
			common.Fight,
			energy,
			tagCalculator,
			agentComparator,
			randomizer,
			getAgentByTagCallback,
			deleteAgentCallback,
			addAgentCallback)
	}
	return population, nil
}
