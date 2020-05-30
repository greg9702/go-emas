package solution

import (
	"errors"
	"go-emas/pkg/randomizer"
	"strconv"

	"github.com/willf/bitset"
)

const mutationRate = 0.5
const BitSetLength = 5

// Solution is used as solution value
type ISolution interface {
	String() string
	Mutate() ISolution
}

type IntSolution struct {
	solution int
}

// NewRandomIntSolution returns solution of type int. The solution value is random and does not exceed maxSolutionValue
func NewRandomIntSolution(maxSolutionValue int) *IntSolution {
	randomizer := randomizer.BaseRand()
	solutionValue, _ := randomizer.RandInt(0, maxSolutionValue)
	return NewIntSolution(solutionValue)
}

// NewIntSolution returns solution of type int with value passed (solution)
func NewIntSolution(solution int) *IntSolution {
	i := IntSolution{solution}
	return &i
}

// Solution returns solution
func (i IntSolution) Solution() int {
	return i.solution
}

// Mutate returns similar ISolution that differs from the original one. It does not modify the original object
func (i IntSolution) Mutate() ISolution {
	solutionDelta, _ := randomizer.BaseRand().RandInt(-int(float32(i.solution)*mutationRate), int(float32(i.solution)*mutationRate))
	return NewIntSolution(i.solution + solutionDelta)
}

// String used to display solution
func (i IntSolution) String() string {
	return strconv.Itoa(i.solution)
}

type BitSetSolution struct {
	solution bitset.BitSet
}

// NewRandomIntSolution returns random solution of type bitset. User has to specify how many bits should be set and they will be randomly splitted across the bitset
func NewRandomBitSetSolution(setBits uint) (*BitSetSolution, error) {
	if setBits > BitSetLength {
		return nil, errors.New("setBits can not exceed bitset length!")
	}
	randomizer := randomizer.BaseRand()
	solutionValue := bitset.New(BitSetLength)
	for solutionValue.Count() < setBits {
		idx, _ := randomizer.RandInt(0, BitSetLength-1)
		solutionValue.Set(uint(idx))
	}
	agentSolution := NewBitSetSolution(*solutionValue)

	return agentSolution, nil
}

// NewIntSolution returns solution of type bitset with value passed (solution)
func NewBitSetSolution(solution bitset.BitSet) *BitSetSolution {
	i := BitSetSolution{solution}
	return &i
}

// Solution returns solution
func (i BitSetSolution) Solution() *bitset.BitSet {
	return &i.solution
}

// TODO make random
func (i BitSetSolution) Mutate() ISolution {
	idx, _ := randomizer.BaseRand().RandInt(0, BitSetLength-1)
	oldSolution := i.Solution().Clone()
	return NewBitSetSolution(*oldSolution.Flip(uint(idx)))
}

// String used to display solution
func (i BitSetSolution) String() string {
	return i.solution.String()
}
