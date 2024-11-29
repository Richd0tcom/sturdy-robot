package handlers

import (
	"fmt"
	"github.com/Richd0tcom/sturdy-robot/internal/handlers/requests"
	"github.com/Richd0tcom/sturdy-robot/internal/service"

	// "github.com/Richd0tcom/sturdy-robot/internal/service"
	"github.com/gin-gonic/gin"
)


func (s *Server) SetupInvoiceHandler()  {
	r:= s.serverRouter.Group("/invoices")

	r.GET("/hello", func(ctx *gin.Context) {
		fmt.Println("dead last")
		ctx.JSON(200, map[string]interface{}{
			"message": "red",
		})
	}) //helo world handlerd
	r.POST("/", s.CreateInvoice)
	r.PATCH("/:id", s.UpdateInvoice)
	r.GET("/analytics", s.GetAnalytics)
	r.GET("/activity", s.GetActivityLog) 
	r.GET("/:id", s.GetInvoice)
	
}

func (s *Server) CreateInvoice(ctx *gin.Context){
	var args requests.CreateInvoiceReq
	ctx.ShouldBindJSON(&args)

	invoice, err := service.CreateNewInvoice(ctx, args, s.store)
	if err!=nil{
		fmt.Println("error", err)
		ctx.JSON(400, gin.H{})
		return 
	}

	ctx.JSON(200, invoice)
}

func (s *Server) UpdateInvoice(ctx *gin.Context) {

}


// see analytics
func (s *Server) GetAnalytics(ctx *gin.Context) {
	var args requests.UserID

	ctx.ShouldBindJSON(&args)

	service.GetAnalytics(ctx, args, s.store)
}


// change invoice settings 

//edit/set reminder
func (s *Server) SetReminder(ctx *gin.Context) {

}

//get invoice 

func (s *Server) GetInvoice(ctx *gin.Context) {

}

func (s *Server) GetAllInvoices(ctx *gin.Context) {

}

func(s *Server) GetInvoicesByStatus(ctx *gin.Context) {

}
// - sender info


// - manage customer

func (s *Server) AddCustomer(ctx *gin.Context) {

}

func (s *Server) RemoveCustomer(ctx *gin.Context) {

}
// - invoice_info
// - currency 


// get invoice items
func (s *Server) GetInvoiceItems(ctx *gin.Context) {

}

// get payment info
func (s *Server) GetPaymentInfo(ctx *gin.Context) {

}

// get activity log
func (s *Server) GetActivityLog(ctx *gin.Context) {

}

//confirm payment
func (s *Server) ConfirmPayment(ctx *gin.Context) {

}
