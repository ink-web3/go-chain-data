package blockchain

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	models "go-chain-data/internal/model"
)

func HandleTransactionEvent(rLog *types.Log, status uint64) error {
	log.Info("ProcessTransactionEvent", "address", rLog.Address, "data", rLog.Data)
	event := &models.Events{
		Address:     rLog.Address.String(),
		Data:        "",
		BlockNumber: rLog.BlockNumber,
		TxHash:      rLog.TxHash.String(),
		TxIndex:     rLog.TxIndex,
		BlockHash:   rLog.BlockHash.String(),
		LogIndex:    rLog.Index,
		Removed:     rLog.Removed,
	}
	err := event.Insert()
	if err != nil {
		log.Error("event.Insert() fail", "err", err)
		return err
	}
	evt, err := event.GetEventByTxHash()
	if err != nil {
		log.Error("event.GetEventByTxHash() fail", "err", err)
		return err
	}
	log.Info("Topics", "topic", rLog.Topics)
	for _, tp := range rLog.Topics {
		topic := &models.Topic{
			EventId: evt.Id,
			Topic:   tp.String(),
		}
		err := topic.Insert()
		if err != nil {
			log.Error("topic.Insert() fail", "err", err)
			return err
		}
	}
	return nil
}
