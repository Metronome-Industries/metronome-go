// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/v3/option"
)

func TestV1PaymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Payments.List(context.TODO(), metronome.V1PaymentListParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		InvoiceID:  "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
		Limit:      metronome.Int(1),
		NextPage:   metronome.String("next_page"),
		Statuses:   []metronome.PaymentStatus{metronome.PaymentStatusPending, metronome.PaymentStatusRequiresIntervention},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PaymentAttempt(t *testing.T) {
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
	_, err := client.V1.Payments.Attempt(context.TODO(), metronome.V1PaymentAttemptParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		InvoiceID:  "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PaymentCancel(t *testing.T) {
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
	_, err := client.V1.Payments.Cancel(context.TODO(), metronome.V1PaymentCancelParams{
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		InvoiceID:  "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
