package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Crypto struct {
	gorm.Model `json:"-"`
	ID         int    `gorm:"primaryKey" json:"id"`
	Symbol     string `json:"symbol"`
	Hot        bool   `json:"hot"`
}

func (c Crypto) fmt() string {
	return fmt.Sprintf("symbol=%s, hot=%v", c.Symbol, c.Hot)
}

func (c Crypto) fmtFull() string {
	return c.fmt()
}

func (c Crypto) Log() string {
	return fmt.Sprintf("a cryptoasset: %s", c.fmt())
}

func (c Crypto) LogFull() string {
	return c.Log()
}
