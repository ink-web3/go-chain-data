package blockchain

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"go-chain-data/global"
	models "go-chain-data/internal/model"
	"math/big"
)

// HandleTransaction 处理交易数据
func HandleTransaction(block *types.Block) error {
	for _, tx := range block.Transactions() {
		receipt, err := global.EthRpcClient.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Error("get transaction fail", "err", err)
		}
		for _, rLog := range receipt.Logs {
			err = HandleTransactionEvent(rLog, receipt.Status)
			if err != nil {
				log.Error("process transaction event fail", "err", err)
			}
		}
		err = ProcessTransaction(tx, block.Number(), receipt.Status)
		if err != nil {
			log.Error("process transaction fail", "err", err)
		}
	}
	return nil
}

func ProcessTransaction(tx *types.Transaction, blockNumber *big.Int, status uint64) error {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		log.Error("Failed to read the sender address", "TxHash", tx.Hash(), "err", err)
		return err
	}
	log.Info("hand transaction", "txHash", tx.Hash().String())
	transaction := &models.Transaction{
		BlockNumber: blockNumber.Uint64(),
		TxHash:      tx.Hash().Hex(),
		From:        from.Hex(),
		Value:       tx.Value().String(),
		Status:      status,
		InputData:   hex.EncodeToString(tx.Data()),
	}
	if tx.To() == nil {
		log.Info("Contract creation found", "Sender", transaction.From, "TxHash", transaction.TxHash)
		toAddress := crypto.CreateAddress(from, tx.Nonce()).Hex()
		transaction.Contract = toAddress
	} else {
		isContract, err := isContractAddress(tx.To().Hex())
		if err != nil {
			return err
		}
		if isContract {
			transaction.Contract = tx.To().Hex()
		} else {
			transaction.To = tx.To().Hex()
		}
	}
	err = transaction.Insert()
	if err != nil {
		log.Error("insert transaction fail", "err", err)
		return err
	}
	return nil
}
