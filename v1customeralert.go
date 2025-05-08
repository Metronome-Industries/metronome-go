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
	"github.com/Metronome-Industries/metronome-go/shared"
)

// V1CustomerAlertService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerAlertService] method instead.
type V1CustomerAlertService struct {
	Options []option.RequestOption
}

// NewV1CustomerAlertService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomerAlertService(opts ...option.RequestOption) (r *V1CustomerAlertService) {
	r = &V1CustomerAlertService{}
	r.Options = opts
	return
}

// Get the customer alert status and alert information for the specified customer
// and alert
func (r *V1CustomerAlertService) Get(ctx context.Context, body V1CustomerAlertGetParams, opts ...option.RequestOption) (res *V1CustomerAlertGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/customer-alerts/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch all customer alert statuses and alert information for a customer
func (r *V1CustomerAlertService) List(ctx context.Context, params V1CustomerAlertListParams, opts ...option.RequestOption) (res *V1CustomerAlertListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/customer-alerts/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Reset state for an alert by customer id and force re-evaluation
func (r *V1CustomerAlertService) Reset(ctx context.Context, body V1CustomerAlertResetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/customer-alerts/reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CustomerAlert struct {
	Alert CustomerAlertAlert `json:"alert,required"`
	// The status of the customer alert. If the alert is archived, null will be
	// returned.
	CustomerStatus CustomerAlertCustomerStatus `json:"customer_status,required,nullable"`
	// If present, indicates the reason the alert was triggered.
	TriggeredBy string            `json:"triggered_by,nullable"`
	JSON        customerAlertJSON `json:"-"`
}

// customerAlertJSON contains the JSON metadata for the struct [CustomerAlert]
type customerAlertJSON struct {
	Alert          apijson.Field
	CustomerStatus apijson.Field
	TriggeredBy    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerAlert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertJSON) RawJSON() string {
	return r.raw
}

type CustomerAlertAlert struct {
	// the Metronome ID of the alert
	ID string `json:"id,required"`
	// Name of the alert
	Name string `json:"name,required"`
	// Status of the alert
	Status CustomerAlertAlertStatus `json:"status,required"`
	// Threshold value of the alert policy
	Threshold float64 `json:"threshold,required"`
	// Type of the alert
	Type CustomerAlertAlertType `json:"type,required"`
	// Timestamp for when the alert was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An array of strings, representing a way to filter the credit grant this alert
	// applies to, by looking at the credit_grant_type field on the credit grant. This
	// field is only defined for CreditPercentage and CreditBalance alerts
	CreditGrantTypeFilters []string              `json:"credit_grant_type_filters"`
	CreditType             shared.CreditTypeData `json:"credit_type,nullable"`
	// A list of custom field filters for alert types that support advanced filtering
	CustomFieldFilters []CustomerAlertAlertCustomFieldFilter `json:"custom_field_filters"`
	// Scopes alert evaluation to a specific presentation group key on individual line
	// items. Only present for spend alerts.
	GroupKeyFilter CustomerAlertAlertGroupKeyFilter `json:"group_key_filter"`
	// Only supported for invoice_total_reached alerts. A list of invoice types to
	// evaluate.
	InvoiceTypesFilter []string `json:"invoice_types_filter"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                 `json:"uniqueness_key"`
	JSON          customerAlertAlertJSON `json:"-"`
}

// customerAlertAlertJSON contains the JSON metadata for the struct
// [CustomerAlertAlert]
type customerAlertAlertJSON struct {
	ID                     apijson.Field
	Name                   apijson.Field
	Status                 apijson.Field
	Threshold              apijson.Field
	Type                   apijson.Field
	UpdatedAt              apijson.Field
	CreditGrantTypeFilters apijson.Field
	CreditType             apijson.Field
	CustomFieldFilters     apijson.Field
	GroupKeyFilter         apijson.Field
	InvoiceTypesFilter     apijson.Field
	UniquenessKey          apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CustomerAlertAlert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertAlertJSON) RawJSON() string {
	return r.raw
}

// Status of the alert
type CustomerAlertAlertStatus string

const (
	CustomerAlertAlertStatusEnabled  CustomerAlertAlertStatus = "enabled"
	CustomerAlertAlertStatusArchived CustomerAlertAlertStatus = "archived"
	CustomerAlertAlertStatusDisabled CustomerAlertAlertStatus = "disabled"
)

func (r CustomerAlertAlertStatus) IsKnown() bool {
	switch r {
	case CustomerAlertAlertStatusEnabled, CustomerAlertAlertStatusArchived, CustomerAlertAlertStatusDisabled:
		return true
	}
	return false
}

// Type of the alert
type CustomerAlertAlertType string

const (
	CustomerAlertAlertTypeLowCreditBalanceReached                           CustomerAlertAlertType = "low_credit_balance_reached"
	CustomerAlertAlertTypeSpendThresholdReached                             CustomerAlertAlertType = "spend_threshold_reached"
	CustomerAlertAlertTypeMonthlyInvoiceTotalSpendThresholdReached          CustomerAlertAlertType = "monthly_invoice_total_spend_threshold_reached"
	CustomerAlertAlertTypeLowRemainingDaysInPlanReached                     CustomerAlertAlertType = "low_remaining_days_in_plan_reached"
	CustomerAlertAlertTypeLowRemainingCreditPercentageReached               CustomerAlertAlertType = "low_remaining_credit_percentage_reached"
	CustomerAlertAlertTypeUsageThresholdReached                             CustomerAlertAlertType = "usage_threshold_reached"
	CustomerAlertAlertTypeLowRemainingDaysForCommitSegmentReached           CustomerAlertAlertType = "low_remaining_days_for_commit_segment_reached"
	CustomerAlertAlertTypeLowRemainingCommitBalanceReached                  CustomerAlertAlertType = "low_remaining_commit_balance_reached"
	CustomerAlertAlertTypeLowRemainingCommitPercentageReached               CustomerAlertAlertType = "low_remaining_commit_percentage_reached"
	CustomerAlertAlertTypeLowRemainingDaysForContractCreditSegmentReached   CustomerAlertAlertType = "low_remaining_days_for_contract_credit_segment_reached"
	CustomerAlertAlertTypeLowRemainingContractCreditBalanceReached          CustomerAlertAlertType = "low_remaining_contract_credit_balance_reached"
	CustomerAlertAlertTypeLowRemainingContractCreditPercentageReached       CustomerAlertAlertType = "low_remaining_contract_credit_percentage_reached"
	CustomerAlertAlertTypeLowRemainingContractCreditAndCommitBalanceReached CustomerAlertAlertType = "low_remaining_contract_credit_and_commit_balance_reached"
	CustomerAlertAlertTypeInvoiceTotalReached                               CustomerAlertAlertType = "invoice_total_reached"
)

func (r CustomerAlertAlertType) IsKnown() bool {
	switch r {
	case CustomerAlertAlertTypeLowCreditBalanceReached, CustomerAlertAlertTypeSpendThresholdReached, CustomerAlertAlertTypeMonthlyInvoiceTotalSpendThresholdReached, CustomerAlertAlertTypeLowRemainingDaysInPlanReached, CustomerAlertAlertTypeLowRemainingCreditPercentageReached, CustomerAlertAlertTypeUsageThresholdReached, CustomerAlertAlertTypeLowRemainingDaysForCommitSegmentReached, CustomerAlertAlertTypeLowRemainingCommitBalanceReached, CustomerAlertAlertTypeLowRemainingCommitPercentageReached, CustomerAlertAlertTypeLowRemainingDaysForContractCreditSegmentReached, CustomerAlertAlertTypeLowRemainingContractCreditBalanceReached, CustomerAlertAlertTypeLowRemainingContractCreditPercentageReached, CustomerAlertAlertTypeLowRemainingContractCreditAndCommitBalanceReached, CustomerAlertAlertTypeInvoiceTotalReached:
		return true
	}
	return false
}

type CustomerAlertAlertCustomFieldFilter struct {
	Entity CustomerAlertAlertCustomFieldFiltersEntity `json:"entity,required"`
	Key    string                                     `json:"key,required"`
	Value  string                                     `json:"value,required"`
	JSON   customerAlertAlertCustomFieldFilterJSON    `json:"-"`
}

// customerAlertAlertCustomFieldFilterJSON contains the JSON metadata for the
// struct [CustomerAlertAlertCustomFieldFilter]
type customerAlertAlertCustomFieldFilterJSON struct {
	Entity      apijson.Field
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertAlertCustomFieldFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertAlertCustomFieldFilterJSON) RawJSON() string {
	return r.raw
}

type CustomerAlertAlertCustomFieldFiltersEntity string

const (
	CustomerAlertAlertCustomFieldFiltersEntityContract       CustomerAlertAlertCustomFieldFiltersEntity = "Contract"
	CustomerAlertAlertCustomFieldFiltersEntityCommit         CustomerAlertAlertCustomFieldFiltersEntity = "Commit"
	CustomerAlertAlertCustomFieldFiltersEntityContractCredit CustomerAlertAlertCustomFieldFiltersEntity = "ContractCredit"
)

func (r CustomerAlertAlertCustomFieldFiltersEntity) IsKnown() bool {
	switch r {
	case CustomerAlertAlertCustomFieldFiltersEntityContract, CustomerAlertAlertCustomFieldFiltersEntityCommit, CustomerAlertAlertCustomFieldFiltersEntityContractCredit:
		return true
	}
	return false
}

// Scopes alert evaluation to a specific presentation group key on individual line
// items. Only present for spend alerts.
type CustomerAlertAlertGroupKeyFilter struct {
	Key   string                               `json:"key,required"`
	Value string                               `json:"value,required"`
	JSON  customerAlertAlertGroupKeyFilterJSON `json:"-"`
}

// customerAlertAlertGroupKeyFilterJSON contains the JSON metadata for the struct
// [CustomerAlertAlertGroupKeyFilter]
type customerAlertAlertGroupKeyFilterJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertAlertGroupKeyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertAlertGroupKeyFilterJSON) RawJSON() string {
	return r.raw
}

// The status of the customer alert. If the alert is archived, null will be
// returned.
type CustomerAlertCustomerStatus string

const (
	CustomerAlertCustomerStatusOk         CustomerAlertCustomerStatus = "ok"
	CustomerAlertCustomerStatusInAlarm    CustomerAlertCustomerStatus = "in_alarm"
	CustomerAlertCustomerStatusEvaluating CustomerAlertCustomerStatus = "evaluating"
)

func (r CustomerAlertCustomerStatus) IsKnown() bool {
	switch r {
	case CustomerAlertCustomerStatusOk, CustomerAlertCustomerStatusInAlarm, CustomerAlertCustomerStatusEvaluating:
		return true
	}
	return false
}

type V1CustomerAlertGetResponse struct {
	Data CustomerAlert                  `json:"data,required"`
	JSON v1CustomerAlertGetResponseJSON `json:"-"`
}

// v1CustomerAlertGetResponseJSON contains the JSON metadata for the struct
// [V1CustomerAlertGetResponse]
type v1CustomerAlertGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerAlertGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerAlertListResponse struct {
	Data     []CustomerAlert                 `json:"data,required"`
	NextPage string                          `json:"next_page,required,nullable"`
	JSON     v1CustomerAlertListResponseJSON `json:"-"`
}

// v1CustomerAlertListResponseJSON contains the JSON metadata for the struct
// [V1CustomerAlertListResponse]
type v1CustomerAlertListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerAlertListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertListResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerAlertGetParams struct {
	// The Metronome ID of the alert
	AlertID param.Field[string] `json:"alert_id,required" format:"uuid"`
	// The Metronome ID of the customer
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// When parallel alerts are enabled during migration, this flag denotes whether to
	// fetch alerts for plans or contracts.
	PlansOrContracts param.Field[V1CustomerAlertGetParamsPlansOrContracts] `json:"plans_or_contracts"`
}

func (r V1CustomerAlertGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When parallel alerts are enabled during migration, this flag denotes whether to
// fetch alerts for plans or contracts.
type V1CustomerAlertGetParamsPlansOrContracts string

const (
	V1CustomerAlertGetParamsPlansOrContractsPlans     V1CustomerAlertGetParamsPlansOrContracts = "PLANS"
	V1CustomerAlertGetParamsPlansOrContractsContracts V1CustomerAlertGetParamsPlansOrContracts = "CONTRACTS"
)

func (r V1CustomerAlertGetParamsPlansOrContracts) IsKnown() bool {
	switch r {
	case V1CustomerAlertGetParamsPlansOrContractsPlans, V1CustomerAlertGetParamsPlansOrContractsContracts:
		return true
	}
	return false
}

type V1CustomerAlertListParams struct {
	// The Metronome ID of the customer
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Optionally filter by alert status. If absent, only enabled alerts will be
	// returned.
	AlertStatuses param.Field[[]V1CustomerAlertListParamsAlertStatus] `json:"alert_statuses"`
}

func (r V1CustomerAlertListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V1CustomerAlertListParams]'s query parameters as
// `url.Values`.
func (r V1CustomerAlertListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerAlertListParamsAlertStatus string

const (
	V1CustomerAlertListParamsAlertStatusEnabled  V1CustomerAlertListParamsAlertStatus = "ENABLED"
	V1CustomerAlertListParamsAlertStatusDisabled V1CustomerAlertListParamsAlertStatus = "DISABLED"
	V1CustomerAlertListParamsAlertStatusArchived V1CustomerAlertListParamsAlertStatus = "ARCHIVED"
)

func (r V1CustomerAlertListParamsAlertStatus) IsKnown() bool {
	switch r {
	case V1CustomerAlertListParamsAlertStatusEnabled, V1CustomerAlertListParamsAlertStatusDisabled, V1CustomerAlertListParamsAlertStatusArchived:
		return true
	}
	return false
}

type V1CustomerAlertResetParams struct {
	// The Metronome ID of the alert
	AlertID param.Field[string] `json:"alert_id,required" format:"uuid"`
	// The Metronome ID of the customer
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r V1CustomerAlertResetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
