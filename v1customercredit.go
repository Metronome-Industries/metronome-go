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
func (r *V1CustomerCreditService) List(ctx context.Context, body V1CustomerCreditListParams, opts ...option.RequestOption) (res *pagination.BodyCursorPage[V1CustomerCreditListResponse], err error) {
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
func (r *V1CustomerCreditService) ListAutoPaging(ctx context.Context, body V1CustomerCreditListParams, opts ...option.RequestOption) *pagination.BodyCursorPageAutoPager[V1CustomerCreditListResponse] {
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
	Data V1CustomerCreditNewResponseData `json:"data,required"`
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

type V1CustomerCreditNewResponseData struct {
	ID   string                              `json:"id,required" format:"uuid"`
	JSON v1CustomerCreditNewResponseDataJSON `json:"-"`
}

// v1CustomerCreditNewResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerCreditNewResponseData]
type v1CustomerCreditNewResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponse struct {
	ID      string                              `json:"id,required" format:"uuid"`
	Product V1CustomerCreditListResponseProduct `json:"product,required"`
	Type    V1CustomerCreditListResponseType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        V1CustomerCreditListResponseAccessSchedule `json:"access_schedule"`
	ApplicableContractIDs []string                                   `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string                                   `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                                   `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance  float64                              `json:"balance"`
	Contract V1CustomerCreditListResponseContract `json:"contract"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration V1CustomerCreditListResponseHierarchyConfiguration `json:"hierarchy_configuration"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []V1CustomerCreditListResponseLedger `json:"ledger"`
	Name   string                               `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64                              `json:"priority"`
	RateType V1CustomerCreditListResponseRateType `json:"rate_type"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []V1CustomerCreditListResponseSpecifier `json:"specifiers"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey string                           `json:"uniqueness_key"`
	JSON          v1CustomerCreditListResponseJSON `json:"-"`
}

// v1CustomerCreditListResponseJSON contains the JSON metadata for the struct
// [V1CustomerCreditListResponse]
type v1CustomerCreditListResponseJSON struct {
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

func (r *V1CustomerCreditListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseProduct struct {
	ID   string                                  `json:"id,required" format:"uuid"`
	Name string                                  `json:"name,required"`
	JSON v1CustomerCreditListResponseProductJSON `json:"-"`
}

// v1CustomerCreditListResponseProductJSON contains the JSON metadata for the
// struct [V1CustomerCreditListResponseProduct]
type v1CustomerCreditListResponseProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseProductJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseType string

const (
	V1CustomerCreditListResponseTypeCredit V1CustomerCreditListResponseType = "CREDIT"
)

func (r V1CustomerCreditListResponseType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseTypeCredit:
		return true
	}
	return false
}

// The schedule that the customer will gain access to the credits.
type V1CustomerCreditListResponseAccessSchedule struct {
	ScheduleItems []V1CustomerCreditListResponseAccessScheduleScheduleItem `json:"schedule_items,required"`
	CreditType    V1CustomerCreditListResponseAccessScheduleCreditType     `json:"credit_type"`
	JSON          v1CustomerCreditListResponseAccessScheduleJSON           `json:"-"`
}

// v1CustomerCreditListResponseAccessScheduleJSON contains the JSON metadata for
// the struct [V1CustomerCreditListResponseAccessSchedule]
type v1CustomerCreditListResponseAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseAccessScheduleScheduleItem struct {
	ID           string                                                     `json:"id,required" format:"uuid"`
	Amount       float64                                                    `json:"amount,required"`
	EndingBefore time.Time                                                  `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                                                  `json:"starting_at,required" format:"date-time"`
	JSON         v1CustomerCreditListResponseAccessScheduleScheduleItemJSON `json:"-"`
}

// v1CustomerCreditListResponseAccessScheduleScheduleItemJSON contains the JSON
// metadata for the struct [V1CustomerCreditListResponseAccessScheduleScheduleItem]
type v1CustomerCreditListResponseAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseAccessScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseAccessScheduleCreditType struct {
	ID   string                                                   `json:"id,required" format:"uuid"`
	Name string                                                   `json:"name,required"`
	JSON v1CustomerCreditListResponseAccessScheduleCreditTypeJSON `json:"-"`
}

// v1CustomerCreditListResponseAccessScheduleCreditTypeJSON contains the JSON
// metadata for the struct [V1CustomerCreditListResponseAccessScheduleCreditType]
type v1CustomerCreditListResponseAccessScheduleCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseAccessScheduleCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseAccessScheduleCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseContract struct {
	ID   string                                   `json:"id,required" format:"uuid"`
	JSON v1CustomerCreditListResponseContractJSON `json:"-"`
}

// v1CustomerCreditListResponseContractJSON contains the JSON metadata for the
// struct [V1CustomerCreditListResponseContract]
type v1CustomerCreditListResponseContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseContractJSON) RawJSON() string {
	return r.raw
}

// Optional configuration for credit hierarchy access control
type V1CustomerCreditListResponseHierarchyConfiguration struct {
	ChildAccess V1CustomerCreditListResponseHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        v1CustomerCreditListResponseHierarchyConfigurationJSON        `json:"-"`
}

// v1CustomerCreditListResponseHierarchyConfigurationJSON contains the JSON
// metadata for the struct [V1CustomerCreditListResponseHierarchyConfiguration]
type v1CustomerCreditListResponseHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseHierarchyConfigurationChildAccess struct {
	Type V1CustomerCreditListResponseHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                       `json:"contract_ids"`
	JSON        v1CustomerCreditListResponseHierarchyConfigurationChildAccessJSON `json:"-"`
	union       V1CustomerCreditListResponseHierarchyConfigurationChildAccessUnion
}

// v1CustomerCreditListResponseHierarchyConfigurationChildAccessJSON contains the
// JSON metadata for the struct
// [V1CustomerCreditListResponseHierarchyConfigurationChildAccess]
type v1CustomerCreditListResponseHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v1CustomerCreditListResponseHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *V1CustomerCreditListResponseHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = V1CustomerCreditListResponseHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [V1CustomerCreditListResponseHierarchyConfigurationChildAccessUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V1CustomerCreditListResponseHierarchyConfigurationChildAccessType],
// [V1CustomerCreditListResponseHierarchyConfigurationChildAccessType],
// [V1CustomerCreditListResponseHierarchyConfigurationChildAccessObject].
func (r V1CustomerCreditListResponseHierarchyConfigurationChildAccess) AsUnion() V1CustomerCreditListResponseHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [V1CustomerCreditListResponseHierarchyConfigurationChildAccessType],
// [V1CustomerCreditListResponseHierarchyConfigurationChildAccessType] or
// [V1CustomerCreditListResponseHierarchyConfigurationChildAccessObject].
type V1CustomerCreditListResponseHierarchyConfigurationChildAccessUnion interface {
	implementsV1CustomerCreditListResponseHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1CustomerCreditListResponseHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseHierarchyConfigurationChildAccessType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseHierarchyConfigurationChildAccessType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseHierarchyConfigurationChildAccessObject{}),
		},
	)
}

type V1CustomerCreditListResponseHierarchyConfigurationChildAccessType struct {
	Type V1CustomerCreditListResponseHierarchyConfigurationChildAccessTypeType `json:"type,required"`
	JSON v1CustomerCreditListResponseHierarchyConfigurationChildAccessTypeJSON `json:"-"`
}

// v1CustomerCreditListResponseHierarchyConfigurationChildAccessTypeJSON contains
// the JSON metadata for the struct
// [V1CustomerCreditListResponseHierarchyConfigurationChildAccessType]
type v1CustomerCreditListResponseHierarchyConfigurationChildAccessTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseHierarchyConfigurationChildAccessType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseHierarchyConfigurationChildAccessTypeJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCreditListResponseHierarchyConfigurationChildAccessType) implementsV1CustomerCreditListResponseHierarchyConfigurationChildAccess() {
}

type V1CustomerCreditListResponseHierarchyConfigurationChildAccessTypeType string

const (
	V1CustomerCreditListResponseHierarchyConfigurationChildAccessTypeTypeAll V1CustomerCreditListResponseHierarchyConfigurationChildAccessTypeType = "ALL"
)

func (r V1CustomerCreditListResponseHierarchyConfigurationChildAccessTypeType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseHierarchyConfigurationChildAccessTypeTypeAll:
		return true
	}
	return false
}

type V1CustomerCreditListResponseHierarchyConfigurationChildAccessObject struct {
	ContractIDs []string                                                                `json:"contract_ids,required" format:"uuid"`
	Type        V1CustomerCreditListResponseHierarchyConfigurationChildAccessObjectType `json:"type,required"`
	JSON        v1CustomerCreditListResponseHierarchyConfigurationChildAccessObjectJSON `json:"-"`
}

// v1CustomerCreditListResponseHierarchyConfigurationChildAccessObjectJSON contains
// the JSON metadata for the struct
// [V1CustomerCreditListResponseHierarchyConfigurationChildAccessObject]
type v1CustomerCreditListResponseHierarchyConfigurationChildAccessObjectJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseHierarchyConfigurationChildAccessObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseHierarchyConfigurationChildAccessObjectJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCreditListResponseHierarchyConfigurationChildAccessObject) implementsV1CustomerCreditListResponseHierarchyConfigurationChildAccess() {
}

type V1CustomerCreditListResponseHierarchyConfigurationChildAccessObjectType string

const (
	V1CustomerCreditListResponseHierarchyConfigurationChildAccessObjectTypeContractIDs V1CustomerCreditListResponseHierarchyConfigurationChildAccessObjectType = "CONTRACT_IDS"
)

func (r V1CustomerCreditListResponseHierarchyConfigurationChildAccessObjectType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseHierarchyConfigurationChildAccessObjectTypeContractIDs:
		return true
	}
	return false
}

type V1CustomerCreditListResponseLedger struct {
	Amount     float64                                `json:"amount,required"`
	Timestamp  time.Time                              `json:"timestamp,required" format:"date-time"`
	Type       V1CustomerCreditListResponseLedgerType `json:"type,required"`
	ContractID string                                 `json:"contract_id" format:"uuid"`
	InvoiceID  string                                 `json:"invoice_id" format:"uuid"`
	Reason     string                                 `json:"reason"`
	SegmentID  string                                 `json:"segment_id" format:"uuid"`
	JSON       v1CustomerCreditListResponseLedgerJSON `json:"-"`
	union      V1CustomerCreditListResponseLedgerUnion
}

// v1CustomerCreditListResponseLedgerJSON contains the JSON metadata for the struct
// [V1CustomerCreditListResponseLedger]
type v1CustomerCreditListResponseLedgerJSON struct {
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

func (r v1CustomerCreditListResponseLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *V1CustomerCreditListResponseLedger) UnmarshalJSON(data []byte) (err error) {
	*r = V1CustomerCreditListResponseLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V1CustomerCreditListResponseLedgerUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V1CustomerCreditListResponseLedgerObject],
// [V1CustomerCreditListResponseLedgerObject],
// [V1CustomerCreditListResponseLedgerObject],
// [V1CustomerCreditListResponseLedgerObject],
// [V1CustomerCreditListResponseLedgerObject],
// [V1CustomerCreditListResponseLedgerObject],
// [V1CustomerCreditListResponseLedgerObject].
func (r V1CustomerCreditListResponseLedger) AsUnion() V1CustomerCreditListResponseLedgerUnion {
	return r.union
}

// Union satisfied by [V1CustomerCreditListResponseLedgerObject],
// [V1CustomerCreditListResponseLedgerObject],
// [V1CustomerCreditListResponseLedgerObject],
// [V1CustomerCreditListResponseLedgerObject],
// [V1CustomerCreditListResponseLedgerObject],
// [V1CustomerCreditListResponseLedgerObject] or
// [V1CustomerCreditListResponseLedgerObject].
type V1CustomerCreditListResponseLedgerUnion interface {
	implementsV1CustomerCreditListResponseLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1CustomerCreditListResponseLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseLedgerObject{}),
		},
	)
}

type V1CustomerCreditListResponseLedgerObject struct {
	Amount    float64                                      `json:"amount,required"`
	SegmentID string                                       `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                    `json:"timestamp,required" format:"date-time"`
	Type      V1CustomerCreditListResponseLedgerObjectType `json:"type,required"`
	JSON      v1CustomerCreditListResponseLedgerObjectJSON `json:"-"`
}

// v1CustomerCreditListResponseLedgerObjectJSON contains the JSON metadata for the
// struct [V1CustomerCreditListResponseLedgerObject]
type v1CustomerCreditListResponseLedgerObjectJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseLedgerObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseLedgerObjectJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCreditListResponseLedgerObject) implementsV1CustomerCreditListResponseLedger() {}

type V1CustomerCreditListResponseLedgerObjectType string

const (
	V1CustomerCreditListResponseLedgerObjectTypeCreditSegmentStart V1CustomerCreditListResponseLedgerObjectType = "CREDIT_SEGMENT_START"
)

func (r V1CustomerCreditListResponseLedgerObjectType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseLedgerObjectTypeCreditSegmentStart:
		return true
	}
	return false
}

type V1CustomerCreditListResponseLedgerType string

const (
	V1CustomerCreditListResponseLedgerTypeCreditSegmentStart              V1CustomerCreditListResponseLedgerType = "CREDIT_SEGMENT_START"
	V1CustomerCreditListResponseLedgerTypeCreditAutomatedInvoiceDeduction V1CustomerCreditListResponseLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	V1CustomerCreditListResponseLedgerTypeCreditExpiration                V1CustomerCreditListResponseLedgerType = "CREDIT_EXPIRATION"
	V1CustomerCreditListResponseLedgerTypeCreditCanceled                  V1CustomerCreditListResponseLedgerType = "CREDIT_CANCELED"
	V1CustomerCreditListResponseLedgerTypeCreditCredited                  V1CustomerCreditListResponseLedgerType = "CREDIT_CREDITED"
	V1CustomerCreditListResponseLedgerTypeCreditManual                    V1CustomerCreditListResponseLedgerType = "CREDIT_MANUAL"
	V1CustomerCreditListResponseLedgerTypeCreditSeatBasedAdjustment       V1CustomerCreditListResponseLedgerType = "CREDIT_SEAT_BASED_ADJUSTMENT"
)

func (r V1CustomerCreditListResponseLedgerType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseLedgerTypeCreditSegmentStart, V1CustomerCreditListResponseLedgerTypeCreditAutomatedInvoiceDeduction, V1CustomerCreditListResponseLedgerTypeCreditExpiration, V1CustomerCreditListResponseLedgerTypeCreditCanceled, V1CustomerCreditListResponseLedgerTypeCreditCredited, V1CustomerCreditListResponseLedgerTypeCreditManual, V1CustomerCreditListResponseLedgerTypeCreditSeatBasedAdjustment:
		return true
	}
	return false
}

type V1CustomerCreditListResponseRateType string

const (
	V1CustomerCreditListResponseRateTypeCommitRate V1CustomerCreditListResponseRateType = "COMMIT_RATE"
	V1CustomerCreditListResponseRateTypeListRate   V1CustomerCreditListResponseRateType = "LIST_RATE"
)

func (r V1CustomerCreditListResponseRateType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseRateTypeCommitRate, V1CustomerCreditListResponseRateTypeListRate:
		return true
	}
	return false
}

type V1CustomerCreditListResponseSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                  `json:"product_tags"`
	JSON        v1CustomerCreditListResponseSpecifierJSON `json:"-"`
}

// v1CustomerCreditListResponseSpecifierJSON contains the JSON metadata for the
// struct [V1CustomerCreditListResponseSpecifier]
type v1CustomerCreditListResponseSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseSpecifierJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditUpdateEndDateResponse struct {
	Data V1CustomerCreditUpdateEndDateResponseData `json:"data,required"`
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

type V1CustomerCreditUpdateEndDateResponseData struct {
	ID   string                                        `json:"id,required" format:"uuid"`
	JSON v1CustomerCreditUpdateEndDateResponseDataJSON `json:"-"`
}

// v1CustomerCreditUpdateEndDateResponseDataJSON contains the JSON metadata for the
// struct [V1CustomerCreditUpdateEndDateResponseData]
type v1CustomerCreditUpdateEndDateResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditUpdateEndDateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditUpdateEndDateResponseDataJSON) RawJSON() string {
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
	Specifiers param.Field[[]V1CustomerCreditNewParamsSpecifier] `json:"specifiers"`
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

type V1CustomerCreditNewParamsSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V1CustomerCreditNewParamsSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
