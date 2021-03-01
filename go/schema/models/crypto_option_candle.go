package models

import (
	"fmt"
	"gorm.io/gorm"
)

type CryptoOptionCandle struct {
	gorm.Model     `json:"-"`
	ID             int          `gorm:"primaryKey" json:"id"`
	CryptoOptionID int          `gorm:"index:idx_crypto_option_candle,unique" json:"crypto_option_id"`
	CryptoOption   CryptoOption `json:"crypto_option"`
	Interval       string       `json:"interval"`
	OpenTime       uint64       `gorm:"index:idx_crypto_option_candle,unique" json:"open_time"`
	CloseTime      uint64       `json:"close_time"`
	Open           float32      `json:"open"`
	High           float32      `json:"high"`
	Low            float32      `json:"low"`
	Close          float32      `json:"close"`
	Volume         int32        `json:"volume"`
}

func (co CryptoOptionCandle) fmt() string {
	return fmt.Sprintf(
		"open_time=%d, close_time=%d, open=%.2f, high=%.2f, low=%.2f, close=%.2f, volume=%d",
		co.OpenTime, co.CloseTime, co.Open, co.High, co.Low, co.Close, co.Volume)
}

func (co CryptoOptionCandle) fmtFull() string {
	return fmt.Sprintf("%s, %s", co.CryptoOption.fmtFull(), co.fmt())
}

func (co CryptoOptionCandle) Log() string {
	return fmt.Sprintf("a cryptoasset option observation: %s", co.fmt())
}

func (co CryptoOptionCandle) LogFull() string {
	return fmt.Sprintf("a cryptoasset option observation: %s", co.fmtFull())
}
