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

// CustomerCreditService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerCreditService] method instead.
type CustomerCreditService struct {
	Options []option.RequestOption
}

// NewCustomerCreditService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerCreditService(opts ...option.RequestOption) (r *CustomerCreditService) {
	r = &CustomerCreditService{}
	r.Options = opts
	return
}

// Create a new credit at the customer level.
func (r *CustomerCreditService) New(ctx context.Context, body CustomerCreditNewParams, opts ...option.RequestOption) (res *CustomerCreditNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/customerCredits/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List credits.
func (r *CustomerCreditService) List(ctx context.Context, body CustomerCreditListParams, opts ...option.RequestOption) (res *CustomerCreditListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/customerCredits/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the end date of a credit
func (r *CustomerCreditService) UpdateEndDate(ctx context.Context, body CustomerCreditUpdateEndDateParams, opts ...option.RequestOption) (res *CustomerCreditUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/customerCredits/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CustomerCreditNewResponse struct {
	Data shared.ID                     `json:"data,required"`
	JSON customerCreditNewResponseJSON `json:"-"`
}

// customerCreditNewResponseJSON contains the JSON metadata for the struct
// [CustomerCreditNewResponse]
type customerCreditNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditNewResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListResponse struct {
	Data     []shared.Credit                `json:"data,required"`
	NextPage string                         `json:"next_page,required,nullable"`
	JSON     customerCreditListResponseJSON `json:"-"`
}

// customerCreditListResponseJSON contains the JSON metadata for the struct
// [CustomerCreditListResponse]
type customerCreditListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditUpdateEndDateResponse struct {
	Data shared.ID                               `json:"data,required"`
	JSON customerCreditUpdateEndDateResponseJSON `json:"-"`
}

// customerCreditUpdateEndDateResponseJSON contains the JSON metadata for the
// struct [CustomerCreditUpdateEndDateResponse]
type customerCreditUpdateEndDateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditUpdateEndDateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditUpdateEndDateResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditNewParams struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule param.Field[CustomerCreditNewParamsAccessSchedule] `json:"access_schedule,required"`
	CustomerID     param.Field[string]                                `json:"customer_id,required" format:"uuid"`
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
	ApplicableProductTags param.Field[[]string]          `json:"applicable_product_tags"`
	CustomFields          param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string] `json:"salesforce_opportunity_id"`
}

func (r CustomerCreditNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the credit to the customer.
type CustomerCreditNewParamsAccessSchedule struct {
	ScheduleItems param.Field[[]CustomerCreditNewParamsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	CreditTypeID  param.Field[string]                                              `json:"credit_type_id" format:"uuid"`
}

func (r CustomerCreditNewParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCreditNewParamsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r CustomerCreditNewParamsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCreditListParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Return only credits that have access schedules that "cover" the provided date
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	CreditID     param.Field[string]    `json:"credit_id" format:"uuid"`
	// Include only credits that have any access before the provided date (exclusive)
	EffectiveBefore param.Field[time.Time] `json:"effective_before" format:"date-time"`
	// Include credits from archived contracts.
	IncludeArchived param.Field[bool] `json:"include_archived"`
	// Include credits on the contract level.
	IncludeContractCredits param.Field[bool] `json:"include_contract_credits"`
	// Include credit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
	// The next page token from a previous response.
	NextPage param.Field[string] `json:"next_page"`
	// Include only credits that have any access on or after the provided date
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r CustomerCreditListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCreditUpdateEndDateParams struct {
	// RFC 3339 timestamp indicating when access to the credit will end and it will no
	// longer be possible to draw it down (exclusive).
	AccessEndingBefore param.Field[time.Time] `json:"access_ending_before,required" format:"date-time"`
	// ID of the commit to update
	CreditID param.Field[string] `json:"credit_id,required" format:"uuid"`
	// ID of the customer whose credit is to be updated
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r CustomerCreditUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
