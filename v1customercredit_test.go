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

func TestV1CustomerCreditNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Credits.New(context.TODO(), metronome.V1CustomerCreditNewParams{
		AccessSchedule: metronome.V1CustomerCreditNewParamsAccessSchedule{
			ScheduleItems: []metronome.V1CustomerCreditNewParamsAccessScheduleScheduleItem{{
				Amount:       1000,
				EndingBefore: time.Now(),
				StartingAt:   time.Now(),
			}},
			CreditTypeID: metronome.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
		},
		CustomerID:            "13117714-3f05-48e5-a6e9-a66093f13b4d",
		Priority:              100,
		ProductID:             "f14d6729-6a44-4b13-9908-9387f1918790",
		ApplicableContractIDs: []string{"string"},
		ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		ApplicableProductTags: []string{"string"},
		CustomFields: map[string]string{
			"foo": "string",
		},
		Description:             metronome.String("description"),
		Name:                    metronome.String("My Credit"),
		NetsuiteSalesOrderID:    metronome.String("netsuite_sales_order_id"),
		RateType:                metronome.V1CustomerCreditNewParamsRateTypeCommitRate,
		SalesforceOpportunityID: metronome.String("salesforce_opportunity_id"),
		Specifiers: []shared.CommitSpecifierInputParam{{
			PresentationGroupValues: map[string]string{
				"foo": "string",
			},
			PricingGroupValues: map[string]string{
				"foo": "string",
			},
			ProductID:   metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProductTags: []string{"string"},
		}},
		UniquenessKey: metronome.String("x"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerCreditListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Credits.List(context.TODO(), metronome.V1CustomerCreditListParams{
		CustomerID:             "13117714-3f05-48e5-a6e9-a66093f13b4d",
		CoveringDate:           metronome.Time(time.Now()),
		CreditID:               metronome.String("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
		EffectiveBefore:        metronome.Time(time.Now()),
		IncludeArchived:        metronome.Bool(true),
		IncludeBalance:         metronome.Bool(true),
		IncludeContractCredits: metronome.Bool(true),
		IncludeLedgers:         metronome.Bool(true),
		Limit:                  metronome.Int(1),
		NextPage:               metronome.String("next_page"),
		StartingAt:             metronome.Time(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerCreditUpdateEndDate(t *testing.T) {
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
	_, err := client.V1.Customers.Credits.UpdateEndDate(context.TODO(), metronome.V1CustomerCreditUpdateEndDateParams{
		AccessEndingBefore: time.Now(),
		CreditID:           "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
		CustomerID:         "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
