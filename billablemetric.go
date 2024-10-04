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
func (r *BillableMetricService) Get(ctx context.Context, query BillableMetricGetParams, opts ...option.RequestOption) (res *BillableMetricGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.BillableMetricID.Value == "" {
		err = errors.New("missing required billable_metric_id parameter")
		return
	}
	path := fmt.Sprintf("billable-metrics/%s", query.BillableMetricID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all billable metrics.
func (r *BillableMetricService) List(ctx context.Context, query BillableMetricListParams, opts ...option.RequestOption) (res *pagination.CursorPage[BillableMetricListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "billable-metrics"
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
func (r *BillableMetricService) ListAutoPaging(ctx context.Context, query BillableMetricListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[BillableMetricListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
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
	EventTypeFilter shared.EventTypeFilter `json:"event_type_filter"`
	// Property names that are used to group usage costs on an invoice. Each entry
	// represents a set of properties used to slice events into distinct buckets.
	GroupKeys [][]string `json:"group_keys"`
	// A list of filters to match events to this billable metric. Each filter defines a
	// rule on an event property. All rules must pass for the event to match the
	// billable metric.
	PropertyFilters []shared.PropertyFilter `json:"property_filters"`
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
	BillableMetricGetResponseDataAggregationTypeCount  BillableMetricGetResponseDataAggregationType = "COUNT"
	BillableMetricGetResponseDataAggregationTypeLatest BillableMetricGetResponseDataAggregationType = "LATEST"
	BillableMetricGetResponseDataAggregationTypeMax    BillableMetricGetResponseDataAggregationType = "MAX"
	BillableMetricGetResponseDataAggregationTypeSum    BillableMetricGetResponseDataAggregationType = "SUM"
	BillableMetricGetResponseDataAggregationTypeUnique BillableMetricGetResponseDataAggregationType = "UNIQUE"
)

func (r BillableMetricGetResponseDataAggregationType) IsKnown() bool {
	switch r {
	case BillableMetricGetResponseDataAggregationTypeCount, BillableMetricGetResponseDataAggregationTypeLatest, BillableMetricGetResponseDataAggregationTypeMax, BillableMetricGetResponseDataAggregationTypeSum, BillableMetricGetResponseDataAggregationTypeUnique:
		return true
	}
	return false
}

type BillableMetricListResponse struct {
	// ID of the billable metric
	ID string `json:"id,required" format:"uuid"`
	// The display name of the billable metric.
	Name string `json:"name,required"`
	// A key that specifies which property of the event is used to aggregate data. This
	// key must be one of the property filter names and is not applicable when the
	// aggregation type is 'count'.
	AggregationKey string `json:"aggregation_key"`
	// Specifies the type of aggregation performed on matching events.
	AggregationType BillableMetricListResponseAggregationType `json:"aggregation_type"`
	CustomFields    map[string]string                         `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter shared.EventTypeFilter `json:"event_type_filter"`
	// Property names that are used to group usage costs on an invoice. Each entry
	// represents a set of properties used to slice events into distinct buckets.
	GroupKeys [][]string `json:"group_keys"`
	// A list of filters to match events to this billable metric. Each filter defines a
	// rule on an event property. All rules must pass for the event to match the
	// billable metric.
	PropertyFilters []shared.PropertyFilter `json:"property_filters"`
	// The SQL query associated with the billable metric
	Sql  string                         `json:"sql"`
	JSON billableMetricListResponseJSON `json:"-"`
}

// billableMetricListResponseJSON contains the JSON metadata for the struct
// [BillableMetricListResponse]
type billableMetricListResponseJSON struct {
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

func (r *BillableMetricListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricListResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of aggregation performed on matching events.
type BillableMetricListResponseAggregationType string

const (
	BillableMetricListResponseAggregationTypeCount  BillableMetricListResponseAggregationType = "COUNT"
	BillableMetricListResponseAggregationTypeLatest BillableMetricListResponseAggregationType = "LATEST"
	BillableMetricListResponseAggregationTypeMax    BillableMetricListResponseAggregationType = "MAX"
	BillableMetricListResponseAggregationTypeSum    BillableMetricListResponseAggregationType = "SUM"
	BillableMetricListResponseAggregationTypeUnique BillableMetricListResponseAggregationType = "UNIQUE"
)

func (r BillableMetricListResponseAggregationType) IsKnown() bool {
	switch r {
	case BillableMetricListResponseAggregationTypeCount, BillableMetricListResponseAggregationTypeLatest, BillableMetricListResponseAggregationTypeMax, BillableMetricListResponseAggregationTypeSum, BillableMetricListResponseAggregationTypeUnique:
		return true
	}
	return false
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
	EventTypeFilter param.Field[shared.EventTypeFilterParam] `json:"event_type_filter"`
	// Property names that are used to group usage costs on an invoice. Each entry
	// represents a set of properties used to slice events into distinct buckets.
	GroupKeys param.Field[[][]string] `json:"group_keys"`
	// A list of filters to match events to this billable metric. Each filter defines a
	// rule on an event property. All rules must pass for the event to match the
	// billable metric.
	PropertyFilters param.Field[[]shared.PropertyFilterParam] `json:"property_filters"`
}

func (r BillableMetricNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of aggregation performed on matching events.
type BillableMetricNewParamsAggregationType string

const (
	BillableMetricNewParamsAggregationTypeCount  BillableMetricNewParamsAggregationType = "COUNT"
	BillableMetricNewParamsAggregationTypeLatest BillableMetricNewParamsAggregationType = "LATEST"
	BillableMetricNewParamsAggregationTypeMax    BillableMetricNewParamsAggregationType = "MAX"
	BillableMetricNewParamsAggregationTypeSum    BillableMetricNewParamsAggregationType = "SUM"
	BillableMetricNewParamsAggregationTypeUnique BillableMetricNewParamsAggregationType = "UNIQUE"
)

func (r BillableMetricNewParamsAggregationType) IsKnown() bool {
	switch r {
	case BillableMetricNewParamsAggregationTypeCount, BillableMetricNewParamsAggregationTypeLatest, BillableMetricNewParamsAggregationTypeMax, BillableMetricNewParamsAggregationTypeSum, BillableMetricNewParamsAggregationTypeUnique:
		return true
	}
	return false
}

type BillableMetricGetParams struct {
	BillableMetricID param.Field[string] `path:"billable_metric_id,required" format:"uuid"`
}

type BillableMetricListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
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
