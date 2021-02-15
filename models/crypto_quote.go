package models

import "gorm.io/gorm"

type CryptoQuote struct {
	gorm.Model `json:"-"`
	ID         int    `gorm:"primaryKey" json:"id"`
	CryptoID   int    `json:"crypto_id"`
	Crypto     Crypto `json:"crypto"`
	Timestamp  uint64 `json:"timestamp"`
}

type CryptoQuoteLevel struct {
	gorm.Model    `json:"-"`
	ID            int         `gorm:"primaryKey" json:"id"`
	CryptoQuoteID int         `json:"crypto_quote_id"`
	CryptoQuote   CryptoQuote `json:"crypto_quote"`
	Level         uint8       `json:"level"`
	IsBid         bool        `json:"is_bid"`
	Price         float64     `json:"price"`
	Size          uint64      `json:"size"`
}
