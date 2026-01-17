// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"

	"github.com/Metronome-Industries/metronome-go/v3/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v3/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"github.com/Metronome-Industries/metronome-go/v3/packages/respjson"
)

// V1ServiceService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ServiceService] method instead.
type V1ServiceService struct {
	Options []option.RequestOption
}

// NewV1ServiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ServiceService(opts ...option.RequestOption) (r V1ServiceService) {
	r = V1ServiceService{}
	r.Options = opts
	return
}

// Gets Metronome's service registry with associated IP addresses for security
// allowlisting and firewall configuration. Use this endpoint to maintain an
// up-to-date list of IPs that your systems should trust for Metronome
// communications. Returns service names and their current IP ranges, with new IPs
// typically appearing 30+ days before first use to ensure smooth allowlist
// updates.
func (r *V1ServiceService) List(ctx context.Context, opts ...option.RequestOption) (res *V1ServiceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/services"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type V1ServiceListResponse struct {
	Services []V1ServiceListResponseService `json:"services,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Services    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ServiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ServiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ServiceListResponseService struct {
	IPs  []string `json:"ips,required"`
	Name string   `json:"name,required"`
	// Any of "makes_connections_from", "accepts_connections_at".
	Usage string `json:"usage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPs         respjson.Field
		Name        respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ServiceListResponseService) RawJSON() string { return r.JSON.raw }
func (r *V1ServiceListResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
