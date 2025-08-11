// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
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
func NewV1CreditGrantService(opts ...option.RequestOption) (r *V1CreditGrantService) {
	r = &V1CreditGrantService{}
	r.Options = opts
	return
}

// Create a new credit grant
func (r *V1CreditGrantService) New(ctx context.Context, body V1CreditGrantNewParams, opts ...option.RequestOption) (res *V1CreditGrantNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/credits/createGrant"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List credit grants. This list does not included voided grants.
func (r *V1CreditGrantService) List(ctx context.Context, params V1CreditGrantListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CreditGrantListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	path := "v1/credits/editGrant"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a list of credit ledger entries. Returns lists of ledgers per customer.
// Ledger entries are returned in chronological order. Ledger entries associated
// with voided credit grants are not included.
func (r *V1CreditGrantService) ListEntries(ctx context.Context, params V1CreditGrantListEntriesParams, opts ...option.RequestOption) (res *V1CreditGrantListEntriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/credits/listEntries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Void a credit grant
func (r *V1CreditGrantService) Void(ctx context.Context, body V1CreditGrantVoidParams, opts ...option.RequestOption) (res *V1CreditGrantVoidResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/credits/voidGrant"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1CreditGrantNewResponse struct {
	Data V1CreditGrantNewResponseData `json:"data,required"`
	JSON v1CreditGrantNewResponseJSON `json:"-"`
}

// v1CreditGrantNewResponseJSON contains the JSON metadata for the struct
// [V1CreditGrantNewResponse]
type v1CreditGrantNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantNewResponseData struct {
	ID   string                           `json:"id,required" format:"uuid"`
	JSON v1CreditGrantNewResponseDataJSON `json:"-"`
}

// v1CreditGrantNewResponseDataJSON contains the JSON metadata for the struct
// [V1CreditGrantNewResponseData]
type v1CreditGrantNewResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantListResponse struct {
	// the Metronome ID of the credit grant
	ID string `json:"id,required" format:"uuid"`
	// The effective balance of the grant as of the end of the customer's current
	// billing period. Expiration deductions will be included only if the grant expires
	// before the end of the current billing period.
	Balance      V1CreditGrantListResponseBalance `json:"balance,required"`
	CustomFields map[string]string                `json:"custom_fields,required"`
	// the Metronome ID of the customer
	CustomerID  string                               `json:"customer_id,required" format:"uuid"`
	Deductions  []V1CreditGrantListResponseDeduction `json:"deductions,required"`
	EffectiveAt time.Time                            `json:"effective_at,required" format:"date-time"`
	ExpiresAt   time.Time                            `json:"expires_at,required" format:"date-time"`
	// the amount of credits initially granted
	GrantAmount V1CreditGrantListResponseGrantAmount `json:"grant_amount,required"`
	Name        string                               `json:"name,required"`
	// the amount paid for this credit grant
	PaidAmount        V1CreditGrantListResponsePaidAmount         `json:"paid_amount,required"`
	PendingDeductions []V1CreditGrantListResponsePendingDeduction `json:"pending_deductions,required"`
	Priority          float64                                     `json:"priority,required"`
	CreditGrantType   string                                      `json:"credit_grant_type,nullable"`
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
	UniquenessKey string                        `json:"uniqueness_key,nullable"`
	JSON          v1CreditGrantListResponseJSON `json:"-"`
}

// v1CreditGrantListResponseJSON contains the JSON metadata for the struct
// [V1CreditGrantListResponse]
type v1CreditGrantListResponseJSON struct {
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

func (r *V1CreditGrantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListResponseJSON) RawJSON() string {
	return r.raw
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
	IncludingPending float64                              `json:"including_pending,required"`
	JSON             v1CreditGrantListResponseBalanceJSON `json:"-"`
}

// v1CreditGrantListResponseBalanceJSON contains the JSON metadata for the struct
// [V1CreditGrantListResponseBalance]
type v1CreditGrantListResponseBalanceJSON struct {
	EffectiveAt      apijson.Field
	ExcludingPending apijson.Field
	IncludingPending apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1CreditGrantListResponseBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListResponseBalanceJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantListResponseDeduction struct {
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
	InvoiceID string                                 `json:"invoice_id,nullable" format:"uuid"`
	JSON      v1CreditGrantListResponseDeductionJSON `json:"-"`
}

// v1CreditGrantListResponseDeductionJSON contains the JSON metadata for the struct
// [V1CreditGrantListResponseDeduction]
type v1CreditGrantListResponseDeductionJSON struct {
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

func (r *V1CreditGrantListResponseDeduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListResponseDeductionJSON) RawJSON() string {
	return r.raw
}

// the amount of credits initially granted
type V1CreditGrantListResponseGrantAmount struct {
	Amount float64 `json:"amount,required"`
	// the credit type for the amount granted
	CreditType V1CreditGrantListResponseGrantAmountCreditType `json:"credit_type,required"`
	JSON       v1CreditGrantListResponseGrantAmountJSON       `json:"-"`
}

// v1CreditGrantListResponseGrantAmountJSON contains the JSON metadata for the
// struct [V1CreditGrantListResponseGrantAmount]
type v1CreditGrantListResponseGrantAmountJSON struct {
	Amount      apijson.Field
	CreditType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantListResponseGrantAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListResponseGrantAmountJSON) RawJSON() string {
	return r.raw
}

// the credit type for the amount granted
type V1CreditGrantListResponseGrantAmountCreditType struct {
	ID   string                                             `json:"id,required" format:"uuid"`
	Name string                                             `json:"name,required"`
	JSON v1CreditGrantListResponseGrantAmountCreditTypeJSON `json:"-"`
}

// v1CreditGrantListResponseGrantAmountCreditTypeJSON contains the JSON metadata
// for the struct [V1CreditGrantListResponseGrantAmountCreditType]
type v1CreditGrantListResponseGrantAmountCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantListResponseGrantAmountCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListResponseGrantAmountCreditTypeJSON) RawJSON() string {
	return r.raw
}

// the amount paid for this credit grant
type V1CreditGrantListResponsePaidAmount struct {
	Amount float64 `json:"amount,required"`
	// the credit type for the amount paid
	CreditType V1CreditGrantListResponsePaidAmountCreditType `json:"credit_type,required"`
	JSON       v1CreditGrantListResponsePaidAmountJSON       `json:"-"`
}

// v1CreditGrantListResponsePaidAmountJSON contains the JSON metadata for the
// struct [V1CreditGrantListResponsePaidAmount]
type v1CreditGrantListResponsePaidAmountJSON struct {
	Amount      apijson.Field
	CreditType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantListResponsePaidAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListResponsePaidAmountJSON) RawJSON() string {
	return r.raw
}

// the credit type for the amount paid
type V1CreditGrantListResponsePaidAmountCreditType struct {
	ID   string                                            `json:"id,required" format:"uuid"`
	Name string                                            `json:"name,required"`
	JSON v1CreditGrantListResponsePaidAmountCreditTypeJSON `json:"-"`
}

// v1CreditGrantListResponsePaidAmountCreditTypeJSON contains the JSON metadata for
// the struct [V1CreditGrantListResponsePaidAmountCreditType]
type v1CreditGrantListResponsePaidAmountCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantListResponsePaidAmountCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListResponsePaidAmountCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantListResponsePendingDeduction struct {
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
	InvoiceID string                                        `json:"invoice_id,nullable" format:"uuid"`
	JSON      v1CreditGrantListResponsePendingDeductionJSON `json:"-"`
}

// v1CreditGrantListResponsePendingDeductionJSON contains the JSON metadata for the
// struct [V1CreditGrantListResponsePendingDeduction]
type v1CreditGrantListResponsePendingDeductionJSON struct {
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

func (r *V1CreditGrantListResponsePendingDeduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListResponsePendingDeductionJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantListResponseProduct struct {
	ID   string                               `json:"id,required"`
	Name string                               `json:"name,required"`
	JSON v1CreditGrantListResponseProductJSON `json:"-"`
}

// v1CreditGrantListResponseProductJSON contains the JSON metadata for the struct
// [V1CreditGrantListResponseProduct]
type v1CreditGrantListResponseProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantListResponseProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListResponseProductJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantEditResponse struct {
	Data V1CreditGrantEditResponseData `json:"data,required"`
	JSON v1CreditGrantEditResponseJSON `json:"-"`
}

// v1CreditGrantEditResponseJSON contains the JSON metadata for the struct
// [V1CreditGrantEditResponse]
type v1CreditGrantEditResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantEditResponseJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantEditResponseData struct {
	ID   string                            `json:"id,required" format:"uuid"`
	JSON v1CreditGrantEditResponseDataJSON `json:"-"`
}

// v1CreditGrantEditResponseDataJSON contains the JSON metadata for the struct
// [V1CreditGrantEditResponseData]
type v1CreditGrantEditResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantEditResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantEditResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantListEntriesResponse struct {
	Data     []V1CreditGrantListEntriesResponseData `json:"data,required"`
	NextPage string                                 `json:"next_page,required,nullable"`
	JSON     v1CreditGrantListEntriesResponseJSON   `json:"-"`
}

// v1CreditGrantListEntriesResponseJSON contains the JSON metadata for the struct
// [V1CreditGrantListEntriesResponse]
type v1CreditGrantListEntriesResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantListEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListEntriesResponseJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantListEntriesResponseData struct {
	CustomerID string                                       `json:"customer_id,required" format:"uuid"`
	Ledgers    []V1CreditGrantListEntriesResponseDataLedger `json:"ledgers,required"`
	JSON       v1CreditGrantListEntriesResponseDataJSON     `json:"-"`
}

// v1CreditGrantListEntriesResponseDataJSON contains the JSON metadata for the
// struct [V1CreditGrantListEntriesResponseData]
type v1CreditGrantListEntriesResponseDataJSON struct {
	CustomerID  apijson.Field
	Ledgers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantListEntriesResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListEntriesResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantListEntriesResponseDataLedger struct {
	CreditType V1CreditGrantListEntriesResponseDataLedgersCreditType `json:"credit_type,required"`
	// the effective balances at the end of the specified time window
	EndingBalance   V1CreditGrantListEntriesResponseDataLedgersEndingBalance   `json:"ending_balance,required"`
	Entries         []V1CreditGrantListEntriesResponseDataLedgersEntry         `json:"entries,required"`
	PendingEntries  []V1CreditGrantListEntriesResponseDataLedgersPendingEntry  `json:"pending_entries,required"`
	StartingBalance V1CreditGrantListEntriesResponseDataLedgersStartingBalance `json:"starting_balance,required"`
	JSON            v1CreditGrantListEntriesResponseDataLedgerJSON             `json:"-"`
}

// v1CreditGrantListEntriesResponseDataLedgerJSON contains the JSON metadata for
// the struct [V1CreditGrantListEntriesResponseDataLedger]
type v1CreditGrantListEntriesResponseDataLedgerJSON struct {
	CreditType      apijson.Field
	EndingBalance   apijson.Field
	Entries         apijson.Field
	PendingEntries  apijson.Field
	StartingBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V1CreditGrantListEntriesResponseDataLedger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListEntriesResponseDataLedgerJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantListEntriesResponseDataLedgersCreditType struct {
	ID   string                                                    `json:"id,required" format:"uuid"`
	Name string                                                    `json:"name,required"`
	JSON v1CreditGrantListEntriesResponseDataLedgersCreditTypeJSON `json:"-"`
}

// v1CreditGrantListEntriesResponseDataLedgersCreditTypeJSON contains the JSON
// metadata for the struct [V1CreditGrantListEntriesResponseDataLedgersCreditType]
type v1CreditGrantListEntriesResponseDataLedgersCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantListEntriesResponseDataLedgersCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListEntriesResponseDataLedgersCreditTypeJSON) RawJSON() string {
	return r.raw
}

// the effective balances at the end of the specified time window
type V1CreditGrantListEntriesResponseDataLedgersEndingBalance struct {
	// the ending_before request parameter (if supplied) or the current billing
	// period's end date
	EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
	// the ending balance, including the balance of all grants that have not expired
	// before the effective_at date and deductions that happened before the
	// effective_at date
	ExcludingPending float64 `json:"excluding_pending,required"`
	// the excluding_pending balance plus any pending invoice deductions and
	// expirations that will happen by the effective_at date
	IncludingPending float64                                                      `json:"including_pending,required"`
	JSON             v1CreditGrantListEntriesResponseDataLedgersEndingBalanceJSON `json:"-"`
}

// v1CreditGrantListEntriesResponseDataLedgersEndingBalanceJSON contains the JSON
// metadata for the struct
// [V1CreditGrantListEntriesResponseDataLedgersEndingBalance]
type v1CreditGrantListEntriesResponseDataLedgersEndingBalanceJSON struct {
	EffectiveAt      apijson.Field
	ExcludingPending apijson.Field
	IncludingPending apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1CreditGrantListEntriesResponseDataLedgersEndingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListEntriesResponseDataLedgersEndingBalanceJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantListEntriesResponseDataLedgersEntry struct {
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
	InvoiceID string                                               `json:"invoice_id,nullable" format:"uuid"`
	JSON      v1CreditGrantListEntriesResponseDataLedgersEntryJSON `json:"-"`
}

// v1CreditGrantListEntriesResponseDataLedgersEntryJSON contains the JSON metadata
// for the struct [V1CreditGrantListEntriesResponseDataLedgersEntry]
type v1CreditGrantListEntriesResponseDataLedgersEntryJSON struct {
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

func (r *V1CreditGrantListEntriesResponseDataLedgersEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListEntriesResponseDataLedgersEntryJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantListEntriesResponseDataLedgersPendingEntry struct {
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
	InvoiceID string                                                      `json:"invoice_id,nullable" format:"uuid"`
	JSON      v1CreditGrantListEntriesResponseDataLedgersPendingEntryJSON `json:"-"`
}

// v1CreditGrantListEntriesResponseDataLedgersPendingEntryJSON contains the JSON
// metadata for the struct
// [V1CreditGrantListEntriesResponseDataLedgersPendingEntry]
type v1CreditGrantListEntriesResponseDataLedgersPendingEntryJSON struct {
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

func (r *V1CreditGrantListEntriesResponseDataLedgersPendingEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListEntriesResponseDataLedgersPendingEntryJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantListEntriesResponseDataLedgersStartingBalance struct {
	// the starting_on request parameter (if supplied) or the first credit grant's
	// effective_at date
	EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
	// the starting balance, including all posted grants, deductions, and expirations
	// that happened at or before the effective_at timestamp
	ExcludingPending float64 `json:"excluding_pending,required"`
	// the excluding_pending balance plus any pending activity that has not been posted
	// at the time of the query
	IncludingPending float64                                                        `json:"including_pending,required"`
	JSON             v1CreditGrantListEntriesResponseDataLedgersStartingBalanceJSON `json:"-"`
}

// v1CreditGrantListEntriesResponseDataLedgersStartingBalanceJSON contains the JSON
// metadata for the struct
// [V1CreditGrantListEntriesResponseDataLedgersStartingBalance]
type v1CreditGrantListEntriesResponseDataLedgersStartingBalanceJSON struct {
	EffectiveAt      apijson.Field
	ExcludingPending apijson.Field
	IncludingPending apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1CreditGrantListEntriesResponseDataLedgersStartingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantListEntriesResponseDataLedgersStartingBalanceJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantVoidResponse struct {
	Data V1CreditGrantVoidResponseData `json:"data,required"`
	JSON v1CreditGrantVoidResponseJSON `json:"-"`
}

// v1CreditGrantVoidResponseJSON contains the JSON metadata for the struct
// [V1CreditGrantVoidResponse]
type v1CreditGrantVoidResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantVoidResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantVoidResponseJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantVoidResponseData struct {
	ID   string                            `json:"id,required" format:"uuid"`
	JSON v1CreditGrantVoidResponseDataJSON `json:"-"`
}

// v1CreditGrantVoidResponseDataJSON contains the JSON metadata for the struct
// [V1CreditGrantVoidResponseData]
type v1CreditGrantVoidResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CreditGrantVoidResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CreditGrantVoidResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CreditGrantNewParams struct {
	// the Metronome ID of the customer
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// The credit grant will only apply to usage or charges dated before this timestamp
	ExpiresAt param.Field[time.Time] `json:"expires_at,required" format:"date-time"`
	// the amount of credits granted
	GrantAmount param.Field[V1CreditGrantNewParamsGrantAmount] `json:"grant_amount,required"`
	// the name of the credit grant as it will appear on invoices
	Name param.Field[string] `json:"name,required"`
	// the amount paid for this credit grant
	PaidAmount      param.Field[V1CreditGrantNewParamsPaidAmount] `json:"paid_amount,required"`
	Priority        param.Field[float64]                          `json:"priority,required"`
	CreditGrantType param.Field[string]                           `json:"credit_grant_type"`
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
	RolloverSettings param.Field[V1CreditGrantNewParamsRolloverSettings] `json:"rollover_settings"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey param.Field[string] `json:"uniqueness_key"`
}

func (r V1CreditGrantNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// the amount of credits granted
type V1CreditGrantNewParamsGrantAmount struct {
	Amount param.Field[float64] `json:"amount,required"`
	// the ID of the pricing unit to be used. Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id,required" format:"uuid"`
}

func (r V1CreditGrantNewParamsGrantAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// the amount paid for this credit grant
type V1CreditGrantNewParamsPaidAmount struct {
	Amount param.Field[float64] `json:"amount,required"`
	// the ID of the pricing unit to be used. Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id,required" format:"uuid"`
}

func (r V1CreditGrantNewParamsPaidAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a rollover for this credit grant so if it expires it rolls over a
// configured amount to a new credit grant. This feature is currently opt-in only.
// Contact Metronome to be added to the beta.
type V1CreditGrantNewParamsRolloverSettings struct {
	// The date to expire the rollover credits.
	ExpiresAt param.Field[time.Time] `json:"expires_at,required" format:"date-time"`
	// The priority to give the rollover credit grant that gets created when a rollover
	// happens.
	Priority param.Field[float64] `json:"priority,required"`
	// Specify how much to rollover to the rollover credit grant
	RolloverAmount param.Field[V1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion] `json:"rollover_amount,required"`
}

func (r V1CreditGrantNewParamsRolloverSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how much to rollover to the rollover credit grant
type V1CreditGrantNewParamsRolloverSettingsRolloverAmount struct {
	// Rollover up to a percentage of the original credit grant amount.
	Type param.Field[V1CreditGrantNewParamsRolloverSettingsRolloverAmountType] `json:"type,required"`
	// The maximum percentage (0-1) of the original credit grant to rollover.
	Value param.Field[float64] `json:"value,required"`
}

func (r V1CreditGrantNewParamsRolloverSettingsRolloverAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1CreditGrantNewParamsRolloverSettingsRolloverAmount) implementsV1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion() {
}

// Specify how much to rollover to the rollover credit grant
//
// Satisfied by [V1CreditGrantNewParamsRolloverSettingsRolloverAmountObject],
// [V1CreditGrantNewParamsRolloverSettingsRolloverAmountObject],
// [V1CreditGrantNewParamsRolloverSettingsRolloverAmount].
type V1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion interface {
	implementsV1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion()
}

type V1CreditGrantNewParamsRolloverSettingsRolloverAmountObject struct {
	// Rollover up to a percentage of the original credit grant amount.
	Type param.Field[V1CreditGrantNewParamsRolloverSettingsRolloverAmountObjectType] `json:"type,required"`
	// The maximum percentage (0-1) of the original credit grant to rollover.
	Value param.Field[float64] `json:"value,required"`
}

func (r V1CreditGrantNewParamsRolloverSettingsRolloverAmountObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1CreditGrantNewParamsRolloverSettingsRolloverAmountObject) implementsV1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion() {
}

// Rollover up to a percentage of the original credit grant amount.
type V1CreditGrantNewParamsRolloverSettingsRolloverAmountObjectType string

const (
	V1CreditGrantNewParamsRolloverSettingsRolloverAmountObjectTypeMaxPercentage V1CreditGrantNewParamsRolloverSettingsRolloverAmountObjectType = "MAX_PERCENTAGE"
)

func (r V1CreditGrantNewParamsRolloverSettingsRolloverAmountObjectType) IsKnown() bool {
	switch r {
	case V1CreditGrantNewParamsRolloverSettingsRolloverAmountObjectTypeMaxPercentage:
		return true
	}
	return false
}

// Rollover up to a percentage of the original credit grant amount.
type V1CreditGrantNewParamsRolloverSettingsRolloverAmountType string

const (
	V1CreditGrantNewParamsRolloverSettingsRolloverAmountTypeMaxPercentage V1CreditGrantNewParamsRolloverSettingsRolloverAmountType = "MAX_PERCENTAGE"
	V1CreditGrantNewParamsRolloverSettingsRolloverAmountTypeMaxAmount     V1CreditGrantNewParamsRolloverSettingsRolloverAmountType = "MAX_AMOUNT"
)

func (r V1CreditGrantNewParamsRolloverSettingsRolloverAmountType) IsKnown() bool {
	switch r {
	case V1CreditGrantNewParamsRolloverSettingsRolloverAmountTypeMaxPercentage, V1CreditGrantNewParamsRolloverSettingsRolloverAmountTypeMaxAmount:
		return true
	}
	return false
}

type V1CreditGrantListParams struct {
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

func (r V1CreditGrantListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V1CreditGrantListParams]'s query parameters as
// `url.Values`.
func (r V1CreditGrantListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CreditGrantEditParams struct {
	// the ID of the credit grant
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// the updated credit grant type
	CreditGrantType param.Field[string] `json:"credit_grant_type"`
	// the updated expiration date for the credit grant
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// the updated name for the credit grant
	Name param.Field[string] `json:"name"`
}

func (r V1CreditGrantEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CreditGrantListEntriesParams struct {
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Ledgers sort order by date, asc or desc. Defaults to asc.
	Sort param.Field[V1CreditGrantListEntriesParamsSort] `query:"sort"`
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

func (r V1CreditGrantListEntriesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V1CreditGrantListEntriesParams]'s query parameters as
// `url.Values`.
func (r V1CreditGrantListEntriesParams) URLQuery() (v url.Values) {
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

func (r V1CreditGrantListEntriesParamsSort) IsKnown() bool {
	switch r {
	case V1CreditGrantListEntriesParamsSortAsc, V1CreditGrantListEntriesParamsSortDesc:
		return true
	}
	return false
}

type V1CreditGrantVoidParams struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// If true, resets the uniqueness key on this grant so it can be re-used
	ReleaseUniquenessKey param.Field[bool] `json:"release_uniqueness_key"`
	// If true, void the purchase invoice associated with the grant
	VoidCreditPurchaseInvoice param.Field[bool] `json:"void_credit_purchase_invoice"`
}

func (r V1CreditGrantVoidParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
