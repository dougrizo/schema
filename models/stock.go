package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	ID     int    `gorm:"primaryKey"`
	Symbol string `json:"symbol"`
	Hot    bool   `json:"hot"`
}
