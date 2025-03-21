// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// AlertService contains methods and other services that help with interacting with
// the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlertService] method instead.
type AlertService struct {
	Options []option.RequestOption
}

// NewAlertService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAlertService(opts ...option.RequestOption) (r *AlertService) {
	r = &AlertService{}
	r.Options = opts
	return
}

// Create a new alert
func (r *AlertService) New(ctx context.Context, body AlertNewParams, opts ...option.RequestOption) (res *AlertNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/alerts/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Archive an existing alert
func (r *AlertService) Archive(ctx context.Context, body AlertArchiveParams, opts ...option.RequestOption) (res *AlertArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/alerts/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AlertNewResponse struct {
	Data shared.ID            `json:"data,required"`
	JSON alertNewResponseJSON `json:"-"`
}

// alertNewResponseJSON contains the JSON metadata for the struct
// [AlertNewResponse]
type alertNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertNewResponseJSON) RawJSON() string {
	return r.raw
}

type AlertArchiveResponse struct {
	Data shared.ID                `json:"data,required"`
	JSON alertArchiveResponseJSON `json:"-"`
}

// alertArchiveResponseJSON contains the JSON metadata for the struct
// [AlertArchiveResponse]
type alertArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type AlertNewParams struct {
	// Type of the alert
	AlertType param.Field[AlertNewParamsAlertType] `json:"alert_type,required"`
	// Name of the alert
	Name param.Field[string] `json:"name,required"`
	// Threshold value of the alert policy. Depending upon the alert type, this number
	// may represent a financial amount, the days remaining, or a percentage reached.
	Threshold param.Field[float64] `json:"threshold,required"`
	// For alerts of type `usage_threshold_reached`, specifies which billable metric to
	// track the usage for.
	BillableMetricID param.Field[string] `json:"billable_metric_id" format:"uuid"`
	// An array of strings, representing a way to filter the credit grant this alert
	// applies to, by looking at the credit_grant_type field on the credit grant. This
	// field is only defined for CreditPercentage and CreditBalance alerts
	CreditGrantTypeFilters param.Field[[]string] `json:"credit_grant_type_filters"`
	CreditTypeID           param.Field[string]   `json:"credit_type_id" format:"uuid"`
	// A list of custom field filters for alert types that support advanced filtering.
	// Only present for contract invoices.
	CustomFieldFilters param.Field[[]AlertNewParamsCustomFieldFilter] `json:"custom_field_filters"`
	// If provided, will create this alert for this specific customer. To create an
	// alert for all customers, do not specify `customer_id` or `plan_id`.
	CustomerID param.Field[string] `json:"customer_id" format:"uuid"`
	// If true, the alert will evaluate immediately on customers that already meet the
	// alert threshold. If false, it will only evaluate on future customers that
	// trigger the alert threshold. Defaults to true.
	EvaluateOnCreate param.Field[bool] `json:"evaluate_on_create"`
	// Scopes alert evaluation to a specific presentation group key on individual line
	// items. Only present for spend alerts.
	GroupKeyFilter param.Field[AlertNewParamsGroupKeyFilter] `json:"group_key_filter"`
	// Only supported for invoice_total_reached alerts. A list of invoice types to
	// evaluate.
	InvoiceTypesFilter param.Field[[]string] `json:"invoice_types_filter"`
	// If provided, will create this alert for this specific plan. To create an alert
	// for all customers, do not specify `customer_id` or `plan_id`.
	PlanID param.Field[string] `json:"plan_id" format:"uuid"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey param.Field[string] `json:"uniqueness_key"`
}

func (r AlertNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the alert
type AlertNewParamsAlertType string

const (
	AlertNewParamsAlertTypeLowCreditBalanceReached                           AlertNewParamsAlertType = "low_credit_balance_reached"
	AlertNewParamsAlertTypeSpendThresholdReached                             AlertNewParamsAlertType = "spend_threshold_reached"
	AlertNewParamsAlertTypeMonthlyInvoiceTotalSpendThresholdReached          AlertNewParamsAlertType = "monthly_invoice_total_spend_threshold_reached"
	AlertNewParamsAlertTypeLowRemainingDaysInPlanReached                     AlertNewParamsAlertType = "low_remaining_days_in_plan_reached"
	AlertNewParamsAlertTypeLowRemainingCreditPercentageReached               AlertNewParamsAlertType = "low_remaining_credit_percentage_reached"
	AlertNewParamsAlertTypeUsageThresholdReached                             AlertNewParamsAlertType = "usage_threshold_reached"
	AlertNewParamsAlertTypeLowRemainingDaysForCommitSegmentReached           AlertNewParamsAlertType = "low_remaining_days_for_commit_segment_reached"
	AlertNewParamsAlertTypeLowRemainingCommitBalanceReached                  AlertNewParamsAlertType = "low_remaining_commit_balance_reached"
	AlertNewParamsAlertTypeLowRemainingCommitPercentageReached               AlertNewParamsAlertType = "low_remaining_commit_percentage_reached"
	AlertNewParamsAlertTypeLowRemainingDaysForContractCreditSegmentReached   AlertNewParamsAlertType = "low_remaining_days_for_contract_credit_segment_reached"
	AlertNewParamsAlertTypeLowRemainingContractCreditBalanceReached          AlertNewParamsAlertType = "low_remaining_contract_credit_balance_reached"
	AlertNewParamsAlertTypeLowRemainingContractCreditPercentageReached       AlertNewParamsAlertType = "low_remaining_contract_credit_percentage_reached"
	AlertNewParamsAlertTypeLowRemainingContractCreditAndCommitBalanceReached AlertNewParamsAlertType = "low_remaining_contract_credit_and_commit_balance_reached"
	AlertNewParamsAlertTypeInvoiceTotalReached                               AlertNewParamsAlertType = "invoice_total_reached"
)

func (r AlertNewParamsAlertType) IsKnown() bool {
	switch r {
	case AlertNewParamsAlertTypeLowCreditBalanceReached, AlertNewParamsAlertTypeSpendThresholdReached, AlertNewParamsAlertTypeMonthlyInvoiceTotalSpendThresholdReached, AlertNewParamsAlertTypeLowRemainingDaysInPlanReached, AlertNewParamsAlertTypeLowRemainingCreditPercentageReached, AlertNewParamsAlertTypeUsageThresholdReached, AlertNewParamsAlertTypeLowRemainingDaysForCommitSegmentReached, AlertNewParamsAlertTypeLowRemainingCommitBalanceReached, AlertNewParamsAlertTypeLowRemainingCommitPercentageReached, AlertNewParamsAlertTypeLowRemainingDaysForContractCreditSegmentReached, AlertNewParamsAlertTypeLowRemainingContractCreditBalanceReached, AlertNewParamsAlertTypeLowRemainingContractCreditPercentageReached, AlertNewParamsAlertTypeLowRemainingContractCreditAndCommitBalanceReached, AlertNewParamsAlertTypeInvoiceTotalReached:
		return true
	}
	return false
}

type AlertNewParamsCustomFieldFilter struct {
	Entity param.Field[AlertNewParamsCustomFieldFiltersEntity] `json:"entity,required"`
	Key    param.Field[string]                                 `json:"key,required"`
	Value  param.Field[string]                                 `json:"value,required"`
}

func (r AlertNewParamsCustomFieldFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertNewParamsCustomFieldFiltersEntity string

const (
	AlertNewParamsCustomFieldFiltersEntityContract       AlertNewParamsCustomFieldFiltersEntity = "Contract"
	AlertNewParamsCustomFieldFiltersEntityCommit         AlertNewParamsCustomFieldFiltersEntity = "Commit"
	AlertNewParamsCustomFieldFiltersEntityContractCredit AlertNewParamsCustomFieldFiltersEntity = "ContractCredit"
)

func (r AlertNewParamsCustomFieldFiltersEntity) IsKnown() bool {
	switch r {
	case AlertNewParamsCustomFieldFiltersEntityContract, AlertNewParamsCustomFieldFiltersEntityCommit, AlertNewParamsCustomFieldFiltersEntityContractCredit:
		return true
	}
	return false
}

// Scopes alert evaluation to a specific presentation group key on individual line
// items. Only present for spend alerts.
type AlertNewParamsGroupKeyFilter struct {
	Key   param.Field[string] `json:"key,required"`
	Value param.Field[string] `json:"value,required"`
}

func (r AlertNewParamsGroupKeyFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertArchiveParams struct {
	// The Metronome ID of the alert
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// If true, resets the uniqueness key on this alert so it can be re-used
	ReleaseUniquenessKey param.Field[bool] `json:"release_uniqueness_key"`
}

func (r AlertArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
