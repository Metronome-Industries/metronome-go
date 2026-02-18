package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	listPricingUnits(client)
}

// listPricingUnits lists all available pricing units (credit types).
func listPricingUnits(client metronome.Client) {
	page, err := client.V1.PricingUnits.List(context.TODO(), metronome.V1PricingUnitListParams{})
	if err != nil {
		panic(err)
	}
	for _, unit := range page.Data {
		fmt.Printf("Pricing unit: %s (ID: %s)\n", unit.Name, unit.ID)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.PricingUnits.ListAutoPaging(context.TODO(), metronome.V1PricingUnitListParams{})
	// for iter.Next() {
	//     unit := iter.Current()
	//     fmt.Printf("Pricing unit: %s\n", unit.Name)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}
