package service

import (
	"github.com/daydoing/trade/config"
	"github.com/sirupsen/logrus"
)

type Context struct {
	Config *config.Config
	Logger *logrus.Logger
}

func NewContext(cfgFile string) (*Context, error) {
	c, err := initViper(cfgFile)
	if err != nil {
		return nil, err
	}

	l, err := initLogger(c.Sentry)
	if err != nil {
		return nil, err
	}

	return &Context{Config: c, Logger: l}, nil
}

func (srv *Context) Gc() {}
