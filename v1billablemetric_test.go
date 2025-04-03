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

func TestV1BillableMetricNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.BillableMetrics.New(context.TODO(), metronome.V1BillableMetricNewParams{
		Name:            metronome.F("CPU Hours"),
		AggregationKey:  metronome.F("cpu_hours"),
		AggregationType: metronome.F(metronome.V1BillableMetricNewParamsAggregationTypeCount),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		EventTypeFilter: metronome.F(shared.EventTypeFilterParam{
			InValues:    metronome.F([]string{"cpu_usage"}),
			NotInValues: metronome.F([]string{"string"}),
		}),
		GroupKeys: metronome.F([][]string{{"region"}, {"machine_type"}}),
		PropertyFilters: metronome.F([]shared.PropertyFilterParam{{
			Name:        metronome.F("cpu_hours"),
			Exists:      metronome.F(true),
			InValues:    metronome.F([]string{"string"}),
			NotInValues: metronome.F([]string{"string"}),
		}, {
			Name:        metronome.F("region"),
			Exists:      metronome.F(true),
			InValues:    metronome.F([]string{"EU", "NA"}),
			NotInValues: metronome.F([]string{"string"}),
		}, {
			Name:        metronome.F("machine_type"),
			Exists:      metronome.F(true),
			InValues:    metronome.F([]string{"slow", "fast"}),
			NotInValues: metronome.F([]string{"string"}),
		}}),
		Sql: metronome.F("sql"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BillableMetricGet(t *testing.T) {
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
	_, err := client.V1.BillableMetrics.Get(context.TODO(), metronome.V1BillableMetricGetParams{
		BillableMetricID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BillableMetricListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.BillableMetrics.List(context.TODO(), metronome.V1BillableMetricListParams{
		IncludeArchived: metronome.F(true),
		Limit:           metronome.F(int64(1)),
		NextPage:        metronome.F("next_page"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BillableMetricArchive(t *testing.T) {
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
	_, err := client.V1.BillableMetrics.Archive(context.TODO(), metronome.V1BillableMetricArchiveParams{
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
