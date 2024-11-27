package service

import (
	"context"

	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"

	//  "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
)

type Invoice db.Invoice

// create new invoice
func CreateNewInvoice(ctx context.Context, args db.CreateInvoiceParams, st db.Store) (db.Invoice, error){

	invoice, err:= st.CreateInvoice(ctx, args)

	

	return invoice, err
}

// see analytics


// change invoice settings 

//edit/set reminder

//get invoice 
// - sender info
// - customer_info
// - invoice_info
// - currency 


// get invoice items

// get payment info

// get activity log




