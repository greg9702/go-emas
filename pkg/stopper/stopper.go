package stopper

import "go-emas/pkg/common"

// IStopper interface for all stopper modules
type IStopper interface {
	Stop() bool
}

// IterationBasedStopper is a stopper which counts iterations
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

// // TopFitnessBasedStopper is a stopper which monitors the top fitness among agents
// type TopFitnessBasedStopper struct {
// }

// // NewTopFitnessBasedStopper creates new TopFitnessBasedStopper object
// func NewTopFitnessBasedStopper() *TopFitnessBasedStopper {
// 	i := TopFitnessBasedStopper{}
// 	return &i
// }

// // Stop returns true when the top solutions among the agents is higher than threshold
// func (i *TopFitnessBasedStopper) Stop() bool {
// 	return i.currentIteration >= common.TopFitnessThreshold
// }
