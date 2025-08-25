// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
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
func NewV1ServiceService(opts ...option.RequestOption) (r *V1ServiceService) {
	r = &V1ServiceService{}
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
	opts = append(r.Options[:], opts...)
	path := "v1/services"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type V1ServiceListResponse struct {
	Services []V1ServiceListResponseService `json:"services,required"`
	JSON     v1ServiceListResponseJSON      `json:"-"`
}

// v1ServiceListResponseJSON contains the JSON metadata for the struct
// [V1ServiceListResponse]
type v1ServiceListResponseJSON struct {
	Services    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ServiceListResponseJSON) RawJSON() string {
	return r.raw
}

type V1ServiceListResponseService struct {
	IPs   []string                           `json:"ips,required"`
	Name  string                             `json:"name,required"`
	Usage V1ServiceListResponseServicesUsage `json:"usage,required"`
	JSON  v1ServiceListResponseServiceJSON   `json:"-"`
}

// v1ServiceListResponseServiceJSON contains the JSON metadata for the struct
// [V1ServiceListResponseService]
type v1ServiceListResponseServiceJSON struct {
	IPs         apijson.Field
	Name        apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ServiceListResponseService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ServiceListResponseServiceJSON) RawJSON() string {
	return r.raw
}

type V1ServiceListResponseServicesUsage string

const (
	V1ServiceListResponseServicesUsageMakesConnectionsFrom V1ServiceListResponseServicesUsage = "makes_connections_from"
	V1ServiceListResponseServicesUsageAcceptsConnectionsAt V1ServiceListResponseServicesUsage = "accepts_connections_at"
)

func (r V1ServiceListResponseServicesUsage) IsKnown() bool {
	switch r {
	case V1ServiceListResponseServicesUsageMakesConnectionsFrom, V1ServiceListResponseServicesUsageAcceptsConnectionsAt:
		return true
	}
	return false
}
