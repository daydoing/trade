package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	symbol  string
	pair    string
	amount  float64
	taker   float64
	maker   float64

	rootCmd = &cobra.Command{
		Use:   "traded",
		Short: "Traded is a quantitative trading terminal",
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default $HOME/.traded.yaml)")
	rootCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "USDT", "the token symbol")
	rootCmd.PersistentFlags().StringVarP(&pair, "pair", "p", "BTCUSDT", "trading pair")
	rootCmd.PersistentFlags().Float64VarP(&amount, "amount", "a", 1000, "the token amount")
	rootCmd.PersistentFlags().Float64VarP(&taker, "taker", "t", 0.001, "taker fee")
	rootCmd.PersistentFlags().Float64VarP(&maker, "maker", "m", 0.001, "maker fee")

	viper.BindPFlag("asset.symbol", rootCmd.PersistentFlags().Lookup("symbol"))
	viper.BindPFlag("asset.amount", rootCmd.PersistentFlags().Lookup("amount"))
	viper.BindPFlag("settings.pair", rootCmd.PersistentFlags().Lookup("pair"))
	viper.BindPFlag("settings.taker", rootCmd.PersistentFlags().Lookup("taker"))
	viper.BindPFlag("settings.maker", rootCmd.PersistentFlags().Lookup("maker"))

	rootCmd.AddCommand(backtestingCommand())
	rootCmd.AddCommand(paperwalletCommand())
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".traded")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

// Execute run root command
func Execute() error {
	return rootCmd.Execute()
}
