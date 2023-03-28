package strategies

import (
	"log"

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
	period        int
	gridNumber    int
	currentGrid   int
	gridQuantity  float64
	totalCost     float64
	totalQuantity float64
	timeframe     string
	order         model.Order
}

func NewTrough(timeframe string, period int, gridNumber int) strategy.HighFrequencyStrategy {
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
		t.gridQuantity = quantity / float64(t.gridNumber)

		if quantity > minQuantity && currentPrice <= lowestPrice*buySwing {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
			if err != nil {
				log.Fatal(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity + t.order.Quantity
			t.totalCost = t.totalCost + t.order.Price*t.order.Quantity
			t.currentGrid++
		}
	} else {
		if quantity >= t.gridQuantity && currentPrice <= t.order.Price*downRate {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
			if err != nil {
				log.Fatal(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity + t.order.Quantity
			t.totalCost = t.totalCost + t.order.Price*t.order.Quantity
			t.currentGrid++
		}
	}

	if t.totalQuantity > 0 {
		avp := t.totalCost / t.totalQuantity
		if t.totalQuantity > minQuantity && currentPrice >= avp*upRate {
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
