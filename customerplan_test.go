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

func TestCustomerPlanListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Plans.List(
		context.TODO(),
		"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		metronome.CustomerPlanListParams{
			Limit:    metronome.F(int64(1)),
			NextPage: metronome.F("next_page"),
		},
	)
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerPlanAddWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Plans.Add(
		context.TODO(),
		"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		metronome.CustomerPlanAddParams{
			PlanID:              metronome.F("d2c06dae-9549-4d7d-bc04-b78dd3d241b8"),
			StartingOn:          metronome.F(time.Now()),
			EndingBefore:        metronome.F(time.Now()),
			NetPaymentTermsDays: metronome.F(0.000000),
			OverageRateAdjustments: metronome.F([]metronome.CustomerPlanAddParamsOverageRateAdjustment{{
				CustomCreditTypeID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				FiatCurrencyCreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ToFiatConversionFactor:   metronome.F(0.000000),
			}, {
				CustomCreditTypeID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				FiatCurrencyCreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ToFiatConversionFactor:   metronome.F(0.000000),
			}, {
				CustomCreditTypeID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				FiatCurrencyCreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ToFiatConversionFactor:   metronome.F(0.000000),
			}}),
			PriceAdjustments: metronome.F([]metronome.CustomerPlanAddParamsPriceAdjustment{{
				ChargeID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				AdjustmentType: metronome.F(metronome.CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypePercentage),
				Value:          metronome.F(0.000000),
				Quantity:       metronome.F(0.000000),
				Tier:           metronome.F(0.000000),
				StartPeriod:    metronome.F(0.000000),
			}, {
				ChargeID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				AdjustmentType: metronome.F(metronome.CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypePercentage),
				Value:          metronome.F(0.000000),
				Quantity:       metronome.F(0.000000),
				Tier:           metronome.F(0.000000),
				StartPeriod:    metronome.F(0.000000),
			}, {
				ChargeID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				AdjustmentType: metronome.F(metronome.CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypePercentage),
				Value:          metronome.F(0.000000),
				Quantity:       metronome.F(0.000000),
				Tier:           metronome.F(0.000000),
				StartPeriod:    metronome.F(0.000000),
			}}),
			TrialSpec: metronome.F(metronome.CustomerPlanAddParamsTrialSpec{
				LengthInDays: metronome.F(0.000000),
				SpendingCap: metronome.F(metronome.CustomerPlanAddParamsTrialSpecSpendingCap{
					CreditTypeID: metronome.F("credit_type_id"),
					Amount:       metronome.F(0.000000),
				}),
			}),
		},
	)
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerPlanEndWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Plans.End(
		context.TODO(),
		"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		"7aa11640-0703-4600-8eb9-293f535a6b74",
		metronome.CustomerPlanEndParams{
			EndingBefore:       metronome.F(time.Now()),
			VoidInvoices:       metronome.F(true),
			VoidStripeInvoices: metronome.F(true),
		},
	)
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerPlanListPriceAdjustmentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Plans.ListPriceAdjustments(
		context.TODO(),
		"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		"7aa11640-0703-4600-8eb9-293f535a6b74",
		metronome.CustomerPlanListPriceAdjustmentsParams{
			Limit:    metronome.F(int64(1)),
			NextPage: metronome.F("next_page"),
		},
	)
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
