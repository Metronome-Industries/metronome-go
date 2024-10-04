// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/pagination"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// CreditGrantService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditGrantService] method instead.
type CreditGrantService struct {
	Options []option.RequestOption
}

// NewCreditGrantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCreditGrantService(opts ...option.RequestOption) (r *CreditGrantService) {
	r = &CreditGrantService{}
	r.Options = opts
	return
}

// Create a new credit grant
func (r *CreditGrantService) New(ctx context.Context, body CreditGrantNewParams, opts ...option.RequestOption) (res *CreditGrantNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "credits/createGrant"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List credit grants. This list does not included voided grants.
func (r *CreditGrantService) List(ctx context.Context, params CreditGrantListParams, opts ...option.RequestOption) (res *pagination.CursorPage[CreditGrantListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "credits/listGrants"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// List credit grants. This list does not included voided grants.
func (r *CreditGrantService) ListAutoPaging(ctx context.Context, params CreditGrantListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CreditGrantListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Edit an existing credit grant
func (r *CreditGrantService) Edit(ctx context.Context, body CreditGrantEditParams, opts ...option.RequestOption) (res *CreditGrantEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "credits/editGrant"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all pricing units (known in the API by the legacy term "credit types").
func (r *CreditGrantService) ListCreditTypes(ctx context.Context, query CreditGrantListCreditTypesParams, opts ...option.RequestOption) (res *pagination.CursorPage[CreditGrantListCreditTypesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "credit-types/list"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List all pricing units (known in the API by the legacy term "credit types").
func (r *CreditGrantService) ListCreditTypesAutoPaging(ctx context.Context, query CreditGrantListCreditTypesParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CreditGrantListCreditTypesResponse] {
	return pagination.NewCursorPageAutoPager(r.ListCreditTypes(ctx, query, opts...))
}

// Fetches a list of credit ledger entries. Returns lists of ledgers per customer.
// Ledger entries are returned in chronological order. Ledger entries associated
// with voided credit grants are not included.
func (r *CreditGrantService) ListEntries(ctx context.Context, params CreditGrantListEntriesParams, opts ...option.RequestOption) (res *CreditGrantListEntriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "credits/listEntries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Void a credit grant
func (r *CreditGrantService) Void(ctx context.Context, body CreditGrantVoidParams, opts ...option.RequestOption) (res *CreditGrantVoidResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "credits/voidGrant"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CreditLedgerEntry struct {
	// an amount representing the change to the customer's credit balance
	Amount    float64 `json:"amount,required"`
	CreatedBy string  `json:"created_by,required"`
	// the credit grant this entry is related to
	CreditGrantID string    `json:"credit_grant_id,required" format:"uuid"`
	EffectiveAt   time.Time `json:"effective_at,required" format:"date-time"`
	Reason        string    `json:"reason,required"`
	// the running balance for this credit type at the time of the ledger entry,
	// including all preceding charges
	RunningBalance float64 `json:"running_balance,required"`
	// if this entry is a deduction, the Metronome ID of the invoice where the credit
	// deduction was consumed; if this entry is a grant, the Metronome ID of the
	// invoice where the grant's paid_amount was charged
	InvoiceID string                `json:"invoice_id,nullable" format:"uuid"`
	JSON      creditLedgerEntryJSON `json:"-"`
}

// creditLedgerEntryJSON contains the JSON metadata for the struct
// [CreditLedgerEntry]
type creditLedgerEntryJSON struct {
	Amount         apijson.Field
	CreatedBy      apijson.Field
	CreditGrantID  apijson.Field
	EffectiveAt    apijson.Field
	Reason         apijson.Field
	RunningBalance apijson.Field
	InvoiceID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CreditLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerEntryJSON) RawJSON() string {
	return r.raw
}

type RolloverAmountMaxAmountParam struct {
	// Rollover up to a fixed amount of the original credit grant amount.
	Type param.Field[RolloverAmountMaxAmountType] `json:"type,required"`
	// The maximum amount to rollover.
	Value param.Field[float64] `json:"value,required"`
}

func (r RolloverAmountMaxAmountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RolloverAmountMaxAmountParam) implementsCreditGrantNewParamsRolloverSettingsRolloverAmountUnion() {
}

// Rollover up to a fixed amount of the original credit grant amount.
type RolloverAmountMaxAmountType string

const (
	RolloverAmountMaxAmountTypeMaxAmount RolloverAmountMaxAmountType = "MAX_AMOUNT"
)

func (r RolloverAmountMaxAmountType) IsKnown() bool {
	switch r {
	case RolloverAmountMaxAmountTypeMaxAmount:
		return true
	}
	return false
}

type RolloverAmountMaxPercentageParam struct {
	// Rollover up to a percentage of the original credit grant amount.
	Type param.Field[RolloverAmountMaxPercentageType] `json:"type,required"`
	// The maximum percentage (0-1) of the original credit grant to rollover.
	Value param.Field[float64] `json:"value,required"`
}

func (r RolloverAmountMaxPercentageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RolloverAmountMaxPercentageParam) implementsCreditGrantNewParamsRolloverSettingsRolloverAmountUnion() {
}

// Rollover up to a percentage of the original credit grant amount.
type RolloverAmountMaxPercentageType string

const (
	RolloverAmountMaxPercentageTypeMaxPercentage RolloverAmountMaxPercentageType = "MAX_PERCENTAGE"
)

func (r RolloverAmountMaxPercentageType) IsKnown() bool {
	switch r {
	case RolloverAmountMaxPercentageTypeMaxPercentage:
		return true
	}
	return false
}

type CreditGrantNewResponse struct {
	Data shared.ID                  `json:"data,required"`
	JSON creditGrantNewResponseJSON `json:"-"`
}

// creditGrantNewResponseJSON contains the JSON metadata for the struct
// [CreditGrantNewResponse]
type creditGrantNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditGrantNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantNewResponseJSON) RawJSON() string {
	return r.raw
}

type CreditGrantListResponse struct {
	// the Metronome ID of the credit grant
	ID string `json:"id,required" format:"uuid"`
	// The effective balance of the grant as of the end of the customer's current
	// billing period. Expiration deductions will be included only if the grant expires
	// before the end of the current billing period.
	Balance      CreditGrantListResponseBalance `json:"balance,required"`
	CustomFields map[string]string              `json:"custom_fields,required"`
	// the Metronome ID of the customer
	CustomerID  string              `json:"customer_id,required" format:"uuid"`
	Deductions  []CreditLedgerEntry `json:"deductions,required"`
	EffectiveAt time.Time           `json:"effective_at,required" format:"date-time"`
	ExpiresAt   time.Time           `json:"expires_at,required" format:"date-time"`
	// the amount of credits initially granted
	GrantAmount CreditGrantListResponseGrantAmount `json:"grant_amount,required"`
	Name        string                             `json:"name,required"`
	// the amount paid for this credit grant
	PaidAmount        CreditGrantListResponsePaidAmount `json:"paid_amount,required"`
	PendingDeductions []CreditLedgerEntry               `json:"pending_deductions,required"`
	Priority          float64                           `json:"priority,required"`
	CreditGrantType   string                            `json:"credit_grant_type,nullable"`
	// the Metronome ID of the invoice with the purchase charge for this credit grant,
	// if applicable
	InvoiceID string `json:"invoice_id,nullable" format:"uuid"`
	// The products which these credits will be applied to. (If unspecified, the
	// credits will be applied to charges for all products.)
	Products []CreditGrantListResponseProduct `json:"products"`
	Reason   string                           `json:"reason,nullable"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                      `json:"uniqueness_key,nullable"`
	JSON          creditGrantListResponseJSON `json:"-"`
}

// creditGrantListResponseJSON contains the JSON metadata for the struct
// [CreditGrantListResponse]
type creditGrantListResponseJSON struct {
	ID                apijson.Field
	Balance           apijson.Field
	CustomFields      apijson.Field
	CustomerID        apijson.Field
	Deductions        apijson.Field
	EffectiveAt       apijson.Field
	ExpiresAt         apijson.Field
	GrantAmount       apijson.Field
	Name              apijson.Field
	PaidAmount        apijson.Field
	PendingDeductions apijson.Field
	Priority          apijson.Field
	CreditGrantType   apijson.Field
	InvoiceID         apijson.Field
	Products          apijson.Field
	Reason            apijson.Field
	UniquenessKey     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CreditGrantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantListResponseJSON) RawJSON() string {
	return r.raw
}

// The effective balance of the grant as of the end of the customer's current
// billing period. Expiration deductions will be included only if the grant expires
// before the end of the current billing period.
type CreditGrantListResponseBalance struct {
	// The end_date of the customer's current billing period.
	EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
	// The grant's current balance including all posted deductions. If the grant has
	// expired, this amount will be 0.
	ExcludingPending float64 `json:"excluding_pending,required"`
	// The grant's current balance including all posted and pending deductions. If the
	// grant expires before the end of the customer's current billing period, this
	// amount will be 0.
	IncludingPending float64                            `json:"including_pending,required"`
	JSON             creditGrantListResponseBalanceJSON `json:"-"`
}

// creditGrantListResponseBalanceJSON contains the JSON metadata for the struct
// [CreditGrantListResponseBalance]
type creditGrantListResponseBalanceJSON struct {
	EffectiveAt      apijson.Field
	ExcludingPending apijson.Field
	IncludingPending apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CreditGrantListResponseBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantListResponseBalanceJSON) RawJSON() string {
	return r.raw
}

// the amount of credits initially granted
type CreditGrantListResponseGrantAmount struct {
	Amount float64 `json:"amount,required"`
	// the credit type for the amount granted
	CreditType shared.CreditType                      `json:"credit_type,required"`
	JSON       creditGrantListResponseGrantAmountJSON `json:"-"`
}

// creditGrantListResponseGrantAmountJSON contains the JSON metadata for the struct
// [CreditGrantListResponseGrantAmount]
type creditGrantListResponseGrantAmountJSON struct {
	Amount      apijson.Field
	CreditType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditGrantListResponseGrantAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantListResponseGrantAmountJSON) RawJSON() string {
	return r.raw
}

// the amount paid for this credit grant
type CreditGrantListResponsePaidAmount struct {
	Amount float64 `json:"amount,required"`
	// the credit type for the amount paid
	CreditType shared.CreditType                     `json:"credit_type,required"`
	JSON       creditGrantListResponsePaidAmountJSON `json:"-"`
}

// creditGrantListResponsePaidAmountJSON contains the JSON metadata for the struct
// [CreditGrantListResponsePaidAmount]
type creditGrantListResponsePaidAmountJSON struct {
	Amount      apijson.Field
	CreditType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditGrantListResponsePaidAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantListResponsePaidAmountJSON) RawJSON() string {
	return r.raw
}

type CreditGrantListResponseProduct struct {
	ID   string                             `json:"id,required"`
	Name string                             `json:"name,required"`
	JSON creditGrantListResponseProductJSON `json:"-"`
}

// creditGrantListResponseProductJSON contains the JSON metadata for the struct
// [CreditGrantListResponseProduct]
type creditGrantListResponseProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditGrantListResponseProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantListResponseProductJSON) RawJSON() string {
	return r.raw
}

type CreditGrantEditResponse struct {
	Data shared.ID                   `json:"data,required"`
	JSON creditGrantEditResponseJSON `json:"-"`
}

// creditGrantEditResponseJSON contains the JSON metadata for the struct
// [CreditGrantEditResponse]
type creditGrantEditResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditGrantEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantEditResponseJSON) RawJSON() string {
	return r.raw
}

type CreditGrantListCreditTypesResponse struct {
	ID         string                                 `json:"id" format:"uuid"`
	IsCurrency bool                                   `json:"is_currency"`
	Name       string                                 `json:"name"`
	JSON       creditGrantListCreditTypesResponseJSON `json:"-"`
}

// creditGrantListCreditTypesResponseJSON contains the JSON metadata for the struct
// [CreditGrantListCreditTypesResponse]
type creditGrantListCreditTypesResponseJSON struct {
	ID          apijson.Field
	IsCurrency  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditGrantListCreditTypesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantListCreditTypesResponseJSON) RawJSON() string {
	return r.raw
}

type CreditGrantListEntriesResponse struct {
	Data     []CreditGrantListEntriesResponseData `json:"data,required"`
	NextPage string                               `json:"next_page,required,nullable"`
	JSON     creditGrantListEntriesResponseJSON   `json:"-"`
}

// creditGrantListEntriesResponseJSON contains the JSON metadata for the struct
// [CreditGrantListEntriesResponse]
type creditGrantListEntriesResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditGrantListEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantListEntriesResponseJSON) RawJSON() string {
	return r.raw
}

type CreditGrantListEntriesResponseData struct {
	CustomerID string                                     `json:"customer_id,required" format:"uuid"`
	Ledgers    []CreditGrantListEntriesResponseDataLedger `json:"ledgers,required"`
	JSON       creditGrantListEntriesResponseDataJSON     `json:"-"`
}

// creditGrantListEntriesResponseDataJSON contains the JSON metadata for the struct
// [CreditGrantListEntriesResponseData]
type creditGrantListEntriesResponseDataJSON struct {
	CustomerID  apijson.Field
	Ledgers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditGrantListEntriesResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantListEntriesResponseDataJSON) RawJSON() string {
	return r.raw
}

type CreditGrantListEntriesResponseDataLedger struct {
	CreditType shared.CreditType `json:"credit_type,required"`
	// the effective balances at the end of the specified time window
	EndingBalance   CreditGrantListEntriesResponseDataLedgersEndingBalance   `json:"ending_balance,required"`
	Entries         []CreditLedgerEntry                                      `json:"entries,required"`
	PendingEntries  []CreditLedgerEntry                                      `json:"pending_entries,required"`
	StartingBalance CreditGrantListEntriesResponseDataLedgersStartingBalance `json:"starting_balance,required"`
	JSON            creditGrantListEntriesResponseDataLedgerJSON             `json:"-"`
}

// creditGrantListEntriesResponseDataLedgerJSON contains the JSON metadata for the
// struct [CreditGrantListEntriesResponseDataLedger]
type creditGrantListEntriesResponseDataLedgerJSON struct {
	CreditType      apijson.Field
	EndingBalance   apijson.Field
	Entries         apijson.Field
	PendingEntries  apijson.Field
	StartingBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CreditGrantListEntriesResponseDataLedger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantListEntriesResponseDataLedgerJSON) RawJSON() string {
	return r.raw
}

// the effective balances at the end of the specified time window
type CreditGrantListEntriesResponseDataLedgersEndingBalance struct {
	// the ending_before request parameter (if supplied) or the current billing
	// period's end date
	EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
	// the ending balance, including the balance of all grants that have not expired
	// before the effective_at date and deductions that happened before the
	// effective_at date
	ExcludingPending float64 `json:"excluding_pending,required"`
	// the excluding_pending balance plus any pending invoice deductions and
	// expirations that will happen by the effective_at date
	IncludingPending float64                                                    `json:"including_pending,required"`
	JSON             creditGrantListEntriesResponseDataLedgersEndingBalanceJSON `json:"-"`
}

// creditGrantListEntriesResponseDataLedgersEndingBalanceJSON contains the JSON
// metadata for the struct [CreditGrantListEntriesResponseDataLedgersEndingBalance]
type creditGrantListEntriesResponseDataLedgersEndingBalanceJSON struct {
	EffectiveAt      apijson.Field
	ExcludingPending apijson.Field
	IncludingPending apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CreditGrantListEntriesResponseDataLedgersEndingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantListEntriesResponseDataLedgersEndingBalanceJSON) RawJSON() string {
	return r.raw
}

type CreditGrantListEntriesResponseDataLedgersStartingBalance struct {
	// the starting_on request parameter (if supplied) or the first credit grant's
	// effective_at date
	EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
	// the starting balance, including all posted grants, deductions, and expirations
	// that happened at or before the effective_at timestamp
	ExcludingPending float64 `json:"excluding_pending,required"`
	// the excluding_pending balance plus any pending activity that has not been posted
	// at the time of the query
	IncludingPending float64                                                      `json:"including_pending,required"`
	JSON             creditGrantListEntriesResponseDataLedgersStartingBalanceJSON `json:"-"`
}

// creditGrantListEntriesResponseDataLedgersStartingBalanceJSON contains the JSON
// metadata for the struct
// [CreditGrantListEntriesResponseDataLedgersStartingBalance]
type creditGrantListEntriesResponseDataLedgersStartingBalanceJSON struct {
	EffectiveAt      apijson.Field
	ExcludingPending apijson.Field
	IncludingPending apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CreditGrantListEntriesResponseDataLedgersStartingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantListEntriesResponseDataLedgersStartingBalanceJSON) RawJSON() string {
	return r.raw
}

type CreditGrantVoidResponse struct {
	Data shared.ID                   `json:"data,required"`
	JSON creditGrantVoidResponseJSON `json:"-"`
}

// creditGrantVoidResponseJSON contains the JSON metadata for the struct
// [CreditGrantVoidResponse]
type creditGrantVoidResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditGrantVoidResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGrantVoidResponseJSON) RawJSON() string {
	return r.raw
}

type CreditGrantNewParams struct {
	// the Metronome ID of the customer
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// The credit grant will only apply to usage or charges dated before this timestamp
	ExpiresAt param.Field[time.Time] `json:"expires_at,required" format:"date-time"`
	// the amount of credits granted
	GrantAmount param.Field[CreditGrantNewParamsGrantAmount] `json:"grant_amount,required"`
	// the name of the credit grant as it will appear on invoices
	Name param.Field[string] `json:"name,required"`
	// the amount paid for this credit grant
	PaidAmount      param.Field[CreditGrantNewParamsPaidAmount] `json:"paid_amount,required"`
	Priority        param.Field[float64]                        `json:"priority,required"`
	CreditGrantType param.Field[string]                         `json:"credit_grant_type"`
	// Custom fields to attach to the credit grant.
	CustomFields param.Field[map[string]string] `json:"custom_fields"`
	// The credit grant will only apply to usage or charges dated on or after this
	// timestamp
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// The date to issue an invoice for the paid_amount.
	InvoiceDate param.Field[time.Time] `json:"invoice_date" format:"date-time"`
	// The product(s) which these credits will be applied to. (If unspecified, the
	// credits will be applied to charges for all products.). The array ordering
	// specified here will be used to determine the order in which credits will be
	// applied to invoice line items
	ProductIDs param.Field[[]string] `json:"product_ids" format:"uuid"`
	Reason     param.Field[string]   `json:"reason"`
	// Configure a rollover for this credit grant so if it expires it rolls over a
	// configured amount to a new credit grant. This feature is currently opt-in only.
	// Contact Metronome to be added to the beta.
	RolloverSettings param.Field[CreditGrantNewParamsRolloverSettings] `json:"rollover_settings"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey param.Field[string] `json:"uniqueness_key"`
}

func (r CreditGrantNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// the amount of credits granted
type CreditGrantNewParamsGrantAmount struct {
	Amount param.Field[float64] `json:"amount,required"`
	// the ID of the pricing unit to be used
	CreditTypeID param.Field[string] `json:"credit_type_id,required" format:"uuid"`
}

func (r CreditGrantNewParamsGrantAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// the amount paid for this credit grant
type CreditGrantNewParamsPaidAmount struct {
	Amount param.Field[float64] `json:"amount,required"`
	// the ID of the pricing unit to be used
	CreditTypeID param.Field[string] `json:"credit_type_id,required" format:"uuid"`
}

func (r CreditGrantNewParamsPaidAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a rollover for this credit grant so if it expires it rolls over a
// configured amount to a new credit grant. This feature is currently opt-in only.
// Contact Metronome to be added to the beta.
type CreditGrantNewParamsRolloverSettings struct {
	// The date to expire the rollover credits.
	ExpiresAt param.Field[time.Time] `json:"expires_at,required" format:"date-time"`
	// The priority to give the rollover credit grant that gets created when a rollover
	// happens.
	Priority param.Field[float64] `json:"priority,required"`
	// Specify how much to rollover to the rollover credit grant
	RolloverAmount param.Field[CreditGrantNewParamsRolloverSettingsRolloverAmountUnion] `json:"rollover_amount,required"`
}

func (r CreditGrantNewParamsRolloverSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how much to rollover to the rollover credit grant
type CreditGrantNewParamsRolloverSettingsRolloverAmount struct {
	// Rollover up to a percentage of the original credit grant amount.
	Type param.Field[CreditGrantNewParamsRolloverSettingsRolloverAmountType] `json:"type,required"`
	// The maximum percentage (0-1) of the original credit grant to rollover.
	Value param.Field[float64] `json:"value,required"`
}

func (r CreditGrantNewParamsRolloverSettingsRolloverAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreditGrantNewParamsRolloverSettingsRolloverAmount) implementsCreditGrantNewParamsRolloverSettingsRolloverAmountUnion() {
}

// Specify how much to rollover to the rollover credit grant
//
// Satisfied by [RolloverAmountMaxPercentageParam], [RolloverAmountMaxAmountParam],
// [CreditGrantNewParamsRolloverSettingsRolloverAmount].
type CreditGrantNewParamsRolloverSettingsRolloverAmountUnion interface {
	implementsCreditGrantNewParamsRolloverSettingsRolloverAmountUnion()
}

// Rollover up to a percentage of the original credit grant amount.
type CreditGrantNewParamsRolloverSettingsRolloverAmountType string

const (
	CreditGrantNewParamsRolloverSettingsRolloverAmountTypeMaxPercentage CreditGrantNewParamsRolloverSettingsRolloverAmountType = "MAX_PERCENTAGE"
	CreditGrantNewParamsRolloverSettingsRolloverAmountTypeMaxAmount     CreditGrantNewParamsRolloverSettingsRolloverAmountType = "MAX_AMOUNT"
)

func (r CreditGrantNewParamsRolloverSettingsRolloverAmountType) IsKnown() bool {
	switch r {
	case CreditGrantNewParamsRolloverSettingsRolloverAmountTypeMaxPercentage, CreditGrantNewParamsRolloverSettingsRolloverAmountTypeMaxAmount:
		return true
	}
	return false
}

type CreditGrantListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// An array of credit grant IDs. If this is specified, neither credit_type_ids nor
	// customer_ids may be specified.
	CreditGrantIDs param.Field[[]string] `json:"credit_grant_ids" format:"uuid"`
	// An array of credit type IDs. This must not be specified if credit_grant_ids is
	// specified.
	CreditTypeIDs param.Field[[]string] `json:"credit_type_ids" format:"uuid"`
	// An array of Metronome customer IDs. This must not be specified if
	// credit_grant_ids is specified.
	CustomerIDs param.Field[[]string] `json:"customer_ids" format:"uuid"`
	// Only return credit grants that are effective before this timestamp (exclusive).
	EffectiveBefore param.Field[time.Time] `json:"effective_before" format:"date-time"`
	// Only return credit grants that expire at or after this timestamp.
	NotExpiringBefore param.Field[time.Time] `json:"not_expiring_before" format:"date-time"`
}

func (r CreditGrantListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [CreditGrantListParams]'s query parameters as `url.Values`.
func (r CreditGrantListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CreditGrantEditParams struct {
	// the ID of the credit grant
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// the updated credit grant type
	CreditGrantType param.Field[string] `json:"credit_grant_type"`
	// the updated expiration date for the credit grant
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// the updated name for the credit grant
	Name param.Field[string] `json:"name"`
}

func (r CreditGrantEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreditGrantListCreditTypesParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [CreditGrantListCreditTypesParams]'s query parameters as
// `url.Values`.
func (r CreditGrantListCreditTypesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CreditGrantListEntriesParams struct {
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// A list of Metronome credit type IDs to fetch ledger entries for. If absent,
	// ledger entries for all credit types will be returned.
	CreditTypeIDs param.Field[[]string] `json:"credit_type_ids" format:"uuid"`
	// A list of Metronome customer IDs to fetch ledger entries for. If absent, ledger
	// entries for all customers will be returned.
	CustomerIDs param.Field[[]string] `json:"customer_ids" format:"uuid"`
	// If supplied, ledger entries will only be returned with an effective_at before
	// this time. This timestamp must not be in the future. If no timestamp is
	// supplied, all entries up to the start of the customer's next billing period will
	// be returned.
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// If supplied, only ledger entries effective at or after this time will be
	// returned.
	StartingOn param.Field[time.Time] `json:"starting_on" format:"date-time"`
}

func (r CreditGrantListEntriesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [CreditGrantListEntriesParams]'s query parameters as
// `url.Values`.
func (r CreditGrantListEntriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CreditGrantVoidParams struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// If true, resets the uniqueness key on this grant so it can be re-used
	ReleaseUniquenessKey param.Field[bool] `json:"release_uniqueness_key"`
	// If true, void the purchase invoice associated with the grant
	VoidCreditPurchaseInvoice param.Field[bool] `json:"void_credit_purchase_invoice"`
}

func (r CreditGrantVoidParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
