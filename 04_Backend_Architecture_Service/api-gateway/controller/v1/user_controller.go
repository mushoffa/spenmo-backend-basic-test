package v1

import (
	"context"
	"net/http"

	"api-gateway/data/model"

	"github.com/gin-gonic/gin"
	"github.com/mushoffa/spenmo-proto/protos"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type UserController interface {
	Register(*gin.Context)
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

	res, err := c.client.RegisterUser(context.Background(), &protos.RegisterUserRequest {
		Name: request.Name,
		PhoneNumber: request.PhoneNumber,
		Email: request.Email,
		Dob: request.DOB,
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
