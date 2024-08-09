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

func TestContractRateCardRateListWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.RateCards.Rates.List(context.TODO(), metronome.ContractRateCardRateListParams{
		At:         metronome.F(time.Now()),
		RateCardID: metronome.F("f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe"),
		Limit:      metronome.F(int64(1)),
		NextPage:   metronome.F("next_page"),
		Selectors: metronome.F([]metronome.ContractRateCardRateListParamsSelector{{
			ProductID:   metronome.F("d6300dbb-882e-4d2d-8dec-5125d16b65d0"),
			ProductTags: metronome.F([]string{"string", "string", "string"}),
			PricingGroupValues: metronome.F(map[string]string{
				"foo": "string",
			}),
			PartialPricingGroupValues: metronome.F(map[string]string{
				"region": "us-west-2",
				"cloud":  "aws",
			}),
		}}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractRateCardRateAddWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.RateCards.Rates.Add(context.TODO(), metronome.ContractRateCardRateAddParams{
		Entitled:     metronome.F(true),
		ProductID:    metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		RateCardID:   metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		RateType:     metronome.F(metronome.ContractRateCardRateAddParamsRateTypeFlat),
		StartingAt:   metronome.F(time.Now()),
		CreditTypeID: metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
		CustomRate: metronome.F(map[string]interface{}{
			"foo": "bar",
		}),
		EndingBefore: metronome.F(time.Now()),
		IsProrated:   metronome.F(true),
		Price:        metronome.F(100.000000),
		PricingGroupValues: metronome.F(map[string]string{
			"foo": "string",
		}),
		Quantity: metronome.F(0.000000),
		Tiers: metronome.F([]metronome.ContractRateCardRateAddParamsTier{{
			Size:  metronome.F(0.000000),
			Price: metronome.F(0.000000),
		}, {
			Size:  metronome.F(0.000000),
			Price: metronome.F(0.000000),
		}, {
			Size:  metronome.F(0.000000),
			Price: metronome.F(0.000000),
		}}),
		UseListPrices: metronome.F(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractRateCardRateAddManyWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.RateCards.Rates.AddMany(context.TODO(), metronome.ContractRateCardRateAddManyParams{
		RateCardID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		Rates: metronome.F([]metronome.ContractRateCardRateAddManyParamsRate{{
			ProductID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
			PricingGroupValues: metronome.F(map[string]string{
				"region": "us-west-2",
				"cloud":  "aws",
			}),
			StartingAt:    metronome.F(time.Now()),
			EndingBefore:  metronome.F(time.Now()),
			Entitled:      metronome.F(true),
			RateType:      metronome.F(metronome.ContractRateCardRateAddManyParamsRatesRateTypeFlat),
			Price:         metronome.F(100.000000),
			CreditTypeID:  metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			Quantity:      metronome.F(0.000000),
			IsProrated:    metronome.F(true),
			UseListPrices: metronome.F(true),
			Tiers: metronome.F([]metronome.ContractRateCardRateAddManyParamsRatesTier{{
				Size:  metronome.F(0.000000),
				Price: metronome.F(0.000000),
			}, {
				Size:  metronome.F(0.000000),
				Price: metronome.F(0.000000),
			}, {
				Size:  metronome.F(0.000000),
				Price: metronome.F(0.000000),
			}}),
			CustomRate: metronome.F(map[string]interface{}{
				"foo": "bar",
			}),
		}, {
			ProductID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
			PricingGroupValues: metronome.F(map[string]string{
				"region": "us-east-2",
				"cloud":  "aws",
			}),
			StartingAt:    metronome.F(time.Now()),
			EndingBefore:  metronome.F(time.Now()),
			Entitled:      metronome.F(true),
			RateType:      metronome.F(metronome.ContractRateCardRateAddManyParamsRatesRateTypeFlat),
			Price:         metronome.F(120.000000),
			CreditTypeID:  metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			Quantity:      metronome.F(0.000000),
			IsProrated:    metronome.F(true),
			UseListPrices: metronome.F(true),
			Tiers: metronome.F([]metronome.ContractRateCardRateAddManyParamsRatesTier{{
				Size:  metronome.F(0.000000),
				Price: metronome.F(0.000000),
			}, {
				Size:  metronome.F(0.000000),
				Price: metronome.F(0.000000),
			}, {
				Size:  metronome.F(0.000000),
				Price: metronome.F(0.000000),
			}}),
			CustomRate: metronome.F(map[string]interface{}{
				"foo": "bar",
			}),
		}}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}