// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/tidwall/gjson"
)

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

func (r Commit) ImplementsV1ContractListBalancesResponseData() {}

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

// Optional configuration for commit hierarchy access control
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
	HierarchyConfiguration ContractWithoutAmendmentsHierarchyConfiguration `json:"hierarchy_configuration"`
	Name                   string                                          `json:"name"`
	NetPaymentTermsDays    float64                                         `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID                 string                                                        `json:"netsuite_sales_order_id"`
	PrepaidBalanceThresholdConfiguration ContractWithoutAmendmentsPrepaidBalanceThresholdConfiguration `json:"prepaid_balance_threshold_configuration"`
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
	SpendThresholdConfiguration     ContractWithoutAmendmentsSpendThresholdConfiguration     `json:"spend_threshold_configuration"`
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

// Either a **parent** configuration with a list of children or a **child**
// configuration with a single parent.
type ContractWithoutAmendmentsHierarchyConfiguration struct {
	// This field can have the runtime type of
	// [[]ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationChild].
	Children interface{} `json:"children"`
	// This field can have the runtime type of
	// [ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationParent].
	Parent interface{}                                         `json:"parent"`
	JSON   contractWithoutAmendmentsHierarchyConfigurationJSON `json:"-"`
	union  ContractWithoutAmendmentsHierarchyConfigurationUnion
}

// contractWithoutAmendmentsHierarchyConfigurationJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsHierarchyConfiguration]
type contractWithoutAmendmentsHierarchyConfigurationJSON struct {
	Children    apijson.Field
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r contractWithoutAmendmentsHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r *ContractWithoutAmendmentsHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	*r = ContractWithoutAmendmentsHierarchyConfiguration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContractWithoutAmendmentsHierarchyConfigurationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfiguration],
// [ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfiguration].
func (r ContractWithoutAmendmentsHierarchyConfiguration) AsUnion() ContractWithoutAmendmentsHierarchyConfigurationUnion {
	return r.union
}

// Either a **parent** configuration with a list of children or a **child**
// configuration with a single parent.
//
// Union satisfied by
// [ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfiguration] or
// [ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfiguration].
type ContractWithoutAmendmentsHierarchyConfigurationUnion interface {
	implementsContractWithoutAmendmentsHierarchyConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContractWithoutAmendmentsHierarchyConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfiguration{}),
		},
	)
}

type ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfiguration struct {
	// List of contracts that belong to this parent.
	Children []ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationChild `json:"children,required"`
	JSON     contractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationJSON    `json:"-"`
}

// contractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfiguration]
type contractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationJSON struct {
	Children    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfiguration) implementsContractWithoutAmendmentsHierarchyConfiguration() {
}

type ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationChild struct {
	ContractID string                                                                               `json:"contract_id,required" format:"uuid"`
	CustomerID string                                                                               `json:"customer_id,required" format:"uuid"`
	JSON       contractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationChildJSON `json:"-"`
}

// contractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationChildJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationChild]
type contractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationChildJSON struct {
	ContractID  apijson.Field
	CustomerID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationChild) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsHierarchyConfigurationParentHierarchyConfigurationChildJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfiguration struct {
	// The single parent contract/customer for this child.
	Parent ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationParent `json:"parent,required"`
	JSON   contractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationJSON   `json:"-"`
}

// contractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfiguration]
type contractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationJSON struct {
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfiguration) implementsContractWithoutAmendmentsHierarchyConfiguration() {
}

// The single parent contract/customer for this child.
type ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationParent struct {
	ContractID string                                                                               `json:"contract_id,required" format:"uuid"`
	CustomerID string                                                                               `json:"customer_id,required" format:"uuid"`
	JSON       contractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationParentJSON `json:"-"`
}

// contractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationParentJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationParent]
type contractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationParentJSON struct {
	ContractID  apijson.Field
	CustomerID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsHierarchyConfigurationChildHierarchyConfigurationParentJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsPrepaidBalanceThresholdConfiguration struct {
	Commit ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                                                                           `json:"is_enabled,required"`
	PaymentGateConfig ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfig `json:"payment_gate_config,required"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount float64 `json:"recharge_to_amount,required"`
	// Specify the threshold amount for the contract. Each time the contract's prepaid
	// balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount,required"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID string                                                            `json:"custom_credit_type_id" format:"uuid"`
	JSON               contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationJSON `json:"-"`
}

// contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationJSON contains the
// JSON metadata for the struct
// [ContractWithoutAmendmentsPrepaidBalanceThresholdConfiguration]
type contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationJSON struct {
	Commit             apijson.Field
	IsEnabled          apijson.Field
	PaymentGateConfig  apijson.Field
	RechargeToAmount   apijson.Field
	ThresholdAmount    apijson.Field
	CustomCreditTypeID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsPrepaidBalanceThresholdConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommit struct {
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
	Specifiers []ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitSpecifier `json:"specifiers"`
	JSON       contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitJSON        `json:"-"`
}

// contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitJSON contains
// the JSON metadata for the struct
// [ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommit]
type contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitJSON struct {
	ProductID             apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Description           apijson.Field
	Name                  apijson.Field
	Specifiers            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                                                         `json:"product_tags"`
	JSON        contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitSpecifierJSON `json:"-"`
}

// contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitSpecifierJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitSpecifier]
type contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationCommitSpecifierJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType `json:"tax_type"`
	JSON    contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigJSON    `json:"-"`
}

// contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfig]
type contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigJSON struct {
	PaymentGateType        apijson.Field
	PrecalculatedTaxConfig apijson.Field
	StripeConfig           apijson.Field
	TaxType                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigJSON) RawJSON() string {
	return r.raw
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType string

const (
	ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone     ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "NONE"
	ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe   ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "STRIPE"
	ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone, ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe, ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName string                                                                                                   `json:"tax_name"`
	JSON    contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON `json:"-"`
}

// contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig]
type contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON struct {
	TaxAmount   apijson.Field
	TaxName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON) RawJSON() string {
	return r.raw
}

// Only applicable if using STRIPE as your payment gate type.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string                                                                              `json:"invoice_metadata"`
	JSON            contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON `json:"-"`
}

// contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig]
type contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON struct {
	PaymentType     apijson.Field
	InvoiceMetadata apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON) RawJSON() string {
	return r.raw
}

// If left blank, will default to INVOICE
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType string

const (
	ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice       ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "INVOICE"
	ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice, ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType string

const (
	ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone          ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "NONE"
	ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe        ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "STRIPE"
	ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok         ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "ANROK"
	ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "PRECALCULATED"
)

func (r ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone, ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe, ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok, ContractWithoutAmendmentsPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated:
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
	HierarchyConfiguration ContractWithoutAmendmentsRecurringCommitsHierarchyConfiguration `json:"hierarchy_configuration"`
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
	Specifiers []ContractWithoutAmendmentsRecurringCommitsSpecifier `json:"specifiers"`
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

// Optional configuration for recurring commit/credit hierarchy access control
type ContractWithoutAmendmentsRecurringCommitsHierarchyConfiguration struct {
	ChildAccess ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationJSON        `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationJSON contains the
// JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfiguration]
type contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess struct {
	Type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                                    `json:"contract_ids"`
	JSON        contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessJSON `json:"-"`
	union       ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessUnion
}

// contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess]
type contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
func (r ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess) AsUnion() ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
// or
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessUnion interface {
	implementsContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs{}),
		},
	)
}

type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType `json:"type,required"`
	JSON contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll]
type contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess() {
}

type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType `json:"type,required"`
	JSON contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
type contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess() {
}

type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs []string                                                                                                            `json:"contract_ids,required" format:"uuid"`
	Type        ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType `json:"type,required"`
	JSON        contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs]
type contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccess() {
}

type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessType string

const (
	ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeAll         ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessType = "ALL"
	ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeNone        ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessType = "NONE"
	ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeContractIDs ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeAll, ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeNone, ContractWithoutAmendmentsRecurringCommitsHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
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

type ContractWithoutAmendmentsRecurringCommitsSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                               `json:"product_tags"`
	JSON        contractWithoutAmendmentsRecurringCommitsSpecifierJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsSpecifierJSON contains the JSON
// metadata for the struct [ContractWithoutAmendmentsRecurringCommitsSpecifier]
type contractWithoutAmendmentsRecurringCommitsSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsSpecifierJSON) RawJSON() string {
	return r.raw
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
	HierarchyConfiguration ContractWithoutAmendmentsRecurringCreditsHierarchyConfiguration `json:"hierarchy_configuration"`
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
	Specifiers []ContractWithoutAmendmentsRecurringCreditsSpecifier `json:"specifiers"`
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

// Optional configuration for recurring commit/credit hierarchy access control
type ContractWithoutAmendmentsRecurringCreditsHierarchyConfiguration struct {
	ChildAccess ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationJSON        `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationJSON contains the
// JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfiguration]
type contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess struct {
	Type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                                    `json:"contract_ids"`
	JSON        contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessJSON `json:"-"`
	union       ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessUnion
}

// contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess]
type contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
func (r ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess) AsUnion() ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
// or
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessUnion interface {
	implementsContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs{}),
		},
	)
}

type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType `json:"type,required"`
	JSON contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll]
type contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess() {
}

type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType `json:"type,required"`
	JSON contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
type contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess() {
}

type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs []string                                                                                                            `json:"contract_ids,required" format:"uuid"`
	Type        ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType `json:"type,required"`
	JSON        contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs]
type contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON) RawJSON() string {
	return r.raw
}

func (r ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccess() {
}

type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessType string

const (
	ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeAll         ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessType = "ALL"
	ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeNone        ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessType = "NONE"
	ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeContractIDs ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeAll, ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeNone, ContractWithoutAmendmentsRecurringCreditsHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
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

type ContractWithoutAmendmentsRecurringCreditsSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                               `json:"product_tags"`
	JSON        contractWithoutAmendmentsRecurringCreditsSpecifierJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsSpecifierJSON contains the JSON
// metadata for the struct [ContractWithoutAmendmentsRecurringCreditsSpecifier]
type contractWithoutAmendmentsRecurringCreditsSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsSpecifierJSON) RawJSON() string {
	return r.raw
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

type ContractWithoutAmendmentsSpendThresholdConfiguration struct {
	Commit ContractWithoutAmendmentsSpendThresholdConfigurationCommit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                                                                  `json:"is_enabled,required"`
	PaymentGateConfig ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfig `json:"payment_gate_config,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64                                                  `json:"threshold_amount,required"`
	JSON            contractWithoutAmendmentsSpendThresholdConfigurationJSON `json:"-"`
}

// contractWithoutAmendmentsSpendThresholdConfigurationJSON contains the JSON
// metadata for the struct [ContractWithoutAmendmentsSpendThresholdConfiguration]
type contractWithoutAmendmentsSpendThresholdConfigurationJSON struct {
	Commit            apijson.Field
	IsEnabled         apijson.Field
	PaymentGateConfig apijson.Field
	ThresholdAmount   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsSpendThresholdConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsSpendThresholdConfigurationJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsSpendThresholdConfigurationCommit struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID   string `json:"product_id,required"`
	Description string `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string                                                         `json:"name"`
	JSON contractWithoutAmendmentsSpendThresholdConfigurationCommitJSON `json:"-"`
}

// contractWithoutAmendmentsSpendThresholdConfigurationCommitJSON contains the JSON
// metadata for the struct
// [ContractWithoutAmendmentsSpendThresholdConfigurationCommit]
type contractWithoutAmendmentsSpendThresholdConfigurationCommitJSON struct {
	ProductID   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsSpendThresholdConfigurationCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsSpendThresholdConfigurationCommitJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateType `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfig `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxType `json:"tax_type"`
	JSON    contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigJSON    `json:"-"`
}

// contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfig]
type contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigJSON struct {
	PaymentGateType        apijson.Field
	PrecalculatedTaxConfig apijson.Field
	StripeConfig           apijson.Field
	TaxType                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigJSON) RawJSON() string {
	return r.raw
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateType string

const (
	ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone     ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "NONE"
	ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe   ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "STRIPE"
	ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone, ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe, ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName string                                                                                          `json:"tax_name"`
	JSON    contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON `json:"-"`
}

// contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig]
type contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON struct {
	TaxAmount   apijson.Field
	TaxName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON) RawJSON() string {
	return r.raw
}

// Only applicable if using STRIPE as your payment gate type.
type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string                                                                     `json:"invoice_metadata"`
	JSON            contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON `json:"-"`
}

// contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON
// contains the JSON metadata for the struct
// [ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfig]
type contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON struct {
	PaymentType     apijson.Field
	InvoiceMetadata apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON) RawJSON() string {
	return r.raw
}

// If left blank, will default to INVOICE
type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType string

const (
	ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice       ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "INVOICE"
	ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice, ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxType string

const (
	ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeNone          ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxType = "NONE"
	ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe        ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxType = "STRIPE"
	ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok         ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxType = "ANROK"
	ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxType = "PRECALCULATED"
)

func (r ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeNone, ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe, ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok, ContractWithoutAmendmentsSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated:
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
	HierarchyConfiguration CreditHierarchyConfiguration `json:"hierarchy_configuration"`
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
	Specifiers []CreditSpecifier `json:"specifiers"`
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

func (r Credit) ImplementsV1ContractListBalancesResponseData() {}

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

// Optional configuration for credit hierarchy access control
type CreditHierarchyConfiguration struct {
	ChildAccess CreditHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        creditHierarchyConfigurationJSON        `json:"-"`
}

// creditHierarchyConfigurationJSON contains the JSON metadata for the struct
// [CreditHierarchyConfiguration]
type creditHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type CreditHierarchyConfigurationChildAccess struct {
	Type CreditHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                 `json:"contract_ids"`
	JSON        creditHierarchyConfigurationChildAccessJSON `json:"-"`
	union       CreditHierarchyConfigurationChildAccessUnion
}

// creditHierarchyConfigurationChildAccessJSON contains the JSON metadata for the
// struct [CreditHierarchyConfigurationChildAccess]
type creditHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r creditHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *CreditHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = CreditHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CreditHierarchyConfigurationChildAccessUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
func (r CreditHierarchyConfigurationChildAccess) AsUnion() CreditHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone] or
// [CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
type CreditHierarchyConfigurationChildAccessUnion interface {
	implementsCreditHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CreditHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs{}),
		},
	)
}

type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType `json:"type,required"`
	JSON creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON `json:"-"`
}

// creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON
// contains the JSON metadata for the struct
// [CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll]
type creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON) RawJSON() string {
	return r.raw
}

func (r CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsCreditHierarchyConfigurationChildAccess() {
}

type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType `json:"type,required"`
	JSON creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON `json:"-"`
}

// creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON
// contains the JSON metadata for the struct
// [CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
type creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON) RawJSON() string {
	return r.raw
}

func (r CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsCreditHierarchyConfigurationChildAccess() {
}

type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs []string                                                                         `json:"contract_ids,required" format:"uuid"`
	Type        CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType `json:"type,required"`
	JSON        creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON `json:"-"`
}

// creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON
// contains the JSON metadata for the struct
// [CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs]
type creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON) RawJSON() string {
	return r.raw
}

func (r CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsCreditHierarchyConfigurationChildAccess() {
}

type CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case CreditHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type CreditHierarchyConfigurationChildAccessType string

const (
	CreditHierarchyConfigurationChildAccessTypeAll         CreditHierarchyConfigurationChildAccessType = "ALL"
	CreditHierarchyConfigurationChildAccessTypeNone        CreditHierarchyConfigurationChildAccessType = "NONE"
	CreditHierarchyConfigurationChildAccessTypeContractIDs CreditHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r CreditHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case CreditHierarchyConfigurationChildAccessTypeAll, CreditHierarchyConfigurationChildAccessTypeNone, CreditHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
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

type CreditSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string            `json:"product_tags"`
	JSON        creditSpecifierJSON `json:"-"`
}

// creditSpecifierJSON contains the JSON metadata for the struct [CreditSpecifier]
type creditSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CreditSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditSpecifierJSON) RawJSON() string {
	return r.raw
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
	CreditType    CreditTypeData                    `json:"credit_type"`
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
