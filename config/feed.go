package config

import "time"

type Feed struct {
	Futures bool      `mapstructure:"futures"`
	Pair    string    `mapstructure:"pair"`
	Start   time.Time `mapstructure:"start"`
	End     time.Time `mapstructure:"end"`
	Path    string    `mapstructure:"path"`
}
