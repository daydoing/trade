package strategies

import (
	"fmt"
	"strconv"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/indicator"
	"github.com/rodrigo-brito/ninjabot/model"
	"github.com/rodrigo-brito/ninjabot/service"
	"github.com/rodrigo-brito/ninjabot/strategy"
	"github.com/rodrigo-brito/ninjabot/tools/log"
)

const (
	minQuantity = 10.0
	buySwing    = 1 + (1.0 / 100)
)

type trough struct {
	period              int
	currentGrid         int
	gridNumber          float64
	gridQuantity        float64
	totalCost           float64
	totalQuantity       float64
	stopLosePoint       float64
	averagePurchaseCost float64
	downRate            float64
	upRate              float64
	timeframe           string
	order               model.Order
}

func NewTrough(timeframe string, period int, gridNumber float64, downRate float64, upRate float64) strategy.HighFrequencyStrategy {
	return &trough{
		timeframe:  timeframe,
		period:     period,
		gridNumber: gridNumber,
		downRate:   1 - (downRate / 100),
		upRate:     1 + (upRate / 100),
	}
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
	t.execStrategy(df, broker)
}

func (t *trough) OnPartialCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execStrategy(df, broker)
}

func (t *trough) execStrategy(df *ninjabot.Dataframe, broker service.Broker) {
	_, quotePosition, err := broker.Position(df.Pair)
	if err != nil {
		log.Fatal(err)
	}

	lowestPrice := df.Metadata["lowestPrice"].Last(0)
	closePrice := df.Close.Last(0)

	if t.currentGrid == 0 {
		t.gridQuantity = quotePosition / t.gridNumber
		if quotePosition > minQuantity && closePrice <= lowestPrice*buySwing {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, t.gridQuantity)
			if err != nil {
				log.Error(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity + t.order.Quantity
			t.totalCost = t.totalCost + t.order.Price*t.order.Quantity
			t.averagePurchaseCost = t.totalCost / t.totalQuantity
			t.stopLosePoint = t.averagePurchaseCost
			t.currentGrid++
		}
	} else {
		discountStr := fmt.Sprintf("%.2f", 1.0-(1.0-t.downRate)*float64(t.currentGrid))
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
			t.stopLosePoint = t.averagePurchaseCost
			t.currentGrid++
		}
	}

	if t.totalCost > minQuantity {
		if closePrice > t.stopLosePoint {
			t.stopLosePoint = closePrice
		}

		if t.stopLosePoint < t.averagePurchaseCost*t.downRate || t.stopLosePoint > t.averagePurchaseCost*t.upRate {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeSell, df.Pair, t.totalCost)
			if err != nil {
				log.Error(err)
			}

			t.order = order
			t.totalQuantity = 0.0
			t.totalCost = 0.0
			t.currentGrid = 0.0
		}
	}
}
