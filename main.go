package main

import (
	"go-chain-data/config"
	"go-chain-data/pkg/blockchain"
	"log"
)

func init() {
	config.SetupConfig()
	config.SetupDBEngine()

	err := config.MigrateDb()
	if err != nil {
		log.Panic("config.MigrateDb error : ", err)
	}
	config.SetupEthClient()
}

func main() {
	blockchain.InitBlock()
	blockchain.SyncTask()
}
