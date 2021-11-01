package postgres

import (
	"user-service/config"
	"user-service/data/datasource"
	"user-service/data/model"

	"github.com/mushoffa/go-library/database"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type UserDB struct {
	database.Database
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewUserDB(cfg *config.Config) (datasource.Database, error) {
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

	userDB := &UserDB{db}
	return userDB, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (db *UserDB) InitializeDatabase() {
	db.AutoMigrate(model.UserDB{})
}
