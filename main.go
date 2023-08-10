package main

import (
	"log"

	"github.com/daydoing/trade/cmd/backtesting"
	"github.com/daydoing/trade/cmd/depth"
	"github.com/daydoing/trade/cmd/download"
	"github.com/daydoing/trade/cmd/futures"
	"github.com/daydoing/trade/cmd/monitor"
	"github.com/daydoing/trade/cmd/paperwallet"
	"github.com/daydoing/trade/cmd/spot"
	"github.com/daydoing/trade/context"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "traded",
		Short: "Traded is a quantitative trading terminal",
	}

	srv, err := context.NewBotContext()
	if err != nil {
		log.Panic(err)
	}

	cmd.AddCommand(backtesting.BacktestingCommand(srv))
	cmd.AddCommand(paperwallet.PaperwalletCommand(srv))
	cmd.AddCommand(download.DownloadCommand(srv))
	cmd.AddCommand(spot.SpotMarketCommand(srv))
	cmd.AddCommand(futures.FuturesMarketCommand(srv))
	cmd.AddCommand(depth.DepthCommand(srv))
	cmd.AddCommand(monitor.MonitorCommand(srv))

	if err := cmd.Execute(); err != nil {
		log.Panic(err)
	}
}
