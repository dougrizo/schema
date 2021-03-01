package models

import (
	"fmt"
	"gorm.io/gorm"
)

type CryptoTrade struct {
	gorm.Model  `json:"-"`
	ID          int     `gorm:"primaryKey" json:"-"`
	CryptoID    int     `gorm:"index:idx_crypto_trade,unique" json:"crypto_id"`
	Crypto      Crypto  `json:"crypto"`
	Timestamp   uint64  `gorm:"index:idx_crypto_trade,unique" json:"timestamp"`
	Price       float32 `json:"price"`
	Size        int32   `json:"size"`
	MarketMaker bool    `json:"market_maker"`
}

func (ct CryptoTrade) fmt() string {
	return fmt.Sprintf("timestamp=%d, price=%.2f, size=%d, market_maker=%v",
		ct.Timestamp, ct.Price, ct.Size, ct.MarketMaker)
}

func (ct CryptoTrade) fmtFull() string {
	return fmt.Sprintf("%s, %s", ct.Crypto.fmtFull(), ct.fmt())
}

func (ct CryptoTrade) Log() string {
	return fmt.Sprintf("a cryptoasset quote: %s", ct.fmt())
}

func (ct CryptoTrade) LogFull() string {
	return fmt.Sprintf("a cryptoasset quote: %s", ct.fmtFull())
}
