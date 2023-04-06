package cmd

import (
	"github.com/daydoing/trade/service"
	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "traded",
		Short: "Traded is a quantitative trading terminal",
	}
)

// Execute run root command
func Execute() error {
	srv, err := service.NewServiceContext(cfgFile)
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default ./traded.yaml)")

	rootCmd.AddCommand(backtestingCommand(srv))
	rootCmd.AddCommand(paperwalletCommand(srv))
	rootCmd.AddCommand(downloadCommand(srv))
	rootCmd.AddCommand(spotMarketCommand(srv))
	rootCmd.AddCommand(futuresMarketCommand(srv))

	return rootCmd.Execute()
}
