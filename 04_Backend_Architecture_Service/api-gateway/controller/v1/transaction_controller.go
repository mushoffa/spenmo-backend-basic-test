package v1

import (
	"context"

	"github.com/gin-gonic/gin"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type TransactionController interface {
	Purchase(*gin.Context)
	Topup(*gin.Context)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type transactionController struct {
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewTransactionController() TransactionController {
	return &transactionController{}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (c *transactionController) Purchase(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (c *transactionController) Topup(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()
}
