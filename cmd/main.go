package main

import (
	"flag"
	"fmt"
	"go-emas/pkg/environment"
	"go-emas/pkg/logger"
	"go-emas/pkg/population_factory"
	"go-emas/pkg/randomizer"
	"go-emas/pkg/stopper"
	"os"
)

const populationSize = 1000

func usage() {
	fmt.Println("usage: go run main.go -logFile <PATH_TO_LOG_FILE> -logLevel <LOGLEVEL>")
}

func main() {

	logFilePath := flag.String("logFile", "", "specify log file path")
	logLevel := flag.Int("logLevel", -1, "specify log level")
	flag.Parse()

	if *logFilePath == "" || *logLevel == -1 {
		usage()
		return
	}

	lf, err := os.OpenFile(*logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic("Failed to open log file")
	}
	defer lf.Close()

	logger.BaseLog().InitLogger(*logLevel, lf)

	var populationFactory = population_factory.NewBasicPopulationFactroy()
	env, err := environment.NewEnvironment(populationSize, populationFactory, stopper.NewIterationBasedStopper(), randomizer.BaseRand())

	if err != nil {
		panic("Environment setup error")
	}
	env.Start()
}
