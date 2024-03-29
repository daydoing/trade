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
	maxRSI     = 70.0
)

type trough struct {
	ctx                  context.BotContext
	gridNumber           int
	currentBuyGridNumber int
	step                 int
	quotePositionSize    float64
	stopLosePoint        float64
	trailingStop         *tools.TrailingStop
}

func NewTrough(srv context.BotContext) strategy.HighFrequencyStrategy {
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
	df.Metadata["rsi"] = indicator.RSI(df.Close, 14)
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
	assetPosition, quotePosition, err := broker.Position(df.Pair)
	if err != nil {
		t.ctx.Logger.Error(err)
		return
	}

	var (
		clossPrice = df.Close.Last(0)
		boll       = df.Metadata["boll"].Last(0)
		atr        = df.Metadata["atr"].Last(0)
	)

	if t.currentBuyGridNumber == 0 {
		t.quotePositionSize = math.Floor(quotePosition / float64(t.gridNumber))
		if t.quotePositionSize > t.ctx.Config.MinQuote {
			c1 := df.Low.Crossunder(df.Metadata["lb"])
			c2 := df.High.Crossover(df.Metadata["ub"])
			if c1 && !c2 {
				_, err = broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.quotePositionSize)
				if err != nil {
					t.ctx.Logger.Error(err)
					return
				}

				t.currentBuyGridNumber++
				t.stopLosePoint = boll - atr*float64(t.currentBuyGridNumber+t.step)

				t.trailingStop.Start(df.Close.Last(0), t.stopLosePoint)
			}
		}
	} else {
		if clossPrice <= t.stopLosePoint {
			if quotePosition >= t.quotePositionSize {
				_, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.quotePositionSize)
				if err != nil {
					t.ctx.Logger.Error(err)
					return
				}

				t.currentBuyGridNumber++
				t.stopLosePoint = boll - atr*float64(t.currentBuyGridNumber+t.step)

				diff := t.stopLosePoint - clossPrice
				if diff < atr {
					t.stopLosePoint = t.stopLosePoint - atr
				}

				t.trailingStop.Start(clossPrice, t.stopLosePoint)
			} else {
				if assetPosition*clossPrice > t.ctx.Config.MinQuote {
					if trailing := t.trailingStop; trailing != nil && trailing.Update(clossPrice) {
						_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, assetPosition)
						if err != nil {
							t.ctx.Logger.Error(err)
							return
						}

						t.ctx.NotifyBot.Notify("Important reminder: the market situation may reverse.")

						t.gridNumber++
						t.currentBuyGridNumber = 0.0
						t.trailingStop.Stop()
					}
				}
			}
		}
	}

	if assetPosition*clossPrice > t.ctx.Config.MinQuote {
		c1 := df.High.Crossover(df.Metadata["ub"])
		c2 := df.Metadata["rsi"].Last(0) >= maxRSI
		if c1 || c2 {
			_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, assetPosition)
			if err != nil {
				t.ctx.Logger.Error(err)
				return
			}

			t.gridNumber = 3
			t.currentBuyGridNumber = 0.0
			t.trailingStop.Stop()
		}
	}
}
