package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	getNamedSchedule(client)
	updateNamedSchedule(client)
}

// getNamedSchedule retrieves a named schedule for a rate card.
func getNamedSchedule(client metronome.Client) {
	resp, err := client.V1.Contracts.NamedSchedules.Get(context.TODO(), metronome.V1ContractNamedScheduleGetParams{
		RateCardID:   "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		ScheduleName: "my-schedule",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Named schedule: %+v\n", resp)
}

// updateNamedSchedule updates a named schedule for a rate card.
func updateNamedSchedule(client metronome.Client) {
	err := client.V1.Contracts.NamedSchedules.Update(context.TODO(), metronome.V1ContractNamedScheduleUpdateParams{
		RateCardID:   "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		ScheduleName: "my-schedule",
		StartingAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Value: map[string]any{
			"my_key": "my_value",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Named schedule updated successfully")
}
