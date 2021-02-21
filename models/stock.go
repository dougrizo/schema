package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	ID     int    `gorm:"primaryKey"`
	Symbol string `json:"symbol"`
	Hot    bool   `json:"hot"`
}

func (s Stock) fmt() string {
	return fmt.Sprintf("symbol=%s, hot=%v", s.Symbol, s.Hot)
}

func (s Stock) fmtFull() string {
	return s.fmt()
}

func (s Stock) Log() string {
	return fmt.Sprintf("a stock: %s", s.fmt())
}

func (s Stock) LogFull() string {
	return fmt.Sprintf("a stock: %s", s.fmtFull())
}
