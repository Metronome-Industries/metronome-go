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
)

func TestV1AlertNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Alerts.New(context.TODO(), metronome.V1AlertNewParams{
		AlertType:              metronome.F(metronome.V1AlertNewParamsAlertTypeSpendThresholdReached),
		Name:                   metronome.F("$100 spend threshold reached"),
		Threshold:              metronome.F(10000.000000),
		BillableMetricID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CreditGrantTypeFilters: metronome.F([]string{"enterprise"}),
		CreditTypeID:           metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
		CustomFieldFilters: metronome.F([]metronome.V1AlertNewParamsCustomFieldFilter{{
			Entity: metronome.F(metronome.V1AlertNewParamsCustomFieldFiltersEntityContract),
			Key:    metronome.F("key"),
			Value:  metronome.F("value"),
		}}),
		CustomerID:       metronome.F("4db51251-61de-4bfe-b9ce-495e244f3491"),
		EvaluateOnCreate: metronome.F(true),
		GroupKeyFilter: metronome.F(metronome.V1AlertNewParamsGroupKeyFilter{
			Key:   metronome.F("key"),
			Value: metronome.F("value"),
		}),
		GroupValues: metronome.F([]metronome.V1AlertNewParamsGroupValue{{
			Key:   metronome.F("key"),
			Value: metronome.F("value"),
		}}),
		InvoiceTypesFilter: metronome.F([]string{"PLAN_ARREARS, SCHEDULED, USAGE, CORRECTION, CREDIT_PURCHASE, or SEAT_PURCHASE"}),
		PlanID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		UniquenessKey:      metronome.F("x"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1AlertArchiveWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Alerts.Archive(context.TODO(), metronome.V1AlertArchiveParams{
		ID:                   metronome.F("8deed800-1b7a-495d-a207-6c52bac54dc9"),
		ReleaseUniquenessKey: metronome.F(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
