package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go-chain-data/global"
	models "go-chain-data/internal/model"
	"log"
	"math/big"
	"time"
)

// InitBlock 初始化第一个区块数据
func InitBlock() {
	block := &models.Blocks{}
	count := block.Counts()
	if count == 0 {
		lastBlockNumber, err := global.EthRpcClient.BlockNumber(context.Background())
		if err != nil {
			log.Panic("InitBlock - BlockNumber err : ", err)
		}
		lastBlock, err := global.EthRpcClient.BlockByNumber(context.Background(), big.NewInt(int64(lastBlockNumber)))

		if err != nil {
			log.Panic("InitBlock - BlockByNumber err : ", err)
		}
		block.BlockHash = lastBlock.Hash().Hex()
		block.BlockHeight = lastBlock.NumberU64()
		block.LatestBlockHeight = lastBlock.NumberU64()
		block.ParentHash = lastBlock.ParentHash().Hex()
		err = block.Insert()
		if err != nil {
			log.Panic("InitBlock - Insert block err : ", err)
		}
	}
}

func SyncTask() {
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			latestBlockNumber, err := global.EthRpcClient.BlockNumber(context.Background())
			if err != nil {
				log.Panic("EthRpcClient.BlockNumber error : ", err)
			}
			var blocks models.Blocks
			latestBlock, err := blocks.GetLatest()
			if err != nil {
				log.Panic("blocks.GetLatest error : ", err)
			}
			if latestBlock.LatestBlockHeight > latestBlockNumber {
				log.Printf("latestBlock.LatestBlockHeight : %v greater than latestBlockNumber : %v \n", latestBlock.LatestBlockHeight, latestBlockNumber)
				continue
			}
			currentBlock, err := global.EthRpcClient.BlockByNumber(context.Background(), big.NewInt(int64(latestBlock.LatestBlockHeight)))
			if err != nil {
				log.Panic("EthRpcClient.BlockByNumber error : ", err)
			}
			log.Printf("get currentBlock blockNumber : %v , blockHash : %v \n", currentBlock.Number(), currentBlock.Hash().Hex())
			err = HandleBlock(currentBlock)
			if err != nil {
				log.Panic("HandleBlock error : ", err)
			}
		}
	}
}

// HandleBlock 处理区块信息
func HandleBlock(currentBlock *types.Block) error {
	block := &models.Blocks{
		BlockHeight:       currentBlock.NumberU64(),
		BlockHash:         currentBlock.Hash().Hex(),
		ParentHash:        currentBlock.ParentHash().Hex(),
		LatestBlockHeight: currentBlock.NumberU64() + 1,
	}
	err := block.Insert()
	if err != nil {
		return err
	}
	err = HandleTransaction(currentBlock)
	if err != nil {
		return err
	}
	return nil
}

// 判断一个地址是否是合约地址
func isContractAddress(address string) (bool, error) {

	addr := common.HexToAddress(address)
	code, err := global.EthRpcClient.CodeAt(context.Background(), addr, nil)
	if err != nil {
		return false, err
	}

	return len(code) > 0, nil
}
