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
		EndingBefore: metronome.F(time.Now()),
		StartingOn:   metronome.F(time.Now()),
		WindowSize:   metronome.F(metronome.V1UsageListParamsWindowSizeHour),
		NextPage:     metronome.F("next_page"),
		BillableMetrics: metronome.F([]metronome.V1UsageListParamsBillableMetric{{
			ID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			GroupBy: metronome.F(metronome.V1UsageListParamsBillableMetricsGroupBy{
				Key:    metronome.F("key"),
				Values: metronome.F([]string{"x"}),
			}),
		}}),
		CustomerIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
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
			CustomerID:    metronome.F("team@example.com"),
			EventType:     metronome.F("heartbeat"),
			Timestamp:     metronome.F("2021-01-01T00:00:00Z"),
			TransactionID: metronome.F("2021-01-01T00:00:00Z_cluster42"),
			Properties: metronome.F(map[string]interface{}{
				"cluster_id":  "bar",
				"cpu_seconds": "bar",
				"region":      "bar",
			}),
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
		BillableMetricID: metronome.F("222796fd-d29c-429e-89b2-549fabda4ed6"),
		CustomerID:       metronome.F("04ca7e72-4229-4a6e-ab11-9f7376fccbcb"),
		WindowSize:       metronome.F(metronome.V1UsageListWithGroupsParamsWindowSizeHour),
		Limit:            metronome.F(int64(1)),
		NextPage:         metronome.F("next_page"),
		CurrentPeriod:    metronome.F(true),
		EndingBefore:     metronome.F(time.Now()),
		GroupBy: metronome.F(metronome.V1UsageListWithGroupsParamsGroupBy{
			Key:    metronome.F("region"),
			Values: metronome.F([]string{"US-East", "US-West", "EU-Central"}),
		}),
		StartingOn: metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
