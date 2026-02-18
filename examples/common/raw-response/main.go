package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/option"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	accessResponseHeaders(client)
}

// accessResponseHeaders demonstrates using option.WithResponseInto to access
// the raw HTTP response alongside the parsed response body. This is useful
// when you need to inspect response headers, status codes, or other HTTP-level
// details that aren't part of the API response body.
func accessResponseHeaders(client metronome.Client) {
	var response *http.Response

	_, err := client.V1.Contracts.New(
		context.TODO(),
		metronome.V1ContractNewParams{
			CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
			StartingAt: time.Now(),
		},
		option.WithResponseInto(&response),
	)
	if err != nil {
		fmt.Printf("Error (expected in example): %s\n", err)
		// Note: response may still be populated even on error
		if response != nil {
			fmt.Printf("Status Code: %d\n", response.StatusCode)
			fmt.Printf("Content-Type: %s\n", response.Header.Get("Content-Type"))
		}
		return
	}

	fmt.Printf("Status Code: %d\n", response.StatusCode)
	fmt.Printf("Headers: %+v\n", response.Header)
}
