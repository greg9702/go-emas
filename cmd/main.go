package main

import (
	"go-emas/pkg/environment"
	"go-emas/pkg/population_factory"
	"go-emas/pkg/randomizer"
	"go-emas/pkg/stopper"
)

const populationSize = 4

func main() {
	var populationFactory = population_factory.NewBasicPopulationFactroy()

	env, err := environment.NewEnvironment(populationSize, populationFactory, stopper.NewIterationBasedStopper(), randomizer.BaseRand())

	if err != nil {
		panic("Environment setup error")
	}
	env.Start()
}
