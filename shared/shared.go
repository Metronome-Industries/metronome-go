// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/tidwall/gjson"
)

type Commit struct {
	ID      string        `json:"id,required" format:"uuid"`
	Product CommitProduct `json:"product,required"`
	Type    CommitType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule CommitAccessSchedule `json:"access_schedule"`
	// (DEPRECATED) Use access_schedule + invoice_schedule instead.
	Amount                float64           `json:"amount"`
	ApplicableContractIDs []string          `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string          `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string          `json:"applicable_product_tags"`
	Contract              CommitContract    `json:"contract"`
	CustomFields          map[string]string `json:"custom_fields"`
	Description           string            `json:"description"`
	// The contract that this commit will be billed on.
	InvoiceContract CommitInvoiceContract `json:"invoice_contract"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule SchedulePointInTime `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger []CommitLedger `json:"ledger"`
	Name   string         `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority         float64              `json:"priority"`
	RolledOverFrom   CommitRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction float64              `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string     `json:"salesforce_opportunity_id"`
	JSON                    commitJSON `json:"-"`
}

// commitJSON contains the JSON metadata for the struct [Commit]
type commitJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	Amount                  apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	InvoiceContract         apijson.Field
	InvoiceSchedule         apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	RolledOverFrom          apijson.Field
	RolloverFraction        apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *Commit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitJSON) RawJSON() string {
	return r.raw
}

func (r Commit) ImplementsContractListBalancesResponseData() {}

type CommitProduct struct {
	ID   string            `json:"id,required" format:"uuid"`
	Name string            `json:"name,required"`
	JSON commitProductJSON `json:"-"`
}

// commitProductJSON contains the JSON metadata for the struct [CommitProduct]
type commitProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitProductJSON) RawJSON() string {
	return r.raw
}

type CommitType string

const (
	CommitTypePrepaid  CommitType = "PREPAID"
	CommitTypePostpaid CommitType = "POSTPAID"
)

func (r CommitType) IsKnown() bool {
	switch r {
	case CommitTypePrepaid, CommitTypePostpaid:
		return true
	}
	return false
}

// The schedule that the customer will gain access to the credits purposed with
// this commit.
type CommitAccessSchedule struct {
	ScheduleItems []CommitAccessScheduleScheduleItem `json:"schedule_items,required"`
	CreditType    CreditType                         `json:"credit_type"`
	JSON          commitAccessScheduleJSON           `json:"-"`
}

// commitAccessScheduleJSON contains the JSON metadata for the struct
// [CommitAccessSchedule]
type commitAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CommitAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type CommitAccessScheduleScheduleItem struct {
	ID           string                               `json:"id,required" format:"uuid"`
	Amount       float64                              `json:"amount,required"`
	EndingBefore time.Time                            `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                            `json:"starting_at,required" format:"date-time"`
	JSON         commitAccessScheduleScheduleItemJSON `json:"-"`
}

// commitAccessScheduleScheduleItemJSON contains the JSON metadata for the struct
// [CommitAccessScheduleScheduleItem]
type commitAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CommitAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitAccessScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type CommitContract struct {
	ID   string             `json:"id,required" format:"uuid"`
	JSON commitContractJSON `json:"-"`
}

// commitContractJSON contains the JSON metadata for the struct [CommitContract]
type commitContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitContractJSON) RawJSON() string {
	return r.raw
}

// The contract that this commit will be billed on.
type CommitInvoiceContract struct {
	ID   string                    `json:"id,required" format:"uuid"`
	JSON commitInvoiceContractJSON `json:"-"`
}

// commitInvoiceContractJSON contains the JSON metadata for the struct
// [CommitInvoiceContract]
type commitInvoiceContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitInvoiceContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitInvoiceContractJSON) RawJSON() string {
	return r.raw
}

type CommitLedger struct {
	Type          CommitLedgerType `json:"type,required"`
	Timestamp     time.Time        `json:"timestamp,required" format:"date-time"`
	Amount        float64          `json:"amount,required"`
	SegmentID     string           `json:"segment_id" format:"uuid"`
	InvoiceID     string           `json:"invoice_id" format:"uuid"`
	NewContractID string           `json:"new_contract_id" format:"uuid"`
	Reason        string           `json:"reason"`
	JSON          commitLedgerJSON `json:"-"`
	union         CommitLedgerUnion
}

// commitLedgerJSON contains the JSON metadata for the struct [CommitLedger]
type commitLedgerJSON struct {
	Type          apijson.Field
	Timestamp     apijson.Field
	Amount        apijson.Field
	SegmentID     apijson.Field
	InvoiceID     apijson.Field
	NewContractID apijson.Field
	Reason        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r commitLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *CommitLedger) UnmarshalJSON(data []byte) (err error) {
	*r = CommitLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CommitLedgerUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.CommitLedgerPrepaidCommitSegmentStartLedgerEntry],
// [shared.CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [shared.CommitLedgerPrepaidCommitRolloverLedgerEntry],
// [shared.CommitLedgerPrepaidCommitExpirationLedgerEntry],
// [shared.CommitLedgerPrepaidCommitCanceledLedgerEntry],
// [shared.CommitLedgerPrepaidCommitCreditedLedgerEntry],
// [shared.CommitLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [shared.CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [shared.CommitLedgerPostpaidCommitRolloverLedgerEntry],
// [shared.CommitLedgerPostpaidCommitTrueupLedgerEntry],
// [shared.CommitLedgerPrepaidCommitManualLedgerEntry],
// [shared.CommitLedgerPostpaidCommitManualLedgerEntry],
// [shared.CommitLedgerPostpaidCommitExpirationLedgerEntry].
func (r CommitLedger) AsUnion() CommitLedgerUnion {
	return r.union
}

// Union satisfied by [shared.CommitLedgerPrepaidCommitSegmentStartLedgerEntry],
// [shared.CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [shared.CommitLedgerPrepaidCommitRolloverLedgerEntry],
// [shared.CommitLedgerPrepaidCommitExpirationLedgerEntry],
// [shared.CommitLedgerPrepaidCommitCanceledLedgerEntry],
// [shared.CommitLedgerPrepaidCommitCreditedLedgerEntry],
// [shared.CommitLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [shared.CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [shared.CommitLedgerPostpaidCommitRolloverLedgerEntry],
// [shared.CommitLedgerPostpaidCommitTrueupLedgerEntry],
// [shared.CommitLedgerPrepaidCommitManualLedgerEntry],
// [shared.CommitLedgerPostpaidCommitManualLedgerEntry] or
// [shared.CommitLedgerPostpaidCommitExpirationLedgerEntry].
type CommitLedgerUnion interface {
	implementsSharedCommitLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CommitLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitRolloverLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitInitialBalanceLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitRolloverLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitTrueupLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitExpirationLedgerEntry{}),
		},
	)
}

type CommitLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                              `json:"amount,required"`
	SegmentID string                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                            `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitSegmentStartLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitSegmentStartLedgerEntryJSON contains the JSON metadata
// for the struct [CommitLedgerPrepaidCommitSegmentStartLedgerEntry]
type commitLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitSegmentStartLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	CommitLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart CommitLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

func (r CommitLedgerPrepaidCommitSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                           `json:"amount,required"`
	InvoiceID string                                                            `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                         `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON contains the
// JSON metadata for the struct
// [CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsSharedCommitLedger() {
}

type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                          `json:"amount,required"`
	NewContractID string                                           `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                           `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                        `json:"timestamp,required" format:"date-time"`
	Type          CommitLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          commitLedgerPrepaidCommitRolloverLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitRolloverLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitRolloverLedgerEntry]
type commitLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitRolloverLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitRolloverLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	CommitLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover CommitLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

func (r CommitLedgerPrepaidCommitRolloverLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                            `json:"amount,required"`
	SegmentID string                                             `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                          `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitExpirationLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitExpirationLedgerEntryJSON contains the JSON metadata
// for the struct [CommitLedgerPrepaidCommitExpirationLedgerEntry]
type commitLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitExpirationLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	CommitLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration CommitLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

func (r CommitLedgerPrepaidCommitExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                          `json:"amount,required"`
	InvoiceID string                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                        `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitCanceledLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitCanceledLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitCanceledLedgerEntry]
type commitLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitCanceledLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	CommitLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled CommitLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

func (r CommitLedgerPrepaidCommitCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                          `json:"amount,required"`
	InvoiceID string                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                        `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitCreditedLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitCreditedLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitCreditedLedgerEntry]
type commitLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitCreditedLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	CommitLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited CommitLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

func (r CommitLedgerPrepaidCommitCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                 `json:"amount,required"`
	Timestamp time.Time                                               `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPostpaidCommitInitialBalanceLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitInitialBalanceLedgerEntryJSON contains the JSON
// metadata for the struct [CommitLedgerPostpaidCommitInitialBalanceLedgerEntry]
type commitLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitInitialBalanceLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	CommitLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

func (r CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                            `json:"amount,required"`
	InvoiceID string                                                             `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                             `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                          `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON contains the
// JSON metadata for the struct
// [CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsSharedCommitLedger() {
}

type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitRolloverLedgerEntry struct {
	Amount        float64                                           `json:"amount,required"`
	NewContractID string                                            `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                            `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                         `json:"timestamp,required" format:"date-time"`
	Type          CommitLedgerPostpaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          commitLedgerPostpaidCommitRolloverLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitRolloverLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPostpaidCommitRolloverLedgerEntry]
type commitLedgerPostpaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitRolloverLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitRolloverLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPostpaidCommitRolloverLedgerEntryType string

const (
	CommitLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover CommitLedgerPostpaidCommitRolloverLedgerEntryType = "POSTPAID_COMMIT_ROLLOVER"
)

func (r CommitLedgerPostpaidCommitRolloverLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                         `json:"amount,required"`
	InvoiceID string                                          `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                       `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPostpaidCommitTrueupLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitTrueupLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPostpaidCommitTrueupLedgerEntry]
type commitLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitTrueupLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitTrueupLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	CommitLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup CommitLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

func (r CommitLedgerPostpaidCommitTrueupLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitManualLedgerEntry struct {
	Amount    float64                                        `json:"amount,required"`
	Reason    string                                         `json:"reason,required"`
	Timestamp time.Time                                      `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitManualLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitManualLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitManualLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitManualLedgerEntry]
type commitLedgerPrepaidCommitManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitManualLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitManualLedgerEntryType string

const (
	CommitLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual CommitLedgerPrepaidCommitManualLedgerEntryType = "PREPAID_COMMIT_MANUAL"
)

func (r CommitLedgerPrepaidCommitManualLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitManualLedgerEntry struct {
	Amount    float64                                         `json:"amount,required"`
	Reason    string                                          `json:"reason,required"`
	Timestamp time.Time                                       `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPostpaidCommitManualLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPostpaidCommitManualLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitManualLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPostpaidCommitManualLedgerEntry]
type commitLedgerPostpaidCommitManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitManualLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPostpaidCommitManualLedgerEntryType string

const (
	CommitLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual CommitLedgerPostpaidCommitManualLedgerEntryType = "POSTPAID_COMMIT_MANUAL"
)

func (r CommitLedgerPostpaidCommitManualLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitExpirationLedgerEntry struct {
	Amount    float64                                             `json:"amount,required"`
	Timestamp time.Time                                           `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPostpaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPostpaidCommitExpirationLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitExpirationLedgerEntryJSON contains the JSON metadata
// for the struct [CommitLedgerPostpaidCommitExpirationLedgerEntry]
type commitLedgerPostpaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitExpirationLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPostpaidCommitExpirationLedgerEntryType string

const (
	CommitLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration CommitLedgerPostpaidCommitExpirationLedgerEntryType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r CommitLedgerPostpaidCommitExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type CommitLedgerType string

const (
	CommitLedgerTypePrepaidCommitSegmentStart               CommitLedgerType = "PREPAID_COMMIT_SEGMENT_START"
	CommitLedgerTypePrepaidCommitAutomatedInvoiceDeduction  CommitLedgerType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	CommitLedgerTypePrepaidCommitRollover                   CommitLedgerType = "PREPAID_COMMIT_ROLLOVER"
	CommitLedgerTypePrepaidCommitExpiration                 CommitLedgerType = "PREPAID_COMMIT_EXPIRATION"
	CommitLedgerTypePrepaidCommitCanceled                   CommitLedgerType = "PREPAID_COMMIT_CANCELED"
	CommitLedgerTypePrepaidCommitCredited                   CommitLedgerType = "PREPAID_COMMIT_CREDITED"
	CommitLedgerTypePostpaidCommitInitialBalance            CommitLedgerType = "POSTPAID_COMMIT_INITIAL_BALANCE"
	CommitLedgerTypePostpaidCommitAutomatedInvoiceDeduction CommitLedgerType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	CommitLedgerTypePostpaidCommitRollover                  CommitLedgerType = "POSTPAID_COMMIT_ROLLOVER"
	CommitLedgerTypePostpaidCommitTrueup                    CommitLedgerType = "POSTPAID_COMMIT_TRUEUP"
	CommitLedgerTypePrepaidCommitManual                     CommitLedgerType = "PREPAID_COMMIT_MANUAL"
	CommitLedgerTypePostpaidCommitManual                    CommitLedgerType = "POSTPAID_COMMIT_MANUAL"
	CommitLedgerTypePostpaidCommitExpiration                CommitLedgerType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r CommitLedgerType) IsKnown() bool {
	switch r {
	case CommitLedgerTypePrepaidCommitSegmentStart, CommitLedgerTypePrepaidCommitAutomatedInvoiceDeduction, CommitLedgerTypePrepaidCommitRollover, CommitLedgerTypePrepaidCommitExpiration, CommitLedgerTypePrepaidCommitCanceled, CommitLedgerTypePrepaidCommitCredited, CommitLedgerTypePostpaidCommitInitialBalance, CommitLedgerTypePostpaidCommitAutomatedInvoiceDeduction, CommitLedgerTypePostpaidCommitRollover, CommitLedgerTypePostpaidCommitTrueup, CommitLedgerTypePrepaidCommitManual, CommitLedgerTypePostpaidCommitManual, CommitLedgerTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type CommitRolledOverFrom struct {
	CommitID   string                   `json:"commit_id,required" format:"uuid"`
	ContractID string                   `json:"contract_id,required" format:"uuid"`
	JSON       commitRolledOverFromJSON `json:"-"`
}

// commitRolledOverFromJSON contains the JSON metadata for the struct
// [CommitRolledOverFrom]
type commitRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitRolledOverFromJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendments struct {
	Commits                []Commit                                        `json:"commits,required"`
	CreatedAt              time.Time                                       `json:"created_at,required" format:"date-time"`
	CreatedBy              string                                          `json:"created_by,required"`
	Overrides              []Override                                      `json:"overrides,required"`
	ScheduledCharges       []ScheduledCharge                               `json:"scheduled_charges,required"`
	StartingAt             time.Time                                       `json:"starting_at,required" format:"date-time"`
	Transitions            []ContractWithoutAmendmentsTransition           `json:"transitions,required"`
	UsageStatementSchedule ContractWithoutAmendmentsUsageStatementSchedule `json:"usage_statement_schedule,required"`
	Credits                []ContractWithoutAmendmentsCredit               `json:"credits"`
	// This field's availability is dependent on your client's configuration.
	Discounts           []Discount `json:"discounts"`
	EndingBefore        time.Time  `json:"ending_before" format:"date-time"`
	Name                string     `json:"name"`
	NetPaymentTermsDays float64    `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []ContractWithoutAmendmentsProfessionalService `json:"professional_services"`
	RateCardID           string                                         `json:"rate_card_id" format:"uuid"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []ContractWithoutAmendmentsResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue float64                              `json:"total_contract_value"`
	UsageFilter        ContractWithoutAmendmentsUsageFilter `json:"usage_filter"`
	JSON               contractWithoutAmendmentsJSON        `json:"-"`
}

// contractWithoutAmendmentsJSON contains the JSON metadata for the struct
// [ContractWithoutAmendments]
type contractWithoutAmendmentsJSON struct {
	Commits                 apijson.Field
	CreatedAt               apijson.Field
	CreatedBy               apijson.Field
	Overrides               apijson.Field
	ScheduledCharges        apijson.Field
	StartingAt              apijson.Field
	Transitions             apijson.Field
	UsageStatementSchedule  apijson.Field
	Credits                 apijson.Field
	Discounts               apijson.Field
	EndingBefore            apijson.Field
	Name                    apijson.Field
	NetPaymentTermsDays     apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	ProfessionalServices    apijson.Field
	RateCardID              apijson.Field
	ResellerRoyalties       apijson.Field
	SalesforceOpportunityID apijson.Field
	TotalContractValue      apijson.Field
	UsageFilter             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractWithoutAmendments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsTransition struct {
	FromContractID string                                   `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                   `json:"to_contract_id,required" format:"uuid"`
	Type           ContractWithoutAmendmentsTransitionsType `json:"type,required"`
	JSON           contractWithoutAmendmentsTransitionJSON  `json:"-"`
}

// contractWithoutAmendmentsTransitionJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsTransition]
type contractWithoutAmendmentsTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsTransitionJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsTransitionsType string

const (
	ContractWithoutAmendmentsTransitionsTypeSupersede ContractWithoutAmendmentsTransitionsType = "SUPERSEDE"
	ContractWithoutAmendmentsTransitionsTypeRenewal   ContractWithoutAmendmentsTransitionsType = "RENEWAL"
)

func (r ContractWithoutAmendmentsTransitionsType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsTransitionsTypeSupersede, ContractWithoutAmendmentsTransitionsTypeRenewal:
		return true
	}
	return false
}

type ContractWithoutAmendmentsUsageStatementSchedule struct {
	Frequency ContractWithoutAmendmentsUsageStatementScheduleFrequency `json:"frequency,required"`
	JSON      contractWithoutAmendmentsUsageStatementScheduleJSON      `json:"-"`
}

// contractWithoutAmendmentsUsageStatementScheduleJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsUsageStatementSchedule]
type contractWithoutAmendmentsUsageStatementScheduleJSON struct {
	Frequency   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageStatementSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsUsageStatementScheduleJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsUsageStatementScheduleFrequency string

const (
	ContractWithoutAmendmentsUsageStatementScheduleFrequencyMonthly   ContractWithoutAmendmentsUsageStatementScheduleFrequency = "MONTHLY"
	ContractWithoutAmendmentsUsageStatementScheduleFrequencyMonthly   ContractWithoutAmendmentsUsageStatementScheduleFrequency = "monthly"
	ContractWithoutAmendmentsUsageStatementScheduleFrequencyQuarterly ContractWithoutAmendmentsUsageStatementScheduleFrequency = "QUARTERLY"
	ContractWithoutAmendmentsUsageStatementScheduleFrequencyQuarterly ContractWithoutAmendmentsUsageStatementScheduleFrequency = "quarterly"
)

func (r ContractWithoutAmendmentsUsageStatementScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsUsageStatementScheduleFrequencyMonthly, ContractWithoutAmendmentsUsageStatementScheduleFrequencyMonthly, ContractWithoutAmendmentsUsageStatementScheduleFrequencyQuarterly, ContractWithoutAmendmentsUsageStatementScheduleFrequencyQuarterly:
		return true
	}
	return false
}

type ContractWithoutAmendmentsCredit struct {
	ID      string                                  `json:"id,required" format:"uuid"`
	Product ContractWithoutAmendmentsCreditsProduct `json:"product,required"`
	Type    ContractWithoutAmendmentsCreditsType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        ContractWithoutAmendmentsCreditsAccessSchedule `json:"access_schedule"`
	ApplicableContractIDs []string                                       `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string                                       `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                                       `json:"applicable_product_tags"`
	Contract              ContractWithoutAmendmentsCreditsContract       `json:"contract"`
	CustomFields          map[string]string                              `json:"custom_fields"`
	Description           string                                         `json:"description"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []ContractWithoutAmendmentsCreditsLedger `json:"ledger"`
	Name   string                                   `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                              `json:"salesforce_opportunity_id"`
	JSON                    contractWithoutAmendmentsCreditJSON `json:"-"`
}

// contractWithoutAmendmentsCreditJSON contains the JSON metadata for the struct
// [ContractWithoutAmendmentsCredit]
type contractWithoutAmendmentsCreditJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsCreditJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsCreditsProduct struct {
	ID   string                                      `json:"id,required" format:"uuid"`
	Name string                                      `json:"name,required"`
	JSON contractWithoutAmendmentsCreditsProductJSON `json:"-"`
}

// contractWithoutAmendmentsCreditsProductJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsCreditsProduct]
type contractWithoutAmendmentsCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsCreditsProductJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsCreditsType string

const (
	ContractWithoutAmendmentsCreditsTypeCredit ContractWithoutAmendmentsCreditsType = "CREDIT"
)

func (r ContractWithoutAmendmentsCreditsType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsCreditsTypeCredit:
		return true
	}
	return false
}

// The schedule that the customer will gain access to the credits.
type ContractWithoutAmendmentsCreditsAccessSchedule struct {
	ScheduleItems []ContractWithoutAmendmentsCreditsAccessScheduleScheduleItem `json:"schedule_items,required"`
	CreditType    CreditType                                                   `json:"credit_type"`
	JSON          contractWithoutAmendmentsCreditsAccessScheduleJSON           `json:"-"`
}

// contractWithoutAmendmentsCreditsAccessScheduleJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsCreditsAccessSchedule]
type contractWithoutAmendmentsCreditsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsCreditsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsCreditsAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsCreditsAccessScheduleScheduleItem struct {
	ID           string                                                         `json:"id,required" format:"uuid"`
	Amount       float64                                                        `json:"amount,required"`
	EndingBefore time.Time                                                      `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                                                      `json:"starting_at,required" format:"date-time"`
	JSON         contractWithoutAmendmentsCreditsAccessScheduleScheduleItemJSON `json:"-"`
}

// contractWithoutAmendmentsCreditsAccessScheduleScheduleItemJSON contains the JSON
// metadata for the struct
// [ContractWithoutAmendmentsCreditsAccessScheduleScheduleItem]
type contractWithoutAmendmentsCreditsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsCreditsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsCreditsAccessScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsCreditsContract struct {
	ID   string                                       `json:"id,required" format:"uuid"`
	JSON contractWithoutAmendmentsCreditsContractJSON `json:"-"`
}

// contractWithoutAmendmentsCreditsContractJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsCreditsContract]
type contractWithoutAmendmentsCreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsCreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsCreditsContractJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsCreditsLedger struct {
	Type      ContractWithoutAmendmentsCreditsLedgerType `json:"type,required"`
	Timestamp time.Time                                  `json:"timestamp,required" format:"date-time"`
	Amount    float64                                    `json:"amount,required"`
	SegmentID string                                     `json:"segment_id" format:"uuid"`
	InvoiceID string                                     `json:"invoice_id" format:"uuid"`
	Reason    string                                     `json:"reason"`
	JSON      contractWithoutAmendmentsCreditsLedgerJSON `json:"-"`
	union     ContractWithoutAmendmentsCreditsLedgerUnion
}

// contractWithoutAmendmentsCreditsLedgerJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsCreditsLedger]
type contractWithoutAmendmentsCreditsLedgerJSON struct {
	Type        apijson.Field
	Timestamp   apijson.Field
	Amount      apijson.Field
	SegmentID   apijson.Field
	InvoiceID   apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r contractWithoutAmendmentsCreditsLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *ContractWithoutAmendmentsCreditsLedger) UnmarshalJSON(data []byte) (err error) {
	*r = ContractWithoutAmendmentsCreditsLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContractWithoutAmendmentsCreditsLedgerUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry],
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntry],
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntry],
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntry],
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntry].
func (r ContractWithoutAmendmentsCreditsLedger) AsUnion() ContractWithoutAmendmentsCreditsLedgerUnion {
	return r.union
}

// Union satisfied by
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry],
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntry],
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntry],
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntry] or
// [shared.ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntry].
type ContractWithoutAmendmentsCreditsLedgerUnion interface {
	implementsSharedContractWithoutAmendmentsCreditsLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContractWithoutAmendmentsCreditsLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntry{}),
		},
	)
}

type ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64                                                                 `json:"amount,required"`
	SegmentID string                                                                  `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                               `json:"timestamp,required" format:"date-time"`
	Type      ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON `json:"-"`
}

// contractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON contains
// the JSON metadata for the struct
// [ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry]
type contractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry) implementsSharedContractWithoutAmendmentsCreditsLedger() {
}

type ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType string

const (
	ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType = "CREDIT_SEGMENT_START"
)

func (r ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart:
		return true
	}
	return false
}

type ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                              `json:"amount,required"`
	InvoiceID string                                                                               `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                            `json:"timestamp,required" format:"date-time"`
	Type      ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// contractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry]
type contractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) implementsSharedContractWithoutAmendmentsCreditsLedger() {
}

type ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntry struct {
	Amount    float64                                                               `json:"amount,required"`
	SegmentID string                                                                `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                             `json:"timestamp,required" format:"date-time"`
	Type      ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntryType `json:"type,required"`
	JSON      contractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON `json:"-"`
}

// contractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON contains
// the JSON metadata for the struct
// [ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntry]
type contractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntry) implementsSharedContractWithoutAmendmentsCreditsLedger() {
}

type ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntryType string

const (
	ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntryType = "CREDIT_EXPIRATION"
)

func (r ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsCreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration:
		return true
	}
	return false
}

type ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntry struct {
	Amount    float64                                                             `json:"amount,required"`
	InvoiceID string                                                              `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                              `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                           `json:"timestamp,required" format:"date-time"`
	Type      ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntryType `json:"type,required"`
	JSON      contractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON `json:"-"`
}

// contractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON contains the
// JSON metadata for the struct
// [ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntry]
type contractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntry) implementsSharedContractWithoutAmendmentsCreditsLedger() {
}

type ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntryType string

const (
	ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntryType = "CREDIT_CANCELED"
)

func (r ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsCreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled:
		return true
	}
	return false
}

type ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntry struct {
	Amount    float64                                                             `json:"amount,required"`
	InvoiceID string                                                              `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                              `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                           `json:"timestamp,required" format:"date-time"`
	Type      ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntryType `json:"type,required"`
	JSON      contractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON `json:"-"`
}

// contractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON contains the
// JSON metadata for the struct
// [ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntry]
type contractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntry) implementsSharedContractWithoutAmendmentsCreditsLedger() {
}

type ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntryType string

const (
	ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntryType = "CREDIT_CREDITED"
)

func (r ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsCreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited:
		return true
	}
	return false
}

type ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntry struct {
	Amount    float64                                                           `json:"amount,required"`
	Reason    string                                                            `json:"reason,required"`
	Timestamp time.Time                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntryType `json:"type,required"`
	JSON      contractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntryJSON `json:"-"`
}

// contractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntryJSON contains the
// JSON metadata for the struct
// [ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntry]
type contractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntry) implementsSharedContractWithoutAmendmentsCreditsLedger() {
}

type ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntryType string

const (
	ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntryTypeCreditManual ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntryType = "CREDIT_MANUAL"
)

func (r ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsCreditsLedgerCreditManualLedgerEntryTypeCreditManual:
		return true
	}
	return false
}

type ContractWithoutAmendmentsCreditsLedgerType string

const (
	ContractWithoutAmendmentsCreditsLedgerTypeCreditSegmentStart              ContractWithoutAmendmentsCreditsLedgerType = "CREDIT_SEGMENT_START"
	ContractWithoutAmendmentsCreditsLedgerTypeCreditAutomatedInvoiceDeduction ContractWithoutAmendmentsCreditsLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	ContractWithoutAmendmentsCreditsLedgerTypeCreditExpiration                ContractWithoutAmendmentsCreditsLedgerType = "CREDIT_EXPIRATION"
	ContractWithoutAmendmentsCreditsLedgerTypeCreditCanceled                  ContractWithoutAmendmentsCreditsLedgerType = "CREDIT_CANCELED"
	ContractWithoutAmendmentsCreditsLedgerTypeCreditCredited                  ContractWithoutAmendmentsCreditsLedgerType = "CREDIT_CREDITED"
	ContractWithoutAmendmentsCreditsLedgerTypeCreditManual                    ContractWithoutAmendmentsCreditsLedgerType = "CREDIT_MANUAL"
)

func (r ContractWithoutAmendmentsCreditsLedgerType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsCreditsLedgerTypeCreditSegmentStart, ContractWithoutAmendmentsCreditsLedgerTypeCreditAutomatedInvoiceDeduction, ContractWithoutAmendmentsCreditsLedgerTypeCreditExpiration, ContractWithoutAmendmentsCreditsLedgerTypeCreditCanceled, ContractWithoutAmendmentsCreditsLedgerTypeCreditCredited, ContractWithoutAmendmentsCreditsLedgerTypeCreditManual:
		return true
	}
	return false
}

type ContractWithoutAmendmentsProfessionalService struct {
	ID string `json:"id,required" format:"uuid"`
	// Maximum amount for the term.
	MaxAmount float64 `json:"max_amount,required"`
	ProductID string  `json:"product_id,required" format:"uuid"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity float64 `json:"quantity,required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified.
	UnitPrice    float64           `json:"unit_price,required"`
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string                                           `json:"netsuite_sales_order_id"`
	JSON                 contractWithoutAmendmentsProfessionalServiceJSON `json:"-"`
}

// contractWithoutAmendmentsProfessionalServiceJSON contains the JSON metadata for
// the struct [ContractWithoutAmendmentsProfessionalService]
type contractWithoutAmendmentsProfessionalServiceJSON struct {
	ID                   apijson.Field
	MaxAmount            apijson.Field
	ProductID            apijson.Field
	Quantity             apijson.Field
	UnitPrice            apijson.Field
	CustomFields         apijson.Field
	Description          apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsProfessionalService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsProfessionalServiceJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsResellerRoyalty struct {
	Fraction              float64                                                `json:"fraction,required"`
	NetsuiteResellerID    string                                                 `json:"netsuite_reseller_id,required"`
	ResellerType          ContractWithoutAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                              `json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  []string                                               `json:"applicable_product_ids"`
	ApplicableProductTags []string                                               `json:"applicable_product_tags"`
	AwsAccountNumber      string                                                 `json:"aws_account_number"`
	AwsOfferID            string                                                 `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                 `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                              `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                 `json:"gcp_account_id"`
	GcpOfferID            string                                                 `json:"gcp_offer_id"`
	ResellerContractValue float64                                                `json:"reseller_contract_value"`
	JSON                  contractWithoutAmendmentsResellerRoyaltyJSON           `json:"-"`
}

// contractWithoutAmendmentsResellerRoyaltyJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsResellerRoyalty]
type contractWithoutAmendmentsResellerRoyaltyJSON struct {
	Fraction              apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerType          apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	ResellerContractValue apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsResellerRoyaltiesResellerType string

const (
	ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAws           ContractWithoutAmendmentsResellerRoyaltiesResellerType = "AWS"
	ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAwsProService ContractWithoutAmendmentsResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcp           ContractWithoutAmendmentsResellerRoyaltiesResellerType = "GCP"
	ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcpProService ContractWithoutAmendmentsResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r ContractWithoutAmendmentsResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAws, ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAwsProService, ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcp, ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

type ContractWithoutAmendmentsUsageFilter struct {
	Current ContractWithoutAmendmentsUsageFilterCurrent  `json:"current,required,nullable"`
	Initial ContractWithoutAmendmentsUsageFilterInitial  `json:"initial,required"`
	Updates []ContractWithoutAmendmentsUsageFilterUpdate `json:"updates,required"`
	JSON    contractWithoutAmendmentsUsageFilterJSON     `json:"-"`
}

// contractWithoutAmendmentsUsageFilterJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsUsageFilter]
type contractWithoutAmendmentsUsageFilterJSON struct {
	Current     apijson.Field
	Initial     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsUsageFilterJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsUsageFilterCurrent struct {
	GroupKey    string                                          `json:"group_key,required"`
	GroupValues []string                                        `json:"group_values,required"`
	StartingAt  time.Time                                       `json:"starting_at" format:"date-time"`
	JSON        contractWithoutAmendmentsUsageFilterCurrentJSON `json:"-"`
}

// contractWithoutAmendmentsUsageFilterCurrentJSON contains the JSON metadata for
// the struct [ContractWithoutAmendmentsUsageFilterCurrent]
type contractWithoutAmendmentsUsageFilterCurrentJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsUsageFilterCurrentJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsUsageFilterInitial struct {
	GroupKey    string                                          `json:"group_key,required"`
	GroupValues []string                                        `json:"group_values,required"`
	StartingAt  time.Time                                       `json:"starting_at" format:"date-time"`
	JSON        contractWithoutAmendmentsUsageFilterInitialJSON `json:"-"`
}

// contractWithoutAmendmentsUsageFilterInitialJSON contains the JSON metadata for
// the struct [ContractWithoutAmendmentsUsageFilterInitial]
type contractWithoutAmendmentsUsageFilterInitialJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsUsageFilterInitialJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsUsageFilterUpdate struct {
	GroupKey    string                                         `json:"group_key,required"`
	GroupValues []string                                       `json:"group_values,required"`
	StartingAt  time.Time                                      `json:"starting_at,required" format:"date-time"`
	JSON        contractWithoutAmendmentsUsageFilterUpdateJSON `json:"-"`
}

// contractWithoutAmendmentsUsageFilterUpdateJSON contains the JSON metadata for
// the struct [ContractWithoutAmendmentsUsageFilterUpdate]
type contractWithoutAmendmentsUsageFilterUpdateJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsUsageFilterUpdateJSON) RawJSON() string {
	return r.raw
}

type CreditType struct {
	ID   string         `json:"id,required" format:"uuid"`
	Name string         `json:"name,required"`
	JSON creditTypeJSON `json:"-"`
}

// creditTypeJSON contains the JSON metadata for the struct [CreditType]
type creditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditTypeJSON) RawJSON() string {
	return r.raw
}

type Discount struct {
	ID       string              `json:"id,required" format:"uuid"`
	Product  DiscountProduct     `json:"product,required"`
	Schedule SchedulePointInTime `json:"schedule,required"`
	Name     string              `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string       `json:"netsuite_sales_order_id"`
	JSON                 discountJSON `json:"-"`
}

// discountJSON contains the JSON metadata for the struct [Discount]
type discountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Discount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discountJSON) RawJSON() string {
	return r.raw
}

type DiscountProduct struct {
	ID   string              `json:"id,required" format:"uuid"`
	Name string              `json:"name,required"`
	JSON discountProductJSON `json:"-"`
}

// discountProductJSON contains the JSON metadata for the struct [DiscountProduct]
type discountProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscountProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discountProductJSON) RawJSON() string {
	return r.raw
}

type ID struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON idJSON `json:"-"`
}

// idJSON contains the JSON metadata for the struct [ID]
type idJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idJSON) RawJSON() string {
	return r.raw
}

type IDParam struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r IDParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Override struct {
	ID                    string     `json:"id,required" format:"uuid"`
	StartingAt            time.Time  `json:"starting_at,required" format:"date-time"`
	ApplicableProductTags []string   `json:"applicable_product_tags"`
	CreditType            CreditType `json:"credit_type"`
	EndingBefore          time.Time  `json:"ending_before" format:"date-time"`
	Entitled              bool       `json:"entitled"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type.
	IsProrated         bool                        `json:"is_prorated"`
	Multiplier         float64                     `json:"multiplier"`
	OverrideSpecifiers []OverrideOverrideSpecifier `json:"override_specifiers"`
	OverrideTiers      []OverrideOverrideTier      `json:"override_tiers"`
	OverwriteRate      OverrideOverwriteRate       `json:"overwrite_rate"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64         `json:"price"`
	Priority float64         `json:"priority"`
	Product  OverrideProduct `json:"product"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64          `json:"quantity"`
	RateType OverrideRateType `json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []OverrideTier `json:"tiers"`
	Type  OverrideType   `json:"type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	Value map[string]interface{} `json:"value"`
	JSON  overrideJSON           `json:"-"`
}

// overrideJSON contains the JSON metadata for the struct [Override]
type overrideJSON struct {
	ID                    apijson.Field
	StartingAt            apijson.Field
	ApplicableProductTags apijson.Field
	CreditType            apijson.Field
	EndingBefore          apijson.Field
	Entitled              apijson.Field
	IsProrated            apijson.Field
	Multiplier            apijson.Field
	OverrideSpecifiers    apijson.Field
	OverrideTiers         apijson.Field
	OverwriteRate         apijson.Field
	Price                 apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	Quantity              apijson.Field
	RateType              apijson.Field
	Tiers                 apijson.Field
	Type                  apijson.Field
	Value                 apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideJSON) RawJSON() string {
	return r.raw
}

type OverrideOverrideSpecifier struct {
	PresentationGroupValues map[string]string             `json:"presentation_group_values"`
	PricingGroupValues      map[string]string             `json:"pricing_group_values"`
	ProductID               string                        `json:"product_id" format:"uuid"`
	ProductTags             []string                      `json:"product_tags"`
	JSON                    overrideOverrideSpecifierJSON `json:"-"`
}

// overrideOverrideSpecifierJSON contains the JSON metadata for the struct
// [OverrideOverrideSpecifier]
type overrideOverrideSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OverrideOverrideSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideOverrideSpecifierJSON) RawJSON() string {
	return r.raw
}

type OverrideOverrideTier struct {
	Multiplier float64                  `json:"multiplier,required"`
	Size       float64                  `json:"size"`
	JSON       overrideOverrideTierJSON `json:"-"`
}

// overrideOverrideTierJSON contains the JSON metadata for the struct
// [OverrideOverrideTier]
type overrideOverrideTierJSON struct {
	Multiplier  apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideOverrideTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideOverrideTierJSON) RawJSON() string {
	return r.raw
}

type OverrideOverwriteRate struct {
	RateType   OverrideOverwriteRateRateType `json:"rate_type,required"`
	CreditType CreditType                    `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]interface{} `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []OverrideOverwriteRateTier `json:"tiers"`
	JSON  overrideOverwriteRateJSON   `json:"-"`
}

// overrideOverwriteRateJSON contains the JSON metadata for the struct
// [OverrideOverwriteRate]
type overrideOverwriteRateJSON struct {
	RateType    apijson.Field
	CreditType  apijson.Field
	CustomRate  apijson.Field
	IsProrated  apijson.Field
	Price       apijson.Field
	Quantity    apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideOverwriteRateJSON) RawJSON() string {
	return r.raw
}

type OverrideOverwriteRateRateType string

const (
	OverrideOverwriteRateRateTypeFlat         OverrideOverwriteRateRateType = "FLAT"
	OverrideOverwriteRateRateTypeFlat         OverrideOverwriteRateRateType = "flat"
	OverrideOverwriteRateRateTypePercentage   OverrideOverwriteRateRateType = "PERCENTAGE"
	OverrideOverwriteRateRateTypePercentage   OverrideOverwriteRateRateType = "percentage"
	OverrideOverwriteRateRateTypeSubscription OverrideOverwriteRateRateType = "SUBSCRIPTION"
	OverrideOverwriteRateRateTypeSubscription OverrideOverwriteRateRateType = "subscription"
	OverrideOverwriteRateRateTypeTiered       OverrideOverwriteRateRateType = "TIERED"
	OverrideOverwriteRateRateTypeTiered       OverrideOverwriteRateRateType = "tiered"
	OverrideOverwriteRateRateTypeCustom       OverrideOverwriteRateRateType = "CUSTOM"
	OverrideOverwriteRateRateTypeCustom       OverrideOverwriteRateRateType = "custom"
)

func (r OverrideOverwriteRateRateType) IsKnown() bool {
	switch r {
	case OverrideOverwriteRateRateTypeFlat, OverrideOverwriteRateRateTypeFlat, OverrideOverwriteRateRateTypePercentage, OverrideOverwriteRateRateTypePercentage, OverrideOverwriteRateRateTypeSubscription, OverrideOverwriteRateRateTypeSubscription, OverrideOverwriteRateRateTypeTiered, OverrideOverwriteRateRateTypeTiered, OverrideOverwriteRateRateTypeCustom, OverrideOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

type OverrideOverwriteRateTier struct {
	Price float64                       `json:"price,required"`
	Size  float64                       `json:"size"`
	JSON  overrideOverwriteRateTierJSON `json:"-"`
}

// overrideOverwriteRateTierJSON contains the JSON metadata for the struct
// [OverrideOverwriteRateTier]
type overrideOverwriteRateTierJSON struct {
	Price       apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideOverwriteRateTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideOverwriteRateTierJSON) RawJSON() string {
	return r.raw
}

type OverrideProduct struct {
	ID   string              `json:"id,required" format:"uuid"`
	Name string              `json:"name,required"`
	JSON overrideProductJSON `json:"-"`
}

// overrideProductJSON contains the JSON metadata for the struct [OverrideProduct]
type overrideProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideProductJSON) RawJSON() string {
	return r.raw
}

type OverrideRateType string

const (
	OverrideRateTypeFlat         OverrideRateType = "FLAT"
	OverrideRateTypeFlat         OverrideRateType = "flat"
	OverrideRateTypePercentage   OverrideRateType = "PERCENTAGE"
	OverrideRateTypePercentage   OverrideRateType = "percentage"
	OverrideRateTypeSubscription OverrideRateType = "SUBSCRIPTION"
	OverrideRateTypeSubscription OverrideRateType = "subscription"
	OverrideRateTypeTiered       OverrideRateType = "TIERED"
	OverrideRateTypeTiered       OverrideRateType = "tiered"
	OverrideRateTypeCustom       OverrideRateType = "CUSTOM"
	OverrideRateTypeCustom       OverrideRateType = "custom"
)

func (r OverrideRateType) IsKnown() bool {
	switch r {
	case OverrideRateTypeFlat, OverrideRateTypeFlat, OverrideRateTypePercentage, OverrideRateTypePercentage, OverrideRateTypeSubscription, OverrideRateTypeSubscription, OverrideRateTypeTiered, OverrideRateTypeTiered, OverrideRateTypeCustom, OverrideRateTypeCustom:
		return true
	}
	return false
}

type OverrideTier struct {
	Price float64          `json:"price,required"`
	Size  float64          `json:"size"`
	JSON  overrideTierJSON `json:"-"`
}

// overrideTierJSON contains the JSON metadata for the struct [OverrideTier]
type overrideTierJSON struct {
	Price       apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideTierJSON) RawJSON() string {
	return r.raw
}

type OverrideType string

const (
	OverrideTypeOverwrite  OverrideType = "OVERWRITE"
	OverrideTypeMultiplier OverrideType = "MULTIPLIER"
	OverrideTypeTiered     OverrideType = "TIERED"
)

func (r OverrideType) IsKnown() bool {
	switch r {
	case OverrideTypeOverwrite, OverrideTypeMultiplier, OverrideTypeTiered:
		return true
	}
	return false
}

type Rate struct {
	RateType   RateRateType `json:"rate_type,required"`
	CreditType CreditType   `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]interface{} `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// if pricing groups are used, this will contain the values used to calculate the
	// price
	PricingGroupValues map[string]string `json:"pricing_group_values"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []RateTier `json:"tiers"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool     `json:"use_list_prices"`
	JSON          rateJSON `json:"-"`
}

// rateJSON contains the JSON metadata for the struct [Rate]
type rateJSON struct {
	RateType           apijson.Field
	CreditType         apijson.Field
	CustomRate         apijson.Field
	IsProrated         apijson.Field
	Price              apijson.Field
	PricingGroupValues apijson.Field
	Quantity           apijson.Field
	Tiers              apijson.Field
	UseListPrices      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Rate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateJSON) RawJSON() string {
	return r.raw
}

type RateRateType string

const (
	RateRateTypeFlat         RateRateType = "FLAT"
	RateRateTypeFlat         RateRateType = "flat"
	RateRateTypePercentage   RateRateType = "PERCENTAGE"
	RateRateTypePercentage   RateRateType = "percentage"
	RateRateTypeSubscription RateRateType = "SUBSCRIPTION"
	RateRateTypeSubscription RateRateType = "subscription"
	RateRateTypeCustom       RateRateType = "CUSTOM"
	RateRateTypeCustom       RateRateType = "custom"
	RateRateTypeTiered       RateRateType = "TIERED"
	RateRateTypeTiered       RateRateType = "tiered"
)

func (r RateRateType) IsKnown() bool {
	switch r {
	case RateRateTypeFlat, RateRateTypeFlat, RateRateTypePercentage, RateRateTypePercentage, RateRateTypeSubscription, RateRateTypeSubscription, RateRateTypeCustom, RateRateTypeCustom, RateRateTypeTiered, RateRateTypeTiered:
		return true
	}
	return false
}

type RateTier struct {
	Price float64      `json:"price,required"`
	Size  float64      `json:"size"`
	JSON  rateTierJSON `json:"-"`
}

// rateTierJSON contains the JSON metadata for the struct [RateTier]
type rateTierJSON struct {
	Price       apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateTierJSON) RawJSON() string {
	return r.raw
}

type SchedulePointInTime struct {
	CreditType    CreditType                        `json:"credit_type"`
	ScheduleItems []SchedulePointInTimeScheduleItem `json:"schedule_items"`
	JSON          schedulePointInTimeJSON           `json:"-"`
}

// schedulePointInTimeJSON contains the JSON metadata for the struct
// [SchedulePointInTime]
type schedulePointInTimeJSON struct {
	CreditType    apijson.Field
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SchedulePointInTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schedulePointInTimeJSON) RawJSON() string {
	return r.raw
}

type SchedulePointInTimeScheduleItem struct {
	ID        string                              `json:"id,required" format:"uuid"`
	Amount    float64                             `json:"amount,required"`
	InvoiceID string                              `json:"invoice_id,required" format:"uuid"`
	Quantity  float64                             `json:"quantity,required"`
	Timestamp time.Time                           `json:"timestamp,required" format:"date-time"`
	UnitPrice float64                             `json:"unit_price,required"`
	JSON      schedulePointInTimeScheduleItemJSON `json:"-"`
}

// schedulePointInTimeScheduleItemJSON contains the JSON metadata for the struct
// [SchedulePointInTimeScheduleItem]
type schedulePointInTimeScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchedulePointInTimeScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schedulePointInTimeScheduleItemJSON) RawJSON() string {
	return r.raw
}

type ScheduledCharge struct {
	ID           string                 `json:"id,required" format:"uuid"`
	Product      ScheduledChargeProduct `json:"product,required"`
	Schedule     SchedulePointInTime    `json:"schedule,required"`
	CustomFields map[string]string      `json:"custom_fields"`
	// displayed on invoices
	Name string `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string              `json:"netsuite_sales_order_id"`
	JSON                 scheduledChargeJSON `json:"-"`
}

// scheduledChargeJSON contains the JSON metadata for the struct [ScheduledCharge]
type scheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	CustomFields         apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduledChargeJSON) RawJSON() string {
	return r.raw
}

type ScheduledChargeProduct struct {
	ID   string                     `json:"id,required" format:"uuid"`
	Name string                     `json:"name,required"`
	JSON scheduledChargeProductJSON `json:"-"`
}

// scheduledChargeProductJSON contains the JSON metadata for the struct
// [ScheduledChargeProduct]
type scheduledChargeProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduledChargeProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduledChargeProductJSON) RawJSON() string {
	return r.raw
}
