package models

import "gorm.io/gorm"

type OptionObservation struct {
	gorm.Model             `json:"-"`
	ID                     int     `gorm:"primaryKey" json:"id"`
	OptionID               int     `json:"option_id"`
	Option                 Option  `json:"option"`
	TradeTimeInLong        uint64  `json:"trade_time_in_long"`
	QuoteTimeInLong        uint64  `json:"quote_time_in_long"`
	Bid                    float64 `json:"bid"`
	Ask                    float64 `json:"ask"`
	Mark                   float64 `json:"mark"`
	BidSize                int64   `json:"bid_size"`
	AskSize                int64   `json:"ask_size"`
	HighPrice              float64 `json:"high_price"`
	LowPrice               float64 `json:"low_price"`
	OpenPrice              float64 `json:"open_price"`
	ClosePrice             float64 `json:"close_price"`
	Volume                 int64   `json:"volume"`
	Volatility             float64 `json:"volatility"`
	Delta                  float64 `json:"delta"`
	Gamma                  float64 `json:"gamma"`
	Theta                  float64 `json:"theta"`
	Vega                   float64 `json:"vega"`
	Rho                    float64 `json:"rho"`
	OpenInterest           int64   `json:"open_interest"`
	TimeValue              float64 `json:"time_value"`
	TheoreticalOptionValue float64 `json:"theoretical_option_value"`
	TheoreticalVolatility  float64 `json:"theoretical_volatility"`
	DaysToExpiration       int64   `json:"days_to_expiration"`
}
