package handlers

import (
	"fmt"
	// "net/http"

	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
	// "github.com/Richd0tcom/sturdy-robot/internal/handlers/requests"
	// "github.com/Richd0tcom/sturdy-robot/internal/service"
	"github.com/Richd0tcom/sturdy-robot/internal/utils"

	// "github.com/Richd0tcom/sturdy-robot/internal/service"
	"github.com/gin-gonic/gin"
)

func (s *Server) SetupProductHandler(r *gin.RouterGroup) {
	r = r.Group("/products")

	r.GET("/hello", func(ctx *gin.Context) {
		fmt.Println("dead last")
		ctx.JSON(200, map[string]interface{}{
			"message": "red",
		})
	}) //helo world handlerd
	// r.POST("/", s.CreateInvoice)
	// r.PATCH("/:id", s.UpdateInvoice)
	r.GET("/", s.FetchProducts)
	// r.GET("/analytics", s.GetAnalytics)
	// r.POST("/reminder", s.SetReminder)
	// r.GET("payment-info", s.GetPaymentInfo)
	// r.GET("/activity", s.GetInvoiceActivityLog)
	// r.GET("/invoice/:id/items", s.GetInvoiceItems)
	// r.GET("/:id", s.GetInvoice)
	

}

// func (s *Server) CreateProduct(ctx *gin.Context) {

// 	var req  requests.CreateProduct

// 	err:= ctx.ShouldBindJSON(&req)

// 	if err != nil {
// 		ctx.JSON(400, gin.H{})
// 		return
// 	}
// 	p, err:= s.store.CreateProduct(ctx, db.CreateProductParams{
// 		CategoryID: utils.ParseUUID(req.CategoryID),
// 		BranchID: utils.ParseUUID(req.BranchID),
// 		Name:

// 	})
// }

func (s *Server) FetchProducts(c *gin.Context) {

	branch_id, ok:=c.Params.Get("branch_id")

	fmt.Println("I was called")
	if !ok {
		c.JSON(400, gin.H{})
		return
	}
	p, err:= s.store.GetProductsByBranchID(c, utils.ParseUUID(branch_id))

	if err!=nil {
		c.JSON(200, []db.Product{})
	}
	c.JSON(200, p)
}