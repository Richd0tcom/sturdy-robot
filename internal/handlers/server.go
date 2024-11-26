package handlers

import (
	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

//serves all http request in the banking application
type Server struct {
	store db.Store
	serverRouter *gin.Engine
}

//Creates a new server and sets up routing to handle request
func NewServer(store db.Store) *Server {

	engine:= gin.Default()
	
	server:= &Server{
		serverRouter: engine,
	}

	server.SetupInvoiceHandler()

	return server
}

//Starts the created sever
func (server *Server) Start (address string) error {
	return server.serverRouter.Run(address)
}

func buildErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}