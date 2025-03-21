// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
)

// V1DashboardService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1DashboardService] method instead.
type V1DashboardService struct {
	Options []option.RequestOption
}

// NewV1DashboardService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1DashboardService(opts ...option.RequestOption) (r *V1DashboardService) {
	r = &V1DashboardService{}
	r.Options = opts
	return
}

// Retrieve an embeddable dashboard url for a customer. The dashboard can be
// embedded using an iframe in a website. This will show information such as usage
// data and customer invoices.
func (r *V1DashboardService) GetEmbeddableURL(ctx context.Context, body V1DashboardGetEmbeddableURLParams, opts ...option.RequestOption) (res *V1DashboardGetEmbeddableURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/dashboards/getEmbeddableUrl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1DashboardGetEmbeddableURLResponse struct {
	Data V1DashboardGetEmbeddableURLResponseData `json:"data,required"`
	JSON v1DashboardGetEmbeddableURLResponseJSON `json:"-"`
}

// v1DashboardGetEmbeddableURLResponseJSON contains the JSON metadata for the
// struct [V1DashboardGetEmbeddableURLResponse]
type v1DashboardGetEmbeddableURLResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1DashboardGetEmbeddableURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1DashboardGetEmbeddableURLResponseJSON) RawJSON() string {
	return r.raw
}

type V1DashboardGetEmbeddableURLResponseData struct {
	URL  string                                      `json:"url"`
	JSON v1DashboardGetEmbeddableURLResponseDataJSON `json:"-"`
}

// v1DashboardGetEmbeddableURLResponseDataJSON contains the JSON metadata for the
// struct [V1DashboardGetEmbeddableURLResponseData]
type v1DashboardGetEmbeddableURLResponseDataJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1DashboardGetEmbeddableURLResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1DashboardGetEmbeddableURLResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1DashboardGetEmbeddableURLParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// The type of dashboard to retrieve.
	Dashboard param.Field[V1DashboardGetEmbeddableURLParamsDashboard] `json:"dashboard,required"`
	// Optional list of billable metric group key overrides
	BmGroupKeyOverrides param.Field[[]V1DashboardGetEmbeddableURLParamsBmGroupKeyOverride] `json:"bm_group_key_overrides"`
	// Optional list of colors to override
	ColorOverrides param.Field[[]V1DashboardGetEmbeddableURLParamsColorOverride] `json:"color_overrides"`
	// Optional dashboard specific options
	DashboardOptions param.Field[[]V1DashboardGetEmbeddableURLParamsDashboardOption] `json:"dashboard_options"`
}

func (r V1DashboardGetEmbeddableURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of dashboard to retrieve.
type V1DashboardGetEmbeddableURLParamsDashboard string

const (
	V1DashboardGetEmbeddableURLParamsDashboardInvoices V1DashboardGetEmbeddableURLParamsDashboard = "invoices"
	V1DashboardGetEmbeddableURLParamsDashboardUsage    V1DashboardGetEmbeddableURLParamsDashboard = "usage"
	V1DashboardGetEmbeddableURLParamsDashboardCredits  V1DashboardGetEmbeddableURLParamsDashboard = "credits"
)

func (r V1DashboardGetEmbeddableURLParamsDashboard) IsKnown() bool {
	switch r {
	case V1DashboardGetEmbeddableURLParamsDashboardInvoices, V1DashboardGetEmbeddableURLParamsDashboardUsage, V1DashboardGetEmbeddableURLParamsDashboardCredits:
		return true
	}
	return false
}

type V1DashboardGetEmbeddableURLParamsBmGroupKeyOverride struct {
	// The name of the billable metric group key.
	GroupKeyName param.Field[string] `json:"group_key_name,required"`
	// The display name for the billable metric group key
	DisplayName param.Field[string] `json:"display_name"`
	// <key, value> pairs of the billable metric group key values and their display
	// names. e.g. {"a": "Asia", "b": "Euro"}
	ValueDisplayNames param.Field[map[string]interface{}] `json:"value_display_names"`
}

func (r V1DashboardGetEmbeddableURLParamsBmGroupKeyOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1DashboardGetEmbeddableURLParamsColorOverride struct {
	// The color to override
	Name param.Field[V1DashboardGetEmbeddableURLParamsColorOverridesName] `json:"name"`
	// Hex value representation of the color
	Value param.Field[string] `json:"value"`
}

func (r V1DashboardGetEmbeddableURLParamsColorOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The color to override
type V1DashboardGetEmbeddableURLParamsColorOverridesName string

const (
	V1DashboardGetEmbeddableURLParamsColorOverridesNameGrayDark       V1DashboardGetEmbeddableURLParamsColorOverridesName = "Gray_dark"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameGrayMedium     V1DashboardGetEmbeddableURLParamsColorOverridesName = "Gray_medium"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameGrayLight      V1DashboardGetEmbeddableURLParamsColorOverridesName = "Gray_light"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameGrayExtralight V1DashboardGetEmbeddableURLParamsColorOverridesName = "Gray_extralight"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameWhite          V1DashboardGetEmbeddableURLParamsColorOverridesName = "White"
	V1DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryMedium  V1DashboardGetEmbeddableURLParamsColorOverridesName = "Primary_medium"
	V1DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryLight   V1DashboardGetEmbeddableURLParamsColorOverridesName = "Primary_light"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine0     V1DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_0"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine1     V1DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_1"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine2     V1DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_2"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine3     V1DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_3"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine4     V1DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_4"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine5     V1DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_5"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine6     V1DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_6"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine7     V1DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_7"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine8     V1DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_8"
	V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine9     V1DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_9"
	V1DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryGreen   V1DashboardGetEmbeddableURLParamsColorOverridesName = "Primary_green"
	V1DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryRed     V1DashboardGetEmbeddableURLParamsColorOverridesName = "Primary_red"
)

func (r V1DashboardGetEmbeddableURLParamsColorOverridesName) IsKnown() bool {
	switch r {
	case V1DashboardGetEmbeddableURLParamsColorOverridesNameGrayDark, V1DashboardGetEmbeddableURLParamsColorOverridesNameGrayMedium, V1DashboardGetEmbeddableURLParamsColorOverridesNameGrayLight, V1DashboardGetEmbeddableURLParamsColorOverridesNameGrayExtralight, V1DashboardGetEmbeddableURLParamsColorOverridesNameWhite, V1DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryMedium, V1DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryLight, V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine0, V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine1, V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine2, V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine3, V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine4, V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine5, V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine6, V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine7, V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine8, V1DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine9, V1DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryGreen, V1DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryRed:
		return true
	}
	return false
}

type V1DashboardGetEmbeddableURLParamsDashboardOption struct {
	// The option key name
	Key param.Field[string] `json:"key,required"`
	// The option value
	Value param.Field[string] `json:"value,required"`
}

func (r V1DashboardGetEmbeddableURLParamsDashboardOption) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
