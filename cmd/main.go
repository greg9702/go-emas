package main

import "go-emas/pkg/environment"

var POPULATION_SIZE = 4

func main() {

	var env environment.Environment = environment.NewEnvironment(POPULATION_SIZE)
	env.ShowMap()
	env.RunExecutor()
	env.ShowMap()
}
