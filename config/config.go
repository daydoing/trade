package config

type Config struct {
	System   `mapstructure:"system"`
	Binance  `mapstructure:"binance"`
	Telegram `mapstructure:"telegram"`
	Feed     `mapstructure:"feed"`
	Strategy `mapstructure:"strategy"`
}