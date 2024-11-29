package service

import (
	"context"
	"encoding/json"

	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
	"github.com/Richd0tcom/sturdy-robot/internal/handlers/requests"
	"github.com/Richd0tcom/sturdy-robot/internal/utils"
)

func RemoveCustomer(ctx context.Context, arg string, st db.Store) error {
	err:= st.DeleteCustomerByID(ctx, utils.ParseUUID(arg))

	if err != nil {
		return err
	}
	return nil
}

func AddCustomer(ctx context.Context, args requests.CreateCustomerReq, st db.Store) (db.Customer, error){
	address, err:= json.Marshal(args.BillingAddress)
	if err != nil {
		return db.Customer{}, err
	}
	customer, err:= st.CreateCustomer(ctx, db.CreateCustomerParams{
		Email: utils.StringToPGText(args.Email),
		Name: args.Name,
		BillingAddress: address,
		BranchID: utils.ParseUUID(args.BranchID),
	})

	if err != nil {
		return db.Customer{}, err
	}

	return customer, nil
}

func GetCustomers(ctx context.Context, branch_id string,  st db.Store) ([]db.Customer, error) {
	customers, err:= st.GetCustomersByBranch(ctx, utils.ParseUUID(branch_id))

	if err != nil {
		return []db.Customer{}, err
	}

	return customers, nil
}

// get activity log
func GetUserActivityLog(ctx context.Context, userID, st db.Store) ([]db.ActivityLog, error) {
	logs, err:= st.GetActivityLogsByUserID(ctx, utils.ParseUUID(userID))

	if err != nil {
		return []db.ActivityLog{}, err
	}

	return logs, nil
}