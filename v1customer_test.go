// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Metronome-Industries/metronome-go/v2"
	"github.com/Metronome-Industries/metronome-go/v2/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/shared"
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
		Name: "Example, Inc.",
		BillingConfig: metronome.V1CustomerNewParamsBillingConfig{
			BillingProviderCustomerID: "billing_provider_customer_id",
			BillingProviderType:       "aws_marketplace",
			AwsIsSubscriptionProduct:  metronome.Bool(true),
			AwsProductCode:            metronome.String("aws_product_code"),
			AwsRegion:                 "af-south-1",
			StripeCollectionMethod:    "charge_automatically",
		},
		CustomFields: map[string]string{
			"foo": "string",
		},
		CustomerBillingProviderConfigurations: []metronome.V1CustomerNewParamsCustomerBillingProviderConfiguration{{
			BillingProvider: "stripe",
			Configuration: map[string]any{
				"stripe_customer_id":       "bar",
				"stripe_collection_method": "bar",
			},
			DeliveryMethod:   "direct_to_billing_provider",
			DeliveryMethodID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			TaxProvider:      "anrok",
		}},
		ExternalID:    metronome.String("x"),
		IngestAliases: []string{"team@example.com"},
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
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
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
		CustomerIDs:          []string{"string"},
		IngestAlias:          metronome.String("ingest_alias"),
		Limit:                metronome.Int(1),
		NextPage:             metronome.String("next_page"),
		OnlyArchived:         metronome.Bool(true),
		SalesforceAccountIDs: []string{"string"},
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
			ID: "8deed800-1b7a-495d-a207-6c52bac54dc9",
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
		CustomerID:      "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		IncludeArchived: metronome.Bool(true),
		Limit:           metronome.Int(1),
		NextPage:        metronome.String("next_page"),
		OnCurrentPlan:   metronome.Bool(true),
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
		CustomerID:   "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		EndingBefore: time.Now(),
		StartingOn:   time.Now(),
		Limit:        metronome.Int(1),
		NextPage:     metronome.String("next_page"),
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
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		Events: []metronome.V1CustomerPreviewEventsParamsEvent{{
			EventType: "heartbeat",
			Properties: map[string]any{
				"cpu_hours":       "bar",
				"memory_gb_hours": "bar",
			},
			Timestamp:     metronome.String("2021-01-01T00:00:00Z"),
			TransactionID: metronome.String("x"),
		}},
		Mode:                 metronome.V1CustomerPreviewEventsParamsModeReplace,
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

func TestV1CustomerGetBillingConfigurationsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.GetBillingConfigurations(context.TODO(), metronome.V1CustomerGetBillingConfigurationsParams{
		CustomerID:      "6a37bb88-8538-48c5-b37b-a41c836328bd",
		IncludeArchived: metronome.Bool(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerSetBillingConfigurations(t *testing.T) {
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
	err := client.V1.Customers.SetBillingConfigurations(context.TODO(), metronome.V1CustomerSetBillingConfigurationsParams{
		Data: []metronome.V1CustomerSetBillingConfigurationsParamsData{{
			BillingProvider: "stripe",
			CustomerID:      "4db51251-61de-4bfe-b9ce-495e244f3491",
			Configuration: map[string]any{
				"stripe_customer_id":             "bar",
				"stripe_collection_method":       "bar",
				"leave_stripe_invoices_in_draft": "bar",
			},
			DeliveryMethod:   "direct_to_billing_provider",
			DeliveryMethodID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			TaxProvider:      "anrok",
		}, {
			BillingProvider: "aws_marketplace",
			CustomerID:      "4db51251-61de-4bfe-b9ce-495e244f3491",
			Configuration: map[string]any{
				"aws_customer_id":  "bar",
				"aws_product_code": "bar",
				"aws_region":       "bar",
			},
			DeliveryMethod:   "direct_to_billing_provider",
			DeliveryMethodID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			TaxProvider:      "anrok",
		}, {
			BillingProvider: "azure_marketplace",
			CustomerID:      "4db51251-61de-4bfe-b9ce-495e244f3491",
			Configuration: map[string]any{
				"azure_subscription_id": "bar",
			},
			DeliveryMethod:   "direct_to_billing_provider",
			DeliveryMethodID: metronome.String("5b9e3072-415b-4842-94f0-0b6700c8b6be"),
			TaxProvider:      "anrok",
		}, {
			BillingProvider: "aws_marketplace",
			CustomerID:      "4db51251-61de-4bfe-b9ce-495e244f3491",
			Configuration: map[string]any{
				"aws_customer_id":             "bar",
				"aws_product_code":            "bar",
				"aws_region":                  "bar",
				"aws_is_subscription_product": "bar",
			},
			DeliveryMethod:   "direct_to_billing_provider",
			DeliveryMethodID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			TaxProvider:      "anrok",
		}},
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
		CustomerID:    "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		IngestAliases: []string{"team@example.com"},
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
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		Name:       "Example, Inc.",
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
		CustomerID:                 "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		LeaveStripeInvoicesInDraft: metronome.Bool(true),
		SalesforceAccountID:        metronome.String("0015500001WO1ZiABL"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
