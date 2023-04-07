package futures

import (
	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/spf13/cobra"

	"github.com/daydoing/trade/context"
	"github.com/daydoing/trade/strategies"
)

func FuturesMarketCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "futures",
		Short: "Running futures trading strategies",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				pairs           = ctx.Config.System.Pairs
				pair            = ctx.Config.Feed.Pair
				strategyName    = ctx.Config.Strategy.Name
				binanceKey      = ctx.Config.Binance.Key
				binanceSecret   = ctx.Config.Binance.Secret
				binanceLeverage = ctx.Config.Binance.Leverage
				telegramToken   = ctx.Config.Telegram.Token
				telegramUser    = ctx.Config.Telegram.UID
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

			strategy, err := strategies.Strategy(strategyName, ctx)
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
