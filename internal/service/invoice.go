package service

import (
	"context"
	"encoding/json"
	"fmt"

	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
	"github.com/Richd0tcom/sturdy-robot/internal/handlers/requests"
	"github.com/Richd0tcom/sturdy-robot/internal/utils"


	"github.com/shopspring/decimal"
)

// create new invoice
func CreateNewInvoice(ctx context.Context, args requests.CreateInvoiceReq, st db.Store) (requests.InvoiceResponse, error){

	sub_total:= decimal.NewFromInt(0)
	for _, item:= range args.InvoiceItems {
		sub_total = sub_total.Add(decimal.NewFromInt(int64(item.Quantity)).Mul(decimal.NewFromFloat(item.UnitPrice))) 
	}
	rem, err:= json.Marshal(args.Reminders)

	fmt.Println("s1",sub_total)

	
	invoice, err:= st.CreateInvoice(ctx,db.CreateInvoiceParams{
		CustomerID: utils.ParseUUID(args.CustomerID),
		CurrencyID: utils.ParseUUID(args.CurrencyID),
		InvoiceNumber: utils.RandomInvoiceNumber(),
		Subtotal: utils.DecimalToPGNumeric(sub_total),
		Status: args.Status,
		Total: utils.DecimalToPGNumeric(sub_total), //change to db generated
		DueDate: utils.ParseDate(args.DueDate),
		Reminders: rem,
		CreatedBy: utils.ParseUUID("f27899b4-f707-41d1-bd6d-07f584c7cea4"),
	})

	if err!=nil {
		fmt.Println(err)
		return requests.InvoiceResponse{}, err
	}


	items:= make([]db.CreateMultipleInvoiceItemsParams, len(args.InvoiceItems))

	for i, item := range args.InvoiceItems {
		sub_total:= decimal.NewFromInt(0)
		sub_total= sub_total.Add(decimal.NewFromInt(int64(item.Quantity)).Mul(decimal.NewFromFloat(item.UnitPrice))) 

		fmt.Println("sub: ",sub_total)
		
		items[i] = db.CreateMultipleInvoiceItemsParams{
			ID: utils.ParseUUID(utils.NewRandomUUID().String()),
			InvoiceID: invoice.ID,
			VersionID: utils.ParseUUID(item.VersionID),
			Quantity: int32(item.Quantity),
			UnitPrice: utils.DecimalToPGNumeric(decimal.NewFromFloat(item.UnitPrice)),
			Subtotal: utils.DecimalToPGNumeric(sub_total),
		}
	}

	_, err= st.CreateMultipleInvoiceItems(ctx, items)

	go func(ctx context.Context){
		st.CreateActivityLog(ctx, db.CreateActivityLogParams{

		})
	}(ctx)

	if err!= nil {
		return requests.InvoiceResponse{}, err
	}

	invoiceItems, err:=st.GetInvoiceItemsByInvoiceID(ctx, invoice.ID)
	return requests.InvoiceResponse{
		Invoice: invoice,
		Items: invoiceItems,
	}, err
}


// see analytics
func GetAnalytics(ctx context.Context,args requests.UserID,  st db.Store) (db.GetTotalsByStatusesRow, error){
	row, err:= st.GetTotalsByStatuses(ctx, utils.ParseUUID(args.ID) )

	if err != nil {
		return db.GetTotalsByStatusesRow{}, err
	}
	return row, nil
}


// change invoice settings 

//edit/set reminder
func SetReminder(ctx context.Context, args requests.Reminders) {

}

//get invoice 
func GetInvoice(ctx context.Context, args requests.InvoiceID, st db.Store ) (requests.InvoiceResponse, error){
	inv, err:= st.GetInvoiceByID(ctx, utils.ParseUUID(args.ID))
	if err != nil {
		return requests.InvoiceResponse{}, err
	}
	sender, err:= st.GetUserById(ctx, inv.CreatedBy)
	if err != nil {
		return requests.InvoiceResponse{}, err
	}

	customer,err:= st.GetCustomerById(ctx, inv.CustomerID)
	if err != nil {
		return requests.InvoiceResponse{}, err
	}

	items, err:= st.GetInvoiceItemsByInvoiceID(ctx, utils.ParseUUID(args.ID))
	if err != nil {
		return requests.InvoiceResponse{}, err
	}

	return requests.InvoiceResponse{
		Invoice: inv,
		SenderInfo: sender,
		CustomerInfo: customer,
		Items: items,
	}, nil
}



// get invoice items
func GetInvoiceItems(ctx context.Context, args requests.InvoiceID, st db.Store) ([]db.InvoiceItem, error) {
	items, err:= st.GetInvoiceItemsByInvoiceID(ctx, utils.ParseUUID(args.ID))

	if err != nil {
		return []db.InvoiceItem{}, err
	}

	return items, nil
}

// get payment info
func GetPaymentInfo(ctx context.Context, args requests.UserID, st db.Store) (db.PaymentInfo, error){
	info, err:= st.GetPaymentInfoByUserID(ctx, utils.ParseUUID(args.ID))

	if err != nil {
		return db.PaymentInfo{}, err
	}

	return info, nil
}

// get activity log
func GetUserActivityLog(ctx context.Context, args requests.UserID, st db.Store) ([]db.ActivityLog, error) {
	logs, err:= st.GetActivityLogsByUserID(ctx, utils.ParseUUID(args.ID))

	if err != nil {
		return []db.ActivityLog{}, err
	}

	return logs, nil

}

func GetInvoiceActivityLog(ctx context.Context, args requests.InvoiceID, st db.Store) ([]db.ActivityLog, error) {
	logs, err:= st.GetActivityLogByEntityID(ctx,  utils.ParseUUID(args.ID))

	if err != nil {
		return []db.ActivityLog{}, err
	}

	return logs, nil
}




