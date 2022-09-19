package controller

import (
	"github.com/gdcolmenares/GinTutorial/request"
	"github.com/gdcolmenares/GinTutorial/response"
	"github.com/gdcolmenares/GinTutorial/service"
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	MakeTransaction(ctx *gin.Context) response.TransactionResponse
}

type controller struct {
	service service.TransactionService
}

func New(service service.TransactionService) TransactionController {
	return &controller{
		service: service,
	}
}

func (c *controller) MakeTransaction(ctx *gin.Context) response.TransactionResponse {
	var request request.TransactionRequest
	ctx.BindJSON(&request)
	var response response.TransactionResponse = c.service.MakeTransaction(request)
	return response
}
