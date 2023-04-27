package uni

import (
	"math/big"
	"strconv"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/examples/contract"
	"github.com/daoleno/uniswapv3-sdk/examples/helper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LPTokenAmount(client *ethclient.Client, tokenID *big.Int, token0, token1 *core.Token) (float64, float64, error) {
	pool, err := helper.ConstructV3Pool(client, token0, token1, uint64(constants.FeeMedium))
	if err != nil {
		return 0.0, 0.0, err
	}

	pmc, err := contract.NewUniswapv3NFTPositionManager(common.HexToAddress(helper.ContractV3NFTPositionManager), client)
	if err != nil {
		return 0.0, 0.0, err
	}

	p, err := pmc.Positions(nil, tokenID)
	if err != nil {
		return 0.0, 0.0, err
	}

	po, err := entities.NewPosition(pool, p.Liquidity, int(p.TickLower.Int64()), int(p.TickUpper.Int64()))
	if err != nil {
		return 0.0, 0.0, err
	}

	a0, err := po.Amount0()
	if err != nil {
		return 0.0, 0.0, err
	}

	a1, err := po.Amount1()
	if err != nil {
		return 0.0, 0.0, err
	}

	bf1 := new(big.Float).Quo(new(big.Float).SetInt(a0.Quotient()), new(big.Float).SetInt64(1e18))
	bf2 := new(big.Float).Quo(new(big.Float).SetInt(a1.Quotient()), new(big.Float).SetInt64(1e6))

	f1, _ := strconv.ParseFloat(bf1.String(), 64)
	f2, _ := strconv.ParseFloat(bf2.String(), 64)

	return f1, f2, nil
}
