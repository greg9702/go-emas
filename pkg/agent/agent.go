package agent

import "strconv"

// Agent struct
type Agent struct {
	id       int
	solution int
}

// NewAgent creates new Agent object
func NewAgent(id int, solution int) *Agent {
	a := Agent{id, solution}
	return &a
}

func (a Agent) Solution() int {
	return a.solution
}

func (a Agent) String() string {
	return "Agent: " + strconv.Itoa(a.id) + " solution: " + strconv.Itoa(a.solution)
}
