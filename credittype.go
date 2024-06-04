// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
)

// CreditTypeService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditTypeService] method instead.
type CreditTypeService struct {
	Options []option.RequestOption
}

// NewCreditTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCreditTypeService(opts ...option.RequestOption) (r *CreditTypeService) {
	r = &CreditTypeService{}
	r.Options = opts
	return
}

// List all pricing units (known in the API by the legacy term "credit types").
func (r *CreditTypeService) List(ctx context.Context, query CreditTypeListParams, opts ...option.RequestOption) (res *CreditTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "credit-types/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CreditTypeListResponse struct {
	Data     []CreditTypeListResponseData `json:"data,required"`
	NextPage string                       `json:"next_page,required,nullable"`
	JSON     creditTypeListResponseJSON   `json:"-"`
}

// creditTypeListResponseJSON contains the JSON metadata for the struct
// [CreditTypeListResponse]
type creditTypeListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditTypeListResponseJSON) RawJSON() string {
	return r.raw
}

type CreditTypeListResponseData struct {
	ID         string                         `json:"id" format:"uuid"`
	IsCurrency bool                           `json:"is_currency"`
	Name       string                         `json:"name"`
	JSON       creditTypeListResponseDataJSON `json:"-"`
}

// creditTypeListResponseDataJSON contains the JSON metadata for the struct
// [CreditTypeListResponseData]
type creditTypeListResponseDataJSON struct {
	ID          apijson.Field
	IsCurrency  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditTypeListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditTypeListResponseDataJSON) RawJSON() string {
	return r.raw
}

type CreditTypeListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [CreditTypeListParams]'s query parameters as `url.Values`.
func (r CreditTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
