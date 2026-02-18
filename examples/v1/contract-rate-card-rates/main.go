package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	listRates(client)
	addRate(client)
	addManyRates(client)
}

// listRates lists rates for a rate card at a specific point in time.
func listRates(client metronome.Client) {
	page, err := client.V1.Contracts.RateCards.Rates.List(context.TODO(), metronome.V1ContractRateCardRateListParams{
		RateCardID: "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe",
		At:         time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		panic(err)
	}
	for _, rate := range page.Data {
		fmt.Printf("Rate: %+v\n", rate)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.Contracts.RateCards.Rates.ListAutoPaging(context.TODO(), metronome.V1ContractRateCardRateListParams{
	//     RateCardID: "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe",
	//     At:         time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	// })
	// for iter.Next() {
	//     rate := iter.Current()
	//     fmt.Printf("Rate: %+v\n", rate)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// addRate adds a single rate to a rate card.
func addRate(client metronome.Client) {
	resp, err := client.V1.Contracts.RateCards.Rates.Add(context.TODO(), metronome.V1ContractRateCardRateAddParams{
		RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		ProductID:  "13117714-3f05-48e5-a6e9-a66093f13b4d",
		RateType:   metronome.V1ContractRateCardRateAddParamsRateTypeFlat,
		StartingAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Entitled:   true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added rate: %+v\n", resp)
}

// addManyRates adds multiple rates to a rate card at once.
func addManyRates(client metronome.Client) {
	resp, err := client.V1.Contracts.RateCards.Rates.AddMany(context.TODO(), metronome.V1ContractRateCardRateAddManyParams{
		RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		Rates: []metronome.V1ContractRateCardRateAddManyParamsRate{{
			ProductID:  "13117714-3f05-48e5-a6e9-a66093f13b4d",
			RateType:   "FLAT",
			StartingAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			Entitled:   true,
		}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added rates: %+v\n", resp)
}
