package futures

import (
	"context"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/spf13/cobra"

	"github.com/daydoing/trade/service"
	"github.com/daydoing/trade/strategies"
)

func FuturesMarketCommand(srv *service.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "futures",
		Short: "Running futures trading strategies",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				ctx             = context.Background()
				pairs           = srv.Config.System.Pairs
				pair            = srv.Config.Feed.Pair
				strategyName    = srv.Config.Strategy.Name
				binanceKey      = srv.Config.Binance.Key
				binanceSecret   = srv.Config.Binance.Secret
				binanceLeverage = srv.Config.Binance.Leverage
				telegramToken   = srv.Config.Telegram.Token
				telegramUser    = srv.Config.Telegram.UID
			)

			settings := ninjabot.Settings{
				Pairs: pairs,
				Telegram: ninjabot.TelegramSettings{
					Enabled: true,
					Token:   telegramToken,
					Users:   []int{telegramUser},
				},
			}

			binance, err := exchange.NewBinanceFuture(
				ctx,
				exchange.WithBinanceFutureCredentials(binanceKey, binanceSecret),
				exchange.WithBinanceFutureLeverage(pair, binanceLeverage, exchange.MarginTypeIsolated),
			)
			if err != nil {
				return err
			}

			strategy, err := strategies.Strategy(strategyName)
			if err != nil {
				return err
			}

			bot, err := ninjabot.NewBot(ctx, settings, binance, strategy)
			if err != nil {
				return err
			}

			return bot.Run(ctx)
		},
	}
}
