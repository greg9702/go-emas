package main

import "go-emas/pkg"

var POPULATION_SIZE = 4

func main() {

	var env pkg.Environment = pkg.NewEnvironment(POPULATION_SIZE)
	env.ShowMap()
	env.RunExecutor()
	env.ShowMap()
}
