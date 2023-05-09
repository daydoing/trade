package config

type Feed struct {
	Futures   bool     `mapstructure:"futures"`
	Timeframe string   `mapstructure:"timeframe"`
	Start     string   `mapstructure:"start"`
	End       string   `mapstructure:"end"`
	Path      string   `mapstructure:"path"`
	Pairs     []string `mapstructure:"pairs"`
}
