// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v3/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"github.com/Metronome-Industries/metronome-go/v3/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/v3/packages/param"
	"github.com/Metronome-Industries/metronome-go/v3/packages/respjson"
)

// Named schedules are used for storing custom data that can change over time.
// Named schedules are often used in custom pricing logic.
//
// V1ContractRateCardNamedScheduleService contains methods and other services that
// help with interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContractRateCardNamedScheduleService] method instead.
type V1ContractRateCardNamedScheduleService struct {
	Options []option.RequestOption
}

// NewV1ContractRateCardNamedScheduleService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1ContractRateCardNamedScheduleService(opts ...option.RequestOption) (r V1ContractRateCardNamedScheduleService) {
	r = V1ContractRateCardNamedScheduleService{}
	r.Options = opts
	return
}

// Get a named schedule for the given contract. This endpoint's availability is
// dependent on your client's configuration.
func (r *V1ContractRateCardNamedScheduleService) Get(ctx context.Context, body V1ContractRateCardNamedScheduleGetParams, opts ...option.RequestOption) (res *V1ContractRateCardNamedScheduleGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contracts/getNamedSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a named schedule for the given contract. This endpoint's availability is
// dependent on your client's configuration.
func (r *V1ContractRateCardNamedScheduleService) Update(ctx context.Context, body V1ContractRateCardNamedScheduleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/contracts/updateNamedSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// List contract-level named schedules for a customer, optionally scoped to a
// single contract.
func (r *V1ContractRateCardNamedScheduleService) List(ctx context.Context, body V1ContractRateCardNamedScheduleListParams, opts ...option.RequestOption) (res *pagination.BodyCursorPage[V1ContractRateCardNamedScheduleListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/contracts/listNamedSchedules"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
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

// List contract-level named schedules for a customer, optionally scoped to a
// single contract.
func (r *V1ContractRateCardNamedScheduleService) ListAutoPaging(ctx context.Context, body V1ContractRateCardNamedScheduleListParams, opts ...option.RequestOption) *pagination.BodyCursorPageAutoPager[V1ContractRateCardNamedScheduleListResponse] {
	return pagination.NewBodyCursorPageAutoPager(r.List(ctx, body, opts...))
}

type V1ContractRateCardNamedScheduleGetResponse struct {
	Data []V1ContractRateCardNamedScheduleGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardNamedScheduleGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardNamedScheduleGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardNamedScheduleGetResponseData struct {
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	Value        any       `json:"value" api:"required"`
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StartingAt   respjson.Field
		Value        respjson.Field
		EndingBefore respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardNamedScheduleGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardNamedScheduleGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardNamedScheduleListResponse struct {
	// ID of the contract this named schedule belongs to.
	ContractID string `json:"contract_id" api:"required" format:"uuid"`
	// Name of the schedule.
	ScheduleName string                                               `json:"schedule_name" api:"required"`
	Segments     []V1ContractRateCardNamedScheduleListResponseSegment `json:"segments" api:"required"`
	// Properties on this named schedule.
	Properties map[string]string `json:"properties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContractID   respjson.Field
		ScheduleName respjson.Field
		Segments     respjson.Field
		Properties   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardNamedScheduleListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardNamedScheduleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardNamedScheduleListResponseSegment struct {
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	Value        any       `json:"value" api:"required"`
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StartingAt   respjson.Field
		Value        respjson.Field
		EndingBefore respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardNamedScheduleListResponseSegment) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardNamedScheduleListResponseSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardNamedScheduleGetParams struct {
	// ID of the contract whose named schedule is to be retrieved
	ContractID string `json:"contract_id" api:"required" format:"uuid"`
	// ID of the customer whose named schedule is to be retrieved
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// The identifier for the schedule to be retrieved
	ScheduleName string `json:"schedule_name" api:"required"`
	// If provided, at most one schedule segment will be returned (the one that covers
	// this date). If not provided, all segments will be returned.
	CoveringDate param.Opt[time.Time] `json:"covering_date,omitzero" format:"date-time"`
	// A set of key-value pairs that qualifies which schedule should be applied when
	// looking up a schedule by name.
	Properties map[string]string `json:"properties,omitzero"`
	paramObj
}

func (r V1ContractRateCardNamedScheduleGetParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardNamedScheduleGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardNamedScheduleGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardNamedScheduleUpdateParams struct {
	// ID of the contract whose named schedule is to be updated
	ContractID string `json:"contract_id" api:"required" format:"uuid"`
	// ID of the customer whose named schedule is to be updated
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// The identifier for the schedule to be updated
	ScheduleName string    `json:"schedule_name" api:"required"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	// The value to set for the named schedule. The structure of this object is
	// specific to the named schedule.
	Value        any                  `json:"value,omitzero" api:"required"`
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// A set of key-value pairs that qualifies which schedule should be applied when
	// looking up a schedule by name.
	Properties map[string]string `json:"properties,omitzero"`
	paramObj
}

func (r V1ContractRateCardNamedScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardNamedScheduleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardNamedScheduleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardNamedScheduleListParams struct {
	// ID of the customer to list contract named schedules for
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Optional ID of a contract to scope the results to a single contract
	ContractID param.Opt[string] `json:"contract_id,omitzero" format:"uuid"`
	// If provided, result contains at most one segment covering this date for each
	// schedule.
	CoveringDate param.Opt[time.Time] `json:"covering_date,omitzero" format:"date-time"`
	// Maximum number of named schedules to return.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Cursor for pagination. Use the value from a previous response's next_page.
	NextPage param.Opt[string] `json:"next_page,omitzero"`
	// Optional filter to scope the results to a single schedule name
	ScheduleName param.Opt[string] `json:"schedule_name,omitzero"`
	// Optional key-value filters. A schedule matches when it contains at least this
	// subset of properties.
	Properties map[string]string `json:"properties,omitzero"`
	paramObj
}

func (r V1ContractRateCardNamedScheduleListParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardNamedScheduleListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardNamedScheduleListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
