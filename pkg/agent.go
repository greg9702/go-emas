package pkg

import "fmt"

type Agent struct {
	id int
}

func (a Agent) Run(deleter func(int)) {

	if a.id == 3 {
		fmt.Println("[Agent] Agent", a.id, "assigned for deletion by itself")
		deleter(a.id)
	} else {
		fmt.Println("[Agent] Agent", a.id, "executed")
	}
}

func NewAgent(id int) Agent {
	var a Agent
	a.id = id
	return a
}
