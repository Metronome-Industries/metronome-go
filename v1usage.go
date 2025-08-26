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
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
)

// V1UsageService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1UsageService] method instead.
type V1UsageService struct {
	Options []option.RequestOption
}

// NewV1UsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1UsageService(opts ...option.RequestOption) (r *V1UsageService) {
	r = &V1UsageService{}
	r.Options = opts
	return
}

// Retrieve aggregated usage data across multiple customers and billable metrics in
// a single query. This batch endpoint enables you to fetch usage patterns at
// scale, broken down by time windows, making it ideal for building analytics
// dashboards, generating reports, and monitoring platform-wide usage trends.
//
// Use this endpoint to:
//
// - Generate platform-wide usage reports for internal teams
// - Monitor aggregate usage trends across your entire customer base
// - Create comparative usage analyses between customers or time periods
// - Support capacity planning with historical usage patterns
//
// Key response fields: An array of UsageBatchAggregate objects containing:
//
//   - customer_id: The customer this usage belongs to
//   - billable_metric_id and billable_metric_name: What was measured
//   - start_timestamp and end_timestamp: Time window for this data point
//   - value: Aggregated usage amount for the period
//   - groups (optional): Usage broken down by group keys with values -next_page:
//     Pagination cursor for large result sets
//
// Usage guidelines:
//
// - Time windows: Set window_size to hour, day, or none (entire period)
// - Required parameters: Must specify starting_on, ending_before, and window_size
// - Filtering options:
//   - customer_ids: Limit to specific customers (omit for all customers)
//   - billable_metrics: Limit to specific metrics (omit for all metrics)
//
// - Pagination: Use next_page cursor to retrieve large datasets
// - Null values: Group values may be null when no usage matches that group
func (r *V1UsageService) List(ctx context.Context, params V1UsageListParams, opts ...option.RequestOption) (res *pagination.CursorPageWithoutLimit[V1UsageListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/usage"
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

// Retrieve aggregated usage data across multiple customers and billable metrics in
// a single query. This batch endpoint enables you to fetch usage patterns at
// scale, broken down by time windows, making it ideal for building analytics
// dashboards, generating reports, and monitoring platform-wide usage trends.
//
// Use this endpoint to:
//
// - Generate platform-wide usage reports for internal teams
// - Monitor aggregate usage trends across your entire customer base
// - Create comparative usage analyses between customers or time periods
// - Support capacity planning with historical usage patterns
//
// Key response fields: An array of UsageBatchAggregate objects containing:
//
//   - customer_id: The customer this usage belongs to
//   - billable_metric_id and billable_metric_name: What was measured
//   - start_timestamp and end_timestamp: Time window for this data point
//   - value: Aggregated usage amount for the period
//   - groups (optional): Usage broken down by group keys with values -next_page:
//     Pagination cursor for large result sets
//
// Usage guidelines:
//
// - Time windows: Set window_size to hour, day, or none (entire period)
// - Required parameters: Must specify starting_on, ending_before, and window_size
// - Filtering options:
//   - customer_ids: Limit to specific customers (omit for all customers)
//   - billable_metrics: Limit to specific metrics (omit for all metrics)
//
// - Pagination: Use next_page cursor to retrieve large datasets
// - Null values: Group values may be null when no usage matches that group
func (r *V1UsageService) ListAutoPaging(ctx context.Context, params V1UsageListParams, opts ...option.RequestOption) *pagination.CursorPageWithoutLimitAutoPager[V1UsageListResponse] {
	return pagination.NewCursorPageWithoutLimitAutoPager(r.List(ctx, params, opts...))
}

// The ingest endpoint is the primary method for sending usage events to Metronome,
// serving as the foundation for all billing calculations in your usage-based
// pricing model. This high-throughput endpoint is designed for real-time streaming
// ingestion, supports backdating 34 days, and is built to handle mission-critical
// usage data with enterprise-grade reliability. Metronome supports 100,000 events
// per second without requiring pre-aggregation or rollups and can scale up from
// there. See
// [Getting usage into Metronome](https://docs.metronome.com/connect-metronome/) to
// learn more about usage events.
//
// Use this endpoint to:\
// Create a customer usage pipeline into Metronome that drives billable metrics, credit
// drawdown, and invoicing. Track customer behavior, resource consumption, and feature
// usage
//
// What happens when you send events:
//
//   - Events are validated and processed in real-time
//   - Events are matched to customers using customer IDs or customer ingest aliases
//   - Events are matched to billable metrics and are immediately available for usage
//     and spend calculations
//
// Usage guidelines:
//
//   - Historical events can be backdated up to 34 days and will immediately impact
//     live customer spend
//   - Duplicate events are automatically detected and ignored (34-day deduplication
//     window)
//
// Event structure: Usage events are simple JSON objects designed for flexibility
// and ease of integration:
//
// ```json
//
//	{
//	  "transaction_id": "2021-01-01T00:00:00Z_cluster42",
//	  "customer_id": "team@example.com",
//	  "event_type": "api_request",
//	  "timestamp": "2021-01-01T00:00:00Z",
//	  "properties": {
//	    "endpoint": "/v1/users",
//	    "method": "POST",
//	    "response_time_ms": 45,
//	    "region": "us-west-2"
//	  }
//	}
//
// ```
//
//   - Transaction ID\
//     The transaction_id serves as your idempotency key, ensuring events are processed
//     exactly once. Metronome maintains a 34-day deduplication window - significantly
//     longer than typical 12-hour windows - enabling robust backfill scenarios without
//     duplicate billing.
//
//   - Best Practices:
//
//   - Use UUIDs for one-time events: uuid4()
//
//   - For heartbeat events, use deterministic IDs
//
//   - Include enough context to avoid collisions across different event sources
//
//   - Customer ID\
//     Identifies which customer should be billed for this usage. Supports two identification
//     methods:
//   - Metronome Customer ID: The UUID returned when creating a customer
//   - Ingest Alias: Your system's identifier (email, account number, etc.)
//
// Ingest aliases enable seamless integration without requiring ID mapping, and
// customers can have multiple aliases for flexibility.
//
//   - Event Type: Categorizes the event type for billable metric matching. Choose
//     descriptive names that aligns with the product surface area.
//
//   - Properties: Flexible metadata also used to match billable metrics or to be
//     used to serve as group keys to create multiple pricing dimensions or breakdown
//     costs by novel properties for end customers or internal finance teams
//     measuring underlying COGs.
func (r *V1UsageService) Ingest(ctx context.Context, body V1UsageIngestParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/ingest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve granular usage data for a specific customer and billable metric, with
// the ability to break down usage by custom grouping dimensions. This endpoint
// enables deep usage analytics by segmenting data across attributes like region,
// user, model type, or any custom dimension defined in your billable metrics.
//
// Use this endpoint to:
//
//   - Analyze usage patterns broken down by specific attributes (region, user,
//     department, etc.)
//   - Build detailed usage dashboards with dimensional filtering
//   - Identify high-usage segments for optimization opportunities
//
// Key response fields: An array of PagedUsageAggregate objects containing:
//
// - starting_on and ending_before: Time window boundaries
// - group_key: The dimension being grouped by (e.g., "region")
// - group_value: The specific value for this group (e.g., "US-East")
// - value: Aggregated usage for this group and time window
// - next_page: Pagination cursor for large datasets
//
// Usage guidelines:
//
//   - Required parameters: Must specify customer_id, billable_metric_id, and
//     window_size
//   - Time windows: Set window_size to hour, day, or none for different
//     granularities
//   - Group filtering: Use group_by to specify:
//   - key: The dimension to group by (must be set on the billable metric as a
//     group key)
//   - values: Optional array to filter to specific values only
//   - Pagination: Use limit and next_page for large result sets
//   - Null handling: group_value may be null for unmatched data
func (r *V1UsageService) ListWithGroups(ctx context.Context, params V1UsageListWithGroupsParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1UsageListWithGroupsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/usage/groups"
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

// Retrieve granular usage data for a specific customer and billable metric, with
// the ability to break down usage by custom grouping dimensions. This endpoint
// enables deep usage analytics by segmenting data across attributes like region,
// user, model type, or any custom dimension defined in your billable metrics.
//
// Use this endpoint to:
//
//   - Analyze usage patterns broken down by specific attributes (region, user,
//     department, etc.)
//   - Build detailed usage dashboards with dimensional filtering
//   - Identify high-usage segments for optimization opportunities
//
// Key response fields: An array of PagedUsageAggregate objects containing:
//
// - starting_on and ending_before: Time window boundaries
// - group_key: The dimension being grouped by (e.g., "region")
// - group_value: The specific value for this group (e.g., "US-East")
// - value: Aggregated usage for this group and time window
// - next_page: Pagination cursor for large datasets
//
// Usage guidelines:
//
//   - Required parameters: Must specify customer_id, billable_metric_id, and
//     window_size
//   - Time windows: Set window_size to hour, day, or none for different
//     granularities
//   - Group filtering: Use group_by to specify:
//   - key: The dimension to group by (must be set on the billable metric as a
//     group key)
//   - values: Optional array to filter to specific values only
//   - Pagination: Use limit and next_page for large result sets
//   - Null handling: group_value may be null for unmatched data
func (r *V1UsageService) ListWithGroupsAutoPaging(ctx context.Context, params V1UsageListWithGroupsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1UsageListWithGroupsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListWithGroups(ctx, params, opts...))
}

// This endpoint retrieves events by transaction ID for events that occurred within
// the last 34 days. It is specifically designed for sampling-based testing
// workflows to detect revenue leakage. The Event Search API provides a critical
// observability tool that validates the integrity of your usage pipeline by
// allowing you to sample raw events and verify their matching against active
// billable metrics.
//
// Why event observability matters for revenue leakage: Silent revenue loss occurs
// when events are dropped, delayed, or fail to match billable metrics due to:
//
// - Upstream system failures
// - Event format changes
// - Misconfigured billable metrics
//
// Use this endpoint to:
//
// - Sample raw events and validate they match the expected billable metrics
// - Build custom leakage detection alerts to prevent silent revenue loss
// - Verify event processing accuracy during system changes or metric updates
// - Debug event matching issues in real-time
//
// Key response fields:
//
// - Complete event details including transaction ID, customer ID, and properties
// - Matched Metronome customer (if any)
// - Matched billable metric information (if any)
// - Processing status and duplicate detection flags
//
// Usage guidelines:\
// ⚠️ Important: This endpoint is heavily rate limited and designed for sampling workflows
// only. Do not use this endpoint to check every event in your system. Instead, implement
// a sampling strategy to randomly validate a subset of events for observability purposes.
func (r *V1UsageService) Search(ctx context.Context, body V1UsageSearchParams, opts ...option.RequestOption) (res *[]V1UsageSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/events/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1UsageListResponse struct {
	BillableMetricID   string    `json:"billable_metric_id,required" format:"uuid"`
	BillableMetricName string    `json:"billable_metric_name,required"`
	CustomerID         string    `json:"customer_id,required" format:"uuid"`
	EndTimestamp       time.Time `json:"end_timestamp,required" format:"date-time"`
	StartTimestamp     time.Time `json:"start_timestamp,required" format:"date-time"`
	Value              float64   `json:"value,required,nullable"`
	// Values will be either a number or null. Null indicates that there were no
	// matches for the group_by value.
	Groups map[string]float64      `json:"groups"`
	JSON   v1UsageListResponseJSON `json:"-"`
}

// v1UsageListResponseJSON contains the JSON metadata for the struct
// [V1UsageListResponse]
type v1UsageListResponseJSON struct {
	BillableMetricID   apijson.Field
	BillableMetricName apijson.Field
	CustomerID         apijson.Field
	EndTimestamp       apijson.Field
	StartTimestamp     apijson.Field
	Value              apijson.Field
	Groups             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V1UsageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1UsageListResponseJSON) RawJSON() string {
	return r.raw
}

type V1UsageListWithGroupsResponse struct {
	EndingBefore time.Time                         `json:"ending_before,required" format:"date-time"`
	GroupKey     string                            `json:"group_key,required,nullable"`
	GroupValue   string                            `json:"group_value,required,nullable"`
	StartingOn   time.Time                         `json:"starting_on,required" format:"date-time"`
	Value        float64                           `json:"value,required,nullable"`
	JSON         v1UsageListWithGroupsResponseJSON `json:"-"`
}

// v1UsageListWithGroupsResponseJSON contains the JSON metadata for the struct
// [V1UsageListWithGroupsResponse]
type v1UsageListWithGroupsResponseJSON struct {
	EndingBefore apijson.Field
	GroupKey     apijson.Field
	GroupValue   apijson.Field
	StartingOn   apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1UsageListWithGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1UsageListWithGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type V1UsageSearchResponse struct {
	ID string `json:"id,required"`
	// The ID of the customer in the ingest event body
	CustomerID             string                                       `json:"customer_id,required"`
	EventType              string                                       `json:"event_type,required"`
	Timestamp              time.Time                                    `json:"timestamp,required" format:"date-time"`
	TransactionID          string                                       `json:"transaction_id,required"`
	IsDuplicate            bool                                         `json:"is_duplicate"`
	MatchedBillableMetrics []V1UsageSearchResponseMatchedBillableMetric `json:"matched_billable_metrics"`
	// The customer the event was matched to if a match was found
	MatchedCustomer V1UsageSearchResponseMatchedCustomer `json:"matched_customer"`
	ProcessedAt     time.Time                            `json:"processed_at" format:"date-time"`
	Properties      map[string]interface{}               `json:"properties"`
	JSON            v1UsageSearchResponseJSON            `json:"-"`
}

// v1UsageSearchResponseJSON contains the JSON metadata for the struct
// [V1UsageSearchResponse]
type v1UsageSearchResponseJSON struct {
	ID                     apijson.Field
	CustomerID             apijson.Field
	EventType              apijson.Field
	Timestamp              apijson.Field
	TransactionID          apijson.Field
	IsDuplicate            apijson.Field
	MatchedBillableMetrics apijson.Field
	MatchedCustomer        apijson.Field
	ProcessedAt            apijson.Field
	Properties             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1UsageSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1UsageSearchResponseJSON) RawJSON() string {
	return r.raw
}

type V1UsageSearchResponseMatchedBillableMetric struct {
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
	AggregationType V1UsageSearchResponseMatchedBillableMetricsAggregationType `json:"aggregation_type"`
	// RFC 3339 timestamp indicating when the billable metric was archived. If not
	// provided, the billable metric is not archived.
	ArchivedAt time.Time `json:"archived_at" format:"date-time"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter V1UsageSearchResponseMatchedBillableMetricsEventTypeFilter `json:"event_type_filter"`
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
	PropertyFilters []V1UsageSearchResponseMatchedBillableMetricsPropertyFilter `json:"property_filters"`
	// The SQL query associated with the billable metric
	Sql  string                                         `json:"sql"`
	JSON v1UsageSearchResponseMatchedBillableMetricJSON `json:"-"`
}

// v1UsageSearchResponseMatchedBillableMetricJSON contains the JSON metadata for
// the struct [V1UsageSearchResponseMatchedBillableMetric]
type v1UsageSearchResponseMatchedBillableMetricJSON struct {
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

func (r *V1UsageSearchResponseMatchedBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1UsageSearchResponseMatchedBillableMetricJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of aggregation performed on matching events.
type V1UsageSearchResponseMatchedBillableMetricsAggregationType string

const (
	V1UsageSearchResponseMatchedBillableMetricsAggregationTypeCount  V1UsageSearchResponseMatchedBillableMetricsAggregationType = "COUNT"
	V1UsageSearchResponseMatchedBillableMetricsAggregationTypeLatest V1UsageSearchResponseMatchedBillableMetricsAggregationType = "LATEST"
	V1UsageSearchResponseMatchedBillableMetricsAggregationTypeMax    V1UsageSearchResponseMatchedBillableMetricsAggregationType = "MAX"
	V1UsageSearchResponseMatchedBillableMetricsAggregationTypeSum    V1UsageSearchResponseMatchedBillableMetricsAggregationType = "SUM"
	V1UsageSearchResponseMatchedBillableMetricsAggregationTypeUnique V1UsageSearchResponseMatchedBillableMetricsAggregationType = "UNIQUE"
)

func (r V1UsageSearchResponseMatchedBillableMetricsAggregationType) IsKnown() bool {
	switch r {
	case V1UsageSearchResponseMatchedBillableMetricsAggregationTypeCount, V1UsageSearchResponseMatchedBillableMetricsAggregationTypeLatest, V1UsageSearchResponseMatchedBillableMetricsAggregationTypeMax, V1UsageSearchResponseMatchedBillableMetricsAggregationTypeSum, V1UsageSearchResponseMatchedBillableMetricsAggregationTypeUnique:
		return true
	}
	return false
}

// An optional filtering rule to match the 'event_type' property of an event.
type V1UsageSearchResponseMatchedBillableMetricsEventTypeFilter struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues []string `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues []string                                                       `json:"not_in_values"`
	JSON        v1UsageSearchResponseMatchedBillableMetricsEventTypeFilterJSON `json:"-"`
}

// v1UsageSearchResponseMatchedBillableMetricsEventTypeFilterJSON contains the JSON
// metadata for the struct
// [V1UsageSearchResponseMatchedBillableMetricsEventTypeFilter]
type v1UsageSearchResponseMatchedBillableMetricsEventTypeFilterJSON struct {
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1UsageSearchResponseMatchedBillableMetricsEventTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1UsageSearchResponseMatchedBillableMetricsEventTypeFilterJSON) RawJSON() string {
	return r.raw
}

type V1UsageSearchResponseMatchedBillableMetricsPropertyFilter struct {
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
	NotInValues []string                                                      `json:"not_in_values"`
	JSON        v1UsageSearchResponseMatchedBillableMetricsPropertyFilterJSON `json:"-"`
}

// v1UsageSearchResponseMatchedBillableMetricsPropertyFilterJSON contains the JSON
// metadata for the struct
// [V1UsageSearchResponseMatchedBillableMetricsPropertyFilter]
type v1UsageSearchResponseMatchedBillableMetricsPropertyFilterJSON struct {
	Name        apijson.Field
	Exists      apijson.Field
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1UsageSearchResponseMatchedBillableMetricsPropertyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1UsageSearchResponseMatchedBillableMetricsPropertyFilterJSON) RawJSON() string {
	return r.raw
}

// The customer the event was matched to if a match was found
type V1UsageSearchResponseMatchedCustomer struct {
	ID   string                                   `json:"id" format:"uuid"`
	Name string                                   `json:"name"`
	JSON v1UsageSearchResponseMatchedCustomerJSON `json:"-"`
}

// v1UsageSearchResponseMatchedCustomerJSON contains the JSON metadata for the
// struct [V1UsageSearchResponseMatchedCustomer]
type v1UsageSearchResponseMatchedCustomerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1UsageSearchResponseMatchedCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1UsageSearchResponseMatchedCustomerJSON) RawJSON() string {
	return r.raw
}

type V1UsageListParams struct {
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	StartingOn   param.Field[time.Time] `json:"starting_on,required" format:"date-time"`
	// A window_size of "day" or "hour" will return the usage for the specified period
	// segmented into daily or hourly aggregates. A window_size of "none" will return a
	// single usage aggregate for the entirety of the specified period.
	WindowSize param.Field[V1UsageListParamsWindowSize] `json:"window_size,required"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// A list of billable metrics to fetch usage for. If absent, all billable metrics
	// will be returned.
	BillableMetrics param.Field[[]V1UsageListParamsBillableMetric] `json:"billable_metrics"`
	// A list of Metronome customer IDs to fetch usage for. If absent, usage for all
	// customers will be returned.
	CustomerIDs param.Field[[]string] `json:"customer_ids" format:"uuid"`
}

func (r V1UsageListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V1UsageListParams]'s query parameters as `url.Values`.
func (r V1UsageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A window_size of "day" or "hour" will return the usage for the specified period
// segmented into daily or hourly aggregates. A window_size of "none" will return a
// single usage aggregate for the entirety of the specified period.
type V1UsageListParamsWindowSize string

const (
	V1UsageListParamsWindowSizeHour V1UsageListParamsWindowSize = "HOUR"
	V1UsageListParamsWindowSizeDay  V1UsageListParamsWindowSize = "DAY"
	V1UsageListParamsWindowSizeNone V1UsageListParamsWindowSize = "NONE"
)

func (r V1UsageListParamsWindowSize) IsKnown() bool {
	switch r {
	case V1UsageListParamsWindowSizeHour, V1UsageListParamsWindowSizeDay, V1UsageListParamsWindowSizeNone:
		return true
	}
	return false
}

type V1UsageListParamsBillableMetric struct {
	ID      param.Field[string]                                  `json:"id,required" format:"uuid"`
	GroupBy param.Field[V1UsageListParamsBillableMetricsGroupBy] `json:"group_by"`
}

func (r V1UsageListParamsBillableMetric) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1UsageListParamsBillableMetricsGroupBy struct {
	// The name of the group_by key to use
	Key param.Field[string] `json:"key,required"`
	// Values of the group_by key to return in the query. If this field is omitted, all
	// available values will be returned, up to a maximum of 200.
	Values param.Field[[]string] `json:"values"`
}

func (r V1UsageListParamsBillableMetricsGroupBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1UsageIngestParams struct {
	Usage []V1UsageIngestParamsUsage `json:"usage"`
}

func (r V1UsageIngestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Usage)
}

type V1UsageIngestParamsUsage struct {
	CustomerID param.Field[string] `json:"customer_id,required"`
	EventType  param.Field[string] `json:"event_type,required"`
	// RFC 3339 formatted
	Timestamp     param.Field[string]                 `json:"timestamp,required"`
	TransactionID param.Field[string]                 `json:"transaction_id,required"`
	Properties    param.Field[map[string]interface{}] `json:"properties"`
}

func (r V1UsageIngestParamsUsage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1UsageListWithGroupsParams struct {
	BillableMetricID param.Field[string] `json:"billable_metric_id,required" format:"uuid"`
	CustomerID       param.Field[string] `json:"customer_id,required" format:"uuid"`
	// A window_size of "day" or "hour" will return the usage for the specified period
	// segmented into daily or hourly aggregates. A window_size of "none" will return a
	// single usage aggregate for the entirety of the specified period.
	WindowSize param.Field[V1UsageListWithGroupsParamsWindowSize] `json:"window_size,required"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// If true, will return the usage for the current billing period. Will return an
	// error if the customer is currently uncontracted or starting_on and ending_before
	// are specified when this is true.
	CurrentPeriod param.Field[bool]                               `json:"current_period"`
	EndingBefore  param.Field[time.Time]                          `json:"ending_before" format:"date-time"`
	GroupBy       param.Field[V1UsageListWithGroupsParamsGroupBy] `json:"group_by"`
	StartingOn    param.Field[time.Time]                          `json:"starting_on" format:"date-time"`
}

func (r V1UsageListWithGroupsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V1UsageListWithGroupsParams]'s query parameters as
// `url.Values`.
func (r V1UsageListWithGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A window_size of "day" or "hour" will return the usage for the specified period
// segmented into daily or hourly aggregates. A window_size of "none" will return a
// single usage aggregate for the entirety of the specified period.
type V1UsageListWithGroupsParamsWindowSize string

const (
	V1UsageListWithGroupsParamsWindowSizeHour V1UsageListWithGroupsParamsWindowSize = "HOUR"
	V1UsageListWithGroupsParamsWindowSizeDay  V1UsageListWithGroupsParamsWindowSize = "DAY"
	V1UsageListWithGroupsParamsWindowSizeNone V1UsageListWithGroupsParamsWindowSize = "NONE"
)

func (r V1UsageListWithGroupsParamsWindowSize) IsKnown() bool {
	switch r {
	case V1UsageListWithGroupsParamsWindowSizeHour, V1UsageListWithGroupsParamsWindowSizeDay, V1UsageListWithGroupsParamsWindowSizeNone:
		return true
	}
	return false
}

type V1UsageListWithGroupsParamsGroupBy struct {
	// The name of the group_by key to use
	Key param.Field[string] `json:"key,required"`
	// Values of the group_by key to return in the query. Omit this if you'd like all
	// values for the key returned.
	Values param.Field[[]string] `json:"values"`
}

func (r V1UsageListWithGroupsParamsGroupBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1UsageSearchParams struct {
	// The transaction IDs of the events to retrieve
	TransactionIDs param.Field[[]string] `json:"transactionIds,required"`
}

func (r V1UsageSearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
