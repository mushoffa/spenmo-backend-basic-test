package v1

import (
	"context"

	"github.com/gin-gonic/gin"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type UserController interface {
	Register(*gin.Context)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type userController struct {
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewUserController() UserController {
	return &userController{}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (c *userController) Register(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()
}
