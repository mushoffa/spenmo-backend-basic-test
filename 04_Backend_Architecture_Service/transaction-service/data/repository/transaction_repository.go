package repository

import (
	"transaction-service/data/datasource"
	domain "transaction-service/domain/repository"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
type transaction struct {
	db datasource.Database
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func NewTransactionRepository(db datasource.Database) domain.TransactionRepository {
	return &transaction{db}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (r *transaction) Initialize() {
	r.db.InitializeDatabase()
}
