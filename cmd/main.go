package main

import (
	"go-emas/pkg/environment"
	"go-emas/pkg/population_factory"
)

const populationSize = 4

func main() {
	var populationFactory = population_factory.NewBasicPopulationFactroy()

	env, err := environment.NewEnvironment(populationSize, populationFactory)

	if err != nil {
		panic("Environment setup error")
	}
	env.Start()
}
