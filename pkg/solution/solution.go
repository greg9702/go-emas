package solution

import "strconv"

// Solution is used as solution value
type Solution interface {
	set(string)
	String() string
	Mutate() Solution
}

type IntSolution struct {
	solution int
}

func NewIntSolution(solution int) *IntSolution {
	i := IntSolution{solution}
	return &i
}

func (i IntSolution) set(string) {
}

func (i IntSolution) Solution() int {
	return i.solution
}

func (i IntSolution) Mutate() Solution {
	return IntSolution{i.solution + 1}
}

func (i IntSolution) String() string {
	return strconv.Itoa(i.solution)
}
