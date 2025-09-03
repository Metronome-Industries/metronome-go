// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type BaseThresholdCommit struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID   string `json:"product_id,required"`
	Description string `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductID   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseThresholdCommit) RawJSON() string { return r.JSON.raw }
func (r *BaseThresholdCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BaseThresholdCommit to a BaseThresholdCommitParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BaseThresholdCommitParam.Overrides()
func (r BaseThresholdCommit) ToParam() BaseThresholdCommitParam {
	return param.Override[BaseThresholdCommitParam](json.RawMessage(r.RawJSON()))
}

// The property ProductID is required.
type BaseThresholdCommitParam struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID   string            `json:"product_id,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r BaseThresholdCommitParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseThresholdCommitParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseThresholdCommitParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseUsageFilter struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupKey    respjson.Field
		GroupValues respjson.Field
		StartingAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseUsageFilter) RawJSON() string { return r.JSON.raw }
func (r *BaseUsageFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BaseUsageFilter to a BaseUsageFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BaseUsageFilterParam.Overrides()
func (r BaseUsageFilter) ToParam() BaseUsageFilterParam {
	return param.Override[BaseUsageFilterParam](json.RawMessage(r.RawJSON()))
}

// The properties GroupKey, GroupValues are required.
type BaseUsageFilterParam struct {
	GroupKey    string               `json:"group_key,required"`
	GroupValues []string             `json:"group_values,omitzero,required"`
	StartingAt  param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r BaseUsageFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseUsageFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseUsageFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Commit struct {
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the commit was created.
	//
	//   - Recurring commits: latter of commit service period date and parent commit
	//     start date
	//   - Rollover commits: when the new contract started
	CreatedAt time.Time     `json:"created_at,required" format:"date-time"`
	Product   CommitProduct `json:"product,required"`
	// Any of "PREPAID", "POSTPAID".
	Type CommitType `json:"type,required"`
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
	Balance  float64        `json:"balance"`
	Contract CommitContract `json:"contract"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
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
	Ledger []CommitLedgerUnion `json:"ledger"`
	Name   string              `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// Any of "COMMIT_RATE", "LIST_RATE".
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
	UniquenessKey string `json:"uniqueness_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CreatedAt               respjson.Field
		Product                 respjson.Field
		Type                    respjson.Field
		AccessSchedule          respjson.Field
		Amount                  respjson.Field
		ApplicableContractIDs   respjson.Field
		ApplicableProductIDs    respjson.Field
		ApplicableProductTags   respjson.Field
		ArchivedAt              respjson.Field
		Balance                 respjson.Field
		Contract                respjson.Field
		CustomFields            respjson.Field
		Description             respjson.Field
		HierarchyConfiguration  respjson.Field
		InvoiceContract         respjson.Field
		InvoiceSchedule         respjson.Field
		Ledger                  respjson.Field
		Name                    respjson.Field
		NetsuiteSalesOrderID    respjson.Field
		Priority                respjson.Field
		RateType                respjson.Field
		RolledOverFrom          respjson.Field
		RolloverFraction        respjson.Field
		SalesforceOpportunityID respjson.Field
		Specifiers              respjson.Field
		UniquenessKey           respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Commit) RawJSON() string { return r.JSON.raw }
func (r *Commit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitProduct) RawJSON() string { return r.JSON.raw }
func (r *CommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitType string

const (
	CommitTypePrepaid  CommitType = "PREPAID"
	CommitTypePostpaid CommitType = "POSTPAID"
)

type CommitContract struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitContract) RawJSON() string { return r.JSON.raw }
func (r *CommitContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The contract that this commit will be billed on.
type CommitInvoiceContract struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitInvoiceContract) RawJSON() string { return r.JSON.raw }
func (r *CommitInvoiceContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CommitLedgerUnion contains all possible properties and values from
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
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CommitLedgerUnion struct {
	Amount        float64   `json:"amount"`
	SegmentID     string    `json:"segment_id"`
	Timestamp     time.Time `json:"timestamp"`
	Type          string    `json:"type"`
	InvoiceID     string    `json:"invoice_id"`
	ContractID    string    `json:"contract_id"`
	NewContractID string    `json:"new_contract_id"`
	Reason        string    `json:"reason"`
	JSON          struct {
		Amount        respjson.Field
		SegmentID     respjson.Field
		Timestamp     respjson.Field
		Type          respjson.Field
		InvoiceID     respjson.Field
		ContractID    respjson.Field
		NewContractID respjson.Field
		Reason        respjson.Field
		raw           string
	} `json:"-"`
}

func (u CommitLedgerUnion) AsCommitLedgerPrepaidCommitSegmentStartLedgerEntry() (v CommitLedgerPrepaidCommitSegmentStartLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry() (v CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPrepaidCommitRolloverLedgerEntry() (v CommitLedgerPrepaidCommitRolloverLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPrepaidCommitExpirationLedgerEntry() (v CommitLedgerPrepaidCommitExpirationLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPrepaidCommitCanceledLedgerEntry() (v CommitLedgerPrepaidCommitCanceledLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPrepaidCommitCreditedLedgerEntry() (v CommitLedgerPrepaidCommitCreditedLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry() (v CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPostpaidCommitInitialBalanceLedgerEntry() (v CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry() (v CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPostpaidCommitRolloverLedgerEntry() (v CommitLedgerPostpaidCommitRolloverLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPostpaidCommitTrueupLedgerEntry() (v CommitLedgerPostpaidCommitTrueupLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPrepaidCommitManualLedgerEntry() (v CommitLedgerPrepaidCommitManualLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPostpaidCommitManualLedgerEntry() (v CommitLedgerPostpaidCommitManualLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitLedgerUnion) AsCommitLedgerPostpaidCommitExpirationLedgerEntry() (v CommitLedgerPostpaidCommitExpirationLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CommitLedgerUnion) RawJSON() string { return u.JSON.raw }

func (r *CommitLedgerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_SEGMENT_START".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPrepaidCommitSegmentStartLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) RawJSON() string {
	return r.JSON.raw
}
func (r *CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64   `json:"amount,required"`
	NewContractID string    `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string    `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_ROLLOVER".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount        respjson.Field
		NewContractID respjson.Field
		SegmentID     respjson.Field
		Timestamp     respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPrepaidCommitRolloverLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_EXPIRATION".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPrepaidCommitExpirationLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_CANCELED".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPrepaidCommitCanceledLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_CREDITED".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPrepaidCommitCreditedLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_SEAT_BASED_ADJUSTMENT".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_INITIAL_BALANCE".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) RawJSON() string {
	return r.JSON.raw
}
func (r *CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPostpaidCommitRolloverLedgerEntry struct {
	Amount        float64   `json:"amount,required"`
	NewContractID string    `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string    `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_ROLLOVER".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount        respjson.Field
		NewContractID respjson.Field
		SegmentID     respjson.Field
		Timestamp     respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPostpaidCommitRolloverLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPostpaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_TRUEUP".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPostpaidCommitTrueupLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPrepaidCommitManualLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	Reason    string    `json:"reason,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_MANUAL".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Reason      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPrepaidCommitManualLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPrepaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPostpaidCommitManualLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	Reason    string    `json:"reason,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_MANUAL".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Reason      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPostpaidCommitManualLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPostpaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitLedgerPostpaidCommitExpirationLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_EXPIRATION".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitLedgerPostpaidCommitExpirationLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CommitLedgerPostpaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitRateType string

const (
	CommitRateTypeCommitRate CommitRateType = "COMMIT_RATE"
	CommitRateTypeListRate   CommitRateType = "LIST_RATE"
)

type CommitRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommitID    respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitRolledOverFrom) RawJSON() string { return r.JSON.raw }
func (r *CommitRolledOverFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitHierarchyConfiguration struct {
	ChildAccess CommitHierarchyConfigurationChildAccessUnion `json:"child_access,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChildAccess respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitHierarchyConfiguration) RawJSON() string { return r.JSON.raw }
func (r *CommitHierarchyConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CommitHierarchyConfiguration to a
// CommitHierarchyConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CommitHierarchyConfigurationParam.Overrides()
func (r CommitHierarchyConfiguration) ToParam() CommitHierarchyConfigurationParam {
	return param.Override[CommitHierarchyConfigurationParam](json.RawMessage(r.RawJSON()))
}

// CommitHierarchyConfigurationChildAccessUnion contains all possible properties
// and values from
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CommitHierarchyConfigurationChildAccessUnion struct {
	Type string `json:"type"`
	// This field is from variant
	// [CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
	ContractIDs []string `json:"contract_ids"`
	JSON        struct {
		Type        respjson.Field
		ContractIDs respjson.Field
		raw         string
	} `json:"-"`
}

func (u CommitHierarchyConfigurationChildAccessUnion) AsCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll() (v CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitHierarchyConfigurationChildAccessUnion) AsCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone() (v CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CommitHierarchyConfigurationChildAccessUnion) AsCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs() (v CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CommitHierarchyConfigurationChildAccessUnion) RawJSON() string { return u.JSON.raw }

func (r *CommitHierarchyConfigurationChildAccessUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	// Any of "ALL".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) RawJSON() string {
	return r.JSON.raw
}
func (r *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	// Any of "NONE".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) RawJSON() string {
	return r.JSON.raw
}
func (r *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs []string `json:"contract_ids,required" format:"uuid"`
	// Any of "CONTRACT_IDS".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContractIDs respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) RawJSON() string {
	return r.JSON.raw
}
func (r *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ChildAccess is required.
type CommitHierarchyConfigurationParam struct {
	ChildAccess CommitHierarchyConfigurationChildAccessUnionParam `json:"child_access,omitzero,required"`
	paramObj
}

func (r CommitHierarchyConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow CommitHierarchyConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommitHierarchyConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CommitHierarchyConfigurationChildAccessUnionParam struct {
	OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll         *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam         `json:",omitzero,inline"`
	OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone        *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam        `json:",omitzero,inline"`
	OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam `json:",omitzero,inline"`
	paramUnion
}

func (u CommitHierarchyConfigurationChildAccessUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll, u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone, u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs)
}
func (u *CommitHierarchyConfigurationChildAccessUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CommitHierarchyConfigurationChildAccessUnionParam) asAny() any {
	if !param.IsOmitted(u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) {
		return u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll
	} else if !param.IsOmitted(u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) {
		return u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone
	} else if !param.IsOmitted(u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) {
		return u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CommitHierarchyConfigurationChildAccessUnionParam) GetContractIDs() []string {
	if vt := u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs; vt != nil {
		return vt.ContractIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CommitHierarchyConfigurationChildAccessUnionParam) GetType() *string {
	if vt := u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The property Type is required.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam struct {
	// Any of "ALL".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam) MarshalJSON() (data []byte, err error) {
	type shadow CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllParam](
		"type", "ALL",
	)
}

// The property Type is required.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam struct {
	// Any of "NONE".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam) MarshalJSON() (data []byte, err error) {
	type shadow CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneParam](
		"type", "NONE",
	)
}

// The properties ContractIDs, Type are required.
type CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam struct {
	ContractIDs []string `json:"contract_ids,omitzero,required" format:"uuid"`
	// Any of "CONTRACT_IDS".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam) MarshalJSON() (data []byte, err error) {
	type shadow CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CommitHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsParam](
		"type", "CONTRACT_IDS",
	)
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type CommitRate struct {
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM".
	RateType CommitRateRateType `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price float64 `json:"price"`
	// Only set for TIERED rate_type.
	Tiers []Tier `json:"tiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RateType    respjson.Field
		Price       respjson.Field
		Tiers       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitRate) RawJSON() string { return r.JSON.raw }
func (r *CommitRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CommitRate to a CommitRateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CommitRateParam.Overrides()
func (r CommitRate) ToParam() CommitRateParam {
	return param.Override[CommitRateParam](json.RawMessage(r.RawJSON()))
}

type CommitRateRateType string

const (
	CommitRateRateTypeFlat         CommitRateRateType = "FLAT"
	CommitRateRateTypePercentage   CommitRateRateType = "PERCENTAGE"
	CommitRateRateTypeSubscription CommitRateRateType = "SUBSCRIPTION"
	CommitRateRateTypeTiered       CommitRateRateType = "TIERED"
	CommitRateRateTypeCustom       CommitRateRateType = "CUSTOM"
)

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
//
// The property RateType is required.
type CommitRateParam struct {
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM".
	RateType CommitRateRateType `json:"rate_type,omitzero,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price param.Opt[float64] `json:"price,omitzero"`
	// Only set for TIERED rate_type.
	Tiers []TierParam `json:"tiers,omitzero"`
	paramObj
}

func (r CommitRateParam) MarshalJSON() (data []byte, err error) {
	type shadow CommitRateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommitRateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string `json:"product_tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PresentationGroupValues respjson.Field
		PricingGroupValues      respjson.Field
		ProductID               respjson.Field
		ProductTags             respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitSpecifier) RawJSON() string { return r.JSON.raw }
func (r *CommitSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitSpecifierInput struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string `json:"product_tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PresentationGroupValues respjson.Field
		PricingGroupValues      respjson.Field
		ProductID               respjson.Field
		ProductTags             respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitSpecifierInput) RawJSON() string { return r.JSON.raw }
func (r *CommitSpecifierInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CommitSpecifierInput to a CommitSpecifierInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CommitSpecifierInputParam.Overrides()
func (r CommitSpecifierInput) ToParam() CommitSpecifierInputParam {
	return param.Override[CommitSpecifierInputParam](json.RawMessage(r.RawJSON()))
}

type CommitSpecifierInputParam struct {
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID               param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	PresentationGroupValues map[string]string `json:"presentation_group_values,omitzero"`
	PricingGroupValues      map[string]string `json:"pricing_group_values,omitzero"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string `json:"product_tags,omitzero"`
	paramObj
}

func (r CommitSpecifierInputParam) MarshalJSON() (data []byte, err error) {
	type shadow CommitSpecifierInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommitSpecifierInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Contract struct {
	ID         string                    `json:"id,required" format:"uuid"`
	Amendments []ContractAmendment       `json:"amendments,required"`
	Current    ContractWithoutAmendments `json:"current,required"`
	CustomerID string                    `json:"customer_id,required" format:"uuid"`
	Initial    ContractWithoutAmendments `json:"initial,required"`
	// RFC 3339 timestamp indicating when the contract was archived. If not returned,
	// the contract is not archived.
	ArchivedAt time.Time `json:"archived_at" format:"date-time"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
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
	//
	// Any of "ALL".
	ScheduledChargesOnUsageInvoices ContractScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	SpendThresholdConfiguration     SpendThresholdConfiguration             `json:"spend_threshold_configuration"`
	// List of subscriptions on the contract.
	Subscriptions []Subscription `json:"subscriptions"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string `json:"uniqueness_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                   respjson.Field
		Amendments                           respjson.Field
		Current                              respjson.Field
		CustomerID                           respjson.Field
		Initial                              respjson.Field
		ArchivedAt                           respjson.Field
		CustomFields                         respjson.Field
		CustomerBillingProviderConfiguration respjson.Field
		PrepaidBalanceThresholdConfiguration respjson.Field
		Priority                             respjson.Field
		ScheduledChargesOnUsageInvoices      respjson.Field
		SpendThresholdConfiguration          respjson.Field
		Subscriptions                        respjson.Field
		UniquenessKey                        respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Contract) RawJSON() string { return r.JSON.raw }
func (r *Contract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	ResellerRoyalties []ContractAmendmentResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Commits                 respjson.Field
		CreatedAt               respjson.Field
		CreatedBy               respjson.Field
		Overrides               respjson.Field
		ScheduledCharges        respjson.Field
		StartingAt              respjson.Field
		Credits                 respjson.Field
		Discounts               respjson.Field
		NetsuiteSalesOrderID    respjson.Field
		ProfessionalServices    respjson.Field
		ResellerRoyalties       respjson.Field
		SalesforceOpportunityID respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractAmendment) RawJSON() string { return r.JSON.raw }
func (r *ContractAmendment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractAmendmentResellerRoyalty struct {
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType          string    `json:"reseller_type,required"`
	AwsAccountNumber      string    `json:"aws_account_number"`
	AwsOfferID            string    `json:"aws_offer_id"`
	AwsPayerReferenceID   string    `json:"aws_payer_reference_id"`
	EndingBefore          time.Time `json:"ending_before,nullable" format:"date-time"`
	Fraction              float64   `json:"fraction"`
	GcpAccountID          string    `json:"gcp_account_id"`
	GcpOfferID            string    `json:"gcp_offer_id"`
	NetsuiteResellerID    string    `json:"netsuite_reseller_id"`
	ResellerContractValue float64   `json:"reseller_contract_value"`
	StartingAt            time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResellerType          respjson.Field
		AwsAccountNumber      respjson.Field
		AwsOfferID            respjson.Field
		AwsPayerReferenceID   respjson.Field
		EndingBefore          respjson.Field
		Fraction              respjson.Field
		GcpAccountID          respjson.Field
		GcpOfferID            respjson.Field
		NetsuiteResellerID    respjson.Field
		ResellerContractValue respjson.Field
		StartingAt            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractAmendmentResellerRoyalty) RawJSON() string { return r.JSON.raw }
func (r *ContractAmendmentResellerRoyalty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The billing provider configuration associated with a contract.
type ContractCustomerBillingProviderConfiguration struct {
	ArchivedAt time.Time `json:"archived_at,required" format:"date-time"`
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace".
	BillingProvider string `json:"billing_provider,required"`
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod string `json:"delivery_method,required"`
	ID             string `json:"id" format:"uuid"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider.
	Configuration map[string]any `json:"configuration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArchivedAt      respjson.Field
		BillingProvider respjson.Field
		DeliveryMethod  respjson.Field
		ID              respjson.Field
		Configuration   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractCustomerBillingProviderConfiguration) RawJSON() string { return r.JSON.raw }
func (r *ContractCustomerBillingProviderConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
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
	HierarchyConfiguration HierarchyConfigurationUnion `json:"hierarchy_configuration"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
	// prices automatically. EXPLICIT prioritization requires specifying priorities for
	// each multiplier; the one with the lowest priority value will be prioritized
	// first.
	//
	// Any of "LOWEST_MULTIPLIER", "EXPLICIT".
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
	//
	// Any of "ALL".
	ScheduledChargesOnUsageInvoices ContractV2ScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	SpendThresholdConfiguration     SpendThresholdConfigurationV2             `json:"spend_threshold_configuration"`
	// List of subscriptions on the contract.
	Subscriptions      []Subscription `json:"subscriptions"`
	TotalContractValue float64        `json:"total_contract_value"`
	// Optional uniqueness key to prevent duplicate contract creations.
	UniquenessKey string `json:"uniqueness_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                   respjson.Field
		Commits                              respjson.Field
		CreatedAt                            respjson.Field
		CreatedBy                            respjson.Field
		CustomerID                           respjson.Field
		Overrides                            respjson.Field
		ScheduledCharges                     respjson.Field
		StartingAt                           respjson.Field
		Transitions                          respjson.Field
		UsageFilter                          respjson.Field
		UsageStatementSchedule               respjson.Field
		ArchivedAt                           respjson.Field
		Credits                              respjson.Field
		CustomFields                         respjson.Field
		CustomerBillingProviderConfiguration respjson.Field
		Discounts                            respjson.Field
		EndingBefore                         respjson.Field
		HasMore                              respjson.Field
		HierarchyConfiguration               respjson.Field
		MultiplierOverridePrioritization     respjson.Field
		Name                                 respjson.Field
		NetPaymentTermsDays                  respjson.Field
		NetsuiteSalesOrderID                 respjson.Field
		PrepaidBalanceThresholdConfiguration respjson.Field
		Priority                             respjson.Field
		ProfessionalServices                 respjson.Field
		RateCardID                           respjson.Field
		RecurringCommits                     respjson.Field
		RecurringCredits                     respjson.Field
		ResellerRoyalties                    respjson.Field
		SalesforceOpportunityID              respjson.Field
		ScheduledChargesOnUsageInvoices      respjson.Field
		SpendThresholdConfiguration          respjson.Field
		Subscriptions                        respjson.Field
		TotalContractValue                   respjson.Field
		UniquenessKey                        respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2) RawJSON() string { return r.JSON.raw }
func (r *ContractV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2Commit struct {
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the commit was created.
	//
	//   - Recurring commits: latter of commit service period date and parent commit
	//     start date
	//   - Rollover commits: when the new contract started
	CreatedAt time.Time               `json:"created_at,required" format:"date-time"`
	Product   ContractV2CommitProduct `json:"product,required"`
	// Any of "PREPAID", "POSTPAID".
	Type string `json:"type,required"`
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
	Balance  float64                  `json:"balance"`
	Contract ContractV2CommitContract `json:"contract"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// The contract that this commit will be billed on.
	InvoiceContract ContractV2CommitInvoiceContract `json:"invoice_contract"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule SchedulePointInTime `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger []ContractV2CommitLedgerUnion `json:"ledger"`
	Name   string                        `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType         string                         `json:"rate_type"`
	RolledOverFrom   ContractV2CommitRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction float64                        `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CreatedAt               respjson.Field
		Product                 respjson.Field
		Type                    respjson.Field
		AccessSchedule          respjson.Field
		ApplicableContractIDs   respjson.Field
		ApplicableProductIDs    respjson.Field
		ApplicableProductTags   respjson.Field
		ArchivedAt              respjson.Field
		Balance                 respjson.Field
		Contract                respjson.Field
		CustomFields            respjson.Field
		Description             respjson.Field
		HierarchyConfiguration  respjson.Field
		InvoiceContract         respjson.Field
		InvoiceSchedule         respjson.Field
		Ledger                  respjson.Field
		Name                    respjson.Field
		NetsuiteSalesOrderID    respjson.Field
		Priority                respjson.Field
		RateType                respjson.Field
		RolledOverFrom          respjson.Field
		RolloverFraction        respjson.Field
		SalesforceOpportunityID respjson.Field
		Specifiers              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2Commit) RawJSON() string { return r.JSON.raw }
func (r *ContractV2Commit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitProduct) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitContract struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitContract) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The contract that this commit will be billed on.
type ContractV2CommitInvoiceContract struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitInvoiceContract) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitInvoiceContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContractV2CommitLedgerUnion contains all possible properties and values from
// [ContractV2CommitLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractV2CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractV2CommitLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractV2CommitLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractV2CommitLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractV2CommitLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractV2CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry],
// [ContractV2CommitLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractV2CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractV2CommitLedgerPostpaidCommitRolloverLedgerEntry],
// [ContractV2CommitLedgerPostpaidCommitTrueupLedgerEntry],
// [ContractV2CommitLedgerPrepaidCommitManualLedgerEntry],
// [ContractV2CommitLedgerPostpaidCommitManualLedgerEntry],
// [ContractV2CommitLedgerPostpaidCommitExpirationLedgerEntry].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ContractV2CommitLedgerUnion struct {
	Amount        float64   `json:"amount"`
	SegmentID     string    `json:"segment_id"`
	Timestamp     time.Time `json:"timestamp"`
	Type          string    `json:"type"`
	InvoiceID     string    `json:"invoice_id"`
	ContractID    string    `json:"contract_id"`
	NewContractID string    `json:"new_contract_id"`
	Reason        string    `json:"reason"`
	JSON          struct {
		Amount        respjson.Field
		SegmentID     respjson.Field
		Timestamp     respjson.Field
		Type          respjson.Field
		InvoiceID     respjson.Field
		ContractID    respjson.Field
		NewContractID respjson.Field
		Reason        respjson.Field
		raw           string
	} `json:"-"`
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPrepaidCommitSegmentStartLedgerEntry() (v ContractV2CommitLedgerPrepaidCommitSegmentStartLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry() (v ContractV2CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPrepaidCommitRolloverLedgerEntry() (v ContractV2CommitLedgerPrepaidCommitRolloverLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPrepaidCommitExpirationLedgerEntry() (v ContractV2CommitLedgerPrepaidCommitExpirationLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPrepaidCommitCanceledLedgerEntry() (v ContractV2CommitLedgerPrepaidCommitCanceledLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPrepaidCommitCreditedLedgerEntry() (v ContractV2CommitLedgerPrepaidCommitCreditedLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry() (v ContractV2CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPostpaidCommitInitialBalanceLedgerEntry() (v ContractV2CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry() (v ContractV2CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPostpaidCommitRolloverLedgerEntry() (v ContractV2CommitLedgerPostpaidCommitRolloverLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPostpaidCommitTrueupLedgerEntry() (v ContractV2CommitLedgerPostpaidCommitTrueupLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPrepaidCommitManualLedgerEntry() (v ContractV2CommitLedgerPrepaidCommitManualLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPostpaidCommitManualLedgerEntry() (v ContractV2CommitLedgerPostpaidCommitManualLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CommitLedgerUnion) AsContractV2CommitLedgerPostpaidCommitExpirationLedgerEntry() (v ContractV2CommitLedgerPostpaidCommitExpirationLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContractV2CommitLedgerUnion) RawJSON() string { return u.JSON.raw }

func (r *ContractV2CommitLedgerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_SEGMENT_START".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPrepaidCommitSegmentStartLedgerEntry) RawJSON() string {
	return r.JSON.raw
}
func (r *ContractV2CommitLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) RawJSON() string {
	return r.JSON.raw
}
func (r *ContractV2CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64   `json:"amount,required"`
	NewContractID string    `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string    `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_ROLLOVER".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount        respjson.Field
		NewContractID respjson.Field
		SegmentID     respjson.Field
		Timestamp     respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPrepaidCommitRolloverLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_EXPIRATION".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPrepaidCommitExpirationLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_CANCELED".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPrepaidCommitCanceledLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_CREDITED".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPrepaidCommitCreditedLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_SEAT_BASED_ADJUSTMENT".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry) RawJSON() string {
	return r.JSON.raw
}
func (r *ContractV2CommitLedgerPrepaidCommitSeatBasedAdjustmentLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_INITIAL_BALANCE".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) RawJSON() string {
	return r.JSON.raw
}
func (r *ContractV2CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) RawJSON() string {
	return r.JSON.raw
}
func (r *ContractV2CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPostpaidCommitRolloverLedgerEntry struct {
	Amount        float64   `json:"amount,required"`
	NewContractID string    `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string    `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_ROLLOVER".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount        respjson.Field
		NewContractID respjson.Field
		SegmentID     respjson.Field
		Timestamp     respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPostpaidCommitRolloverLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitLedgerPostpaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_TRUEUP".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPostpaidCommitTrueupLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPrepaidCommitManualLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	Reason    string    `json:"reason,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "PREPAID_COMMIT_MANUAL".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Reason      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPrepaidCommitManualLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitLedgerPrepaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPostpaidCommitManualLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	Reason    string    `json:"reason,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_MANUAL".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Reason      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPostpaidCommitManualLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitLedgerPostpaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitLedgerPostpaidCommitExpirationLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "POSTPAID_COMMIT_EXPIRATION".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitLedgerPostpaidCommitExpirationLedgerEntry) RawJSON() string {
	return r.JSON.raw
}
func (r *ContractV2CommitLedgerPostpaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CommitRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommitID    respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CommitRolledOverFrom) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CommitRolledOverFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2Override struct {
	ID                    string                                `json:"id,required" format:"uuid"`
	StartingAt            time.Time                             `json:"starting_at,required" format:"date-time"`
	ApplicableProductTags []string                              `json:"applicable_product_tags"`
	EndingBefore          time.Time                             `json:"ending_before" format:"date-time"`
	Entitled              bool                                  `json:"entitled"`
	IsCommitSpecific      bool                                  `json:"is_commit_specific"`
	Multiplier            float64                               `json:"multiplier"`
	OverrideSpecifiers    []ContractV2OverrideOverrideSpecifier `json:"override_specifiers"`
	OverrideTiers         []OverrideTier                        `json:"override_tiers"`
	OverwriteRate         OverwriteRate                         `json:"overwrite_rate"`
	Priority              float64                               `json:"priority"`
	Product               ContractV2OverrideProduct             `json:"product"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	Target string `json:"target"`
	// Any of "OVERWRITE", "MULTIPLIER", "TIERED".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		StartingAt            respjson.Field
		ApplicableProductTags respjson.Field
		EndingBefore          respjson.Field
		Entitled              respjson.Field
		IsCommitSpecific      respjson.Field
		Multiplier            respjson.Field
		OverrideSpecifiers    respjson.Field
		OverrideTiers         respjson.Field
		OverwriteRate         respjson.Field
		Priority              respjson.Field
		Product               respjson.Field
		Target                respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2Override) RawJSON() string { return r.JSON.raw }
func (r *ContractV2Override) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2OverrideOverrideSpecifier struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency        string            `json:"billing_frequency"`
	CommitIDs               []string          `json:"commit_ids"`
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	ProductID               string            `json:"product_id" format:"uuid"`
	ProductTags             []string          `json:"product_tags"`
	RecurringCommitIDs      []string          `json:"recurring_commit_ids"`
	RecurringCreditIDs      []string          `json:"recurring_credit_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFrequency        respjson.Field
		CommitIDs               respjson.Field
		PresentationGroupValues respjson.Field
		PricingGroupValues      respjson.Field
		ProductID               respjson.Field
		ProductTags             respjson.Field
		RecurringCommitIDs      respjson.Field
		RecurringCreditIDs      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2OverrideOverrideSpecifier) RawJSON() string { return r.JSON.raw }
func (r *ContractV2OverrideOverrideSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2OverrideProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2OverrideProduct) RawJSON() string { return r.JSON.raw }
func (r *ContractV2OverrideProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2Transition struct {
	FromContractID string `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string `json:"to_contract_id,required" format:"uuid"`
	// Any of "SUPERSEDE", "RENEWAL".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromContractID respjson.Field
		ToContractID   respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2Transition) RawJSON() string { return r.JSON.raw }
func (r *ContractV2Transition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2UsageFilter struct {
	GroupKey    string   `json:"group_key,required"`
	GroupValues []string `json:"group_values,required"`
	// This will match contract starting_at value if usage filter is active from the
	// beginning of the contract.
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// This will match contract ending_before value if usage filter is active until the
	// end of the contract. It will be undefined if the contract is open-ended.
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupKey     respjson.Field
		GroupValues  respjson.Field
		StartingAt   respjson.Field
		EndingBefore respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2UsageFilter) RawJSON() string { return r.JSON.raw }
func (r *ContractV2UsageFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2UsageStatementSchedule struct {
	// Contract usage statements follow a selected cadence based on this date.
	BillingAnchorDate time.Time `json:"billing_anchor_date,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingAnchorDate respjson.Field
		Frequency         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2UsageStatementSchedule) RawJSON() string { return r.JSON.raw }
func (r *ContractV2UsageStatementSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2Credit struct {
	ID      string                  `json:"id,required" format:"uuid"`
	Product ContractV2CreditProduct `json:"product,required"`
	// Any of "CREDIT".
	Type string `json:"type,required"`
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
	Balance  float64                  `json:"balance"`
	Contract ContractV2CreditContract `json:"contract"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []ContractV2CreditLedgerUnion `json:"ledger"`
	Name   string                        `json:"name"`
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
	Specifiers []CommitSpecifier `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Product                 respjson.Field
		Type                    respjson.Field
		AccessSchedule          respjson.Field
		ApplicableContractIDs   respjson.Field
		ApplicableProductIDs    respjson.Field
		ApplicableProductTags   respjson.Field
		Balance                 respjson.Field
		Contract                respjson.Field
		CustomFields            respjson.Field
		Description             respjson.Field
		HierarchyConfiguration  respjson.Field
		Ledger                  respjson.Field
		Name                    respjson.Field
		NetsuiteSalesOrderID    respjson.Field
		Priority                respjson.Field
		SalesforceOpportunityID respjson.Field
		Specifiers              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2Credit) RawJSON() string { return r.JSON.raw }
func (r *ContractV2Credit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CreditProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CreditProduct) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CreditContract struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CreditContract) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CreditContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContractV2CreditLedgerUnion contains all possible properties and values from
// [ContractV2CreditLedgerCreditSegmentStartLedgerEntry],
// [ContractV2CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [ContractV2CreditLedgerCreditExpirationLedgerEntry],
// [ContractV2CreditLedgerCreditCanceledLedgerEntry],
// [ContractV2CreditLedgerCreditCreditedLedgerEntry],
// [ContractV2CreditLedgerCreditManualLedgerEntry],
// [ContractV2CreditLedgerCreditSeatBasedAdjustmentLedgerEntry].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ContractV2CreditLedgerUnion struct {
	Amount     float64   `json:"amount"`
	SegmentID  string    `json:"segment_id"`
	Timestamp  time.Time `json:"timestamp"`
	Type       string    `json:"type"`
	InvoiceID  string    `json:"invoice_id"`
	ContractID string    `json:"contract_id"`
	// This field is from variant [ContractV2CreditLedgerCreditManualLedgerEntry].
	Reason string `json:"reason"`
	JSON   struct {
		Amount     respjson.Field
		SegmentID  respjson.Field
		Timestamp  respjson.Field
		Type       respjson.Field
		InvoiceID  respjson.Field
		ContractID respjson.Field
		Reason     respjson.Field
		raw        string
	} `json:"-"`
}

func (u ContractV2CreditLedgerUnion) AsContractV2CreditLedgerCreditSegmentStartLedgerEntry() (v ContractV2CreditLedgerCreditSegmentStartLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CreditLedgerUnion) AsContractV2CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry() (v ContractV2CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CreditLedgerUnion) AsContractV2CreditLedgerCreditExpirationLedgerEntry() (v ContractV2CreditLedgerCreditExpirationLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CreditLedgerUnion) AsContractV2CreditLedgerCreditCanceledLedgerEntry() (v ContractV2CreditLedgerCreditCanceledLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CreditLedgerUnion) AsContractV2CreditLedgerCreditCreditedLedgerEntry() (v ContractV2CreditLedgerCreditCreditedLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CreditLedgerUnion) AsContractV2CreditLedgerCreditManualLedgerEntry() (v ContractV2CreditLedgerCreditManualLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContractV2CreditLedgerUnion) AsContractV2CreditLedgerCreditSeatBasedAdjustmentLedgerEntry() (v ContractV2CreditLedgerCreditSeatBasedAdjustmentLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContractV2CreditLedgerUnion) RawJSON() string { return u.JSON.raw }

func (r *ContractV2CreditLedgerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CreditLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_SEGMENT_START".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CreditLedgerCreditSegmentStartLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CreditLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_AUTOMATED_INVOICE_DEDUCTION".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) RawJSON() string {
	return r.JSON.raw
}
func (r *ContractV2CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CreditLedgerCreditExpirationLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_EXPIRATION".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CreditLedgerCreditExpirationLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CreditLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CreditLedgerCreditCanceledLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_CANCELED".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CreditLedgerCreditCanceledLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CreditLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CreditLedgerCreditCreditedLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_CREDITED".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CreditLedgerCreditCreditedLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CreditLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CreditLedgerCreditManualLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	Reason    string    `json:"reason,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_MANUAL".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Reason      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CreditLedgerCreditManualLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CreditLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2CreditLedgerCreditSeatBasedAdjustmentLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_SEAT_BASED_ADJUSTMENT".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CreditLedgerCreditSeatBasedAdjustmentLedgerEntry) RawJSON() string {
	return r.JSON.raw
}
func (r *ContractV2CreditLedgerCreditSeatBasedAdjustmentLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This field's availability is dependent on your client's configuration.
type ContractV2CustomerBillingProviderConfiguration struct {
	// ID of Customer's billing provider configuration.
	ID string `json:"id,required" format:"uuid"`
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace".
	BillingProvider string `json:"billing_provider,required"`
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod string `json:"delivery_method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		BillingProvider respjson.Field
		DeliveryMethod  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2CustomerBillingProviderConfiguration) RawJSON() string { return r.JSON.raw }
func (r *ContractV2CustomerBillingProviderConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	Credits bool `json:"credits,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commits     respjson.Field
		Credits     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2HasMore) RawJSON() string { return r.JSON.raw }
func (r *ContractV2HasMore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type ContractV2RecurringCommit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount ContractV2RecurringCommitAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration ContractV2RecurringCommitCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                          `json:"priority,required"`
	Product  ContractV2RecurringCommitProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                          `json:"applicable_product_tags"`
	Contract              ContractV2RecurringCommitContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount ContractV2RecurringCommitInvoiceAmount `json:"invoice_amount"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	RecurrenceFrequency string `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig RecurringCommitSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessAmount           respjson.Field
		CommitDuration         respjson.Field
		Priority               respjson.Field
		Product                respjson.Field
		RateType               respjson.Field
		StartingAt             respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Contract               respjson.Field
		Description            respjson.Field
		EndingBefore           respjson.Field
		HierarchyConfiguration respjson.Field
		InvoiceAmount          respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Proration              respjson.Field
		RecurrenceFrequency    respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		SubscriptionConfig     respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2RecurringCommit) RawJSON() string { return r.JSON.raw }
func (r *ContractV2RecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type ContractV2RecurringCommitAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
	Quantity     float64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		UnitPrice    respjson.Field
		Quantity     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2RecurringCommitAccessAmount) RawJSON() string { return r.JSON.raw }
func (r *ContractV2RecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time the created commits will be valid for
type ContractV2RecurringCommitCommitDuration struct {
	Value float64 `json:"value,required"`
	// Any of "PERIODS".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2RecurringCommitCommitDuration) RawJSON() string { return r.JSON.raw }
func (r *ContractV2RecurringCommitCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2RecurringCommitProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2RecurringCommitProduct) RawJSON() string { return r.JSON.raw }
func (r *ContractV2RecurringCommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2RecurringCommitContract struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2RecurringCommitContract) RawJSON() string { return r.JSON.raw }
func (r *ContractV2RecurringCommitContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount the customer should be billed for the commit. Not required.
type ContractV2RecurringCommitInvoiceAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64 `json:"quantity,required"`
	UnitPrice    float64 `json:"unit_price,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		Quantity     respjson.Field
		UnitPrice    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2RecurringCommitInvoiceAmount) RawJSON() string { return r.JSON.raw }
func (r *ContractV2RecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2RecurringCredit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount ContractV2RecurringCreditAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration ContractV2RecurringCreditCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                          `json:"priority,required"`
	Product  ContractV2RecurringCreditProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                          `json:"applicable_product_tags"`
	Contract              ContractV2RecurringCreditContract `json:"contract"`
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
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	RecurrenceFrequency string `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig RecurringCommitSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessAmount           respjson.Field
		CommitDuration         respjson.Field
		Priority               respjson.Field
		Product                respjson.Field
		RateType               respjson.Field
		StartingAt             respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Contract               respjson.Field
		Description            respjson.Field
		EndingBefore           respjson.Field
		HierarchyConfiguration respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Proration              respjson.Field
		RecurrenceFrequency    respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		SubscriptionConfig     respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2RecurringCredit) RawJSON() string { return r.JSON.raw }
func (r *ContractV2RecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type ContractV2RecurringCreditAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
	Quantity     float64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		UnitPrice    respjson.Field
		Quantity     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2RecurringCreditAccessAmount) RawJSON() string { return r.JSON.raw }
func (r *ContractV2RecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time the created commits will be valid for
type ContractV2RecurringCreditCommitDuration struct {
	Value float64 `json:"value,required"`
	// Any of "PERIODS".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2RecurringCreditCommitDuration) RawJSON() string { return r.JSON.raw }
func (r *ContractV2RecurringCreditCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2RecurringCreditProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2RecurringCreditProduct) RawJSON() string { return r.JSON.raw }
func (r *ContractV2RecurringCreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2RecurringCreditContract struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2RecurringCreditContract) RawJSON() string { return r.JSON.raw }
func (r *ContractV2RecurringCreditContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2ResellerRoyalty struct {
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType string                             `json:"reseller_type,required"`
	Segments     []ContractV2ResellerRoyaltySegment `json:"segments,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResellerType respjson.Field
		Segments     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2ResellerRoyalty) RawJSON() string { return r.JSON.raw }
func (r *ContractV2ResellerRoyalty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractV2ResellerRoyaltySegment struct {
	Fraction           float64 `json:"fraction,required"`
	NetsuiteResellerID string  `json:"netsuite_reseller_id,required"`
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType          string    `json:"reseller_type,required"`
	StartingAt            time.Time `json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  []string  `json:"applicable_product_ids"`
	ApplicableProductTags []string  `json:"applicable_product_tags"`
	AwsAccountNumber      string    `json:"aws_account_number"`
	AwsOfferID            string    `json:"aws_offer_id"`
	AwsPayerReferenceID   string    `json:"aws_payer_reference_id"`
	EndingBefore          time.Time `json:"ending_before" format:"date-time"`
	GcpAccountID          string    `json:"gcp_account_id"`
	GcpOfferID            string    `json:"gcp_offer_id"`
	ResellerContractValue float64   `json:"reseller_contract_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fraction              respjson.Field
		NetsuiteResellerID    respjson.Field
		ResellerType          respjson.Field
		StartingAt            respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		AwsAccountNumber      respjson.Field
		AwsOfferID            respjson.Field
		AwsPayerReferenceID   respjson.Field
		EndingBefore          respjson.Field
		GcpAccountID          respjson.Field
		GcpOfferID            respjson.Field
		ResellerContractValue respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractV2ResellerRoyaltySegment) RawJSON() string { return r.JSON.raw }
func (r *ContractV2ResellerRoyaltySegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	HierarchyConfiguration HierarchyConfigurationUnion `json:"hierarchy_configuration"`
	Name                   string                      `json:"name"`
	NetPaymentTermsDays    float64                     `json:"net_payment_terms_days"`
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
	//
	// Any of "ALL".
	ScheduledChargesOnUsageInvoices ContractWithoutAmendmentsScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	SpendThresholdConfiguration     SpendThresholdConfiguration                              `json:"spend_threshold_configuration"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue float64                              `json:"total_contract_value"`
	UsageFilter        ContractWithoutAmendmentsUsageFilter `json:"usage_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commits                              respjson.Field
		CreatedAt                            respjson.Field
		CreatedBy                            respjson.Field
		Overrides                            respjson.Field
		ScheduledCharges                     respjson.Field
		StartingAt                           respjson.Field
		Transitions                          respjson.Field
		UsageStatementSchedule               respjson.Field
		Credits                              respjson.Field
		Discounts                            respjson.Field
		EndingBefore                         respjson.Field
		HierarchyConfiguration               respjson.Field
		Name                                 respjson.Field
		NetPaymentTermsDays                  respjson.Field
		NetsuiteSalesOrderID                 respjson.Field
		PrepaidBalanceThresholdConfiguration respjson.Field
		ProfessionalServices                 respjson.Field
		RateCardID                           respjson.Field
		RecurringCommits                     respjson.Field
		RecurringCredits                     respjson.Field
		ResellerRoyalties                    respjson.Field
		SalesforceOpportunityID              respjson.Field
		ScheduledChargesOnUsageInvoices      respjson.Field
		SpendThresholdConfiguration          respjson.Field
		TotalContractValue                   respjson.Field
		UsageFilter                          respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendments) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendments) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsTransition struct {
	FromContractID string `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string `json:"to_contract_id,required" format:"uuid"`
	// Any of "SUPERSEDE", "RENEWAL".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromContractID respjson.Field
		ToContractID   respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsTransition) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsTransition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsUsageStatementSchedule struct {
	// Contract usage statements follow a selected cadence based on this date.
	BillingAnchorDate time.Time `json:"billing_anchor_date,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingAnchorDate respjson.Field
		Frequency         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsUsageStatementSchedule) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsUsageStatementSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsRecurringCommit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount ContractWithoutAmendmentsRecurringCommitAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration ContractWithoutAmendmentsRecurringCommitCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                         `json:"priority,required"`
	Product  ContractWithoutAmendmentsRecurringCommitProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                         `json:"applicable_product_tags"`
	Contract              ContractWithoutAmendmentsRecurringCommitContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount ContractWithoutAmendmentsRecurringCommitInvoiceAmount `json:"invoice_amount"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	RecurrenceFrequency string `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig RecurringCommitSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessAmount           respjson.Field
		CommitDuration         respjson.Field
		Priority               respjson.Field
		Product                respjson.Field
		RateType               respjson.Field
		StartingAt             respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Contract               respjson.Field
		Description            respjson.Field
		EndingBefore           respjson.Field
		HierarchyConfiguration respjson.Field
		InvoiceAmount          respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Proration              respjson.Field
		RecurrenceFrequency    respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		SubscriptionConfig     respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsRecurringCommit) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type ContractWithoutAmendmentsRecurringCommitAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
	Quantity     float64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		UnitPrice    respjson.Field
		Quantity     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsRecurringCommitAccessAmount) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time the created commits will be valid for
type ContractWithoutAmendmentsRecurringCommitCommitDuration struct {
	Value float64 `json:"value,required"`
	// Any of "PERIODS".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsRecurringCommitCommitDuration) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsRecurringCommitCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsRecurringCommitProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsRecurringCommitProduct) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsRecurringCommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsRecurringCommitContract struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsRecurringCommitContract) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsRecurringCommitContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount the customer should be billed for the commit. Not required.
type ContractWithoutAmendmentsRecurringCommitInvoiceAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64 `json:"quantity,required"`
	UnitPrice    float64 `json:"unit_price,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		Quantity     respjson.Field
		UnitPrice    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsRecurringCommitInvoiceAmount) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsRecurringCredit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount ContractWithoutAmendmentsRecurringCreditAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration ContractWithoutAmendmentsRecurringCreditCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                         `json:"priority,required"`
	Product  ContractWithoutAmendmentsRecurringCreditProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                         `json:"applicable_product_tags"`
	Contract              ContractWithoutAmendmentsRecurringCreditContract `json:"contract"`
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
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	RecurrenceFrequency string `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig RecurringCommitSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessAmount           respjson.Field
		CommitDuration         respjson.Field
		Priority               respjson.Field
		Product                respjson.Field
		RateType               respjson.Field
		StartingAt             respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Contract               respjson.Field
		Description            respjson.Field
		EndingBefore           respjson.Field
		HierarchyConfiguration respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Proration              respjson.Field
		RecurrenceFrequency    respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		SubscriptionConfig     respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsRecurringCredit) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type ContractWithoutAmendmentsRecurringCreditAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
	Quantity     float64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		UnitPrice    respjson.Field
		Quantity     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsRecurringCreditAccessAmount) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time the created commits will be valid for
type ContractWithoutAmendmentsRecurringCreditCommitDuration struct {
	Value float64 `json:"value,required"`
	// Any of "PERIODS".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsRecurringCreditCommitDuration) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsRecurringCreditCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsRecurringCreditProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsRecurringCreditProduct) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsRecurringCreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsRecurringCreditContract struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsRecurringCreditContract) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsRecurringCreditContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsResellerRoyalty struct {
	Fraction           float64 `json:"fraction,required"`
	NetsuiteResellerID string  `json:"netsuite_reseller_id,required"`
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType          string    `json:"reseller_type,required"`
	StartingAt            time.Time `json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  []string  `json:"applicable_product_ids"`
	ApplicableProductTags []string  `json:"applicable_product_tags"`
	AwsAccountNumber      string    `json:"aws_account_number"`
	AwsOfferID            string    `json:"aws_offer_id"`
	AwsPayerReferenceID   string    `json:"aws_payer_reference_id"`
	EndingBefore          time.Time `json:"ending_before" format:"date-time"`
	GcpAccountID          string    `json:"gcp_account_id"`
	GcpOfferID            string    `json:"gcp_offer_id"`
	ResellerContractValue float64   `json:"reseller_contract_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fraction              respjson.Field
		NetsuiteResellerID    respjson.Field
		ResellerType          respjson.Field
		StartingAt            respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		AwsAccountNumber      respjson.Field
		AwsOfferID            respjson.Field
		AwsPayerReferenceID   respjson.Field
		EndingBefore          respjson.Field
		GcpAccountID          respjson.Field
		GcpOfferID            respjson.Field
		ResellerContractValue respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsResellerRoyalty) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type ContractWithoutAmendmentsUsageFilter struct {
	Current BaseUsageFilter                              `json:"current,required"`
	Initial BaseUsageFilter                              `json:"initial,required"`
	Updates []ContractWithoutAmendmentsUsageFilterUpdate `json:"updates,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Initial     respjson.Field
		Updates     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsUsageFilter) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsUsageFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsUsageFilterUpdate struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupKey    respjson.Field
		GroupValues respjson.Field
		StartingAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractWithoutAmendmentsUsageFilterUpdate) RawJSON() string { return r.JSON.raw }
func (r *ContractWithoutAmendmentsUsageFilterUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Credit struct {
	ID      string        `json:"id,required" format:"uuid"`
	Product CreditProduct `json:"product,required"`
	// Any of "CREDIT".
	Type CreditType `json:"type,required"`
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
	Balance  float64        `json:"balance"`
	Contract CreditContract `json:"contract"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []CreditLedgerUnion `json:"ledger"`
	Name   string              `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// Any of "COMMIT_RATE", "LIST_RATE".
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
	UniquenessKey string `json:"uniqueness_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Product                 respjson.Field
		Type                    respjson.Field
		AccessSchedule          respjson.Field
		ApplicableContractIDs   respjson.Field
		ApplicableProductIDs    respjson.Field
		ApplicableProductTags   respjson.Field
		Balance                 respjson.Field
		Contract                respjson.Field
		CustomFields            respjson.Field
		Description             respjson.Field
		HierarchyConfiguration  respjson.Field
		Ledger                  respjson.Field
		Name                    respjson.Field
		NetsuiteSalesOrderID    respjson.Field
		Priority                respjson.Field
		RateType                respjson.Field
		SalesforceOpportunityID respjson.Field
		Specifiers              respjson.Field
		UniquenessKey           respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Credit) RawJSON() string { return r.JSON.raw }
func (r *Credit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditProduct) RawJSON() string { return r.JSON.raw }
func (r *CreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditType string

const (
	CreditTypeCredit CreditType = "CREDIT"
)

type CreditContract struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditContract) RawJSON() string { return r.JSON.raw }
func (r *CreditContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CreditLedgerUnion contains all possible properties and values from
// [CreditLedgerCreditSegmentStartLedgerEntry],
// [CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [CreditLedgerCreditExpirationLedgerEntry],
// [CreditLedgerCreditCanceledLedgerEntry],
// [CreditLedgerCreditCreditedLedgerEntry], [CreditLedgerCreditManualLedgerEntry],
// [CreditLedgerCreditSeatBasedAdjustmentLedgerEntry].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CreditLedgerUnion struct {
	Amount     float64   `json:"amount"`
	SegmentID  string    `json:"segment_id"`
	Timestamp  time.Time `json:"timestamp"`
	Type       string    `json:"type"`
	InvoiceID  string    `json:"invoice_id"`
	ContractID string    `json:"contract_id"`
	// This field is from variant [CreditLedgerCreditManualLedgerEntry].
	Reason string `json:"reason"`
	JSON   struct {
		Amount     respjson.Field
		SegmentID  respjson.Field
		Timestamp  respjson.Field
		Type       respjson.Field
		InvoiceID  respjson.Field
		ContractID respjson.Field
		Reason     respjson.Field
		raw        string
	} `json:"-"`
}

func (u CreditLedgerUnion) AsCreditLedgerCreditSegmentStartLedgerEntry() (v CreditLedgerCreditSegmentStartLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreditLedgerUnion) AsCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry() (v CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreditLedgerUnion) AsCreditLedgerCreditExpirationLedgerEntry() (v CreditLedgerCreditExpirationLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreditLedgerUnion) AsCreditLedgerCreditCanceledLedgerEntry() (v CreditLedgerCreditCanceledLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreditLedgerUnion) AsCreditLedgerCreditCreditedLedgerEntry() (v CreditLedgerCreditCreditedLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreditLedgerUnion) AsCreditLedgerCreditManualLedgerEntry() (v CreditLedgerCreditManualLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreditLedgerUnion) AsCreditLedgerCreditSeatBasedAdjustmentLedgerEntry() (v CreditLedgerCreditSeatBasedAdjustmentLedgerEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CreditLedgerUnion) RawJSON() string { return u.JSON.raw }

func (r *CreditLedgerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_SEGMENT_START".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditLedgerCreditSegmentStartLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CreditLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_AUTOMATED_INVOICE_DEDUCTION".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditLedgerCreditExpirationLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_EXPIRATION".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditLedgerCreditExpirationLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CreditLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditLedgerCreditCanceledLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_CANCELED".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditLedgerCreditCanceledLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CreditLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditLedgerCreditCreditedLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_CREDITED".
	Type       string `json:"type,required"`
	ContractID string `json:"contract_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ContractID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditLedgerCreditCreditedLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CreditLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditLedgerCreditManualLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	Reason    string    `json:"reason,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_MANUAL".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Reason      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditLedgerCreditManualLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CreditLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditLedgerCreditSeatBasedAdjustmentLedgerEntry struct {
	Amount    float64   `json:"amount,required"`
	SegmentID string    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "CREDIT_SEAT_BASED_ADJUSTMENT".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		SegmentID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditLedgerCreditSeatBasedAdjustmentLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CreditLedgerCreditSeatBasedAdjustmentLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditRateType string

const (
	CreditRateTypeCommitRate CreditRateType = "COMMIT_RATE"
	CreditRateTypeListRate   CreditRateType = "LIST_RATE"
)

type CreditTypeData struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditTypeData) RawJSON() string { return r.JSON.raw }
func (r *CreditTypeData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Discount struct {
	ID       string              `json:"id,required" format:"uuid"`
	Product  DiscountProduct     `json:"product,required"`
	Schedule SchedulePointInTime `json:"schedule,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Name         string            `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Product              respjson.Field
		Schedule             respjson.Field
		CustomFields         respjson.Field
		Name                 respjson.Field
		NetsuiteSalesOrderID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Discount) RawJSON() string { return r.JSON.raw }
func (r *Discount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DiscountProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiscountProduct) RawJSON() string { return r.JSON.raw }
func (r *DiscountProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	NotInValues []string `json:"not_in_values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InValues    respjson.Field
		NotInValues respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventTypeFilter) RawJSON() string { return r.JSON.raw }
func (r *EventTypeFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EventTypeFilter to a EventTypeFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EventTypeFilterParam.Overrides()
func (r EventTypeFilter) ToParam() EventTypeFilterParam {
	return param.Override[EventTypeFilterParam](json.RawMessage(r.RawJSON()))
}

// An optional filtering rule to match the 'event_type' property of an event.
type EventTypeFilterParam struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues []string `json:"in_values,omitzero"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues []string `json:"not_in_values,omitzero"`
	paramObj
}

func (r EventTypeFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow EventTypeFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventTypeFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HierarchyConfigurationUnion contains all possible properties and values from
// [HierarchyConfigurationParentHierarchyConfiguration],
// [HierarchyConfigurationChildHierarchyConfiguration].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type HierarchyConfigurationUnion struct {
	// This field is from variant [HierarchyConfigurationParentHierarchyConfiguration].
	Children []HierarchyConfigurationParentHierarchyConfigurationChild `json:"children"`
	// This field is from variant [HierarchyConfigurationChildHierarchyConfiguration].
	Parent HierarchyConfigurationChildHierarchyConfigurationParent `json:"parent"`
	JSON   struct {
		Children respjson.Field
		Parent   respjson.Field
		raw      string
	} `json:"-"`
}

func (u HierarchyConfigurationUnion) AsHierarchyConfigurationParentHierarchyConfiguration() (v HierarchyConfigurationParentHierarchyConfiguration) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HierarchyConfigurationUnion) AsHierarchyConfigurationChildHierarchyConfiguration() (v HierarchyConfigurationChildHierarchyConfiguration) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u HierarchyConfigurationUnion) RawJSON() string { return u.JSON.raw }

func (r *HierarchyConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HierarchyConfigurationParentHierarchyConfiguration struct {
	// List of contracts that belong to this parent.
	Children []HierarchyConfigurationParentHierarchyConfigurationChild `json:"children,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HierarchyConfigurationParentHierarchyConfiguration) RawJSON() string { return r.JSON.raw }
func (r *HierarchyConfigurationParentHierarchyConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HierarchyConfigurationParentHierarchyConfigurationChild struct {
	ContractID string `json:"contract_id,required" format:"uuid"`
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContractID  respjson.Field
		CustomerID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HierarchyConfigurationParentHierarchyConfigurationChild) RawJSON() string { return r.JSON.raw }
func (r *HierarchyConfigurationParentHierarchyConfigurationChild) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HierarchyConfigurationChildHierarchyConfiguration struct {
	// The single parent contract/customer for this child.
	Parent HierarchyConfigurationChildHierarchyConfigurationParent `json:"parent,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parent      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HierarchyConfigurationChildHierarchyConfiguration) RawJSON() string { return r.JSON.raw }
func (r *HierarchyConfigurationChildHierarchyConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The single parent contract/customer for this child.
type HierarchyConfigurationChildHierarchyConfigurationParent struct {
	ContractID string `json:"contract_id,required" format:"uuid"`
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContractID  respjson.Field
		CustomerID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HierarchyConfigurationChildHierarchyConfigurationParent) RawJSON() string { return r.JSON.raw }
func (r *HierarchyConfigurationChildHierarchyConfigurationParent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ID struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ID) RawJSON() string { return r.JSON.raw }
func (r *ID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ID to a IDParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IDParam.Overrides()
func (r ID) ToParam() IDParam {
	return param.Override[IDParam](json.RawMessage(r.RawJSON()))
}

// The property ID is required.
type IDParam struct {
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r IDParam) MarshalJSON() (data []byte, err error) {
	type shadow IDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	OverwriteRate      OverwriteRate               `json:"overwrite_rate"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64         `json:"price"`
	Priority float64         `json:"priority"`
	Product  OverrideProduct `json:"product"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM".
	RateType OverrideRateType `json:"rate_type"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	Target OverrideTarget `json:"target"`
	// Only set for TIERED rate_type.
	Tiers []Tier `json:"tiers"`
	// Any of "OVERWRITE", "MULTIPLIER", "TIERED".
	Type OverrideType `json:"type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	Value map[string]any `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		StartingAt            respjson.Field
		ApplicableProductTags respjson.Field
		CreditType            respjson.Field
		EndingBefore          respjson.Field
		Entitled              respjson.Field
		IsCommitSpecific      respjson.Field
		IsProrated            respjson.Field
		Multiplier            respjson.Field
		OverrideSpecifiers    respjson.Field
		OverrideTiers         respjson.Field
		OverwriteRate         respjson.Field
		Price                 respjson.Field
		Priority              respjson.Field
		Product               respjson.Field
		Quantity              respjson.Field
		RateType              respjson.Field
		Target                respjson.Field
		Tiers                 respjson.Field
		Type                  respjson.Field
		Value                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Override) RawJSON() string { return r.JSON.raw }
func (r *Override) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideOverrideSpecifier struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency        string            `json:"billing_frequency"`
	CommitIDs               []string          `json:"commit_ids"`
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	ProductID               string            `json:"product_id" format:"uuid"`
	ProductTags             []string          `json:"product_tags"`
	RecurringCommitIDs      []string          `json:"recurring_commit_ids"`
	RecurringCreditIDs      []string          `json:"recurring_credit_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFrequency        respjson.Field
		CommitIDs               respjson.Field
		PresentationGroupValues respjson.Field
		PricingGroupValues      respjson.Field
		ProductID               respjson.Field
		ProductTags             respjson.Field
		RecurringCommitIDs      respjson.Field
		RecurringCreditIDs      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OverrideOverrideSpecifier) RawJSON() string { return r.JSON.raw }
func (r *OverrideOverrideSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OverrideProduct) RawJSON() string { return r.JSON.raw }
func (r *OverrideProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideRateType string

const (
	OverrideRateTypeFlat         OverrideRateType = "FLAT"
	OverrideRateTypePercentage   OverrideRateType = "PERCENTAGE"
	OverrideRateTypeSubscription OverrideRateType = "SUBSCRIPTION"
	OverrideRateTypeTiered       OverrideRateType = "TIERED"
	OverrideRateTypeCustom       OverrideRateType = "CUSTOM"
)

type OverrideTarget string

const (
	OverrideTargetCommitRate OverrideTarget = "COMMIT_RATE"
	OverrideTargetListRate   OverrideTarget = "LIST_RATE"
)

type OverrideType string

const (
	OverrideTypeOverwrite  OverrideType = "OVERWRITE"
	OverrideTypeMultiplier OverrideType = "MULTIPLIER"
	OverrideTypeTiered     OverrideType = "TIERED"
)

type OverrideTier struct {
	Multiplier float64 `json:"multiplier,required"`
	Size       float64 `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Multiplier  respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OverrideTier) RawJSON() string { return r.JSON.raw }
func (r *OverrideTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OverwriteRate struct {
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM".
	RateType   OverwriteRateRateType `json:"rate_type,required"`
	CreditType CreditTypeData        `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]any `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []Tier `json:"tiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RateType    respjson.Field
		CreditType  respjson.Field
		CustomRate  respjson.Field
		IsProrated  respjson.Field
		Price       respjson.Field
		Quantity    respjson.Field
		Tiers       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OverwriteRate) RawJSON() string { return r.JSON.raw }
func (r *OverwriteRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OverwriteRateRateType string

const (
	OverwriteRateRateTypeFlat         OverwriteRateRateType = "FLAT"
	OverwriteRateRateTypePercentage   OverwriteRateRateType = "PERCENTAGE"
	OverwriteRateRateTypeSubscription OverwriteRateRateType = "SUBSCRIPTION"
	OverwriteRateRateTypeTiered       OverwriteRateRateType = "TIERED"
	OverwriteRateRateTypeCustom       OverwriteRateRateType = "CUSTOM"
)

type PaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	//
	// Any of "NONE", "STRIPE", "EXTERNAL".
	PaymentGateType PaymentGateConfigPaymentGateType `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig PaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig PaymentGateConfigStripeConfig `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	//
	// Any of "NONE", "STRIPE", "ANROK", "PRECALCULATED".
	TaxType PaymentGateConfigTaxType `json:"tax_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentGateType        respjson.Field
		PrecalculatedTaxConfig respjson.Field
		StripeConfig           respjson.Field
		TaxType                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentGateConfig) RawJSON() string { return r.JSON.raw }
func (r *PaymentGateConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PaymentGateConfig to a PaymentGateConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PaymentGateConfigParam.Overrides()
func (r PaymentGateConfig) ToParam() PaymentGateConfigParam {
	return param.Override[PaymentGateConfigParam](json.RawMessage(r.RawJSON()))
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

// Only applicable if using PRECALCULATED as your tax type.
type PaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName string `json:"tax_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TaxAmount   respjson.Field
		TaxName     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentGateConfigPrecalculatedTaxConfig) RawJSON() string { return r.JSON.raw }
func (r *PaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only applicable if using STRIPE as your payment gate type.
type PaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	//
	// Any of "INVOICE", "PAYMENT_INTENT".
	PaymentType string `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string `json:"invoice_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentType     respjson.Field
		InvoiceMetadata respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentGateConfigStripeConfig) RawJSON() string { return r.JSON.raw }
func (r *PaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// The property PaymentGateType is required.
type PaymentGateConfigParam struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	//
	// Any of "NONE", "STRIPE", "EXTERNAL".
	PaymentGateType PaymentGateConfigPaymentGateType `json:"payment_gate_type,omitzero,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig PaymentGateConfigPrecalculatedTaxConfigParam `json:"precalculated_tax_config,omitzero"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig PaymentGateConfigStripeConfigParam `json:"stripe_config,omitzero"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	//
	// Any of "NONE", "STRIPE", "ANROK", "PRECALCULATED".
	TaxType PaymentGateConfigTaxType `json:"tax_type,omitzero"`
	paramObj
}

func (r PaymentGateConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow PaymentGateConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentGateConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only applicable if using PRECALCULATED as your tax type.
//
// The property TaxAmount is required.
type PaymentGateConfigPrecalculatedTaxConfigParam struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Opt[string] `json:"tax_name,omitzero"`
	paramObj
}

func (r PaymentGateConfigPrecalculatedTaxConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow PaymentGateConfigPrecalculatedTaxConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentGateConfigPrecalculatedTaxConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only applicable if using STRIPE as your payment gate type.
//
// The property PaymentType is required.
type PaymentGateConfigStripeConfigParam struct {
	// If left blank, will default to INVOICE
	//
	// Any of "INVOICE", "PAYMENT_INTENT".
	PaymentType string `json:"payment_type,omitzero,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string `json:"invoice_metadata,omitzero"`
	paramObj
}

func (r PaymentGateConfigStripeConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow PaymentGateConfigStripeConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentGateConfigStripeConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PaymentGateConfigStripeConfigParam](
		"payment_type", "INVOICE", "PAYMENT_INTENT",
	)
}

type PaymentGateConfigV2 struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	//
	// Any of "NONE", "STRIPE", "EXTERNAL".
	PaymentGateType PaymentGateConfigV2PaymentGateType `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig PaymentGateConfigV2PrecalculatedTaxConfig `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig PaymentGateConfigV2StripeConfig `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	//
	// Any of "NONE", "STRIPE", "ANROK", "PRECALCULATED".
	TaxType PaymentGateConfigV2TaxType `json:"tax_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentGateType        respjson.Field
		PrecalculatedTaxConfig respjson.Field
		StripeConfig           respjson.Field
		TaxType                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentGateConfigV2) RawJSON() string { return r.JSON.raw }
func (r *PaymentGateConfigV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PaymentGateConfigV2 to a PaymentGateConfigV2Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PaymentGateConfigV2Param.Overrides()
func (r PaymentGateConfigV2) ToParam() PaymentGateConfigV2Param {
	return param.Override[PaymentGateConfigV2Param](json.RawMessage(r.RawJSON()))
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

// Only applicable if using PRECALCULATED as your tax type.
type PaymentGateConfigV2PrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName string `json:"tax_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TaxAmount   respjson.Field
		TaxName     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentGateConfigV2PrecalculatedTaxConfig) RawJSON() string { return r.JSON.raw }
func (r *PaymentGateConfigV2PrecalculatedTaxConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only applicable if using STRIPE as your payment gateway type.
type PaymentGateConfigV2StripeConfig struct {
	// If left blank, will default to INVOICE
	//
	// Any of "INVOICE", "PAYMENT_INTENT".
	PaymentType string `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string `json:"invoice_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentType     respjson.Field
		InvoiceMetadata respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentGateConfigV2StripeConfig) RawJSON() string { return r.JSON.raw }
func (r *PaymentGateConfigV2StripeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// The property PaymentGateType is required.
type PaymentGateConfigV2Param struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	//
	// Any of "NONE", "STRIPE", "EXTERNAL".
	PaymentGateType PaymentGateConfigV2PaymentGateType `json:"payment_gate_type,omitzero,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig PaymentGateConfigV2PrecalculatedTaxConfigParam `json:"precalculated_tax_config,omitzero"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig PaymentGateConfigV2StripeConfigParam `json:"stripe_config,omitzero"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	//
	// Any of "NONE", "STRIPE", "ANROK", "PRECALCULATED".
	TaxType PaymentGateConfigV2TaxType `json:"tax_type,omitzero"`
	paramObj
}

func (r PaymentGateConfigV2Param) MarshalJSON() (data []byte, err error) {
	type shadow PaymentGateConfigV2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentGateConfigV2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only applicable if using PRECALCULATED as your tax type.
//
// The property TaxAmount is required.
type PaymentGateConfigV2PrecalculatedTaxConfigParam struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Opt[string] `json:"tax_name,omitzero"`
	paramObj
}

func (r PaymentGateConfigV2PrecalculatedTaxConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow PaymentGateConfigV2PrecalculatedTaxConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentGateConfigV2PrecalculatedTaxConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only applicable if using STRIPE as your payment gateway type.
//
// The property PaymentType is required.
type PaymentGateConfigV2StripeConfigParam struct {
	// If left blank, will default to INVOICE
	//
	// Any of "INVOICE", "PAYMENT_INTENT".
	PaymentType string `json:"payment_type,omitzero,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string `json:"invoice_metadata,omitzero"`
	paramObj
}

func (r PaymentGateConfigV2StripeConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow PaymentGateConfigV2StripeConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentGateConfigV2StripeConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PaymentGateConfigV2StripeConfigParam](
		"payment_type", "INVOICE", "PAYMENT_INTENT",
	)
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
	CustomCreditTypeID string `json:"custom_credit_type_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit             respjson.Field
		IsEnabled          respjson.Field
		PaymentGateConfig  respjson.Field
		RechargeToAmount   respjson.Field
		ThresholdAmount    respjson.Field
		CustomCreditTypeID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrepaidBalanceThresholdConfiguration) RawJSON() string { return r.JSON.raw }
func (r *PrepaidBalanceThresholdConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PrepaidBalanceThresholdConfiguration to a
// PrepaidBalanceThresholdConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PrepaidBalanceThresholdConfigurationParam.Overrides()
func (r PrepaidBalanceThresholdConfiguration) ToParam() PrepaidBalanceThresholdConfigurationParam {
	return param.Override[PrepaidBalanceThresholdConfigurationParam](json.RawMessage(r.RawJSON()))
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
	Specifiers []CommitSpecifierInput `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		Specifiers            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
	BaseThresholdCommit
}

// Returns the unmodified JSON received from the API
func (r PrepaidBalanceThresholdConfigurationCommit) RawJSON() string { return r.JSON.raw }
func (r *PrepaidBalanceThresholdConfigurationCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Commit, IsEnabled, PaymentGateConfig, RechargeToAmount,
// ThresholdAmount are required.
type PrepaidBalanceThresholdConfigurationParam struct {
	Commit PrepaidBalanceThresholdConfigurationCommitParam `json:"commit,omitzero,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                   `json:"is_enabled,required"`
	PaymentGateConfig PaymentGateConfigParam `json:"payment_gate_config,omitzero,required"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount float64 `json:"recharge_to_amount,required"`
	// Specify the threshold amount for the contract. Each time the contract's prepaid
	// balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount,required"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID param.Opt[string] `json:"custom_credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r PrepaidBalanceThresholdConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow PrepaidBalanceThresholdConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PrepaidBalanceThresholdConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PrepaidBalanceThresholdConfigurationCommitParam struct {
	// Which products the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []CommitSpecifierInputParam `json:"specifiers,omitzero"`
	BaseThresholdCommitParam
}

func (r PrepaidBalanceThresholdConfigurationCommitParam) MarshalJSON() (data []byte, err error) {
	type shadow PrepaidBalanceThresholdConfigurationCommitParam
	return param.MarshalObject(r, (*shadow)(&r))
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
	CustomCreditTypeID string `json:"custom_credit_type_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit             respjson.Field
		IsEnabled          respjson.Field
		PaymentGateConfig  respjson.Field
		RechargeToAmount   respjson.Field
		ThresholdAmount    respjson.Field
		CustomCreditTypeID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrepaidBalanceThresholdConfigurationV2) RawJSON() string { return r.JSON.raw }
func (r *PrepaidBalanceThresholdConfigurationV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PrepaidBalanceThresholdConfigurationV2 to a
// PrepaidBalanceThresholdConfigurationV2Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PrepaidBalanceThresholdConfigurationV2Param.Overrides()
func (r PrepaidBalanceThresholdConfigurationV2) ToParam() PrepaidBalanceThresholdConfigurationV2Param {
	return param.Override[PrepaidBalanceThresholdConfigurationV2Param](json.RawMessage(r.RawJSON()))
}

type PrepaidBalanceThresholdConfigurationV2Commit struct {
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
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []CommitSpecifierInput `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		Specifiers            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
	UpdateBaseThresholdCommit
}

// Returns the unmodified JSON received from the API
func (r PrepaidBalanceThresholdConfigurationV2Commit) RawJSON() string { return r.JSON.raw }
func (r *PrepaidBalanceThresholdConfigurationV2Commit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Commit, IsEnabled, PaymentGateConfig, RechargeToAmount,
// ThresholdAmount are required.
type PrepaidBalanceThresholdConfigurationV2Param struct {
	Commit PrepaidBalanceThresholdConfigurationV2CommitParam `json:"commit,omitzero,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                     `json:"is_enabled,required"`
	PaymentGateConfig PaymentGateConfigV2Param `json:"payment_gate_config,omitzero,required"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount float64 `json:"recharge_to_amount,required"`
	// Specify the threshold amount for the contract. Each time the contract's balance
	// lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount,required"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID param.Opt[string] `json:"custom_credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r PrepaidBalanceThresholdConfigurationV2Param) MarshalJSON() (data []byte, err error) {
	type shadow PrepaidBalanceThresholdConfigurationV2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PrepaidBalanceThresholdConfigurationV2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PrepaidBalanceThresholdConfigurationV2CommitParam struct {
	// Which products the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []CommitSpecifierInputParam `json:"specifiers,omitzero"`
	UpdateBaseThresholdCommitParam
}

func (r PrepaidBalanceThresholdConfigurationV2CommitParam) MarshalJSON() (data []byte, err error) {
	type shadow PrepaidBalanceThresholdConfigurationV2CommitParam
	return param.MarshalObject(r, (*shadow)(&r))
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
	NotInValues []string `json:"not_in_values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Exists      respjson.Field
		InValues    respjson.Field
		NotInValues respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyFilter) RawJSON() string { return r.JSON.raw }
func (r *PropertyFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PropertyFilter to a PropertyFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PropertyFilterParam.Overrides()
func (r PropertyFilter) ToParam() PropertyFilterParam {
	return param.Override[PropertyFilterParam](json.RawMessage(r.RawJSON()))
}

// The property Name is required.
type PropertyFilterParam struct {
	// The name of the event property.
	Name string `json:"name,required"`
	// Determines whether the property must exist in the event. If true, only events
	// with this property will pass the filter. If false, only events without this
	// property will pass the filter. If null or omitted, the existence of the property
	// is optional.
	Exists param.Opt[bool] `json:"exists,omitzero"`
	// Specifies the allowed values for the property to match an event. An event will
	// pass the filter only if its property value is included in this list. If
	// undefined, all property values will pass the filter. Must be non-empty if
	// present.
	InValues []string `json:"in_values,omitzero"`
	// Specifies the values that prevent an event from matching the filter. An event
	// will not pass the filter if its property value is included in this list. If null
	// or empty, all property values will pass the filter. Must be non-empty if
	// present.
	NotInValues []string `json:"not_in_values,omitzero"`
	paramObj
}

func (r PropertyFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PropertyFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PropertyFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	UnitPrice float64 `json:"unit_price,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		MaxAmount            respjson.Field
		ProductID            respjson.Field
		Quantity             respjson.Field
		UnitPrice            respjson.Field
		CustomFields         respjson.Field
		Description          respjson.Field
		NetsuiteSalesOrderID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProService) RawJSON() string { return r.JSON.raw }
func (r *ProService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Rate struct {
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "CUSTOM", "TIERED".
	RateType   RateRateType   `json:"rate_type,required"`
	CreditType CreditTypeData `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]any `json:"custom_rate"`
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
	UseListPrices bool `json:"use_list_prices"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RateType           respjson.Field
		CreditType         respjson.Field
		CustomRate         respjson.Field
		IsProrated         respjson.Field
		Price              respjson.Field
		PricingGroupValues respjson.Field
		Quantity           respjson.Field
		Tiers              respjson.Field
		UseListPrices      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Rate) RawJSON() string { return r.JSON.raw }
func (r *Rate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateRateType string

const (
	RateRateTypeFlat         RateRateType = "FLAT"
	RateRateTypePercentage   RateRateType = "PERCENTAGE"
	RateRateTypeSubscription RateRateType = "SUBSCRIPTION"
	RateRateTypeCustom       RateRateType = "CUSTOM"
	RateRateTypeTiered       RateRateType = "TIERED"
)

type RecurringCommitSubscriptionConfig struct {
	// Any of "INDIVIDUAL", "POOLED".
	Allocation              RecurringCommitSubscriptionConfigAllocation              `json:"allocation,required"`
	ApplySeatIncreaseConfig RecurringCommitSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,required"`
	SubscriptionID          string                                                   `json:"subscription_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allocation              respjson.Field
		ApplySeatIncreaseConfig respjson.Field
		SubscriptionID          respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecurringCommitSubscriptionConfig) RawJSON() string { return r.JSON.raw }
func (r *RecurringCommitSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurringCommitSubscriptionConfigAllocation string

const (
	RecurringCommitSubscriptionConfigAllocationIndividual RecurringCommitSubscriptionConfigAllocation = "INDIVIDUAL"
	RecurringCommitSubscriptionConfigAllocationPooled     RecurringCommitSubscriptionConfigAllocation = "POOLED"
)

type RecurringCommitSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsProrated  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecurringCommitSubscriptionConfigApplySeatIncreaseConfig) RawJSON() string { return r.JSON.raw }
func (r *RecurringCommitSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduledCharge struct {
	ID         string                 `json:"id,required" format:"uuid"`
	Product    ScheduledChargeProduct `json:"product,required"`
	Schedule   SchedulePointInTime    `json:"schedule,required"`
	ArchivedAt time.Time              `json:"archived_at" format:"date-time"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	// displayed on invoices
	Name string `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Product              respjson.Field
		Schedule             respjson.Field
		ArchivedAt           respjson.Field
		CustomFields         respjson.Field
		Name                 respjson.Field
		NetsuiteSalesOrderID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduledCharge) RawJSON() string { return r.JSON.raw }
func (r *ScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduledChargeProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduledChargeProduct) RawJSON() string { return r.JSON.raw }
func (r *ScheduledChargeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduleDuration struct {
	ScheduleItems []ScheduleDurationScheduleItem `json:"schedule_items,required"`
	CreditType    CreditTypeData                 `json:"credit_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScheduleItems respjson.Field
		CreditType    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleDuration) RawJSON() string { return r.JSON.raw }
func (r *ScheduleDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduleDurationScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Amount       respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleDurationScheduleItem) RawJSON() string { return r.JSON.raw }
func (r *ScheduleDurationScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SchedulePointInTime struct {
	CreditType CreditTypeData `json:"credit_type"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice  bool                              `json:"do_not_invoice"`
	ScheduleItems []SchedulePointInTimeScheduleItem `json:"schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType    respjson.Field
		DoNotInvoice  respjson.Field
		ScheduleItems respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchedulePointInTime) RawJSON() string { return r.JSON.raw }
func (r *SchedulePointInTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SchedulePointInTimeScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	InvoiceID string    `json:"invoice_id,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Quantity    respjson.Field
		Timestamp   respjson.Field
		UnitPrice   respjson.Field
		InvoiceID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchedulePointInTimeScheduleItem) RawJSON() string { return r.JSON.raw }
func (r *SchedulePointInTimeScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	ThresholdAmount float64 `json:"threshold_amount,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit            respjson.Field
		IsEnabled         respjson.Field
		PaymentGateConfig respjson.Field
		ThresholdAmount   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendThresholdConfiguration) RawJSON() string { return r.JSON.raw }
func (r *SpendThresholdConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SpendThresholdConfiguration to a
// SpendThresholdConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SpendThresholdConfigurationParam.Overrides()
func (r SpendThresholdConfiguration) ToParam() SpendThresholdConfigurationParam {
	return param.Override[SpendThresholdConfigurationParam](json.RawMessage(r.RawJSON()))
}

// The properties Commit, IsEnabled, PaymentGateConfig, ThresholdAmount are
// required.
type SpendThresholdConfigurationParam struct {
	Commit BaseThresholdCommitParam `json:"commit,omitzero,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                   `json:"is_enabled,required"`
	PaymentGateConfig PaymentGateConfigParam `json:"payment_gate_config,omitzero,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount,required"`
	paramObj
}

func (r SpendThresholdConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow SpendThresholdConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendThresholdConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendThresholdConfigurationV2 struct {
	Commit UpdateBaseThresholdCommit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                `json:"is_enabled,required"`
	PaymentGateConfig PaymentGateConfigV2 `json:"payment_gate_config,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit            respjson.Field
		IsEnabled         respjson.Field
		PaymentGateConfig respjson.Field
		ThresholdAmount   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendThresholdConfigurationV2) RawJSON() string { return r.JSON.raw }
func (r *SpendThresholdConfigurationV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SpendThresholdConfigurationV2 to a
// SpendThresholdConfigurationV2Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SpendThresholdConfigurationV2Param.Overrides()
func (r SpendThresholdConfigurationV2) ToParam() SpendThresholdConfigurationV2Param {
	return param.Override[SpendThresholdConfigurationV2Param](json.RawMessage(r.RawJSON()))
}

// The properties Commit, IsEnabled, PaymentGateConfig, ThresholdAmount are
// required.
type SpendThresholdConfigurationV2Param struct {
	Commit UpdateBaseThresholdCommitParam `json:"commit,omitzero,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                     `json:"is_enabled,required"`
	PaymentGateConfig PaymentGateConfigV2Param `json:"payment_gate_config,omitzero,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount,required"`
	paramObj
}

func (r SpendThresholdConfigurationV2Param) MarshalJSON() (data []byte, err error) {
	type shadow SpendThresholdConfigurationV2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendThresholdConfigurationV2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Subscription struct {
	// Any of "ADVANCE", "ARREARS".
	CollectionSchedule SubscriptionCollectionSchedule `json:"collection_schedule,required"`
	Proration          SubscriptionProration          `json:"proration,required"`
	// Determines how the subscription's quantity is controlled. Defaults to
	// QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified
	// directly on the subscription. `initial_quantity` must be provided with this
	// option. Compatible with recurring commits/credits that use POOLED allocation.
	//
	// Any of "SEAT_BASED", "QUANTITY_ONLY".
	QuantityManagementMode SubscriptionQuantityManagementMode `json:"quantity_management_mode,required"`
	// List of quantity schedule items for the subscription. Only includes the current
	// quantity and future quantity changes.
	QuantitySchedule []SubscriptionQuantitySchedule `json:"quantity_schedule,required"`
	StartingAt       time.Time                      `json:"starting_at,required" format:"date-time"`
	SubscriptionRate SubscriptionSubscriptionRate   `json:"subscription_rate,required"`
	ID               string                         `json:"id" format:"uuid"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields     map[string]string `json:"custom_fields"`
	Description      string            `json:"description"`
	EndingBefore     time.Time         `json:"ending_before" format:"date-time"`
	FiatCreditTypeID string            `json:"fiat_credit_type_id" format:"uuid"`
	Name             string            `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionSchedule     respjson.Field
		Proration              respjson.Field
		QuantityManagementMode respjson.Field
		QuantitySchedule       respjson.Field
		StartingAt             respjson.Field
		SubscriptionRate       respjson.Field
		ID                     respjson.Field
		CustomFields           respjson.Field
		Description            respjson.Field
		EndingBefore           respjson.Field
		FiatCreditTypeID       respjson.Field
		Name                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subscription) RawJSON() string { return r.JSON.raw }
func (r *Subscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionCollectionSchedule string

const (
	SubscriptionCollectionScheduleAdvance SubscriptionCollectionSchedule = "ADVANCE"
	SubscriptionCollectionScheduleArrears SubscriptionCollectionSchedule = "ARREARS"
)

type SubscriptionProration struct {
	// Any of "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE".
	InvoiceBehavior string `json:"invoice_behavior,required"`
	IsProrated      bool   `json:"is_prorated,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InvoiceBehavior respjson.Field
		IsProrated      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionProration) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionProration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determines how the subscription's quantity is controlled. Defaults to
// QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified
// directly on the subscription. `initial_quantity` must be provided with this
// option. Compatible with recurring commits/credits that use POOLED allocation.
type SubscriptionQuantityManagementMode string

const (
	SubscriptionQuantityManagementModeSeatBased    SubscriptionQuantityManagementMode = "SEAT_BASED"
	SubscriptionQuantityManagementModeQuantityOnly SubscriptionQuantityManagementMode = "QUANTITY_ONLY"
)

type SubscriptionQuantitySchedule struct {
	Quantity     float64   `json:"quantity,required"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity     respjson.Field
		StartingAt   respjson.Field
		EndingBefore respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionQuantitySchedule) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionQuantitySchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionSubscriptionRate struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string                              `json:"billing_frequency,required"`
	Product          SubscriptionSubscriptionRateProduct `json:"product,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFrequency respjson.Field
		Product          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionSubscriptionRate) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionSubscriptionRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionSubscriptionRateProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionSubscriptionRateProduct) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionSubscriptionRateProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Tier struct {
	Price float64 `json:"price,required"`
	Size  float64 `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Price       respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tier) RawJSON() string { return r.JSON.raw }
func (r *Tier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Tier to a TierParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TierParam.Overrides()
func (r Tier) ToParam() TierParam {
	return param.Override[TierParam](json.RawMessage(r.RawJSON()))
}

// The property Price is required.
type TierParam struct {
	Price float64            `json:"price,required"`
	Size  param.Opt[float64] `json:"size,omitzero"`
	paramObj
}

func (r TierParam) MarshalJSON() (data []byte, err error) {
	type shadow TierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateBaseThresholdCommit struct {
	Description string `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string `json:"name"`
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID string `json:"product_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		ProductID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UpdateBaseThresholdCommit) RawJSON() string { return r.JSON.raw }
func (r *UpdateBaseThresholdCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UpdateBaseThresholdCommit to a
// UpdateBaseThresholdCommitParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UpdateBaseThresholdCommitParam.Overrides()
func (r UpdateBaseThresholdCommit) ToParam() UpdateBaseThresholdCommitParam {
	return param.Override[UpdateBaseThresholdCommitParam](json.RawMessage(r.RawJSON()))
}

type UpdateBaseThresholdCommitParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name param.Opt[string] `json:"name,omitzero"`
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID param.Opt[string] `json:"product_id,omitzero"`
	paramObj
}

func (r UpdateBaseThresholdCommitParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateBaseThresholdCommitParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateBaseThresholdCommitParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
