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

// NotificationOffsetService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationOffsetService] method instead.
type NotificationOffsetService struct {
	Options []option.RequestOption
}

// NewNotificationOffsetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationOffsetService(opts ...option.RequestOption) (r NotificationOffsetService) {
	r = NotificationOffsetService{}
	r.Options = opts
	return
}

// List offset lifecycle event notification configurations. These are user-created
// notifications that trigger at a specified time offset relative to lifecycle
// events.
func (r *NotificationOffsetService) List(ctx context.Context, body NotificationOffsetListParams, opts ...option.RequestOption) (res *NotificationOffsetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/offset/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type NotificationOffsetListResponse struct {
	Data   []NotificationOffsetListResponseData `json:"data,required"`
	Cursor string                               `json:"cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Cursor      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationOffsetListResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationOffsetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationOffsetListResponseData struct {
	// ID for this offset notification configuration
	ID string `json:"id,required" format:"uuid"`
	// RFC 3339 timestamp when this notification configuration was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Who created this notification configuration
	CreatedBy string `json:"created_by,required"`
	// The environment type where this notification configuration was created.
	EnvironmentType string `json:"environment_type,required"`
	// The name for this offset notification configuration.
	Name   string                                   `json:"name,required"`
	Policy NotificationOffsetListResponseDataPolicy `json:"policy,required"`
	// Indicates this is an offset lifecycle event notification
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		CreatedBy       respjson.Field
		EnvironmentType respjson.Field
		Name            respjson.Field
		Policy          respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationOffsetListResponseData) RawJSON() string { return r.JSON.raw }
func (r *NotificationOffsetListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationOffsetListResponseDataPolicy struct {
	// ISO-8601 duration string indicating how much time before or after the base event
	// this notification should be sent. Positive values indicate notifications after
	// the event, negative values indicate notifications before the event. Examples:
	// "P1D" (1 day after), "-PT2H" (2 hours before)
	Offset string `json:"offset,required"`
	// The type of lifecycle event that this offset is based on.
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationOffsetListResponseDataPolicy) RawJSON() string { return r.JSON.raw }
func (r *NotificationOffsetListResponseDataPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationOffsetListParams struct {
	Cursor param.Opt[string]  `json:"cursor,omitzero"`
	Limit  param.Opt[float64] `json:"limit,omitzero"`
	paramObj
}

func (r NotificationOffsetListParams) MarshalJSON() (data []byte, err error) {
	type shadow NotificationOffsetListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationOffsetListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
