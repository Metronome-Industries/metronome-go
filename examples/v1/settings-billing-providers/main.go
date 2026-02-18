package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	createBillingProvider(client)
	listBillingProviders(client)
}

// createBillingProvider sets up a new billing provider configuration.
func createBillingProvider(client metronome.Client) {
	resp, err := client.V1.Settings.BillingProviders.New(context.TODO(), metronome.V1SettingBillingProviderNewParams{
		BillingProvider: metronome.V1SettingBillingProviderNewParamsBillingProviderAwsMarketplace,
		Configuration: map[string]any{
			"aws_external_id":  "external-id-123",
			"aws_iam_role_arn": "arn:aws:iam::123456789012:role/MetronomeRole",
		},
		DeliveryMethod: metronome.V1SettingBillingProviderNewParamsDeliveryMethodDirectToBillingProvider,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created billing provider: %+v\n", resp)
}

// listBillingProviders lists all configured billing providers.
func listBillingProviders(client metronome.Client) {
	resp, err := client.V1.Settings.BillingProviders.List(context.TODO(), metronome.V1SettingBillingProviderListParams{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Billing providers: %+v\n", resp)
}
