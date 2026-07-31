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
)

func TestV1CustomerPlanListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Plans.List(context.TODO(), metronome.V1CustomerPlanListParams{
		CustomerID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		Limit:      metronome.Int(1),
		NextPage:   metronome.String("next_page"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerPlanAddWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Plans.Add(context.TODO(), metronome.V1CustomerPlanAddParams{
		CustomerID:          "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		PlanID:              "d2c06dae-9549-4d7d-bc04-b78dd3d241b8",
		StartingOn:          time.Now(),
		EndingBefore:        metronome.Time(time.Now()),
		NetPaymentTermsDays: metronome.Float(0),
		OverageRateAdjustments: []metronome.V1CustomerPlanAddParamsOverageRateAdjustment{{
			CustomCreditTypeID:       "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			FiatCurrencyCreditTypeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ToFiatConversionFactor:   0,
		}},
		PriceAdjustments: []metronome.V1CustomerPlanAddParamsPriceAdjustment{{
			AdjustmentType: "percentage",
			ChargeID:       "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			StartPeriod:    0,
			Quantity:       metronome.Float(0),
			Tier:           metronome.Float(0),
			Value:          metronome.Float(0),
		}},
		TrialSpec: metronome.V1CustomerPlanAddParamsTrialSpec{
			LengthInDays: 0,
			SpendingCap: metronome.V1CustomerPlanAddParamsTrialSpecSpendingCap{
				Amount:       0,
				CreditTypeID: "credit_type_id",
			},
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

func TestV1CustomerPlanEndWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Plans.End(context.TODO(), metronome.V1CustomerPlanEndParams{
		CustomerID:         "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerPlanID:     "7aa11640-0703-4600-8eb9-293f535a6b74",
		EndingBefore:       metronome.Time(time.Now()),
		VoidInvoices:       metronome.Bool(true),
		VoidStripeInvoices: metronome.Bool(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerPlanListPriceAdjustmentsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Plans.ListPriceAdjustments(context.TODO(), metronome.V1CustomerPlanListPriceAdjustmentsParams{
		CustomerID:     "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerPlanID: "7aa11640-0703-4600-8eb9-293f535a6b74",
		Limit:          metronome.Int(1),
		NextPage:       metronome.String("next_page"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
