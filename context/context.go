package context

import (
	"context"

	"github.com/daoleno/uniswapv3-sdk/examples/helper"
	"github.com/daydoing/trade/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type Context struct {
	context.Context
	Config      *config.Config
	Logger      *logrus.Logger
	NotifyBot   *NotifyBot
	ChainClient *ethclient.Client
	ChainWallet *helper.Wallet
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

	n, err := initNotifyBot(c.System.Telegram.Token, c.System.Telegram.UID)
	if err != nil {
		return Context{}, err
	}

	return Context{Context: context.Background(), Config: c, Logger: l, NotifyBot: n}, nil
}

func (srv *Context) Gc() {}
