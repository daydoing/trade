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

	client, err := ethclient.Dial(c.RpcURL)
	if err != nil {
		return Context{}, err
	}

	n, err := initNotifyBot(c.System.Telegram.Token, c.System.Telegram.UID)
	if err != nil {
		return Context{}, err
	}

	/*
		wallet := helper.InitWallet(c.PrivateKey)
		if wallet == nil {
			return Context{}, errors.New("incorrect wallet key")
		}

	*/
	return Context{Context: context.Background(), Config: c, Logger: l, ChainClient: client, NotifyBot: n}, nil
}

func (srv *Context) Gc() {}
