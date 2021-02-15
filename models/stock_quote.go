package models

import "gorm.io/gorm"

type StockQuote struct {
	gorm.Model `json:"-"`
	ID         int     `gorm:"primaryKey" json:"id"`
	StockID    int     `json:"stock_id"`
	Stock      Stock   `json:"stock"`
	Timestamp  uint64  `json:"timestamp"`
	BidPrice   float64 `json:"bid_price"`
	BidSize    int     `json:"bid_size"`
	AskPrice   float64 `json:"ask_price"`
	AskSize    int     `json:"ask_size"`
}
