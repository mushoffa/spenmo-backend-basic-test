package repository

import (
	"user-service/data/datasource"
	"user-service/domain/entity"
	domain "user-service/domain/repository"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type user struct {
	db datasource.Database
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewUserRepository(db datasource.Database) domain.UserRepository {
	return &user{db}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (r *user) Initialize() {
	r.db.InitializeDatabase()
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (r *user) Create(user *entity.User) error {
	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (r *user) FindByID(id string) (*entity.User, error) {
	return nil, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (r *user) FindByPhoneNumber(phoneNumber string) (*entity.User, error) {
	return nil, nil
}
