// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
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
func NewV1DashboardService(opts ...option.RequestOption) (r V1DashboardService) {
	r = V1DashboardService{}
	r.Options = opts
	return
}

// Generate secure, embeddable dashboard URLs that allow you to seamlessly
// integrate Metronome's billing visualizations directly into your application.
// This endpoint creates authenticated iframe-ready URLs for customer-specific
// dashboards, providing a white-labeled billing experience without building custom
// UI.
//
// ### Use this endpoint to:
//
// - Embed billing dashboards directly in your customer portal or admin interface
// - Provide self-service access to invoices, usage data, and credit balances
// - Build white-labeled billing experiences with minimal development effort
//
// ### Key response fields:
//
// - A secure, time-limited URL that can be embedded in an iframe
// - The URL includes authentication tokens and configuration parameters
// - URLs are customer-specific and respect your security settings
//
// ### Usage guidelines:
//
// - Dashboard types: Choose from `invoices`, `usage`, or `commits_and_credits`
// - Customization options:
//   - `dashboard_options`: Configure whether you want invoices to show zero usage
//     line items
//   - `color_overrides`: Match your brand's color palette
//   - `bm_group_key_overrides`: Customize how dimensions are displayed (for the
//     usage embeddable dashboard)
//
// - Iframe implementation: Embed the returned URL directly in an iframe element
// - Responsive design: Dashboards automatically adapt to container dimensions
func (r *V1DashboardService) GetEmbeddableURL(ctx context.Context, body V1DashboardGetEmbeddableURLParams, opts ...option.RequestOption) (res *V1DashboardGetEmbeddableURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/dashboards/getEmbeddableUrl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1DashboardGetEmbeddableURLResponse struct {
	Data V1DashboardGetEmbeddableURLResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DashboardGetEmbeddableURLResponse) RawJSON() string { return r.JSON.raw }
func (r *V1DashboardGetEmbeddableURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DashboardGetEmbeddableURLResponseData struct {
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DashboardGetEmbeddableURLResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1DashboardGetEmbeddableURLResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DashboardGetEmbeddableURLParams struct {
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// The type of dashboard to retrieve.
	//
	// Any of "invoices", "usage", "credits", "commits_and_credits".
	Dashboard V1DashboardGetEmbeddableURLParamsDashboard `json:"dashboard,omitzero,required"`
	// Optional list of billable metric group key overrides
	BmGroupKeyOverrides []V1DashboardGetEmbeddableURLParamsBmGroupKeyOverride `json:"bm_group_key_overrides,omitzero"`
	// Optional list of colors to override
	ColorOverrides []V1DashboardGetEmbeddableURLParamsColorOverride `json:"color_overrides,omitzero"`
	// Optional dashboard specific options
	DashboardOptions []V1DashboardGetEmbeddableURLParamsDashboardOption `json:"dashboard_options,omitzero"`
	paramObj
}

func (r V1DashboardGetEmbeddableURLParams) MarshalJSON() (data []byte, err error) {
	type shadow V1DashboardGetEmbeddableURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1DashboardGetEmbeddableURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of dashboard to retrieve.
type V1DashboardGetEmbeddableURLParamsDashboard string

const (
	V1DashboardGetEmbeddableURLParamsDashboardInvoices          V1DashboardGetEmbeddableURLParamsDashboard = "invoices"
	V1DashboardGetEmbeddableURLParamsDashboardUsage             V1DashboardGetEmbeddableURLParamsDashboard = "usage"
	V1DashboardGetEmbeddableURLParamsDashboardCredits           V1DashboardGetEmbeddableURLParamsDashboard = "credits"
	V1DashboardGetEmbeddableURLParamsDashboardCommitsAndCredits V1DashboardGetEmbeddableURLParamsDashboard = "commits_and_credits"
)

// The property GroupKeyName is required.
type V1DashboardGetEmbeddableURLParamsBmGroupKeyOverride struct {
	// The name of the billable metric group key.
	GroupKeyName string `json:"group_key_name,required"`
	// The display name for the billable metric group key
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// <key, value> pairs of the billable metric group key values and their display
	// names. e.g. {"a": "Asia", "b": "Euro"}
	ValueDisplayNames map[string]any `json:"value_display_names,omitzero"`
	paramObj
}

func (r V1DashboardGetEmbeddableURLParamsBmGroupKeyOverride) MarshalJSON() (data []byte, err error) {
	type shadow V1DashboardGetEmbeddableURLParamsBmGroupKeyOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1DashboardGetEmbeddableURLParamsBmGroupKeyOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DashboardGetEmbeddableURLParamsColorOverride struct {
	// Hex value representation of the color
	Value param.Opt[string] `json:"value,omitzero"`
	// The color to override
	//
	// Any of "Gray_dark", "Gray_medium", "Gray_light", "Gray_extralight", "White",
	// "Primary_medium", "Primary_light", "UsageLine_0", "UsageLine_1", "UsageLine_2",
	// "UsageLine_3", "UsageLine_4", "UsageLine_5", "UsageLine_6", "UsageLine_7",
	// "UsageLine_8", "UsageLine_9", "Primary_green", "Primary_red", "Progress_bar",
	// "Progress_bar_background".
	Name string `json:"name,omitzero"`
	paramObj
}

func (r V1DashboardGetEmbeddableURLParamsColorOverride) MarshalJSON() (data []byte, err error) {
	type shadow V1DashboardGetEmbeddableURLParamsColorOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1DashboardGetEmbeddableURLParamsColorOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1DashboardGetEmbeddableURLParamsColorOverride](
		"name", "Gray_dark", "Gray_medium", "Gray_light", "Gray_extralight", "White", "Primary_medium", "Primary_light", "UsageLine_0", "UsageLine_1", "UsageLine_2", "UsageLine_3", "UsageLine_4", "UsageLine_5", "UsageLine_6", "UsageLine_7", "UsageLine_8", "UsageLine_9", "Primary_green", "Primary_red", "Progress_bar", "Progress_bar_background",
	)
}

// The properties Key, Value are required.
type V1DashboardGetEmbeddableURLParamsDashboardOption struct {
	// The option key name
	Key string `json:"key,required"`
	// The option value
	Value string `json:"value,required"`
	paramObj
}

func (r V1DashboardGetEmbeddableURLParamsDashboardOption) MarshalJSON() (data []byte, err error) {
	type shadow V1DashboardGetEmbeddableURLParamsDashboardOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1DashboardGetEmbeddableURLParamsDashboardOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
