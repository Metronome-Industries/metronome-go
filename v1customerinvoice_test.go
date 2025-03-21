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

func TestV1CustomerInvoiceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Invoices.Get(context.TODO(), metronome.V1CustomerInvoiceGetParams{
		CustomerID:           metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		InvoiceID:            metronome.F("6a37bb88-8538-48c5-b37b-a41c836328bd"),
		SkipZeroQtyLineItems: metronome.F(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerInvoiceListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Invoices.List(context.TODO(), metronome.V1CustomerInvoiceListParams{
		CustomerID:           metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CreditTypeID:         metronome.F("credit_type_id"),
		EndingBefore:         metronome.F(time.Now()),
		Limit:                metronome.F(int64(1)),
		NextPage:             metronome.F("next_page"),
		SkipZeroQtyLineItems: metronome.F(true),
		Sort:                 metronome.F(metronome.V1CustomerInvoiceListParamsSortDateAsc),
		StartingOn:           metronome.F(time.Now()),
		Status:               metronome.F("status"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerInvoiceAddCharge(t *testing.T) {
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
	_, err := client.V1.Customers.Invoices.AddCharge(context.TODO(), metronome.V1CustomerInvoiceAddChargeParams{
		CustomerID:            metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		ChargeID:              metronome.F("5ae4b726-1ebe-439c-9190-9831760ba195"),
		CustomerPlanID:        metronome.F("a23b3cf4-47fb-4c3f-bb3d-9e64f7704015"),
		Description:           metronome.F("One time charge"),
		InvoiceStartTimestamp: metronome.F(time.Now()),
		Price:                 metronome.F(250.000000),
		Quantity:              metronome.F(1.000000),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerInvoiceListBreakdownsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Invoices.ListBreakdowns(context.TODO(), metronome.V1CustomerInvoiceListBreakdownsParams{
		CustomerID:           metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		EndingBefore:         metronome.F(time.Now()),
		StartingOn:           metronome.F(time.Now()),
		CreditTypeID:         metronome.F("credit_type_id"),
		Limit:                metronome.F(int64(1)),
		NextPage:             metronome.F("next_page"),
		SkipZeroQtyLineItems: metronome.F(true),
		Sort:                 metronome.F(metronome.V1CustomerInvoiceListBreakdownsParamsSortDateAsc),
		Status:               metronome.F("status"),
		WindowSize:           metronome.F(metronome.V1CustomerInvoiceListBreakdownsParamsWindowSizeHour),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
