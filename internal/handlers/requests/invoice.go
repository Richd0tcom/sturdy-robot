package requests

import db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"


type CreateInvoiceReq struct {
	CurrencyID     string `json:"currency_id"`
    BranchID       string `json:"branch_id"`
    CustomerID     string `json:"customer_id"`
    DueDate        string `json:"due_date"`
    Reminders      []string `json:"reminders"`
    PaymentInfoID  string `json:"payment_info_id"`
    Status         string `json:"status"`
    InvoiceItems   []InvoiceItem `json:"invoice_items"`
}

type UpdateInvoiceReq struct {
	CurrencyID     string `json:"currency_id"`
    CustomerID     string `json:"customer_id"`
    DueDate        string `json:"due_date"`
    Reminders      []string `json:"reminders"`
    PaymentInfoID  string `json:"payment_info_id"`
    Status         string `json:"status"`
    InvoiceItems   []InvoiceItem `json:"invoice_items"`
}

type InvoiceItem struct {
    VersionID      string  `json:"version_id"`
    Quantity       int     `json:"quantity"`
    UnitPrice      float64 `json:"unit_price"`
}

type InvoiceResponse struct {
    Currency db.Currency `json:"currency"`
    Invoice db.Invoice `json:"invoice"`
    SenderInfo db.User `json:"sender_info"`
    CustomerInfo db.Customer `json:"customer_info"`
    Items []db.GetInvoiceItemsByInvoiceIDRow `json:"items"`
}

type InvoiceID struct {
    ID string `uri:"invoice_id" binding:"required,uuid"`
}

type UpdateReminders struct {
    InvoiceID string
    Reminders []string
}

type ConfirmPaymentRequest struct {
    InvoiceID string `json:"invoice_id" binding:"required"`

    PaymentDetails 
}

type PaymentDetails struct {
    PaymentMethod string `json:"payment_method"`
    PaymentAmount float64 `json:"payment_amount"`
    PaymentRef string `json:"payment_ref"`
    
}