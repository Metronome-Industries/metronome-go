// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	shimjson "github.com/Metronome-Industries/metronome-go/internal/encoding/json"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// V1CustomerService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerService] method instead.
type V1CustomerService struct {
	Options        []option.RequestOption
	Alerts         V1CustomerAlertService
	Plans          V1CustomerPlanService
	Invoices       V1CustomerInvoiceService
	BillingConfig  V1CustomerBillingConfigService
	Commits        V1CustomerCommitService
	Credits        V1CustomerCreditService
	NamedSchedules V1CustomerNamedScheduleService
}

// NewV1CustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomerService(opts ...option.RequestOption) (r V1CustomerService) {
	r = V1CustomerService{}
	r.Options = opts
	r.Alerts = NewV1CustomerAlertService(opts...)
	r.Plans = NewV1CustomerPlanService(opts...)
	r.Invoices = NewV1CustomerInvoiceService(opts...)
	r.BillingConfig = NewV1CustomerBillingConfigService(opts...)
	r.Commits = NewV1CustomerCommitService(opts...)
	r.Credits = NewV1CustomerCreditService(opts...)
	r.NamedSchedules = NewV1CustomerNamedScheduleService(opts...)
	return
}

// Create a new customer in Metronome and optionally the billing configuration
// (recommended) which dictates where invoices for the customer will be sent or
// where payment will be collected.
//
// ### Use this endpoint to:
//
// Execute your customer provisioning workflows for either PLG motions, where
// customers originate in your platform, or SLG motions, where customers originate
// in your sales system.
//
// ### Key response fields:
//
// This end-point returns the `customer_id` created by the request. This id can be
// used to fetch relevant billing configurations and create contracts.
//
// ### Example workflow:
//
//   - Generally, Metronome recommends first creating the customer in the downstream
//     payment / ERP system when payment method is collected and then creating the
//     customer in Metronome using the response (i.e. `customer_id`) from the
//     downstream system. If you do not create a billing configuration on customer
//     creation, you can add it later.
//   - Once a customer is created, you can then create a contract for the customer.
//     In the contract creation process, you will need to add the customer billing
//     configuration to the contract to ensure Metronome invoices the customer
//     correctly. This is because a customer can have multiple configurations.
//   - As part of the customer creation process, set the ingest alias for the
//     customer which will ensure usage is accurately mapped to the customer. Ingest
//     aliases can be added or changed after the creation process as well.
//
// ### Usage guidelines:
//
// For details on different billing configurations for different systems, review
// the `/setCustomerBillingConfiguration` end-point.
func (r *V1CustomerService) New(ctx context.Context, body V1CustomerNewParams, opts ...option.RequestOption) (res *V1CustomerNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get detailed information for a specific customer by their Metronome ID. Returns
// customer profile data including name, creation date, ingest aliases,
// configuration settings, and custom fields. Use this endpoint to fetch complete
// customer details for billing operations or account management.
//
// Note: If searching for a customer billing configuration, use the
// `/getCustomerBillingConfigurations` endpoint.
func (r *V1CustomerService) Get(ctx context.Context, query V1CustomerGetParams, opts ...option.RequestOption) (res *V1CustomerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s", query.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets a paginated list of all customers in your Metronome account. Use this
// endpoint to browse your customer base, implement customer search functionality,
// or sync customer data with external systems. Returns customer details including
// IDs, names, and configuration settings. Supports filtering and pagination
// parameters for efficient data retrieval.
func (r *V1CustomerService) List(ctx context.Context, query V1CustomerListParams, opts ...option.RequestOption) (res *pagination.CursorPage[CustomerDetail], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/customers"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Gets a paginated list of all customers in your Metronome account. Use this
// endpoint to browse your customer base, implement customer search functionality,
// or sync customer data with external systems. Returns customer details including
// IDs, names, and configuration settings. Supports filtering and pagination
// parameters for efficient data retrieval.
func (r *V1CustomerService) ListAutoPaging(ctx context.Context, query V1CustomerListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CustomerDetail] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Use this endpoint to archive a customer while preserving auditability. Archiving
// a customer will automatically archive all contracts as of the current date and
// void all corresponding invoices. Use this endpoint if a customer is onboarded by
// mistake.
//
// ### Usage guidelines:
//
//   - Once a customer is archived, it cannot be unarchived.
//   - Archived customers can still be viewed through the API or the UI for audit
//     purposes.
//   - Ingest aliases remain idempotent for archived customers. In order to reuse an
//     ingest alias, first remove the ingest alias from the customer prior to
//     archiving.
//   - Any notifications associated with the customer will no longer be triggered.
func (r *V1CustomerService) Archive(ctx context.Context, body V1CustomerArchiveParams, opts ...option.RequestOption) (res *V1CustomerArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/customers/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get all billable metrics available for a specific customer. Supports pagination
// and filtering by current plan status or archived metrics. Use this endpoint to
// see which metrics are being tracked for billing calculations for a given
// customer.
func (r *V1CustomerService) ListBillableMetrics(ctx context.Context, params V1CustomerListBillableMetricsParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerListBillableMetricsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/billable-metrics", params.CustomerID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Get all billable metrics available for a specific customer. Supports pagination
// and filtering by current plan status or archived metrics. Use this endpoint to
// see which metrics are being tracked for billing calculations for a given
// customer.
func (r *V1CustomerService) ListBillableMetricsAutoPaging(ctx context.Context, params V1CustomerListBillableMetricsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CustomerListBillableMetricsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListBillableMetrics(ctx, params, opts...))
}

// Fetch daily pending costs for the specified customer, broken down by credit type
// and line items. Note: this is not supported for customers whose plan includes a
// UNIQUE-type billable metric.
func (r *V1CustomerService) ListCosts(ctx context.Context, params V1CustomerListCostsParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerListCostsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/costs", params.CustomerID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Fetch daily pending costs for the specified customer, broken down by credit type
// and line items. Note: this is not supported for customers whose plan includes a
// UNIQUE-type billable metric.
func (r *V1CustomerService) ListCostsAutoPaging(ctx context.Context, params V1CustomerListCostsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CustomerListCostsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListCosts(ctx, params, opts...))
}

// Preview how a set of events will affect a customer's invoices. Generates draft
// invoices for a customer using their current contract configuration and the
// provided events. This is useful for testing how new events will affect the
// customer's invoices before they are actually processed.
func (r *V1CustomerService) PreviewEvents(ctx context.Context, params V1CustomerPreviewEventsParams, opts ...option.RequestOption) (res *V1CustomerPreviewEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/previewEvents", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns all billing configurations previously set for the customer. Use during
// the contract provisioning process to fetch the
// `billing_provider_configuration_id` needed to set the contract billing
// configuration.
func (r *V1CustomerService) GetBillingConfigurations(ctx context.Context, body V1CustomerGetBillingConfigurationsParams, opts ...option.RequestOption) (res *V1CustomerGetBillingConfigurationsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/getCustomerBillingProviderConfigurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a billing configuration for a customer. Once created, these
// configurations are available to associate to a contract and dictates which
// downstream system to collect payment in or send the invoice to. You can create
// multiple configurations per customer. The configuration formats are distinct for
// each downstream provider.
//
// ### Use this endpoint to:
//
//   - Add the initial configuration to an existing customer. Once created, the
//     billing configuration can then be associated to the customer's contract.
//   - Add a new configuration to an existing customer. This might be used as part of
//     an upgrade or downgrade workflow where the customer was previously billed
//     through system A (e.g. Stripe) but will now be billed through system B (e.g.
//     AWS). Once created, the new configuration can then be associated to the
//     customer's contract.
//   - Multiple configurations can be added per destination. For example, you can
//     create two Stripe billing configurations for a Metronome customer that each
//     have a distinct `collection_method`.
//
// ### Delivery method options:
//
//   - `direct_to_billing_provider`: Use when Metronome should send invoices directly
//     to the billing provider's API (e.g., Stripe, NetSuite). This is the most
//     common method for automated billing workflows.
//   - `tackle`: Use specifically for AWS Marketplace transactions that require
//     Tackle's co-selling platform for partner attribution and commission tracking.
//   - `aws_sqs`: Use when you want invoice data delivered to an AWS SQS queue for
//     custom processing before sending to your billing system.
//   - `aws_sns`: Use when you want invoice notifications published to an AWS SNS
//     topic for event-driven billing workflows.
//
// ### Key response fields:
//
// The id for the customer billing configuration. This id can be used to associate
// the billing configuration to a contract.
//
// ### Usage guidelines:
//
// Must use the `delivery_method_id` if you have multiple Stripe accounts connected
// to Metronome.
func (r *V1CustomerService) SetBillingConfigurations(ctx context.Context, body V1CustomerSetBillingConfigurationsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/setCustomerBillingProviderConfigurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Sets the ingest aliases for a customer. Use this endpoint to associate a
// Metronome customer with an internal ID for easier tracking between systems.
// Ingest aliases can be used in the `customer_id` field when sending usage events
// to Metronome.
//
// ### Usage guidelines:
//
//   - This call is idempotent and fully replaces the set of ingest aliases for the
//     given customer.
//   - Switching an ingest alias from one customer to another will associate all
//     corresponding usage to the new customer.
//   - Use multiple ingest aliases to model child organizations within a single
//     Metronome customer.
func (r *V1CustomerService) SetIngestAliases(ctx context.Context, params V1CustomerSetIngestAliasesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/setIngestAliases", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Updates the display name for a customer record. Use this to correct customer
// names, update business names after rebranding, or maintain accurate customer
// information for invoicing and reporting. Returns the updated customer object
// with the new name applied immediately across all billing documents and
// interfaces.
func (r *V1CustomerService) SetName(ctx context.Context, params V1CustomerSetNameParams, opts ...option.RequestOption) (res *V1CustomerSetNameResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/setName", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update configuration settings for a specific customer, such as external system
// integrations (e.g., Salesforce account ID) and other customer-specific billing
// parameters. Use this endpoint to modify customer configurations without
// affecting core customer data like name or ingest aliases.
func (r *V1CustomerService) UpdateConfig(ctx context.Context, params V1CustomerUpdateConfigParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/updateConfig", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type Customer struct {
	// the Metronome ID of the customer
	ID string `json:"id,required" format:"uuid"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string `json:"ingest_aliases,required"`
	Name          string   `json:"name,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ExternalID    respjson.Field
		IngestAliases respjson.Field
		Name          respjson.Field
		CustomFields  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Customer) RawJSON() string { return r.JSON.raw }
func (r *Customer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerDetail struct {
	// the Metronome ID of the customer
	ID string `json:"id,required" format:"uuid"`
	// RFC 3339 timestamp indicating when the customer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields   map[string]string            `json:"custom_fields,required"`
	CustomerConfig CustomerDetailCustomerConfig `json:"customer_config,required"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string `json:"ingest_aliases,required"`
	Name          string   `json:"name,required"`
	// RFC 3339 timestamp indicating when the customer was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// RFC 3339 timestamp indicating when the customer was archived. Null if the
	// customer is active.
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// This field's availability is dependent on your client's configuration.
	CurrentBillableStatus CustomerDetailCurrentBillableStatus `json:"current_billable_status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CustomFields          respjson.Field
		CustomerConfig        respjson.Field
		ExternalID            respjson.Field
		IngestAliases         respjson.Field
		Name                  respjson.Field
		UpdatedAt             respjson.Field
		ArchivedAt            respjson.Field
		CurrentBillableStatus respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerDetail) RawJSON() string { return r.JSON.raw }
func (r *CustomerDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerDetailCustomerConfig struct {
	// The Salesforce account ID for the customer
	SalesforceAccountID string `json:"salesforce_account_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SalesforceAccountID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerDetailCustomerConfig) RawJSON() string { return r.JSON.raw }
func (r *CustomerDetailCustomerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This field's availability is dependent on your client's configuration.
type CustomerDetailCurrentBillableStatus struct {
	// Any of "billable", "unbillable".
	Value       string    `json:"value,required"`
	EffectiveAt time.Time `json:"effective_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		EffectiveAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerDetailCurrentBillableStatus) RawJSON() string { return r.JSON.raw }
func (r *CustomerDetailCurrentBillableStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerNewResponse struct {
	Data Customer `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerGetResponse struct {
	Data CustomerDetail `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerArchiveResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerListBillableMetricsResponse struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// (DEPRECATED) use aggregation_type instead
	Aggregate string `json:"aggregate"`
	// (DEPRECATED) use aggregation_key instead
	AggregateKeys []string `json:"aggregate_keys"`
	// A key that specifies which property of the event is used to aggregate data. This
	// key must be one of the property filter names and is not applicable when the
	// aggregation type is 'count'.
	AggregationKey string `json:"aggregation_key"`
	// Specifies the type of aggregation performed on matching events.
	//
	// Any of "COUNT", "LATEST", "MAX", "SUM", "UNIQUE".
	AggregationType V1CustomerListBillableMetricsResponseAggregationType `json:"aggregation_type"`
	// RFC 3339 timestamp indicating when the billable metric was archived. If not
	// provided, the billable metric is not archived.
	ArchivedAt time.Time `json:"archived_at" format:"date-time"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter shared.EventTypeFilter `json:"event_type_filter"`
	// (DEPRECATED) use property_filters & event_type_filter instead
	Filter map[string]any `json:"filter"`
	// (DEPRECATED) use group_keys instead
	GroupBy []string `json:"group_by"`
	// Property names that are used to group usage costs on an invoice. Each entry
	// represents a set of properties used to slice events into distinct buckets.
	GroupKeys [][]string `json:"group_keys"`
	// A list of filters to match events to this billable metric. Each filter defines a
	// rule on an event property. All rules must pass for the event to match the
	// billable metric.
	PropertyFilters []shared.PropertyFilter `json:"property_filters"`
	// The SQL query associated with the billable metric
	Sql string `json:"sql"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Name            respjson.Field
		Aggregate       respjson.Field
		AggregateKeys   respjson.Field
		AggregationKey  respjson.Field
		AggregationType respjson.Field
		ArchivedAt      respjson.Field
		CustomFields    respjson.Field
		EventTypeFilter respjson.Field
		Filter          respjson.Field
		GroupBy         respjson.Field
		GroupKeys       respjson.Field
		PropertyFilters respjson.Field
		Sql             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListBillableMetricsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListBillableMetricsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the type of aggregation performed on matching events.
type V1CustomerListBillableMetricsResponseAggregationType string

const (
	V1CustomerListBillableMetricsResponseAggregationTypeCount  V1CustomerListBillableMetricsResponseAggregationType = "COUNT"
	V1CustomerListBillableMetricsResponseAggregationTypeLatest V1CustomerListBillableMetricsResponseAggregationType = "LATEST"
	V1CustomerListBillableMetricsResponseAggregationTypeMax    V1CustomerListBillableMetricsResponseAggregationType = "MAX"
	V1CustomerListBillableMetricsResponseAggregationTypeSum    V1CustomerListBillableMetricsResponseAggregationType = "SUM"
	V1CustomerListBillableMetricsResponseAggregationTypeUnique V1CustomerListBillableMetricsResponseAggregationType = "UNIQUE"
)

type V1CustomerListCostsResponse struct {
	CreditTypes    map[string]V1CustomerListCostsResponseCreditType `json:"credit_types,required"`
	EndTimestamp   time.Time                                        `json:"end_timestamp,required" format:"date-time"`
	StartTimestamp time.Time                                        `json:"start_timestamp,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypes    respjson.Field
		EndTimestamp   respjson.Field
		StartTimestamp respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListCostsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListCostsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerListCostsResponseCreditType struct {
	Cost              float64                                                  `json:"cost"`
	LineItemBreakdown []V1CustomerListCostsResponseCreditTypeLineItemBreakdown `json:"line_item_breakdown"`
	Name              string                                                   `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost              respjson.Field
		LineItemBreakdown respjson.Field
		Name              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListCostsResponseCreditType) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListCostsResponseCreditType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerListCostsResponseCreditTypeLineItemBreakdown struct {
	Cost       float64 `json:"cost,required"`
	Name       string  `json:"name,required"`
	GroupKey   string  `json:"group_key"`
	GroupValue string  `json:"group_value,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		Name        respjson.Field
		GroupKey    respjson.Field
		GroupValue  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListCostsResponseCreditTypeLineItemBreakdown) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListCostsResponseCreditTypeLineItemBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPreviewEventsResponse struct {
	Data []Invoice `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPreviewEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPreviewEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerGetBillingConfigurationsResponse struct {
	Data []V1CustomerGetBillingConfigurationsResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerGetBillingConfigurationsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerGetBillingConfigurationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerGetBillingConfigurationsResponseData struct {
	// ID of this configuration; can be provided as the
	// billing_provider_configuration_id when creating a contract.
	ID         string    `json:"id,required" format:"uuid"`
	ArchivedAt time.Time `json:"archived_at,required" format:"date-time"`
	// The billing provider set for this configuration.
	//
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace".
	BillingProvider string `json:"billing_provider,required"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider.
	Configuration map[string]any `json:"configuration,required"`
	CustomerID    string         `json:"customer_id,required" format:"uuid"`
	// The method to use for delivering invoices to this customer.
	//
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod string `json:"delivery_method,required"`
	// Configuration for the delivery method. The structure of this object is specific
	// to the delivery method.
	DeliveryMethodConfiguration map[string]any `json:"delivery_method_configuration,required"`
	// ID of the delivery method to use for this customer.
	DeliveryMethodID string `json:"delivery_method_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		ArchivedAt                  respjson.Field
		BillingProvider             respjson.Field
		Configuration               respjson.Field
		CustomerID                  respjson.Field
		DeliveryMethod              respjson.Field
		DeliveryMethodConfiguration respjson.Field
		DeliveryMethodID            respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerGetBillingConfigurationsResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerGetBillingConfigurationsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerSetNameResponse struct {
	Data Customer `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerSetNameResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerSetNameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerNewParams struct {
	// This will be truncated to 160 characters if the provided name is longer.
	Name string `json:"name,required"`
	// (deprecated, use ingest_aliases instead) an alias that can be used to refer to
	// this customer in usage events
	ExternalID    param.Opt[string]                `json:"external_id,omitzero"`
	BillingConfig V1CustomerNewParamsBillingConfig `json:"billing_config,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields                          map[string]string                                         `json:"custom_fields,omitzero"`
	CustomerBillingProviderConfigurations []V1CustomerNewParamsCustomerBillingProviderConfiguration `json:"customer_billing_provider_configurations,omitzero"`
	// Aliases that can be used to refer to this customer in usage events
	IngestAliases []string `json:"ingest_aliases,omitzero"`
	paramObj
}

func (r V1CustomerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BillingProviderCustomerID, BillingProviderType are required.
type V1CustomerNewParamsBillingConfig struct {
	BillingProviderCustomerID string `json:"billing_provider_customer_id,required"`
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace".
	BillingProviderType string `json:"billing_provider_type,omitzero,required"`
	// True if the aws_product_code is a SAAS subscription product, false otherwise.
	AwsIsSubscriptionProduct param.Opt[bool]   `json:"aws_is_subscription_product,omitzero"`
	AwsProductCode           param.Opt[string] `json:"aws_product_code,omitzero"`
	// Any of "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2",
	// "ap-northeast-3", "ap-south-1", "ap-southeast-1", "ap-southeast-2",
	// "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1",
	// "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "me-south-1", "sa-east-1",
	// "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1",
	// "us-west-2".
	AwsRegion string `json:"aws_region,omitzero"`
	// Any of "charge_automatically", "send_invoice", "auto_charge_payment_intent",
	// "manually_charge_payment_intent".
	StripeCollectionMethod string `json:"stripe_collection_method,omitzero"`
	paramObj
}

func (r V1CustomerNewParamsBillingConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerNewParamsBillingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerNewParamsBillingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerNewParamsBillingConfig](
		"billing_provider_type", "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace", "quickbooks_online", "workday", "gcp_marketplace",
	)
	apijson.RegisterFieldValidator[V1CustomerNewParamsBillingConfig](
		"aws_region", "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-southeast-1", "ap-southeast-2", "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1", "us-west-2",
	)
	apijson.RegisterFieldValidator[V1CustomerNewParamsBillingConfig](
		"stripe_collection_method", "charge_automatically", "send_invoice", "auto_charge_payment_intent", "manually_charge_payment_intent",
	)
}

// The property BillingProvider is required.
type V1CustomerNewParamsCustomerBillingProviderConfiguration struct {
	// The billing provider set for this configuration.
	//
	// Any of "aws_marketplace", "azure_marketplace", "gcp_marketplace", "stripe",
	// "netsuite".
	BillingProvider string `json:"billing_provider,omitzero,required"`
	// ID of the delivery method to use for this customer. If not provided, the
	// `delivery_method` must be provided.
	DeliveryMethodID param.Opt[string] `json:"delivery_method_id,omitzero" format:"uuid"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider and delivery provider combination. Defaults to an empty
	// object, however, for most billing provider + delivery method combinations, it
	// will not be a valid configuration.
	Configuration map[string]any `json:"configuration,omitzero"`
	// The method to use for delivering invoices to this customer. If not provided, the
	// `delivery_method_id` must be provided.
	//
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod string `json:"delivery_method,omitzero"`
	// Specifies which tax provider Metronome should use for tax calculation when
	// billing through Stripe. This is only supported for Stripe billing provider
	// configurations with auto_charge_payment_intent or manual_charge_payment_intent
	// collection methods.
	//
	// Any of "anrok", "avalara", "stripe".
	TaxProvider string `json:"tax_provider,omitzero"`
	paramObj
}

func (r V1CustomerNewParamsCustomerBillingProviderConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerNewParamsCustomerBillingProviderConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerNewParamsCustomerBillingProviderConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerNewParamsCustomerBillingProviderConfiguration](
		"billing_provider", "aws_marketplace", "azure_marketplace", "gcp_marketplace", "stripe", "netsuite",
	)
	apijson.RegisterFieldValidator[V1CustomerNewParamsCustomerBillingProviderConfiguration](
		"delivery_method", "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns",
	)
	apijson.RegisterFieldValidator[V1CustomerNewParamsCustomerBillingProviderConfiguration](
		"tax_provider", "anrok", "avalara", "stripe",
	)
}

type V1CustomerGetParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	paramObj
}

type V1CustomerListParams struct {
	// Filter the customer list by ingest_alias
	IngestAlias param.Opt[string] `query:"ingest_alias,omitzero" json:"-"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// Filter the customer list to only return archived customers. By default, only
	// active customers are returned.
	OnlyArchived param.Opt[bool] `query:"only_archived,omitzero" json:"-"`
	// Filter the customer list by customer_id. Up to 100 ids can be provided.
	CustomerIDs []string `query:"customer_ids,omitzero" json:"-"`
	// Filter the customer list by salesforce_account_id. Up to 100 ids can be
	// provided.
	SalesforceAccountIDs []string `query:"salesforce_account_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerListParams]'s query parameters as `url.Values`.
func (r V1CustomerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerArchiveParams struct {
	ID shared.IDParam
	paramObj
}

func (r V1CustomerArchiveParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ID)
}
func (r *V1CustomerArchiveParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ID)
}

type V1CustomerListBillableMetricsParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	// If true, the list of returned metrics will include archived metrics
	IncludeArchived param.Opt[bool] `query:"include_archived,omitzero" json:"-"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// If true, the list of metrics will be filtered to just ones that are on the
	// customer's current plan
	OnCurrentPlan param.Opt[bool] `query:"on_current_plan,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerListBillableMetricsParams]'s query parameters as
// `url.Values`.
func (r V1CustomerListBillableMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerListCostsParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `query:"ending_before,required" format:"date-time" json:"-"`
	// RFC 3339 timestamp (inclusive)
	StartingOn time.Time `query:"starting_on,required" format:"date-time" json:"-"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerListCostsParams]'s query parameters as
// `url.Values`.
func (r V1CustomerListCostsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerPreviewEventsParams struct {
	CustomerID string                               `path:"customer_id,required" format:"uuid" json:"-"`
	Events     []V1CustomerPreviewEventsParamsEvent `json:"events,omitzero,required"`
	// If set, all zero quantity line items will be filtered out of the response.
	SkipZeroQtyLineItems param.Opt[bool] `json:"skip_zero_qty_line_items,omitzero"`
	// If set to "replace", the preview will be generated as if those were the only
	// events for the specified customer. If set to "merge", the events will be merged
	// with any existing events for the specified customer. Defaults to "replace".
	//
	// Any of "replace", "merge".
	Mode V1CustomerPreviewEventsParamsMode `json:"mode,omitzero"`
	paramObj
}

func (r V1CustomerPreviewEventsParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPreviewEventsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPreviewEventsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property EventType is required.
type V1CustomerPreviewEventsParamsEvent struct {
	EventType string `json:"event_type,required"`
	// This has no effect for preview events, but may be set for consistency with Event
	// objects. They will be processed even if they do not match the customer's ID or
	// ingest aliases.
	CustomerID param.Opt[string] `json:"customer_id,omitzero"`
	// RFC 3339 formatted. If not provided, the current time will be used.
	Timestamp param.Opt[string] `json:"timestamp,omitzero"`
	// This has no effect for preview events, but may be set for consistency with Event
	// objects. Duplicate transaction_ids are NOT filtered out, even within the same
	// request.
	TransactionID param.Opt[string] `json:"transaction_id,omitzero"`
	Properties    map[string]any    `json:"properties,omitzero"`
	paramObj
}

func (r V1CustomerPreviewEventsParamsEvent) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPreviewEventsParamsEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPreviewEventsParamsEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// If set to "replace", the preview will be generated as if those were the only
// events for the specified customer. If set to "merge", the events will be merged
// with any existing events for the specified customer. Defaults to "replace".
type V1CustomerPreviewEventsParamsMode string

const (
	V1CustomerPreviewEventsParamsModeReplace V1CustomerPreviewEventsParamsMode = "replace"
	V1CustomerPreviewEventsParamsModeMerge   V1CustomerPreviewEventsParamsMode = "merge"
)

type V1CustomerGetBillingConfigurationsParams struct {
	CustomerID      string          `json:"customer_id,required" format:"uuid"`
	IncludeArchived param.Opt[bool] `json:"include_archived,omitzero"`
	paramObj
}

func (r V1CustomerGetBillingConfigurationsParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerGetBillingConfigurationsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerGetBillingConfigurationsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerSetBillingConfigurationsParams struct {
	Data []V1CustomerSetBillingConfigurationsParamsData `json:"data,omitzero,required"`
	paramObj
}

func (r V1CustomerSetBillingConfigurationsParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerSetBillingConfigurationsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerSetBillingConfigurationsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BillingProvider, CustomerID are required.
type V1CustomerSetBillingConfigurationsParamsData struct {
	// The billing provider set for this configuration.
	//
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace".
	BillingProvider string `json:"billing_provider,omitzero,required"`
	CustomerID      string `json:"customer_id,required" format:"uuid"`
	// ID of the delivery method to use for this customer. If not provided, the
	// `delivery_method` must be provided.
	DeliveryMethodID param.Opt[string] `json:"delivery_method_id,omitzero" format:"uuid"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider and delivery method combination. Defaults to an empty
	// object, however, for most billing provider + delivery method combinations, it
	// will not be a valid configuration. For AWS marketplace configurations, the
	// aws_is_subscription_product flag can be used to indicate a product with
	// usage-based pricing. More information can be found
	// [here](https://docs.metronome.com/invoice-customers/solutions/marketplaces/invoice-aws/#provision-aws-marketplace-customers-in-metronome).
	Configuration map[string]any `json:"configuration,omitzero"`
	// The method to use for delivering invoices to this customer. If not provided, the
	// `delivery_method_id` must be provided.
	//
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod string `json:"delivery_method,omitzero"`
	// Specifies which tax provider Metronome should use for tax calculation when
	// billing through Stripe. This is only supported for Stripe billing provider
	// configurations with auto_charge_payment_intent or manual_charge_payment_intent
	// collection methods.
	//
	// Any of "anrok", "avalara", "stripe".
	TaxProvider string `json:"tax_provider,omitzero"`
	paramObj
}

func (r V1CustomerSetBillingConfigurationsParamsData) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerSetBillingConfigurationsParamsData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerSetBillingConfigurationsParamsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerSetBillingConfigurationsParamsData](
		"billing_provider", "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace", "quickbooks_online", "workday", "gcp_marketplace",
	)
	apijson.RegisterFieldValidator[V1CustomerSetBillingConfigurationsParamsData](
		"delivery_method", "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns",
	)
	apijson.RegisterFieldValidator[V1CustomerSetBillingConfigurationsParamsData](
		"tax_provider", "anrok", "avalara", "stripe",
	)
}

type V1CustomerSetIngestAliasesParams struct {
	CustomerID    string   `path:"customer_id,required" format:"uuid" json:"-"`
	IngestAliases []string `json:"ingest_aliases,omitzero,required"`
	paramObj
}

func (r V1CustomerSetIngestAliasesParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerSetIngestAliasesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerSetIngestAliasesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerSetNameParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	// The new name for the customer. This will be truncated to 160 characters if the
	// provided name is longer.
	Name string `json:"name,required"`
	paramObj
}

func (r V1CustomerSetNameParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerSetNameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerSetNameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerUpdateConfigParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	// Leave in draft or set to auto-advance on invoices sent to Stripe. Falls back to
	// the client-level config if unset, which defaults to true if unset.
	LeaveStripeInvoicesInDraft param.Opt[bool] `json:"leave_stripe_invoices_in_draft,omitzero"`
	// The Salesforce account ID for the customer
	SalesforceAccountID param.Opt[string] `json:"salesforce_account_id,omitzero"`
	paramObj
}

func (r V1CustomerUpdateConfigParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerUpdateConfigParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerUpdateConfigParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
