package service

import (
	"context"

	"user-service/domain/entity"
	"user-service/domain/usecase"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mushoffa/spenmo-proto/protos"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type user struct {
	protos.UnimplementedUserServiceServer
	u usecase.UserUsecase
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewUserService(u usecase.UserUsecase) protos.UserServiceServer {
	return &user{u: u}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (s *user) RegisterUser(ctx context.Context, request *protos.RegisterUserRequest) (*protos.RegisterUserResponse, error) {
	user := entity.User{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		DOB:         request.Dob,
	}

	if err := s.u.RegisterUser(&user); err != nil {
		return nil, err
	}

	response := protos.RegisterUserResponse{
		User: &protos.User{
			Id:          user.ID,
			Created:     user.Created.String(),
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
			Dob:         user.DOB,
		},
	}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (s *user) GetAllUsers(ctx context.Context, request *empty.Empty) (*protos.GetAllUsersResponse, error) {
	_users, err := s.u.GetAllUsers()
	if err != nil {
		return nil, err
	}

	users := []*protos.User{}

	for _, _user := range _users {
		user := &protos.User{
			Id:          _user.ID,
			Created:     _user.Created.String(),
			Name:        _user.Name,
			PhoneNumber: _user.PhoneNumber,
			Email:       _user.Email,
			Dob:         _user.DOB,
		}

		users = append(users, user)
	}

	response := protos.GetAllUsersResponse{
		Users: users,
	}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (s *user) InquiryUser(ctx context.Context, request *protos.InquiryUserRequest) (*protos.InquiryUserResponse, error) {
	user, err := s.u.InquiryUser(request.PhoneNumber)
	if err != nil {
		return nil, err
	}

	response := protos.InquiryUserResponse{
		User: &protos.User{
			Id:          user.ID,
			Created:     user.Created.String(),
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
			Dob:         user.DOB,
		},
	}

	return &response, nil
}
