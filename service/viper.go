package service

import (
	"os"

	"github.com/daydoing/trade/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initViper(cfgFile string) (c *config.Config, err error) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("traded")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// unmarshal config to struct
	if err = viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return c, err
}
