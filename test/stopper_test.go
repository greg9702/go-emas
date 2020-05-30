package test

import (
	"go-emas/pkg/common"
	"go-emas/pkg/stopper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStopper(t *testing.T) {

	t.Run("Test Fitness Based Stopper", func(t *testing.T) {
		mockTopFitnessObserver := &MockTopFitnessObserver{}
		sut := stopper.NewTopFitnessBasedStopper(mockTopFitnessObserver)

		testParams := []struct {
			topFitness int
			shouldStop bool
		}{
			{common.TopFitnessThreshold - 1, false},
			{common.TopFitnessThreshold, false},
			{common.TopFitnessThreshold + 1, true},
		}

		for _, param := range testParams {
			mockTopFitnessObserver.On("TopFitness").Return(param.topFitness).Once()

			assert.Equal(t, sut.Stop(), param.shouldStop, "Fitness based stopper should stop only when some agent's solution has strongly bigger fitness than TopFitnessThreshold")
		}
	})

}
