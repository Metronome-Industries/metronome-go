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

func TestContractNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.New(context.TODO(), metronome.ContractNewParams{
		CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		StartingAt: metronome.F(time.Now()),
		BillingProviderConfiguration: metronome.F(metronome.ContractNewParamsBillingProviderConfiguration{
			BillingProvider:                metronome.F(metronome.ContractNewParamsBillingProviderConfigurationBillingProviderAwsMarketplace),
			BillingProviderConfigurationID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DeliveryMethod:                 metronome.F(metronome.ContractNewParamsBillingProviderConfigurationDeliveryMethodDirectToBillingProvider),
		}),
		Commits: metronome.F([]metronome.ContractNewParamsCommit{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Type:      metronome.F(metronome.ContractNewParamsCommitsTypePrepaid),
			AccessSchedule: metronome.F(metronome.ContractNewParamsCommitsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					EndingBefore: metronome.F(time.Now()),
					StartingAt:   metronome.F(time.Now()),
				}}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			Amount:                metronome.F(0.000000),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Description: metronome.F("description"),
			InvoiceSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			Priority:             metronome.F(0.000000),
			RateType:             metronome.F(metronome.ContractNewParamsCommitsRateTypeCommitRate),
			RolloverFraction:     metronome.F(0.000000),
			TemporaryID:          metronome.F("temporary_id"),
		}}),
		Credits: metronome.F([]metronome.ContractNewParamsCredit{{
			AccessSchedule: metronome.F(metronome.ContractNewParamsCreditsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCreditsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					EndingBefore: metronome.F(time.Now()),
					StartingAt:   metronome.F(time.Now()),
				}}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Description:          metronome.F("description"),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			Priority:             metronome.F(0.000000),
			RateType:             metronome.F(metronome.ContractNewParamsCreditsRateTypeCommitRate),
		}}),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		Discounts: metronome.F([]metronome.ContractNewParamsDiscount{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Schedule: metronome.F(metronome.ContractNewParamsDiscountsSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsDiscountsScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		EndingBefore:                     metronome.F(time.Now()),
		MultiplierOverridePrioritization: metronome.F(metronome.ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier),
		Name:                             metronome.F("name"),
		NetPaymentTermsDays:              metronome.F(0.000000),
		NetsuiteSalesOrderID:             metronome.F("netsuite_sales_order_id"),
		Overrides: metronome.F([]metronome.ContractNewParamsOverride{{
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductTags: metronome.F([]string{"string"}),
			EndingBefore:          metronome.F(time.Now()),
			Entitled:              metronome.F(true),
			IsCommitSpecific:      metronome.F(true),
			Multiplier:            metronome.F(0.000000),
			OverrideSpecifiers: metronome.F([]metronome.ContractNewParamsOverridesOverrideSpecifier{{
				CommitIDs: metronome.F([]string{"string"}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				ProductID:          metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags:        metronome.F([]string{"string"}),
				RecurringCommitIDs: metronome.F([]string{"string"}),
				RecurringCreditIDs: metronome.F([]string{"string"}),
			}}),
			OverwriteRate: metronome.F(metronome.ContractNewParamsOverridesOverwriteRate{
				RateType:     metronome.F(metronome.ContractNewParamsOverridesOverwriteRateRateTypeFlat),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				CustomRate: metronome.F(map[string]interface{}{
					"foo": "bar",
				}),
				IsProrated: metronome.F(true),
				Price:      metronome.F(0.000000),
				Quantity:   metronome.F(0.000000),
				Tiers: metronome.F([]shared.TierParam{{
					Price: metronome.F(0.000000),
					Size:  metronome.F(0.000000),
				}}),
			}),
			Priority:  metronome.F(0.000000),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Target:    metronome.F(metronome.ContractNewParamsOverridesTargetCommitRate),
			Tiers: metronome.F([]metronome.ContractNewParamsOverridesTier{{
				Multiplier: metronome.F(0.000000),
				Size:       metronome.F(0.000000),
			}}),
			Type: metronome.F(metronome.ContractNewParamsOverridesTypeOverwrite),
		}}),
		ProfessionalServices: metronome.F([]metronome.ContractNewParamsProfessionalService{{
			MaxAmount: metronome.F(0.000000),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Quantity:  metronome.F(0.000000),
			UnitPrice: metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Description:          metronome.F("description"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		RateCardAlias: metronome.F("rate_card_alias"),
		RateCardID:    metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		RecurringCommits: metronome.F([]metronome.ContractNewParamsRecurringCommit{{
			AccessAmount: metronome.F(metronome.ContractNewParamsRecurringCommitsAccessAmount{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Quantity:     metronome.F(0.000000),
				UnitPrice:    metronome.F(0.000000),
			}),
			CommitDuration: metronome.F(metronome.ContractNewParamsRecurringCommitsCommitDuration{
				Unit:  metronome.F(metronome.ContractNewParamsRecurringCommitsCommitDurationUnitPeriods),
				Value: metronome.F(0.000000),
			}),
			Priority:              metronome.F(0.000000),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			Description:           metronome.F("description"),
			EndingBefore:          metronome.F(time.Now()),
			InvoiceAmount: metronome.F(metronome.ContractNewParamsRecurringCommitsInvoiceAmount{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Quantity:     metronome.F(0.000000),
				UnitPrice:    metronome.F(0.000000),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			RateType:             metronome.F(metronome.ContractNewParamsRecurringCommitsRateTypeCommitRate),
			RolloverFraction:     metronome.F(0.000000),
			TemporaryID:          metronome.F("temporary_id"),
		}}),
		RecurringCredits: metronome.F([]metronome.ContractNewParamsRecurringCredit{{
			AccessAmount: metronome.F(metronome.ContractNewParamsRecurringCreditsAccessAmount{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Quantity:     metronome.F(0.000000),
				UnitPrice:    metronome.F(0.000000),
			}),
			CommitDuration: metronome.F(metronome.ContractNewParamsRecurringCreditsCommitDuration{
				Unit:  metronome.F(metronome.ContractNewParamsRecurringCreditsCommitDurationUnitPeriods),
				Value: metronome.F(0.000000),
			}),
			Priority:              metronome.F(0.000000),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			Description:           metronome.F("description"),
			EndingBefore:          metronome.F(time.Now()),
			Name:                  metronome.F("x"),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			RateType:              metronome.F(metronome.ContractNewParamsRecurringCreditsRateTypeCommitRate),
			RolloverFraction:      metronome.F(0.000000),
			TemporaryID:           metronome.F("temporary_id"),
		}}),
		ResellerRoyalties: metronome.F([]metronome.ContractNewParamsResellerRoyalty{{
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("netsuite_reseller_id"),
			ResellerType:          metronome.F(metronome.ContractNewParamsResellerRoyaltiesResellerTypeAws),
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			AwsOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("aws_account_number"),
				AwsOfferID:          metronome.F("aws_offer_id"),
				AwsPayerReferenceID: metronome.F("aws_payer_reference_id"),
			}),
			EndingBefore: metronome.F(time.Now()),
			GcpOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("gcp_account_id"),
				GcpOfferID:   metronome.F("gcp_offer_id"),
			}),
			ResellerContractValue: metronome.F(0.000000),
		}}),
		SalesforceOpportunityID: metronome.F("salesforce_opportunity_id"),
		ScheduledCharges: metronome.F([]metronome.ContractNewParamsScheduledCharge{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Schedule: metronome.F(metronome.ContractNewParamsScheduledChargesSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		ScheduledChargesOnUsageInvoices: metronome.F(metronome.ContractNewParamsScheduledChargesOnUsageInvoicesAll),
		TotalContractValue:              metronome.F(0.000000),
		Transition: metronome.F(metronome.ContractNewParamsTransition{
			FromContractID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Type:           metronome.F(metronome.ContractNewParamsTransitionTypeSupersede),
			FutureInvoiceBehavior: metronome.F(metronome.ContractNewParamsTransitionFutureInvoiceBehavior{
				Trueup: metronome.F(metronome.ContractNewParamsTransitionFutureInvoiceBehaviorTrueupRemove),
			}),
		}),
		UniquenessKey: metronome.F("x"),
		UsageFilter: metronome.F(shared.BaseUsageFilterParam{
			GroupKey:    metronome.F("group_key"),
			GroupValues: metronome.F([]string{"string"}),
			StartingAt:  metronome.F(time.Now()),
		}),
		UsageStatementSchedule: metronome.F(metronome.ContractNewParamsUsageStatementSchedule{
			Frequency:                   metronome.F(metronome.ContractNewParamsUsageStatementScheduleFrequencyMonthly),
			BillingAnchorDate:           metronome.F(time.Now()),
			Day:                         metronome.F(metronome.ContractNewParamsUsageStatementScheduleDayFirstOfMonth),
			InvoiceGenerationStartingAt: metronome.F(time.Now()),
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

func TestContractGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.Get(context.TODO(), metronome.ContractGetParams{
		ContractID:     metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID:     metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		IncludeBalance: metronome.F(true),
		IncludeLedgers: metronome.F(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractListWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.List(context.TODO(), metronome.ContractListParams{
		CustomerID:      metronome.F("9b85c1c1-5238-4f2a-a409-61412905e1e1"),
		CoveringDate:    metronome.F(time.Now()),
		IncludeArchived: metronome.F(true),
		IncludeBalance:  metronome.F(true),
		IncludeLedgers:  metronome.F(true),
		StartingAt:      metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractAddManualBalanceEntryWithOptionalParams(t *testing.T) {
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
	err := client.Contracts.AddManualBalanceEntry(context.TODO(), metronome.ContractAddManualBalanceEntryParams{
		ID:         metronome.F("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
		Amount:     metronome.F(-1000.000000),
		CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		Reason:     metronome.F("Reason for entry"),
		SegmentID:  metronome.F("66368e29-3f97-4d15-a6e9-120897f0070a"),
		ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		Timestamp:  metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractAmendWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.Amend(context.TODO(), metronome.ContractAmendParams{
		ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		StartingAt: metronome.F(time.Now()),
		Commits: metronome.F([]metronome.ContractAmendParamsCommit{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Type:      metronome.F(metronome.ContractAmendParamsCommitsTypePrepaid),
			AccessSchedule: metronome.F(metronome.ContractAmendParamsCommitsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					EndingBefore: metronome.F(time.Now()),
					StartingAt:   metronome.F(time.Now()),
				}}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			Amount:                metronome.F(0.000000),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Description: metronome.F("description"),
			InvoiceSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			Priority:             metronome.F(0.000000),
			RateType:             metronome.F(metronome.ContractAmendParamsCommitsRateTypeCommitRate),
			RolloverFraction:     metronome.F(0.000000),
			TemporaryID:          metronome.F("temporary_id"),
		}}),
		Credits: metronome.F([]metronome.ContractAmendParamsCredit{{
			AccessSchedule: metronome.F(metronome.ContractAmendParamsCreditsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCreditsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					EndingBefore: metronome.F(time.Now()),
					StartingAt:   metronome.F(time.Now()),
				}}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Description:          metronome.F("description"),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			Priority:             metronome.F(0.000000),
			RateType:             metronome.F(metronome.ContractAmendParamsCreditsRateTypeCommitRate),
		}}),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		Discounts: metronome.F([]metronome.ContractAmendParamsDiscount{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Schedule: metronome.F(metronome.ContractAmendParamsDiscountsSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsDiscountsScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		Overrides: metronome.F([]metronome.ContractAmendParamsOverride{{
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductTags: metronome.F([]string{"string"}),
			EndingBefore:          metronome.F(time.Now()),
			Entitled:              metronome.F(true),
			IsCommitSpecific:      metronome.F(true),
			Multiplier:            metronome.F(0.000000),
			OverrideSpecifiers: metronome.F([]metronome.ContractAmendParamsOverridesOverrideSpecifier{{
				CommitIDs: metronome.F([]string{"string"}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				ProductID:          metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags:        metronome.F([]string{"string"}),
				RecurringCommitIDs: metronome.F([]string{"string"}),
				RecurringCreditIDs: metronome.F([]string{"string"}),
			}}),
			OverwriteRate: metronome.F(metronome.ContractAmendParamsOverridesOverwriteRate{
				RateType:     metronome.F(metronome.ContractAmendParamsOverridesOverwriteRateRateTypeFlat),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				CustomRate: metronome.F(map[string]interface{}{
					"foo": "bar",
				}),
				IsProrated: metronome.F(true),
				Price:      metronome.F(0.000000),
				Quantity:   metronome.F(0.000000),
				Tiers: metronome.F([]shared.TierParam{{
					Price: metronome.F(0.000000),
					Size:  metronome.F(0.000000),
				}}),
			}),
			Priority:  metronome.F(0.000000),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Target:    metronome.F(metronome.ContractAmendParamsOverridesTargetCommitRate),
			Tiers: metronome.F([]metronome.ContractAmendParamsOverridesTier{{
				Multiplier: metronome.F(0.000000),
				Size:       metronome.F(0.000000),
			}}),
			Type: metronome.F(metronome.ContractAmendParamsOverridesTypeOverwrite),
		}}),
		ProfessionalServices: metronome.F([]metronome.ContractAmendParamsProfessionalService{{
			MaxAmount: metronome.F(0.000000),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Quantity:  metronome.F(0.000000),
			UnitPrice: metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Description:          metronome.F("description"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		ResellerRoyalties: metronome.F([]metronome.ContractAmendParamsResellerRoyalty{{
			ResellerType:          metronome.F(metronome.ContractAmendParamsResellerRoyaltiesResellerTypeAws),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			AwsOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("aws_account_number"),
				AwsOfferID:          metronome.F("aws_offer_id"),
				AwsPayerReferenceID: metronome.F("aws_payer_reference_id"),
			}),
			EndingBefore: metronome.F(time.Now()),
			Fraction:     metronome.F(0.000000),
			GcpOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("gcp_account_id"),
				GcpOfferID:   metronome.F("gcp_offer_id"),
			}),
			NetsuiteResellerID:    metronome.F("netsuite_reseller_id"),
			ResellerContractValue: metronome.F(0.000000),
			StartingAt:            metronome.F(time.Now()),
		}}),
		SalesforceOpportunityID: metronome.F("salesforce_opportunity_id"),
		ScheduledCharges: metronome.F([]metronome.ContractAmendParamsScheduledCharge{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Schedule: metronome.F(metronome.ContractAmendParamsScheduledChargesSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		TotalContractValue: metronome.F(0.000000),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractArchive(t *testing.T) {
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
	_, err := client.Contracts.Archive(context.TODO(), metronome.ContractArchiveParams{
		ContractID:   metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID:   metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		VoidInvoices: metronome.F(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractNewHistoricalInvoices(t *testing.T) {
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
	_, err := client.Contracts.NewHistoricalInvoices(context.TODO(), metronome.ContractNewHistoricalInvoicesParams{
		Invoices: metronome.F([]metronome.ContractNewHistoricalInvoicesParamsInvoice{{
			ContractID:         metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
			CreditTypeID:       metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			CustomerID:         metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
			ExclusiveEndDate:   metronome.F(time.Now()),
			InclusiveStartDate: metronome.F(time.Now()),
			IssueDate:          metronome.F(time.Now()),
			UsageLineItems: metronome.F([]metronome.ContractNewHistoricalInvoicesParamsInvoicesUsageLineItem{{
				ExclusiveEndDate:   metronome.F(time.Now()),
				InclusiveStartDate: metronome.F(time.Now()),
				ProductID:          metronome.F("f14d6729-6a44-4b13-9908-9387f1918790"),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				Quantity: metronome.F(100.000000),
				SubtotalsWithQuantity: metronome.F([]metronome.ContractNewHistoricalInvoicesParamsInvoicesUsageLineItemsSubtotalsWithQuantity{{
					ExclusiveEndDate:   metronome.F(time.Now()),
					InclusiveStartDate: metronome.F(time.Now()),
					Quantity:           metronome.F(0.000000),
				}}),
			}}),
			BillableStatus:       metronome.F(metronome.ContractNewHistoricalInvoicesParamsInvoicesBillableStatusBillable),
			BreakdownGranularity: metronome.F(metronome.ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularityHour),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}}),
		Preview: metronome.F(false),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractListBalancesWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.ListBalances(context.TODO(), metronome.ContractListBalancesParams{
		CustomerID:              metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		ID:                      metronome.F("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
		CoveringDate:            metronome.F(time.Now()),
		EffectiveBefore:         metronome.F(time.Now()),
		IncludeArchived:         metronome.F(true),
		IncludeBalance:          metronome.F(true),
		IncludeContractBalances: metronome.F(true),
		IncludeLedgers:          metronome.F(true),
		NextPage:                metronome.F("next_page"),
		StartingAt:              metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractGetRateScheduleWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.GetRateSchedule(context.TODO(), metronome.ContractGetRateScheduleParams{
		ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		Limit:      metronome.F(int64(1)),
		NextPage:   metronome.F("next_page"),
		At:         metronome.F(time.Now()),
		Selectors: metronome.F([]metronome.ContractGetRateScheduleParamsSelector{{
			PartialPricingGroupValues: metronome.F(map[string]string{
				"region": "us-west-2",
				"cloud":  "aws",
			}),
			PricingGroupValues: metronome.F(map[string]string{
				"foo": "string",
			}),
			ProductID:   metronome.F("d6300dbb-882e-4d2d-8dec-5125d16b65d0"),
			ProductTags: metronome.F([]string{"string"}),
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

func TestContractScheduleProServicesInvoiceWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.ScheduleProServicesInvoice(context.TODO(), metronome.ContractScheduleProServicesInvoiceParams{
		ContractID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CustomerID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IssuedAt:   metronome.F(time.Now()),
		LineItems: metronome.F([]metronome.ContractScheduleProServicesInvoiceParamsLineItem{{
			ProfessionalServiceID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AmendmentID:                 metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Amount:                      metronome.F(0.000000),
			Metadata:                    metronome.F("metadata"),
			NetsuiteInvoiceBillingEnd:   metronome.F(time.Now()),
			NetsuiteInvoiceBillingStart: metronome.F(time.Now()),
			Quantity:                    metronome.F(0.000000),
			UnitPrice:                   metronome.F(0.000000),
		}}),
		NetsuiteInvoiceHeaderEnd:   metronome.F(time.Now()),
		NetsuiteInvoiceHeaderStart: metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractSetUsageFilter(t *testing.T) {
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
	err := client.Contracts.SetUsageFilter(context.TODO(), metronome.ContractSetUsageFilterParams{
		ContractID:  metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID:  metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		GroupKey:    metronome.F("business_subscription_id"),
		GroupValues: metronome.F([]string{"ID-1", "ID-2"}),
		StartingAt:  metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractUpdateEndDateWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.UpdateEndDate(context.TODO(), metronome.ContractUpdateEndDateParams{
		ContractID:                        metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID:                        metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		AllowEndingBeforeFinalizedInvoice: metronome.F(true),
		EndingBefore:                      metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
