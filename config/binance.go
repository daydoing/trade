package config

type Binance struct {
	Key      string  `mapstructure:"key"`
	Secret   string  `mapstructure:"secret"`
	Taker    float64 `mapstructure:"taker"`
	Maker    float64 `mapstructure:"maker"`
	Leverage int     `mapstructure:"leverage"`
}
