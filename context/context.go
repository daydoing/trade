package context

import (
	"context"

	"github.com/daydoing/trade/config"
	"github.com/sirupsen/logrus"
)

type Context struct {
	context.Context
	Config *config.Config
	Logger *logrus.Logger
}

func NewContext() (Context, error) {
	c, err := initViper()
	if err != nil {
		return Context{}, err
	}

	l, err := initLogger(c.Sentry)
	if err != nil {
		return Context{}, err
	}

	return Context{Context: context.Background(), Config: c, Logger: l}, nil
}

func (srv *Context) Gc() {}
