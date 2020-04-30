package pkg

import "fmt"

type Environment struct {
	population map[int]Agent
}

func NewEnvironment(size int) Environment {
	var e Environment
	e.population = make(map[int]Agent)
	e.initMap(size)
	return e
}

func (e Environment) ShowMap() {
	fmt.Println("[Environment] ", e.population)
}

func (e Environment) initMap(size int) {
	for i := 0; i < size; i++ {
		e.population[i] = NewAgent(i)
	}
}

func (e Environment) deleteFromMap(id int) {

	if _, ok := e.population[id]; ok {
		delete(e.population, id)
		fmt.Println("[Environment] Deleted element", id)
	} else {
		fmt.Println("[Environment] Not found element", id)
	}
}

func (e Environment) RunExecutor() {
	fmt.Println("[Environment] Executor started")

	for _, val := range e.population {
		val.Run(e.deleteFromMap)
	}

	fmt.Println("[Environment] Executor finished")
}
