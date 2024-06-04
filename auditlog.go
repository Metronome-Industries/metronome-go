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
func (r *AuditLogService) List(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) (res *AuditLogListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "auditLogs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AuditLogListResponse struct {
	Data []AuditLogListResponseData `json:"data,required"`
	// The next_page parameter is always returned to support ongoing log retrieval. It
	// enables continuous querying, even when some requests return no new data. Save
	// the next_page token from each response and use it for future requests to ensure
	// no logs are missed. This setup is ideal for regular updates via automated
	// processes, like cron jobs, to fetch logs continuously as they become available.
	// When you receive an empty data array, it indicates a temporary absence of new
	// logs, but subsequent requests might return new data.
	NextPage string                   `json:"next_page,required,nullable"`
	JSON     auditLogListResponseJSON `json:"-"`
}

// auditLogListResponseJSON contains the JSON metadata for the struct
// [AuditLogListResponse]
type auditLogListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogListResponseJSON) RawJSON() string {
	return r.raw
}

type AuditLogListResponseData struct {
	ID           string                         `json:"id,required"`
	Timestamp    time.Time                      `json:"timestamp,required" format:"date-time"`
	Action       string                         `json:"action"`
	Actor        AuditLogListResponseDataActor  `json:"actor"`
	Description  string                         `json:"description"`
	ResourceID   string                         `json:"resource_id"`
	ResourceType string                         `json:"resource_type"`
	Status       AuditLogListResponseDataStatus `json:"status"`
	JSON         auditLogListResponseDataJSON   `json:"-"`
}

// auditLogListResponseDataJSON contains the JSON metadata for the struct
// [AuditLogListResponseData]
type auditLogListResponseDataJSON struct {
	ID           apijson.Field
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

func (r *AuditLogListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogListResponseDataJSON) RawJSON() string {
	return r.raw
}

type AuditLogListResponseDataActor struct {
	ID    string                            `json:"id,required"`
	Name  string                            `json:"name,required"`
	Email string                            `json:"email"`
	JSON  auditLogListResponseDataActorJSON `json:"-"`
}

// auditLogListResponseDataActorJSON contains the JSON metadata for the struct
// [AuditLogListResponseDataActor]
type auditLogListResponseDataActorJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogListResponseDataActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogListResponseDataActorJSON) RawJSON() string {
	return r.raw
}

type AuditLogListResponseDataStatus string

const (
	AuditLogListResponseDataStatusSuccess AuditLogListResponseDataStatus = "success"
	AuditLogListResponseDataStatusFailure AuditLogListResponseDataStatus = "failure"
	AuditLogListResponseDataStatusPending AuditLogListResponseDataStatus = "pending"
)

func (r AuditLogListResponseDataStatus) IsKnown() bool {
	switch r {
	case AuditLogListResponseDataStatusSuccess, AuditLogListResponseDataStatusFailure, AuditLogListResponseDataStatusPending:
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
