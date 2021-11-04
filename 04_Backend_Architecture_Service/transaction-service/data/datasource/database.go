package datasource

import (
	"github.com/mushoffa/go-library/database"
)

type Database interface {
	database.Database
	InitializeDatabase()
}
