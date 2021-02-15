package schema

import (
	"github.com/dougrizo/schema/models"
	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

func Migrate(db gorm.DB, l hclog.Logger) error {
	err := db.AutoMigrate(
		&models.Crypto{},
		&models.CryptoOption{},
		&models.CryptoOptionObservation{},
		&models.CryptoOptionQuote{},
		&models.CryptoOptionTransaction{},
		&models.CryptoTrade{},
		&models.CryptoTransaction{},
		&models.Option{},
		&models.OptionObservation{},
		&models.OptionTransaction{},
		&models.Portfolio{},
		&models.Stock{},
		&models.StockQuote{},
		&models.StockTrade{},
		&models.StockTransaction{},
		&models.Portfolio{},
		&models.PortfolioSnapshot{},
		&models.CryptoPortfolioComponent{},
		&models.CryptoOptionPortfolioComponent{},
		&models.StockPortfolioComponent{},
		&models.OptionPortfolioComponent{},
	)

	if err != nil {
		l.Error("Failed to perform migrations")
		return err
	}
	l.Info("Successfully performed migrations")
	return nil
}
