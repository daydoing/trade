package paperwallet

import (
	"context"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/rodrigo-brito/ninjabot/storage"
	"github.com/spf13/cobra"

	"github.com/daydoing/trade/service"
	"github.com/daydoing/trade/strategies"
)

func PaperwalletCommand(srv *service.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "paperwallet",
		Short: "Paperwallet is a simulated run of trading strategies",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				ctx          = context.Background()
				pairs        = srv.Config.System.Pairs
				maker        = srv.Config.Binance.Maker
				taker        = srv.Config.Binance.Taker
				baseCoin     = srv.Config.System.BaseCoin
				amount       = srv.Config.System.Amount
				strategyName = srv.Config.Strategy.Name
			)

			settings := ninjabot.Settings{
				Pairs: pairs,
			}

			strategy, err := strategies.Strategy(strategyName)
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
