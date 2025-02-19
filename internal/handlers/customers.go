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

func (s *Server) SetupCustomerHandler(r *gin.RouterGroup) {

	r = r.Group("/customers")

	r.GET("/hello", func(ctx *gin.Context) {
		fmt.Println("dead last")
		ctx.JSON(200, map[string]interface{}{
			"message": "red",
		})
	}) //helo world handlerd
	// r.POST("/", s.CreateInvoice)
	// r.PATCH("/:id", s.UpdateInvoice)
	r.GET("/", s.GetCustomers)
	// r.GET("/analytics", s.GetAnalytics)
	// r.POST("/reminder", s.SetReminder)
	// r.GET("payment-info", s.GetPaymentInfo)
	// r.GET("/activity", s.GetInvoiceActivityLog)
	// r.GET("/invoice/:id/items", s.GetInvoiceItems)
	// r.GET("/:id", s.GetInvoice)
	

}

func (s *Server) GetCustomers(c *gin.Context) {
	branch_id, ok:=c.Params.Get("branch_id")

	if !ok {
		c.JSON(400, gin.H{})
		return
	}
	customers, err:=s.store.GetCustomersByBranch(c, utils.ParseUUID(branch_id))
	if err!=nil {
		c.JSON(200, []db.Customer{})
	}
	c.JSON(200, customers)
}