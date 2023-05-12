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
		okx       = common.HexToAddress("0x6cC5F688a315f3dC28A7781717a9A798a59fDA7b")
		binance14 = common.HexToAddress("0x28C6c06298d514Db089934071355E5743bf21d60")
		binance15 = common.HexToAddress("0x21a31Ee1afC51d94C2eFcCAa2092aD1028285549")
		binance16 = common.HexToAddress("0xDFd5293D8e347dFe59E90eFd55b2956a1343963d")
		kucoin    = common.HexToAddress("0xf16E9B0D03470827A95CDfd0Cb8a8A3b46969B91")
		cexio     = common.HexToAddress("0xc9f5296Eb3ac266c94568D790b6e91ebA7D76a11")
		bybit     = common.HexToAddress("0xf89d7b9c864f589bbF53a82105107622B35EaA40")
		mexc      = common.HexToAddress("0x75e89d5979E4f6Fba9F97c104c2F0AFB3F1dcB88")
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
		Topics:    [][]common.Hash{{iface.Events["Transfer"].ID}},
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
			if v.Address == contractAddress && v.Topics[0] == iface.Events["Transfer"].ID {
				transferEvent := struct {
					From  common.Address
					To    common.Address
					Value *big.Int
				}{}

				err := iface.UnpackIntoInterface(&transferEvent, "Transfer", v.Data)
				if err != nil {
					log.Fatalf("Failed to parse transfer log: %v", err)
				}

				transferEvent.From = common.HexToAddress(v.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(v.Topics[2].Hex())

				switch transferEvent.To {
				case binance14, binance15, binance16:
					fmt.Println("dydx tokens transferred to binance:", transferEvent.Value.String())
				case okx:
					fmt.Println("dydx tokens transferred to okx:", transferEvent.Value.String())
				case kucoin:
					fmt.Println("dydx tokens transferred to kucoin:", transferEvent.Value.String())
				case cexio:
					fmt.Println("dydx tokens transferred to cex.io:", transferEvent.Value.String())
				case bybit:
					fmt.Println("dydx tokens transferred to bybit:", transferEvent.Value.String())
				case mexc:
					fmt.Println("dydx tokens transferred to mexc:", transferEvent.Value.String())
				}
			}
		}
	}
}
