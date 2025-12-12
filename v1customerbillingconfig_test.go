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
		CustomerID:                "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		BillingProviderType:       metronome.V1CustomerBillingConfigNewParamsBillingProviderTypeStripe,
		BillingProviderCustomerID: "cus_AJ6y20bjkOOayM",
		AwsCustomerAccountID:      metronome.String("aws_customer_account_id"),
		AwsCustomerID:             metronome.String("aws_customer_id"),
		AwsProductCode:            metronome.String("aws_product_code"),
		AwsRegion:                 metronome.V1CustomerBillingConfigNewParamsAwsRegionAfSouth1,
		StripeCollectionMethod:    metronome.V1CustomerBillingConfigNewParamsStripeCollectionMethodChargeAutomatically,
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
		CustomerID:          "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		BillingProviderType: metronome.V1CustomerBillingConfigGetParamsBillingProviderTypeStripe,
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
		CustomerID:          "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		BillingProviderType: metronome.V1CustomerBillingConfigDeleteParamsBillingProviderTypeStripe,
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
