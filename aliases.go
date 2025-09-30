// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"github.com/Metronome-Industries/metronome-go/internal/apierror"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

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

// Equals "PREPAID"
const CommitTypePrepaid = shared.CommitTypePrepaid

// Equals "POSTPAID"
const CommitTypePostpaid = shared.CommitTypePostpaid

// This is an alias to an internal type.
type CommitContract = shared.CommitContract

// The contract that this commit will be billed on.
//
// This is an alias to an internal type.
type CommitInvoiceContract = shared.CommitInvoiceContract

// This is an alias to an internal type.
type CommitLedgerUnion = shared.CommitLedgerUnion

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitSegmentStartLedgerEntry = shared.CommitLedgerPrepaidCommitSegmentStartLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry = shared.CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitRolloverLedgerEntry = shared.CommitLedgerPrepaidCommitRolloverLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitExpirationLedgerEntry = shared.CommitLedgerPrepaidCommitExpirationLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitCanceledLedgerEntry = shared.CommitLedgerPrepaidCommitCanceledLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitCreditedLedgerEntry = shared.CommitLedgerPrepaidCommitCreditedLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry = shared.CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitInitialBalanceLedgerEntry = shared.CommitLedgerPostpaidCommitInitialBalanceLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry = shared.CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitRolloverLedgerEntry = shared.CommitLedgerPostpaidCommitRolloverLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitTrueupLedgerEntry = shared.CommitLedgerPostpaidCommitTrueupLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitManualLedgerEntry = shared.CommitLedgerPrepaidCommitManualLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitManualLedgerEntry = shared.CommitLedgerPostpaidCommitManualLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitExpirationLedgerEntry = shared.CommitLedgerPostpaidCommitExpirationLedgerEntry

// This is an alias to an internal type.
type CommitRateType = shared.CommitRateType

// Equals "COMMIT_RATE"
const CommitRateTypeCommitRate = shared.CommitRateTypeCommitRate

// Equals "LIST_RATE"
const CommitRateTypeListRate = shared.CommitRateTypeListRate

// This is an alias to an internal type.
type CommitRolledOverFrom = shared.CommitRolledOverFrom

// This is an alias to an internal type.
type CommitHierarchyConfiguration = shared.CommitHierarchyConfiguration

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessUnion = shared.CommitHierarchyConfigurationChildAccessUnion

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone

// This is an alias to an internal type.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs = shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs

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

// Equals "FLAT"
const CommitRateRateTypeFlat = shared.CommitRateRateTypeFlat

// Equals "PERCENTAGE"
const CommitRateRateTypePercentage = shared.CommitRateRateTypePercentage

// Equals "SUBSCRIPTION"
const CommitRateRateTypeSubscription = shared.CommitRateRateTypeSubscription

// Equals "TIERED"
const CommitRateRateTypeTiered = shared.CommitRateRateTypeTiered

// Equals "CUSTOM"
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
type ContractAmendmentResellerRoyalty = shared.ContractAmendmentResellerRoyalty

// The billing provider configuration associated with a contract.
//
// This is an alias to an internal type.
type ContractCustomerBillingProviderConfiguration = shared.ContractCustomerBillingProviderConfiguration

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
//
// This is an alias to an internal type.
type ContractScheduledChargesOnUsageInvoices = shared.ContractScheduledChargesOnUsageInvoices

// Equals "ALL"
const ContractScheduledChargesOnUsageInvoicesAll = shared.ContractScheduledChargesOnUsageInvoicesAll

// This is an alias to an internal type.
type ContractV2 = shared.ContractV2

// This is an alias to an internal type.
type ContractV2Commit = shared.ContractV2Commit

// This is an alias to an internal type.
type ContractV2CommitProduct = shared.ContractV2CommitProduct

// This is an alias to an internal type.
type ContractV2CommitContract = shared.ContractV2CommitContract

// The contract that this commit will be billed on.
//
// This is an alias to an internal type.
type ContractV2CommitInvoiceContract = shared.ContractV2CommitInvoiceContract

// This is an alias to an internal type.
type ContractV2CommitLedgerUnion = shared.ContractV2CommitLedgerUnion

// This is an alias to an internal type.
type ContractV2CommitLedgerPrepaidCommitSegmentStartLedgerEntry = shared.ContractV2CommitLedgerPrepaidCommitSegmentStartLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry = shared.ContractV2CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPrepaidCommitRolloverLedgerEntry = shared.ContractV2CommitLedgerPrepaidCommitRolloverLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPrepaidCommitExpirationLedgerEntry = shared.ContractV2CommitLedgerPrepaidCommitExpirationLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPrepaidCommitCanceledLedgerEntry = shared.ContractV2CommitLedgerPrepaidCommitCanceledLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPrepaidCommitCreditedLedgerEntry = shared.ContractV2CommitLedgerPrepaidCommitCreditedLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry = shared.ContractV2CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPostpaidCommitInitialBalanceLedgerEntry = shared.ContractV2CommitLedgerPostpaidCommitInitialBalanceLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry = shared.ContractV2CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPostpaidCommitRolloverLedgerEntry = shared.ContractV2CommitLedgerPostpaidCommitRolloverLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPostpaidCommitTrueupLedgerEntry = shared.ContractV2CommitLedgerPostpaidCommitTrueupLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPrepaidCommitManualLedgerEntry = shared.ContractV2CommitLedgerPrepaidCommitManualLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPostpaidCommitManualLedgerEntry = shared.ContractV2CommitLedgerPostpaidCommitManualLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitLedgerPostpaidCommitExpirationLedgerEntry = shared.ContractV2CommitLedgerPostpaidCommitExpirationLedgerEntry

// This is an alias to an internal type.
type ContractV2CommitRolledOverFrom = shared.ContractV2CommitRolledOverFrom

// This is an alias to an internal type.
type ContractV2Override = shared.ContractV2Override

// This is an alias to an internal type.
type ContractV2OverrideOverrideSpecifier = shared.ContractV2OverrideOverrideSpecifier

// This is an alias to an internal type.
type ContractV2OverrideOverwriteRate = shared.ContractV2OverrideOverwriteRate

// This is an alias to an internal type.
type ContractV2OverrideProduct = shared.ContractV2OverrideProduct

// This is an alias to an internal type.
type ContractV2Transition = shared.ContractV2Transition

// This is an alias to an internal type.
type ContractV2UsageFilter = shared.ContractV2UsageFilter

// This is an alias to an internal type.
type ContractV2UsageStatementSchedule = shared.ContractV2UsageStatementSchedule

// This is an alias to an internal type.
type ContractV2Credit = shared.ContractV2Credit

// This is an alias to an internal type.
type ContractV2CreditProduct = shared.ContractV2CreditProduct

// This is an alias to an internal type.
type ContractV2CreditContract = shared.ContractV2CreditContract

// This is an alias to an internal type.
type ContractV2CreditLedgerUnion = shared.ContractV2CreditLedgerUnion

// This is an alias to an internal type.
type ContractV2CreditLedgerCreditSegmentStartLedgerEntry = shared.ContractV2CreditLedgerCreditSegmentStartLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry = shared.ContractV2CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditLedgerCreditExpirationLedgerEntry = shared.ContractV2CreditLedgerCreditExpirationLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditLedgerCreditCanceledLedgerEntry = shared.ContractV2CreditLedgerCreditCanceledLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditLedgerCreditCreditedLedgerEntry = shared.ContractV2CreditLedgerCreditCreditedLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditLedgerCreditManualLedgerEntry = shared.ContractV2CreditLedgerCreditManualLedgerEntry

// This is an alias to an internal type.
type ContractV2CreditLedgerCreditSeatBasedAdjustmentLedgerEntry = shared.ContractV2CreditLedgerCreditSeatBasedAdjustmentLedgerEntry

// This field's availability is dependent on your client's configuration.
//
// This is an alias to an internal type.
type ContractV2CustomerBillingProviderConfiguration = shared.ContractV2CustomerBillingProviderConfiguration

// Indicates whether there are more items than the limit for this endpoint. Use the
// respective list endpoints to get the full lists.
//
// This is an alias to an internal type.
type ContractV2HasMore = shared.ContractV2HasMore

// Either a **parent** configuration with a list of children or a **child**
// configuration with a single parent.
//
// This is an alias to an internal type.
type ContractV2HierarchyConfigurationUnion = shared.ContractV2HierarchyConfigurationUnion

// This is an alias to an internal type.
type ContractV2HierarchyConfigurationParentHierarchyConfiguration = shared.ContractV2HierarchyConfigurationParentHierarchyConfiguration

// This is an alias to an internal type.
type ContractV2HierarchyConfigurationParentHierarchyConfigurationChild = shared.ContractV2HierarchyConfigurationParentHierarchyConfigurationChild

// This is an alias to an internal type.
type ContractV2HierarchyConfigurationParentHierarchyConfigurationParentBehavior = shared.ContractV2HierarchyConfigurationParentHierarchyConfigurationParentBehavior

// This is an alias to an internal type.
type ContractV2HierarchyConfigurationChildHierarchyConfigurationV2 = shared.ContractV2HierarchyConfigurationChildHierarchyConfigurationV2

// The single parent contract/customer for this child.
//
// This is an alias to an internal type.
type ContractV2HierarchyConfigurationChildHierarchyConfigurationV2Parent = shared.ContractV2HierarchyConfigurationChildHierarchyConfigurationV2Parent

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
// prices automatically. EXPLICIT prioritization requires specifying priorities for
// each multiplier; the one with the lowest priority value will be prioritized
// first.
//
// This is an alias to an internal type.
type ContractV2MultiplierOverridePrioritization = shared.ContractV2MultiplierOverridePrioritization

// Equals "LOWEST_MULTIPLIER"
const ContractV2MultiplierOverridePrioritizationLowestMultiplier = shared.ContractV2MultiplierOverridePrioritizationLowestMultiplier

// Equals "EXPLICIT"
const ContractV2MultiplierOverridePrioritizationExplicit = shared.ContractV2MultiplierOverridePrioritizationExplicit

// This is an alias to an internal type.
type ContractV2RecurringCommit = shared.ContractV2RecurringCommit

// The amount of commit to grant.
//
// This is an alias to an internal type.
type ContractV2RecurringCommitAccessAmount = shared.ContractV2RecurringCommitAccessAmount

// The amount of time the created commits will be valid for
//
// This is an alias to an internal type.
type ContractV2RecurringCommitCommitDuration = shared.ContractV2RecurringCommitCommitDuration

// This is an alias to an internal type.
type ContractV2RecurringCommitProduct = shared.ContractV2RecurringCommitProduct

// This is an alias to an internal type.
type ContractV2RecurringCommitContract = shared.ContractV2RecurringCommitContract

// The amount the customer should be billed for the commit. Not required.
//
// This is an alias to an internal type.
type ContractV2RecurringCommitInvoiceAmount = shared.ContractV2RecurringCommitInvoiceAmount

// This is an alias to an internal type.
type ContractV2RecurringCredit = shared.ContractV2RecurringCredit

// The amount of commit to grant.
//
// This is an alias to an internal type.
type ContractV2RecurringCreditAccessAmount = shared.ContractV2RecurringCreditAccessAmount

// The amount of time the created commits will be valid for
//
// This is an alias to an internal type.
type ContractV2RecurringCreditCommitDuration = shared.ContractV2RecurringCreditCommitDuration

// This is an alias to an internal type.
type ContractV2RecurringCreditProduct = shared.ContractV2RecurringCreditProduct

// This is an alias to an internal type.
type ContractV2RecurringCreditContract = shared.ContractV2RecurringCreditContract

// This is an alias to an internal type.
type ContractV2ResellerRoyalty = shared.ContractV2ResellerRoyalty

// This is an alias to an internal type.
type ContractV2ResellerRoyaltySegment = shared.ContractV2ResellerRoyaltySegment

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
//
// This is an alias to an internal type.
type ContractV2ScheduledChargesOnUsageInvoices = shared.ContractV2ScheduledChargesOnUsageInvoices

// Equals "ALL"
const ContractV2ScheduledChargesOnUsageInvoicesAll = shared.ContractV2ScheduledChargesOnUsageInvoicesAll

// This is an alias to an internal type.
type ContractWithoutAmendments = shared.ContractWithoutAmendments

// This is an alias to an internal type.
type ContractWithoutAmendmentsTransition = shared.ContractWithoutAmendmentsTransition

// This is an alias to an internal type.
type ContractWithoutAmendmentsUsageStatementSchedule = shared.ContractWithoutAmendmentsUsageStatementSchedule

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommit = shared.ContractWithoutAmendmentsRecurringCommit

// The amount of commit to grant.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitAccessAmount = shared.ContractWithoutAmendmentsRecurringCommitAccessAmount

// The amount of time the created commits will be valid for
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitCommitDuration = shared.ContractWithoutAmendmentsRecurringCommitCommitDuration

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitProduct = shared.ContractWithoutAmendmentsRecurringCommitProduct

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitContract = shared.ContractWithoutAmendmentsRecurringCommitContract

// The amount the customer should be billed for the commit. Not required.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCommitInvoiceAmount = shared.ContractWithoutAmendmentsRecurringCommitInvoiceAmount

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCredit = shared.ContractWithoutAmendmentsRecurringCredit

// The amount of commit to grant.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditAccessAmount = shared.ContractWithoutAmendmentsRecurringCreditAccessAmount

// The amount of time the created commits will be valid for
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditCommitDuration = shared.ContractWithoutAmendmentsRecurringCreditCommitDuration

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditProduct = shared.ContractWithoutAmendmentsRecurringCreditProduct

// This is an alias to an internal type.
type ContractWithoutAmendmentsRecurringCreditContract = shared.ContractWithoutAmendmentsRecurringCreditContract

// This is an alias to an internal type.
type ContractWithoutAmendmentsResellerRoyalty = shared.ContractWithoutAmendmentsResellerRoyalty

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
//
// This is an alias to an internal type.
type ContractWithoutAmendmentsScheduledChargesOnUsageInvoices = shared.ContractWithoutAmendmentsScheduledChargesOnUsageInvoices

// Equals "ALL"
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

// Equals "CREDIT"
const CreditTypeCredit = shared.CreditTypeCredit

// This is an alias to an internal type.
type CreditContract = shared.CreditContract

// This is an alias to an internal type.
type CreditLedgerUnion = shared.CreditLedgerUnion

// This is an alias to an internal type.
type CreditLedgerCreditSegmentStartLedgerEntry = shared.CreditLedgerCreditSegmentStartLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry = shared.CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditExpirationLedgerEntry = shared.CreditLedgerCreditExpirationLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditCanceledLedgerEntry = shared.CreditLedgerCreditCanceledLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditCreditedLedgerEntry = shared.CreditLedgerCreditCreditedLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditManualLedgerEntry = shared.CreditLedgerCreditManualLedgerEntry

// This is an alias to an internal type.
type CreditLedgerCreditSeatBasedAdjustmentLedgerEntry = shared.CreditLedgerCreditSeatBasedAdjustmentLedgerEntry

// This is an alias to an internal type.
type CreditRateType = shared.CreditRateType

// Equals "COMMIT_RATE"
const CreditRateTypeCommitRate = shared.CreditRateTypeCommitRate

// Equals "LIST_RATE"
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
type HierarchyConfigurationUnion = shared.HierarchyConfigurationUnion

// This is an alias to an internal type.
type HierarchyConfigurationParentHierarchyConfiguration = shared.HierarchyConfigurationParentHierarchyConfiguration

// This is an alias to an internal type.
type HierarchyConfigurationParentHierarchyConfigurationChild = shared.HierarchyConfigurationParentHierarchyConfigurationChild

// This is an alias to an internal type.
type HierarchyConfigurationParentHierarchyConfigurationParentBehavior = shared.HierarchyConfigurationParentHierarchyConfigurationParentBehavior

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
type OverrideProduct = shared.OverrideProduct

// This is an alias to an internal type.
type OverrideRateType = shared.OverrideRateType

// Equals "FLAT"
const OverrideRateTypeFlat = shared.OverrideRateTypeFlat

// Equals "PERCENTAGE"
const OverrideRateTypePercentage = shared.OverrideRateTypePercentage

// Equals "SUBSCRIPTION"
const OverrideRateTypeSubscription = shared.OverrideRateTypeSubscription

// Equals "TIERED"
const OverrideRateTypeTiered = shared.OverrideRateTypeTiered

// Equals "CUSTOM"
const OverrideRateTypeCustom = shared.OverrideRateTypeCustom

// This is an alias to an internal type.
type OverrideTarget = shared.OverrideTarget

// Equals "COMMIT_RATE"
const OverrideTargetCommitRate = shared.OverrideTargetCommitRate

// Equals "LIST_RATE"
const OverrideTargetListRate = shared.OverrideTargetListRate

// This is an alias to an internal type.
type OverrideType = shared.OverrideType

// Equals "OVERWRITE"
const OverrideTypeOverwrite = shared.OverrideTypeOverwrite

// Equals "MULTIPLIER"
const OverrideTypeMultiplier = shared.OverrideTypeMultiplier

// Equals "TIERED"
const OverrideTypeTiered = shared.OverrideTypeTiered

// This is an alias to an internal type.
type OverrideTier = shared.OverrideTier

// This is an alias to an internal type.
type OverwriteRate = shared.OverwriteRate

// This is an alias to an internal type.
type OverwriteRateRateType = shared.OverwriteRateRateType

// Equals "FLAT"
const OverwriteRateRateTypeFlat = shared.OverwriteRateRateTypeFlat

// Equals "PERCENTAGE"
const OverwriteRateRateTypePercentage = shared.OverwriteRateRateTypePercentage

// Equals "SUBSCRIPTION"
const OverwriteRateRateTypeSubscription = shared.OverwriteRateRateTypeSubscription

// Equals "TIERED"
const OverwriteRateRateTypeTiered = shared.OverwriteRateRateTypeTiered

// Equals "CUSTOM"
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

// Equals "NONE"
const PaymentGateConfigPaymentGateTypeNone = shared.PaymentGateConfigPaymentGateTypeNone

// Equals "STRIPE"
const PaymentGateConfigPaymentGateTypeStripe = shared.PaymentGateConfigPaymentGateTypeStripe

// Equals "EXTERNAL"
const PaymentGateConfigPaymentGateTypeExternal = shared.PaymentGateConfigPaymentGateTypeExternal

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type PaymentGateConfigPrecalculatedTaxConfig = shared.PaymentGateConfigPrecalculatedTaxConfig

// Only applicable if using STRIPE as your payment gate type.
//
// This is an alias to an internal type.
type PaymentGateConfigStripeConfig = shared.PaymentGateConfigStripeConfig

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
//
// This is an alias to an internal type.
type PaymentGateConfigTaxType = shared.PaymentGateConfigTaxType

// Equals "NONE"
const PaymentGateConfigTaxTypeNone = shared.PaymentGateConfigTaxTypeNone

// Equals "STRIPE"
const PaymentGateConfigTaxTypeStripe = shared.PaymentGateConfigTaxTypeStripe

// Equals "ANROK"
const PaymentGateConfigTaxTypeAnrok = shared.PaymentGateConfigTaxTypeAnrok

// Equals "PRECALCULATED"
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

// Equals "NONE"
const PaymentGateConfigV2PaymentGateTypeNone = shared.PaymentGateConfigV2PaymentGateTypeNone

// Equals "STRIPE"
const PaymentGateConfigV2PaymentGateTypeStripe = shared.PaymentGateConfigV2PaymentGateTypeStripe

// Equals "EXTERNAL"
const PaymentGateConfigV2PaymentGateTypeExternal = shared.PaymentGateConfigV2PaymentGateTypeExternal

// Only applicable if using PRECALCULATED as your tax type.
//
// This is an alias to an internal type.
type PaymentGateConfigV2PrecalculatedTaxConfig = shared.PaymentGateConfigV2PrecalculatedTaxConfig

// Only applicable if using STRIPE as your payment gateway type.
//
// This is an alias to an internal type.
type PaymentGateConfigV2StripeConfig = shared.PaymentGateConfigV2StripeConfig

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
//
// This is an alias to an internal type.
type PaymentGateConfigV2TaxType = shared.PaymentGateConfigV2TaxType

// Equals "NONE"
const PaymentGateConfigV2TaxTypeNone = shared.PaymentGateConfigV2TaxTypeNone

// Equals "STRIPE"
const PaymentGateConfigV2TaxTypeStripe = shared.PaymentGateConfigV2TaxTypeStripe

// Equals "ANROK"
const PaymentGateConfigV2TaxTypeAnrok = shared.PaymentGateConfigV2TaxTypeAnrok

// Equals "PRECALCULATED"
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

// Equals "FLAT"
const RateRateTypeFlat = shared.RateRateTypeFlat

// Equals "PERCENTAGE"
const RateRateTypePercentage = shared.RateRateTypePercentage

// Equals "SUBSCRIPTION"
const RateRateTypeSubscription = shared.RateRateTypeSubscription

// Equals "CUSTOM"
const RateRateTypeCustom = shared.RateRateTypeCustom

// Equals "TIERED"
const RateRateTypeTiered = shared.RateRateTypeTiered

// This is an alias to an internal type.
type RecurringCommitSubscriptionConfig = shared.RecurringCommitSubscriptionConfig

// This is an alias to an internal type.
type RecurringCommitSubscriptionConfigAllocation = shared.RecurringCommitSubscriptionConfigAllocation

// Equals "INDIVIDUAL"
const RecurringCommitSubscriptionConfigAllocationIndividual = shared.RecurringCommitSubscriptionConfigAllocationIndividual

// Equals "POOLED"
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

// Equals "ADVANCE"
const SubscriptionCollectionScheduleAdvance = shared.SubscriptionCollectionScheduleAdvance

// Equals "ARREARS"
const SubscriptionCollectionScheduleArrears = shared.SubscriptionCollectionScheduleArrears

// This is an alias to an internal type.
type SubscriptionProration = shared.SubscriptionProration

// Determines how the subscription's quantity is controlled. Defaults to
// QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified
// directly on the subscription. `initial_quantity` must be provided with this
// option. Compatible with recurring commits/credits that use POOLED allocation.
//
// This is an alias to an internal type.
type SubscriptionQuantityManagementMode = shared.SubscriptionQuantityManagementMode

// Equals "SEAT_BASED"
const SubscriptionQuantityManagementModeSeatBased = shared.SubscriptionQuantityManagementModeSeatBased

// Equals "QUANTITY_ONLY"
const SubscriptionQuantityManagementModeQuantityOnly = shared.SubscriptionQuantityManagementModeQuantityOnly

// This is an alias to an internal type.
type SubscriptionQuantitySchedule = shared.SubscriptionQuantitySchedule

// This is an alias to an internal type.
type SubscriptionSubscriptionRate = shared.SubscriptionSubscriptionRate

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
