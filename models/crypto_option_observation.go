package models

import "gorm.io/gorm"

type CryptoOptionObservation struct {
	gorm.Model     `json:"-"`
	ID             int          `gorm:"primaryKey" json:"id"`
	CryptoOptionID int          `json:"crypto_option_id"`
	CryptoOption   CryptoOption `json:"crypto_option"`
	Interval       string       `json:"interval"`
	OpenTime       uint64       `json:"open_time"`
	CloseTime      uint64       `json:"close_time"`
	Open           float64      `json:"open"`
	High           float64      `json:"high"`
	Low            float64      `json:"low"`
	Close          float64      `json:"close"`
	Volume         uint64       `json:"volume"`
}
