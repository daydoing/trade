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

func paperwalletCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "paperwallet",
		Short: "Paperwallet is a simulated run of trading strategies",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				ctx          = context.Background()
				feedPair     = viper.GetString("pair")
				maker        = viper.GetFloat64("maker")
				taker        = viper.GetFloat64("taker")
				assetSymbol  = viper.GetString("symbol")
				assetAmount  = viper.GetFloat64("amount")
				strategyName = viper.GetString("strategy")
			)

			settings := ninjabot.Settings{
				Pairs: []string{feedPair},
			}

			strategy, err := strategies.NewStrategy(strategyName)
			if err != nil {
				return err
			}

			binance, err := exchange.NewBinance(ctx)
			if err != nil {
				return err
			}

			storage, err := storage.FromMemory()
			if err != nil {
				return err
			}

			wallet := exchange.NewPaperWallet(
				ctx,
				feedPair,
				exchange.WithPaperFee(maker, taker),
				exchange.WithPaperAsset(assetSymbol, assetAmount),
				exchange.WithDataFeed(binance),
			)

			bot, err := ninjabot.NewBot(
				ctx,
				settings,
				wallet,
				strategy,
				ninjabot.WithStorage(storage),
				ninjabot.WithPaperWallet(wallet),
			)
			if err != nil {
				return err
			}

			return bot.Run(ctx)
		},
	}
}
