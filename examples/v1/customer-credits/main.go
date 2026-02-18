package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	createCredit(client)
	listCredits(client)
	updateCreditEndDate(client)
}

// createCredit creates a new credit for a customer.
func createCredit(client metronome.Client) {
	resp, err := client.V1.Customers.Credits.New(context.TODO(), metronome.V1CustomerCreditNewParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		Priority:   100,
		ProductID:  "f14d6729-6a44-4b13-9908-9387f1918790",
		AccessSchedule: metronome.V1CustomerCreditNewParamsAccessSchedule{
			ScheduleItems: []metronome.V1CustomerCreditNewParamsAccessScheduleScheduleItem{{
				Amount:       1000,
				StartingAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				EndingBefore: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			}},
		},
		Name:        metronome.String("Annual Credit"),
		Description: metronome.String("Credit for 2024"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created credit: %+v\n", resp)
}

// listCredits lists all credits for a customer.
func listCredits(client metronome.Client) {
	page, err := client.V1.Customers.Credits.List(context.TODO(), metronome.V1CustomerCreditListParams{
		CustomerID:     "13117714-3f05-48e5-a6e9-a66093f13b4d",
		IncludeBalance: metronome.Bool(true),
	})
	if err != nil {
		panic(err)
	}
	for _, credit := range page.Data {
		fmt.Printf("Credit: %s (Type: %s)\n", credit.ID, credit.Type)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.Customers.Credits.ListAutoPaging(context.TODO(), metronome.V1CustomerCreditListParams{
	//     CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	// })
	// for iter.Next() {
	//     credit := iter.Current()
	//     fmt.Printf("Credit: %s\n", credit.ID)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// updateCreditEndDate updates the end date of a credit.
func updateCreditEndDate(client metronome.Client) {
	resp, err := client.V1.Customers.Credits.UpdateEndDate(context.TODO(), metronome.V1CustomerCreditUpdateEndDateParams{
		CreditID:           "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
		CustomerID:         "13117714-3f05-48e5-a6e9-a66093f13b4d",
		AccessEndingBefore: time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated credit end date: %+v\n", resp)
}
