package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestNewCryptoModel(t *testing.T) {
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

	crypto := Crypto{
		Symbol: testSymbol,
		Hot:    testHot,
	}
	db.Create(crypto)

	var crypto2 Crypto
	result := db.First(&crypto2)
	if result.RowsAffected != 1 {
		t.Errorf("found %d rows in database instead of 1", result.RowsAffected)
		return
	}
	if result.Error != nil {
		t.Error(err)
		return
	}

	if crypto2.Symbol != testSymbol {
		t.Error(testErr(crypto2.Symbol, testSymbol, "symbol"))
		return
	}
	if crypto2.Hot != testHot {
		t.Error(testErr(crypto2.Hot, testHot, "hot"))
		return
	}

	t.Log("All tests passed!")
}
