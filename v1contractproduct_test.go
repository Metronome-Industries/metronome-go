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
		Name:                   metronome.F("My Product"),
		Type:                   metronome.F(metronome.V1ContractProductNewParamsTypeFixed),
		BillableMetricID:       metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		CompositeProductIDs:    metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		CompositeTags:          metronome.F([]string{"string"}),
		ExcludeFreeUsage:       metronome.F(true),
		IsRefundable:           metronome.F(true),
		NetsuiteInternalItemID: metronome.F("netsuite_internal_item_id"),
		NetsuiteOverageItemID:  metronome.F("netsuite_overage_item_id"),
		PresentationGroupKey:   metronome.F([]string{"string"}),
		PricingGroupKey:        metronome.F([]string{"string"}),
		QuantityConversion: metronome.F(metronome.V1ContractProductNewParamsQuantityConversion{
			ConversionFactor: metronome.F(0.000000),
			Operation:        metronome.F(metronome.V1ContractProductNewParamsQuantityConversionOperationMultiply),
			Name:             metronome.F("name"),
		}),
		QuantityRounding: metronome.F(metronome.V1ContractProductNewParamsQuantityRounding{
			DecimalPlaces:  metronome.F(0.000000),
			RoundingMethod: metronome.F(metronome.V1ContractProductNewParamsQuantityRoundingRoundingMethodRoundUp),
		}),
		Tags: metronome.F([]string{"string"}),
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
		ID: metronome.F("d84e7f4e-7a70-4fe4-be02-7a5027beffcc"),
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
		ProductID:              metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		StartingAt:             metronome.F(time.Now()),
		BillableMetricID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CompositeProductIDs:    metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		CompositeTags:          metronome.F([]string{"string"}),
		ExcludeFreeUsage:       metronome.F(true),
		IsRefundable:           metronome.F(true),
		Name:                   metronome.F("My Updated Product"),
		NetsuiteInternalItemID: metronome.F("netsuite_internal_item_id"),
		NetsuiteOverageItemID:  metronome.F("netsuite_overage_item_id"),
		PresentationGroupKey:   metronome.F([]string{"string"}),
		PricingGroupKey:        metronome.F([]string{"string"}),
		QuantityConversion: metronome.F(metronome.V1ContractProductUpdateParamsQuantityConversion{
			ConversionFactor: metronome.F(0.000000),
			Operation:        metronome.F(metronome.V1ContractProductUpdateParamsQuantityConversionOperationMultiply),
			Name:             metronome.F("name"),
		}),
		QuantityRounding: metronome.F(metronome.V1ContractProductUpdateParamsQuantityRounding{
			DecimalPlaces:  metronome.F(0.000000),
			RoundingMethod: metronome.F(metronome.V1ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundUp),
		}),
		Tags: metronome.F([]string{"string"}),
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
		Limit:         metronome.F(int64(1)),
		NextPage:      metronome.F("next_page"),
		ArchiveFilter: metronome.F(metronome.V1ContractProductListParamsArchiveFilterArchived),
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
		ProductID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
