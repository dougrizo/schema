package models

import "gorm.io/gorm"

type CryptoOptionQuote struct {
	gorm.Model     `json:"-"`
	ID             int          `gorm:"primaryKey" json:"id"`
	CryptoOptionID int          `json:"crypto_option_id"`
	CryptoOption   CryptoOption `json:"crypto_option"`
	Timestamp      uint64       `json:"timestamp"`
}

type CryptoOptionQuoteLevel struct {
	gorm.Model          `json:"-"`
	ID                  int               `gorm:"primaryKey" json:"id"`
	CryptoOptionQuoteID int               `json:"crypto_option_quote_id"`
	CryptoOptionQuote   CryptoOptionQuote `json:"crypto_option_quote"`
	Level               uint8             `json:"level"`
	IsBid               bool              `json:"is_bid"`
	Price               float64           `json:"price"`
	Size                uint64            `json:"size"`
}
