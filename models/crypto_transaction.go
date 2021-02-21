package models

import (
	"fmt"
	"gorm.io/gorm"
)

type CryptoTransaction struct {
	gorm.Model `json:"-"`
	ID         int     `gorm:"primaryKey" json:"id"`
	CryptoID   int     `json:"crypto_id"`
	Crypto     Crypto  `json:"crypto"`
	Timestamp  uint64  `json:"timestamp"`
	Price      float32 `json:"price"`
	Size       int32   `json:"size"`
}

func (ct CryptoTransaction) fmt() string {
	return fmt.Sprintf("timestamp=%d, price=%.2f, size=%d", ct.Timestamp, ct.Price, ct.Size)
}

func (ct CryptoTransaction) fmtFull() string {
	return fmt.Sprintf("%s, %s", ct.Crypto.fmtFull(), ct.fmt())
}

func (ct CryptoTransaction) Log() string {
	return fmt.Sprintf("a cryptoasset transaction: %s", ct.fmt())
}

func (ct CryptoTransaction) LogFull() string {
	return fmt.Sprintf("a cryptoasset transaction: %s", ct.fmtFull())
}
