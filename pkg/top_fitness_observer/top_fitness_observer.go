package top_fitness_observer

import (
	"go-emas/pkg/i_agent"
	"go-emas/pkg/solution"
)

type ITopFitnessObserver interface {
	Update(agent i_agent.IAgent)
	TopFitness() int
}

type TopFitnessObserver struct {
	topFitness  int
	topSolution solution.ISolution
	agentId     int64
}

func NewTopFitnessObserver() *TopFitnessObserver {
	t := TopFitnessObserver{0, solution.NewIntSolution(0), 0}
	return &t
}

func (t *TopFitnessObserver) Update(agent i_agent.IAgent) {
	if agent.Fitness() > t.topFitness {
		t.topFitness = agent.Fitness()
		t.topSolution = agent.Solution()
		t.agentId = agent.ID()
	}
}

func (t *TopFitnessObserver) TopFitness() int {
	return t.topFitness
}
