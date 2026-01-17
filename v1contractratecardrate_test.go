// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"github.com/Metronome-Industries/metronome-go/v3/shared"
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
		At:         time.Now(),
		RateCardID: "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe",
		Limit:      metronome.Int(1),
		NextPage:   metronome.String("next_page"),
		Selectors: []metronome.V1ContractRateCardRateListParamsSelector{{
			BillingFrequency: "MONTHLY",
			PartialPricingGroupValues: map[string]string{
				"region": "us-west-2",
				"cloud":  "aws",
			},
			PricingGroupValues: map[string]string{
				"foo": "string",
			},
			ProductID:   metronome.String("d6300dbb-882e-4d2d-8dec-5125d16b65d0"),
			ProductTags: []string{"string"},
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
		Entitled:         true,
		ProductID:        "13117714-3f05-48e5-a6e9-a66093f13b4d",
		RateCardID:       "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		RateType:         metronome.V1ContractRateCardRateAddParamsRateTypeFlat,
		StartingAt:       time.Now(),
		BillingFrequency: metronome.V1ContractRateCardRateAddParamsBillingFrequencyMonthly,
		CommitRate: shared.CommitRateParam{
			RateType: shared.CommitRateRateTypeFlat,
			Price:    metronome.Float(0),
			Tiers: []shared.TierParam{{
				Price: 0,
				Size:  metronome.Float(0),
			}},
		},
		CreditTypeID: metronome.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
		CustomRate: map[string]any{
			"foo": "bar",
		},
		EndingBefore: metronome.Time(time.Now()),
		IsProrated:   metronome.Bool(true),
		Price:        metronome.Float(100),
		PricingGroupValues: map[string]string{
			"foo": "string",
		},
		Quantity: metronome.Float(0),
		Tiers: []shared.TierParam{{
			Price: 0,
			Size:  metronome.Float(0),
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
		RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		Rates: []metronome.V1ContractRateCardRateAddManyParamsRate{{
			Entitled:         true,
			ProductID:        "13117714-3f05-48e5-a6e9-a66093f13b4d",
			RateType:         "FLAT",
			StartingAt:       time.Now(),
			BillingFrequency: "MONTHLY",
			CommitRate: shared.CommitRateParam{
				RateType: shared.CommitRateRateTypeFlat,
				Price:    metronome.Float(0),
				Tiers: []shared.TierParam{{
					Price: 0,
					Size:  metronome.Float(0),
				}},
			},
			CreditTypeID: metronome.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			CustomRate: map[string]any{
				"foo": "bar",
			},
			EndingBefore: metronome.Time(time.Now()),
			IsProrated:   metronome.Bool(true),
			Price:        metronome.Float(100),
			PricingGroupValues: map[string]string{
				"region": "us-west-2",
				"cloud":  "aws",
			},
			Quantity: metronome.Float(0),
			Tiers: []shared.TierParam{{
				Price: 0,
				Size:  metronome.Float(0),
			}},
		}, {
			Entitled:         true,
			ProductID:        "13117714-3f05-48e5-a6e9-a66093f13b4d",
			RateType:         "FLAT",
			StartingAt:       time.Now(),
			BillingFrequency: "MONTHLY",
			CommitRate: shared.CommitRateParam{
				RateType: shared.CommitRateRateTypeFlat,
				Price:    metronome.Float(0),
				Tiers: []shared.TierParam{{
					Price: 0,
					Size:  metronome.Float(0),
				}},
			},
			CreditTypeID: metronome.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			CustomRate: map[string]any{
				"foo": "bar",
			},
			EndingBefore: metronome.Time(time.Now()),
			IsProrated:   metronome.Bool(true),
			Price:        metronome.Float(120),
			PricingGroupValues: map[string]string{
				"region": "us-east-2",
				"cloud":  "aws",
			},
			Quantity: metronome.Float(0),
			Tiers: []shared.TierParam{{
				Price: 0,
				Size:  metronome.Float(0),
			}},
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
