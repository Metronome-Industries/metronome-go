// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// V1CustomerCreditService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerCreditService] method instead.
type V1CustomerCreditService struct {
	Options []option.RequestOption
}

// NewV1CustomerCreditService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerCreditService(opts ...option.RequestOption) (r *V1CustomerCreditService) {
	r = &V1CustomerCreditService{}
	r.Options = opts
	return
}

// Creates customer-level credits that provide spending allowances or free credit
// balances for customers across their Metronome usage. Note: In most cases, you
// should add credits directly to customer contracts using the contract/create or
// contract/edit APIs.
//
// When to use this endpoint: Use this endpoint when you need to provision credits
// directly at the customer level that can be applied across multiple contracts or
// scoped to specific contracts. Customer-level credits are ideal for:
//
// - Customer onboarding incentives that apply globally
// - Flexible spending allowances that aren't tied to a single contract
// - Migration scenarios where you need to preserve existing customer balances
//
// Scoping Flexibility: Customer-level credits can be configured in two ways:
//
//   - Contract-specific: Use the applicable_contract_ids field to limit the credit
//     to specific contracts
//   - Cross-contract: Leave applicable_contract_ids empty to allow the credit to be
//     used across all of the customer's contracts
//
// Product Targeting: Credits can be scoped to specific products using
// applicable_product_ids or applicable_product_tags, or left unrestricted to apply
// to all products.
//
// Priority Considerations: When multiple credits are applicable, the one with the
// lower priority value will be consumed first. If there is a tie, contract level
// commits and credits will be applied before customer level commits and credits.
// Plan your priority scheme carefully to ensure credits are applied in the desired
// order.
//
// Access Schedule Required: You must provide an access_schedule that defines when
// and how much credit becomes available to the customer over time. This usually is
// aligned to the contract schedule or starts immediately and is set to expire in
// the future.
//
// Usage Guidelines:\
// ⚠️ Preferred Alternative: In most cases, you should add credits directly to contracts
// using the contract/create or contract/edit APIs instead of creating customer-level
// credits. Contract-level credits provide better organization, and are easier for finance
// teams to recognize revenue, and are the recommended approach for most use cases.
func (r *V1CustomerCreditService) New(ctx context.Context, body V1CustomerCreditNewParams, opts ...option.RequestOption) (res *V1CustomerCreditNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCredits/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a detailed list of all credits available to a customer, including
// promotional credits and contract-specific credits. This endpoint provides
// comprehensive visibility into credit balances, access schedules, and usage
// rules, enabling you to build credit management interfaces and track available
// funding.
//
// Use this endpoint to:
//
//   - Display all available credits in customer billing dashboards
//   - Show credit balances and expiration dates
//   - Track credit usage history with optional ledger details
//   - Build credit management and reporting tools
//   - Monitor promotional credit utilization • Support customer inquiries about
//     available credits
//
// Key response fields: An array of Credit objects containing:
//
//   - Credit details: Name, priority, and which applicable products/tags it applies
//     to
//   - Product ID: The product_id of the credit. This is for external mapping into
//     your quote-to-cash stack, not the product it applies to.
//   - Access schedule: When credits become available and expire
//   - Optional ledger entries: Transaction history (if include_ledgers=true)
//   - Balance information: Current available amount (if include_balance=true)
//   - Metadata: Custom fields and usage specifiers
//
// Usage guidelines:
//
// - Pagination: Results limited to 25 commits per page; use next_page for more
// - Date filtering options:
//   - covering_date: Credits active on a specific date
//   - starting_at: Credits with access on/after a date
//   - effective_before: Credits with access before a date (exclusive)
//
// - Scope options:
//   - include_contract_credits: Include contract-level credits (not just
//     customer-level)
//   - include_archived: Include archived credits and credits from archived
//     contracts
//
// - Performance considerations:
//   - include_ledgers: Adds detailed transaction history (slower)
//   - include_balance: Adds current balance calculation (slower)
//
// - Optional filtering: Use credit_id to retrieve a specific commit
func (r *V1CustomerCreditService) List(ctx context.Context, body V1CustomerCreditListParams, opts ...option.RequestOption) (res *pagination.BodyCursorPage[shared.Credit], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/contracts/customerCredits/list"
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

// Retrieve a detailed list of all credits available to a customer, including
// promotional credits and contract-specific credits. This endpoint provides
// comprehensive visibility into credit balances, access schedules, and usage
// rules, enabling you to build credit management interfaces and track available
// funding.
//
// Use this endpoint to:
//
//   - Display all available credits in customer billing dashboards
//   - Show credit balances and expiration dates
//   - Track credit usage history with optional ledger details
//   - Build credit management and reporting tools
//   - Monitor promotional credit utilization • Support customer inquiries about
//     available credits
//
// Key response fields: An array of Credit objects containing:
//
//   - Credit details: Name, priority, and which applicable products/tags it applies
//     to
//   - Product ID: The product_id of the credit. This is for external mapping into
//     your quote-to-cash stack, not the product it applies to.
//   - Access schedule: When credits become available and expire
//   - Optional ledger entries: Transaction history (if include_ledgers=true)
//   - Balance information: Current available amount (if include_balance=true)
//   - Metadata: Custom fields and usage specifiers
//
// Usage guidelines:
//
// - Pagination: Results limited to 25 commits per page; use next_page for more
// - Date filtering options:
//   - covering_date: Credits active on a specific date
//   - starting_at: Credits with access on/after a date
//   - effective_before: Credits with access before a date (exclusive)
//
// - Scope options:
//   - include_contract_credits: Include contract-level credits (not just
//     customer-level)
//   - include_archived: Include archived credits and credits from archived
//     contracts
//
// - Performance considerations:
//   - include_ledgers: Adds detailed transaction history (slower)
//   - include_balance: Adds current balance calculation (slower)
//
// - Optional filtering: Use credit_id to retrieve a specific commit
func (r *V1CustomerCreditService) ListAutoPaging(ctx context.Context, body V1CustomerCreditListParams, opts ...option.RequestOption) *pagination.BodyCursorPageAutoPager[shared.Credit] {
	return pagination.NewBodyCursorPageAutoPager(r.List(ctx, body, opts...))
}

// Shortens the end date of an existing customer credit to terminate it earlier
// than originally scheduled. Only allows moving end dates forward (earlier), not
// extending them.
//
// Note: To extend credit end dates or make comprehensive edits, use the 'edit
// credit' endpoint instead.
func (r *V1CustomerCreditService) UpdateEndDate(ctx context.Context, body V1CustomerCreditUpdateEndDateParams, opts ...option.RequestOption) (res *V1CustomerCreditUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCredits/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1CustomerCreditNewResponse struct {
	Data shared.ID                       `json:"data,required"`
	JSON v1CustomerCreditNewResponseJSON `json:"-"`
}

// v1CustomerCreditNewResponseJSON contains the JSON metadata for the struct
// [V1CustomerCreditNewResponse]
type v1CustomerCreditNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditUpdateEndDateResponse struct {
	Data shared.ID                                 `json:"data,required"`
	JSON v1CustomerCreditUpdateEndDateResponseJSON `json:"-"`
}

// v1CustomerCreditUpdateEndDateResponseJSON contains the JSON metadata for the
// struct [V1CustomerCreditUpdateEndDateResponse]
type v1CustomerCreditUpdateEndDateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditUpdateEndDateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditUpdateEndDateResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditNewParams struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule param.Field[V1CustomerCreditNewParamsAccessSchedule] `json:"access_schedule,required"`
	CustomerID     param.Field[string]                                  `json:"customer_id,required" format:"uuid"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority  param.Field[float64] `json:"priority,required"`
	ProductID param.Field[string]  `json:"product_id,required" format:"uuid"`
	// Which contract the credit applies to. If not provided, the credit applies to all
	// contracts.
	ApplicableContractIDs param.Field[[]string] `json:"applicable_contract_ids"`
	// Which products the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string]                            `json:"netsuite_sales_order_id"`
	RateType             param.Field[V1CustomerCreditNewParamsRateType] `json:"rate_type"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string] `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers param.Field[[]shared.CommitSpecifierInputParam] `json:"specifiers"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey param.Field[string] `json:"uniqueness_key"`
}

func (r V1CustomerCreditNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the credit to the customer.
type V1CustomerCreditNewParamsAccessSchedule struct {
	ScheduleItems param.Field[[]V1CustomerCreditNewParamsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r V1CustomerCreditNewParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCreditNewParamsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V1CustomerCreditNewParamsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCreditNewParamsRateType string

const (
	V1CustomerCreditNewParamsRateTypeCommitRate V1CustomerCreditNewParamsRateType = "COMMIT_RATE"
	V1CustomerCreditNewParamsRateTypeListRate   V1CustomerCreditNewParamsRateType = "LIST_RATE"
)

func (r V1CustomerCreditNewParamsRateType) IsKnown() bool {
	switch r {
	case V1CustomerCreditNewParamsRateTypeCommitRate, V1CustomerCreditNewParamsRateTypeListRate:
		return true
	}
	return false
}

type V1CustomerCreditListParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Return only credits that have access schedules that "cover" the provided date
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	CreditID     param.Field[string]    `json:"credit_id" format:"uuid"`
	// Include only credits that have any access before the provided date (exclusive)
	EffectiveBefore param.Field[time.Time] `json:"effective_before" format:"date-time"`
	// Include archived credits and credits from archived contracts.
	IncludeArchived param.Field[bool] `json:"include_archived"`
	// Include the balance in the response. Setting this flag may cause the query to be
	// slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include credits on the contract level.
	IncludeContractCredits param.Field[bool] `json:"include_contract_credits"`
	// Include credit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
	// The maximum number of commits to return. Defaults to 25.
	Limit param.Field[int64] `json:"limit"`
	// The next page token from a previous response.
	NextPage param.Field[string] `json:"next_page"`
	// Include only credits that have any access on or after the provided date
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V1CustomerCreditListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCreditUpdateEndDateParams struct {
	// RFC 3339 timestamp indicating when access to the credit will end and it will no
	// longer be possible to draw it down (exclusive).
	AccessEndingBefore param.Field[time.Time] `json:"access_ending_before,required" format:"date-time"`
	// ID of the commit to update
	CreditID param.Field[string] `json:"credit_id,required" format:"uuid"`
	// ID of the customer whose credit is to be updated
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r V1CustomerCreditUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
