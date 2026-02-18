package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	createContract(client)
	getContract(client)
	listContracts(client)
	addManualBalanceEntry(client)
	amendContract(client)
	archiveContract(client)
	createHistoricalInvoices(client)
	getNetBalance(client)
	listBalances(client)
	getRateSchedule(client)
	getSubscriptionQuantityHistory(client)
	scheduleProServicesInvoice(client)
	setUsageFilter(client)
	updateEndDate(client)
}

// createContract creates a new contract for a customer.
func createContract(client metronome.Client) {
	resp, err := client.V1.Contracts.New(context.TODO(), metronome.V1ContractNewParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		StartingAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created contract: %+v\n", resp)
}

// getContract retrieves a specific contract.
func getContract(client metronome.Client) {
	resp, err := client.V1.Contracts.Get(context.TODO(), metronome.V1ContractGetParams{
		ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Contract: %+v\n", resp)
}

// listContracts lists all contracts for a customer.
func listContracts(client metronome.Client) {
	resp, err := client.V1.Contracts.List(context.TODO(), metronome.V1ContractListParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Contracts: %+v\n", resp)
}

// addManualBalanceEntry adds a manual balance ledger entry to a contract.
func addManualBalanceEntry(client metronome.Client) {
	err := client.V1.Contracts.AddManualBalanceEntry(context.TODO(), metronome.V1ContractAddManualBalanceEntryParams{
		ID:         "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
		Amount:     -1000,
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		Reason:     "Manual adjustment for overcharge",
		SegmentID:  "66368e29-3f97-4d15-a6e9-120897f0070a",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Manual balance entry added successfully")
}

// amendContract amends an existing contract.
func amendContract(client metronome.Client) {
	resp, err := client.V1.Contracts.Amend(context.TODO(), metronome.V1ContractAmendParams{
		ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		StartingAt: time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Amended contract: %+v\n", resp)
}

// archiveContract archives a contract.
func archiveContract(client metronome.Client) {
	resp, err := client.V1.Contracts.Archive(context.TODO(), metronome.V1ContractArchiveParams{
		ContractID:   "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:   "13117714-3f05-48e5-a6e9-a66093f13b4d",
		VoidInvoices: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Archived contract: %+v\n", resp)
}

// createHistoricalInvoices creates historical invoices for a contract.
func createHistoricalInvoices(client metronome.Client) {
	resp, err := client.V1.Contracts.NewHistoricalInvoices(context.TODO(), metronome.V1ContractNewHistoricalInvoicesParams{
		Invoices: []metronome.V1ContractNewHistoricalInvoicesParamsInvoice{{
			ContractID:         "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
			CreditTypeID:       "2714e483-4ff1-48e4-9e25-ac732e8f24f2",
			CustomerID:         "13117714-3f05-48e5-a6e9-a66093f13b4d",
			InclusiveStartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			ExclusiveEndDate:   time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
			IssueDate:          time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
		}},
		Preview: false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Historical invoices: %+v\n", resp)
}

// getNetBalance retrieves the net balance for a customer's contracts.
func getNetBalance(client metronome.Client) {
	resp, err := client.V1.Contracts.GetNetBalance(context.TODO(), metronome.V1ContractGetNetBalanceParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Net balance: %+v\n", resp)
}

// listBalances lists balance entries for a customer's contracts.
func listBalances(client metronome.Client) {
	page, err := client.V1.Contracts.ListBalances(context.TODO(), metronome.V1ContractListBalancesParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		panic(err)
	}
	for _, balance := range page.Data {
		fmt.Printf("Balance: %+v\n", balance)
	}

	// To automatically paginate, use ListBalancesAutoPaging:
	//
	// iter := client.V1.Contracts.ListBalancesAutoPaging(context.TODO(), metronome.V1ContractListBalancesParams{
	//     CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	// })
	// for iter.Next() {
	//     balance := iter.Current()
	//     fmt.Printf("Balance: %+v\n", balance)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// getRateSchedule retrieves the rate schedule for a contract.
func getRateSchedule(client metronome.Client) {
	resp, err := client.V1.Contracts.GetRateSchedule(context.TODO(), metronome.V1ContractGetRateScheduleParams{
		ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rate schedule: %+v\n", resp)
}

// getSubscriptionQuantityHistory retrieves the history of subscription quantity changes.
func getSubscriptionQuantityHistory(client metronome.Client) {
	resp, err := client.V1.Contracts.GetSubscriptionQuantityHistory(context.TODO(), metronome.V1ContractGetSubscriptionQuantityHistoryParams{
		ContractID:     "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:     "13117714-3f05-48e5-a6e9-a66093f13b4d",
		SubscriptionID: "1a824d53-bde6-4d82-96d7-6347ff227d5c",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Subscription quantity history: %+v\n", resp)
}

// scheduleProServicesInvoice schedules a professional services invoice.
func scheduleProServicesInvoice(client metronome.Client) {
	resp, err := client.V1.Contracts.ScheduleProServicesInvoice(context.TODO(), metronome.V1ContractScheduleProServicesInvoiceParams{
		ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		IssuedAt:   time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
		LineItems: []metronome.V1ContractScheduleProServicesInvoiceParamsLineItem{{
			ProfessionalServiceID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Scheduled pro services invoice: %+v\n", resp)
}

// setUsageFilter sets a usage filter on a contract.
func setUsageFilter(client metronome.Client) {
	err := client.V1.Contracts.SetUsageFilter(context.TODO(), metronome.V1ContractSetUsageFilterParams{
		ContractID:  "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:  "13117714-3f05-48e5-a6e9-a66093f13b4d",
		GroupKey:    "business_subscription_id",
		GroupValues: []string{"ID-1", "ID-2"},
		StartingAt:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Usage filter set successfully")
}

// updateEndDate updates the end date of a contract.
func updateEndDate(client metronome.Client) {
	resp, err := client.V1.Contracts.UpdateEndDate(context.TODO(), metronome.V1ContractUpdateEndDateParams{
		ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated contract end date: %+v\n", resp)
}
