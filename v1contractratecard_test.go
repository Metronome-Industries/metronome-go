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
		Name: metronome.F("My Rate Card"),
		Aliases: metronome.F([]metronome.V1ContractRateCardNewParamsAlias{{
			Name:         metronome.F("my-rate-card"),
			EndingBefore: metronome.F(time.Now()),
			StartingAt:   metronome.F(time.Now()),
		}}),
		CreditTypeConversions: metronome.F([]metronome.V1ContractRateCardNewParamsCreditTypeConversion{{
			CustomCreditTypeID:  metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			FiatPerCustomCredit: metronome.F(2.000000),
		}}),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		Description:      metronome.F("My Rate Card Description"),
		FiatCreditTypeID: metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
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
		ID: metronome.F("f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe"),
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
		RateCardID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		Aliases: metronome.F([]metronome.V1ContractRateCardUpdateParamsAlias{{
			Name:         metronome.F("name"),
			EndingBefore: metronome.F(time.Now()),
			StartingAt:   metronome.F(time.Now()),
		}}),
		Description: metronome.F("My Updated Rate Card Description"),
		Name:        metronome.F("My Updated Rate Card"),
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
		Limit:    metronome.F(int64(1)),
		NextPage: metronome.F("next_page"),
		Body:     map[string]interface{}{},
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
		ID: metronome.F("12b21470-4570-40df-8998-449d0b0bc52f"),
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
		RateCardID:   metronome.F("f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe"),
		StartingAt:   metronome.F(time.Now()),
		Limit:        metronome.F(int64(1)),
		NextPage:     metronome.F("next_page"),
		EndingBefore: metronome.F(time.Now()),
		Selectors: metronome.F([]metronome.V1ContractRateCardGetRateScheduleParamsSelector{{
			BillingFrequency: metronome.F(metronome.V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequencyMonthly),
			PartialPricingGroupValues: metronome.F(map[string]string{
				"region": "us-west-2",
				"cloud":  "aws",
			}),
			PricingGroupValues: metronome.F(map[string]string{
				"foo": "string",
			}),
			ProductID: metronome.F("d6300dbb-882e-4d2d-8dec-5125d16b65d0"),
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
