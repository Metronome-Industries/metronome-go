// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
  "context"
  "encoding/json"
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

// V2NotificationOffsetService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2NotificationOffsetService] method instead.
type V2NotificationOffsetService struct {
Options []option.RequestOption
}

// NewV2NotificationOffsetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2NotificationOffsetService(opts ...option.RequestOption) (r V2NotificationOffsetService) {
  r = V2NotificationOffsetService{}
  r.Options = opts
  return
}

// Create an offset lifecycle event notification configuration. The lifecycle event
// type is inferred from the policy.type field.
func (r *V2NotificationOffsetService) New(ctx context.Context, body V2NotificationOffsetNewParams, opts ...option.RequestOption) (res *V2NotificationOffsetNewResponse, err error) {
  opts = slices.Concat(r.Options, opts)
  path := "v2/notifications/create"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return res, err
}

// Retrieve a specific offset lifecycle event notification configuration by ID.
func (r *V2NotificationOffsetService) Get(ctx context.Context, body V2NotificationOffsetGetParams, opts ...option.RequestOption) (res *V2NotificationOffsetGetResponse, err error) {
  opts = slices.Concat(r.Options, opts)
  path := "v2/notifications/get"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return res, err
}

// List offset lifecycle event notification configurations. These are user-created
// notifications that trigger at a specified time offset relative to lifecycle
// events. Returns a maximum of 400 results per request.
func (r *V2NotificationOffsetService) List(ctx context.Context, body V2NotificationOffsetListParams, opts ...option.RequestOption) (res *pagination.BodyCursorPageCursorField[LifecycleEventOffsetNotificationConfig], err error) {
  var raw *http.Response
  opts = slices.Concat(r.Options, opts)
  opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
  path := "v2/notifications/offset/list"
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

// List offset lifecycle event notification configurations. These are user-created
// notifications that trigger at a specified time offset relative to lifecycle
// events. Returns a maximum of 400 results per request.
func (r *V2NotificationOffsetService) ListAutoPaging(ctx context.Context, body V2NotificationOffsetListParams, opts ...option.RequestOption) (*pagination.BodyCursorPageCursorFieldAutoPager[LifecycleEventOffsetNotificationConfig]) {
  return pagination.NewBodyCursorPageCursorFieldAutoPager(r.List(ctx, body, opts...))
}

// Archive an offset lifecycle event notification configuration. Archived
// notifications are not processed.
func (r *V2NotificationOffsetService) Archive(ctx context.Context, body V2NotificationOffsetArchiveParams, opts ...option.RequestOption) (res *V2NotificationOffsetArchiveResponse, err error) {
  opts = slices.Concat(r.Options, opts)
  path := "v2/notifications/archive"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return res, err
}

// Edit an existing offset lifecycle event notification configuration.
func (r *V2NotificationOffsetService) Edit(ctx context.Context, body V2NotificationOffsetEditParams, opts ...option.RequestOption) (res *V2NotificationOffsetEditResponse, err error) {
  opts = slices.Concat(r.Options, opts)
  path := "v2/notifications/edit"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return res, err
}

type V2NotificationOffsetNewResponse struct {
Data LifecycleEventOffsetNotificationConfig `json:"data" api:"required"`
// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
JSON struct {
              Data respjson.Field
              ExtraFields map[string]respjson.Field
              raw string
            } `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2NotificationOffsetNewResponse) RawJSON() (string) { return r.JSON.raw }
func (r *V2NotificationOffsetNewResponse) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

type V2NotificationOffsetGetResponse struct {
Data LifecycleEventOffsetNotificationConfig `json:"data" api:"required"`
// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
JSON struct {
              Data respjson.Field
              ExtraFields map[string]respjson.Field
              raw string
            } `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2NotificationOffsetGetResponse) RawJSON() (string) { return r.JSON.raw }
func (r *V2NotificationOffsetGetResponse) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

type V2NotificationOffsetArchiveResponse struct {
Data LifecycleEventOffsetNotificationConfig `json:"data" api:"required"`
// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
JSON struct {
              Data respjson.Field
              ExtraFields map[string]respjson.Field
              raw string
            } `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2NotificationOffsetArchiveResponse) RawJSON() (string) { return r.JSON.raw }
func (r *V2NotificationOffsetArchiveResponse) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

type V2NotificationOffsetEditResponse struct {
Data V2NotificationOffsetEditResponseDataUnion `json:"data" api:"required"`
// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
JSON struct {
              Data respjson.Field
              ExtraFields map[string]respjson.Field
              raw string
            } `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2NotificationOffsetEditResponse) RawJSON() (string) { return r.JSON.raw }
func (r *V2NotificationOffsetEditResponse) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

// V2NotificationOffsetEditResponseDataUnion contains all possible properties and
// values from [LifecycleEventSystemNotificationConfig],
// [LifecycleEventOffsetNotificationConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V2NotificationOffsetEditResponseDataUnion struct {
// This field is a union of [LifecycleEventSystemNotificationConfigPolicy],
// [LifecycleEventOffsetNotificationConfigPolicy]
Policy V2NotificationOffsetEditResponseDataUnionPolicy `json:"policy"`
Type string `json:"type"`
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
JSON struct { Policy respjson.Field
Type respjson.Field
IsEnabled respjson.Field
ID respjson.Field
ArchivedAt respjson.Field
CreatedAt respjson.Field
CreatedBy respjson.Field
EnvironmentType respjson.Field
Name respjson.Field
raw string } `json:"-"`
}

func (u V2NotificationOffsetEditResponseDataUnion) AsLifecycleEventSystemNotificationConfig() (v LifecycleEventSystemNotificationConfig) {
  apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
  return
}

func (u V2NotificationOffsetEditResponseDataUnion) AsLifecycleEventOffsetNotificationConfig() (v LifecycleEventOffsetNotificationConfig) {
  apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
  return
}

// Returns the unmodified JSON received from the API
func (u V2NotificationOffsetEditResponseDataUnion) RawJSON() (string) { return u.JSON.raw }

func (r *V2NotificationOffsetEditResponseDataUnion) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

// V2NotificationOffsetEditResponseDataUnionPolicy is an implicit subunion of
// [V2NotificationOffsetEditResponseDataUnion].
// V2NotificationOffsetEditResponseDataUnionPolicy provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V2NotificationOffsetEditResponseDataUnion].
type V2NotificationOffsetEditResponseDataUnionPolicy struct {
Type string `json:"type"`
// This field is from variant [LifecycleEventOffsetNotificationConfigPolicy].
Offset string `json:"offset"`
JSON struct { Type respjson.Field
Offset respjson.Field
raw string } `json:"-"`
}

func (r *V2NotificationOffsetEditResponseDataUnionPolicy) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

type V2NotificationOffsetNewParams struct {
// The name for this offset notification configuration.
Name string `json:"name" api:"required"`
// The offset lifecycle event policy that defines when and how this notification
// should be triggered. The lifecycle event type is inferred from the policy.type
// field.
Policy V2NotificationOffsetNewParamsPolicy `json:"policy,omitzero" api:"required"`
// Optional uniqueness key to prevent duplicate notification configurations.
UniquenessKey param.Opt[string] `json:"uniqueness_key,omitzero"`
paramObj
}

func (r V2NotificationOffsetNewParams) MarshalJSON() (data []byte, err error) {
  type shadow V2NotificationOffsetNewParams
  return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationOffsetNewParams) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

// The offset lifecycle event policy that defines when and how this notification
// should be triggered. The lifecycle event type is inferred from the policy.type
// field.
//
// The properties Offset, Type are required.
type V2NotificationOffsetNewParamsPolicy struct {
// ISO-8601 duration string indicating how much time before or after the base event
// this notification should be sent. Positive values indicate notifications after
// the event, negative values indicate notifications before the event. Examples:
// "P1D" (1 day after), "-PT2H" (2 hours before)
Offset string `json:"offset" api:"required"`
// The type of lifecycle event that this offset is based on.
Type string `json:"type" api:"required"`
paramObj
}

func (r V2NotificationOffsetNewParamsPolicy) MarshalJSON() (data []byte, err error) {
  type shadow V2NotificationOffsetNewParamsPolicy
  return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationOffsetNewParamsPolicy) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

type V2NotificationOffsetGetParams struct {
// The ID of the notification configuration to retrieve
ID string `json:"id" api:"required" format:"uuid"`
paramObj
}

func (r V2NotificationOffsetGetParams) MarshalJSON() (data []byte, err error) {
  type shadow V2NotificationOffsetGetParams
  return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationOffsetGetParams) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

type V2NotificationOffsetListParams struct {
Cursor param.Opt[string] `json:"cursor,omitzero"`
Limit param.Opt[float64] `json:"limit,omitzero"`
// Filter options for the notification configurations. If not provided, defaults to
// NOT_ARCHIVED.
//
// Any of "ARCHIVED", "NOT_ARCHIVED", "ALL".
ArchiveFilter V2NotificationOffsetListParamsArchiveFilter `json:"archive_filter,omitzero"`
paramObj
}

func (r V2NotificationOffsetListParams) MarshalJSON() (data []byte, err error) {
  type shadow V2NotificationOffsetListParams
  return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationOffsetListParams) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

// Filter options for the notification configurations. If not provided, defaults to
// NOT_ARCHIVED.
type V2NotificationOffsetListParamsArchiveFilter string

const (
    V2NotificationOffsetListParamsArchiveFilterArchived V2NotificationOffsetListParamsArchiveFilter = "ARCHIVED"
    V2NotificationOffsetListParamsArchiveFilterNotArchived V2NotificationOffsetListParamsArchiveFilter = "NOT_ARCHIVED"
    V2NotificationOffsetListParamsArchiveFilterAll V2NotificationOffsetListParamsArchiveFilter = "ALL"
  )

type V2NotificationOffsetArchiveParams struct {
// The ID of the offset lifecycle event notification configuration to archive.
ID string `json:"id" api:"required" format:"uuid"`
paramObj
}

func (r V2NotificationOffsetArchiveParams) MarshalJSON() (data []byte, err error) {
  type shadow V2NotificationOffsetArchiveParams
  return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationOffsetArchiveParams) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

type V2NotificationOffsetEditParams struct {
// Updated policy configuration. The policy.type must match the existing lifecycle
// event type.
Policy V2NotificationOffsetEditParamsPolicyUnion `json:"policy,omitzero" api:"required"`
// The ID of the notification configuration to edit. Not provided when updating the
// configuration for system events
ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
// Set to true to enable webhook messages for the notification indicated in the
// policy, false to disable. Only supported by system lifecycle events.
IsEnabled param.Opt[bool] `json:"is_enabled,omitzero"`
paramObj
}

func (r V2NotificationOffsetEditParams) MarshalJSON() (data []byte, err error) {
  type shadow V2NotificationOffsetEditParams
  return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationOffsetEditParams) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V2NotificationOffsetEditParamsPolicyUnion struct {
OfV2NotificationOffsetEditsPolicyLifecycleEventOffsetPolicy *V2NotificationOffsetEditParamsPolicyLifecycleEventOffsetPolicy `json:",omitzero,inline"`
OfV2NotificationOffsetEditsPolicyLifecycleEventSystemPolicy *V2NotificationOffsetEditParamsPolicyLifecycleEventSystemPolicy `json:",omitzero,inline"`
paramUnion
}

func (u V2NotificationOffsetEditParamsPolicyUnion) MarshalJSON() ([]byte, error) {
  return param.MarshalUnion(u, u.OfV2NotificationOffsetEditsPolicyLifecycleEventOffsetPolicy, u.OfV2NotificationOffsetEditsPolicyLifecycleEventSystemPolicy)
}
func (u *V2NotificationOffsetEditParamsPolicyUnion) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, u)
}

func (u *V2NotificationOffsetEditParamsPolicyUnion) asAny() (any) {
  if !param.IsOmitted(u.OfV2NotificationOffsetEditsPolicyLifecycleEventOffsetPolicy) {
    return u.OfV2NotificationOffsetEditsPolicyLifecycleEventOffsetPolicy
  } else if  !param.IsOmitted(u.OfV2NotificationOffsetEditsPolicyLifecycleEventSystemPolicy) {
    return u.OfV2NotificationOffsetEditsPolicyLifecycleEventSystemPolicy
  }
  return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V2NotificationOffsetEditParamsPolicyUnion) GetOffset() (*string) {
  if vt := u.OfV2NotificationOffsetEditsPolicyLifecycleEventOffsetPolicy; vt != nil {
    return &vt.Offset
  }
  return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V2NotificationOffsetEditParamsPolicyUnion) GetType() (*string) {
  if vt := u.OfV2NotificationOffsetEditsPolicyLifecycleEventOffsetPolicy; vt != nil {
    return (*string)(&vt.Type)
    } else if vt := u.OfV2NotificationOffsetEditsPolicyLifecycleEventSystemPolicy; vt != nil {
      return (*string)(&vt.Type)
  }
  return nil
}

// The properties Offset, Type are required.
type V2NotificationOffsetEditParamsPolicyLifecycleEventOffsetPolicy struct {
// ISO-8601 duration string indicating how much time before or after the base event
// this notification should be sent. Positive values indicate notifications after
// the event, negative values indicate notifications before the event. Examples:
// "P1D" (1 day after), "-PT2H" (2 hours before)
Offset string `json:"offset" api:"required"`
// The type of lifecycle event that this offset is based on.
Type string `json:"type" api:"required"`
paramObj
}

func (r V2NotificationOffsetEditParamsPolicyLifecycleEventOffsetPolicy) MarshalJSON() (data []byte, err error) {
  type shadow V2NotificationOffsetEditParamsPolicyLifecycleEventOffsetPolicy
  return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationOffsetEditParamsPolicyLifecycleEventOffsetPolicy) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type V2NotificationOffsetEditParamsPolicyLifecycleEventSystemPolicy struct {
// The type of lifecycle event (e.g., "contract.create", "contract.start")
Type string `json:"type" api:"required"`
paramObj
}

func (r V2NotificationOffsetEditParamsPolicyLifecycleEventSystemPolicy) MarshalJSON() (data []byte, err error) {
  type shadow V2NotificationOffsetEditParamsPolicyLifecycleEventSystemPolicy
  return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2NotificationOffsetEditParamsPolicyLifecycleEventSystemPolicy) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}
