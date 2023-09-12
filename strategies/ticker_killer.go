package strategies

import (
	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/indicator"
	"github.com/rodrigo-brito/ninjabot/model"
	"github.com/rodrigo-brito/ninjabot/service"
	"github.com/rodrigo-brito/ninjabot/strategy"

	"github.com/daydoing/trade/context"
)

const (
	crossunderLowerband = "crossunderLowerband"
	crossunder3atr      = "crossunder3atr"
	rsi30               = "rsi30"
)

type (
	tickerKiller struct {
		ctx           context.BotContext
		quote         float64
		rsi           float64
		atr           float64
		upperband     float64
		middleband    float64
		lowerband     float64
		stopLosePoint float64
		buyConditions []condition
	}

	condition struct {
		name     string
		weight   float64
		order    *model.Order
		callback func(df *ninjabot.Dataframe, broker service.Broker, quote float64) (model.Order, error)
	}
)

func NewTickerKiller(srv context.BotContext) strategy.HighFrequencyStrategy {
	tk := &tickerKiller{ctx: srv}

	c1 := condition{name: crossunderLowerband, weight: 0.2}
	c1.callback = func(df *ninjabot.Dataframe, broker service.Broker, quote float64) (model.Order, error) {
		size := quote * c1.weight / tk.lowerband
		if tk.lowerband*size >= tk.ctx.Config.MinQuote {
			order, err := broker.CreateOrderLimit(ninjabot.SideTypeBuy, df.Pair, size, tk.lowerband)
			if err != nil {
				return model.Order{}, err
			}

			return order, nil
		}

		return model.Order{}, nil
	}

	c2 := condition{name: crossunder3atr, weight: 0.3}
	c2.callback = func(df *ninjabot.Dataframe, broker service.Broker, quote float64) (model.Order, error) {
		size := quote * c2.weight / (tk.middleband - tk.atr*3)
		if quote*c2.weight >= tk.ctx.Config.MinQuote {
			order, err := broker.CreateOrderLimit(ninjabot.SideTypeBuy, df.Pair, size, tk.middleband-tk.atr*3)
			if err != nil {
				return model.Order{}, err
			}

			return order, nil
		}

		return model.Order{}, nil

	}

	c3 := condition{name: rsi30, weight: 0.5}
	c3.callback = func(df *ninjabot.Dataframe, broker service.Broker, quote float64) (model.Order, error) {
		if quote*c3.weight >= tk.ctx.Config.MinQuote {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, quote*c3.weight-10)
			if err != nil {
				return model.Order{}, err
			}

			return order, nil
		}
		return model.Order{}, nil
	}

	tk.buyConditions = []condition{c1, c2, c3}

	return tk
}

func (t *tickerKiller) Timeframe() string {
	return t.ctx.Config.Strategy.Timeframe
}

func (t *tickerKiller) WarmupPeriod() int {
	return t.ctx.Config.Strategy.Period
}

func (t *tickerKiller) Indicators(df *ninjabot.Dataframe) []strategy.ChartIndicator {
	df.Metadata["rsi"] = indicator.RSI(df.Close, 14)
	df.Metadata["atr"] = indicator.ATR(df.High, df.Low, df.Close, bbPeriod/2)
	df.Metadata["upperband"], df.Metadata["middleband"], df.Metadata["lowerband"] = indicator.BB(df.Close, bbPeriod, deviation, indicator.TypeEMA)

	t.rsi = df.Metadata["rsi"].Last(0)
	t.upperband = df.Metadata["upperband"].Last(0)
	t.middleband = df.Metadata["middleband"].Last(0)
	t.lowerband = df.Metadata["lowerband"].Last(0)
	t.atr = df.Metadata["atr"].Last(0)

	return []strategy.ChartIndicator{}
}

// source data's timeframe must less then t.timeframe, otherwise it's will be panic for HighFrequencyStrategy
func (t *tickerKiller) OnCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execLongStrategy(df, broker)
}

func (t *tickerKiller) OnPartialCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execLongStrategy(df, broker)
}

func (t *tickerKiller) execLongStrategy(df *ninjabot.Dataframe, broker service.Broker) {
	asset, quote, err := broker.Position(df.Pair)
	if err != nil {
		t.ctx.Logger.Error(err)
		return
	}

	if t.quote == 0 {
		t.quote = quote
	}

	for k, v := range t.buyConditions {
		if v.order != nil {
			continue
		}

		order, err := v.callback(df, broker, t.quote)
		if err != nil {
			t.ctx.Logger.Error(err)
			return
		}

		t.buyConditions[k].order = &order
	}

	price := df.Close.Last(0)
	if asset*price > t.ctx.Config.MinQuote {
		c1 := df.High.Crossover(df.Metadata["upperband"])
		c2 := t.rsi >= 70.0

		if c1 || c2 {
			_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, asset)
			if err != nil {
				t.ctx.Logger.Error(err)
				return
			}

			for k := range t.buyConditions {
				order := t.buyConditions[k].order
				if order != nil {
					err := broker.Cancel(*order)
					if err != nil {
						t.ctx.Logger.Error(err)
						return
					}
					t.buyConditions[k].order = nil
				}
			}

			t.quote = 0.0
		}
	}
}
