package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	createBillingConfig(client)
	getBillingConfig(client)
	deleteBillingConfig(client)
}

// createBillingConfig creates a billing provider configuration for a customer.
func createBillingConfig(client metronome.Client) {
	err := client.V1.Customers.BillingConfig.New(context.TODO(), metronome.V1CustomerBillingConfigNewParams{
		CustomerID:                "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		BillingProviderType:       metronome.V1CustomerBillingConfigNewParamsBillingProviderTypeStripe,
		BillingProviderCustomerID: "cus_AJ6y20bjkOOayM",
		StripeCollectionMethod:    metronome.V1CustomerBillingConfigNewParamsStripeCollectionMethodChargeAutomatically,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Billing config created successfully")
}

// getBillingConfig retrieves a billing provider configuration for a customer.
func getBillingConfig(client metronome.Client) {
	resp, err := client.V1.Customers.BillingConfig.Get(context.TODO(), metronome.V1CustomerBillingConfigGetParams{
		CustomerID:          "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		BillingProviderType: metronome.V1CustomerBillingConfigGetParamsBillingProviderTypeStripe,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Billing config: %+v\n", resp)
}

// deleteBillingConfig deletes a billing provider configuration for a customer.
func deleteBillingConfig(client metronome.Client) {
	err := client.V1.Customers.BillingConfig.Delete(context.TODO(), metronome.V1CustomerBillingConfigDeleteParams{
		CustomerID:          "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		BillingProviderType: metronome.V1CustomerBillingConfigDeleteParamsBillingProviderTypeStripe,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Billing config deleted successfully")
}
