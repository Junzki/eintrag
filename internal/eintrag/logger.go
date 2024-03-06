package eintrag

import (
	"github.com/sirupsen/logrus"
)

var LOG *logrus.Logger = InitLogger()

func InitLogger() *logrus.Logger {
	logger := logrus.New()
	return logger
}
