package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"localhost.com/deep-learnning-algorithmes/chapter1_perceptron_model"
)

// tips
var (
	appName    = "deep-learnning-algorithmes"
	appVersion = "v1.0.0"
	v          *viper.Viper
	logger     *logrus.Logger
)

func init() {
	tips()
	initViper()
	//printConfig()
	initLogger()

	exit()
}

func main() {
	chapter1_perceptron_model.Perceptron()
}
