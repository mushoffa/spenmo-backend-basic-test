package domain

import "user-service/domain/entity"

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type UserRepository interface {
	Initialize()
	Create(*entity.User) error
	FindByID(string) (*entity.User, error)
	FindByPhoneNumber(string) (*entity.User, error)
}
