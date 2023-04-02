package cmd

import (
	"context"

	"github.com/rodrigo-brito/ninjabot/download"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/rodrigo-brito/ninjabot/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func downloadCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "download",
		Short: "Download historical data for user backtesting",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				ctx       = context.Background()
				pair      = viper.GetString("pair")
				timeframe = viper.GetString("timeframe")
				file      = viper.GetString("file")
				start     = viper.GetTime("start")
				end       = viper.GetTime("end")
				futures   = viper.GetBool("futures")
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
			if days := viper.GetInt("days"); days > 0 {
				options = append(options, download.WithDays(days))
			}

			if !start.IsZero() && !end.IsZero() {
				options = append(options, download.WithInterval(start, end))
			}

			return download.NewDownloader(exc).Download(ctx, pair, timeframe, file, options...)
		},
	}
}
