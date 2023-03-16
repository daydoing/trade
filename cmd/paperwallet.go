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

var (
	pc = &cobra.Command{
		Use:   "paperwallet",
		Short: "Paperwallet is a simulated run of trading strategies",
	}
)

func paperwalletCommand() *cobra.Command {
	pc.RunE = func(cmd *cobra.Command, args []string) error {
		var (
			ctx         = context.Background()
			feedPair    = viper.GetString("settings.pair")
			maker       = viper.GetFloat64("settings.maker")
			taker       = viper.GetFloat64("settings.taker")
			assetSymbol = viper.GetString("asset.symbol")
			assetAmount = viper.GetFloat64("asset.amount")
		)

		settings := ninjabot.Settings{
			Pairs: []string{feedPair},
		}

		strategy := strategies.NewStrategy()

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
	}

	return pc
}
