package service

import (
	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"

	"github.com/daydoing/trade/config"
)

func initLogger(sentry config.Sentry) (*logrus.Logger, error) {
	hook, err := logrus_sentry.NewSentryHook(sentry.DSN, []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	})
	if err != nil {
		return nil, err
	}

	hook.StacktraceConfiguration.Enable = true
	hook.Timeout = 0

	logger := logrus.New()
	logger.AddHook(hook)

	return logger, nil
}
