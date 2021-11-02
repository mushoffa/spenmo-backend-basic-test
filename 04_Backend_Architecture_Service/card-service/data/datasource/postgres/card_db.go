package postgres

import (
	"card-service/config"
	"card-service/data/datasource"
	"card-service/data/model"

	"github.com/mushoffa/go-library/database"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type CardDB struct {
	database.Database
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func NewCardDB(cfg *config.Config) (datasource.Database, error) {
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

	cardDB := &CardDB{db}
	return cardDB, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (db *CardDB) InitializeDatabase() {
	db.AutoMigrate(model.CardDB{})
}
