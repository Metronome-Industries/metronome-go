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

// AuditLogService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditLogService] method instead.
type AuditLogService struct {
	Options []option.RequestOption
}

// NewAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditLogService(opts ...option.RequestOption) (r *AuditLogService) {
	r = &AuditLogService{}
	r.Options = opts
	return
}

// Retrieves a range of audit logs. If no further audit logs are currently
// available, the data array will be empty. As new audit logs are created,
// subsequent requests using the same next_page value will be in the returned data
// array, ensuring a continuous and uninterrupted reading of audit logs.
func (r *AuditLogService) List(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) (res *pagination.CursorPage[AuditLogListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "auditLogs"
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
func (r *AuditLogService) ListAutoPaging(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AuditLogListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type AuditLogListResponse struct {
	ID           string                      `json:"id,required"`
	Request      AuditLogListResponseRequest `json:"request,required"`
	Timestamp    time.Time                   `json:"timestamp,required" format:"date-time"`
	Action       string                      `json:"action"`
	Actor        AuditLogListResponseActor   `json:"actor"`
	Description  string                      `json:"description"`
	ResourceID   string                      `json:"resource_id"`
	ResourceType string                      `json:"resource_type"`
	Status       AuditLogListResponseStatus  `json:"status"`
	JSON         auditLogListResponseJSON    `json:"-"`
}

// auditLogListResponseJSON contains the JSON metadata for the struct
// [AuditLogListResponse]
type auditLogListResponseJSON struct {
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

func (r *AuditLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogListResponseJSON) RawJSON() string {
	return r.raw
}

type AuditLogListResponseRequest struct {
	ID        string                          `json:"id,required"`
	IP        string                          `json:"ip"`
	UserAgent string                          `json:"user_agent"`
	JSON      auditLogListResponseRequestJSON `json:"-"`
}

// auditLogListResponseRequestJSON contains the JSON metadata for the struct
// [AuditLogListResponseRequest]
type auditLogListResponseRequestJSON struct {
	ID          apijson.Field
	IP          apijson.Field
	UserAgent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogListResponseRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogListResponseRequestJSON) RawJSON() string {
	return r.raw
}

type AuditLogListResponseActor struct {
	ID    string                        `json:"id,required"`
	Name  string                        `json:"name,required"`
	Email string                        `json:"email"`
	JSON  auditLogListResponseActorJSON `json:"-"`
}

// auditLogListResponseActorJSON contains the JSON metadata for the struct
// [AuditLogListResponseActor]
type auditLogListResponseActorJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogListResponseActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogListResponseActorJSON) RawJSON() string {
	return r.raw
}

type AuditLogListResponseStatus string

const (
	AuditLogListResponseStatusSuccess AuditLogListResponseStatus = "success"
	AuditLogListResponseStatusFailure AuditLogListResponseStatus = "failure"
	AuditLogListResponseStatusPending AuditLogListResponseStatus = "pending"
)

func (r AuditLogListResponseStatus) IsKnown() bool {
	switch r {
	case AuditLogListResponseStatusSuccess, AuditLogListResponseStatusFailure, AuditLogListResponseStatusPending:
		return true
	}
	return false
}

type AuditLogListParams struct {
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
	Sort param.Field[AuditLogListParamsSort] `query:"sort"`
	// RFC 3339 timestamp of the earliest audit log to return. Cannot be used with
	// 'next_page'.
	StartingOn param.Field[time.Time] `query:"starting_on" format:"date-time"`
}

// URLQuery serializes [AuditLogListParams]'s query parameters as `url.Values`.
func (r AuditLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order by timestamp, e.g. date_asc or date_desc. Defaults to date_asc.
type AuditLogListParamsSort string

const (
	AuditLogListParamsSortDateAsc  AuditLogListParamsSort = "date_asc"
	AuditLogListParamsSortDateDesc AuditLogListParamsSort = "date_desc"
)

func (r AuditLogListParamsSort) IsKnown() bool {
	switch r {
	case AuditLogListParamsSortDateAsc, AuditLogListParamsSortDateDesc:
		return true
	}
	return false
}
