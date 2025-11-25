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

func TestV1ContractRateCardNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.RateCards.New(context.TODO(), metronome.V1ContractRateCardNewParams{
		Name: "My Rate Card",
		Aliases: []metronome.V1ContractRateCardNewParamsAlias{{
			Name:         "my-rate-card",
			EndingBefore: metronome.Time(time.Now()),
			StartingAt:   metronome.Time(time.Now()),
		}},
		CreditTypeConversions: []metronome.V1ContractRateCardNewParamsCreditTypeConversion{{
			CustomCreditTypeID:  "2714e483-4ff1-48e4-9e25-ac732e8f24f2",
			FiatPerCustomCredit: 2,
		}},
		CustomFields: map[string]string{
			"foo": "string",
		},
		Description:      metronome.String("My Rate Card Description"),
		FiatCreditTypeID: metronome.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContractRateCardGet(t *testing.T) {
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
	_, err := client.V1.Contracts.RateCards.Get(context.TODO(), metronome.V1ContractRateCardGetParams{
		ID: shared.IDParam{
			ID: "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe",
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

func TestV1ContractRateCardUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.RateCards.Update(context.TODO(), metronome.V1ContractRateCardUpdateParams{
		RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		Aliases: []metronome.V1ContractRateCardUpdateParamsAlias{{
			Name:         "name",
			EndingBefore: metronome.Time(time.Now()),
			StartingAt:   metronome.Time(time.Now()),
		}},
		Description: metronome.String("My Updated Rate Card Description"),
		Name:        metronome.String("My Updated Rate Card"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContractRateCardListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.RateCards.List(context.TODO(), metronome.V1ContractRateCardListParams{
		Limit:    metronome.Int(1),
		NextPage: metronome.String("next_page"),
		Body:     map[string]any{},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContractRateCardArchive(t *testing.T) {
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
	_, err := client.V1.Contracts.RateCards.Archive(context.TODO(), metronome.V1ContractRateCardArchiveParams{
		ID: shared.IDParam{
			ID: "12b21470-4570-40df-8998-449d0b0bc52f",
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

func TestV1ContractRateCardGetRateScheduleWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.RateCards.GetRateSchedule(context.TODO(), metronome.V1ContractRateCardGetRateScheduleParams{
		RateCardID:   "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe",
		StartingAt:   time.Now(),
		Limit:        metronome.Int(1),
		NextPage:     metronome.String("next_page"),
		EndingBefore: metronome.Time(time.Now()),
		Selectors: []metronome.V1ContractRateCardGetRateScheduleParamsSelector{{
			BillingFrequency: "MONTHLY",
			PartialPricingGroupValues: map[string]string{
				"region": "us-west-2",
				"cloud":  "aws",
			},
			PricingGroupValues: map[string]string{
				"foo": "string",
			},
			ProductID: metronome.String("d6300dbb-882e-4d2d-8dec-5125d16b65d0"),
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
