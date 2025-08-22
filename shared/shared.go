// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/tidwall/gjson"
)

type BaseThresholdCommit struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID   string `json:"product_id,required"`
	Description string `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string                  `json:"name"`
	JSON baseThresholdCommitJSON `json:"-"`
}

// baseThresholdCommitJSON contains the JSON metadata for the struct
// [BaseThresholdCommit]
type baseThresholdCommitJSON struct {
	ProductID   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BaseThresholdCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r baseThresholdCommitJSON) RawJSON() string {
	return r.raw
}

type BaseThresholdCommitParam struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID   param.Field[string] `json:"product_id,required"`
	Description param.Field[string] `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name param.Field[string] `json:"name"`
}

func (r BaseThresholdCommitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BaseUsageFilter struct {
	GroupKey    string              `json:"group_key,required"`
	GroupValues []string            `json:"group_values,required"`
	StartingAt  time.Time           `json:"starting_at" format:"date-time"`
	JSON        baseUsageFilterJSON `json:"-"`
}

// baseUsageFilterJSON contains the JSON metadata for the struct [BaseUsageFilter]
type baseUsageFilterJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BaseUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r baseUsageFilterJSON) RawJSON() string {
	return r.raw
}

type BaseUsageFilterParam struct {
	GroupKey    param.Field[string]    `json:"group_key,required"`
	GroupValues param.Field[[]string]  `json:"group_values,required"`
	StartingAt  param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r BaseUsageFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Commit struct {
	ID      string        `json:"id,required" format:"uuid"`
	Product CommitProduct `json:"product,required"`
	Type    CommitType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule ScheduleDuration `json:"access_schedule"`
	// (DEPRECATED) Use access_schedule + invoice_schedule instead.
	Amount                float64  `json:"amount"`
	ApplicableContractIDs []string `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string `json:"applicable_product_tags"`
	// RFC 3339 timestamp indicating when the commit was archived. If not provided, the
	// commit is not archived.
	ArchivedAt time.Time `json:"archived_at" format:"date-time"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64           `json:"balance"`
	Contract     CommitContract    `json:"contract"`
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
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
	RateType         CommitRateType       `json:"rate_type"`
	RolledOverFrom   CommitRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction float64              `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier `json:"specifiers"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey string     `json:"uniqueness_key"`
	JSON          commitJSON `json:"-"`
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
	ArchivedAt              apijson.Field
	Balance                 apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	HierarchyConfiguration  apijson.Field
	InvoiceContract         apijson.Field
	InvoiceSchedule         apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	RateType                apijson.Field
	RolledOverFrom          apijson.Field
	RolloverFraction        apijson.Field
	SalesforceOpportunityID apijson.Field
	Specifiers              apijson.Field
	UniquenessKey           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *Commit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitJSON) RawJSON() string {
	return r.raw
}

func (r Commit) ImplementsV1ContractListBalancesResponse() {}

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
	Amount        float64          `json:"amount,required"`
	Timestamp     time.Time        `json:"timestamp,required" format:"date-time"`
	Type          CommitLedgerType `json:"type,required"`
	ContractID    string           `json:"contract_id" format:"uuid"`
	InvoiceID     string           `json:"invoice_id" format:"uuid"`
	NewContractID string           `json:"new_contract_id" format:"uuid"`
	Reason        string           `json:"reason"`
	SegmentID     string           `json:"segment_id" format:"uuid"`
	JSON          commitLedgerJSON `json:"-"`
	union         CommitLedgerUnion
}

// commitLedgerJSON contains the JSON metadata for the struct [CommitLedger]
type commitLedgerJSON struct {
	Amount        apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	ContractID    apijson.Field
	InvoiceID     apijson.Field
	NewContractID apijson.Field
	Reason        apijson.Field
	SegmentID     apijson.Field
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
// [CommitLedgerPrepaidCommitSegmentStartLedgerEntry],
// [CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [CommitLedgerPrepaidCommitRolloverLedgerEntry],
// [CommitLedgerPrepaidCommitExpirationLedgerEntry],
// [CommitLedgerPrepaidCommitCanceledLedgerEntry],
// [CommitLedgerPrepaidCommitCreditedLedgerEntry],
// [CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry],
// [CommitLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [CommitLedgerPostpaidCommitRolloverLedgerEntry],
// [CommitLedgerPostpaidCommitTrueupLedgerEntry],
// [CommitLedgerPrepaidCommitManualLedgerEntry],
// [CommitLedgerPostpaidCommitManualLedgerEntry],
// [CommitLedgerPostpaidCommitExpirationLedgerEntry].
func (r CommitLedger) AsUnion() CommitLedgerUnion {
	return r.union
}

// Union satisfied by [CommitLedgerPrepaidCommitSegmentStartLedgerEntry],
// [CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [CommitLedgerPrepaidCommitRolloverLedgerEntry],
// [CommitLedgerPrepaidCommitExpirationLedgerEntry],
// [CommitLedgerPrepaidCommitCanceledLedgerEntry],
// [CommitLedgerPrepaidCommitCreditedLedgerEntry],
// [CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry],
// [CommitLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [CommitLedgerPostpaidCommitRolloverLedgerEntry],
// [CommitLedgerPostpaidCommitTrueupLedgerEntry],
// [CommitLedgerPrepaidCommitManualLedgerEntry],
// [CommitLedgerPostpaidCommitManualLedgerEntry] or
// [CommitLedgerPostpaidCommitExpirationLedgerEntry].
type CommitLedgerUnion interface {
	implementsCommitLedger()
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
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry{}),
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

func (r CommitLedgerPrepaidCommitSegmentStartLedgerEntry) implementsCommitLedger() {}

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
	Amount     float64                                                           `json:"amount,required"`
	InvoiceID  string                                                            `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                                         `json:"timestamp,required" format:"date-time"`
	Type       CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	ContractID string                                                            `json:"contract_id" format:"uuid"`
	JSON       commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
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
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsCommitLedger() {}

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

func (r CommitLedgerPrepaidCommitRolloverLedgerEntry) implementsCommitLedger() {}

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

func (r CommitLedgerPrepaidCommitExpirationLedgerEntry) implementsCommitLedger() {}

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
	Amount     float64                                          `json:"amount,required"`
	InvoiceID  string                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                           `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                        `json:"timestamp,required" format:"date-time"`
	Type       CommitLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	ContractID string                                           `json:"contract_id" format:"uuid"`
	JSON       commitLedgerPrepaidCommitCanceledLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitCanceledLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitCanceledLedgerEntry]
type commitLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitCanceledLedgerEntry) implementsCommitLedger() {}

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
	Amount     float64                                          `json:"amount,required"`
	InvoiceID  string                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                           `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                        `json:"timestamp,required" format:"date-time"`
	Type       CommitLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	ContractID string                                           `json:"contract_id" format:"uuid"`
	JSON       commitLedgerPrepaidCommitCreditedLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitCreditedLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitCreditedLedgerEntry]
type commitLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitCreditedLedgerEntry) implementsCommitLedger() {}

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

type CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry struct {
	Amount    float64                                                     `json:"amount,required"`
	SegmentID string                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                   `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryJSON contains the JSON
// metadata for the struct
// [CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry]
type commitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType string

const (
	CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryTypePrepaidCommitSeatBasedAdjustment CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType = "PREPAID_COMMIT_SEAT_BASED_ADJUSTMENT"
)

func (r CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryTypePrepaidCommitSeatBasedAdjustment:
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

func (r CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsCommitLedger() {}

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
	Amount     float64                                                            `json:"amount,required"`
	InvoiceID  string                                                             `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                                             `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                                          `json:"timestamp,required" format:"date-time"`
	Type       CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	ContractID string                                                             `json:"contract_id" format:"uuid"`
	JSON       commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
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
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsCommitLedger() {}

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

func (r CommitLedgerPostpaidCommitRolloverLedgerEntry) implementsCommitLedger() {}

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
	Amount     float64                                         `json:"amount,required"`
	InvoiceID  string                                          `json:"invoice_id,required" format:"uuid"`
	Timestamp  time.Time                                       `json:"timestamp,required" format:"date-time"`
	Type       CommitLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	ContractID string                                          `json:"contract_id" format:"uuid"`
	JSON       commitLedgerPostpaidCommitTrueupLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitTrueupLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPostpaidCommitTrueupLedgerEntry]
type commitLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitTrueupLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitTrueupLedgerEntry) implementsCommitLedger() {}

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

func (r CommitLedgerPrepaidCommitManualLedgerEntry) implementsCommitLedger() {}

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

func (r CommitLedgerPostpaidCommitManualLedgerEntry) implementsCommitLedger() {}

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

func (r CommitLedgerPostpaidCommitExpirationLedgerEntry) implementsCommitLedger() {}

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
	CommitLedgerTypePrepaidCommitSeatBasedAdjustment        CommitLedgerType = "PREPAID_COMMIT_SEAT_BASED_ADJUSTMENT"
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
	case CommitLedgerTypePrepaidCommitSegmentStart, CommitLedgerTypePrepaidCommitAutomatedInvoiceDeduction, CommitLedgerTypePrepaidCommitRollover, CommitLedgerTypePrepaidCommitExpiration, CommitLedgerTypePrepaidCommitCanceled, CommitLedgerTypePrepaidCommitCredited, CommitLedgerTypePrepaidCommitSeatBasedAdjustment, CommitLedgerTypePostpaidCommitInitialBalance, CommitLedgerTypePostpaidCommitAutomatedInvoiceDeduction, CommitLedgerTypePostpaidCommitRollover, CommitLedgerTypePostpaidCommitTrueup, CommitLedgerTypePrepaidCommitManual, CommitLedgerTypePostpaidCommitManual, CommitLedgerTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type CommitRateType string

const (
	CommitRateTypeCommitRate CommitRateType = "COMMIT_RATE"
	CommitRateTypeListRate   CommitRateType = "LIST_RATE"
)

func (r CommitRateType) IsKnown() bool {
	switch r {
	case CommitRateTypeCommitRate, CommitRateTypeListRate:
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

type CommitHierarchyConfiguration struct {
	ChildAccess CommitHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        commitHierarchyConfigurationJSON        `json:"-"`
}

// commitHierarchyConfigurationJSON contains the JSON metadata for the struct
// [CommitHierarchyConfiguration]
type commitHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type CommitHierarchyConfigurationChildAccess struct {
	Type CommitHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                 `json:"contract_ids"`
	JSON        commitHierarchyConfigurationChildAccessJSON `json:"-"`
	union       CommitHierarchyConfigurationChildAccessUnion
}

// commitHierarchyConfigurationChildAccessJSON contains the JSON metadata for the
// struct [CommitHierarchyConfigurationChildAccess]
type commitHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r commitHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *CommitHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = CommitHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CommitHierarchyConfigurationChildAccessUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
func (r CommitHierarchyConfigurationChildAccess) AsUnion() CommitHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone] or
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
type CommitHierarchyConfigurationChildAccessUnion interface {
	implementsCommitHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CommitHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs{}),
		},
	)
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType `json:"type,required"`
	JSON commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON `json:"-"`
}

// commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON
// contains the JSON metadata for the struct
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll]
type commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON) RawJSON() string {
	return r.raw
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsCommitHierarchyConfigurationChildAccess() {
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType `json:"type,required"`
	JSON commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON `json:"-"`
}

// commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON
// contains the JSON metadata for the struct
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
type commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON) RawJSON() string {
	return r.raw
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsCommitHierarchyConfigurationChildAccess() {
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs []string                                                                         `json:"contract_ids,required" format:"uuid"`
	Type        CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType `json:"type,required"`
	JSON        commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON `json:"-"`
}

// commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON
// contains the JSON metadata for the struct
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs]
type commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON) RawJSON() string {
	return r.raw
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsCommitHierarchyConfigurationChildAccess() {
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type CommitHierarchyConfigurationChildAccessType string

const (
	CommitHierarchyConfigurationChildAccessTypeAll         CommitHierarchyConfigurationChildAccessType = "ALL"
	CommitHierarchyConfigurationChildAccessTypeNone        CommitHierarchyConfigurationChildAccessType = "NONE"
	CommitHierarchyConfigurationChildAccessTypeContractIDs CommitHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r CommitHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case CommitHierarchyConfigurationChildAccessTypeAll, CommitHierarchyConfigurationChildAccessTypeNone, CommitHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
}

type CommitHierarchyConfigurationParam struct {
	ChildAccess param.Field[CommitHierarchyConfigurationChildAccessUnionParam] `json:"child_access,required"`
}

func (r CommitHierarchyConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommitHierarchyConfigurationChildAccessParam struct {
	Type        param.Field[CommitHierarchyConfigurationChildAccessType] `json:"type,required"`
	ContractIDs param.Field[interface{}]                                 `json:"contract_ids"`
}

func (r CommitHierarchyConfigurationChildAccessParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CommitHierarchyConfigurationChildAccessParam) implementsCommitHierarchyConfigurationChildAccessUnionParam() {
}

// Satisfied by
// [shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam],
// [shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam],
// [shared.CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam],
// [CommitHierarchyConfigurationChildAccessParam].
type CommitHierarchyConfigurationChildAccessUnionParam interface {
	implementsCommitHierarchyConfigurationChildAccessUnionParam()
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam struct {
	Type param.Field[CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType] `json:"type,required"`
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam) implementsCommitHierarchyConfigurationChildAccessUnionParam() {
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam struct {
	Type param.Field[CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType] `json:"type,required"`
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam) implementsCommitHierarchyConfigurationChildAccessUnionParam() {
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam struct {
	ContractIDs param.Field[[]string]                                                                         `json:"contract_ids,required" format:"uuid"`
	Type        param.Field[CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType] `json:"type,required"`
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam) implementsCommitHierarchyConfigurationChildAccessUnionParam() {
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type CommitRate struct {
	RateType CommitRateRateType `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price float64 `json:"price"`
	// Only set for TIERED rate_type.
	Tiers []Tier         `json:"tiers"`
	JSON  commitRateJSON `json:"-"`
}

// commitRateJSON contains the JSON metadata for the struct [CommitRate]
type commitRateJSON struct {
	RateType    apijson.Field
	Price       apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitRateJSON) RawJSON() string {
	return r.raw
}

type CommitRateRateType string

const (
	CommitRateRateTypeFlat         CommitRateRateType = "FLAT"
	CommitRateRateTypePercentage   CommitRateRateType = "PERCENTAGE"
	CommitRateRateTypeSubscription CommitRateRateType = "SUBSCRIPTION"
	CommitRateRateTypeTiered       CommitRateRateType = "TIERED"
	CommitRateRateTypeCustom       CommitRateRateType = "CUSTOM"
)

func (r CommitRateRateType) IsKnown() bool {
	switch r {
	case CommitRateRateTypeFlat, CommitRateRateTypePercentage, CommitRateRateTypeSubscription, CommitRateRateTypeTiered, CommitRateRateTypeCustom:
		return true
	}
	return false
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type CommitRateParam struct {
	RateType param.Field[CommitRateRateType] `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price param.Field[float64] `json:"price"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]TierParam] `json:"tiers"`
}

func (r CommitRateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommitSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string            `json:"product_tags"`
	JSON        commitSpecifierJSON `json:"-"`
}

// commitSpecifierJSON contains the JSON metadata for the struct [CommitSpecifier]
type commitSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CommitSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitSpecifierJSON) RawJSON() string {
	return r.raw
}

type CommitSpecifierInput struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                 `json:"product_tags"`
	JSON        commitSpecifierInputJSON `json:"-"`
}

// commitSpecifierInputJSON contains the JSON metadata for the struct
// [CommitSpecifierInput]
type commitSpecifierInputJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CommitSpecifierInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitSpecifierInputJSON) RawJSON() string {
	return r.raw
}

type CommitSpecifierInputParam struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r CommitSpecifierInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Contract struct {
	ID         string                    `json:"id,required" format:"uuid"`
	Amendments []ContractAmendment       `json:"amendments,required"`
	Current    ContractWithoutAmendments `json:"current,required"`
	CustomerID string                    `json:"customer_id,required" format:"uuid"`
	Initial    ContractWithoutAmendments `json:"initial,required"`
	// RFC 3339 timestamp indicating when the contract was archived. If not returned,
	// the contract is not archived.
	ArchivedAt   time.Time         `json:"archived_at" format:"date-time"`
	CustomFields map[string]string `json:"custom_fields"`
	// The billing provider configuration associated with a contract.
	CustomerBillingProviderConfiguration ContractCustomerBillingProviderConfiguration `json:"customer_billing_provider_configuration"`
	PrepaidBalanceThresholdConfiguration PrepaidBalanceThresholdConfiguration         `json:"prepaid_balance_threshold_configuration"`
	// Priority of the contract.
	Priority float64 `json:"priority"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices ContractScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	SpendThresholdConfiguration     SpendThresholdConfiguration             `json:"spend_threshold_configuration"`
	// List of subscriptions on the contract.
	Subscriptions []Subscription `json:"subscriptions"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string       `json:"uniqueness_key"`
	JSON          contractJSON `json:"-"`
}

// contractJSON contains the JSON metadata for the struct [Contract]
type contractJSON struct {
	ID                                   apijson.Field
	Amendments                           apijson.Field
	Current                              apijson.Field
	CustomerID                           apijson.Field
	Initial                              apijson.Field
	ArchivedAt                           apijson.Field
	CustomFields                         apijson.Field
	CustomerBillingProviderConfiguration apijson.Field
	PrepaidBalanceThresholdConfiguration apijson.Field
	Priority                             apijson.Field
	ScheduledChargesOnUsageInvoices      apijson.Field
	SpendThresholdConfiguration          apijson.Field
	Subscriptions                        apijson.Field
	UniquenessKey                        apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *Contract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractJSON) RawJSON() string {
	return r.raw
}

type ContractAmendment struct {
	ID               string            `json:"id,required" format:"uuid"`
	Commits          []Commit          `json:"commits,required"`
	CreatedAt        time.Time         `json:"created_at,required" format:"date-time"`
	CreatedBy        string            `json:"created_by,required"`
	Overrides        []Override        `json:"overrides,required"`
	ScheduledCharges []ScheduledCharge `json:"scheduled_charges,required"`
	StartingAt       time.Time         `json:"starting_at,required" format:"date-time"`
	Credits          []Credit          `json:"credits"`
	// This field's availability is dependent on your client's configuration.
	Discounts []Discount `json:"discounts"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []ProService `json:"professional_services"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []ContractAmendmentsResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                `json:"salesforce_opportunity_id"`
	JSON                    contractAmendmentJSON `json:"-"`
}

// contractAmendmentJSON contains the JSON metadata for the struct
// [ContractAmendment]
type contractAmendmentJSON struct {
	ID                      apijson.Field
	Commits                 apijson.Field
	CreatedAt               apijson.Field
	CreatedBy               apijson.Field
	Overrides               apijson.Field
	ScheduledCharges        apijson.Field
	StartingAt              apijson.Field
	Credits                 apijson.Field
	Discounts               apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	ProfessionalServices    apijson.Field
	ResellerRoyalties       apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractAmendment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractAmendmentJSON) RawJSON() string {
	return r.raw
}

type ContractAmendmentsResellerRoyalty struct {
	ResellerType          ContractAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
	AwsAccountNumber      string                                          `json:"aws_account_number"`
	AwsOfferID            string                                          `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                          `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                       `json:"ending_before,nullable" format:"date-time"`
	Fraction              float64                                         `json:"fraction"`
	GcpAccountID          string                                          `json:"gcp_account_id"`
	GcpOfferID            string                                          `json:"gcp_offer_id"`
	NetsuiteResellerID    string                                          `json:"netsuite_reseller_id"`
	ResellerContractValue float64                                         `json:"reseller_contract_value"`
	StartingAt            time.Time                                       `json:"starting_at" format:"date-time"`
	JSON                  contractAmendmentsResellerRoyaltyJSON           `json:"-"`
}

// contractAmendmentsResellerRoyaltyJSON contains the JSON metadata for the struct
// [ContractAmendmentsResellerRoyalty]
type contractAmendmentsResellerRoyaltyJSON struct {
	ResellerType          apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	Fraction              apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerContractValue apijson.Field
	StartingAt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractAmendmentsResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type ContractAmendmentsResellerRoyaltiesResellerType string

const (
	ContractAmendmentsResellerRoyaltiesResellerTypeAws           ContractAmendmentsResellerRoyaltiesResellerType = "AWS"
	ContractAmendmentsResellerRoyaltiesResellerTypeAwsProService ContractAmendmentsResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	ContractAmendmentsResellerRoyaltiesResellerTypeGcp           ContractAmendmentsResellerRoyaltiesResellerType = "GCP"
	ContractAmendmentsResellerRoyaltiesResellerTypeGcpProService ContractAmendmentsResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r ContractAmendmentsResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case ContractAmendmentsResellerRoyaltiesResellerTypeAws, ContractAmendmentsResellerRoyaltiesResellerTypeAwsProService, ContractAmendmentsResellerRoyaltiesResellerTypeGcp, ContractAmendmentsResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

// The billing provider configuration associated with a contract.
type ContractCustomerBillingProviderConfiguration struct {
	BillingProvider ContractCustomerBillingProviderConfigurationBillingProvider `json:"billing_provider,required"`
	DeliveryMethod  ContractCustomerBillingProviderConfigurationDeliveryMethod  `json:"delivery_method,required"`
	ID              string                                                      `json:"id" format:"uuid"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider.
	Configuration map[string]interface{}                           `json:"configuration"`
	JSON          contractCustomerBillingProviderConfigurationJSON `json:"-"`
}

// contractCustomerBillingProviderConfigurationJSON contains the JSON metadata for
// the struct [ContractCustomerBillingProviderConfiguration]
type contractCustomerBillingProviderConfigurationJSON struct {
	BillingProvider apijson.Field
	DeliveryMethod  apijson.Field
	ID              apijson.Field
	Configuration   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ContractCustomerBillingProviderConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractCustomerBillingProviderConfigurationJSON) RawJSON() string {
	return r.raw
}

type ContractCustomerBillingProviderConfigurationBillingProvider string

const (
	ContractCustomerBillingProviderConfigurationBillingProviderAwsMarketplace   ContractCustomerBillingProviderConfigurationBillingProvider = "aws_marketplace"
	ContractCustomerBillingProviderConfigurationBillingProviderStripe           ContractCustomerBillingProviderConfigurationBillingProvider = "stripe"
	ContractCustomerBillingProviderConfigurationBillingProviderNetsuite         ContractCustomerBillingProviderConfigurationBillingProvider = "netsuite"
	ContractCustomerBillingProviderConfigurationBillingProviderCustom           ContractCustomerBillingProviderConfigurationBillingProvider = "custom"
	ContractCustomerBillingProviderConfigurationBillingProviderAzureMarketplace ContractCustomerBillingProviderConfigurationBillingProvider = "azure_marketplace"
	ContractCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline ContractCustomerBillingProviderConfigurationBillingProvider = "quickbooks_online"
	ContractCustomerBillingProviderConfigurationBillingProviderWorkday          ContractCustomerBillingProviderConfigurationBillingProvider = "workday"
	ContractCustomerBillingProviderConfigurationBillingProviderGcpMarketplace   ContractCustomerBillingProviderConfigurationBillingProvider = "gcp_marketplace"
)

func (r ContractCustomerBillingProviderConfigurationBillingProvider) IsKnown() bool {
	switch r {
	case ContractCustomerBillingProviderConfigurationBillingProviderAwsMarketplace, ContractCustomerBillingProviderConfigurationBillingProviderStripe, ContractCustomerBillingProviderConfigurationBillingProviderNetsuite, ContractCustomerBillingProviderConfigurationBillingProviderCustom, ContractCustomerBillingProviderConfigurationBillingProviderAzureMarketplace, ContractCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline, ContractCustomerBillingProviderConfigurationBillingProviderWorkday, ContractCustomerBillingProviderConfigurationBillingProviderGcpMarketplace:
		return true
	}
	return false
}

type ContractCustomerBillingProviderConfigurationDeliveryMethod string

const (
	ContractCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider ContractCustomerBillingProviderConfigurationDeliveryMethod = "direct_to_billing_provider"
	ContractCustomerBillingProviderConfigurationDeliveryMethodAwsSqs                  ContractCustomerBillingProviderConfigurationDeliveryMethod = "aws_sqs"
	ContractCustomerBillingProviderConfigurationDeliveryMethodTackle                  ContractCustomerBillingProviderConfigurationDeliveryMethod = "tackle"
	ContractCustomerBillingProviderConfigurationDeliveryMethodAwsSns                  ContractCustomerBillingProviderConfigurationDeliveryMethod = "aws_sns"
)

func (r ContractCustomerBillingProviderConfigurationDeliveryMethod) IsKnown() bool {
	switch r {
	case ContractCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider, ContractCustomerBillingProviderConfigurationDeliveryMethodAwsSqs, ContractCustomerBillingProviderConfigurationDeliveryMethodTackle, ContractCustomerBillingProviderConfigurationDeliveryMethodAwsSns:
		return true
	}
	return false
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type ContractScheduledChargesOnUsageInvoices string

const (
	ContractScheduledChargesOnUsageInvoicesAll ContractScheduledChargesOnUsageInvoices = "ALL"
)

func (r ContractScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case ContractScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
}

type ContractV2 struct {
	ID                     string                           `json:"id,required" format:"uuid"`
	Commits                []ContractV2Commit               `json:"commits,required"`
	CreatedAt              time.Time                        `json:"created_at,required" format:"date-time"`
	CreatedBy              string                           `json:"created_by,required"`
	CustomerID             string                           `json:"customer_id,required" format:"uuid"`
	Overrides              []ContractV2Override             `json:"overrides,required"`
	ScheduledCharges       []ScheduledCharge                `json:"scheduled_charges,required"`
	StartingAt             time.Time                        `json:"starting_at,required" format:"date-time"`
	Transitions            []ContractV2Transition           `json:"transitions,required"`
	UsageFilter            []ContractV2UsageFilter          `json:"usage_filter,required"`
	UsageStatementSchedule ContractV2UsageStatementSchedule `json:"usage_statement_schedule,required"`
	ArchivedAt             time.Time                        `json:"archived_at" format:"date-time"`
	Credits                []ContractV2Credit               `json:"credits"`
	CustomFields           map[string]string                `json:"custom_fields"`
	// This field's availability is dependent on your client's configuration.
	CustomerBillingProviderConfiguration ContractV2CustomerBillingProviderConfiguration `json:"customer_billing_provider_configuration"`
	// This field's availability is dependent on your client's configuration.
	Discounts    []Discount `json:"discounts"`
	EndingBefore time.Time  `json:"ending_before" format:"date-time"`
	// Indicates whether there are more items than the limit for this endpoint. Use the
	// respective list endpoints to get the full lists.
	HasMore ContractV2HasMore `json:"has_more"`
	// Either a **parent** configuration with a list of children or a **child**
	// configuration with a single parent.
	HierarchyConfiguration HierarchyConfiguration `json:"hierarchy_configuration"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
	// prices automatically. EXPLICIT prioritization requires specifying priorities for
	// each multiplier; the one with the lowest priority value will be prioritized
	// first.
	MultiplierOverridePrioritization ContractV2MultiplierOverridePrioritization `json:"multiplier_override_prioritization"`
	Name                             string                                     `json:"name"`
	NetPaymentTermsDays              float64                                    `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID                 string                                 `json:"netsuite_sales_order_id"`
	PrepaidBalanceThresholdConfiguration PrepaidBalanceThresholdConfigurationV2 `json:"prepaid_balance_threshold_configuration"`
	// Priority of the contract.
	Priority float64 `json:"priority"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []ProService                `json:"professional_services"`
	RateCardID           string                      `json:"rate_card_id" format:"uuid"`
	RecurringCommits     []ContractV2RecurringCommit `json:"recurring_commits"`
	RecurringCredits     []ContractV2RecurringCredit `json:"recurring_credits"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []ContractV2ResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices ContractV2ScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	SpendThresholdConfiguration     SpendThresholdConfigurationV2             `json:"spend_threshold_configuration"`
	// List of subscriptions on the contract.
	Subscriptions      []Subscription `json:"subscriptions"`
	TotalContractValue float64        `json:"total_contract_value"`
	// Optional uniqueness key to prevent duplicate contract creations.
	UniquenessKey string         `json:"uniqueness_key"`
	JSON          contractV2JSON `json:"-"`
}

// contractV2JSON contains the JSON metadata for the struct [ContractV2]
type contractV2JSON struct {
	ID                                   apijson.Field
	Commits                              apijson.Field
	CreatedAt                            apijson.Field
	CreatedBy                            apijson.Field
	CustomerID                           apijson.Field
	Overrides                            apijson.Field
	ScheduledCharges                     apijson.Field
	StartingAt                           apijson.Field
	Transitions                          apijson.Field
	UsageFilter                          apijson.Field
	UsageStatementSchedule               apijson.Field
	ArchivedAt                           apijson.Field
	Credits                              apijson.Field
	CustomFields                         apijson.Field
	CustomerBillingProviderConfiguration apijson.Field
	Discounts                            apijson.Field
	EndingBefore                         apijson.Field
	HasMore                              apijson.Field
	HierarchyConfiguration               apijson.Field
	MultiplierOverridePrioritization     apijson.Field
	Name                                 apijson.Field
	NetPaymentTermsDays                  apijson.Field
	NetsuiteSalesOrderID                 apijson.Field
	PrepaidBalanceThresholdConfiguration apijson.Field
	Priority                             apijson.Field
	ProfessionalServices                 apijson.Field
	RateCardID                           apijson.Field
	RecurringCommits                     apijson.Field
	RecurringCredits                     apijson.Field
	ResellerRoyalties                    apijson.Field
	SalesforceOpportunityID              apijson.Field
	ScheduledChargesOnUsageInvoices      apijson.Field
	SpendThresholdConfiguration          apijson.Field
	Subscriptions                        apijson.Field
	TotalContractValue                   apijson.Field
	UniquenessKey                        apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *ContractV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2JSON) RawJSON() string {
	return r.raw
}

type ContractV2Commit struct {
	ID      string                   `json:"id,required" format:"uuid"`
	Product ContractV2CommitsProduct `json:"product,required"`
	Type    ContractV2CommitsType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule        ScheduleDuration `json:"access_schedule"`
	ApplicableContractIDs []string         `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string         `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string         `json:"applicable_product_tags"`
	ArchivedAt            time.Time        `json:"archived_at" format:"date-time"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64                   `json:"balance"`
	Contract     ContractV2CommitsContract `json:"contract"`
	CustomFields map[string]string         `json:"custom_fields"`
	Description  string                    `json:"description"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// The contract that this commit will be billed on.
	InvoiceContract ContractV2CommitsInvoiceContract `json:"invoice_contract"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule SchedulePointInTime `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger []ContractV2CommitsLedger `json:"ledger"`
	Name   string                    `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority         float64                         `json:"priority"`
	RateType         ContractV2CommitsRateType       `json:"rate_type"`
	RolledOverFrom   ContractV2CommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction float64                         `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier    `json:"specifiers"`
	JSON       contractV2CommitJSON `json:"-"`
}

// contractV2CommitJSON contains the JSON metadata for the struct
// [ContractV2Commit]
type contractV2CommitJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	ArchivedAt              apijson.Field
	Balance                 apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	HierarchyConfiguration  apijson.Field
	InvoiceContract         apijson.Field
	InvoiceSchedule         apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	RateType                apijson.Field
	RolledOverFrom          apijson.Field
	RolloverFraction        apijson.Field
	SalesforceOpportunityID apijson.Field
	Specifiers              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractV2Commit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitJSON) RawJSON() string {
	return r.raw
}

type ContractV2CommitsProduct struct {
	ID   string                       `json:"id,required" format:"uuid"`
	Name string                       `json:"name,required"`
	JSON contractV2CommitsProductJSON `json:"-"`
}

// contractV2CommitsProductJSON contains the JSON metadata for the struct
// [ContractV2CommitsProduct]
type contractV2CommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsProductJSON) RawJSON() string {
	return r.raw
}

type ContractV2CommitsType string

const (
	ContractV2CommitsTypePrepaid  ContractV2CommitsType = "PREPAID"
	ContractV2CommitsTypePostpaid ContractV2CommitsType = "POSTPAID"
)

func (r ContractV2CommitsType) IsKnown() bool {
	switch r {
	case ContractV2CommitsTypePrepaid, ContractV2CommitsTypePostpaid:
		return true
	}
	return false
}

type ContractV2CommitsContract struct {
	ID   string                        `json:"id,required" format:"uuid"`
	JSON contractV2CommitsContractJSON `json:"-"`
}

// contractV2CommitsContractJSON contains the JSON metadata for the struct
// [ContractV2CommitsContract]
type contractV2CommitsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsContractJSON) RawJSON() string {
	return r.raw
}

// The contract that this commit will be billed on.
type ContractV2CommitsInvoiceContract struct {
	ID   string                               `json:"id,required" format:"uuid"`
	JSON contractV2CommitsInvoiceContractJSON `json:"-"`
}

// contractV2CommitsInvoiceContractJSON contains the JSON metadata for the struct
// [ContractV2CommitsInvoiceContract]
type contractV2CommitsInvoiceContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsInvoiceContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsInvoiceContractJSON) RawJSON() string {
	return r.raw
}

type ContractV2CommitsLedger struct {
	Amount        float64                     `json:"amount,required"`
	Timestamp     time.Time                   `json:"timestamp,required" format:"date-time"`
	Type          ContractV2CommitsLedgerType `json:"type,required"`
	ContractID    string                      `json:"contract_id" format:"uuid"`
	InvoiceID     string                      `json:"invoice_id" format:"uuid"`
	NewContractID string                      `json:"new_contract_id" format:"uuid"`
	Reason        string                      `json:"reason"`
	SegmentID     string                      `json:"segment_id" format:"uuid"`
	JSON          contractV2CommitsLedgerJSON `json:"-"`
	union         ContractV2CommitsLedgerUnion
}

// contractV2CommitsLedgerJSON contains the JSON metadata for the struct
// [ContractV2CommitsLedger]
type contractV2CommitsLedgerJSON struct {
	Amount        apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	ContractID    apijson.Field
	InvoiceID     apijson.Field
	NewContractID apijson.Field
	Reason        apijson.Field
	SegmentID     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r contractV2CommitsLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *ContractV2CommitsLedger) UnmarshalJSON(data []byte) (err error) {
	*r = ContractV2CommitsLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContractV2CommitsLedgerUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry],
// [ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntry],
// [ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitManualLedgerEntry],
// [ContractV2CommitsLedgerPostpaidCommitManualLedgerEntry],
// [ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntry].
func (r ContractV2CommitsLedger) AsUnion() ContractV2CommitsLedgerUnion {
	return r.union
}

// Union satisfied by
// [ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry],
// [ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntry],
// [ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntry],
// [ContractV2CommitsLedgerPrepaidCommitManualLedgerEntry],
// [ContractV2CommitsLedgerPostpaidCommitManualLedgerEntry] or
// [ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntry].
type ContractV2CommitsLedgerUnion interface {
	implementsContractV2CommitsLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContractV2CommitsLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPrepaidCommitManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPostpaidCommitManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntry{}),
		},
	)
}

type ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                         `json:"amount,required"`
	SegmentID string                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON contains the
// JSON metadata for the struct
// [ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractV2CommitsLedger() {
}

type ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

func (r ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount     float64                                                                      `json:"amount,required"`
	InvoiceID  string                                                                       `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                                                       `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                                                    `json:"timestamp,required" format:"date-time"`
	Type       ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	ContractID string                                                                       `json:"contract_id" format:"uuid"`
	JSON       contractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractV2CommitsLedger() {
}

type ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                     `json:"amount,required"`
	NewContractID string                                                      `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                   `json:"timestamp,required" format:"date-time"`
	Type          ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryJSON contains the JSON
// metadata for the struct
// [ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractV2CommitsLedger() {
}

type ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

func (r ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                       `json:"amount,required"`
	SegmentID string                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                     `json:"timestamp,required" format:"date-time"`
	Type      ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryJSON contains the JSON
// metadata for the struct
// [ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractV2CommitsLedger() {
}

type ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

func (r ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount     float64                                                     `json:"amount,required"`
	InvoiceID  string                                                      `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                                   `json:"timestamp,required" format:"date-time"`
	Type       ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	ContractID string                                                      `json:"contract_id" format:"uuid"`
	JSON       contractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryJSON contains the JSON
// metadata for the struct
// [ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractV2CommitsLedger() {
}

type ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

func (r ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount     float64                                                     `json:"amount,required"`
	InvoiceID  string                                                      `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                                   `json:"timestamp,required" format:"date-time"`
	Type       ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	ContractID string                                                      `json:"contract_id" format:"uuid"`
	JSON       contractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryJSON contains the JSON
// metadata for the struct
// [ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractV2CommitsLedger() {
}

type ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

func (r ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry struct {
	Amount    float64                                                                `json:"amount,required"`
	SegmentID string                                                                 `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                              `json:"timestamp,required" format:"date-time"`
	Type      ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType `json:"type,required"`
	JSON      contractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryJSON contains
// the JSON metadata for the struct
// [ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry]
type contractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry) implementsContractV2CommitsLedger() {
}

type ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType string

const (
	ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryTypePrepaidCommitSeatBasedAdjustment ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType = "PREPAID_COMMIT_SEAT_BASED_ADJUSTMENT"
)

func (r ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntryTypePrepaidCommitSeatBasedAdjustment:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                            `json:"amount,required"`
	Timestamp time.Time                                                          `json:"timestamp,required" format:"date-time"`
	Type      ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON contains the
// JSON metadata for the struct
// [ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractV2CommitsLedger() {
}

type ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

func (r ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount     float64                                                                       `json:"amount,required"`
	InvoiceID  string                                                                        `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                                                     `json:"timestamp,required" format:"date-time"`
	Type       ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	ContractID string                                                                        `json:"contract_id" format:"uuid"`
	JSON       contractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractV2CommitsLedger() {
}

type ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                      `json:"amount,required"`
	NewContractID string                                                       `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                       `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                    `json:"timestamp,required" format:"date-time"`
	Type          ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryJSON contains the JSON
// metadata for the struct
// [ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntry]
type contractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntry) implementsContractV2CommitsLedger() {
}

type ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryType string

const (
	ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryType = "POSTPAID_COMMIT_ROLLOVER"
)

func (r ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount     float64                                                    `json:"amount,required"`
	InvoiceID  string                                                     `json:"invoice_id,required" format:"uuid"`
	Timestamp  time.Time                                                  `json:"timestamp,required" format:"date-time"`
	Type       ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	ContractID string                                                     `json:"contract_id" format:"uuid"`
	JSON       contractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryJSON contains the JSON
// metadata for the struct [ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractV2CommitsLedger() {}

type ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

func (r ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPrepaidCommitManualLedgerEntry struct {
	Amount    float64                                                   `json:"amount,required"`
	Reason    string                                                    `json:"reason,required"`
	Timestamp time.Time                                                 `json:"timestamp,required" format:"date-time"`
	Type      ContractV2CommitsLedgerPrepaidCommitManualLedgerEntryType `json:"type,required"`
	JSON      contractV2CommitsLedgerPrepaidCommitManualLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPrepaidCommitManualLedgerEntryJSON contains the JSON
// metadata for the struct [ContractV2CommitsLedgerPrepaidCommitManualLedgerEntry]
type contractV2CommitsLedgerPrepaidCommitManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPrepaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPrepaidCommitManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPrepaidCommitManualLedgerEntry) implementsContractV2CommitsLedger() {}

type ContractV2CommitsLedgerPrepaidCommitManualLedgerEntryType string

const (
	ContractV2CommitsLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual ContractV2CommitsLedgerPrepaidCommitManualLedgerEntryType = "PREPAID_COMMIT_MANUAL"
)

func (r ContractV2CommitsLedgerPrepaidCommitManualLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPostpaidCommitManualLedgerEntry struct {
	Amount    float64                                                    `json:"amount,required"`
	Reason    string                                                     `json:"reason,required"`
	Timestamp time.Time                                                  `json:"timestamp,required" format:"date-time"`
	Type      ContractV2CommitsLedgerPostpaidCommitManualLedgerEntryType `json:"type,required"`
	JSON      contractV2CommitsLedgerPostpaidCommitManualLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPostpaidCommitManualLedgerEntryJSON contains the JSON
// metadata for the struct [ContractV2CommitsLedgerPostpaidCommitManualLedgerEntry]
type contractV2CommitsLedgerPostpaidCommitManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPostpaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPostpaidCommitManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPostpaidCommitManualLedgerEntry) implementsContractV2CommitsLedger() {}

type ContractV2CommitsLedgerPostpaidCommitManualLedgerEntryType string

const (
	ContractV2CommitsLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual ContractV2CommitsLedgerPostpaidCommitManualLedgerEntryType = "POSTPAID_COMMIT_MANUAL"
)

func (r ContractV2CommitsLedgerPostpaidCommitManualLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual:
		return true
	}
	return false
}

type ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                        `json:"amount,required"`
	Timestamp time.Time                                                      `json:"timestamp,required" format:"date-time"`
	Type      ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryJSON `json:"-"`
}

// contractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryJSON contains the JSON
// metadata for the struct
// [ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntry]
type contractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntry) implementsContractV2CommitsLedger() {
}

type ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryType string

const (
	ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type ContractV2CommitsLedgerType string

const (
	ContractV2CommitsLedgerTypePrepaidCommitSegmentStart               ContractV2CommitsLedgerType = "PREPAID_COMMIT_SEGMENT_START"
	ContractV2CommitsLedgerTypePrepaidCommitAutomatedInvoiceDeduction  ContractV2CommitsLedgerType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	ContractV2CommitsLedgerTypePrepaidCommitRollover                   ContractV2CommitsLedgerType = "PREPAID_COMMIT_ROLLOVER"
	ContractV2CommitsLedgerTypePrepaidCommitExpiration                 ContractV2CommitsLedgerType = "PREPAID_COMMIT_EXPIRATION"
	ContractV2CommitsLedgerTypePrepaidCommitCanceled                   ContractV2CommitsLedgerType = "PREPAID_COMMIT_CANCELED"
	ContractV2CommitsLedgerTypePrepaidCommitCredited                   ContractV2CommitsLedgerType = "PREPAID_COMMIT_CREDITED"
	ContractV2CommitsLedgerTypePrepaidCommitSeatBasedAdjustment        ContractV2CommitsLedgerType = "PREPAID_COMMIT_SEAT_BASED_ADJUSTMENT"
	ContractV2CommitsLedgerTypePostpaidCommitInitialBalance            ContractV2CommitsLedgerType = "POSTPAID_COMMIT_INITIAL_BALANCE"
	ContractV2CommitsLedgerTypePostpaidCommitAutomatedInvoiceDeduction ContractV2CommitsLedgerType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	ContractV2CommitsLedgerTypePostpaidCommitRollover                  ContractV2CommitsLedgerType = "POSTPAID_COMMIT_ROLLOVER"
	ContractV2CommitsLedgerTypePostpaidCommitTrueup                    ContractV2CommitsLedgerType = "POSTPAID_COMMIT_TRUEUP"
	ContractV2CommitsLedgerTypePrepaidCommitManual                     ContractV2CommitsLedgerType = "PREPAID_COMMIT_MANUAL"
	ContractV2CommitsLedgerTypePostpaidCommitManual                    ContractV2CommitsLedgerType = "POSTPAID_COMMIT_MANUAL"
	ContractV2CommitsLedgerTypePostpaidCommitExpiration                ContractV2CommitsLedgerType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r ContractV2CommitsLedgerType) IsKnown() bool {
	switch r {
	case ContractV2CommitsLedgerTypePrepaidCommitSegmentStart, ContractV2CommitsLedgerTypePrepaidCommitAutomatedInvoiceDeduction, ContractV2CommitsLedgerTypePrepaidCommitRollover, ContractV2CommitsLedgerTypePrepaidCommitExpiration, ContractV2CommitsLedgerTypePrepaidCommitCanceled, ContractV2CommitsLedgerTypePrepaidCommitCredited, ContractV2CommitsLedgerTypePrepaidCommitSeatBasedAdjustment, ContractV2CommitsLedgerTypePostpaidCommitInitialBalance, ContractV2CommitsLedgerTypePostpaidCommitAutomatedInvoiceDeduction, ContractV2CommitsLedgerTypePostpaidCommitRollover, ContractV2CommitsLedgerTypePostpaidCommitTrueup, ContractV2CommitsLedgerTypePrepaidCommitManual, ContractV2CommitsLedgerTypePostpaidCommitManual, ContractV2CommitsLedgerTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type ContractV2CommitsRateType string

const (
	ContractV2CommitsRateTypeCommitRate ContractV2CommitsRateType = "COMMIT_RATE"
	ContractV2CommitsRateTypeListRate   ContractV2CommitsRateType = "LIST_RATE"
)

func (r ContractV2CommitsRateType) IsKnown() bool {
	switch r {
	case ContractV2CommitsRateTypeCommitRate, ContractV2CommitsRateTypeListRate:
		return true
	}
	return false
}

type ContractV2CommitsRolledOverFrom struct {
	CommitID   string                              `json:"commit_id,required" format:"uuid"`
	ContractID string                              `json:"contract_id,required" format:"uuid"`
	JSON       contractV2CommitsRolledOverFromJSON `json:"-"`
}

// contractV2CommitsRolledOverFromJSON contains the JSON metadata for the struct
// [ContractV2CommitsRolledOverFrom]
type contractV2CommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CommitsRolledOverFromJSON) RawJSON() string {
	return r.raw
}

type ContractV2Override struct {
	ID                    string                                 `json:"id,required" format:"uuid"`
	StartingAt            time.Time                              `json:"starting_at,required" format:"date-time"`
	ApplicableProductTags []string                               `json:"applicable_product_tags"`
	EndingBefore          time.Time                              `json:"ending_before" format:"date-time"`
	Entitled              bool                                   `json:"entitled"`
	IsCommitSpecific      bool                                   `json:"is_commit_specific"`
	Multiplier            float64                                `json:"multiplier"`
	OverrideSpecifiers    []ContractV2OverridesOverrideSpecifier `json:"override_specifiers"`
	OverrideTiers         []OverrideTier                         `json:"override_tiers"`
	OverwriteRate         ContractV2OverridesOverwriteRate       `json:"overwrite_rate"`
	Priority              float64                                `json:"priority"`
	Product               ContractV2OverridesProduct             `json:"product"`
	Target                ContractV2OverridesTarget              `json:"target"`
	Type                  ContractV2OverridesType                `json:"type"`
	JSON                  contractV2OverrideJSON                 `json:"-"`
}

// contractV2OverrideJSON contains the JSON metadata for the struct
// [ContractV2Override]
type contractV2OverrideJSON struct {
	ID                    apijson.Field
	StartingAt            apijson.Field
	ApplicableProductTags apijson.Field
	EndingBefore          apijson.Field
	Entitled              apijson.Field
	IsCommitSpecific      apijson.Field
	Multiplier            apijson.Field
	OverrideSpecifiers    apijson.Field
	OverrideTiers         apijson.Field
	OverwriteRate         apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	Target                apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractV2Override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2OverrideJSON) RawJSON() string {
	return r.raw
}

type ContractV2OverridesOverrideSpecifier struct {
	BillingFrequency        ContractV2OverridesOverrideSpecifiersBillingFrequency `json:"billing_frequency"`
	CommitIDs               []string                                              `json:"commit_ids"`
	PresentationGroupValues map[string]string                                     `json:"presentation_group_values"`
	PricingGroupValues      map[string]string                                     `json:"pricing_group_values"`
	ProductID               string                                                `json:"product_id" format:"uuid"`
	ProductTags             []string                                              `json:"product_tags"`
	RecurringCommitIDs      []string                                              `json:"recurring_commit_ids"`
	RecurringCreditIDs      []string                                              `json:"recurring_credit_ids"`
	JSON                    contractV2OverridesOverrideSpecifierJSON              `json:"-"`
}

// contractV2OverridesOverrideSpecifierJSON contains the JSON metadata for the
// struct [ContractV2OverridesOverrideSpecifier]
type contractV2OverridesOverrideSpecifierJSON struct {
	BillingFrequency        apijson.Field
	CommitIDs               apijson.Field
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	RecurringCommitIDs      apijson.Field
	RecurringCreditIDs      apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractV2OverridesOverrideSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2OverridesOverrideSpecifierJSON) RawJSON() string {
	return r.raw
}

type ContractV2OverridesOverrideSpecifiersBillingFrequency string

const (
	ContractV2OverridesOverrideSpecifiersBillingFrequencyMonthly   ContractV2OverridesOverrideSpecifiersBillingFrequency = "MONTHLY"
	ContractV2OverridesOverrideSpecifiersBillingFrequencyQuarterly ContractV2OverridesOverrideSpecifiersBillingFrequency = "QUARTERLY"
	ContractV2OverridesOverrideSpecifiersBillingFrequencyAnnual    ContractV2OverridesOverrideSpecifiersBillingFrequency = "ANNUAL"
	ContractV2OverridesOverrideSpecifiersBillingFrequencyWeekly    ContractV2OverridesOverrideSpecifiersBillingFrequency = "WEEKLY"
)

func (r ContractV2OverridesOverrideSpecifiersBillingFrequency) IsKnown() bool {
	switch r {
	case ContractV2OverridesOverrideSpecifiersBillingFrequencyMonthly, ContractV2OverridesOverrideSpecifiersBillingFrequencyQuarterly, ContractV2OverridesOverrideSpecifiersBillingFrequencyAnnual, ContractV2OverridesOverrideSpecifiersBillingFrequencyWeekly:
		return true
	}
	return false
}

type ContractV2OverridesOverwriteRate struct {
	RateType   ContractV2OverridesOverwriteRateRateType `json:"rate_type,required"`
	CreditType CreditTypeData                           `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]interface{} `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []Tier                               `json:"tiers"`
	JSON  contractV2OverridesOverwriteRateJSON `json:"-"`
}

// contractV2OverridesOverwriteRateJSON contains the JSON metadata for the struct
// [ContractV2OverridesOverwriteRate]
type contractV2OverridesOverwriteRateJSON struct {
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

func (r *ContractV2OverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2OverridesOverwriteRateJSON) RawJSON() string {
	return r.raw
}

type ContractV2OverridesOverwriteRateRateType string

const (
	ContractV2OverridesOverwriteRateRateTypeFlat         ContractV2OverridesOverwriteRateRateType = "FLAT"
	ContractV2OverridesOverwriteRateRateTypePercentage   ContractV2OverridesOverwriteRateRateType = "PERCENTAGE"
	ContractV2OverridesOverwriteRateRateTypeSubscription ContractV2OverridesOverwriteRateRateType = "SUBSCRIPTION"
	ContractV2OverridesOverwriteRateRateTypeTiered       ContractV2OverridesOverwriteRateRateType = "TIERED"
	ContractV2OverridesOverwriteRateRateTypeCustom       ContractV2OverridesOverwriteRateRateType = "CUSTOM"
)

func (r ContractV2OverridesOverwriteRateRateType) IsKnown() bool {
	switch r {
	case ContractV2OverridesOverwriteRateRateTypeFlat, ContractV2OverridesOverwriteRateRateTypePercentage, ContractV2OverridesOverwriteRateRateTypeSubscription, ContractV2OverridesOverwriteRateRateTypeTiered, ContractV2OverridesOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

type ContractV2OverridesProduct struct {
	ID   string                         `json:"id,required" format:"uuid"`
	Name string                         `json:"name,required"`
	JSON contractV2OverridesProductJSON `json:"-"`
}

// contractV2OverridesProductJSON contains the JSON metadata for the struct
// [ContractV2OverridesProduct]
type contractV2OverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2OverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2OverridesProductJSON) RawJSON() string {
	return r.raw
}

type ContractV2OverridesTarget string

const (
	ContractV2OverridesTargetCommitRate ContractV2OverridesTarget = "COMMIT_RATE"
	ContractV2OverridesTargetListRate   ContractV2OverridesTarget = "LIST_RATE"
)

func (r ContractV2OverridesTarget) IsKnown() bool {
	switch r {
	case ContractV2OverridesTargetCommitRate, ContractV2OverridesTargetListRate:
		return true
	}
	return false
}

type ContractV2OverridesType string

const (
	ContractV2OverridesTypeOverwrite  ContractV2OverridesType = "OVERWRITE"
	ContractV2OverridesTypeMultiplier ContractV2OverridesType = "MULTIPLIER"
	ContractV2OverridesTypeTiered     ContractV2OverridesType = "TIERED"
)

func (r ContractV2OverridesType) IsKnown() bool {
	switch r {
	case ContractV2OverridesTypeOverwrite, ContractV2OverridesTypeMultiplier, ContractV2OverridesTypeTiered:
		return true
	}
	return false
}

type ContractV2Transition struct {
	FromContractID string                    `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                    `json:"to_contract_id,required" format:"uuid"`
	Type           ContractV2TransitionsType `json:"type,required"`
	JSON           contractV2TransitionJSON  `json:"-"`
}

// contractV2TransitionJSON contains the JSON metadata for the struct
// [ContractV2Transition]
type contractV2TransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractV2Transition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2TransitionJSON) RawJSON() string {
	return r.raw
}

type ContractV2TransitionsType string

const (
	ContractV2TransitionsTypeSupersede ContractV2TransitionsType = "SUPERSEDE"
	ContractV2TransitionsTypeRenewal   ContractV2TransitionsType = "RENEWAL"
)

func (r ContractV2TransitionsType) IsKnown() bool {
	switch r {
	case ContractV2TransitionsTypeSupersede, ContractV2TransitionsTypeRenewal:
		return true
	}
	return false
}

type ContractV2UsageFilter struct {
	GroupKey    string   `json:"group_key,required"`
	GroupValues []string `json:"group_values,required"`
	// This will match contract starting_at value if usage filter is active from the
	// beginning of the contract.
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// This will match contract ending_before value if usage filter is active until the
	// end of the contract. It will be undefined if the contract is open-ended.
	EndingBefore time.Time                 `json:"ending_before" format:"date-time"`
	JSON         contractV2UsageFilterJSON `json:"-"`
}

// contractV2UsageFilterJSON contains the JSON metadata for the struct
// [ContractV2UsageFilter]
type contractV2UsageFilterJSON struct {
	GroupKey     apijson.Field
	GroupValues  apijson.Field
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractV2UsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2UsageFilterJSON) RawJSON() string {
	return r.raw
}

type ContractV2UsageStatementSchedule struct {
	// Contract usage statements follow a selected cadence based on this date.
	BillingAnchorDate time.Time                                 `json:"billing_anchor_date,required" format:"date-time"`
	Frequency         ContractV2UsageStatementScheduleFrequency `json:"frequency,required"`
	JSON              contractV2UsageStatementScheduleJSON      `json:"-"`
}

// contractV2UsageStatementScheduleJSON contains the JSON metadata for the struct
// [ContractV2UsageStatementSchedule]
type contractV2UsageStatementScheduleJSON struct {
	BillingAnchorDate apijson.Field
	Frequency         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ContractV2UsageStatementSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2UsageStatementScheduleJSON) RawJSON() string {
	return r.raw
}

type ContractV2UsageStatementScheduleFrequency string

const (
	ContractV2UsageStatementScheduleFrequencyMonthly   ContractV2UsageStatementScheduleFrequency = "MONTHLY"
	ContractV2UsageStatementScheduleFrequencyQuarterly ContractV2UsageStatementScheduleFrequency = "QUARTERLY"
	ContractV2UsageStatementScheduleFrequencyAnnual    ContractV2UsageStatementScheduleFrequency = "ANNUAL"
	ContractV2UsageStatementScheduleFrequencyWeekly    ContractV2UsageStatementScheduleFrequency = "WEEKLY"
)

func (r ContractV2UsageStatementScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractV2UsageStatementScheduleFrequencyMonthly, ContractV2UsageStatementScheduleFrequencyQuarterly, ContractV2UsageStatementScheduleFrequencyAnnual, ContractV2UsageStatementScheduleFrequencyWeekly:
		return true
	}
	return false
}

type ContractV2Credit struct {
	ID      string                   `json:"id,required" format:"uuid"`
	Product ContractV2CreditsProduct `json:"product,required"`
	Type    ContractV2CreditsType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        ScheduleDuration `json:"access_schedule"`
	ApplicableContractIDs []string         `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string         `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string         `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64                   `json:"balance"`
	Contract     ContractV2CreditsContract `json:"contract"`
	CustomFields map[string]string         `json:"custom_fields"`
	Description  string                    `json:"description"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []ContractV2CreditsLedger `json:"ledger"`
	Name   string                    `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier    `json:"specifiers"`
	JSON       contractV2CreditJSON `json:"-"`
}

// contractV2CreditJSON contains the JSON metadata for the struct
// [ContractV2Credit]
type contractV2CreditJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Balance                 apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	HierarchyConfiguration  apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	SalesforceOpportunityID apijson.Field
	Specifiers              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractV2Credit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CreditJSON) RawJSON() string {
	return r.raw
}

type ContractV2CreditsProduct struct {
	ID   string                       `json:"id,required" format:"uuid"`
	Name string                       `json:"name,required"`
	JSON contractV2CreditsProductJSON `json:"-"`
}

// contractV2CreditsProductJSON contains the JSON metadata for the struct
// [ContractV2CreditsProduct]
type contractV2CreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CreditsProductJSON) RawJSON() string {
	return r.raw
}

type ContractV2CreditsType string

const (
	ContractV2CreditsTypeCredit ContractV2CreditsType = "CREDIT"
)

func (r ContractV2CreditsType) IsKnown() bool {
	switch r {
	case ContractV2CreditsTypeCredit:
		return true
	}
	return false
}

type ContractV2CreditsContract struct {
	ID   string                        `json:"id,required" format:"uuid"`
	JSON contractV2CreditsContractJSON `json:"-"`
}

// contractV2CreditsContractJSON contains the JSON metadata for the struct
// [ContractV2CreditsContract]
type contractV2CreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CreditsContractJSON) RawJSON() string {
	return r.raw
}

type ContractV2CreditsLedger struct {
	Amount     float64                     `json:"amount,required"`
	Timestamp  time.Time                   `json:"timestamp,required" format:"date-time"`
	Type       ContractV2CreditsLedgerType `json:"type,required"`
	ContractID string                      `json:"contract_id" format:"uuid"`
	InvoiceID  string                      `json:"invoice_id" format:"uuid"`
	Reason     string                      `json:"reason"`
	SegmentID  string                      `json:"segment_id" format:"uuid"`
	JSON       contractV2CreditsLedgerJSON `json:"-"`
	union      ContractV2CreditsLedgerUnion
}

// contractV2CreditsLedgerJSON contains the JSON metadata for the struct
// [ContractV2CreditsLedger]
type contractV2CreditsLedgerJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	InvoiceID   apijson.Field
	Reason      apijson.Field
	SegmentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r contractV2CreditsLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *ContractV2CreditsLedger) UnmarshalJSON(data []byte) (err error) {
	*r = ContractV2CreditsLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContractV2CreditsLedgerUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ContractV2CreditsLedgerCreditSegmentStartLedgerEntry],
// [ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [ContractV2CreditsLedgerCreditExpirationLedgerEntry],
// [ContractV2CreditsLedgerCreditCanceledLedgerEntry],
// [ContractV2CreditsLedgerCreditCreditedLedgerEntry],
// [ContractV2CreditsLedgerCreditManualLedgerEntry],
// [ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntry].
func (r ContractV2CreditsLedger) AsUnion() ContractV2CreditsLedgerUnion {
	return r.union
}

// Union satisfied by [ContractV2CreditsLedgerCreditSegmentStartLedgerEntry],
// [ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [ContractV2CreditsLedgerCreditExpirationLedgerEntry],
// [ContractV2CreditsLedgerCreditCanceledLedgerEntry],
// [ContractV2CreditsLedgerCreditCreditedLedgerEntry],
// [ContractV2CreditsLedgerCreditManualLedgerEntry] or
// [ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntry].
type ContractV2CreditsLedgerUnion interface {
	implementsContractV2CreditsLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContractV2CreditsLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CreditsLedgerCreditSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CreditsLedgerCreditExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CreditsLedgerCreditCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CreditsLedgerCreditCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CreditsLedgerCreditManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntry{}),
		},
	)
}

type ContractV2CreditsLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64                                                  `json:"amount,required"`
	SegmentID string                                                   `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                `json:"timestamp,required" format:"date-time"`
	Type      ContractV2CreditsLedgerCreditSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractV2CreditsLedgerCreditSegmentStartLedgerEntryJSON `json:"-"`
}

// contractV2CreditsLedgerCreditSegmentStartLedgerEntryJSON contains the JSON
// metadata for the struct [ContractV2CreditsLedgerCreditSegmentStartLedgerEntry]
type contractV2CreditsLedgerCreditSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CreditsLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CreditsLedgerCreditSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CreditsLedgerCreditSegmentStartLedgerEntry) implementsContractV2CreditsLedger() {}

type ContractV2CreditsLedgerCreditSegmentStartLedgerEntryType string

const (
	ContractV2CreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart ContractV2CreditsLedgerCreditSegmentStartLedgerEntryType = "CREDIT_SEGMENT_START"
)

func (r ContractV2CreditsLedgerCreditSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart:
		return true
	}
	return false
}

type ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount     float64                                                               `json:"amount,required"`
	InvoiceID  string                                                                `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                                                `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                                             `json:"timestamp,required" format:"date-time"`
	Type       ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	ContractID string                                                                `json:"contract_id" format:"uuid"`
	JSON       contractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// contractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON contains
// the JSON metadata for the struct
// [ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry]
type contractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) implementsContractV2CreditsLedger() {
}

type ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type ContractV2CreditsLedgerCreditExpirationLedgerEntry struct {
	Amount    float64                                                `json:"amount,required"`
	SegmentID string                                                 `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                              `json:"timestamp,required" format:"date-time"`
	Type      ContractV2CreditsLedgerCreditExpirationLedgerEntryType `json:"type,required"`
	JSON      contractV2CreditsLedgerCreditExpirationLedgerEntryJSON `json:"-"`
}

// contractV2CreditsLedgerCreditExpirationLedgerEntryJSON contains the JSON
// metadata for the struct [ContractV2CreditsLedgerCreditExpirationLedgerEntry]
type contractV2CreditsLedgerCreditExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CreditsLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CreditsLedgerCreditExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CreditsLedgerCreditExpirationLedgerEntry) implementsContractV2CreditsLedger() {}

type ContractV2CreditsLedgerCreditExpirationLedgerEntryType string

const (
	ContractV2CreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration ContractV2CreditsLedgerCreditExpirationLedgerEntryType = "CREDIT_EXPIRATION"
)

func (r ContractV2CreditsLedgerCreditExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration:
		return true
	}
	return false
}

type ContractV2CreditsLedgerCreditCanceledLedgerEntry struct {
	Amount     float64                                              `json:"amount,required"`
	InvoiceID  string                                               `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                               `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                            `json:"timestamp,required" format:"date-time"`
	Type       ContractV2CreditsLedgerCreditCanceledLedgerEntryType `json:"type,required"`
	ContractID string                                               `json:"contract_id" format:"uuid"`
	JSON       contractV2CreditsLedgerCreditCanceledLedgerEntryJSON `json:"-"`
}

// contractV2CreditsLedgerCreditCanceledLedgerEntryJSON contains the JSON metadata
// for the struct [ContractV2CreditsLedgerCreditCanceledLedgerEntry]
type contractV2CreditsLedgerCreditCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CreditsLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CreditsLedgerCreditCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CreditsLedgerCreditCanceledLedgerEntry) implementsContractV2CreditsLedger() {}

type ContractV2CreditsLedgerCreditCanceledLedgerEntryType string

const (
	ContractV2CreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled ContractV2CreditsLedgerCreditCanceledLedgerEntryType = "CREDIT_CANCELED"
)

func (r ContractV2CreditsLedgerCreditCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled:
		return true
	}
	return false
}

type ContractV2CreditsLedgerCreditCreditedLedgerEntry struct {
	Amount     float64                                              `json:"amount,required"`
	InvoiceID  string                                               `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                               `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                            `json:"timestamp,required" format:"date-time"`
	Type       ContractV2CreditsLedgerCreditCreditedLedgerEntryType `json:"type,required"`
	ContractID string                                               `json:"contract_id" format:"uuid"`
	JSON       contractV2CreditsLedgerCreditCreditedLedgerEntryJSON `json:"-"`
}

// contractV2CreditsLedgerCreditCreditedLedgerEntryJSON contains the JSON metadata
// for the struct [ContractV2CreditsLedgerCreditCreditedLedgerEntry]
type contractV2CreditsLedgerCreditCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CreditsLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CreditsLedgerCreditCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CreditsLedgerCreditCreditedLedgerEntry) implementsContractV2CreditsLedger() {}

type ContractV2CreditsLedgerCreditCreditedLedgerEntryType string

const (
	ContractV2CreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited ContractV2CreditsLedgerCreditCreditedLedgerEntryType = "CREDIT_CREDITED"
)

func (r ContractV2CreditsLedgerCreditCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited:
		return true
	}
	return false
}

type ContractV2CreditsLedgerCreditManualLedgerEntry struct {
	Amount    float64                                            `json:"amount,required"`
	Reason    string                                             `json:"reason,required"`
	Timestamp time.Time                                          `json:"timestamp,required" format:"date-time"`
	Type      ContractV2CreditsLedgerCreditManualLedgerEntryType `json:"type,required"`
	JSON      contractV2CreditsLedgerCreditManualLedgerEntryJSON `json:"-"`
}

// contractV2CreditsLedgerCreditManualLedgerEntryJSON contains the JSON metadata
// for the struct [ContractV2CreditsLedgerCreditManualLedgerEntry]
type contractV2CreditsLedgerCreditManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CreditsLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CreditsLedgerCreditManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CreditsLedgerCreditManualLedgerEntry) implementsContractV2CreditsLedger() {}

type ContractV2CreditsLedgerCreditManualLedgerEntryType string

const (
	ContractV2CreditsLedgerCreditManualLedgerEntryTypeCreditManual ContractV2CreditsLedgerCreditManualLedgerEntryType = "CREDIT_MANUAL"
)

func (r ContractV2CreditsLedgerCreditManualLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CreditsLedgerCreditManualLedgerEntryTypeCreditManual:
		return true
	}
	return false
}

type ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntry struct {
	Amount    float64                                                         `json:"amount,required"`
	SegmentID string                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryType `json:"type,required"`
	JSON      contractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryJSON `json:"-"`
}

// contractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryJSON contains the
// JSON metadata for the struct
// [ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntry]
type contractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntry) implementsContractV2CreditsLedger() {
}

type ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryType string

const (
	ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryTypeCreditSeatBasedAdjustment ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryType = "CREDIT_SEAT_BASED_ADJUSTMENT"
)

func (r ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractV2CreditsLedgerCreditSeatBasedAdjustmentLedgerEntryTypeCreditSeatBasedAdjustment:
		return true
	}
	return false
}

type ContractV2CreditsLedgerType string

const (
	ContractV2CreditsLedgerTypeCreditSegmentStart              ContractV2CreditsLedgerType = "CREDIT_SEGMENT_START"
	ContractV2CreditsLedgerTypeCreditAutomatedInvoiceDeduction ContractV2CreditsLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	ContractV2CreditsLedgerTypeCreditExpiration                ContractV2CreditsLedgerType = "CREDIT_EXPIRATION"
	ContractV2CreditsLedgerTypeCreditCanceled                  ContractV2CreditsLedgerType = "CREDIT_CANCELED"
	ContractV2CreditsLedgerTypeCreditCredited                  ContractV2CreditsLedgerType = "CREDIT_CREDITED"
	ContractV2CreditsLedgerTypeCreditManual                    ContractV2CreditsLedgerType = "CREDIT_MANUAL"
	ContractV2CreditsLedgerTypeCreditSeatBasedAdjustment       ContractV2CreditsLedgerType = "CREDIT_SEAT_BASED_ADJUSTMENT"
)

func (r ContractV2CreditsLedgerType) IsKnown() bool {
	switch r {
	case ContractV2CreditsLedgerTypeCreditSegmentStart, ContractV2CreditsLedgerTypeCreditAutomatedInvoiceDeduction, ContractV2CreditsLedgerTypeCreditExpiration, ContractV2CreditsLedgerTypeCreditCanceled, ContractV2CreditsLedgerTypeCreditCredited, ContractV2CreditsLedgerTypeCreditManual, ContractV2CreditsLedgerTypeCreditSeatBasedAdjustment:
		return true
	}
	return false
}

// This field's availability is dependent on your client's configuration.
type ContractV2CustomerBillingProviderConfiguration struct {
	// ID of Customer's billing provider configuration.
	ID              string                                                        `json:"id,required" format:"uuid"`
	BillingProvider ContractV2CustomerBillingProviderConfigurationBillingProvider `json:"billing_provider,required"`
	DeliveryMethod  ContractV2CustomerBillingProviderConfigurationDeliveryMethod  `json:"delivery_method,required"`
	JSON            contractV2CustomerBillingProviderConfigurationJSON            `json:"-"`
}

// contractV2CustomerBillingProviderConfigurationJSON contains the JSON metadata
// for the struct [ContractV2CustomerBillingProviderConfiguration]
type contractV2CustomerBillingProviderConfigurationJSON struct {
	ID              apijson.Field
	BillingProvider apijson.Field
	DeliveryMethod  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ContractV2CustomerBillingProviderConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2CustomerBillingProviderConfigurationJSON) RawJSON() string {
	return r.raw
}

type ContractV2CustomerBillingProviderConfigurationBillingProvider string

const (
	ContractV2CustomerBillingProviderConfigurationBillingProviderAwsMarketplace   ContractV2CustomerBillingProviderConfigurationBillingProvider = "aws_marketplace"
	ContractV2CustomerBillingProviderConfigurationBillingProviderStripe           ContractV2CustomerBillingProviderConfigurationBillingProvider = "stripe"
	ContractV2CustomerBillingProviderConfigurationBillingProviderNetsuite         ContractV2CustomerBillingProviderConfigurationBillingProvider = "netsuite"
	ContractV2CustomerBillingProviderConfigurationBillingProviderCustom           ContractV2CustomerBillingProviderConfigurationBillingProvider = "custom"
	ContractV2CustomerBillingProviderConfigurationBillingProviderAzureMarketplace ContractV2CustomerBillingProviderConfigurationBillingProvider = "azure_marketplace"
	ContractV2CustomerBillingProviderConfigurationBillingProviderQuickbooksOnline ContractV2CustomerBillingProviderConfigurationBillingProvider = "quickbooks_online"
	ContractV2CustomerBillingProviderConfigurationBillingProviderWorkday          ContractV2CustomerBillingProviderConfigurationBillingProvider = "workday"
	ContractV2CustomerBillingProviderConfigurationBillingProviderGcpMarketplace   ContractV2CustomerBillingProviderConfigurationBillingProvider = "gcp_marketplace"
)

func (r ContractV2CustomerBillingProviderConfigurationBillingProvider) IsKnown() bool {
	switch r {
	case ContractV2CustomerBillingProviderConfigurationBillingProviderAwsMarketplace, ContractV2CustomerBillingProviderConfigurationBillingProviderStripe, ContractV2CustomerBillingProviderConfigurationBillingProviderNetsuite, ContractV2CustomerBillingProviderConfigurationBillingProviderCustom, ContractV2CustomerBillingProviderConfigurationBillingProviderAzureMarketplace, ContractV2CustomerBillingProviderConfigurationBillingProviderQuickbooksOnline, ContractV2CustomerBillingProviderConfigurationBillingProviderWorkday, ContractV2CustomerBillingProviderConfigurationBillingProviderGcpMarketplace:
		return true
	}
	return false
}

type ContractV2CustomerBillingProviderConfigurationDeliveryMethod string

const (
	ContractV2CustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider ContractV2CustomerBillingProviderConfigurationDeliveryMethod = "direct_to_billing_provider"
	ContractV2CustomerBillingProviderConfigurationDeliveryMethodAwsSqs                  ContractV2CustomerBillingProviderConfigurationDeliveryMethod = "aws_sqs"
	ContractV2CustomerBillingProviderConfigurationDeliveryMethodTackle                  ContractV2CustomerBillingProviderConfigurationDeliveryMethod = "tackle"
	ContractV2CustomerBillingProviderConfigurationDeliveryMethodAwsSns                  ContractV2CustomerBillingProviderConfigurationDeliveryMethod = "aws_sns"
)

func (r ContractV2CustomerBillingProviderConfigurationDeliveryMethod) IsKnown() bool {
	switch r {
	case ContractV2CustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider, ContractV2CustomerBillingProviderConfigurationDeliveryMethodAwsSqs, ContractV2CustomerBillingProviderConfigurationDeliveryMethodTackle, ContractV2CustomerBillingProviderConfigurationDeliveryMethodAwsSns:
		return true
	}
	return false
}

// Indicates whether there are more items than the limit for this endpoint. Use the
// respective list endpoints to get the full lists.
type ContractV2HasMore struct {
	// Whether there are more commits on this contract than the limit for this
	// endpoint. Use the /contracts/customerCommits/list endpoint to get the full list
	// of commits.
	Commits bool `json:"commits,required"`
	// Whether there are more credits on this contract than the limit for this
	// endpoint. Use the /contracts/customerCredits/list endpoint to get the full list
	// of credits.
	Credits bool                  `json:"credits,required"`
	JSON    contractV2HasMoreJSON `json:"-"`
}

// contractV2HasMoreJSON contains the JSON metadata for the struct
// [ContractV2HasMore]
type contractV2HasMoreJSON struct {
	Commits     apijson.Field
	Credits     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2HasMore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2HasMoreJSON) RawJSON() string {
	return r.raw
}

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
// prices automatically. EXPLICIT prioritization requires specifying priorities for
// each multiplier; the one with the lowest priority value will be prioritized
// first.
type ContractV2MultiplierOverridePrioritization string

const (
	ContractV2MultiplierOverridePrioritizationLowestMultiplier ContractV2MultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	ContractV2MultiplierOverridePrioritizationExplicit         ContractV2MultiplierOverridePrioritization = "EXPLICIT"
)

func (r ContractV2MultiplierOverridePrioritization) IsKnown() bool {
	switch r {
	case ContractV2MultiplierOverridePrioritizationLowestMultiplier, ContractV2MultiplierOverridePrioritizationExplicit:
		return true
	}
	return false
}

type ContractV2RecurringCommit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount ContractV2RecurringCommitsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration ContractV2RecurringCommitsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                           `json:"priority,required"`
	Product  ContractV2RecurringCommitsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType ContractV2RecurringCommitsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                           `json:"applicable_product_tags"`
	Contract              ContractV2RecurringCommitsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount ContractV2RecurringCommitsInvoiceAmount `json:"invoice_amount"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration ContractV2RecurringCommitsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	RecurrenceFrequency ContractV2RecurringCommitsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig ContractV2RecurringCommitsSubscriptionConfig `json:"subscription_config"`
	JSON               contractV2RecurringCommitJSON                `json:"-"`
}

// contractV2RecurringCommitJSON contains the JSON metadata for the struct
// [ContractV2RecurringCommit]
type contractV2RecurringCommitJSON struct {
	ID                     apijson.Field
	AccessAmount           apijson.Field
	CommitDuration         apijson.Field
	Priority               apijson.Field
	Product                apijson.Field
	RateType               apijson.Field
	StartingAt             apijson.Field
	ApplicableProductIDs   apijson.Field
	ApplicableProductTags  apijson.Field
	Contract               apijson.Field
	Description            apijson.Field
	EndingBefore           apijson.Field
	HierarchyConfiguration apijson.Field
	InvoiceAmount          apijson.Field
	Name                   apijson.Field
	NetsuiteSalesOrderID   apijson.Field
	Proration              apijson.Field
	RecurrenceFrequency    apijson.Field
	RolloverFraction       apijson.Field
	Specifiers             apijson.Field
	SubscriptionConfig     apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ContractV2RecurringCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCommitJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type ContractV2RecurringCommitsAccessAmount struct {
	CreditTypeID string                                     `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64                                    `json:"unit_price,required"`
	Quantity     float64                                    `json:"quantity"`
	JSON         contractV2RecurringCommitsAccessAmountJSON `json:"-"`
}

// contractV2RecurringCommitsAccessAmountJSON contains the JSON metadata for the
// struct [ContractV2RecurringCommitsAccessAmount]
type contractV2RecurringCommitsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	UnitPrice    apijson.Field
	Quantity     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractV2RecurringCommitsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCommitsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type ContractV2RecurringCommitsCommitDuration struct {
	Value float64                                      `json:"value,required"`
	Unit  ContractV2RecurringCommitsCommitDurationUnit `json:"unit"`
	JSON  contractV2RecurringCommitsCommitDurationJSON `json:"-"`
}

// contractV2RecurringCommitsCommitDurationJSON contains the JSON metadata for the
// struct [ContractV2RecurringCommitsCommitDuration]
type contractV2RecurringCommitsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2RecurringCommitsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCommitsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type ContractV2RecurringCommitsCommitDurationUnit string

const (
	ContractV2RecurringCommitsCommitDurationUnitPeriods ContractV2RecurringCommitsCommitDurationUnit = "PERIODS"
)

func (r ContractV2RecurringCommitsCommitDurationUnit) IsKnown() bool {
	switch r {
	case ContractV2RecurringCommitsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type ContractV2RecurringCommitsProduct struct {
	ID   string                                `json:"id,required" format:"uuid"`
	Name string                                `json:"name,required"`
	JSON contractV2RecurringCommitsProductJSON `json:"-"`
}

// contractV2RecurringCommitsProductJSON contains the JSON metadata for the struct
// [ContractV2RecurringCommitsProduct]
type contractV2RecurringCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2RecurringCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCommitsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type ContractV2RecurringCommitsRateType string

const (
	ContractV2RecurringCommitsRateTypeCommitRate ContractV2RecurringCommitsRateType = "COMMIT_RATE"
	ContractV2RecurringCommitsRateTypeListRate   ContractV2RecurringCommitsRateType = "LIST_RATE"
)

func (r ContractV2RecurringCommitsRateType) IsKnown() bool {
	switch r {
	case ContractV2RecurringCommitsRateTypeCommitRate, ContractV2RecurringCommitsRateTypeListRate:
		return true
	}
	return false
}

type ContractV2RecurringCommitsContract struct {
	ID   string                                 `json:"id,required" format:"uuid"`
	JSON contractV2RecurringCommitsContractJSON `json:"-"`
}

// contractV2RecurringCommitsContractJSON contains the JSON metadata for the struct
// [ContractV2RecurringCommitsContract]
type contractV2RecurringCommitsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2RecurringCommitsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCommitsContractJSON) RawJSON() string {
	return r.raw
}

// The amount the customer should be billed for the commit. Not required.
type ContractV2RecurringCommitsInvoiceAmount struct {
	CreditTypeID string                                      `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                     `json:"quantity,required"`
	UnitPrice    float64                                     `json:"unit_price,required"`
	JSON         contractV2RecurringCommitsInvoiceAmountJSON `json:"-"`
}

// contractV2RecurringCommitsInvoiceAmountJSON contains the JSON metadata for the
// struct [ContractV2RecurringCommitsInvoiceAmount]
type contractV2RecurringCommitsInvoiceAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractV2RecurringCommitsInvoiceAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCommitsInvoiceAmountJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type ContractV2RecurringCommitsProration string

const (
	ContractV2RecurringCommitsProrationNone         ContractV2RecurringCommitsProration = "NONE"
	ContractV2RecurringCommitsProrationFirst        ContractV2RecurringCommitsProration = "FIRST"
	ContractV2RecurringCommitsProrationLast         ContractV2RecurringCommitsProration = "LAST"
	ContractV2RecurringCommitsProrationFirstAndLast ContractV2RecurringCommitsProration = "FIRST_AND_LAST"
)

func (r ContractV2RecurringCommitsProration) IsKnown() bool {
	switch r {
	case ContractV2RecurringCommitsProrationNone, ContractV2RecurringCommitsProrationFirst, ContractV2RecurringCommitsProrationLast, ContractV2RecurringCommitsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's starting_at rather than the usage
// invoice dates.
type ContractV2RecurringCommitsRecurrenceFrequency string

const (
	ContractV2RecurringCommitsRecurrenceFrequencyMonthly   ContractV2RecurringCommitsRecurrenceFrequency = "MONTHLY"
	ContractV2RecurringCommitsRecurrenceFrequencyQuarterly ContractV2RecurringCommitsRecurrenceFrequency = "QUARTERLY"
	ContractV2RecurringCommitsRecurrenceFrequencyAnnual    ContractV2RecurringCommitsRecurrenceFrequency = "ANNUAL"
	ContractV2RecurringCommitsRecurrenceFrequencyWeekly    ContractV2RecurringCommitsRecurrenceFrequency = "WEEKLY"
)

func (r ContractV2RecurringCommitsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case ContractV2RecurringCommitsRecurrenceFrequencyMonthly, ContractV2RecurringCommitsRecurrenceFrequencyQuarterly, ContractV2RecurringCommitsRecurrenceFrequencyAnnual, ContractV2RecurringCommitsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

// Attach a subscription to the recurring commit/credit.
type ContractV2RecurringCommitsSubscriptionConfig struct {
	Allocation              ContractV2RecurringCommitsSubscriptionConfigAllocation              `json:"allocation,required"`
	ApplySeatIncreaseConfig ContractV2RecurringCommitsSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,required"`
	SubscriptionID          string                                                              `json:"subscription_id,required" format:"uuid"`
	JSON                    contractV2RecurringCommitsSubscriptionConfigJSON                    `json:"-"`
}

// contractV2RecurringCommitsSubscriptionConfigJSON contains the JSON metadata for
// the struct [ContractV2RecurringCommitsSubscriptionConfig]
type contractV2RecurringCommitsSubscriptionConfigJSON struct {
	Allocation              apijson.Field
	ApplySeatIncreaseConfig apijson.Field
	SubscriptionID          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractV2RecurringCommitsSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCommitsSubscriptionConfigJSON) RawJSON() string {
	return r.raw
}

type ContractV2RecurringCommitsSubscriptionConfigAllocation string

const (
	ContractV2RecurringCommitsSubscriptionConfigAllocationIndividual ContractV2RecurringCommitsSubscriptionConfigAllocation = "INDIVIDUAL"
	ContractV2RecurringCommitsSubscriptionConfigAllocationPooled     ContractV2RecurringCommitsSubscriptionConfigAllocation = "POOLED"
)

func (r ContractV2RecurringCommitsSubscriptionConfigAllocation) IsKnown() bool {
	switch r {
	case ContractV2RecurringCommitsSubscriptionConfigAllocationIndividual, ContractV2RecurringCommitsSubscriptionConfigAllocationPooled:
		return true
	}
	return false
}

type ContractV2RecurringCommitsSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool                                                                    `json:"is_prorated,required"`
	JSON       contractV2RecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON `json:"-"`
}

// contractV2RecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON contains
// the JSON metadata for the struct
// [ContractV2RecurringCommitsSubscriptionConfigApplySeatIncreaseConfig]
type contractV2RecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON struct {
	IsProrated  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2RecurringCommitsSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON) RawJSON() string {
	return r.raw
}

type ContractV2RecurringCredit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount ContractV2RecurringCreditsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration ContractV2RecurringCreditsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                           `json:"priority,required"`
	Product  ContractV2RecurringCreditsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType ContractV2RecurringCreditsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                           `json:"applicable_product_tags"`
	Contract              ContractV2RecurringCreditsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration ContractV2RecurringCreditsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	RecurrenceFrequency ContractV2RecurringCreditsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig ContractV2RecurringCreditsSubscriptionConfig `json:"subscription_config"`
	JSON               contractV2RecurringCreditJSON                `json:"-"`
}

// contractV2RecurringCreditJSON contains the JSON metadata for the struct
// [ContractV2RecurringCredit]
type contractV2RecurringCreditJSON struct {
	ID                     apijson.Field
	AccessAmount           apijson.Field
	CommitDuration         apijson.Field
	Priority               apijson.Field
	Product                apijson.Field
	RateType               apijson.Field
	StartingAt             apijson.Field
	ApplicableProductIDs   apijson.Field
	ApplicableProductTags  apijson.Field
	Contract               apijson.Field
	Description            apijson.Field
	EndingBefore           apijson.Field
	HierarchyConfiguration apijson.Field
	Name                   apijson.Field
	NetsuiteSalesOrderID   apijson.Field
	Proration              apijson.Field
	RecurrenceFrequency    apijson.Field
	RolloverFraction       apijson.Field
	Specifiers             apijson.Field
	SubscriptionConfig     apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ContractV2RecurringCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCreditJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type ContractV2RecurringCreditsAccessAmount struct {
	CreditTypeID string                                     `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64                                    `json:"unit_price,required"`
	Quantity     float64                                    `json:"quantity"`
	JSON         contractV2RecurringCreditsAccessAmountJSON `json:"-"`
}

// contractV2RecurringCreditsAccessAmountJSON contains the JSON metadata for the
// struct [ContractV2RecurringCreditsAccessAmount]
type contractV2RecurringCreditsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	UnitPrice    apijson.Field
	Quantity     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractV2RecurringCreditsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCreditsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type ContractV2RecurringCreditsCommitDuration struct {
	Value float64                                      `json:"value,required"`
	Unit  ContractV2RecurringCreditsCommitDurationUnit `json:"unit"`
	JSON  contractV2RecurringCreditsCommitDurationJSON `json:"-"`
}

// contractV2RecurringCreditsCommitDurationJSON contains the JSON metadata for the
// struct [ContractV2RecurringCreditsCommitDuration]
type contractV2RecurringCreditsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2RecurringCreditsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCreditsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type ContractV2RecurringCreditsCommitDurationUnit string

const (
	ContractV2RecurringCreditsCommitDurationUnitPeriods ContractV2RecurringCreditsCommitDurationUnit = "PERIODS"
)

func (r ContractV2RecurringCreditsCommitDurationUnit) IsKnown() bool {
	switch r {
	case ContractV2RecurringCreditsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type ContractV2RecurringCreditsProduct struct {
	ID   string                                `json:"id,required" format:"uuid"`
	Name string                                `json:"name,required"`
	JSON contractV2RecurringCreditsProductJSON `json:"-"`
}

// contractV2RecurringCreditsProductJSON contains the JSON metadata for the struct
// [ContractV2RecurringCreditsProduct]
type contractV2RecurringCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2RecurringCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCreditsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type ContractV2RecurringCreditsRateType string

const (
	ContractV2RecurringCreditsRateTypeCommitRate ContractV2RecurringCreditsRateType = "COMMIT_RATE"
	ContractV2RecurringCreditsRateTypeListRate   ContractV2RecurringCreditsRateType = "LIST_RATE"
)

func (r ContractV2RecurringCreditsRateType) IsKnown() bool {
	switch r {
	case ContractV2RecurringCreditsRateTypeCommitRate, ContractV2RecurringCreditsRateTypeListRate:
		return true
	}
	return false
}

type ContractV2RecurringCreditsContract struct {
	ID   string                                 `json:"id,required" format:"uuid"`
	JSON contractV2RecurringCreditsContractJSON `json:"-"`
}

// contractV2RecurringCreditsContractJSON contains the JSON metadata for the struct
// [ContractV2RecurringCreditsContract]
type contractV2RecurringCreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2RecurringCreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCreditsContractJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type ContractV2RecurringCreditsProration string

const (
	ContractV2RecurringCreditsProrationNone         ContractV2RecurringCreditsProration = "NONE"
	ContractV2RecurringCreditsProrationFirst        ContractV2RecurringCreditsProration = "FIRST"
	ContractV2RecurringCreditsProrationLast         ContractV2RecurringCreditsProration = "LAST"
	ContractV2RecurringCreditsProrationFirstAndLast ContractV2RecurringCreditsProration = "FIRST_AND_LAST"
)

func (r ContractV2RecurringCreditsProration) IsKnown() bool {
	switch r {
	case ContractV2RecurringCreditsProrationNone, ContractV2RecurringCreditsProrationFirst, ContractV2RecurringCreditsProrationLast, ContractV2RecurringCreditsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's starting_at rather than the usage
// invoice dates.
type ContractV2RecurringCreditsRecurrenceFrequency string

const (
	ContractV2RecurringCreditsRecurrenceFrequencyMonthly   ContractV2RecurringCreditsRecurrenceFrequency = "MONTHLY"
	ContractV2RecurringCreditsRecurrenceFrequencyQuarterly ContractV2RecurringCreditsRecurrenceFrequency = "QUARTERLY"
	ContractV2RecurringCreditsRecurrenceFrequencyAnnual    ContractV2RecurringCreditsRecurrenceFrequency = "ANNUAL"
	ContractV2RecurringCreditsRecurrenceFrequencyWeekly    ContractV2RecurringCreditsRecurrenceFrequency = "WEEKLY"
)

func (r ContractV2RecurringCreditsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case ContractV2RecurringCreditsRecurrenceFrequencyMonthly, ContractV2RecurringCreditsRecurrenceFrequencyQuarterly, ContractV2RecurringCreditsRecurrenceFrequencyAnnual, ContractV2RecurringCreditsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

// Attach a subscription to the recurring commit/credit.
type ContractV2RecurringCreditsSubscriptionConfig struct {
	Allocation              ContractV2RecurringCreditsSubscriptionConfigAllocation              `json:"allocation,required"`
	ApplySeatIncreaseConfig ContractV2RecurringCreditsSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,required"`
	SubscriptionID          string                                                              `json:"subscription_id,required" format:"uuid"`
	JSON                    contractV2RecurringCreditsSubscriptionConfigJSON                    `json:"-"`
}

// contractV2RecurringCreditsSubscriptionConfigJSON contains the JSON metadata for
// the struct [ContractV2RecurringCreditsSubscriptionConfig]
type contractV2RecurringCreditsSubscriptionConfigJSON struct {
	Allocation              apijson.Field
	ApplySeatIncreaseConfig apijson.Field
	SubscriptionID          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractV2RecurringCreditsSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCreditsSubscriptionConfigJSON) RawJSON() string {
	return r.raw
}

type ContractV2RecurringCreditsSubscriptionConfigAllocation string

const (
	ContractV2RecurringCreditsSubscriptionConfigAllocationIndividual ContractV2RecurringCreditsSubscriptionConfigAllocation = "INDIVIDUAL"
	ContractV2RecurringCreditsSubscriptionConfigAllocationPooled     ContractV2RecurringCreditsSubscriptionConfigAllocation = "POOLED"
)

func (r ContractV2RecurringCreditsSubscriptionConfigAllocation) IsKnown() bool {
	switch r {
	case ContractV2RecurringCreditsSubscriptionConfigAllocationIndividual, ContractV2RecurringCreditsSubscriptionConfigAllocationPooled:
		return true
	}
	return false
}

type ContractV2RecurringCreditsSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool                                                                    `json:"is_prorated,required"`
	JSON       contractV2RecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON `json:"-"`
}

// contractV2RecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON contains
// the JSON metadata for the struct
// [ContractV2RecurringCreditsSubscriptionConfigApplySeatIncreaseConfig]
type contractV2RecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON struct {
	IsProrated  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractV2RecurringCreditsSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2RecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON) RawJSON() string {
	return r.raw
}

type ContractV2ResellerRoyalty struct {
	ResellerType ContractV2ResellerRoyaltiesResellerType `json:"reseller_type,required"`
	Segments     []ContractV2ResellerRoyaltiesSegment    `json:"segments,required"`
	JSON         contractV2ResellerRoyaltyJSON           `json:"-"`
}

// contractV2ResellerRoyaltyJSON contains the JSON metadata for the struct
// [ContractV2ResellerRoyalty]
type contractV2ResellerRoyaltyJSON struct {
	ResellerType apijson.Field
	Segments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractV2ResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2ResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type ContractV2ResellerRoyaltiesResellerType string

const (
	ContractV2ResellerRoyaltiesResellerTypeAws           ContractV2ResellerRoyaltiesResellerType = "AWS"
	ContractV2ResellerRoyaltiesResellerTypeAwsProService ContractV2ResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	ContractV2ResellerRoyaltiesResellerTypeGcp           ContractV2ResellerRoyaltiesResellerType = "GCP"
	ContractV2ResellerRoyaltiesResellerTypeGcpProService ContractV2ResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r ContractV2ResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case ContractV2ResellerRoyaltiesResellerTypeAws, ContractV2ResellerRoyaltiesResellerTypeAwsProService, ContractV2ResellerRoyaltiesResellerTypeGcp, ContractV2ResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

type ContractV2ResellerRoyaltiesSegment struct {
	Fraction              float64                                         `json:"fraction,required"`
	NetsuiteResellerID    string                                          `json:"netsuite_reseller_id,required"`
	ResellerType          ContractV2ResellerRoyaltiesSegmentsResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                       `json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  []string                                        `json:"applicable_product_ids"`
	ApplicableProductTags []string                                        `json:"applicable_product_tags"`
	AwsAccountNumber      string                                          `json:"aws_account_number"`
	AwsOfferID            string                                          `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                          `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                       `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                          `json:"gcp_account_id"`
	GcpOfferID            string                                          `json:"gcp_offer_id"`
	ResellerContractValue float64                                         `json:"reseller_contract_value"`
	JSON                  contractV2ResellerRoyaltiesSegmentJSON          `json:"-"`
}

// contractV2ResellerRoyaltiesSegmentJSON contains the JSON metadata for the struct
// [ContractV2ResellerRoyaltiesSegment]
type contractV2ResellerRoyaltiesSegmentJSON struct {
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

func (r *ContractV2ResellerRoyaltiesSegment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractV2ResellerRoyaltiesSegmentJSON) RawJSON() string {
	return r.raw
}

type ContractV2ResellerRoyaltiesSegmentsResellerType string

const (
	ContractV2ResellerRoyaltiesSegmentsResellerTypeAws           ContractV2ResellerRoyaltiesSegmentsResellerType = "AWS"
	ContractV2ResellerRoyaltiesSegmentsResellerTypeAwsProService ContractV2ResellerRoyaltiesSegmentsResellerType = "AWS_PRO_SERVICE"
	ContractV2ResellerRoyaltiesSegmentsResellerTypeGcp           ContractV2ResellerRoyaltiesSegmentsResellerType = "GCP"
	ContractV2ResellerRoyaltiesSegmentsResellerTypeGcpProService ContractV2ResellerRoyaltiesSegmentsResellerType = "GCP_PRO_SERVICE"
)

func (r ContractV2ResellerRoyaltiesSegmentsResellerType) IsKnown() bool {
	switch r {
	case ContractV2ResellerRoyaltiesSegmentsResellerTypeAws, ContractV2ResellerRoyaltiesSegmentsResellerTypeAwsProService, ContractV2ResellerRoyaltiesSegmentsResellerTypeGcp, ContractV2ResellerRoyaltiesSegmentsResellerTypeGcpProService:
		return true
	}
	return false
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type ContractV2ScheduledChargesOnUsageInvoices string

const (
	ContractV2ScheduledChargesOnUsageInvoicesAll ContractV2ScheduledChargesOnUsageInvoices = "ALL"
)

func (r ContractV2ScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case ContractV2ScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
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
	Credits                []Credit                                        `json:"credits"`
	// This field's availability is dependent on your client's configuration.
	Discounts    []Discount `json:"discounts"`
	EndingBefore time.Time  `json:"ending_before" format:"date-time"`
	// Either a **parent** configuration with a list of children or a **child**
	// configuration with a single parent.
	HierarchyConfiguration HierarchyConfiguration `json:"hierarchy_configuration"`
	Name                   string                 `json:"name"`
	NetPaymentTermsDays    float64                `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID                 string                               `json:"netsuite_sales_order_id"`
	PrepaidBalanceThresholdConfiguration PrepaidBalanceThresholdConfiguration `json:"prepaid_balance_threshold_configuration"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []ProService                               `json:"professional_services"`
	RateCardID           string                                     `json:"rate_card_id" format:"uuid"`
	RecurringCommits     []ContractWithoutAmendmentsRecurringCommit `json:"recurring_commits"`
	RecurringCredits     []ContractWithoutAmendmentsRecurringCredit `json:"recurring_credits"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []ContractWithoutAmendmentsResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices ContractWithoutAmendmentsScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	SpendThresholdConfiguration     SpendThresholdConfiguration                              `json:"spend_threshold_configuration"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue float64                              `json:"total_contract_value"`
	UsageFilter        ContractWithoutAmendmentsUsageFilter `json:"usage_filter"`
	JSON               contractWithoutAmendmentsJSON        `json:"-"`
}

// contractWithoutAmendmentsJSON contains the JSON metadata for the struct
// [ContractWithoutAmendments]
type contractWithoutAmendmentsJSON struct {
	Commits                              apijson.Field
	CreatedAt                            apijson.Field
	CreatedBy                            apijson.Field
	Overrides                            apijson.Field
	ScheduledCharges                     apijson.Field
	StartingAt                           apijson.Field
	Transitions                          apijson.Field
	UsageStatementSchedule               apijson.Field
	Credits                              apijson.Field
	Discounts                            apijson.Field
	EndingBefore                         apijson.Field
	HierarchyConfiguration               apijson.Field
	Name                                 apijson.Field
	NetPaymentTermsDays                  apijson.Field
	NetsuiteSalesOrderID                 apijson.Field
	PrepaidBalanceThresholdConfiguration apijson.Field
	ProfessionalServices                 apijson.Field
	RateCardID                           apijson.Field
	RecurringCommits                     apijson.Field
	RecurringCredits                     apijson.Field
	ResellerRoyalties                    apijson.Field
	SalesforceOpportunityID              apijson.Field
	ScheduledChargesOnUsageInvoices      apijson.Field
	SpendThresholdConfiguration          apijson.Field
	TotalContractValue                   apijson.Field
	UsageFilter                          apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
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
	// Contract usage statements follow a selected cadence based on this date.
	BillingAnchorDate time.Time                                                `json:"billing_anchor_date,required" format:"date-time"`
	Frequency         ContractWithoutAmendmentsUsageStatementScheduleFrequency `json:"frequency,required"`
	JSON              contractWithoutAmendmentsUsageStatementScheduleJSON      `json:"-"`
}

// contractWithoutAmendmentsUsageStatementScheduleJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsUsageStatementSchedule]
type contractWithoutAmendmentsUsageStatementScheduleJSON struct {
	BillingAnchorDate apijson.Field
	Frequency         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
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
	ContractWithoutAmendmentsUsageStatementScheduleFrequencyQuarterly ContractWithoutAmendmentsUsageStatementScheduleFrequency = "QUARTERLY"
	ContractWithoutAmendmentsUsageStatementScheduleFrequencyAnnual    ContractWithoutAmendmentsUsageStatementScheduleFrequency = "ANNUAL"
	ContractWithoutAmendmentsUsageStatementScheduleFrequencyWeekly    ContractWithoutAmendmentsUsageStatementScheduleFrequency = "WEEKLY"
)

func (r ContractWithoutAmendmentsUsageStatementScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsUsageStatementScheduleFrequencyMonthly, ContractWithoutAmendmentsUsageStatementScheduleFrequencyQuarterly, ContractWithoutAmendmentsUsageStatementScheduleFrequencyAnnual, ContractWithoutAmendmentsUsageStatementScheduleFrequencyWeekly:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCommit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount ContractWithoutAmendmentsRecurringCommitsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration ContractWithoutAmendmentsRecurringCommitsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                          `json:"priority,required"`
	Product  ContractWithoutAmendmentsRecurringCommitsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType ContractWithoutAmendmentsRecurringCommitsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                          `json:"applicable_product_tags"`
	Contract              ContractWithoutAmendmentsRecurringCommitsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount ContractWithoutAmendmentsRecurringCommitsInvoiceAmount `json:"invoice_amount"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration ContractWithoutAmendmentsRecurringCommitsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	RecurrenceFrequency ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig ContractWithoutAmendmentsRecurringCommitsSubscriptionConfig `json:"subscription_config"`
	JSON               contractWithoutAmendmentsRecurringCommitJSON                `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsRecurringCommit]
type contractWithoutAmendmentsRecurringCommitJSON struct {
	ID                     apijson.Field
	AccessAmount           apijson.Field
	CommitDuration         apijson.Field
	Priority               apijson.Field
	Product                apijson.Field
	RateType               apijson.Field
	StartingAt             apijson.Field
	ApplicableProductIDs   apijson.Field
	ApplicableProductTags  apijson.Field
	Contract               apijson.Field
	Description            apijson.Field
	EndingBefore           apijson.Field
	HierarchyConfiguration apijson.Field
	InvoiceAmount          apijson.Field
	Name                   apijson.Field
	NetsuiteSalesOrderID   apijson.Field
	Proration              apijson.Field
	RecurrenceFrequency    apijson.Field
	RolloverFraction       apijson.Field
	Specifiers             apijson.Field
	SubscriptionConfig     apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type ContractWithoutAmendmentsRecurringCommitsAccessAmount struct {
	CreditTypeID string                                                    `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64                                                   `json:"unit_price,required"`
	Quantity     float64                                                   `json:"quantity"`
	JSON         contractWithoutAmendmentsRecurringCommitsAccessAmountJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsAccessAmountJSON contains the JSON
// metadata for the struct [ContractWithoutAmendmentsRecurringCommitsAccessAmount]
type contractWithoutAmendmentsRecurringCommitsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	UnitPrice    apijson.Field
	Quantity     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type ContractWithoutAmendmentsRecurringCommitsCommitDuration struct {
	Value float64                                                     `json:"value,required"`
	Unit  ContractWithoutAmendmentsRecurringCommitsCommitDurationUnit `json:"unit"`
	JSON  contractWithoutAmendmentsRecurringCommitsCommitDurationJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsCommitDurationJSON contains the JSON
// metadata for the struct
// [ContractWithoutAmendmentsRecurringCommitsCommitDuration]
type contractWithoutAmendmentsRecurringCommitsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsRecurringCommitsCommitDurationUnit string

const (
	ContractWithoutAmendmentsRecurringCommitsCommitDurationUnitPeriods ContractWithoutAmendmentsRecurringCommitsCommitDurationUnit = "PERIODS"
)

func (r ContractWithoutAmendmentsRecurringCommitsCommitDurationUnit) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCommitsProduct struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	Name string                                               `json:"name,required"`
	JSON contractWithoutAmendmentsRecurringCommitsProductJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsProductJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsRecurringCommitsProduct]
type contractWithoutAmendmentsRecurringCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type ContractWithoutAmendmentsRecurringCommitsRateType string

const (
	ContractWithoutAmendmentsRecurringCommitsRateTypeCommitRate ContractWithoutAmendmentsRecurringCommitsRateType = "COMMIT_RATE"
	ContractWithoutAmendmentsRecurringCommitsRateTypeListRate   ContractWithoutAmendmentsRecurringCommitsRateType = "LIST_RATE"
)

func (r ContractWithoutAmendmentsRecurringCommitsRateType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsRateTypeCommitRate, ContractWithoutAmendmentsRecurringCommitsRateTypeListRate:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCommitsContract struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	JSON contractWithoutAmendmentsRecurringCommitsContractJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsContractJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsRecurringCommitsContract]
type contractWithoutAmendmentsRecurringCommitsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsContractJSON) RawJSON() string {
	return r.raw
}

// The amount the customer should be billed for the commit. Not required.
type ContractWithoutAmendmentsRecurringCommitsInvoiceAmount struct {
	CreditTypeID string                                                     `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                    `json:"quantity,required"`
	UnitPrice    float64                                                    `json:"unit_price,required"`
	JSON         contractWithoutAmendmentsRecurringCommitsInvoiceAmountJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsInvoiceAmountJSON contains the JSON
// metadata for the struct [ContractWithoutAmendmentsRecurringCommitsInvoiceAmount]
type contractWithoutAmendmentsRecurringCommitsInvoiceAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsInvoiceAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsInvoiceAmountJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type ContractWithoutAmendmentsRecurringCommitsProration string

const (
	ContractWithoutAmendmentsRecurringCommitsProrationNone         ContractWithoutAmendmentsRecurringCommitsProration = "NONE"
	ContractWithoutAmendmentsRecurringCommitsProrationFirst        ContractWithoutAmendmentsRecurringCommitsProration = "FIRST"
	ContractWithoutAmendmentsRecurringCommitsProrationLast         ContractWithoutAmendmentsRecurringCommitsProration = "LAST"
	ContractWithoutAmendmentsRecurringCommitsProrationFirstAndLast ContractWithoutAmendmentsRecurringCommitsProration = "FIRST_AND_LAST"
)

func (r ContractWithoutAmendmentsRecurringCommitsProration) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsProrationNone, ContractWithoutAmendmentsRecurringCommitsProrationFirst, ContractWithoutAmendmentsRecurringCommitsProrationLast, ContractWithoutAmendmentsRecurringCommitsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's starting_at rather than the usage
// invoice dates.
type ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency string

const (
	ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyMonthly   ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency = "MONTHLY"
	ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyQuarterly ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency = "QUARTERLY"
	ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyAnnual    ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency = "ANNUAL"
	ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyWeekly    ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency = "WEEKLY"
)

func (r ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyMonthly, ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyQuarterly, ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyAnnual, ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

// Attach a subscription to the recurring commit/credit.
type ContractWithoutAmendmentsRecurringCommitsSubscriptionConfig struct {
	Allocation              ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocation              `json:"allocation,required"`
	ApplySeatIncreaseConfig ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,required"`
	SubscriptionID          string                                                                             `json:"subscription_id,required" format:"uuid"`
	JSON                    contractWithoutAmendmentsRecurringCommitsSubscriptionConfigJSON                    `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsSubscriptionConfigJSON contains the
// JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCommitsSubscriptionConfig]
type contractWithoutAmendmentsRecurringCommitsSubscriptionConfigJSON struct {
	Allocation              apijson.Field
	ApplySeatIncreaseConfig apijson.Field
	SubscriptionID          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsSubscriptionConfigJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocation string

const (
	ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocationIndividual ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocation = "INDIVIDUAL"
	ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocationPooled     ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocation = "POOLED"
)

func (r ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocation) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocationIndividual, ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigAllocationPooled:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool                                                                                   `json:"is_prorated,required"`
	JSON       contractWithoutAmendmentsRecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig]
type contractWithoutAmendmentsRecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON struct {
	IsProrated  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsRecurringCredit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount ContractWithoutAmendmentsRecurringCreditsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration ContractWithoutAmendmentsRecurringCreditsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                          `json:"priority,required"`
	Product  ContractWithoutAmendmentsRecurringCreditsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType ContractWithoutAmendmentsRecurringCreditsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                          `json:"applicable_product_tags"`
	Contract              ContractWithoutAmendmentsRecurringCreditsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration ContractWithoutAmendmentsRecurringCreditsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	RecurrenceFrequency ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig ContractWithoutAmendmentsRecurringCreditsSubscriptionConfig `json:"subscription_config"`
	JSON               contractWithoutAmendmentsRecurringCreditJSON                `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsRecurringCredit]
type contractWithoutAmendmentsRecurringCreditJSON struct {
	ID                     apijson.Field
	AccessAmount           apijson.Field
	CommitDuration         apijson.Field
	Priority               apijson.Field
	Product                apijson.Field
	RateType               apijson.Field
	StartingAt             apijson.Field
	ApplicableProductIDs   apijson.Field
	ApplicableProductTags  apijson.Field
	Contract               apijson.Field
	Description            apijson.Field
	EndingBefore           apijson.Field
	HierarchyConfiguration apijson.Field
	Name                   apijson.Field
	NetsuiteSalesOrderID   apijson.Field
	Proration              apijson.Field
	RecurrenceFrequency    apijson.Field
	RolloverFraction       apijson.Field
	Specifiers             apijson.Field
	SubscriptionConfig     apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type ContractWithoutAmendmentsRecurringCreditsAccessAmount struct {
	CreditTypeID string                                                    `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64                                                   `json:"unit_price,required"`
	Quantity     float64                                                   `json:"quantity"`
	JSON         contractWithoutAmendmentsRecurringCreditsAccessAmountJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsAccessAmountJSON contains the JSON
// metadata for the struct [ContractWithoutAmendmentsRecurringCreditsAccessAmount]
type contractWithoutAmendmentsRecurringCreditsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	UnitPrice    apijson.Field
	Quantity     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type ContractWithoutAmendmentsRecurringCreditsCommitDuration struct {
	Value float64                                                     `json:"value,required"`
	Unit  ContractWithoutAmendmentsRecurringCreditsCommitDurationUnit `json:"unit"`
	JSON  contractWithoutAmendmentsRecurringCreditsCommitDurationJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsCommitDurationJSON contains the JSON
// metadata for the struct
// [ContractWithoutAmendmentsRecurringCreditsCommitDuration]
type contractWithoutAmendmentsRecurringCreditsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsRecurringCreditsCommitDurationUnit string

const (
	ContractWithoutAmendmentsRecurringCreditsCommitDurationUnitPeriods ContractWithoutAmendmentsRecurringCreditsCommitDurationUnit = "PERIODS"
)

func (r ContractWithoutAmendmentsRecurringCreditsCommitDurationUnit) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCreditsProduct struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	Name string                                               `json:"name,required"`
	JSON contractWithoutAmendmentsRecurringCreditsProductJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsProductJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsRecurringCreditsProduct]
type contractWithoutAmendmentsRecurringCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type ContractWithoutAmendmentsRecurringCreditsRateType string

const (
	ContractWithoutAmendmentsRecurringCreditsRateTypeCommitRate ContractWithoutAmendmentsRecurringCreditsRateType = "COMMIT_RATE"
	ContractWithoutAmendmentsRecurringCreditsRateTypeListRate   ContractWithoutAmendmentsRecurringCreditsRateType = "LIST_RATE"
)

func (r ContractWithoutAmendmentsRecurringCreditsRateType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsRateTypeCommitRate, ContractWithoutAmendmentsRecurringCreditsRateTypeListRate:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCreditsContract struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	JSON contractWithoutAmendmentsRecurringCreditsContractJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsContractJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsRecurringCreditsContract]
type contractWithoutAmendmentsRecurringCreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsContractJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type ContractWithoutAmendmentsRecurringCreditsProration string

const (
	ContractWithoutAmendmentsRecurringCreditsProrationNone         ContractWithoutAmendmentsRecurringCreditsProration = "NONE"
	ContractWithoutAmendmentsRecurringCreditsProrationFirst        ContractWithoutAmendmentsRecurringCreditsProration = "FIRST"
	ContractWithoutAmendmentsRecurringCreditsProrationLast         ContractWithoutAmendmentsRecurringCreditsProration = "LAST"
	ContractWithoutAmendmentsRecurringCreditsProrationFirstAndLast ContractWithoutAmendmentsRecurringCreditsProration = "FIRST_AND_LAST"
)

func (r ContractWithoutAmendmentsRecurringCreditsProration) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsProrationNone, ContractWithoutAmendmentsRecurringCreditsProrationFirst, ContractWithoutAmendmentsRecurringCreditsProrationLast, ContractWithoutAmendmentsRecurringCreditsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's starting_at rather than the usage
// invoice dates.
type ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency string

const (
	ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyMonthly   ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency = "MONTHLY"
	ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyQuarterly ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency = "QUARTERLY"
	ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyAnnual    ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency = "ANNUAL"
	ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyWeekly    ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency = "WEEKLY"
)

func (r ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyMonthly, ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyQuarterly, ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyAnnual, ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

// Attach a subscription to the recurring commit/credit.
type ContractWithoutAmendmentsRecurringCreditsSubscriptionConfig struct {
	Allocation              ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocation              `json:"allocation,required"`
	ApplySeatIncreaseConfig ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,required"`
	SubscriptionID          string                                                                             `json:"subscription_id,required" format:"uuid"`
	JSON                    contractWithoutAmendmentsRecurringCreditsSubscriptionConfigJSON                    `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsSubscriptionConfigJSON contains the
// JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCreditsSubscriptionConfig]
type contractWithoutAmendmentsRecurringCreditsSubscriptionConfigJSON struct {
	Allocation              apijson.Field
	ApplySeatIncreaseConfig apijson.Field
	SubscriptionID          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsSubscriptionConfigJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocation string

const (
	ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocationIndividual ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocation = "INDIVIDUAL"
	ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocationPooled     ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocation = "POOLED"
)

func (r ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocation) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocationIndividual, ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigAllocationPooled:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool                                                                                   `json:"is_prorated,required"`
	JSON       contractWithoutAmendmentsRecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig]
type contractWithoutAmendmentsRecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON struct {
	IsProrated  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON) RawJSON() string {
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

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type ContractWithoutAmendmentsScheduledChargesOnUsageInvoices string

const (
	ContractWithoutAmendmentsScheduledChargesOnUsageInvoicesAll ContractWithoutAmendmentsScheduledChargesOnUsageInvoices = "ALL"
)

func (r ContractWithoutAmendmentsScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
}

type ContractWithoutAmendmentsUsageFilter struct {
	Current BaseUsageFilter                              `json:"current,required,nullable"`
	Initial BaseUsageFilter                              `json:"initial,required"`
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

type Credit struct {
	ID      string        `json:"id,required" format:"uuid"`
	Product CreditProduct `json:"product,required"`
	Type    CreditType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        ScheduleDuration `json:"access_schedule"`
	ApplicableContractIDs []string         `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string         `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string         `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64           `json:"balance"`
	Contract     CreditContract    `json:"contract"`
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []CreditLedger `json:"ledger"`
	Name   string         `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64        `json:"priority"`
	RateType CreditRateType `json:"rate_type"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier `json:"specifiers"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey string     `json:"uniqueness_key"`
	JSON          creditJSON `json:"-"`
}

// creditJSON contains the JSON metadata for the struct [Credit]
type creditJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Balance                 apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	HierarchyConfiguration  apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	RateType                apijson.Field
	SalesforceOpportunityID apijson.Field
	Specifiers              apijson.Field
	UniquenessKey           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *Credit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditJSON) RawJSON() string {
	return r.raw
}

func (r Credit) ImplementsV1ContractListBalancesResponse() {}

type CreditProduct struct {
	ID   string            `json:"id,required" format:"uuid"`
	Name string            `json:"name,required"`
	JSON creditProductJSON `json:"-"`
}

// creditProductJSON contains the JSON metadata for the struct [CreditProduct]
type creditProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditProductJSON) RawJSON() string {
	return r.raw
}

type CreditType string

const (
	CreditTypeCredit CreditType = "CREDIT"
)

func (r CreditType) IsKnown() bool {
	switch r {
	case CreditTypeCredit:
		return true
	}
	return false
}

type CreditContract struct {
	ID   string             `json:"id,required" format:"uuid"`
	JSON creditContractJSON `json:"-"`
}

// creditContractJSON contains the JSON metadata for the struct [CreditContract]
type creditContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditContractJSON) RawJSON() string {
	return r.raw
}

type CreditLedger struct {
	Amount     float64          `json:"amount,required"`
	Timestamp  time.Time        `json:"timestamp,required" format:"date-time"`
	Type       CreditLedgerType `json:"type,required"`
	ContractID string           `json:"contract_id" format:"uuid"`
	InvoiceID  string           `json:"invoice_id" format:"uuid"`
	Reason     string           `json:"reason"`
	SegmentID  string           `json:"segment_id" format:"uuid"`
	JSON       creditLedgerJSON `json:"-"`
	union      CreditLedgerUnion
}

// creditLedgerJSON contains the JSON metadata for the struct [CreditLedger]
type creditLedgerJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	InvoiceID   apijson.Field
	Reason      apijson.Field
	SegmentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r creditLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *CreditLedger) UnmarshalJSON(data []byte) (err error) {
	*r = CreditLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CreditLedgerUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [CreditLedgerCreditSegmentStartLedgerEntry],
// [CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [CreditLedgerCreditExpirationLedgerEntry],
// [CreditLedgerCreditCanceledLedgerEntry],
// [CreditLedgerCreditCreditedLedgerEntry], [CreditLedgerCreditManualLedgerEntry],
// [CreditLedgerCreditSeatBasedAdjustmentLedgerEntry].
func (r CreditLedger) AsUnion() CreditLedgerUnion {
	return r.union
}

// Union satisfied by [CreditLedgerCreditSegmentStartLedgerEntry],
// [CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [CreditLedgerCreditExpirationLedgerEntry],
// [CreditLedgerCreditCanceledLedgerEntry],
// [CreditLedgerCreditCreditedLedgerEntry], [CreditLedgerCreditManualLedgerEntry]
// or [CreditLedgerCreditSeatBasedAdjustmentLedgerEntry].
type CreditLedgerUnion interface {
	implementsCreditLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CreditLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditSeatBasedAdjustmentLedgerEntry{}),
		},
	)
}

type CreditLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64                                       `json:"amount,required"`
	SegmentID string                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                     `json:"timestamp,required" format:"date-time"`
	Type      CreditLedgerCreditSegmentStartLedgerEntryType `json:"type,required"`
	JSON      creditLedgerCreditSegmentStartLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditSegmentStartLedgerEntryJSON contains the JSON metadata for the
// struct [CreditLedgerCreditSegmentStartLedgerEntry]
type creditLedgerCreditSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditSegmentStartLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditSegmentStartLedgerEntryType string

const (
	CreditLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart CreditLedgerCreditSegmentStartLedgerEntryType = "CREDIT_SEGMENT_START"
)

func (r CreditLedgerCreditSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart:
		return true
	}
	return false
}

type CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount     float64                                                    `json:"amount,required"`
	InvoiceID  string                                                     `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                                  `json:"timestamp,required" format:"date-time"`
	Type       CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	ContractID string                                                     `json:"contract_id" format:"uuid"`
	JSON       creditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON contains the JSON
// metadata for the struct [CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry]
type creditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType string

const (
	CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type CreditLedgerCreditExpirationLedgerEntry struct {
	Amount    float64                                     `json:"amount,required"`
	SegmentID string                                      `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                   `json:"timestamp,required" format:"date-time"`
	Type      CreditLedgerCreditExpirationLedgerEntryType `json:"type,required"`
	JSON      creditLedgerCreditExpirationLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditExpirationLedgerEntryJSON contains the JSON metadata for the
// struct [CreditLedgerCreditExpirationLedgerEntry]
type creditLedgerCreditExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditExpirationLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditExpirationLedgerEntryType string

const (
	CreditLedgerCreditExpirationLedgerEntryTypeCreditExpiration CreditLedgerCreditExpirationLedgerEntryType = "CREDIT_EXPIRATION"
)

func (r CreditLedgerCreditExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditExpirationLedgerEntryTypeCreditExpiration:
		return true
	}
	return false
}

type CreditLedgerCreditCanceledLedgerEntry struct {
	Amount     float64                                   `json:"amount,required"`
	InvoiceID  string                                    `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                    `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                 `json:"timestamp,required" format:"date-time"`
	Type       CreditLedgerCreditCanceledLedgerEntryType `json:"type,required"`
	ContractID string                                    `json:"contract_id" format:"uuid"`
	JSON       creditLedgerCreditCanceledLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditCanceledLedgerEntryJSON contains the JSON metadata for the
// struct [CreditLedgerCreditCanceledLedgerEntry]
type creditLedgerCreditCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditCanceledLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditCanceledLedgerEntryType string

const (
	CreditLedgerCreditCanceledLedgerEntryTypeCreditCanceled CreditLedgerCreditCanceledLedgerEntryType = "CREDIT_CANCELED"
)

func (r CreditLedgerCreditCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditCanceledLedgerEntryTypeCreditCanceled:
		return true
	}
	return false
}

type CreditLedgerCreditCreditedLedgerEntry struct {
	Amount     float64                                   `json:"amount,required"`
	InvoiceID  string                                    `json:"invoice_id,required" format:"uuid"`
	SegmentID  string                                    `json:"segment_id,required" format:"uuid"`
	Timestamp  time.Time                                 `json:"timestamp,required" format:"date-time"`
	Type       CreditLedgerCreditCreditedLedgerEntryType `json:"type,required"`
	ContractID string                                    `json:"contract_id" format:"uuid"`
	JSON       creditLedgerCreditCreditedLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditCreditedLedgerEntryJSON contains the JSON metadata for the
// struct [CreditLedgerCreditCreditedLedgerEntry]
type creditLedgerCreditCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditCreditedLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditCreditedLedgerEntryType string

const (
	CreditLedgerCreditCreditedLedgerEntryTypeCreditCredited CreditLedgerCreditCreditedLedgerEntryType = "CREDIT_CREDITED"
)

func (r CreditLedgerCreditCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditCreditedLedgerEntryTypeCreditCredited:
		return true
	}
	return false
}

type CreditLedgerCreditManualLedgerEntry struct {
	Amount    float64                                 `json:"amount,required"`
	Reason    string                                  `json:"reason,required"`
	Timestamp time.Time                               `json:"timestamp,required" format:"date-time"`
	Type      CreditLedgerCreditManualLedgerEntryType `json:"type,required"`
	JSON      creditLedgerCreditManualLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditManualLedgerEntryJSON contains the JSON metadata for the
// struct [CreditLedgerCreditManualLedgerEntry]
type creditLedgerCreditManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditManualLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditManualLedgerEntryType string

const (
	CreditLedgerCreditManualLedgerEntryTypeCreditManual CreditLedgerCreditManualLedgerEntryType = "CREDIT_MANUAL"
)

func (r CreditLedgerCreditManualLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditManualLedgerEntryTypeCreditManual:
		return true
	}
	return false
}

type CreditLedgerCreditSeatBasedAdjustmentLedgerEntry struct {
	Amount    float64                                              `json:"amount,required"`
	SegmentID string                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                            `json:"timestamp,required" format:"date-time"`
	Type      CreditLedgerCreditSeatBasedAdjustmentLedgerEntryType `json:"type,required"`
	JSON      creditLedgerCreditSeatBasedAdjustmentLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditSeatBasedAdjustmentLedgerEntryJSON contains the JSON metadata
// for the struct [CreditLedgerCreditSeatBasedAdjustmentLedgerEntry]
type creditLedgerCreditSeatBasedAdjustmentLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditSeatBasedAdjustmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditSeatBasedAdjustmentLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditSeatBasedAdjustmentLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditSeatBasedAdjustmentLedgerEntryType string

const (
	CreditLedgerCreditSeatBasedAdjustmentLedgerEntryTypeCreditSeatBasedAdjustment CreditLedgerCreditSeatBasedAdjustmentLedgerEntryType = "CREDIT_SEAT_BASED_ADJUSTMENT"
)

func (r CreditLedgerCreditSeatBasedAdjustmentLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditSeatBasedAdjustmentLedgerEntryTypeCreditSeatBasedAdjustment:
		return true
	}
	return false
}

type CreditLedgerType string

const (
	CreditLedgerTypeCreditSegmentStart              CreditLedgerType = "CREDIT_SEGMENT_START"
	CreditLedgerTypeCreditAutomatedInvoiceDeduction CreditLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	CreditLedgerTypeCreditExpiration                CreditLedgerType = "CREDIT_EXPIRATION"
	CreditLedgerTypeCreditCanceled                  CreditLedgerType = "CREDIT_CANCELED"
	CreditLedgerTypeCreditCredited                  CreditLedgerType = "CREDIT_CREDITED"
	CreditLedgerTypeCreditManual                    CreditLedgerType = "CREDIT_MANUAL"
	CreditLedgerTypeCreditSeatBasedAdjustment       CreditLedgerType = "CREDIT_SEAT_BASED_ADJUSTMENT"
)

func (r CreditLedgerType) IsKnown() bool {
	switch r {
	case CreditLedgerTypeCreditSegmentStart, CreditLedgerTypeCreditAutomatedInvoiceDeduction, CreditLedgerTypeCreditExpiration, CreditLedgerTypeCreditCanceled, CreditLedgerTypeCreditCredited, CreditLedgerTypeCreditManual, CreditLedgerTypeCreditSeatBasedAdjustment:
		return true
	}
	return false
}

type CreditRateType string

const (
	CreditRateTypeCommitRate CreditRateType = "COMMIT_RATE"
	CreditRateTypeListRate   CreditRateType = "LIST_RATE"
)

func (r CreditRateType) IsKnown() bool {
	switch r {
	case CreditRateTypeCommitRate, CreditRateTypeListRate:
		return true
	}
	return false
}

type CreditTypeData struct {
	ID   string             `json:"id,required" format:"uuid"`
	Name string             `json:"name,required"`
	JSON creditTypeDataJSON `json:"-"`
}

// creditTypeDataJSON contains the JSON metadata for the struct [CreditTypeData]
type creditTypeDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditTypeData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditTypeDataJSON) RawJSON() string {
	return r.raw
}

type Discount struct {
	ID           string              `json:"id,required" format:"uuid"`
	Product      DiscountProduct     `json:"product,required"`
	Schedule     SchedulePointInTime `json:"schedule,required"`
	CustomFields map[string]string   `json:"custom_fields"`
	Name         string              `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string       `json:"netsuite_sales_order_id"`
	JSON                 discountJSON `json:"-"`
}

// discountJSON contains the JSON metadata for the struct [Discount]
type discountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	CustomFields         apijson.Field
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

// An optional filtering rule to match the 'event_type' property of an event.
type EventTypeFilter struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues []string `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues []string            `json:"not_in_values"`
	JSON        eventTypeFilterJSON `json:"-"`
}

// eventTypeFilterJSON contains the JSON metadata for the struct [EventTypeFilter]
type eventTypeFilterJSON struct {
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventTypeFilterJSON) RawJSON() string {
	return r.raw
}

// An optional filtering rule to match the 'event_type' property of an event.
type EventTypeFilterParam struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues param.Field[[]string] `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues param.Field[[]string] `json:"not_in_values"`
}

func (r EventTypeFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either a **parent** configuration with a list of children or a **child**
// configuration with a single parent.
type HierarchyConfiguration struct {
	// This field can have the runtime type of
	// [[]HierarchyConfigurationParentHierarchyConfigurationChild].
	Children interface{} `json:"children"`
	// This field can have the runtime type of
	// [HierarchyConfigurationChildHierarchyConfigurationParent].
	Parent interface{}                `json:"parent"`
	JSON   hierarchyConfigurationJSON `json:"-"`
	union  HierarchyConfigurationUnion
}

// hierarchyConfigurationJSON contains the JSON metadata for the struct
// [HierarchyConfiguration]
type hierarchyConfigurationJSON struct {
	Children    apijson.Field
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r hierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r *HierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	*r = HierarchyConfiguration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [HierarchyConfigurationUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [HierarchyConfigurationParentHierarchyConfiguration],
// [HierarchyConfigurationChildHierarchyConfiguration].
func (r HierarchyConfiguration) AsUnion() HierarchyConfigurationUnion {
	return r.union
}

// Either a **parent** configuration with a list of children or a **child**
// configuration with a single parent.
//
// Union satisfied by [HierarchyConfigurationParentHierarchyConfiguration] or
// [HierarchyConfigurationChildHierarchyConfiguration].
type HierarchyConfigurationUnion interface {
	implementsHierarchyConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HierarchyConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HierarchyConfigurationParentHierarchyConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HierarchyConfigurationChildHierarchyConfiguration{}),
		},
	)
}

type HierarchyConfigurationParentHierarchyConfiguration struct {
	// List of contracts that belong to this parent.
	Children []HierarchyConfigurationParentHierarchyConfigurationChild `json:"children,required"`
	JSON     hierarchyConfigurationParentHierarchyConfigurationJSON    `json:"-"`
}

// hierarchyConfigurationParentHierarchyConfigurationJSON contains the JSON
// metadata for the struct [HierarchyConfigurationParentHierarchyConfiguration]
type hierarchyConfigurationParentHierarchyConfigurationJSON struct {
	Children    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HierarchyConfigurationParentHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hierarchyConfigurationParentHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r HierarchyConfigurationParentHierarchyConfiguration) implementsHierarchyConfiguration() {}

type HierarchyConfigurationParentHierarchyConfigurationChild struct {
	ContractID string                                                      `json:"contract_id,required" format:"uuid"`
	CustomerID string                                                      `json:"customer_id,required" format:"uuid"`
	JSON       hierarchyConfigurationParentHierarchyConfigurationChildJSON `json:"-"`
}

// hierarchyConfigurationParentHierarchyConfigurationChildJSON contains the JSON
// metadata for the struct
// [HierarchyConfigurationParentHierarchyConfigurationChild]
type hierarchyConfigurationParentHierarchyConfigurationChildJSON struct {
	ContractID  apijson.Field
	CustomerID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HierarchyConfigurationParentHierarchyConfigurationChild) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hierarchyConfigurationParentHierarchyConfigurationChildJSON) RawJSON() string {
	return r.raw
}

type HierarchyConfigurationChildHierarchyConfiguration struct {
	// The single parent contract/customer for this child.
	Parent HierarchyConfigurationChildHierarchyConfigurationParent `json:"parent,required"`
	JSON   hierarchyConfigurationChildHierarchyConfigurationJSON   `json:"-"`
}

// hierarchyConfigurationChildHierarchyConfigurationJSON contains the JSON metadata
// for the struct [HierarchyConfigurationChildHierarchyConfiguration]
type hierarchyConfigurationChildHierarchyConfigurationJSON struct {
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HierarchyConfigurationChildHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hierarchyConfigurationChildHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r HierarchyConfigurationChildHierarchyConfiguration) implementsHierarchyConfiguration() {}

// The single parent contract/customer for this child.
type HierarchyConfigurationChildHierarchyConfigurationParent struct {
	ContractID string                                                      `json:"contract_id,required" format:"uuid"`
	CustomerID string                                                      `json:"customer_id,required" format:"uuid"`
	JSON       hierarchyConfigurationChildHierarchyConfigurationParentJSON `json:"-"`
}

// hierarchyConfigurationChildHierarchyConfigurationParentJSON contains the JSON
// metadata for the struct
// [HierarchyConfigurationChildHierarchyConfigurationParent]
type hierarchyConfigurationChildHierarchyConfigurationParentJSON struct {
	ContractID  apijson.Field
	CustomerID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HierarchyConfigurationChildHierarchyConfigurationParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hierarchyConfigurationChildHierarchyConfigurationParentJSON) RawJSON() string {
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
	ID                    string         `json:"id,required" format:"uuid"`
	StartingAt            time.Time      `json:"starting_at,required" format:"date-time"`
	ApplicableProductTags []string       `json:"applicable_product_tags"`
	CreditType            CreditTypeData `json:"credit_type"`
	EndingBefore          time.Time      `json:"ending_before" format:"date-time"`
	Entitled              bool           `json:"entitled"`
	IsCommitSpecific      bool           `json:"is_commit_specific"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated         bool                        `json:"is_prorated"`
	Multiplier         float64                     `json:"multiplier"`
	OverrideSpecifiers []OverrideOverrideSpecifier `json:"override_specifiers"`
	OverrideTiers      []OverrideTier              `json:"override_tiers"`
	OverwriteRate      OverrideOverwriteRate       `json:"overwrite_rate"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64         `json:"price"`
	Priority float64         `json:"priority"`
	Product  OverrideProduct `json:"product"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64          `json:"quantity"`
	RateType OverrideRateType `json:"rate_type"`
	Target   OverrideTarget   `json:"target"`
	// Only set for TIERED rate_type.
	Tiers []Tier       `json:"tiers"`
	Type  OverrideType `json:"type"`
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
	IsCommitSpecific      apijson.Field
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
	Target                apijson.Field
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
	BillingFrequency        OverrideOverrideSpecifiersBillingFrequency `json:"billing_frequency"`
	CommitIDs               []string                                   `json:"commit_ids"`
	PresentationGroupValues map[string]string                          `json:"presentation_group_values"`
	PricingGroupValues      map[string]string                          `json:"pricing_group_values"`
	ProductID               string                                     `json:"product_id" format:"uuid"`
	ProductTags             []string                                   `json:"product_tags"`
	RecurringCommitIDs      []string                                   `json:"recurring_commit_ids"`
	RecurringCreditIDs      []string                                   `json:"recurring_credit_ids"`
	JSON                    overrideOverrideSpecifierJSON              `json:"-"`
}

// overrideOverrideSpecifierJSON contains the JSON metadata for the struct
// [OverrideOverrideSpecifier]
type overrideOverrideSpecifierJSON struct {
	BillingFrequency        apijson.Field
	CommitIDs               apijson.Field
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	RecurringCommitIDs      apijson.Field
	RecurringCreditIDs      apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OverrideOverrideSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideOverrideSpecifierJSON) RawJSON() string {
	return r.raw
}

type OverrideOverrideSpecifiersBillingFrequency string

const (
	OverrideOverrideSpecifiersBillingFrequencyMonthly   OverrideOverrideSpecifiersBillingFrequency = "MONTHLY"
	OverrideOverrideSpecifiersBillingFrequencyQuarterly OverrideOverrideSpecifiersBillingFrequency = "QUARTERLY"
	OverrideOverrideSpecifiersBillingFrequencyAnnual    OverrideOverrideSpecifiersBillingFrequency = "ANNUAL"
	OverrideOverrideSpecifiersBillingFrequencyWeekly    OverrideOverrideSpecifiersBillingFrequency = "WEEKLY"
)

func (r OverrideOverrideSpecifiersBillingFrequency) IsKnown() bool {
	switch r {
	case OverrideOverrideSpecifiersBillingFrequencyMonthly, OverrideOverrideSpecifiersBillingFrequencyQuarterly, OverrideOverrideSpecifiersBillingFrequencyAnnual, OverrideOverrideSpecifiersBillingFrequencyWeekly:
		return true
	}
	return false
}

type OverrideOverwriteRate struct {
	RateType   OverrideOverwriteRateRateType `json:"rate_type,required"`
	CreditType CreditTypeData                `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]interface{} `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []Tier                    `json:"tiers"`
	JSON  overrideOverwriteRateJSON `json:"-"`
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
	OverrideOverwriteRateRateTypePercentage   OverrideOverwriteRateRateType = "PERCENTAGE"
	OverrideOverwriteRateRateTypeSubscription OverrideOverwriteRateRateType = "SUBSCRIPTION"
	OverrideOverwriteRateRateTypeTiered       OverrideOverwriteRateRateType = "TIERED"
	OverrideOverwriteRateRateTypeCustom       OverrideOverwriteRateRateType = "CUSTOM"
)

func (r OverrideOverwriteRateRateType) IsKnown() bool {
	switch r {
	case OverrideOverwriteRateRateTypeFlat, OverrideOverwriteRateRateTypePercentage, OverrideOverwriteRateRateTypeSubscription, OverrideOverwriteRateRateTypeTiered, OverrideOverwriteRateRateTypeCustom:
		return true
	}
	return false
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
	OverrideRateTypePercentage   OverrideRateType = "PERCENTAGE"
	OverrideRateTypeSubscription OverrideRateType = "SUBSCRIPTION"
	OverrideRateTypeTiered       OverrideRateType = "TIERED"
	OverrideRateTypeCustom       OverrideRateType = "CUSTOM"
)

func (r OverrideRateType) IsKnown() bool {
	switch r {
	case OverrideRateTypeFlat, OverrideRateTypePercentage, OverrideRateTypeSubscription, OverrideRateTypeTiered, OverrideRateTypeCustom:
		return true
	}
	return false
}

type OverrideTarget string

const (
	OverrideTargetCommitRate OverrideTarget = "COMMIT_RATE"
	OverrideTargetListRate   OverrideTarget = "LIST_RATE"
)

func (r OverrideTarget) IsKnown() bool {
	switch r {
	case OverrideTargetCommitRate, OverrideTargetListRate:
		return true
	}
	return false
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

type OverrideTier struct {
	Multiplier float64          `json:"multiplier,required"`
	Size       float64          `json:"size"`
	JSON       overrideTierJSON `json:"-"`
}

// overrideTierJSON contains the JSON metadata for the struct [OverrideTier]
type overrideTierJSON struct {
	Multiplier  apijson.Field
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

type PaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType PaymentGateConfigPaymentGateType `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig PaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig PaymentGateConfigStripeConfig `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType PaymentGateConfigTaxType `json:"tax_type"`
	JSON    paymentGateConfigJSON    `json:"-"`
}

// paymentGateConfigJSON contains the JSON metadata for the struct
// [PaymentGateConfig]
type paymentGateConfigJSON struct {
	PaymentGateType        apijson.Field
	PrecalculatedTaxConfig apijson.Field
	StripeConfig           apijson.Field
	TaxType                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PaymentGateConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentGateConfigJSON) RawJSON() string {
	return r.raw
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type PaymentGateConfigPaymentGateType string

const (
	PaymentGateConfigPaymentGateTypeNone     PaymentGateConfigPaymentGateType = "NONE"
	PaymentGateConfigPaymentGateTypeStripe   PaymentGateConfigPaymentGateType = "STRIPE"
	PaymentGateConfigPaymentGateTypeExternal PaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r PaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case PaymentGateConfigPaymentGateTypeNone, PaymentGateConfigPaymentGateTypeStripe, PaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type PaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName string                                      `json:"tax_name"`
	JSON    paymentGateConfigPrecalculatedTaxConfigJSON `json:"-"`
}

// paymentGateConfigPrecalculatedTaxConfigJSON contains the JSON metadata for the
// struct [PaymentGateConfigPrecalculatedTaxConfig]
type paymentGateConfigPrecalculatedTaxConfigJSON struct {
	TaxAmount   apijson.Field
	TaxName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentGateConfigPrecalculatedTaxConfigJSON) RawJSON() string {
	return r.raw
}

// Only applicable if using STRIPE as your payment gate type.
type PaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType PaymentGateConfigStripeConfigPaymentType `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string                 `json:"invoice_metadata"`
	JSON            paymentGateConfigStripeConfigJSON `json:"-"`
}

// paymentGateConfigStripeConfigJSON contains the JSON metadata for the struct
// [PaymentGateConfigStripeConfig]
type paymentGateConfigStripeConfigJSON struct {
	PaymentType     apijson.Field
	InvoiceMetadata apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentGateConfigStripeConfigJSON) RawJSON() string {
	return r.raw
}

// If left blank, will default to INVOICE
type PaymentGateConfigStripeConfigPaymentType string

const (
	PaymentGateConfigStripeConfigPaymentTypeInvoice       PaymentGateConfigStripeConfigPaymentType = "INVOICE"
	PaymentGateConfigStripeConfigPaymentTypePaymentIntent PaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r PaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case PaymentGateConfigStripeConfigPaymentTypeInvoice, PaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type PaymentGateConfigTaxType string

const (
	PaymentGateConfigTaxTypeNone          PaymentGateConfigTaxType = "NONE"
	PaymentGateConfigTaxTypeStripe        PaymentGateConfigTaxType = "STRIPE"
	PaymentGateConfigTaxTypeAnrok         PaymentGateConfigTaxType = "ANROK"
	PaymentGateConfigTaxTypePrecalculated PaymentGateConfigTaxType = "PRECALCULATED"
)

func (r PaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case PaymentGateConfigTaxTypeNone, PaymentGateConfigTaxTypeStripe, PaymentGateConfigTaxTypeAnrok, PaymentGateConfigTaxTypePrecalculated:
		return true
	}
	return false
}

type PaymentGateConfigParam struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType param.Field[PaymentGateConfigPaymentGateType] `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig param.Field[PaymentGateConfigPrecalculatedTaxConfigParam] `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig param.Field[PaymentGateConfigStripeConfigParam] `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType param.Field[PaymentGateConfigTaxType] `json:"tax_type"`
}

func (r PaymentGateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only applicable if using PRECALCULATED as your tax type.
type PaymentGateConfigPrecalculatedTaxConfigParam struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount param.Field[float64] `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Field[string] `json:"tax_name"`
}

func (r PaymentGateConfigPrecalculatedTaxConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only applicable if using STRIPE as your payment gate type.
type PaymentGateConfigStripeConfigParam struct {
	// If left blank, will default to INVOICE
	PaymentType param.Field[PaymentGateConfigStripeConfigPaymentType] `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata param.Field[map[string]string] `json:"invoice_metadata"`
}

func (r PaymentGateConfigStripeConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentGateConfigV2 struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType PaymentGateConfigV2PaymentGateType `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig PaymentGateConfigV2PrecalculatedTaxConfig `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig PaymentGateConfigV2StripeConfig `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType PaymentGateConfigV2TaxType `json:"tax_type"`
	JSON    paymentGateConfigV2JSON    `json:"-"`
}

// paymentGateConfigV2JSON contains the JSON metadata for the struct
// [PaymentGateConfigV2]
type paymentGateConfigV2JSON struct {
	PaymentGateType        apijson.Field
	PrecalculatedTaxConfig apijson.Field
	StripeConfig           apijson.Field
	TaxType                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PaymentGateConfigV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentGateConfigV2JSON) RawJSON() string {
	return r.raw
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type PaymentGateConfigV2PaymentGateType string

const (
	PaymentGateConfigV2PaymentGateTypeNone     PaymentGateConfigV2PaymentGateType = "NONE"
	PaymentGateConfigV2PaymentGateTypeStripe   PaymentGateConfigV2PaymentGateType = "STRIPE"
	PaymentGateConfigV2PaymentGateTypeExternal PaymentGateConfigV2PaymentGateType = "EXTERNAL"
)

func (r PaymentGateConfigV2PaymentGateType) IsKnown() bool {
	switch r {
	case PaymentGateConfigV2PaymentGateTypeNone, PaymentGateConfigV2PaymentGateTypeStripe, PaymentGateConfigV2PaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type PaymentGateConfigV2PrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName string                                        `json:"tax_name"`
	JSON    paymentGateConfigV2PrecalculatedTaxConfigJSON `json:"-"`
}

// paymentGateConfigV2PrecalculatedTaxConfigJSON contains the JSON metadata for the
// struct [PaymentGateConfigV2PrecalculatedTaxConfig]
type paymentGateConfigV2PrecalculatedTaxConfigJSON struct {
	TaxAmount   apijson.Field
	TaxName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentGateConfigV2PrecalculatedTaxConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentGateConfigV2PrecalculatedTaxConfigJSON) RawJSON() string {
	return r.raw
}

// Only applicable if using STRIPE as your payment gateway type.
type PaymentGateConfigV2StripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType PaymentGateConfigV2StripeConfigPaymentType `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string                   `json:"invoice_metadata"`
	JSON            paymentGateConfigV2StripeConfigJSON `json:"-"`
}

// paymentGateConfigV2StripeConfigJSON contains the JSON metadata for the struct
// [PaymentGateConfigV2StripeConfig]
type paymentGateConfigV2StripeConfigJSON struct {
	PaymentType     apijson.Field
	InvoiceMetadata apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PaymentGateConfigV2StripeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentGateConfigV2StripeConfigJSON) RawJSON() string {
	return r.raw
}

// If left blank, will default to INVOICE
type PaymentGateConfigV2StripeConfigPaymentType string

const (
	PaymentGateConfigV2StripeConfigPaymentTypeInvoice       PaymentGateConfigV2StripeConfigPaymentType = "INVOICE"
	PaymentGateConfigV2StripeConfigPaymentTypePaymentIntent PaymentGateConfigV2StripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r PaymentGateConfigV2StripeConfigPaymentType) IsKnown() bool {
	switch r {
	case PaymentGateConfigV2StripeConfigPaymentTypeInvoice, PaymentGateConfigV2StripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type PaymentGateConfigV2TaxType string

const (
	PaymentGateConfigV2TaxTypeNone          PaymentGateConfigV2TaxType = "NONE"
	PaymentGateConfigV2TaxTypeStripe        PaymentGateConfigV2TaxType = "STRIPE"
	PaymentGateConfigV2TaxTypeAnrok         PaymentGateConfigV2TaxType = "ANROK"
	PaymentGateConfigV2TaxTypePrecalculated PaymentGateConfigV2TaxType = "PRECALCULATED"
)

func (r PaymentGateConfigV2TaxType) IsKnown() bool {
	switch r {
	case PaymentGateConfigV2TaxTypeNone, PaymentGateConfigV2TaxTypeStripe, PaymentGateConfigV2TaxTypeAnrok, PaymentGateConfigV2TaxTypePrecalculated:
		return true
	}
	return false
}

type PaymentGateConfigV2Param struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType param.Field[PaymentGateConfigV2PaymentGateType] `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig param.Field[PaymentGateConfigV2PrecalculatedTaxConfigParam] `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig param.Field[PaymentGateConfigV2StripeConfigParam] `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType param.Field[PaymentGateConfigV2TaxType] `json:"tax_type"`
}

func (r PaymentGateConfigV2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only applicable if using PRECALCULATED as your tax type.
type PaymentGateConfigV2PrecalculatedTaxConfigParam struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount param.Field[float64] `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Field[string] `json:"tax_name"`
}

func (r PaymentGateConfigV2PrecalculatedTaxConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only applicable if using STRIPE as your payment gateway type.
type PaymentGateConfigV2StripeConfigParam struct {
	// If left blank, will default to INVOICE
	PaymentType param.Field[PaymentGateConfigV2StripeConfigPaymentType] `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata param.Field[map[string]string] `json:"invoice_metadata"`
}

func (r PaymentGateConfigV2StripeConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrepaidBalanceThresholdConfiguration struct {
	Commit PrepaidBalanceThresholdConfigurationCommit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool              `json:"is_enabled,required"`
	PaymentGateConfig PaymentGateConfig `json:"payment_gate_config,required"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount float64 `json:"recharge_to_amount,required"`
	// Specify the threshold amount for the contract. Each time the contract's prepaid
	// balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount,required"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID string                                   `json:"custom_credit_type_id" format:"uuid"`
	JSON               prepaidBalanceThresholdConfigurationJSON `json:"-"`
}

// prepaidBalanceThresholdConfigurationJSON contains the JSON metadata for the
// struct [PrepaidBalanceThresholdConfiguration]
type prepaidBalanceThresholdConfigurationJSON struct {
	Commit             apijson.Field
	IsEnabled          apijson.Field
	PaymentGateConfig  apijson.Field
	RechargeToAmount   apijson.Field
	ThresholdAmount    apijson.Field
	CustomCreditTypeID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PrepaidBalanceThresholdConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prepaidBalanceThresholdConfigurationJSON) RawJSON() string {
	return r.raw
}

type PrepaidBalanceThresholdConfigurationCommit struct {
	// Which products the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Which tags the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []CommitSpecifierInput                         `json:"specifiers"`
	JSON       prepaidBalanceThresholdConfigurationCommitJSON `json:"-"`
	BaseThresholdCommit
}

// prepaidBalanceThresholdConfigurationCommitJSON contains the JSON metadata for
// the struct [PrepaidBalanceThresholdConfigurationCommit]
type prepaidBalanceThresholdConfigurationCommitJSON struct {
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Specifiers            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PrepaidBalanceThresholdConfigurationCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prepaidBalanceThresholdConfigurationCommitJSON) RawJSON() string {
	return r.raw
}

type PrepaidBalanceThresholdConfigurationParam struct {
	Commit param.Field[PrepaidBalanceThresholdConfigurationCommitParam] `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         param.Field[bool]                   `json:"is_enabled,required"`
	PaymentGateConfig param.Field[PaymentGateConfigParam] `json:"payment_gate_config,required"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount param.Field[float64] `json:"recharge_to_amount,required"`
	// Specify the threshold amount for the contract. Each time the contract's prepaid
	// balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount param.Field[float64] `json:"threshold_amount,required"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID param.Field[string] `json:"custom_credit_type_id" format:"uuid"`
}

func (r PrepaidBalanceThresholdConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrepaidBalanceThresholdConfigurationCommitParam struct {
	// Which products the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers param.Field[[]CommitSpecifierInputParam] `json:"specifiers"`
	BaseThresholdCommitParam
}

func (r PrepaidBalanceThresholdConfigurationCommitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrepaidBalanceThresholdConfigurationV2 struct {
	Commit PrepaidBalanceThresholdConfigurationV2Commit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                `json:"is_enabled,required"`
	PaymentGateConfig PaymentGateConfigV2 `json:"payment_gate_config,required"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount float64 `json:"recharge_to_amount,required"`
	// Specify the threshold amount for the contract. Each time the contract's balance
	// lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount,required"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID string                                     `json:"custom_credit_type_id" format:"uuid"`
	JSON               prepaidBalanceThresholdConfigurationV2JSON `json:"-"`
}

// prepaidBalanceThresholdConfigurationV2JSON contains the JSON metadata for the
// struct [PrepaidBalanceThresholdConfigurationV2]
type prepaidBalanceThresholdConfigurationV2JSON struct {
	Commit             apijson.Field
	IsEnabled          apijson.Field
	PaymentGateConfig  apijson.Field
	RechargeToAmount   apijson.Field
	ThresholdAmount    apijson.Field
	CustomCreditTypeID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PrepaidBalanceThresholdConfigurationV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prepaidBalanceThresholdConfigurationV2JSON) RawJSON() string {
	return r.raw
}

type PrepaidBalanceThresholdConfigurationV2Commit struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID string `json:"product_id,required"`
	// Which products the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Which tags the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags"`
	Description           string   `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string `json:"name"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []CommitSpecifierInput                           `json:"specifiers"`
	JSON       prepaidBalanceThresholdConfigurationV2CommitJSON `json:"-"`
}

// prepaidBalanceThresholdConfigurationV2CommitJSON contains the JSON metadata for
// the struct [PrepaidBalanceThresholdConfigurationV2Commit]
type prepaidBalanceThresholdConfigurationV2CommitJSON struct {
	ProductID             apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Description           apijson.Field
	Name                  apijson.Field
	Specifiers            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PrepaidBalanceThresholdConfigurationV2Commit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prepaidBalanceThresholdConfigurationV2CommitJSON) RawJSON() string {
	return r.raw
}

type PrepaidBalanceThresholdConfigurationV2Param struct {
	Commit param.Field[PrepaidBalanceThresholdConfigurationV2CommitParam] `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         param.Field[bool]                     `json:"is_enabled,required"`
	PaymentGateConfig param.Field[PaymentGateConfigV2Param] `json:"payment_gate_config,required"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount param.Field[float64] `json:"recharge_to_amount,required"`
	// Specify the threshold amount for the contract. Each time the contract's balance
	// lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount param.Field[float64] `json:"threshold_amount,required"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID param.Field[string] `json:"custom_credit_type_id" format:"uuid"`
}

func (r PrepaidBalanceThresholdConfigurationV2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrepaidBalanceThresholdConfigurationV2CommitParam struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID param.Field[string] `json:"product_id,required"`
	// Which products the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	Description           param.Field[string]   `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name param.Field[string] `json:"name"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers param.Field[[]CommitSpecifierInputParam] `json:"specifiers"`
}

func (r PrepaidBalanceThresholdConfigurationV2CommitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PropertyFilter struct {
	// The name of the event property.
	Name string `json:"name,required"`
	// Determines whether the property must exist in the event. If true, only events
	// with this property will pass the filter. If false, only events without this
	// property will pass the filter. If null or omitted, the existence of the property
	// is optional.
	Exists bool `json:"exists"`
	// Specifies the allowed values for the property to match an event. An event will
	// pass the filter only if its property value is included in this list. If
	// undefined, all property values will pass the filter. Must be non-empty if
	// present.
	InValues []string `json:"in_values"`
	// Specifies the values that prevent an event from matching the filter. An event
	// will not pass the filter if its property value is included in this list. If null
	// or empty, all property values will pass the filter. Must be non-empty if
	// present.
	NotInValues []string           `json:"not_in_values"`
	JSON        propertyFilterJSON `json:"-"`
}

// propertyFilterJSON contains the JSON metadata for the struct [PropertyFilter]
type propertyFilterJSON struct {
	Name        apijson.Field
	Exists      apijson.Field
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PropertyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r propertyFilterJSON) RawJSON() string {
	return r.raw
}

type PropertyFilterParam struct {
	// The name of the event property.
	Name param.Field[string] `json:"name,required"`
	// Determines whether the property must exist in the event. If true, only events
	// with this property will pass the filter. If false, only events without this
	// property will pass the filter. If null or omitted, the existence of the property
	// is optional.
	Exists param.Field[bool] `json:"exists"`
	// Specifies the allowed values for the property to match an event. An event will
	// pass the filter only if its property value is included in this list. If
	// undefined, all property values will pass the filter. Must be non-empty if
	// present.
	InValues param.Field[[]string] `json:"in_values"`
	// Specifies the values that prevent an event from matching the filter. An event
	// will not pass the filter if its property value is included in this list. If null
	// or empty, all property values will pass the filter. Must be non-empty if
	// present.
	NotInValues param.Field[[]string] `json:"not_in_values"`
}

func (r PropertyFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProService struct {
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
	NetsuiteSalesOrderID string         `json:"netsuite_sales_order_id"`
	JSON                 proServiceJSON `json:"-"`
}

// proServiceJSON contains the JSON metadata for the struct [ProService]
type proServiceJSON struct {
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

func (r *ProService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proServiceJSON) RawJSON() string {
	return r.raw
}

type Rate struct {
	RateType   RateRateType   `json:"rate_type,required"`
	CreditType CreditTypeData `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]interface{} `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
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
	Tiers []Tier `json:"tiers"`
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
	RateRateTypePercentage   RateRateType = "PERCENTAGE"
	RateRateTypeSubscription RateRateType = "SUBSCRIPTION"
	RateRateTypeCustom       RateRateType = "CUSTOM"
	RateRateTypeTiered       RateRateType = "TIERED"
)

func (r RateRateType) IsKnown() bool {
	switch r {
	case RateRateTypeFlat, RateRateTypePercentage, RateRateTypeSubscription, RateRateTypeCustom, RateRateTypeTiered:
		return true
	}
	return false
}

type ScheduledCharge struct {
	ID           string                 `json:"id,required" format:"uuid"`
	Product      ScheduledChargeProduct `json:"product,required"`
	Schedule     SchedulePointInTime    `json:"schedule,required"`
	ArchivedAt   time.Time              `json:"archived_at" format:"date-time"`
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
	ArchivedAt           apijson.Field
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

type ScheduleDuration struct {
	ScheduleItems []ScheduleDurationScheduleItem `json:"schedule_items,required"`
	CreditType    CreditTypeData                 `json:"credit_type"`
	JSON          scheduleDurationJSON           `json:"-"`
}

// scheduleDurationJSON contains the JSON metadata for the struct
// [ScheduleDuration]
type scheduleDurationJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScheduleDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleDurationJSON) RawJSON() string {
	return r.raw
}

type ScheduleDurationScheduleItem struct {
	ID           string                           `json:"id,required" format:"uuid"`
	Amount       float64                          `json:"amount,required"`
	EndingBefore time.Time                        `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                        `json:"starting_at,required" format:"date-time"`
	JSON         scheduleDurationScheduleItemJSON `json:"-"`
}

// scheduleDurationScheduleItemJSON contains the JSON metadata for the struct
// [ScheduleDurationScheduleItem]
type scheduleDurationScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScheduleDurationScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleDurationScheduleItemJSON) RawJSON() string {
	return r.raw
}

type SchedulePointInTime struct {
	CreditType CreditTypeData `json:"credit_type"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice  bool                              `json:"do_not_invoice"`
	ScheduleItems []SchedulePointInTimeScheduleItem `json:"schedule_items"`
	JSON          schedulePointInTimeJSON           `json:"-"`
}

// schedulePointInTimeJSON contains the JSON metadata for the struct
// [SchedulePointInTime]
type schedulePointInTimeJSON struct {
	CreditType    apijson.Field
	DoNotInvoice  apijson.Field
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
	Quantity  float64                             `json:"quantity,required"`
	Timestamp time.Time                           `json:"timestamp,required" format:"date-time"`
	UnitPrice float64                             `json:"unit_price,required"`
	InvoiceID string                              `json:"invoice_id,nullable" format:"uuid"`
	JSON      schedulePointInTimeScheduleItemJSON `json:"-"`
}

// schedulePointInTimeScheduleItemJSON contains the JSON metadata for the struct
// [SchedulePointInTimeScheduleItem]
type schedulePointInTimeScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	InvoiceID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchedulePointInTimeScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schedulePointInTimeScheduleItemJSON) RawJSON() string {
	return r.raw
}

type SpendThresholdConfiguration struct {
	Commit BaseThresholdCommit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool              `json:"is_enabled,required"`
	PaymentGateConfig PaymentGateConfig `json:"payment_gate_config,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64                         `json:"threshold_amount,required"`
	JSON            spendThresholdConfigurationJSON `json:"-"`
}

// spendThresholdConfigurationJSON contains the JSON metadata for the struct
// [SpendThresholdConfiguration]
type spendThresholdConfigurationJSON struct {
	Commit            apijson.Field
	IsEnabled         apijson.Field
	PaymentGateConfig apijson.Field
	ThresholdAmount   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendThresholdConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendThresholdConfigurationJSON) RawJSON() string {
	return r.raw
}

type SpendThresholdConfigurationParam struct {
	Commit param.Field[BaseThresholdCommitParam] `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         param.Field[bool]                   `json:"is_enabled,required"`
	PaymentGateConfig param.Field[PaymentGateConfigParam] `json:"payment_gate_config,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount param.Field[float64] `json:"threshold_amount,required"`
}

func (r SpendThresholdConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SpendThresholdConfigurationV2 struct {
	Commit SpendThresholdConfigurationV2Commit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                `json:"is_enabled,required"`
	PaymentGateConfig PaymentGateConfigV2 `json:"payment_gate_config,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64                           `json:"threshold_amount,required"`
	JSON            spendThresholdConfigurationV2JSON `json:"-"`
}

// spendThresholdConfigurationV2JSON contains the JSON metadata for the struct
// [SpendThresholdConfigurationV2]
type spendThresholdConfigurationV2JSON struct {
	Commit            apijson.Field
	IsEnabled         apijson.Field
	PaymentGateConfig apijson.Field
	ThresholdAmount   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendThresholdConfigurationV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendThresholdConfigurationV2JSON) RawJSON() string {
	return r.raw
}

type SpendThresholdConfigurationV2Commit struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID   string `json:"product_id,required"`
	Description string `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string                                  `json:"name"`
	JSON spendThresholdConfigurationV2CommitJSON `json:"-"`
}

// spendThresholdConfigurationV2CommitJSON contains the JSON metadata for the
// struct [SpendThresholdConfigurationV2Commit]
type spendThresholdConfigurationV2CommitJSON struct {
	ProductID   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendThresholdConfigurationV2Commit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendThresholdConfigurationV2CommitJSON) RawJSON() string {
	return r.raw
}

type SpendThresholdConfigurationV2Param struct {
	Commit param.Field[SpendThresholdConfigurationV2CommitParam] `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         param.Field[bool]                     `json:"is_enabled,required"`
	PaymentGateConfig param.Field[PaymentGateConfigV2Param] `json:"payment_gate_config,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount param.Field[float64] `json:"threshold_amount,required"`
}

func (r SpendThresholdConfigurationV2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SpendThresholdConfigurationV2CommitParam struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID   param.Field[string] `json:"product_id,required"`
	Description param.Field[string] `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name param.Field[string] `json:"name"`
}

func (r SpendThresholdConfigurationV2CommitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Subscription struct {
	CollectionSchedule SubscriptionCollectionSchedule `json:"collection_schedule,required"`
	Proration          SubscriptionProration          `json:"proration,required"`
	// List of quantity schedule items for the subscription. Only includes the current
	// quantity and future quantity changes.
	QuantitySchedule []SubscriptionQuantitySchedule `json:"quantity_schedule,required"`
	StartingAt       time.Time                      `json:"starting_at,required" format:"date-time"`
	SubscriptionRate SubscriptionSubscriptionRate   `json:"subscription_rate,required"`
	ID               string                         `json:"id" format:"uuid"`
	CustomFields     map[string]string              `json:"custom_fields"`
	Description      string                         `json:"description"`
	EndingBefore     time.Time                      `json:"ending_before" format:"date-time"`
	FiatCreditTypeID string                         `json:"fiat_credit_type_id" format:"uuid"`
	Name             string                         `json:"name"`
	JSON             subscriptionJSON               `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	CollectionSchedule apijson.Field
	Proration          apijson.Field
	QuantitySchedule   apijson.Field
	StartingAt         apijson.Field
	SubscriptionRate   apijson.Field
	ID                 apijson.Field
	CustomFields       apijson.Field
	Description        apijson.Field
	EndingBefore       apijson.Field
	FiatCreditTypeID   apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionCollectionSchedule string

const (
	SubscriptionCollectionScheduleAdvance SubscriptionCollectionSchedule = "ADVANCE"
	SubscriptionCollectionScheduleArrears SubscriptionCollectionSchedule = "ARREARS"
)

func (r SubscriptionCollectionSchedule) IsKnown() bool {
	switch r {
	case SubscriptionCollectionScheduleAdvance, SubscriptionCollectionScheduleArrears:
		return true
	}
	return false
}

type SubscriptionProration struct {
	InvoiceBehavior SubscriptionProrationInvoiceBehavior `json:"invoice_behavior,required"`
	IsProrated      bool                                 `json:"is_prorated,required"`
	JSON            subscriptionProrationJSON            `json:"-"`
}

// subscriptionProrationJSON contains the JSON metadata for the struct
// [SubscriptionProration]
type subscriptionProrationJSON struct {
	InvoiceBehavior apijson.Field
	IsProrated      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SubscriptionProration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionProrationJSON) RawJSON() string {
	return r.raw
}

type SubscriptionProrationInvoiceBehavior string

const (
	SubscriptionProrationInvoiceBehaviorBillImmediately          SubscriptionProrationInvoiceBehavior = "BILL_IMMEDIATELY"
	SubscriptionProrationInvoiceBehaviorBillOnNextCollectionDate SubscriptionProrationInvoiceBehavior = "BILL_ON_NEXT_COLLECTION_DATE"
)

func (r SubscriptionProrationInvoiceBehavior) IsKnown() bool {
	switch r {
	case SubscriptionProrationInvoiceBehaviorBillImmediately, SubscriptionProrationInvoiceBehaviorBillOnNextCollectionDate:
		return true
	}
	return false
}

type SubscriptionQuantitySchedule struct {
	Quantity     float64                          `json:"quantity,required"`
	StartingAt   time.Time                        `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time                        `json:"ending_before" format:"date-time"`
	JSON         subscriptionQuantityScheduleJSON `json:"-"`
}

// subscriptionQuantityScheduleJSON contains the JSON metadata for the struct
// [SubscriptionQuantitySchedule]
type subscriptionQuantityScheduleJSON struct {
	Quantity     apijson.Field
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionQuantitySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionQuantityScheduleJSON) RawJSON() string {
	return r.raw
}

type SubscriptionSubscriptionRate struct {
	BillingFrequency SubscriptionSubscriptionRateBillingFrequency `json:"billing_frequency,required"`
	Product          SubscriptionSubscriptionRateProduct          `json:"product,required"`
	JSON             subscriptionSubscriptionRateJSON             `json:"-"`
}

// subscriptionSubscriptionRateJSON contains the JSON metadata for the struct
// [SubscriptionSubscriptionRate]
type subscriptionSubscriptionRateJSON struct {
	BillingFrequency apijson.Field
	Product          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubscriptionSubscriptionRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSubscriptionRateJSON) RawJSON() string {
	return r.raw
}

type SubscriptionSubscriptionRateBillingFrequency string

const (
	SubscriptionSubscriptionRateBillingFrequencyMonthly   SubscriptionSubscriptionRateBillingFrequency = "MONTHLY"
	SubscriptionSubscriptionRateBillingFrequencyQuarterly SubscriptionSubscriptionRateBillingFrequency = "QUARTERLY"
	SubscriptionSubscriptionRateBillingFrequencyAnnual    SubscriptionSubscriptionRateBillingFrequency = "ANNUAL"
	SubscriptionSubscriptionRateBillingFrequencyWeekly    SubscriptionSubscriptionRateBillingFrequency = "WEEKLY"
)

func (r SubscriptionSubscriptionRateBillingFrequency) IsKnown() bool {
	switch r {
	case SubscriptionSubscriptionRateBillingFrequencyMonthly, SubscriptionSubscriptionRateBillingFrequencyQuarterly, SubscriptionSubscriptionRateBillingFrequencyAnnual, SubscriptionSubscriptionRateBillingFrequencyWeekly:
		return true
	}
	return false
}

type SubscriptionSubscriptionRateProduct struct {
	ID   string                                  `json:"id,required" format:"uuid"`
	Name string                                  `json:"name,required"`
	JSON subscriptionSubscriptionRateProductJSON `json:"-"`
}

// subscriptionSubscriptionRateProductJSON contains the JSON metadata for the
// struct [SubscriptionSubscriptionRateProduct]
type subscriptionSubscriptionRateProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionSubscriptionRateProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSubscriptionRateProductJSON) RawJSON() string {
	return r.raw
}

type Tier struct {
	Price float64  `json:"price,required"`
	Size  float64  `json:"size"`
	JSON  tierJSON `json:"-"`
}

// tierJSON contains the JSON metadata for the struct [Tier]
type tierJSON struct {
	Price       apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Tier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tierJSON) RawJSON() string {
	return r.raw
}

type TierParam struct {
	Price param.Field[float64] `json:"price,required"`
	Size  param.Field[float64] `json:"size"`
}

func (r TierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
