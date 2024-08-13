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

func TestCustomerCommitNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Commits.New(context.TODO(), metronome.CustomerCommitNewParams{
		AccessSchedule: metronome.F(metronome.CustomerCommitNewParamsAccessSchedule{
			ScheduleItems: metronome.F([]metronome.CustomerCommitNewParamsAccessScheduleScheduleItem{{
				Amount:       metronome.F(1000.000000),
				EndingBefore: metronome.F(time.Now()),
				StartingAt:   metronome.F(time.Now()),
			}}),
			CreditTypeID: metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
		}),
		CustomerID:            metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		Priority:              metronome.F(100.000000),
		ProductID:             metronome.F("f14d6729-6a44-4b13-9908-9387f1918790"),
		Type:                  metronome.F(metronome.CustomerCommitNewParamsTypePrepaid),
		ApplicableContractIDs: metronome.F([]string{"string", "string", "string"}),
		ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		Description:       metronome.F("description"),
		InvoiceContractID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		InvoiceSchedule: metronome.F(metronome.CustomerCommitNewParamsInvoiceSchedule{
			CreditTypeID: metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			RecurringSchedule: metronome.F(metronome.CustomerCommitNewParamsInvoiceScheduleRecurringSchedule{
				AmountDistribution: metronome.F(metronome.CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				EndingBefore:       metronome.F(time.Now()),
				Frequency:          metronome.F(metronome.CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyMonthly),
				StartingAt:         metronome.F(time.Now()),
				Amount:             metronome.F(0.000000),
				Quantity:           metronome.F(0.000000),
				UnitPrice:          metronome.F(0.000000),
			}),
			ScheduleItems: metronome.F([]metronome.CustomerCommitNewParamsInvoiceScheduleScheduleItem{{
				Timestamp: metronome.F(time.Now()),
				Amount:    metronome.F(10000000.000000),
				Quantity:  metronome.F(1.000000),
				UnitPrice: metronome.F(10000000.000000),
			}}),
		}),
		Name:                    metronome.F("My Commit"),
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

func TestCustomerCommitListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Commits.List(context.TODO(), metronome.CustomerCommitListParams{
		CustomerID:             metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		CommitID:               metronome.F("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
		CoveringDate:           metronome.F(time.Now()),
		EffectiveBefore:        metronome.F(time.Now()),
		IncludeArchived:        metronome.F(true),
		IncludeContractCommits: metronome.F(true),
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

func TestCustomerCommitUpdateEndDateWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Commits.UpdateEndDate(context.TODO(), metronome.CustomerCommitUpdateEndDateParams{
		CommitID:             metronome.F("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
		CustomerID:           metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		AccessEndingBefore:   metronome.F(time.Now()),
		InvoicesEndingBefore: metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
