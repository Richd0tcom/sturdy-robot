package handlers

import (
	"github.com/Richd0tcom/sturdy-robot/internal/handlers/requests"
	"github.com/Richd0tcom/sturdy-robot/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) SetupBranchHandler()  {
	r:= s.serverRouter.Group("/branches")

	
	r.POST("/customer", s.AddCustomer)
	r.DELETE("/customer", s.RemoveCustomer)
}

// - manage customer

func (s *Server) AddCustomer(ctx *gin.Context) {
	
}

func (s *Server) RemoveCustomer(ctx *gin.Context) {
	var arg requests.DeleteCustomerReq

	utils.ExtractTokenIDs(ctx)

	ctx.ShouldBindJSON(&arg)
}

func (s *Server) GetActivityLog(ctx *gin.Context) {
	
}