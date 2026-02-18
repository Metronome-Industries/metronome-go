package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	regenerateInvoice(client)
	voidInvoice(client)
}

// regenerateInvoice regenerates an invoice.
func regenerateInvoice(client metronome.Client) {
	resp, err := client.V1.Invoices.Regenerate(context.TODO(), metronome.V1InvoiceRegenerateParams{
		ID: "6a37bb88-8538-48c5-b37b-a41c836328bd",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Regenerated invoice: %+v\n", resp)
}

// voidInvoice voids an invoice.
func voidInvoice(client metronome.Client) {
	resp, err := client.V1.Invoices.Void(context.TODO(), metronome.V1InvoiceVoidParams{
		ID: "6a37bb88-8538-48c5-b37b-a41c836328bd",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Voided invoice: %+v\n", resp)
}
