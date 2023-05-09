package download

import (
	"path/filepath"
	"strings"
	"time"

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
				timeframe = ctx.Config.Feed.Timeframe
				pairs     = ctx.Config.Feed.Pairs
				path      = ctx.Config.Feed.Path
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

			startTime, err := time.Parse("2006-01-02 15:04:05", start)
			if err != nil {
				return err
			}

			endTime, err := time.Parse("2006-01-02 15:04:05", end)
			if err != nil {
				return err
			}

			var options []download.Option
			if !startTime.IsZero() && !endTime.IsZero() {
				options = append(options, download.WithInterval(startTime, endTime))
			}

			ec := make(chan error)

			for _, pair := range pairs {
				output := filepath.Join(path, strings.Join([]string{pair, timeframe}, "_")) + ".csv"
				go func(pair string) {
					ec <- download.NewDownloader(exc).Download(ctx, pair, timeframe, output, options...)
				}(pair)

				err := <-ec
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}
