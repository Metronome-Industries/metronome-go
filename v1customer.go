// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
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
	Alerts         *V1CustomerAlertService
	Plans          *V1CustomerPlanService
	Invoices       *V1CustomerInvoiceService
	BillingConfig  *V1CustomerBillingConfigService
	Commits        *V1CustomerCommitService
	Credits        *V1CustomerCreditService
	NamedSchedules *V1CustomerNamedScheduleService
}

// NewV1CustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomerService(opts ...option.RequestOption) (r *V1CustomerService) {
	r = &V1CustomerService{}
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
// Use this endpoint to:\
// Execute your customer provisioning workflows for either PLG motions, where customers
// originate in your platform, or SLG motions, where customers originate in your sales
// system.
//
//   - Key response fields: This end-point returns the customer_id created by the
//     request. This id can be used to fetch relevant billing configurations and
//     create contracts.
//
// Example workflow:
//
//   - Generally, Metronome recommends first creating the customer in the downstream
//     payment / ERP system when payment method is collected and then creating the
//     customer in Metronome using the response (i.e. customer_id) from the
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
// Usage guidelines:\
// For details on different billing configurations for different systems, review the
// /setCustomerBillingConfiguration end-point.
func (r *V1CustomerService) New(ctx context.Context, body V1CustomerNewParams, opts ...option.RequestOption) (res *V1CustomerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
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
// /getCustomerBillingConfigurations end-point.
func (r *V1CustomerService) Get(ctx context.Context, query V1CustomerGetParams, opts ...option.RequestOption) (res *V1CustomerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.CustomerID.Value == "" {
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
	opts = append(r.Options[:], opts...)
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
// Usage guidelines:
//
//   - Once a customer is archived, it cannot be unarchived.
//   - Archived customers can still be viewed through the API or the UI for audit
//     purposes.
//   - Ingest aliases remain idempotent for archived customers. In order to reuse an
//     ingest alias, first remove the ingest alias from the customer prior to
//     archiving.
//   - Any alerts associated with the customer will no longer be triggered.
func (r *V1CustomerService) Archive(ctx context.Context, body V1CustomerArchiveParams, opts ...option.RequestOption) (res *V1CustomerArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID.Value == "" {
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID.Value == "" {
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

// Preview how a set of events will affect a customer's invoice. Generates a draft
// invoice for a customer using their current contract configuration and the
// provided events. This is useful for testing how new events will affect the
// customer's invoice before they are actually processed.
func (r *V1CustomerService) PreviewEvents(ctx context.Context, params V1CustomerPreviewEventsParams, opts ...option.RequestOption) (res *V1CustomerPreviewEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.CustomerID.Value == "" {
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
func (r *V1CustomerService) GetCustomerBillingConfigurations(ctx context.Context, body V1CustomerGetCustomerBillingConfigurationsParams, opts ...option.RequestOption) (res *V1CustomerGetCustomerBillingConfigurationsResponse, err error) {
	opts = append(r.Options[:], opts...)
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
// Use this endpoint to:
//
//   - Add the initial configuration to an existing customer. Once created, the
//     billing configuration can then be associated to the customer's contract.
//   - Add a new configuration to an existing customer. This might be used as part of
//     an upgrade or downgrade workflow where the customer was previously billed
//     through system A (e.g. Stripe) but will now be billed through system B (e.g.
//     AWS). Once created, the new configuration can then be associated to the
//     customer's contract.
//
// Delivery Method Options:
//
//   - direct_to_billing_provider: Use when Metronome should send invoices directly
//     to the billing provider's API (e.g., Stripe, NetSuite). This is the most
//     common method for automated billing workflows.
//   - tackle: Use specifically for AWS Marketplace transactions that require
//     Tackle's co-selling platform for partner attribution and commission tracking.
//   - aws_sqs: Use when you want invoice data delivered to an AWS SQS queue for
//     custom processing before sending to your billing system.
//   - aws_sns: Use when you want invoice notifications published to an AWS SNS topic
//     for event-driven billing workflows.
//
// Key response fields: The id for the customer billing configuration. This id can
// be used to associate the billing configuration to a contract.
//
// Usage guidelines:\
// Must use the delivery_method_id if you have multiple Stripe accounts connected to
// Metronome.
func (r *V1CustomerService) SetCustomerBillingConfigurations(ctx context.Context, body V1CustomerSetCustomerBillingConfigurationsParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/setCustomerBillingProviderConfigurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Sets the ingest aliases for a customer. Use this endpoint to associate a
// Metronome customer with an internal ID for easier tracking between systems.
// Ingest aliases can be used in the customer_id field when sending usage events to
// Metronome.
//
// Usage guidelines:
//
//   - This call is idempotent and fully replaces the set of ingest aliases for the
//     given customer.
//   - Switching an ingest alias from one customer to another will associate all
//     corresponding usage to the new customer.
//   - Use multiple ingest aliases to model child organizations within a single
//     Metronome customer.
func (r *V1CustomerService) SetIngestAliases(ctx context.Context, params V1CustomerSetIngestAliasesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.CustomerID.Value == "" {
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
	opts = append(r.Options[:], opts...)
	if params.CustomerID.Value == "" {
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.CustomerID.Value == "" {
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
	JSON         customerJSON      `json:"-"`
}

// customerJSON contains the JSON metadata for the struct [Customer]
type customerJSON struct {
	ID            apijson.Field
	ExternalID    apijson.Field
	IngestAliases apijson.Field
	Name          apijson.Field
	CustomFields  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Customer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerJSON) RawJSON() string {
	return r.raw
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
	// RFC 3339 timestamp indicating when the customer was archived. Null if the
	// customer is active.
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// This field's availability is dependent on your client's configuration.
	CurrentBillableStatus CustomerDetailCurrentBillableStatus `json:"current_billable_status"`
	JSON                  customerDetailJSON                  `json:"-"`
}

// customerDetailJSON contains the JSON metadata for the struct [CustomerDetail]
type customerDetailJSON struct {
	ID                    apijson.Field
	CreatedAt             apijson.Field
	CustomFields          apijson.Field
	CustomerConfig        apijson.Field
	ExternalID            apijson.Field
	IngestAliases         apijson.Field
	Name                  apijson.Field
	ArchivedAt            apijson.Field
	CurrentBillableStatus apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CustomerDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerDetailJSON) RawJSON() string {
	return r.raw
}

type CustomerDetailCustomerConfig struct {
	// The Salesforce account ID for the customer
	SalesforceAccountID string                           `json:"salesforce_account_id,required,nullable"`
	JSON                customerDetailCustomerConfigJSON `json:"-"`
}

// customerDetailCustomerConfigJSON contains the JSON metadata for the struct
// [CustomerDetailCustomerConfig]
type customerDetailCustomerConfigJSON struct {
	SalesforceAccountID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerDetailCustomerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerDetailCustomerConfigJSON) RawJSON() string {
	return r.raw
}

// This field's availability is dependent on your client's configuration.
type CustomerDetailCurrentBillableStatus struct {
	Value       CustomerDetailCurrentBillableStatusValue `json:"value,required"`
	EffectiveAt time.Time                                `json:"effective_at,nullable" format:"date-time"`
	JSON        customerDetailCurrentBillableStatusJSON  `json:"-"`
}

// customerDetailCurrentBillableStatusJSON contains the JSON metadata for the
// struct [CustomerDetailCurrentBillableStatus]
type customerDetailCurrentBillableStatusJSON struct {
	Value       apijson.Field
	EffectiveAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerDetailCurrentBillableStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerDetailCurrentBillableStatusJSON) RawJSON() string {
	return r.raw
}

type CustomerDetailCurrentBillableStatusValue string

const (
	CustomerDetailCurrentBillableStatusValueBillable   CustomerDetailCurrentBillableStatusValue = "billable"
	CustomerDetailCurrentBillableStatusValueUnbillable CustomerDetailCurrentBillableStatusValue = "unbillable"
)

func (r CustomerDetailCurrentBillableStatusValue) IsKnown() bool {
	switch r {
	case CustomerDetailCurrentBillableStatusValueBillable, CustomerDetailCurrentBillableStatusValueUnbillable:
		return true
	}
	return false
}

type V1CustomerNewResponse struct {
	Data Customer                  `json:"data,required"`
	JSON v1CustomerNewResponseJSON `json:"-"`
}

// v1CustomerNewResponseJSON contains the JSON metadata for the struct
// [V1CustomerNewResponse]
type v1CustomerNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerGetResponse struct {
	Data CustomerDetail            `json:"data,required"`
	JSON v1CustomerGetResponseJSON `json:"-"`
}

// v1CustomerGetResponseJSON contains the JSON metadata for the struct
// [V1CustomerGetResponse]
type v1CustomerGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerArchiveResponse struct {
	Data shared.ID                     `json:"data,required"`
	JSON v1CustomerArchiveResponseJSON `json:"-"`
}

// v1CustomerArchiveResponseJSON contains the JSON metadata for the struct
// [V1CustomerArchiveResponse]
type v1CustomerArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerArchiveResponseJSON) RawJSON() string {
	return r.raw
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
	AggregationType V1CustomerListBillableMetricsResponseAggregationType `json:"aggregation_type"`
	// RFC 3339 timestamp indicating when the billable metric was archived. If not
	// provided, the billable metric is not archived.
	ArchivedAt time.Time `json:"archived_at" format:"date-time"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter shared.EventTypeFilter `json:"event_type_filter"`
	// (DEPRECATED) use property_filters & event_type_filter instead
	Filter map[string]interface{} `json:"filter"`
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
	Sql  string                                    `json:"sql"`
	JSON v1CustomerListBillableMetricsResponseJSON `json:"-"`
}

// v1CustomerListBillableMetricsResponseJSON contains the JSON metadata for the
// struct [V1CustomerListBillableMetricsResponse]
type v1CustomerListBillableMetricsResponseJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	Aggregate       apijson.Field
	AggregateKeys   apijson.Field
	AggregationKey  apijson.Field
	AggregationType apijson.Field
	ArchivedAt      apijson.Field
	CustomFields    apijson.Field
	EventTypeFilter apijson.Field
	Filter          apijson.Field
	GroupBy         apijson.Field
	GroupKeys       apijson.Field
	PropertyFilters apijson.Field
	Sql             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V1CustomerListBillableMetricsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListBillableMetricsResponseJSON) RawJSON() string {
	return r.raw
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

func (r V1CustomerListBillableMetricsResponseAggregationType) IsKnown() bool {
	switch r {
	case V1CustomerListBillableMetricsResponseAggregationTypeCount, V1CustomerListBillableMetricsResponseAggregationTypeLatest, V1CustomerListBillableMetricsResponseAggregationTypeMax, V1CustomerListBillableMetricsResponseAggregationTypeSum, V1CustomerListBillableMetricsResponseAggregationTypeUnique:
		return true
	}
	return false
}

type V1CustomerListCostsResponse struct {
	CreditTypes    map[string]V1CustomerListCostsResponseCreditType `json:"credit_types,required"`
	EndTimestamp   time.Time                                        `json:"end_timestamp,required" format:"date-time"`
	StartTimestamp time.Time                                        `json:"start_timestamp,required" format:"date-time"`
	JSON           v1CustomerListCostsResponseJSON                  `json:"-"`
}

// v1CustomerListCostsResponseJSON contains the JSON metadata for the struct
// [V1CustomerListCostsResponse]
type v1CustomerListCostsResponseJSON struct {
	CreditTypes    apijson.Field
	EndTimestamp   apijson.Field
	StartTimestamp apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1CustomerListCostsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListCostsResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerListCostsResponseCreditType struct {
	Cost              float64                                                   `json:"cost"`
	LineItemBreakdown []V1CustomerListCostsResponseCreditTypesLineItemBreakdown `json:"line_item_breakdown"`
	Name              string                                                    `json:"name"`
	JSON              v1CustomerListCostsResponseCreditTypeJSON                 `json:"-"`
}

// v1CustomerListCostsResponseCreditTypeJSON contains the JSON metadata for the
// struct [V1CustomerListCostsResponseCreditType]
type v1CustomerListCostsResponseCreditTypeJSON struct {
	Cost              apijson.Field
	LineItemBreakdown apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *V1CustomerListCostsResponseCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListCostsResponseCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerListCostsResponseCreditTypesLineItemBreakdown struct {
	Cost       float64                                                     `json:"cost,required"`
	Name       string                                                      `json:"name,required"`
	GroupKey   string                                                      `json:"group_key"`
	GroupValue string                                                      `json:"group_value,nullable"`
	JSON       v1CustomerListCostsResponseCreditTypesLineItemBreakdownJSON `json:"-"`
}

// v1CustomerListCostsResponseCreditTypesLineItemBreakdownJSON contains the JSON
// metadata for the struct
// [V1CustomerListCostsResponseCreditTypesLineItemBreakdown]
type v1CustomerListCostsResponseCreditTypesLineItemBreakdownJSON struct {
	Cost        apijson.Field
	Name        apijson.Field
	GroupKey    apijson.Field
	GroupValue  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerListCostsResponseCreditTypesLineItemBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListCostsResponseCreditTypesLineItemBreakdownJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponse struct {
	Data Invoice                             `json:"data,required"`
	JSON v1CustomerPreviewEventsResponseJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseJSON contains the JSON metadata for the struct
// [V1CustomerPreviewEventsResponse]
type v1CustomerPreviewEventsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerGetCustomerBillingConfigurationsResponse struct {
	Data []V1CustomerGetCustomerBillingConfigurationsResponseData `json:"data,required"`
	JSON v1CustomerGetCustomerBillingConfigurationsResponseJSON   `json:"-"`
}

// v1CustomerGetCustomerBillingConfigurationsResponseJSON contains the JSON
// metadata for the struct [V1CustomerGetCustomerBillingConfigurationsResponse]
type v1CustomerGetCustomerBillingConfigurationsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerGetCustomerBillingConfigurationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerGetCustomerBillingConfigurationsResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerGetCustomerBillingConfigurationsResponseData struct {
	// ID of this configuration; can be provided as the
	// billing_provider_configuration_id when creating a contract.
	ID string `json:"id,required" format:"uuid"`
	// The billing provider set for this configuration.
	BillingProvider V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProvider `json:"billing_provider,required"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider.
	Configuration map[string]interface{} `json:"configuration,required"`
	CustomerID    string                 `json:"customer_id,required" format:"uuid"`
	// The method to use for delivering invoices to this customer.
	DeliveryMethod V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethod `json:"delivery_method,required"`
	// Configuration for the delivery method. The structure of this object is specific
	// to the delivery method.
	DeliveryMethodConfiguration map[string]interface{} `json:"delivery_method_configuration,required"`
	// ID of the delivery method to use for this customer.
	DeliveryMethodID string                                                     `json:"delivery_method_id,required" format:"uuid"`
	JSON             v1CustomerGetCustomerBillingConfigurationsResponseDataJSON `json:"-"`
}

// v1CustomerGetCustomerBillingConfigurationsResponseDataJSON contains the JSON
// metadata for the struct [V1CustomerGetCustomerBillingConfigurationsResponseData]
type v1CustomerGetCustomerBillingConfigurationsResponseDataJSON struct {
	ID                          apijson.Field
	BillingProvider             apijson.Field
	Configuration               apijson.Field
	CustomerID                  apijson.Field
	DeliveryMethod              apijson.Field
	DeliveryMethodConfiguration apijson.Field
	DeliveryMethodID            apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *V1CustomerGetCustomerBillingConfigurationsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerGetCustomerBillingConfigurationsResponseDataJSON) RawJSON() string {
	return r.raw
}

// The billing provider set for this configuration.
type V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProvider string

const (
	V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderAwsMarketplace   V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProvider = "aws_marketplace"
	V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderStripe           V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProvider = "stripe"
	V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderNetsuite         V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProvider = "netsuite"
	V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderCustom           V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProvider = "custom"
	V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderAzureMarketplace V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProvider = "azure_marketplace"
	V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderQuickbooksOnline V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProvider = "quickbooks_online"
	V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderWorkday          V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProvider = "workday"
	V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderGcpMarketplace   V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProvider = "gcp_marketplace"
)

func (r V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProvider) IsKnown() bool {
	switch r {
	case V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderAwsMarketplace, V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderStripe, V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderNetsuite, V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderCustom, V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderAzureMarketplace, V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderQuickbooksOnline, V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderWorkday, V1CustomerGetCustomerBillingConfigurationsResponseDataBillingProviderGcpMarketplace:
		return true
	}
	return false
}

// The method to use for delivering invoices to this customer.
type V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethod string

const (
	V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethodDirectToBillingProvider V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethod = "direct_to_billing_provider"
	V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethodAwsSqs                  V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethod = "aws_sqs"
	V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethodTackle                  V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethod = "tackle"
	V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethodAwsSns                  V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethod = "aws_sns"
)

func (r V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethod) IsKnown() bool {
	switch r {
	case V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethodDirectToBillingProvider, V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethodAwsSqs, V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethodTackle, V1CustomerGetCustomerBillingConfigurationsResponseDataDeliveryMethodAwsSns:
		return true
	}
	return false
}

type V1CustomerSetNameResponse struct {
	Data Customer                      `json:"data,required"`
	JSON v1CustomerSetNameResponseJSON `json:"-"`
}

// v1CustomerSetNameResponseJSON contains the JSON metadata for the struct
// [V1CustomerSetNameResponse]
type v1CustomerSetNameResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerSetNameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerSetNameResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerNewParams struct {
	// This will be truncated to 160 characters if the provided name is longer.
	Name          param.Field[string]                           `json:"name,required"`
	BillingConfig param.Field[V1CustomerNewParamsBillingConfig] `json:"billing_config"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields                          param.Field[map[string]string]                                         `json:"custom_fields"`
	CustomerBillingProviderConfigurations param.Field[[]V1CustomerNewParamsCustomerBillingProviderConfiguration] `json:"customer_billing_provider_configurations"`
	// (deprecated, use ingest_aliases instead) an alias that can be used to refer to
	// this customer in usage events
	ExternalID param.Field[string] `json:"external_id"`
	// Aliases that can be used to refer to this customer in usage events
	IngestAliases param.Field[[]string] `json:"ingest_aliases"`
}

func (r V1CustomerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerNewParamsBillingConfig struct {
	BillingProviderCustomerID param.Field[string]                                              `json:"billing_provider_customer_id,required"`
	BillingProviderType       param.Field[V1CustomerNewParamsBillingConfigBillingProviderType] `json:"billing_provider_type,required"`
	// True if the aws_product_code is a SAAS subscription product, false otherwise.
	AwsIsSubscriptionProduct param.Field[bool]                                                   `json:"aws_is_subscription_product"`
	AwsProductCode           param.Field[string]                                                 `json:"aws_product_code"`
	AwsRegion                param.Field[V1CustomerNewParamsBillingConfigAwsRegion]              `json:"aws_region"`
	StripeCollectionMethod   param.Field[V1CustomerNewParamsBillingConfigStripeCollectionMethod] `json:"stripe_collection_method"`
}

func (r V1CustomerNewParamsBillingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerNewParamsBillingConfigBillingProviderType string

const (
	V1CustomerNewParamsBillingConfigBillingProviderTypeAwsMarketplace   V1CustomerNewParamsBillingConfigBillingProviderType = "aws_marketplace"
	V1CustomerNewParamsBillingConfigBillingProviderTypeStripe           V1CustomerNewParamsBillingConfigBillingProviderType = "stripe"
	V1CustomerNewParamsBillingConfigBillingProviderTypeNetsuite         V1CustomerNewParamsBillingConfigBillingProviderType = "netsuite"
	V1CustomerNewParamsBillingConfigBillingProviderTypeCustom           V1CustomerNewParamsBillingConfigBillingProviderType = "custom"
	V1CustomerNewParamsBillingConfigBillingProviderTypeAzureMarketplace V1CustomerNewParamsBillingConfigBillingProviderType = "azure_marketplace"
	V1CustomerNewParamsBillingConfigBillingProviderTypeQuickbooksOnline V1CustomerNewParamsBillingConfigBillingProviderType = "quickbooks_online"
	V1CustomerNewParamsBillingConfigBillingProviderTypeWorkday          V1CustomerNewParamsBillingConfigBillingProviderType = "workday"
	V1CustomerNewParamsBillingConfigBillingProviderTypeGcpMarketplace   V1CustomerNewParamsBillingConfigBillingProviderType = "gcp_marketplace"
)

func (r V1CustomerNewParamsBillingConfigBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerNewParamsBillingConfigBillingProviderTypeAwsMarketplace, V1CustomerNewParamsBillingConfigBillingProviderTypeStripe, V1CustomerNewParamsBillingConfigBillingProviderTypeNetsuite, V1CustomerNewParamsBillingConfigBillingProviderTypeCustom, V1CustomerNewParamsBillingConfigBillingProviderTypeAzureMarketplace, V1CustomerNewParamsBillingConfigBillingProviderTypeQuickbooksOnline, V1CustomerNewParamsBillingConfigBillingProviderTypeWorkday, V1CustomerNewParamsBillingConfigBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type V1CustomerNewParamsBillingConfigAwsRegion string

const (
	V1CustomerNewParamsBillingConfigAwsRegionAfSouth1     V1CustomerNewParamsBillingConfigAwsRegion = "af-south-1"
	V1CustomerNewParamsBillingConfigAwsRegionApEast1      V1CustomerNewParamsBillingConfigAwsRegion = "ap-east-1"
	V1CustomerNewParamsBillingConfigAwsRegionApNortheast1 V1CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-1"
	V1CustomerNewParamsBillingConfigAwsRegionApNortheast2 V1CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-2"
	V1CustomerNewParamsBillingConfigAwsRegionApNortheast3 V1CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-3"
	V1CustomerNewParamsBillingConfigAwsRegionApSouth1     V1CustomerNewParamsBillingConfigAwsRegion = "ap-south-1"
	V1CustomerNewParamsBillingConfigAwsRegionApSoutheast1 V1CustomerNewParamsBillingConfigAwsRegion = "ap-southeast-1"
	V1CustomerNewParamsBillingConfigAwsRegionApSoutheast2 V1CustomerNewParamsBillingConfigAwsRegion = "ap-southeast-2"
	V1CustomerNewParamsBillingConfigAwsRegionCaCentral1   V1CustomerNewParamsBillingConfigAwsRegion = "ca-central-1"
	V1CustomerNewParamsBillingConfigAwsRegionCnNorth1     V1CustomerNewParamsBillingConfigAwsRegion = "cn-north-1"
	V1CustomerNewParamsBillingConfigAwsRegionCnNorthwest1 V1CustomerNewParamsBillingConfigAwsRegion = "cn-northwest-1"
	V1CustomerNewParamsBillingConfigAwsRegionEuCentral1   V1CustomerNewParamsBillingConfigAwsRegion = "eu-central-1"
	V1CustomerNewParamsBillingConfigAwsRegionEuNorth1     V1CustomerNewParamsBillingConfigAwsRegion = "eu-north-1"
	V1CustomerNewParamsBillingConfigAwsRegionEuSouth1     V1CustomerNewParamsBillingConfigAwsRegion = "eu-south-1"
	V1CustomerNewParamsBillingConfigAwsRegionEuWest1      V1CustomerNewParamsBillingConfigAwsRegion = "eu-west-1"
	V1CustomerNewParamsBillingConfigAwsRegionEuWest2      V1CustomerNewParamsBillingConfigAwsRegion = "eu-west-2"
	V1CustomerNewParamsBillingConfigAwsRegionEuWest3      V1CustomerNewParamsBillingConfigAwsRegion = "eu-west-3"
	V1CustomerNewParamsBillingConfigAwsRegionMeSouth1     V1CustomerNewParamsBillingConfigAwsRegion = "me-south-1"
	V1CustomerNewParamsBillingConfigAwsRegionSaEast1      V1CustomerNewParamsBillingConfigAwsRegion = "sa-east-1"
	V1CustomerNewParamsBillingConfigAwsRegionUsEast1      V1CustomerNewParamsBillingConfigAwsRegion = "us-east-1"
	V1CustomerNewParamsBillingConfigAwsRegionUsEast2      V1CustomerNewParamsBillingConfigAwsRegion = "us-east-2"
	V1CustomerNewParamsBillingConfigAwsRegionUsGovEast1   V1CustomerNewParamsBillingConfigAwsRegion = "us-gov-east-1"
	V1CustomerNewParamsBillingConfigAwsRegionUsGovWest1   V1CustomerNewParamsBillingConfigAwsRegion = "us-gov-west-1"
	V1CustomerNewParamsBillingConfigAwsRegionUsWest1      V1CustomerNewParamsBillingConfigAwsRegion = "us-west-1"
	V1CustomerNewParamsBillingConfigAwsRegionUsWest2      V1CustomerNewParamsBillingConfigAwsRegion = "us-west-2"
)

func (r V1CustomerNewParamsBillingConfigAwsRegion) IsKnown() bool {
	switch r {
	case V1CustomerNewParamsBillingConfigAwsRegionAfSouth1, V1CustomerNewParamsBillingConfigAwsRegionApEast1, V1CustomerNewParamsBillingConfigAwsRegionApNortheast1, V1CustomerNewParamsBillingConfigAwsRegionApNortheast2, V1CustomerNewParamsBillingConfigAwsRegionApNortheast3, V1CustomerNewParamsBillingConfigAwsRegionApSouth1, V1CustomerNewParamsBillingConfigAwsRegionApSoutheast1, V1CustomerNewParamsBillingConfigAwsRegionApSoutheast2, V1CustomerNewParamsBillingConfigAwsRegionCaCentral1, V1CustomerNewParamsBillingConfigAwsRegionCnNorth1, V1CustomerNewParamsBillingConfigAwsRegionCnNorthwest1, V1CustomerNewParamsBillingConfigAwsRegionEuCentral1, V1CustomerNewParamsBillingConfigAwsRegionEuNorth1, V1CustomerNewParamsBillingConfigAwsRegionEuSouth1, V1CustomerNewParamsBillingConfigAwsRegionEuWest1, V1CustomerNewParamsBillingConfigAwsRegionEuWest2, V1CustomerNewParamsBillingConfigAwsRegionEuWest3, V1CustomerNewParamsBillingConfigAwsRegionMeSouth1, V1CustomerNewParamsBillingConfigAwsRegionSaEast1, V1CustomerNewParamsBillingConfigAwsRegionUsEast1, V1CustomerNewParamsBillingConfigAwsRegionUsEast2, V1CustomerNewParamsBillingConfigAwsRegionUsGovEast1, V1CustomerNewParamsBillingConfigAwsRegionUsGovWest1, V1CustomerNewParamsBillingConfigAwsRegionUsWest1, V1CustomerNewParamsBillingConfigAwsRegionUsWest2:
		return true
	}
	return false
}

type V1CustomerNewParamsBillingConfigStripeCollectionMethod string

const (
	V1CustomerNewParamsBillingConfigStripeCollectionMethodChargeAutomatically V1CustomerNewParamsBillingConfigStripeCollectionMethod = "charge_automatically"
	V1CustomerNewParamsBillingConfigStripeCollectionMethodSendInvoice         V1CustomerNewParamsBillingConfigStripeCollectionMethod = "send_invoice"
)

func (r V1CustomerNewParamsBillingConfigStripeCollectionMethod) IsKnown() bool {
	switch r {
	case V1CustomerNewParamsBillingConfigStripeCollectionMethodChargeAutomatically, V1CustomerNewParamsBillingConfigStripeCollectionMethodSendInvoice:
		return true
	}
	return false
}

type V1CustomerNewParamsCustomerBillingProviderConfiguration struct {
	// The billing provider set for this configuration.
	BillingProvider param.Field[V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider] `json:"billing_provider,required"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider and delivery provider combination. Defaults to an empty
	// object, however, for most billing provider + delivery method combinations, it
	// will not be a valid configuration.
	Configuration param.Field[map[string]interface{}] `json:"configuration"`
	// The method to use for delivering invoices to this customer. If not provided, the
	// `delivery_method_id` must be provided.
	DeliveryMethod param.Field[V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod] `json:"delivery_method"`
	// ID of the delivery method to use for this customer. If not provided, the
	// `delivery_method` must be provided.
	DeliveryMethodID param.Field[string] `json:"delivery_method_id" format:"uuid"`
}

func (r V1CustomerNewParamsCustomerBillingProviderConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The billing provider set for this configuration.
type V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider string

const (
	V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderAwsMarketplace   V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider = "aws_marketplace"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderAzureMarketplace V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider = "azure_marketplace"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderGcpMarketplace   V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider = "gcp_marketplace"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderStripe           V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider = "stripe"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderNetsuite         V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider = "netsuite"
)

func (r V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider) IsKnown() bool {
	switch r {
	case V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderAwsMarketplace, V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderAzureMarketplace, V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderGcpMarketplace, V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderStripe, V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderNetsuite:
		return true
	}
	return false
}

// The method to use for delivering invoices to this customer. If not provided, the
// `delivery_method_id` must be provided.
type V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod string

const (
	V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodDirectToBillingProvider V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod = "direct_to_billing_provider"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodAwsSqs                  V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod = "aws_sqs"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodTackle                  V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod = "tackle"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodAwsSns                  V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod = "aws_sns"
)

func (r V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod) IsKnown() bool {
	switch r {
	case V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodDirectToBillingProvider, V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodAwsSqs, V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodTackle, V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodAwsSns:
		return true
	}
	return false
}

type V1CustomerGetParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
}

type V1CustomerListParams struct {
	// Filter the customer list by customer_id. Up to 100 ids can be provided.
	CustomerIDs param.Field[[]string] `query:"customer_ids"`
	// Filter the customer list by ingest_alias
	IngestAlias param.Field[string] `query:"ingest_alias"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Filter the customer list to only return archived customers. By default, only
	// active customers are returned.
	OnlyArchived param.Field[bool] `query:"only_archived"`
	// Filter the customer list by salesforce_account_id. Up to 100 ids can be
	// provided.
	SalesforceAccountIDs param.Field[[]string] `query:"salesforce_account_ids"`
}

// URLQuery serializes [V1CustomerListParams]'s query parameters as `url.Values`.
func (r V1CustomerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerArchiveParams struct {
	ID shared.IDParam `json:"id,required"`
}

func (r V1CustomerArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ID)
}

type V1CustomerListBillableMetricsParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// If true, the list of returned metrics will include archived metrics
	IncludeArchived param.Field[bool] `query:"include_archived"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// If true, the list of metrics will be filtered to just ones that are on the
	// customer's current plan
	OnCurrentPlan param.Field[bool] `query:"on_current_plan"`
}

// URLQuery serializes [V1CustomerListBillableMetricsParams]'s query parameters as
// `url.Values`.
func (r V1CustomerListBillableMetricsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerListCostsParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `query:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingOn param.Field[time.Time] `query:"starting_on,required" format:"date-time"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [V1CustomerListCostsParams]'s query parameters as
// `url.Values`.
func (r V1CustomerListCostsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerPreviewEventsParams struct {
	CustomerID param.Field[string]                               `path:"customer_id,required" format:"uuid"`
	Events     param.Field[[]V1CustomerPreviewEventsParamsEvent] `json:"events,required"`
	// If set to "replace", the preview will be generated as if those were the only
	// events for the specified customer. If set to "merge", the events will be merged
	// with any existing events for the specified customer. Defaults to "replace".
	Mode param.Field[V1CustomerPreviewEventsParamsMode] `json:"mode"`
	// If set, all zero quantity line items will be filtered out of the response.
	SkipZeroQtyLineItems param.Field[bool] `json:"skip_zero_qty_line_items"`
}

func (r V1CustomerPreviewEventsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerPreviewEventsParamsEvent struct {
	EventType param.Field[string] `json:"event_type,required"`
	// This has no effect for preview events, but may be set for consistency with Event
	// objects. They will be processed even if they do not match the customer's ID or
	// ingest aliases.
	CustomerID param.Field[string]                 `json:"customer_id"`
	Properties param.Field[map[string]interface{}] `json:"properties"`
	// RFC 3339 formatted. If not provided, the current time will be used.
	Timestamp param.Field[string] `json:"timestamp"`
	// This has no effect for preview events, but may be set for consistency with Event
	// objects. Duplicate transaction_ids are NOT filtered out, even within the same
	// request.
	TransactionID param.Field[string] `json:"transaction_id"`
}

func (r V1CustomerPreviewEventsParamsEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set to "replace", the preview will be generated as if those were the only
// events for the specified customer. If set to "merge", the events will be merged
// with any existing events for the specified customer. Defaults to "replace".
type V1CustomerPreviewEventsParamsMode string

const (
	V1CustomerPreviewEventsParamsModeReplace V1CustomerPreviewEventsParamsMode = "replace"
	V1CustomerPreviewEventsParamsModeMerge   V1CustomerPreviewEventsParamsMode = "merge"
)

func (r V1CustomerPreviewEventsParamsMode) IsKnown() bool {
	switch r {
	case V1CustomerPreviewEventsParamsModeReplace, V1CustomerPreviewEventsParamsModeMerge:
		return true
	}
	return false
}

type V1CustomerGetCustomerBillingConfigurationsParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r V1CustomerGetCustomerBillingConfigurationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerSetCustomerBillingConfigurationsParams struct {
	Data param.Field[[]V1CustomerSetCustomerBillingConfigurationsParamsData] `json:"data,required"`
}

func (r V1CustomerSetCustomerBillingConfigurationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerSetCustomerBillingConfigurationsParamsData struct {
	// The billing provider set for this configuration.
	BillingProvider param.Field[V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProvider] `json:"billing_provider,required"`
	CustomerID      param.Field[string]                                                              `json:"customer_id,required" format:"uuid"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider and delivery method combination. Defaults to an empty
	// object, however, for most billing provider + delivery method combinations, it
	// will not be a valid configuration. For AWS marketplace configurations, the
	// aws_is_subscription_product flag can be used to indicate a product with
	// usage-based pricing. More information can be found
	// [here](https://docs.metronome.com/invoice-customers/solutions/marketplaces/invoice-aws/#provision-aws-marketplace-customers-in-metronome).
	Configuration param.Field[map[string]interface{}] `json:"configuration"`
	// The method to use for delivering invoices to this customer. If not provided, the
	// `delivery_method_id` must be provided.
	DeliveryMethod param.Field[V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethod] `json:"delivery_method"`
	// ID of the delivery method to use for this customer. If not provided, the
	// `delivery_method` must be provided.
	DeliveryMethodID param.Field[string] `json:"delivery_method_id" format:"uuid"`
}

func (r V1CustomerSetCustomerBillingConfigurationsParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The billing provider set for this configuration.
type V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProvider string

const (
	V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderAwsMarketplace   V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProvider = "aws_marketplace"
	V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderStripe           V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProvider = "stripe"
	V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderNetsuite         V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProvider = "netsuite"
	V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderCustom           V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProvider = "custom"
	V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderAzureMarketplace V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProvider = "azure_marketplace"
	V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderQuickbooksOnline V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProvider = "quickbooks_online"
	V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderWorkday          V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProvider = "workday"
	V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderGcpMarketplace   V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProvider = "gcp_marketplace"
)

func (r V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProvider) IsKnown() bool {
	switch r {
	case V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderAwsMarketplace, V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderStripe, V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderNetsuite, V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderCustom, V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderAzureMarketplace, V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderQuickbooksOnline, V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderWorkday, V1CustomerSetCustomerBillingConfigurationsParamsDataBillingProviderGcpMarketplace:
		return true
	}
	return false
}

// The method to use for delivering invoices to this customer. If not provided, the
// `delivery_method_id` must be provided.
type V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethod string

const (
	V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethodDirectToBillingProvider V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethod = "direct_to_billing_provider"
	V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethodAwsSqs                  V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethod = "aws_sqs"
	V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethodTackle                  V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethod = "tackle"
	V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethodAwsSns                  V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethod = "aws_sns"
)

func (r V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethod) IsKnown() bool {
	switch r {
	case V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethodDirectToBillingProvider, V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethodAwsSqs, V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethodTackle, V1CustomerSetCustomerBillingConfigurationsParamsDataDeliveryMethodAwsSns:
		return true
	}
	return false
}

type V1CustomerSetIngestAliasesParams struct {
	CustomerID    param.Field[string]   `path:"customer_id,required" format:"uuid"`
	IngestAliases param.Field[[]string] `json:"ingest_aliases,required"`
}

func (r V1CustomerSetIngestAliasesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerSetNameParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// The new name for the customer. This will be truncated to 160 characters if the
	// provided name is longer.
	Name param.Field[string] `json:"name,required"`
}

func (r V1CustomerSetNameParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerUpdateConfigParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// Leave in draft or set to auto-advance on invoices sent to Stripe. Falls back to
	// the client-level config if unset, which defaults to true if unset.
	LeaveStripeInvoicesInDraft param.Field[bool] `json:"leave_stripe_invoices_in_draft"`
	// The Salesforce account ID for the customer
	SalesforceAccountID param.Field[string] `json:"salesforce_account_id"`
}

func (r V1CustomerUpdateConfigParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
