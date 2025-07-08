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
	"github.com/Metronome-Industries/metronome-go/shared"
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

// Fetch aggregated usage data for multiple customers and billable-metrics, broken
// into intervals of the specified length.
func (r *V1UsageService) List(ctx context.Context, params V1UsageListParams, opts ...option.RequestOption) (res *V1UsageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Send usage events to Metronome. The body of this request is expected to be a
// JSON array of between 1 and 100 usage events. Compressed request bodies are
// supported with a `Content-Encoding: gzip` header. See
// [Getting usage into Metronome](https://docs.metronome.com/connect-metronome/) to
// learn more about usage events.
func (r *V1UsageService) Ingest(ctx context.Context, body V1UsageIngestParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/ingest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Fetch aggregated usage data for the specified customer, billable-metric, and
// optional group, broken into intervals of the specified length.
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

// Fetch aggregated usage data for the specified customer, billable-metric, and
// optional group, broken into intervals of the specified length.
func (r *V1UsageService) ListWithGroupsAutoPaging(ctx context.Context, params V1UsageListWithGroupsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1UsageListWithGroupsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListWithGroups(ctx, params, opts...))
}

// Find events to match to customers, billable metrics, etc. We only look for
// transactions that occurred in the last 34 days.
func (r *V1UsageService) Search(ctx context.Context, body V1UsageSearchParams, opts ...option.RequestOption) (res *[]V1UsageSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/events/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1UsageListResponse struct {
	Data     []V1UsageListResponseData `json:"data,required"`
	NextPage string                    `json:"next_page,required,nullable"`
	JSON     v1UsageListResponseJSON   `json:"-"`
}

// v1UsageListResponseJSON contains the JSON metadata for the struct
// [V1UsageListResponse]
type v1UsageListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1UsageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1UsageListResponseJSON) RawJSON() string {
	return r.raw
}

type V1UsageListResponseData struct {
	BillableMetricID   string    `json:"billable_metric_id,required" format:"uuid"`
	BillableMetricName string    `json:"billable_metric_name,required"`
	CustomerID         string    `json:"customer_id,required" format:"uuid"`
	EndTimestamp       time.Time `json:"end_timestamp,required" format:"date-time"`
	StartTimestamp     time.Time `json:"start_timestamp,required" format:"date-time"`
	Value              float64   `json:"value,required,nullable"`
	// Values will be either a number or null. Null indicates that there were no
	// matches for the group_by value.
	Groups map[string]float64          `json:"groups"`
	JSON   v1UsageListResponseDataJSON `json:"-"`
}

// v1UsageListResponseDataJSON contains the JSON metadata for the struct
// [V1UsageListResponseData]
type v1UsageListResponseDataJSON struct {
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

func (r *V1UsageListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1UsageListResponseDataJSON) RawJSON() string {
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
	ArchivedAt   time.Time         `json:"archived_at" format:"date-time"`
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
