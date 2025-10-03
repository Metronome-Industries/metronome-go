// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
)

// NotificationService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationService] method instead.
type NotificationService struct {
	Options []option.RequestOption
	Offset  NotificationOffsetService
	System  NotificationSystemService
}

// NewNotificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNotificationService(opts ...option.RequestOption) (r NotificationService) {
	r = NotificationService{}
	r.Options = opts
	r.Offset = NewNotificationOffsetService(opts...)
	r.System = NewNotificationSystemService(opts...)
	return
}

// Create an offset lifecycle event notification configuration. The lifecycle event
// type is inferred from the policy.type field.
func (r *NotificationService) New(ctx context.Context, body NotificationNewParams, opts ...option.RequestOption) (res *NotificationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific offset lifecycle event notification configuration by ID.
func (r *NotificationService) Get(ctx context.Context, body NotificationGetParams, opts ...option.RequestOption) (res *NotificationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Edit an existing offset lifecycle event notification configuration.
func (r *NotificationService) Update(ctx context.Context, body NotificationUpdateParams, opts ...option.RequestOption) (res *NotificationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Archive an offset lifecycle event notification configuration. Archived
// notifications are not processed.
func (r *NotificationService) Archive(ctx context.Context, body NotificationArchiveParams, opts ...option.RequestOption) (res *NotificationArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type NotificationNewResponse struct {
	Data NotificationNewResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationNewResponseData struct {
	// ID for this offset notification configuration
	ID string `json:"id,required" format:"uuid"`
	// When this notification configuration was archived
	ArchivedAt time.Time `json:"archived_at,required" format:"date-time"`
	// RFC 3339 timestamp when this notification configuration was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Who created this notification configuration
	CreatedBy string `json:"created_by,required"`
	// The environment type where this notification configuration was created.
	EnvironmentType string `json:"environment_type,required"`
	// The name for this offset notification configuration.
	Name   string                            `json:"name,required"`
	Policy NotificationNewResponseDataPolicy `json:"policy,required"`
	// Indicates this is an offset lifecycle event notification
	Type string `json:"type,required"`
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
func (r NotificationNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *NotificationNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationNewResponseDataPolicy struct {
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
func (r NotificationNewResponseDataPolicy) RawJSON() string { return r.JSON.raw }
func (r *NotificationNewResponseDataPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetResponse struct {
	Data NotificationGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetResponseData struct {
	// ID for this offset notification configuration
	ID string `json:"id,required" format:"uuid"`
	// When this notification configuration was archived
	ArchivedAt time.Time `json:"archived_at,required" format:"date-time"`
	// RFC 3339 timestamp when this notification configuration was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Who created this notification configuration
	CreatedBy string `json:"created_by,required"`
	// The environment type where this notification configuration was created.
	EnvironmentType string `json:"environment_type,required"`
	// The name for this offset notification configuration.
	Name   string                            `json:"name,required"`
	Policy NotificationGetResponseDataPolicy `json:"policy,required"`
	// Indicates this is an offset lifecycle event notification
	Type string `json:"type,required"`
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
func (r NotificationGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetResponseDataPolicy struct {
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
func (r NotificationGetResponseDataPolicy) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetResponseDataPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationUpdateResponse struct {
	Data NotificationUpdateResponseDataUnion `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NotificationUpdateResponseDataUnion contains all possible properties and values
// from [NotificationUpdateResponseDataLifecycleEventSystemNotificationConfig],
// [NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type NotificationUpdateResponseDataUnion struct {
	// This field is a union of
	// [NotificationUpdateResponseDataLifecycleEventSystemNotificationConfigPolicy],
	// [NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfigPolicy]
	Policy NotificationUpdateResponseDataUnionPolicy `json:"policy"`
	Type   string                                    `json:"type"`
	// This field is from variant
	// [NotificationUpdateResponseDataLifecycleEventSystemNotificationConfig].
	IsEnabled bool `json:"is_enabled"`
	// This field is from variant
	// [NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig].
	ID string `json:"id"`
	// This field is from variant
	// [NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig].
	ArchivedAt time.Time `json:"archived_at"`
	// This field is from variant
	// [NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant
	// [NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig].
	CreatedBy string `json:"created_by"`
	// This field is from variant
	// [NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig].
	EnvironmentType string `json:"environment_type"`
	// This field is from variant
	// [NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig].
	Name string `json:"name"`
	JSON struct {
		Policy          respjson.Field
		Type            respjson.Field
		IsEnabled       respjson.Field
		ID              respjson.Field
		ArchivedAt      respjson.Field
		CreatedAt       respjson.Field
		CreatedBy       respjson.Field
		EnvironmentType respjson.Field
		Name            respjson.Field
		raw             string
	} `json:"-"`
}

func (u NotificationUpdateResponseDataUnion) AsNotificationUpdateResponseDataLifecycleEventSystemNotificationConfig() (v NotificationUpdateResponseDataLifecycleEventSystemNotificationConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NotificationUpdateResponseDataUnion) AsNotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig() (v NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NotificationUpdateResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *NotificationUpdateResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NotificationUpdateResponseDataUnionPolicy is an implicit subunion of
// [NotificationUpdateResponseDataUnion]. NotificationUpdateResponseDataUnionPolicy
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [NotificationUpdateResponseDataUnion].
type NotificationUpdateResponseDataUnionPolicy struct {
	Type string `json:"type"`
	// This field is from variant
	// [NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfigPolicy].
	Offset string `json:"offset"`
	JSON   struct {
		Type   respjson.Field
		Offset respjson.Field
		raw    string
	} `json:"-"`
}

func (r *NotificationUpdateResponseDataUnionPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationUpdateResponseDataLifecycleEventSystemNotificationConfig struct {
	Policy NotificationUpdateResponseDataLifecycleEventSystemNotificationConfigPolicy `json:"policy,required"`
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
func (r NotificationUpdateResponseDataLifecycleEventSystemNotificationConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationUpdateResponseDataLifecycleEventSystemNotificationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationUpdateResponseDataLifecycleEventSystemNotificationConfigPolicy struct {
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
func (r NotificationUpdateResponseDataLifecycleEventSystemNotificationConfigPolicy) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationUpdateResponseDataLifecycleEventSystemNotificationConfigPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig struct {
	// ID for this offset notification configuration
	ID string `json:"id,required" format:"uuid"`
	// When this notification configuration was archived
	ArchivedAt time.Time `json:"archived_at,required" format:"date-time"`
	// RFC 3339 timestamp when this notification configuration was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Who created this notification configuration
	CreatedBy string `json:"created_by,required"`
	// The environment type where this notification configuration was created.
	EnvironmentType string `json:"environment_type,required"`
	// The name for this offset notification configuration.
	Name   string                                                                     `json:"name,required"`
	Policy NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfigPolicy `json:"policy,required"`
	// Indicates this is an offset lifecycle event notification
	Type string `json:"type,required"`
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
func (r NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfigPolicy struct {
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
func (r NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfigPolicy) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationUpdateResponseDataLifecycleEventOffsetNotificationConfigPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationArchiveResponse struct {
	Data NotificationArchiveResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationArchiveResponseData struct {
	// ID for this offset notification configuration
	ID string `json:"id,required" format:"uuid"`
	// When this notification configuration was archived
	ArchivedAt time.Time `json:"archived_at,required" format:"date-time"`
	// RFC 3339 timestamp when this notification configuration was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Who created this notification configuration
	CreatedBy string `json:"created_by,required"`
	// The environment type where this notification configuration was created.
	EnvironmentType string `json:"environment_type,required"`
	// The name for this offset notification configuration.
	Name   string                                `json:"name,required"`
	Policy NotificationArchiveResponseDataPolicy `json:"policy,required"`
	// Indicates this is an offset lifecycle event notification
	Type string `json:"type,required"`
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
func (r NotificationArchiveResponseData) RawJSON() string { return r.JSON.raw }
func (r *NotificationArchiveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationArchiveResponseDataPolicy struct {
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
func (r NotificationArchiveResponseDataPolicy) RawJSON() string { return r.JSON.raw }
func (r *NotificationArchiveResponseDataPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationNewParams struct {
	// The name for this offset notification configuration.
	Name string `json:"name,required"`
	// The offset lifecycle event policy that defines when and how this notification
	// should be triggered. The lifecycle event type is inferred from the policy.type
	// field.
	Policy NotificationNewParamsPolicy `json:"policy,omitzero,required"`
	// Optional uniqueness key to prevent duplicate notification configurations.
	UniquenessKey param.Opt[string] `json:"uniqueness_key,omitzero"`
	paramObj
}

func (r NotificationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NotificationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The offset lifecycle event policy that defines when and how this notification
// should be triggered. The lifecycle event type is inferred from the policy.type
// field.
//
// The properties Offset, Type are required.
type NotificationNewParamsPolicy struct {
	// ISO-8601 duration string indicating how much time before or after the base event
	// this notification should be sent. Positive values indicate notifications after
	// the event, negative values indicate notifications before the event. Examples:
	// "P1D" (1 day after), "-PT2H" (2 hours before)
	Offset string `json:"offset,required"`
	// The type of lifecycle event that this offset is based on.
	Type string `json:"type,required"`
	paramObj
}

func (r NotificationNewParamsPolicy) MarshalJSON() (data []byte, err error) {
	type shadow NotificationNewParamsPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationNewParamsPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetParams struct {
	// The ID of the notification configuration to retrieve
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r NotificationGetParams) MarshalJSON() (data []byte, err error) {
	type shadow NotificationGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationUpdateParams struct {
	// Updated policy configuration. The policy.type must match the existing lifecycle
	// event type.
	Policy NotificationUpdateParamsPolicyUnion `json:"policy,omitzero,required"`
	// The ID of the notification configuration to edit. Not provided when updating the
	// configuration for system events
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// Set to true to enable webhook messages for the notification indicated in the
	// policy, false to disable. Only supported by system lifecycle events.
	IsEnabled param.Opt[bool] `json:"is_enabled,omitzero"`
	paramObj
}

func (r NotificationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NotificationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NotificationUpdateParamsPolicyUnion struct {
	OfNotificationUpdatesPolicyLifecycleEventOffsetPolicy *NotificationUpdateParamsPolicyLifecycleEventOffsetPolicy `json:",omitzero,inline"`
	OfNotificationUpdatesPolicyLifecycleEventSystemPolicy *NotificationUpdateParamsPolicyLifecycleEventSystemPolicy `json:",omitzero,inline"`
	paramUnion
}

func (u NotificationUpdateParamsPolicyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNotificationUpdatesPolicyLifecycleEventOffsetPolicy, u.OfNotificationUpdatesPolicyLifecycleEventSystemPolicy)
}
func (u *NotificationUpdateParamsPolicyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *NotificationUpdateParamsPolicyUnion) asAny() any {
	if !param.IsOmitted(u.OfNotificationUpdatesPolicyLifecycleEventOffsetPolicy) {
		return u.OfNotificationUpdatesPolicyLifecycleEventOffsetPolicy
	} else if !param.IsOmitted(u.OfNotificationUpdatesPolicyLifecycleEventSystemPolicy) {
		return u.OfNotificationUpdatesPolicyLifecycleEventSystemPolicy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u NotificationUpdateParamsPolicyUnion) GetOffset() *string {
	if vt := u.OfNotificationUpdatesPolicyLifecycleEventOffsetPolicy; vt != nil {
		return &vt.Offset
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u NotificationUpdateParamsPolicyUnion) GetType() *string {
	if vt := u.OfNotificationUpdatesPolicyLifecycleEventOffsetPolicy; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfNotificationUpdatesPolicyLifecycleEventSystemPolicy; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The properties Offset, Type are required.
type NotificationUpdateParamsPolicyLifecycleEventOffsetPolicy struct {
	// ISO-8601 duration string indicating how much time before or after the base event
	// this notification should be sent. Positive values indicate notifications after
	// the event, negative values indicate notifications before the event. Examples:
	// "P1D" (1 day after), "-PT2H" (2 hours before)
	Offset string `json:"offset,required"`
	// The type of lifecycle event that this offset is based on.
	Type string `json:"type,required"`
	paramObj
}

func (r NotificationUpdateParamsPolicyLifecycleEventOffsetPolicy) MarshalJSON() (data []byte, err error) {
	type shadow NotificationUpdateParamsPolicyLifecycleEventOffsetPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationUpdateParamsPolicyLifecycleEventOffsetPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type NotificationUpdateParamsPolicyLifecycleEventSystemPolicy struct {
	// The type of lifecycle event (e.g., "contract.create", "contract.start")
	Type string `json:"type,required"`
	paramObj
}

func (r NotificationUpdateParamsPolicyLifecycleEventSystemPolicy) MarshalJSON() (data []byte, err error) {
	type shadow NotificationUpdateParamsPolicyLifecycleEventSystemPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationUpdateParamsPolicyLifecycleEventSystemPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationArchiveParams struct {
	// The ID of the offset lifecycle event notification configuration to archive.
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r NotificationArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow NotificationArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
