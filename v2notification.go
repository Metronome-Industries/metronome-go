// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v2/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v2/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/packages/param"
	"github.com/Metronome-Industries/metronome-go/v2/packages/respjson"
)

// V2NotificationService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2NotificationService] method instead.
type V2NotificationService struct {
	Options []option.RequestOption
}

// NewV2NotificationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2NotificationService(opts ...option.RequestOption) (r V2NotificationService) {
	r = V2NotificationService{}
	r.Options = opts
	return
}

// Create an offset lifecycle event notification configuration. The lifecycle event
// type is inferred from the policy.type field.
func (r *V2NotificationService) New(ctx context.Context, body V2NotificationNewParams, opts ...option.RequestOption) (res *V2NotificationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific offset lifecycle event notification configuration by ID.
func (r *V2NotificationService) Get(ctx context.Context, body V2NotificationGetParams, opts ...option.RequestOption) (res *V2NotificationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Archive an offset lifecycle event notification configuration. Archived
// notifications are not processed.
func (r *V2NotificationService) Archive(ctx context.Context, body V2NotificationArchiveParams, opts ...option.RequestOption) (res *V2NotificationArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Edit an existing offset lifecycle event notification configuration.
func (r *V2NotificationService) Edit(ctx context.Context, body V2NotificationEditParams, opts ...option.RequestOption) (res *V2NotificationEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List offset lifecycle event notification configurations. These are user-created
// notifications that trigger at a specified time offset relative to lifecycle
// events. Returns a maximum of 400 results per request.
func (r *V2NotificationService) ListOffset(ctx context.Context, body V2NotificationListOffsetParams, opts ...option.RequestOption) (res *V2NotificationListOffsetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/offset/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List available system lifecycle event types for notifications. These are
// read-only event types that can be used when creating offset notifications.
func (r *V2NotificationService) ListSystem(ctx context.Context, opts ...option.RequestOption) (res *V2NotificationListSystemResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notifications/system/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type LifecycleEventOffsetNotificationConfig struct {
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
	Name   string                     `json:"name,required"`
	Policy LifecycleEventOffsetPolicy `json:"policy,required"`
	// Indicates this is an offset lifecycle event notification
	//
	// Any of "OFFSET_LIFECYCLE_EVENT".
	Type LifecycleEventOffsetNotificationConfigType `json:"type,required"`
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

// Indicates this is an offset lifecycle event notification
type LifecycleEventOffsetNotificationConfigType string

const (
	LifecycleEventOffsetNotificationConfigTypeOffsetLifecycleEvent LifecycleEventOffsetNotificationConfigType = "OFFSET_LIFECYCLE_EVENT"
)

type LifecycleEventOffsetPolicy struct {
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
func (r LifecycleEventOffsetPolicy) RawJSON() string { return r.JSON.raw }
func (r *LifecycleEventOffsetPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LifecycleEventOffsetPolicy to a
// LifecycleEventOffsetPolicyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LifecycleEventOffsetPolicyParam.Overrides()
func (r LifecycleEventOffsetPolicy) ToParam() LifecycleEventOffsetPolicyParam {
	return param.Override[LifecycleEventOffsetPolicyParam](json.RawMessage(r.RawJSON()))
}

// The properties Offset, Type are required.
type LifecycleEventOffsetPolicyParam struct {
	// ISO-8601 duration string indicating how much time before or after the base event
	// this notification should be sent. Positive values indicate notifications after
	// the event, negative values indicate notifications before the event. Examples:
	// "P1D" (1 day after), "-PT2H" (2 hours before)
	Offset string `json:"offset,required"`
	// The type of lifecycle event that this offset is based on.
	Type string `json:"type,required"`
	paramObj
}

func (r LifecycleEventOffsetPolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow LifecycleEventOffsetPolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LifecycleEventOffsetPolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LifecycleEventSystemNotificationConfig struct {
	Policy LifecycleEventSystemPolicy `json:"policy,required"`
	// Indicates this is a system lifecycle event notification
	//
	// Any of "SYSTEM_LIFECYCLE_EVENT".
	Type LifecycleEventSystemNotificationConfigType `json:"type,required"`
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

// Indicates this is a system lifecycle event notification
type LifecycleEventSystemNotificationConfigType string

const (
	LifecycleEventSystemNotificationConfigTypeSystemLifecycleEvent LifecycleEventSystemNotificationConfigType = "SYSTEM_LIFECYCLE_EVENT"
)

type LifecycleEventSystemPolicy struct {
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
func (r LifecycleEventSystemPolicy) RawJSON() string { return r.JSON.raw }
func (r *LifecycleEventSystemPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LifecycleEventSystemPolicy to a
// LifecycleEventSystemPolicyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LifecycleEventSystemPolicyParam.Overrides()
func (r LifecycleEventSystemPolicy) ToParam() LifecycleEventSystemPolicyParam {
	return param.Override[LifecycleEventSystemPolicyParam](json.RawMessage(r.RawJSON()))
}

// The property Type is required.
type LifecycleEventSystemPolicyParam struct {
	// The type of lifecycle event (e.g., "contract.create", "contract.start")
	Type string `json:"type,required"`
	paramObj
}

func (r LifecycleEventSystemPolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow LifecycleEventSystemPolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LifecycleEventSystemPolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2NotificationNewResponse struct {
	Data LifecycleEventOffsetNotificationConfig `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2NotificationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V2NotificationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2NotificationGetResponse struct {
	Data LifecycleEventOffsetNotificationConfig `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2NotificationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V2NotificationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2NotificationArchiveResponse struct {
	Data LifecycleEventOffsetNotificationConfig `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2NotificationArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V2NotificationArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2NotificationEditResponse struct {
	Data V2NotificationEditResponseDataUnion `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2NotificationEditResponse) RawJSON() string { return r.JSON.raw }
func (r *V2NotificationEditResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V2NotificationEditResponseDataUnion contains all possible properties and values
// from [LifecycleEventSystemNotificationConfig],
// [LifecycleEventOffsetNotificationConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V2NotificationEditResponseDataUnion struct {
	// This field is a union of [LifecycleEventSystemPolicy],
	// [LifecycleEventOffsetPolicy]
	Policy V2NotificationEditResponseDataUnionPolicy `json:"policy"`
	Type   string                                    `json:"type"`
	// This field is from variant [LifecycleEventSystemNotificationConfig].
	IsEnabled bool `json:"is_enabled"`
	// This field is from variant [LifecycleEventOffsetNotificationConfig].
	ID string `json:"id"`
	// This field is from variant [LifecycleEventOffsetNotificationConfig].
	ArchivedAt time.Time `json:"archived_at"`
	// This field is from variant [LifecycleEventOffsetNotificationConfig].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [LifecycleEventOffsetNotificationConfig].
	CreatedBy string `json:"created_by"`
	// This field is from variant [LifecycleEventOffsetNotificationConfig].
	EnvironmentType string `json:"environment_type"`
	// This field is from variant [LifecycleEventOffsetNotificationConfig].
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

func (u V2NotificationEditResponseDataUnion) AsLifecycleEventSystemNotificationConfig() (v LifecycleEventSystemNotificationConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V2NotificationEditResponseDataUnion) AsLifecycleEventOffsetNotificationConfig() (v LifecycleEventOffsetNotificationConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V2NotificationEditResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V2NotificationEditResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V2NotificationEditResponseDataUnionPolicy is an implicit subunion of
// [V2NotificationEditResponseDataUnion]. V2NotificationEditResponseDataUnionPolicy
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V2NotificationEditResponseDataUnion].
type V2NotificationEditResponseDataUnionPolicy struct {
	Type string `json:"type"`
	// This field is from variant [LifecycleEventOffsetPolicy].
	Offset string `json:"offset"`
	JSON   struct {
		Type   respjson.Field
		Offset respjson.Field
		raw    string
	} `json:"-"`
}

func (r *V2NotificationEditResponseDataUnionPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2NotificationListOffsetResponse struct {
	Data   []LifecycleEventOffsetNotificationConfig `json:"data,required"`
	Cursor string                                   `json:"cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Cursor      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2NotificationListOffsetResponse) RawJSON() string { return r.JSON.raw }
func (r *V2NotificationListOffsetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2NotificationListSystemResponse struct {
	Data   []LifecycleEventSystemNotificationConfig `json:"data,required"`
	Cursor string                                   `json:"cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Cursor      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2NotificationListSystemResponse) RawJSON() string { return r.JSON.raw }
func (r *V2NotificationListSystemResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2NotificationNewParams struct {
	// The name for this offset notification configuration.
	Name string `json:"name,required"`
	// The offset lifecycle event policy that defines when and how this notification
	// should be triggered. The lifecycle event type is inferred from the policy.type
	// field.
	Policy LifecycleEventOffsetPolicyParam `json:"policy,omitzero,required"`
	// Optional uniqueness key to prevent duplicate notification configurations.
	UniquenessKey param.Opt[string] `json:"uniqueness_key,omitzero"`
	paramObj
}

func (r V2NotificationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V2NotificationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2NotificationGetParams struct {
	// The ID of the notification configuration to retrieve
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r V2NotificationGetParams) MarshalJSON() (data []byte, err error) {
	type shadow V2NotificationGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2NotificationArchiveParams struct {
	// The ID of the offset lifecycle event notification configuration to archive.
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r V2NotificationArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V2NotificationArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2NotificationEditParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfObject *V2NotificationEditParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfV2NotificationEditsBodyObject *V2NotificationEditParamsBodyObject `json:",inline"`

	paramObj
}

func (u V2NotificationEditParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfObject, u.OfV2NotificationEditsBodyObject)
}
func (r *V2NotificationEditParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Policy, Type are required.
type V2NotificationEditParamsBodyObject struct {
	Policy LifecycleEventSystemPolicyParam `json:"policy,omitzero,required"`
	// Indicates this is a system lifecycle event notification
	//
	// Any of "SYSTEM_LIFECYCLE_EVENT".
	Type string `json:"type,omitzero,required"`
	// Set to true to enable webhook messages for the notification indicated in the
	// policy, false to disable. Only supported by system lifecycle events.
	IsEnabled param.Opt[bool] `json:"is_enabled,omitzero"`
	paramObj
}

func (r V2NotificationEditParamsBodyObject) MarshalJSON() (data []byte, err error) {
	type shadow V2NotificationEditParamsBodyObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationEditParamsBodyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2NotificationEditParamsBodyObject](
		"type", "SYSTEM_LIFECYCLE_EVENT",
	)
}

type V2NotificationListOffsetParams struct {
	Cursor param.Opt[string]  `json:"cursor,omitzero"`
	Limit  param.Opt[float64] `json:"limit,omitzero"`
	// Filter options for the notification configurations. If not provided, defaults to
	// NOT_ARCHIVED.
	//
	// Any of "ARCHIVED", "NOT_ARCHIVED", "ALL".
	ArchiveFilter V2NotificationListOffsetParamsArchiveFilter `json:"archive_filter,omitzero"`
	paramObj
}

func (r V2NotificationListOffsetParams) MarshalJSON() (data []byte, err error) {
	type shadow V2NotificationListOffsetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationListOffsetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for the notification configurations. If not provided, defaults to
// NOT_ARCHIVED.
type V2NotificationListOffsetParamsArchiveFilter string

const (
	V2NotificationListOffsetParamsArchiveFilterArchived    V2NotificationListOffsetParamsArchiveFilter = "ARCHIVED"
	V2NotificationListOffsetParamsArchiveFilterNotArchived V2NotificationListOffsetParamsArchiveFilter = "NOT_ARCHIVED"
	V2NotificationListOffsetParamsArchiveFilterAll         V2NotificationListOffsetParamsArchiveFilter = "ALL"
)
