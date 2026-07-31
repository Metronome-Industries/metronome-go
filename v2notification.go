// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"time"

	"github.com/Metronome-Industries/metronome-go/v3/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"github.com/Metronome-Industries/metronome-go/v3/packages/respjson"
)

// V2NotificationService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2NotificationService] method instead.
type V2NotificationService struct {
	Options []option.RequestOption
	Offset  V2NotificationOffsetService
	System  V2NotificationSystemService
}

// NewV2NotificationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2NotificationService(opts ...option.RequestOption) (r V2NotificationService) {
	r = V2NotificationService{}
	r.Options = opts
	r.Offset = NewV2NotificationOffsetService(opts...)
	r.System = NewV2NotificationSystemService(opts...)
	return
}

type LifecycleEventOffsetNotificationConfig struct {
	// ID for this offset notification configuration
	ID string `json:"id" api:"required" format:"uuid"`
	// When this notification configuration was archived
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
	// RFC 3339 timestamp when this notification configuration was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Who created this notification configuration
	CreatedBy string `json:"created_by" api:"required"`
	// The environment type where this notification configuration was created.
	EnvironmentType string `json:"environment_type" api:"required"`
	// The name for this offset notification configuration.
	Name   string                                       `json:"name" api:"required"`
	Policy LifecycleEventOffsetNotificationConfigPolicy `json:"policy" api:"required"`
	// Indicates this is an offset lifecycle event notification
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ArchivedAt      respjson.Field
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
func (r LifecycleEventOffsetNotificationConfig) RawJSON() string { return r.JSON.raw }
func (r *LifecycleEventOffsetNotificationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LifecycleEventOffsetNotificationConfigPolicy struct {
	// ISO-8601 duration string indicating how much time before or after the base event
	// this notification should be sent. Positive values indicate notifications after
	// the event, negative values indicate notifications before the event. Examples:
	// "P1D" (1 day after), "-PT2H" (2 hours before)
	Offset string `json:"offset" api:"required"`
	// The type of lifecycle event that this offset is based on.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LifecycleEventOffsetNotificationConfigPolicy) RawJSON() string { return r.JSON.raw }
func (r *LifecycleEventOffsetNotificationConfigPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LifecycleEventSystemNotificationConfig struct {
	Policy LifecycleEventSystemNotificationConfigPolicy `json:"policy" api:"required"`
	// Indicates this is a system lifecycle event notification
	Type string `json:"type" api:"required"`
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
func (r LifecycleEventSystemNotificationConfig) RawJSON() string { return r.JSON.raw }
func (r *LifecycleEventSystemNotificationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LifecycleEventSystemNotificationConfigPolicy struct {
	// The type of lifecycle event (e.g., "contract.create", "contract.start")
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LifecycleEventSystemNotificationConfigPolicy) RawJSON() string { return r.JSON.raw }
func (r *LifecycleEventSystemNotificationConfigPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
