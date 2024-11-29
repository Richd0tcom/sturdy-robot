package handlers

import (
	"fmt"
	"net/http"

	"github.com/Richd0tcom/sturdy-robot/internal/handlers/requests"
	"github.com/Richd0tcom/sturdy-robot/internal/service"
	"github.com/Richd0tcom/sturdy-robot/internal/utils"

	// "github.com/Richd0tcom/sturdy-robot/internal/service"
	"github.com/gin-gonic/gin"
)

func (s *Server) SetupInvoiceHandler() {
	r := s.serverRouter.Group("/invoices")

	r.GET("/hello", func(ctx *gin.Context) {
		fmt.Println("dead last")
		ctx.JSON(200, map[string]interface{}{
			"message": "red",
		})
	}) //helo world handlerd
	r.POST("/", s.CreateInvoice)
	r.PATCH("/:id", s.UpdateInvoice)
	r.GET("/all", s.GetAllInvoices)
	r.GET("/analytics", s.GetAnalytics)
	r.POST("/reminder", s.SetReminder)
	r.GET("payment-info", s.GetPaymentInfo)
	r.GET("/activity", s.GetInvoiceActivityLog)
	r.GET("/:id/items", s.GetInvoiceItems)
	r.GET("/:id", s.GetInvoice)
	

}

func (s *Server) CreateInvoice(ctx *gin.Context) {
	var args requests.CreateInvoiceReq

	claims, err := utils.ExtractTokenIDs(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{})
		return
	}
	userID := claims["id"].(string)
	ctx.ShouldBindJSON(&args)

	invoice, err := service.CreateNewInvoice(ctx, userID, args, s.store)
	if err != nil {
		fmt.Println("error", err)
		ctx.JSON(400, gin.H{})
		return
	}

	ctx.JSON(200, invoice)
}

func (s *Server) UpdateInvoice(ctx *gin.Context) {
	//TODO
}

// see analytics
func (s *Server) GetAnalytics(ctx *gin.Context) {

	claims, err := utils.ExtractTokenIDs(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{})
		return
	}
	userID := claims["id"].(string)

	analytics, err := service.GetAnalytics(ctx, userID, s.store)

	if err != nil {
		ctx.JSON(400, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, analytics)
}

// change invoice settings

// edit/set reminder
func (s *Server) SetReminder(ctx *gin.Context) {
	//TODO:
}

//get invoice

func (s *Server) GetInvoice(ctx *gin.Context) {
	//TODO(Auth): Validate invoice against user_id

	invoice_id, ok := ctx.Params.Get("invoice_id")
	if !ok {
		ctx.JSON(400, gin.H{})
		return
	}
	invoice, err := service.GetInvoice(ctx, invoice_id, s.store)

	if err != nil {
		ctx.JSON(400, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, invoice)

}

func (s *Server) GetAllInvoices(ctx *gin.Context) {

	claims, err := utils.ExtractTokenIDs(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{})
		return
	}
	userID := claims["id"].(string)
	invoices, err := service.GetAllInvoicesByUser(ctx, userID, s.store)

	if err != nil {
		ctx.JSON(400, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, invoices)
}

func (s *Server) GetInvoicesByStatus(ctx *gin.Context) {
	//TODO
}

// get invoice items
func (s *Server) GetInvoiceItems(ctx *gin.Context) {
	//TODO(Auth): Validate invoice against user_id

	invoice_id, ok := ctx.Params.Get("invoice_id")
	if !ok {
		ctx.JSON(400, gin.H{})
		return
	}

	items, err := service.GetInvoiceItems(ctx, invoice_id, s.store)
	if err != nil {
		ctx.JSON(400, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

// get payment info
func (s *Server) GetPaymentInfo(ctx *gin.Context) {
	claims, err := utils.ExtractTokenIDs(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{})
		return
	}
	userID := claims["id"].(string)

	info, err:= service.GetPaymentInfo(ctx, userID, s.store)

	if err != nil {
		ctx.JSON(400, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, info)
}

// get activity log
func (s *Server) GetInvoiceActivityLog(ctx *gin.Context) {
	invoice_id, ok := ctx.Params.Get("invoice_id")
	if !ok {
		ctx.JSON(400, gin.H{})
		return
	}
	logs, err:= service.GetInvoiceActivityLog(ctx, invoice_id, s.store)
	if err != nil {
		ctx.JSON(400, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, logs)
}

// confirm payment
func (s *Server) ConfirmPayment(ctx *gin.Context) {
	//TODO
}
