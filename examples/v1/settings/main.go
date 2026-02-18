package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	upsertAvalaraCredentials(client)
}

// upsertAvalaraCredentials configures Avalara tax provider credentials.
func upsertAvalaraCredentials(client metronome.Client) {
	resp, err := client.V1.Settings.UpsertAvalaraCredentials(context.TODO(), metronome.V1SettingUpsertAvalaraCredentialsParams{
		AvalaraEnvironment: metronome.V1SettingUpsertAvalaraCredentialsParamsAvalaraEnvironmentProduction,
		AvalaraUsername:    "test@metronome.com",
		AvalaraPassword:    "my_password_123",
		DeliveryMethodIDs:  []string{"9a906ebb-fbc7-42e8-8e29-53bfd2db3aca"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Avalara credentials configured: %+v\n", resp)
}
