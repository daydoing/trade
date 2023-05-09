package backtesting

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/rodrigo-brito/ninjabot/plot"
	"github.com/rodrigo-brito/ninjabot/storage"
	"github.com/spf13/cobra"

	"github.com/daydoing/trade/context"
	"github.com/daydoing/trade/strategies"
)

func BacktestingCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "backtesting",
		Short: "Backtesting is to backtest the trading strategy",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				pairs        = ctx.Config.System.Pairs
				baseCoin     = ctx.Config.System.BaseCoin
				amount       = ctx.Config.System.Amount
				feedPairs    = ctx.Config.Feed.Pairs
				path         = ctx.Config.Feed.Path
				timeframe    = ctx.Config.Feed.Timeframe
				maker        = ctx.Config.Binance.Maker
				taker        = ctx.Config.Binance.Taker
				strategyName = ctx.Config.Strategy.Name
			)

			settings := ninjabot.Settings{
				Pairs: pairs,
			}

			strategy, err := strategies.Strategy(strategyName, ctx)
			if err != nil {
				return err
			}

			pairFeed := []exchange.PairFeed{}
			for _, pair := range feedPairs {
				pairFeed = append(pairFeed, exchange.PairFeed{
					Pair:      pair,
					File:      filepath.Join(path, strings.Join([]string{pair, timeframe}, "_")) + ".csv",
					Timeframe: timeframe,
				})
			}

			csvFeed, err := exchange.NewCSVFeed(
				strategy.Timeframe(),
				pairFeed...,
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
				baseCoin,
				exchange.WithPaperFee(maker, taker),
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
				ctx,
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

			if err := bot.Run(ctx); err != nil {
				return err
			}

			bot.Summary()

			return chart.Start()
		},
	}
}
