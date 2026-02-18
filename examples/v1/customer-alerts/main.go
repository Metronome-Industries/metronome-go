package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	getCustomerAlert(client)
	listCustomerAlerts(client)
	resetCustomerAlert(client)
}

// getCustomerAlert retrieves a specific alert for a customer.
func getCustomerAlert(client metronome.Client) {
	resp, err := client.V1.Customers.Alerts.Get(context.TODO(), metronome.V1CustomerAlertGetParams{
		AlertID:    "8deed800-1b7a-495d-a207-6c52bac54dc9",
		CustomerID: "9b85c1c1-5238-4f2a-a409-61412905e1e1",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Customer alert: %+v\n", resp)
}

// listCustomerAlerts lists all alerts for a customer.
func listCustomerAlerts(client metronome.Client) {
	page, err := client.V1.Customers.Alerts.List(context.TODO(), metronome.V1CustomerAlertListParams{
		CustomerID:    "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		AlertStatuses: []string{"ENABLED"},
	})
	if err != nil {
		panic(err)
	}
	for _, alert := range page.Data {
		fmt.Printf("Alert: %s (Status: %s)\n", alert.Alert.Name, alert.CustomerStatus)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.Customers.Alerts.ListAutoPaging(context.TODO(), metronome.V1CustomerAlertListParams{
	//     CustomerID: "9b85c1c1-5238-4f2a-a409-61412905e1e1",
	// })
	// for iter.Next() {
	//     alert := iter.Current()
	//     fmt.Printf("Alert: %s\n", alert.ID)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// resetCustomerAlert resets the state of a customer alert.
func resetCustomerAlert(client metronome.Client) {
	err := client.V1.Customers.Alerts.Reset(context.TODO(), metronome.V1CustomerAlertResetParams{
		AlertID:    "5e8691bf-b22a-4672-922d-f80eee940f01",
		CustomerID: "4c83caf3-8af4-44e2-9aeb-e290531726d9",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Customer alert reset successfully")
}
