package strategies

import (
	"errors"

	"github.com/rodrigo-brito/ninjabot/strategy"

	"github.com/daydoing/trade/context"
)

const (
	EMA          = "ema"
	TROUGH       = "trough"
	TROUGH_SHORT = "trough_short"
)

func Strategy(name string, srv context.Context) (strategy.Strategy, error) {
	switch name {
	case EMA:
		return NewCrossEMA(srv), nil
	case TROUGH:
		return NewTrough(srv), nil
	case TROUGH_SHORT:
		return NewTroughShort(srv), nil
	default:
		return nil, errors.New("unsupported trading strategy")
	}
}
