// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/shared"
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
func NewV1CustomerCommitService(opts ...option.RequestOption) (r V1CustomerCommitService) {
	r = V1CustomerCommitService{}
	r.Options = opts
	return
}

// Creates customer-level commits that establish spending commitments for customers
// across their Metronome usage. Commits represent contracted spending obligations
// that can be either prepaid (paid upfront) or postpaid (billed later). Note: In
// most cases, you should add commitments directly to customer contracts using the
// contract/create or contract/edit APIs.
//
// ### Use this endpoint to:
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
// ### Usage guidelines:
//
// ⚠️ Preferred Alternative: In most cases, you should add commits directly to
// contracts using the create contract or edit contract APIs instead of creating
// customer-level commits. Contract-level commits provide better organization and
// are the recommended approach for standard use cases.
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
// ### Use this endpoint to:
//
// - Display commitment balances and utilization in customer dashboards
// - Track prepaid commitment drawdown and remaining balances
// - Monitor postpaid commitment progress toward minimum thresholds
// - Build commitment tracking and forecasting tools
// - Show commitment history with optional ledger details
// - Manage rollover balances between contract periods
//
// ### Key response fields:
//
// An array of Commit objects containing:
//
// - Commit type: PREPAID (pay upfront) or POSTPAID (pay at true-up)
// - Rate type: COMMIT_RATE (discounted) or LIST_RATE (standard pricing)
// - Access schedule: When commitment funds become available
// - Invoice schedule: When the customer is billed
// - Product targeting: Which product(s) usage is eligible to draw from this commit
// - Optional ledger entries: Transaction history (if `include_ledgers=true`)
// - Balance information: Current available amount (if `include_balance=true`)
// - Rollover settings: Fraction of unused amount that carries forward
//
// ### Usage guidelines:
//
// - Pagination: Results limited to 25 commits per page; use 'next_page' for more
// - Date filtering options:
//   - covering_date: Commits active on a specific date
//   - starting_at: Commits with access on/after a date
//   - effective_before: Commits with access before a date (exclusive)
//
// - Scope options:
//   - `include_contract_commits`: Include contract-level commits (not just
//     customer-level)
//   - `include_archived`: Include archived commits and commits from archived
//     contracts
//
// - Performance considerations:
//   - include_ledgers: Adds detailed transaction history (slower)
//   - include_balance: Adds current balance calculation (slower)
//
// - Optional filtering: Use commit_id to retrieve a specific commit
func (r *V1CustomerCommitService) List(ctx context.Context, body V1CustomerCommitListParams, opts ...option.RequestOption) (res *pagination.BodyCursorPage[shared.Commit], err error) {
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
// ### Use this endpoint to:
//
// - Display commitment balances and utilization in customer dashboards
// - Track prepaid commitment drawdown and remaining balances
// - Monitor postpaid commitment progress toward minimum thresholds
// - Build commitment tracking and forecasting tools
// - Show commitment history with optional ledger details
// - Manage rollover balances between contract periods
//
// ### Key response fields:
//
// An array of Commit objects containing:
//
// - Commit type: PREPAID (pay upfront) or POSTPAID (pay at true-up)
// - Rate type: COMMIT_RATE (discounted) or LIST_RATE (standard pricing)
// - Access schedule: When commitment funds become available
// - Invoice schedule: When the customer is billed
// - Product targeting: Which product(s) usage is eligible to draw from this commit
// - Optional ledger entries: Transaction history (if `include_ledgers=true`)
// - Balance information: Current available amount (if `include_balance=true`)
// - Rollover settings: Fraction of unused amount that carries forward
//
// ### Usage guidelines:
//
// - Pagination: Results limited to 25 commits per page; use 'next_page' for more
// - Date filtering options:
//   - covering_date: Commits active on a specific date
//   - starting_at: Commits with access on/after a date
//   - effective_before: Commits with access before a date (exclusive)
//
// - Scope options:
//   - `include_contract_commits`: Include contract-level commits (not just
//     customer-level)
//   - `include_archived`: Include archived commits and commits from archived
//     contracts
//
// - Performance considerations:
//   - include_ledgers: Adds detailed transaction history (slower)
//   - include_balance: Adds current balance calculation (slower)
//
// - Optional filtering: Use commit_id to retrieve a specific commit
func (r *V1CustomerCommitService) ListAutoPaging(ctx context.Context, body V1CustomerCommitListParams, opts ...option.RequestOption) *pagination.BodyCursorPageAutoPager[shared.Commit] {
	return pagination.NewBodyCursorPageAutoPager(r.List(ctx, body, opts...))
}

// Shortens the end date of a prepaid commit to terminate it earlier than
// originally scheduled. Use this endpoint when you need to cancel or reduce the
// duration of an existing prepaid commit. Only works with prepaid commit types and
// can only move the end date forward (earlier), not extend it.
//
// ### Usage guidelines:
//
// To extend commit end dates or make other comprehensive edits, use the 'edit
// commit' endpoint instead.
func (r *V1CustomerCommitService) UpdateEndDate(ctx context.Context, body V1CustomerCommitUpdateEndDateParams, opts ...option.RequestOption) (res *V1CustomerCommitUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCommits/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1CustomerCommitNewResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerCommitNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerCommitNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerCommitUpdateEndDateResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerCommitUpdateEndDateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerCommitUpdateEndDateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerCommitNewParams struct {
	// Schedule for distributing the commit to the customer. For "POSTPAID" commits
	// only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule V1CustomerCommitNewParamsAccessSchedule `json:"access_schedule,omitzero,required"`
	CustomerID     string                                  `json:"customer_id,required" format:"uuid"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority,required"`
	// ID of the fixed product associated with the commit. This is required because
	// products are used to invoice the commit amount.
	ProductID string `json:"product_id,required" format:"uuid"`
	// Any of "PREPAID", "POSTPAID".
	Type V1CustomerCommitNewParamsType `json:"type,omitzero,required"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Opt[string] `json:"description,omitzero"`
	// The contract that this commit will be billed on. This is required for "POSTPAID"
	// commits and for "PREPAID" commits unless there is no invoice schedule above
	// (i.e., the commit is 'free').
	InvoiceContractID param.Opt[string] `json:"invoice_contract_id,omitzero" format:"uuid"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Opt[string] `json:"salesforce_opportunity_id,omitzero"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey param.Opt[string] `json:"uniqueness_key,omitzero"`
	// Which contract the commit applies to. If not provided, the commit applies to all
	// contracts.
	ApplicableContractIDs []string `json:"applicable_contract_ids,omitzero"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this
	// time and only one schedule item is allowed; the total must match
	// accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this
	// will be a "complimentary" commit with no invoice.
	InvoiceSchedule V1CustomerCommitNewParamsInvoiceSchedule `json:"invoice_schedule,omitzero"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType V1CustomerCommitNewParamsRateType `json:"rate_type,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	paramObj
}

func (r V1CustomerCommitNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerCommitNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerCommitNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schedule for distributing the commit to the customer. For "POSTPAID" commits
// only one schedule item is allowed and amount must match invoice_schedule total.
//
// The property ScheduleItems is required.
type V1CustomerCommitNewParamsAccessSchedule struct {
	ScheduleItems []V1CustomerCommitNewParamsAccessScheduleScheduleItem `json:"schedule_items,omitzero,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r V1CustomerCommitNewParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerCommitNewParamsAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerCommitNewParamsAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, EndingBefore, StartingAt are required.
type V1CustomerCommitNewParamsAccessScheduleScheduleItem struct {
	Amount float64 `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	paramObj
}

func (r V1CustomerCommitNewParamsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerCommitNewParamsAccessScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerCommitNewParamsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerCommitNewParamsType string

const (
	V1CustomerCommitNewParamsTypePrepaid  V1CustomerCommitNewParamsType = "PREPAID"
	V1CustomerCommitNewParamsTypePostpaid V1CustomerCommitNewParamsType = "POSTPAID"
)

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match
// accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this
// will be a "complimentary" commit with no invoice.
type V1CustomerCommitNewParamsInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Opt[bool] `json:"do_not_invoice,omitzero"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule `json:"recurring_schedule,omitzero"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V1CustomerCommitNewParamsInvoiceScheduleScheduleItem `json:"schedule_items,omitzero"`
	paramObj
}

func (r V1CustomerCommitNewParamsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerCommitNewParamsInvoiceSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerCommitNewParamsInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
//
// The properties AmountDistribution, EndingBefore, Frequency, StartingAt are
// required.
type V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution,omitzero,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL".
	Frequency string `json:"frequency,omitzero,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule](
		"amount_distribution", "DIVIDED", "DIVIDED_ROUNDED", "EACH",
	)
	apijson.RegisterFieldValidator[V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL",
	)
}

// The property Timestamp is required.
type V1CustomerCommitNewParamsInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1CustomerCommitNewParamsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerCommitNewParamsInvoiceScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerCommitNewParamsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerCommitNewParamsRateType string

const (
	V1CustomerCommitNewParamsRateTypeCommitRate V1CustomerCommitNewParamsRateType = "COMMIT_RATE"
	V1CustomerCommitNewParamsRateTypeListRate   V1CustomerCommitNewParamsRateType = "LIST_RATE"
)

type V1CustomerCommitListParams struct {
	CustomerID string            `json:"customer_id,required" format:"uuid"`
	CommitID   param.Opt[string] `json:"commit_id,omitzero" format:"uuid"`
	// Include only commits that have access schedules that "cover" the provided date
	CoveringDate param.Opt[time.Time] `json:"covering_date,omitzero" format:"date-time"`
	// Include only commits that have any access before the provided date (exclusive)
	EffectiveBefore param.Opt[time.Time] `json:"effective_before,omitzero" format:"date-time"`
	// Include archived commits and commits from archived contracts.
	IncludeArchived param.Opt[bool] `json:"include_archived,omitzero"`
	// Include the balance in the response. Setting this flag may cause the query to be
	// slower.
	IncludeBalance param.Opt[bool] `json:"include_balance,omitzero"`
	// Include commits on the contract level.
	IncludeContractCommits param.Opt[bool] `json:"include_contract_commits,omitzero"`
	// Include commit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Opt[bool] `json:"include_ledgers,omitzero"`
	// The maximum number of commits to return. Defaults to 25.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// The next page token from a previous response.
	NextPage param.Opt[string] `json:"next_page,omitzero"`
	// Include only commits that have any access on or after the provided date
	StartingAt param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V1CustomerCommitListParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerCommitListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerCommitListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerCommitUpdateEndDateParams struct {
	// ID of the commit to update. Only supports "PREPAID" commits.
	CommitID string `json:"commit_id,required" format:"uuid"`
	// ID of the customer whose commit is to be updated
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// RFC 3339 timestamp indicating when access to the commit will end and it will no
	// longer be possible to draw it down (exclusive). If not provided, the access will
	// not be updated.
	AccessEndingBefore param.Opt[time.Time] `json:"access_ending_before,omitzero" format:"date-time"`
	// RFC 3339 timestamp indicating when the commit will stop being invoiced
	// (exclusive). If not provided, the invoice schedule will not be updated.
	InvoicesEndingBefore param.Opt[time.Time] `json:"invoices_ending_before,omitzero" format:"date-time"`
	paramObj
}

func (r V1CustomerCommitUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerCommitUpdateEndDateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerCommitUpdateEndDateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
