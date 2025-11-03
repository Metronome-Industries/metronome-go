// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"

	"github.com/Metronome-Industries/metronome-go/v2/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v2/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/packages/param"
	"github.com/Metronome-Industries/metronome-go/v2/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/v2/shared"
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
func NewV1AlertService(opts ...option.RequestOption) (r V1AlertService) {
	r = V1AlertService{}
	r.Options = opts
	return
}

// Create a new threshold notification to monitor customer spending, balances, and
// billing metrics in real-time. Metronome's notification system provides
// industry-leading speed with immediate evaluation capabilities, enabling you to
// proactively manage customer accounts and prevent revenue leakage.
//
// This endpoint creates configurable threshold notifications that continuously
// monitor various billing thresholds including spend limits, credit balances,
// commitment utilization, and invoice totals. Threshold notifications can be
// configured globally for all customers or targeted to specific customer accounts.
//
// ### Use this endpoint to:
//
//   - Proactively monitor customer spending patterns to prevent unexpected overages
//     or credit exhaustion
//   - Automate notifications when customers approach commitment limits or credit
//     thresholds
//   - Enable real-time intervention for accounts at risk of churn or payment issues
//   - Scale billing operations by automating threshold-based workflows and
//     notifications
//
// ### Key response fields:
//
// A successful response returns a CustomerAlert object containing:
//
//   - The threshold notification configuration with its unique ID and current status
//   - The customer's evaluation status (ok, in_alarm, or evaluating)
//   - Threshold notification metadata including type, threshold values, and update
//     timestamps
//
// ### Usage guidelines:
//
//   - Immediate evaluation: Set `evaluate_on_create` : `true` (default) for instant
//     evaluation against existing customers
//   - Uniqueness constraints: Each threshold notification must have a unique
//     `uniqueness_key` within your organization. Use `release_uniqueness_key` :
//     `true` when archiving to reuse keys
//   - Notification type requirements: Different threshold notification types require
//     specific fields (e.g., `billable_metric_id` for usage notifications,
//     `credit_type_id` for credit-based threshold notifications)
//   - Webhook delivery: Threshold notifications trigger webhook notifications for
//     real-time integration with your systems. Configure webhook endpoints before
//     creating threshold notifications
//   - Performance at scale: Metronome's event-driven architecture processes
//     threshold notification evaluations in real-time as usage events stream in,
//     unlike competitors who rely on periodic polling or batch evaluation cycles
func (r *V1AlertService) New(ctx context.Context, body V1AlertNewParams, opts ...option.RequestOption) (res *V1AlertNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/alerts/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Permanently disable a threshold notification and remove it from active
// monitoring across all customers. Archived threshold notifications stop
// evaluating immediately and can optionally release their uniqueness key for reuse
// in future threshold notification configurations.
//
// ### Use this endpoint to:
//
//   - Decommission threshold notifications that are no longer needed
//   - Clean up test or deprecated threshold notification configurations
//   - Free up uniqueness keys for reuse with new threshold notifications
//   - Stop threshold notification evaluations without losing historical
//     configuration data
//   - Disable outdated monitoring rules during pricing model transitions
//
// ### Key response fields:
//
// - data: Object containing the archived threshold notification's ID
//
// ### Usage guidelines:
//
//   - Irreversible for evaluation: Archived threshold notifications cannot be
//     re-enabled; create a new threshold notification to resume monitoring
//   - Uniqueness key handling: Set `release_uniqueness_key` : `true` to reuse the
//     key in future threshold notifications
//   - Immediate effect: Threshold notification evaluation stops instantly across all
//     customers
//   - Historical preservation: Archive operation maintains threshold notification
//     history and configuration for compliance and auditing
func (r *V1AlertService) Archive(ctx context.Context, body V1AlertArchiveParams, opts ...option.RequestOption) (res *V1AlertArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/alerts/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1AlertNewResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AlertNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AlertNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AlertArchiveResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AlertArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AlertArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AlertNewParams struct {
	// Type of the threshold notification
	//
	// Any of "low_credit_balance_reached", "spend_threshold_reached",
	// "monthly_invoice_total_spend_threshold_reached",
	// "low_remaining_days_in_plan_reached", "low_remaining_credit_percentage_reached",
	// "usage_threshold_reached", "low_remaining_days_for_commit_segment_reached",
	// "low_remaining_commit_balance_reached",
	// "low_remaining_commit_percentage_reached",
	// "low_remaining_days_for_contract_credit_segment_reached",
	// "low_remaining_contract_credit_balance_reached",
	// "low_remaining_contract_credit_percentage_reached",
	// "low_remaining_contract_credit_and_commit_balance_reached",
	// "invoice_total_reached", "low_remaining_seat_balance_reached".
	AlertType V1AlertNewParamsAlertType `json:"alert_type,omitzero,required"`
	// Name of the threshold notification
	Name string `json:"name,required"`
	// Threshold value of the notification policy. Depending upon the notification
	// type, this number may represent a financial amount, the days remaining, or a
	// percentage reached.
	Threshold float64 `json:"threshold,required"`
	// For threshold notifications of type `usage_threshold_reached`, specifies which
	// billable metric to track the usage for.
	BillableMetricID param.Opt[string] `json:"billable_metric_id,omitzero" format:"uuid"`
	// ID of the credit's currency, defaults to USD. If the specific notification type
	// requires a pricing unit/currency, find the ID in the
	// [Metronome app](https://app.metronome.com/offering/pricing-units).
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// If provided, will create this threshold notification for this specific customer.
	// To create a notification for all customers, do not specify a `customer_id`.
	CustomerID param.Opt[string] `json:"customer_id,omitzero" format:"uuid"`
	// If true, the threshold notification will evaluate immediately on customers that
	// already meet the notification threshold. If false, it will only evaluate on
	// future customers that trigger the threshold. Defaults to true.
	EvaluateOnCreate param.Opt[bool] `json:"evaluate_on_create,omitzero"`
	// If provided, will create this threshold notification for this specific plan. To
	// create a notification for all customers, do not specify a `plan_id`.
	PlanID param.Opt[string] `json:"plan_id,omitzero" format:"uuid"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey param.Opt[string] `json:"uniqueness_key,omitzero"`
	// An array of strings, representing a way to filter the credit grant this
	// threshold notification applies to, by looking at the credit_grant_type field on
	// the credit grant. This field is only defined for CreditPercentage and
	// CreditBalance notifications
	CreditGrantTypeFilters []string `json:"credit_grant_type_filters,omitzero"`
	// A list of custom field filters for threshold notification types that support
	// advanced filtering. Only present for contract invoices.
	CustomFieldFilters []V1AlertNewParamsCustomFieldFilter `json:"custom_field_filters,omitzero"`
	// Only present for `spend_threshold_reached` notifications. Scope notification to
	// a specific group key on individual line items.
	GroupValues []V1AlertNewParamsGroupValue `json:"group_values,omitzero"`
	// Only supported for invoice_total_reached threshold notifications. A list of
	// invoice types to evaluate.
	InvoiceTypesFilter []string `json:"invoice_types_filter,omitzero"`
	paramObj
}

func (r V1AlertNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1AlertNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AlertNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the threshold notification
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
	V1AlertNewParamsAlertTypeLowRemainingSeatBalanceReached                    V1AlertNewParamsAlertType = "low_remaining_seat_balance_reached"
)

// The properties Entity, Key, Value are required.
type V1AlertNewParamsCustomFieldFilter struct {
	// Any of "Contract", "Commit", "ContractCredit".
	Entity string `json:"entity,omitzero,required"`
	Key    string `json:"key,required"`
	Value  string `json:"value,required"`
	paramObj
}

func (r V1AlertNewParamsCustomFieldFilter) MarshalJSON() (data []byte, err error) {
	type shadow V1AlertNewParamsCustomFieldFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AlertNewParamsCustomFieldFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AlertNewParamsCustomFieldFilter](
		"entity", "Contract", "Commit", "ContractCredit",
	)
}

// The property Key is required.
type V1AlertNewParamsGroupValue struct {
	Key   string            `json:"key,required"`
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r V1AlertNewParamsGroupValue) MarshalJSON() (data []byte, err error) {
	type shadow V1AlertNewParamsGroupValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AlertNewParamsGroupValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AlertArchiveParams struct {
	// The Metronome ID of the threshold notification
	ID string `json:"id,required" format:"uuid"`
	// If true, resets the uniqueness key on this threshold notification so it can be
	// re-used
	ReleaseUniquenessKey param.Opt[bool] `json:"release_uniqueness_key,omitzero"`
	paramObj
}

func (r V1AlertArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1AlertArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AlertArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
