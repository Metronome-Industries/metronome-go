package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	handleAPIError(client)
}

// handleAPIError demonstrates how to inspect API errors returned by the SDK.
// When the API returns a non-success status code, the error can be unwrapped
// to *metronome.Error which provides access to the status code, request/response
// dumps, and the raw JSON error body.
func handleAPIError(client metronome.Client) {
	_, err := client.V1.Contracts.New(context.TODO(), metronome.V1ContractNewParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		StartingAt: time.Now(),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			fmt.Printf("Status Code: %d\n", apierr.StatusCode)
			fmt.Printf("Request:\n%s\n", apierr.DumpRequest(true))
			fmt.Printf("Response:\n%s\n", apierr.DumpResponse(true))
			fmt.Printf("Raw JSON: %s\n", apierr.RawJSON())
		}
		// Non-API errors (e.g. network failures) are returned unwrapped.
		// For example, HTTP transport errors may be *url.Error wrapping *net.OpError.
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Println("Contract created successfully")
}
