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

// CustomerCommitService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerCommitService] method instead.
type CustomerCommitService struct {
	Options []option.RequestOption
}

// NewCustomerCommitService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerCommitService(opts ...option.RequestOption) (r *CustomerCommitService) {
	r = &CustomerCommitService{}
	r.Options = opts
	return
}

// Create a new commit at the customer level.
func (r *CustomerCommitService) New(ctx context.Context, body CustomerCommitNewParams, opts ...option.RequestOption) (res *CustomerCommitNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/customerCommits/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List commits.
func (r *CustomerCommitService) List(ctx context.Context, body CustomerCommitListParams, opts ...option.RequestOption) (res *CustomerCommitListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/customerCommits/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the end date of a PREPAID commit
func (r *CustomerCommitService) UpdateEndDate(ctx context.Context, body CustomerCommitUpdateEndDateParams, opts ...option.RequestOption) (res *CustomerCommitUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/customerCommits/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CustomerCommitNewResponse struct {
	Data shared.ID                     `json:"data,required"`
	JSON customerCommitNewResponseJSON `json:"-"`
}

// customerCommitNewResponseJSON contains the JSON metadata for the struct
// [CustomerCommitNewResponse]
type customerCommitNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCommitNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCommitNewResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCommitListResponse struct {
	Data     []shared.Commit                `json:"data,required"`
	NextPage string                         `json:"next_page,required,nullable"`
	JSON     customerCommitListResponseJSON `json:"-"`
}

// customerCommitListResponseJSON contains the JSON metadata for the struct
// [CustomerCommitListResponse]
type customerCommitListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCommitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCommitListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCommitUpdateEndDateResponse struct {
	Data shared.ID                               `json:"data,required"`
	JSON customerCommitUpdateEndDateResponseJSON `json:"-"`
}

// customerCommitUpdateEndDateResponseJSON contains the JSON metadata for the
// struct [CustomerCommitUpdateEndDateResponse]
type customerCommitUpdateEndDateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCommitUpdateEndDateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCommitUpdateEndDateResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCommitNewParams struct {
	// Schedule for distributing the commit to the customer. For "POSTPAID" commits
	// only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule param.Field[CustomerCommitNewParamsAccessSchedule] `json:"access_schedule,required"`
	CustomerID     param.Field[string]                                `json:"customer_id,required" format:"uuid"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority param.Field[float64] `json:"priority,required"`
	// ID of the fixed product associated with the commit. This is required because
	// products are used to invoice the commit amount.
	ProductID param.Field[string]                      `json:"product_id,required" format:"uuid"`
	Type      param.Field[CustomerCommitNewParamsType] `json:"type,required"`
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
	InvoiceSchedule param.Field[CustomerCommitNewParamsInvoiceSchedule] `json:"invoice_schedule"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string]                          `json:"netsuite_sales_order_id"`
	RateType             param.Field[CustomerCommitNewParamsRateType] `json:"rate_type"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string] `json:"salesforce_opportunity_id"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey param.Field[string] `json:"uniqueness_key"`
}

func (r CustomerCommitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the commit to the customer. For "POSTPAID" commits
// only one schedule item is allowed and amount must match invoice_schedule total.
type CustomerCommitNewParamsAccessSchedule struct {
	ScheduleItems param.Field[[]CustomerCommitNewParamsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r CustomerCommitNewParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCommitNewParamsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r CustomerCommitNewParamsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCommitNewParamsType string

const (
	CustomerCommitNewParamsTypePrepaid  CustomerCommitNewParamsType = "PREPAID"
	CustomerCommitNewParamsTypePostpaid CustomerCommitNewParamsType = "POSTPAID"
)

func (r CustomerCommitNewParamsType) IsKnown() bool {
	switch r {
	case CustomerCommitNewParamsTypePrepaid, CustomerCommitNewParamsTypePostpaid:
		return true
	}
	return false
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match
// accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this
// will be a "complimentary" commit with no invoice.
type CustomerCommitNewParamsInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[CustomerCommitNewParamsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]CustomerCommitNewParamsInvoiceScheduleScheduleItem] `json:"schedule_items"`
}

func (r CustomerCommitNewParamsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type CustomerCommitNewParamsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                        `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r CustomerCommitNewParamsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
	CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDivided        CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionEach           CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDivided, CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency string

const (
	CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyMonthly    CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyQuarterly  CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencySemiAnnual CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyAnnual     CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyMonthly, CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyQuarterly, CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type CustomerCommitNewParamsInvoiceScheduleScheduleItem struct {
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

func (r CustomerCommitNewParamsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCommitNewParamsRateType string

const (
	CustomerCommitNewParamsRateTypeCommitRate CustomerCommitNewParamsRateType = "COMMIT_RATE"
	CustomerCommitNewParamsRateTypeListRate   CustomerCommitNewParamsRateType = "LIST_RATE"
)

func (r CustomerCommitNewParamsRateType) IsKnown() bool {
	switch r {
	case CustomerCommitNewParamsRateTypeCommitRate, CustomerCommitNewParamsRateTypeListRate:
		return true
	}
	return false
}

type CustomerCommitListParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	CommitID   param.Field[string] `json:"commit_id" format:"uuid"`
	// Include only commits that have access schedules that "cover" the provided date
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	// Include only commits that have any access before the provided date (exclusive)
	EffectiveBefore param.Field[time.Time] `json:"effective_before" format:"date-time"`
	// Include commits from archived contracts.
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

func (r CustomerCommitListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCommitUpdateEndDateParams struct {
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

func (r CustomerCommitUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
