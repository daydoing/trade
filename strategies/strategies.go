package strategies

import (
	"errors"

	"github.com/rodrigo-brito/ninjabot/strategy"

	"github.com/daydoing/trade/service"
)

const (
	EMA    = "ema"
	TROUGH = "trough"
)

func Strategy(name string, srv *service.Context) (strategy.Strategy, error) {
	switch name {
	case EMA:
		return NewCrossEMA(srv), nil
	case TROUGH:
		return NewTrough(srv), nil
	default:
		return nil, errors.New("unsupported trading strategy")
	}
}
