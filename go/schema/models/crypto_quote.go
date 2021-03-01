package models

import (
	"fmt"
	"gorm.io/gorm"
)

type CryptoQuote struct {
	gorm.Model `json:"-"`
	ID         int    `gorm:"primaryKey" json:"id"`
	CryptoID   int    `gorm:"index:idx_crypto_quote" json:"crypto_id"`
	Crypto     Crypto `json:"crypto"`
	Timestamp  uint64 `gorm:"index:idx_crypto_quote" json:"timestamp"`
}

func (cq CryptoQuote) fmt() string {
	return fmt.Sprintf("timestamp=%d", cq.Timestamp)
}

func (cq CryptoQuote) fmtFull() string {
	return fmt.Sprintf("%s, %s", cq.Crypto.fmtFull(), cq.fmt())
}

func (cq CryptoQuote) Log() string {
	return fmt.Sprintf("a cryptoasset quote: %s", cq.fmt())
}

func (cq CryptoQuote) LogFull() string {
	return fmt.Sprintf("a cryptoasset quote: %s", cq.fmtFull())
}

type CryptoQuoteLevel struct {
	gorm.Model    `json:"-"`
	ID            int         `gorm:"primaryKey" json:"id"`
	CryptoQuoteID int         `gorm:"index:idx_crypto_quote_level,unique" json:"crypto_quote_id"`
	CryptoQuote   CryptoQuote `json:"crypto_quote"`
	Level         uint8       `gorm:"index:idx_crypto_quote_level,unique" json:"level"`
	IsBid         bool        `gorm:"index:idx_crypto_quote_level,unique" json:"is_bid"`
	Price         float32     `json:"price"`
	Size          float32     `json:"size"`
}

func (cq CryptoQuoteLevel) fmt() string {
	return fmt.Sprintf("level=%d, is_bid=%v, price=%.6f, size=%.6f", cq.Level, cq.IsBid, cq.Price, cq.Size)
}

func (cq CryptoQuoteLevel) fmtFull() string {
	return fmt.Sprintf("%s, %s", cq.CryptoQuote.fmtFull(), cq.fmt())
}

func (cq CryptoQuoteLevel) Log() string {
	return fmt.Sprintf("a crypto quote level: %s", cq.fmt())
}

func (cq CryptoQuoteLevel) LogFull() string {
	return fmt.Sprintf("a crypto quote level: %s", cq.fmtFull())
}
