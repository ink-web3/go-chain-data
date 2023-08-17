package global

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"go-chain-data/config/setting"
	"gorm.io/gorm"
)

var (
	DbConfig         *setting.DbConfig
	BlockChainConfig *setting.BlockChainConfig
	DBEngine         *gorm.DB
	EthRpcClient     *ethclient.Client
)
