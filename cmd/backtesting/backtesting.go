package backtesting

import (
	"log"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/rodrigo-brito/ninjabot/plot"
	"github.com/rodrigo-brito/ninjabot/storage"
	"github.com/spf13/cobra"

	"github.com/daydoing/trade/context"
	"github.com/daydoing/trade/strategies"
)

func BacktestingCommand(srv context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "backtesting",
		Short: "Backtesting is to backtest the trading strategy",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				pairs        = srv.Config.System.Pairs
				pair         = srv.Config.Feed.Pair
				file         = srv.Config.Feed.Path
				timeframe    = srv.Config.Strategy.Timeframe
				baseCoin     = srv.Config.System.BaseCoin
				amount       = srv.Config.System.Amount
				strategyName = srv.Config.Strategy.Name
			)

			settings := ninjabot.Settings{
				Pairs: pairs,
			}

			strategy, err := strategies.Strategy(strategyName, srv)
			if err != nil {
				return err
			}

			csvFeed, err := exchange.NewCSVFeed(
				strategy.Timeframe(),
				exchange.PairFeed{
					Pair:      pair,
					File:      file,
					Timeframe: timeframe,
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
				srv,
				baseCoin,
				exchange.WithPaperAsset(baseCoin, amount),
				exchange.WithDataFeed(csvFeed),
			)

			// create a chart  with indicators from the strategy and a custom additional RSI indicator
			chart, err := plot.NewChart(
				plot.WithStrategyIndicators(strategy),
				plot.WithPaperWallet(wallet),
			)
			if err != nil {
				log.Fatal(err)
			}

			bot, err := ninjabot.NewBot(
				srv,
				settings,
				wallet,
				strategy,
				ninjabot.WithBacktest(wallet),
				ninjabot.WithStorage(storage),
				ninjabot.WithCandleSubscription(chart),
				ninjabot.WithOrderSubscription(chart),
			)
			if err != nil {
				return err
			}

			if err := bot.Run(srv); err != nil {
				return err
			}

			bot.Summary()

			return chart.Start()
		},
	}
}
