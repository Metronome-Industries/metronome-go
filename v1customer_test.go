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

func TestV1CustomerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.New(context.TODO(), metronome.V1CustomerNewParams{
		Name: metronome.F("Example, Inc."),
		BillingConfig: metronome.F(metronome.V1CustomerNewParamsBillingConfig{
			BillingProviderCustomerID: metronome.F("billing_provider_customer_id"),
			BillingProviderType:       metronome.F(metronome.V1CustomerNewParamsBillingConfigBillingProviderTypeAwsMarketplace),
			AwsIsSubscriptionProduct:  metronome.F(true),
			AwsProductCode:            metronome.F("aws_product_code"),
			AwsRegion:                 metronome.F(metronome.V1CustomerNewParamsBillingConfigAwsRegionAfSouth1),
			StripeCollectionMethod:    metronome.F(metronome.V1CustomerNewParamsBillingConfigStripeCollectionMethodChargeAutomatically),
		}),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		CustomerBillingProviderConfigurations: metronome.F([]metronome.V1CustomerNewParamsCustomerBillingProviderConfiguration{{
			BillingProvider: metronome.F(metronome.V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderStripe),
			Configuration: metronome.F(map[string]interface{}{
				"stripe_customer_id":       "bar",
				"stripe_collection_method": "bar",
			}),
			DeliveryMethod:   metronome.F(metronome.V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodDirectToBillingProvider),
			DeliveryMethodID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}}),
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

func TestV1CustomerGet(t *testing.T) {
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
	_, err := client.V1.Customers.Get(context.TODO(), metronome.V1CustomerGetParams{
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

func TestV1CustomerListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.List(context.TODO(), metronome.V1CustomerListParams{
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

func TestV1CustomerArchive(t *testing.T) {
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
	_, err := client.V1.Customers.Archive(context.TODO(), metronome.V1CustomerArchiveParams{
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

func TestV1CustomerListBillableMetricsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.ListBillableMetrics(context.TODO(), metronome.V1CustomerListBillableMetricsParams{
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

func TestV1CustomerListCostsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.ListCosts(context.TODO(), metronome.V1CustomerListCostsParams{
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

func TestV1CustomerPreviewEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.PreviewEvents(context.TODO(), metronome.V1CustomerPreviewEventsParams{
		CustomerID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		Events: metronome.F([]metronome.V1CustomerPreviewEventsParamsEvent{{
			EventType:  metronome.F("heartbeat"),
			CustomerID: metronome.F("x"),
			Properties: metronome.F(map[string]interface{}{
				"cpu_hours":       "bar",
				"memory_gb_hours": "bar",
			}),
			Timestamp:     metronome.F("2021-01-01T00:00:00Z"),
			TransactionID: metronome.F("x"),
		}}),
		Mode:                 metronome.F(metronome.V1CustomerPreviewEventsParamsModeReplace),
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

func TestV1CustomerSetIngestAliases(t *testing.T) {
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
	err := client.V1.Customers.SetIngestAliases(context.TODO(), metronome.V1CustomerSetIngestAliasesParams{
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

func TestV1CustomerSetName(t *testing.T) {
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
	_, err := client.V1.Customers.SetName(context.TODO(), metronome.V1CustomerSetNameParams{
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

func TestV1CustomerUpdateConfigWithOptionalParams(t *testing.T) {
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
	err := client.V1.Customers.UpdateConfig(context.TODO(), metronome.V1CustomerUpdateConfigParams{
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
