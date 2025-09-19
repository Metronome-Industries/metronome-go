// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
)

// V1CustomerNamedScheduleService contains methods and other services that help
// with interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerNamedScheduleService] method instead.
type V1CustomerNamedScheduleService struct {
	Options []option.RequestOption
}

// NewV1CustomerNamedScheduleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerNamedScheduleService(opts ...option.RequestOption) (r V1CustomerNamedScheduleService) {
	r = V1CustomerNamedScheduleService{}
	r.Options = opts
	return
}

// Get a named schedule for the given customer. This endpoint's availability is
// dependent on your client's configuration.
func (r *V1CustomerNamedScheduleService) Get(ctx context.Context, body V1CustomerNamedScheduleGetParams, opts ...option.RequestOption) (res *V1CustomerNamedScheduleGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/customers/getNamedSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a named schedule for the given customer. This endpoint's availability is
// dependent on your client's configuration.
func (r *V1CustomerNamedScheduleService) Update(ctx context.Context, body V1CustomerNamedScheduleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/customers/updateNamedSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type V1CustomerNamedScheduleGetResponse struct {
	Data []V1CustomerNamedScheduleGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerNamedScheduleGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerNamedScheduleGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerNamedScheduleGetResponseData struct {
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	Value        any       `json:"value,required"`
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
func (r V1CustomerNamedScheduleGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerNamedScheduleGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerNamedScheduleGetParams struct {
	// ID of the customer whose named schedule is to be retrieved
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// The identifier for the schedule to be retrieved
	ScheduleName string `json:"schedule_name,required"`
	// If provided, at most one schedule segment will be returned (the one that covers
	// this date). If not provided, all segments will be returned.
	CoveringDate param.Opt[time.Time] `json:"covering_date,omitzero" format:"date-time"`
	paramObj
}

func (r V1CustomerNamedScheduleGetParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerNamedScheduleGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerNamedScheduleGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerNamedScheduleUpdateParams struct {
	// ID of the customer whose named schedule is to be updated
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// The identifier for the schedule to be updated
	ScheduleName string    `json:"schedule_name,required"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	// The value to set for the named schedule. The structure of this object is
	// specific to the named schedule.
	Value        any                  `json:"value,omitzero,required"`
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	paramObj
}

func (r V1CustomerNamedScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerNamedScheduleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerNamedScheduleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
