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

// V1AlertService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AlertService] method instead.
type V1AlertService struct {
	Options []option.RequestOption
}

// NewV1AlertService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1AlertService(opts ...option.RequestOption) (r *V1AlertService) {
	r = &V1AlertService{}
	r.Options = opts
	return
}

// Create a new alert to monitor customer spending, balances, and billing metrics
// in real-time. Metronome's alert system provides industry-leading speed with
// immediate evaluation capabilities, enabling you to proactively manage customer
// accounts and prevent revenue leakage.
//
// This endpoint creates configurable alerts that continuously monitor various
// billing thresholds including spend limits, credit balances, commitment
// utilization, and invoice totals. Alerts can be configured globally for all
// customers or targeted to specific customer accounts. Custom fields can be used
// on certain alert types to further target alerts to groups of customers.
//
// Use this endpoint to:
//
//   - Proactively monitor customer spending patterns to prevent unexpected overages
//     or credit exhaustion
//   - Automate notifications when customers approach commitment limits or credit
//     thresholds
//   - Enable real-time intervention for accounts at risk of churn or payment issues
//   - Scale billing operations by automating threshold-based workflows and
//     notifications
//
// Key response fields: A successful response returns a CustomerAlert object
// containing:
//
// - The alert configuration with its unique ID and current status
// - The customer's evaluation status (ok, in_alarm, or evaluating)
// - Alert metadata including type, threshold values, and update timestamps
//
// Usage guidelines:
//
//   - Immediate evaluation: Set evaluate_on_create: true (default) for instant
//     evaluation against existing customers
//   - Uniqueness constraints: Each alert must have a unique uniqueness_key within
//     your organization. Use release_uniqueness_key: true when archiving to reuse
//     keys
//   - Alert type requirements: Different alert types require specific fields (e.g.,
//     billable_metric_id for usage alerts, credit_type_id for credit-based alerts)
//   - Webhook delivery: Alerts trigger webhook notifications for real-time
//     integration with your systems. Configure webhook endpoints before creating
//     alerts
//   - Performance at scale: Metronome's event-driven architecture processes alert
//     evaluations in real-time as usage events stream in, unlike competitors who
//     rely on periodic polling or batch evaluation cycles
func (r *V1AlertService) New(ctx context.Context, body V1AlertNewParams, opts ...option.RequestOption) (res *V1AlertNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/alerts/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Permanently disable an alert and remove it from active monitoring across all
// customers. Archived alerts stop evaluating immediately and can optionally
// release their uniqueness key for reuse in future alert configurations.
//
// Use this endpoint to:
//
// - Decommission alerts that are no longer needed
// - Clean up test or deprecated alert configurations
// - Free up uniqueness keys for reuse with new alerts
// - Stop alert evaluations without losing historical configuration data
// - Disable outdated monitoring rules during pricing model transitions
//
// Key response fields:
//
// - data: Object containing the archived alert's ID
// - Alert evaluation stops immediately for all affected customers
// - Historical alert data and configurations remain accessible for audit purposes
//
// Usage guidelines:
//
//   - Irreversible for evaluation: Archived alerts cannot be re-enabled; create a
//     new alert to resume monitoring
//   - Uniqueness key handling: Set release_uniqueness_key: true to reuse the key in
//     future alerts
//   - Immediate effect: Alert evaluation stops instantly across all customers
//   - Historical preservation: Archive operation maintains alert history and
//     configuration for compliance and auditing
func (r *V1AlertService) Archive(ctx context.Context, body V1AlertArchiveParams, opts ...option.RequestOption) (res *V1AlertArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/alerts/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1AlertNewResponse struct {
	Data shared.ID              `json:"data,required"`
	JSON v1AlertNewResponseJSON `json:"-"`
}

// v1AlertNewResponseJSON contains the JSON metadata for the struct
// [V1AlertNewResponse]
type v1AlertNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1AlertNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1AlertNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1AlertArchiveResponse struct {
	Data shared.ID                  `json:"data,required"`
	JSON v1AlertArchiveResponseJSON `json:"-"`
}

// v1AlertArchiveResponseJSON contains the JSON metadata for the struct
// [V1AlertArchiveResponse]
type v1AlertArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1AlertArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1AlertArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type V1AlertNewParams struct {
	// Type of the alert
	AlertType param.Field[V1AlertNewParamsAlertType] `json:"alert_type,required"`
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
	// ID of the credit's currency, defaults to USD. If the specific alert type
	// requires a pricing unit/currency, find the ID in the
	// [Metronome app](https://app.metronome.com/offering/pricing-units).
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// A list of custom field filters for alert types that support advanced filtering.
	// Only present for contract invoices.
	CustomFieldFilters param.Field[[]V1AlertNewParamsCustomFieldFilter] `json:"custom_field_filters"`
	// If provided, will create this alert for this specific customer. To create an
	// alert for all customers, do not specify a `customer_id`.
	CustomerID param.Field[string] `json:"customer_id" format:"uuid"`
	// If true, the alert will evaluate immediately on customers that already meet the
	// alert threshold. If false, it will only evaluate on future customers that
	// trigger the alert threshold. Defaults to true.
	EvaluateOnCreate param.Field[bool] `json:"evaluate_on_create"`
	// Only present for `spend_threshold_reached` alerts. Scope alert to a specific
	// group key on individual line items.
	GroupValues param.Field[[]V1AlertNewParamsGroupValue] `json:"group_values"`
	// Only supported for invoice_total_reached alerts. A list of invoice types to
	// evaluate.
	InvoiceTypesFilter param.Field[[]string] `json:"invoice_types_filter"`
	// If provided, will create this alert for this specific plan. To create an alert
	// for all customers, do not specify a `plan_id`.
	PlanID param.Field[string] `json:"plan_id" format:"uuid"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey param.Field[string] `json:"uniqueness_key"`
}

func (r V1AlertNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the alert
type V1AlertNewParamsAlertType string

const (
	V1AlertNewParamsAlertTypeLowCreditBalanceReached                           V1AlertNewParamsAlertType = "low_credit_balance_reached"
	V1AlertNewParamsAlertTypeSpendThresholdReached                             V1AlertNewParamsAlertType = "spend_threshold_reached"
	V1AlertNewParamsAlertTypeMonthlyInvoiceTotalSpendThresholdReached          V1AlertNewParamsAlertType = "monthly_invoice_total_spend_threshold_reached"
	V1AlertNewParamsAlertTypeLowRemainingDaysInPlanReached                     V1AlertNewParamsAlertType = "low_remaining_days_in_plan_reached"
	V1AlertNewParamsAlertTypeLowRemainingCreditPercentageReached               V1AlertNewParamsAlertType = "low_remaining_credit_percentage_reached"
	V1AlertNewParamsAlertTypeUsageThresholdReached                             V1AlertNewParamsAlertType = "usage_threshold_reached"
	V1AlertNewParamsAlertTypeLowRemainingDaysForCommitSegmentReached           V1AlertNewParamsAlertType = "low_remaining_days_for_commit_segment_reached"
	V1AlertNewParamsAlertTypeLowRemainingCommitBalanceReached                  V1AlertNewParamsAlertType = "low_remaining_commit_balance_reached"
	V1AlertNewParamsAlertTypeLowRemainingCommitPercentageReached               V1AlertNewParamsAlertType = "low_remaining_commit_percentage_reached"
	V1AlertNewParamsAlertTypeLowRemainingDaysForContractCreditSegmentReached   V1AlertNewParamsAlertType = "low_remaining_days_for_contract_credit_segment_reached"
	V1AlertNewParamsAlertTypeLowRemainingContractCreditBalanceReached          V1AlertNewParamsAlertType = "low_remaining_contract_credit_balance_reached"
	V1AlertNewParamsAlertTypeLowRemainingContractCreditPercentageReached       V1AlertNewParamsAlertType = "low_remaining_contract_credit_percentage_reached"
	V1AlertNewParamsAlertTypeLowRemainingContractCreditAndCommitBalanceReached V1AlertNewParamsAlertType = "low_remaining_contract_credit_and_commit_balance_reached"
	V1AlertNewParamsAlertTypeInvoiceTotalReached                               V1AlertNewParamsAlertType = "invoice_total_reached"
)

func (r V1AlertNewParamsAlertType) IsKnown() bool {
	switch r {
	case V1AlertNewParamsAlertTypeLowCreditBalanceReached, V1AlertNewParamsAlertTypeSpendThresholdReached, V1AlertNewParamsAlertTypeMonthlyInvoiceTotalSpendThresholdReached, V1AlertNewParamsAlertTypeLowRemainingDaysInPlanReached, V1AlertNewParamsAlertTypeLowRemainingCreditPercentageReached, V1AlertNewParamsAlertTypeUsageThresholdReached, V1AlertNewParamsAlertTypeLowRemainingDaysForCommitSegmentReached, V1AlertNewParamsAlertTypeLowRemainingCommitBalanceReached, V1AlertNewParamsAlertTypeLowRemainingCommitPercentageReached, V1AlertNewParamsAlertTypeLowRemainingDaysForContractCreditSegmentReached, V1AlertNewParamsAlertTypeLowRemainingContractCreditBalanceReached, V1AlertNewParamsAlertTypeLowRemainingContractCreditPercentageReached, V1AlertNewParamsAlertTypeLowRemainingContractCreditAndCommitBalanceReached, V1AlertNewParamsAlertTypeInvoiceTotalReached:
		return true
	}
	return false
}

type V1AlertNewParamsCustomFieldFilter struct {
	Entity param.Field[V1AlertNewParamsCustomFieldFiltersEntity] `json:"entity,required"`
	Key    param.Field[string]                                   `json:"key,required"`
	Value  param.Field[string]                                   `json:"value,required"`
}

func (r V1AlertNewParamsCustomFieldFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1AlertNewParamsCustomFieldFiltersEntity string

const (
	V1AlertNewParamsCustomFieldFiltersEntityContract       V1AlertNewParamsCustomFieldFiltersEntity = "Contract"
	V1AlertNewParamsCustomFieldFiltersEntityCommit         V1AlertNewParamsCustomFieldFiltersEntity = "Commit"
	V1AlertNewParamsCustomFieldFiltersEntityContractCredit V1AlertNewParamsCustomFieldFiltersEntity = "ContractCredit"
)

func (r V1AlertNewParamsCustomFieldFiltersEntity) IsKnown() bool {
	switch r {
	case V1AlertNewParamsCustomFieldFiltersEntityContract, V1AlertNewParamsCustomFieldFiltersEntityCommit, V1AlertNewParamsCustomFieldFiltersEntityContractCredit:
		return true
	}
	return false
}

type V1AlertNewParamsGroupValue struct {
	Key   param.Field[string] `json:"key,required"`
	Value param.Field[string] `json:"value"`
}

func (r V1AlertNewParamsGroupValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1AlertArchiveParams struct {
	// The Metronome ID of the alert
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// If true, resets the uniqueness key on this alert so it can be re-used
	ReleaseUniquenessKey param.Field[bool] `json:"release_uniqueness_key"`
}

func (r V1AlertArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
