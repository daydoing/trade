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
	bbPeriod   = 20
	deviation  = 3
	gridNumber = 3
)

type trough struct {
	ctx                context.Context
	currentGrid        int
	gridNumber         int
	gridQuantity       float64
	stopLosePoint      float64
	takeProfitPoint    float64
	interruptExecution bool
	trailingStop       *tools.TrailingStop
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
	df.Metadata["atr"] = indicator.ATR(df.High, df.Low, df.Close, bbPeriod/2)
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
	t.execLongStrategy(df, broker)
}

func (t *trough) OnPartialCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execLongStrategy(df, broker)
}

func (t *trough) execLongStrategy(df *ninjabot.Dataframe, broker service.Broker) {
	c1 := df.High.Crossover(df.Metadata["ub"])
	if c1 {
		t.interruptExecution = false
	}

	if t.interruptExecution {
		return
	}

	assetPosition, quotePosition, err := broker.Position(df.Pair)
	if err != nil {
		t.ctx.Logger.Error(err)
	}

	step := 1
	if t.currentGrid == 0 {
		if quotePosition > t.ctx.Config.MinQuote {
			t.gridQuantity = math.Floor(quotePosition / float64(t.gridNumber))

			c1 := df.Low.Crossunder(df.Metadata["lb"])
			if c1 {
				_, err = broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
				if err != nil {
					t.ctx.Logger.Error(err)
				}

				t.currentGrid++
				t.stopLosePoint = df.Metadata["boll"].Last(0) - df.Metadata["atr"].Last(0)*float64(t.currentGrid+step)
				t.takeProfitPoint = df.Metadata["boll"].Last(0) + df.Metadata["atr"].Last(0)*float64(t.currentGrid+step)
				t.trailingStop.Start(df.Low.Last(0), t.stopLosePoint)
			}
		}
	} else {
		if quotePosition >= t.gridQuantity && df.Close.Last(0) <= t.stopLosePoint {
			_, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
			if err != nil {
				t.ctx.Logger.Error(err)
			}

			t.currentGrid++
			t.stopLosePoint = df.Metadata["boll"].Last(0) - df.Metadata["atr"].Last(0)*float64(t.currentGrid+step)
			t.takeProfitPoint = df.Metadata["boll"].Last(0) + df.Metadata["atr"].Last(0)*float64(t.currentGrid+step)

			diff := t.stopLosePoint - df.Close.Last(0)
			if diff < df.Metadata["atr"].Last(0) {
				t.stopLosePoint = t.stopLosePoint - df.Metadata["atr"].Last(0)
			}

			t.trailingStop.Start(df.Low.Last(0), t.stopLosePoint)
		}
	}

	if assetPosition*df.Close.Last(0) > t.ctx.Config.MinQuote {
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

		if df.Close.Last(0) < t.stopLosePoint && quotePosition < t.gridQuantity {
			if trailing := t.trailingStop; trailing != nil && trailing.Update(df.Close.Last(0)) {
				_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, assetPosition)
				if err != nil {
					t.ctx.Logger.Error(err)
				}

				t.ctx.Logger.Info("Important reminder: the market situation may reverse.")

				t.currentGrid = 0.0
				t.trailingStop.Stop()
				t.interruptExecution = true
			}
		}
	}
}
