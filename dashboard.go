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

// DashboardService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDashboardService] method instead.
type DashboardService struct {
	Options []option.RequestOption
}

// NewDashboardService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDashboardService(opts ...option.RequestOption) (r *DashboardService) {
	r = &DashboardService{}
	r.Options = opts
	return
}

// Retrieve an embeddable dashboard url for a customer. The dashboard can be
// embedded using an iframe in a website. This will show information such as usage
// data and customer invoices.
func (r *DashboardService) GetEmbeddableURL(ctx context.Context, body DashboardGetEmbeddableURLParams, opts ...option.RequestOption) (res *DashboardGetEmbeddableURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "dashboards/getEmbeddableUrl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DashboardGetEmbeddableURLResponse struct {
	Data DashboardGetEmbeddableURLResponseData `json:"data,required"`
	JSON dashboardGetEmbeddableURLResponseJSON `json:"-"`
}

// dashboardGetEmbeddableURLResponseJSON contains the JSON metadata for the struct
// [DashboardGetEmbeddableURLResponse]
type dashboardGetEmbeddableURLResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardGetEmbeddableURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dashboardGetEmbeddableURLResponseJSON) RawJSON() string {
	return r.raw
}

type DashboardGetEmbeddableURLResponseData struct {
	URL  string                                    `json:"url"`
	JSON dashboardGetEmbeddableURLResponseDataJSON `json:"-"`
}

// dashboardGetEmbeddableURLResponseDataJSON contains the JSON metadata for the
// struct [DashboardGetEmbeddableURLResponseData]
type dashboardGetEmbeddableURLResponseDataJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardGetEmbeddableURLResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dashboardGetEmbeddableURLResponseDataJSON) RawJSON() string {
	return r.raw
}

type DashboardGetEmbeddableURLParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// The type of dashboard to retrieve.
	Dashboard param.Field[DashboardGetEmbeddableURLParamsDashboard] `json:"dashboard,required"`
	// Optional list of colors to override
	ColorOverrides param.Field[[]DashboardGetEmbeddableURLParamsColorOverride] `json:"color_overrides"`
	// Optional dashboard specific options
	DashboardOptions param.Field[[]DashboardGetEmbeddableURLParamsDashboardOption] `json:"dashboard_options"`
}

func (r DashboardGetEmbeddableURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of dashboard to retrieve.
type DashboardGetEmbeddableURLParamsDashboard string

const (
	DashboardGetEmbeddableURLParamsDashboardInvoices DashboardGetEmbeddableURLParamsDashboard = "invoices"
	DashboardGetEmbeddableURLParamsDashboardUsage    DashboardGetEmbeddableURLParamsDashboard = "usage"
	DashboardGetEmbeddableURLParamsDashboardCredits  DashboardGetEmbeddableURLParamsDashboard = "credits"
)

func (r DashboardGetEmbeddableURLParamsDashboard) IsKnown() bool {
	switch r {
	case DashboardGetEmbeddableURLParamsDashboardInvoices, DashboardGetEmbeddableURLParamsDashboardUsage, DashboardGetEmbeddableURLParamsDashboardCredits:
		return true
	}
	return false
}

type DashboardGetEmbeddableURLParamsColorOverride struct {
	// The color to override
	Name param.Field[DashboardGetEmbeddableURLParamsColorOverridesName] `json:"name"`
	// Hex value representation of the color
	Value param.Field[string] `json:"value"`
}

func (r DashboardGetEmbeddableURLParamsColorOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The color to override
type DashboardGetEmbeddableURLParamsColorOverridesName string

const (
	DashboardGetEmbeddableURLParamsColorOverridesNameGrayDark       DashboardGetEmbeddableURLParamsColorOverridesName = "Gray_dark"
	DashboardGetEmbeddableURLParamsColorOverridesNameGrayMedium     DashboardGetEmbeddableURLParamsColorOverridesName = "Gray_medium"
	DashboardGetEmbeddableURLParamsColorOverridesNameGrayLight      DashboardGetEmbeddableURLParamsColorOverridesName = "Gray_light"
	DashboardGetEmbeddableURLParamsColorOverridesNameGrayExtralight DashboardGetEmbeddableURLParamsColorOverridesName = "Gray_extralight"
	DashboardGetEmbeddableURLParamsColorOverridesNameWhite          DashboardGetEmbeddableURLParamsColorOverridesName = "White"
	DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryMedium  DashboardGetEmbeddableURLParamsColorOverridesName = "Primary_medium"
	DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryLight   DashboardGetEmbeddableURLParamsColorOverridesName = "Primary_light"
	DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine0     DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_0"
	DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine1     DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_1"
	DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine2     DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_2"
	DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine3     DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_3"
	DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine4     DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_4"
	DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine5     DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_5"
	DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine6     DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_6"
	DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine7     DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_7"
	DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine8     DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_8"
	DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine9     DashboardGetEmbeddableURLParamsColorOverridesName = "UsageLine_9"
	DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryGreen   DashboardGetEmbeddableURLParamsColorOverridesName = "Primary_green"
	DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryRed     DashboardGetEmbeddableURLParamsColorOverridesName = "Primary_red"
)

func (r DashboardGetEmbeddableURLParamsColorOverridesName) IsKnown() bool {
	switch r {
	case DashboardGetEmbeddableURLParamsColorOverridesNameGrayDark, DashboardGetEmbeddableURLParamsColorOverridesNameGrayMedium, DashboardGetEmbeddableURLParamsColorOverridesNameGrayLight, DashboardGetEmbeddableURLParamsColorOverridesNameGrayExtralight, DashboardGetEmbeddableURLParamsColorOverridesNameWhite, DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryMedium, DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryLight, DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine0, DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine1, DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine2, DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine3, DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine4, DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine5, DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine6, DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine7, DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine8, DashboardGetEmbeddableURLParamsColorOverridesNameUsageLine9, DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryGreen, DashboardGetEmbeddableURLParamsColorOverridesNamePrimaryRed:
		return true
	}
	return false
}

type DashboardGetEmbeddableURLParamsDashboardOption struct {
	// The option key name
	Key param.Field[string] `json:"key,required"`
	// The option value
	Value param.Field[string] `json:"value,required"`
}

func (r DashboardGetEmbeddableURLParamsDashboardOption) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
