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
	_, err := client.Customers.Plans.List(context.TODO(), metronome.CustomerPlanListParams{
		CustomerID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		Limit:      metronome.F(int64(1)),
		NextPage:   metronome.F("next_page"),
	})
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
	_, err := client.Customers.Plans.Add(context.TODO(), metronome.CustomerPlanAddParams{
		CustomerID:          metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
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
			AdjustmentType: metronome.F(metronome.CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypePercentage),
			ChargeID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			StartPeriod:    metronome.F(0.000000),
			Quantity:       metronome.F(0.000000),
			Tier:           metronome.F(0.000000),
			Value:          metronome.F(0.000000),
		}, {
			AdjustmentType: metronome.F(metronome.CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypePercentage),
			ChargeID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			StartPeriod:    metronome.F(0.000000),
			Quantity:       metronome.F(0.000000),
			Tier:           metronome.F(0.000000),
			Value:          metronome.F(0.000000),
		}, {
			AdjustmentType: metronome.F(metronome.CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypePercentage),
			ChargeID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			StartPeriod:    metronome.F(0.000000),
			Quantity:       metronome.F(0.000000),
			Tier:           metronome.F(0.000000),
			Value:          metronome.F(0.000000),
		}}),
		TrialSpec: metronome.F(metronome.CustomerPlanAddParamsTrialSpec{
			LengthInDays: metronome.F(0.000000),
			SpendingCap: metronome.F(metronome.CustomerPlanAddParamsTrialSpecSpendingCap{
				Amount:       metronome.F(0.000000),
				CreditTypeID: metronome.F("credit_type_id"),
			}),
		}),
	})
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
	_, err := client.Customers.Plans.End(context.TODO(), metronome.CustomerPlanEndParams{
		CustomerID:         metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerPlanID:     metronome.F("7aa11640-0703-4600-8eb9-293f535a6b74"),
		EndingBefore:       metronome.F(time.Now()),
		VoidInvoices:       metronome.F(true),
		VoidStripeInvoices: metronome.F(true),
	})
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
	_, err := client.Customers.Plans.ListPriceAdjustments(context.TODO(), metronome.CustomerPlanListPriceAdjustmentsParams{
		CustomerID:     metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerPlanID: metronome.F("7aa11640-0703-4600-8eb9-293f535a6b74"),
		Limit:          metronome.F(int64(1)),
		NextPage:       metronome.F("next_page"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
