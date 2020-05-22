package environment

import (
	"errors"
	"fmt"
	"go-emas/pkg/i_agent"
	"go-emas/pkg/population_factory"
	"go-emas/pkg/stopper"
	"strconv"
)

// IEnvironment interface for environments
type IEnvironment interface {
	Start() error
	PopulationSize() int
	DeleteFromPopulation(id int64) error
	AddToPopulation(agent i_agent.IAgent) error
	ShowMap()
	GetAgentByTag(tag string) i_agent.IAgent
}

// Environment is struct representing environment
type Environment struct {
	population          map[int64]i_agent.IAgent
	agentsBeforeActions map[string][]i_agent.IAgent
	stopper             stopper.IStopper
}

// NewEnvironment creates new Environment object
func NewEnvironment(populationSize int, populationFactory population_factory.IPopulationFactory) (*Environment, error) {

	var e = Environment{
		population: make(map[int64]i_agent.IAgent),
		stopper:    stopper.NewIterationBasedStopper(),
	}

	population, err := populationFactory.CreatePopulation(populationSize,
		e.GetAgentByTag,
		e.DeleteFromPopulation,
		e.AddToPopulation)

	if err != nil {
		return nil, err
	}

	e.population = population

	return &e, nil
}

// Start is an entry point method
func (e *Environment) Start() error {

	var i int = 0

	for {
		i++

		e.tagAgents()

		e.agentsBeforeActions = make(map[string][]i_agent.IAgent)
		for _, agent := range e.population {
			e.agentsBeforeActions[agent.ActionTag()] = append(e.agentsBeforeActions[agent.ActionTag()], agent)
		}

		e.executeActions()
		e.ShowMap()
		fmt.Println("Running...")

		if e.stopper.Stop(i) {
			fmt.Println("Stop condition met")
			break
		}
	}

	return nil
}

// PopulationSize return current size of population
func (e *Environment) PopulationSize() int {
	// TODO something like pupulationMutex?
	return len(e.population)
}

// GetAgentByTag - return random Agent which has a given tag and has not yet performed its action in the turn
func (e *Environment) GetAgentByTag(tag string) i_agent.IAgent {
	for k := range e.population {
		if e.population[k].ActionTag() == tag {
			return e.population[k]
		}
	}
	return nil
}

// DeleteFromPopulation used to delete agent from map by id
// passing as callback to Agent
func (e *Environment) DeleteFromPopulation(id int64) error {
	// TODO use pupulationMutex
	_, ok := e.population[id]
	if ok {
		delete(e.population, id)
	} else {
		return errors.New("Element with " + strconv.FormatInt(id, 10) + " id do not exist")
	}
	return nil
}

// AddToPopulation adds new record to population
func (e *Environment) AddToPopulation(agent i_agent.IAgent) error {
	agent.SetId(int64(len(e.population))) // TODO improve IDs generation
	_, ok := e.population[agent.ID()]
	if ok {
		return errors.New("Element with " + strconv.FormatInt(agent.ID(), 10) + " id already exists")
	}
	e.population[agent.ID()] = agent
	return nil
}

// ShowMap is a helper used to display current state of a population
func (e *Environment) ShowMap() {
	fmt.Println("[Environment] ", e.population)
}

func (e *Environment) tagAgents() {
	for _, agent := range e.population {
		agent.Tag()
	}
}

func (e *Environment) executeActions() {
	// for len(e.agentsBeforeActions) > 0 {
	// 	action := getArbitraryAction(e.agentsBeforeActions)
	// 	fmt.Println(action)
	// 	fmt.Println(e.agentsBeforeActions[action][0])
	// 	e.agentsBeforeActions[action][0].Execute()
	// 	e.agentsBeforeActions[action][0]
	// }
}

func getArbitraryAction(agents map[string][]i_agent.IAgent) string {
	for k := range agents {
		return k
	}
	return ""
}
