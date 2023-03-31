package strategies

import (
	"errors"

	"github.com/rodrigo-brito/ninjabot/strategy"
)

const (
	EMA    = "ema"
	TROUGH = "trough"
)

func NewStrategy(strategy string) (strategy.Strategy, error) {
	switch strategy {
	case EMA:
		return NewCrossEMA(), nil
	case TROUGH:
		return NewTrough("5m", 9, 2.0, 1.5, 3.0), nil
	default:
		return nil, errors.New("unsupported trading strategy")
	}
}
