package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Option struct {
	gorm.Model     `json:"-"`
	ID             int     `gorm:"primaryKey" json:"id"`
	StockID        int     `gorm:"index:idx_option,unique" json:"stock_id"`
	Stock          Stock   `json:"stock"`
	Expiration     uint64  `gorm:"index:idx_option,unique" json:"expiration"`
	Strike         float32 `gorm:"index:idx_option,unique" json:"strike"`
	IsPut          bool    `gorm:"index:idx_option,unique" json:"is_put"`
	ExpirationType string  `json:"expiration_type"`
	NonStandard    bool    `json:"non_standard"`
	Mini           bool    `json:"mini"`
}

func (o Option) fmt() string {
	return fmt.Sprintf("expiration=%d, strike=%.2f, is_put=%v, expiration_type=%s, non_standard=%v, mini=%v",
		o.Expiration, o.Strike, o.IsPut, o.ExpirationType, o.NonStandard, o.Mini)
}

func (o Option) fmtFull() string {
	return fmt.Sprintf("%s, %s", o.Stock.fmtFull(), o.fmt())
}

func (o Option) Log() string {
	return fmt.Sprintf("an option: %s", o.fmt())
}

func (o Option) LogFull() string {
	return fmt.Sprintf("an option: %s", o.fmtFull())
}
