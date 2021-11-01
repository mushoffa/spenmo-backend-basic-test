package usecase

import domain "user-service/domain/repository"

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type UserUsecase interface {
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
