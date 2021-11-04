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

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 03/11/2021
func (r *Router) initializeV1Router() {
	v1 := r.Gin.Group("/v1")

	cards := v1.Group("/cards")
	cards.POST("/", r.v1.Card.Create)
	cards.GET("/:id", r.v1.Card.GetByID)
	cards.GET("/", r.v1.Card.GetAll)
	cards.POST("/link", r.v1.Card.Link)

	wallets := v1.Group("/wallets")
	wallets.POST("/", r.v1.Wallet.Create)
	wallets.GET("/", r.v1.Wallet.GetAll)
	wallets.GET("/user/:id", r.v1.Wallet.GetByAccountNumber)
	wallets.GET("/:id", r.v1.Wallet.GetByID)

	users := v1.Group("/users")
	users.POST("/", r.v1.User.Register)
	users.GET("/", r.v1.User.GetAll)

	transactions := v1.Group("/transactions")
	transactions.POST("/", r.v1.Transaction.Create)
	transactions.POST("/purchase", r.v1.Transaction.Purchase)
	transactions.POST("/topup", r.v1.Transaction.Topup)
}
