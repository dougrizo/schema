package models

import (
	"gorm.io/gorm"
)

type Crypto struct {
	gorm.Model `json:"-"`
	ID         int    `gorm:"primaryKey" json:"id"`
	Symbol     string `json:"symbol"`
	Hot        bool   `json:"hot"`
}
