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
	ctx                   context.BotContext
	period                int
	gridNumber            int
	currentSellGridNumber int
	step                  int
	quotePositionSize     float64
	stopLosePoint         float64
	timeframe             string
	trailingStop          *tools.TrailingStop
}

func NewTroughShort(srv context.BotContext) strategy.HighFrequencyStrategy {
	return &troughShort{
		ctx:          srv,
		step:         1,
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
	assetPosition, quotePosition, err := broker.Position(df.Pair)
	if err != nil {
		t.ctx.Logger.Error(err)
		return
	}

	var (
		leverage   = t.ctx.Config.Leverage
		highPrice  = df.High.Last(0)
		clossPrice = df.Close.Last(0)
		boll       = df.Metadata["boll"].Last(0)
		atr        = df.Metadata["atr"].Last(0)
	)

	absAssetPosition := math.Abs(assetPosition)
	quotePosition = quotePosition*float64(leverage) - absAssetPosition*clossPrice

	if t.currentSellGridNumber == 0 {
		t.quotePositionSize = math.Floor(quotePosition / float64(t.gridNumber))
		if t.quotePositionSize > t.ctx.Config.MinQuote {
			c1 := df.High.Crossover(df.Metadata["ub"])
			c2 := df.Low.Crossunder(df.Metadata["lb"])
			if c1 && !c2 {
				_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, t.quotePositionSize/clossPrice)
				if err != nil {
					t.ctx.Logger.Error(err)
					return
				}

				t.currentSellGridNumber++
				t.stopLosePoint = boll + atr*float64(t.currentSellGridNumber+t.step)
				t.trailingStop.Start(df.Low.Last(0), t.stopLosePoint)
			}
		}
	} else {
		if clossPrice >= t.stopLosePoint {
			if quotePosition >= t.quotePositionSize {
				_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, t.quotePositionSize/clossPrice)
				if err != nil {
					t.ctx.Logger.Error(err)
					return
				}

				t.currentSellGridNumber++
				t.stopLosePoint = boll + atr*float64(t.currentSellGridNumber+t.step)

				diff := clossPrice - t.stopLosePoint
				if diff < atr {
					t.stopLosePoint = t.stopLosePoint + atr
				}

				t.trailingStop.Start(df.Low.Last(0), t.stopLosePoint)
			} else {
				if absAssetPosition*clossPrice > t.ctx.Config.MinQuote {
					if trailing := t.trailingStop; trailing != nil && trailing.Update(highPrice) {
						_, err := broker.CreateOrderMarket(ninjabot.SideTypeBuy, df.Pair, absAssetPosition)
						if err != nil {
							t.ctx.Logger.Error(err)
							return
						}

						t.ctx.NotifyBot.Notify("Important reminder: the market situation may reverse.")

						t.gridNumber++
						t.currentSellGridNumber = 0.0
						t.trailingStop.Stop()
					}
				}
			}
		}
	}

	if absAssetPosition*clossPrice > t.ctx.Config.MinQuote {
		c1 := df.Low.Crossunder(df.Metadata["lb"])
		if c1 {
			_, err := broker.CreateOrderMarket(ninjabot.SideTypeBuy, df.Pair, absAssetPosition)
			if err != nil {
				t.ctx.Logger.Error(err)
				return
			}

			t.gridNumber = 3
			t.currentSellGridNumber = 0.0
			t.trailingStop.Stop()
		}
	}
}
