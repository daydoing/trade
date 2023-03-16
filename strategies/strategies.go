package strategies

import (
	"errors"

	"github.com/rodrigo-brito/ninjabot/strategy"
)

const (
	EMA = "ema"
)

func NewStrategy(strategy string) (strategy.Strategy, error) {
	switch strategy {
	case EMA:
		return new(CrossEMA), nil
	default:
		return nil, errors.New("unsupported trading strategy")
	}
}
