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
)

// CustomerAlertService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerAlertService] method instead.
type CustomerAlertService struct {
	Options []option.RequestOption
}

// NewCustomerAlertService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerAlertService(opts ...option.RequestOption) (r *CustomerAlertService) {
	r = &CustomerAlertService{}
	r.Options = opts
	return
}

// Get the customer alert status and alert information for the specified customer
// and alert
func (r *CustomerAlertService) Get(ctx context.Context, body CustomerAlertGetParams, opts ...option.RequestOption) (res *CustomerAlertGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer-alerts/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch all customer alert statuses and alert information for a customer
func (r *CustomerAlertService) List(ctx context.Context, params CustomerAlertListParams, opts ...option.RequestOption) (res *CustomerAlertListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer-alerts/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Reset state for an alert by customer id and force re-evaluation
func (r *CustomerAlertService) Reset(ctx context.Context, body CustomerAlertResetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "customer-alerts/reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CustomerAlertGetResponse struct {
	Data CustomerAlertGetResponseData `json:"data,required"`
	JSON customerAlertGetResponseJSON `json:"-"`
}

// customerAlertGetResponseJSON contains the JSON metadata for the struct
// [CustomerAlertGetResponse]
type customerAlertGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertGetResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerAlertGetResponseData struct {
	Alert CustomerAlertGetResponseDataAlert `json:"alert,required"`
	// The status of the customer alert. If the alert is archived, null will be
	// returned.
	CustomerStatus CustomerAlertGetResponseDataCustomerStatus `json:"customer_status,required,nullable"`
	// If present, indicates the reason the alert was triggered.
	TriggeredBy string                           `json:"triggered_by,nullable"`
	JSON        customerAlertGetResponseDataJSON `json:"-"`
}

// customerAlertGetResponseDataJSON contains the JSON metadata for the struct
// [CustomerAlertGetResponseData]
type customerAlertGetResponseDataJSON struct {
	Alert          apijson.Field
	CustomerStatus apijson.Field
	TriggeredBy    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerAlertGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type CustomerAlertGetResponseDataAlert struct {
	// the Metronome ID of the alert
	ID string `json:"id,required"`
	// Name of the alert
	Name string `json:"name,required"`
	// Status of the alert
	Status CustomerAlertGetResponseDataAlertStatus `json:"status,required"`
	// Threshold value of the alert policy
	Threshold float64 `json:"threshold,required"`
	// Type of the alert
	Type CustomerAlertGetResponseDataAlertType `json:"type,required"`
	// Timestamp for when the alert was last updated
	UpdatedAt  time.Time                                   `json:"updated_at,required" format:"date-time"`
	CreditType CustomerAlertGetResponseDataAlertCreditType `json:"credit_type,nullable"`
	// A list of custom field filters for alert types that support advanced filtering
	CustomFieldFilters []CustomerAlertGetResponseDataAlertCustomFieldFilter `json:"custom_field_filters"`
	// Scopes alert evaluation to a specific presentation group key on individual line
	// items. Only present for spend alerts.
	GroupKeyFilter CustomerAlertGetResponseDataAlertGroupKeyFilter `json:"group_key_filter"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                                `json:"uniqueness_key"`
	JSON          customerAlertGetResponseDataAlertJSON `json:"-"`
}

// customerAlertGetResponseDataAlertJSON contains the JSON metadata for the struct
// [CustomerAlertGetResponseDataAlert]
type customerAlertGetResponseDataAlertJSON struct {
	ID                 apijson.Field
	Name               apijson.Field
	Status             apijson.Field
	Threshold          apijson.Field
	Type               apijson.Field
	UpdatedAt          apijson.Field
	CreditType         apijson.Field
	CustomFieldFilters apijson.Field
	GroupKeyFilter     apijson.Field
	UniquenessKey      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerAlertGetResponseDataAlert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertGetResponseDataAlertJSON) RawJSON() string {
	return r.raw
}

// Status of the alert
type CustomerAlertGetResponseDataAlertStatus string

const (
	CustomerAlertGetResponseDataAlertStatusEnabled  CustomerAlertGetResponseDataAlertStatus = "enabled"
	CustomerAlertGetResponseDataAlertStatusArchived CustomerAlertGetResponseDataAlertStatus = "archived"
	CustomerAlertGetResponseDataAlertStatusDisabled CustomerAlertGetResponseDataAlertStatus = "disabled"
)

func (r CustomerAlertGetResponseDataAlertStatus) IsKnown() bool {
	switch r {
	case CustomerAlertGetResponseDataAlertStatusEnabled, CustomerAlertGetResponseDataAlertStatusArchived, CustomerAlertGetResponseDataAlertStatusDisabled:
		return true
	}
	return false
}

// Type of the alert
type CustomerAlertGetResponseDataAlertType string

const (
	CustomerAlertGetResponseDataAlertTypeLowCreditBalanceReached                         CustomerAlertGetResponseDataAlertType = "low_credit_balance_reached"
	CustomerAlertGetResponseDataAlertTypeSpendThresholdReached                           CustomerAlertGetResponseDataAlertType = "spend_threshold_reached"
	CustomerAlertGetResponseDataAlertTypeMonthlyInvoiceTotalSpendThresholdReached        CustomerAlertGetResponseDataAlertType = "monthly_invoice_total_spend_threshold_reached"
	CustomerAlertGetResponseDataAlertTypeLowRemainingDaysInPlanReached                   CustomerAlertGetResponseDataAlertType = "low_remaining_days_in_plan_reached"
	CustomerAlertGetResponseDataAlertTypeLowRemainingCreditPercentageReached             CustomerAlertGetResponseDataAlertType = "low_remaining_credit_percentage_reached"
	CustomerAlertGetResponseDataAlertTypeUsageThresholdReached                           CustomerAlertGetResponseDataAlertType = "usage_threshold_reached"
	CustomerAlertGetResponseDataAlertTypeLowRemainingDaysForCommitSegmentReached         CustomerAlertGetResponseDataAlertType = "low_remaining_days_for_commit_segment_reached"
	CustomerAlertGetResponseDataAlertTypeLowRemainingCommitBalanceReached                CustomerAlertGetResponseDataAlertType = "low_remaining_commit_balance_reached"
	CustomerAlertGetResponseDataAlertTypeLowRemainingCommitPercentageReached             CustomerAlertGetResponseDataAlertType = "low_remaining_commit_percentage_reached"
	CustomerAlertGetResponseDataAlertTypeLowRemainingDaysForContractCreditSegmentReached CustomerAlertGetResponseDataAlertType = "low_remaining_days_for_contract_credit_segment_reached"
	CustomerAlertGetResponseDataAlertTypeLowRemainingContractCreditBalanceReached        CustomerAlertGetResponseDataAlertType = "low_remaining_contract_credit_balance_reached"
	CustomerAlertGetResponseDataAlertTypeLowRemainingContractCreditPercentageReached     CustomerAlertGetResponseDataAlertType = "low_remaining_contract_credit_percentage_reached"
)

func (r CustomerAlertGetResponseDataAlertType) IsKnown() bool {
	switch r {
	case CustomerAlertGetResponseDataAlertTypeLowCreditBalanceReached, CustomerAlertGetResponseDataAlertTypeSpendThresholdReached, CustomerAlertGetResponseDataAlertTypeMonthlyInvoiceTotalSpendThresholdReached, CustomerAlertGetResponseDataAlertTypeLowRemainingDaysInPlanReached, CustomerAlertGetResponseDataAlertTypeLowRemainingCreditPercentageReached, CustomerAlertGetResponseDataAlertTypeUsageThresholdReached, CustomerAlertGetResponseDataAlertTypeLowRemainingDaysForCommitSegmentReached, CustomerAlertGetResponseDataAlertTypeLowRemainingCommitBalanceReached, CustomerAlertGetResponseDataAlertTypeLowRemainingCommitPercentageReached, CustomerAlertGetResponseDataAlertTypeLowRemainingDaysForContractCreditSegmentReached, CustomerAlertGetResponseDataAlertTypeLowRemainingContractCreditBalanceReached, CustomerAlertGetResponseDataAlertTypeLowRemainingContractCreditPercentageReached:
		return true
	}
	return false
}

type CustomerAlertGetResponseDataAlertCreditType struct {
	ID   string                                          `json:"id,required" format:"uuid"`
	Name string                                          `json:"name,required"`
	JSON customerAlertGetResponseDataAlertCreditTypeJSON `json:"-"`
}

// customerAlertGetResponseDataAlertCreditTypeJSON contains the JSON metadata for
// the struct [CustomerAlertGetResponseDataAlertCreditType]
type customerAlertGetResponseDataAlertCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertGetResponseDataAlertCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertGetResponseDataAlertCreditTypeJSON) RawJSON() string {
	return r.raw
}

type CustomerAlertGetResponseDataAlertCustomFieldFilter struct {
	Entity CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity `json:"entity,required"`
	Key    string                                                    `json:"key,required"`
	Value  string                                                    `json:"value,required"`
	JSON   customerAlertGetResponseDataAlertCustomFieldFilterJSON    `json:"-"`
}

// customerAlertGetResponseDataAlertCustomFieldFilterJSON contains the JSON
// metadata for the struct [CustomerAlertGetResponseDataAlertCustomFieldFilter]
type customerAlertGetResponseDataAlertCustomFieldFilterJSON struct {
	Entity      apijson.Field
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertGetResponseDataAlertCustomFieldFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertGetResponseDataAlertCustomFieldFilterJSON) RawJSON() string {
	return r.raw
}

type CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity string

const (
	CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityContract       CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity = "Contract"
	CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityCommit         CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity = "Commit"
	CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityContractCredit CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity = "ContractCredit"
)

func (r CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity) IsKnown() bool {
	switch r {
	case CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityContract, CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityCommit, CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityContractCredit:
		return true
	}
	return false
}

// Scopes alert evaluation to a specific presentation group key on individual line
// items. Only present for spend alerts.
type CustomerAlertGetResponseDataAlertGroupKeyFilter struct {
	Key   string                                              `json:"key,required"`
	Value string                                              `json:"value,required"`
	JSON  customerAlertGetResponseDataAlertGroupKeyFilterJSON `json:"-"`
}

// customerAlertGetResponseDataAlertGroupKeyFilterJSON contains the JSON metadata
// for the struct [CustomerAlertGetResponseDataAlertGroupKeyFilter]
type customerAlertGetResponseDataAlertGroupKeyFilterJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertGetResponseDataAlertGroupKeyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertGetResponseDataAlertGroupKeyFilterJSON) RawJSON() string {
	return r.raw
}

// The status of the customer alert. If the alert is archived, null will be
// returned.
type CustomerAlertGetResponseDataCustomerStatus string

const (
	CustomerAlertGetResponseDataCustomerStatusOk         CustomerAlertGetResponseDataCustomerStatus = "ok"
	CustomerAlertGetResponseDataCustomerStatusInAlarm    CustomerAlertGetResponseDataCustomerStatus = "in_alarm"
	CustomerAlertGetResponseDataCustomerStatusEvaluating CustomerAlertGetResponseDataCustomerStatus = "evaluating"
)

func (r CustomerAlertGetResponseDataCustomerStatus) IsKnown() bool {
	switch r {
	case CustomerAlertGetResponseDataCustomerStatusOk, CustomerAlertGetResponseDataCustomerStatusInAlarm, CustomerAlertGetResponseDataCustomerStatusEvaluating:
		return true
	}
	return false
}

type CustomerAlertListResponse struct {
	Data     []CustomerAlertListResponseData `json:"data,required"`
	NextPage string                          `json:"next_page,required,nullable"`
	JSON     customerAlertListResponseJSON   `json:"-"`
}

// customerAlertListResponseJSON contains the JSON metadata for the struct
// [CustomerAlertListResponse]
type customerAlertListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerAlertListResponseData struct {
	Alert CustomerAlertListResponseDataAlert `json:"alert,required"`
	// The status of the customer alert. If the alert is archived, null will be
	// returned.
	CustomerStatus CustomerAlertListResponseDataCustomerStatus `json:"customer_status,required,nullable"`
	// If present, indicates the reason the alert was triggered.
	TriggeredBy string                            `json:"triggered_by,nullable"`
	JSON        customerAlertListResponseDataJSON `json:"-"`
}

// customerAlertListResponseDataJSON contains the JSON metadata for the struct
// [CustomerAlertListResponseData]
type customerAlertListResponseDataJSON struct {
	Alert          apijson.Field
	CustomerStatus apijson.Field
	TriggeredBy    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerAlertListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertListResponseDataJSON) RawJSON() string {
	return r.raw
}

type CustomerAlertListResponseDataAlert struct {
	// the Metronome ID of the alert
	ID string `json:"id,required"`
	// Name of the alert
	Name string `json:"name,required"`
	// Status of the alert
	Status CustomerAlertListResponseDataAlertStatus `json:"status,required"`
	// Threshold value of the alert policy
	Threshold float64 `json:"threshold,required"`
	// Type of the alert
	Type CustomerAlertListResponseDataAlertType `json:"type,required"`
	// Timestamp for when the alert was last updated
	UpdatedAt  time.Time                                    `json:"updated_at,required" format:"date-time"`
	CreditType CustomerAlertListResponseDataAlertCreditType `json:"credit_type,nullable"`
	// A list of custom field filters for alert types that support advanced filtering
	CustomFieldFilters []CustomerAlertListResponseDataAlertCustomFieldFilter `json:"custom_field_filters"`
	// Scopes alert evaluation to a specific presentation group key on individual line
	// items. Only present for spend alerts.
	GroupKeyFilter CustomerAlertListResponseDataAlertGroupKeyFilter `json:"group_key_filter"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                                 `json:"uniqueness_key"`
	JSON          customerAlertListResponseDataAlertJSON `json:"-"`
}

// customerAlertListResponseDataAlertJSON contains the JSON metadata for the struct
// [CustomerAlertListResponseDataAlert]
type customerAlertListResponseDataAlertJSON struct {
	ID                 apijson.Field
	Name               apijson.Field
	Status             apijson.Field
	Threshold          apijson.Field
	Type               apijson.Field
	UpdatedAt          apijson.Field
	CreditType         apijson.Field
	CustomFieldFilters apijson.Field
	GroupKeyFilter     apijson.Field
	UniquenessKey      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerAlertListResponseDataAlert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertListResponseDataAlertJSON) RawJSON() string {
	return r.raw
}

// Status of the alert
type CustomerAlertListResponseDataAlertStatus string

const (
	CustomerAlertListResponseDataAlertStatusEnabled  CustomerAlertListResponseDataAlertStatus = "enabled"
	CustomerAlertListResponseDataAlertStatusArchived CustomerAlertListResponseDataAlertStatus = "archived"
	CustomerAlertListResponseDataAlertStatusDisabled CustomerAlertListResponseDataAlertStatus = "disabled"
)

func (r CustomerAlertListResponseDataAlertStatus) IsKnown() bool {
	switch r {
	case CustomerAlertListResponseDataAlertStatusEnabled, CustomerAlertListResponseDataAlertStatusArchived, CustomerAlertListResponseDataAlertStatusDisabled:
		return true
	}
	return false
}

// Type of the alert
type CustomerAlertListResponseDataAlertType string

const (
	CustomerAlertListResponseDataAlertTypeLowCreditBalanceReached                         CustomerAlertListResponseDataAlertType = "low_credit_balance_reached"
	CustomerAlertListResponseDataAlertTypeSpendThresholdReached                           CustomerAlertListResponseDataAlertType = "spend_threshold_reached"
	CustomerAlertListResponseDataAlertTypeMonthlyInvoiceTotalSpendThresholdReached        CustomerAlertListResponseDataAlertType = "monthly_invoice_total_spend_threshold_reached"
	CustomerAlertListResponseDataAlertTypeLowRemainingDaysInPlanReached                   CustomerAlertListResponseDataAlertType = "low_remaining_days_in_plan_reached"
	CustomerAlertListResponseDataAlertTypeLowRemainingCreditPercentageReached             CustomerAlertListResponseDataAlertType = "low_remaining_credit_percentage_reached"
	CustomerAlertListResponseDataAlertTypeUsageThresholdReached                           CustomerAlertListResponseDataAlertType = "usage_threshold_reached"
	CustomerAlertListResponseDataAlertTypeLowRemainingDaysForCommitSegmentReached         CustomerAlertListResponseDataAlertType = "low_remaining_days_for_commit_segment_reached"
	CustomerAlertListResponseDataAlertTypeLowRemainingCommitBalanceReached                CustomerAlertListResponseDataAlertType = "low_remaining_commit_balance_reached"
	CustomerAlertListResponseDataAlertTypeLowRemainingCommitPercentageReached             CustomerAlertListResponseDataAlertType = "low_remaining_commit_percentage_reached"
	CustomerAlertListResponseDataAlertTypeLowRemainingDaysForContractCreditSegmentReached CustomerAlertListResponseDataAlertType = "low_remaining_days_for_contract_credit_segment_reached"
	CustomerAlertListResponseDataAlertTypeLowRemainingContractCreditBalanceReached        CustomerAlertListResponseDataAlertType = "low_remaining_contract_credit_balance_reached"
	CustomerAlertListResponseDataAlertTypeLowRemainingContractCreditPercentageReached     CustomerAlertListResponseDataAlertType = "low_remaining_contract_credit_percentage_reached"
)

func (r CustomerAlertListResponseDataAlertType) IsKnown() bool {
	switch r {
	case CustomerAlertListResponseDataAlertTypeLowCreditBalanceReached, CustomerAlertListResponseDataAlertTypeSpendThresholdReached, CustomerAlertListResponseDataAlertTypeMonthlyInvoiceTotalSpendThresholdReached, CustomerAlertListResponseDataAlertTypeLowRemainingDaysInPlanReached, CustomerAlertListResponseDataAlertTypeLowRemainingCreditPercentageReached, CustomerAlertListResponseDataAlertTypeUsageThresholdReached, CustomerAlertListResponseDataAlertTypeLowRemainingDaysForCommitSegmentReached, CustomerAlertListResponseDataAlertTypeLowRemainingCommitBalanceReached, CustomerAlertListResponseDataAlertTypeLowRemainingCommitPercentageReached, CustomerAlertListResponseDataAlertTypeLowRemainingDaysForContractCreditSegmentReached, CustomerAlertListResponseDataAlertTypeLowRemainingContractCreditBalanceReached, CustomerAlertListResponseDataAlertTypeLowRemainingContractCreditPercentageReached:
		return true
	}
	return false
}

type CustomerAlertListResponseDataAlertCreditType struct {
	ID   string                                           `json:"id,required" format:"uuid"`
	Name string                                           `json:"name,required"`
	JSON customerAlertListResponseDataAlertCreditTypeJSON `json:"-"`
}

// customerAlertListResponseDataAlertCreditTypeJSON contains the JSON metadata for
// the struct [CustomerAlertListResponseDataAlertCreditType]
type customerAlertListResponseDataAlertCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertListResponseDataAlertCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertListResponseDataAlertCreditTypeJSON) RawJSON() string {
	return r.raw
}

type CustomerAlertListResponseDataAlertCustomFieldFilter struct {
	Entity CustomerAlertListResponseDataAlertCustomFieldFiltersEntity `json:"entity,required"`
	Key    string                                                     `json:"key,required"`
	Value  string                                                     `json:"value,required"`
	JSON   customerAlertListResponseDataAlertCustomFieldFilterJSON    `json:"-"`
}

// customerAlertListResponseDataAlertCustomFieldFilterJSON contains the JSON
// metadata for the struct [CustomerAlertListResponseDataAlertCustomFieldFilter]
type customerAlertListResponseDataAlertCustomFieldFilterJSON struct {
	Entity      apijson.Field
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertListResponseDataAlertCustomFieldFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertListResponseDataAlertCustomFieldFilterJSON) RawJSON() string {
	return r.raw
}

type CustomerAlertListResponseDataAlertCustomFieldFiltersEntity string

const (
	CustomerAlertListResponseDataAlertCustomFieldFiltersEntityContract       CustomerAlertListResponseDataAlertCustomFieldFiltersEntity = "Contract"
	CustomerAlertListResponseDataAlertCustomFieldFiltersEntityCommit         CustomerAlertListResponseDataAlertCustomFieldFiltersEntity = "Commit"
	CustomerAlertListResponseDataAlertCustomFieldFiltersEntityContractCredit CustomerAlertListResponseDataAlertCustomFieldFiltersEntity = "ContractCredit"
)

func (r CustomerAlertListResponseDataAlertCustomFieldFiltersEntity) IsKnown() bool {
	switch r {
	case CustomerAlertListResponseDataAlertCustomFieldFiltersEntityContract, CustomerAlertListResponseDataAlertCustomFieldFiltersEntityCommit, CustomerAlertListResponseDataAlertCustomFieldFiltersEntityContractCredit:
		return true
	}
	return false
}

// Scopes alert evaluation to a specific presentation group key on individual line
// items. Only present for spend alerts.
type CustomerAlertListResponseDataAlertGroupKeyFilter struct {
	Key   string                                               `json:"key,required"`
	Value string                                               `json:"value,required"`
	JSON  customerAlertListResponseDataAlertGroupKeyFilterJSON `json:"-"`
}

// customerAlertListResponseDataAlertGroupKeyFilterJSON contains the JSON metadata
// for the struct [CustomerAlertListResponseDataAlertGroupKeyFilter]
type customerAlertListResponseDataAlertGroupKeyFilterJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertListResponseDataAlertGroupKeyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAlertListResponseDataAlertGroupKeyFilterJSON) RawJSON() string {
	return r.raw
}

// The status of the customer alert. If the alert is archived, null will be
// returned.
type CustomerAlertListResponseDataCustomerStatus string

const (
	CustomerAlertListResponseDataCustomerStatusOk         CustomerAlertListResponseDataCustomerStatus = "ok"
	CustomerAlertListResponseDataCustomerStatusInAlarm    CustomerAlertListResponseDataCustomerStatus = "in_alarm"
	CustomerAlertListResponseDataCustomerStatusEvaluating CustomerAlertListResponseDataCustomerStatus = "evaluating"
)

func (r CustomerAlertListResponseDataCustomerStatus) IsKnown() bool {
	switch r {
	case CustomerAlertListResponseDataCustomerStatusOk, CustomerAlertListResponseDataCustomerStatusInAlarm, CustomerAlertListResponseDataCustomerStatusEvaluating:
		return true
	}
	return false
}

type CustomerAlertGetParams struct {
	// The Metronome ID of the alert
	AlertID param.Field[string] `json:"alert_id,required" format:"uuid"`
	// The Metronome ID of the customer
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r CustomerAlertGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerAlertListParams struct {
	// The Metronome ID of the customer
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Optionally filter by alert status. If absent, only enabled alerts will be
	// returned.
	AlertStatuses param.Field[[]CustomerAlertListParamsAlertStatus] `json:"alert_statuses"`
}

func (r CustomerAlertListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [CustomerAlertListParams]'s query parameters as
// `url.Values`.
func (r CustomerAlertListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerAlertListParamsAlertStatus string

const (
	CustomerAlertListParamsAlertStatusEnabled  CustomerAlertListParamsAlertStatus = "enabled"
	CustomerAlertListParamsAlertStatusDisabled CustomerAlertListParamsAlertStatus = "disabled"
	CustomerAlertListParamsAlertStatusArchived CustomerAlertListParamsAlertStatus = "archived"
	CustomerAlertListParamsAlertStatusEnabled  CustomerAlertListParamsAlertStatus = "ENABLED"
	CustomerAlertListParamsAlertStatusDisabled CustomerAlertListParamsAlertStatus = "DISABLED"
	CustomerAlertListParamsAlertStatusArchived CustomerAlertListParamsAlertStatus = "ARCHIVED"
	CustomerAlertListParamsAlertStatusEnabled  CustomerAlertListParamsAlertStatus = "Enabled"
	CustomerAlertListParamsAlertStatusDisabled CustomerAlertListParamsAlertStatus = "Disabled"
	CustomerAlertListParamsAlertStatusArchived CustomerAlertListParamsAlertStatus = "Archived"
)

func (r CustomerAlertListParamsAlertStatus) IsKnown() bool {
	switch r {
	case CustomerAlertListParamsAlertStatusEnabled, CustomerAlertListParamsAlertStatusDisabled, CustomerAlertListParamsAlertStatusArchived, CustomerAlertListParamsAlertStatusEnabled, CustomerAlertListParamsAlertStatusDisabled, CustomerAlertListParamsAlertStatusArchived, CustomerAlertListParamsAlertStatusEnabled, CustomerAlertListParamsAlertStatusDisabled, CustomerAlertListParamsAlertStatusArchived:
		return true
	}
	return false
}

type CustomerAlertResetParams struct {
	// The Metronome ID of the alert
	AlertID param.Field[string] `json:"alert_id,required" format:"uuid"`
	// The Metronome ID of the customer
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r CustomerAlertResetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
