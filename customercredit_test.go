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

func TestCustomerCreditNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.New(context.TODO(), metronome.CustomerCreditNewParams{
		AccessSchedule: metronome.F(metronome.CustomerCreditNewParamsAccessSchedule{
			CreditTypeID: metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			ScheduleItems: metronome.F([]metronome.CustomerCreditNewParamsAccessScheduleScheduleItem{{
				Amount:       metronome.F(1000.000000),
				StartingAt:   metronome.F(time.Now()),
				EndingBefore: metronome.F(time.Now()),
			}}),
		}),
		CustomerID:            metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		Priority:              metronome.F(100.000000),
		ProductID:             metronome.F("f14d6729-6a44-4b13-9908-9387f1918790"),
		ApplicableContractIDs: metronome.F([]string{"string", "string", "string"}),
		ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		Description:             metronome.F("description"),
		Name:                    metronome.F("My Credit"),
		NetsuiteSalesOrderID:    metronome.F("netsuite_sales_order_id"),
		SalesforceOpportunityID: metronome.F("salesforce_opportunity_id"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerCreditListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.List(context.TODO(), metronome.CustomerCreditListParams{
		CustomerID:             metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		CoveringDate:           metronome.F(time.Now()),
		CreditID:               metronome.F("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
		EffectiveBefore:        metronome.F(time.Now()),
		IncludeArchived:        metronome.F(true),
		IncludeContractCredits: metronome.F(true),
		IncludeLedgers:         metronome.F(true),
		NextPage:               metronome.F("next_page"),
		StartingAt:             metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerCreditUpdateEndDate(t *testing.T) {
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
	_, err := client.Customers.Credits.UpdateEndDate(context.TODO(), metronome.CustomerCreditUpdateEndDateParams{
		AccessEndingBefore: metronome.F(time.Now()),
		CreditID:           metronome.F("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
		CustomerID:         metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
