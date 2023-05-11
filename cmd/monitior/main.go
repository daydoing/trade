package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	var (
		okx            = common.HexToAddress("0x6cC5F688a315f3dC28A7781717a9A798a59fDA7b")
		binanceAddress = common.HexToAddress("0x28C6c06298d514Db089934071355E5743bf21d60")
		kuCoin         = common.HexToAddress("0xf16E9B0D03470827A95CDfd0Cb8a8A3b46969B91")
		cexio          = common.HexToAddress("0xc9f5296Eb3ac266c94568D790b6e91ebA7D76a11")
	)

	// Solidity事件的接口
	iface, _ := abi.JSON(strings.NewReader(`
        [
            {
                "anonymous": false,
                "inputs": [
                    {
                        "indexed": true,
                        "internalType": "address",
                        "name": "from",
                        "type": "address"
                    },
                    {
                        "indexed": true,
                        "internalType": "address",
                        "name": "to",
                        "type": "address"
                    },
                    {
                        "indexed": false,
                        "internalType": "uint256",
                        "name": "value",
                        "type": "uint256"
                    }
                ],
                "name": "Transfer",
                "type": "event"
            }
        ]
    `))

	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/130ac87f7aef49d497402ef015078324")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}

	contractAddress := common.HexToAddress("0x92D6C1e31e14520e676a687F0a93788B716BEff5")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case v := <-logs:
			if v.Address == contractAddress && v.Topics[0] == common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef") {
				// 解码Transfer事件
				transferEvent := struct {
					From  common.Address
					To    common.Address
					Value *big.Int
				}{}

				err := iface.UnpackIntoInterface(&transferEvent, "Transfer", v.Data)
				if err != nil {
					log.Fatalf("Failed to parse transfer log: %v", err)
				}

				switch transferEvent.To {
				case binanceAddress:
					fmt.Println("dydx tokens transferred to binance:", transferEvent.Value.String())
				case okx:
					fmt.Println("dydx tokens transferred to okx:", transferEvent.Value.String())
				case kuCoin:
					fmt.Println("dydx tokens transferred to kucoin:", transferEvent.Value.String())
				case cexio:
					fmt.Println("dydx tokens transferred to cex.io:", transferEvent.Value.String())
				}
			}
		}
	}
}
