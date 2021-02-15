package models

import "gorm.io/gorm"

type CryptoTransaction struct {
	gorm.Model `json:"-"`
	ID         int     `gorm:"primaryKey" json:"id"`
	CryptoID   int     `json:"crypto_id"`
	Crypto     Crypto  `json:"crypto"`
	Timestamp  uint64  `json:"timestamp"`
	Price      float64 `json:"price"`
	Size       uint64  `json:"size"`
}
