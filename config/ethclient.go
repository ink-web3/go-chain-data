package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"go-chain-data/global"
)

func NewEthRpcClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(global.BlockChainConfig.RpcUrl)
	if err != nil {
		return nil, err
	}
	return client, nil
}
