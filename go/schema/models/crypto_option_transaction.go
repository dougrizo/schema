package models

import (
	"fmt"
	"gorm.io/gorm"
)

type CryptoOptionTransaction struct {
	gorm.Model     `json:"-"`
	ID             int          `gorm:"primaryKey" json:"id"`
	CryptoOptionID int          `gorm:"index:idx_crypto_option_transaction,unique" json:"crypto_option_id"`
	CryptoOption   CryptoOption `json:"crypto_option"`
	Timestamp      uint64       `gorm:"index:idx_crypto_option_transaction,unique" json:"timestamp"`
	Price          float32      `json:"price"`
	Size           int32        `json:"size"`
}

func (co CryptoOptionTransaction) fmt() string {
	return fmt.Sprintf("timestamp=%d, price=%.2f, size=%d", co.Timestamp, co.Price, co.Size)
}

func (co CryptoOptionTransaction) fmtFull() string {
	return fmt.Sprintf("%s, %s", co.CryptoOption.fmtFull(), co.fmt())
}

func (co CryptoOptionTransaction) Log() string {
	return fmt.Sprintf("a cryptoasset option transaction: %s", co.fmt())
}

func (co CryptoOptionTransaction) LogFull() string {
	return fmt.Sprintf("a cryptoasset option transaction: %s", co.fmtFull())
}
