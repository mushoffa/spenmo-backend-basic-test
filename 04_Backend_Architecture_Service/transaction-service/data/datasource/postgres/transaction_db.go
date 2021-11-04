package postgres

import (
	"transaction-service/config"
	"transaction-service/data/datasource"
	"transaction-service/data/model"

	"github.com/mushoffa/go-library/database"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
type TransactionDB struct {
	database.Database
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func NewTransactionDB(cfg *config.Config) (datasource.Database, error) {
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

	transactionDB := &TransactionDB{db}
	return transactionDB, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (db *TransactionDB) InitializeDatabase() {
	db.AutoMigrate(model.TransactionDB{})
	db.AutoMigrate(model.LimitDB{})
}
