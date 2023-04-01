package cmd

import (
	"context"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/daydoing/trade/strategies"
)

func futuresMarketCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "futures",
		Short: "Running futures trading strategies",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				ctx           = context.Background()
				feedPair      = viper.GetString("pair")
				strategyName  = viper.GetString("strategy")
				apiKey        = viper.GetString("binance.key")
				secretKey     = viper.GetString("binance.secret")
				leverage      = viper.GetInt("binance.leverage")
				telegramToken = viper.GetString("telegram.token")
				telegramUser  = viper.GetInt("telegram.uid")
			)

			settings := ninjabot.Settings{
				Pairs: []string{feedPair},
				Telegram: ninjabot.TelegramSettings{
					Enabled: true,
					Token:   telegramToken,
					Users:   []int{telegramUser},
				},
			}

			binance, err := exchange.NewBinanceFuture(
				ctx,
				exchange.WithBinanceFutureCredentials(apiKey, secretKey),
				exchange.WithBinanceFutureLeverage(feedPair, leverage, exchange.MarginTypeIsolated),
			)
			if err != nil {
				return err
			}

			strategy, err := strategies.NewStrategy(strategyName)
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
