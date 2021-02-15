package models

import "gorm.io/gorm"

type CryptoOption struct {
	gorm.Model `json:"-"`
	ID         int     `gorm:"primaryKey" json:"id"`
	CryptoID   int     `json:"crypto_id"`
	Crypto     Crypto  `json:"crypto"`
	Expiration uint64  `json:"expiration"`
	Strike     float64 `json:"strike"`
	IsPut      bool    `json:"is_put"`
}
