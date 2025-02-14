package service

// //pdf service here

// import (
// 	"fmt"
// 	// "os"
// 	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
// 	"github.com/jung-kurt/gofpdf"
// )

// type PDFInvoice struct {
//     Invoice struct {
//         db.Invoice
//     }
//     Customer struct {
//         db.Customer
//     }
//     Sender struct {
//         db.User
//     }
//     Items       []InvoiceItem
// }

// type InvoiceItem struct {
//     db.GetInvoiceItemsByInvoiceIDRow
// }

// func GenerateInvoicePDF(invoice *PDFInvoice) error {
//     pdf := gofpdf.New("P", "mm", "A4", "")
//     pdf.AddPage()

//     // Set fonts
//     pdf.SetFont("Arial", "B", 16)
//     pdf.Cell(190, 10, "Invoice")
//     pdf.Ln(10)

//     // Invoice Number
//     pdf.SetFont("Arial", "", 10)
//     pdf.Cell(0, 10, fmt.Sprintf("Invoice Number: %s", invoice.Invoice.InvoiceNumber))
//     pdf.Ln(10)

//     // Sender Details
//     pdf.Cell(0, 10, "From:")
//     pdf.Ln(5)
//     pdf.Cell(0, 10, invoice.Sender.Name)
//     pdf.Ln(5)
//     pdf.Cell(0, 10, invoice.Sender.Address.String)
//     pdf.Ln(5)
//     pdf.Cell(0, 10, invoice.Sender.Email)
//     pdf.Ln(10)

//     // Customer Details
//     pdf.Cell(0, 10, "Bill To:")
//     pdf.Ln(5)
//     pdf.Cell(0, 10, invoice.Customer.Name)
//     pdf.Ln(5)
//     pdf.Cell(0, 10, string(invoice.Customer.BillingAddress)) //TODO: proper formatting
//     pdf.Ln(5)
//     pdf.Cell(0, 10, invoice.Customer.Email.String)
//     pdf.Ln(10)

//     // Items Table Header
//     pdf.SetFillColor(230, 230, 230)
//     pdf.SetFont("Arial", "B", 10)
//     headers := []string{"Description", "Quantity", "Unit Price", "Total"}
//     colWidths := []float64{90, 30, 30, 40}
    
//     for i, header := range headers {
//         pdf.CellFormat(colWidths[i], 7, header, "1", 0, "C", true, 0, "")
//     }
//     pdf.Ln(7)

//     // Items Table Body
//     pdf.SetFont("Arial", "", 10)
//     for _, item := range invoice.Items {
//         pdf.CellFormat(colWidths[0], 7, item.Name, "1", 0, "L", false, 0, "")
//         pdf.CellFormat(colWidths[1], 7, fmt.Sprintf("%d", item.Quantity), "1", 0, "R", false, 0, "")
//         pdf.CellFormat(colWidths[2], 7, fmt.Sprintf("$%.2f", item.UnitPrice), "1", 0, "R", false, 0, "")
//         pdf.CellFormat(colWidths[3], 7, fmt.Sprintf("$%.2f", item.Subtotal), "1", 0, "R", false, 0, "")
//         pdf.Ln(7)
//     }

//     // Totals
//     pdf.Ln(10)
//     pdf.SetFont("Arial", "B", 10)
//     pdf.Cell(130, 7, "Subtotal:")
//     pdf.CellFormat(60, 7, fmt.Sprintf("$%.2f", invoice.Invoice.Subtotal), "1", 1, "R", false, 0, "")
    
//     pdf.Cell(130, 7, "Discount:")
//     pdf.CellFormat(60, 7, fmt.Sprintf("$%.2f", invoice.Invoice.Discount), "1", 1, "R", false, 0, "")
    
//     pdf.Cell(130, 7, "Total:")
//     pdf.CellFormat(60, 7, fmt.Sprintf("$%.2f", invoice.Invoice.Total), "1", 1, "R", true, 0, "")

//     // Save PDF
//     return pdf.OutputFileAndClose("invoice.pdf")
// }

// // func ExampleInvoiceGeneration() {
// //     invoice := &PDFInvoice{
// //         InvoiceNumber: "INV-2024-001",
// //         Customer: struct {
// //             Name    string
// //             Address string
// //             Email   string
// //         }{
// //             Name:    "John Doe",
// //             Address: "123 Customer St, Cityville",
// //             Email:   "john.doe@example.com",
// //         },
// //         Sender: struct {
// //             Name    string
// //             Address string
// //             Email   string
// //         }{
// //             Name:    "Your Company",
// //             Address: "456 Business Rd, Townsville",
// //             Email:   "billing@yourcompany.com",
// //         },
// //         Items: []InvoiceItem{
// //             {
// //                 Name: "Product A",
// //                 Quantity:    2,
// //                 UnitPrice:   50.00,
// //                 Total:       100.00,
// //             },
// //             {
// //                 Name: "Product B",
// //                 Quantity:    1,
// //                 UnitPrice:   75.00,
// //                 Total:       75.00,
// //             },
// //         },
// //         SubTotal: 175.00,
// //         Discount: 10.00,
// //         Total:    165.00,
// //     }

// //     err := GenerateInvoicePDF(invoice)
// //     if err != nil {
// //         fmt.Println("Error generating PDF:", err)
// //     }
// // }

