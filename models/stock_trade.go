package models

import (
	"fmt"
	"gorm.io/gorm"
)

type StockTrade struct {
	gorm.Model `json:"-"`
	ID         int     `json:"id"`
	StockID    int     `gorm:"index:idx_stock_trade,unique" json:"stock_id"`
	Stock      Stock   `json:"stock"`
	Timestamp  uint64  `gorm:"index:idx_stock_trade,unique" json:"timestamp"`
	Price      float32 `json:"price"`
	Size       int32   `json:"size"`
}

func (st StockTrade) fmt() string {
	return fmt.Sprintf("timestamp=%d, price=%.2f, size=%d", st.Timestamp, st.Price, st.Size)
}

func (st StockTrade) fmtFull() string {
	return fmt.Sprintf("%s, %s", st.Stock.fmtFull(), st.fmt())
}

func (st StockTrade) Log() string {
	return fmt.Sprintf("a stock trade: %s", st.fmt())
}

func (st StockTrade) LogFull() string {
	return fmt.Sprintf("a stock trade: %s", st.fmtFull())
}
