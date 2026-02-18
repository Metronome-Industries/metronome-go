package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/option"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	undocumentedEndpoint(client)
	undocumentedParams(client)
	undocumentedResponseFields(client)
}

// undocumentedEndpoint demonstrates making requests to API endpoints that
// aren't yet covered by the SDK's typed methods. The client provides Get,
// Post, Put, Patch, and Delete methods that accept arbitrary params and
// response types. Request options like retries still apply.
func undocumentedEndpoint(client metronome.Client) {
	// params can be an io.Reader, []byte, encoding/json serializable object,
	// or a Params struct defined in this library.
	params := map[string]any{
		"key": "value",
	}

	// result can be []byte, *http.Response, encoding/json deserializable object,
	// or a model defined in this library.
	var result *http.Response

	err := client.Post(context.Background(), "/custom/endpoint", params, &result)
	if err != nil {
		fmt.Printf("Error (expected in example): %s\n", err)
		return
	}
	fmt.Printf("Status: %d\n", result.StatusCode)
}

// undocumentedParams demonstrates sending undocumented request parameters
// alongside typed params. Use option.WithJSONSet for body parameters (using
// sjson path syntax) and option.WithQuerySet for query parameters. These are
// useful when the API accepts parameters not yet reflected in the SDK.
func undocumentedParams(client metronome.Client) {
	customers, err := client.V1.Customers.List(
		context.TODO(),
		metronome.V1CustomerListParams{},
		// Add an undocumented query parameter
		option.WithQuery("custom_filter", "active"),
	)
	if err != nil {
		fmt.Printf("Error (expected in example): %s\n", err)
		return
	}
	fmt.Printf("Customers: %+v\n", customers.Data)
}

// undocumentedResponseFields demonstrates accessing response fields that
// aren't part of the typed response struct. Any extra JSON properties in the
// response are captured in the JSON.ExtraFields map, and the full raw JSON
// is available via RawJSON().
func undocumentedResponseFields(client metronome.Client) {
	page, err := client.V1.Customers.List(context.TODO(), metronome.V1CustomerListParams{})
	if err != nil {
		panic(err)
	}

	for _, customer := range page.Data {
		// Access the raw JSON of the entire response object
		fmt.Printf("Raw JSON: %s\n", customer.RawJSON())

		// Access undocumented fields from the JSON response
		for key, field := range customer.JSON.ExtraFields {
			fmt.Printf("Extra field %q: %s\n", key, field.Raw())
		}

		// Check if an optional field was actually present in the response
		if customer.JSON.Name.Valid() {
			fmt.Printf("Name is present: %s\n", customer.Name)
		}
	}
}
