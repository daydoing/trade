package cmd

import (
	"context"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/rodrigo-brito/ninjabot/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/daydoing/trade/strategies"
)

func backtestingCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "backtesting",
		Short: "Backtesting is to backtest the trading strategy",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				ctx         = context.Background()
				feedPair    = viper.GetString("pair")
				feedFile    = viper.GetString("file")
				assetSymbol = viper.GetString("symbol")
				assetAmount = viper.GetFloat64("amount")
			)

			settings := ninjabot.Settings{
				Pairs: []string{feedPair},
			}

			strategy := strategies.NewStrategy()

			csvFeed, err := exchange.NewCSVFeed(
				strategy.Timeframe(),
				exchange.PairFeed{
					Pair:      feedPair,
					File:      feedFile,
					Timeframe: strategy.Timeframe(), // specify the dataset timeframe
				},
			)
			if err != nil {
				return err
			}

			storage, err := storage.FromMemory()
			if err != nil {
				return err
			}

			wallet := exchange.NewPaperWallet(
				ctx,
				assetSymbol,
				exchange.WithPaperAsset(assetSymbol, assetAmount),
				exchange.WithDataFeed(csvFeed),
			)

			bot, err := ninjabot.NewBot(
				ctx,
				settings,
				wallet,
				strategy,
				ninjabot.WithBacktest(wallet),
				ninjabot.WithStorage(storage),
			)
			if err != nil {
				return err
			}

			if err := bot.Run(ctx); err != nil {
				return err
			}

			bot.Summary()

			return nil
		},
	}
}
