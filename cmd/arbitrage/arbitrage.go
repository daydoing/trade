package arbitrage

import (
	"github.com/daydoing/trade/context"

	"github.com/spf13/cobra"
)

func ArbitrageCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "arbitrage",
		Short: "Arbitraging Uniswap's trading fees",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
