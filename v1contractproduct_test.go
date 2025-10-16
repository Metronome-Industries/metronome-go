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

func TestV1ContractProductNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.Products.New(context.TODO(), metronome.V1ContractProductNewParams{
		Name:                "My Product",
		Type:                metronome.V1ContractProductNewParamsTypeUsage,
		BillableMetricID:    metronome.String("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		CompositeProductIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		CompositeTags:       []string{"string"},
		CustomFields: map[string]string{
			"foo": "string",
		},
		ExcludeFreeUsage:       metronome.Bool(true),
		IsRefundable:           metronome.Bool(true),
		NetsuiteInternalItemID: metronome.String("netsuite_internal_item_id"),
		NetsuiteOverageItemID:  metronome.String("netsuite_overage_item_id"),
		PresentationGroupKey:   []string{"string"},
		PricingGroupKey:        []string{"string"},
		QuantityConversion: metronome.QuantityConversionParam{
			ConversionFactor: 0,
			Operation:        metronome.QuantityConversionOperationMultiply,
			Name:             metronome.String("name"),
		},
		QuantityRounding: metronome.QuantityRoundingParam{
			DecimalPlaces:  0,
			RoundingMethod: metronome.QuantityRoundingRoundingMethodRoundUp,
		},
		Tags: []string{"string"},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContractProductGet(t *testing.T) {
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
	_, err := client.V1.Contracts.Products.Get(context.TODO(), metronome.V1ContractProductGetParams{
		ID: shared.IDParam{
			ID: "d84e7f4e-7a70-4fe4-be02-7a5027beffcc",
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

func TestV1ContractProductUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.Products.Update(context.TODO(), metronome.V1ContractProductUpdateParams{
		ProductID:              "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		StartingAt:             time.Now(),
		BillableMetricID:       metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CompositeProductIDs:    []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		CompositeTags:          []string{"string"},
		ExcludeFreeUsage:       metronome.Bool(true),
		IsRefundable:           metronome.Bool(true),
		Name:                   metronome.String("My Updated Product"),
		NetsuiteInternalItemID: metronome.String("netsuite_internal_item_id"),
		NetsuiteOverageItemID:  metronome.String("netsuite_overage_item_id"),
		PresentationGroupKey:   []string{"string"},
		PricingGroupKey:        []string{"string"},
		QuantityConversion: metronome.QuantityConversionParam{
			ConversionFactor: 0,
			Operation:        metronome.QuantityConversionOperationMultiply,
			Name:             metronome.String("name"),
		},
		QuantityRounding: metronome.QuantityRoundingParam{
			DecimalPlaces:  0,
			RoundingMethod: metronome.QuantityRoundingRoundingMethodRoundUp,
		},
		Tags: []string{"string"},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContractProductListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.Products.List(context.TODO(), metronome.V1ContractProductListParams{
		Limit:         metronome.Int(1),
		NextPage:      metronome.String("next_page"),
		ArchiveFilter: metronome.V1ContractProductListParamsArchiveFilterNotArchived,
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContractProductArchive(t *testing.T) {
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
	_, err := client.V1.Contracts.Products.Archive(context.TODO(), metronome.V1ContractProductArchiveParams{
		ProductID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
