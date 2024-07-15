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

func TestCustomerInvoiceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Invoices.Get(
		context.TODO(),
		"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		"6a37bb88-8538-48c5-b37b-a41c836328bd",
		metronome.CustomerInvoiceGetParams{
			SkipZeroQtyLineItems: metronome.F(true),
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

func TestCustomerInvoiceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Invoices.List(
		context.TODO(),
		"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		metronome.CustomerInvoiceListParams{
			CreditTypeID:         metronome.F("credit_type_id"),
			EndingBefore:         metronome.F(time.Now()),
			Limit:                metronome.F(int64(1)),
			NextPage:             metronome.F("next_page"),
			SkipZeroQtyLineItems: metronome.F(true),
			Sort:                 metronome.F(metronome.CustomerInvoiceListParamsSortDateAsc),
			StartingOn:           metronome.F(time.Now()),
			Status:               metronome.F("status"),
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

func TestCustomerInvoiceAddCharge(t *testing.T) {
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
	_, err := client.Customers.Invoices.AddCharge(
		context.TODO(),
		"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		metronome.CustomerInvoiceAddChargeParams{
			ChargeID:              metronome.F("5ae4b726-1ebe-439c-9190-9831760ba195"),
			CustomerPlanID:        metronome.F("a23b3cf4-47fb-4c3f-bb3d-9e64f7704015"),
			Description:           metronome.F("One time charge"),
			InvoiceStartTimestamp: metronome.F(time.Now()),
			Price:                 metronome.F(250.000000),
			Quantity:              metronome.F(1.000000),
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
