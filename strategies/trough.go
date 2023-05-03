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
	ctx                  context.Context
	gridNumber           int
	currentBuyGridNumber int
	step                 int
	quotePositionSize    float64
	assetPositionSize    float64
	stopLosePoint        float64
	takeProfitPoint      float64
	interruptExecution   bool
	trailingStop         *tools.TrailingStop
}

func NewTrough(srv context.Context) strategy.HighFrequencyStrategy {
	return &trough{
		ctx:          srv,
		step:         1,
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
			GroupName: "BB",
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
			GroupName: "BB",
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
			GroupName: "BB",
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

	if quotePosition > t.ctx.Config.MinQuote {
		if t.currentBuyGridNumber == 0 {
			t.quotePositionSize = math.Floor(quotePosition / float64(t.gridNumber))
			if t.quotePositionSize > t.ctx.Config.MinQuote {
				c1 := df.Low.Crossunder(df.Metadata["lb"])
				if c1 {
					_, err = broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.quotePositionSize)
					if err != nil {
						t.ctx.Logger.Error(err)
					}

					t.currentBuyGridNumber++
					t.stopLosePoint = df.Metadata["boll"].Last(0) - df.Metadata["atr"].Last(0)*float64(t.currentBuyGridNumber+t.step)
					t.takeProfitPoint = df.Metadata["boll"].Last(0) + df.Metadata["atr"].Last(0)*float64(t.currentBuyGridNumber+t.step)

					t.trailingStop.Start(df.Close.Last(0), t.stopLosePoint)
				}
			}
		} else {
			if df.Close.Last(0) <= t.stopLosePoint {
				if quotePosition >= t.quotePositionSize {
					_, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.quotePositionSize)
					if err != nil {
						t.ctx.Logger.Error(err)
					}

					t.currentBuyGridNumber++
					t.stopLosePoint = df.Metadata["boll"].Last(0) - df.Metadata["atr"].Last(0)*float64(t.currentBuyGridNumber+t.step)
					t.takeProfitPoint = df.Metadata["boll"].Last(0) + df.Metadata["atr"].Last(0)*float64(t.currentBuyGridNumber+t.step)

					diff := t.stopLosePoint - df.Close.Last(0)
					if diff < df.Metadata["atr"].Last(0) {
						t.stopLosePoint = t.stopLosePoint - df.Metadata["atr"].Last(0)
					}

					t.trailingStop.Start(df.Close.Last(0), t.stopLosePoint)
				} else {
					if assetPosition*df.Close.Last(0) > t.ctx.Config.MinQuote {
						if trailing := t.trailingStop; trailing != nil && trailing.Update(df.Close.Last(0)) {
							_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, assetPosition)
							if err != nil {
								t.ctx.Logger.Error(err)
							}

							t.ctx.Logger.Info("Important reminder: the market situation may reverse.")

							t.currentBuyGridNumber = 0.0
							t.interruptExecution = true
							t.trailingStop.Stop()
						}
					}
				}
			}
		}
	}

	if assetPosition*df.Close.Last(0) > t.ctx.Config.MinQuote {
		c1 := df.High.Crossover(df.Metadata["ub"])
		c2 := df.Close.Last(0) > t.takeProfitPoint
		if c1 || c2 {
			t.assetPositionSize = assetPosition / float64(t.currentBuyGridNumber)
			_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, t.assetPositionSize)
			if err != nil {
				t.ctx.Logger.Error(err)
			}

			t.currentBuyGridNumber = 0.0
			t.trailingStop.Stop()
		}
	}
}
