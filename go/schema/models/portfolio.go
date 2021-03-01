package models

import "gorm.io/gorm"

type Portfolio struct {
	gorm.Model  `json:"-"`
	ID          int   `json:"id" gorm:"primaryKey"`
	Value       int64 `json:"value"`
	BuyingPower int64 `json:"buying_power"`
}

type PortfolioSnapshot struct {
	gorm.Model  `json:"-"`
	ID          int       `json:"id" gorm:"primaryKey"`
	PortfolioID int       `json:"portfolio_id"`
	Portfolio   Portfolio `json:"portfolio"`
	Value       int64     `json:"value"`
	BuyingPower int64     `json:"buying_power"`
	Timestamp   uint64    `json:"timestamp"`
}

type CryptoPortfolioComponent struct {
	gorm.Model      `json:"-"`
	ID              int       `json:"id" gorm:"primaryKey"`
	PortfolioID     int       `json:"portfolio_id"`
	Portfolio       Portfolio `json:"portfolio"`
	CryptoID        int       `json:"crypto_id"`
	Crypto          Crypto    `json:"crypto"`
	Price           float64   `json:"price"`
	Size            uint64    `json:"size"`
	UsedBuyingPower float64   `json:"used_buying_power"`
}

type CryptoOptionPortfolioComponent struct {
	gorm.Model      `json:"-"`
	ID              int          `json:"id" gorm:"primaryKey"`
	PortfolioID     int          `json:"portfolio_id"`
	Portfolio       Portfolio    `json:"portfolio"`
	CryptoOptionID  int          `json:"crypto_option_id"`
	CryptoOption    CryptoOption `json:"crypto_option"`
	Price           float64      `json:"price"`
	Size            uint64       `json:"size"`
	UsedBuyingPower float64      `json:"used_buying_power"`
}

type StockPortfolioComponent struct {
	gorm.Model      `json:"-"`
	ID              int       `json:"id" gorm:"primaryKey"`
	PortfolioID     int       `json:"portfolio_id"`
	Portfolio       Portfolio `json:"portfolio"`
	StockID         int       `json:"stock_id"`
	Stock           `json:"stock"`
	Price           float64 `json:"price"`
	Size            uint64  `json:"size"`
	UsedBuyingPower float64 `json:"used_buying_power"`
}

type OptionPortfolioComponent struct {
	gorm.Model      `json:"-"`
	ID              int       `json:"id" gorm:"primaryKey"`
	PortfolioID     int       `json:"portfolio_id"`
	Portfolio       Portfolio `json:"portfolio"`
	OptionID        int       `json:"option_id"`
	Option          `json:"option"`
	Price           float64 `json:"price"`
	Size            uint64  `json:"size"`
	UsedBuyingPower float64 `json:"used_buying_power"`
}
