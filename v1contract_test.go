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

func TestV1ContractNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.New(context.TODO(), metronome.V1ContractNewParams{
		CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		StartingAt: metronome.F(time.Now()),
		BillingProviderConfiguration: metronome.F(metronome.V1ContractNewParamsBillingProviderConfiguration{
			BillingProvider:                metronome.F(metronome.V1ContractNewParamsBillingProviderConfigurationBillingProviderAwsMarketplace),
			BillingProviderConfigurationID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DeliveryMethod:                 metronome.F(metronome.V1ContractNewParamsBillingProviderConfigurationDeliveryMethodDirectToBillingProvider),
		}),
		Commits: metronome.F([]metronome.V1ContractNewParamsCommit{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Type:      metronome.F(metronome.V1ContractNewParamsCommitsTypePrepaid),
			AccessSchedule: metronome.F(metronome.V1ContractNewParamsCommitsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.V1ContractNewParamsCommitsAccessScheduleScheduleItem{{
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
			InvoiceSchedule: metronome.F(metronome.V1ContractNewParamsCommitsInvoiceSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.V1ContractNewParamsCommitsInvoiceScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.V1ContractNewParamsCommitsInvoiceScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			Priority:             metronome.F(0.000000),
			RateType:             metronome.F(metronome.V1ContractNewParamsCommitsRateTypeCommitRate),
			RolloverFraction:     metronome.F(0.000000),
			TemporaryID:          metronome.F("temporary_id"),
		}}),
		Credits: metronome.F([]metronome.V1ContractNewParamsCredit{{
			AccessSchedule: metronome.F(metronome.V1ContractNewParamsCreditsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.V1ContractNewParamsCreditsAccessScheduleScheduleItem{{
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
			RateType:             metronome.F(metronome.V1ContractNewParamsCreditsRateTypeCommitRate),
		}}),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		Discounts: metronome.F([]metronome.V1ContractNewParamsDiscount{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Schedule: metronome.F(metronome.V1ContractNewParamsDiscountsSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.V1ContractNewParamsDiscountsScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.V1ContractNewParamsDiscountsScheduleScheduleItem{{
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
		MultiplierOverridePrioritization: metronome.F(metronome.V1ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier),
		Name:                             metronome.F("name"),
		NetPaymentTermsDays:              metronome.F(0.000000),
		NetsuiteSalesOrderID:             metronome.F("netsuite_sales_order_id"),
		Overrides: metronome.F([]metronome.V1ContractNewParamsOverride{{
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductTags: metronome.F([]string{"string"}),
			EndingBefore:          metronome.F(time.Now()),
			Entitled:              metronome.F(true),
			IsCommitSpecific:      metronome.F(true),
			Multiplier:            metronome.F(0.000000),
			OverrideSpecifiers: metronome.F([]metronome.V1ContractNewParamsOverridesOverrideSpecifier{{
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
			OverwriteRate: metronome.F(metronome.V1ContractNewParamsOverridesOverwriteRate{
				RateType:     metronome.F(metronome.V1ContractNewParamsOverridesOverwriteRateRateTypeFlat),
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
			Target:    metronome.F(metronome.V1ContractNewParamsOverridesTargetCommitRate),
			Tiers: metronome.F([]metronome.V1ContractNewParamsOverridesTier{{
				Multiplier: metronome.F(0.000000),
				Size:       metronome.F(0.000000),
			}}),
			Type: metronome.F(metronome.V1ContractNewParamsOverridesTypeOverwrite),
		}}),
		ProfessionalServices: metronome.F([]metronome.V1ContractNewParamsProfessionalService{{
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
		RecurringCommits: metronome.F([]metronome.V1ContractNewParamsRecurringCommit{{
			AccessAmount: metronome.F(metronome.V1ContractNewParamsRecurringCommitsAccessAmount{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Quantity:     metronome.F(0.000000),
				UnitPrice:    metronome.F(0.000000),
			}),
			CommitDuration: metronome.F(metronome.V1ContractNewParamsRecurringCommitsCommitDuration{
				Unit:  metronome.F(metronome.V1ContractNewParamsRecurringCommitsCommitDurationUnitPeriods),
				Value: metronome.F(0.000000),
			}),
			Priority:              metronome.F(0.000000),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			Description:           metronome.F("description"),
			EndingBefore:          metronome.F(time.Now()),
			InvoiceAmount: metronome.F(metronome.V1ContractNewParamsRecurringCommitsInvoiceAmount{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Quantity:     metronome.F(0.000000),
				UnitPrice:    metronome.F(0.000000),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			Proration:            metronome.F(metronome.V1ContractNewParamsRecurringCommitsProrationNone),
			RateType:             metronome.F(metronome.V1ContractNewParamsRecurringCommitsRateTypeCommitRate),
			RecurrenceFrequency:  metronome.F(metronome.V1ContractNewParamsRecurringCommitsRecurrenceFrequencyMonthly),
			RolloverFraction:     metronome.F(0.000000),
			TemporaryID:          metronome.F("temporary_id"),
		}}),
		RecurringCredits: metronome.F([]metronome.V1ContractNewParamsRecurringCredit{{
			AccessAmount: metronome.F(metronome.V1ContractNewParamsRecurringCreditsAccessAmount{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Quantity:     metronome.F(0.000000),
				UnitPrice:    metronome.F(0.000000),
			}),
			CommitDuration: metronome.F(metronome.V1ContractNewParamsRecurringCreditsCommitDuration{
				Unit:  metronome.F(metronome.V1ContractNewParamsRecurringCreditsCommitDurationUnitPeriods),
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
			Proration:             metronome.F(metronome.V1ContractNewParamsRecurringCreditsProrationNone),
			RateType:              metronome.F(metronome.V1ContractNewParamsRecurringCreditsRateTypeCommitRate),
			RecurrenceFrequency:   metronome.F(metronome.V1ContractNewParamsRecurringCreditsRecurrenceFrequencyMonthly),
			RolloverFraction:      metronome.F(0.000000),
			TemporaryID:           metronome.F("temporary_id"),
		}}),
		ResellerRoyalties: metronome.F([]metronome.V1ContractNewParamsResellerRoyalty{{
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("netsuite_reseller_id"),
			ResellerType:          metronome.F(metronome.V1ContractNewParamsResellerRoyaltiesResellerTypeAws),
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			AwsOptions: metronome.F(metronome.V1ContractNewParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("aws_account_number"),
				AwsOfferID:          metronome.F("aws_offer_id"),
				AwsPayerReferenceID: metronome.F("aws_payer_reference_id"),
			}),
			EndingBefore: metronome.F(time.Now()),
			GcpOptions: metronome.F(metronome.V1ContractNewParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("gcp_account_id"),
				GcpOfferID:   metronome.F("gcp_offer_id"),
			}),
			ResellerContractValue: metronome.F(0.000000),
		}}),
		SalesforceOpportunityID: metronome.F("salesforce_opportunity_id"),
		ScheduledCharges: metronome.F([]metronome.V1ContractNewParamsScheduledCharge{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Schedule: metronome.F(metronome.V1ContractNewParamsScheduledChargesSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.V1ContractNewParamsScheduledChargesScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.V1ContractNewParamsScheduledChargesScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		ScheduledChargesOnUsageInvoices: metronome.F(metronome.V1ContractNewParamsScheduledChargesOnUsageInvoicesAll),
		ThresholdBillingConfiguration: metronome.F(metronome.V1ContractNewParamsThresholdBillingConfiguration{
			CreditBalanceThresholdConfiguration: metronome.F(metronome.V1ContractNewParamsThresholdBillingConfigurationCreditBalanceThresholdConfiguration{
				Commit: metronome.F(metronome.V1ContractNewParamsThresholdBillingConfigurationCreditBalanceThresholdConfigurationCommit{
					ProductID:             metronome.F("product_id"),
					ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
					ApplicableProductTags: metronome.F([]string{"string"}),
					Description:           metronome.F("description"),
					Name:                  metronome.F("name"),
				}),
				IsEnabled:        metronome.F(true),
				RechargeToAmount: metronome.F(0.000000),
				ThresholdAmount:  metronome.F(0.000000),
			}),
			SpendThresholdConfiguration: metronome.F(metronome.V1ContractNewParamsThresholdBillingConfigurationSpendThresholdConfiguration{
				Commit: metronome.F(metronome.V1ContractNewParamsThresholdBillingConfigurationSpendThresholdConfigurationCommit{
					ProductID:             metronome.F("product_id"),
					ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
					ApplicableProductTags: metronome.F([]string{"string"}),
					Description:           metronome.F("description"),
					Name:                  metronome.F("name"),
				}),
				IsEnabled:       metronome.F(true),
				ThresholdAmount: metronome.F(0.000000),
			}),
		}),
		TotalContractValue: metronome.F(0.000000),
		Transition: metronome.F(metronome.V1ContractNewParamsTransition{
			FromContractID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Type:           metronome.F(metronome.V1ContractNewParamsTransitionTypeSupersede),
			FutureInvoiceBehavior: metronome.F(metronome.V1ContractNewParamsTransitionFutureInvoiceBehavior{
				Trueup: metronome.F(metronome.V1ContractNewParamsTransitionFutureInvoiceBehaviorTrueupRemove),
			}),
		}),
		UniquenessKey: metronome.F("x"),
		UsageFilter: metronome.F(shared.BaseUsageFilterParam{
			GroupKey:    metronome.F("group_key"),
			GroupValues: metronome.F([]string{"string"}),
			StartingAt:  metronome.F(time.Now()),
		}),
		UsageStatementSchedule: metronome.F(metronome.V1ContractNewParamsUsageStatementSchedule{
			Frequency:                   metronome.F(metronome.V1ContractNewParamsUsageStatementScheduleFrequencyMonthly),
			BillingAnchorDate:           metronome.F(time.Now()),
			Day:                         metronome.F(metronome.V1ContractNewParamsUsageStatementScheduleDayFirstOfMonth),
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

func TestV1ContractGetWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.Get(context.TODO(), metronome.V1ContractGetParams{
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

func TestV1ContractListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.List(context.TODO(), metronome.V1ContractListParams{
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

func TestV1ContractAddManualBalanceEntryWithOptionalParams(t *testing.T) {
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
	err := client.V1.Contracts.AddManualBalanceEntry(context.TODO(), metronome.V1ContractAddManualBalanceEntryParams{
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

func TestV1ContractAmendWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.Amend(context.TODO(), metronome.V1ContractAmendParams{
		ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		StartingAt: metronome.F(time.Now()),
		Commits: metronome.F([]metronome.V1ContractAmendParamsCommit{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Type:      metronome.F(metronome.V1ContractAmendParamsCommitsTypePrepaid),
			AccessSchedule: metronome.F(metronome.V1ContractAmendParamsCommitsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.V1ContractAmendParamsCommitsAccessScheduleScheduleItem{{
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
			InvoiceSchedule: metronome.F(metronome.V1ContractAmendParamsCommitsInvoiceSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.V1ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.V1ContractAmendParamsCommitsInvoiceScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			Priority:             metronome.F(0.000000),
			RateType:             metronome.F(metronome.V1ContractAmendParamsCommitsRateTypeCommitRate),
			RolloverFraction:     metronome.F(0.000000),
			TemporaryID:          metronome.F("temporary_id"),
		}}),
		Credits: metronome.F([]metronome.V1ContractAmendParamsCredit{{
			AccessSchedule: metronome.F(metronome.V1ContractAmendParamsCreditsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.V1ContractAmendParamsCreditsAccessScheduleScheduleItem{{
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
			RateType:             metronome.F(metronome.V1ContractAmendParamsCreditsRateTypeCommitRate),
		}}),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		Discounts: metronome.F([]metronome.V1ContractAmendParamsDiscount{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Schedule: metronome.F(metronome.V1ContractAmendParamsDiscountsSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.V1ContractAmendParamsDiscountsScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.V1ContractAmendParamsDiscountsScheduleScheduleItem{{
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
		Overrides: metronome.F([]metronome.V1ContractAmendParamsOverride{{
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductTags: metronome.F([]string{"string"}),
			EndingBefore:          metronome.F(time.Now()),
			Entitled:              metronome.F(true),
			IsCommitSpecific:      metronome.F(true),
			Multiplier:            metronome.F(0.000000),
			OverrideSpecifiers: metronome.F([]metronome.V1ContractAmendParamsOverridesOverrideSpecifier{{
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
			OverwriteRate: metronome.F(metronome.V1ContractAmendParamsOverridesOverwriteRate{
				RateType:     metronome.F(metronome.V1ContractAmendParamsOverridesOverwriteRateRateTypeFlat),
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
			Target:    metronome.F(metronome.V1ContractAmendParamsOverridesTargetCommitRate),
			Tiers: metronome.F([]metronome.V1ContractAmendParamsOverridesTier{{
				Multiplier: metronome.F(0.000000),
				Size:       metronome.F(0.000000),
			}}),
			Type: metronome.F(metronome.V1ContractAmendParamsOverridesTypeOverwrite),
		}}),
		ProfessionalServices: metronome.F([]metronome.V1ContractAmendParamsProfessionalService{{
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
		ResellerRoyalties: metronome.F([]metronome.V1ContractAmendParamsResellerRoyalty{{
			ResellerType:          metronome.F(metronome.V1ContractAmendParamsResellerRoyaltiesResellerTypeAws),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			AwsOptions: metronome.F(metronome.V1ContractAmendParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("aws_account_number"),
				AwsOfferID:          metronome.F("aws_offer_id"),
				AwsPayerReferenceID: metronome.F("aws_payer_reference_id"),
			}),
			EndingBefore: metronome.F(time.Now()),
			Fraction:     metronome.F(0.000000),
			GcpOptions: metronome.F(metronome.V1ContractAmendParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("gcp_account_id"),
				GcpOfferID:   metronome.F("gcp_offer_id"),
			}),
			NetsuiteResellerID:    metronome.F("netsuite_reseller_id"),
			ResellerContractValue: metronome.F(0.000000),
			StartingAt:            metronome.F(time.Now()),
		}}),
		SalesforceOpportunityID: metronome.F("salesforce_opportunity_id"),
		ScheduledCharges: metronome.F([]metronome.V1ContractAmendParamsScheduledCharge{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Schedule: metronome.F(metronome.V1ContractAmendParamsScheduledChargesSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.V1ContractAmendParamsScheduledChargesScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.V1ContractAmendParamsScheduledChargesScheduleScheduleItem{{
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

func TestV1ContractArchive(t *testing.T) {
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
	_, err := client.V1.Contracts.Archive(context.TODO(), metronome.V1ContractArchiveParams{
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

func TestV1ContractNewHistoricalInvoices(t *testing.T) {
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
	_, err := client.V1.Contracts.NewHistoricalInvoices(context.TODO(), metronome.V1ContractNewHistoricalInvoicesParams{
		Invoices: metronome.F([]metronome.V1ContractNewHistoricalInvoicesParamsInvoice{{
			ContractID:         metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
			CreditTypeID:       metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			CustomerID:         metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
			ExclusiveEndDate:   metronome.F(time.Now()),
			InclusiveStartDate: metronome.F(time.Now()),
			IssueDate:          metronome.F(time.Now()),
			UsageLineItems: metronome.F([]metronome.V1ContractNewHistoricalInvoicesParamsInvoicesUsageLineItem{{
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
				SubtotalsWithQuantity: metronome.F([]metronome.V1ContractNewHistoricalInvoicesParamsInvoicesUsageLineItemsSubtotalsWithQuantity{{
					ExclusiveEndDate:   metronome.F(time.Now()),
					InclusiveStartDate: metronome.F(time.Now()),
					Quantity:           metronome.F(0.000000),
				}}),
			}}),
			BillableStatus:       metronome.F(metronome.V1ContractNewHistoricalInvoicesParamsInvoicesBillableStatusBillable),
			BreakdownGranularity: metronome.F(metronome.V1ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularityHour),
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

func TestV1ContractListBalancesWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.ListBalances(context.TODO(), metronome.V1ContractListBalancesParams{
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

func TestV1ContractGetRateScheduleWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.GetRateSchedule(context.TODO(), metronome.V1ContractGetRateScheduleParams{
		ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		Limit:      metronome.F(int64(1)),
		NextPage:   metronome.F("next_page"),
		At:         metronome.F(time.Now()),
		Selectors: metronome.F([]metronome.V1ContractGetRateScheduleParamsSelector{{
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

func TestV1ContractScheduleProServicesInvoiceWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.ScheduleProServicesInvoice(context.TODO(), metronome.V1ContractScheduleProServicesInvoiceParams{
		ContractID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CustomerID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IssuedAt:   metronome.F(time.Now()),
		LineItems: metronome.F([]metronome.V1ContractScheduleProServicesInvoiceParamsLineItem{{
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

func TestV1ContractSetUsageFilter(t *testing.T) {
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
	err := client.V1.Contracts.SetUsageFilter(context.TODO(), metronome.V1ContractSetUsageFilterParams{
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

func TestV1ContractUpdateEndDateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.UpdateEndDate(context.TODO(), metronome.V1ContractUpdateEndDateParams{
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
