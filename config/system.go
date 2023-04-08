package config

type System struct {
	BaseCoin string   `mapstructure:"base_coin"`
	Amount   float64  `mapstructure:"amount"`
	MinQuote float64  `mapstructure:"min_quote"`
	Pairs    []string `mapstructure:"pairs"`
}
