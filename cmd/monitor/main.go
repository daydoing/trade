package monitor

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/daydoing/trade/cmd/monitor/utils"
	"github.com/daydoing/trade/context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

func MonitorCommand(ctx context.BotContext) *cobra.Command {
	return &cobra.Command{
		Use:   "monitor",
		Short: "monitoring on-chain transfers",
		RunE: func(cmd *cobra.Command, args []string) error {
			bot := ctx.NotifyBot
			rawurl := fmt.Sprintf("wss://mainnet.infura.io/ws/v3/%s", ctx.Config.InfuraKey)
			client, err := ethclient.Dial(rawurl)
			if err != nil {
				log.Fatalf("Failed to connect to the Ethereum network: %v", err)
			}

			tokens := map[string]string{
				"0x92D6C1e31e14520e676a687F0a93788B716BEff5": "dydx",
			}

			pairs := map[string]map[string]string{
				"0x92D6C1e31e14520e676a687F0a93788B716BEff5": {
					"0x6cC5F688a315f3dC28A7781717a9A798a59fDA7b": "okx",
					"0x28C6c06298d514Db089934071355E5743bf21d60": "binance14",
					"0x21a31Ee1afC51d94C2eFcCAa2092aD1028285549": "binance15",
					"0xDFd5293D8e347dFe59E90eFd55b2956a1343963d": "binance16",
					"0xf16E9B0D03470827A95CDfd0Cb8a8A3b46969B91": "kucoin",
					"0xf89d7b9c864f589bbF53a82105107622B35EaA40": "bybit",
					"0x75e89d5979E4f6Fba9F97c104c2F0AFB3F1dcB88": "mexc",
				},
			}

			iface, _ := abi.JSON(strings.NewReader(utils.ABI))

			errs := make(chan error)
			for key, item := range pairs {
				go func(key string, item map[string]string) {
					contractAddress := common.HexToAddress(key)
					query := ethereum.FilterQuery{
						Addresses: []common.Address{contractAddress},
						Topics:    [][]common.Hash{{iface.Events["Transfer"].ID}},
					}

					logs := make(chan types.Log)
					sub, err := client.SubscribeFilterLogs(ctx.Context, query, logs)
					if err != nil {
						errs <- err
						return
					}

					for {
						select {
						case <-sub.Err():
							continue
						case v := <-logs:
							if v.Address == contractAddress && v.Topics[0] == iface.Events["Transfer"].ID {
								transferEvent := struct {
									From  common.Address
									To    common.Address
									Value *big.Int
								}{}

								err := iface.UnpackIntoInterface(&transferEvent, "Transfer", v.Data)
								if err != nil {
									errs <- err
									return
								}

								transferEvent.From = common.HexToAddress(v.Topics[1].Hex())
								transferEvent.To = common.HexToAddress(v.Topics[2].Hex())

								token := tokens[contractAddress.String()]
								if to, ok := item[transferEvent.To.String()]; ok {
									result := big.NewInt(0)
									result.Div(transferEvent.Value, big.NewInt(1e+18))
									if result.Cmp(big.NewInt(1000)) == 1 {
										from := transferEvent.From.String()
										if v, ok := item[transferEvent.From.String()]; ok {
											from = v
										}

										content := fmt.Sprintf("%s from %s transferred to %s, number:%s", token, from, to, result.String())
										bot.Notify(content)
									}
								}
							}
						}
					}
				}(key, item)
			}

			err = <-errs

			content := fmt.Sprintf("error:%s, the bot is stoped", err.Error())
			bot.Notify(content)

			return err
		},
	}
}
