package stopper

import "go-emas/pkg/common"

// IStopper interface for all stopper modules
type IStopper interface {
	Stop() bool
}

// IterationBasedStopper is a stopper which use
type IterationBasedStopper struct {
	currentIteration int
}

// NewIterationBasedStopper creates new IterationBasedStopper object
func NewIterationBasedStopper() *IterationBasedStopper {
	i := IterationBasedStopper{0}
	return &i
}

// Stop returns true when iteration greater or equal maxIters
func (i *IterationBasedStopper) Stop() bool {
	i.currentIteration++
	return i.currentIteration >= common.MaxIters
}
