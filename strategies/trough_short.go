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
	sbbPeriod   = 20
	sdeviation  = 3
	sgridNumber = 3
)

type troughShort struct {
	ctx                context.Context
	period             int
	currentGrid        int
	gridNumber         float64
	gridQuantity       float64
	stopLosePoint      float64
	takeProfitPoint    float64
	timeframe          string
	interruptExecution bool
	trailingStop       *tools.TrailingStop
}

func NewTroughShort(srv context.Context) strategy.HighFrequencyStrategy {
	return &troughShort{
		ctx:          srv,
		timeframe:    srv.Config.Strategy.Timeframe,
		period:       srv.Config.Strategy.Period,
		gridNumber:   sgridNumber,
		trailingStop: tools.NewTrailingStop(),
	}
}

func (t *troughShort) Timeframe() string {
	return t.timeframe
}

func (t *troughShort) WarmupPeriod() int {
	return t.period
}

func (t *troughShort) Indicators(df *ninjabot.Dataframe) []strategy.ChartIndicator {
	df.Metadata["atr"] = indicator.ATR(df.High, df.Low, df.Close, bbPeriod/2)
	df.Metadata["ub"], df.Metadata["boll"], df.Metadata["lb"] = indicator.BB(df.Close, sbbPeriod, sdeviation, indicator.TypeEMA)

	return []strategy.ChartIndicator{
		{
			Overlay:   true,
			GroupName: "UB",
			Time:      df.Time,
			Warmup:    t.period,
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
			Warmup:    t.period,
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
			Warmup:    t.period,
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
func (t *troughShort) OnCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execShortStrategy(df, broker)
}

func (t *troughShort) OnPartialCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execShortStrategy(df, broker)
}

func (t *troughShort) execShortStrategy(df *ninjabot.Dataframe, broker service.Broker) {
	c1 := df.Low.Crossunder(df.Metadata["lb"])
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

	absAssetPosition := math.Abs(assetPosition)

	step := 1
	if t.currentGrid == 0 {
		if quotePosition > t.ctx.Config.MinQuote {
			t.gridQuantity = math.Floor(quotePosition / t.gridNumber)

			c1 := df.High.Crossover(df.Metadata["ub"])
			if c1 {
				_, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeSell, df.Pair, t.gridQuantity)
				if err != nil {
					t.ctx.Logger.Error(err)
				}

				t.currentGrid++
				t.stopLosePoint = df.Metadata["boll"].Last(0) + df.Metadata["atr"].Last(0)*float64(t.currentGrid+step)
				t.takeProfitPoint = df.Metadata["boll"].Last(0) - df.Metadata["atr"].Last(0)*float64(t.currentGrid+step)
				t.trailingStop.Start(df.Low.Last(0), t.stopLosePoint)
			}
		}
	} else {
		if quotePosition >= t.gridQuantity && df.Close.Last(0) >= t.stopLosePoint {
			_, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeSell, df.Pair, t.gridQuantity)
			if err != nil {
				t.ctx.Logger.Error(err)
			}

			t.currentGrid++
			t.stopLosePoint = df.Metadata["boll"].Last(0) + df.Metadata["atr"].Last(0)*float64(t.currentGrid+step)
			t.takeProfitPoint = df.Metadata["boll"].Last(0) - df.Metadata["atr"].Last(0)*float64(t.currentGrid+step)
			t.trailingStop.Start(df.Low.Last(0), t.stopLosePoint)
		}
	}

	if absAssetPosition > t.ctx.Config.MinQuote {
		c1 := df.Low.Crossunder(df.Metadata["lb"])
		c2 := df.Close.Last(0) <= t.takeProfitPoint
		if c1 || c2 {
			_, err := broker.CreateOrderMarket(ninjabot.SideTypeBuy, df.Pair, absAssetPosition)
			if err != nil {
				t.ctx.Logger.Error(err)
			}

			t.currentGrid = 0.0
			t.trailingStop.Stop()
		}

		if df.High.Last(0) > t.stopLosePoint && quotePosition < t.gridQuantity {
			if trailing := t.trailingStop; trailing != nil && trailing.Update(df.High.Last(0)) {
				_, err := broker.CreateOrderMarket(ninjabot.SideTypeBuy, df.Pair, absAssetPosition)
				if err != nil {
					t.ctx.Logger.Error(err)
				}

				t.currentGrid = 0.0
				t.trailingStop.Stop()
				t.interruptExecution = true
			}
		}
	}
}
