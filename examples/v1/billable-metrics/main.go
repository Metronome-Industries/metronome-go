package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/shared"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	createBillableMetric(client)
	getBillableMetric(client)
	listBillableMetrics(client)
	archiveBillableMetric(client)
}

// createBillableMetric creates a new billable metric.
func createBillableMetric(client metronome.Client) {
	resp, err := client.V1.BillableMetrics.New(context.TODO(), metronome.V1BillableMetricNewParams{
		Name:            "CPU Hours",
		AggregationType: metronome.V1BillableMetricNewParamsAggregationTypeSum,
		AggregationKey:  metronome.String("cpu_hours"),
		EventTypeFilter: shared.EventTypeFilterParam{
			InValues: []string{"cpu_usage"},
		},
		GroupKeys: [][]string{{"region"}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created billable metric: %+v\n", resp)
}

// getBillableMetric retrieves a billable metric by ID.
func getBillableMetric(client metronome.Client) {
	resp, err := client.V1.BillableMetrics.Get(context.TODO(), metronome.V1BillableMetricGetParams{
		BillableMetricID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Billable metric: %+v\n", resp)
}

// listBillableMetrics lists all billable metrics.
func listBillableMetrics(client metronome.Client) {
	page, err := client.V1.BillableMetrics.List(context.TODO(), metronome.V1BillableMetricListParams{})
	if err != nil {
		panic(err)
	}
	for _, metric := range page.Data {
		fmt.Printf("Metric: %s (ID: %s)\n", metric.Name, metric.ID)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.BillableMetrics.ListAutoPaging(context.TODO(), metronome.V1BillableMetricListParams{})
	// for iter.Next() {
	//     metric := iter.Current()
	//     fmt.Printf("Metric: %s\n", metric.Name)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// archiveBillableMetric archives a billable metric.
func archiveBillableMetric(client metronome.Client) {
	resp, err := client.V1.BillableMetrics.Archive(context.TODO(), metronome.V1BillableMetricArchiveParams{
		ID: shared.IDParam{
			ID: "8deed800-1b7a-495d-a207-6c52bac54dc9",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Archived billable metric: %+v\n", resp)
}
