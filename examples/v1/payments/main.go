package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	listPayments(client)
	attemptPayment(client)
	cancelPayment(client)
}

// listPayments lists payments for a customer.
func listPayments(client metronome.Client) {
	page, err := client.V1.Payments.List(context.TODO(), metronome.V1PaymentListParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		panic(err)
	}
	for _, payment := range page.Data {
		fmt.Printf("Payment: %s (Status: %s)\n", payment.ID, payment.Status)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.Payments.ListAutoPaging(context.TODO(), metronome.V1PaymentListParams{
	//     CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	// })
	// for iter.Next() {
	//     payment := iter.Current()
	//     fmt.Printf("Payment: %s\n", payment.ID)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// attemptPayment triggers a payment attempt for a specific invoice.
func attemptPayment(client metronome.Client) {
	resp, err := client.V1.Payments.Attempt(context.TODO(), metronome.V1PaymentAttemptParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		InvoiceID:  "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Payment attempt: %+v\n", resp)
}

// cancelPayment cancels a pending payment.
func cancelPayment(client metronome.Client) {
	resp, err := client.V1.Payments.Cancel(context.TODO(), metronome.V1PaymentCancelParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		InvoiceID:  "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Cancelled payment: %+v\n", resp)
}
