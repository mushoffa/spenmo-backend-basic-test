package service

import "user-service/domain/usecase"

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type user struct {
	u usecase.UserUsecase
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewUserService(u usecase.UserUsecase) {

}
