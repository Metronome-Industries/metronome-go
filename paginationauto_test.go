// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"os"
	"testing"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/v3/option"
)

func TestAutoPagination(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	iter := client.V1.Contracts.Products.ListAutoPaging(context.TODO(), metronome.V1ContractProductListParams{})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		product := iter.Current()
		t.Logf("%+v\n", product.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
