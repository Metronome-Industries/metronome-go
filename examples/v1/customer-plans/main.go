package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	listCustomerPlans(client)
	addPlanToCustomer(client)
	endCustomerPlan(client)
	listPriceAdjustments(client)
}

// listCustomerPlans lists all plans associated with a customer.
func listCustomerPlans(client metronome.Client) {
	page, err := client.V1.Customers.Plans.List(context.TODO(), metronome.V1CustomerPlanListParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	})
	if err != nil {
		panic(err)
	}
	for _, plan := range page.Data {
		fmt.Printf("Plan: %+v\n", plan)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.Customers.Plans.ListAutoPaging(context.TODO(), metronome.V1CustomerPlanListParams{
	//     CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	// })
	// for iter.Next() {
	//     plan := iter.Current()
	//     fmt.Printf("Plan: %+v\n", plan)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// addPlanToCustomer adds a plan to a customer with a start date.
func addPlanToCustomer(client metronome.Client) {
	resp, err := client.V1.Customers.Plans.Add(context.TODO(), metronome.V1CustomerPlanAddParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		PlanID:     "d2c06dae-9549-4d7d-bc04-b78dd3d241b8",
		StartingOn: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added plan to customer: %+v\n", resp)
}

// endCustomerPlan ends a customer's plan.
func endCustomerPlan(client metronome.Client) {
	resp, err := client.V1.Customers.Plans.End(context.TODO(), metronome.V1CustomerPlanEndParams{
		CustomerID:     "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerPlanID: "7aa11640-0703-4600-8eb9-293f535a6b74",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ended customer plan: %+v\n", resp)
}

// listPriceAdjustments lists price adjustments for a customer plan.
func listPriceAdjustments(client metronome.Client) {
	page, err := client.V1.Customers.Plans.ListPriceAdjustments(context.TODO(), metronome.V1CustomerPlanListPriceAdjustmentsParams{
		CustomerID:     "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerPlanID: "7aa11640-0703-4600-8eb9-293f535a6b74",
	})
	if err != nil {
		panic(err)
	}
	for _, adj := range page.Data {
		fmt.Printf("Price adjustment: %+v\n", adj)
	}

	// To automatically paginate, use ListPriceAdjustmentsAutoPaging:
	//
	// iter := client.V1.Customers.Plans.ListPriceAdjustmentsAutoPaging(context.TODO(), metronome.V1CustomerPlanListPriceAdjustmentsParams{
	//     CustomerID:     "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	//     CustomerPlanID: "7aa11640-0703-4600-8eb9-293f535a6b74",
	// })
	// for iter.Next() {
	//     adj := iter.Current()
	//     fmt.Printf("Adjustment: %+v\n", adj)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}
