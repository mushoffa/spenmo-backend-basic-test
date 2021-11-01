package usecase

import (
	"errors"
	// "fmt"

	"user-service/domain/entity"
	domain "user-service/domain/repository"

	"github.com/jinzhu/gorm"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type UserUsecase interface {
	RegisterUser(*entity.User) error
	InquiryUser(string) (*entity.User, error)
	IsUserExist(string) (bool, error)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type user struct {
	r domain.UserRepository
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewUserUsecase(r domain.UserRepository) UserUsecase {
	return &user{r}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (u *user) RegisterUser(user *entity.User) error {

	if err := user.Validate(); err != nil {
		return err
	}

	isExist, err := u.IsUserExist(user.PhoneNumber)
	if err != nil {
		return err
	}

	if isExist {
		return entity.ErrUserIsExist
	}

	if err := u.r.Create(user); err != nil {
		return err
	}

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (u *user) InquiryUser(phoneNumber string) (*entity.User, error) {
	return nil, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (u *user) IsUserExist(phoneNumber string) (bool, error) {
	_, err := u.r.FindByPhoneNumber(phoneNumber)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}