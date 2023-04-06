package config

type System struct {
	BaseCoin string   `mapstructure:"baseCoin"`
	Amount   float64  `mapstructure:"amount"`
	Pairs    []string `mapstructure:"pairs"`
}
