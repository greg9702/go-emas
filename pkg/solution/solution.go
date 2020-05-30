package solution

import (
	"errors"
	"go-emas/pkg/common"
	"go-emas/pkg/randomizer"

	"strconv"

	"github.com/willf/bitset"
)

// ISolution is a general type representing agent's solution
type ISolution interface {
	String() string
	Mutate() ISolution
}

type IntSolution struct {
	solution int
}

// NewRandomIntSolution returns solution of type int. The solution value is random and does not exceed maxSolutionValue
func NewRandomIntSolution(maxSolutionValue int) (*IntSolution, error) {
	if maxSolutionValue < 0 {
		return nil, errors.New("maxSolutionValue must be a positive number")
	}
	randomizer := randomizer.BaseRand()
	solutionValue, _ := randomizer.RandInt(0, maxSolutionValue)
	return NewIntSolution(solutionValue), nil
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
	solutionDelta, _ := randomizer.BaseRand().RandInt(-int(float32(i.solution)*common.MutationRate), int(float32(i.solution)*common.MutationRate))
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
	if setBits > common.BitSetLength {
		return nil, errors.New("setBits can not exceed bitset length!")
	}
	randomizer := randomizer.BaseRand()
	solutionValue := bitset.New(common.BitSetLength)
	for solutionValue.Count() < setBits {
		idx, _ := randomizer.RandInt(0, common.BitSetLength-1)
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

// Mutate returns similar ISolution that differs from the original one. It does not modify the original object
func (i BitSetSolution) Mutate() ISolution {
	idx, _ := randomizer.BaseRand().RandInt(0, common.BitSetLength-1)
	oldSolution := i.Solution().Clone()
	return NewBitSetSolution(*oldSolution.Flip(uint(idx)))
}

// String used to display solution - displays indexes of set bits
func (i BitSetSolution) String() string {
	// Alternative - display bits TODO decide what is much dercriptive
	bitsView := i.solution.DumpAsBits()
	return bitsView[len(bitsView)-common.BitSetLength-1:]
	// return i.solution.String()
}
