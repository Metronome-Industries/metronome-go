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
)

// V1BillableMetricService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BillableMetricService] method instead.
type V1BillableMetricService struct {
	Options []option.RequestOption
}

// NewV1BillableMetricService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1BillableMetricService(opts ...option.RequestOption) (r *V1BillableMetricService) {
	r = &V1BillableMetricService{}
	r.Options = opts
	return
}

// Creates a new Billable Metric.
func (r *V1BillableMetricService) New(ctx context.Context, body V1BillableMetricNewParams, opts ...option.RequestOption) (res *V1BillableMetricNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/billable-metrics/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a billable metric.
func (r *V1BillableMetricService) Get(ctx context.Context, query V1BillableMetricGetParams, opts ...option.RequestOption) (res *V1BillableMetricGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.BillableMetricID.Value == "" {
		err = errors.New("missing required billable_metric_id parameter")
		return
	}
	path := fmt.Sprintf("v1/billable-metrics/%s", query.BillableMetricID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all billable metrics.
func (r *V1BillableMetricService) List(ctx context.Context, query V1BillableMetricListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1BillableMetricListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/billable-metrics"
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

// List all billable metrics.
func (r *V1BillableMetricService) ListAutoPaging(ctx context.Context, query V1BillableMetricListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1BillableMetricListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Archive an existing billable metric.
func (r *V1BillableMetricService) Archive(ctx context.Context, body V1BillableMetricArchiveParams, opts ...option.RequestOption) (res *V1BillableMetricArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/billable-metrics/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1BillableMetricNewResponse struct {
	Data V1BillableMetricNewResponseData `json:"data,required"`
	JSON v1BillableMetricNewResponseJSON `json:"-"`
}

// v1BillableMetricNewResponseJSON contains the JSON metadata for the struct
// [V1BillableMetricNewResponse]
type v1BillableMetricNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1BillableMetricNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1BillableMetricNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1BillableMetricNewResponseData struct {
	ID   string                              `json:"id,required" format:"uuid"`
	JSON v1BillableMetricNewResponseDataJSON `json:"-"`
}

// v1BillableMetricNewResponseDataJSON contains the JSON metadata for the struct
// [V1BillableMetricNewResponseData]
type v1BillableMetricNewResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1BillableMetricNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1BillableMetricNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1BillableMetricGetResponse struct {
	Data V1BillableMetricGetResponseData `json:"data,required"`
	JSON v1BillableMetricGetResponseJSON `json:"-"`
}

// v1BillableMetricGetResponseJSON contains the JSON metadata for the struct
// [V1BillableMetricGetResponse]
type v1BillableMetricGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1BillableMetricGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1BillableMetricGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1BillableMetricGetResponseData struct {
	// ID of the billable metric
	ID string `json:"id,required" format:"uuid"`
	// The display name of the billable metric.
	Name string `json:"name,required"`
	// A key that specifies which property of the event is used to aggregate data. This
	// key must be one of the property filter names and is not applicable when the
	// aggregation type is 'count'.
	AggregationKey string `json:"aggregation_key"`
	// Specifies the type of aggregation performed on matching events.
	AggregationType V1BillableMetricGetResponseDataAggregationType `json:"aggregation_type"`
	// RFC 3339 timestamp indicating when the billable metric was archived. If not
	// provided, the billable metric is not archived.
	ArchivedAt   time.Time         `json:"archived_at" format:"date-time"`
	CustomFields map[string]string `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter V1BillableMetricGetResponseDataEventTypeFilter `json:"event_type_filter"`
	// Property names that are used to group usage costs on an invoice. Each entry
	// represents a set of properties used to slice events into distinct buckets.
	GroupKeys [][]string `json:"group_keys"`
	// A list of filters to match events to this billable metric. Each filter defines a
	// rule on an event property. All rules must pass for the event to match the
	// billable metric.
	PropertyFilters []V1BillableMetricGetResponseDataPropertyFilter `json:"property_filters"`
	// The SQL query associated with the billable metric
	Sql  string                              `json:"sql"`
	JSON v1BillableMetricGetResponseDataJSON `json:"-"`
}

// v1BillableMetricGetResponseDataJSON contains the JSON metadata for the struct
// [V1BillableMetricGetResponseData]
type v1BillableMetricGetResponseDataJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	AggregationKey  apijson.Field
	AggregationType apijson.Field
	ArchivedAt      apijson.Field
	CustomFields    apijson.Field
	EventTypeFilter apijson.Field
	GroupKeys       apijson.Field
	PropertyFilters apijson.Field
	Sql             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V1BillableMetricGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1BillableMetricGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of aggregation performed on matching events.
type V1BillableMetricGetResponseDataAggregationType string

const (
	V1BillableMetricGetResponseDataAggregationTypeCount  V1BillableMetricGetResponseDataAggregationType = "COUNT"
	V1BillableMetricGetResponseDataAggregationTypeLatest V1BillableMetricGetResponseDataAggregationType = "LATEST"
	V1BillableMetricGetResponseDataAggregationTypeMax    V1BillableMetricGetResponseDataAggregationType = "MAX"
	V1BillableMetricGetResponseDataAggregationTypeSum    V1BillableMetricGetResponseDataAggregationType = "SUM"
	V1BillableMetricGetResponseDataAggregationTypeUnique V1BillableMetricGetResponseDataAggregationType = "UNIQUE"
)

func (r V1BillableMetricGetResponseDataAggregationType) IsKnown() bool {
	switch r {
	case V1BillableMetricGetResponseDataAggregationTypeCount, V1BillableMetricGetResponseDataAggregationTypeLatest, V1BillableMetricGetResponseDataAggregationTypeMax, V1BillableMetricGetResponseDataAggregationTypeSum, V1BillableMetricGetResponseDataAggregationTypeUnique:
		return true
	}
	return false
}

// An optional filtering rule to match the 'event_type' property of an event.
type V1BillableMetricGetResponseDataEventTypeFilter struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues []string `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues []string                                           `json:"not_in_values"`
	JSON        v1BillableMetricGetResponseDataEventTypeFilterJSON `json:"-"`
}

// v1BillableMetricGetResponseDataEventTypeFilterJSON contains the JSON metadata
// for the struct [V1BillableMetricGetResponseDataEventTypeFilter]
type v1BillableMetricGetResponseDataEventTypeFilterJSON struct {
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1BillableMetricGetResponseDataEventTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1BillableMetricGetResponseDataEventTypeFilterJSON) RawJSON() string {
	return r.raw
}

type V1BillableMetricGetResponseDataPropertyFilter struct {
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
	NotInValues []string                                          `json:"not_in_values"`
	JSON        v1BillableMetricGetResponseDataPropertyFilterJSON `json:"-"`
}

// v1BillableMetricGetResponseDataPropertyFilterJSON contains the JSON metadata for
// the struct [V1BillableMetricGetResponseDataPropertyFilter]
type v1BillableMetricGetResponseDataPropertyFilterJSON struct {
	Name        apijson.Field
	Exists      apijson.Field
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1BillableMetricGetResponseDataPropertyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1BillableMetricGetResponseDataPropertyFilterJSON) RawJSON() string {
	return r.raw
}

type V1BillableMetricListResponse struct {
	// ID of the billable metric
	ID string `json:"id,required" format:"uuid"`
	// The display name of the billable metric.
	Name string `json:"name,required"`
	// A key that specifies which property of the event is used to aggregate data. This
	// key must be one of the property filter names and is not applicable when the
	// aggregation type is 'count'.
	AggregationKey string `json:"aggregation_key"`
	// Specifies the type of aggregation performed on matching events.
	AggregationType V1BillableMetricListResponseAggregationType `json:"aggregation_type"`
	// RFC 3339 timestamp indicating when the billable metric was archived. If not
	// provided, the billable metric is not archived.
	ArchivedAt   time.Time         `json:"archived_at" format:"date-time"`
	CustomFields map[string]string `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter V1BillableMetricListResponseEventTypeFilter `json:"event_type_filter"`
	// Property names that are used to group usage costs on an invoice. Each entry
	// represents a set of properties used to slice events into distinct buckets.
	GroupKeys [][]string `json:"group_keys"`
	// A list of filters to match events to this billable metric. Each filter defines a
	// rule on an event property. All rules must pass for the event to match the
	// billable metric.
	PropertyFilters []V1BillableMetricListResponsePropertyFilter `json:"property_filters"`
	// The SQL query associated with the billable metric
	Sql  string                           `json:"sql"`
	JSON v1BillableMetricListResponseJSON `json:"-"`
}

// v1BillableMetricListResponseJSON contains the JSON metadata for the struct
// [V1BillableMetricListResponse]
type v1BillableMetricListResponseJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	AggregationKey  apijson.Field
	AggregationType apijson.Field
	ArchivedAt      apijson.Field
	CustomFields    apijson.Field
	EventTypeFilter apijson.Field
	GroupKeys       apijson.Field
	PropertyFilters apijson.Field
	Sql             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V1BillableMetricListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1BillableMetricListResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of aggregation performed on matching events.
type V1BillableMetricListResponseAggregationType string

const (
	V1BillableMetricListResponseAggregationTypeCount  V1BillableMetricListResponseAggregationType = "COUNT"
	V1BillableMetricListResponseAggregationTypeLatest V1BillableMetricListResponseAggregationType = "LATEST"
	V1BillableMetricListResponseAggregationTypeMax    V1BillableMetricListResponseAggregationType = "MAX"
	V1BillableMetricListResponseAggregationTypeSum    V1BillableMetricListResponseAggregationType = "SUM"
	V1BillableMetricListResponseAggregationTypeUnique V1BillableMetricListResponseAggregationType = "UNIQUE"
)

func (r V1BillableMetricListResponseAggregationType) IsKnown() bool {
	switch r {
	case V1BillableMetricListResponseAggregationTypeCount, V1BillableMetricListResponseAggregationTypeLatest, V1BillableMetricListResponseAggregationTypeMax, V1BillableMetricListResponseAggregationTypeSum, V1BillableMetricListResponseAggregationTypeUnique:
		return true
	}
	return false
}

// An optional filtering rule to match the 'event_type' property of an event.
type V1BillableMetricListResponseEventTypeFilter struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues []string `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues []string                                        `json:"not_in_values"`
	JSON        v1BillableMetricListResponseEventTypeFilterJSON `json:"-"`
}

// v1BillableMetricListResponseEventTypeFilterJSON contains the JSON metadata for
// the struct [V1BillableMetricListResponseEventTypeFilter]
type v1BillableMetricListResponseEventTypeFilterJSON struct {
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1BillableMetricListResponseEventTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1BillableMetricListResponseEventTypeFilterJSON) RawJSON() string {
	return r.raw
}

type V1BillableMetricListResponsePropertyFilter struct {
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
	NotInValues []string                                       `json:"not_in_values"`
	JSON        v1BillableMetricListResponsePropertyFilterJSON `json:"-"`
}

// v1BillableMetricListResponsePropertyFilterJSON contains the JSON metadata for
// the struct [V1BillableMetricListResponsePropertyFilter]
type v1BillableMetricListResponsePropertyFilterJSON struct {
	Name        apijson.Field
	Exists      apijson.Field
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1BillableMetricListResponsePropertyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1BillableMetricListResponsePropertyFilterJSON) RawJSON() string {
	return r.raw
}

type V1BillableMetricArchiveResponse struct {
	Data V1BillableMetricArchiveResponseData `json:"data,required"`
	JSON v1BillableMetricArchiveResponseJSON `json:"-"`
}

// v1BillableMetricArchiveResponseJSON contains the JSON metadata for the struct
// [V1BillableMetricArchiveResponse]
type v1BillableMetricArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1BillableMetricArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1BillableMetricArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type V1BillableMetricArchiveResponseData struct {
	ID   string                                  `json:"id,required" format:"uuid"`
	JSON v1BillableMetricArchiveResponseDataJSON `json:"-"`
}

// v1BillableMetricArchiveResponseDataJSON contains the JSON metadata for the
// struct [V1BillableMetricArchiveResponseData]
type v1BillableMetricArchiveResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1BillableMetricArchiveResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1BillableMetricArchiveResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1BillableMetricNewParams struct {
	// The display name of the billable metric.
	Name param.Field[string] `json:"name,required"`
	// Specifies the type of aggregation performed on matching events. Required if
	// `sql` is not provided.
	AggregationKey param.Field[string] `json:"aggregation_key"`
	// Specifies the type of aggregation performed on matching events.
	AggregationType param.Field[V1BillableMetricNewParamsAggregationType] `json:"aggregation_type"`
	// Custom fields to attach to the billable metric.
	CustomFields param.Field[map[string]string] `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter param.Field[V1BillableMetricNewParamsEventTypeFilter] `json:"event_type_filter"`
	// Property names that are used to group usage costs on an invoice. Each entry
	// represents a set of properties used to slice events into distinct buckets.
	GroupKeys param.Field[[][]string] `json:"group_keys"`
	// A list of filters to match events to this billable metric. Each filter defines a
	// rule on an event property. All rules must pass for the event to match the
	// billable metric.
	PropertyFilters param.Field[[]V1BillableMetricNewParamsPropertyFilter] `json:"property_filters"`
	// The SQL query associated with the billable metric. This field is mutually
	// exclusive with aggregation_type, event_type_filter, property_filters,
	// aggregation_key, and group_keys. If provided, these other fields must be
	// omitted.
	Sql param.Field[string] `json:"sql"`
}

func (r V1BillableMetricNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of aggregation performed on matching events.
type V1BillableMetricNewParamsAggregationType string

const (
	V1BillableMetricNewParamsAggregationTypeCount  V1BillableMetricNewParamsAggregationType = "COUNT"
	V1BillableMetricNewParamsAggregationTypeLatest V1BillableMetricNewParamsAggregationType = "LATEST"
	V1BillableMetricNewParamsAggregationTypeMax    V1BillableMetricNewParamsAggregationType = "MAX"
	V1BillableMetricNewParamsAggregationTypeSum    V1BillableMetricNewParamsAggregationType = "SUM"
	V1BillableMetricNewParamsAggregationTypeUnique V1BillableMetricNewParamsAggregationType = "UNIQUE"
)

func (r V1BillableMetricNewParamsAggregationType) IsKnown() bool {
	switch r {
	case V1BillableMetricNewParamsAggregationTypeCount, V1BillableMetricNewParamsAggregationTypeLatest, V1BillableMetricNewParamsAggregationTypeMax, V1BillableMetricNewParamsAggregationTypeSum, V1BillableMetricNewParamsAggregationTypeUnique:
		return true
	}
	return false
}

// An optional filtering rule to match the 'event_type' property of an event.
type V1BillableMetricNewParamsEventTypeFilter struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues param.Field[[]string] `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues param.Field[[]string] `json:"not_in_values"`
}

func (r V1BillableMetricNewParamsEventTypeFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1BillableMetricNewParamsPropertyFilter struct {
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

func (r V1BillableMetricNewParamsPropertyFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1BillableMetricGetParams struct {
	BillableMetricID param.Field[string] `path:"billable_metric_id,required" format:"uuid"`
}

type V1BillableMetricListParams struct {
	// If true, the list of returned metrics will include archived metrics
	IncludeArchived param.Field[bool] `query:"include_archived"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [V1BillableMetricListParams]'s query parameters as
// `url.Values`.
func (r V1BillableMetricListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1BillableMetricArchiveParams struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V1BillableMetricArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
