package strategies

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/indicator"
	"github.com/rodrigo-brito/ninjabot/model"
	"github.com/rodrigo-brito/ninjabot/service"
	"github.com/rodrigo-brito/ninjabot/strategy"
)

const (
	minQuantity = 10.0
	buySwing    = 1 + (1.0 / 100)
	downRate    = 1 - (5.0 / 100)
	upRate      = 1 + (5.0 / 100)
)

type trough struct {
	period          int
	currentGrid     int
	gridNumber      float64
	gridQuantity    float64
	totalCost       float64
	totalQuantity   float64
	stopLosePoint   float64
	takeProfitPoint float64
	timeframe       string
	order           model.Order
}

func NewTrough(timeframe string, period int, gridNumber float64) strategy.HighFrequencyStrategy {
	return &trough{timeframe: timeframe, period: period, gridNumber: gridNumber}
}

func (t *trough) Timeframe() string {
	return t.timeframe
}

func (t *trough) WarmupPeriod() int {
	return t.period
}

func (t *trough) Indicators(df *ninjabot.Dataframe) []strategy.ChartIndicator {
	df.Metadata["lowestPrice"] = indicator.Min(df.Low, t.WarmupPeriod())

	return nil
}

func (t *trough) OnCandle(df *ninjabot.Dataframe, broker service.Broker) {
	_, quantity, err := broker.Position(df.Pair)
	if err != nil {
		log.Fatal(err)
	}

	lowestPrice := df.Metadata["lowestPrice"].Last(0)
	currentPrice := df.Close.Last(0)

	if t.currentGrid == 0 {
		t.gridQuantity = quantity / t.gridNumber

		if quantity > minQuantity && currentPrice <= lowestPrice*buySwing {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
			if err != nil {
				log.Fatal(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity + t.order.Quantity
			t.totalCost = t.totalCost + t.order.Price*t.order.Quantity
			t.stopLosePoint = (t.totalCost / t.totalQuantity) * downRate
			t.takeProfitPoint = (t.totalCost / t.totalQuantity) * upRate
			t.currentGrid++
		}
	} else {
		discountStr := fmt.Sprintf("%.2f", downRate-(1.0-downRate)*float64(t.currentGrid-1))
		discount, _ := strconv.ParseFloat(discountStr, 64)

		if quantity >= t.gridQuantity && currentPrice <= t.order.Price*discount {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
			if err != nil {
				log.Fatal(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity + t.order.Quantity
			t.totalCost = t.totalCost + t.order.Price*t.order.Quantity
			t.stopLosePoint = (t.totalCost / t.totalQuantity) * downRate
			t.takeProfitPoint = (t.totalCost / t.totalQuantity) * upRate
			t.currentGrid++
		}
	}

	if t.totalQuantity > minQuantity {
		if currentPrice <= t.stopLosePoint {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeSell, df.Pair, t.totalQuantity)
			if err != nil {
				log.Fatal(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity - t.order.Quantity
			t.totalCost = t.totalCost - t.order.Price*t.order.Quantity
			t.currentGrid = 0
		}

		if currentPrice >= t.takeProfitPoint {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeSell, df.Pair, t.totalQuantity)
			if err != nil {
				log.Fatal(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity - t.order.Quantity
			t.totalCost = t.totalCost - t.order.Price*t.order.Quantity
			t.currentGrid = 0
		}
	}
}

func (t *trough) OnPartialCandle(df *ninjabot.Dataframe, broker service.Broker) {}