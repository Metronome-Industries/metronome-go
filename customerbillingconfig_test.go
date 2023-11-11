// File generated from our OpenAPI spec by Stainless.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/metronome/metronome-go"
	"github.com/metronome/metronome-go/internal/testutil"
	"github.com/metronome/metronome-go/option"
)

func TestCustomerBillingConfigNewWithOptionalParams(t *testing.T) {
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
	err := client.Customers.BillingConfig.New(
		context.TODO(),
		"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		metronome.CustomerBillingConfigNewParamsBillingProviderTypeStripe,
		metronome.CustomerBillingConfigNewParams{
			BillingProviderCustomerID: metronome.F("cus_AJ6y20bjkOOayM"),
			AwsProductCode:            metronome.F("string"),
			AwsRegion:                 metronome.F(metronome.CustomerBillingConfigNewParamsAwsRegionAfSouth1),
			StripeCollectionMethod:    metronome.F(metronome.CustomerBillingConfigNewParamsStripeCollectionMethodChargeAutomatically),
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

func TestCustomerBillingConfigGet(t *testing.T) {
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
	_, err := client.Customers.BillingConfig.Get(
		context.TODO(),
		"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		metronome.CustomerBillingConfigGetParamsBillingProviderTypeStripe,
	)
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerBillingConfigDelete(t *testing.T) {
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
	err := client.Customers.BillingConfig.Delete(
		context.TODO(),
		"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		metronome.CustomerBillingConfigDeleteParamsBillingProviderTypeStripe,
	)
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
