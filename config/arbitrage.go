package config

type Arbitrage struct {
	RpcURL     string `mapstructure:"rpc_url"`
	PrivateKey string `mapstructure:"private_key"`
	ChainID    int64  `mapstructure:"chain_id"`
}
