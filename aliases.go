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
type CommitRolledOverFrom = shared.CommitRolledOverFrom

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
type CreditTypeData = shared.CreditTypeData

// This is an alias to an internal type.
type CreditTypeDataParam = shared.CreditTypeDataParam

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
