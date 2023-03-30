package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "traded",
		Short: "Traded is a quantitative trading terminal",
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default $HOME/.traded.yaml)")

	rootCmd.AddCommand(backtestingCommand())
	rootCmd.AddCommand(paperwalletCommand())
	rootCmd.AddCommand(downloadCommand())
	rootCmd.AddCommand(walletCommand())
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
