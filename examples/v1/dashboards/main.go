package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	getEmbeddableURL(client)
}

// getEmbeddableURL generates an embeddable dashboard URL for a customer.
func getEmbeddableURL(client metronome.Client) {
	resp, err := client.V1.Dashboards.GetEmbeddableURL(context.TODO(), metronome.V1DashboardGetEmbeddableURLParams{
		CustomerID: "4db51251-61de-4bfe-b9ce-495e244f3491",
		Dashboard:  metronome.V1DashboardGetEmbeddableURLParamsDashboardInvoices,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Embeddable URL: %+v\n", resp)
}
