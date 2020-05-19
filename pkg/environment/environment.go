package environment

import (
	"errors"
	"fmt"
	"go-emas/pkg/agent"
	"go-emas/pkg/common_types"
	"go-emas/pkg/population_factory"
	"strconv"
)

// IEnvironment interface for environments
type IEnvironment interface {
	Start() error
	PopulationSize() int
	DeleteFromPopulation(id int) error
	AddToPopulation(agent agent.IAgent) error
	ShowMap()
	GetAgentByTag(tag common_types.ActionTag) agent.IAgent
}

// Environment is struct representing environment
type Environment struct {
	populationSize int
	population     map[int]agent.Agent
}

// NewEnvironment creates new Environment object
func NewEnvironment(size int, populationFactory population_factory.IPopulationFactory) *Environment {
	population, _ := populationFactory.CreatePopulation(size)
	var e = Environment{size, population}
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
		return errors.New("Element with " + strconv.Itoa(id) + "id do not exist")
	}
	return nil
}

// ShowMap is a helper used to display current state of a population
func (e Environment) ShowMap() {
	fmt.Println("[Environment] ", e.population)
}
