// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"

	"github.com/Metronome-Industries/metronome-go/v3/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v3/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"github.com/Metronome-Industries/metronome-go/v3/packages/respjson"
)

// V2NotificationSystemService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2NotificationSystemService] method instead.
type V2NotificationSystemService struct {
	Options []option.RequestOption
}

// NewV2NotificationSystemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2NotificationSystemService(opts ...option.RequestOption) (r V2NotificationSystemService) {
	r = V2NotificationSystemService{}
	r.Options = opts
	return
}

// List available system lifecycle event types for notifications. These are
// read-only event types that can be used when creating offset notifications.
func (r *V2NotificationSystemService) List(ctx context.Context, opts ...option.RequestOption) (res *V2NotificationSystemListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/system/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type V2NotificationSystemListResponse struct {
	Data   []LifecycleEventSystemNotificationConfig `json:"data" api:"required"`
	Cursor string                                   `json:"cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Cursor      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2NotificationSystemListResponse) RawJSON() string { return r.JSON.raw }
func (r *V2NotificationSystemListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
