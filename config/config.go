package config

import (
	"github.com/spf13/viper"
	"go-chain-data/global"
	"log"
)

type Config struct {
	vp *viper.Viper
}

func NewConfig() (*Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Config{vp}, nil
}

func (config *Config) ReadSection(k string, v interface{}) error {
	err := config.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

func SetupConfig() {
	conf, err := NewConfig()
	if err != nil {
		log.Panic("NewConfig error : ", err)
	}
	err = conf.ReadSection("Database", &global.DbConfig)
	if err != nil {
		log.Panic("ReadSection - Database error : ", err)
	}
	err = conf.ReadSection("BlockChain", &global.BlockChainConfig)
	if err != nil {
		log.Panic("ReadSection - BlockChain error : ", err)
	}
}
func SetupDBEngine() {
	var err error
	global.DBEngine, err = NewDBEngine(global.DbConfig)
	if err != nil {
		log.Panic("NewDBEngine error : ", err)
	}
}

func SetupEthClient() {
	var err error
	global.EthRpcClient, err = NewEthRpcClient()
	if err != nil {
		log.Panic("NewEthRpcClient error : ", err)
	}
}
