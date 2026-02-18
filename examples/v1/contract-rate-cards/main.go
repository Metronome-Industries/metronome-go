package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/shared"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	createRateCard(client)
	getRateCard(client)
	updateRateCard(client)
	listRateCards(client)
	archiveRateCard(client)
	getRateSchedule(client)
}

// createRateCard creates a new rate card.
func createRateCard(client metronome.Client) {
	resp, err := client.V1.Contracts.RateCards.New(context.TODO(), metronome.V1ContractRateCardNewParams{
		Name: "My Rate Card",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created rate card: %+v\n", resp)
}

// getRateCard retrieves a rate card by ID.
func getRateCard(client metronome.Client) {
	resp, err := client.V1.Contracts.RateCards.Get(context.TODO(), metronome.V1ContractRateCardGetParams{
		ID: shared.IDParam{
			ID: "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rate card: %+v\n", resp)
}

// updateRateCard updates a rate card.
func updateRateCard(client metronome.Client) {
	resp, err := client.V1.Contracts.RateCards.Update(context.TODO(), metronome.V1ContractRateCardUpdateParams{
		RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		Name:       metronome.String("Updated Rate Card Name"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated rate card: %+v\n", resp)
}

// listRateCards lists all rate cards.
func listRateCards(client metronome.Client) {
	page, err := client.V1.Contracts.RateCards.List(context.TODO(), metronome.V1ContractRateCardListParams{})
	if err != nil {
		panic(err)
	}
	for _, card := range page.Data {
		fmt.Printf("Rate card: %+v\n", card)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.Contracts.RateCards.ListAutoPaging(context.TODO(), metronome.V1ContractRateCardListParams{})
	// for iter.Next() {
	//     card := iter.Current()
	//     fmt.Printf("Rate card: %+v\n", card)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// archiveRateCard archives a rate card.
func archiveRateCard(client metronome.Client) {
	resp, err := client.V1.Contracts.RateCards.Archive(context.TODO(), metronome.V1ContractRateCardArchiveParams{
		ID: shared.IDParam{
			ID: "12b21470-4570-40df-8998-449d0b0bc52f",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Archived rate card: %+v\n", resp)
}

// getRateSchedule retrieves the rate schedule for a rate card.
func getRateSchedule(client metronome.Client) {
	resp, err := client.V1.Contracts.RateCards.GetRateSchedule(context.TODO(), metronome.V1ContractRateCardGetRateScheduleParams{
		RateCardID: "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe",
		StartingAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rate schedule: %+v\n", resp)
}
