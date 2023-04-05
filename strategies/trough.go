package strategies

import (
	"fmt"
	"math"
	"strconv"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/indicator"
	"github.com/rodrigo-brito/ninjabot/model"
	"github.com/rodrigo-brito/ninjabot/service"
	"github.com/rodrigo-brito/ninjabot/strategy"
	"github.com/rodrigo-brito/ninjabot/tools"
	"github.com/rodrigo-brito/ninjabot/tools/log"
)

const (
	minQuote  = 10.0
	bbPeriod  = 20
	deviation = 3
)

type trough struct {
	period              int
	currentGrid         int
	gridNumber          float64
	gridQuantity        float64
	totalCost           float64
	totalQuantity       float64
	averagePurchaseCost float64
	drawdown            float64
	timeframe           string
	order               model.Order
	trailingStop        *tools.TrailingStop
}

func NewTrough(timeframe string, period int, gridNumber float64, drawdown float64) strategy.HighFrequencyStrategy {
	return &trough{
		timeframe:    timeframe,
		period:       period,
		gridNumber:   gridNumber,
		drawdown:     drawdown,
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
	df.Metadata["minPrice"] = indicator.Min(df.Low, t.WarmupPeriod())
	df.Metadata["maxPrice"] = indicator.Max(df.High, t.WarmupPeriod())
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

func (t *trough) OnCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execStrategy(df, broker)
}

func (t *trough) OnPartialCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execStrategy(df, broker)
}

func (t *trough) execStrategy(df *ninjabot.Dataframe, broker service.Broker) {
	assetPosition, quotePosition, err := broker.Position(df.Pair)
	if err != nil {
		log.Fatal(err)
	}

	var (
		closePrice = df.Close.Last(0)
	)

	if t.currentGrid == 0 && quotePosition > minQuote {
		t.gridQuantity = math.Floor(quotePosition / t.gridNumber)

		//1.最低价格下穿布林线下轨
		c1 := df.Low.Crossunder(df.Metadata["lb"])

		if c1 {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
			if err != nil {
				log.Error(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity + t.order.Quantity
			t.totalCost = t.totalCost + t.order.Price*t.order.Quantity
			t.averagePurchaseCost = t.totalCost / t.totalQuantity
			t.currentGrid++

			t.trailingStop.Start(t.averagePurchaseCost, t.averagePurchaseCost*(1-t.drawdown/100))
		}
	} else {
		discountStr := fmt.Sprintf("%.2f", 1.0-(t.drawdown/100)*float64(t.currentGrid))
		discount, _ := strconv.ParseFloat(discountStr, 64)

		if quotePosition >= t.gridQuantity && closePrice <= t.order.Price*discount {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
			if err != nil {
				log.Error(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity + t.order.Quantity
			t.totalCost = t.totalCost + t.order.Price*t.order.Quantity
			t.averagePurchaseCost = t.totalCost / t.totalQuantity
			t.currentGrid++

			t.trailingStop.Start(t.averagePurchaseCost, t.averagePurchaseCost*(1-t.drawdown/100))
		}
	}

	if t.totalCost > minQuote {
		if closePrice < t.averagePurchaseCost && quotePosition < t.gridQuantity {
			if trailing := t.trailingStop; trailing != nil && trailing.Update(closePrice) {
				order, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, assetPosition)
				if err != nil {
					log.Error(err)
				}

				t.order = order
				t.totalQuantity = 0.0
				t.totalCost = 0.0
				t.currentGrid = 0.0

				t.trailingStop.Stop()
			}
		}

		if closePrice > t.averagePurchaseCost {
			if trailing := t.trailingStop; trailing != nil && trailing.Update(closePrice) {
				order, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, assetPosition)
				if err != nil {
					log.Error(err)
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
