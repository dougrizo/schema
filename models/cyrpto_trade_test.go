package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestNewCryptoTradeModel(t *testing.T) {
	var err error

	db, err := gorm.Open(sqlite.Open(testDatabase), &gorm.Config{})
	if err != nil {
		t.Error(err)
		return
	}

	err = db.AutoMigrate(&Crypto{})
	if err != nil {
		t.Error(err)
		return
	}

	err = db.AutoMigrate(&CryptoTrade{})
	if err != nil {
		t.Error(err)
		return
	}

	crypto := &Crypto{
		Symbol: testSymbol,
		Hot:    testHot,
	}
	db.Create(crypto)

	cryptoTrade := &CryptoTrade{
		CryptoID:    (*crypto).ID,
		Crypto:      *crypto,
		Timestamp:   testTimestamp,
		Price:       testPrice,
		Size:        testSize,
		MarketMaker: testMarketMaker,
	}
	db.Create(cryptoTrade)

	var cryptoTrade2 CryptoTrade
	result := db.First(&cryptoTrade2)
	if result.RowsAffected != 1 {
		t.Errorf("found %d rows in database instead of 1", result.RowsAffected)
	}
	if result.Error != nil {
		t.Error(err)
		return
	}

	if cryptoTrade2.Timestamp != testTimestamp {
		t.Error(testErr(cryptoTrade2.Timestamp, testTimestamp, "timestamp"))
		return
	}
	if cryptoTrade2.Price != testPrice {
		t.Error(testErr(cryptoTrade2.Price, testPrice, "price"))
		return
	}
	if cryptoTrade2.Size != testSize {
		t.Error(testErr(cryptoTrade2.Size, testSize, "size"))
		return
	}
	if cryptoTrade2.MarketMaker != testMarketMaker {
		t.Error(testErr(cryptoTrade2.MarketMaker, testMarketMaker, "market_maker"))
	}

	t.Log("All tests passed!")
}
