// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v2/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v2/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/v2/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/v2/packages/param"
	"github.com/Metronome-Industries/metronome-go/v2/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/v2/shared"
)

// V1CreditGrantService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CreditGrantService] method instead.
type V1CreditGrantService struct {
	Options []option.RequestOption
}

// NewV1CreditGrantService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CreditGrantService(opts ...option.RequestOption) (r V1CreditGrantService) {
	r = V1CreditGrantService{}
	r.Options = opts
	return
}

// Create a new credit grant
func (r *V1CreditGrantService) New(ctx context.Context, body V1CreditGrantNewParams, opts ...option.RequestOption) (res *V1CreditGrantNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/credits/createGrant"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List credit grants. This list does not included voided grants.
func (r *V1CreditGrantService) List(ctx context.Context, params V1CreditGrantListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CreditGrantListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/credits/listGrants"
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
func (r *V1CreditGrantService) ListAutoPaging(ctx context.Context, params V1CreditGrantListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CreditGrantListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Edit an existing credit grant
func (r *V1CreditGrantService) Edit(ctx context.Context, body V1CreditGrantEditParams, opts ...option.RequestOption) (res *V1CreditGrantEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/credits/editGrant"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a list of credit ledger entries. Returns lists of ledgers per customer.
// Ledger entries are returned in chronological order. Ledger entries associated
// with voided credit grants are not included.
func (r *V1CreditGrantService) ListEntries(ctx context.Context, params V1CreditGrantListEntriesParams, opts ...option.RequestOption) (res *pagination.CursorPageWithoutLimit[V1CreditGrantListEntriesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/credits/listEntries"
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

// Fetches a list of credit ledger entries. Returns lists of ledgers per customer.
// Ledger entries are returned in chronological order. Ledger entries associated
// with voided credit grants are not included.
func (r *V1CreditGrantService) ListEntriesAutoPaging(ctx context.Context, params V1CreditGrantListEntriesParams, opts ...option.RequestOption) *pagination.CursorPageWithoutLimitAutoPager[V1CreditGrantListEntriesResponse] {
	return pagination.NewCursorPageWithoutLimitAutoPager(r.ListEntries(ctx, params, opts...))
}

// Void a credit grant
func (r *V1CreditGrantService) Void(ctx context.Context, body V1CreditGrantVoidParams, opts ...option.RequestOption) (res *V1CreditGrantVoidResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/credits/voidGrant"
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
	InvoiceID string `json:"invoice_id,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount         respjson.Field
		CreatedBy      respjson.Field
		CreditGrantID  respjson.Field
		EffectiveAt    respjson.Field
		Reason         respjson.Field
		RunningBalance respjson.Field
		InvoiceID      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditLedgerEntry) RawJSON() string { return r.JSON.raw }
func (r *CreditLedgerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, Value are required.
type RolloverAmountMaxAmountParam struct {
	// Rollover up to a fixed amount of the original credit grant amount.
	//
	// Any of "MAX_AMOUNT".
	Type RolloverAmountMaxAmountType `json:"type,omitzero,required"`
	// The maximum amount to rollover.
	Value float64 `json:"value,required"`
	paramObj
}

func (r RolloverAmountMaxAmountParam) MarshalJSON() (data []byte, err error) {
	type shadow RolloverAmountMaxAmountParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RolloverAmountMaxAmountParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rollover up to a fixed amount of the original credit grant amount.
type RolloverAmountMaxAmountType string

const (
	RolloverAmountMaxAmountTypeMaxAmount RolloverAmountMaxAmountType = "MAX_AMOUNT"
)

// The properties Type, Value are required.
type RolloverAmountMaxPercentageParam struct {
	// Rollover up to a percentage of the original credit grant amount.
	//
	// Any of "MAX_PERCENTAGE".
	Type RolloverAmountMaxPercentageType `json:"type,omitzero,required"`
	// The maximum percentage (0-1) of the original credit grant to rollover.
	Value float64 `json:"value,required"`
	paramObj
}

func (r RolloverAmountMaxPercentageParam) MarshalJSON() (data []byte, err error) {
	type shadow RolloverAmountMaxPercentageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RolloverAmountMaxPercentageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rollover up to a percentage of the original credit grant amount.
type RolloverAmountMaxPercentageType string

const (
	RolloverAmountMaxPercentageTypeMaxPercentage RolloverAmountMaxPercentageType = "MAX_PERCENTAGE"
)

type V1CreditGrantNewResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditGrantListResponse struct {
	// the Metronome ID of the credit grant
	ID string `json:"id,required" format:"uuid"`
	// The effective balance of the grant as of the end of the customer's current
	// billing period. Expiration deductions will be included only if the grant expires
	// before the end of the current billing period.
	Balance V1CreditGrantListResponseBalance `json:"balance,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,required"`
	// the Metronome ID of the customer
	CustomerID  string              `json:"customer_id,required" format:"uuid"`
	Deductions  []CreditLedgerEntry `json:"deductions,required"`
	EffectiveAt time.Time           `json:"effective_at,required" format:"date-time"`
	ExpiresAt   time.Time           `json:"expires_at,required" format:"date-time"`
	// the amount of credits initially granted
	GrantAmount V1CreditGrantListResponseGrantAmount `json:"grant_amount,required"`
	Name        string                               `json:"name,required"`
	// the amount paid for this credit grant
	PaidAmount        V1CreditGrantListResponsePaidAmount `json:"paid_amount,required"`
	PendingDeductions []CreditLedgerEntry                 `json:"pending_deductions,required"`
	Priority          float64                             `json:"priority,required"`
	CreditGrantType   string                              `json:"credit_grant_type,nullable"`
	// the Metronome ID of the invoice with the purchase charge for this credit grant,
	// if applicable
	InvoiceID string `json:"invoice_id,nullable" format:"uuid"`
	// The products which these credits will be applied to. (If unspecified, the
	// credits will be applied to charges for all products.)
	Products []V1CreditGrantListResponseProduct `json:"products"`
	Reason   string                             `json:"reason,nullable"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string `json:"uniqueness_key,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Balance           respjson.Field
		CustomFields      respjson.Field
		CustomerID        respjson.Field
		Deductions        respjson.Field
		EffectiveAt       respjson.Field
		ExpiresAt         respjson.Field
		GrantAmount       respjson.Field
		Name              respjson.Field
		PaidAmount        respjson.Field
		PendingDeductions respjson.Field
		Priority          respjson.Field
		CreditGrantType   respjson.Field
		InvoiceID         respjson.Field
		Products          respjson.Field
		Reason            respjson.Field
		UniquenessKey     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The effective balance of the grant as of the end of the customer's current
// billing period. Expiration deductions will be included only if the grant expires
// before the end of the current billing period.
type V1CreditGrantListResponseBalance struct {
	// The end_date of the customer's current billing period.
	EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
	// The grant's current balance including all posted deductions. If the grant has
	// expired, this amount will be 0.
	ExcludingPending float64 `json:"excluding_pending,required"`
	// The grant's current balance including all posted and pending deductions. If the
	// grant expires before the end of the customer's current billing period, this
	// amount will be 0.
	IncludingPending float64 `json:"including_pending,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EffectiveAt      respjson.Field
		ExcludingPending respjson.Field
		IncludingPending respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantListResponseBalance) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListResponseBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// the amount of credits initially granted
type V1CreditGrantListResponseGrantAmount struct {
	Amount float64 `json:"amount,required"`
	// the credit type for the amount granted
	CreditType shared.CreditTypeData `json:"credit_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		CreditType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantListResponseGrantAmount) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListResponseGrantAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// the amount paid for this credit grant
type V1CreditGrantListResponsePaidAmount struct {
	Amount float64 `json:"amount,required"`
	// the credit type for the amount paid
	CreditType shared.CreditTypeData `json:"credit_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		CreditType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantListResponsePaidAmount) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListResponsePaidAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditGrantListResponseProduct struct {
	ID   string `json:"id,required"`
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
func (r V1CreditGrantListResponseProduct) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListResponseProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditGrantEditResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantEditResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantEditResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditGrantListEntriesResponse struct {
	CustomerID string                                   `json:"customer_id,required" format:"uuid"`
	Ledgers    []V1CreditGrantListEntriesResponseLedger `json:"ledgers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID  respjson.Field
		Ledgers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantListEntriesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListEntriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditGrantListEntriesResponseLedger struct {
	CreditType shared.CreditTypeData `json:"credit_type,required"`
	// the effective balances at the end of the specified time window
	EndingBalance   V1CreditGrantListEntriesResponseLedgerEndingBalance   `json:"ending_balance,required"`
	Entries         []CreditLedgerEntry                                   `json:"entries,required"`
	PendingEntries  []CreditLedgerEntry                                   `json:"pending_entries,required"`
	StartingBalance V1CreditGrantListEntriesResponseLedgerStartingBalance `json:"starting_balance,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType      respjson.Field
		EndingBalance   respjson.Field
		Entries         respjson.Field
		PendingEntries  respjson.Field
		StartingBalance respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantListEntriesResponseLedger) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListEntriesResponseLedger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// the effective balances at the end of the specified time window
type V1CreditGrantListEntriesResponseLedgerEndingBalance struct {
	// the ending_before request parameter (if supplied) or the current billing
	// period's end date
	EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
	// the ending balance, including the balance of all grants that have not expired
	// before the effective_at date and deductions that happened before the
	// effective_at date
	ExcludingPending float64 `json:"excluding_pending,required"`
	// the excluding_pending balance plus any pending invoice deductions and
	// expirations that will happen by the effective_at date
	IncludingPending float64 `json:"including_pending,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EffectiveAt      respjson.Field
		ExcludingPending respjson.Field
		IncludingPending respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantListEntriesResponseLedgerEndingBalance) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListEntriesResponseLedgerEndingBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditGrantListEntriesResponseLedgerStartingBalance struct {
	// the starting_on request parameter (if supplied) or the first credit grant's
	// effective_at date
	EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
	// the starting balance, including all posted grants, deductions, and expirations
	// that happened at or before the effective_at timestamp
	ExcludingPending float64 `json:"excluding_pending,required"`
	// the excluding_pending balance plus any pending activity that has not been posted
	// at the time of the query
	IncludingPending float64 `json:"including_pending,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EffectiveAt      respjson.Field
		ExcludingPending respjson.Field
		IncludingPending respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantListEntriesResponseLedgerStartingBalance) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListEntriesResponseLedgerStartingBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditGrantVoidResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantVoidResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantVoidResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditGrantNewParams struct {
	// the Metronome ID of the customer
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// The credit grant will only apply to usage or charges dated before this timestamp
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// the amount of credits granted
	GrantAmount V1CreditGrantNewParamsGrantAmount `json:"grant_amount,omitzero,required"`
	// the name of the credit grant as it will appear on invoices
	Name string `json:"name,required"`
	// the amount paid for this credit grant
	PaidAmount      V1CreditGrantNewParamsPaidAmount `json:"paid_amount,omitzero,required"`
	Priority        float64                          `json:"priority,required"`
	CreditGrantType param.Opt[string]                `json:"credit_grant_type,omitzero"`
	// The credit grant will only apply to usage or charges dated on or after this
	// timestamp
	EffectiveAt param.Opt[time.Time] `json:"effective_at,omitzero" format:"date-time"`
	// The date to issue an invoice for the paid_amount.
	InvoiceDate param.Opt[time.Time] `json:"invoice_date,omitzero" format:"date-time"`
	Reason      param.Opt[string]    `json:"reason,omitzero"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey param.Opt[string] `json:"uniqueness_key,omitzero"`
	// Custom fields to attach to the credit grant.
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// The product(s) which these credits will be applied to. (If unspecified, the
	// credits will be applied to charges for all products.). The array ordering
	// specified here will be used to determine the order in which credits will be
	// applied to invoice line items
	ProductIDs []string `json:"product_ids,omitzero" format:"uuid"`
	// Configure a rollover for this credit grant so if it expires it rolls over a
	// configured amount to a new credit grant. This feature is currently opt-in only.
	// Contact Metronome to be added to the beta.
	RolloverSettings V1CreditGrantNewParamsRolloverSettings `json:"rollover_settings,omitzero"`
	paramObj
}

func (r V1CreditGrantNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// the amount of credits granted
//
// The properties Amount, CreditTypeID are required.
type V1CreditGrantNewParamsGrantAmount struct {
	Amount float64 `json:"amount,required"`
	// the ID of the pricing unit to be used. Defaults to USD (cents) if not passed.
	CreditTypeID string `json:"credit_type_id,required" format:"uuid"`
	paramObj
}

func (r V1CreditGrantNewParamsGrantAmount) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantNewParamsGrantAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantNewParamsGrantAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// the amount paid for this credit grant
//
// The properties Amount, CreditTypeID are required.
type V1CreditGrantNewParamsPaidAmount struct {
	Amount float64 `json:"amount,required"`
	// the ID of the pricing unit to be used. Defaults to USD (cents) if not passed.
	CreditTypeID string `json:"credit_type_id,required" format:"uuid"`
	paramObj
}

func (r V1CreditGrantNewParamsPaidAmount) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantNewParamsPaidAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantNewParamsPaidAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a rollover for this credit grant so if it expires it rolls over a
// configured amount to a new credit grant. This feature is currently opt-in only.
// Contact Metronome to be added to the beta.
//
// The properties ExpiresAt, Priority, RolloverAmount are required.
type V1CreditGrantNewParamsRolloverSettings struct {
	// The date to expire the rollover credits.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// The priority to give the rollover credit grant that gets created when a rollover
	// happens.
	Priority float64 `json:"priority,required"`
	// Specify how much to rollover to the rollover credit grant
	RolloverAmount V1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion `json:"rollover_amount,omitzero,required"`
	paramObj
}

func (r V1CreditGrantNewParamsRolloverSettings) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantNewParamsRolloverSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantNewParamsRolloverSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion struct {
	OfRolloverAmountMaxPercentage *RolloverAmountMaxPercentageParam `json:",omitzero,inline"`
	OfRolloverAmountMaxAmount     *RolloverAmountMaxAmountParam     `json:",omitzero,inline"`
	paramUnion
}

func (u V1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRolloverAmountMaxPercentage, u.OfRolloverAmountMaxAmount)
}
func (u *V1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion) asAny() any {
	if !param.IsOmitted(u.OfRolloverAmountMaxPercentage) {
		return u.OfRolloverAmountMaxPercentage
	} else if !param.IsOmitted(u.OfRolloverAmountMaxAmount) {
		return u.OfRolloverAmountMaxAmount
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion) GetType() *string {
	if vt := u.OfRolloverAmountMaxPercentage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRolloverAmountMaxAmount; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion) GetValue() *float64 {
	if vt := u.OfRolloverAmountMaxPercentage; vt != nil {
		return (*float64)(&vt.Value)
	} else if vt := u.OfRolloverAmountMaxAmount; vt != nil {
		return (*float64)(&vt.Value)
	}
	return nil
}

type V1CreditGrantListParams struct {
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// Only return credit grants that are effective before this timestamp (exclusive).
	EffectiveBefore param.Opt[time.Time] `json:"effective_before,omitzero" format:"date-time"`
	// Only return credit grants that expire at or after this timestamp.
	NotExpiringBefore param.Opt[time.Time] `json:"not_expiring_before,omitzero" format:"date-time"`
	// An array of credit grant IDs. If this is specified, neither credit_type_ids nor
	// customer_ids may be specified.
	CreditGrantIDs []string `json:"credit_grant_ids,omitzero" format:"uuid"`
	// An array of credit type IDs. This must not be specified if credit_grant_ids is
	// specified.
	CreditTypeIDs []string `json:"credit_type_ids,omitzero" format:"uuid"`
	// An array of Metronome customer IDs. This must not be specified if
	// credit_grant_ids is specified.
	CustomerIDs []string `json:"customer_ids,omitzero" format:"uuid"`
	paramObj
}

func (r V1CreditGrantListParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1CreditGrantListParams]'s query parameters as
// `url.Values`.
func (r V1CreditGrantListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CreditGrantEditParams struct {
	// the ID of the credit grant
	ID string `json:"id,required" format:"uuid"`
	// the updated credit grant type
	CreditGrantType param.Opt[string] `json:"credit_grant_type,omitzero"`
	// the updated expiration date for the credit grant
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// the updated name for the credit grant
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r V1CreditGrantEditParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantEditParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantEditParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditGrantListEntriesParams struct {
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// If supplied, ledger entries will only be returned with an effective_at before
	// this time. This timestamp must not be in the future. If no timestamp is
	// supplied, all entries up to the start of the customer's next billing period will
	// be returned.
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// If supplied, only ledger entries effective at or after this time will be
	// returned.
	StartingOn param.Opt[time.Time] `json:"starting_on,omitzero" format:"date-time"`
	// Ledgers sort order by date, asc or desc. Defaults to asc.
	//
	// Any of "asc", "desc".
	Sort V1CreditGrantListEntriesParamsSort `query:"sort,omitzero" json:"-"`
	// A list of Metronome credit type IDs to fetch ledger entries for. If absent,
	// ledger entries for all credit types will be returned.
	CreditTypeIDs []string `json:"credit_type_ids,omitzero" format:"uuid"`
	// A list of Metronome customer IDs to fetch ledger entries for. If absent, ledger
	// entries for all customers will be returned.
	CustomerIDs []string `json:"customer_ids,omitzero" format:"uuid"`
	paramObj
}

func (r V1CreditGrantListEntriesParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantListEntriesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantListEntriesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1CreditGrantListEntriesParams]'s query parameters as
// `url.Values`.
func (r V1CreditGrantListEntriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Ledgers sort order by date, asc or desc. Defaults to asc.
type V1CreditGrantListEntriesParamsSort string

const (
	V1CreditGrantListEntriesParamsSortAsc  V1CreditGrantListEntriesParamsSort = "asc"
	V1CreditGrantListEntriesParamsSortDesc V1CreditGrantListEntriesParamsSort = "desc"
)

type V1CreditGrantVoidParams struct {
	ID string `json:"id,required" format:"uuid"`
	// If true, resets the uniqueness key on this grant so it can be re-used
	ReleaseUniquenessKey param.Opt[bool] `json:"release_uniqueness_key,omitzero"`
	// If true, void the purchase invoice associated with the grant
	VoidCreditPurchaseInvoice param.Opt[bool] `json:"void_credit_purchase_invoice,omitzero"`
	paramObj
}

func (r V1CreditGrantVoidParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantVoidParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantVoidParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
