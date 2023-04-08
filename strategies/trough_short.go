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
	sbbPeriod   = 20
	sdeviation  = 3
	sgridNumber = 3.0
	sdrawdown   = 3.0
)

type troughShort struct {
	ctx                 context.Context
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

func NewTroughShort(srv context.Context) strategy.HighFrequencyStrategy {
	return &troughShort{
		ctx:          srv,
		timeframe:    srv.Config.Strategy.Timeframe,
		period:       srv.Config.Strategy.Period,
		gridNumber:   sgridNumber,
		drawdown:     sdrawdown,
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

func (t *troughShort) OnCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execStrategy(df, broker)
}

func (t *troughShort) OnPartialCandle(df *ninjabot.Dataframe, broker service.Broker) {
	t.execStrategy(df, broker)
}

func (t *troughShort) execStrategy(df *ninjabot.Dataframe, broker service.Broker) {
	assetPosition, quotePosition, err := broker.Position(df.Pair)
	if err != nil {
		t.ctx.Logger.Error(err)
	}

	var (
		minQuote   = t.ctx.Config.MinQuote
		closePrice = df.Close.Last(0)
	)

	if t.currentGrid == 0 && quotePosition > minQuote {
		t.gridQuantity = math.Floor(quotePosition / t.gridNumber)

		c1 := df.High.Crossover(df.Metadata["ub"])
		if c1 {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeSell, df.Pair, t.gridQuantity)
			if err != nil {
				t.ctx.Logger.Error(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity + t.order.Quantity
			t.totalCost = t.totalCost + t.order.Price*t.order.Quantity
			t.averagePurchaseCost = t.totalCost / t.totalQuantity
			t.currentGrid++

			t.trailingStop.Start(t.averagePurchaseCost, t.averagePurchaseCost*(1+t.drawdown/100))
		}
	} else {
		if quotePosition >= t.gridQuantity && closePrice >= t.order.Price*(1.0+t.drawdown/100) {
			order, err := broker.CreateOrderMarketQuote(ninjabot.SideTypeSell, df.Pair, t.gridQuantity)
			if err != nil {
				t.ctx.Logger.Error(err)
			}

			t.order = order
			t.totalQuantity = t.totalQuantity + t.order.Quantity
			t.totalCost = t.totalCost + t.order.Price*t.order.Quantity
			t.averagePurchaseCost = t.totalCost / t.totalQuantity
			t.currentGrid++

			t.trailingStop.Start(t.averagePurchaseCost, t.averagePurchaseCost*(1+t.drawdown/100))
		}
	}

	if t.totalCost > minQuote {
		absAssetPosition := math.Abs(assetPosition)

		c1 := df.Low.Crossunder(df.Metadata["lb"])
		c2 := closePrice < t.averagePurchaseCost
		if c1 && c2 {
			order, err := broker.CreateOrderMarket(ninjabot.SideTypeBuy, df.Pair, absAssetPosition)
			if err != nil {
				t.ctx.Logger.Error(err)
			}

			t.order = order
			t.totalQuantity = 0.0
			t.totalCost = 0.0
			t.currentGrid = 0.0
			t.trailingStop.Stop()
		}

		if closePrice > t.averagePurchaseCost && quotePosition < t.gridQuantity {
			if trailing := t.trailingStop; trailing != nil && trailing.Update(closePrice) {
				order, err := broker.CreateOrderMarket(ninjabot.SideTypeBuy, df.Pair, absAssetPosition)
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
