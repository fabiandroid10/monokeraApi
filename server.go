package main

import (
	"github.com/gdcolmenares/GinTutorial/controller"
	"github.com/gdcolmenares/GinTutorial/service"
	"github.com/gin-gonic/gin"
)

var (
	transactionService    service.TransactionService       = service.New()
	transactionController controller.TransactionController = controller.New(transactionService)
)

func main() {
	server := gin.Default()

	server.POST("/purchase", func(ctx *gin.Context) {
		ctx.JSON(200, transactionController.MakeTransaction(ctx))
	})

	server.Run(":80")
}
