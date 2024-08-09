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
			BillingProvider: metronome.F(metronome.ContractNewParamsBillingProviderConfigurationBillingProviderAwsMarketplace),
			DeliveryMethod:  metronome.F(metronome.ContractNewParamsBillingProviderConfigurationDeliveryMethodDirectToBillingProvider),
		}),
		Commits: metronome.F([]metronome.ContractNewParamsCommit{{
			Type:      metronome.F(metronome.ContractNewParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractNewParamsCommitsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			InvoiceSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:                metronome.F(0.000000),
			Description:           metronome.F("description"),
			RolloverFraction:      metronome.F(0.000000),
			Priority:              metronome.F(0.000000),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Type:      metronome.F(metronome.ContractNewParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractNewParamsCommitsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			InvoiceSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:                metronome.F(0.000000),
			Description:           metronome.F("description"),
			RolloverFraction:      metronome.F(0.000000),
			Priority:              metronome.F(0.000000),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Type:      metronome.F(metronome.ContractNewParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractNewParamsCommitsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			InvoiceSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:                metronome.F(0.000000),
			Description:           metronome.F("description"),
			RolloverFraction:      metronome.F(0.000000),
			Priority:              metronome.F(0.000000),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}}),
		Credits: metronome.F([]metronome.ContractNewParamsCredit{{
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractNewParamsCreditsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCreditsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			Description:           metronome.F("description"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			Priority:              metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractNewParamsCreditsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCreditsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			Description:           metronome.F("description"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			Priority:              metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractNewParamsCreditsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsCreditsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			Description:           metronome.F("description"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			Priority:              metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}}),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		Discounts: metronome.F([]metronome.ContractNewParamsDiscount{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractNewParamsDiscountsSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsDiscountsScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractNewParamsDiscountsSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsDiscountsScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractNewParamsDiscountsSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsDiscountsScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		EndingBefore:                     metronome.F(time.Now()),
		MultiplierOverridePrioritization: metronome.F(metronome.ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier),
		Name:                             metronome.F("name"),
		NetPaymentTermsDays:              metronome.F(0.000000),
		NetsuiteSalesOrderID:             metronome.F("netsuite_sales_order_id"),
		Overrides: metronome.F([]metronome.ContractNewParamsOverride{{
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractNewParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			Priority:     metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractNewParamsOverridesOverwriteRate{
				RateType:   metronome.F(metronome.ContractNewParamsOverridesOverwriteRateRateTypeFlat),
				Price:      metronome.F(0.000000),
				Quantity:   metronome.F(0.000000),
				IsProrated: metronome.F(true),
				Tiers: metronome.F([]metronome.ContractNewParamsOverridesOverwriteRateTier{{
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}}),
				CustomRate: metronome.F(map[string]interface{}{
					"foo": "bar",
				}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			OverrideSpecifiers: metronome.F([]metronome.ContractNewParamsOverridesOverrideSpecifier{{
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}}),
		}, {
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractNewParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			Priority:     metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractNewParamsOverridesOverwriteRate{
				RateType:   metronome.F(metronome.ContractNewParamsOverridesOverwriteRateRateTypeFlat),
				Price:      metronome.F(0.000000),
				Quantity:   metronome.F(0.000000),
				IsProrated: metronome.F(true),
				Tiers: metronome.F([]metronome.ContractNewParamsOverridesOverwriteRateTier{{
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}}),
				CustomRate: metronome.F(map[string]interface{}{
					"foo": "bar",
				}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			OverrideSpecifiers: metronome.F([]metronome.ContractNewParamsOverridesOverrideSpecifier{{
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}}),
		}, {
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractNewParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			Priority:     metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractNewParamsOverridesOverwriteRate{
				RateType:   metronome.F(metronome.ContractNewParamsOverridesOverwriteRateRateTypeFlat),
				Price:      metronome.F(0.000000),
				Quantity:   metronome.F(0.000000),
				IsProrated: metronome.F(true),
				Tiers: metronome.F([]metronome.ContractNewParamsOverridesOverwriteRateTier{{
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}}),
				CustomRate: metronome.F(map[string]interface{}{
					"foo": "bar",
				}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			OverrideSpecifiers: metronome.F([]metronome.ContractNewParamsOverridesOverrideSpecifier{{
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}}),
		}}),
		ProfessionalServices: metronome.F([]metronome.ContractNewParamsProfessionalService{{
			Description:          metronome.F("description"),
			ProductID:            metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			UnitPrice:            metronome.F(0.000000),
			Quantity:             metronome.F(0.000000),
			MaxAmount:            metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Description:          metronome.F("description"),
			ProductID:            metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			UnitPrice:            metronome.F(0.000000),
			Quantity:             metronome.F(0.000000),
			MaxAmount:            metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Description:          metronome.F("description"),
			ProductID:            metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			UnitPrice:            metronome.F(0.000000),
			Quantity:             metronome.F(0.000000),
			MaxAmount:            metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}}),
		RateCardAlias: metronome.F("rate_card_alias"),
		RateCardID:    metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		ResellerRoyalties: metronome.F([]metronome.ContractNewParamsResellerRoyalty{{
			ResellerType:          metronome.F(metronome.ContractNewParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("netsuite_reseller_id"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("aws_account_number"),
				AwsPayerReferenceID: metronome.F("aws_payer_reference_id"),
				AwsOfferID:          metronome.F("aws_offer_id"),
			}),
			GcpOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("gcp_account_id"),
				GcpOfferID:   metronome.F("gcp_offer_id"),
			}),
		}, {
			ResellerType:          metronome.F(metronome.ContractNewParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("netsuite_reseller_id"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("aws_account_number"),
				AwsPayerReferenceID: metronome.F("aws_payer_reference_id"),
				AwsOfferID:          metronome.F("aws_offer_id"),
			}),
			GcpOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("gcp_account_id"),
				GcpOfferID:   metronome.F("gcp_offer_id"),
			}),
		}, {
			ResellerType:          metronome.F(metronome.ContractNewParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("netsuite_reseller_id"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("aws_account_number"),
				AwsPayerReferenceID: metronome.F("aws_payer_reference_id"),
				AwsOfferID:          metronome.F("aws_offer_id"),
			}),
			GcpOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("gcp_account_id"),
				GcpOfferID:   metronome.F("gcp_offer_id"),
			}),
		}}),
		SalesforceOpportunityID: metronome.F("salesforce_opportunity_id"),
		ScheduledCharges: metronome.F([]metronome.ContractNewParamsScheduledCharge{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractNewParamsScheduledChargesSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractNewParamsScheduledChargesSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractNewParamsScheduledChargesSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		TotalContractValue: metronome.F(0.000000),
		Transition: metronome.F(metronome.ContractNewParamsTransition{
			Type:           metronome.F(metronome.ContractNewParamsTransitionTypeSupersede),
			FromContractID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			FutureInvoiceBehavior: metronome.F(metronome.ContractNewParamsTransitionFutureInvoiceBehavior{
				Trueup: metronome.F(metronome.ContractNewParamsTransitionFutureInvoiceBehaviorTrueupRemove),
			}),
		}),
		UniquenessKey: metronome.F("x"),
		UsageFilter: metronome.F(metronome.ContractNewParamsUsageFilter{
			GroupKey:    metronome.F("group_key"),
			GroupValues: metronome.F([]string{"string", "string", "string"}),
			StartingAt:  metronome.F(time.Now()),
		}),
		UsageStatementSchedule: metronome.F(metronome.ContractNewParamsUsageStatementSchedule{
			Frequency:                   metronome.F(metronome.ContractNewParamsUsageStatementScheduleFrequencyMonthly),
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
			Type:      metronome.F(metronome.ContractAmendParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractAmendParamsCommitsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			InvoiceSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:                metronome.F(0.000000),
			Description:           metronome.F("description"),
			RolloverFraction:      metronome.F(0.000000),
			Priority:              metronome.F(0.000000),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Type:      metronome.F(metronome.ContractAmendParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractAmendParamsCommitsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			InvoiceSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:                metronome.F(0.000000),
			Description:           metronome.F("description"),
			RolloverFraction:      metronome.F(0.000000),
			Priority:              metronome.F(0.000000),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Type:      metronome.F(metronome.ContractAmendParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractAmendParamsCommitsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			InvoiceSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:                metronome.F(0.000000),
			Description:           metronome.F("description"),
			RolloverFraction:      metronome.F(0.000000),
			Priority:              metronome.F(0.000000),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}}),
		Credits: metronome.F([]metronome.ContractAmendParamsCredit{{
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractAmendParamsCreditsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCreditsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			Description:           metronome.F("description"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			Priority:              metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractAmendParamsCreditsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCreditsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			Description:           metronome.F("description"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			Priority:              metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractAmendParamsCreditsAccessSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsCreditsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			Description:           metronome.F("description"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			Priority:              metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}}),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		Discounts: metronome.F([]metronome.ContractAmendParamsDiscount{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendParamsDiscountsSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsDiscountsScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendParamsDiscountsSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsDiscountsScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendParamsDiscountsSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsDiscountsScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		Overrides: metronome.F([]metronome.ContractAmendParamsOverride{{
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractAmendParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			Priority:     metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractAmendParamsOverridesOverwriteRate{
				RateType:   metronome.F(metronome.ContractAmendParamsOverridesOverwriteRateRateTypeFlat),
				Price:      metronome.F(0.000000),
				Quantity:   metronome.F(0.000000),
				IsProrated: metronome.F(true),
				Tiers: metronome.F([]metronome.ContractAmendParamsOverridesOverwriteRateTier{{
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}}),
				CustomRate: metronome.F(map[string]interface{}{
					"foo": "bar",
				}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			OverrideSpecifiers: metronome.F([]metronome.ContractAmendParamsOverridesOverrideSpecifier{{
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}}),
		}, {
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractAmendParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			Priority:     metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractAmendParamsOverridesOverwriteRate{
				RateType:   metronome.F(metronome.ContractAmendParamsOverridesOverwriteRateRateTypeFlat),
				Price:      metronome.F(0.000000),
				Quantity:   metronome.F(0.000000),
				IsProrated: metronome.F(true),
				Tiers: metronome.F([]metronome.ContractAmendParamsOverridesOverwriteRateTier{{
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}}),
				CustomRate: metronome.F(map[string]interface{}{
					"foo": "bar",
				}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			OverrideSpecifiers: metronome.F([]metronome.ContractAmendParamsOverridesOverrideSpecifier{{
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}}),
		}, {
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractAmendParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			Priority:     metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractAmendParamsOverridesOverwriteRate{
				RateType:   metronome.F(metronome.ContractAmendParamsOverridesOverwriteRateRateTypeFlat),
				Price:      metronome.F(0.000000),
				Quantity:   metronome.F(0.000000),
				IsProrated: metronome.F(true),
				Tiers: metronome.F([]metronome.ContractAmendParamsOverridesOverwriteRateTier{{
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}, {
					Size:  metronome.F(0.000000),
					Price: metronome.F(0.000000),
				}}),
				CustomRate: metronome.F(map[string]interface{}{
					"foo": "bar",
				}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			OverrideSpecifiers: metronome.F([]metronome.ContractAmendParamsOverridesOverrideSpecifier{{
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}, {
				ProductID:   metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags: metronome.F([]string{"string", "string", "string"}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
			}}),
		}}),
		ProfessionalServices: metronome.F([]metronome.ContractAmendParamsProfessionalService{{
			Description:          metronome.F("description"),
			ProductID:            metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			UnitPrice:            metronome.F(0.000000),
			Quantity:             metronome.F(0.000000),
			MaxAmount:            metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Description:          metronome.F("description"),
			ProductID:            metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			UnitPrice:            metronome.F(0.000000),
			Quantity:             metronome.F(0.000000),
			MaxAmount:            metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}, {
			Description:          metronome.F("description"),
			ProductID:            metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			UnitPrice:            metronome.F(0.000000),
			Quantity:             metronome.F(0.000000),
			MaxAmount:            metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
		}}),
		ResellerRoyalties: metronome.F([]metronome.ContractAmendParamsResellerRoyalty{{
			ResellerType:          metronome.F(metronome.ContractAmendParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("netsuite_reseller_id"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("aws_account_number"),
				AwsPayerReferenceID: metronome.F("aws_payer_reference_id"),
				AwsOfferID:          metronome.F("aws_offer_id"),
			}),
			GcpOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("gcp_account_id"),
				GcpOfferID:   metronome.F("gcp_offer_id"),
			}),
		}, {
			ResellerType:          metronome.F(metronome.ContractAmendParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("netsuite_reseller_id"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("aws_account_number"),
				AwsPayerReferenceID: metronome.F("aws_payer_reference_id"),
				AwsOfferID:          metronome.F("aws_offer_id"),
			}),
			GcpOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("gcp_account_id"),
				GcpOfferID:   metronome.F("gcp_offer_id"),
			}),
		}, {
			ResellerType:          metronome.F(metronome.ContractAmendParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("netsuite_reseller_id"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string", "string", "string"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("aws_account_number"),
				AwsPayerReferenceID: metronome.F("aws_payer_reference_id"),
				AwsOfferID:          metronome.F("aws_offer_id"),
			}),
			GcpOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("gcp_account_id"),
				GcpOfferID:   metronome.F("gcp_offer_id"),
			}),
		}}),
		SalesforceOpportunityID: metronome.F("salesforce_opportunity_id"),
		ScheduledCharges: metronome.F([]metronome.ContractAmendParamsScheduledCharge{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendParamsScheduledChargesSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendParamsScheduledChargesSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendParamsScheduledChargesSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ScheduleItems: metronome.F([]metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
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
			ProductID:   metronome.F("d6300dbb-882e-4d2d-8dec-5125d16b65d0"),
			ProductTags: metronome.F([]string{"string", "string", "string"}),
			PricingGroupValues: metronome.F(map[string]string{
				"foo": "string",
			}),
			PartialPricingGroupValues: metronome.F(map[string]string{
				"region": "us-west-2",
				"cloud":  "aws",
			}),
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
			UnitPrice:                   metronome.F(0.000000),
			Quantity:                    metronome.F(0.000000),
			Amount:                      metronome.F(0.000000),
			NetsuiteInvoiceBillingStart: metronome.F(time.Now()),
			NetsuiteInvoiceBillingEnd:   metronome.F(time.Now()),
			Metadata:                    metronome.F("metadata"),
		}, {
			ProfessionalServiceID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AmendmentID:                 metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			UnitPrice:                   metronome.F(0.000000),
			Quantity:                    metronome.F(0.000000),
			Amount:                      metronome.F(0.000000),
			NetsuiteInvoiceBillingStart: metronome.F(time.Now()),
			NetsuiteInvoiceBillingEnd:   metronome.F(time.Now()),
			Metadata:                    metronome.F("metadata"),
		}, {
			ProfessionalServiceID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AmendmentID:                 metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			UnitPrice:                   metronome.F(0.000000),
			Quantity:                    metronome.F(0.000000),
			Amount:                      metronome.F(0.000000),
			NetsuiteInvoiceBillingStart: metronome.F(time.Now()),
			NetsuiteInvoiceBillingEnd:   metronome.F(time.Now()),
			Metadata:                    metronome.F("metadata"),
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
		ContractID:   metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID:   metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		EndingBefore: metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
