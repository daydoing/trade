package depth

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/common"
	"github.com/daydoing/trade/context"
	"github.com/rodrigo-brito/ninjabot/tools/log"
	"github.com/spf13/cobra"
)

func DepthCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "arbitrage",
		Short: "Arbitraging Uniswap's trading fees",
		RunE: func(cmd *cobra.Command, args []string) error {
			const symbol = "DYDXUSDT"

			depthHandler := func(event *binance.WsDepthEvent) {
				bids, asks := event.Bids, event.Asks

				if len(bids) == 0 || len(asks) == 0 {
					return
				}

				// 计算买卖盘的加权平均价格和数量
				bidPrice, bidQuantity, err := calculateWeightedAverage(bids)
				if err != nil {
					log.Error("parse common.PriceLevel error:", err.Error())
					return
				}
				askPrice, askQuantity, err := calculateWeightedAverage(asks)
				if err != nil {
					log.Error("parse common.PriceLevel error:", err.Error())
					return
				}

				// 计算买卖盘价格和数量的差值
				priceDiff := bidPrice - askPrice
				quantityDiff := bidQuantity - askQuantity

				if priceDiff > 0 && quantityDiff > 0 {
					fmt.Println("buy single")
				}

				if priceDiff < 0 && quantityDiff < 0 {
					fmt.Println("sell single")
				}
			}

			errHandler := func(err error) {
				log.Error("websocket error:", err.Error())
			}

			done, _, err := binance.WsDepthServe(symbol, depthHandler, errHandler)
			if err != nil {
				panic(err)
			}

			<-done

			return nil
		},
	}
}

// 计算加权平均价格和数量
func calculateWeightedAverage(levels []common.PriceLevel) (float64, float64, error) {
	totalQuantity := 0.0
	totalPrice := 0.0

	for _, v := range levels {
		price, quantity, err := v.Parse()
		if err != nil {
			return 0.0, 0.0, err
		}

		totalQuantity += quantity
		totalPrice += price * quantity
	}

	averagePrice := totalPrice / totalQuantity

	return averagePrice, totalQuantity, nil
}
