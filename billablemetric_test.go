// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
)

func TestBillableMetricNewWithOptionalParams(t *testing.T) {
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
	_, err := client.BillableMetrics.New(context.TODO(), metronome.BillableMetricNewParams{
		AggregationType: metronome.F(metronome.BillableMetricNewParamsAggregationTypeCount),
		Name:            metronome.F("CPU Hours"),
		AggregationKey:  metronome.F("cpu_hours"),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		EventTypeFilter: metronome.F(metronome.BillableMetricNewParamsEventTypeFilter{
			InValues:    metronome.F([]string{"cpu_usage"}),
			NotInValues: metronome.F([]string{"string", "string", "string"}),
		}),
		GroupKeys: metronome.F([][]string{{"region"}, {"machine_type"}}),
		PropertyFilters: metronome.F([]metronome.BillableMetricNewParamsPropertyFilter{{
			Name:        metronome.F("cpu_hours"),
			Exists:      metronome.F(true),
			InValues:    metronome.F([]string{"string", "string", "string"}),
			NotInValues: metronome.F([]string{"string", "string", "string"}),
		}, {
			Name:        metronome.F("region"),
			Exists:      metronome.F(true),
			InValues:    metronome.F([]string{"EU", "NA"}),
			NotInValues: metronome.F([]string{"string", "string", "string"}),
		}, {
			Name:        metronome.F("machine_type"),
			Exists:      metronome.F(true),
			InValues:    metronome.F([]string{"slow", "fast"}),
			NotInValues: metronome.F([]string{"string", "string", "string"}),
		}}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBillableMetricGet(t *testing.T) {
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
	_, err := client.BillableMetrics.Get(context.TODO(), "13117714-3f05-48e5-a6e9-a66093f13b4d")
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBillableMetricListWithOptionalParams(t *testing.T) {
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
	_, err := client.BillableMetrics.List(
		context.TODO(),
		"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		metronome.BillableMetricListParams{
			Limit:         metronome.F(int64(1)),
			NextPage:      metronome.F("next_page"),
			OnCurrentPlan: metronome.F(true),
		},
	)
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBillableMetricArchive(t *testing.T) {
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
	_, err := client.BillableMetrics.Archive(context.TODO(), metronome.BillableMetricArchiveParams{
		ID: shared.IDParam{
			ID: metronome.F("8deed800-1b7a-495d-a207-6c52bac54dc9"),
		},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
