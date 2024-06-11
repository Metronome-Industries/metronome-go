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
	"github.com/Metronome-Industries/metronome-go/internal/pagination"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// CustomerService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerService] method instead.
type CustomerService struct {
	Options       []option.RequestOption
	Alerts        *CustomerAlertService
	Plans         *CustomerPlanService
	Invoices      *CustomerInvoiceService
	BillingConfig *CustomerBillingConfigService
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r *CustomerService) {
	r = &CustomerService{}
	r.Options = opts
	r.Alerts = NewCustomerAlertService(opts...)
	r.Plans = NewCustomerPlanService(opts...)
	r.Invoices = NewCustomerInvoiceService(opts...)
	r.BillingConfig = NewCustomerBillingConfigService(opts...)
	return
}

// Create a new customer
func (r *CustomerService) New(ctx context.Context, body CustomerNewParams, opts ...option.RequestOption) (res *CustomerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a customer by Metronome ID.
func (r *CustomerService) Get(ctx context.Context, customerID string, opts ...option.RequestOption) (res *CustomerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all customers.
func (r *CustomerService) List(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) (res *pagination.CursorPage[CustomerDetail], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "customers"
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

// List all customers.
func (r *CustomerService) ListAutoPaging(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CustomerDetail] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Archive a customer
func (r *CustomerService) Archive(ctx context.Context, body CustomerArchiveParams, opts ...option.RequestOption) (res *CustomerArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customers/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get all billable metrics for a given customer.
func (r *CustomerService) ListBillableMetrics(ctx context.Context, customerID string, query CustomerListBillableMetricsParams, opts ...option.RequestOption) (res *pagination.CursorPage[CustomerListBillableMetricsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("customers/%s/billable-metrics", customerID)
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

// Get all billable metrics for a given customer.
func (r *CustomerService) ListBillableMetricsAutoPaging(ctx context.Context, customerID string, query CustomerListBillableMetricsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CustomerListBillableMetricsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListBillableMetrics(ctx, customerID, query, opts...))
}

// Fetch daily pending costs for the specified customer, broken down by credit type
// and line items. Note: this is not supported for customers whose plan includes a
// UNIQUE-type billable metric.
func (r *CustomerService) ListCosts(ctx context.Context, customerID string, query CustomerListCostsParams, opts ...option.RequestOption) (res *pagination.CursorPage[CustomerListCostsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("customers/%s/costs", customerID)
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

// Fetch daily pending costs for the specified customer, broken down by credit type
// and line items. Note: this is not supported for customers whose plan includes a
// UNIQUE-type billable metric.
func (r *CustomerService) ListCostsAutoPaging(ctx context.Context, customerID string, query CustomerListCostsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CustomerListCostsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListCosts(ctx, customerID, query, opts...))
}

// Sets the ingest aliases for a customer. Ingest aliases can be used in the
// `customer_id` field when sending usage events to Metronome. This call is
// idempotent. It fully replaces the set of ingest aliases for the given customer.
func (r *CustomerService) SetIngestAliases(ctx context.Context, customerID string, body CustomerSetIngestAliasesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/setIngestAliases", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Updates the specified customer's name.
func (r *CustomerService) SetName(ctx context.Context, customerID string, body CustomerSetNameParams, opts ...option.RequestOption) (res *CustomerSetNameResponse, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/setName", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates the specified customer's config.
func (r *CustomerService) UpdateConfig(ctx context.Context, customerID string, body CustomerUpdateConfigParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/updateConfig", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
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
	IngestAliases []string          `json:"ingest_aliases,required"`
	Name          string            `json:"name,required"`
	CustomFields  map[string]string `json:"custom_fields"`
	JSON          customerJSON      `json:"-"`
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
	ID                    string                              `json:"id,required" format:"uuid"`
	CurrentBillableStatus CustomerDetailCurrentBillableStatus `json:"current_billable_status,required"`
	CustomFields          map[string]string                   `json:"custom_fields,required"`
	CustomerConfig        CustomerDetailCustomerConfig        `json:"customer_config,required"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string           `json:"ingest_aliases,required"`
	Name          string             `json:"name,required"`
	JSON          customerDetailJSON `json:"-"`
}

// customerDetailJSON contains the JSON metadata for the struct [CustomerDetail]
type customerDetailJSON struct {
	ID                    apijson.Field
	CurrentBillableStatus apijson.Field
	CustomFields          apijson.Field
	CustomerConfig        apijson.Field
	ExternalID            apijson.Field
	IngestAliases         apijson.Field
	Name                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CustomerDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerDetailJSON) RawJSON() string {
	return r.raw
}

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

type CustomerNewResponse struct {
	Data Customer                `json:"data,required"`
	JSON customerNewResponseJSON `json:"-"`
}

// customerNewResponseJSON contains the JSON metadata for the struct
// [CustomerNewResponse]
type customerNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerNewResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerGetResponse struct {
	Data CustomerDetail          `json:"data,required"`
	JSON customerGetResponseJSON `json:"-"`
}

// customerGetResponseJSON contains the JSON metadata for the struct
// [CustomerGetResponse]
type customerGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerGetResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerArchiveResponse struct {
	Data shared.ID                   `json:"data,required"`
	JSON customerArchiveResponseJSON `json:"-"`
}

// customerArchiveResponseJSON contains the JSON metadata for the struct
// [CustomerArchiveResponse]
type customerArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerListBillableMetricsResponse struct {
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
	AggregationType CustomerListBillableMetricsResponseAggregationType `json:"aggregation_type"`
	CustomFields    map[string]string                                  `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter CustomerListBillableMetricsResponseEventTypeFilter `json:"event_type_filter"`
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
	PropertyFilters []CustomerListBillableMetricsResponsePropertyFilter `json:"property_filters"`
	JSON            customerListBillableMetricsResponseJSON             `json:"-"`
}

// customerListBillableMetricsResponseJSON contains the JSON metadata for the
// struct [CustomerListBillableMetricsResponse]
type customerListBillableMetricsResponseJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	Aggregate       apijson.Field
	AggregateKeys   apijson.Field
	AggregationKey  apijson.Field
	AggregationType apijson.Field
	CustomFields    apijson.Field
	EventTypeFilter apijson.Field
	Filter          apijson.Field
	GroupBy         apijson.Field
	GroupKeys       apijson.Field
	PropertyFilters apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CustomerListBillableMetricsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerListBillableMetricsResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of aggregation performed on matching events.
type CustomerListBillableMetricsResponseAggregationType string

const (
	CustomerListBillableMetricsResponseAggregationTypeCount  CustomerListBillableMetricsResponseAggregationType = "count"
	CustomerListBillableMetricsResponseAggregationTypeCount  CustomerListBillableMetricsResponseAggregationType = "Count"
	CustomerListBillableMetricsResponseAggregationTypeCount  CustomerListBillableMetricsResponseAggregationType = "COUNT"
	CustomerListBillableMetricsResponseAggregationTypeLatest CustomerListBillableMetricsResponseAggregationType = "latest"
	CustomerListBillableMetricsResponseAggregationTypeLatest CustomerListBillableMetricsResponseAggregationType = "Latest"
	CustomerListBillableMetricsResponseAggregationTypeLatest CustomerListBillableMetricsResponseAggregationType = "LATEST"
	CustomerListBillableMetricsResponseAggregationTypeMax    CustomerListBillableMetricsResponseAggregationType = "max"
	CustomerListBillableMetricsResponseAggregationTypeMax    CustomerListBillableMetricsResponseAggregationType = "Max"
	CustomerListBillableMetricsResponseAggregationTypeMax    CustomerListBillableMetricsResponseAggregationType = "MAX"
	CustomerListBillableMetricsResponseAggregationTypeSum    CustomerListBillableMetricsResponseAggregationType = "sum"
	CustomerListBillableMetricsResponseAggregationTypeSum    CustomerListBillableMetricsResponseAggregationType = "Sum"
	CustomerListBillableMetricsResponseAggregationTypeSum    CustomerListBillableMetricsResponseAggregationType = "SUM"
	CustomerListBillableMetricsResponseAggregationTypeUnique CustomerListBillableMetricsResponseAggregationType = "unique"
	CustomerListBillableMetricsResponseAggregationTypeUnique CustomerListBillableMetricsResponseAggregationType = "Unique"
	CustomerListBillableMetricsResponseAggregationTypeUnique CustomerListBillableMetricsResponseAggregationType = "UNIQUE"
)

func (r CustomerListBillableMetricsResponseAggregationType) IsKnown() bool {
	switch r {
	case CustomerListBillableMetricsResponseAggregationTypeCount, CustomerListBillableMetricsResponseAggregationTypeCount, CustomerListBillableMetricsResponseAggregationTypeCount, CustomerListBillableMetricsResponseAggregationTypeLatest, CustomerListBillableMetricsResponseAggregationTypeLatest, CustomerListBillableMetricsResponseAggregationTypeLatest, CustomerListBillableMetricsResponseAggregationTypeMax, CustomerListBillableMetricsResponseAggregationTypeMax, CustomerListBillableMetricsResponseAggregationTypeMax, CustomerListBillableMetricsResponseAggregationTypeSum, CustomerListBillableMetricsResponseAggregationTypeSum, CustomerListBillableMetricsResponseAggregationTypeSum, CustomerListBillableMetricsResponseAggregationTypeUnique, CustomerListBillableMetricsResponseAggregationTypeUnique, CustomerListBillableMetricsResponseAggregationTypeUnique:
		return true
	}
	return false
}

// An optional filtering rule to match the 'event_type' property of an event.
type CustomerListBillableMetricsResponseEventTypeFilter struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues []string `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues []string                                               `json:"not_in_values"`
	JSON        customerListBillableMetricsResponseEventTypeFilterJSON `json:"-"`
}

// customerListBillableMetricsResponseEventTypeFilterJSON contains the JSON
// metadata for the struct [CustomerListBillableMetricsResponseEventTypeFilter]
type customerListBillableMetricsResponseEventTypeFilterJSON struct {
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerListBillableMetricsResponseEventTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerListBillableMetricsResponseEventTypeFilterJSON) RawJSON() string {
	return r.raw
}

type CustomerListBillableMetricsResponsePropertyFilter struct {
	// The name of the event property.
	Name string `json:"name,required"`
	// Determines whether the property must exist in the event. If true, only events
	// with this property will pass the filter. If false, only events without this
	// property will pass the filter. If null or omitted, the existence of the property
	// is optional.
	Exists bool `json:"exists"`
	// Specifies the allowed values for the property to match an event. An event will
	// pass the filter only if its property value is included in this list. If
	// undefined, all property values will pass the filter. Must be non-empty if
	// present.
	InValues []string `json:"in_values"`
	// Specifies the values that prevent an event from matching the filter. An event
	// will not pass the filter if its property value is included in this list. If null
	// or empty, all property values will pass the filter. Must be non-empty if
	// present.
	NotInValues []string                                              `json:"not_in_values"`
	JSON        customerListBillableMetricsResponsePropertyFilterJSON `json:"-"`
}

// customerListBillableMetricsResponsePropertyFilterJSON contains the JSON metadata
// for the struct [CustomerListBillableMetricsResponsePropertyFilter]
type customerListBillableMetricsResponsePropertyFilterJSON struct {
	Name        apijson.Field
	Exists      apijson.Field
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerListBillableMetricsResponsePropertyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerListBillableMetricsResponsePropertyFilterJSON) RawJSON() string {
	return r.raw
}

type CustomerListCostsResponse struct {
	CreditTypes    map[string]CustomerListCostsResponseCreditType `json:"credit_types,required"`
	EndTimestamp   time.Time                                      `json:"end_timestamp,required" format:"date-time"`
	StartTimestamp time.Time                                      `json:"start_timestamp,required" format:"date-time"`
	JSON           customerListCostsResponseJSON                  `json:"-"`
}

// customerListCostsResponseJSON contains the JSON metadata for the struct
// [CustomerListCostsResponse]
type customerListCostsResponseJSON struct {
	CreditTypes    apijson.Field
	EndTimestamp   apijson.Field
	StartTimestamp apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerListCostsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerListCostsResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerListCostsResponseCreditType struct {
	Cost              float64                                                 `json:"cost"`
	LineItemBreakdown []CustomerListCostsResponseCreditTypesLineItemBreakdown `json:"line_item_breakdown"`
	Name              string                                                  `json:"name"`
	JSON              customerListCostsResponseCreditTypeJSON                 `json:"-"`
}

// customerListCostsResponseCreditTypeJSON contains the JSON metadata for the
// struct [CustomerListCostsResponseCreditType]
type customerListCostsResponseCreditTypeJSON struct {
	Cost              apijson.Field
	LineItemBreakdown apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerListCostsResponseCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerListCostsResponseCreditTypeJSON) RawJSON() string {
	return r.raw
}

type CustomerListCostsResponseCreditTypesLineItemBreakdown struct {
	Cost       float64                                                   `json:"cost,required"`
	Name       string                                                    `json:"name,required"`
	GroupKey   string                                                    `json:"group_key"`
	GroupValue string                                                    `json:"group_value,nullable"`
	JSON       customerListCostsResponseCreditTypesLineItemBreakdownJSON `json:"-"`
}

// customerListCostsResponseCreditTypesLineItemBreakdownJSON contains the JSON
// metadata for the struct [CustomerListCostsResponseCreditTypesLineItemBreakdown]
type customerListCostsResponseCreditTypesLineItemBreakdownJSON struct {
	Cost        apijson.Field
	Name        apijson.Field
	GroupKey    apijson.Field
	GroupValue  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerListCostsResponseCreditTypesLineItemBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerListCostsResponseCreditTypesLineItemBreakdownJSON) RawJSON() string {
	return r.raw
}

type CustomerSetNameResponse struct {
	Data Customer                    `json:"data,required"`
	JSON customerSetNameResponseJSON `json:"-"`
}

// customerSetNameResponseJSON contains the JSON metadata for the struct
// [CustomerSetNameResponse]
type customerSetNameResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerSetNameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerSetNameResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerNewParams struct {
	Name          param.Field[string]                         `json:"name,required"`
	BillingConfig param.Field[CustomerNewParamsBillingConfig] `json:"billing_config"`
	CustomFields  param.Field[map[string]string]              `json:"custom_fields"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome ID or ingest
	// alias) that can be used in usage events
	ExternalID param.Field[string] `json:"external_id"`
	// Aliases that can be used to refer to this customer in usage events
	IngestAliases param.Field[[]string] `json:"ingest_aliases"`
}

func (r CustomerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParamsBillingConfig struct {
	BillingProviderCustomerID param.Field[string]                                               `json:"billing_provider_customer_id,required"`
	BillingProviderType       param.Field[CustomerNewParamsBillingConfigBillingProviderType]    `json:"billing_provider_type,required"`
	AwsProductCode            param.Field[string]                                               `json:"aws_product_code"`
	AwsRegion                 param.Field[CustomerNewParamsBillingConfigAwsRegion]              `json:"aws_region"`
	StripeCollectionMethod    param.Field[CustomerNewParamsBillingConfigStripeCollectionMethod] `json:"stripe_collection_method"`
}

func (r CustomerNewParamsBillingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParamsBillingConfigBillingProviderType string

const (
	CustomerNewParamsBillingConfigBillingProviderTypeAwsMarketplace   CustomerNewParamsBillingConfigBillingProviderType = "aws_marketplace"
	CustomerNewParamsBillingConfigBillingProviderTypeStripe           CustomerNewParamsBillingConfigBillingProviderType = "stripe"
	CustomerNewParamsBillingConfigBillingProviderTypeNetsuite         CustomerNewParamsBillingConfigBillingProviderType = "netsuite"
	CustomerNewParamsBillingConfigBillingProviderTypeCustom           CustomerNewParamsBillingConfigBillingProviderType = "custom"
	CustomerNewParamsBillingConfigBillingProviderTypeAzureMarketplace CustomerNewParamsBillingConfigBillingProviderType = "azure_marketplace"
	CustomerNewParamsBillingConfigBillingProviderTypeQuickbooksOnline CustomerNewParamsBillingConfigBillingProviderType = "quickbooks_online"
	CustomerNewParamsBillingConfigBillingProviderTypeWorkday          CustomerNewParamsBillingConfigBillingProviderType = "workday"
	CustomerNewParamsBillingConfigBillingProviderTypeGcpMarketplace   CustomerNewParamsBillingConfigBillingProviderType = "gcp_marketplace"
)

func (r CustomerNewParamsBillingConfigBillingProviderType) IsKnown() bool {
	switch r {
	case CustomerNewParamsBillingConfigBillingProviderTypeAwsMarketplace, CustomerNewParamsBillingConfigBillingProviderTypeStripe, CustomerNewParamsBillingConfigBillingProviderTypeNetsuite, CustomerNewParamsBillingConfigBillingProviderTypeCustom, CustomerNewParamsBillingConfigBillingProviderTypeAzureMarketplace, CustomerNewParamsBillingConfigBillingProviderTypeQuickbooksOnline, CustomerNewParamsBillingConfigBillingProviderTypeWorkday, CustomerNewParamsBillingConfigBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type CustomerNewParamsBillingConfigAwsRegion string

const (
	CustomerNewParamsBillingConfigAwsRegionAfSouth1     CustomerNewParamsBillingConfigAwsRegion = "af-south-1"
	CustomerNewParamsBillingConfigAwsRegionApEast1      CustomerNewParamsBillingConfigAwsRegion = "ap-east-1"
	CustomerNewParamsBillingConfigAwsRegionApNortheast1 CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-1"
	CustomerNewParamsBillingConfigAwsRegionApNortheast2 CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-2"
	CustomerNewParamsBillingConfigAwsRegionApNortheast3 CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-3"
	CustomerNewParamsBillingConfigAwsRegionApSouth1     CustomerNewParamsBillingConfigAwsRegion = "ap-south-1"
	CustomerNewParamsBillingConfigAwsRegionApSoutheast1 CustomerNewParamsBillingConfigAwsRegion = "ap-southeast-1"
	CustomerNewParamsBillingConfigAwsRegionApSoutheast2 CustomerNewParamsBillingConfigAwsRegion = "ap-southeast-2"
	CustomerNewParamsBillingConfigAwsRegionCaCentral1   CustomerNewParamsBillingConfigAwsRegion = "ca-central-1"
	CustomerNewParamsBillingConfigAwsRegionCnNorth1     CustomerNewParamsBillingConfigAwsRegion = "cn-north-1"
	CustomerNewParamsBillingConfigAwsRegionCnNorthwest1 CustomerNewParamsBillingConfigAwsRegion = "cn-northwest-1"
	CustomerNewParamsBillingConfigAwsRegionEuCentral1   CustomerNewParamsBillingConfigAwsRegion = "eu-central-1"
	CustomerNewParamsBillingConfigAwsRegionEuNorth1     CustomerNewParamsBillingConfigAwsRegion = "eu-north-1"
	CustomerNewParamsBillingConfigAwsRegionEuSouth1     CustomerNewParamsBillingConfigAwsRegion = "eu-south-1"
	CustomerNewParamsBillingConfigAwsRegionEuWest1      CustomerNewParamsBillingConfigAwsRegion = "eu-west-1"
	CustomerNewParamsBillingConfigAwsRegionEuWest2      CustomerNewParamsBillingConfigAwsRegion = "eu-west-2"
	CustomerNewParamsBillingConfigAwsRegionEuWest3      CustomerNewParamsBillingConfigAwsRegion = "eu-west-3"
	CustomerNewParamsBillingConfigAwsRegionMeSouth1     CustomerNewParamsBillingConfigAwsRegion = "me-south-1"
	CustomerNewParamsBillingConfigAwsRegionSaEast1      CustomerNewParamsBillingConfigAwsRegion = "sa-east-1"
	CustomerNewParamsBillingConfigAwsRegionUsEast1      CustomerNewParamsBillingConfigAwsRegion = "us-east-1"
	CustomerNewParamsBillingConfigAwsRegionUsEast2      CustomerNewParamsBillingConfigAwsRegion = "us-east-2"
	CustomerNewParamsBillingConfigAwsRegionUsGovEast1   CustomerNewParamsBillingConfigAwsRegion = "us-gov-east-1"
	CustomerNewParamsBillingConfigAwsRegionUsGovWest1   CustomerNewParamsBillingConfigAwsRegion = "us-gov-west-1"
	CustomerNewParamsBillingConfigAwsRegionUsWest1      CustomerNewParamsBillingConfigAwsRegion = "us-west-1"
	CustomerNewParamsBillingConfigAwsRegionUsWest2      CustomerNewParamsBillingConfigAwsRegion = "us-west-2"
)

func (r CustomerNewParamsBillingConfigAwsRegion) IsKnown() bool {
	switch r {
	case CustomerNewParamsBillingConfigAwsRegionAfSouth1, CustomerNewParamsBillingConfigAwsRegionApEast1, CustomerNewParamsBillingConfigAwsRegionApNortheast1, CustomerNewParamsBillingConfigAwsRegionApNortheast2, CustomerNewParamsBillingConfigAwsRegionApNortheast3, CustomerNewParamsBillingConfigAwsRegionApSouth1, CustomerNewParamsBillingConfigAwsRegionApSoutheast1, CustomerNewParamsBillingConfigAwsRegionApSoutheast2, CustomerNewParamsBillingConfigAwsRegionCaCentral1, CustomerNewParamsBillingConfigAwsRegionCnNorth1, CustomerNewParamsBillingConfigAwsRegionCnNorthwest1, CustomerNewParamsBillingConfigAwsRegionEuCentral1, CustomerNewParamsBillingConfigAwsRegionEuNorth1, CustomerNewParamsBillingConfigAwsRegionEuSouth1, CustomerNewParamsBillingConfigAwsRegionEuWest1, CustomerNewParamsBillingConfigAwsRegionEuWest2, CustomerNewParamsBillingConfigAwsRegionEuWest3, CustomerNewParamsBillingConfigAwsRegionMeSouth1, CustomerNewParamsBillingConfigAwsRegionSaEast1, CustomerNewParamsBillingConfigAwsRegionUsEast1, CustomerNewParamsBillingConfigAwsRegionUsEast2, CustomerNewParamsBillingConfigAwsRegionUsGovEast1, CustomerNewParamsBillingConfigAwsRegionUsGovWest1, CustomerNewParamsBillingConfigAwsRegionUsWest1, CustomerNewParamsBillingConfigAwsRegionUsWest2:
		return true
	}
	return false
}

type CustomerNewParamsBillingConfigStripeCollectionMethod string

const (
	CustomerNewParamsBillingConfigStripeCollectionMethodChargeAutomatically CustomerNewParamsBillingConfigStripeCollectionMethod = "charge_automatically"
	CustomerNewParamsBillingConfigStripeCollectionMethodSendInvoice         CustomerNewParamsBillingConfigStripeCollectionMethod = "send_invoice"
)

func (r CustomerNewParamsBillingConfigStripeCollectionMethod) IsKnown() bool {
	switch r {
	case CustomerNewParamsBillingConfigStripeCollectionMethodChargeAutomatically, CustomerNewParamsBillingConfigStripeCollectionMethodSendInvoice:
		return true
	}
	return false
}

type CustomerListParams struct {
	// Filter the customer list by customer_id. Up to 100 ids can be provided.
	CustomerIDs param.Field[[]string] `query:"customer_ids"`
	// Filter the customer list by ingest_alias
	IngestAlias param.Field[string] `query:"ingest_alias"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Filter the customer list by only archived customers.
	OnlyArchived param.Field[bool] `query:"only_archived"`
	// Filter the customer list by salesforce_account_id. Up to 100 ids can be
	// provided.
	SalesforceAccountIDs param.Field[[]string] `query:"salesforce_account_ids"`
}

// URLQuery serializes [CustomerListParams]'s query parameters as `url.Values`.
func (r CustomerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerArchiveParams struct {
	ID shared.IDParam `json:"id,required"`
}

func (r CustomerArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ID)
}

type CustomerListBillableMetricsParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// If true, the list of metrics will be filtered to just ones that are on the
	// customer's current plan
	OnCurrentPlan param.Field[bool] `query:"on_current_plan"`
}

// URLQuery serializes [CustomerListBillableMetricsParams]'s query parameters as
// `url.Values`.
func (r CustomerListBillableMetricsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerListCostsParams struct {
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `query:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingOn param.Field[time.Time] `query:"starting_on,required" format:"date-time"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [CustomerListCostsParams]'s query parameters as
// `url.Values`.
func (r CustomerListCostsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerSetIngestAliasesParams struct {
	IngestAliases param.Field[[]string] `json:"ingest_aliases,required"`
}

func (r CustomerSetIngestAliasesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerSetNameParams struct {
	// The new name for the customer
	Name param.Field[string] `json:"name,required"`
}

func (r CustomerSetNameParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateConfigParams struct {
	// Leave in draft or set to auto-advance on invoices sent to Stripe. Falls back to
	// the client-level config if unset, which defaults to true if unset.
	LeaveStripeInvoicesInDraft param.Field[bool] `json:"leave_stripe_invoices_in_draft"`
	// The Salesforce account ID for the customer
	SalesforceAccountID param.Field[string] `json:"salesforce_account_id"`
}

func (r CustomerUpdateConfigParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
