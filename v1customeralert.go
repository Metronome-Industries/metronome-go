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

type V1CustomerAlertGetResponse struct {
	Data V1CustomerAlertGetResponseData `json:"data,required"`
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

type V1CustomerAlertGetResponseData struct {
	Alert V1CustomerAlertGetResponseDataAlert `json:"alert,required"`
	// The status of the customer alert. If the alert is archived, null will be
	// returned.
	CustomerStatus V1CustomerAlertGetResponseDataCustomerStatus `json:"customer_status,required,nullable"`
	// If present, indicates the reason the alert was triggered.
	TriggeredBy string                             `json:"triggered_by,nullable"`
	JSON        v1CustomerAlertGetResponseDataJSON `json:"-"`
}

// v1CustomerAlertGetResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerAlertGetResponseData]
type v1CustomerAlertGetResponseDataJSON struct {
	Alert          apijson.Field
	CustomerStatus apijson.Field
	TriggeredBy    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1CustomerAlertGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerAlertGetResponseDataAlert struct {
	// the Metronome ID of the alert
	ID string `json:"id,required"`
	// Name of the alert
	Name string `json:"name,required"`
	// Status of the alert
	Status V1CustomerAlertGetResponseDataAlertStatus `json:"status,required"`
	// Threshold value of the alert policy
	Threshold float64 `json:"threshold,required"`
	// Type of the alert
	Type V1CustomerAlertGetResponseDataAlertType `json:"type,required"`
	// Timestamp for when the alert was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An array of strings, representing a way to filter the credit grant this alert
	// applies to, by looking at the credit_grant_type field on the credit grant. This
	// field is only defined for CreditPercentage and CreditBalance alerts
	CreditGrantTypeFilters []string                                      `json:"credit_grant_type_filters"`
	CreditType             V1CustomerAlertGetResponseDataAlertCreditType `json:"credit_type,nullable"`
	// A list of custom field filters for alert types that support advanced filtering
	CustomFieldFilters []V1CustomerAlertGetResponseDataAlertCustomFieldFilter `json:"custom_field_filters"`
	// Scopes alert evaluation to a specific presentation group key on individual line
	// items. Only present for spend alerts.
	GroupKeyFilter V1CustomerAlertGetResponseDataAlertGroupKeyFilter `json:"group_key_filter"`
	// Only present for `spend_threshold_reached` alerts. Scope alert to a specific
	// group key on individual line items.
	GroupValues []V1CustomerAlertGetResponseDataAlertGroupValue `json:"group_values"`
	// Only supported for invoice_total_reached alerts. A list of invoice types to
	// evaluate.
	InvoiceTypesFilter []string `json:"invoice_types_filter"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                                  `json:"uniqueness_key"`
	JSON          v1CustomerAlertGetResponseDataAlertJSON `json:"-"`
}

// v1CustomerAlertGetResponseDataAlertJSON contains the JSON metadata for the
// struct [V1CustomerAlertGetResponseDataAlert]
type v1CustomerAlertGetResponseDataAlertJSON struct {
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
	GroupValues            apijson.Field
	InvoiceTypesFilter     apijson.Field
	UniquenessKey          apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1CustomerAlertGetResponseDataAlert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertGetResponseDataAlertJSON) RawJSON() string {
	return r.raw
}

// Status of the alert
type V1CustomerAlertGetResponseDataAlertStatus string

const (
	V1CustomerAlertGetResponseDataAlertStatusEnabled  V1CustomerAlertGetResponseDataAlertStatus = "enabled"
	V1CustomerAlertGetResponseDataAlertStatusArchived V1CustomerAlertGetResponseDataAlertStatus = "archived"
	V1CustomerAlertGetResponseDataAlertStatusDisabled V1CustomerAlertGetResponseDataAlertStatus = "disabled"
)

func (r V1CustomerAlertGetResponseDataAlertStatus) IsKnown() bool {
	switch r {
	case V1CustomerAlertGetResponseDataAlertStatusEnabled, V1CustomerAlertGetResponseDataAlertStatusArchived, V1CustomerAlertGetResponseDataAlertStatusDisabled:
		return true
	}
	return false
}

// Type of the alert
type V1CustomerAlertGetResponseDataAlertType string

const (
	V1CustomerAlertGetResponseDataAlertTypeLowCreditBalanceReached                           V1CustomerAlertGetResponseDataAlertType = "low_credit_balance_reached"
	V1CustomerAlertGetResponseDataAlertTypeSpendThresholdReached                             V1CustomerAlertGetResponseDataAlertType = "spend_threshold_reached"
	V1CustomerAlertGetResponseDataAlertTypeMonthlyInvoiceTotalSpendThresholdReached          V1CustomerAlertGetResponseDataAlertType = "monthly_invoice_total_spend_threshold_reached"
	V1CustomerAlertGetResponseDataAlertTypeLowRemainingDaysInPlanReached                     V1CustomerAlertGetResponseDataAlertType = "low_remaining_days_in_plan_reached"
	V1CustomerAlertGetResponseDataAlertTypeLowRemainingCreditPercentageReached               V1CustomerAlertGetResponseDataAlertType = "low_remaining_credit_percentage_reached"
	V1CustomerAlertGetResponseDataAlertTypeUsageThresholdReached                             V1CustomerAlertGetResponseDataAlertType = "usage_threshold_reached"
	V1CustomerAlertGetResponseDataAlertTypeLowRemainingDaysForCommitSegmentReached           V1CustomerAlertGetResponseDataAlertType = "low_remaining_days_for_commit_segment_reached"
	V1CustomerAlertGetResponseDataAlertTypeLowRemainingCommitBalanceReached                  V1CustomerAlertGetResponseDataAlertType = "low_remaining_commit_balance_reached"
	V1CustomerAlertGetResponseDataAlertTypeLowRemainingCommitPercentageReached               V1CustomerAlertGetResponseDataAlertType = "low_remaining_commit_percentage_reached"
	V1CustomerAlertGetResponseDataAlertTypeLowRemainingDaysForContractCreditSegmentReached   V1CustomerAlertGetResponseDataAlertType = "low_remaining_days_for_contract_credit_segment_reached"
	V1CustomerAlertGetResponseDataAlertTypeLowRemainingContractCreditBalanceReached          V1CustomerAlertGetResponseDataAlertType = "low_remaining_contract_credit_balance_reached"
	V1CustomerAlertGetResponseDataAlertTypeLowRemainingContractCreditPercentageReached       V1CustomerAlertGetResponseDataAlertType = "low_remaining_contract_credit_percentage_reached"
	V1CustomerAlertGetResponseDataAlertTypeLowRemainingContractCreditAndCommitBalanceReached V1CustomerAlertGetResponseDataAlertType = "low_remaining_contract_credit_and_commit_balance_reached"
	V1CustomerAlertGetResponseDataAlertTypeInvoiceTotalReached                               V1CustomerAlertGetResponseDataAlertType = "invoice_total_reached"
)

func (r V1CustomerAlertGetResponseDataAlertType) IsKnown() bool {
	switch r {
	case V1CustomerAlertGetResponseDataAlertTypeLowCreditBalanceReached, V1CustomerAlertGetResponseDataAlertTypeSpendThresholdReached, V1CustomerAlertGetResponseDataAlertTypeMonthlyInvoiceTotalSpendThresholdReached, V1CustomerAlertGetResponseDataAlertTypeLowRemainingDaysInPlanReached, V1CustomerAlertGetResponseDataAlertTypeLowRemainingCreditPercentageReached, V1CustomerAlertGetResponseDataAlertTypeUsageThresholdReached, V1CustomerAlertGetResponseDataAlertTypeLowRemainingDaysForCommitSegmentReached, V1CustomerAlertGetResponseDataAlertTypeLowRemainingCommitBalanceReached, V1CustomerAlertGetResponseDataAlertTypeLowRemainingCommitPercentageReached, V1CustomerAlertGetResponseDataAlertTypeLowRemainingDaysForContractCreditSegmentReached, V1CustomerAlertGetResponseDataAlertTypeLowRemainingContractCreditBalanceReached, V1CustomerAlertGetResponseDataAlertTypeLowRemainingContractCreditPercentageReached, V1CustomerAlertGetResponseDataAlertTypeLowRemainingContractCreditAndCommitBalanceReached, V1CustomerAlertGetResponseDataAlertTypeInvoiceTotalReached:
		return true
	}
	return false
}

type V1CustomerAlertGetResponseDataAlertCreditType struct {
	ID   string                                            `json:"id,required" format:"uuid"`
	Name string                                            `json:"name,required"`
	JSON v1CustomerAlertGetResponseDataAlertCreditTypeJSON `json:"-"`
}

// v1CustomerAlertGetResponseDataAlertCreditTypeJSON contains the JSON metadata for
// the struct [V1CustomerAlertGetResponseDataAlertCreditType]
type v1CustomerAlertGetResponseDataAlertCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerAlertGetResponseDataAlertCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertGetResponseDataAlertCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerAlertGetResponseDataAlertCustomFieldFilter struct {
	Entity V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity `json:"entity,required"`
	Key    string                                                      `json:"key,required"`
	Value  string                                                      `json:"value,required"`
	JSON   v1CustomerAlertGetResponseDataAlertCustomFieldFilterJSON    `json:"-"`
}

// v1CustomerAlertGetResponseDataAlertCustomFieldFilterJSON contains the JSON
// metadata for the struct [V1CustomerAlertGetResponseDataAlertCustomFieldFilter]
type v1CustomerAlertGetResponseDataAlertCustomFieldFilterJSON struct {
	Entity      apijson.Field
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerAlertGetResponseDataAlertCustomFieldFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertGetResponseDataAlertCustomFieldFilterJSON) RawJSON() string {
	return r.raw
}

type V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity string

const (
	V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityContract       V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity = "Contract"
	V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityCommit         V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity = "Commit"
	V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityContractCredit V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity = "ContractCredit"
)

func (r V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntity) IsKnown() bool {
	switch r {
	case V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityContract, V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityCommit, V1CustomerAlertGetResponseDataAlertCustomFieldFiltersEntityContractCredit:
		return true
	}
	return false
}

// Scopes alert evaluation to a specific presentation group key on individual line
// items. Only present for spend alerts.
type V1CustomerAlertGetResponseDataAlertGroupKeyFilter struct {
	Key   string                                                `json:"key,required"`
	Value string                                                `json:"value,required"`
	JSON  v1CustomerAlertGetResponseDataAlertGroupKeyFilterJSON `json:"-"`
}

// v1CustomerAlertGetResponseDataAlertGroupKeyFilterJSON contains the JSON metadata
// for the struct [V1CustomerAlertGetResponseDataAlertGroupKeyFilter]
type v1CustomerAlertGetResponseDataAlertGroupKeyFilterJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerAlertGetResponseDataAlertGroupKeyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertGetResponseDataAlertGroupKeyFilterJSON) RawJSON() string {
	return r.raw
}

type V1CustomerAlertGetResponseDataAlertGroupValue struct {
	Key   string                                            `json:"key,required"`
	Value string                                            `json:"value,required"`
	JSON  v1CustomerAlertGetResponseDataAlertGroupValueJSON `json:"-"`
}

// v1CustomerAlertGetResponseDataAlertGroupValueJSON contains the JSON metadata for
// the struct [V1CustomerAlertGetResponseDataAlertGroupValue]
type v1CustomerAlertGetResponseDataAlertGroupValueJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerAlertGetResponseDataAlertGroupValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertGetResponseDataAlertGroupValueJSON) RawJSON() string {
	return r.raw
}

// The status of the customer alert. If the alert is archived, null will be
// returned.
type V1CustomerAlertGetResponseDataCustomerStatus string

const (
	V1CustomerAlertGetResponseDataCustomerStatusOk         V1CustomerAlertGetResponseDataCustomerStatus = "ok"
	V1CustomerAlertGetResponseDataCustomerStatusInAlarm    V1CustomerAlertGetResponseDataCustomerStatus = "in_alarm"
	V1CustomerAlertGetResponseDataCustomerStatusEvaluating V1CustomerAlertGetResponseDataCustomerStatus = "evaluating"
)

func (r V1CustomerAlertGetResponseDataCustomerStatus) IsKnown() bool {
	switch r {
	case V1CustomerAlertGetResponseDataCustomerStatusOk, V1CustomerAlertGetResponseDataCustomerStatusInAlarm, V1CustomerAlertGetResponseDataCustomerStatusEvaluating:
		return true
	}
	return false
}

type V1CustomerAlertListResponse struct {
	Data     []V1CustomerAlertListResponseData `json:"data,required"`
	NextPage string                            `json:"next_page,required,nullable"`
	JSON     v1CustomerAlertListResponseJSON   `json:"-"`
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

type V1CustomerAlertListResponseData struct {
	Alert V1CustomerAlertListResponseDataAlert `json:"alert,required"`
	// The status of the customer alert. If the alert is archived, null will be
	// returned.
	CustomerStatus V1CustomerAlertListResponseDataCustomerStatus `json:"customer_status,required,nullable"`
	// If present, indicates the reason the alert was triggered.
	TriggeredBy string                              `json:"triggered_by,nullable"`
	JSON        v1CustomerAlertListResponseDataJSON `json:"-"`
}

// v1CustomerAlertListResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerAlertListResponseData]
type v1CustomerAlertListResponseDataJSON struct {
	Alert          apijson.Field
	CustomerStatus apijson.Field
	TriggeredBy    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1CustomerAlertListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertListResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerAlertListResponseDataAlert struct {
	// the Metronome ID of the alert
	ID string `json:"id,required"`
	// Name of the alert
	Name string `json:"name,required"`
	// Status of the alert
	Status V1CustomerAlertListResponseDataAlertStatus `json:"status,required"`
	// Threshold value of the alert policy
	Threshold float64 `json:"threshold,required"`
	// Type of the alert
	Type V1CustomerAlertListResponseDataAlertType `json:"type,required"`
	// Timestamp for when the alert was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An array of strings, representing a way to filter the credit grant this alert
	// applies to, by looking at the credit_grant_type field on the credit grant. This
	// field is only defined for CreditPercentage and CreditBalance alerts
	CreditGrantTypeFilters []string                                       `json:"credit_grant_type_filters"`
	CreditType             V1CustomerAlertListResponseDataAlertCreditType `json:"credit_type,nullable"`
	// A list of custom field filters for alert types that support advanced filtering
	CustomFieldFilters []V1CustomerAlertListResponseDataAlertCustomFieldFilter `json:"custom_field_filters"`
	// Scopes alert evaluation to a specific presentation group key on individual line
	// items. Only present for spend alerts.
	GroupKeyFilter V1CustomerAlertListResponseDataAlertGroupKeyFilter `json:"group_key_filter"`
	// Only present for `spend_threshold_reached` alerts. Scope alert to a specific
	// group key on individual line items.
	GroupValues []V1CustomerAlertListResponseDataAlertGroupValue `json:"group_values"`
	// Only supported for invoice_total_reached alerts. A list of invoice types to
	// evaluate.
	InvoiceTypesFilter []string `json:"invoice_types_filter"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                                   `json:"uniqueness_key"`
	JSON          v1CustomerAlertListResponseDataAlertJSON `json:"-"`
}

// v1CustomerAlertListResponseDataAlertJSON contains the JSON metadata for the
// struct [V1CustomerAlertListResponseDataAlert]
type v1CustomerAlertListResponseDataAlertJSON struct {
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
	GroupValues            apijson.Field
	InvoiceTypesFilter     apijson.Field
	UniquenessKey          apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1CustomerAlertListResponseDataAlert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertListResponseDataAlertJSON) RawJSON() string {
	return r.raw
}

// Status of the alert
type V1CustomerAlertListResponseDataAlertStatus string

const (
	V1CustomerAlertListResponseDataAlertStatusEnabled  V1CustomerAlertListResponseDataAlertStatus = "enabled"
	V1CustomerAlertListResponseDataAlertStatusArchived V1CustomerAlertListResponseDataAlertStatus = "archived"
	V1CustomerAlertListResponseDataAlertStatusDisabled V1CustomerAlertListResponseDataAlertStatus = "disabled"
)

func (r V1CustomerAlertListResponseDataAlertStatus) IsKnown() bool {
	switch r {
	case V1CustomerAlertListResponseDataAlertStatusEnabled, V1CustomerAlertListResponseDataAlertStatusArchived, V1CustomerAlertListResponseDataAlertStatusDisabled:
		return true
	}
	return false
}

// Type of the alert
type V1CustomerAlertListResponseDataAlertType string

const (
	V1CustomerAlertListResponseDataAlertTypeLowCreditBalanceReached                           V1CustomerAlertListResponseDataAlertType = "low_credit_balance_reached"
	V1CustomerAlertListResponseDataAlertTypeSpendThresholdReached                             V1CustomerAlertListResponseDataAlertType = "spend_threshold_reached"
	V1CustomerAlertListResponseDataAlertTypeMonthlyInvoiceTotalSpendThresholdReached          V1CustomerAlertListResponseDataAlertType = "monthly_invoice_total_spend_threshold_reached"
	V1CustomerAlertListResponseDataAlertTypeLowRemainingDaysInPlanReached                     V1CustomerAlertListResponseDataAlertType = "low_remaining_days_in_plan_reached"
	V1CustomerAlertListResponseDataAlertTypeLowRemainingCreditPercentageReached               V1CustomerAlertListResponseDataAlertType = "low_remaining_credit_percentage_reached"
	V1CustomerAlertListResponseDataAlertTypeUsageThresholdReached                             V1CustomerAlertListResponseDataAlertType = "usage_threshold_reached"
	V1CustomerAlertListResponseDataAlertTypeLowRemainingDaysForCommitSegmentReached           V1CustomerAlertListResponseDataAlertType = "low_remaining_days_for_commit_segment_reached"
	V1CustomerAlertListResponseDataAlertTypeLowRemainingCommitBalanceReached                  V1CustomerAlertListResponseDataAlertType = "low_remaining_commit_balance_reached"
	V1CustomerAlertListResponseDataAlertTypeLowRemainingCommitPercentageReached               V1CustomerAlertListResponseDataAlertType = "low_remaining_commit_percentage_reached"
	V1CustomerAlertListResponseDataAlertTypeLowRemainingDaysForContractCreditSegmentReached   V1CustomerAlertListResponseDataAlertType = "low_remaining_days_for_contract_credit_segment_reached"
	V1CustomerAlertListResponseDataAlertTypeLowRemainingContractCreditBalanceReached          V1CustomerAlertListResponseDataAlertType = "low_remaining_contract_credit_balance_reached"
	V1CustomerAlertListResponseDataAlertTypeLowRemainingContractCreditPercentageReached       V1CustomerAlertListResponseDataAlertType = "low_remaining_contract_credit_percentage_reached"
	V1CustomerAlertListResponseDataAlertTypeLowRemainingContractCreditAndCommitBalanceReached V1CustomerAlertListResponseDataAlertType = "low_remaining_contract_credit_and_commit_balance_reached"
	V1CustomerAlertListResponseDataAlertTypeInvoiceTotalReached                               V1CustomerAlertListResponseDataAlertType = "invoice_total_reached"
)

func (r V1CustomerAlertListResponseDataAlertType) IsKnown() bool {
	switch r {
	case V1CustomerAlertListResponseDataAlertTypeLowCreditBalanceReached, V1CustomerAlertListResponseDataAlertTypeSpendThresholdReached, V1CustomerAlertListResponseDataAlertTypeMonthlyInvoiceTotalSpendThresholdReached, V1CustomerAlertListResponseDataAlertTypeLowRemainingDaysInPlanReached, V1CustomerAlertListResponseDataAlertTypeLowRemainingCreditPercentageReached, V1CustomerAlertListResponseDataAlertTypeUsageThresholdReached, V1CustomerAlertListResponseDataAlertTypeLowRemainingDaysForCommitSegmentReached, V1CustomerAlertListResponseDataAlertTypeLowRemainingCommitBalanceReached, V1CustomerAlertListResponseDataAlertTypeLowRemainingCommitPercentageReached, V1CustomerAlertListResponseDataAlertTypeLowRemainingDaysForContractCreditSegmentReached, V1CustomerAlertListResponseDataAlertTypeLowRemainingContractCreditBalanceReached, V1CustomerAlertListResponseDataAlertTypeLowRemainingContractCreditPercentageReached, V1CustomerAlertListResponseDataAlertTypeLowRemainingContractCreditAndCommitBalanceReached, V1CustomerAlertListResponseDataAlertTypeInvoiceTotalReached:
		return true
	}
	return false
}

type V1CustomerAlertListResponseDataAlertCreditType struct {
	ID   string                                             `json:"id,required" format:"uuid"`
	Name string                                             `json:"name,required"`
	JSON v1CustomerAlertListResponseDataAlertCreditTypeJSON `json:"-"`
}

// v1CustomerAlertListResponseDataAlertCreditTypeJSON contains the JSON metadata
// for the struct [V1CustomerAlertListResponseDataAlertCreditType]
type v1CustomerAlertListResponseDataAlertCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerAlertListResponseDataAlertCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertListResponseDataAlertCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerAlertListResponseDataAlertCustomFieldFilter struct {
	Entity V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntity `json:"entity,required"`
	Key    string                                                       `json:"key,required"`
	Value  string                                                       `json:"value,required"`
	JSON   v1CustomerAlertListResponseDataAlertCustomFieldFilterJSON    `json:"-"`
}

// v1CustomerAlertListResponseDataAlertCustomFieldFilterJSON contains the JSON
// metadata for the struct [V1CustomerAlertListResponseDataAlertCustomFieldFilter]
type v1CustomerAlertListResponseDataAlertCustomFieldFilterJSON struct {
	Entity      apijson.Field
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerAlertListResponseDataAlertCustomFieldFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertListResponseDataAlertCustomFieldFilterJSON) RawJSON() string {
	return r.raw
}

type V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntity string

const (
	V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntityContract       V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntity = "Contract"
	V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntityCommit         V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntity = "Commit"
	V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntityContractCredit V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntity = "ContractCredit"
)

func (r V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntity) IsKnown() bool {
	switch r {
	case V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntityContract, V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntityCommit, V1CustomerAlertListResponseDataAlertCustomFieldFiltersEntityContractCredit:
		return true
	}
	return false
}

// Scopes alert evaluation to a specific presentation group key on individual line
// items. Only present for spend alerts.
type V1CustomerAlertListResponseDataAlertGroupKeyFilter struct {
	Key   string                                                 `json:"key,required"`
	Value string                                                 `json:"value,required"`
	JSON  v1CustomerAlertListResponseDataAlertGroupKeyFilterJSON `json:"-"`
}

// v1CustomerAlertListResponseDataAlertGroupKeyFilterJSON contains the JSON
// metadata for the struct [V1CustomerAlertListResponseDataAlertGroupKeyFilter]
type v1CustomerAlertListResponseDataAlertGroupKeyFilterJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerAlertListResponseDataAlertGroupKeyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertListResponseDataAlertGroupKeyFilterJSON) RawJSON() string {
	return r.raw
}

type V1CustomerAlertListResponseDataAlertGroupValue struct {
	Key   string                                             `json:"key,required"`
	Value string                                             `json:"value,required"`
	JSON  v1CustomerAlertListResponseDataAlertGroupValueJSON `json:"-"`
}

// v1CustomerAlertListResponseDataAlertGroupValueJSON contains the JSON metadata
// for the struct [V1CustomerAlertListResponseDataAlertGroupValue]
type v1CustomerAlertListResponseDataAlertGroupValueJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerAlertListResponseDataAlertGroupValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerAlertListResponseDataAlertGroupValueJSON) RawJSON() string {
	return r.raw
}

// The status of the customer alert. If the alert is archived, null will be
// returned.
type V1CustomerAlertListResponseDataCustomerStatus string

const (
	V1CustomerAlertListResponseDataCustomerStatusOk         V1CustomerAlertListResponseDataCustomerStatus = "ok"
	V1CustomerAlertListResponseDataCustomerStatusInAlarm    V1CustomerAlertListResponseDataCustomerStatus = "in_alarm"
	V1CustomerAlertListResponseDataCustomerStatusEvaluating V1CustomerAlertListResponseDataCustomerStatus = "evaluating"
)

func (r V1CustomerAlertListResponseDataCustomerStatus) IsKnown() bool {
	switch r {
	case V1CustomerAlertListResponseDataCustomerStatusOk, V1CustomerAlertListResponseDataCustomerStatusInAlarm, V1CustomerAlertListResponseDataCustomerStatusEvaluating:
		return true
	}
	return false
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
