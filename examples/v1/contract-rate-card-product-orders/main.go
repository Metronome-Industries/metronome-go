package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	moveProducts(client)
	setProductOrder(client)
}

// moveProducts moves products to new positions in a rate card's product order.
func moveProducts(client metronome.Client) {
	resp, err := client.V1.Contracts.RateCards.ProductOrders.Update(context.TODO(), metronome.V1ContractRateCardProductOrderUpdateParams{
		RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		ProductMoves: []metronome.V1ContractRateCardProductOrderUpdateParamsProductMove{{
			ProductID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
			Position:  0,
		}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated product order: %+v\n", resp)
}

// setProductOrder sets the complete product order for a rate card.
func setProductOrder(client metronome.Client) {
	resp, err := client.V1.Contracts.RateCards.ProductOrders.Set(context.TODO(), metronome.V1ContractRateCardProductOrderSetParams{
		RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		ProductOrder: []string{
			"13117714-3f05-48e5-a6e9-a66093f13b4d",
			"b086f2f4-9851-4466-9ca0-30d53e6a42ac",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Set product order: %+v\n", resp)
}
