package context

import (
	"strings"

	"github.com/daydoing/trade/config"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func initViper() (c *config.Config, err error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	viper.SetEnvPrefix("traded")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.BindEnv("system.base_coin")
	viper.BindEnv("system.amount")
	viper.BindEnv("system.min_quote")
	viper.BindEnv("system.pairs")

	viper.BindEnv("binance.key")
	viper.BindEnv("binance.secret")
	viper.BindEnv("binance.taker")
	viper.BindEnv("binance.maker")
	viper.BindEnv("binance.leverage")

	viper.BindEnv("telegram.uid")
	viper.BindEnv("telegram.token")

	viper.BindEnv("feed.futures")
	viper.BindEnv("feed.pair")
	viper.BindEnv("feed.start")
	viper.BindEnv("feed.end")
	viper.BindEnv("feed.path")

	viper.BindEnv("strategy.name")
	viper.BindEnv("strategy.timeframe")
	viper.BindEnv("strategy.period")

	viper.BindEnv("sentry.dsn")

	if err = viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return c, err
}
