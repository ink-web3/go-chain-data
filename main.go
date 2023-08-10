package main

import (
	"go-chain-data/global"
	models "go-chain-data/internal/model"
	"log"
)

func main() {
	log.Println(global.BlockChainConfig.RpcUrl)

	block := models.Blocks{
		BlockHeight:       1,
		BlockHash:         "hash",
		ParentHash:        "parentHash",
		LatestBlockHeight: 2,
	}
	err := block.Insert()
	if err != nil {
		log.Panic("block.Insert error : ", err)
	}
}
