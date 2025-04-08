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

func TestV1CustomerBillingConfigNewWithOptionalParams(t *testing.T) {
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
	err := client.V1.Customers.BillingConfig.New(context.TODO(), metronome.V1CustomerBillingConfigNewParams{
		CustomerID:                metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		BillingProviderType:       metronome.F(metronome.V1CustomerBillingConfigNewParamsBillingProviderTypeStripe),
		BillingProviderCustomerID: metronome.F("cus_AJ6y20bjkOOayM"),
		AwsProductCode:            metronome.F("aws_product_code"),
		AwsRegion:                 metronome.F(metronome.V1CustomerBillingConfigNewParamsAwsRegionAfSouth1),
		StripeCollectionMethod:    metronome.F(metronome.V1CustomerBillingConfigNewParamsStripeCollectionMethodChargeAutomatically),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerBillingConfigGet(t *testing.T) {
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
	_, err := client.V1.Customers.BillingConfig.Get(context.TODO(), metronome.V1CustomerBillingConfigGetParams{
		CustomerID:          metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		BillingProviderType: metronome.F(metronome.V1CustomerBillingConfigGetParamsBillingProviderTypeStripe),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerBillingConfigDelete(t *testing.T) {
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
	err := client.V1.Customers.BillingConfig.Delete(context.TODO(), metronome.V1CustomerBillingConfigDeleteParams{
		CustomerID:          metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		BillingProviderType: metronome.F(metronome.V1CustomerBillingConfigDeleteParamsBillingProviderTypeStripe),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
