package environment

import (
	"errors"
	"fmt"
	"go-emas/pkg/agent"
	"go-emas/pkg/agent_factory"
)

// Environment is struct representing environment
type Environment struct {
	population   map[int]agent.Agent
	agentFactory *agent_factory.IAgnetFactory
}

// NewEnvironment creates new Environment object
func NewEnvironment(size int, agentFactory *agent_factory.IAgnetFactory) *Environment {
	var e = Environment{size, agentFactory}
	return &e
}

// PopulationSize return current size of poulation
func (e Environment) PopulationSize() int {
	// TODO something like pupulationMutex?
	return len(e.population)
}

// DeleteFromPopulation used to delete agent from map by id
// passing as callback to Agent
func (e Environment) DeleteFromPopulation(id int) error {
	// TODO use pupulationMutex
	_, ok := e.population[id]
	if ok {
		delete(e.population, id)
	} else {
		return errors.New("Element with %d id do not exist", id)
	}
}

// ShowMap is a helper used to display current state of a population
func (e Environment) ShowMap() {
	fmt.Println("[Environment] ", e.population)
}
