// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/option"
)

func TestV1UsageListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Usage.List(context.TODO(), metronome.V1UsageListParams{
		EndingBefore: time.Now(),
		StartingOn:   time.Now(),
		WindowSize:   metronome.V1UsageListParamsWindowSizeHour,
		NextPage:     metronome.String("next_page"),
		BillableMetrics: []metronome.V1UsageListParamsBillableMetric{{
			ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			GroupBy: metronome.V1UsageListParamsBillableMetricGroupBy{
				Key:    "key",
				Values: []string{"x"},
			},
		}},
		CustomerIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1UsageIngestWithOptionalParams(t *testing.T) {
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
			CustomerID:    "team@example.com",
			EventType:     "heartbeat",
			Timestamp:     "2021-01-01T00:00:00Z",
			TransactionID: "2021-01-01T00:00:00Z_cluster42",
			Properties: map[string]any{
				"cluster_id":  "bar",
				"cpu_seconds": "bar",
				"region":      "bar",
			},
		}},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1UsageListWithGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Usage.ListWithGroups(context.TODO(), metronome.V1UsageListWithGroupsParams{
		BillableMetricID: "222796fd-d29c-429e-89b2-549fabda4ed6",
		CustomerID:       "04ca7e72-4229-4a6e-ab11-9f7376fccbcb",
		WindowSize:       metronome.V1UsageListWithGroupsParamsWindowSizeHour,
		Limit:            metronome.Int(1),
		NextPage:         metronome.String("next_page"),
		CurrentPeriod:    metronome.Bool(true),
		EndingBefore:     metronome.Time(time.Now()),
		GroupBy: metronome.V1UsageListWithGroupsParamsGroupBy{
			Key:    "region",
			Values: []string{"US-East", "US-West", "EU-Central"},
		},
		StartingOn: metronome.Time(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1UsageSearch(t *testing.T) {
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
	_, err := client.V1.Usage.Search(context.TODO(), metronome.V1UsageSearchParams{
		TransactionIDs: []string{"2021-01-01T00:00:00Z_cluster42"},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
