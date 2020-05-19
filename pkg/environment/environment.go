package environment

import (
	"errors"
	"fmt"
	"go-emas/pkg/agent"
	"go-emas/pkg/common_types"
	"go-emas/pkg/population_factory"
	"go-emas/pkg/stopper"
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
	population map[int]agent.IAgent
	stopper    stopper.IStopper
}

// NewEnvironment creates new Environment object
func NewEnvironment(populationSize int, populationFactory population_factory.IPopulationFactory) (*Environment, error) {

	population, err := populationFactory.CreatePopulation(populationSize)

	if err != nil {
		return nil, err
	}

	var e = Environment{
		population: population,
		stopper:    stopper.NewIterationBasedStopper(),
	}

	return &e, nil
}

// Start is an entry point method
func (e Environment) Start() error {

	var i int = 0

	for {
		i++

		// e.tagAgents()
		// e.executeActions()

		fmt.Println("Running...")

		if e.stopper.Stop(i) {
			fmt.Println("Stop condition met")
			break
		}
	}

	return nil
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
		return errors.New("Element with " + strconv.Itoa(id) + " id do not exist")
	}
	return nil
}

// AddToPopulation adds new record to population
func (e Environment) AddToPopulation(agent agent.IAgent) error {
	_, ok := e.population[agent.ID()]
	if ok {
		return errors.New("Element with " + strconv.Itoa(agent.ID()) + " id already exists")
	}
	e.population[agent.ID()] = agent
	return nil
}

// ShowMap is a helper used to display current state of a population
func (e Environment) ShowMap() {
	fmt.Println("[Environment] ", e.population)
}

func (e Environment) tagAgents() {
	// for _, agent := range e.population {
	// 	agent.Tag()
	// }
}

func (e Environment) executeActions() {
	for _, agent := range e.population {
		agent.Execute()
	}
}
