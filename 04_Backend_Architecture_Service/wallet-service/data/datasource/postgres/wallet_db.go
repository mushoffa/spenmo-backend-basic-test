package postgres

import (
	"wallet-service/config"
	"wallet-service/data/datasource"
	"wallet-service/data/model"

	"github.com/mushoffa/go-library/database"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type WalletDB struct {
	database.Database
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewWalletDB(cfg *config.Config) (datasource.Database, error) {
	db, err := database.NewPostgres(
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		false,
	)

	if err != nil {
		return nil, err
	}

	walletDB := &WalletDB{db}
	return walletDB, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (db *WalletDB) InitializeDatabase() {
	db.AutoMigrate(model.WalletDB{})
	// db.AutoMigrate(model.UserDB{})
}
