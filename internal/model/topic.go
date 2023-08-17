package models

import (
	"go-chain-data/global"
	"gorm.io/gorm"
)

type Topic struct {
	Id      uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	EventId uint64 `json:"event_id"`
	Topic   string `json:"topic" gorm:"type:longtext" `
	*gorm.Model
}

func (tc *Topic) TableName() string {
	return "topics"
}

func (tc *Topic) Insert() error {
	if err := global.DBEngine.Create(&tc).Error; err != nil {
		return err
	}
	return nil
}
