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

func TestV1CustomerCommitNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Commits.New(context.TODO(), metronome.V1CustomerCommitNewParams{
		AccessSchedule: metronome.V1CustomerCommitNewParamsAccessSchedule{
			ScheduleItems: []metronome.V1CustomerCommitNewParamsAccessScheduleScheduleItem{{
				Amount:       1000,
				EndingBefore: time.Now(),
				StartingAt:   time.Now(),
			}},
			CreditTypeID: metronome.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
		},
		CustomerID:            "13117714-3f05-48e5-a6e9-a66093f13b4d",
		Priority:              100,
		ProductID:             "f14d6729-6a44-4b13-9908-9387f1918790",
		Type:                  metronome.V1CustomerCommitNewParamsTypePrepaid,
		ApplicableContractIDs: []string{"string"},
		ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		ApplicableProductTags: []string{"string"},
		CustomFields: map[string]string{
			"foo": "string",
		},
		Description:       metronome.String("description"),
		InvoiceContractID: metronome.String("e57d6929-c2f1-4796-a9a8-63cedefe848d"),
		InvoiceSchedule: metronome.V1CustomerCommitNewParamsInvoiceSchedule{
			CreditTypeID: metronome.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			DoNotInvoice: metronome.Bool(false),
			RecurringSchedule: metronome.V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule{
				AmountDistribution: "DIVIDED",
				EndingBefore:       time.Now(),
				Frequency:          "MONTHLY",
				StartingAt:         time.Now(),
				Amount:             metronome.Float(0),
				Quantity:           metronome.Float(0),
				UnitPrice:          metronome.Float(0),
			},
			ScheduleItems: []metronome.V1CustomerCommitNewParamsInvoiceScheduleScheduleItem{{
				Timestamp: time.Now(),
				Amount:    metronome.Float(0),
				Quantity:  metronome.Float(1),
				UnitPrice: metronome.Float(10000000),
			}},
		},
		Name:                    metronome.String("My Commit"),
		NetsuiteSalesOrderID:    metronome.String("netsuite_sales_order_id"),
		RateType:                metronome.V1CustomerCommitNewParamsRateTypeCommitRate,
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

func TestV1CustomerCommitListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Commits.List(context.TODO(), metronome.V1CustomerCommitListParams{
		CustomerID:             "13117714-3f05-48e5-a6e9-a66093f13b4d",
		CommitID:               metronome.String("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
		CoveringDate:           metronome.Time(time.Now()),
		EffectiveBefore:        metronome.Time(time.Now()),
		IncludeArchived:        metronome.Bool(true),
		IncludeBalance:         metronome.Bool(true),
		IncludeContractCommits: metronome.Bool(true),
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

func TestV1CustomerCommitUpdateEndDateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Commits.UpdateEndDate(context.TODO(), metronome.V1CustomerCommitUpdateEndDateParams{
		CommitID:             "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
		CustomerID:           "13117714-3f05-48e5-a6e9-a66093f13b4d",
		AccessEndingBefore:   metronome.Time(time.Now()),
		InvoicesEndingBefore: metronome.Time(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
