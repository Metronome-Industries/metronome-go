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

// Only applicable if using Stripe as your payment gateway through Metronome.
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
type ContractWithoutAmendmentsRecurringCommitsSpecifier = shared.ContractWithoutAmendmentsRecurringCommitsSpecifier

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
type ContractWithoutAmendmentsRecurringCreditsSpecifier = shared.ContractWithoutAmendmentsRecurringCreditsSpecifier

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

// Only applicable if using Stripe as your payment gateway through Metronome.
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
