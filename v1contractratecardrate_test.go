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

func TestV1ContractRateCardRateListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.RateCards.Rates.List(context.TODO(), metronome.V1ContractRateCardRateListParams{
		At:         metronome.F(time.Now()),
		RateCardID: metronome.F("f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe"),
		Limit:      metronome.F(int64(1)),
		NextPage:   metronome.F("next_page"),
		Selectors: metronome.F([]metronome.V1ContractRateCardRateListParamsSelector{{
			PartialPricingGroupValues: metronome.F(map[string]string{
				"region": "us-west-2",
				"cloud":  "aws",
			}),
			PricingGroupValues: metronome.F(map[string]string{
				"foo": "string",
			}),
			ProductID:   metronome.F("d6300dbb-882e-4d2d-8dec-5125d16b65d0"),
			ProductTags: metronome.F([]string{"string"}),
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

func TestV1ContractRateCardRateAddWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.RateCards.Rates.Add(context.TODO(), metronome.V1ContractRateCardRateAddParams{
		Entitled:   metronome.F(true),
		ProductID:  metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		RateCardID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		RateType:   metronome.F(metronome.V1ContractRateCardRateAddParamsRateTypeFlat),
		StartingAt: metronome.F(time.Now()),
		CommitRate: metronome.F(metronome.V1ContractRateCardRateAddParamsCommitRate{
			RateType: metronome.F(metronome.V1ContractRateCardRateAddParamsCommitRateRateTypeFlat),
			Price:    metronome.F(0.000000),
			Tiers: metronome.F([]shared.TierParam{{
				Price: metronome.F(0.000000),
				Size:  metronome.F(0.000000),
			}}),
		}),
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
		Tiers: metronome.F([]shared.TierParam{{
			Price: metronome.F(0.000000),
			Size:  metronome.F(0.000000),
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

func TestV1ContractRateCardRateAddMany(t *testing.T) {
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
	_, err := client.V1.Contracts.RateCards.Rates.AddMany(context.TODO(), metronome.V1ContractRateCardRateAddManyParams{
		RateCardID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		Rates: metronome.F([]metronome.V1ContractRateCardRateAddManyParamsRate{{
			Entitled:   metronome.F(true),
			ProductID:  metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
			RateType:   metronome.F(metronome.V1ContractRateCardRateAddManyParamsRatesRateTypeFlat),
			StartingAt: metronome.F(time.Now()),
			CommitRate: metronome.F(metronome.V1ContractRateCardRateAddManyParamsRatesCommitRate{
				RateType: metronome.F(metronome.V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypeFlat),
				Price:    metronome.F(0.000000),
				Tiers: metronome.F([]shared.TierParam{{
					Price: metronome.F(0.000000),
					Size:  metronome.F(0.000000),
				}}),
			}),
			CreditTypeID: metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			CustomRate: metronome.F(map[string]interface{}{
				"foo": "bar",
			}),
			EndingBefore: metronome.F(time.Now()),
			IsProrated:   metronome.F(true),
			Price:        metronome.F(100.000000),
			PricingGroupValues: metronome.F(map[string]string{
				"region": "us-west-2",
				"cloud":  "aws",
			}),
			Quantity: metronome.F(0.000000),
			Tiers: metronome.F([]shared.TierParam{{
				Price: metronome.F(0.000000),
				Size:  metronome.F(0.000000),
			}}),
			UseListPrices: metronome.F(true),
		}, {
			Entitled:   metronome.F(true),
			ProductID:  metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
			RateType:   metronome.F(metronome.V1ContractRateCardRateAddManyParamsRatesRateTypeFlat),
			StartingAt: metronome.F(time.Now()),
			CommitRate: metronome.F(metronome.V1ContractRateCardRateAddManyParamsRatesCommitRate{
				RateType: metronome.F(metronome.V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypeFlat),
				Price:    metronome.F(0.000000),
				Tiers: metronome.F([]shared.TierParam{{
					Price: metronome.F(0.000000),
					Size:  metronome.F(0.000000),
				}}),
			}),
			CreditTypeID: metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			CustomRate: metronome.F(map[string]interface{}{
				"foo": "bar",
			}),
			EndingBefore: metronome.F(time.Now()),
			IsProrated:   metronome.F(true),
			Price:        metronome.F(120.000000),
			PricingGroupValues: metronome.F(map[string]string{
				"region": "us-east-2",
				"cloud":  "aws",
			}),
			Quantity: metronome.F(0.000000),
			Tiers: metronome.F([]shared.TierParam{{
				Price: metronome.F(0.000000),
				Size:  metronome.F(0.000000),
			}}),
			UseListPrices: metronome.F(true),
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
