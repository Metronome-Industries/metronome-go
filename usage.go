// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/pagination"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
)

// UsageService contains methods and other services that help with interacting with
// the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageService] method instead.
type UsageService struct {
	Options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r *UsageService) {
	r = &UsageService{}
	r.Options = opts
	return
}

// Fetch aggregated usage data for multiple customers and billable-metrics, broken
// into intervals of the specified length.
func (r *UsageService) List(ctx context.Context, params UsageListParams, opts ...option.RequestOption) (res *UsageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Send usage events to Metronome. The body of this request is expected to be a
// JSON array of between 1 and 100 usage events. Compressed request bodies are
// supported with a `Content-Encoding: gzip` header. See
// [Getting usage into Metronome](https://docs.metronome.com/getting-usage-data-into-metronome/overview)
// to learn more about usage events.
func (r *UsageService) Ingest(ctx context.Context, body UsageIngestParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ingest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Fetch aggregated usage data for the specified customer, billable-metric, and
// optional group, broken into intervals of the specified length.
func (r *UsageService) ListWithGroups(ctx context.Context, params UsageListWithGroupsParams, opts ...option.RequestOption) (res *pagination.CursorPage[UsageListWithGroupsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "usage/groups"
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
func (r *UsageService) ListWithGroupsAutoPaging(ctx context.Context, params UsageListWithGroupsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[UsageListWithGroupsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListWithGroups(ctx, params, opts...))
}

type UsageListResponse struct {
	Data     []UsageListResponseData `json:"data,required"`
	NextPage string                  `json:"next_page,required,nullable"`
	JSON     usageListResponseJSON   `json:"-"`
}

// usageListResponseJSON contains the JSON metadata for the struct
// [UsageListResponse]
type usageListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageListResponseJSON) RawJSON() string {
	return r.raw
}

type UsageListResponseData struct {
	BillableMetricID   string    `json:"billable_metric_id,required" format:"uuid"`
	BillableMetricName string    `json:"billable_metric_name,required"`
	CustomerID         string    `json:"customer_id,required" format:"uuid"`
	EndTimestamp       time.Time `json:"end_timestamp,required" format:"date-time"`
	StartTimestamp     time.Time `json:"start_timestamp,required" format:"date-time"`
	Value              float64   `json:"value,required,nullable"`
	// Values will be either a number or null. Null indicates that there were no
	// matches for the group_by value.
	Groups map[string]float64        `json:"groups"`
	JSON   usageListResponseDataJSON `json:"-"`
}

// usageListResponseDataJSON contains the JSON metadata for the struct
// [UsageListResponseData]
type usageListResponseDataJSON struct {
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

func (r *UsageListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageListResponseDataJSON) RawJSON() string {
	return r.raw
}

type UsageListWithGroupsResponse struct {
	EndingBefore time.Time                       `json:"ending_before,required" format:"date-time"`
	GroupKey     string                          `json:"group_key,required,nullable"`
	GroupValue   string                          `json:"group_value,required,nullable"`
	StartingOn   time.Time                       `json:"starting_on,required" format:"date-time"`
	Value        float64                         `json:"value,required,nullable"`
	JSON         usageListWithGroupsResponseJSON `json:"-"`
}

// usageListWithGroupsResponseJSON contains the JSON metadata for the struct
// [UsageListWithGroupsResponse]
type usageListWithGroupsResponseJSON struct {
	EndingBefore apijson.Field
	GroupKey     apijson.Field
	GroupValue   apijson.Field
	StartingOn   apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UsageListWithGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageListWithGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type UsageListParams struct {
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	StartingOn   param.Field[time.Time] `json:"starting_on,required" format:"date-time"`
	// A window_size of "day" or "hour" will return the usage for the specified period
	// segmented into daily or hourly aggregates. A window_size of "none" will return a
	// single usage aggregate for the entirety of the specified period.
	WindowSize param.Field[UsageListParamsWindowSize] `json:"window_size,required"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// A list of billable metrics to fetch usage for. If absent, all billable metrics
	// will be returned.
	BillableMetrics param.Field[[]UsageListParamsBillableMetric] `json:"billable_metrics"`
	// A list of Metronome customer IDs to fetch usage for. If absent, usage for all
	// customers will be returned.
	CustomerIDs param.Field[[]string] `json:"customer_ids" format:"uuid"`
}

func (r UsageListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [UsageListParams]'s query parameters as `url.Values`.
func (r UsageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A window_size of "day" or "hour" will return the usage for the specified period
// segmented into daily or hourly aggregates. A window_size of "none" will return a
// single usage aggregate for the entirety of the specified period.
type UsageListParamsWindowSize string

const (
	UsageListParamsWindowSizeHour UsageListParamsWindowSize = "hour"
	UsageListParamsWindowSizeDay  UsageListParamsWindowSize = "day"
	UsageListParamsWindowSizeNone UsageListParamsWindowSize = "none"
	UsageListParamsWindowSizeHour UsageListParamsWindowSize = "HOUR"
	UsageListParamsWindowSizeDay  UsageListParamsWindowSize = "DAY"
	UsageListParamsWindowSizeNone UsageListParamsWindowSize = "NONE"
	UsageListParamsWindowSizeHour UsageListParamsWindowSize = "Hour"
	UsageListParamsWindowSizeDay  UsageListParamsWindowSize = "Day"
	UsageListParamsWindowSizeNone UsageListParamsWindowSize = "None"
)

func (r UsageListParamsWindowSize) IsKnown() bool {
	switch r {
	case UsageListParamsWindowSizeHour, UsageListParamsWindowSizeDay, UsageListParamsWindowSizeNone, UsageListParamsWindowSizeHour, UsageListParamsWindowSizeDay, UsageListParamsWindowSizeNone, UsageListParamsWindowSizeHour, UsageListParamsWindowSizeDay, UsageListParamsWindowSizeNone:
		return true
	}
	return false
}

type UsageListParamsBillableMetric struct {
	ID      param.Field[string]                                `json:"id,required" format:"uuid"`
	GroupBy param.Field[UsageListParamsBillableMetricsGroupBy] `json:"group_by"`
}

func (r UsageListParamsBillableMetric) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageListParamsBillableMetricsGroupBy struct {
	// The name of the group_by key to use
	Key param.Field[string] `json:"key,required"`
	// Values of the group_by key to return in the query. If this field is omitted, all
	// available values will be returned, up to a maximum of 200.
	Values param.Field[[]string] `json:"values"`
}

func (r UsageListParamsBillableMetricsGroupBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageIngestParams struct {
	Usage []UsageIngestParamsUsage `json:"usage,required"`
}

func (r UsageIngestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Usage)
}

type UsageIngestParamsUsage struct {
	CustomerID param.Field[string] `json:"customer_id,required"`
	EventType  param.Field[string] `json:"event_type,required"`
	// RFC 3339 formatted
	Timestamp     param.Field[string]                 `json:"timestamp,required"`
	TransactionID param.Field[string]                 `json:"transaction_id,required"`
	Properties    param.Field[map[string]interface{}] `json:"properties"`
}

func (r UsageIngestParamsUsage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageListWithGroupsParams struct {
	BillableMetricID param.Field[string] `json:"billable_metric_id,required" format:"uuid"`
	CustomerID       param.Field[string] `json:"customer_id,required" format:"uuid"`
	// A window_size of "day" or "hour" will return the usage for the specified period
	// segmented into daily or hourly aggregates. A window_size of "none" will return a
	// single usage aggregate for the entirety of the specified period.
	WindowSize param.Field[UsageListWithGroupsParamsWindowSize] `json:"window_size,required"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// If true, will return the usage for the current billing period. Will return an
	// error if the customer is currently uncontracted or starting_on and ending_before
	// are specified when this is true.
	CurrentPeriod param.Field[bool]                             `json:"current_period"`
	EndingBefore  param.Field[time.Time]                        `json:"ending_before" format:"date-time"`
	GroupBy       param.Field[UsageListWithGroupsParamsGroupBy] `json:"group_by"`
	StartingOn    param.Field[time.Time]                        `json:"starting_on" format:"date-time"`
}

func (r UsageListWithGroupsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [UsageListWithGroupsParams]'s query parameters as
// `url.Values`.
func (r UsageListWithGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A window_size of "day" or "hour" will return the usage for the specified period
// segmented into daily or hourly aggregates. A window_size of "none" will return a
// single usage aggregate for the entirety of the specified period.
type UsageListWithGroupsParamsWindowSize string

const (
	UsageListWithGroupsParamsWindowSizeHour UsageListWithGroupsParamsWindowSize = "hour"
	UsageListWithGroupsParamsWindowSizeDay  UsageListWithGroupsParamsWindowSize = "day"
	UsageListWithGroupsParamsWindowSizeNone UsageListWithGroupsParamsWindowSize = "none"
	UsageListWithGroupsParamsWindowSizeHour UsageListWithGroupsParamsWindowSize = "HOUR"
	UsageListWithGroupsParamsWindowSizeDay  UsageListWithGroupsParamsWindowSize = "DAY"
	UsageListWithGroupsParamsWindowSizeNone UsageListWithGroupsParamsWindowSize = "NONE"
	UsageListWithGroupsParamsWindowSizeHour UsageListWithGroupsParamsWindowSize = "Hour"
	UsageListWithGroupsParamsWindowSizeDay  UsageListWithGroupsParamsWindowSize = "Day"
	UsageListWithGroupsParamsWindowSizeNone UsageListWithGroupsParamsWindowSize = "None"
)

func (r UsageListWithGroupsParamsWindowSize) IsKnown() bool {
	switch r {
	case UsageListWithGroupsParamsWindowSizeHour, UsageListWithGroupsParamsWindowSizeDay, UsageListWithGroupsParamsWindowSizeNone, UsageListWithGroupsParamsWindowSizeHour, UsageListWithGroupsParamsWindowSizeDay, UsageListWithGroupsParamsWindowSizeNone, UsageListWithGroupsParamsWindowSizeHour, UsageListWithGroupsParamsWindowSizeDay, UsageListWithGroupsParamsWindowSizeNone:
		return true
	}
	return false
}

type UsageListWithGroupsParamsGroupBy struct {
	// The name of the group_by key to use
	Key param.Field[string] `json:"key,required"`
	// Values of the group_by key to return in the query. Omit this if you'd like all
	// values for the key returned.
	Values param.Field[[]string] `json:"values"`
}

func (r UsageListWithGroupsParamsGroupBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
