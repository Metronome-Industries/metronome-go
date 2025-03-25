// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"os"
	"testing"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/option"
)

func TestUsage(t *testing.T) {
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
	err := client.V1.Usage.Ingest(context.TODO(), metronome.V1UsageIngestParams{
		Usage: []metronome.V1UsageIngestParamsUsage{{
			TransactionID: metronome.F("90e9401f-0f8c-4cd3-9a9f-d6beb56d8d72"),
			CustomerID:    metronome.F("team@example.com"),
			EventType:     metronome.F("heartbeat"),
			Timestamp:     metronome.F("2024-01-01T00:00:00Z"),
			Properties: metronome.F(map[string]interface{}{
				"cluster_id":  "42",
				"cpu_seconds": 60,
				"region":      "Europe",
			}),
		}},
	})
	if err != nil {
		t.Error(err)
	}
}
