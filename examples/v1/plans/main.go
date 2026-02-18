package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	listPlans(client)
	getPlanDetails(client)
	listPlanCharges(client)
	listPlanCustomers(client)
}

// listPlans retrieves all plans in your Metronome account.
func listPlans(client metronome.Client) {
	page, err := client.V1.Plans.List(context.TODO(), metronome.V1PlanListParams{
		Limit: metronome.Int(10),
	})
	if err != nil {
		panic(err)
	}
	for _, plan := range page.Data {
		fmt.Printf("Plan: %s (ID: %s)\n", plan.Name, plan.ID)
	}

	// To automatically paginate through all results, use ListAutoPaging:
	//
	// iter := client.V1.Plans.ListAutoPaging(context.TODO(), metronome.V1PlanListParams{})
	// for iter.Next() {
	//     plan := iter.Current()
	//     fmt.Printf("Plan: %s\n", plan.Name)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// getPlanDetails retrieves the details of a specific plan.
func getPlanDetails(client metronome.Client) {
	details, err := client.V1.Plans.GetDetails(context.TODO(), metronome.V1PlanGetDetailsParams{
		PlanID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Plan details: %+v\n", details)
}

// listPlanCharges lists all charges associated with a plan.
func listPlanCharges(client metronome.Client) {
	page, err := client.V1.Plans.ListCharges(context.TODO(), metronome.V1PlanListChargesParams{
		PlanID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	})
	if err != nil {
		panic(err)
	}
	for _, charge := range page.Data {
		fmt.Printf("Charge: %+v\n", charge)
	}

	// To automatically paginate, use ListChargesAutoPaging:
	//
	// iter := client.V1.Plans.ListChargesAutoPaging(context.TODO(), metronome.V1PlanListChargesParams{
	//     PlanID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	// })
	// for iter.Next() {
	//     charge := iter.Current()
	//     fmt.Printf("Charge: %+v\n", charge)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// listPlanCustomers lists all customers on a specific plan.
func listPlanCustomers(client metronome.Client) {
	page, err := client.V1.Plans.ListCustomers(context.TODO(), metronome.V1PlanListCustomersParams{
		PlanID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		Status: metronome.V1PlanListCustomersParamsStatusActive,
	})
	if err != nil {
		panic(err)
	}
	for _, customer := range page.Data {
		fmt.Printf("Customer: %+v\n", customer)
	}

	// To automatically paginate, use ListCustomersAutoPaging:
	//
	// iter := client.V1.Plans.ListCustomersAutoPaging(context.TODO(), metronome.V1PlanListCustomersParams{
	//     PlanID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	// })
	// for iter.Next() {
	//     customer := iter.Current()
	//     fmt.Printf("Customer: %+v\n", customer)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}
