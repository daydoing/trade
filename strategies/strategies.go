package strategies

import (
	"errors"

	"github.com/rodrigo-brito/ninjabot/strategy"

	"github.com/daydoing/trade/context"
	"github.com/daydoing/trade/strategies/arbitrage"
)

const (
	EMA          = "ema"
	ARBITRAGE    = "arbitrage"
	TROUGH       = "trough"
	TROUGH_SHORT = "trough_short"
)

func Strategy(name string, srv context.BotContext) (strategy.Strategy, error) {
	switch name {
	case EMA:
		return NewCrossEMA(srv), nil
	case TROUGH:
		return NewTrough(srv), nil
	case TROUGH_SHORT:
		return NewTroughShort(srv), nil
	case ARBITRAGE:
		return arbitrage.NewArbitrage(srv)
	default:
		return nil, errors.New("unsupported trading strategy")
	}
}
