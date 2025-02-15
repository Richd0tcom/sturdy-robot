package handlers

import (
	"net/http"

	"github.com/Richd0tcom/sturdy-robot/internal/handlers/requests"
	"github.com/Richd0tcom/sturdy-robot/internal/service"
	"github.com/Richd0tcom/sturdy-robot/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) SetupBranchHandler()  {
	r:= s.ServerRouter.Group("/branches")
	
	r.POST("/customer", s.AddCustomer)
	r.GET("/customer")
	r.DELETE("/customer", s.RemoveCustomer)
	r.GET("/activity", s.GetActivityLog)
}

// - manage customer

func (s *Server) AddCustomer(ctx *gin.Context) {
	var arg requests.CreateCustomerReq

	err:= ctx.ShouldBindJSON(&arg)
	if err != nil {
		ctx.JSON(400, buildErrorResponse(err))
		return 
	}
	customer, err:= service.AddCustomer(ctx, arg, s.store)

	if err != nil {
		ctx.JSON(400, buildErrorResponse(err))
		return 
	}
	ctx.JSON(http.StatusOK, customer)
}

func (s *Server) RemoveCustomer(ctx *gin.Context) {
	var arg requests.DeleteCustomerReq

	err:=ctx.ShouldBindJSON(&arg)
	if err != nil {
		ctx.JSON(400, buildErrorResponse(err))
		return 
	}

	err= service.RemoveCustomer(ctx, arg.ID, s.store)

	if err != nil {
		ctx.JSON(400, buildErrorResponse(err))
		return 
	}
	ctx.JSON(http.StatusOK, buildSuccessResponse(gin.H{}))
}

func (s *Server) GetActivityLog(ctx *gin.Context) {
	claims, err := utils.ExtractTokenIDs(ctx)
	if err != nil {
		ctx.JSON(400, buildErrorResponse(err))
		return
	}
	userID := claims["user_id"].(string)
	logs, err:= service.GetUserActivityLog(ctx, userID, s.store)

	if err != nil {
		ctx.JSON(400, buildErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, logs)
}

func (s *Server) GetBranches(c *gin.Context){
	
}
