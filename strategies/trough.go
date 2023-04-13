package strategies

import (
	"math"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/indicator"
	"github.com/rodrigo-brito/ninjabot/service"
	"github.com/rodrigo-brito/ninjabot/strategy"
	"github.com/rodrigo-brito/ninjabot/tools"

	"github.com/daydoing/trade/context"
)

const (
	bbPeriod    = 20
	deviation   = 3
	gridNumber  = 3.0
	downPercent = 2.5
	upPercent   = 3.5
)

type trough struct {
	ctx             context.Context
	currentGrid     int
	gridNumber      float64
	gridQuantity    float64
	stopLosePoint   float64
	takeProfitPoint float64
	trailingStop    *tools.TrailingStop
}

func NewTrough(srv context.Context) strategy.HighFrequencyStrategy {
	return &trough{
		ctx:          srv,
		gridNumber:   gridNumber,
		trailingStop: tools.NewTrailingStop(),
	}
}

func (t *trough) Timeframe() string {
	return t.ctx.Config.Strategy.Timeframe
}

func (t *trough) WarmupPeriod() int {
	return t.ctx.Config.Strategy.Period
}

func (t *trough) Indicators(df *ninjabot.Dataframe) []strategy.ChartIndicator {
	df.Metadata["atr"] = indicator.ATR(df.High, df.Low, df.Close, bbPeriod)
	df.Metadata["rsi"] = indicator.RSI(df.Close, bbPeriod)
	df.Metadata["ub"], df.Metadata["boll"], df.Metadata["lb"] = indicator.BB(df.Close, bbPeriod, deviation, indicator.TypeEMA)

	return []strategy.ChartIndicator{
		{
			Overlay:   true,
			GroupName: "UB",
			Time:      df.Time,
			Warmup:    t.ctx.Config.Strategy.Period,
			Metrics: []strategy.IndicatorMetric{
				{
					Values: df.Metadata["ub"],
					Name:   "UB",
					Color:  "blue",
					Style:  strategy.StyleLine,
				},
			},
		},
		{
			Overlay:   true,
			GroupName: "BOLL",
			Time:      df.Time,
			Warmup:    t.ctx.Config.Strategy.Period,
			Metrics: []strategy.IndicatorMetric{
				{
					Values: df.Metadata["boll"],
					Name:   "BOLL",
					Color:  "red",
					Style:  strategy.StyleLine,
				},
			},
		},
		{
			Overlay:   true,
			GroupName: "LB",
			Time:      df.Time,
			Warmup:    t.ctx.Config.Strategy.Period,
			Metrics: []strategy.IndicatorMetric{
				{
					Values: df.Metadata["lb"],
					Name:   "LB",
					Color:  "blue",
					Style:  strategy.StyleLine,
				},
			},
		},
	}
}

// source data's timeframe must less then t.timeframe, otherwise it's will be panic for HighFrequencyStrategy
func (t *trough) OnCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execStrategy(df, broker)
}

func (t *trough) OnPartialCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execStrategy(df, broker)
}

func (t *trough) execStrategy(df *ninjabot.Dataframe, broker service.Broker) {
	assetPosition, quotePosition, err := broker.Position(df.Pair)
	if err != nil {
		t.ctx.Logger.Error(err)
	}

	if t.currentGrid == 0 {
		if quotePosition > t.ctx.Config.MinQuote {
			t.gridQuantity = math.Floor(quotePosition / t.gridNumber)

			c1 := df.Low.Crossunder(df.Metadata["lb"])
			if c1 {
				_, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
				if err != nil {
					t.ctx.Logger.Error(err)
				}

				t.stopLosePoint = df.Low.Last(0) - df.Metadata["atr"].Last(0)*downPercent
				t.takeProfitPoint = df.High.Last(0) + df.Metadata["atr"].Last(0)*upPercent
				t.currentGrid++

				t.trailingStop.Start(df.Low.Last(0), t.stopLosePoint)
			}
		}
	} else {
		if quotePosition >= t.gridQuantity && df.Close.Last(0) <= t.stopLosePoint {
			_, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
			if err != nil {
				t.ctx.Logger.Error(err)
			}

			t.stopLosePoint = df.Low.Last(0) - df.Metadata["atr"].Last(0)*downPercent
			t.takeProfitPoint = df.High.Last(0) + df.Metadata["atr"].Last(0)*upPercent
			t.currentGrid++

			t.trailingStop.Start(df.Low.Last(0), t.stopLosePoint)
		}
	}

	if assetPosition > t.ctx.Config.MinQuote {
		c1 := df.High.Crossover(df.Metadata["ub"])
		c2 := df.High.Last(0) >= t.takeProfitPoint
		if c1 || c2 {
			_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, assetPosition)
			if err != nil {
				t.ctx.Logger.Error(err)
			}

			t.currentGrid = 0.0
			t.trailingStop.Stop()
		}

		if df.Low.Last(0) < t.stopLosePoint && quotePosition < t.gridQuantity {
			if trailing := t.trailingStop; trailing != nil && trailing.Update(df.Low.Last(0)) {
				_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, assetPosition)
				if err != nil {
					t.ctx.Logger.Error(err)
				}

				t.currentGrid = 0.0
				t.trailingStop.Stop()
			}
		}
	}
}
