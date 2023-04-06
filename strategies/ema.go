package strategies

import (
	"log"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/indicator"
	"github.com/rodrigo-brito/ninjabot/service"
	"github.com/rodrigo-brito/ninjabot/strategy"
)

type crossEMA struct{}

func NewCrossEMA() strategy.Strategy {
	return &crossEMA{}
}

func (e *crossEMA) Timeframe() string {
	return "1d" // examples: 1m, 5m, 15m, 30m, 1h, 4h, 12h, 1d, 1w
}

func (e *crossEMA) WarmupPeriod() int {
	return 9 // warmup period, to preload indicators
}

func (e *crossEMA) Indicators(df *ninjabot.Dataframe) []strategy.ChartIndicator {
	// define a custom indicator, Exponential Moving Average of 9 periods
	df.Metadata["ema9"] = indicator.EMA(df.Close, 9)

	// (Optional) you can return a list of indicators to include in the final chart
	return []strategy.ChartIndicator{
		{
			Overlay:   true,
			Time:      df.Time,
			GroupName: "EMA",
			Metrics: []strategy.IndicatorMetric{
				{
					Values: df.Metadata["ema9"],
					Name:   "EMA 9",
					Color:  "red",
					Style:  strategy.StyleLine,
				},
			},
		},
	}
}

func (e *crossEMA) OnCandle(df *ninjabot.Dataframe, broker service.Broker) {
	// Get the quote and assets information
	assetPosition, quotePosition, err := broker.Position(df.Pair)
	if err != nil {
		log.Fatal(err)
	}

	// Check if we have more than 10 USDT available in the wallet and the buy signal is triggered
	if quotePosition > 10 && df.Close.Crossover(df.Metadata["ema9"]) {
		_, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeBuy, df.Pair, quotePosition*0.99)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Check if we have position in the pair and the sell signal is triggered
	if assetPosition > 0 &&
		df.Close.Crossunder(df.Metadata["ema9"]) {
		_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, assetPosition)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	RegisterStrategy("ema", NewCrossEMA())
}
