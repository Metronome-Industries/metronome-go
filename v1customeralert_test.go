// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Metronome-Industries/metronome-go/v2"
	"github.com/Metronome-Industries/metronome-go/v2/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/v2/option"
)

func TestV1CustomerAlertGetWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Alerts.Get(context.TODO(), metronome.V1CustomerAlertGetParams{
		AlertID:    "8deed800-1b7a-495d-a207-6c52bac54dc9",
		CustomerID: "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		GroupValues: []metronome.V1CustomerAlertGetParamsGroupValue{{
			Key:   "key",
			Value: "value",
		}},
		PlansOrContracts: metronome.V1CustomerAlertGetParamsPlansOrContractsPlans,
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerAlertListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Alerts.List(context.TODO(), metronome.V1CustomerAlertListParams{
		CustomerID:    "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		NextPage:      metronome.String("next_page"),
		AlertStatuses: []string{"ENABLED"},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerAlertReset(t *testing.T) {
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
	err := client.V1.Customers.Alerts.Reset(context.TODO(), metronome.V1CustomerAlertResetParams{
		AlertID:    "5e8691bf-b22a-4672-922d-f80eee940f01",
		CustomerID: "4c83caf3-8af4-44e2-9aeb-e290531726d9",
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
