package spot

import (
	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/spf13/cobra"

	"github.com/daydoing/trade/context"
	"github.com/daydoing/trade/strategies"
)

func SpotMarketCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "spot",
		Short: "Running spot trading strategies",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				pairs         = ctx.Config.System.Pairs
				strategyName  = ctx.Config.Strategy.Name
				binanceKey    = ctx.Config.Binance.Key
				binanceSecret = ctx.Config.Binance.Secret
				telegramToken = ctx.Config.Telegram.Token
				telegramUser  = ctx.Config.Telegram.UID
			)

			settings := ninjabot.Settings{
				Pairs: pairs,
				Telegram: ninjabot.TelegramSettings{
					Enabled: true,
					Token:   telegramToken,
					Users:   []int{telegramUser},
				},
			}

			binance, err := exchange.NewBinance(ctx, exchange.WithBinanceCredentials(binanceKey, binanceSecret))
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
