package agent

// Agent struct
type Agent struct {
	id int
}

// NewAgent creates new Agent object
func NewAgent(id int) *Agent {
	a := Agent{id}
	return &a
}
