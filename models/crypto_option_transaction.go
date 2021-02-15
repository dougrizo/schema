package models

import "gorm.io/gorm"

type CryptoOptionTransaction struct {
	gorm.Model     `json:"-"`
	ID             int          `gorm:"primaryKey" json:"id"`
	CryptoOptionID int          `json:"crypto_option_id"`
	CryptoOption   CryptoOption `json:"crypto_option"`
	Timestamp      uint64       `json:"timestamp"`
	Price          float64      `json:"price"`
	Size           uint64       `json:"size"`
}
