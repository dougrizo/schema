package models

import (
	"fmt"
	"gorm.io/gorm"
)

type CryptoOption struct {
	gorm.Model    `json:"-"`
	ID            int         `gorm:"primaryKey" json:"id"`
	CryptoID      int         `json:"crypto_id"`
	Crypto        Crypto      `json:"crypto"`
	CryptoQuoteID int         `json:"crypto_quote_id"` // Optionally associate crypto quote
	CryptoQuote   CryptoQuote `json:"crypto_quote"`
	CryptoTradeID int         `json:"crypto_trade_id"` // Optionally associate crypto trade
	CryptoTrade   CryptoTrade `json:"crypto_trade"`
	Expiration    uint64      `json:"expiration"`
	Strike        float32     `json:"strike"`
	IsPut         bool        `json:"is_put"`
}

func (co CryptoOption) fmt() string {
	return fmt.Sprintf("expiration=%d, strike=%.2f, is_put=%v", co.Expiration, co.Strike, co.IsPut)
}

func (co CryptoOption) fmtFull() string {
	return fmt.Sprintf("%s, %s", co.Crypto.fmtFull(), co.fmt())
}

func (co CryptoOption) Log() string {
	return fmt.Sprintf("a cryptoasset option: %s", co.fmt())
}

func (co CryptoOption) LogFull() string {
	return fmt.Sprintf("a cryptoasset option: %s", co.fmtFull())
}
