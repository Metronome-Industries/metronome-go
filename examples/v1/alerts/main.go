package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	createAlert(client)
	archiveAlert(client)
}

// createAlert creates a new threshold alert notification.
func createAlert(client metronome.Client) {
	resp, err := client.V1.Alerts.New(context.TODO(), metronome.V1AlertNewParams{
		AlertType: metronome.V1AlertNewParamsAlertTypeSpendThresholdReached,
		Name:      "$100 spend threshold reached",
		Threshold: 10000,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created alert: %+v\n", resp)
}

// archiveAlert disables and archives an alert.
func archiveAlert(client metronome.Client) {
	resp, err := client.V1.Alerts.Archive(context.TODO(), metronome.V1AlertArchiveParams{
		ID: "8deed800-1b7a-495d-a207-6c52bac54dc9",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Archived alert: %+v\n", resp)
}
