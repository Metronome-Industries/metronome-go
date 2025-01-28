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
	_, err := client.Usage.Ingest(context.TODO(), metronome.UsageIngestParams{
		Usage: []metronome.UsageIngestParamsUsage{{
			CustomerID:    metronome.F("team@example.com"),
			EventType:     metronome.F("heartbeat"),
			Timestamp:     metronome.F("2021-01-01T00:00:00Z"),
			TransactionID: metronome.F("2021-01-01T00:00:00Z_cluster42"),
		}},
	})
	if err != nil {
		t.Error(err)
	}
}
