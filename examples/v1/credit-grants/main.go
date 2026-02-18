package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	createCreditGrant(client)
	listCreditGrants(client)
	editCreditGrant(client)
	listCreditEntries(client)
	voidCreditGrant(client)
}

// createCreditGrant creates a new credit grant for a customer.
func createCreditGrant(client metronome.Client) {
	resp, err := client.V1.CreditGrants.New(context.TODO(), metronome.V1CreditGrantNewParams{
		CustomerID: "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		ExpiresAt:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		GrantAmount: metronome.V1CreditGrantNewParamsGrantAmount{
			Amount:       1000,
			CreditTypeID: "5ae401dc-a648-4b49-9ac3-391bb5bc4d7b",
		},
		Name: "Acme Corp Promotional Credit Grant",
		PaidAmount: metronome.V1CreditGrantNewParamsPaidAmount{
			Amount:       5000,
			CreditTypeID: "2714e483-4ff1-48e4-9e25-ac732e8f24f2",
		},
		Priority: 0.5,
		Reason:   metronome.String("Incentivize new customer"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created credit grant: %+v\n", resp)
}

// listCreditGrants lists all credit grants, optionally filtered by customer.
func listCreditGrants(client metronome.Client) {
	page, err := client.V1.CreditGrants.List(context.TODO(), metronome.V1CreditGrantListParams{
		CustomerIDs: []string{"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"},
	})
	if err != nil {
		panic(err)
	}
	for _, grant := range page.Data {
		fmt.Printf("Credit grant: %s (Balance: %+v)\n", grant.ID, grant.Balance)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.CreditGrants.ListAutoPaging(context.TODO(), metronome.V1CreditGrantListParams{})
	// for iter.Next() {
	//     grant := iter.Current()
	//     fmt.Printf("Credit grant: %s\n", grant.ID)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// editCreditGrant edits an existing credit grant.
func editCreditGrant(client metronome.Client) {
	resp, err := client.V1.CreditGrants.Edit(context.TODO(), metronome.V1CreditGrantEditParams{
		ID:        "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		ExpiresAt: metronome.Time(time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)),
		Name:      metronome.String("Updated Credit Grant Name"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Edited credit grant: %+v\n", resp)
}

// listCreditEntries lists credit ledger entries.
func listCreditEntries(client metronome.Client) {
	page, err := client.V1.CreditGrants.ListEntries(context.TODO(), metronome.V1CreditGrantListEntriesParams{
		CustomerIDs: []string{"6a37bb88-8538-48c5-b37b-a41c836328bd"},
	})
	if err != nil {
		panic(err)
	}
	for _, entry := range page.Data {
		fmt.Printf("Credit entry: %+v\n", entry)
	}

	// To automatically paginate, use ListEntriesAutoPaging:
	//
	// iter := client.V1.CreditGrants.ListEntriesAutoPaging(context.TODO(), metronome.V1CreditGrantListEntriesParams{
	//     CustomerIDs: []string{"6a37bb88-8538-48c5-b37b-a41c836328bd"},
	// })
	// for iter.Next() {
	//     entry := iter.Current()
	//     fmt.Printf("Entry: %+v\n", entry)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// voidCreditGrant voids a credit grant.
func voidCreditGrant(client metronome.Client) {
	resp, err := client.V1.CreditGrants.Void(context.TODO(), metronome.V1CreditGrantVoidParams{
		ID: "9b85c1c1-5238-4f2a-a409-61412905e1e1",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Voided credit grant: %+v\n", resp)
}
