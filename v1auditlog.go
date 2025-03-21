// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
)

// V1AuditLogService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AuditLogService] method instead.
type V1AuditLogService struct {
	Options []option.RequestOption
}

// NewV1AuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1AuditLogService(opts ...option.RequestOption) (r *V1AuditLogService) {
	r = &V1AuditLogService{}
	r.Options = opts
	return
}

// Retrieves a range of audit logs. If no further audit logs are currently
// available, the data array will be empty. As new audit logs are created,
// subsequent requests using the same next_page value will be in the returned data
// array, ensuring a continuous and uninterrupted reading of audit logs.
func (r *V1AuditLogService) List(ctx context.Context, query V1AuditLogListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1AuditLogListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/auditLogs"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Retrieves a range of audit logs. If no further audit logs are currently
// available, the data array will be empty. As new audit logs are created,
// subsequent requests using the same next_page value will be in the returned data
// array, ensuring a continuous and uninterrupted reading of audit logs.
func (r *V1AuditLogService) ListAutoPaging(ctx context.Context, query V1AuditLogListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1AuditLogListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type V1AuditLogListResponse struct {
	ID           string                        `json:"id,required"`
	Request      V1AuditLogListResponseRequest `json:"request,required"`
	Timestamp    time.Time                     `json:"timestamp,required" format:"date-time"`
	Action       string                        `json:"action"`
	Actor        V1AuditLogListResponseActor   `json:"actor"`
	Description  string                        `json:"description"`
	ResourceID   string                        `json:"resource_id"`
	ResourceType string                        `json:"resource_type"`
	Status       V1AuditLogListResponseStatus  `json:"status"`
	JSON         v1AuditLogListResponseJSON    `json:"-"`
}

// v1AuditLogListResponseJSON contains the JSON metadata for the struct
// [V1AuditLogListResponse]
type v1AuditLogListResponseJSON struct {
	ID           apijson.Field
	Request      apijson.Field
	Timestamp    apijson.Field
	Action       apijson.Field
	Actor        apijson.Field
	Description  apijson.Field
	ResourceID   apijson.Field
	ResourceType apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1AuditLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1AuditLogListResponseJSON) RawJSON() string {
	return r.raw
}

type V1AuditLogListResponseRequest struct {
	ID        string                            `json:"id,required"`
	IP        string                            `json:"ip"`
	UserAgent string                            `json:"user_agent"`
	JSON      v1AuditLogListResponseRequestJSON `json:"-"`
}

// v1AuditLogListResponseRequestJSON contains the JSON metadata for the struct
// [V1AuditLogListResponseRequest]
type v1AuditLogListResponseRequestJSON struct {
	ID          apijson.Field
	IP          apijson.Field
	UserAgent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1AuditLogListResponseRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1AuditLogListResponseRequestJSON) RawJSON() string {
	return r.raw
}

type V1AuditLogListResponseActor struct {
	ID    string                          `json:"id,required"`
	Name  string                          `json:"name,required"`
	Email string                          `json:"email"`
	JSON  v1AuditLogListResponseActorJSON `json:"-"`
}

// v1AuditLogListResponseActorJSON contains the JSON metadata for the struct
// [V1AuditLogListResponseActor]
type v1AuditLogListResponseActorJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1AuditLogListResponseActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1AuditLogListResponseActorJSON) RawJSON() string {
	return r.raw
}

type V1AuditLogListResponseStatus string

const (
	V1AuditLogListResponseStatusSuccess V1AuditLogListResponseStatus = "success"
	V1AuditLogListResponseStatusFailure V1AuditLogListResponseStatus = "failure"
	V1AuditLogListResponseStatusPending V1AuditLogListResponseStatus = "pending"
)

func (r V1AuditLogListResponseStatus) IsKnown() bool {
	switch r {
	case V1AuditLogListResponseStatusSuccess, V1AuditLogListResponseStatusFailure, V1AuditLogListResponseStatusPending:
		return true
	}
	return false
}

type V1AuditLogListParams struct {
	// RFC 3339 timestamp (exclusive). Cannot be used with 'next_page'.
	EndingBefore param.Field[time.Time] `query:"ending_before" format:"date-time"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Optional parameter that can be used to filter which audit logs are returned. If
	// you specify resource_id, you must also specify resource_type.
	ResourceID param.Field[string] `query:"resource_id"`
	// Optional parameter that can be used to filter which audit logs are returned. If
	// you specify resource_type, you must also specify resource_id.
	ResourceType param.Field[string] `query:"resource_type"`
	// Sort order by timestamp, e.g. date_asc or date_desc. Defaults to date_asc.
	Sort param.Field[V1AuditLogListParamsSort] `query:"sort"`
	// RFC 3339 timestamp of the earliest audit log to return. Cannot be used with
	// 'next_page'.
	StartingOn param.Field[time.Time] `query:"starting_on" format:"date-time"`
}

// URLQuery serializes [V1AuditLogListParams]'s query parameters as `url.Values`.
func (r V1AuditLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order by timestamp, e.g. date_asc or date_desc. Defaults to date_asc.
type V1AuditLogListParamsSort string

const (
	V1AuditLogListParamsSortDateAsc  V1AuditLogListParamsSort = "date_asc"
	V1AuditLogListParamsSortDateDesc V1AuditLogListParamsSort = "date_desc"
)

func (r V1AuditLogListParamsSort) IsKnown() bool {
	switch r {
	case V1AuditLogListParamsSortDateAsc, V1AuditLogListParamsSortDateDesc:
		return true
	}
	return false
}
