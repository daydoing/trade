package service

import "github.com/daydoing/trade/config"

type Context struct {
	Config *config.Config
}

func NewContext(cfgFile string) (*Context, error) {
	c, err := initViper(cfgFile)
	if err != nil {
		return nil, err
	}

	return &Context{Config: c}, nil
}

func (srv *Context) Gc() {}
