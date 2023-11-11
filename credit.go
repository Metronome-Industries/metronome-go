// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "github.com/metronome/metronome-go/internal/shared"
  "github.com/metronome/metronome-go/internal/apijson"
  "time"
  "github.com/metronome/metronome-go/internal/param"
  "net/url"
  "github.com/metronome/metronome-go/internal/apiquery"
  "context"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/requestconfig"
)

// CreditService contains methods and other services that help with interacting
// with the metronome API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewCreditService] method instead.
type CreditService struct {
Options []option.RequestOption
}

// NewCreditService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCreditService(opts ...option.RequestOption) (r *CreditService) {
  r = &CreditService{}
  r.Options = opts
  return
}

// Create a new credit grant
func (r *CreditService) NewGrant(ctx context.Context, body CreditNewGrantParams, opts ...option.RequestOption) (res *CreditNewGrantResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "credits/createGrant"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Edit an existing credit grant
func (r *CreditService) EditGrant(ctx context.Context, body CreditEditGrantParams, opts ...option.RequestOption) (res *CreditEditGrantResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "credits/editGrant"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Fetches a list of credit ledger entries. Returns lists of ledgers per customer.
// Ledger entries are returned in reverse chronological order. Ledger entries
// associated with voided credit grants are not included.
func (r *CreditService) ListEntries(ctx context.Context, params CreditListEntriesParams, opts ...option.RequestOption) (res *CreditListEntriesResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "credits/listEntries"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
  return
}

// List credit grants. This list does not included voided grants.
func (r *CreditService) ListGrants(ctx context.Context, params CreditListGrantsParams, opts ...option.RequestOption) (res *CreditListGrantsResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "credits/listGrants"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
  return
}

// Void a credit grant
func (r *CreditService) VoidGrant(ctx context.Context, body CreditVoidGrantParams, opts ...option.RequestOption) (res *CreditVoidGrantResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "credits/voidGrant"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

type CreditNewGrantResponse struct {
Data shared.ID `json:"data,required"`
JSON creditNewGrantResponseJSON
}

// creditNewGrantResponseJSON contains the JSON metadata for the struct
// [CreditNewGrantResponse]
type creditNewGrantResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditNewGrantResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditEditGrantResponse struct {
Data shared.ID `json:"data,required"`
JSON creditEditGrantResponseJSON
}

// creditEditGrantResponseJSON contains the JSON metadata for the struct
// [CreditEditGrantResponse]
type creditEditGrantResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditEditGrantResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListEntriesResponse struct {
Data []CreditListEntriesResponseData `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON creditListEntriesResponseJSON
}

// creditListEntriesResponseJSON contains the JSON metadata for the struct
// [CreditListEntriesResponse]
type creditListEntriesResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListEntriesResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListEntriesResponseData struct {
CustomerID string `json:"customer_id,required" format:"uuid"`
Ledgers []CreditListEntriesResponseDataLedger `json:"ledgers,required"`
JSON creditListEntriesResponseDataJSON
}

// creditListEntriesResponseDataJSON contains the JSON metadata for the struct
// [CreditListEntriesResponseData]
type creditListEntriesResponseDataJSON struct {
CustomerID apijson.Field
Ledgers apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListEntriesResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListEntriesResponseDataLedger struct {
CreditType CreditListEntriesResponseDataLedgersCreditType `json:"credit_type,required"`
// the effective balances at the end of the specified time window
EndingBalance CreditListEntriesResponseDataLedgersEndingBalance `json:"ending_balance,required"`
Entries []CreditListEntriesResponseDataLedgersEntry `json:"entries,required"`
PendingEntries []CreditListEntriesResponseDataLedgersPendingEntry `json:"pending_entries,required"`
StartingBalance CreditListEntriesResponseDataLedgersStartingBalance `json:"starting_balance,required"`
JSON creditListEntriesResponseDataLedgerJSON
}

// creditListEntriesResponseDataLedgerJSON contains the JSON metadata for the
// struct [CreditListEntriesResponseDataLedger]
type creditListEntriesResponseDataLedgerJSON struct {
CreditType apijson.Field
EndingBalance apijson.Field
Entries apijson.Field
PendingEntries apijson.Field
StartingBalance apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListEntriesResponseDataLedger) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListEntriesResponseDataLedgersCreditType struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON creditListEntriesResponseDataLedgersCreditTypeJSON
}

// creditListEntriesResponseDataLedgersCreditTypeJSON contains the JSON metadata
// for the struct [CreditListEntriesResponseDataLedgersCreditType]
type creditListEntriesResponseDataLedgersCreditTypeJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListEntriesResponseDataLedgersCreditType) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

// the effective balances at the end of the specified time window
type CreditListEntriesResponseDataLedgersEndingBalance struct {
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
JSON creditListEntriesResponseDataLedgersEndingBalanceJSON
}

// creditListEntriesResponseDataLedgersEndingBalanceJSON contains the JSON metadata
// for the struct [CreditListEntriesResponseDataLedgersEndingBalance]
type creditListEntriesResponseDataLedgersEndingBalanceJSON struct {
EffectiveAt apijson.Field
ExcludingPending apijson.Field
IncludingPending apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListEntriesResponseDataLedgersEndingBalance) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListEntriesResponseDataLedgersEntry struct {
// an amount representing the change to the customer's credit balance
Amount float64 `json:"amount,required"`
CreatedBy string `json:"created_by,required"`
// the credit grant this entry is related to
CreditGrantID string `json:"credit_grant_id,required" format:"uuid"`
EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
Reason string `json:"reason,required"`
// the running balance for this credit type at the time of the ledger entry,
// including all preceding charges
RunningBalance float64 `json:"running_balance,required"`
// if this entry is a deduction, the Metronome ID of the invoice where the credit
// deduction was consumed; if this entry is a grant, the Metronome ID of the
// invoice where the grant's paid_amount was charged
InvoiceID string `json:"invoice_id,nullable" format:"uuid"`
JSON creditListEntriesResponseDataLedgersEntryJSON
}

// creditListEntriesResponseDataLedgersEntryJSON contains the JSON metadata for the
// struct [CreditListEntriesResponseDataLedgersEntry]
type creditListEntriesResponseDataLedgersEntryJSON struct {
Amount apijson.Field
CreatedBy apijson.Field
CreditGrantID apijson.Field
EffectiveAt apijson.Field
Reason apijson.Field
RunningBalance apijson.Field
InvoiceID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListEntriesResponseDataLedgersEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListEntriesResponseDataLedgersPendingEntry struct {
// an amount representing the change to the customer's credit balance
Amount float64 `json:"amount,required"`
CreatedBy string `json:"created_by,required"`
// the credit grant this entry is related to
CreditGrantID string `json:"credit_grant_id,required" format:"uuid"`
EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
Reason string `json:"reason,required"`
// the running balance for this credit type at the time of the ledger entry,
// including all preceding charges
RunningBalance float64 `json:"running_balance,required"`
// if this entry is a deduction, the Metronome ID of the invoice where the credit
// deduction was consumed; if this entry is a grant, the Metronome ID of the
// invoice where the grant's paid_amount was charged
InvoiceID string `json:"invoice_id,nullable" format:"uuid"`
JSON creditListEntriesResponseDataLedgersPendingEntryJSON
}

// creditListEntriesResponseDataLedgersPendingEntryJSON contains the JSON metadata
// for the struct [CreditListEntriesResponseDataLedgersPendingEntry]
type creditListEntriesResponseDataLedgersPendingEntryJSON struct {
Amount apijson.Field
CreatedBy apijson.Field
CreditGrantID apijson.Field
EffectiveAt apijson.Field
Reason apijson.Field
RunningBalance apijson.Field
InvoiceID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListEntriesResponseDataLedgersPendingEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListEntriesResponseDataLedgersStartingBalance struct {
// the starting_on request parameter (if supplied) or the first credit grant's
// effective_at date
EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
// the starting balance, including all posted grants, deductions, and expirations
// that happened at or before the effective_at timestamp
ExcludingPending float64 `json:"excluding_pending,required"`
// the excluding_pending balance plus any pending activity that has not been posted
// at the time of the query
IncludingPending float64 `json:"including_pending,required"`
JSON creditListEntriesResponseDataLedgersStartingBalanceJSON
}

// creditListEntriesResponseDataLedgersStartingBalanceJSON contains the JSON
// metadata for the struct [CreditListEntriesResponseDataLedgersStartingBalance]
type creditListEntriesResponseDataLedgersStartingBalanceJSON struct {
EffectiveAt apijson.Field
ExcludingPending apijson.Field
IncludingPending apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListEntriesResponseDataLedgersStartingBalance) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListGrantsResponse struct {
Data []CreditListGrantsResponseData `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON creditListGrantsResponseJSON
}

// creditListGrantsResponseJSON contains the JSON metadata for the struct
// [CreditListGrantsResponse]
type creditListGrantsResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListGrantsResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListGrantsResponseData struct {
// the Metronome ID of the credit grant
ID string `json:"id,required" format:"uuid"`
// The effective balance of the grant as of the end of the customer's current
// billing period. Expiration deductions will be included only if the grant expires
// before the end of the current billing period.
Balance CreditListGrantsResponseDataBalance `json:"balance,required"`
CustomFields map[string]string `json:"custom_fields,required"`
// the Metronome ID of the customer
CustomerID string `json:"customer_id,required" format:"uuid"`
Deductions []CreditListGrantsResponseDataDeduction `json:"deductions,required"`
EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
// the amount of credits initially granted
GrantAmount CreditListGrantsResponseDataGrantAmount `json:"grant_amount,required"`
Name string `json:"name,required"`
// the amount paid for this credit grant
PaidAmount CreditListGrantsResponseDataPaidAmount `json:"paid_amount,required"`
PendingDeductions []CreditListGrantsResponseDataPendingDeduction `json:"pending_deductions,required"`
Priority float64 `json:"priority,required"`
// the Metronome ID of the invoice with the purchase charge for this credit grant,
// if applicable
InvoiceID string `json:"invoice_id,nullable" format:"uuid"`
// The products which these credits will be applied to. (If unspecified, the
// credits will be applied to charges for all products.)
Products []CreditListGrantsResponseDataProduct `json:"products"`
Reason string `json:"reason,nullable"`
JSON creditListGrantsResponseDataJSON
}

// creditListGrantsResponseDataJSON contains the JSON metadata for the struct
// [CreditListGrantsResponseData]
type creditListGrantsResponseDataJSON struct {
ID apijson.Field
Balance apijson.Field
CustomFields apijson.Field
CustomerID apijson.Field
Deductions apijson.Field
EffectiveAt apijson.Field
ExpiresAt apijson.Field
GrantAmount apijson.Field
Name apijson.Field
PaidAmount apijson.Field
PendingDeductions apijson.Field
Priority apijson.Field
InvoiceID apijson.Field
Products apijson.Field
Reason apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListGrantsResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

// The effective balance of the grant as of the end of the customer's current
// billing period. Expiration deductions will be included only if the grant expires
// before the end of the current billing period.
type CreditListGrantsResponseDataBalance struct {
// The end_date of the customer's current billing period.
EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
// The grant's current balance including all posted deductions. If the grant has
// expired, this amount will be 0.
ExcludingPending float64 `json:"excluding_pending,required"`
// The grant's current balance including all posted and pending deductions. If the
// grant expires before the end of the customer's current billing period, this
// amount will be 0.
IncludingPending float64 `json:"including_pending,required"`
JSON creditListGrantsResponseDataBalanceJSON
}

// creditListGrantsResponseDataBalanceJSON contains the JSON metadata for the
// struct [CreditListGrantsResponseDataBalance]
type creditListGrantsResponseDataBalanceJSON struct {
EffectiveAt apijson.Field
ExcludingPending apijson.Field
IncludingPending apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListGrantsResponseDataBalance) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListGrantsResponseDataDeduction struct {
// an amount representing the change to the customer's credit balance
Amount float64 `json:"amount,required"`
CreatedBy string `json:"created_by,required"`
// the credit grant this entry is related to
CreditGrantID string `json:"credit_grant_id,required" format:"uuid"`
EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
Reason string `json:"reason,required"`
// the running balance for this credit type at the time of the ledger entry,
// including all preceding charges
RunningBalance float64 `json:"running_balance,required"`
// if this entry is a deduction, the Metronome ID of the invoice where the credit
// deduction was consumed; if this entry is a grant, the Metronome ID of the
// invoice where the grant's paid_amount was charged
InvoiceID string `json:"invoice_id,nullable" format:"uuid"`
JSON creditListGrantsResponseDataDeductionJSON
}

// creditListGrantsResponseDataDeductionJSON contains the JSON metadata for the
// struct [CreditListGrantsResponseDataDeduction]
type creditListGrantsResponseDataDeductionJSON struct {
Amount apijson.Field
CreatedBy apijson.Field
CreditGrantID apijson.Field
EffectiveAt apijson.Field
Reason apijson.Field
RunningBalance apijson.Field
InvoiceID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListGrantsResponseDataDeduction) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

// the amount of credits initially granted
type CreditListGrantsResponseDataGrantAmount struct {
Amount float64 `json:"amount,required"`
// the credit type for the amount granted
CreditType CreditListGrantsResponseDataGrantAmountCreditType `json:"credit_type,required"`
JSON creditListGrantsResponseDataGrantAmountJSON
}

// creditListGrantsResponseDataGrantAmountJSON contains the JSON metadata for the
// struct [CreditListGrantsResponseDataGrantAmount]
type creditListGrantsResponseDataGrantAmountJSON struct {
Amount apijson.Field
CreditType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListGrantsResponseDataGrantAmount) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

// the credit type for the amount granted
type CreditListGrantsResponseDataGrantAmountCreditType struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON creditListGrantsResponseDataGrantAmountCreditTypeJSON
}

// creditListGrantsResponseDataGrantAmountCreditTypeJSON contains the JSON metadata
// for the struct [CreditListGrantsResponseDataGrantAmountCreditType]
type creditListGrantsResponseDataGrantAmountCreditTypeJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListGrantsResponseDataGrantAmountCreditType) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

// the amount paid for this credit grant
type CreditListGrantsResponseDataPaidAmount struct {
Amount float64 `json:"amount,required"`
// the credit type for the amount paid
CreditType CreditListGrantsResponseDataPaidAmountCreditType `json:"credit_type,required"`
JSON creditListGrantsResponseDataPaidAmountJSON
}

// creditListGrantsResponseDataPaidAmountJSON contains the JSON metadata for the
// struct [CreditListGrantsResponseDataPaidAmount]
type creditListGrantsResponseDataPaidAmountJSON struct {
Amount apijson.Field
CreditType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListGrantsResponseDataPaidAmount) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

// the credit type for the amount paid
type CreditListGrantsResponseDataPaidAmountCreditType struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON creditListGrantsResponseDataPaidAmountCreditTypeJSON
}

// creditListGrantsResponseDataPaidAmountCreditTypeJSON contains the JSON metadata
// for the struct [CreditListGrantsResponseDataPaidAmountCreditType]
type creditListGrantsResponseDataPaidAmountCreditTypeJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListGrantsResponseDataPaidAmountCreditType) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListGrantsResponseDataPendingDeduction struct {
// an amount representing the change to the customer's credit balance
Amount float64 `json:"amount,required"`
CreatedBy string `json:"created_by,required"`
// the credit grant this entry is related to
CreditGrantID string `json:"credit_grant_id,required" format:"uuid"`
EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
Reason string `json:"reason,required"`
// the running balance for this credit type at the time of the ledger entry,
// including all preceding charges
RunningBalance float64 `json:"running_balance,required"`
// if this entry is a deduction, the Metronome ID of the invoice where the credit
// deduction was consumed; if this entry is a grant, the Metronome ID of the
// invoice where the grant's paid_amount was charged
InvoiceID string `json:"invoice_id,nullable" format:"uuid"`
JSON creditListGrantsResponseDataPendingDeductionJSON
}

// creditListGrantsResponseDataPendingDeductionJSON contains the JSON metadata for
// the struct [CreditListGrantsResponseDataPendingDeduction]
type creditListGrantsResponseDataPendingDeductionJSON struct {
Amount apijson.Field
CreatedBy apijson.Field
CreditGrantID apijson.Field
EffectiveAt apijson.Field
Reason apijson.Field
RunningBalance apijson.Field
InvoiceID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListGrantsResponseDataPendingDeduction) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditListGrantsResponseDataProduct struct {
ID string `json:"id,required"`
Name string `json:"name,required"`
JSON creditListGrantsResponseDataProductJSON
}

// creditListGrantsResponseDataProductJSON contains the JSON metadata for the
// struct [CreditListGrantsResponseDataProduct]
type creditListGrantsResponseDataProductJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditListGrantsResponseDataProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditVoidGrantResponse struct {
Data shared.ID `json:"data,required"`
JSON creditVoidGrantResponseJSON
}

// creditVoidGrantResponseJSON contains the JSON metadata for the struct
// [CreditVoidGrantResponse]
type creditVoidGrantResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CreditVoidGrantResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CreditNewGrantParams struct {
// the Metronome ID of the customer
CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
// The credit grant will only apply to billing periods that end before this
// timestamp.
ExpiresAt param.Field[time.Time] `json:"expires_at,required" format:"date-time"`
// the amount of credits granted
GrantAmount param.Field[CreditNewGrantParamsGrantAmount] `json:"grant_amount,required"`
// the name of the credit grant as it will appear on invoices
Name param.Field[string] `json:"name,required"`
// the amount paid for this credit grant
PaidAmount param.Field[CreditNewGrantParamsPaidAmount] `json:"paid_amount,required"`
Priority param.Field[float64] `json:"priority,required"`
// The credit grant will only apply to billing periods that end at or after this
// timestamp.
EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
// The date to issue an invoice for the paid_amount.
InvoiceDate param.Field[time.Time] `json:"invoice_date" format:"date-time"`
// The product(s) which these credits will be applied to. (If unspecified, the
// credits will be applied to charges for all products.). The array ordering
// specified here will be used to determine the order in which credits will be
// applied to invoice line items
ProductIDs param.Field[[]string] `json:"product_ids" format:"uuid"`
Reason param.Field[string] `json:"reason"`
}

func (r CreditNewGrantParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// the amount of credits granted
type CreditNewGrantParamsGrantAmount struct {
Amount param.Field[float64] `json:"amount,required"`
CreditTypeID param.Field[string] `json:"credit_type_id,required" format:"uuid"`
}

func (r CreditNewGrantParamsGrantAmount) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// the amount paid for this credit grant
type CreditNewGrantParamsPaidAmount struct {
Amount param.Field[float64] `json:"amount,required"`
CreditTypeID param.Field[string] `json:"credit_type_id,required" format:"uuid"`
}

func (r CreditNewGrantParamsPaidAmount) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type CreditEditGrantParams struct {
// the ID of the credit grant
ID param.Field[string] `json:"id,required" format:"uuid"`
// the updated expiration date for the credit grant
ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
// the updated name for the credit grant
Name param.Field[string] `json:"name"`
}

func (r CreditEditGrantParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type CreditListEntriesParams struct {
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

func (r CreditListEntriesParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// URLQuery serializes [CreditListEntriesParams]'s query parameters as
// `url.Values`.
func (r CreditListEntriesParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

type CreditListGrantsParams struct {
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

func (r CreditListGrantsParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// URLQuery serializes [CreditListGrantsParams]'s query parameters as `url.Values`.
func (r CreditListGrantsParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

type CreditVoidGrantParams struct {
ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r CreditVoidGrantParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}