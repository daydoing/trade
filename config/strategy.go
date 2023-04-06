package config

type Strategy struct {
	Name      string `mapstructure:"name"`
	Timeframe string `mapstructure:"timeframe"`
	Period    int    `mapstructure:"period"`
}
