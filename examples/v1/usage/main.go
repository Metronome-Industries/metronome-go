package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	ingestUsage(client)
	listUsage(client)
	listUsageWithGroups(client)
	searchUsage(client)
}

// ingestUsage sends usage events to Metronome for billing.
func ingestUsage(client metronome.Client) {
	err := client.V1.Usage.Ingest(context.TODO(), metronome.V1UsageIngestParams{
		Usage: []metronome.V1UsageIngestParamsUsage{{
			TransactionID: "90e9401f-0f8c-4cd3-9a9f-d6beb56d8d72",
			CustomerID:    "team@example.com",
			EventType:     "heartbeat",
			Timestamp:     "2024-01-01T00:00:00Z",
			Properties: map[string]any{
				"cluster_id":  "42",
				"cpu_seconds": 60,
				"region":      "Europe",
			},
		}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Usage ingested successfully")
}

// listUsage retrieves aggregated usage data for a time window.
func listUsage(client metronome.Client) {
	page, err := client.V1.Usage.List(context.TODO(), metronome.V1UsageListParams{
		StartingOn:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndingBefore: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
		WindowSize:   metronome.V1UsageListParamsWindowSizeDay,
		CustomerIDs:  []string{"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"},
	})
	if err != nil {
		panic(err)
	}
	for _, usage := range page.Data {
		fmt.Printf("Metric: %s, Customer: %s, Value: %f\n",
			usage.BillableMetricName, usage.CustomerID, usage.Value)
	}

	// To automatically paginate through all results, use ListAutoPaging:
	//
	// iter := client.V1.Usage.ListAutoPaging(context.TODO(), metronome.V1UsageListParams{...})
	// for iter.Next() {
	//     usage := iter.Current()
	//     fmt.Printf("Value: %f\n", usage.Value)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// listUsageWithGroups retrieves usage data grouped by a property key.
func listUsageWithGroups(client metronome.Client) {
	page, err := client.V1.Usage.ListWithGroups(context.TODO(), metronome.V1UsageListWithGroupsParams{
		BillableMetricID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:       "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		WindowSize:       metronome.V1UsageListWithGroupsParamsWindowSizeDay,
		GroupBy: metronome.V1UsageListWithGroupsParamsGroupBy{
			Key:    "region",
			Values: []string{"US", "Europe"},
		},
	})
	if err != nil {
		panic(err)
	}
	for _, usage := range page.Data {
		fmt.Printf("Group: %s=%s, Value: %f\n",
			usage.GroupKey, usage.GroupValue, usage.Value)
	}

	// To automatically paginate, use ListWithGroupsAutoPaging:
	//
	// iter := client.V1.Usage.ListWithGroupsAutoPaging(context.TODO(), metronome.V1UsageListWithGroupsParams{...})
	// for iter.Next() {
	//     usage := iter.Current()
	//     fmt.Printf("Value: %f\n", usage.Value)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// searchUsage searches for specific usage events by transaction ID.
func searchUsage(client metronome.Client) {
	events, err := client.V1.Usage.Search(context.TODO(), metronome.V1UsageSearchParams{
		TransactionIDs: []string{"90e9401f-0f8c-4cd3-9a9f-d6beb56d8d72"},
	})
	if err != nil {
		panic(err)
	}
	for _, event := range *events {
		fmt.Printf("Event: %s, Type: %s, Customer: %s\n",
			event.TransactionID, event.EventType, event.CustomerID)
	}
}
