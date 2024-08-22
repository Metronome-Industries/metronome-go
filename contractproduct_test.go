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

func TestContractProductNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.Products.New(context.TODO(), metronome.ContractProductNewParams{
		Name:                   metronome.F("My Product"),
		Type:                   metronome.F(metronome.ContractProductNewParamsTypeFixed),
		BillableMetricID:       metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		CompositeProductIDs:    metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		CompositeTags:          metronome.F([]string{"string", "string", "string"}),
		ExcludeFreeUsage:       metronome.F(true),
		IsRefundable:           metronome.F(true),
		NetsuiteInternalItemID: metronome.F("netsuite_internal_item_id"),
		NetsuiteOverageItemID:  metronome.F("netsuite_overage_item_id"),
		PresentationGroupKey:   metronome.F([]string{"string", "string", "string"}),
		PricingGroupKey:        metronome.F([]string{"string", "string", "string"}),
		QuantityConversion: metronome.F(metronome.QuantityConversionParam{
			ConversionFactor: metronome.F(0.000000),
			Operation:        metronome.F(metronome.QuantityConversionOperationMultiply),
			Name:             metronome.F("name"),
		}),
		QuantityRounding: metronome.F(metronome.QuantityRoundingParam{
			DecimalPlaces:  metronome.F(0.000000),
			RoundingMethod: metronome.F(metronome.QuantityRoundingRoundingMethodRoundUp),
		}),
		Tags: metronome.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractProductGet(t *testing.T) {
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
	_, err := client.Contracts.Products.Get(context.TODO(), metronome.ContractProductGetParams{
		ID: shared.IDParam{
			ID: metronome.F("d84e7f4e-7a70-4fe4-be02-7a5027beffcc"),
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

func TestContractProductUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.Products.Update(context.TODO(), metronome.ContractProductUpdateParams{
		ProductID:              metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		StartingAt:             metronome.F(time.Now()),
		BillableMetricID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CompositeProductIDs:    metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		CompositeTags:          metronome.F([]string{"string", "string", "string"}),
		ExcludeFreeUsage:       metronome.F(true),
		IsRefundable:           metronome.F(true),
		Name:                   metronome.F("My Updated Product"),
		NetsuiteInternalItemID: metronome.F("netsuite_internal_item_id"),
		NetsuiteOverageItemID:  metronome.F("netsuite_overage_item_id"),
		PresentationGroupKey:   metronome.F([]string{"string", "string", "string"}),
		PricingGroupKey:        metronome.F([]string{"string", "string", "string"}),
		QuantityConversion: metronome.F(metronome.QuantityConversionParam{
			ConversionFactor: metronome.F(0.000000),
			Operation:        metronome.F(metronome.QuantityConversionOperationMultiply),
			Name:             metronome.F("name"),
		}),
		QuantityRounding: metronome.F(metronome.QuantityRoundingParam{
			DecimalPlaces:  metronome.F(0.000000),
			RoundingMethod: metronome.F(metronome.QuantityRoundingRoundingMethodRoundUp),
		}),
		Tags: metronome.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractProductListWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.Products.List(context.TODO(), metronome.ContractProductListParams{
		Limit:         metronome.F(int64(1)),
		NextPage:      metronome.F("next_page"),
		ArchiveFilter: metronome.F(metronome.ContractProductListParamsArchiveFilterArchived),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractProductArchive(t *testing.T) {
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
	_, err := client.Contracts.Products.Archive(context.TODO(), metronome.ContractProductArchiveParams{
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
