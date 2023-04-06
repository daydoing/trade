package config

import "time"

type Feed struct {
	Start time.Time
	End   time.Time
	Path  string
}
