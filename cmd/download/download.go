package download

import (
	"github.com/rodrigo-brito/ninjabot/download"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/rodrigo-brito/ninjabot/service"
	"github.com/spf13/cobra"

	"github.com/daydoing/trade/context"
)

func DownloadCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "download",
		Short: "Download historical data for user backtesting",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				futures   = ctx.Config.Feed.Futures
				start     = ctx.Config.Feed.Start
				end       = ctx.Config.Feed.End
				timeframe = ctx.Config.Strategy.Timeframe
				pair      = ctx.Config.Feed.Pair
				output    = ctx.Config.Feed.Path
			)

			var (
				exc service.Feeder
				err error
			)

			if futures {
				// fetch data from binance futures
				exc, err = exchange.NewBinanceFuture(ctx)
				if err != nil {
					return err
				}
			} else {
				// fetch data from binance spot
				exc, err = exchange.NewBinance(ctx)
				if err != nil {
					return err
				}
			}

			var options []download.Option
			if !start.IsZero() && !end.IsZero() {
				options = append(options, download.WithInterval(start, end))
			}

			return download.NewDownloader(exc).Download(ctx, pair, timeframe, output, options...)
		},
	}
}
