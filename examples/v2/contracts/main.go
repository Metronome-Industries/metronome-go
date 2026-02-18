package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	getContract(client)
	listContracts(client)
	editContract(client)
	editCommit(client)
	editCredit(client)
	getEditHistory(client)
}

// getContract retrieves a contract using the V2 API.
func getContract(client metronome.Client) {
	resp, err := client.V2.Contracts.Get(context.TODO(), metronome.V2ContractGetParams{
		ContractID:     "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:     "13117714-3f05-48e5-a6e9-a66093f13b4d",
		IncludeBalance: metronome.Bool(true),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Contract: %+v\n", resp)
}

// listContracts lists contracts for a customer using the V2 API.
func listContracts(client metronome.Client) {
	resp, err := client.V2.Contracts.List(context.TODO(), metronome.V2ContractListParams{
		CustomerID:     "13117714-3f05-48e5-a6e9-a66093f13b4d",
		IncludeBalance: metronome.Bool(true),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Contracts: %+v\n", resp)
}

// editContract edits a contract using the V2 API.
func editContract(client metronome.Client) {
	resp, err := client.V2.Contracts.Edit(context.TODO(), metronome.V2ContractEditParams{
		ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Edited contract: %+v\n", resp)
}

// editCommit edits a commit within a contract using the V2 API.
func editCommit(client metronome.Client) {
	resp, err := client.V2.Contracts.EditCommit(context.TODO(), metronome.V2ContractEditCommitParams{
		CommitID:   "5e7e82cf-ccb7-428c-a96f-a8e4f67af822",
		CustomerID: "4c91c473-fc12-445a-9c38-40421d47023f",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Edited commit: %+v\n", resp)
}

// editCredit edits a credit within a contract using the V2 API.
func editCredit(client metronome.Client) {
	resp, err := client.V2.Contracts.EditCredit(context.TODO(), metronome.V2ContractEditCreditParams{
		CreditID:   "5e7e82cf-ccb7-428c-a96f-a8e4f67af822",
		CustomerID: "4c91c473-fc12-445a-9c38-40421d47023f",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Edited credit: %+v\n", resp)
}

// getEditHistory retrieves the edit history for a contract using the V2 API.
func getEditHistory(client metronome.Client) {
	resp, err := client.V2.Contracts.GetEditHistory(context.TODO(), metronome.V2ContractGetEditHistoryParams{
		ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Edit history: %+v\n", resp)
}
