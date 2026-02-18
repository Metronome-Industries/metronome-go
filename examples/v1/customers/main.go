package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/shared"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	createCustomer(client)
	getCustomer(client)
	listCustomers(client)
	archiveCustomer(client)
	listBillableMetrics(client)
	listCosts(client)
	previewEvents(client)
	getBillingConfigurations(client)
	setBillingConfigurations(client)
	setIngestAliases(client)
	setName(client)
	updateConfig(client)
}

// createCustomer creates a new customer in Metronome.
func createCustomer(client metronome.Client) {
	customer, err := client.V1.Customers.New(context.TODO(), metronome.V1CustomerNewParams{
		Name:          "Acme, Inc.",
		IngestAliases: []string{"acme@example.com"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created customer: %+v\n", customer)
}

// getCustomer retrieves a customer by their ID.
func getCustomer(client metronome.Client) {
	customer, err := client.V1.Customers.Get(context.TODO(), metronome.V1CustomerGetParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Customer: %s (ID: %s)\n", customer.Data.Name, customer.Data.ID)
}

// listCustomers lists all customers with optional filtering.
func listCustomers(client metronome.Client) {
	page, err := client.V1.Customers.List(context.TODO(), metronome.V1CustomerListParams{
		Limit: metronome.Int(10),
	})
	if err != nil {
		panic(err)
	}
	for _, customer := range page.Data {
		fmt.Printf("Customer: %s (ID: %s)\n", customer.Name, customer.ID)
	}

	// To automatically paginate through all results, use ListAutoPaging:
	//
	// iter := client.V1.Customers.ListAutoPaging(context.TODO(), metronome.V1CustomerListParams{})
	// for iter.Next() {
	//     customer := iter.Current()
	//     fmt.Printf("Customer: %s\n", customer.Name)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// archiveCustomer archives a customer by their ID.
func archiveCustomer(client metronome.Client) {
	resp, err := client.V1.Customers.Archive(context.TODO(), metronome.V1CustomerArchiveParams{
		ID: shared.IDParam{
			ID: "8deed800-1b7a-495d-a207-6c52bac54dc9",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Archived customer: %+v\n", resp)
}

// listBillableMetrics lists billable metrics associated with a customer.
func listBillableMetrics(client metronome.Client) {
	page, err := client.V1.Customers.ListBillableMetrics(context.TODO(), metronome.V1CustomerListBillableMetricsParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	})
	if err != nil {
		panic(err)
	}
	for _, metric := range page.Data {
		fmt.Printf("Metric: %s (ID: %s)\n", metric.Name, metric.ID)
	}

	// To automatically paginate, use ListBillableMetricsAutoPaging:
	//
	// iter := client.V1.Customers.ListBillableMetricsAutoPaging(context.TODO(), metronome.V1CustomerListBillableMetricsParams{
	//     CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	// })
	// for iter.Next() {
	//     metric := iter.Current()
	//     fmt.Printf("Metric: %s\n", metric.Name)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// listCosts retrieves costs for a customer over a time range.
func listCosts(client metronome.Client) {
	page, err := client.V1.Customers.ListCosts(context.TODO(), metronome.V1CustomerListCostsParams{
		CustomerID:   "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		StartingOn:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndingBefore: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		panic(err)
	}
	for _, cost := range page.Data {
		fmt.Printf("Cost: %+v\n", cost)
	}

	// To automatically paginate, use ListCostsAutoPaging:
	//
	// iter := client.V1.Customers.ListCostsAutoPaging(context.TODO(), metronome.V1CustomerListCostsParams{...})
	// for iter.Next() {
	//     cost := iter.Current()
	//     fmt.Printf("Cost: %+v\n", cost)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// previewEvents previews how usage events would be processed for a customer.
func previewEvents(client metronome.Client) {
	resp, err := client.V1.Customers.PreviewEvents(context.TODO(), metronome.V1CustomerPreviewEventsParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		Events: []metronome.V1CustomerPreviewEventsParamsEvent{{
			EventType: "heartbeat",
			Properties: map[string]any{
				"cpu_hours":       "12",
				"memory_gb_hours": "8",
			},
			Timestamp:     metronome.String("2024-01-01T00:00:00Z"),
			TransactionID: metronome.String("preview-txn-001"),
		}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Preview result: %+v\n", resp)
}

// getBillingConfigurations retrieves billing provider configurations for a customer.
func getBillingConfigurations(client metronome.Client) {
	resp, err := client.V1.Customers.GetBillingConfigurations(context.TODO(), metronome.V1CustomerGetBillingConfigurationsParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Billing configurations: %+v\n", resp)
}

// setBillingConfigurations configures billing providers for a customer.
func setBillingConfigurations(client metronome.Client) {
	resp, err := client.V1.Customers.SetBillingConfigurations(context.TODO(), metronome.V1CustomerSetBillingConfigurationsParams{
		Data: []metronome.V1CustomerSetBillingConfigurationsParamsData{{
			BillingProvider: "stripe",
			CustomerID:      "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
			Configuration: map[string]any{
				"stripe_customer_id":       "cus_abc123",
				"stripe_collection_method": "charge_automatically",
			},
			DeliveryMethod: "direct_to_billing_provider",
		}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Set billing configurations: %+v\n", resp)
}

// setIngestAliases sets the ingest aliases for a customer.
func setIngestAliases(client metronome.Client) {
	err := client.V1.Customers.SetIngestAliases(context.TODO(), metronome.V1CustomerSetIngestAliasesParams{
		CustomerID:    "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		IngestAliases: []string{"team@example.com", "billing@example.com"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Ingest aliases updated successfully")
}

// setName updates the name of a customer.
func setName(client metronome.Client) {
	resp, err := client.V1.Customers.SetName(context.TODO(), metronome.V1CustomerSetNameParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		Name:       "Acme Corporation",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated customer name: %+v\n", resp)
}

// updateConfig updates configuration settings for a customer.
func updateConfig(client metronome.Client) {
	err := client.V1.Customers.UpdateConfig(context.TODO(), metronome.V1CustomerUpdateConfigParams{
		CustomerID:                 "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		LeaveStripeInvoicesInDraft: metronome.Bool(true),
		SalesforceAccountID:        metronome.String("0015500001WO1ZiABL"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Customer config updated successfully")
}
