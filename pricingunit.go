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
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
)

// PricingUnitService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingUnitService] method instead.
type PricingUnitService struct {
	Options []option.RequestOption
}

// NewPricingUnitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPricingUnitService(opts ...option.RequestOption) (r *PricingUnitService) {
	r = &PricingUnitService{}
	r.Options = opts
	return
}

// List all pricing units (known in the API by the legacy term "credit types").
func (r *PricingUnitService) List(ctx context.Context, query PricingUnitListParams, opts ...option.RequestOption) (res *pagination.CursorPage[PricingUnitListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "credit-types/list"
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

// List all pricing units (known in the API by the legacy term "credit types").
func (r *PricingUnitService) ListAutoPaging(ctx context.Context, query PricingUnitListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[PricingUnitListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type PricingUnitListResponse struct {
	ID         string                      `json:"id" format:"uuid"`
	IsCurrency bool                        `json:"is_currency"`
	Name       string                      `json:"name"`
	JSON       pricingUnitListResponseJSON `json:"-"`
}

// pricingUnitListResponseJSON contains the JSON metadata for the struct
// [PricingUnitListResponse]
type pricingUnitListResponseJSON struct {
	ID          apijson.Field
	IsCurrency  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PricingUnitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricingUnitListResponseJSON) RawJSON() string {
	return r.raw
}

type PricingUnitListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [PricingUnitListParams]'s query parameters as `url.Values`.
func (r PricingUnitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
