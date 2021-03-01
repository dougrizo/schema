package models

import (
	"fmt"
	"gorm.io/gorm"
)

type OptionTransaction struct {
	gorm.Model `json:"-"`
	ID         int     `gorm:"primaryKey" json:"id"`
	OptionID   int     `gorm:"index:idx_option_transaction,unique" json:"option_id"`
	Option     Option  `json:"option"`
	Timestamp  uint64  `gorm:"index:idx_option_transaction,unique" json:"timestamp"`
	Price      float32 `json:"price"`
	Size       int32   `json:"volume"`
}

func (ot OptionTransaction) fmt() string {
	return fmt.Sprintf("timestamp=%d, price=%.2f, size=%d", ot.Timestamp, ot.Price, ot.Size)
}

func (ot OptionTransaction) fmtFull() string {
	return fmt.Sprintf("%s, %s", ot.Option.fmtFull(), ot.fmt())
}

func (ot OptionTransaction) Log() string {
	return fmt.Sprintf("an option transaction: %s", ot.fmt())
}

func (ot OptionTransaction) LogFull() string {
	return fmt.Sprintf("an option transaction: %s", ot.fmtFull())
}
