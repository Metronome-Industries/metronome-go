// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
)

// NotificationSystemService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationSystemService] method instead.
type NotificationSystemService struct {
	Options []option.RequestOption
}

// NewNotificationSystemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationSystemService(opts ...option.RequestOption) (r NotificationSystemService) {
	r = NotificationSystemService{}
	r.Options = opts
	return
}

// List available system lifecycle event types for notifications. These are
// read-only event types that can be used when creating offset notifications.
func (r *NotificationSystemService) List(ctx context.Context, opts ...option.RequestOption) (res *NotificationSystemListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/system/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type NotificationSystemListResponse struct {
	Data   []NotificationSystemListResponseData `json:"data,required"`
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
func (r NotificationSystemListResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationSystemListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationSystemListResponseData struct {
	Policy NotificationSystemListResponseDataPolicy `json:"policy,required"`
	// Indicates this is a system lifecycle event notification
	Type string `json:"type,required"`
	// Whether or not webhook publishing for this lifecycle event is enabled
	IsEnabled bool `json:"is_enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Policy      respjson.Field
		Type        respjson.Field
		IsEnabled   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationSystemListResponseData) RawJSON() string { return r.JSON.raw }
func (r *NotificationSystemListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationSystemListResponseDataPolicy struct {
	// The type of lifecycle event (e.g., "contract.create", "contract.start")
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationSystemListResponseDataPolicy) RawJSON() string { return r.JSON.raw }
func (r *NotificationSystemListResponseDataPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
