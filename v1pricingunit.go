// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/Metronome-Industries/metronome-go/v3/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v3/internal/apiquery"
	shimjson "github.com/Metronome-Industries/metronome-go/v3/internal/encoding/json"
	"github.com/Metronome-Industries/metronome-go/v3/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"github.com/Metronome-Industries/metronome-go/v3/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/v3/packages/param"
	"github.com/Metronome-Industries/metronome-go/v3/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/v3/shared"
)

// Use these endpoints to configure a billing API key, a webhook secret, or invoice
// finalization behavior.
//
// V1PricingUnitService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PricingUnitService] method instead.
type V1PricingUnitService struct {
	Options []option.RequestOption
}

// NewV1PricingUnitService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1PricingUnitService(opts ...option.RequestOption) (r V1PricingUnitService) {
	r = V1PricingUnitService{}
	r.Options = opts
	return
}

// Create a custom pricing unit. Custom pricing units can be used to charge for
// usage in a non-fiat pricing unit, for example AI credits.
func (r *V1PricingUnitService) New(ctx context.Context, body V1PricingUnitNewParams, opts ...option.RequestOption) (res *V1PricingUnitNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/credit-types/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List all pricing units. All fiat currency types (for example, USD or GBP) will
// be included, as well as any custom pricing units that were configured. Custom
// pricing units can be used to charge for usage in a non-fiat pricing unit, for
// example AI credits.
//
// Note: The USD (cents) pricing unit is 2714e483-4ff1-48e4-9e25-ac732e8f24f2.
func (r *V1PricingUnitService) List(ctx context.Context, query V1PricingUnitListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1PricingUnitListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/credit-types/list"
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

// List all pricing units. All fiat currency types (for example, USD or GBP) will
// be included, as well as any custom pricing units that were configured. Custom
// pricing units can be used to charge for usage in a non-fiat pricing unit, for
// example AI credits.
//
// Note: The USD (cents) pricing unit is 2714e483-4ff1-48e4-9e25-ac732e8f24f2.
func (r *V1PricingUnitService) ListAutoPaging(ctx context.Context, query V1PricingUnitListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1PricingUnitListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Archive a custom pricing unit. Once archived, it will no longer appear in
// pricing unit selectors by default.
func (r *V1PricingUnitService) Archive(ctx context.Context, body V1PricingUnitArchiveParams, opts ...option.RequestOption) (res *V1PricingUnitArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/credit-types/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type V1PricingUnitNewResponse struct {
	Data shared.ID `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PricingUnitNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PricingUnitNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PricingUnitListResponse struct {
	ID         string `json:"id" format:"uuid"`
	IsCurrency bool   `json:"is_currency"`
	Name       string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsCurrency  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PricingUnitListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PricingUnitListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PricingUnitArchiveResponse struct {
	Data shared.ID `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PricingUnitArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PricingUnitArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PricingUnitNewParams struct {
	// The name of the custom pricing unit. This will appear on invoices.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r V1PricingUnitNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PricingUnitNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PricingUnitNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PricingUnitListParams struct {
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1PricingUnitListParams]'s query parameters as
// `url.Values`.
func (r V1PricingUnitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PricingUnitArchiveParams struct {
	ID shared.IDParam
	paramObj
}

func (r V1PricingUnitArchiveParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ID)
}
func (r *V1PricingUnitArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
