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
// @Updated 03/11/2021
type TransactionController interface {
	Create(*gin.Context)
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
// @Updated 03/11/2021
func NewTransactionController(client protos.TransactionServiceClient) TransactionController {
	return &transactionController{client}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (c *transactionController) Create(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	meta := model.Meta{}
	request := model.CreateBillRequest{}
	response := model.BaseApiResponse{Meta: meta.GetDefaultError()}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// res, err := c.client.Purchase(context.Background(), &protos.PurchaseRequest{})
	// if err != nil {
	// 	response.Meta.Message = err.Error()
	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	response.Data = "res"
	meta.Message = "Success"
	response.Meta = meta.GetDefaultSuccess()
	ctx.JSON(http.StatusOK, response)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 04/11/2021
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
		Pan: request.PAN,
		Amount: request.Amount,
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
		AccountNumber: request.AccountNumber,
		WalletId:      request.WalletID,
		Amount:        request.Amount,
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
