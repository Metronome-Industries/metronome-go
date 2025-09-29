// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
)

func TestPackageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Packages.New(context.TODO(), metronome.PackageNewParams{
		Name:              "My package",
		BillingAnchorDate: metronome.PackageNewParamsBillingAnchorDateContractStartDate,
		BillingProvider:   metronome.PackageNewParamsBillingProviderStripe,
		Commits: []metronome.PackageNewParamsCommit{{
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Type:      "PREPAID",
			AccessSchedule: metronome.PackageNewParamsCommitAccessSchedule{
				ScheduleItems: []metronome.PackageNewParamsCommitAccessScheduleScheduleItem{{
					Amount: 0,
					Duration: metronome.PackageNewParamsCommitAccessScheduleScheduleItemDuration{
						Unit:  "DAYS",
						Value: 1,
					},
					StartingAtOffset: metronome.PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset{
						Unit:  "DAYS",
						Value: 1,
					},
				}},
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Description: metronome.String("description"),
			InvoiceSchedule: metronome.PackageNewParamsCommitInvoiceSchedule{
				ScheduleItems: []metronome.PackageNewParamsCommitInvoiceScheduleScheduleItem{{
					DateOffset: metronome.PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset{
						Unit:  "DAYS",
						Value: 1,
					},
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					UnitPrice: metronome.Float(0),
				}},
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			Name: metronome.String("x"),
			PaymentGateConfig: metronome.PackageNewParamsCommitPaymentGateConfig{
				PaymentGateType: "NONE",
				PrecalculatedTaxConfig: metronome.PackageNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig{
					TaxAmount: 0,
					TaxName:   metronome.String("tax_name"),
				},
				StripeConfig: metronome.PackageNewParamsCommitPaymentGateConfigStripeConfig{
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
		ContractName: metronome.String("contract_name"),
		Credits: []metronome.PackageNewParamsCredit{{
			AccessSchedule: metronome.PackageNewParamsCreditAccessSchedule{
				ScheduleItems: []metronome.PackageNewParamsCreditAccessScheduleScheduleItem{{
					Amount: 0,
					Duration: metronome.PackageNewParamsCreditAccessScheduleScheduleItemDuration{
						Unit:  "DAYS",
						Value: 1,
					},
					StartingAtOffset: metronome.PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset{
						Unit:  "DAYS",
						Value: 1,
					},
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
			Name:        metronome.String("x"),
			Priority:    metronome.Float(0),
			RateType:    "COMMIT_RATE",
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
		DeliveryMethod: metronome.PackageNewParamsDeliveryMethodDirectToBillingProvider,
		Duration: metronome.PackageNewParamsDuration{
			Unit:  "DAYS",
			Value: 1,
		},
		MultiplierOverridePrioritization: metronome.PackageNewParamsMultiplierOverridePrioritizationLowestMultiplier,
		NetPaymentTermsDays:              metronome.Float(0),
		Overrides: []metronome.PackageNewParamsOverride{{
			OverrideSpecifiers: []metronome.PackageNewParamsOverrideOverrideSpecifier{{
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
			StartingOffset: metronome.PackageNewParamsOverrideStartingOffset{
				Unit:  "DAYS",
				Value: 1,
			},
			Duration: metronome.PackageNewParamsOverrideDuration{
				Unit:  "DAYS",
				Value: 1,
			},
			Entitled:         metronome.Bool(true),
			IsCommitSpecific: metronome.Bool(true),
			Multiplier:       metronome.Float(0),
			OverwriteRate: metronome.PackageNewParamsOverrideOverwriteRate{
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
			Priority: metronome.Float(0),
			Target:   "COMMIT_RATE",
			Tiers: []metronome.PackageNewParamsOverrideTier{{
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
		Priority:      metronome.Float(0),
		RateCardAlias: metronome.String("rate_card_alias"),
		RateCardID:    metronome.String("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		RecurringCommits: []metronome.PackageNewParamsRecurringCommit{{
			AccessAmount: metronome.PackageNewParamsRecurringCommitAccessAmount{
				CreditTypeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				UnitPrice:    0,
				Quantity:     metronome.Float(0),
			},
			CommitDuration: metronome.PackageNewParamsRecurringCommitCommitDuration{
				Value: 0,
				Unit:  "PERIODS",
			},
			Priority:  0,
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			StartingOffset: metronome.PackageNewParamsRecurringCommitStartingOffset{
				Unit:  "DAYS",
				Value: 1,
			},
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			Description:           metronome.String("description"),
			EndingOffset: metronome.PackageNewParamsRecurringCommitEndingOffset{
				Unit:  "DAYS",
				Value: 1,
			},
			InvoiceAmount: metronome.PackageNewParamsRecurringCommitInvoiceAmount{
				CreditTypeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Quantity:     0,
				UnitPrice:    0,
			},
			Name:                metronome.String("x"),
			Proration:           "NONE",
			RateType:            "COMMIT_RATE",
			RecurrenceFrequency: "MONTHLY",
			RolloverFraction:    metronome.Float(0),
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
			SubscriptionConfig: metronome.PackageNewParamsRecurringCommitSubscriptionConfig{
				ApplySeatIncreaseConfig: metronome.PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig{
					IsProrated: true,
				},
				SubscriptionID: "subscription_id",
				Allocation:     "INDIVIDUAL",
			},
			TemporaryID: metronome.String("temporary_id"),
		}},
		RecurringCredits: []metronome.PackageNewParamsRecurringCredit{{
			AccessAmount: metronome.PackageNewParamsRecurringCreditAccessAmount{
				CreditTypeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				UnitPrice:    0,
				Quantity:     metronome.Float(0),
			},
			CommitDuration: metronome.PackageNewParamsRecurringCreditCommitDuration{
				Value: 0,
				Unit:  "PERIODS",
			},
			Priority:  0,
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			StartingOffset: metronome.PackageNewParamsRecurringCreditStartingOffset{
				Unit:  "DAYS",
				Value: 1,
			},
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			Description:           metronome.String("description"),
			EndingOffset: metronome.PackageNewParamsRecurringCreditEndingOffset{
				Unit:  "DAYS",
				Value: 1,
			},
			Name:                metronome.String("x"),
			Proration:           "NONE",
			RateType:            "COMMIT_RATE",
			RecurrenceFrequency: "MONTHLY",
			RolloverFraction:    metronome.Float(0),
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
			SubscriptionConfig: metronome.PackageNewParamsRecurringCreditSubscriptionConfig{
				ApplySeatIncreaseConfig: metronome.PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig{
					IsProrated: true,
				},
				SubscriptionID: "subscription_id",
				Allocation:     "INDIVIDUAL",
			},
			TemporaryID: metronome.String("temporary_id"),
		}},
		ScheduledCharges: []metronome.PackageNewParamsScheduledCharge{{
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Schedule: metronome.PackageNewParamsScheduledChargeSchedule{
				ScheduleItems: []metronome.PackageNewParamsScheduledChargeScheduleScheduleItem{{
					DateOffset: metronome.PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset{
						Unit:  "DAYS",
						Value: 1,
					},
					Amount:    metronome.Float(0),
					Quantity:  metronome.Float(0),
					UnitPrice: metronome.Float(0),
				}},
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Name: metronome.String("x"),
		}},
		ScheduledChargesOnUsageInvoices: metronome.PackageNewParamsScheduledChargesOnUsageInvoicesAll,
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
		Subscriptions: []metronome.PackageNewParamsSubscription{{
			CollectionSchedule: "ADVANCE",
			InitialQuantity:    0,
			Proration: metronome.PackageNewParamsSubscriptionProration{
				InvoiceBehavior: "BILL_IMMEDIATELY",
				IsProrated:      metronome.Bool(true),
			},
			SubscriptionRate: metronome.PackageNewParamsSubscriptionSubscriptionRate{
				BillingFrequency: "MONTHLY",
				ProductID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Description: metronome.String("description"),
			Duration: metronome.PackageNewParamsSubscriptionDuration{
				Unit:  "DAYS",
				Value: 1,
			},
			Name: metronome.String("name"),
			StartingOffset: metronome.PackageNewParamsSubscriptionStartingOffset{
				Unit:  "DAYS",
				Value: 1,
			},
			TemporaryID: metronome.String("temporary_id"),
		}},
		UniquenessKey: metronome.String("x"),
		UsageStatementSchedule: metronome.PackageNewParamsUsageStatementSchedule{
			Frequency: "MONTHLY",
			Day:       "FIRST_OF_MONTH",
			InvoiceGenerationStartingAtOffset: metronome.PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset{
				Unit:  "DAYS",
				Value: 1,
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
