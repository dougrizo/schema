package models

import (
	"gorm.io/gorm"
)

type Option struct {
	gorm.Model     `json:"-"`
	ID             int     `gorm:"primaryKey" json:"id"`
	StockID        int     `json:"stock_id"`
	Stock          Stock   `json:"stock"`
	Expiration     uint64  `json:"expiration"`
	Strike         float64 `json:"strike"`
	IsPut          bool    `json:"is_put"`
	ExpirationType string  `json:"expiration_type"`
	NonStandard    bool    `json:"non_standard"`
	Mini           bool    `json:"mini"`
}
