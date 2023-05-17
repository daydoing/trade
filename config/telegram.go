package config

type Telegram struct {
	UID   []int  `mapstructure:"uid"`
	Token string `mapstructure:"token"`
}
