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
		return NewTrough("15m", 9, 5.0, 2.5, 5.0), nil
	default:
		return nil, errors.New("unsupported trading strategy")
	}
}
