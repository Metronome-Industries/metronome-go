package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/option"
)

func main() {
	clientWithOptions()
	perRequestOptions()
	timeoutWithContext()
	retryConfiguration()
}

// clientWithOptions demonstrates configuring the client with default request
// options that apply to every request. Options set at the client level can be
// overridden on a per-request basis.
func clientWithOptions() {
	client := metronome.NewClient(
		// Add a custom header to every request
		option.WithHeader("X-Custom-Header", "my-value"),
		// Disable retries globally
		option.WithMaxRetries(0),
	)

	customers, err := client.V1.Customers.List(context.TODO(), metronome.V1CustomerListParams{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Customers: %+v\n", customers.Data)
}

// perRequestOptions demonstrates overriding client-level options for a specific
// request. This is useful when a particular request needs different retry
// behavior, additional headers, or undocumented parameters.
func perRequestOptions() {
	client := metronome.NewClient()

	_, err := client.V1.Contracts.New(
		context.TODO(),
		metronome.V1ContractNewParams{
			CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
			StartingAt: time.Now(),
		},
		// Override retries for this specific request
		option.WithMaxRetries(5),
		// Add an undocumented field to the JSON body using sjson syntax
		option.WithJSONSet("some.json.path", map[string]string{"my": "object"}),
		// Enable debug logging for this request
		option.WithDebugLog(nil),
	)
	if err != nil {
		fmt.Printf("Error (expected in example): %s\n", err)
		return
	}
	fmt.Println("Contract created")
}

// timeoutWithContext demonstrates setting timeouts using Go's context package.
// The context timeout covers the entire request lifecycle including all retries.
// For per-retry timeouts, use option.WithRequestTimeout().
func timeoutWithContext() {
	client := metronome.NewClient()

	// This timeout covers the entire request lifecycle, including all retries.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	_, err := client.V1.Contracts.New(
		ctx,
		metronome.V1ContractNewParams{
			CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
			StartingAt: time.Now(),
		},
		// This sets the per-retry timeout (each retry gets up to 20 seconds)
		option.WithRequestTimeout(20*time.Second),
	)
	if err != nil {
		fmt.Printf("Error (expected in example): %s\n", err)
		return
	}
	fmt.Println("Contract created")
}

// retryConfiguration demonstrates configuring automatic retry behavior.
// By default, the SDK retries connection errors, 408, 409, 429, and >=500
// status codes up to 2 times with exponential backoff.
func retryConfiguration() {
	// Disable retries entirely at the client level
	client := metronome.NewClient(
		option.WithMaxRetries(0),
	)

	// Or override per-request
	_, err := client.V1.Contracts.New(
		context.TODO(),
		metronome.V1ContractNewParams{
			CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
			StartingAt: time.Now(),
		},
		option.WithMaxRetries(5), // retry up to 5 times for this request
	)
	if err != nil {
		fmt.Printf("Error (expected in example): %s\n", err)
		return
	}
	fmt.Println("Contract created")
}
