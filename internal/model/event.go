package models

import (
	"go-chain-data/global"
	"gorm.io/gorm"
)

type Events struct {
	Id          uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT" `
	Address     string `json:"address" gorm:"type:char(42)" `
	Data        string `json:"data" gorm:"type:longtext" `
	BlockNumber uint64 `json:"block_number"`
	TxHash      string `json:"tx_hash" gorm:"type:char(66)" `
	TxIndex     uint   `json:"tx_index" `
	BlockHash   string `json:"block_hash" gorm:"type:varchar(256)" `
	LogIndex    uint   `json:"log_index"`
	Removed     bool   `json:"removed"`
	*gorm.Model
}

func (e *Events) TableName() string {
	return "events"
}

func (e *Events) Insert() error {
	if err := global.DBEngine.Create(&e).Error; err != nil {
		return err
	}
	return nil
}

func (e *Events) GetEventByTxHash() (*Events, error) {
	var event Events
	if err := global.DBEngine.Where("tx_hash = ?", e.TxHash).First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
