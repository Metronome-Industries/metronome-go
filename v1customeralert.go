// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v2/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v2/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/v2/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/v2/packages/param"
	"github.com/Metronome-Industries/metronome-go/v2/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/v2/shared"
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
func NewV1CustomerAlertService(opts ...option.RequestOption) (r V1CustomerAlertService) {
	r = V1CustomerAlertService{}
	r.Options = opts
	return
}

// Retrieve the real-time evaluation status for a specific threshold
// notification-customer pair. This endpoint provides instant visibility into
// whether a customer has triggered a threshold notification condition, enabling
// you to monitor account health and take proactive action based on current
// threshold notification states.
//
// ### Use this endpoint to:
//
//   - Check if a specific customer is currently violating an threshold notification
//     (`in_alarm` status)
//   - Verify threshold notification configuration details and threshold values for a
//     customer
//   - Monitor the evaluation state of newly created or recently modified threshold
//     notification
//   - Build dashboards or automated workflows that respond to specific threshold
//     notification conditions
//   - Validate threshold notification behavior before deploying to production
//     customers
//   - Integrate threshold notification status checks into customer support tools or
//     admin interfaces
//
// ### Key response fields:
//
// A CustomerAlert object containing:
//
// - `customer_status`: The current evaluation state
//
//   - `ok` - Customer is within acceptable thresholds
//   - `in_alarm` - Customer has breached the threshold for the notification
//   - `evaluating` - Notification is currently being evaluated (typically during
//     initial setup)
//   - `null` - Notification has been archived
//   - `triggered_by`: Additional context about what caused the notification to
//     trigger (when applicable)
//   - alert: Complete threshold notification configuration including:
//   - Notification ID, name, and type
//   - Current threshold values and credit type information
//   - Notification status (enabled, disabled, or archived)
//   - Last update timestamp
//   - Any applied filters (credit grant types, custom fields, group values)
//
// ### Usage guidelines:
//
//   - Customer status: Returns the current evaluation state, not historical data.
//     For threshold notification history, use webhook notifications or event logs
//   - Required parameters: Both customer_id and alert_id must be valid UUIDs that
//     exist in your organization
//   - Archived notifications: Returns null for customer_status if the notification
//     has been archived, but still includes the notification configuration details
//   - Performance considerations: This endpoint queries live evaluation state,
//     making it ideal for real-time monitoring but not for bulk status checks
//   - Integration patterns: Best used for on-demand status checks in response to
//     user actions or as part of targeted monitoring workflows
//   - Error handling: Returns 404 if either the customer or alert_id doesn't exist
//     or isn't accessible to your organization
func (r *V1CustomerAlertService) Get(ctx context.Context, body V1CustomerAlertGetParams, opts ...option.RequestOption) (res *V1CustomerAlertGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/customer-alerts/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve all threshold notification configurations and their current statuses
// for a specific customer in a single API call. This endpoint provides a
// comprehensive view of all threshold notification monitoring a customer account.
//
// ### Use this endpoint to:
//
//   - Display all active threshold notifications for a customer in dashboards or
//     admin panels
//   - Quickly identify which threshold notifications a customer is currently
//     triggering
//   - Audit threshold notification coverage for specific accounts
//   - Filter threshold notifications by status (enabled, disabled, or archived)
//
// ### Key response fields:
//
// - data: Array of CustomerAlert objects, each containing:
//   - Current evaluation status (`ok`, `in_alarm`, `evaluating`, or `null`)
//   - Complete threshold notification configuration and threshold details
//   - Threshold notification metadata including type, name, and last update time
//
// - next_page: Pagination cursor for retrieving additional results
//
// ### Usage guidelines:
//
//   - Default behavior: Returns only enabled threshold notifications unless
//     `alert_statuses` filter is specified
//   - Pagination: Use the `next_page` cursor to retrieve all results for customers
//     with many notifications
//   - Performance: Efficiently retrieves multiple threshold notification statuses in
//     a single request instead of making individual calls
//   - Filtering: Pass the `alert_statuses` array to include disabled or archived
//     threshold notifications in results
func (r *V1CustomerAlertService) List(ctx context.Context, params V1CustomerAlertListParams, opts ...option.RequestOption) (res *pagination.CursorPageWithoutLimit[CustomerAlert], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/customer-alerts/list"
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

// Retrieve all threshold notification configurations and their current statuses
// for a specific customer in a single API call. This endpoint provides a
// comprehensive view of all threshold notification monitoring a customer account.
//
// ### Use this endpoint to:
//
//   - Display all active threshold notifications for a customer in dashboards or
//     admin panels
//   - Quickly identify which threshold notifications a customer is currently
//     triggering
//   - Audit threshold notification coverage for specific accounts
//   - Filter threshold notifications by status (enabled, disabled, or archived)
//
// ### Key response fields:
//
// - data: Array of CustomerAlert objects, each containing:
//   - Current evaluation status (`ok`, `in_alarm`, `evaluating`, or `null`)
//   - Complete threshold notification configuration and threshold details
//   - Threshold notification metadata including type, name, and last update time
//
// - next_page: Pagination cursor for retrieving additional results
//
// ### Usage guidelines:
//
//   - Default behavior: Returns only enabled threshold notifications unless
//     `alert_statuses` filter is specified
//   - Pagination: Use the `next_page` cursor to retrieve all results for customers
//     with many notifications
//   - Performance: Efficiently retrieves multiple threshold notification statuses in
//     a single request instead of making individual calls
//   - Filtering: Pass the `alert_statuses` array to include disabled or archived
//     threshold notifications in results
func (r *V1CustomerAlertService) ListAutoPaging(ctx context.Context, params V1CustomerAlertListParams, opts ...option.RequestOption) *pagination.CursorPageWithoutLimitAutoPager[CustomerAlert] {
	return pagination.NewCursorPageWithoutLimitAutoPager(r.List(ctx, params, opts...))
}

// Force an immediate re-evaluation of a specific threshold notification for a
// customer, clearing any previous state and triggering a fresh assessment against
// current thresholds. This endpoint ensures threshold notification accuracy after
// configuration changes or data corrections.
//
// ### Use this endpoint to:
//
//   - Clear false positive threshold notifications after fixing data issues
//   - Re-evaluate threshold notifications after adjusting customer balances or
//     credits
//   - Test threshold notification behavior during development and debugging
//   - Resolve stuck threshold notification that may be in an incorrect state
//   - Trigger immediate evaluation after threshold modifications
//
// ### Key response fields:
//
//   - 200 Success: Confirmation that the threshold notification has been reset and
//     re-evaluation initiated
//   - No response body is returned - the operation completes asynchronously
//
// ### Usage guidelines:
//
//   - Immediate effect: Triggers re-evaluation instantly, which may result in new
//     webhook notifications if thresholds are breached
//   - State clearing: Removes any cached evaluation state, ensuring a fresh
//     assessment
//   - Use sparingly: Intended for exceptional cases, not routine operations
//   - Asynchronous processing: The reset completes immediately, but re-evaluation
//     happens in the background
func (r *V1CustomerAlertService) Reset(ctx context.Context, body V1CustomerAlertResetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/customer-alerts/reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CustomerAlert struct {
	Alert CustomerAlertAlert `json:"alert,required"`
	// The status of the threshold notification. If the notification is archived, null
	// will be returned.
	//
	// Any of "ok", "in_alarm", "evaluating".
	CustomerStatus CustomerAlertCustomerStatus `json:"customer_status,required"`
	// If present, indicates the reason the threshold notification was triggered.
	TriggeredBy string `json:"triggered_by,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alert          respjson.Field
		CustomerStatus respjson.Field
		TriggeredBy    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerAlert) RawJSON() string { return r.JSON.raw }
func (r *CustomerAlert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAlertAlert struct {
	// the Metronome ID of the threshold notification
	ID string `json:"id,required"`
	// Name of the threshold notification
	Name string `json:"name,required"`
	// Status of the threshold notification
	//
	// Any of "enabled", "archived", "disabled".
	Status string `json:"status,required"`
	// Threshold value of the notification policy
	Threshold float64 `json:"threshold,required"`
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
	// "low_remaining_seat_balance_reached", "invoice_total_reached".
	Type string `json:"type,required"`
	// Timestamp for when the threshold notification was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An array of strings, representing a way to filter the credit grant this
	// threshold notification applies to, by looking at the credit_grant_type field on
	// the credit grant. This field is only defined for CreditPercentage and
	// CreditBalance notifications
	CreditGrantTypeFilters []string              `json:"credit_grant_type_filters"`
	CreditType             shared.CreditTypeData `json:"credit_type,nullable"`
	// A list of custom field filters for notification types that support advanced
	// filtering
	CustomFieldFilters []CustomerAlertAlertCustomFieldFilter `json:"custom_field_filters"`
	// Scopes threshold notification evaluation to a specific presentation group key on
	// individual line items. Only present for spend notifications.
	GroupKeyFilter CustomerAlertAlertGroupKeyFilter `json:"group_key_filter"`
	// Only present for `spend_threshold_reached` notifications. Scope notification to
	// a specific group key on individual line items.
	GroupValues []CustomerAlertAlertGroupValue `json:"group_values"`
	// Only supported for invoice_total_reached threshold notifications. A list of
	// invoice types to evaluate.
	InvoiceTypesFilter []string `json:"invoice_types_filter"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string `json:"uniqueness_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Name                   respjson.Field
		Status                 respjson.Field
		Threshold              respjson.Field
		Type                   respjson.Field
		UpdatedAt              respjson.Field
		CreditGrantTypeFilters respjson.Field
		CreditType             respjson.Field
		CustomFieldFilters     respjson.Field
		GroupKeyFilter         respjson.Field
		GroupValues            respjson.Field
		InvoiceTypesFilter     respjson.Field
		UniquenessKey          respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerAlertAlert) RawJSON() string { return r.JSON.raw }
func (r *CustomerAlertAlert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAlertAlertCustomFieldFilter struct {
	// Any of "Contract", "Commit", "ContractCredit".
	Entity string `json:"entity,required"`
	Key    string `json:"key,required"`
	Value  string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entity      respjson.Field
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerAlertAlertCustomFieldFilter) RawJSON() string { return r.JSON.raw }
func (r *CustomerAlertAlertCustomFieldFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scopes threshold notification evaluation to a specific presentation group key on
// individual line items. Only present for spend notifications.
type CustomerAlertAlertGroupKeyFilter struct {
	Key   string `json:"key,required"`
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerAlertAlertGroupKeyFilter) RawJSON() string { return r.JSON.raw }
func (r *CustomerAlertAlertGroupKeyFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAlertAlertGroupValue struct {
	Key   string `json:"key,required"`
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerAlertAlertGroupValue) RawJSON() string { return r.JSON.raw }
func (r *CustomerAlertAlertGroupValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the threshold notification. If the notification is archived, null
// will be returned.
type CustomerAlertCustomerStatus string

const (
	CustomerAlertCustomerStatusOk         CustomerAlertCustomerStatus = "ok"
	CustomerAlertCustomerStatusInAlarm    CustomerAlertCustomerStatus = "in_alarm"
	CustomerAlertCustomerStatusEvaluating CustomerAlertCustomerStatus = "evaluating"
)

type V1CustomerAlertGetResponse struct {
	Data CustomerAlert `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerAlertGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerAlertGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerAlertGetParams struct {
	// The Metronome ID of the threshold notification
	AlertID string `json:"alert_id,required" format:"uuid"`
	// The Metronome ID of the customer
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// Only present for `spend_threshold_reached` notifications. Retrieve the
	// notification for a specific group key-value pair.
	GroupValues []V1CustomerAlertGetParamsGroupValue `json:"group_values,omitzero"`
	// When parallel threshold notifications are enabled during migration, this flag
	// denotes whether to fetch notifications for plans or contracts.
	//
	// Any of "PLANS", "CONTRACTS".
	PlansOrContracts V1CustomerAlertGetParamsPlansOrContracts `json:"plans_or_contracts,omitzero"`
	paramObj
}

func (r V1CustomerAlertGetParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerAlertGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerAlertGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scopes threshold notification evaluation to a specific presentation group key on
// individual line items. Only present for spend notifications.
//
// The properties Key, Value are required.
type V1CustomerAlertGetParamsGroupValue struct {
	Key   string `json:"key,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r V1CustomerAlertGetParamsGroupValue) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerAlertGetParamsGroupValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerAlertGetParamsGroupValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When parallel threshold notifications are enabled during migration, this flag
// denotes whether to fetch notifications for plans or contracts.
type V1CustomerAlertGetParamsPlansOrContracts string

const (
	V1CustomerAlertGetParamsPlansOrContractsPlans     V1CustomerAlertGetParamsPlansOrContracts = "PLANS"
	V1CustomerAlertGetParamsPlansOrContractsContracts V1CustomerAlertGetParamsPlansOrContracts = "CONTRACTS"
)

type V1CustomerAlertListParams struct {
	// The Metronome ID of the customer
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// Optionally filter by threshold notification status. If absent, only enabled
	// notifications will be returned.
	//
	// Any of "ENABLED", "DISABLED", "ARCHIVED".
	AlertStatuses []string `json:"alert_statuses,omitzero"`
	paramObj
}

func (r V1CustomerAlertListParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerAlertListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerAlertListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1CustomerAlertListParams]'s query parameters as
// `url.Values`.
func (r V1CustomerAlertListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerAlertResetParams struct {
	// The Metronome ID of the threshold notification
	AlertID string `json:"alert_id,required" format:"uuid"`
	// The Metronome ID of the customer
	CustomerID string `json:"customer_id,required" format:"uuid"`
	paramObj
}

func (r V1CustomerAlertResetParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerAlertResetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerAlertResetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
