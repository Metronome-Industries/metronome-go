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
func NewV1CustomerCommitService(opts ...option.RequestOption) (r *V1CustomerCommitService) {
	r = &V1CustomerCommitService{}
	r.Options = opts
	return
}

// Create a new commit at the customer level.
func (r *V1CustomerCommitService) New(ctx context.Context, body V1CustomerCommitNewParams, opts ...option.RequestOption) (res *V1CustomerCommitNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCommits/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List commits.
func (r *V1CustomerCommitService) List(ctx context.Context, body V1CustomerCommitListParams, opts ...option.RequestOption) (res *V1CustomerCommitListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCommits/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Pull forward the end date of a prepaid commit. Use the "edit a commit" endpoint
// to extend the end date of a prepaid commit, or to make other edits to the
// commit.
func (r *V1CustomerCommitService) UpdateEndDate(ctx context.Context, body V1CustomerCommitUpdateEndDateParams, opts ...option.RequestOption) (res *V1CustomerCommitUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCommits/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1CustomerCommitNewResponse struct {
	Data shared.ID                       `json:"data,required"`
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

type V1CustomerCommitListResponse struct {
	Data     []shared.Commit                  `json:"data,required"`
	NextPage string                           `json:"next_page,required,nullable"`
	JSON     v1CustomerCommitListResponseJSON `json:"-"`
}

// v1CustomerCommitListResponseJSON contains the JSON metadata for the struct
// [V1CustomerCommitListResponse]
type v1CustomerCommitListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitUpdateEndDateResponse struct {
	Data shared.ID                                 `json:"data,required"`
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
	// Which products the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags param.Field[[]string]          `json:"applicable_product_tags"`
	CustomFields          param.Field[map[string]string] `json:"custom_fields"`
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
