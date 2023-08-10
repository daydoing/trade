package context

import (
	"context"

	"github.com/daoleno/uniswapv3-sdk/examples/helper"
	"github.com/daydoing/trade/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type BotContext struct {
	context.Context
	Config      *config.Config
	Logger      *logrus.Logger
	NotifyBot   *NotifyBot
	ChainClient *ethclient.Client
	ChainWallet *helper.Wallet
}

func NewBotContext() (BotContext, error) {
	c, err := initViper()
	if err != nil {
		return BotContext{}, err
	}

	l, err := initLogger(c.Sentry)
	if err != nil {
		return BotContext{}, err
	}

	n, err := initNotifyBot(c.System.Telegram.Token, c.System.Telegram.UID)
	if err != nil {
		return BotContext{}, err
	}

	return BotContext{Context: context.Background(), Config: c, Logger: l, NotifyBot: n}, nil
}

func (srv *BotContext) Gc() {}
