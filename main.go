package main

import (
	"fmt"
	"os"

	"github.com/daydoing/trade/cmd/backtesting"
	"github.com/daydoing/trade/cmd/download"
	"github.com/daydoing/trade/cmd/futures"
	"github.com/daydoing/trade/cmd/paperwallet"
	"github.com/daydoing/trade/cmd/spot"
	"github.com/daydoing/trade/service"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "traded",
		Short: "Traded is a quantitative trading terminal",
	}

	srv, err := service.NewContext()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmd.AddCommand(backtesting.BacktestingCommand(srv))
	cmd.AddCommand(paperwallet.PaperwalletCommand(srv))
	cmd.AddCommand(download.DownloadCommand(srv))
	cmd.AddCommand(spot.SpotMarketCommand(srv))
	cmd.AddCommand(futures.FuturesMarketCommand(srv))

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
