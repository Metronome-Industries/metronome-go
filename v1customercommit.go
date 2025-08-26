// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
	"github.com/tidwall/gjson"
)

// V1CustomerCommitService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerCommitService] method instead.
type V1CustomerCommitService struct {
	Options []option.RequestOption
}

// NewV1CustomerCommitService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerCommitService(opts ...option.RequestOption) (r *V1CustomerCommitService) {
	r = &V1CustomerCommitService{}
	r.Options = opts
	return
}

// Creates customer-level commits that establish spending commitments for customers
// across their Metronome usage. Commits represent contracted spending obligations
// that can be either prepaid (paid upfront) or postpaid (billed later). Note: In
// most cases, you should add commitments directly to customer contracts using the
// contract/create or contract/edit APIs.
//
// Use this endpoint to:\
//
// Use this endpoint when you need to establish customer-level spending commitments
// that can be applied across multiple contracts or scoped to specific contracts.
// Customer-level commits are ideal for:
//
// - Enterprise-wide minimum spending agreements that span multiple contracts
// - Multi-contract volume commitments with shared spending pools
// - Cross-contract discount tiers based on aggregate usage
//
// Commit type Requirements: You must specify either "prepaid" or "postpaid" as the
// commit type:
//
//   - Prepaid commits: Customer pays upfront; invoice_schedule is optional (if
//     omitted, creates a commit without an invoice)
//   - Postpaid commits: Customer pays when the commitment expires (the end of the
//     access_schedule); invoice_schedule is required and must match access_schedule
//     totals.
//
// Billing configuration:
//
//   - invoice_contract_id is required for postpaid commits and for prepaid commits
//     with billing (only optional for free prepaid commits)
//   - For postpaid commits: access_schedule and invoice_schedule must have matching
//     amounts
//   - For postpaid commits: only one schedule item is allowed in both schedules.
//
// Scoping flexibility: Customer-level commits can be configured in a few ways:
//
//   - Contract-specific: Use the applicable_contract_ids field to limit the commit
//     to specific contracts
//   - Cross-contract: Leave applicable_contract_ids empty to allow the commit to be
//     used across all of the customer's contracts
//
// Product targeting: Commits can be scoped to specific products using
// applicable_product_ids, applicable_product_tags, or specifiers, or left
// unrestricted to apply to all products.
//
// Priority considerations: When multiple commits are applicable, the one with the
// lower priority value will be consumed first. If there is a tie, contract level
// commits and credits will be applied before customer level commits and credits.
// Plan your priority scheme carefully to ensure commits are applied in the desired
// order.
//
// Usage guidelines:\
// ⚠️ Preferred Alternative: In most cases, you should add commits directly to contracts
// using the create contract or edit contract APIs instead of creating customer-level
// commits. Contract-level commits provide better organization and are the recommended
// approach for standard use cases.
func (r *V1CustomerCommitService) New(ctx context.Context, body V1CustomerCommitNewParams, opts ...option.RequestOption) (res *V1CustomerCommitNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCommits/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve all commit agreements for a customer, including both prepaid and
// postpaid commitments. This endpoint provides comprehensive visibility into
// contractual spending obligations, enabling you to track commitment utilization
// and manage customer contracts effectively.
//
// Use this endpoint to:
//
// - Display commitment balances and utilization in customer dashboards
// - Track prepaid commitment drawdown and remaining balances
// - Monitor postpaid commitment progress toward minimum thresholds
// - Build commitment tracking and forecasting tools
// - Show commitment history with optional ledger details
// - Manage rollover balances between contract periods
//
// Key response fields: An array of Commit objects containing:
//
// - Commit type: PREPAID (pay upfront) or POSTPAID (pay at true-up)
// - Rate type: COMMIT_RATE (discounted) or LIST_RATE (standard pricing)
// - Access schedule: When commitment funds become available
// - Invoice schedule: When the customer is billed
// - Product targeting: Which product(s) usage is eligible to draw from this commit
// - Optional ledger entries: Transaction history (if include_ledgers=true)
// - Balance information: Current available amount (if include_balance=true)
// - Rollover settings: Fraction of unused amount that carries forward
//
// Usage guidelines:
//
// - Pagination: Results limited to 25 commits per page; use next_page for more
// - Date filtering options:
//   - covering_date: Commits active on a specific date
//   - starting_at: Commits with access on/after a date
//   - effective_before: Commits with access before a date (exclusive)
//
// - Scope options:
//   - include_contract_commits: Include contract-level commits (not just
//     customer-level)
//   - include_archived: Include archived commits and commits from archived
//     contracts
//
// - Performance considerations:
//   - include_ledgers: Adds detailed transaction history (slower)
//   - include_balance: Adds current balance calculation (slower)
//
// - Optional filtering: Use commit_id to retrieve a specific commit
func (r *V1CustomerCommitService) List(ctx context.Context, body V1CustomerCommitListParams, opts ...option.RequestOption) (res *pagination.BodyCursorPage[V1CustomerCommitListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/contracts/customerCommits/list"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve all commit agreements for a customer, including both prepaid and
// postpaid commitments. This endpoint provides comprehensive visibility into
// contractual spending obligations, enabling you to track commitment utilization
// and manage customer contracts effectively.
//
// Use this endpoint to:
//
// - Display commitment balances and utilization in customer dashboards
// - Track prepaid commitment drawdown and remaining balances
// - Monitor postpaid commitment progress toward minimum thresholds
// - Build commitment tracking and forecasting tools
// - Show commitment history with optional ledger details
// - Manage rollover balances between contract periods
//
// Key response fields: An array of Commit objects containing:
//
// - Commit type: PREPAID (pay upfront) or POSTPAID (pay at true-up)
// - Rate type: COMMIT_RATE (discounted) or LIST_RATE (standard pricing)
// - Access schedule: When commitment funds become available
// - Invoice schedule: When the customer is billed
// - Product targeting: Which product(s) usage is eligible to draw from this commit
// - Optional ledger entries: Transaction history (if include_ledgers=true)
// - Balance information: Current available amount (if include_balance=true)
// - Rollover settings: Fraction of unused amount that carries forward
//
// Usage guidelines:
//
// - Pagination: Results limited to 25 commits per page; use next_page for more
// - Date filtering options:
//   - covering_date: Commits active on a specific date
//   - starting_at: Commits with access on/after a date
//   - effective_before: Commits with access before a date (exclusive)
//
// - Scope options:
//   - include_contract_commits: Include contract-level commits (not just
//     customer-level)
//   - include_archived: Include archived commits and commits from archived
//     contracts
//
// - Performance considerations:
//   - include_ledgers: Adds detailed transaction history (slower)
//   - include_balance: Adds current balance calculation (slower)
//
// - Optional filtering: Use commit_id to retrieve a specific commit
func (r *V1CustomerCommitService) ListAutoPaging(ctx context.Context, body V1CustomerCommitListParams, opts ...option.RequestOption) *pagination.BodyCursorPageAutoPager[V1CustomerCommitListResponse] {
	return pagination.NewBodyCursorPageAutoPager(r.List(ctx, body, opts...))
}

// Shortens the end date of a prepaid commit to terminate it earlier than
// originally scheduled. Use this endpoint when you need to cancel or reduce the
// duration of an existing prepaid commit. Only works with prepaid commit types and
// can only move the end date forward (earlier), not extend it.
//
// Usage guidelines:\ To extend commit end dates or make other comprehensive edits,
// use the 'edit commit' endpoint instead.
func (r *V1CustomerCommitService) UpdateEndDate(ctx context.Context, body V1CustomerCommitUpdateEndDateParams, opts ...option.RequestOption) (res *V1CustomerCommitUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCommits/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1CustomerCommitNewResponse struct {
	Data V1CustomerCommitNewResponseData `json:"data,required"`
	JSON v1CustomerCommitNewResponseJSON `json:"-"`
}

// v1CustomerCommitNewResponseJSON contains the JSON metadata for the struct
// [V1CustomerCommitNewResponse]
type v1CustomerCommitNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitNewResponseData struct {
	ID   string                              `json:"id,required" format:"uuid"`
	JSON v1CustomerCommitNewResponseDataJSON `json:"-"`
}

// v1CustomerCommitNewResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerCommitNewResponseData]
type v1CustomerCommitNewResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponse struct {
	ID      string                              `json:"id,required" format:"uuid"`
	Product V1CustomerCommitListResponseProduct `json:"product,required"`
	Type    V1CustomerCommitListResponseType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule V1CustomerCommitListResponseAccessSchedule `json:"access_schedule"`
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
	Balance  float64                              `json:"balance"`
	Contract V1CustomerCommitListResponseContract `json:"contract"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration V1CustomerCommitListResponseHierarchyConfiguration `json:"hierarchy_configuration"`
	// The contract that this commit will be billed on.
	InvoiceContract V1CustomerCommitListResponseInvoiceContract `json:"invoice_contract"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule V1CustomerCommitListResponseInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger []V1CustomerCommitListResponseLedger `json:"ledger"`
	Name   string                               `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority         float64                                    `json:"priority"`
	RateType         V1CustomerCommitListResponseRateType       `json:"rate_type"`
	RolledOverFrom   V1CustomerCommitListResponseRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction float64                                    `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []V1CustomerCommitListResponseSpecifier `json:"specifiers"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey string                           `json:"uniqueness_key"`
	JSON          v1CustomerCommitListResponseJSON `json:"-"`
}

// v1CustomerCommitListResponseJSON contains the JSON metadata for the struct
// [V1CustomerCommitListResponse]
type v1CustomerCommitListResponseJSON struct {
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

func (r *V1CustomerCommitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseProduct struct {
	ID   string                                  `json:"id,required" format:"uuid"`
	Name string                                  `json:"name,required"`
	JSON v1CustomerCommitListResponseProductJSON `json:"-"`
}

// v1CustomerCommitListResponseProductJSON contains the JSON metadata for the
// struct [V1CustomerCommitListResponseProduct]
type v1CustomerCommitListResponseProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseProductJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseType string

const (
	V1CustomerCommitListResponseTypePrepaid  V1CustomerCommitListResponseType = "PREPAID"
	V1CustomerCommitListResponseTypePostpaid V1CustomerCommitListResponseType = "POSTPAID"
)

func (r V1CustomerCommitListResponseType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseTypePrepaid, V1CustomerCommitListResponseTypePostpaid:
		return true
	}
	return false
}

// The schedule that the customer will gain access to the credits purposed with
// this commit.
type V1CustomerCommitListResponseAccessSchedule struct {
	ScheduleItems []V1CustomerCommitListResponseAccessScheduleScheduleItem `json:"schedule_items,required"`
	CreditType    V1CustomerCommitListResponseAccessScheduleCreditType     `json:"credit_type"`
	JSON          v1CustomerCommitListResponseAccessScheduleJSON           `json:"-"`
}

// v1CustomerCommitListResponseAccessScheduleJSON contains the JSON metadata for
// the struct [V1CustomerCommitListResponseAccessSchedule]
type v1CustomerCommitListResponseAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseAccessScheduleScheduleItem struct {
	ID           string                                                     `json:"id,required" format:"uuid"`
	Amount       float64                                                    `json:"amount,required"`
	EndingBefore time.Time                                                  `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                                                  `json:"starting_at,required" format:"date-time"`
	JSON         v1CustomerCommitListResponseAccessScheduleScheduleItemJSON `json:"-"`
}

// v1CustomerCommitListResponseAccessScheduleScheduleItemJSON contains the JSON
// metadata for the struct [V1CustomerCommitListResponseAccessScheduleScheduleItem]
type v1CustomerCommitListResponseAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseAccessScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseAccessScheduleCreditType struct {
	ID   string                                                   `json:"id,required" format:"uuid"`
	Name string                                                   `json:"name,required"`
	JSON v1CustomerCommitListResponseAccessScheduleCreditTypeJSON `json:"-"`
}

// v1CustomerCommitListResponseAccessScheduleCreditTypeJSON contains the JSON
// metadata for the struct [V1CustomerCommitListResponseAccessScheduleCreditType]
type v1CustomerCommitListResponseAccessScheduleCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseAccessScheduleCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseAccessScheduleCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseContract struct {
	ID   string                                   `json:"id,required" format:"uuid"`
	JSON v1CustomerCommitListResponseContractJSON `json:"-"`
}

// v1CustomerCommitListResponseContractJSON contains the JSON metadata for the
// struct [V1CustomerCommitListResponseContract]
type v1CustomerCommitListResponseContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseContractJSON) RawJSON() string {
	return r.raw
}

// Optional configuration for commit hierarchy access control
type V1CustomerCommitListResponseHierarchyConfiguration struct {
	ChildAccess V1CustomerCommitListResponseHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        v1CustomerCommitListResponseHierarchyConfigurationJSON        `json:"-"`
}

// v1CustomerCommitListResponseHierarchyConfigurationJSON contains the JSON
// metadata for the struct [V1CustomerCommitListResponseHierarchyConfiguration]
type v1CustomerCommitListResponseHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseHierarchyConfigurationChildAccess struct {
	Type V1CustomerCommitListResponseHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                       `json:"contract_ids"`
	JSON        v1CustomerCommitListResponseHierarchyConfigurationChildAccessJSON `json:"-"`
	union       V1CustomerCommitListResponseHierarchyConfigurationChildAccessUnion
}

// v1CustomerCommitListResponseHierarchyConfigurationChildAccessJSON contains the
// JSON metadata for the struct
// [V1CustomerCommitListResponseHierarchyConfigurationChildAccess]
type v1CustomerCommitListResponseHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v1CustomerCommitListResponseHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *V1CustomerCommitListResponseHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = V1CustomerCommitListResponseHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [V1CustomerCommitListResponseHierarchyConfigurationChildAccessUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V1CustomerCommitListResponseHierarchyConfigurationChildAccessType],
// [V1CustomerCommitListResponseHierarchyConfigurationChildAccessType],
// [V1CustomerCommitListResponseHierarchyConfigurationChildAccessObject].
func (r V1CustomerCommitListResponseHierarchyConfigurationChildAccess) AsUnion() V1CustomerCommitListResponseHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [V1CustomerCommitListResponseHierarchyConfigurationChildAccessType],
// [V1CustomerCommitListResponseHierarchyConfigurationChildAccessType] or
// [V1CustomerCommitListResponseHierarchyConfigurationChildAccessObject].
type V1CustomerCommitListResponseHierarchyConfigurationChildAccessUnion interface {
	implementsV1CustomerCommitListResponseHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1CustomerCommitListResponseHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseHierarchyConfigurationChildAccessType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseHierarchyConfigurationChildAccessType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseHierarchyConfigurationChildAccessObject{}),
		},
	)
}

type V1CustomerCommitListResponseHierarchyConfigurationChildAccessType struct {
	Type V1CustomerCommitListResponseHierarchyConfigurationChildAccessTypeType `json:"type,required"`
	JSON v1CustomerCommitListResponseHierarchyConfigurationChildAccessTypeJSON `json:"-"`
}

// v1CustomerCommitListResponseHierarchyConfigurationChildAccessTypeJSON contains
// the JSON metadata for the struct
// [V1CustomerCommitListResponseHierarchyConfigurationChildAccessType]
type v1CustomerCommitListResponseHierarchyConfigurationChildAccessTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseHierarchyConfigurationChildAccessType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseHierarchyConfigurationChildAccessTypeJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCommitListResponseHierarchyConfigurationChildAccessType) implementsV1CustomerCommitListResponseHierarchyConfigurationChildAccess() {
}

type V1CustomerCommitListResponseHierarchyConfigurationChildAccessTypeType string

const (
	V1CustomerCommitListResponseHierarchyConfigurationChildAccessTypeTypeAll V1CustomerCommitListResponseHierarchyConfigurationChildAccessTypeType = "ALL"
)

func (r V1CustomerCommitListResponseHierarchyConfigurationChildAccessTypeType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseHierarchyConfigurationChildAccessTypeTypeAll:
		return true
	}
	return false
}

type V1CustomerCommitListResponseHierarchyConfigurationChildAccessObject struct {
	ContractIDs []string                                                                `json:"contract_ids,required" format:"uuid"`
	Type        V1CustomerCommitListResponseHierarchyConfigurationChildAccessObjectType `json:"type,required"`
	JSON        v1CustomerCommitListResponseHierarchyConfigurationChildAccessObjectJSON `json:"-"`
}

// v1CustomerCommitListResponseHierarchyConfigurationChildAccessObjectJSON contains
// the JSON metadata for the struct
// [V1CustomerCommitListResponseHierarchyConfigurationChildAccessObject]
type v1CustomerCommitListResponseHierarchyConfigurationChildAccessObjectJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseHierarchyConfigurationChildAccessObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseHierarchyConfigurationChildAccessObjectJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCommitListResponseHierarchyConfigurationChildAccessObject) implementsV1CustomerCommitListResponseHierarchyConfigurationChildAccess() {
}

type V1CustomerCommitListResponseHierarchyConfigurationChildAccessObjectType string

const (
	V1CustomerCommitListResponseHierarchyConfigurationChildAccessObjectTypeContractIDs V1CustomerCommitListResponseHierarchyConfigurationChildAccessObjectType = "CONTRACT_IDS"
)

func (r V1CustomerCommitListResponseHierarchyConfigurationChildAccessObjectType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseHierarchyConfigurationChildAccessObjectTypeContractIDs:
		return true
	}
	return false
}

// The contract that this commit will be billed on.
type V1CustomerCommitListResponseInvoiceContract struct {
	ID   string                                          `json:"id,required" format:"uuid"`
	JSON v1CustomerCommitListResponseInvoiceContractJSON `json:"-"`
}

// v1CustomerCommitListResponseInvoiceContractJSON contains the JSON metadata for
// the struct [V1CustomerCommitListResponseInvoiceContract]
type v1CustomerCommitListResponseInvoiceContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseInvoiceContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseInvoiceContractJSON) RawJSON() string {
	return r.raw
}

// The schedule that the customer will be invoiced for this commit.
type V1CustomerCommitListResponseInvoiceSchedule struct {
	CreditType V1CustomerCommitListResponseInvoiceScheduleCreditType `json:"credit_type"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice  bool                                                      `json:"do_not_invoice"`
	ScheduleItems []V1CustomerCommitListResponseInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          v1CustomerCommitListResponseInvoiceScheduleJSON           `json:"-"`
}

// v1CustomerCommitListResponseInvoiceScheduleJSON contains the JSON metadata for
// the struct [V1CustomerCommitListResponseInvoiceSchedule]
type v1CustomerCommitListResponseInvoiceScheduleJSON struct {
	CreditType    apijson.Field
	DoNotInvoice  apijson.Field
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseInvoiceScheduleJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseInvoiceScheduleCreditType struct {
	ID   string                                                    `json:"id,required" format:"uuid"`
	Name string                                                    `json:"name,required"`
	JSON v1CustomerCommitListResponseInvoiceScheduleCreditTypeJSON `json:"-"`
}

// v1CustomerCommitListResponseInvoiceScheduleCreditTypeJSON contains the JSON
// metadata for the struct [V1CustomerCommitListResponseInvoiceScheduleCreditType]
type v1CustomerCommitListResponseInvoiceScheduleCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseInvoiceScheduleCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseInvoiceScheduleCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseInvoiceScheduleScheduleItem struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	Amount    float64                                                     `json:"amount,required"`
	Quantity  float64                                                     `json:"quantity,required"`
	Timestamp time.Time                                                   `json:"timestamp,required" format:"date-time"`
	UnitPrice float64                                                     `json:"unit_price,required"`
	InvoiceID string                                                      `json:"invoice_id,nullable" format:"uuid"`
	JSON      v1CustomerCommitListResponseInvoiceScheduleScheduleItemJSON `json:"-"`
}

// v1CustomerCommitListResponseInvoiceScheduleScheduleItemJSON contains the JSON
// metadata for the struct
// [V1CustomerCommitListResponseInvoiceScheduleScheduleItem]
type v1CustomerCommitListResponseInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	InvoiceID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseInvoiceScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseLedger struct {
	Amount        float64                                `json:"amount,required"`
	Timestamp     time.Time                              `json:"timestamp,required" format:"date-time"`
	Type          V1CustomerCommitListResponseLedgerType `json:"type,required"`
	ContractID    string                                 `json:"contract_id" format:"uuid"`
	InvoiceID     string                                 `json:"invoice_id" format:"uuid"`
	NewContractID string                                 `json:"new_contract_id" format:"uuid"`
	Reason        string                                 `json:"reason"`
	SegmentID     string                                 `json:"segment_id" format:"uuid"`
	JSON          v1CustomerCommitListResponseLedgerJSON `json:"-"`
	union         V1CustomerCommitListResponseLedgerUnion
}

// v1CustomerCommitListResponseLedgerJSON contains the JSON metadata for the struct
// [V1CustomerCommitListResponseLedger]
type v1CustomerCommitListResponseLedgerJSON struct {
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

func (r v1CustomerCommitListResponseLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *V1CustomerCommitListResponseLedger) UnmarshalJSON(data []byte) (err error) {
	*r = V1CustomerCommitListResponseLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V1CustomerCommitListResponseLedgerUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject].
func (r V1CustomerCommitListResponseLedger) AsUnion() V1CustomerCommitListResponseLedgerUnion {
	return r.union
}

// Union satisfied by [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject],
// [V1CustomerCommitListResponseLedgerObject] or
// [V1CustomerCommitListResponseLedgerObject].
type V1CustomerCommitListResponseLedgerUnion interface {
	implementsV1CustomerCommitListResponseLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1CustomerCommitListResponseLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseLedgerObject{}),
		},
	)
}

type V1CustomerCommitListResponseLedgerObject struct {
	Amount    float64                                      `json:"amount,required"`
	SegmentID string                                       `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                    `json:"timestamp,required" format:"date-time"`
	Type      V1CustomerCommitListResponseLedgerObjectType `json:"type,required"`
	JSON      v1CustomerCommitListResponseLedgerObjectJSON `json:"-"`
}

// v1CustomerCommitListResponseLedgerObjectJSON contains the JSON metadata for the
// struct [V1CustomerCommitListResponseLedgerObject]
type v1CustomerCommitListResponseLedgerObjectJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseLedgerObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseLedgerObjectJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCommitListResponseLedgerObject) implementsV1CustomerCommitListResponseLedger() {}

type V1CustomerCommitListResponseLedgerObjectType string

const (
	V1CustomerCommitListResponseLedgerObjectTypePrepaidCommitSegmentStart V1CustomerCommitListResponseLedgerObjectType = "PREPAID_COMMIT_SEGMENT_START"
)

func (r V1CustomerCommitListResponseLedgerObjectType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseLedgerObjectTypePrepaidCommitSegmentStart:
		return true
	}
	return false
}

type V1CustomerCommitListResponseLedgerType string

const (
	V1CustomerCommitListResponseLedgerTypePrepaidCommitSegmentStart               V1CustomerCommitListResponseLedgerType = "PREPAID_COMMIT_SEGMENT_START"
	V1CustomerCommitListResponseLedgerTypePrepaidCommitAutomatedInvoiceDeduction  V1CustomerCommitListResponseLedgerType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	V1CustomerCommitListResponseLedgerTypePrepaidCommitRollover                   V1CustomerCommitListResponseLedgerType = "PREPAID_COMMIT_ROLLOVER"
	V1CustomerCommitListResponseLedgerTypePrepaidCommitExpiration                 V1CustomerCommitListResponseLedgerType = "PREPAID_COMMIT_EXPIRATION"
	V1CustomerCommitListResponseLedgerTypePrepaidCommitCanceled                   V1CustomerCommitListResponseLedgerType = "PREPAID_COMMIT_CANCELED"
	V1CustomerCommitListResponseLedgerTypePrepaidCommitCredited                   V1CustomerCommitListResponseLedgerType = "PREPAID_COMMIT_CREDITED"
	V1CustomerCommitListResponseLedgerTypePrepaidCommitSeatBasedAdjustment        V1CustomerCommitListResponseLedgerType = "PREPAID_COMMIT_SEAT_BASED_ADJUSTMENT"
	V1CustomerCommitListResponseLedgerTypePostpaidCommitInitialBalance            V1CustomerCommitListResponseLedgerType = "POSTPAID_COMMIT_INITIAL_BALANCE"
	V1CustomerCommitListResponseLedgerTypePostpaidCommitAutomatedInvoiceDeduction V1CustomerCommitListResponseLedgerType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	V1CustomerCommitListResponseLedgerTypePostpaidCommitRollover                  V1CustomerCommitListResponseLedgerType = "POSTPAID_COMMIT_ROLLOVER"
	V1CustomerCommitListResponseLedgerTypePostpaidCommitTrueup                    V1CustomerCommitListResponseLedgerType = "POSTPAID_COMMIT_TRUEUP"
	V1CustomerCommitListResponseLedgerTypePrepaidCommitManual                     V1CustomerCommitListResponseLedgerType = "PREPAID_COMMIT_MANUAL"
	V1CustomerCommitListResponseLedgerTypePostpaidCommitManual                    V1CustomerCommitListResponseLedgerType = "POSTPAID_COMMIT_MANUAL"
	V1CustomerCommitListResponseLedgerTypePostpaidCommitExpiration                V1CustomerCommitListResponseLedgerType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r V1CustomerCommitListResponseLedgerType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseLedgerTypePrepaidCommitSegmentStart, V1CustomerCommitListResponseLedgerTypePrepaidCommitAutomatedInvoiceDeduction, V1CustomerCommitListResponseLedgerTypePrepaidCommitRollover, V1CustomerCommitListResponseLedgerTypePrepaidCommitExpiration, V1CustomerCommitListResponseLedgerTypePrepaidCommitCanceled, V1CustomerCommitListResponseLedgerTypePrepaidCommitCredited, V1CustomerCommitListResponseLedgerTypePrepaidCommitSeatBasedAdjustment, V1CustomerCommitListResponseLedgerTypePostpaidCommitInitialBalance, V1CustomerCommitListResponseLedgerTypePostpaidCommitAutomatedInvoiceDeduction, V1CustomerCommitListResponseLedgerTypePostpaidCommitRollover, V1CustomerCommitListResponseLedgerTypePostpaidCommitTrueup, V1CustomerCommitListResponseLedgerTypePrepaidCommitManual, V1CustomerCommitListResponseLedgerTypePostpaidCommitManual, V1CustomerCommitListResponseLedgerTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type V1CustomerCommitListResponseRateType string

const (
	V1CustomerCommitListResponseRateTypeCommitRate V1CustomerCommitListResponseRateType = "COMMIT_RATE"
	V1CustomerCommitListResponseRateTypeListRate   V1CustomerCommitListResponseRateType = "LIST_RATE"
)

func (r V1CustomerCommitListResponseRateType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseRateTypeCommitRate, V1CustomerCommitListResponseRateTypeListRate:
		return true
	}
	return false
}

type V1CustomerCommitListResponseRolledOverFrom struct {
	CommitID   string                                         `json:"commit_id,required" format:"uuid"`
	ContractID string                                         `json:"contract_id,required" format:"uuid"`
	JSON       v1CustomerCommitListResponseRolledOverFromJSON `json:"-"`
}

// v1CustomerCommitListResponseRolledOverFromJSON contains the JSON metadata for
// the struct [V1CustomerCommitListResponseRolledOverFrom]
type v1CustomerCommitListResponseRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseRolledOverFromJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                  `json:"product_tags"`
	JSON        v1CustomerCommitListResponseSpecifierJSON `json:"-"`
}

// v1CustomerCommitListResponseSpecifierJSON contains the JSON metadata for the
// struct [V1CustomerCommitListResponseSpecifier]
type v1CustomerCommitListResponseSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseSpecifierJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitUpdateEndDateResponse struct {
	Data V1CustomerCommitUpdateEndDateResponseData `json:"data,required"`
	JSON v1CustomerCommitUpdateEndDateResponseJSON `json:"-"`
}

// v1CustomerCommitUpdateEndDateResponseJSON contains the JSON metadata for the
// struct [V1CustomerCommitUpdateEndDateResponse]
type v1CustomerCommitUpdateEndDateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitUpdateEndDateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitUpdateEndDateResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitUpdateEndDateResponseData struct {
	ID   string                                        `json:"id,required" format:"uuid"`
	JSON v1CustomerCommitUpdateEndDateResponseDataJSON `json:"-"`
}

// v1CustomerCommitUpdateEndDateResponseDataJSON contains the JSON metadata for the
// struct [V1CustomerCommitUpdateEndDateResponseData]
type v1CustomerCommitUpdateEndDateResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitUpdateEndDateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitUpdateEndDateResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitNewParams struct {
	// Schedule for distributing the commit to the customer. For "POSTPAID" commits
	// only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule param.Field[V1CustomerCommitNewParamsAccessSchedule] `json:"access_schedule,required"`
	CustomerID     param.Field[string]                                  `json:"customer_id,required" format:"uuid"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority param.Field[float64] `json:"priority,required"`
	// ID of the fixed product associated with the commit. This is required because
	// products are used to invoice the commit amount.
	ProductID param.Field[string]                        `json:"product_id,required" format:"uuid"`
	Type      param.Field[V1CustomerCommitNewParamsType] `json:"type,required"`
	// Which contract the commit applies to. If not provided, the commit applies to all
	// contracts.
	ApplicableContractIDs param.Field[[]string] `json:"applicable_contract_ids"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// The contract that this commit will be billed on. This is required for "POSTPAID"
	// commits and for "PREPAID" commits unless there is no invoice schedule above
	// (i.e., the commit is 'free').
	InvoiceContractID param.Field[string] `json:"invoice_contract_id" format:"uuid"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this
	// time and only one schedule item is allowed; the total must match
	// accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this
	// will be a "complimentary" commit with no invoice.
	InvoiceSchedule param.Field[V1CustomerCommitNewParamsInvoiceSchedule] `json:"invoice_schedule"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string]                            `json:"netsuite_sales_order_id"`
	RateType             param.Field[V1CustomerCommitNewParamsRateType] `json:"rate_type"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string] `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers param.Field[[]V1CustomerCommitNewParamsSpecifier] `json:"specifiers"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey param.Field[string] `json:"uniqueness_key"`
}

func (r V1CustomerCommitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the commit to the customer. For "POSTPAID" commits
// only one schedule item is allowed and amount must match invoice_schedule total.
type V1CustomerCommitNewParamsAccessSchedule struct {
	ScheduleItems param.Field[[]V1CustomerCommitNewParamsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r V1CustomerCommitNewParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitNewParamsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V1CustomerCommitNewParamsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitNewParamsType string

const (
	V1CustomerCommitNewParamsTypePrepaid  V1CustomerCommitNewParamsType = "PREPAID"
	V1CustomerCommitNewParamsTypePostpaid V1CustomerCommitNewParamsType = "POSTPAID"
)

func (r V1CustomerCommitNewParamsType) IsKnown() bool {
	switch r {
	case V1CustomerCommitNewParamsTypePrepaid, V1CustomerCommitNewParamsTypePostpaid:
		return true
	}
	return false
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match
// accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this
// will be a "complimentary" commit with no invoice.
type V1CustomerCommitNewParamsInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Field[bool] `json:"do_not_invoice"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]V1CustomerCommitNewParamsInvoiceScheduleScheduleItem] `json:"schedule_items"`
}

func (r V1CustomerCommitNewParamsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                          `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDivided        V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionEach           V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDivided, V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency string

const (
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyMonthly    V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyQuarterly  V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencySemiAnnual V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyAnnual     V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyMonthly, V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyQuarterly, V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type V1CustomerCommitNewParamsInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1CustomerCommitNewParamsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitNewParamsRateType string

const (
	V1CustomerCommitNewParamsRateTypeCommitRate V1CustomerCommitNewParamsRateType = "COMMIT_RATE"
	V1CustomerCommitNewParamsRateTypeListRate   V1CustomerCommitNewParamsRateType = "LIST_RATE"
)

func (r V1CustomerCommitNewParamsRateType) IsKnown() bool {
	switch r {
	case V1CustomerCommitNewParamsRateTypeCommitRate, V1CustomerCommitNewParamsRateTypeListRate:
		return true
	}
	return false
}

type V1CustomerCommitNewParamsSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V1CustomerCommitNewParamsSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitListParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	CommitID   param.Field[string] `json:"commit_id" format:"uuid"`
	// Include only commits that have access schedules that "cover" the provided date
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	// Include only commits that have any access before the provided date (exclusive)
	EffectiveBefore param.Field[time.Time] `json:"effective_before" format:"date-time"`
	// Include archived commits and commits from archived contracts.
	IncludeArchived param.Field[bool] `json:"include_archived"`
	// Include the balance in the response. Setting this flag may cause the query to be
	// slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include commits on the contract level.
	IncludeContractCommits param.Field[bool] `json:"include_contract_commits"`
	// Include commit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
	// The maximum number of commits to return. Defaults to 25.
	Limit param.Field[int64] `json:"limit"`
	// The next page token from a previous response.
	NextPage param.Field[string] `json:"next_page"`
	// Include only commits that have any access on or after the provided date
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V1CustomerCommitListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitUpdateEndDateParams struct {
	// ID of the commit to update. Only supports "PREPAID" commits.
	CommitID param.Field[string] `json:"commit_id,required" format:"uuid"`
	// ID of the customer whose commit is to be updated
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// RFC 3339 timestamp indicating when access to the commit will end and it will no
	// longer be possible to draw it down (exclusive). If not provided, the access will
	// not be updated.
	AccessEndingBefore param.Field[time.Time] `json:"access_ending_before" format:"date-time"`
	// RFC 3339 timestamp indicating when the commit will stop being invoiced
	// (exclusive). If not provided, the invoice schedule will not be updated.
	InvoicesEndingBefore param.Field[time.Time] `json:"invoices_ending_before" format:"date-time"`
}

func (r V1CustomerCommitUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
