package solution

import (
	"errors"
	"fmt"
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
	solutionDelta, _ := randomizer.BaseRand().RandInt(-int(float64(i.solution)*common.MutationRate), int(float64(i.solution)*common.MutationRate))
	return NewIntSolution(i.solution + solutionDelta)
}

// String used to display solution
func (i IntSolution) String() string {
	return "Solution: " + strconv.Itoa(i.solution)
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
	bitsView := i.solution.DumpAsBits()
	return "Solution: " + bitsView[len(bitsView)-common.BitSetLength-1:]
}

type PairSolution struct {
	x1 float64
	x2 float64
}

// NewPairSolution returns solution of type pair with value passed (solution)
func NewPairSolution(x1 float64, x2 float64) *PairSolution {
	i := PairSolution{x1, x2}
	return &i
}

// NewRandomPairSolution returns random solution of type pair. User has to specify the range from the points will be
func NewRandomPairSolution(min float64, max float64) (*PairSolution, error) {
	if min > max {
		return nil, errors.New("Error in PairSolution creation - min > max!")
	}
	randomizer := randomizer.BaseRand()

	x1, _ := randomizer.RandFloat64(min, max)
	x2, _ := randomizer.RandFloat64(min, max)

	agentSolution := NewPairSolution(x1, x2)
	return agentSolution, nil
}

// Solution returns solution
func (i PairSolution) Solution() (float64, float64) {
	return i.x1, i.x2
}

// Mutate returns similar ISolution that differs from the original one. It does not modify the original object
func (i PairSolution) Mutate() ISolution {
	// x1Delta, _ := randomizer.BaseRand().RandFloat64(-(float64(i.x1) * common.MutationRate), (float64(i.x1) * common.MutationRate))
	// x2Delta, _ := randomizer.BaseRand().RandFloat64(-(float64(i.x2) * common.MutationRate), (float64(i.x2) * common.MutationRate))
	x1Delta, _ := randomizer.BaseRand().RandFloat64(-common.MaxMutationDelta, common.MaxMutationDelta)
	x2Delta, _ := randomizer.BaseRand().RandFloat64(-common.MaxMutationDelta, common.MaxMutationDelta)

	return NewPairSolution(i.x1+x1Delta, i.x2+x2Delta)
}

// String used to display solution - displays indexes of set bits
func (i PairSolution) String() string {
	return "x1: " + fmt.Sprintf("%f", i.x1) + " x2: " + fmt.Sprintf("%f", i.x2)
}
