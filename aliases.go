// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"github.com/Metronome-Industries/metronome-go/internal/apierror"
	"github.com/Metronome-Industries/metronome-go/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type BaseUsageFilter = shared.BaseUsageFilter

// This is an alias to an internal type.
type BaseUsageFilterParam = shared.BaseUsageFilterParam

// This is an alias to an internal type.
type Commit = shared.Commit

// This is an alias to an internal type.
type CommitProduct = shared.CommitProduct

// This is an alias to an internal type.
type CommitType = shared.CommitType

// This is an alias to an internal value.
const CommitTypePrepaid = shared.CommitTypePrepaid

// This is an alias to an internal value.
const CommitTypePostpaid = shared.CommitTypePostpaid

// This is an alias to an internal type.
type CommitContract = shared.CommitContract

// Optional configuration for commit hierarchy access control
//
// This is an alias to an internal type.
type CommitHierarchyConfiguration = shared.CommitHierarchyConfiguration

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccess = shared.CommitHierarchyConfigurationChildAccess

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType

// This is an alias to an internal value.
const CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType

// This is an alias to an internal value.
const CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType

// This is an alias to an internal value.
const CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessType = shared.CommitHierarchyConfigurationChildAccessType

// This is an alias to an internal value.
const CommitHierarchyConfigurationChildAccessTypeAll = shared.CommitHierarchyConfigurationChildAccessTypeAll

// This is an alias to an internal value.
const CommitHierarchyConfigurationChildAccessTypeNone = shared.CommitHierarchyConfigurationChildAccessTypeNone

// This is an alias to an internal value.
const CommitHierarchyConfigurationChildAccessTypeContractIDs = shared.CommitHierarchyConfigurationChildAccessTypeContractIDs

// The contract that this commit will be billed on.
//
// This is an alias to an internal type.
type CommitInvoiceContract = shared.CommitInvoiceContract

// This is an alias to an internal type.
type CommitLedger = shared.CommitLedger

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitSegmentStartLedgerEntry = shared.CommitLedgerPrepaidCommitSegmentStartLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitSegmentStartLedgerEntryType = shared.CommitLedgerPrepaidCommitSegmentStartLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart = shared.CommitLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry = shared.CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = shared.CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction = shared.CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitRolloverLedgerEntry = shared.CommitLedgerPrepaidCommitRolloverLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitRolloverLedgerEntryType = shared.CommitLedgerPrepaidCommitRolloverLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover = shared.CommitLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitExpirationLedgerEntry = shared.CommitLedgerPrepaidCommitExpirationLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitExpirationLedgerEntryType = shared.CommitLedgerPrepaidCommitExpirationLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration = shared.CommitLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitCanceledLedgerEntry = shared.CommitLedgerPrepaidCommitCanceledLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitCanceledLedgerEntryType = shared.CommitLedgerPrepaidCommitCanceledLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled = shared.CommitLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitCreditedLedgerEntry = shared.CommitLedgerPrepaidCommitCreditedLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitCreditedLedgerEntryType = shared.CommitLedgerPrepaidCommitCreditedLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited = shared.CommitLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry = shared.CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType = shared.CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryTypePrepaidCommitSeatBasedAdjustment = shared.CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryTypePrepaidCommitSeatBasedAdjustment

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitInitialBalanceLedgerEntry = shared.CommitLedgerPostpaidCommitInitialBalanceLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType = shared.CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance = shared.CommitLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry = shared.CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = shared.CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction = shared.CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitRolloverLedgerEntry = shared.CommitLedgerPostpaidCommitRolloverLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitRolloverLedgerEntryType = shared.CommitLedgerPostpaidCommitRolloverLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover = shared.CommitLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitTrueupLedgerEntry = shared.CommitLedgerPostpaidCommitTrueupLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitTrueupLedgerEntryType = shared.CommitLedgerPostpaidCommitTrueupLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup = shared.CommitLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitManualLedgerEntry = shared.CommitLedgerPrepaidCommitManualLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitManualLedgerEntryType = shared.CommitLedgerPrepaidCommitManualLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual = shared.CommitLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitManualLedgerEntry = shared.CommitLedgerPostpaidCommitManualLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitManualLedgerEntryType = shared.CommitLedgerPostpaidCommitManualLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual = shared.CommitLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitExpirationLedgerEntry = shared.CommitLedgerPostpaidCommitExpirationLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitExpirationLedgerEntryType = shared.CommitLedgerPostpaidCommitExpirationLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration = shared.CommitLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration

// This is an alias to an internal type.
type CommitLedgerType = shared.CommitLedgerType

// This is an alias to an internal value.
const CommitLedgerTypePrepaidCommitSegmentStart = shared.CommitLedgerTypePrepaidCommitSegmentStart

// This is an alias to an internal value.
const CommitLedgerTypePrepaidCommitAutomatedInvoiceDeduction = shared.CommitLedgerTypePrepaidCommitAutomatedInvoiceDeduction

// This is an alias to an internal value.
const CommitLedgerTypePrepaidCommitRollover = shared.CommitLedgerTypePrepaidCommitRollover

// This is an alias to an internal value.
const CommitLedgerTypePrepaidCommitExpiration = shared.CommitLedgerTypePrepaidCommitExpiration

// This is an alias to an internal value.
const CommitLedgerTypePrepaidCommitCanceled = shared.CommitLedgerTypePrepaidCommitCanceled

// This is an alias to an internal value.
const CommitLedgerTypePrepaidCommitCredited = shared.CommitLedgerTypePrepaidCommitCredited

// This is an alias to an internal value.
const CommitLedgerTypePrepaidCommitSeatBasedAdjustment = shared.CommitLedgerTypePrepaidCommitSeatBasedAdjustment

// This is an alias to an internal value.
const CommitLedgerTypePostpaidCommitInitialBalance = shared.CommitLedgerTypePostpaidCommitInitialBalance

// This is an alias to an internal value.
const CommitLedgerTypePostpaidCommitAutomatedInvoiceDeduction = shared.CommitLedgerTypePostpaidCommitAutomatedInvoiceDeduction

// This is an alias to an internal value.
const CommitLedgerTypePostpaidCommitRollover = shared.CommitLedgerTypePostpaidCommitRollover

// This is an alias to an internal value.
const CommitLedgerTypePostpaidCommitTrueup = shared.CommitLedgerTypePostpaidCommitTrueup

// This is an alias to an internal value.
const CommitLedgerTypePrepaidCommitManual = shared.CommitLedgerTypePrepaidCommitManual

// This is an alias to an internal value.
const CommitLedgerTypePostpaidCommitManual = shared.CommitLedgerTypePostpaidCommitManual

// This is an alias to an internal value.
const CommitLedgerTypePostpaidCommitExpiration = shared.CommitLedgerTypePostpaidCommitExpiration

// This is an alias to an internal type.
type CommitRateType = shared.CommitRateType

// This is an alias to an internal value.
const CommitRateTypeCommitRate = shared.CommitRateTypeCommitRate

// This is an alias to an internal value.
const CommitRateTypeListRate = shared.CommitRateTypeListRate

// This is an alias to an internal type.
type CommitRolledOverFrom = shared.CommitRolledOverFrom

// This is an alias to an internal type.
type CommitSpecifier = shared.CommitSpecifier

// This is an alias to an internal type.
type Contract = shared.Contract

// This is an alias to an internal type.
type ContractAmendment = shared.ContractAmendment

// This is an alias to an internal type.
type ContractAmendmentsResellerRoyalty = shared.ContractAmendmentsResellerRoyalty

// This is an alias to an internal type.
type ContractAmendmentsResellerRoyaltiesResellerType = shared.ContractAmendmentsResellerRoyaltiesResellerType

// This is an alias to an internal value.
const ContractAmendmentsResellerRoyaltiesResellerTypeAws = shared.ContractAmendmentsResellerRoyaltiesResellerTypeAws

// This is an alias to an internal value.
const ContractAmendmentsResellerRoyaltiesResellerTypeAwsProService = shared.ContractAmendmentsResellerRoyaltiesResellerTypeAwsProService

// This is an alias to an internal value.
const ContractAmendmentsResellerRoyaltiesResellerTypeGcp = shared.ContractAmendmentsResellerRoyaltiesResellerTypeGcp

// This is an alias to an internal value.
const ContractAmendmentsResellerRoyaltiesResellerTypeGcpProService = shared.ContractAmendmentsResellerRoyaltiesResellerTypeGcpProService

// The billing provider configuration associated with a contract.
//
// This is an alias to an internal type.
type ContractCustomerBillingProviderConfiguration = shared.ContractCustomerBillingProviderConfiguration

// This is an alias to an internal type.
type ContractCustomerBillingProviderConfigurationBillingProvider = shared.ContractCustomerBillingProviderConfigurationBillingProvider

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationBillingProviderAwsMarketplace = shared.ContractCustomerBillingProviderConfigurationBillingProviderAwsMarketplace

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationBillingProviderStripe = shared.ContractCustomerBillingProviderConfigurationBillingProviderStripe

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationBillingProviderNetsuite = shared.ContractCustomerBillingProviderConfigurationBillingProviderNetsuite

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationBillingProviderCustom = shared.ContractCustomerBillingProviderConfigurationBillingProviderCustom

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationBillingProviderAzureMarketplace = shared.ContractCustomerBillingProviderConfigurationBillingProviderAzureMarketplace

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline = shared.ContractCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationBillingProviderWorkday = shared.ContractCustomerBillingProviderConfigurationBillingProviderWorkday

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationBillingProviderGcpMarketplace = shared.ContractCustomerBillingProviderConfigurationBillingProviderGcpMarketplace

// This is an alias to an internal type.
type ContractCustomerBillingProviderConfigurationDeliveryMethod = shared.ContractCustomerBillingProviderConfigurationDeliveryMethod

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider = shared.ContractCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationDeliveryMethodAwsSqs = shared.ContractCustomerBillingProviderConfigurationDeliveryMethodAwsSqs

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationDeliveryMethodTackle = shared.ContractCustomerBillingProviderConfigurationDeliveryMethodTackle

// This is an alias to an internal value.
const ContractCustomerBillingProviderConfigurationDeliveryMethodAwsSns = shared.ContractCustomerBillingProviderConfigurationDeliveryMethodAwsSns

// This is an alias to an internal type.
type ContractPrepaidBalanceThresholdConfiguration = shared.ContractPrepaidBalanceThresholdConfiguration

// This is an alias to an internal type.
type ContractPrepaidBalanceThresholdConfigurationCommit = shared.ContractPrepaidBalanceThresholdConfigurationCommit

// This is an alias to an internal type.
type ContractPrepaidBalanceThresholdConfigurationCommitSpecifier = shared.ContractPrepaidBalanceThresholdConfigurationCommitSpecifier

// This is an alias to an internal type.
type ContractPrepaidBalanceThresholdConfigurationPaymentGateConfig = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfig

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
//
// This is an alias to an internal type.
type ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType

// This is an alias to an internal value.
const ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone

// This is an alias to an internal value.
const ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe

// This is an alias to an internal value.
const ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig

// Only applicable if using STRIPE as your payment gate type.
//
// This is an alias to an internal type.
type ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig

// If left blank, will default to INVOICE
//
// This is an alias to an internal type.
type ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType

// This is an alias to an internal value.
const ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice

// This is an alias to an internal value.
const ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
//
// This is an alias to an internal type.
type ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType

// This is an alias to an internal value.
const ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone

// This is an alias to an internal value.
const ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe

// This is an alias to an internal value.
const ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok

// This is an alias to an internal value.
const ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated = shared.ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
//
// This is an alias to an internal type.
type ContractScheduledChargesOnUsageInvoices = shared.ContractScheduledChargesOnUsageInvoices

// This is an alias to an internal value.
const ContractScheduledChargesOnUsageInvoicesAll = shared.ContractScheduledChargesOnUsageInvoicesAll

// This is an alias to an internal type.
type ContractSpendThresholdConfiguration = shared.ContractSpendThresholdConfiguration

// This is an alias to an internal type.
type ContractSpendThresholdConfigurationCommit = shared.ContractSpendThresholdConfigurationCommit

// This is an alias to an internal type.
type ContractSpendThresholdConfigurationPaymentGateConfig = shared.ContractSpendThresholdConfigurationPaymentGateConfig

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
//
// This is an alias to an internal type.
type ContractSpendThresholdConfigurationPaymentGateConfigPaymentGateType = shared.ContractSpendThresholdConfigurationPaymentGateConfigPaymentGateType

// This is an alias to an internal value.
const ContractSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone = shared.ContractSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone

// This is an alias to an internal value.
const ContractSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe = shared.ContractSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe

// This is an alias to an internal value.
const ContractSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal = shared.ContractSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type ContractSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig = shared.ContractSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig

// Only applicable if using STRIPE as your payment gate type.
//
// This is an alias to an internal type.
type ContractSpendThresholdConfigurationPaymentGateConfigStripeConfig = shared.ContractSpendThresholdConfigurationPaymentGateConfigStripeConfig

// If left blank, will default to INVOICE
//
// This is an alias to an internal type.
type ContractSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = shared.ContractSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType

// This is an alias to an internal value.
const ContractSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice = shared.ContractSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice

// This is an alias to an internal value.
const ContractSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent = shared.ContractSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
//
// This is an alias to an internal type.
type ContractSpendThresholdConfigurationPaymentGateConfigTaxType = shared.ContractSpendThresholdConfigurationPaymentGateConfigTaxType

// This is an alias to an internal value.
const ContractSpendThresholdConfigurationPaymentGateConfigTaxTypeNone = shared.ContractSpendThresholdConfigurationPaymentGateConfigTaxTypeNone

// This is an alias to an internal value.
const ContractSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe = shared.ContractSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe

// This is an alias to an internal value.
const ContractSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok = shared.ContractSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok

// This is an alias to an internal value.
const ContractSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated = shared.ContractSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated

// This is an alias to an internal type.
type ContractSubscription = shared.ContractSubscription

// This is an alias to an internal type.
type ContractSubscriptionsCollectionSchedule = shared.ContractSubscriptionsCollectionSchedule

// This is an alias to an internal value.
const ContractSubscriptionsCollectionScheduleAdvance = shared.ContractSubscriptionsCollectionScheduleAdvance

// This is an alias to an internal value.
const ContractSubscriptionsCollectionScheduleArrears = shared.ContractSubscriptionsCollectionScheduleArrears

// This is an alias to an internal type.
type ContractSubscriptionsProration = shared.ContractSubscriptionsProration

// This is an alias to an internal type.
type ContractSubscriptionsProrationInvoiceBehavior = shared.ContractSubscriptionsProrationInvoiceBehavior

// This is an alias to an internal value.
const ContractSubscriptionsProrationInvoiceBehaviorBillImmediately = shared.ContractSubscriptionsProrationInvoiceBehaviorBillImmediately

// This is an alias to an internal value.
const ContractSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate = shared.ContractSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate

// This is an alias to an internal type.
type ContractSubscriptionsQuantitySchedule = shared.ContractSubscriptionsQuantitySchedule

// This is an alias to an internal type.
type ContractSubscriptionsSubscriptionRate = shared.ContractSubscriptionsSubscriptionRate

// This is an alias to an internal type.
type ContractSubscriptionsSubscriptionRateBillingFrequency = shared.ContractSubscriptionsSubscriptionRateBillingFrequency

// This is an alias to an internal value.
const ContractSubscriptionsSubscriptionRateBillingFrequencyMonthly = shared.ContractSubscriptionsSubscriptionRateBillingFrequencyMonthly

// This is an alias to an internal value.
const ContractSubscriptionsSubscriptionRateBillingFrequencyQuarterly = shared.ContractSubscriptionsSubscriptionRateBillingFrequencyQuarterly

// This is an alias to an internal value.
const ContractSubscriptionsSubscriptionRateBillingFrequencyAnnual = shared.ContractSubscriptionsSubscriptionRateBillingFrequencyAnnual

// This is an alias to an internal value.
const ContractSubscriptionsSubscriptionRateBillingFrequencyWeekly = shared.ContractSubscriptionsSubscriptionRateBillingFrequencyWeekly

// This is an alias to an internal type.
type ContractSubscriptionsSubscriptionRateProduct = shared.ContractSubscriptionsSubscriptionRateProduct

// This is an alias to an internal type.
type ContractV2 = shared.ContractV2

// This is an alias to an internal type.
type ContractV2Commit = shared.ContractV2Commit

// This is an alias to an internal type.
type ContractV2CommitsProduct = shared.ContractV2CommitsProduct

// This is an alias to an internal type.
type ContractV2CommitsType = shared.ContractV2CommitsType

// This is an alias to an internal value.
const ContractV2CommitsTypePrepaid = shared.ContractV2CommitsTypePrepaid

// This is an alias to an internal value.
const ContractV2CommitsTypePostpaid = shared.ContractV2CommitsTypePostpaid

// This is an alias to an internal type.
type ContractV2CommitsContract = shared.ContractV2CommitsContract

// Optional configuration for commit hierarchy access control
//
// This is an alias to an internal type.
type ContractV2CommitsHierarchyConfiguration = shared.ContractV2CommitsHierarchyConfiguration

// This is an alias to an internal type.
type ContractV2CommitsHierarchyConfigurationChildAccess = shared.ContractV2CommitsHierarchyConfigurationChildAccess

// This is an alias to an internal type.
type ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll = shared.ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll

// This is an alias to an internal type.
type ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = shared.ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType

// This is an alias to an internal value.
const ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll = shared.ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll

// This is an alias to an internal type.
type ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone = shared.ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone

// This is an alias to an internal type.
type ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = shared.ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType

// This is an alias to an internal value.
const ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone = shared.ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone

// This is an alias to an internal type.
type ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs = shared.ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs

// This is an alias to an internal type.
type ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = shared.ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType

// This is an alias to an internal value.
const ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs = shared.ContractV2CommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs

// This is an alias to an internal type.
type ContractV2CommitsHierarchyConfigurationChildAccessType = shared.ContractV2CommitsHierarchyConfigurationChildAccessType

// This is an alias to an internal value.
const ContractV2CommitsHierarchyConfigurationChildAccessTypeAll = shared.ContractV2CommitsHierarchyConfigurationChildAccessTypeAll

// This is an alias to an internal value.
const ContractV2CommitsHierarchyConfigurationChildAccessTypeNone = shared.ContractV2CommitsHierarchyConfigurationChildAccessTypeNone

// This is an alias to an internal value.
const ContractV2CommitsHierarchyConfigurationChildAccessTypeContractIDs = shared.ContractV2CommitsHierarchyConfigurationChildAccessTypeContractIDs

// The contract that this commit will be billed on.
//
// This is an alias to an internal type.
type ContractV2CommitsInvoiceContract = shared.ContractV2CommitsInvoiceContract

// This is an alias to an internal type.
type ContractV2CommitsLedger = shared.ContractV2CommitsLedger

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntry = shared.ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = shared.ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart = shared.ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry = shared.ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = shared.ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction = shared.ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntry = shared.ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryType = shared.ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover = shared.ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntry = shared.ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryType = shared.ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration = shared.ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntry = shared.ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryType = shared.ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled = shared.ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntry = shared.ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryType = shared.ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited = shared.ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry = shared.ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType = shared.ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryTypePrepaidCommitSeatBasedAdjustment = shared.ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryTypePrepaidCommitSeatBasedAdjustment

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntry = shared.ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = shared.ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance = shared.ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry = shared.ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = shared.ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction = shared.ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntry = shared.ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryType = shared.ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover = shared.ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntry = shared.ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryType = shared.ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup = shared.ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitManualLedgerEntry = shared.ContractV2CommitsLedgerPrepaidCommitManualLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPrepaidCommitManualLedgerEntryType = shared.ContractV2CommitsLedgerPrepaidCommitManualLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual = shared.ContractV2CommitsLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitManualLedgerEntry = shared.ContractV2CommitsLedgerPostpaidCommitManualLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitManualLedgerEntryType = shared.ContractV2CommitsLedgerPostpaidCommitManualLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual = shared.ContractV2CommitsLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntry = shared.ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryType = shared.ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryType

// This is an alias to an internal value.
const ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration = shared.ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration

// This is an alias to an internal type.
type ContractV2CommitsLedgerType = shared.ContractV2CommitsLedgerType

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePrepaidCommitSegmentStart = shared.ContractV2CommitsLedgerTypePrepaidCommitSegmentStart

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePrepaidCommitAutomatedInvoiceDeduction = shared.ContractV2CommitsLedgerTypePrepaidCommitAutomatedInvoiceDeduction

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePrepaidCommitRollover = shared.ContractV2CommitsLedgerTypePrepaidCommitRollover

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePrepaidCommitExpiration = shared.ContractV2CommitsLedgerTypePrepaidCommitExpiration

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePrepaidCommitCanceled = shared.ContractV2CommitsLedgerTypePrepaidCommitCanceled

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePrepaidCommitCredited = shared.ContractV2CommitsLedgerTypePrepaidCommitCredited

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePrepaidCommitSeatBasedAdjustment = shared.ContractV2CommitsLedgerTypePrepaidCommitSeatBasedAdjustment

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePostpaidCommitInitialBalance = shared.ContractV2CommitsLedgerTypePostpaidCommitInitialBalance

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePostpaidCommitAutomatedInvoiceDeduction = shared.ContractV2CommitsLedgerTypePostpaidCommitAutomatedInvoiceDeduction

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePostpaidCommitRollover = shared.ContractV2CommitsLedgerTypePostpaidCommitRollover

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePostpaidCommitTrueup = shared.ContractV2CommitsLedgerTypePostpaidCommitTrueup

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePrepaidCommitManual = shared.ContractV2CommitsLedgerTypePrepaidCommitManual

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePostpaidCommitManual = shared.ContractV2CommitsLedgerTypePostpaidCommitManual

// This is an alias to an internal value.
const ContractV2CommitsLedgerTypePostpaidCommitExpiration = shared.ContractV2CommitsLedgerTypePostpaidCommitExpiration

// This is an alias to an internal type.
type ContractV2CommitsRateType = shared.ContractV2CommitsRateType

// This is an alias to an internal value.
const ContractV2CommitsRateTypeCommitRate = shared.ContractV2CommitsRateTypeCommitRate

// This is an alias to an internal value.
const ContractV2CommitsRateTypeListRate = shared.ContractV2CommitsRateTypeListRate

// This is an alias to an internal type.
type ContractV2CommitsRolledOverFrom = shared.ContractV2CommitsRolledOverFrom

// This is an alias to an internal type.
type ContractV2CommitsSpecifier = shared.ContractV2CommitsSpecifier

// This is an alias to an internal type.
type ContractV2Override = shared.ContractV2Override

// This is an alias to an internal type.
type ContractV2OverridesOverrideSpecifier = shared.ContractV2OverridesOverrideSpecifier

// This is an alias to an internal type.
type ContractV2OverridesOverrideSpecifiersBillingFrequency = shared.ContractV2OverridesOverrideSpecifiersBillingFrequency

// This is an alias to an internal value.
const ContractV2OverridesOverrideSpecifiersBillingFrequencyMonthly = shared.ContractV2OverridesOverrideSpecifiersBillingFrequencyMonthly

// This is an alias to an internal value.
const ContractV2OverridesOverrideSpecifiersBillingFrequencyQuarterly = shared.ContractV2OverridesOverrideSpecifiersBillingFrequencyQuarterly

// This is an alias to an internal value.
const ContractV2OverridesOverrideSpecifiersBillingFrequencyAnnual = shared.ContractV2OverridesOverrideSpecifiersBillingFrequencyAnnual

// This is an alias to an internal value.
const ContractV2OverridesOverrideSpecifiersBillingFrequencyWeekly = shared.ContractV2OverridesOverrideSpecifiersBillingFrequencyWeekly

// This is an alias to an internal type.
type ContractV2OverridesOverrideTier = shared.ContractV2OverridesOverrideTier

// This is an alias to an internal type.
type ContractV2OverridesOverwriteRate = shared.ContractV2OverridesOverwriteRate

// This is an alias to an internal type.
type ContractV2OverridesOverwriteRateRateType = shared.ContractV2OverridesOverwriteRateRateType

// This is an alias to an internal value.
const ContractV2OverridesOverwriteRateRateTypeFlat = shared.ContractV2OverridesOverwriteRateRateTypeFlat

// This is an alias to an internal value.
const ContractV2OverridesOverwriteRateRateTypePercentage = shared.ContractV2OverridesOverwriteRateRateTypePercentage

// This is an alias to an internal value.
const ContractV2OverridesOverwriteRateRateTypeSubscription = shared.ContractV2OverridesOverwriteRateRateTypeSubscription

// This is an alias to an internal value.
const ContractV2OverridesOverwriteRateRateTypeTiered = shared.ContractV2OverridesOverwriteRateRateTypeTiered

// This is an alias to an internal value.
const ContractV2OverridesOverwriteRateRateTypeCustom = shared.ContractV2OverridesOverwriteRateRateTypeCustom

// This is an alias to an internal type.
type ContractV2OverridesProduct = shared.ContractV2OverridesProduct

// This is an alias to an internal type.
type ContractV2OverridesTarget = shared.ContractV2OverridesTarget

// This is an alias to an internal value.
const ContractV2OverridesTargetCommitRate = shared.ContractV2OverridesTargetCommitRate

// This is an alias to an internal value.
const ContractV2OverridesTargetListRate = shared.ContractV2OverridesTargetListRate

// This is an alias to an internal type.
type ContractV2OverridesType = shared.ContractV2OverridesType

// This is an alias to an internal value.
const ContractV2OverridesTypeOverwrite = shared.ContractV2OverridesTypeOverwrite

// This is an alias to an internal value.
const ContractV2OverridesTypeMultiplier = shared.ContractV2OverridesTypeMultiplier

// This is an alias to an internal value.
const ContractV2OverridesTypeTiered = shared.ContractV2OverridesTypeTiered

// This is an alias to an internal type.
type ContractV2Transition = shared.ContractV2Transition

// This is an alias to an internal type.
type ContractV2TransitionsType = shared.ContractV2TransitionsType

// This is an alias to an internal value.
const ContractV2TransitionsTypeSupersede = shared.ContractV2TransitionsTypeSupersede

// This is an alias to an internal value.
const ContractV2TransitionsTypeRenewal = shared.ContractV2TransitionsTypeRenewal

// This is an alias to an internal type.
type ContractV2UsageFilter = shared.ContractV2UsageFilter

// This is an alias to an internal type.
type ContractV2UsageStatementSchedule = shared.ContractV2UsageStatementSchedule

// This is an alias to an internal type.
type ContractV2UsageStatementScheduleFrequency = shared.ContractV2UsageStatementScheduleFrequency

// This is an alias to an internal value.
const ContractV2UsageStatementScheduleFrequencyMonthly = shared.ContractV2UsageStatementScheduleFrequencyMonthly

// This is an alias to an internal value.
const ContractV2UsageStatementScheduleFrequencyQuarterly = shared.ContractV2UsageStatementScheduleFrequencyQuarterly

// This is an alias to an internal value.
const ContractV2UsageStatementScheduleFrequencyAnnual = shared.ContractV2UsageStatementScheduleFrequencyAnnual

// This is an alias to an internal value.
const ContractV2UsageStatementScheduleFrequencyWeekly = shared.ContractV2UsageStatementScheduleFrequencyWeekly

// This is an alias to an internal type.
type ContractV2Credit = shared.ContractV2Credit

// This is an alias to an internal type.
type ContractV2CreditsProduct = shared.ContractV2CreditsProduct

// This is an alias to an internal type.
type ContractV2CreditsType = shared.ContractV2CreditsType

// This is an alias to an internal value.
const ContractV2CreditsTypeCredit = shared.ContractV2CreditsTypeCredit

// This is an alias to an internal type.
type ContractV2CreditsContract = shared.ContractV2CreditsContract

// Optional configuration for credit hierarchy access control
//
// This is an alias to an internal type.
type ContractV2CreditsHierarchyConfiguration = shared.ContractV2CreditsHierarchyConfiguration

// This is an alias to an internal type.
type ContractV2CreditsHierarchyConfigurationChildAccess = shared.ContractV2CreditsHierarchyConfigurationChildAccess

// This is an alias to an internal type.
type ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll = shared.ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll

// This is an alias to an internal type.
type ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = shared.ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType

// This is an alias to an internal value.
const ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll = shared.ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll

// This is an alias to an internal type.
type ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone = shared.ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone

// This is an alias to an internal type.
type ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = shared.ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType

// This is an alias to an internal value.
const ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone = shared.ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone

// This is an alias to an internal type.
type ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs = shared.ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs

// This is an alias to an internal type.
type ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = shared.ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType

// This is an alias to an internal value.
const ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs = shared.ContractV2CreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs

// This is an alias to an internal type.
type ContractV2CreditsHierarchyConfigurationChildAccessType = shared.ContractV2CreditsHierarchyConfigurationChildAccessType

// This is an alias to an internal value.
const ContractV2CreditsHierarchyConfigurationChildAccessTypeAll = shared.ContractV2CreditsHierarchyConfigurationChildAccessTypeAll

// This is an alias to an internal value.
const ContractV2CreditsHierarchyConfigurationChildAccessTypeNone = shared.ContractV2CreditsHierarchyConfigurationChildAccessTypeNone

// This is an alias to an internal value.
const ContractV2CreditsHierarchyConfigurationChildAccessTypeContractIDs = shared.ContractV2CreditsHierarchyConfigurationChildAccessTypeContractIDs

// This is an alias to an internal type.
type ContractV2CreditsLedger = shared.ContractV2CreditsLedger

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditSegmentStartLedgerEntry = shared.ContractV2CreditsLedgerCreditSegmentStartLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditSegmentStartLedgerEntryType = shared.ContractV2CreditsLedgerCreditSegmentStartLedgerEntryType

// This is an alias to an internal value.
const ContractV2CreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart = shared.ContractV2CreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry = shared.ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = shared.ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType

// This is an alias to an internal value.
const ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction = shared.ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditExpirationLedgerEntry = shared.ContractV2CreditsLedgerCreditExpirationLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditExpirationLedgerEntryType = shared.ContractV2CreditsLedgerCreditExpirationLedgerEntryType

// This is an alias to an internal value.
const ContractV2CreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration = shared.ContractV2CreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditCanceledLedgerEntry = shared.ContractV2CreditsLedgerCreditCanceledLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditCanceledLedgerEntryType = shared.ContractV2CreditsLedgerCreditCanceledLedgerEntryType

// This is an alias to an internal value.
const ContractV2CreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled = shared.ContractV2CreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditCreditedLedgerEntry = shared.ContractV2CreditsLedgerCreditCreditedLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditCreditedLedgerEntryType = shared.ContractV2CreditsLedgerCreditCreditedLedgerEntryType

// This is an alias to an internal value.
const ContractV2CreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited = shared.ContractV2CreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditManualLedgerEntry = shared.ContractV2CreditsLedgerCreditManualLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditManualLedgerEntryType = shared.ContractV2CreditsLedgerCreditManualLedgerEntryType

// This is an alias to an internal value.
const ContractV2CreditsLedgerCreditManualLedgerEntryTypeCreditManual = shared.ContractV2CreditsLedgerCreditManualLedgerEntryTypeCreditManual

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntry = shared.ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryType = shared.ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryType

// This is an alias to an internal value.
const ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryTypeCreditSeatBasedAdjustment = shared.ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryTypeCreditSeatBasedAdjustment

// This is an alias to an internal type.
type ContractV2CreditsLedgerType = shared.ContractV2CreditsLedgerType

// This is an alias to an internal value.
const ContractV2CreditsLedgerTypeCreditSegmentStart = shared.ContractV2CreditsLedgerTypeCreditSegmentStart

// This is an alias to an internal value.
const ContractV2CreditsLedgerTypeCreditAutomatedInvoiceDeduction = shared.ContractV2CreditsLedgerTypeCreditAutomatedInvoiceDeduction

// This is an alias to an internal value.
const ContractV2CreditsLedgerTypeCreditExpiration = shared.ContractV2CreditsLedgerTypeCreditExpiration

// This is an alias to an internal value.
const ContractV2CreditsLedgerTypeCreditCanceled = shared.ContractV2CreditsLedgerTypeCreditCanceled

// This is an alias to an internal value.
const ContractV2CreditsLedgerTypeCreditCredited = shared.ContractV2CreditsLedgerTypeCreditCredited

// This is an alias to an internal value.
const ContractV2CreditsLedgerTypeCreditManual = shared.ContractV2CreditsLedgerTypeCreditManual

// This is an alias to an internal value.
const ContractV2CreditsLedgerTypeCreditSeatBasedAdjustment = shared.ContractV2CreditsLedgerTypeCreditSeatBasedAdjustment

// This is an alias to an internal type.
type ContractV2CreditsSpecifier = shared.ContractV2CreditsSpecifier

// This field's availability is dependent on your client's configuration.
//
// This is an alias to an internal type.
type ContractV2CustomerBillingProviderConfiguration = shared.ContractV2CustomerBillingProviderConfiguration

// This is an alias to an internal type.
type ContractV2CustomerBillingProviderConfigurationBillingProvider = shared.ContractV2CustomerBillingProviderConfigurationBillingProvider

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationBillingProviderAwsMarketplace = shared.ContractV2CustomerBillingProviderConfigurationBillingProviderAwsMarketplace

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationBillingProviderStripe = shared.ContractV2CustomerBillingProviderConfigurationBillingProviderStripe

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationBillingProviderNetsuite = shared.ContractV2CustomerBillingProviderConfigurationBillingProviderNetsuite

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationBillingProviderCustom = shared.ContractV2CustomerBillingProviderConfigurationBillingProviderCustom

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationBillingProviderAzureMarketplace = shared.ContractV2CustomerBillingProviderConfigurationBillingProviderAzureMarketplace

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationBillingProviderQuickbooksOnline = shared.ContractV2CustomerBillingProviderConfigurationBillingProviderQuickbooksOnline

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationBillingProviderWorkday = shared.ContractV2CustomerBillingProviderConfigurationBillingProviderWorkday

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationBillingProviderGcpMarketplace = shared.ContractV2CustomerBillingProviderConfigurationBillingProviderGcpMarketplace

// This is an alias to an internal type.
type ContractV2CustomerBillingProviderConfigurationDeliveryMethod = shared.ContractV2CustomerBillingProviderConfigurationDeliveryMethod

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider = shared.ContractV2CustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationDeliveryMethodAwsSqs = shared.ContractV2CustomerBillingProviderConfigurationDeliveryMethodAwsSqs

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationDeliveryMethodTackle = shared.ContractV2CustomerBillingProviderConfigurationDeliveryMethodTackle

// This is an alias to an internal value.
const ContractV2CustomerBillingProviderConfigurationDeliveryMethodAwsSns = shared.ContractV2CustomerBillingProviderConfigurationDeliveryMethodAwsSns

// Indicates whether there are more items than the limit for this endpoint. Use the
// respective list endpoints to get the full lists.
//
// This is an alias to an internal type.
type ContractV2HasMore = shared.ContractV2HasMore

// Either a **parent** configuration with a list of children or a **child**
// configuration with a single parent.
//
// This is an alias to an internal type.
type ContractV2HierarchyConfiguration = shared.ContractV2HierarchyConfiguration

// This is an alias to an internal type.
type ContractV2HierarchyConfigurationParentHierarchyConfiguration = shared.ContractV2HierarchyConfigurationParentHierarchyConfiguration

// This is an alias to an internal type.
type ContractV2HierarchyConfigurationParentHierarchyConfigurationChild = shared.ContractV2HierarchyConfigurationParentHierarchyConfigurationChild

// This is an alias to an internal type.
type ContractV2HierarchyConfigurationChildHierarchyConfiguration = shared.ContractV2HierarchyConfigurationChildHierarchyConfiguration

// The single parent contract/customer for this child.
//
// This is an alias to an internal type.
type ContractV2HierarchyConfigurationChildHierarchyConfigurationParent = shared.ContractV2HierarchyConfigurationChildHierarchyConfigurationParent

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
// prices automatically. EXPLICIT prioritization requires specifying priorities for
// each multiplier; the one with the lowest priority value will be prioritized
// first.
//
// This is an alias to an internal type.
type ContractV2MultiplierOverridePrioritization = shared.ContractV2MultiplierOverridePrioritization

// This is an alias to an internal value.
const ContractV2MultiplierOverridePrioritizationLowestMultiplier = shared.ContractV2MultiplierOverridePrioritizationLowestMultiplier

// This is an alias to an internal value.
const ContractV2MultiplierOverridePrioritizationExplicit = shared.ContractV2MultiplierOverridePrioritizationExplicit

// This is an alias to an internal type.
type ContractV2PrepaidBalanceThresholdConfiguration = shared.ContractV2PrepaidBalanceThresholdConfiguration

// This is an alias to an internal type.
type ContractV2PrepaidBalanceThresholdConfigurationCommit = shared.ContractV2PrepaidBalanceThresholdConfigurationCommit

// This is an alias to an internal type.
type ContractV2PrepaidBalanceThresholdConfigurationCommitSpecifier = shared.ContractV2PrepaidBalanceThresholdConfigurationCommitSpecifier

// This is an alias to an internal type.
type ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfig = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfig

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
//
// This is an alias to an internal type.
type ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType

// This is an alias to an internal value.
const ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone

// This is an alias to an internal value.
const ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe

// This is an alias to an internal value.
const ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig

// Only applicable if using STRIPE as your payment gateway type.
//
// This is an alias to an internal type.
type ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig

// If left blank, will default to INVOICE
//
// This is an alias to an internal type.
type ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType

// This is an alias to an internal value.
const ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice

// This is an alias to an internal value.
const ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
//
// This is an alias to an internal type.
type ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType

// This is an alias to an internal value.
const ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone

// This is an alias to an internal value.
const ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe

// This is an alias to an internal value.
const ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok

// This is an alias to an internal value.
const ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated = shared.ContractV2PrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated

// This is an alias to an internal type.
type ContractV2RecurringCommit = shared.ContractV2RecurringCommit

// The amount of commit to grant.
//
// This is an alias to an internal type.
type ContractV2RecurringCommitsAccessAmount = shared.ContractV2RecurringCommitsAccessAmount

// The amount of time the created commits will be valid for
//
// This is an alias to an internal type.
type ContractV2RecurringCommitsCommitDuration = shared.ContractV2RecurringCommitsCommitDuration

// This is an alias to an internal type.
type ContractV2RecurringCommitsCommitDurationUnit = shared.ContractV2RecurringCommitsCommitDurationUnit

// This is an alias to an internal value.
const ContractV2RecurringCommitsCommitDurationUnitPeriods = shared.ContractV2RecurringCommitsCommitDurationUnitPeriods

// This is an alias to an internal type.
type ContractV2RecurringCommitsProduct = shared.ContractV2RecurringCommitsProduct

// Whether the created commits will use the commit rate or list rate
//
// This is an alias to an internal type.
type ContractV2RecurringCommitsRateType = shared.ContractV2RecurringCommitsRateType

// This is an alias to an internal value.
const ContractV2RecurringCommitsRateTypeCommitRate = shared.ContractV2RecurringCommitsRateTypeCommitRate

// This is an alias to an internal value.
const ContractV2RecurringCommitsRateTypeListRate = shared.ContractV2RecurringCommitsRateTypeListRate

// This is an alias to an internal type.
type ContractV2RecurringCommitsContract = shared.ContractV2RecurringCommitsContract

// Optional configuration for recurring credit hierarchy access control
//
// This is an alias to an internal type.
type ContractV2RecurringCommitsHierarchyConfiguration = shared.ContractV2RecurringCommitsHierarchyConfiguration

// This is an alias to an internal type.
type ContractV2RecurringCommitsHierarchyConfigurationChildAccess = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccess

// This is an alias to an internal type.
type ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll

// This is an alias to an internal type.
type ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType

// This is an alias to an internal value.
const ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll

// This is an alias to an internal type.
type ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone

// This is an alias to an internal type.
type ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType

// This is an alias to an internal value.
const ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone

// This is an alias to an internal type.
type ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs

// This is an alias to an internal type.
type ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType

// This is an alias to an internal value.
const ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs

// This is an alias to an internal type.
type ContractV2RecurringCommitsHierarchyConfigurationChildAccessType = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessType

// This is an alias to an internal value.
const ContractV2RecurringCommitsHierarchyConfigurationChildAccessTypeAll = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessTypeAll

// This is an alias to an internal value.
const ContractV2RecurringCommitsHierarchyConfigurationChildAccessTypeNone = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessTypeNone

// This is an alias to an internal value.
const ContractV2RecurringCommitsHierarchyConfigurationChildAccessTypeContractIDs = shared.ContractV2RecurringCommitsHierarchyConfigurationChildAccessTypeContractIDs

// The amount the customer should be billed for the commit. Not required.
//
// This is an alias to an internal type.
type ContractV2RecurringCommitsInvoiceAmount = shared.ContractV2RecurringCommitsInvoiceAmount

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
//
// This is an alias to an internal type.
type ContractV2RecurringCommitsProration = shared.ContractV2RecurringCommitsProration

// This is an alias to an internal value.
const ContractV2RecurringCommitsProrationNone = shared.ContractV2RecurringCommitsProrationNone

// This is an alias to an internal value.
const ContractV2RecurringCommitsProrationFirst = shared.ContractV2RecurringCommitsProrationFirst

// This is an alias to an internal value.
const ContractV2RecurringCommitsProrationLast = shared.ContractV2RecurringCommitsProrationLast

// This is an alias to an internal value.
const ContractV2RecurringCommitsProrationFirstAndLast = shared.ContractV2RecurringCommitsProrationFirstAndLast

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's starting_at rather than the usage
// invoice dates.
//
// This is an alias to an internal type.
type ContractV2RecurringCommitsRecurrenceFrequency = shared.ContractV2RecurringCommitsRecurrenceFrequency

// This is an alias to an internal value.
const ContractV2RecurringCommitsRecurrenceFrequencyMonthly = shared.ContractV2RecurringCommitsRecurrenceFrequencyMonthly

// This is an alias to an internal value.
const ContractV2RecurringCommitsRecurrenceFrequencyQuarterly = shared.ContractV2RecurringCommitsRecurrenceFrequencyQuarterly

// This is an alias to an internal value.
const ContractV2RecurringCommitsRecurrenceFrequencyAnnual = shared.ContractV2RecurringCommitsRecurrenceFrequencyAnnual

// This is an alias to an internal value.
const ContractV2RecurringCommitsRecurrenceFrequencyWeekly = shared.ContractV2RecurringCommitsRecurrenceFrequencyWeekly

// This is an alias to an internal type.
type ContractV2RecurringCommitsSpecifier = shared.ContractV2RecurringCommitsSpecifier

// Attach a subscription to the recurring commit/credit.
//
// This is an alias to an internal type.
type ContractV2RecurringCommitsSubscriptionConfig = shared.ContractV2RecurringCommitsSubscriptionConfig

// This is an alias to an internal type.
type ContractV2RecurringCommitsSubscriptionConfigAllocation = shared.ContractV2RecurringCommitsSubscriptionConfigAllocation

// This is an alias to an internal value.
const ContractV2RecurringCommitsSubscriptionConfigAllocationIndividual = shared.ContractV2RecurringCommitsSubscriptionConfigAllocationIndividual

// This is an alias to an internal value.
const ContractV2RecurringCommitsSubscriptionConfigAllocationPooled = shared.ContractV2RecurringCommitsSubscriptionConfigAllocationPooled

// This is an alias to an internal type.
type ContractV2RecurringCommitsSubscriptionConfigApplySeatIncreaseConfig = shared.ContractV2RecurringCommitsSubscriptionConfigApplySeatIncreaseConfig

// This is an alias to an internal type.
type ContractV2RecurringCredit = shared.ContractV2RecurringCredit

// The amount of commit to grant.
//
// This is an alias to an internal type.
type ContractV2RecurringCreditsAccessAmount = shared.ContractV2RecurringCreditsAccessAmount

// The amount of time the created commits will be valid for
//
// This is an alias to an internal type.
type ContractV2RecurringCreditsCommitDuration = shared.ContractV2RecurringCreditsCommitDuration

// This is an alias to an internal type.
type ContractV2RecurringCreditsCommitDurationUnit = shared.ContractV2RecurringCreditsCommitDurationUnit

// This is an alias to an internal value.
const ContractV2RecurringCreditsCommitDurationUnitPeriods = shared.ContractV2RecurringCreditsCommitDurationUnitPeriods

// This is an alias to an internal type.
type ContractV2RecurringCreditsProduct = shared.ContractV2RecurringCreditsProduct

// Whether the created commits will use the commit rate or list rate
//
// This is an alias to an internal type.
type ContractV2RecurringCreditsRateType = shared.ContractV2RecurringCreditsRateType

// This is an alias to an internal value.
const ContractV2RecurringCreditsRateTypeCommitRate = shared.ContractV2RecurringCreditsRateTypeCommitRate

// This is an alias to an internal value.
const ContractV2RecurringCreditsRateTypeListRate = shared.ContractV2RecurringCreditsRateTypeListRate

// This is an alias to an internal type.
type ContractV2RecurringCreditsContract = shared.ContractV2RecurringCreditsContract

// Optional configuration for recurring credit hierarchy access control
//
// This is an alias to an internal type.
type ContractV2RecurringCreditsHierarchyConfiguration = shared.ContractV2RecurringCreditsHierarchyConfiguration

// This is an alias to an internal type.
type ContractV2RecurringCreditsHierarchyConfigurationChildAccess = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccess

// This is an alias to an internal type.
type ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll

// This is an alias to an internal type.
type ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType

// This is an alias to an internal value.
const ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll

// This is an alias to an internal type.
type ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone

// This is an alias to an internal type.
type ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType

// This is an alias to an internal value.
const ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone

// This is an alias to an internal type.
type ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs

// This is an alias to an internal type.
type ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType

// This is an alias to an internal value.
const ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs

// This is an alias to an internal type.
type ContractV2RecurringCreditsHierarchyConfigurationChildAccessType = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessType

// This is an alias to an internal value.
const ContractV2RecurringCreditsHierarchyConfigurationChildAccessTypeAll = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessTypeAll

// This is an alias to an internal value.
const ContractV2RecurringCreditsHierarchyConfigurationChildAccessTypeNone = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessTypeNone

// This is an alias to an internal value.
const ContractV2RecurringCreditsHierarchyConfigurationChildAccessTypeContractIDs = shared.ContractV2RecurringCreditsHierarchyConfigurationChildAccessTypeContractIDs

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
//
// This is an alias to an internal type.
type ContractV2RecurringCreditsProration = shared.ContractV2RecurringCreditsProration

// This is an alias to an internal value.
const ContractV2RecurringCreditsProrationNone = shared.ContractV2RecurringCreditsProrationNone

// This is an alias to an internal value.
const ContractV2RecurringCreditsProrationFirst = shared.ContractV2RecurringCreditsProrationFirst

// This is an alias to an internal value.
const ContractV2RecurringCreditsProrationLast = shared.ContractV2RecurringCreditsProrationLast

// This is an alias to an internal value.
const ContractV2RecurringCreditsProrationFirstAndLast = shared.ContractV2RecurringCreditsProrationFirstAndLast

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's starting_at rather than the usage
// invoice dates.
//
// This is an alias to an internal type.
type ContractV2RecurringCreditsRecurrenceFrequency = shared.ContractV2RecurringCreditsRecurrenceFrequency

// This is an alias to an internal value.
const ContractV2RecurringCreditsRecurrenceFrequencyMonthly = shared.ContractV2RecurringCreditsRecurrenceFrequencyMonthly

// This is an alias to an internal value.
const ContractV2RecurringCreditsRecurrenceFrequencyQuarterly = shared.ContractV2RecurringCreditsRecurrenceFrequencyQuarterly

// This is an alias to an internal value.
const ContractV2RecurringCreditsRecurrenceFrequencyAnnual = shared.ContractV2RecurringCreditsRecurrenceFrequencyAnnual

// This is an alias to an internal value.
const ContractV2RecurringCreditsRecurrenceFrequencyWeekly = shared.ContractV2RecurringCreditsRecurrenceFrequencyWeekly

// This is an alias to an internal type.
type ContractV2RecurringCreditsSpecifier = shared.ContractV2RecurringCreditsSpecifier

// Attach a subscription to the recurring commit/credit.
//
// This is an alias to an internal type.
type ContractV2RecurringCreditsSubscriptionConfig = shared.ContractV2RecurringCreditsSubscriptionConfig

// This is an alias to an internal type.
type ContractV2RecurringCreditsSubscriptionConfigAllocation = shared.ContractV2RecurringCreditsSubscriptionConfigAllocation

// This is an alias to an internal value.
const ContractV2RecurringCreditsSubscriptionConfigAllocationIndividual = shared.ContractV2RecurringCreditsSubscriptionConfigAllocationIndividual

// This is an alias to an internal value.
const ContractV2RecurringCreditsSubscriptionConfigAllocationPooled = shared.ContractV2RecurringCreditsSubscriptionConfigAllocationPooled

// This is an alias to an internal type.
type ContractV2RecurringCreditsSubscriptionConfigApplySeatIncreaseConfig = shared.ContractV2RecurringCreditsSubscriptionConfigApplySeatIncreaseConfig

// This is an alias to an internal type.
type ContractV2ResellerRoyalty = shared.ContractV2ResellerRoyalty

// This is an alias to an internal type.
type ContractV2ResellerRoyaltiesResellerType = shared.ContractV2ResellerRoyaltiesResellerType

// This is an alias to an internal value.
const ContractV2ResellerRoyaltiesResellerTypeAws = shared.ContractV2ResellerRoyaltiesResellerTypeAws

// This is an alias to an internal value.
const ContractV2ResellerRoyaltiesResellerTypeAwsProService = shared.ContractV2ResellerRoyaltiesResellerTypeAwsProService

// This is an alias to an internal value.
const ContractV2ResellerRoyaltiesResellerTypeGcp = shared.ContractV2ResellerRoyaltiesResellerTypeGcp

// This is an alias to an internal value.
const ContractV2ResellerRoyaltiesResellerTypeGcpProService = shared.ContractV2ResellerRoyaltiesResellerTypeGcpProService

// This is an alias to an internal type.
type ContractV2ResellerRoyaltiesSegment = shared.ContractV2ResellerRoyaltiesSegment

// This is an alias to an internal type.
type ContractV2ResellerRoyaltiesSegmentsResellerType = shared.ContractV2ResellerRoyaltiesSegmentsResellerType

// This is an alias to an internal value.
const ContractV2ResellerRoyaltiesSegmentsResellerTypeAws = shared.ContractV2ResellerRoyaltiesSegmentsResellerTypeAws

// This is an alias to an internal value.
const ContractV2ResellerRoyaltiesSegmentsResellerTypeAwsProService = shared.ContractV2ResellerRoyaltiesSegmentsResellerTypeAwsProService

// This is an alias to an internal value.
const ContractV2ResellerRoyaltiesSegmentsResellerTypeGcp = shared.ContractV2ResellerRoyaltiesSegmentsResellerTypeGcp

// This is an alias to an internal value.
const ContractV2ResellerRoyaltiesSegmentsResellerTypeGcpProService = shared.ContractV2ResellerRoyaltiesSegmentsResellerTypeGcpProService

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
//
// This is an alias to an internal type.
type ContractV2ScheduledChargesOnUsageInvoices = shared.ContractV2ScheduledChargesOnUsageInvoices

// This is an alias to an internal value.
const ContractV2ScheduledChargesOnUsageInvoicesAll = shared.ContractV2ScheduledChargesOnUsageInvoicesAll

// This is an alias to an internal type.
type ContractV2SpendThresholdConfiguration = shared.ContractV2SpendThresholdConfiguration

// This is an alias to an internal type.
type ContractV2SpendThresholdConfigurationCommit = shared.ContractV2SpendThresholdConfigurationCommit

// This is an alias to an internal type.
type ContractV2SpendThresholdConfigurationPaymentGateConfig = shared.ContractV2SpendThresholdConfigurationPaymentGateConfig

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
//
// This is an alias to an internal type.
type ContractV2SpendThresholdConfigurationPaymentGateConfigPaymentGateType = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigPaymentGateType

// This is an alias to an internal value.
const ContractV2SpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone

// This is an alias to an internal value.
const ContractV2SpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe

// This is an alias to an internal value.
const ContractV2SpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type ContractV2SpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig

// Only applicable if using STRIPE as your payment gateway type.
//
// This is an alias to an internal type.
type ContractV2SpendThresholdConfigurationPaymentGateConfigStripeConfig = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigStripeConfig

// If left blank, will default to INVOICE
//
// This is an alias to an internal type.
type ContractV2SpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType

// This is an alias to an internal value.
const ContractV2SpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice

// This is an alias to an internal value.
const ContractV2SpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
//
// This is an alias to an internal type.
type ContractV2SpendThresholdConfigurationPaymentGateConfigTaxType = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigTaxType

// This is an alias to an internal value.
const ContractV2SpendThresholdConfigurationPaymentGateConfigTaxTypeNone = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigTaxTypeNone

// This is an alias to an internal value.
const ContractV2SpendThresholdConfigurationPaymentGateConfigTaxTypeStripe = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigTaxTypeStripe

// This is an alias to an internal value.
const ContractV2SpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok

// This is an alias to an internal value.
const ContractV2SpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated = shared.ContractV2SpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated

// This is an alias to an internal type.
type ContractV2Subscription = shared.ContractV2Subscription

// This is an alias to an internal type.
type ContractV2SubscriptionsCollectionSchedule = shared.ContractV2SubscriptionsCollectionSchedule

// This is an alias to an internal value.
const ContractV2SubscriptionsCollectionScheduleAdvance = shared.ContractV2SubscriptionsCollectionScheduleAdvance

// This is an alias to an internal value.
const ContractV2SubscriptionsCollectionScheduleArrears = shared.ContractV2SubscriptionsCollectionScheduleArrears

// This is an alias to an internal type.
type ContractV2SubscriptionsProration = shared.ContractV2SubscriptionsProration

// This is an alias to an internal type.
type ContractV2SubscriptionsProrationInvoiceBehavior = shared.ContractV2SubscriptionsProrationInvoiceBehavior

// This is an alias to an internal value.
const ContractV2SubscriptionsProrationInvoiceBehaviorBillImmediately = shared.ContractV2SubscriptionsProrationInvoiceBehaviorBillImmediately

// This is an alias to an internal value.
const ContractV2SubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate = shared.ContractV2SubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate

// This is an alias to an internal type.
type ContractV2SubscriptionsQuantitySchedule = shared.ContractV2SubscriptionsQuantitySchedule

// This is an alias to an internal type.
type ContractV2SubscriptionsSubscriptionRate = shared.ContractV2SubscriptionsSubscriptionRate

// This is an alias to an internal type.
type ContractV2SubscriptionsSubscriptionRateBillingFrequency = shared.ContractV2SubscriptionsSubscriptionRateBillingFrequency

// This is an alias to an internal value.
const ContractV2SubscriptionsSubscriptionRateBillingFrequencyMonthly = shared.ContractV2SubscriptionsSubscriptionRateBillingFrequencyMonthly

// This is an alias to an internal value.
const ContractV2SubscriptionsSubscriptionRateBillingFrequencyQuarterly = shared.ContractV2SubscriptionsSubscriptionRateBillingFrequencyQuarterly

// This is an alias to an internal value.
const ContractV2SubscriptionsSubscriptionRateBillingFrequencyAnnual = shared.ContractV2SubscriptionsSubscriptionRateBillingFrequencyAnnual

// This is an alias to an internal value.
const ContractV2SubscriptionsSubscriptionRateBillingFrequencyWeekly = shared.ContractV2SubscriptionsSubscriptionRateBillingFrequencyWeekly

// This is an alias to an internal type.
type ContractV2SubscriptionsSubscriptionRateProduct = shared.ContractV2SubscriptionsSubscriptionRateProduct

// This is an alias to an internal type.
type ContractWithoutAmendments = shared.ContractWithoutAmendments

// This is an alias to an internal type.
type ContractWithoutAmendmentsTransition = shared.ContractWithoutAmendmentsTransition

// This is an alias to an internal type.
type ContractWithoutAmendmentsTransitionsType = shared.ContractWithoutAmendmentsTransitionsType

// This is an alias to an internal value.
const ContractWithoutAmendmentsTransitionsTypeSupersede = shared.ContractWithoutAmendmentsTransitionsTypeSupersede

// This is an alias to an internal value.
const ContractWithoutAmendmentsTransitionsTypeRenewal = shared.ContractWithoutAmendmentsTransitionsTypeRenewal

// This is an alias to an internal type.
type ContractWithoutAmendmentsUsageStatementSchedule = shared.ContractWithoutAmendmentsUsageStatementSchedule

// This is an alias to an internal type.
type ContractWithoutAmendmentsUsageStatementScheduleFrequency = shared.ContractWithoutAmendmentsUsageStatementScheduleFrequency

// This is an alias to an internal value.
const ContractWithoutAmendmentsUsageStatementScheduleFrequencyMonthly = shared.ContractWithoutAmendmentsUsageStatementScheduleFrequencyMonthly

// This is an alias to an internal value.
const ContractWithoutAmendmentsUsageStatementScheduleFrequencyQuarterly = shared.ContractWithoutAmendmentsUsageStatementScheduleFrequencyQuarterly

// This is an alias to an internal value.
const ContractWithoutAmendmentsUsageStatementScheduleFrequencyAnnual = shared.ContractWithoutAmendmentsUsageStatementScheduleFrequencyAnnual

// This is an alias to an internal value.
const ContractWithoutAmendmentsUsageStatementScheduleFrequencyWeekly = shared.ContractWithoutAmendmentsUsageStatementScheduleFrequencyWeekly

// Either a **parent** configuration with a list of children or a **child**
// configuration with a single parent.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsHierarchyConfiguration = shared.ContractWithoutAmendmentsHierarchyConfiguration

// This is an alias to an internal type.
type ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfiguration = shared.ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfiguration

// This is an alias to an internal type.
type ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationChild = shared.ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationChild

// This is an alias to an internal type.
type ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfiguration = shared.ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfiguration

// The single parent contract/customer for this child.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationParent = shared.ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationParent

// This is an alias to an internal type.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfiguration = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfiguration

// This is an alias to an internal type.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommit = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommit

// This is an alias to an internal type.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitSpecifier = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitSpecifier

// This is an alias to an internal type.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfig = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfig

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType

// This is an alias to an internal value.
const ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone

// This is an alias to an internal value.
const ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe

// This is an alias to an internal value.
const ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig

// Only applicable if using STRIPE as your payment gate type.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig

// If left blank, will default to INVOICE
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType

// This is an alias to an internal value.
const ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice

// This is an alias to an internal value.
const ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType

// This is an alias to an internal value.
const ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone

// This is an alias to an internal value.
const ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe

// This is an alias to an internal value.
const ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok

// This is an alias to an internal value.
const ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated = shared.ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommit = shared.ContractWithoutAmendmentsRecurringCommit

// The amount of commit to grant.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsAccessAmount = shared.ContractWithoutAmendmentsRecurringCommitsAccessAmount

// The amount of time the created commits will be valid for
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsCommitDuration = shared.ContractWithoutAmendmentsRecurringCommitsCommitDuration

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsCommitDurationUnit = shared.ContractWithoutAmendmentsRecurringCommitsCommitDurationUnit

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsCommitDurationUnitPeriods = shared.ContractWithoutAmendmentsRecurringCommitsCommitDurationUnitPeriods

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsProduct = shared.ContractWithoutAmendmentsRecurringCommitsProduct

// Whether the created commits will use the commit rate or list rate
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsRateType = shared.ContractWithoutAmendmentsRecurringCommitsRateType

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsRateTypeCommitRate = shared.ContractWithoutAmendmentsRecurringCommitsRateTypeCommitRate

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsRateTypeListRate = shared.ContractWithoutAmendmentsRecurringCommitsRateTypeListRate

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsContract = shared.ContractWithoutAmendmentsRecurringCommitsContract

// Optional configuration for recurring commit/credit hierarchy access control
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsHierarchyConfiguration = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfiguration

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessType = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessType

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeAll = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeAll

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeNone = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeNone

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeContractIDs = shared.ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeContractIDs

// The amount the customer should be billed for the commit. Not required.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsInvoiceAmount = shared.ContractWithoutAmendmentsRecurringCommitsInvoiceAmount

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsProration = shared.ContractWithoutAmendmentsRecurringCommitsProration

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsProrationNone = shared.ContractWithoutAmendmentsRecurringCommitsProrationNone

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsProrationFirst = shared.ContractWithoutAmendmentsRecurringCommitsProrationFirst

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsProrationLast = shared.ContractWithoutAmendmentsRecurringCommitsProrationLast

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsProrationFirstAndLast = shared.ContractWithoutAmendmentsRecurringCommitsProrationFirstAndLast

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's starting_at rather than the usage
// invoice dates.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency = shared.ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyMonthly = shared.ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyMonthly

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyQuarterly = shared.ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyQuarterly

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyAnnual = shared.ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyAnnual

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyWeekly = shared.ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyWeekly

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsSpecifier = shared.ContractWithoutAmendmentsRecurringCommitsSpecifier

// Attach a subscription to the recurring commit/credit.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsSubscriptionConfig = shared.ContractWithoutAmendmentsRecurringCommitsSubscriptionConfig

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocation = shared.ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocation

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocationIndividual = shared.ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocationIndividual

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocationPooled = shared.ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocationPooled

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig = shared.ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCredit = shared.ContractWithoutAmendmentsRecurringCredit

// The amount of commit to grant.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsAccessAmount = shared.ContractWithoutAmendmentsRecurringCreditsAccessAmount

// The amount of time the created commits will be valid for
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsCommitDuration = shared.ContractWithoutAmendmentsRecurringCreditsCommitDuration

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsCommitDurationUnit = shared.ContractWithoutAmendmentsRecurringCreditsCommitDurationUnit

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsCommitDurationUnitPeriods = shared.ContractWithoutAmendmentsRecurringCreditsCommitDurationUnitPeriods

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsProduct = shared.ContractWithoutAmendmentsRecurringCreditsProduct

// Whether the created commits will use the commit rate or list rate
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsRateType = shared.ContractWithoutAmendmentsRecurringCreditsRateType

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsRateTypeCommitRate = shared.ContractWithoutAmendmentsRecurringCreditsRateTypeCommitRate

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsRateTypeListRate = shared.ContractWithoutAmendmentsRecurringCreditsRateTypeListRate

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsContract = shared.ContractWithoutAmendmentsRecurringCreditsContract

// Optional configuration for recurring commit/credit hierarchy access control
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsHierarchyConfiguration = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfiguration

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessType = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessType

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeAll = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeAll

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeNone = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeNone

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeContractIDs = shared.ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeContractIDs

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsProration = shared.ContractWithoutAmendmentsRecurringCreditsProration

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsProrationNone = shared.ContractWithoutAmendmentsRecurringCreditsProrationNone

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsProrationFirst = shared.ContractWithoutAmendmentsRecurringCreditsProrationFirst

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsProrationLast = shared.ContractWithoutAmendmentsRecurringCreditsProrationLast

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsProrationFirstAndLast = shared.ContractWithoutAmendmentsRecurringCreditsProrationFirstAndLast

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's starting_at rather than the usage
// invoice dates.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency = shared.ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyMonthly = shared.ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyMonthly

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyQuarterly = shared.ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyQuarterly

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyAnnual = shared.ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyAnnual

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyWeekly = shared.ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyWeekly

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsSpecifier = shared.ContractWithoutAmendmentsRecurringCreditsSpecifier

// Attach a subscription to the recurring commit/credit.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsSubscriptionConfig = shared.ContractWithoutAmendmentsRecurringCreditsSubscriptionConfig

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocation = shared.ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocation

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocationIndividual = shared.ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocationIndividual

// This is an alias to an internal value.
const ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocationPooled = shared.ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocationPooled

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig = shared.ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig

// This is an alias to an internal type.
type ContractWithoutAmendmentsResellerRoyalty = shared.ContractWithoutAmendmentsResellerRoyalty

// This is an alias to an internal type.
type ContractWithoutAmendmentsResellerRoyaltiesResellerType = shared.ContractWithoutAmendmentsResellerRoyaltiesResellerType

// This is an alias to an internal value.
const ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAws = shared.ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAws

// This is an alias to an internal value.
const ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAwsProService = shared.ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAwsProService

// This is an alias to an internal value.
const ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcp = shared.ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcp

// This is an alias to an internal value.
const ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcpProService = shared.ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcpProService

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsScheduledChargesOnUsageInvoices = shared.ContractWithoutAmendmentsScheduledChargesOnUsageInvoices

// This is an alias to an internal value.
const ContractWithoutAmendmentsScheduledChargesOnUsageInvoicesAll = shared.ContractWithoutAmendmentsScheduledChargesOnUsageInvoicesAll

// This is an alias to an internal type.
type ContractWithoutAmendmentsSpendThresholdConfiguration = shared.ContractWithoutAmendmentsSpendThresholdConfiguration

// This is an alias to an internal type.
type ContractWithoutAmendmentsSpendThresholdConfigurationCommit = shared.ContractWithoutAmendmentsSpendThresholdConfigurationCommit

// This is an alias to an internal type.
type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfig = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfig

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateType = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateType

// This is an alias to an internal value.
const ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone

// This is an alias to an internal value.
const ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe

// This is an alias to an internal value.
const ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig

// Only applicable if using STRIPE as your payment gate type.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfig = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfig

// If left blank, will default to INVOICE
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType

// This is an alias to an internal value.
const ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice

// This is an alias to an internal value.
const ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxType = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxType

// This is an alias to an internal value.
const ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeNone = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeNone

// This is an alias to an internal value.
const ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe

// This is an alias to an internal value.
const ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok

// This is an alias to an internal value.
const ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated = shared.ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated

// This is an alias to an internal type.
type ContractWithoutAmendmentsUsageFilter = shared.ContractWithoutAmendmentsUsageFilter

// This is an alias to an internal type.
type ContractWithoutAmendmentsUsageFilterUpdate = shared.ContractWithoutAmendmentsUsageFilterUpdate

// This is an alias to an internal type.
type Credit = shared.Credit

// This is an alias to an internal type.
type CreditProduct = shared.CreditProduct

// This is an alias to an internal type.
type CreditType = shared.CreditType

// This is an alias to an internal value.
const CreditTypeCredit = shared.CreditTypeCredit

// This is an alias to an internal type.
type CreditContract = shared.CreditContract

// Optional configuration for credit hierarchy access control
//
// This is an alias to an internal type.
type CreditHierarchyConfiguration = shared.CreditHierarchyConfiguration

// This is an alias to an internal type.
type CreditHierarchyConfigurationChildAccess = shared.CreditHierarchyConfigurationChildAccess

// This is an alias to an internal type.
type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll = shared.CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll

// This is an alias to an internal type.
type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = shared.CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType

// This is an alias to an internal value.
const CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll = shared.CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll

// This is an alias to an internal type.
type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone = shared.CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone

// This is an alias to an internal type.
type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = shared.CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType

// This is an alias to an internal value.
const CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone = shared.CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone

// This is an alias to an internal type.
type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs = shared.CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs

// This is an alias to an internal type.
type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = shared.CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType

// This is an alias to an internal value.
const CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs = shared.CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs

// This is an alias to an internal type.
type CreditHierarchyConfigurationChildAccessType = shared.CreditHierarchyConfigurationChildAccessType

// This is an alias to an internal value.
const CreditHierarchyConfigurationChildAccessTypeAll = shared.CreditHierarchyConfigurationChildAccessTypeAll

// This is an alias to an internal value.
const CreditHierarchyConfigurationChildAccessTypeNone = shared.CreditHierarchyConfigurationChildAccessTypeNone

// This is an alias to an internal value.
const CreditHierarchyConfigurationChildAccessTypeContractIDs = shared.CreditHierarchyConfigurationChildAccessTypeContractIDs

// This is an alias to an internal type.
type CreditLedger = shared.CreditLedger

// This is an alias to an internal type.
type CreditLedgerCreditSegmentStartLedgerEntry = shared.CreditLedgerCreditSegmentStartLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditSegmentStartLedgerEntryType = shared.CreditLedgerCreditSegmentStartLedgerEntryType

// This is an alias to an internal value.
const CreditLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart = shared.CreditLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart

// This is an alias to an internal type.
type CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry = shared.CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = shared.CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType

// This is an alias to an internal value.
const CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction = shared.CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction

// This is an alias to an internal type.
type CreditLedgerCreditExpirationLedgerEntry = shared.CreditLedgerCreditExpirationLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditExpirationLedgerEntryType = shared.CreditLedgerCreditExpirationLedgerEntryType

// This is an alias to an internal value.
const CreditLedgerCreditExpirationLedgerEntryTypeCreditExpiration = shared.CreditLedgerCreditExpirationLedgerEntryTypeCreditExpiration

// This is an alias to an internal type.
type CreditLedgerCreditCanceledLedgerEntry = shared.CreditLedgerCreditCanceledLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditCanceledLedgerEntryType = shared.CreditLedgerCreditCanceledLedgerEntryType

// This is an alias to an internal value.
const CreditLedgerCreditCanceledLedgerEntryTypeCreditCanceled = shared.CreditLedgerCreditCanceledLedgerEntryTypeCreditCanceled

// This is an alias to an internal type.
type CreditLedgerCreditCreditedLedgerEntry = shared.CreditLedgerCreditCreditedLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditCreditedLedgerEntryType = shared.CreditLedgerCreditCreditedLedgerEntryType

// This is an alias to an internal value.
const CreditLedgerCreditCreditedLedgerEntryTypeCreditCredited = shared.CreditLedgerCreditCreditedLedgerEntryTypeCreditCredited

// This is an alias to an internal type.
type CreditLedgerCreditManualLedgerEntry = shared.CreditLedgerCreditManualLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditManualLedgerEntryType = shared.CreditLedgerCreditManualLedgerEntryType

// This is an alias to an internal value.
const CreditLedgerCreditManualLedgerEntryTypeCreditManual = shared.CreditLedgerCreditManualLedgerEntryTypeCreditManual

// This is an alias to an internal type.
type CreditLedgerCreditSeatBasedAdjustmentLedgerEntry = shared.CreditLedgerCreditSeatBasedAdjustmentLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditSeatBasedAdjustmentLedgerEntryType = shared.CreditLedgerCreditSeatBasedAdjustmentLedgerEntryType

// This is an alias to an internal value.
const CreditLedgerCreditSeatBasedAdjustmentLedgerEntryTypeCreditSeatBasedAdjustment = shared.CreditLedgerCreditSeatBasedAdjustmentLedgerEntryTypeCreditSeatBasedAdjustment

// This is an alias to an internal type.
type CreditLedgerType = shared.CreditLedgerType

// This is an alias to an internal value.
const CreditLedgerTypeCreditSegmentStart = shared.CreditLedgerTypeCreditSegmentStart

// This is an alias to an internal value.
const CreditLedgerTypeCreditAutomatedInvoiceDeduction = shared.CreditLedgerTypeCreditAutomatedInvoiceDeduction

// This is an alias to an internal value.
const CreditLedgerTypeCreditExpiration = shared.CreditLedgerTypeCreditExpiration

// This is an alias to an internal value.
const CreditLedgerTypeCreditCanceled = shared.CreditLedgerTypeCreditCanceled

// This is an alias to an internal value.
const CreditLedgerTypeCreditCredited = shared.CreditLedgerTypeCreditCredited

// This is an alias to an internal value.
const CreditLedgerTypeCreditManual = shared.CreditLedgerTypeCreditManual

// This is an alias to an internal value.
const CreditLedgerTypeCreditSeatBasedAdjustment = shared.CreditLedgerTypeCreditSeatBasedAdjustment

// This is an alias to an internal type.
type CreditRateType = shared.CreditRateType

// This is an alias to an internal value.
const CreditRateTypeCommitRate = shared.CreditRateTypeCommitRate

// This is an alias to an internal value.
const CreditRateTypeListRate = shared.CreditRateTypeListRate

// This is an alias to an internal type.
type CreditSpecifier = shared.CreditSpecifier

// This is an alias to an internal type.
type CreditTypeData = shared.CreditTypeData

// This is an alias to an internal type.
type Discount = shared.Discount

// This is an alias to an internal type.
type DiscountProduct = shared.DiscountProduct

// An optional filtering rule to match the 'event_type' property of an event.
//
// This is an alias to an internal type.
type EventTypeFilter = shared.EventTypeFilter

// An optional filtering rule to match the 'event_type' property of an event.
//
// This is an alias to an internal type.
type EventTypeFilterParam = shared.EventTypeFilterParam

// This is an alias to an internal type.
type ID = shared.ID

// This is an alias to an internal type.
type IDParam = shared.IDParam

// This is an alias to an internal type.
type Override = shared.Override

// This is an alias to an internal type.
type OverrideOverrideSpecifier = shared.OverrideOverrideSpecifier

// This is an alias to an internal type.
type OverrideOverrideSpecifiersBillingFrequency = shared.OverrideOverrideSpecifiersBillingFrequency

// This is an alias to an internal value.
const OverrideOverrideSpecifiersBillingFrequencyMonthly = shared.OverrideOverrideSpecifiersBillingFrequencyMonthly

// This is an alias to an internal value.
const OverrideOverrideSpecifiersBillingFrequencyQuarterly = shared.OverrideOverrideSpecifiersBillingFrequencyQuarterly

// This is an alias to an internal value.
const OverrideOverrideSpecifiersBillingFrequencyAnnual = shared.OverrideOverrideSpecifiersBillingFrequencyAnnual

// This is an alias to an internal value.
const OverrideOverrideSpecifiersBillingFrequencyWeekly = shared.OverrideOverrideSpecifiersBillingFrequencyWeekly

// This is an alias to an internal type.
type OverrideOverrideTier = shared.OverrideOverrideTier

// This is an alias to an internal type.
type OverrideOverwriteRate = shared.OverrideOverwriteRate

// This is an alias to an internal type.
type OverrideOverwriteRateRateType = shared.OverrideOverwriteRateRateType

// This is an alias to an internal value.
const OverrideOverwriteRateRateTypeFlat = shared.OverrideOverwriteRateRateTypeFlat

// This is an alias to an internal value.
const OverrideOverwriteRateRateTypePercentage = shared.OverrideOverwriteRateRateTypePercentage

// This is an alias to an internal value.
const OverrideOverwriteRateRateTypeSubscription = shared.OverrideOverwriteRateRateTypeSubscription

// This is an alias to an internal value.
const OverrideOverwriteRateRateTypeTiered = shared.OverrideOverwriteRateRateTypeTiered

// This is an alias to an internal value.
const OverrideOverwriteRateRateTypeCustom = shared.OverrideOverwriteRateRateTypeCustom

// This is an alias to an internal type.
type OverrideProduct = shared.OverrideProduct

// This is an alias to an internal type.
type OverrideRateType = shared.OverrideRateType

// This is an alias to an internal value.
const OverrideRateTypeFlat = shared.OverrideRateTypeFlat

// This is an alias to an internal value.
const OverrideRateTypePercentage = shared.OverrideRateTypePercentage

// This is an alias to an internal value.
const OverrideRateTypeSubscription = shared.OverrideRateTypeSubscription

// This is an alias to an internal value.
const OverrideRateTypeTiered = shared.OverrideRateTypeTiered

// This is an alias to an internal value.
const OverrideRateTypeCustom = shared.OverrideRateTypeCustom

// This is an alias to an internal type.
type OverrideTarget = shared.OverrideTarget

// This is an alias to an internal value.
const OverrideTargetCommitRate = shared.OverrideTargetCommitRate

// This is an alias to an internal value.
const OverrideTargetListRate = shared.OverrideTargetListRate

// This is an alias to an internal type.
type OverrideType = shared.OverrideType

// This is an alias to an internal value.
const OverrideTypeOverwrite = shared.OverrideTypeOverwrite

// This is an alias to an internal value.
const OverrideTypeMultiplier = shared.OverrideTypeMultiplier

// This is an alias to an internal value.
const OverrideTypeTiered = shared.OverrideTypeTiered

// This is an alias to an internal type.
type PropertyFilter = shared.PropertyFilter

// This is an alias to an internal type.
type PropertyFilterParam = shared.PropertyFilterParam

// This is an alias to an internal type.
type ProService = shared.ProService

// This is an alias to an internal type.
type Rate = shared.Rate

// This is an alias to an internal type.
type RateRateType = shared.RateRateType

// This is an alias to an internal value.
const RateRateTypeFlat = shared.RateRateTypeFlat

// This is an alias to an internal value.
const RateRateTypePercentage = shared.RateRateTypePercentage

// This is an alias to an internal value.
const RateRateTypeSubscription = shared.RateRateTypeSubscription

// This is an alias to an internal value.
const RateRateTypeCustom = shared.RateRateTypeCustom

// This is an alias to an internal value.
const RateRateTypeTiered = shared.RateRateTypeTiered

// This is an alias to an internal type.
type ScheduledCharge = shared.ScheduledCharge

// This is an alias to an internal type.
type ScheduledChargeProduct = shared.ScheduledChargeProduct

// This is an alias to an internal type.
type ScheduleDuration = shared.ScheduleDuration

// This is an alias to an internal type.
type ScheduleDurationScheduleItem = shared.ScheduleDurationScheduleItem

// This is an alias to an internal type.
type SchedulePointInTime = shared.SchedulePointInTime

// This is an alias to an internal type.
type SchedulePointInTimeScheduleItem = shared.SchedulePointInTimeScheduleItem

// This is an alias to an internal type.
type Tier = shared.Tier

// This is an alias to an internal type.
type TierParam = shared.TierParam
