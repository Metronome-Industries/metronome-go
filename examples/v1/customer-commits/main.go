package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	createCommit(client)
	listCommits(client)
	updateCommitEndDate(client)
}

// createCommit creates a new prepaid commit for a customer.
func createCommit(client metronome.Client) {
	resp, err := client.V1.Customers.Commits.New(context.TODO(), metronome.V1CustomerCommitNewParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		Type:       metronome.V1CustomerCommitNewParamsTypePrepaid,
		Priority:   100,
		ProductID:  "f14d6729-6a44-4b13-9908-9387f1918790",
		AccessSchedule: metronome.V1CustomerCommitNewParamsAccessSchedule{
			ScheduleItems: []metronome.V1CustomerCommitNewParamsAccessScheduleScheduleItem{{
				Amount:       1000,
				StartingAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				EndingBefore: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			}},
		},
		Name:        metronome.String("Annual Prepaid Commit"),
		Description: metronome.String("Prepaid commit for 2024"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created commit: %+v\n", resp)
}

// listCommits lists all commits for a customer.
func listCommits(client metronome.Client) {
	page, err := client.V1.Customers.Commits.List(context.TODO(), metronome.V1CustomerCommitListParams{
		CustomerID:     "13117714-3f05-48e5-a6e9-a66093f13b4d",
		IncludeBalance: metronome.Bool(true),
	})
	if err != nil {
		panic(err)
	}
	for _, commit := range page.Data {
		fmt.Printf("Commit: %s (Type: %s)\n", commit.ID, commit.Type)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.Customers.Commits.ListAutoPaging(context.TODO(), metronome.V1CustomerCommitListParams{
	//     CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	// })
	// for iter.Next() {
	//     commit := iter.Current()
	//     fmt.Printf("Commit: %s\n", commit.ID)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// updateCommitEndDate updates the end date of a commit.
func updateCommitEndDate(client metronome.Client) {
	resp, err := client.V1.Customers.Commits.UpdateEndDate(context.TODO(), metronome.V1CustomerCommitUpdateEndDateParams{
		CommitID:           "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
		CustomerID:         "13117714-3f05-48e5-a6e9-a66093f13b4d",
		AccessEndingBefore: metronome.Time(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated commit end date: %+v\n", resp)
}
