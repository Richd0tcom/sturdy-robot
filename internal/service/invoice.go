package service

import (
	"context"
	"encoding/json"
	"fmt"

	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
	"github.com/Richd0tcom/sturdy-robot/internal/handlers/requests"
	"github.com/Richd0tcom/sturdy-robot/internal/utils"
	"github.com/jackc/pgx/v5"

	"github.com/shopspring/decimal"
)

// create new invoice
func CreateNewInvoice(ctx context.Context, userID string, args requests.CreateInvoiceReq, st db.Store) (requests.InvoiceResponse, error) {

	sub_total := decimal.NewFromInt(0)
	for _, item := range args.InvoiceItems {
		sub_total = sub_total.Add(decimal.NewFromInt(int64(item.Quantity)).Mul(decimal.NewFromFloat(item.UnitPrice)))
	}
	rem, err := json.Marshal(args.Reminders)

	if err != nil {
		fmt.Println(err)
		return requests.InvoiceResponse{}, err
	}

	var invoice requests.InvoiceResponse

	//with transaction
	err = st.ExecTx(ctx, func(q *db.Queries) error {

		invoice.Invoice, err = q.CreateInvoice(ctx, db.CreateInvoiceParams{
			CustomerID:    utils.ParseUUID(args.CustomerID),
			CurrencyID:    utils.ParseUUID(args.CurrencyID),
			InvoiceNumber: utils.RandomInvoiceNumber(),
			Subtotal:      utils.DecimalToPGNumeric(sub_total),
			Status:        args.Status,
			Total:         utils.DecimalToPGNumeric(sub_total), //change to db generated
			DueDate:       utils.ParseDate(args.DueDate),
			Reminders:     rem,
			CreatedBy:     utils.ParseUUID(userID),
			PaymentInfo:   utils.ParseUUID(args.PaymentInfoID),
		})

		if err != nil {
			return err
		}

		items := make([]db.CreateMultipleInvoiceItemsParams, len(args.InvoiceItems))

		for i, item := range args.InvoiceItems {
			sub_total := decimal.NewFromInt(0)
			sub_total = sub_total.Add(decimal.NewFromInt(int64(item.Quantity)).Mul(decimal.NewFromFloat(item.UnitPrice)))

			items[i] = db.CreateMultipleInvoiceItemsParams{
				ID:        utils.ParseUUID(utils.NewRandomUUID().String()),
				InvoiceID: invoice.Invoice.ID,
				VersionID: utils.ParseUUID(item.VersionID),
				Quantity:  int32(item.Quantity),
				UnitPrice: utils.DecimalToPGNumeric(decimal.NewFromFloat(item.UnitPrice)),
				Subtotal:  utils.DecimalToPGNumeric(sub_total),
			}
		}

		_, err = st.CreateMultipleInvoiceItems(ctx, items)

		if err != nil {
			return err
		}

		invoice.Items, err = st.GetInvoiceItemsByInvoiceID(ctx, invoice.Invoice.ID)

		if args.Status != "draft" {
			_, err = q.CreateActivityLog(ctx, db.CreateActivityLogParams{
				EntityType: "invoice",
				EntityID:   invoice.Invoice.ID,
				Action:     "Create Invoice",
			})

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return requests.InvoiceResponse{}, err
	}

	return invoice, err
}

// ,
func UpdateInvoice(ctx context.Context, invoiceID string, args requests.UpdateInvoiceReq, st db.Store) (requests.InvoiceResponse, error) {
	//TODO: dynamic patching

	invoice, err := st.GetInvoiceByID(ctx, utils.ParseUUID(invoiceID))

	if err != nil {
		fmt.Println(err)
		return requests.InvoiceResponse{}, err
	}

	invoiceToSet := db.UpdateInvoiceParams{
		ID: invoice.ID,
	}

	//PS I hate doing this too

		if args.CurrencyID != "" {
			invoiceToSet.CurrencyID = utils.ParseUUID(args.CurrencyID)
		} else {
			invoiceToSet.CurrencyID = invoice.CurrencyID
		}

		if args.DueDate != "" {
			invoiceToSet.DueDate = utils.ParseDate(args.DueDate)

		} else {
			invoiceToSet.DueDate = invoice.DueDate
		}

		if args.PaymentInfoID != "" {
			invoiceToSet.PaymentInfo = utils.ParseUUID(args.PaymentInfoID)
		} else {
			invoiceToSet.PaymentInfo = invoice.PaymentInfo
		}

		if args.CustomerID != "" {
			invoiceToSet.CustomerID = utils.ParseUUID(args.CustomerID)
		} else {
			invoiceToSet.CustomerID = invoice.CustomerID
		}

		sub_total := decimal.NewFromInt(0)
		for _, item := range args.InvoiceItems {
			sub_total = sub_total.Add(decimal.NewFromInt(int64(item.Quantity)).Mul(decimal.NewFromFloat(item.UnitPrice)))
		}
		rem, err := json.Marshal(args.Reminders)

		if err != nil {
			fmt.Println(err)
			return requests.InvoiceResponse{}, err
		}
		invoiceToSet.Reminders = rem
		invoiceToSet.Subtotal = utils.DecimalToPGNumeric(sub_total)
		invoiceToSet.Total = utils.DecimalToPGNumeric(sub_total)
	

	// sort and recalculate invoice items


	if err != nil {
		fmt.Println(err)
		return requests.InvoiceResponse{}, err
	}

	st.ExecTx(ctx, func(q *db.Queries) error {

		_, err:= st.UpdateInvoice(ctx, invoiceToSet)

		if err != nil {
			return err
		}

		err=st.DeleteItemsByInvoiceId(ctx, invoice.ID)
		if err != nil {
			return err
		}


		items := make([]db.CreateMultipleInvoiceItemsParams, len(args.InvoiceItems))

		for i, item := range args.InvoiceItems {
			sub_total := decimal.NewFromInt(0)
			sub_total = sub_total.Add(decimal.NewFromInt(int64(item.Quantity)).Mul(decimal.NewFromFloat(item.UnitPrice)))

			items[i] = db.CreateMultipleInvoiceItemsParams{
				ID:        utils.ParseUUID(utils.NewRandomUUID().String()),
				InvoiceID: invoice.ID,
				VersionID: utils.ParseUUID(item.VersionID),
				Quantity:  int32(item.Quantity),
				UnitPrice: utils.DecimalToPGNumeric(decimal.NewFromFloat(item.UnitPrice)),
				Subtotal:  utils.DecimalToPGNumeric(sub_total),
			}
		}

		_, err = st.CreateMultipleInvoiceItems(ctx, items)

		if err != nil {
			return err
		}

		return nil
	})
}

// see analytics
func GetAnalytics(ctx context.Context, userID string, st db.Store) (db.GetTotalsByStatusesRow, error) {
	row, err := st.GetTotalsByStatuses(ctx, utils.ParseUUID(userID))

	if err != nil {
		return db.GetTotalsByStatusesRow{}, err
	}
	return row, nil
}

// change invoice settings

// edit/set reminder
func SetReminder(ctx context.Context, args requests.UpdateReminders, st db.Store) {

}

// get invoice
func GetInvoice(ctx context.Context, invoice_id string, st db.Store) (requests.InvoiceResponse, error) {
	inv, err := st.GetInvoiceByID(ctx, utils.ParseUUID(invoice_id))
	if err != nil {
		return requests.InvoiceResponse{}, err
	}
	sender, err := st.GetUserById(ctx, inv.CreatedBy)
	if err != nil {
		return requests.InvoiceResponse{}, err
	}

	customer, err := st.GetCustomerById(ctx, inv.CustomerID)
	if err != nil {
		return requests.InvoiceResponse{}, err
	}

	items, err := st.GetInvoiceItemsByInvoiceID(ctx, utils.ParseUUID(invoice_id))
	if err != nil {
		return requests.InvoiceResponse{}, err
	}

	return requests.InvoiceResponse{
		Invoice:      inv,
		SenderInfo:   sender,
		CustomerInfo: customer,
		Items:        items,
	}, nil
}

func GetAllInvoicesByUser(ctx context.Context, user_id string, st db.Store) ([]db.Invoice, error) {
	invs, err := st.GetInvoicesCreatedByUser(ctx, utils.ParseUUID(user_id))
	if err != nil {
		return []db.Invoice{}, err
	}
	return invs, nil
}

// get invoice items
func GetInvoiceItems(ctx context.Context, invoice_id string, st db.Store) ([]db.GetInvoiceItemsByInvoiceIDRow, error) {
	items, err := st.GetInvoiceItemsByInvoiceID(ctx, utils.ParseUUID(invoice_id))

	if err != nil {
		return []db.GetInvoiceItemsByInvoiceIDRow{}, err
	}

	return items, nil
}

// get payment info
func GetPaymentInfo(ctx context.Context, userID string, st db.Store) (db.PaymentInfo, error) {
	info, err := st.GetPaymentInfoByUserID(ctx, utils.ParseUUID(userID))

	if err != nil {
		return db.PaymentInfo{}, err
	}

	return info, nil
}

func GetInvoiceActivityLog(ctx context.Context, args string, st db.Store) ([]db.ActivityLog, error) {
	logs, err := st.GetActivityLogByEntityID(ctx, utils.ParseUUID(args))

	if err != nil {
		return []db.ActivityLog{}, err
	}

	return logs, nil
}

func PrintPDF(ctx context.Context, invoice_id string, st db.Store) {
	//TODO
}
