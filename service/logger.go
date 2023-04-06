package service

import (
	"github.com/sirupsen/logrus"
)

func initLogger() (*logrus.Logger, error) {
	return logrus.New(), nil
}
