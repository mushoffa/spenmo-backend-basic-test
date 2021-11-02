package v1

import (
	"context"

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
// @Updated
func (c *userController) Register(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()
}
