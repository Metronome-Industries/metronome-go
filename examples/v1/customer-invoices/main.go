package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	getInvoice(client)
	listInvoices(client)
	addCharge(client)
	listBreakdowns(client)
	getInvoicePdf(client)
}

// getInvoice retrieves a specific invoice for a customer.
func getInvoice(client metronome.Client) {
	resp, err := client.V1.Customers.Invoices.Get(context.TODO(), metronome.V1CustomerInvoiceGetParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		InvoiceID:  "6a37bb88-8538-48c5-b37b-a41c836328bd",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Invoice: %+v\n", resp)
}

// listInvoices lists all invoices for a customer.
func listInvoices(client metronome.Client) {
	page, err := client.V1.Customers.Invoices.List(context.TODO(), metronome.V1CustomerInvoiceListParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	})
	if err != nil {
		panic(err)
	}
	for _, invoice := range page.Data {
		fmt.Printf("Invoice: %s (Status: %s)\n", invoice.ID, invoice.Status)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.Customers.Invoices.ListAutoPaging(context.TODO(), metronome.V1CustomerInvoiceListParams{
	//     CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	// })
	// for iter.Next() {
	//     invoice := iter.Current()
	//     fmt.Printf("Invoice: %s\n", invoice.ID)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// addCharge adds a one-time charge to a customer's invoice.
func addCharge(client metronome.Client) {
	resp, err := client.V1.Customers.Invoices.AddCharge(context.TODO(), metronome.V1CustomerInvoiceAddChargeParams{
		CustomerID:            "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		ChargeID:              "5ae4b726-1ebe-439c-9190-9831760ba195",
		CustomerPlanID:        "a23b3cf4-47fb-4c3f-bb3d-9e64f7704015",
		Description:           "One time charge",
		InvoiceStartTimestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Price:                 250,
		Quantity:              1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added charge: %+v\n", resp)
}

// listBreakdowns lists invoice breakdowns for a customer.
func listBreakdowns(client metronome.Client) {
	page, err := client.V1.Customers.Invoices.ListBreakdowns(context.TODO(), metronome.V1CustomerInvoiceListBreakdownsParams{
		CustomerID:   "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		StartingOn:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndingBefore: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		panic(err)
	}
	for _, breakdown := range page.Data {
		fmt.Printf("Breakdown: %+v\n", breakdown)
	}

	// To automatically paginate, use ListBreakdownsAutoPaging:
	//
	// iter := client.V1.Customers.Invoices.ListBreakdownsAutoPaging(context.TODO(), metronome.V1CustomerInvoiceListBreakdownsParams{...})
	// for iter.Next() {
	//     breakdown := iter.Current()
	//     fmt.Printf("Breakdown: %+v\n", breakdown)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// getInvoicePdf downloads the PDF for a specific invoice.
func getInvoicePdf(client metronome.Client) {
	resp, err := client.V1.Customers.Invoices.GetPdf(context.TODO(), metronome.V1CustomerInvoiceGetPdfParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		InvoiceID:  "6a37bb88-8538-48c5-b37b-a41c836328bd",
	})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	pdfBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Downloaded invoice PDF: %d bytes\n", len(pdfBytes))
}
