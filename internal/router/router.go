package router

import (
	"github.com/Richd0tcom/sturdy-robot/internal/handlers"
)


func SetupRouter(srv *handlers.Server) {

	branchGroup:= srv.ServerRouter.Group("/branch/:branch_id")

	srv.SetupBranchHandler()
	srv.SetupCustomerHandler(branchGroup)
	srv.SetupInvoiceHandler(branchGroup)
	srv.SetupProductHandler(branchGroup)
}