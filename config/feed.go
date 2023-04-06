package config

import "time"

type Feed struct {
	Futures bool
	Pair    string
	Start   time.Time
	End     time.Time
	Path    string
}
