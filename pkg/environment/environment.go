package environment

import (
	"bufio"
	"errors"
	"fmt"
	"go-emas/pkg/common_types"
	"go-emas/pkg/i_agent"
	"go-emas/pkg/population_factory"
	"go-emas/pkg/randomizer"
	"go-emas/pkg/stopper"
	"os"
	"strconv"
)

// IEnvironment interface for environments
type IEnvironment interface {
	Start() error
	PopulationSize() int
	DeleteFromPopulation(id int64) error
	AddToPopulation(agent i_agent.IAgent) error
	ShowMap()
	GetAgentByTag(tag string) (i_agent.IAgent, error)
}

// Environment is struct representing environment
type Environment struct {
	population          map[int64]i_agent.IAgent
	agentsBeforeActions map[string][]i_agent.IAgent
	stopper             stopper.IStopper
	randomizer          randomizer.IRandomizer
}

// NewEnvironment creates new Environment object
func NewEnvironment(populationSize int,
	populationFactory population_factory.IPopulationFactory,
	stopper stopper.IStopper,
	randomizer randomizer.IRandomizer) (*Environment, error) {

	var e = Environment{
		population: make(map[int64]i_agent.IAgent),
		stopper:    stopper,
		randomizer: randomizer,
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
		fmt.Printf("############ Iteration number: %d ############\n", i)

		fmt.Println("--------------------------")
		fmt.Println("Agents before iteration: ")
		fmt.Println("--------------------------")

		e.ShowMap()

		fmt.Println("--------------------------")
		fmt.Println("Events: ")
		fmt.Println("--------------------------")

		e.TagAgents()
		e.executeActions()

		if e.stopper.Stop(i) {
			fmt.Println("Stop condition met")
			break
		}
		i++
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println("")
	}

	return nil
}

// PopulationSize return current size of population
func (e *Environment) PopulationSize() int {
	// TODO something like pupulationMutex?
	return len(e.population)
}

// removeAgentFromWaitingQueue mark agent action as done
func (e *Environment) removeAgentFromWaitingQueue(action string, agentIndex int) {
	e.agentsBeforeActions[action] = append(e.agentsBeforeActions[action][:agentIndex], e.agentsBeforeActions[action][agentIndex+1:]...)
}

// GetAgentByTag return random agent which has a given tag and has not yet performed its action in the turn. Mark the action of agent as done
func (e *Environment) GetAgentByTag(actionTag string) (i_agent.IAgent, error) {
	if len(e.agentsBeforeActions[actionTag]) > 0 {
		agentIndex, _ := randomizer.BaseRand().RandInt(0, len(e.agentsBeforeActions[actionTag])-1)
		agent := e.agentsBeforeActions[actionTag][agentIndex]
		e.removeAgentFromWaitingQueue(actionTag, agentIndex)
		return agent, nil
	}
	return nil, errors.New("There is no agent to perform action: " + actionTag)
}

// DeleteFromPopulation used to delete agent from map by id passed as callback to Agent
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
	agent.SetID(e.getMaxAgentID() + 1) // TODO improve IDs generation
	_, ok := e.population[agent.ID()]
	if ok {
		return errors.New("Element with " + strconv.FormatInt(agent.ID(), 10) + " id already exists")
	}
	e.population[agent.ID()] = agent
	return nil
}

// ShowMap is a helper used to display current state of a population
func (e *Environment) ShowMap() {
	for _, v := range e.population {
		fmt.Println(v)
	}
}

// TagAgents each agent tags itself. Then all agents are marked to perform actions
func (e *Environment) TagAgents() {
	for _, agent := range e.population {
		agent.Tag()
	}
	e.agentsBeforeActions = make(map[string][]i_agent.IAgent)
	for _, agent := range e.population {
		e.agentsBeforeActions[agent.ActionTag()] = append(e.agentsBeforeActions[agent.ActionTag()], agent)
	}
}

func (e *Environment) executeActions() {
	actions := []string{common_types.Death, common_types.Reproduction, common_types.Fight}
	for _, action := range actions {
		for len(e.agentsBeforeActions[action]) > 0 {
			currentExecutor := e.agentsBeforeActions[action][0]
			e.agentsBeforeActions[action] = append(e.agentsBeforeActions[action][:0], e.agentsBeforeActions[action][1:]...)
			currentExecutor.Execute()
		}
	}
}

func (e *Environment) getMaxAgentID() int64 {
	var maxID int64 = 0
	for id := range e.population {
		if id > maxID {
			maxID = id
		}
	}
	return maxID
}
