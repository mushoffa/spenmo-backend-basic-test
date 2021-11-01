package router

import (
	v1 "api-gateway/controller/v1"

	"github.com/gin-gonic/gin"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type Router struct {
	Gin *gin.Engine
	v1  *v1.V1Controller
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewRouter(v1 *v1.V1Controller) *Router {
	return &Router{
		Gin: gin.Default(),
		v1:  v1,
	}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (r *Router) InitializeRouter() {
	r.initializeV1Router()
}

func (r *Router) initializeV1Router() {
	v1 := r.Gin.Group("/v1")

	cards := v1.Group("/cards")
	wallets := v1.Group("/wallets")
	users := v1.Group("/users")
	transactions := v1.Group("/transactions")
}
