package domain

import "user-service/domain/entity"

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 03/11/2021
type UserRepository interface {
	Initialize()
	Create(*entity.User) error
	FindByID(string) (*entity.User, error)
	FindByPhoneNumber(string) (*entity.User, error)
	FindAll() ([]entity.User, error)
}
