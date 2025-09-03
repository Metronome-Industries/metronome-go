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

func TestV2ContractGetWithOptionalParams(t *testing.T) {
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
	_, err := client.V2.Contracts.Get(context.TODO(), metronome.V2ContractGetParams{
		ContractID:     "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:     "13117714-3f05-48e5-a6e9-a66093f13b4d",
		AsOfDate:       metronome.Time(time.Now()),
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

func TestV2ContractListWithOptionalParams(t *testing.T) {
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
	_, err := client.V2.Contracts.List(context.TODO(), metronome.V2ContractListParams{
		CustomerID:      "13117714-3f05-48e5-a6e9-a66093f13b4d",
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

func TestV2ContractEditWithOptionalParams(t *testing.T) {
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
	_, err := client.V2.Contracts.Edit(context.TODO(), metronome.V2ContractEditParams{
		ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
		AddCommits: []metronome.V2ContractEditParamsAddCommit{{
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Type:      "PREPAID",
			AccessSchedule: metronome.V2ContractEditParamsAddCommitAccessSchedule{
				ScheduleItems: []metronome.V2ContractEditParamsAddCommitAccessScheduleScheduleItem{{
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
			InvoiceSchedule: metronome.V2ContractEditParamsAddCommitInvoiceSchedule{
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				DoNotInvoice: metronome.Bool(true),
				RecurringSchedule: metronome.V2ContractEditParamsAddCommitInvoiceScheduleRecurringSchedule{
					AmountDistribution: "DIVIDED",
					EndingBefore:       time.Now(),
					Frequency:          "MONTHLY",
					StartingAt:         time.Now(),
					Amount:             metronome.Float(0),
					Quantity:           metronome.Float(0),
					UnitPrice:          metronome.Float(0),
				},
				ScheduleItems: []metronome.V2ContractEditParamsAddCommitInvoiceScheduleScheduleItem{{
					Timestamp: time.Now(),
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					UnitPrice: metronome.Float(0),
				}},
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
			PaymentGateConfig: metronome.V2ContractEditParamsAddCommitPaymentGateConfig{
				PaymentGateType: "NONE",
				PrecalculatedTaxConfig: metronome.V2ContractEditParamsAddCommitPaymentGateConfigPrecalculatedTaxConfig{
					TaxAmount: 0,
					TaxName:   metronome.String("tax_name"),
				},
				StripeConfig: metronome.V2ContractEditParamsAddCommitPaymentGateConfigStripeConfig{
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
		AddCredits: []metronome.V2ContractEditParamsAddCredit{{
			AccessSchedule: metronome.V2ContractEditParamsAddCreditAccessSchedule{
				ScheduleItems: []metronome.V2ContractEditParamsAddCreditAccessScheduleScheduleItem{{
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
		AddDiscounts: []metronome.V2ContractEditParamsAddDiscount{{
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Schedule: metronome.V2ContractEditParamsAddDiscountSchedule{
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				DoNotInvoice: metronome.Bool(true),
				RecurringSchedule: metronome.V2ContractEditParamsAddDiscountScheduleRecurringSchedule{
					AmountDistribution: "DIVIDED",
					EndingBefore:       time.Now(),
					Frequency:          "MONTHLY",
					StartingAt:         time.Now(),
					Amount:             metronome.Float(0),
					Quantity:           metronome.Float(0),
					UnitPrice:          metronome.Float(0),
				},
				ScheduleItems: []metronome.V2ContractEditParamsAddDiscountScheduleScheduleItem{{
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
		AddOverrides: []metronome.V2ContractEditParamsAddOverride{{
			StartingAt:            time.Now(),
			ApplicableProductTags: []string{"string"},
			EndingBefore:          metronome.Time(time.Now()),
			Entitled:              metronome.Bool(true),
			IsCommitSpecific:      metronome.Bool(true),
			Multiplier:            metronome.Float(2),
			OverrideSpecifiers: []metronome.V2ContractEditParamsAddOverrideOverrideSpecifier{{
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
			OverwriteRate: metronome.V2ContractEditParamsAddOverrideOverwriteRate{
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
			Priority:  metronome.Float(100),
			ProductID: metronome.String("d4fc086c-d8e5-4091-a235-fbba5da4ec14"),
			Target:    "COMMIT_RATE",
			Tiers: []metronome.V2ContractEditParamsAddOverrideTier{{
				Multiplier: 0,
				Size:       metronome.Float(0),
			}},
			Type: "MULTIPLIER",
		}},
		AddPrepaidBalanceThresholdConfiguration: shared.PrepaidBalanceThresholdConfigurationV2Param{
			Commit: shared.PrepaidBalanceThresholdConfigurationV2CommitParam{
				UpdateBaseThresholdCommitParam: shared.UpdateBaseThresholdCommitParam{
					Description: metronome.String("description"),
					Name:        metronome.String("name"),
					ProductID:   metronome.String("product_id"),
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
			PaymentGateConfig: shared.PaymentGateConfigV2Param{
				PaymentGateType: shared.PaymentGateConfigV2PaymentGateTypeNone,
				PrecalculatedTaxConfig: shared.PaymentGateConfigV2PrecalculatedTaxConfigParam{
					TaxAmount: 0,
					TaxName:   metronome.String("tax_name"),
				},
				StripeConfig: shared.PaymentGateConfigV2StripeConfigParam{
					PaymentType: "INVOICE",
					InvoiceMetadata: map[string]string{
						"foo": "string",
					},
				},
				TaxType: shared.PaymentGateConfigV2TaxTypeNone,
			},
			RechargeToAmount:   0,
			ThresholdAmount:    0,
			CustomCreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
		AddProfessionalServices: []metronome.V2ContractEditParamsAddProfessionalService{{
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
		AddRecurringCommits: []metronome.V2ContractEditParamsAddRecurringCommit{{
			AccessAmount: metronome.V2ContractEditParamsAddRecurringCommitAccessAmount{
				CreditTypeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				UnitPrice:    0,
				Quantity:     metronome.Float(0),
			},
			CommitDuration: metronome.V2ContractEditParamsAddRecurringCommitCommitDuration{
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
			InvoiceAmount: metronome.V2ContractEditParamsAddRecurringCommitInvoiceAmount{
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
			SubscriptionConfig: metronome.V2ContractEditParamsAddRecurringCommitSubscriptionConfig{
				ApplySeatIncreaseConfig: metronome.V2ContractEditParamsAddRecurringCommitSubscriptionConfigApplySeatIncreaseConfig{
					IsProrated: true,
				},
				SubscriptionID: "subscription_id",
				Allocation:     "POOLED",
			},
			TemporaryID: metronome.String("temporary_id"),
		}},
		AddRecurringCredits: []metronome.V2ContractEditParamsAddRecurringCredit{{
			AccessAmount: metronome.V2ContractEditParamsAddRecurringCreditAccessAmount{
				CreditTypeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				UnitPrice:    0,
				Quantity:     metronome.Float(0),
			},
			CommitDuration: metronome.V2ContractEditParamsAddRecurringCreditCommitDuration{
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
			SubscriptionConfig: metronome.V2ContractEditParamsAddRecurringCreditSubscriptionConfig{
				ApplySeatIncreaseConfig: metronome.V2ContractEditParamsAddRecurringCreditSubscriptionConfigApplySeatIncreaseConfig{
					IsProrated: true,
				},
				SubscriptionID: "subscription_id",
				Allocation:     "POOLED",
			},
			TemporaryID: metronome.String("temporary_id"),
		}},
		AddResellerRoyalties: []metronome.V2ContractEditParamsAddResellerRoyalty{{
			ResellerType:          "AWS",
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			AwsOptions: metronome.V2ContractEditParamsAddResellerRoyaltyAwsOptions{
				AwsAccountNumber:    metronome.String("aws_account_number"),
				AwsOfferID:          metronome.String("aws_offer_id"),
				AwsPayerReferenceID: metronome.String("aws_payer_reference_id"),
			},
			EndingBefore: metronome.Time(time.Now()),
			Fraction:     metronome.Float(0),
			GcpOptions: metronome.V2ContractEditParamsAddResellerRoyaltyGcpOptions{
				GcpAccountID: metronome.String("gcp_account_id"),
				GcpOfferID:   metronome.String("gcp_offer_id"),
			},
			NetsuiteResellerID:    metronome.String("netsuite_reseller_id"),
			ResellerContractValue: metronome.Float(0),
			StartingAt:            metronome.Time(time.Now()),
		}},
		AddScheduledCharges: []metronome.V2ContractEditParamsAddScheduledCharge{{
			ProductID: "2e30f074-d04c-412e-a134-851ebfa5ceb2",
			Schedule: metronome.V2ContractEditParamsAddScheduledChargeSchedule{
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				DoNotInvoice: metronome.Bool(true),
				RecurringSchedule: metronome.V2ContractEditParamsAddScheduledChargeScheduleRecurringSchedule{
					AmountDistribution: "DIVIDED",
					EndingBefore:       time.Now(),
					Frequency:          "MONTHLY",
					StartingAt:         time.Now(),
					Amount:             metronome.Float(0),
					Quantity:           metronome.Float(0),
					UnitPrice:          metronome.Float(0),
				},
				ScheduleItems: []metronome.V2ContractEditParamsAddScheduledChargeScheduleScheduleItem{{
					Timestamp: time.Now(),
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(1),
					UnitPrice: metronome.Float(1000000),
				}},
			},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Name:                 metronome.String("x"),
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
		}},
		AddSpendThresholdConfiguration: shared.SpendThresholdConfigurationV2Param{
			Commit: shared.UpdateBaseThresholdCommitParam{
				Description: metronome.String("description"),
				Name:        metronome.String("name"),
				ProductID:   metronome.String("product_id"),
			},
			IsEnabled: true,
			PaymentGateConfig: shared.PaymentGateConfigV2Param{
				PaymentGateType: shared.PaymentGateConfigV2PaymentGateTypeNone,
				PrecalculatedTaxConfig: shared.PaymentGateConfigV2PrecalculatedTaxConfigParam{
					TaxAmount: 0,
					TaxName:   metronome.String("tax_name"),
				},
				StripeConfig: shared.PaymentGateConfigV2StripeConfigParam{
					PaymentType: "INVOICE",
					InvoiceMetadata: map[string]string{
						"foo": "string",
					},
				},
				TaxType: shared.PaymentGateConfigV2TaxTypeNone,
			},
			ThresholdAmount: 0,
		},
		AddSubscriptions: []metronome.V2ContractEditParamsAddSubscription{{
			CollectionSchedule: "ADVANCE",
			Proration: metronome.V2ContractEditParamsAddSubscriptionProration{
				InvoiceBehavior: "BILL_IMMEDIATELY",
				IsProrated:      metronome.Bool(true),
			},
			SubscriptionRate: metronome.V2ContractEditParamsAddSubscriptionSubscriptionRate{
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
			StartingAt:             metronome.Time(time.Now()),
			TemporaryID:            metronome.String("temporary_id"),
		}},
		AllowContractEndingBeforeFinalizedInvoice: metronome.Bool(true),
		ArchiveCommits: []metronome.V2ContractEditParamsArchiveCommit{{
			ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		}},
		ArchiveCredits: []metronome.V2ContractEditParamsArchiveCredit{{
			ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		}},
		ArchiveScheduledCharges: []metronome.V2ContractEditParamsArchiveScheduledCharge{{
			ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		}},
		RemoveOverrides: []metronome.V2ContractEditParamsRemoveOverride{{
			ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		}},
		UniquenessKey: metronome.String("x"),
		UpdateCommits: []metronome.V2ContractEditParamsUpdateCommit{{
			CommitID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			AccessSchedule: metronome.V2ContractEditParamsUpdateCommitAccessSchedule{
				AddScheduleItems: []metronome.V2ContractEditParamsUpdateCommitAccessScheduleAddScheduleItem{{
					Amount:       0,
					EndingBefore: time.Now(),
					StartingAt:   time.Now(),
				}},
				RemoveScheduleItems: []metronome.V2ContractEditParamsUpdateCommitAccessScheduleRemoveScheduleItem{{
					ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				}},
				UpdateScheduleItems: []metronome.V2ContractEditParamsUpdateCommitAccessScheduleUpdateScheduleItem{{
					ID:           "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					Amount:       metronome.Float(0),
					EndingBefore: metronome.Time(time.Now()),
					StartingAt:   metronome.Time(time.Now()),
				}},
			},
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			HierarchyConfiguration: shared.CommitHierarchyConfigurationParam{
				ChildAccess: shared.CommitHierarchyConfigurationChildAccessUnionParam{
					OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll: &shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam{
						Type: "ALL",
					},
				},
			},
			InvoiceSchedule: metronome.V2ContractEditParamsUpdateCommitInvoiceSchedule{
				AddScheduleItems: []metronome.V2ContractEditParamsUpdateCommitInvoiceScheduleAddScheduleItem{{
					Timestamp: time.Now(),
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					UnitPrice: metronome.Float(0),
				}},
				RemoveScheduleItems: []metronome.V2ContractEditParamsUpdateCommitInvoiceScheduleRemoveScheduleItem{{
					ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				}},
				UpdateScheduleItems: []metronome.V2ContractEditParamsUpdateCommitInvoiceScheduleUpdateScheduleItem{{
					ID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					Timestamp: metronome.Time(time.Now()),
					UnitPrice: metronome.Float(0),
				}},
			},
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
			Priority:             metronome.Float(0),
			ProductID:            metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			RolloverFraction:     metronome.Float(0),
		}},
		UpdateContractEndDate: metronome.Time(time.Now()),
		UpdateContractName:    metronome.String("update_contract_name"),
		UpdateCredits: []metronome.V2ContractEditParamsUpdateCredit{{
			CreditID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			AccessSchedule: metronome.V2ContractEditParamsUpdateCreditAccessSchedule{
				AddScheduleItems: []metronome.V2ContractEditParamsUpdateCreditAccessScheduleAddScheduleItem{{
					Amount:       0,
					EndingBefore: time.Now(),
					StartingAt:   time.Now(),
				}},
				RemoveScheduleItems: []metronome.V2ContractEditParamsUpdateCreditAccessScheduleRemoveScheduleItem{{
					ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				}},
				UpdateScheduleItems: []metronome.V2ContractEditParamsUpdateCreditAccessScheduleUpdateScheduleItem{{
					ID:           "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					Amount:       metronome.Float(0),
					EndingBefore: metronome.Time(time.Now()),
					StartingAt:   metronome.Time(time.Now()),
				}},
			},
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			HierarchyConfiguration: shared.CommitHierarchyConfigurationParam{
				ChildAccess: shared.CommitHierarchyConfigurationChildAccessUnionParam{
					OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll: &shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam{
						Type: "ALL",
					},
				},
			},
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
			Priority:             metronome.Float(0),
			ProductID:            metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}},
		UpdatePrepaidBalanceThresholdConfiguration: metronome.V2ContractEditParamsUpdatePrepaidBalanceThresholdConfiguration{
			Commit: metronome.V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommit{
				UpdateBaseThresholdCommitParam: shared.UpdateBaseThresholdCommitParam{
					Description: metronome.String("description"),
					Name:        metronome.String("name"),
					ProductID:   metronome.String("product_id"),
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
			CustomCreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			IsEnabled:          metronome.Bool(true),
			PaymentGateConfig: shared.PaymentGateConfigV2Param{
				PaymentGateType: shared.PaymentGateConfigV2PaymentGateTypeNone,
				PrecalculatedTaxConfig: shared.PaymentGateConfigV2PrecalculatedTaxConfigParam{
					TaxAmount: 0,
					TaxName:   metronome.String("tax_name"),
				},
				StripeConfig: shared.PaymentGateConfigV2StripeConfigParam{
					PaymentType: "INVOICE",
					InvoiceMetadata: map[string]string{
						"foo": "string",
					},
				},
				TaxType: shared.PaymentGateConfigV2TaxTypeNone,
			},
			RechargeToAmount: metronome.Float(0),
			ThresholdAmount:  metronome.Float(0),
		},
		UpdateRecurringCommits: []metronome.V2ContractEditParamsUpdateRecurringCommit{{
			RecurringCommitID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			AccessAmount: metronome.V2ContractEditParamsUpdateRecurringCommitAccessAmount{
				Quantity:  metronome.Float(0),
				UnitPrice: metronome.Float(0),
			},
			EndingBefore: metronome.Time(time.Now()),
			InvoiceAmount: metronome.V2ContractEditParamsUpdateRecurringCommitInvoiceAmount{
				Quantity:  metronome.Float(0),
				UnitPrice: metronome.Float(0),
			},
		}},
		UpdateRecurringCredits: []metronome.V2ContractEditParamsUpdateRecurringCredit{{
			RecurringCreditID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			AccessAmount: metronome.V2ContractEditParamsUpdateRecurringCreditAccessAmount{
				Quantity:  metronome.Float(0),
				UnitPrice: metronome.Float(0),
			},
			EndingBefore: metronome.Time(time.Now()),
		}},
		UpdateScheduledCharges: []metronome.V2ContractEditParamsUpdateScheduledCharge{{
			ScheduledChargeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			InvoiceSchedule: metronome.V2ContractEditParamsUpdateScheduledChargeInvoiceSchedule{
				AddScheduleItems: []metronome.V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleAddScheduleItem{{
					Timestamp: time.Now(),
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					UnitPrice: metronome.Float(0),
				}},
				RemoveScheduleItems: []metronome.V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem{{
					ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				}},
				UpdateScheduleItems: []metronome.V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem{{
					ID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					Timestamp: metronome.Time(time.Now()),
					UnitPrice: metronome.Float(0),
				}},
			},
			NetsuiteSalesOrderID: metronome.String("netsuite_sales_order_id"),
		}},
		UpdateSpendThresholdConfiguration: metronome.V2ContractEditParamsUpdateSpendThresholdConfiguration{
			Commit: shared.UpdateBaseThresholdCommitParam{
				Description: metronome.String("description"),
				Name:        metronome.String("name"),
				ProductID:   metronome.String("product_id"),
			},
			IsEnabled: metronome.Bool(true),
			PaymentGateConfig: shared.PaymentGateConfigV2Param{
				PaymentGateType: shared.PaymentGateConfigV2PaymentGateTypeNone,
				PrecalculatedTaxConfig: shared.PaymentGateConfigV2PrecalculatedTaxConfigParam{
					TaxAmount: 0,
					TaxName:   metronome.String("tax_name"),
				},
				StripeConfig: shared.PaymentGateConfigV2StripeConfigParam{
					PaymentType: "INVOICE",
					InvoiceMetadata: map[string]string{
						"foo": "string",
					},
				},
				TaxType: shared.PaymentGateConfigV2TaxTypeNone,
			},
			ThresholdAmount: metronome.Float(0),
		},
		UpdateSubscriptions: []metronome.V2ContractEditParamsUpdateSubscription{{
			SubscriptionID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			EndingBefore:   metronome.Time(time.Now()),
			QuantityUpdates: []metronome.V2ContractEditParamsUpdateSubscriptionQuantityUpdate{{
				StartingAt:    time.Now(),
				Quantity:      metronome.Float(0),
				QuantityDelta: metronome.Float(0),
			}},
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

func TestV2ContractEditCommitWithOptionalParams(t *testing.T) {
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
	_, err := client.V2.Contracts.EditCommit(context.TODO(), metronome.V2ContractEditCommitParams{
		CommitID:   "5e7e82cf-ccb7-428c-a96f-a8e4f67af822",
		CustomerID: "4c91c473-fc12-445a-9c38-40421d47023f",
		AccessSchedule: metronome.V2ContractEditCommitParamsAccessSchedule{
			AddScheduleItems: []metronome.V2ContractEditCommitParamsAccessScheduleAddScheduleItem{{
				Amount:       0,
				EndingBefore: time.Now(),
				StartingAt:   time.Now(),
			}},
			RemoveScheduleItems: []metronome.V2ContractEditCommitParamsAccessScheduleRemoveScheduleItem{{
				ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			}},
			UpdateScheduleItems: []metronome.V2ContractEditCommitParamsAccessScheduleUpdateScheduleItem{{
				ID:           "d5edbd32-c744-48cb-9475-a9bca0e6fa39",
				Amount:       metronome.Float(0),
				EndingBefore: metronome.Time(time.Now()),
				StartingAt:   metronome.Time(time.Now()),
			}},
		},
		ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		ApplicableProductTags: []string{"string"},
		InvoiceContractID:     metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		InvoiceSchedule: metronome.V2ContractEditCommitParamsInvoiceSchedule{
			AddScheduleItems: []metronome.V2ContractEditCommitParamsInvoiceScheduleAddScheduleItem{{
				Timestamp: time.Now(),
				Amount:    metronome.Float(0),
				Quantity:  metronome.Float(0),
				UnitPrice: metronome.Float(0),
			}},
			RemoveScheduleItems: []metronome.V2ContractEditCommitParamsInvoiceScheduleRemoveScheduleItem{{
				ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			}},
			UpdateScheduleItems: []metronome.V2ContractEditCommitParamsInvoiceScheduleUpdateScheduleItem{{
				ID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Amount:    metronome.Float(0),
				Quantity:  metronome.Float(0),
				Timestamp: metronome.Time(time.Now()),
				UnitPrice: metronome.Float(0),
			}},
		},
		Priority:  metronome.Float(0),
		ProductID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ContractEditCreditWithOptionalParams(t *testing.T) {
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
	_, err := client.V2.Contracts.EditCredit(context.TODO(), metronome.V2ContractEditCreditParams{
		CreditID:   "5e7e82cf-ccb7-428c-a96f-a8e4f67af822",
		CustomerID: "4c91c473-fc12-445a-9c38-40421d47023f",
		AccessSchedule: metronome.V2ContractEditCreditParamsAccessSchedule{
			AddScheduleItems: []metronome.V2ContractEditCreditParamsAccessScheduleAddScheduleItem{{
				Amount:       0,
				EndingBefore: time.Now(),
				StartingAt:   time.Now(),
			}},
			RemoveScheduleItems: []metronome.V2ContractEditCreditParamsAccessScheduleRemoveScheduleItem{{
				ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			}},
			UpdateScheduleItems: []metronome.V2ContractEditCreditParamsAccessScheduleUpdateScheduleItem{{
				ID:           "d5edbd32-c744-48cb-9475-a9bca0e6fa39",
				Amount:       metronome.Float(0),
				EndingBefore: metronome.Time(time.Now()),
				StartingAt:   metronome.Time(time.Now()),
			}},
		},
		ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		ApplicableProductTags: []string{"string"},
		Priority:              metronome.Float(0),
		ProductID:             metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ContractGetEditHistory(t *testing.T) {
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
	_, err := client.V2.Contracts.GetEditHistory(context.TODO(), metronome.V2ContractGetEditHistoryParams{
		ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
