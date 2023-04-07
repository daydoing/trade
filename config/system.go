package config

type System struct {
	BaseCoin string   `mapstructure:"baseCoin"`
	Amount   float64  `mapstructure:"amount"`
	MinQuote float64  `mapstructure:"minQuote"`
	Pairs    []string `mapstructure:"pairs"`
}
