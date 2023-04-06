package strategies

import (
	"errors"

	"github.com/rodrigo-brito/ninjabot/strategy"
)

var strategies = map[string]strategy.Strategy{}

func RegisterStrategy(name string, strategy strategy.Strategy) {
	strategies[name] = strategy
}

func Strategy(name string) (strategy.Strategy, error) {
	if v, ok := strategies[name]; ok {
		return v, nil
	}

	return nil, errors.New("unsupported trading strategy")
}
