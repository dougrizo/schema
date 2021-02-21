package models

import (
	"fmt"
	"gorm.io/gorm"
)

type CryptoOptionQuote struct {
	gorm.Model     `json:"-"`
	ID             int          `gorm:"primaryKey" json:"id"`
	CryptoOptionID int          `json:"crypto_option_id"`
	CryptoOption   CryptoOption `json:"crypto_option"`
	Timestamp      uint64       `json:"timestamp"`
}

func (co CryptoOptionQuote) fmt() string {
	return fmt.Sprintf("timestamp=%d", co.Timestamp)
}

func (co CryptoOptionQuote) fmtFull() string {
	return fmt.Sprintf("%s, %s", co.CryptoOption.fmtFull(), co.fmt())
}

func (co CryptoOptionQuote) log() string {
	return fmt.Sprintf("a cryptoasset option quote: %s", co.fmt())
}

func (co CryptoOptionQuote) logFull() string {
	return fmt.Sprintf("a cryptoasset option quote: %s", co.fmtFull())
}

type CryptoOptionQuoteLevel struct {
	gorm.Model          `json:"-"`
	ID                  int               `gorm:"primaryKey" json:"id"`
	CryptoOptionQuoteID int               `json:"crypto_option_quote_id"`
	CryptoOptionQuote   CryptoOptionQuote `json:"crypto_option_quote"`
	Level               uint8             `json:"level"`
	IsBid               bool              `json:"is_bid"`
	Price               float32           `json:"price"`
	Size                int32             `json:"size"`
}

func (co CryptoOptionQuoteLevel) fmt() string {
	return fmt.Sprintf("level=%d, is_bid=%v, price=%.2f, size=%d", co.Level, co.IsBid, co.Price, co.Size)
}

func (co CryptoOptionQuoteLevel) fmtFull() string {
	return fmt.Sprintf("%s, %s", co.CryptoOptionQuote.logFull(), co.fmt())
}

func (co CryptoOptionQuoteLevel) Log() string {
	return fmt.Sprintf("a cryptoasset option quote level: %s", co.fmt())
}

func (co CryptoOptionQuoteLevel) LogFull() string {
	return fmt.Sprintf("a cryptoasset option quote level: %s", co.fmtFull())
}
