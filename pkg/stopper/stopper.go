package stopper

// IStopper interface for all stopper modules
type IStopper interface {
	Stop(iteration int) bool
}

const maxIters = 1000

// IterationBasedStopper is a stopper which use
type IterationBasedStopper struct {
}

// NewIterationBasedStopper creates new IterationBasedStopper object
func NewIterationBasedStopper() *IterationBasedStopper {
	i := IterationBasedStopper{}
	return &i
}

// Stop returns true when iteration greater or equal maxIters
func (i *IterationBasedStopper) Stop(iteration int) bool {
	return iteration >= maxIters
}
