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

	createProduct(client)
	getProduct(client)
	updateProduct(client)
	listProducts(client)
	archiveProduct(client)
}

// createProduct creates a new contract product.
func createProduct(client metronome.Client) {
	resp, err := client.V1.Contracts.Products.New(context.TODO(), metronome.V1ContractProductNewParams{
		Name: "CPU Hours",
		Type: metronome.V1ContractProductNewParamsTypeUsage,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created product: %+v\n", resp)
}

// getProduct retrieves a contract product by ID.
func getProduct(client metronome.Client) {
	resp, err := client.V1.Contracts.Products.Get(context.TODO(), metronome.V1ContractProductGetParams{
		ID: shared.IDParam{
			ID: "d84e7f4e-7a70-4fe4-be02-7a5027beffcc",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %+v\n", resp)
}

// updateProduct updates a contract product.
func updateProduct(client metronome.Client) {
	resp, err := client.V1.Contracts.Products.Update(context.TODO(), metronome.V1ContractProductUpdateParams{
		ProductID:  "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		StartingAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Name:       metronome.String("CPU Hours (Updated)"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated product: %+v\n", resp)
}

// listProducts lists all contract products.
func listProducts(client metronome.Client) {
	page, err := client.V1.Contracts.Products.List(context.TODO(), metronome.V1ContractProductListParams{})
	if err != nil {
		panic(err)
	}
	for _, product := range page.Data {
		fmt.Printf("Product: %+v\n", product)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.Contracts.Products.ListAutoPaging(context.TODO(), metronome.V1ContractProductListParams{})
	// for iter.Next() {
	//     product := iter.Current()
	//     fmt.Printf("Product: %+v\n", product)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// archiveProduct archives a contract product.
func archiveProduct(client metronome.Client) {
	resp, err := client.V1.Contracts.Products.Archive(context.TODO(), metronome.V1ContractProductArchiveParams{
		ProductID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Archived product: %+v\n", resp)
}
