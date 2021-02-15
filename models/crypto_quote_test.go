package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestNewCryptoQuoteModel(t *testing.T) {
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

	err = db.AutoMigrate(&CryptoQuote{})
	if err != nil {
		t.Error(err)
		return
	}

	crypto := &Crypto{
		Symbol: testSymbol,
		Hot:    testHot,
	}
	db.Create(crypto)

	cryptoQuote := &CryptoQuote{
		Crypto:    *crypto,
		Timestamp: testTimestamp,
	}
	db.Create(cryptoQuote)

	var cryptoQuote2 CryptoQuote
	result := db.First(&cryptoQuote2)
	if result.RowsAffected != 1 {
		t.Errorf("found %d rows in database instead of 1", result.RowsAffected)
	}
	if result.Error != nil {
		t.Error(err)
		return
	}

	if cryptoQuote2.Timestamp != testTimestamp {
		t.Error(testErr(cryptoQuote2.Timestamp, testTimestamp, "timestamp"))
		return
	}

	t.Log("All tests passed!")
}

func TestNewCryptoQuoteLevelModel(t *testing.T) {
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

	err = db.AutoMigrate(&CryptoQuote{})
	if err != nil {
		t.Error(err)
		return
	}

	err = db.AutoMigrate(&CryptoQuoteLevel{})
	if err != nil {
		t.Error(err)
		return
	}

	crypto := &Crypto{
		Symbol: testSymbol,
		Hot:    testHot,
	}
	db.Create(crypto)

	cryptoQuote := CryptoQuote{
		Crypto:    *crypto,
		Timestamp: testTimestamp,
	}
	db.Create(cryptoQuote)

	cryptoQuoteLevel := CryptoQuoteLevel{
		CryptoQuote: cryptoQuote,
		Level:       testLevel,
		IsBid:       testIsBid,
		Price:       testPrice,
		Size:        testSize,
	}
	db.Create(cryptoQuoteLevel)

	var cryptoQuoteLevel2 CryptoQuoteLevel
	result := db.First(&cryptoQuoteLevel2)
	if result.RowsAffected != 1 {
		t.Errorf("found %d rows in database instead of 1", result.RowsAffected)
	}
	if result.Error != nil {
		t.Error(err)
		return
	}

	if cryptoQuoteLevel2.IsBid != testIsBid {
		t.Error(testErr(cryptoQuoteLevel2.IsBid, testIsBid, "is_bid"))
		return
	}
	if cryptoQuoteLevel2.Price != testPrice {
		t.Error(testErr(cryptoQuoteLevel2.Price, testPrice, "price"))
		return
	}
	if cryptoQuoteLevel2.Size != testSize {
		t.Error(testErr(cryptoQuoteLevel2.Size, testSize, "size"))
		return
	}

	t.Log("All tests passed!")
}
