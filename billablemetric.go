// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/pagination"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// BillableMetricService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillableMetricService] method instead.
type BillableMetricService struct {
	Options []option.RequestOption
}

// NewBillableMetricService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBillableMetricService(opts ...option.RequestOption) (r *BillableMetricService) {
	r = &BillableMetricService{}
	r.Options = opts
	return
}

// Creates a new Billable Metric.
func (r *BillableMetricService) New(ctx context.Context, body BillableMetricNewParams, opts ...option.RequestOption) (res *BillableMetricNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "billable-metrics/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a billable metric.
func (r *BillableMetricService) Get(ctx context.Context, billableMetricID string, opts ...option.RequestOption) (res *BillableMetricGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if billableMetricID == "" {
		err = errors.New("missing required billable_metric_id parameter")
		return
	}
	path := fmt.Sprintf("billable-metrics/%s", billableMetricID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all billable metrics for a given customer.
func (r *BillableMetricService) List(ctx context.Context, customerID string, query BillableMetricListParams, opts ...option.RequestOption) (res *pagination.CursorPage[BillableMetricListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
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
func (r *BillableMetricService) ListAutoPaging(ctx context.Context, customerID string, query BillableMetricListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[BillableMetricListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, customerID, query, opts...))
}

// Archive an existing billable metric.
func (r *BillableMetricService) Archive(ctx context.Context, body BillableMetricArchiveParams, opts ...option.RequestOption) (res *BillableMetricArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "billable-metrics/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BillableMetricNewResponse struct {
	Data shared.ID                     `json:"data,required"`
	JSON billableMetricNewResponseJSON `json:"-"`
}

// billableMetricNewResponseJSON contains the JSON metadata for the struct
// [BillableMetricNewResponse]
type billableMetricNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricNewResponseJSON) RawJSON() string {
	return r.raw
}

type BillableMetricGetResponse struct {
	Data BillableMetricGetResponseData `json:"data,required"`
	JSON billableMetricGetResponseJSON `json:"-"`
}

// billableMetricGetResponseJSON contains the JSON metadata for the struct
// [BillableMetricGetResponse]
type billableMetricGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricGetResponseJSON) RawJSON() string {
	return r.raw
}

type BillableMetricGetResponseData struct {
	// ID of the billable metric
	ID string `json:"id,required" format:"uuid"`
	// The display name of the billable metric.
	Name string `json:"name,required"`
	// A key that specifies which property of the event is used to aggregate data. This
	// key must be one of the property filter names and is not applicable when the
	// aggregation type is 'count'.
	AggregationKey string `json:"aggregation_key"`
	// Specifies the type of aggregation performed on matching events.
	AggregationType BillableMetricGetResponseDataAggregationType `json:"aggregation_type"`
	CustomFields    map[string]string                            `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter BillableMetricGetResponseDataEventTypeFilter `json:"event_type_filter"`
	// Property names that are used to group usage costs on an invoice. Each entry
	// represents a set of properties used to slice events into distinct buckets.
	GroupKeys [][]string `json:"group_keys"`
	// A list of filters to match events to this billable metric. Each filter defines a
	// rule on an event property. All rules must pass for the event to match the
	// billable metric.
	PropertyFilters []BillableMetricGetResponseDataPropertyFilter `json:"property_filters"`
	// The SQL query associated with the billable metric
	Sql  string                            `json:"sql"`
	JSON billableMetricGetResponseDataJSON `json:"-"`
}

// billableMetricGetResponseDataJSON contains the JSON metadata for the struct
// [BillableMetricGetResponseData]
type billableMetricGetResponseDataJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	AggregationKey  apijson.Field
	AggregationType apijson.Field
	CustomFields    apijson.Field
	EventTypeFilter apijson.Field
	GroupKeys       apijson.Field
	PropertyFilters apijson.Field
	Sql             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BillableMetricGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of aggregation performed on matching events.
type BillableMetricGetResponseDataAggregationType string

const (
	BillableMetricGetResponseDataAggregationTypeCount  BillableMetricGetResponseDataAggregationType = "count"
	BillableMetricGetResponseDataAggregationTypeCount  BillableMetricGetResponseDataAggregationType = "Count"
	BillableMetricGetResponseDataAggregationTypeCount  BillableMetricGetResponseDataAggregationType = "COUNT"
	BillableMetricGetResponseDataAggregationTypeLatest BillableMetricGetResponseDataAggregationType = "latest"
	BillableMetricGetResponseDataAggregationTypeLatest BillableMetricGetResponseDataAggregationType = "Latest"
	BillableMetricGetResponseDataAggregationTypeLatest BillableMetricGetResponseDataAggregationType = "LATEST"
	BillableMetricGetResponseDataAggregationTypeMax    BillableMetricGetResponseDataAggregationType = "max"
	BillableMetricGetResponseDataAggregationTypeMax    BillableMetricGetResponseDataAggregationType = "Max"
	BillableMetricGetResponseDataAggregationTypeMax    BillableMetricGetResponseDataAggregationType = "MAX"
	BillableMetricGetResponseDataAggregationTypeSum    BillableMetricGetResponseDataAggregationType = "sum"
	BillableMetricGetResponseDataAggregationTypeSum    BillableMetricGetResponseDataAggregationType = "Sum"
	BillableMetricGetResponseDataAggregationTypeSum    BillableMetricGetResponseDataAggregationType = "SUM"
	BillableMetricGetResponseDataAggregationTypeUnique BillableMetricGetResponseDataAggregationType = "unique"
	BillableMetricGetResponseDataAggregationTypeUnique BillableMetricGetResponseDataAggregationType = "Unique"
	BillableMetricGetResponseDataAggregationTypeUnique BillableMetricGetResponseDataAggregationType = "UNIQUE"
)

func (r BillableMetricGetResponseDataAggregationType) IsKnown() bool {
	switch r {
	case BillableMetricGetResponseDataAggregationTypeCount, BillableMetricGetResponseDataAggregationTypeCount, BillableMetricGetResponseDataAggregationTypeCount, BillableMetricGetResponseDataAggregationTypeLatest, BillableMetricGetResponseDataAggregationTypeLatest, BillableMetricGetResponseDataAggregationTypeLatest, BillableMetricGetResponseDataAggregationTypeMax, BillableMetricGetResponseDataAggregationTypeMax, BillableMetricGetResponseDataAggregationTypeMax, BillableMetricGetResponseDataAggregationTypeSum, BillableMetricGetResponseDataAggregationTypeSum, BillableMetricGetResponseDataAggregationTypeSum, BillableMetricGetResponseDataAggregationTypeUnique, BillableMetricGetResponseDataAggregationTypeUnique, BillableMetricGetResponseDataAggregationTypeUnique:
		return true
	}
	return false
}

// An optional filtering rule to match the 'event_type' property of an event.
type BillableMetricGetResponseDataEventTypeFilter struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues []string `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues []string                                         `json:"not_in_values"`
	JSON        billableMetricGetResponseDataEventTypeFilterJSON `json:"-"`
}

// billableMetricGetResponseDataEventTypeFilterJSON contains the JSON metadata for
// the struct [BillableMetricGetResponseDataEventTypeFilter]
type billableMetricGetResponseDataEventTypeFilterJSON struct {
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricGetResponseDataEventTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricGetResponseDataEventTypeFilterJSON) RawJSON() string {
	return r.raw
}

type BillableMetricGetResponseDataPropertyFilter struct {
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
	NotInValues []string                                        `json:"not_in_values"`
	JSON        billableMetricGetResponseDataPropertyFilterJSON `json:"-"`
}

// billableMetricGetResponseDataPropertyFilterJSON contains the JSON metadata for
// the struct [BillableMetricGetResponseDataPropertyFilter]
type billableMetricGetResponseDataPropertyFilterJSON struct {
	Name        apijson.Field
	Exists      apijson.Field
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricGetResponseDataPropertyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricGetResponseDataPropertyFilterJSON) RawJSON() string {
	return r.raw
}

type BillableMetricListResponse struct {
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
	AggregationType BillableMetricListResponseAggregationType `json:"aggregation_type"`
	CustomFields    map[string]string                         `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter BillableMetricListResponseEventTypeFilter `json:"event_type_filter"`
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
	PropertyFilters []BillableMetricListResponsePropertyFilter `json:"property_filters"`
	// The SQL query associated with the billable metric
	Sql  string                         `json:"sql"`
	JSON billableMetricListResponseJSON `json:"-"`
}

// billableMetricListResponseJSON contains the JSON metadata for the struct
// [BillableMetricListResponse]
type billableMetricListResponseJSON struct {
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
	Sql             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BillableMetricListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricListResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of aggregation performed on matching events.
type BillableMetricListResponseAggregationType string

const (
	BillableMetricListResponseAggregationTypeCount  BillableMetricListResponseAggregationType = "count"
	BillableMetricListResponseAggregationTypeCount  BillableMetricListResponseAggregationType = "Count"
	BillableMetricListResponseAggregationTypeCount  BillableMetricListResponseAggregationType = "COUNT"
	BillableMetricListResponseAggregationTypeLatest BillableMetricListResponseAggregationType = "latest"
	BillableMetricListResponseAggregationTypeLatest BillableMetricListResponseAggregationType = "Latest"
	BillableMetricListResponseAggregationTypeLatest BillableMetricListResponseAggregationType = "LATEST"
	BillableMetricListResponseAggregationTypeMax    BillableMetricListResponseAggregationType = "max"
	BillableMetricListResponseAggregationTypeMax    BillableMetricListResponseAggregationType = "Max"
	BillableMetricListResponseAggregationTypeMax    BillableMetricListResponseAggregationType = "MAX"
	BillableMetricListResponseAggregationTypeSum    BillableMetricListResponseAggregationType = "sum"
	BillableMetricListResponseAggregationTypeSum    BillableMetricListResponseAggregationType = "Sum"
	BillableMetricListResponseAggregationTypeSum    BillableMetricListResponseAggregationType = "SUM"
	BillableMetricListResponseAggregationTypeUnique BillableMetricListResponseAggregationType = "unique"
	BillableMetricListResponseAggregationTypeUnique BillableMetricListResponseAggregationType = "Unique"
	BillableMetricListResponseAggregationTypeUnique BillableMetricListResponseAggregationType = "UNIQUE"
)

func (r BillableMetricListResponseAggregationType) IsKnown() bool {
	switch r {
	case BillableMetricListResponseAggregationTypeCount, BillableMetricListResponseAggregationTypeCount, BillableMetricListResponseAggregationTypeCount, BillableMetricListResponseAggregationTypeLatest, BillableMetricListResponseAggregationTypeLatest, BillableMetricListResponseAggregationTypeLatest, BillableMetricListResponseAggregationTypeMax, BillableMetricListResponseAggregationTypeMax, BillableMetricListResponseAggregationTypeMax, BillableMetricListResponseAggregationTypeSum, BillableMetricListResponseAggregationTypeSum, BillableMetricListResponseAggregationTypeSum, BillableMetricListResponseAggregationTypeUnique, BillableMetricListResponseAggregationTypeUnique, BillableMetricListResponseAggregationTypeUnique:
		return true
	}
	return false
}

// An optional filtering rule to match the 'event_type' property of an event.
type BillableMetricListResponseEventTypeFilter struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues []string `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues []string                                      `json:"not_in_values"`
	JSON        billableMetricListResponseEventTypeFilterJSON `json:"-"`
}

// billableMetricListResponseEventTypeFilterJSON contains the JSON metadata for the
// struct [BillableMetricListResponseEventTypeFilter]
type billableMetricListResponseEventTypeFilterJSON struct {
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricListResponseEventTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricListResponseEventTypeFilterJSON) RawJSON() string {
	return r.raw
}

type BillableMetricListResponsePropertyFilter struct {
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
	NotInValues []string                                     `json:"not_in_values"`
	JSON        billableMetricListResponsePropertyFilterJSON `json:"-"`
}

// billableMetricListResponsePropertyFilterJSON contains the JSON metadata for the
// struct [BillableMetricListResponsePropertyFilter]
type billableMetricListResponsePropertyFilterJSON struct {
	Name        apijson.Field
	Exists      apijson.Field
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricListResponsePropertyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricListResponsePropertyFilterJSON) RawJSON() string {
	return r.raw
}

type BillableMetricArchiveResponse struct {
	Data shared.ID                         `json:"data,required"`
	JSON billableMetricArchiveResponseJSON `json:"-"`
}

// billableMetricArchiveResponseJSON contains the JSON metadata for the struct
// [BillableMetricArchiveResponse]
type billableMetricArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type BillableMetricNewParams struct {
	// Specifies the type of aggregation performed on matching events.
	AggregationType param.Field[BillableMetricNewParamsAggregationType] `json:"aggregation_type,required"`
	// The display name of the billable metric.
	Name param.Field[string] `json:"name,required"`
	// A key that specifies which property of the event is used to aggregate data. This
	// key must be one of the property filter names and is not applicable when the
	// aggregation type is 'count'.
	AggregationKey param.Field[string] `json:"aggregation_key"`
	// Custom fields to attach to the billable metric.
	CustomFields param.Field[map[string]string] `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter param.Field[BillableMetricNewParamsEventTypeFilter] `json:"event_type_filter"`
	// Property names that are used to group usage costs on an invoice. Each entry
	// represents a set of properties used to slice events into distinct buckets.
	GroupKeys param.Field[[][]string] `json:"group_keys"`
	// A list of filters to match events to this billable metric. Each filter defines a
	// rule on an event property. All rules must pass for the event to match the
	// billable metric.
	PropertyFilters param.Field[[]BillableMetricNewParamsPropertyFilter] `json:"property_filters"`
}

func (r BillableMetricNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of aggregation performed on matching events.
type BillableMetricNewParamsAggregationType string

const (
	BillableMetricNewParamsAggregationTypeCount  BillableMetricNewParamsAggregationType = "count"
	BillableMetricNewParamsAggregationTypeCount  BillableMetricNewParamsAggregationType = "Count"
	BillableMetricNewParamsAggregationTypeCount  BillableMetricNewParamsAggregationType = "COUNT"
	BillableMetricNewParamsAggregationTypeLatest BillableMetricNewParamsAggregationType = "latest"
	BillableMetricNewParamsAggregationTypeLatest BillableMetricNewParamsAggregationType = "Latest"
	BillableMetricNewParamsAggregationTypeLatest BillableMetricNewParamsAggregationType = "LATEST"
	BillableMetricNewParamsAggregationTypeMax    BillableMetricNewParamsAggregationType = "max"
	BillableMetricNewParamsAggregationTypeMax    BillableMetricNewParamsAggregationType = "Max"
	BillableMetricNewParamsAggregationTypeMax    BillableMetricNewParamsAggregationType = "MAX"
	BillableMetricNewParamsAggregationTypeSum    BillableMetricNewParamsAggregationType = "sum"
	BillableMetricNewParamsAggregationTypeSum    BillableMetricNewParamsAggregationType = "Sum"
	BillableMetricNewParamsAggregationTypeSum    BillableMetricNewParamsAggregationType = "SUM"
	BillableMetricNewParamsAggregationTypeUnique BillableMetricNewParamsAggregationType = "unique"
	BillableMetricNewParamsAggregationTypeUnique BillableMetricNewParamsAggregationType = "Unique"
	BillableMetricNewParamsAggregationTypeUnique BillableMetricNewParamsAggregationType = "UNIQUE"
)

func (r BillableMetricNewParamsAggregationType) IsKnown() bool {
	switch r {
	case BillableMetricNewParamsAggregationTypeCount, BillableMetricNewParamsAggregationTypeCount, BillableMetricNewParamsAggregationTypeCount, BillableMetricNewParamsAggregationTypeLatest, BillableMetricNewParamsAggregationTypeLatest, BillableMetricNewParamsAggregationTypeLatest, BillableMetricNewParamsAggregationTypeMax, BillableMetricNewParamsAggregationTypeMax, BillableMetricNewParamsAggregationTypeMax, BillableMetricNewParamsAggregationTypeSum, BillableMetricNewParamsAggregationTypeSum, BillableMetricNewParamsAggregationTypeSum, BillableMetricNewParamsAggregationTypeUnique, BillableMetricNewParamsAggregationTypeUnique, BillableMetricNewParamsAggregationTypeUnique:
		return true
	}
	return false
}

// An optional filtering rule to match the 'event_type' property of an event.
type BillableMetricNewParamsEventTypeFilter struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues param.Field[[]string] `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues param.Field[[]string] `json:"not_in_values"`
}

func (r BillableMetricNewParamsEventTypeFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillableMetricNewParamsPropertyFilter struct {
	// The name of the event property.
	Name param.Field[string] `json:"name,required"`
	// Determines whether the property must exist in the event. If true, only events
	// with this property will pass the filter. If false, only events without this
	// property will pass the filter. If null or omitted, the existence of the property
	// is optional.
	Exists param.Field[bool] `json:"exists"`
	// Specifies the allowed values for the property to match an event. An event will
	// pass the filter only if its property value is included in this list. If
	// undefined, all property values will pass the filter. Must be non-empty if
	// present.
	InValues param.Field[[]string] `json:"in_values"`
	// Specifies the values that prevent an event from matching the filter. An event
	// will not pass the filter if its property value is included in this list. If null
	// or empty, all property values will pass the filter. Must be non-empty if
	// present.
	NotInValues param.Field[[]string] `json:"not_in_values"`
}

func (r BillableMetricNewParamsPropertyFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillableMetricListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// If true, the list of metrics will be filtered to just ones that are on the
	// customer's current plan
	OnCurrentPlan param.Field[bool] `query:"on_current_plan"`
}

// URLQuery serializes [BillableMetricListParams]'s query parameters as
// `url.Values`.
func (r BillableMetricListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillableMetricArchiveParams struct {
	ID shared.IDParam `json:"id,required"`
}

func (r BillableMetricArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ID)
}
