package global

import (
	"go-chain-data/config"
	"go-chain-data/config/setting"
	"gorm.io/gorm"
	"log"
)

var (
	DbConfig         *setting.DbConfig
	BlockChainConfig *setting.BlockChainConfig
	DBEngine         *gorm.DB
)

func init() {
	setupConfig()
	setupDBEngine()
}

func setupConfig() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Panic("config2.NewConfig error : ", err)
	}
	err = conf.ReadSection("Database", &DbConfig)
	if err != nil {
		log.Panic("ReadSection - Database error : ", err)
	}
	err = conf.ReadSection("BlockChain", &BlockChainConfig)
	if err != nil {
		log.Panic("ReadSection - BlockChain error : ", err)
	}
}
func setupDBEngine() {
	var err error
	DBEngine, err = config.NewDBEngine(DbConfig)
	if err != nil {
		log.Panic("config.NewDBEngine error : ", err)
	}
}
