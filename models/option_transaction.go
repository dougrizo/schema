package models

import "gorm.io/gorm"

type OptionTransaction struct {
	gorm.Model          `json:"-"`
	ID                  int     `gorm:"primaryKey" json:"id"`
	OptionID            int     `json:"option_id"`
	Option              Option  `json:"option"`
	Timestamp           uint64  `json:"timestamp"`
	Price               float64 `json:"price"`
	Size                int64   `json:"volume"`
	MetaLastUpdated     uint64  `json:"meta_last_updated"`
	MetaLastUpdatedUser string  `json:"meta_last_updated_user"`
}
