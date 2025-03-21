// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
)

// CustomerNamedScheduleService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerNamedScheduleService] method instead.
type CustomerNamedScheduleService struct {
	Options []option.RequestOption
}

// NewCustomerNamedScheduleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomerNamedScheduleService(opts ...option.RequestOption) (r *CustomerNamedScheduleService) {
	r = &CustomerNamedScheduleService{}
	r.Options = opts
	return
}

// Get a named schedule for the given customer. This endpoint's availability is
// dependent on your client's configuration.
func (r *CustomerNamedScheduleService) Get(ctx context.Context, body CustomerNamedScheduleGetParams, opts ...option.RequestOption) (res *CustomerNamedScheduleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/customers/getNamedSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a named schedule for the given customer. This endpoint's availability is
// dependent on your client's configuration.
func (r *CustomerNamedScheduleService) Update(ctx context.Context, body CustomerNamedScheduleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/customers/updateNamedSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CustomerNamedScheduleGetResponse struct {
	Data []CustomerNamedScheduleGetResponseData `json:"data,required"`
	JSON customerNamedScheduleGetResponseJSON   `json:"-"`
}

// customerNamedScheduleGetResponseJSON contains the JSON metadata for the struct
// [CustomerNamedScheduleGetResponse]
type customerNamedScheduleGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerNamedScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerNamedScheduleGetResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerNamedScheduleGetResponseData struct {
	StartingAt   time.Time                                `json:"starting_at,required" format:"date-time"`
	Value        interface{}                              `json:"value,required"`
	EndingBefore time.Time                                `json:"ending_before" format:"date-time"`
	JSON         customerNamedScheduleGetResponseDataJSON `json:"-"`
}

// customerNamedScheduleGetResponseDataJSON contains the JSON metadata for the
// struct [CustomerNamedScheduleGetResponseData]
type customerNamedScheduleGetResponseDataJSON struct {
	StartingAt   apijson.Field
	Value        apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomerNamedScheduleGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerNamedScheduleGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type CustomerNamedScheduleGetParams struct {
	// ID of the customer whose named schedule is to be retrieved
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// The identifier for the schedule to be retrieved
	ScheduleName param.Field[string] `json:"schedule_name,required"`
	// If provided, at most one schedule segment will be returned (the one that covers
	// this date). If not provided, all segments will be returned.
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
}

func (r CustomerNamedScheduleGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNamedScheduleUpdateParams struct {
	// ID of the customer whose named schedule is to be updated
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// The identifier for the schedule to be updated
	ScheduleName param.Field[string]    `json:"schedule_name,required"`
	StartingAt   param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// The value to set for the named schedule. The structure of this object is
	// specific to the named schedule.
	Value        param.Field[interface{}] `json:"value,required"`
	EndingBefore param.Field[time.Time]   `json:"ending_before" format:"date-time"`
}

func (r CustomerNamedScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
