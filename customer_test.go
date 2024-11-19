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
	"github.com/Metronome-Industries/metronome-go/shared"
)

func TestCustomerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.New(context.TODO(), metronome.CustomerNewParams{
		Name: metronome.F("Example, Inc."),
		BillingConfig: metronome.F(metronome.CustomerNewParamsBillingConfig{
			BillingProviderCustomerID: metronome.F("billing_provider_customer_id"),
			BillingProviderType:       metronome.F(metronome.CustomerNewParamsBillingConfigBillingProviderTypeAwsMarketplace),
			AwsProductCode:            metronome.F("aws_product_code"),
			AwsRegion:                 metronome.F(metronome.CustomerNewParamsBillingConfigAwsRegionAfSouth1),
			StripeCollectionMethod:    metronome.F(metronome.CustomerNewParamsBillingConfigStripeCollectionMethodChargeAutomatically),
		}),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		ExternalID:    metronome.F("x"),
		IngestAliases: metronome.F([]string{"team@example.com"}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerGet(t *testing.T) {
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
	_, err := client.Customers.Get(context.TODO(), metronome.CustomerGetParams{
		CustomerID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.List(context.TODO(), metronome.CustomerListParams{
		CustomerIDs:          metronome.F([]string{"string"}),
		IngestAlias:          metronome.F("ingest_alias"),
		Limit:                metronome.F(int64(1)),
		NextPage:             metronome.F("next_page"),
		OnlyArchived:         metronome.F(true),
		SalesforceAccountIDs: metronome.F([]string{"string"}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerArchive(t *testing.T) {
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
	_, err := client.Customers.Archive(context.TODO(), metronome.CustomerArchiveParams{
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

func TestCustomerListBillableMetricsWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.ListBillableMetrics(context.TODO(), metronome.CustomerListBillableMetricsParams{
		CustomerID:      metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		IncludeArchived: metronome.F(true),
		Limit:           metronome.F(int64(1)),
		NextPage:        metronome.F("next_page"),
		OnCurrentPlan:   metronome.F(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerListCostsWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.ListCosts(context.TODO(), metronome.CustomerListCostsParams{
		CustomerID:   metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		EndingBefore: metronome.F(time.Now()),
		StartingOn:   metronome.F(time.Now()),
		Limit:        metronome.F(int64(1)),
		NextPage:     metronome.F("next_page"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerSetIngestAliases(t *testing.T) {
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
	err := client.Customers.SetIngestAliases(context.TODO(), metronome.CustomerSetIngestAliasesParams{
		CustomerID:    metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		IngestAliases: metronome.F([]string{"team@example.com"}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerSetName(t *testing.T) {
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
	_, err := client.Customers.SetName(context.TODO(), metronome.CustomerSetNameParams{
		CustomerID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		Name:       metronome.F("Example, Inc."),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerUpdateConfigWithOptionalParams(t *testing.T) {
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
	err := client.Customers.UpdateConfig(context.TODO(), metronome.CustomerUpdateConfigParams{
		CustomerID:                 metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		LeaveStripeInvoicesInDraft: metronome.F(true),
		SalesforceAccountID:        metronome.F("0015500001WO1ZiABL"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
