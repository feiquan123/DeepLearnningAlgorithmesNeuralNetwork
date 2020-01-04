package chapter1_perceptron_model

import "github.com/sirupsen/logrus"

var (
	logger *logrus.Logger
)

//log
func SetLog(log *logrus.Logger) {
	logger = log
}
