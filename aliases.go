// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"github.com/Metronome-Industries/metronome-go/internal/apierror"
	"github.com/Metronome-Industries/metronome-go/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type BaseThresholdCommit = shared.BaseThresholdCommit

// This is an alias to an internal type.
type BaseThresholdCommitParam = shared.BaseThresholdCommitParam

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

// This is an alias to an internal type.
type CommitHierarchyConfigurationParam = shared.CommitHierarchyConfigurationParam

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessUnionParam = shared.CommitHierarchyConfigurationChildAccessUnionParam

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
//
// This is an alias to an internal type.
type CommitRate = shared.CommitRate

// This is an alias to an internal type.
type CommitRateRateType = shared.CommitRateRateType

// This is an alias to an internal value.
const CommitRateRateTypeFlat = shared.CommitRateRateTypeFlat

// This is an alias to an internal value.
const CommitRateRateTypePercentage = shared.CommitRateRateTypePercentage

// This is an alias to an internal value.
const CommitRateRateTypeSubscription = shared.CommitRateRateTypeSubscription

// This is an alias to an internal value.
const CommitRateRateTypeTiered = shared.CommitRateRateTypeTiered

// This is an alias to an internal value.
const CommitRateRateTypeCustom = shared.CommitRateRateTypeCustom

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
//
// This is an alias to an internal type.
type CommitRateParam = shared.CommitRateParam

// This is an alias to an internal type.
type CommitSpecifier = shared.CommitSpecifier

// This is an alias to an internal type.
type CommitSpecifierInput = shared.CommitSpecifierInput

// This is an alias to an internal type.
type CommitSpecifierInputParam = shared.CommitSpecifierInputParam

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

// Either a **parent** configuration with a list of children or a **child**
// configuration with a single parent.
//
// This is an alias to an internal type.
type HierarchyConfiguration = shared.HierarchyConfiguration

// This is an alias to an internal type.
type HierarchyConfigurationParentHierarchyConfiguration = shared.HierarchyConfigurationParentHierarchyConfiguration

// This is an alias to an internal type.
type HierarchyConfigurationParentHierarchyConfigurationChild = shared.HierarchyConfigurationParentHierarchyConfigurationChild

// This is an alias to an internal type.
type HierarchyConfigurationChildHierarchyConfiguration = shared.HierarchyConfigurationChildHierarchyConfiguration

// The single parent contract/customer for this child.
//
// This is an alias to an internal type.
type HierarchyConfigurationChildHierarchyConfigurationParent = shared.HierarchyConfigurationChildHierarchyConfigurationParent

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
type OverrideTier = shared.OverrideTier

// This is an alias to an internal type.
type OverwriteRate = shared.OverwriteRate

// This is an alias to an internal type.
type OverwriteRateRateType = shared.OverwriteRateRateType

// This is an alias to an internal value.
const OverwriteRateRateTypeFlat = shared.OverwriteRateRateTypeFlat

// This is an alias to an internal value.
const OverwriteRateRateTypePercentage = shared.OverwriteRateRateTypePercentage

// This is an alias to an internal value.
const OverwriteRateRateTypeSubscription = shared.OverwriteRateRateTypeSubscription

// This is an alias to an internal value.
const OverwriteRateRateTypeTiered = shared.OverwriteRateRateTypeTiered

// This is an alias to an internal value.
const OverwriteRateRateTypeCustom = shared.OverwriteRateRateTypeCustom

// This is an alias to an internal type.
type PaymentGateConfig = shared.PaymentGateConfig

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
//
// This is an alias to an internal type.
type PaymentGateConfigPaymentGateType = shared.PaymentGateConfigPaymentGateType

// This is an alias to an internal value.
const PaymentGateConfigPaymentGateTypeNone = shared.PaymentGateConfigPaymentGateTypeNone

// This is an alias to an internal value.
const PaymentGateConfigPaymentGateTypeStripe = shared.PaymentGateConfigPaymentGateTypeStripe

// This is an alias to an internal value.
const PaymentGateConfigPaymentGateTypeExternal = shared.PaymentGateConfigPaymentGateTypeExternal

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type PaymentGateConfigPrecalculatedTaxConfig = shared.PaymentGateConfigPrecalculatedTaxConfig

// Only applicable if using STRIPE as your payment gate type.
//
// This is an alias to an internal type.
type PaymentGateConfigStripeConfig = shared.PaymentGateConfigStripeConfig

// If left blank, will default to INVOICE
//
// This is an alias to an internal type.
type PaymentGateConfigStripeConfigPaymentType = shared.PaymentGateConfigStripeConfigPaymentType

// This is an alias to an internal value.
const PaymentGateConfigStripeConfigPaymentTypeInvoice = shared.PaymentGateConfigStripeConfigPaymentTypeInvoice

// This is an alias to an internal value.
const PaymentGateConfigStripeConfigPaymentTypePaymentIntent = shared.PaymentGateConfigStripeConfigPaymentTypePaymentIntent

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
//
// This is an alias to an internal type.
type PaymentGateConfigTaxType = shared.PaymentGateConfigTaxType

// This is an alias to an internal value.
const PaymentGateConfigTaxTypeNone = shared.PaymentGateConfigTaxTypeNone

// This is an alias to an internal value.
const PaymentGateConfigTaxTypeStripe = shared.PaymentGateConfigTaxTypeStripe

// This is an alias to an internal value.
const PaymentGateConfigTaxTypeAnrok = shared.PaymentGateConfigTaxTypeAnrok

// This is an alias to an internal value.
const PaymentGateConfigTaxTypePrecalculated = shared.PaymentGateConfigTaxTypePrecalculated

// This is an alias to an internal type.
type PaymentGateConfigParam = shared.PaymentGateConfigParam

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type PaymentGateConfigPrecalculatedTaxConfigParam = shared.PaymentGateConfigPrecalculatedTaxConfigParam

// Only applicable if using STRIPE as your payment gate type.
//
// This is an alias to an internal type.
type PaymentGateConfigStripeConfigParam = shared.PaymentGateConfigStripeConfigParam

// This is an alias to an internal type.
type PaymentGateConfigV2 = shared.PaymentGateConfigV2

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
//
// This is an alias to an internal type.
type PaymentGateConfigV2PaymentGateType = shared.PaymentGateConfigV2PaymentGateType

// This is an alias to an internal value.
const PaymentGateConfigV2PaymentGateTypeNone = shared.PaymentGateConfigV2PaymentGateTypeNone

// This is an alias to an internal value.
const PaymentGateConfigV2PaymentGateTypeStripe = shared.PaymentGateConfigV2PaymentGateTypeStripe

// This is an alias to an internal value.
const PaymentGateConfigV2PaymentGateTypeExternal = shared.PaymentGateConfigV2PaymentGateTypeExternal

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type PaymentGateConfigV2PrecalculatedTaxConfig = shared.PaymentGateConfigV2PrecalculatedTaxConfig

// Only applicable if using STRIPE as your payment gateway type.
//
// This is an alias to an internal type.
type PaymentGateConfigV2StripeConfig = shared.PaymentGateConfigV2StripeConfig

// If left blank, will default to INVOICE
//
// This is an alias to an internal type.
type PaymentGateConfigV2StripeConfigPaymentType = shared.PaymentGateConfigV2StripeConfigPaymentType

// This is an alias to an internal value.
const PaymentGateConfigV2StripeConfigPaymentTypeInvoice = shared.PaymentGateConfigV2StripeConfigPaymentTypeInvoice

// This is an alias to an internal value.
const PaymentGateConfigV2StripeConfigPaymentTypePaymentIntent = shared.PaymentGateConfigV2StripeConfigPaymentTypePaymentIntent

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
//
// This is an alias to an internal type.
type PaymentGateConfigV2TaxType = shared.PaymentGateConfigV2TaxType

// This is an alias to an internal value.
const PaymentGateConfigV2TaxTypeNone = shared.PaymentGateConfigV2TaxTypeNone

// This is an alias to an internal value.
const PaymentGateConfigV2TaxTypeStripe = shared.PaymentGateConfigV2TaxTypeStripe

// This is an alias to an internal value.
const PaymentGateConfigV2TaxTypeAnrok = shared.PaymentGateConfigV2TaxTypeAnrok

// This is an alias to an internal value.
const PaymentGateConfigV2TaxTypePrecalculated = shared.PaymentGateConfigV2TaxTypePrecalculated

// This is an alias to an internal type.
type PaymentGateConfigV2Param = shared.PaymentGateConfigV2Param

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type PaymentGateConfigV2PrecalculatedTaxConfigParam = shared.PaymentGateConfigV2PrecalculatedTaxConfigParam

// Only applicable if using STRIPE as your payment gateway type.
//
// This is an alias to an internal type.
type PaymentGateConfigV2StripeConfigParam = shared.PaymentGateConfigV2StripeConfigParam

// This is an alias to an internal type.
type PrepaidBalanceThresholdConfiguration = shared.PrepaidBalanceThresholdConfiguration

// This is an alias to an internal type.
type PrepaidBalanceThresholdConfigurationCommit = shared.PrepaidBalanceThresholdConfigurationCommit

// This is an alias to an internal type.
type PrepaidBalanceThresholdConfigurationParam = shared.PrepaidBalanceThresholdConfigurationParam

// This is an alias to an internal type.
type PrepaidBalanceThresholdConfigurationCommitParam = shared.PrepaidBalanceThresholdConfigurationCommitParam

// This is an alias to an internal type.
type PrepaidBalanceThresholdConfigurationV2 = shared.PrepaidBalanceThresholdConfigurationV2

// This is an alias to an internal type.
type PrepaidBalanceThresholdConfigurationV2Commit = shared.PrepaidBalanceThresholdConfigurationV2Commit

// This is an alias to an internal type.
type PrepaidBalanceThresholdConfigurationV2Param = shared.PrepaidBalanceThresholdConfigurationV2Param

// This is an alias to an internal type.
type PrepaidBalanceThresholdConfigurationV2CommitParam = shared.PrepaidBalanceThresholdConfigurationV2CommitParam

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
type RecurringCommitSubscriptionConfig = shared.RecurringCommitSubscriptionConfig

// This is an alias to an internal type.
type RecurringCommitSubscriptionConfigAllocation = shared.RecurringCommitSubscriptionConfigAllocation

// This is an alias to an internal value.
const RecurringCommitSubscriptionConfigAllocationIndividual = shared.RecurringCommitSubscriptionConfigAllocationIndividual

// This is an alias to an internal value.
const RecurringCommitSubscriptionConfigAllocationPooled = shared.RecurringCommitSubscriptionConfigAllocationPooled

// This is an alias to an internal type.
type RecurringCommitSubscriptionConfigApplySeatIncreaseConfig = shared.RecurringCommitSubscriptionConfigApplySeatIncreaseConfig

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
type SpendThresholdConfiguration = shared.SpendThresholdConfiguration

// This is an alias to an internal type.
type SpendThresholdConfigurationParam = shared.SpendThresholdConfigurationParam

// This is an alias to an internal type.
type SpendThresholdConfigurationV2 = shared.SpendThresholdConfigurationV2

// This is an alias to an internal type.
type SpendThresholdConfigurationV2Param = shared.SpendThresholdConfigurationV2Param

// This is an alias to an internal type.
type Subscription = shared.Subscription

// This is an alias to an internal type.
type SubscriptionCollectionSchedule = shared.SubscriptionCollectionSchedule

// This is an alias to an internal value.
const SubscriptionCollectionScheduleAdvance = shared.SubscriptionCollectionScheduleAdvance

// This is an alias to an internal value.
const SubscriptionCollectionScheduleArrears = shared.SubscriptionCollectionScheduleArrears

// This is an alias to an internal type.
type SubscriptionProration = shared.SubscriptionProration

// This is an alias to an internal type.
type SubscriptionProrationInvoiceBehavior = shared.SubscriptionProrationInvoiceBehavior

// This is an alias to an internal value.
const SubscriptionProrationInvoiceBehaviorBillImmediately = shared.SubscriptionProrationInvoiceBehaviorBillImmediately

// This is an alias to an internal value.
const SubscriptionProrationInvoiceBehaviorBillOnNextCollectionDate = shared.SubscriptionProrationInvoiceBehaviorBillOnNextCollectionDate

// This is an alias to an internal type.
type SubscriptionQuantitySchedule = shared.SubscriptionQuantitySchedule

// This is an alias to an internal type.
type SubscriptionSubscriptionRate = shared.SubscriptionSubscriptionRate

// This is an alias to an internal type.
type SubscriptionSubscriptionRateBillingFrequency = shared.SubscriptionSubscriptionRateBillingFrequency

// This is an alias to an internal value.
const SubscriptionSubscriptionRateBillingFrequencyMonthly = shared.SubscriptionSubscriptionRateBillingFrequencyMonthly

// This is an alias to an internal value.
const SubscriptionSubscriptionRateBillingFrequencyQuarterly = shared.SubscriptionSubscriptionRateBillingFrequencyQuarterly

// This is an alias to an internal value.
const SubscriptionSubscriptionRateBillingFrequencyAnnual = shared.SubscriptionSubscriptionRateBillingFrequencyAnnual

// This is an alias to an internal value.
const SubscriptionSubscriptionRateBillingFrequencyWeekly = shared.SubscriptionSubscriptionRateBillingFrequencyWeekly

// This is an alias to an internal type.
type SubscriptionSubscriptionRateProduct = shared.SubscriptionSubscriptionRateProduct

// This is an alias to an internal type.
type Tier = shared.Tier

// This is an alias to an internal type.
type TierParam = shared.TierParam

// This is an alias to an internal type.
type UpdateBaseThresholdCommit = shared.UpdateBaseThresholdCommit

// This is an alias to an internal type.
type UpdateBaseThresholdCommitParam = shared.UpdateBaseThresholdCommitParam
