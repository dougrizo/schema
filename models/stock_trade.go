package models

import "gorm.io/gorm"

type StockTrade struct {
	gorm.Model `json:"-"`
	ID         int     `json:"id"`
	StockID    int     `json:"stock_id"`
	Stock      Stock   `json:"stock"`
	Timestamp  uint64  `json:"timestamp"`
	Price      float64 `json:"price"`
	Size       int64   `json:"size"`
}
