package schema

import (
	"github.com/dougrizo/schema/models"
	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

func Migrate(db gorm.DB, l hclog.Logger) error {
	err := db.AutoMigrate(
		&models.Crypto{},
		&models.CryptoTrade{},
		&models.CryptoQuote{},
		&models.CryptoQuoteLevel{},
		&models.CryptoTransaction{},
		&models.CryptoPortfolioComponent{},

		&models.CryptoOption{},
		&models.CryptoOptionObservation{},
		&models.CryptoOptionQuote{},
		&models.CryptoOptionQuoteLevel{},
		&models.CryptoOptionTransaction{},
		&models.CryptoOptionPortfolioComponent{},

		&models.Option{},
		&models.OptionObservation{},
		&models.OptionTransaction{},
		&models.OptionPortfolioComponent{},

		&models.Stock{},
		&models.StockQuote{},
		&models.StockTrade{},
		&models.StockTransaction{},
		&models.StockPortfolioComponent{},

		&models.Portfolio{},
		&models.PortfolioSnapshot{},
	)

	if err != nil {
		l.Error("Failed to perform migrations")
		return err
	}
	l.Info("Successfully performed migrations")
	return nil
}
