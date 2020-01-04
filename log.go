package main

import (
	"localhost.com/deep-learnning-algorithmes/chapter1_perceptron_model"
	"localhost.com/deep-learnning-algorithmes/utils"
)

// set log
func initLogger() {
	logger = utils.NewLoggerFromMap(log_config)
	chapter1_perceptron_model.SetLog(logger)
}
