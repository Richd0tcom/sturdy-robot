package handlers

import (
	"fmt"

	"github.com/Richd0tcom/sturdy-robot/internal/service"
	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"

	// "github.com/Richd0tcom/sturdy-robot/internal/service"
	"github.com/gin-gonic/gin"
)


func (s *Server) SetupInvoiceHandler()  {
	r:= s.serverRouter.Group("/")

	r.GET("/hello", func(ctx *gin.Context) {
		fmt.Println("dead last")
		ctx.JSON(200, map[string]interface{}{
			"message": "red",
		})
	}) //helo world handlerd
	r.GET("/invoice", s.CreateInvoice)
}

func (s *Server) CreateInvoice(ctx *gin.Context){
	var args db.CreateInvoiceParams
	ctx.ShouldBindJSON(&args)
	
	
	invoice, err := service.CreateNewInvoice(ctx, args, s.store)
	if err!=nil{
		fmt.Println("error")
		return 
	}

	ctx.JSON(200, map[string]interface{}{
		"message": invoice,
	})

}
// func (s *Server) UpdateInvoice(ctx *gin.Context)
