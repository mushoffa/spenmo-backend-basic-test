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
// @Updated 02/11/2021
type WalletController interface {
	Create(*gin.Context)
	GetByID(*gin.Context)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type walletController struct {
	client protos.WalletServiceClient
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewWalletController(client protos.WalletServiceClient) WalletController {
	return &walletController{client}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (c *walletController) Create(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	meta := model.Meta{}
	request := model.CreateWalletRequest{}
	response := model.BaseApiResponse{Meta: meta.GetDefaultError()}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.client.CreateWallet(context.Background(), &protos.CreateWalletRequest{
		UserId: request.UserID,
		Name: request.Name,
		MaxLimit: request.MaxLimit,
	})
	if err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = res.Wallet
	meta.Message = "Success"
	response.Meta = meta.GetDefaultSuccess()
	ctx.JSON(http.StatusOK, response)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 
func (c *walletController) GetByID(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	meta := model.Meta{}
	response := model.BaseApiResponse{Meta: meta.GetDefaultError()}

	id := ctx.Param("id")

	res, err := c.client.InquiryBalance(context.Background(), &protos.InquiryBalanceRequest{
		Id: id,
	})
	if err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = res.Wallet
	meta.Message = "Success"
	response.Meta = meta.GetDefaultSuccess()
	ctx.JSON(http.StatusOK, response)
}