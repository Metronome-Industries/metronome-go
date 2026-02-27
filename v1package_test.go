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

func TestV1PackageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Packages.New(context.TODO(), metronome.V1PackageNewParams{
		Name: "My package",
		Aliases: []metronome.V1PackageNewParamsAlias{{
			Name:         "name",
			EndingBefore: metronome.Time(time.Now()),
			StartingAt:   metronome.Time(time.Now()),
		}},
		BillingAnchorDate: metronome.V1PackageNewParamsBillingAnchorDateContractStartDate,
		BillingProvider:   metronome.V1PackageNewParamsBillingProviderStripe,
		Commits: []metronome.V1PackageNewParamsCommit{{
			AccessSchedule: metronome.V1PackageNewParamsCommitAccessSchedule{
				ScheduleItems: []metronome.V1PackageNewParamsCommitAccessScheduleScheduleItem{{
					Amount: 0,
					Duration: metronome.V1PackageNewParamsCommitAccessScheduleScheduleItemDuration{
						Unit:  "DAYS",
						Value: 0,
					},
					StartingAtOffset: metronome.V1PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset{
						Unit:  "DAYS",
						Value: 0,
					},
				}},
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			ProductID:             "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Type:                  "PREPAID",
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Description: metronome.String("description"),
			InvoiceSchedule: metronome.V1PackageNewParamsCommitInvoiceSchedule{
				ScheduleItems: []metronome.V1PackageNewParamsCommitInvoiceScheduleScheduleItem{{
					DateOffset: metronome.V1PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset{
						Unit:  "DAYS",
						Value: 0,
					},
					Quantity:  0,
					UnitPrice: 0,
				}},
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				DoNotInvoice: metronome.Bool(true),
			},
			Name:             metronome.String("x"),
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
		Credits: []metronome.V1PackageNewParamsCredit{{
			AccessSchedule: metronome.V1PackageNewParamsCreditAccessSchedule{
				ScheduleItems: []metronome.V1PackageNewParamsCreditAccessScheduleScheduleItem{{
					Amount: 0,
					Duration: metronome.V1PackageNewParamsCreditAccessScheduleScheduleItemDuration{
						Unit:  "DAYS",
						Value: 0,
					},
					StartingAtOffset: metronome.V1PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset{
						Unit:  "DAYS",
						Value: 0,
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
		DeliveryMethod: metronome.V1PackageNewParamsDeliveryMethodDirectToBillingProvider,
		Duration: metronome.V1PackageNewParamsDuration{
			Unit:  "DAYS",
			Value: 0,
		},
		MultiplierOverridePrioritization: metronome.V1PackageNewParamsMultiplierOverridePrioritizationLowestMultiplier,
		NetPaymentTermsDays:              metronome.Float(0),
		ContractOverrides: []metronome.V1PackageNewParamsOverride{{
			OverrideSpecifiers: []metronome.V1PackageNewParamsOverrideOverrideSpecifier{{
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
			StartingAtOffset: metronome.V1PackageNewParamsOverrideStartingAtOffset{
				Unit:  "DAYS",
				Value: 0,
			},
			Duration: metronome.V1PackageNewParamsOverrideDuration{
				Unit:  "DAYS",
				Value: 0,
			},
			Entitled:         metronome.Bool(true),
			IsCommitSpecific: metronome.Bool(true),
			Multiplier:       metronome.Float(0),
			OverwriteRate: metronome.V1PackageNewParamsOverrideOverwriteRate{
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
			Tiers: []metronome.V1PackageNewParamsOverrideTier{{
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
		RateCardAlias: metronome.String("rate_card_alias"),
		RateCardID:    metronome.String("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		RecurringCommits: []metronome.V1PackageNewParamsRecurringCommit{{
			AccessAmount: metronome.V1PackageNewParamsRecurringCommitAccessAmount{
				CreditTypeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				UnitPrice:    0,
				Quantity:     metronome.Float(0),
			},
			CommitDuration: metronome.V1PackageNewParamsRecurringCommitCommitDuration{
				Value: 0,
				Unit:  "PERIODS",
			},
			Priority:  0,
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			StartingAtOffset: metronome.V1PackageNewParamsRecurringCommitStartingAtOffset{
				Unit:  "DAYS",
				Value: 0,
			},
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			Description:           metronome.String("description"),
			Duration: metronome.V1PackageNewParamsRecurringCommitDuration{
				Unit:  "DAYS",
				Value: 0,
			},
			InvoiceAmount: metronome.V1PackageNewParamsRecurringCommitInvoiceAmount{
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
			SubscriptionConfig: metronome.V1PackageNewParamsRecurringCommitSubscriptionConfig{
				ApplySeatIncreaseConfig: metronome.V1PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig{
					IsProrated: true,
				},
				SubscriptionID: "subscription_id",
				Allocation:     "INDIVIDUAL",
			},
			TemporaryID: metronome.String("temporary_id"),
		}},
		RecurringCredits: []metronome.V1PackageNewParamsRecurringCredit{{
			AccessAmount: metronome.V1PackageNewParamsRecurringCreditAccessAmount{
				CreditTypeID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				UnitPrice:    0,
				Quantity:     metronome.Float(0),
			},
			CommitDuration: metronome.V1PackageNewParamsRecurringCreditCommitDuration{
				Value: 0,
				Unit:  "PERIODS",
			},
			Priority:  0,
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			StartingAtOffset: metronome.V1PackageNewParamsRecurringCreditStartingAtOffset{
				Unit:  "DAYS",
				Value: 0,
			},
			ApplicableProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			ApplicableProductTags: []string{"string"},
			Description:           metronome.String("description"),
			Duration: metronome.V1PackageNewParamsRecurringCreditDuration{
				Unit:  "DAYS",
				Value: 0,
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
			SubscriptionConfig: metronome.V1PackageNewParamsRecurringCreditSubscriptionConfig{
				ApplySeatIncreaseConfig: metronome.V1PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig{
					IsProrated: true,
				},
				SubscriptionID: "subscription_id",
				Allocation:     "INDIVIDUAL",
			},
			TemporaryID: metronome.String("temporary_id"),
		}},
		ScheduledCharges: []metronome.V1PackageNewParamsScheduledCharge{{
			ProductID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Schedule: metronome.V1PackageNewParamsScheduledChargeSchedule{
				ScheduleItems: []metronome.V1PackageNewParamsScheduledChargeScheduleScheduleItem{{
					DateOffset: metronome.V1PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset{
						Unit:  "DAYS",
						Value: 0,
					},
					Quantity:  0,
					UnitPrice: 0,
				}},
				CreditTypeID: metronome.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Name: metronome.String("x"),
		}},
		ScheduledChargesOnUsageInvoices: metronome.V1PackageNewParamsScheduledChargesOnUsageInvoicesAll,
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
		Subscriptions: []metronome.V1PackageNewParamsSubscription{{
			CollectionSchedule: "ADVANCE",
			Proration: metronome.V1PackageNewParamsSubscriptionProration{
				InvoiceBehavior: "BILL_IMMEDIATELY",
				IsProrated:      metronome.Bool(true),
			},
			SubscriptionRate: metronome.V1PackageNewParamsSubscriptionSubscriptionRate{
				BillingFrequency: "MONTHLY",
				ProductID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			},
			CustomFields: map[string]string{
				"foo": "string",
			},
			Description: metronome.String("description"),
			Duration: metronome.V1PackageNewParamsSubscriptionDuration{
				Unit:  "DAYS",
				Value: 0,
			},
			InitialQuantity:        metronome.Float(0),
			Name:                   metronome.String("name"),
			QuantityManagementMode: "SEAT_BASED",
			SeatConfig: metronome.V1PackageNewParamsSubscriptionSeatConfig{
				SeatGroupKey:           "seat_group_key",
				InitialUnassignedSeats: metronome.Float(0),
			},
			StartingAtOffset: metronome.V1PackageNewParamsSubscriptionStartingAtOffset{
				Unit:  "DAYS",
				Value: 0,
			},
			TemporaryID: metronome.String("temporary_id"),
		}},
		UniquenessKey: metronome.String("x"),
		UsageStatementSchedule: metronome.V1PackageNewParamsUsageStatementSchedule{
			Frequency: "MONTHLY",
			Day:       "FIRST_OF_MONTH",
			InvoiceGenerationStartingAtOffset: metronome.V1PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset{
				Unit:  "DAYS",
				Value: 0,
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

func TestV1PackageGet(t *testing.T) {
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
	_, err := client.V1.Packages.Get(context.TODO(), metronome.V1PackageGetParams{
		PackageID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PackageListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Packages.List(context.TODO(), metronome.V1PackageListParams{
		Limit:         metronome.Int(1),
		NextPage:      metronome.String("next_page"),
		ArchiveFilter: metronome.V1PackageListParamsArchiveFilterArchived,
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PackageArchive(t *testing.T) {
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
	_, err := client.V1.Packages.Archive(context.TODO(), metronome.V1PackageArchiveParams{
		PackageID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PackageListContractsOnPackageWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Packages.ListContractsOnPackage(context.TODO(), metronome.V1PackageListContractsOnPackageParams{
		PackageID:       "13117714-3f05-48e5-a6e9-a66093f13b4d",
		Limit:           metronome.Int(1),
		NextPage:        metronome.String("next_page"),
		CoveringDate:    metronome.Time(time.Now()),
		IncludeArchived: metronome.Bool(true),
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
