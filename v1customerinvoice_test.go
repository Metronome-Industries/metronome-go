// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
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
		CustomerID:           "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		InvoiceID:            "6a37bb88-8538-48c5-b37b-a41c836328bd",
		SkipZeroQtyLineItems: metronome.Bool(true),
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
		CustomerID:           "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CreditTypeID:         metronome.String("credit_type_id"),
		EndingBefore:         metronome.Time(time.Now()),
		Limit:                metronome.Int(1),
		NextPage:             metronome.String("next_page"),
		SkipZeroQtyLineItems: metronome.Bool(true),
		Sort:                 metronome.V1CustomerInvoiceListParamsSortDateAsc,
		StartingOn:           metronome.Time(time.Now()),
		Status:               metronome.String("status"),
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
		CustomerID:            "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		ChargeID:              "5ae4b726-1ebe-439c-9190-9831760ba195",
		CustomerPlanID:        "a23b3cf4-47fb-4c3f-bb3d-9e64f7704015",
		Description:           "One time charge",
		InvoiceStartTimestamp: time.Now(),
		Price:                 250,
		Quantity:              1,
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
		CustomerID:           "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		EndingBefore:         time.Now(),
		StartingOn:           time.Now(),
		CreditTypeID:         metronome.String("credit_type_id"),
		Limit:                metronome.Int(1),
		NextPage:             metronome.String("next_page"),
		SkipZeroQtyLineItems: metronome.Bool(true),
		Sort:                 metronome.V1CustomerInvoiceListBreakdownsParamsSortDateAsc,
		Status:               metronome.String("status"),
		WindowSize:           metronome.V1CustomerInvoiceListBreakdownsParamsWindowSizeHour,
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerInvoiceGetPdf(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	resp, err := client.V1.Customers.Invoices.GetPdf(context.TODO(), metronome.V1CustomerInvoiceGetPdfParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		InvoiceID:  "6a37bb88-8538-48c5-b37b-a41c836328bd",
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
