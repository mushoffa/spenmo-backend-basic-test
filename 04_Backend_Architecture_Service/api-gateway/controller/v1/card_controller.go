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
// @Updated
type CardController interface {
	Create(*gin.Context)
	GetByID(*gin.Context)
	GetAll(*gin.Context)
	Link(*gin.Context)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 02/11/2021
type cardController struct {
	client protos.CardServiceClient
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 02/11/2021
func NewCardController(client protos.CardServiceClient) CardController {
	return &cardController{client}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 
func (c *cardController) Create(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	meta := model.Meta{}
	request := model.CreateCardRequest{}
	response := model.BaseApiResponse{Meta: meta.GetDefaultError()}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.client.CreateCard(context.Background(), &protos.CreateCardRequest {
		Pan: request.PAN,
		ExpiryDate: request.ExpiryDate,
		CardType: request.CardType,
		LimitDaily: request.LimitDaily,
		LimitMonthly: request.LimitMonthly,
	})
	if err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = res.Card
	meta.Message = "Success"
	response.Meta = meta.GetDefaultSuccess()
	ctx.JSON(http.StatusOK, response)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 
func (c *cardController) GetByID(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	meta := model.Meta{}
	response := model.BaseApiResponse{Meta: meta.GetDefaultError()}

	id := ctx.Param("id")

	res, err := c.client.InquiryCard(context.Background(), &protos.InquiryCardRequest{
		Id: id,
	})
	if err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = res.Card
	meta.Message = "Success"
	response.Meta = meta.GetDefaultSuccess()
	ctx.JSON(http.StatusOK, response)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 
func (c *cardController) GetAll(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	meta := model.Meta{}
	response := model.BaseApiResponse{Meta: meta.GetDefaultError()}

	res, err := c.client.GetAllCards(context.Background(), &empty.Empty{})
	if err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = res.Cards
	meta.Message = "Success"
	response.Meta = meta.GetDefaultSuccess()
	ctx.JSON(http.StatusOK, response)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 
func (c *cardController) Link(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	meta := model.Meta{}
	request := model.LinkCardRequest{}
	response := model.BaseApiResponse{Meta: meta.GetDefaultError()}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.client.LinkCard(context.Background(), &protos.LinkCardRequest {
		Id: request.CardID,
		WalletId: request.WalletID,
		UserId: request.UserID,
	})
	if err != nil {
		response.Meta.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = res.Card
	meta.Message = "Success"
	response.Meta = meta.GetDefaultSuccess()
	ctx.JSON(http.StatusOK, response)
}