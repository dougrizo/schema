package models

import (
	"fmt"
	"gorm.io/gorm"
)

type StockQuote struct {
	gorm.Model `json:"-"`
	ID         int     `gorm:"primaryKey" json:"id"`
	StockID    int     `json:"stock_id"`
	Stock      Stock   `json:"stock"`
	Timestamp  uint64  `json:"timestamp"`
	Bid        float32 `json:"bid"`
	BidSize    int     `json:"bid_size"`
	Ask        float32 `json:"ask"`
	AskSize    int     `json:"ask_size"`
}

func (sq StockQuote) fmt() string {
	return fmt.Sprintf(
		"timestamp=%d, bid=%.2f, bid_size=%d, ask=%.2f, ask_size=%d",
		sq.Timestamp, sq.Bid, sq.BidSize, sq.Ask, sq.AskSize)
}

func (sq StockQuote) fmtFull() string {
	return fmt.Sprintf("%s, %s", sq.Stock.fmtFull(), sq.fmt())
}

func (sq StockQuote) Log() string {
	return fmt.Sprintf("a stock quote: %s", sq.fmt())
}

func (sq StockQuote) LogFull() string {
	return fmt.Sprintf("a stock quote: %s", sq.fmtFull())
}
