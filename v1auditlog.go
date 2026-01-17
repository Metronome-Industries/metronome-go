// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v3/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/v3/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"github.com/Metronome-Industries/metronome-go/v3/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/v3/packages/param"
	"github.com/Metronome-Industries/metronome-go/v3/packages/respjson"
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
func NewV1AuditLogService(opts ...option.RequestOption) (r V1AuditLogService) {
	r = V1AuditLogService{}
	r.Options = opts
	return
}

// Get a comprehensive audit trail of all operations performed in your Metronome
// account, whether initiated through the API, web interface, or automated
// processes. This endpoint provides detailed logs of who did what and when,
// enabling compliance reporting, security monitoring, and operational
// troubleshooting across all interaction channels.
//
// ### Use this endpoint to:
//
// - Monitor all account activity for security and compliance purposes
// - Track configuration changes regardless of source (API, UI, or system)
// - Investigate issues by reviewing historical operations
//
// ### Key response fields:
//
// An array of AuditLog objects containing:
//
//   - id: Unique identifier for the audit log entry
//   - timestamp: When the action occurred (RFC 3339 format)
//   - actor: Information about who performed the action
//   - request: Details including request ID, IP address, and user agent
//   - `resource_type`: The type of resource affected (e.g., customer, contract,
//     invoice)
//   - `resource_id`: The specific resource identifier
//   - `action`: The operation performed
//   - `next_page`: Cursor for continuous log retrieval
//
// ### Usage guidelines:
//
//   - Continuous retrieval: The next_page token enables uninterrupted log
//     streaming—save it between requests to ensure no logs are missed
//   - Empty responses: An empty data array means no new logs yet; continue polling
//     with the same next_page token
//   - Date filtering:
//   - `starting_on`: Earliest logs to return (inclusive)
//   - `ending_before`: Latest logs to return (exclusive)
//   - Cannot be used with `next_page`
//   - Resource filtering: Must specify both `resource_type` and `resource_id`
//     together
//   - Sort order: Default is `date_asc`; use `date_desc` for newest first
func (r *V1AuditLogService) List(ctx context.Context, query V1AuditLogListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1AuditLogListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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

// Get a comprehensive audit trail of all operations performed in your Metronome
// account, whether initiated through the API, web interface, or automated
// processes. This endpoint provides detailed logs of who did what and when,
// enabling compliance reporting, security monitoring, and operational
// troubleshooting across all interaction channels.
//
// ### Use this endpoint to:
//
// - Monitor all account activity for security and compliance purposes
// - Track configuration changes regardless of source (API, UI, or system)
// - Investigate issues by reviewing historical operations
//
// ### Key response fields:
//
// An array of AuditLog objects containing:
//
//   - id: Unique identifier for the audit log entry
//   - timestamp: When the action occurred (RFC 3339 format)
//   - actor: Information about who performed the action
//   - request: Details including request ID, IP address, and user agent
//   - `resource_type`: The type of resource affected (e.g., customer, contract,
//     invoice)
//   - `resource_id`: The specific resource identifier
//   - `action`: The operation performed
//   - `next_page`: Cursor for continuous log retrieval
//
// ### Usage guidelines:
//
//   - Continuous retrieval: The next_page token enables uninterrupted log
//     streaming—save it between requests to ensure no logs are missed
//   - Empty responses: An empty data array means no new logs yet; continue polling
//     with the same next_page token
//   - Date filtering:
//   - `starting_on`: Earliest logs to return (inclusive)
//   - `ending_before`: Latest logs to return (exclusive)
//   - Cannot be used with `next_page`
//   - Resource filtering: Must specify both `resource_type` and `resource_id`
//     together
//   - Sort order: Default is `date_asc`; use `date_desc` for newest first
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
	// Any of "success", "failure", "pending".
	Status V1AuditLogListResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Request      respjson.Field
		Timestamp    respjson.Field
		Action       respjson.Field
		Actor        respjson.Field
		Description  respjson.Field
		ResourceID   respjson.Field
		ResourceType respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AuditLogListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AuditLogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AuditLogListResponseRequest struct {
	ID        string `json:"id,required"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IP          respjson.Field
		UserAgent   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AuditLogListResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *V1AuditLogListResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AuditLogListResponseActor struct {
	ID    string `json:"id,required"`
	Name  string `json:"name,required"`
	Email string `json:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AuditLogListResponseActor) RawJSON() string { return r.JSON.raw }
func (r *V1AuditLogListResponseActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AuditLogListResponseStatus string

const (
	V1AuditLogListResponseStatusSuccess V1AuditLogListResponseStatus = "success"
	V1AuditLogListResponseStatusFailure V1AuditLogListResponseStatus = "failure"
	V1AuditLogListResponseStatusPending V1AuditLogListResponseStatus = "pending"
)

type V1AuditLogListParams struct {
	// RFC 3339 timestamp (exclusive). Cannot be used with 'next_page'.
	EndingBefore param.Opt[time.Time] `query:"ending_before,omitzero" format:"date-time" json:"-"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// Optional parameter that can be used to filter which audit logs are returned. If
	// you specify resource_id, you must also specify resource_type.
	ResourceID param.Opt[string] `query:"resource_id,omitzero" json:"-"`
	// Optional parameter that can be used to filter which audit logs are returned. If
	// you specify resource_type, you must also specify resource_id.
	ResourceType param.Opt[string] `query:"resource_type,omitzero" json:"-"`
	// RFC 3339 timestamp of the earliest audit log to return. Cannot be used with
	// 'next_page'.
	StartingOn param.Opt[time.Time] `query:"starting_on,omitzero" format:"date-time" json:"-"`
	// Sort order by timestamp, e.g. date_asc or date_desc. Defaults to date_asc.
	//
	// Any of "date_asc", "date_desc".
	Sort V1AuditLogListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1AuditLogListParams]'s query parameters as `url.Values`.
func (r V1AuditLogListParams) URLQuery() (v url.Values, err error) {
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
