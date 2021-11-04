package repository

import (
	"time"

	"user-service/data/datasource"
	"user-service/data/model"
	"user-service/domain/entity"
	domain "user-service/domain/repository"

	"github.com/google/uuid"
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
// @Updated 02/11/2021
func (r *user) Create(user *entity.User) error {
	id := uuid.New().String()
	created := time.Now()

	userDB := model.UserDB{
		ID:          id,
		Created:     created,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		DOB:         user.DOB,
	}

	if err := r.db.Create(&userDB); err != nil {
		return err
	}

	user.ID = id
	user.Created = created

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (r *user) FindByID(id string) (*entity.User, error) {
	user := entity.User{}

	if err := r.db.FindByID("id", id, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (r *user) FindByPhoneNumber(phoneNumber string) (*entity.User, error) {
	user := entity.User{}

	if err := r.db.FindByID("phone_number", phoneNumber, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (r *user) FindAll() ([]entity.User, error) {
	users := []entity.User{}

	if err := r.db.FindAll(&users); err != nil {
		return nil, err
	}

	return users, nil
}
