// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
)

// ServiceService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceService] method instead.
type ServiceService struct {
	Options []option.RequestOption
}

// NewServiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewServiceService(opts ...option.RequestOption) (r *ServiceService) {
	r = &ServiceService{}
	r.Options = opts
	return
}

// Fetches a list of services used by Metronome and the associated IP addresses. IP
// addresses are not necessarily unique between services. In most cases, IP
// addresses will appear in the list at least 30 days before they are used for the
// first time.
func (r *ServiceService) List(ctx context.Context, opts ...option.RequestOption) (res *ServiceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "services"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ServiceListResponse struct {
	Services []ServiceListResponseService `json:"services,required"`
	JSON     serviceListResponseJSON      `json:"-"`
}

// serviceListResponseJSON contains the JSON metadata for the struct
// [ServiceListResponse]
type serviceListResponseJSON struct {
	Services    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseJSON) RawJSON() string {
	return r.raw
}

type ServiceListResponseService struct {
	IPs   []string                         `json:"ips,required"`
	Name  string                           `json:"name,required"`
	Usage ServiceListResponseServicesUsage `json:"usage,required"`
	JSON  serviceListResponseServiceJSON   `json:"-"`
}

// serviceListResponseServiceJSON contains the JSON metadata for the struct
// [ServiceListResponseService]
type serviceListResponseServiceJSON struct {
	IPs         apijson.Field
	Name        apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponseService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseServiceJSON) RawJSON() string {
	return r.raw
}

type ServiceListResponseServicesUsage string

const (
	ServiceListResponseServicesUsageMakesConnectionsFrom ServiceListResponseServicesUsage = "makes_connections_from"
	ServiceListResponseServicesUsageAcceptsConnectionsAt ServiceListResponseServicesUsage = "accepts_connections_at"
)

func (r ServiceListResponseServicesUsage) IsKnown() bool {
	switch r {
	case ServiceListResponseServicesUsageMakesConnectionsFrom, ServiceListResponseServicesUsageAcceptsConnectionsAt:
		return true
	}
	return false
}
