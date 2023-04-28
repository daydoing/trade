package arbitrage

import (
	"math"
	"math/big"

	"github.com/daydoing/trade/context"
	"github.com/daydoing/trade/strategies/arbitrage/uni"
	"github.com/ethereum/go-ethereum/common"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/service"
	"github.com/rodrigo-brito/ninjabot/strategy"
)

type Arbitrage struct {
	context.Context
	tokenID           *big.Int
	token0            *core.Token
	token1            *core.Token
	openPositionPrice float64
}

func NewArbitrage(ctx context.Context) (strategy.Strategy, error) {
	return &Arbitrage{
		Context:           ctx,
		tokenID:           big.NewInt(573925),
		token0:            core.NewToken(42161, common.HexToAddress("0x912CE59144191C1204E64559FE8253a0e49E6548"), 18, "ARB", "Arbitrum Coin"),
		token1:            core.NewToken(42161, common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8"), 6, "USDC", "USD Coin"),
		openPositionPrice: 1.41,
	}, nil
}

func (a *Arbitrage) Timeframe() string {
	return a.Config.Strategy.Timeframe
}

func (a *Arbitrage) WarmupPeriod() int {
	return a.Config.Strategy.Period
}

func (a *Arbitrage) Indicators(df *ninjabot.Dataframe) []strategy.ChartIndicator {
	return []strategy.ChartIndicator{}
}

func (a *Arbitrage) OnCandle(df *ninjabot.Dataframe, broker service.Broker) {
	a1, _, err := uni.LPTokenAmount(a.ChainClient, a.tokenID, a.token0, a.token1)
	if err != nil {
		a.Logger.Error(err)
	}

	assetPosition, _, err := broker.Position(df.Pair)
	if err != nil {
		a.Logger.Error(err)
	}

	close := df.Close.Last(0)
	diff := math.Abs(assetPosition - a1)

	if diff*close > a.Config.MinQuote {
		if a1 > assetPosition {
			_, err := broker.CreateOrderMarket(ninjabot.SideTypeSell, df.Pair, diff)
			if err != nil {
				a.Logger.Error(err)
			}
		}

		if a1 < assetPosition && close < a.openPositionPrice {
			_, err := broker.CreateOrderMarket(ninjabot.SideTypeBuy, df.Pair, diff)
			if err != nil {
				a.Logger.Error(err)
			}
		}
	}
}
