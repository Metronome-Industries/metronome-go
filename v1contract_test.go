// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Metronome-Industries/metronome-go/v2"
	"github.com/Metronome-Industries/metronome-go/v2/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/shared"
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
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		StartingAt: time.Now(),
		BillingProviderConfiguration: metronome.V1ContractNewParamsBillingProviderConfiguration{
			BillingProvider:                "stripe",
			BillingProviderConfigurationID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DeliveryMethod:                 "direct_to_billing_provider",
		},
		Commits: []metronome.V1ContractNewParamsCommit{{
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Type:      "PREPAID",
			AccessSchedule: metronome.V1ContractNewParamsCommitAccessSchedule{
				ScheduleItems: []metronome.V1ContractNewParamsCommitAccessScheduleScheduleItem{{
					Amount:       0,
					EndingBefore: time.Now(),
					StartingAt:   time.Now(),
				}},
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			Amount:                metronome.Float(0),
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Description: metronome.String("description"),
			HierarchyConfiguration: shared.CommitHierarchyConfigurationParam{
				ChildAccess: shared.CommitHierarchyConfigurationChildAccessUnionParam{
					OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll: &shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam{
						Type: "ALL",
					},
				},
			},
			InvoiceSchedule: metronome.V1ContractNewParamsCommitInvoiceSchedule{
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				DoNotInvoice: metronome.Bool(true),
				RecurringSchedule: metronome.V1ContractNewParamsCommitInvoiceScheduleRecurringSchedule{
					AmountDistribution: "DIVIDED",
					EndingBefore:       time.Now(),
					Frequency:          "MONTHLY",
					StartingAt:         time.Now(),
					Amount:             metronome.Float(0),
					Quantity:           metronome.Float(0),
					UnitPrice:          metronome.Float(0),
				},
				ScheduleItems: []metronome.V1ContractNewParamsCommitInvoiceScheduleScheduleItem{{
					Timestamp: time.Now(),
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					UnitPrice: metronome.Float(0),
				}},
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
			PaymentGateConfig: metronome.V1ContractNewParamsCommitPaymentGateConfig{
				PaymentGateType: "NONE",
				PrecalculatedTaxConfig: metronome.V1ContractNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig{
					TaxAmount: 0,
					TaxName:   metronome.String("tax_name"),
				},
				StripeConfig: metronome.V1ContractNewParamsCommitPaymentGateConfigStripeConfig{
					PaymentType: "INVOICE",
					InvoiceMetadata: map[string]string{
						"foo": "string",
					},
					OnSessionPayment: metronome.Bool(true),
				},
				TaxType: "NONE",
			},
			Priority:         metronome.Float(0),
			RateType:         "COMMIT_RATE",
			RolloverFraction: metronome.Float(0),
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
			TemporaryID: metronome.String("temporary_id"),
		}},
		Credits: []metronome.V1ContractNewParamsCredit{{
			AccessSchedule: metronome.V1ContractNewParamsCreditAccessSchedule{
				ScheduleItems: []metronome.V1ContractNewParamsCreditAccessScheduleScheduleItem{{
					Amount:       0,
					EndingBefore: time.Now(),
					StartingAt:   time.Now(),
				}},
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			ProductID:             "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Description: metronome.String("description"),
			HierarchyConfiguration: shared.CommitHierarchyConfigurationParam{
				ChildAccess: shared.CommitHierarchyConfigurationChildAccessUnionParam{
					OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll: &shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam{
						Type: "ALL",
					},
				},
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
			Priority:             metronome.Float(0),
			RateType:             "COMMIT_RATE",
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
		}},
		CustomFields: map[string]string{
			"foo": "string",
		},
		Discounts: []metronome.V1ContractNewParamsDiscount{{
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Schedule: metronome.V1ContractNewParamsDiscountSchedule{
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				DoNotInvoice: metronome.Bool(true),
				RecurringSchedule: metronome.V1ContractNewParamsDiscountScheduleRecurringSchedule{
					AmountDistribution: "DIVIDED",
					EndingBefore:       time.Now(),
					Frequency:          "MONTHLY",
					StartingAt:         time.Now(),
					Amount:             metronome.Float(0),
					Quantity:           metronome.Float(0),
					UnitPrice:          metronome.Float(0),
				},
				ScheduleItems: []metronome.V1ContractNewParamsDiscountScheduleScheduleItem{{
					Timestamp: time.Now(),
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					UnitPrice: metronome.Float(0),
				}},
			},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
		}},
		EndingBefore: metronome.Time(time.Now()),
		HierarchyConfiguration: metronome.V1ContractNewParamsHierarchyConfiguration{
			Parent: metronome.V1ContractNewParamsHierarchyConfigurationParent{
				ContractID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				CustomerID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			},
			ParentBehavior: metronome.V1ContractNewParamsHierarchyConfigurationParentBehavior{
				InvoiceConsolidationType: "CONCATENATE",
			},
			Payer:                  "SELF",
			UsageStatementBehavior: "CONSOLIDATE",
		},
		MultiplierOverridePrioritization: metronome.V1ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier,
		Name:                             metronome.String("name"),
		NetPaymentTermsDays:              metronome.Float(0),
		NetsuiteSalesOrderID:             metronome.String("netsuite_sales_order_id"),
		ContractOverrides: []metronome.V1ContractNewParamsOverride{{
			StartingAt:            time.Now(),
			ApplicableProductTags: []string{"string"},
			EndingBefore:          metronome.Time(time.Now()),
			Entitled:              metronome.Bool(true),
			IsCommitSpecific:      metronome.Bool(true),
			Multiplier:            metronome.Float(0),
			OverrideSpecifiers: []metronome.V1ContractNewParamsOverrideOverrideSpecifier{{
				BillingFrequency: "MONTHLY",
				CommitIDs:        []string{"string"},
				PresentationGroupValues: map[string]string{
					"foo": "string",
				},
				PricingGroupValues: map[string]string{
					"foo": "string",
				},
				ProductID:          metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags:        []string{"string"},
				RecurringCommitIDs: []string{"string"},
				RecurringCreditIDs: []string{"string"},
			}},
			OverwriteRate: metronome.V1ContractNewParamsOverrideOverwriteRate{
				RateType:     "FLAT",
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				CustomRate: map[string]any{
					"foo": "bar",
				},
				IsProrated: metronome.Bool(true),
				Price:      metronome.Float(0),
				Quantity:   metronome.Float(0),
				Tiers: []shared.TierParam{{
					Price: 0,
					Size:  metronome.Float(0),
				}},
			},
			Priority:  metronome.Float(0),
			ProductID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Target:    "COMMIT_RATE",
			Tiers: []metronome.V1ContractNewParamsOverrideTier{{
				Multiplier: 0,
				Size:       metronome.Float(0),
			}},
			Type: "OVERWRITE",
		}},
		PrepaidBalanceThresholdConfiguration: shared.PrepaidBalanceThresholdConfigurationParam{
			Commit: shared.PrepaidBalanceThresholdConfigurationCommitParam{
				BaseThresholdCommitParam: shared.BaseThresholdCommitParam{
					ProductID:   "product_id",
					Description: metronome.String("description"),
					Name:        metronome.String("name"),
				},
				ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
				ApplicableProductTags: []string{"string"},
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
			},
			IsEnabled: true,
			PaymentGateConfig: shared.PaymentGateConfigParam{
				PaymentGateType: shared.PaymentGateConfigPaymentGateTypeNone,
				PrecalculatedTaxConfig: shared.PaymentGateConfigPrecalculatedTaxConfigParam{
					TaxAmount: 0,
					TaxName:   metronome.String("tax_name"),
				},
				StripeConfig: shared.PaymentGateConfigStripeConfigParam{
					PaymentType: "INVOICE",
					InvoiceMetadata: map[string]string{
						"foo": "string",
					},
				},
				TaxType: shared.PaymentGateConfigTaxTypeNone,
			},
			RechargeToAmount:   0,
			ThresholdAmount:    0,
			CustomCreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
		Priority: metronome.Float(0),
		ProfessionalServices: []metronome.V1ContractNewParamsProfessionalService{{
			MaxAmount: 0,
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Quantity:  0,
			UnitPrice: 0,
			CustomFields: map[string]string{
				"foo": "string",
			},
			Description:          metronome.String("description"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
		}},
		RateCardAlias: metronome.String("rate_card_alias"),
		RateCardID:    metronome.String("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		RecurringCommits: []metronome.V1ContractNewParamsRecurringCommit{{
			AccessAmount: metronome.V1ContractNewParamsRecurringCommitAccessAmount{
				CreditTypeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				UnitPrice:    0,
				Quantity:     metronome.Float(0),
			},
			CommitDuration: metronome.V1ContractNewParamsRecurringCommitCommitDuration{
				Value: 0,
				Unit:  "PERIODS",
			},
			Priority:              0,
			ProductID:             "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			StartingAt:            time.Now(),
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			Description:           metronome.String("description"),
			EndingBefore:          metronome.Time(time.Now()),
			HierarchyConfiguration: shared.CommitHierarchyConfigurationParam{
				ChildAccess: shared.CommitHierarchyConfigurationChildAccessUnionParam{
					OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll: &shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam{
						Type: "ALL",
					},
				},
			},
			InvoiceAmount: metronome.V1ContractNewParamsRecurringCommitInvoiceAmount{
				CreditTypeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Quantity:     0,
				UnitPrice:    0,
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
			Proration:            "NONE",
			RateType:             "COMMIT_RATE",
			RecurrenceFrequency:  "MONTHLY",
			RolloverFraction:     metronome.Float(0),
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
			SubscriptionConfig: metronome.V1ContractNewParamsRecurringCommitSubscriptionConfig{
				ApplySeatIncreaseConfig: metronome.V1ContractNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig{
					IsProrated: true,
				},
				SubscriptionID: "subscription_id",
				Allocation:     "INDIVIDUAL",
			},
			TemporaryID: metronome.String("temporary_id"),
		}},
		RecurringCredits: []metronome.V1ContractNewParamsRecurringCredit{{
			AccessAmount: metronome.V1ContractNewParamsRecurringCreditAccessAmount{
				CreditTypeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				UnitPrice:    0,
				Quantity:     metronome.Float(0),
			},
			CommitDuration: metronome.V1ContractNewParamsRecurringCreditCommitDuration{
				Value: 0,
				Unit:  "PERIODS",
			},
			Priority:              0,
			ProductID:             "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			StartingAt:            time.Now(),
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			Description:           metronome.String("description"),
			EndingBefore:          metronome.Time(time.Now()),
			HierarchyConfiguration: shared.CommitHierarchyConfigurationParam{
				ChildAccess: shared.CommitHierarchyConfigurationChildAccessUnionParam{
					OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll: &shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam{
						Type: "ALL",
					},
				},
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
			Proration:            "NONE",
			RateType:             "COMMIT_RATE",
			RecurrenceFrequency:  "MONTHLY",
			RolloverFraction:     metronome.Float(0),
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
			SubscriptionConfig: metronome.V1ContractNewParamsRecurringCreditSubscriptionConfig{
				ApplySeatIncreaseConfig: metronome.V1ContractNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig{
					IsProrated: true,
				},
				SubscriptionID: "subscription_id",
				Allocation:     "INDIVIDUAL",
			},
			TemporaryID: metronome.String("temporary_id"),
		}},
		ResellerRoyalties: []metronome.V1ContractNewParamsResellerRoyalty{{
			Fraction:              0,
			NetsuiteResellerID:    "netsuite_reseller_id",
			ResellerType:          "AWS",
			StartingAt:            time.Now(),
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			AwsOptions: metronome.V1ContractNewParamsResellerRoyaltyAwsOptions{
				AwsAccountNumber:    metronome.String("aws_account_number"),
				AwsOfferID:          metronome.String("aws_offer_id"),
				AwsPayerReferenceID: metronome.String("aws_payer_reference_id"),
			},
			EndingBefore: metronome.Time(time.Now()),
			GcpOptions: metronome.V1ContractNewParamsResellerRoyaltyGcpOptions{
				GcpAccountID: metronome.String("gcp_account_id"),
				GcpOfferID:   metronome.String("gcp_offer_id"),
			},
			ResellerContractValue: metronome.Float(0),
		}},
		RevenueSystemConfiguration: metronome.V1ContractNewParamsRevenueSystemConfiguration{
			DeliveryMethod:               "direct_to_billing_provider",
			Provider:                     "netsuite",
			RevenueSystemConfigurationID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
		SalesforceOpportunityID: metronome.String("salesforce_opportunity_id"),
		ScheduledCharges: []metronome.V1ContractNewParamsScheduledCharge{{
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Schedule: metronome.V1ContractNewParamsScheduledChargeSchedule{
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				DoNotInvoice: metronome.Bool(true),
				RecurringSchedule: metronome.V1ContractNewParamsScheduledChargeScheduleRecurringSchedule{
					AmountDistribution: "DIVIDED",
					EndingBefore:       time.Now(),
					Frequency:          "MONTHLY",
					StartingAt:         time.Now(),
					Amount:             metronome.Float(0),
					Quantity:           metronome.Float(0),
					UnitPrice:          metronome.Float(0),
				},
				ScheduleItems: []metronome.V1ContractNewParamsScheduledChargeScheduleScheduleItem{{
					Timestamp: time.Now(),
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					UnitPrice: metronome.Float(0),
				}},
			},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
		}},
		ScheduledChargesOnUsageInvoices: metronome.V1ContractNewParamsScheduledChargesOnUsageInvoicesAll,
		SpendThresholdConfiguration: shared.SpendThresholdConfigurationParam{
			Commit: shared.BaseThresholdCommitParam{
				ProductID:   "product_id",
				Description: metronome.String("description"),
				Name:        metronome.String("name"),
			},
			IsEnabled: true,
			PaymentGateConfig: shared.PaymentGateConfigParam{
				PaymentGateType: shared.PaymentGateConfigPaymentGateTypeNone,
				PrecalculatedTaxConfig: shared.PaymentGateConfigPrecalculatedTaxConfigParam{
					TaxAmount: 0,
					TaxName:   metronome.String("tax_name"),
				},
				StripeConfig: shared.PaymentGateConfigStripeConfigParam{
					PaymentType: "INVOICE",
					InvoiceMetadata: map[string]string{
						"foo": "string",
					},
				},
				TaxType: shared.PaymentGateConfigTaxTypeNone,
			},
			ThresholdAmount: 0,
		},
		Subscriptions: []metronome.V1ContractNewParamsSubscription{{
			CollectionSchedule: "ADVANCE",
			Proration: metronome.V1ContractNewParamsSubscriptionProration{
				InvoiceBehavior: "BILL_IMMEDIATELY",
				IsProrated:      metronome.Bool(true),
			},
			SubscriptionRate: metronome.V1ContractNewParamsSubscriptionSubscriptionRate{
				BillingFrequency: "MONTHLY",
				ProductID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Description:            metronome.String("description"),
			EndingBefore:           metronome.Time(time.Now()),
			InitialQuantity:        metronome.Float(0),
			Name:                   metronome.String("name"),
			QuantityManagementMode: "SEAT_BASED",
			SeatConfig: metronome.V1ContractNewParamsSubscriptionSeatConfig{
				InitialSeatIDs:         []string{"string"},
				SeatGroupKey:           "seat_group_key",
				InitialUnassignedSeats: metronome.Float(0),
			},
			StartingAt:  metronome.Time(time.Now()),
			TemporaryID: metronome.String("temporary_id"),
		}},
		TotalContractValue: metronome.Float(0),
		Transition: metronome.V1ContractNewParamsTransition{
			FromContractID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Type:           "SUPERSEDE",
			FutureInvoiceBehavior: metronome.V1ContractNewParamsTransitionFutureInvoiceBehavior{
				Trueup: "REMOVE",
			},
		},
		UniquenessKey: metronome.String("x"),
		UsageFilter: shared.BaseUsageFilterParam{
			GroupKey:    "group_key",
			GroupValues: []string{"string"},
			StartingAt:  metronome.Time(time.Now()),
		},
		UsageStatementSchedule: metronome.V1ContractNewParamsUsageStatementSchedule{
			Frequency:                   "MONTHLY",
			BillingAnchorDate:           metronome.Time(time.Now()),
			Day:                         "FIRST_OF_MONTH",
			InvoiceGenerationStartingAt: metronome.Time(time.Now()),
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
		ContractID:     "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:     "13117714-3f05-48e5-a6e9-a66093f13b4d",
		IncludeBalance: metronome.Bool(true),
		IncludeLedgers: metronome.Bool(true),
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
		CustomerID:      "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		CoveringDate:    metronome.Time(time.Now()),
		IncludeArchived: metronome.Bool(true),
		IncludeBalance:  metronome.Bool(true),
		IncludeLedgers:  metronome.Bool(true),
		StartingAt:      metronome.Time(time.Now()),
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
		ID:         "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
		Amount:     -1000,
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		Reason:     "Reason for entry",
		SegmentID:  "66368e29-3f97-4d15-a6e9-120897f0070a",
		ContractID: metronome.String("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		PerGroupAmounts: map[string]float64{
			"foo": 0,
		},
		Timestamp: metronome.Time(time.Now()),
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
		ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		StartingAt: time.Now(),
		Commits: []metronome.V1ContractAmendParamsCommit{{
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Type:      "PREPAID",
			AccessSchedule: metronome.V1ContractAmendParamsCommitAccessSchedule{
				ScheduleItems: []metronome.V1ContractAmendParamsCommitAccessScheduleScheduleItem{{
					Amount:       0,
					EndingBefore: time.Now(),
					StartingAt:   time.Now(),
				}},
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			Amount:                metronome.Float(0),
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Description: metronome.String("description"),
			HierarchyConfiguration: shared.CommitHierarchyConfigurationParam{
				ChildAccess: shared.CommitHierarchyConfigurationChildAccessUnionParam{
					OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll: &shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam{
						Type: "ALL",
					},
				},
			},
			InvoiceSchedule: metronome.V1ContractAmendParamsCommitInvoiceSchedule{
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				DoNotInvoice: metronome.Bool(true),
				RecurringSchedule: metronome.V1ContractAmendParamsCommitInvoiceScheduleRecurringSchedule{
					AmountDistribution: "DIVIDED",
					EndingBefore:       time.Now(),
					Frequency:          "MONTHLY",
					StartingAt:         time.Now(),
					Amount:             metronome.Float(0),
					Quantity:           metronome.Float(0),
					UnitPrice:          metronome.Float(0),
				},
				ScheduleItems: []metronome.V1ContractAmendParamsCommitInvoiceScheduleScheduleItem{{
					Timestamp: time.Now(),
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					UnitPrice: metronome.Float(0),
				}},
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
			PaymentGateConfig: metronome.V1ContractAmendParamsCommitPaymentGateConfig{
				PaymentGateType: "NONE",
				PrecalculatedTaxConfig: metronome.V1ContractAmendParamsCommitPaymentGateConfigPrecalculatedTaxConfig{
					TaxAmount: 0,
					TaxName:   metronome.String("tax_name"),
				},
				StripeConfig: metronome.V1ContractAmendParamsCommitPaymentGateConfigStripeConfig{
					PaymentType: "INVOICE",
					InvoiceMetadata: map[string]string{
						"foo": "string",
					},
					OnSessionPayment: metronome.Bool(true),
				},
				TaxType: "NONE",
			},
			Priority:         metronome.Float(0),
			RateType:         "COMMIT_RATE",
			RolloverFraction: metronome.Float(0),
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
			TemporaryID: metronome.String("temporary_id"),
		}},
		Credits: []metronome.V1ContractAmendParamsCredit{{
			AccessSchedule: metronome.V1ContractAmendParamsCreditAccessSchedule{
				ScheduleItems: []metronome.V1ContractAmendParamsCreditAccessScheduleScheduleItem{{
					Amount:       0,
					EndingBefore: time.Now(),
					StartingAt:   time.Now(),
				}},
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			ProductID:             "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Description: metronome.String("description"),
			HierarchyConfiguration: shared.CommitHierarchyConfigurationParam{
				ChildAccess: shared.CommitHierarchyConfigurationChildAccessUnionParam{
					OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll: &shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam{
						Type: "ALL",
					},
				},
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
			Priority:             metronome.Float(0),
			RateType:             "COMMIT_RATE",
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
		}},
		CustomFields: map[string]string{
			"foo": "string",
		},
		Discounts: []metronome.V1ContractAmendParamsDiscount{{
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Schedule: metronome.V1ContractAmendParamsDiscountSchedule{
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				DoNotInvoice: metronome.Bool(true),
				RecurringSchedule: metronome.V1ContractAmendParamsDiscountScheduleRecurringSchedule{
					AmountDistribution: "DIVIDED",
					EndingBefore:       time.Now(),
					Frequency:          "MONTHLY",
					StartingAt:         time.Now(),
					Amount:             metronome.Float(0),
					Quantity:           metronome.Float(0),
					UnitPrice:          metronome.Float(0),
				},
				ScheduleItems: []metronome.V1ContractAmendParamsDiscountScheduleScheduleItem{{
					Timestamp: time.Now(),
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					UnitPrice: metronome.Float(0),
				}},
			},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
		}},
		NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
		ContractOverrides: []metronome.V1ContractAmendParamsOverride{{
			StartingAt:            time.Now(),
			ApplicableProductTags: []string{"string"},
			EndingBefore:          metronome.Time(time.Now()),
			Entitled:              metronome.Bool(true),
			IsCommitSpecific:      metronome.Bool(true),
			Multiplier:            metronome.Float(0),
			OverrideSpecifiers: []metronome.V1ContractAmendParamsOverrideOverrideSpecifier{{
				BillingFrequency: "MONTHLY",
				CommitIDs:        []string{"string"},
				PresentationGroupValues: map[string]string{
					"foo": "string",
				},
				PricingGroupValues: map[string]string{
					"foo": "string",
				},
				ProductID:          metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags:        []string{"string"},
				RecurringCommitIDs: []string{"string"},
				RecurringCreditIDs: []string{"string"},
			}},
			OverwriteRate: metronome.V1ContractAmendParamsOverrideOverwriteRate{
				RateType:     "FLAT",
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				CustomRate: map[string]any{
					"foo": "bar",
				},
				IsProrated: metronome.Bool(true),
				Price:      metronome.Float(0),
				Quantity:   metronome.Float(0),
				Tiers: []shared.TierParam{{
					Price: 0,
					Size:  metronome.Float(0),
				}},
			},
			Priority:  metronome.Float(0),
			ProductID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Target:    "COMMIT_RATE",
			Tiers: []metronome.V1ContractAmendParamsOverrideTier{{
				Multiplier: 0,
				Size:       metronome.Float(0),
			}},
			Type: "OVERWRITE",
		}},
		ProfessionalServices: []metronome.V1ContractAmendParamsProfessionalService{{
			MaxAmount: 0,
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Quantity:  0,
			UnitPrice: 0,
			CustomFields: map[string]string{
				"foo": "string",
			},
			Description:          metronome.String("description"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
		}},
		ResellerRoyalties: []metronome.V1ContractAmendParamsResellerRoyalty{{
			ResellerType:          "AWS",
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			AwsOptions: metronome.V1ContractAmendParamsResellerRoyaltyAwsOptions{
				AwsAccountNumber:    metronome.String("aws_account_number"),
				AwsOfferID:          metronome.String("aws_offer_id"),
				AwsPayerReferenceID: metronome.String("aws_payer_reference_id"),
			},
			EndingBefore: metronome.Time(time.Now()),
			Fraction:     metronome.Float(0),
			GcpOptions: metronome.V1ContractAmendParamsResellerRoyaltyGcpOptions{
				GcpAccountID: metronome.String("gcp_account_id"),
				GcpOfferID:   metronome.String("gcp_offer_id"),
			},
			NetsuiteResellerID:    metronome.String("netsuite_reseller_id"),
			ResellerContractValue: metronome.Float(0),
			StartingAt:            metronome.Time(time.Now()),
		}},
		SalesforceOpportunityID: metronome.String("salesforce_opportunity_id"),
		ScheduledCharges: []metronome.V1ContractAmendParamsScheduledCharge{{
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Schedule: metronome.V1ContractAmendParamsScheduledChargeSchedule{
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				DoNotInvoice: metronome.Bool(true),
				RecurringSchedule: metronome.V1ContractAmendParamsScheduledChargeScheduleRecurringSchedule{
					AmountDistribution: "DIVIDED",
					EndingBefore:       time.Now(),
					Frequency:          "MONTHLY",
					StartingAt:         time.Now(),
					Amount:             metronome.Float(0),
					Quantity:           metronome.Float(0),
					UnitPrice:          metronome.Float(0),
				},
				ScheduleItems: []metronome.V1ContractAmendParamsScheduledChargeScheduleScheduleItem{{
					Timestamp: time.Now(),
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					UnitPrice: metronome.Float(0),
				}},
			},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
		}},
		TotalContractValue: metronome.Float(0),
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
		ContractID:   "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:   "13117714-3f05-48e5-a6e9-a66093f13b4d",
		VoidInvoices: true,
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
		Invoices: []metronome.V1ContractNewHistoricalInvoicesParamsInvoice{{
			ContractID:         "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
			CreditTypeID:       "2714e483-4ff1-48e4-9e25-ac732e8f24f2",
			CustomerID:         "13117714-3f05-48e5-a6e9-a66093f13b4d",
			ExclusiveEndDate:   time.Now(),
			InclusiveStartDate: time.Now(),
			IssueDate:          time.Now(),
			UsageLineItems: []metronome.V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItem{{
				ExclusiveEndDate:   time.Now(),
				InclusiveStartDate: time.Now(),
				ProductID:          "f14d6729-6a44-4b13-9908-9387f1918790",
				PresentationGroupValues: map[string]string{
					"foo": "string",
				},
				PricingGroupValues: map[string]string{
					"foo": "string",
				},
				Quantity: metronome.Float(100),
				SubtotalsWithQuantity: []metronome.V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItemSubtotalsWithQuantity{{
					ExclusiveEndDate:   time.Now(),
					InclusiveStartDate: time.Now(),
					Quantity:           0,
				}},
			}},
			BillableStatus:       "billable",
			BreakdownGranularity: "HOUR",
			CustomFields: map[string]string{
				"foo": "string",
			},
		}},
		Preview: false,
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
		CustomerID:              "13117714-3f05-48e5-a6e9-a66093f13b4d",
		ID:                      metronome.String("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
		CoveringDate:            metronome.Time(time.Now()),
		EffectiveBefore:         metronome.Time(time.Now()),
		ExcludeZeroBalances:     metronome.Bool(true),
		IncludeArchived:         metronome.Bool(true),
		IncludeBalance:          metronome.Bool(true),
		IncludeContractBalances: metronome.Bool(true),
		IncludeLedgers:          metronome.Bool(true),
		Limit:                   metronome.Int(1),
		NextPage:                metronome.String("next_page"),
		StartingAt:              metronome.Time(time.Now()),
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
		ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		Limit:      metronome.Int(1),
		NextPage:   metronome.String("next_page"),
		At:         metronome.Time(time.Now()),
		Selectors: []metronome.V1ContractGetRateScheduleParamsSelector{{
			BillingFrequency: "MONTHLY",
			PartialPricingGroupValues: map[string]string{
				"region": "us-west-2",
				"cloud":  "aws",
			},
			PricingGroupValues: map[string]string{
				"foo": "string",
			},
			ProductID:   metronome.String("d6300dbb-882e-4d2d-8dec-5125d16b65d0"),
			ProductTags: []string{"string"},
		}},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContractGetSubscriptionQuantityHistory(t *testing.T) {
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
	_, err := client.V1.Contracts.GetSubscriptionQuantityHistory(context.TODO(), metronome.V1ContractGetSubscriptionQuantityHistoryParams{
		ContractID:     "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:     "13117714-3f05-48e5-a6e9-a66093f13b4d",
		SubscriptionID: "1a824d53-bde6-4d82-96d7-6347ff227d5c",
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
		ContractID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		CustomerID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		IssuedAt:   time.Now(),
		LineItems: []metronome.V1ContractScheduleProServicesInvoiceParamsLineItem{{
			ProfessionalServiceID:       "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			AmendmentID:                 metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Amount:                      metronome.Float(0),
			Metadata:                    metronome.String("metadata"),
			NetsuiteInvoiceBillingEnd:   metronome.Time(time.Now()),
			NetsuiteInvoiceBillingStart: metronome.Time(time.Now()),
			Quantity:                    metronome.Float(0),
			UnitPrice:                   metronome.Float(0),
		}},
		NetsuiteInvoiceHeaderEnd:   metronome.Time(time.Now()),
		NetsuiteInvoiceHeaderStart: metronome.Time(time.Now()),
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
		ContractID:  "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:  "13117714-3f05-48e5-a6e9-a66093f13b4d",
		GroupKey:    "business_subscription_id",
		GroupValues: []string{"ID-1", "ID-2"},
		StartingAt:  time.Now(),
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
		ContractID:                        "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:                        "13117714-3f05-48e5-a6e9-a66093f13b4d",
		AllowEndingBeforeFinalizedInvoice: metronome.Bool(true),
		EndingBefore:                      metronome.Time(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
