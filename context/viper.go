package context

import (
	"os"

	"github.com/daydoing/trade/config"
	"github.com/spf13/viper"
)

func initViper() (c *config.Config, err error) {
	home, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName("traded")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// unmarshal config to struct
	if err = viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return c, err
}
