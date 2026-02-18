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

// getNamedSchedule retrieves a named schedule for a customer.
func getNamedSchedule(client metronome.Client) {
	resp, err := client.V1.Customers.NamedSchedules.Get(context.TODO(), metronome.V1CustomerNamedScheduleGetParams{
		CustomerID:   "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		ScheduleName: "my-schedule",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Named schedule: %+v\n", resp)
}

// updateNamedSchedule updates a named schedule for a customer.
func updateNamedSchedule(client metronome.Client) {
	err := client.V1.Customers.NamedSchedules.Update(context.TODO(), metronome.V1CustomerNamedScheduleUpdateParams{
		CustomerID:   "9b85c1c1-5238-4f2a-a409-61412905e1e1",
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
