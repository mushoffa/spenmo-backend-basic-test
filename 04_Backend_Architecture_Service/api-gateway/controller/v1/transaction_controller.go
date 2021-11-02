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
type TransactionController interface {
	Purchase(*gin.Context)
	Topup(*gin.Context)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type transactionController struct {
	client protos.TransactionServiceClient
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewTransactionController() TransactionController {
	return &transactionController{}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 02/11/2021
func (c *transactionController) Purchase(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	meta := model.Meta{}
	request := model.PurchaseRequest{}
	response := model.BaseApiResponse{Meta: meta.GetDefaultError()}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.client.Purchase(context.Background(), &protos.PurchaseRequest{
	
	})
	if err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = res
	meta.Message = "Success"
	response.Meta = meta.GetDefaultSuccess()
	ctx.JSON(http.StatusOK, response)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 02/11/2021
func (c *transactionController) Topup(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	meta := model.Meta{}
	request := model.TopupRequest{}
	response := model.BaseApiResponse{Meta: meta.GetDefaultError()}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.client.Topup(context.Background(), &protos.TopupRequest{
	
	})
	if err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = res
	meta.Message = "Success"
	response.Meta = meta.GetDefaultSuccess()
	ctx.JSON(http.StatusOK, response)
}
