package strategies

import (
	"log"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/indicator"
	"github.com/rodrigo-brito/ninjabot/model"
	"github.com/rodrigo-brito/ninjabot/service"
	"github.com/rodrigo-brito/ninjabot/strategy"
)

type trough struct {
	timeframe string
	period    int
	order     model.Order
}

func NewTrough(timeframe string, period int) strategy.HighFrequencyStrategy {
	return &trough{timeframe: timeframe, period: period}
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
	lowestPrice := df.Metadata["lowestPrice"].Last(0)
	currentPrice := df.Close.Last(0)

	asset, quote, err := broker.Position(df.Pair)
	if err != nil {
		log.Fatal(err)
	}

	if quote > 10.0 && asset*df.Close.Last(0) < 10 && currentPrice <= lowestPrice*1.01 {
		order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, quote)
		if err != nil {
			log.Fatal(err)
		}

		t.order = order
	}

	if asset > 0 {
		if currentPrice <= t.order.Price*0.95 {
			_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, asset)
			if err != nil {
				log.Fatal(err)
			}
		}

		if currentPrice >= t.order.Price {
			t.order.Price = currentPrice
		}
	}
}

func (t *trough) OnPartialCandle(df *ninjabot.Dataframe, broker service.Broker) {}
