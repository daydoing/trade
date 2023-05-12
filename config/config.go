package config

type Config struct {
	System    `mapstructure:"system"`
	Binance   `mapstructure:"binance"`
	Telegram  `mapstructure:"telegram"`
	Feed      `mapstructure:"feed"`
	Strategy  `mapstructure:"strategy"`
	Sentry    `mapstructure:"sentry"`
	Arbitrage `mapstructure:"arbitrage"`
	Monitor   `mapstructure:"monitor"`
}
