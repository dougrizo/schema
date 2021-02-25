package models

import (
	"fmt"
	"gorm.io/gorm"
)

type StockCandle struct {
	gorm.Model `json:"-"`
	ID         int     `gorm:"primaryKey" json:"id"`
	StockID    int     `gorm:"index:idx_stock_candle,unique" json:"stock_id"`
	Stock      Stock   `json:"stock"`
	Timestamp  uint64  `gorm:"index:idx_stock_candle,unique" json:"timestamp"`
	Open       float32 `json:"open"`
	High       float32 `json:"high"`
	Low        float32 `json:"low"`
	Close      float32 `json:"close"`
	Volume     int32   `json:"volume"`
}

func (sc StockCandle) fmt() string {
	return fmt.Sprintf(
		"timestamp=%d, open=%.2f, high=%.2f, low=%.2f, close=%.2f, volume=%d",
		sc.Timestamp, sc.Open, sc.High, sc.Low, sc.Close, sc.Volume)
}

func (sc StockCandle) fmtFull() string {
	return fmt.Sprintf("%s, %s", sc.Stock.fmtFull(), sc.fmt())
}

func (sc StockCandle) Log() string {
	return fmt.Sprintf("a stock candle: %s", sc.fmt())
}

func (sc StockCandle) LogFull() string {
	return fmt.Sprintf("a stock cande: %s", sc.fmtFull())
}
