package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	autoPaging(client)
	manualPaging(client)
}

// autoPaging demonstrates automatic pagination using the AutoPaging iterator.
// The iterator automatically fetches subsequent pages as needed, allowing you
// to iterate through all results without manually managing page tokens.
func autoPaging(client metronome.Client) {
	iter := client.V1.Customers.ListAutoPaging(context.TODO(), metronome.V1CustomerListParams{})

	// Automatically fetches more pages as needed.
	for iter.Next() {
		customer := iter.Current()
		fmt.Printf("Customer: %s (ID: %s)\n", customer.Name, customer.ID)
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

// manualPaging demonstrates manual page-by-page iteration using GetNextPage().
// This gives you more control over pagination, allowing you to process one page
// at a time and decide when to fetch the next page.
func manualPaging(client metronome.Client) {
	page, err := client.V1.Customers.List(context.TODO(), metronome.V1CustomerListParams{})
	for page != nil {
		for _, customer := range page.Data {
			fmt.Printf("Customer: %s (ID: %s)\n", customer.Name, customer.ID)
		}
		page, err = page.GetNextPage()
	}
	if err != nil {
		panic(err)
	}
}
