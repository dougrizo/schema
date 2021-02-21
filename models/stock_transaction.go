package models

import (
	"fmt"
	"gorm.io/gorm"
)

type StockTransaction struct {
	gorm.Model `json:"-"`
	ID         int     `gorm:"primaryKey" json:"id"`
	StockID    int     `json:"stock_id"`
	Stock      Stock   `json:"stock"`
	Timestamp  uint64  `json:"timestamp"`
	Price      float32 `json:"price"`
	Size       int32   `json:"volume"`
}

func (st StockTransaction) fmt() string {
	return fmt.Sprintf("timestamp=%d, price=%.2f, size=%d", st.Timestamp, st.Price, st.Size)
}

func (st StockTransaction) fmtFull() string {
	return fmt.Sprintf("%s, %s", st.Stock.fmtFull(), st.fmt())
}

func (st StockTransaction) Log() string {
	return fmt.Sprintf("a stock transaction: %s", st.fmt())
}

func (st StockTransaction) LogFull() string {
	return fmt.Sprintf("a stock transaction: %s", st.fmtFull())
}
