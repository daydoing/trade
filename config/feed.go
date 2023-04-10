package config

type Feed struct {
	Futures   bool   `mapstructure:"futures"`
	Pair      string `mapstructure:"pair"`
	Timeframe string `mapstructure:"timeframe"`
	Start     string `mapstructure:"start"`
	End       string `mapstructure:"end"`
	Path      string `mapstructure:"path"`
}
