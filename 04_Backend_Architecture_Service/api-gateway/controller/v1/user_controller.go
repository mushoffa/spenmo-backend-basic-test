package v1

import (
	"context"
	"net/http"

	"api-gateway/data/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mushoffa/spenmo-proto/protos"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 03/11/2021
type UserController interface {
	Register(*gin.Context)
	GetAll(*gin.Context)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 02/11/2021
type userController struct {
	client protos.UserServiceClient
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 02/11/2021
func NewUserController(client protos.UserServiceClient) UserController {
	return &userController{client}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 02/11/2021
func (c *userController) Register(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	meta := model.Meta{}
	request := model.RegisterUserRequest{}
	response := model.BaseApiResponse{Meta: meta.GetDefaultError()}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.client.RegisterUser(context.Background(), &protos.RegisterUserRequest{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Dob:         request.DOB,
	})
	if err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = res.User
	meta.Message = "Success"
	response.Meta = meta.GetDefaultSuccess()
	ctx.JSON(http.StatusOK, response)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (c *userController) GetAll(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	meta := model.Meta{}
	response := model.BaseApiResponse{Meta: meta.GetDefaultError()}

	res, err := c.client.GetAllUsers(context.Background(), &empty.Empty{})
	if err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = res.Users
	meta.Message = "Success"
	response.Meta = meta.GetDefaultSuccess()
	ctx.JSON(http.StatusOK, response)
}
