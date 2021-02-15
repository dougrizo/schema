package models

import (
	"gorm.io/gorm"
)

type CryptoTrade struct {
	gorm.Model  `json:"-"`
	ID          int     `gorm:"primaryKey" json:"-"`
	CryptoID    int     `json:"crypto_id"`
	Crypto      Crypto  `json:"crypto"`
	Timestamp   uint64  `json:"timestamp"`
	Price       float64 `json:"price"`
	Size        int64   `json:"size"`
	MarketMaker bool    `json:"market_maker"`
}
