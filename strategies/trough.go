package strategies

import (
	"math"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/indicator"
	"github.com/rodrigo-brito/ninjabot/model"
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
	ctx                 context.Context
	period              int
	currentGrid         int
	gridNumber          float64
	gridQuantity        float64
	totalCost           float64
	totalQuantity       float64
	averagePurchaseCost float64
	stopLosePoint       float64
	takeProfitPoint     float64
	timeframe           string
	order               model.Order
	trailingStop        *tools.TrailingStop
}

func NewTrough(srv context.Context) strategy.HighFrequencyStrategy {
	return &trough{
		ctx:          srv,
		timeframe:    srv.Config.Strategy.Timeframe,
		period:       srv.Config.Strategy.Period,
		gridNumber:   gridNumber,
		trailingStop: tools.NewTrailingStop(),
	}
}

func (t *trough) Timeframe() string {
	return t.timeframe
}

func (t *trough) WarmupPeriod() int {
	return t.period
}

func (t *trough) Indicators(df *ninjabot.Dataframe) []strategy.ChartIndicator {
	df.Metadata["atr"] = indicator.ATR(df.High, df.Low, df.Close, bbPeriod)
	df.Metadata["ub"], df.Metadata["boll"], df.Metadata["lb"] = indicator.BB(df.Close, bbPeriod, deviation, indicator.TypeEMA)

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
					Color:  "orange",
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
func (t *trough) OnCandle(df *ninjabot.Dataframe, broker service.Broker) {}

func (t *trough) OnPartialCandle(df *ninjabot.Dataframe, broker service.Broker) {
	assetPosition, quotePosition, err := broker.Position(df.Pair)
	if err != nil {
		t.ctx.Logger.Error(err)
	}

	var (
		minQuote   = t.ctx.Config.MinQuote
		closePrice = df.Close.Last(0)
	)

	if t.currentGrid == 0 {
		if quotePosition > minQuote {
			t.gridQuantity = math.Floor(quotePosition / t.gridNumber)

			c1 := df.Low.Crossunder(df.Metadata["lb"])
			if c1 {
				order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
				if err != nil {
					t.ctx.Logger.Error(err)
				}

				t.order = order
				t.totalQuantity = t.totalQuantity + t.order.Quantity
				t.totalCost = t.totalCost + t.order.Price*t.order.Quantity
				t.averagePurchaseCost = t.totalCost / t.totalQuantity
				t.stopLosePoint = t.averagePurchaseCost - df.Metadata["atr"].Last(0)*downPercent
				t.takeProfitPoint = t.averagePurchaseCost + df.Metadata["atr"].Last(0)*upPercent
				t.currentGrid++

				t.trailingStop.Start(t.averagePurchaseCost, t.stopLosePoint)
			}
		}
	} else {
		if quotePosition >= t.gridQuantity && closePrice <= t.stopLosePoint {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
			if err != nil {
				t.ctx.Logger.Error(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity + t.order.Quantity
			t.totalCost = t.totalCost + t.order.Price*t.order.Quantity
			t.averagePurchaseCost = t.totalCost / t.totalQuantity
			t.stopLosePoint = t.averagePurchaseCost - df.Metadata["atr"].Last(0)*downPercent
			t.takeProfitPoint = t.averagePurchaseCost + df.Metadata["atr"].Last(0)*upPercent
			t.currentGrid++

			t.trailingStop.Start(t.averagePurchaseCost, t.stopLosePoint)
		}
	}

	if t.totalCost > minQuote {
		c1 := df.High.Crossover(df.Metadata["ub"])
		c2 := closePrice >= t.takeProfitPoint
		if c1 || c2 {
			order, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, assetPosition)
			if err != nil {
				t.ctx.Logger.Error(err)
			}

			t.order = order
			t.totalQuantity = 0.0
			t.totalCost = 0.0
			t.currentGrid = 0.0
			t.trailingStop.Stop()
		}

		if closePrice < t.averagePurchaseCost && quotePosition < t.gridQuantity {
			if trailing := t.trailingStop; trailing != nil && trailing.Update(closePrice) {
				order, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, assetPosition)
				if err != nil {
					t.ctx.Logger.Error(err)
				}

				t.order = order
				t.totalQuantity = 0.0
				t.totalCost = 0.0
				t.currentGrid = 0.0

				t.trailingStop.Stop()
			}
		}
	}
}
