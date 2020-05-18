package main

import (
	"go-emas/pkg/environment"
	"go-emas/pkg/population_factory"
)

var POPULATION_SIZE = 4

func main() {
	var populationFactory = population_factory.NewBasicPopulationFactroy()

	var env environment.Environment = *environment.NewEnvironment(POPULATION_SIZE, populationFactory)
	env.ShowMap()
}
