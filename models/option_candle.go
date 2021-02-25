package models

import (
	"fmt"
	"gorm.io/gorm"
)

type OptionCandle struct {
	gorm.Model             `json:"-"`
	ID                     int         `gorm:"primaryKey" json:"id"`
	OptionID               int         `gorm:"index:idx_option_candle,unique" json:"option_id"`
	Option                 Option      `json:"option"`
	StockCandleID          int         `json:"stock_candle_id"` // Optionally associate stock candle.
	StockCandle            StockCandle `json:"stock_candle"`
	StockQuoteID           int         `json:"stock_quote_id"` // Optionally associate stock quote
	StockQuote             StockQuote  `json:"stock_quote"`
	StockTradeID           int         `json:"stock_trade_id"` // Optionally associate stock trade
	StockTrade             StockTrade  `json:"stock_trade"`
	TradeTimeInLong        uint64      `gorm:"index:idx_option_candle,unique" json:"trade_time_in_long"`
	QuoteTimeInLong        uint64      `gorm:"index:idx_option_candle,unique" json:"quote_time_in_long"`
	Bid                    float32     `json:"bid"`
	Ask                    float32     `json:"ask"`
	Mark                   float32     `json:"mark"`
	BidSize                int16       `json:"bid_size"`
	AskSize                int16       `json:"ask_size"`
	High                   float32     `json:"high"`
	Low                    float32     `json:"low"`
	Open                   float32     `json:"open"`
	Close                  float32     `json:"close"`
	Volume                 int32       `json:"volume"`
	Volatility             float32     `json:"volatility"`
	Delta                  float32     `json:"delta"`
	Gamma                  float32     `json:"gamma"`
	Theta                  float32     `json:"theta"`
	Vega                   float32     `json:"vega"`
	Rho                    float32     `json:"rho"`
	OpenInterest           int32       `json:"open_interest"`
	TimeValue              float32     `json:"time_value"`
	TheoreticalOptionValue float32     `json:"theoretical_option_value"`
	TheoreticalVolatility  float32     `json:"theoretical_volatility"`
	DaysToExpiration       int32       `json:"days_to_expiration"`
}

func (oc OptionCandle) fmt() string {
	return fmt.Sprintf(
		"trade_time=%d, quote_time=%d, bid=%.2f, ask=%.2f, mark=%.2f, bid_size=%d, ask_size=%d, high=%.2f, low=%.2f, open=%.2f, close=%.2f, volume=%d, volatility=%.2f, delta=%.2f, gamma=%.2f, theta=%.2f, vega=%.2f, rho=%.2f, open_interest=%d, time_value=%.2f, theoretical_option_value=%.2f, theoretical_volatility=%.2f, days_to_expiration=%d",
		oc.TradeTimeInLong, oc.QuoteTimeInLong, oc.Bid, oc.Ask, oc.Mark, oc.BidSize, oc.AskSize, oc.High, oc.Low, oc.Open, oc.Close, oc.Volume, oc.Volatility, oc.Delta, oc.Gamma, oc.Theta, oc.Vega, oc.Rho, oc.OpenInterest, oc.TimeValue, oc.TheoreticalOptionValue, oc.TheoreticalVolatility, oc.DaysToExpiration)
}

func (oc OptionCandle) fmtFull() string {
	return fmt.Sprintf("%s, %s", oc.Option.fmtFull(), oc.fmt())
}

func (oc OptionCandle) Log() string {
	return fmt.Sprintf("an option candle: %s", oc.fmt())
}

func (oc OptionCandle) LogFull() string {
	return fmt.Sprintf("an option candle: %s", oc.fmtFull())
}
