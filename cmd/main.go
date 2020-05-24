package main

import (
	"flag"
	"go-emas/pkg/environment"
	"go-emas/pkg/logger"
	"go-emas/pkg/population_factory"
	"go-emas/pkg/randomizer"
	"go-emas/pkg/stopper"
	"os"
)

const populationSize = 4

func main() {

	logFilePath := flag.String("logFile", "/dev/null", "specify log file path")
	logLevel := flag.Int("logLevel", logger.InfoLogs|logger.DebugLogs|logger.WarningLogs|logger.ErrorLogs, "specify log level")
	flag.Parse()

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
