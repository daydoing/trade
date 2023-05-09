package paperwallet

import (
	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/rodrigo-brito/ninjabot/storage"
	"github.com/spf13/cobra"

	"github.com/daydoing/trade/context"
	"github.com/daydoing/trade/strategies"
)

func PaperwalletCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "paperwallet",
		Short: "Paperwallet is a simulated run of trading strategies",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				pairs        = ctx.Config.System.Pairs
				maker        = ctx.Config.Binance.Maker
				taker        = ctx.Config.Binance.Taker
				baseCoin     = ctx.Config.System.BaseCoin
				amount       = ctx.Config.System.Amount
				strategyName = ctx.Config.Strategy.Name
			)

			settings := ninjabot.Settings{
				Pairs: pairs,
			}

			strategy, err := strategies.Strategy(strategyName, ctx)
			if err != nil {
				return err
			}

			binance, err := exchange.NewBinance(ctx, exchange.WithTestNet())
			if err != nil {
				return err
			}

			storage, err := storage.FromMemory()
			if err != nil {
				return err
			}

			wallet := exchange.NewPaperWallet(
				ctx,
				baseCoin,
				exchange.WithPaperFee(maker, taker),
				exchange.WithPaperAsset(baseCoin, amount),
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
