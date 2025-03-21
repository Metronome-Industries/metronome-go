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

// V1InvoiceService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InvoiceService] method instead.
type V1InvoiceService struct {
	Options []option.RequestOption
}

// NewV1InvoiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1InvoiceService(opts ...option.RequestOption) (r *V1InvoiceService) {
	r = &V1InvoiceService{}
	r.Options = opts
	return
}

// Regenerate a voided contract invoice
func (r *V1InvoiceService) Regenerate(ctx context.Context, body V1InvoiceRegenerateParams, opts ...option.RequestOption) (res *V1InvoiceRegenerateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/invoices/regenerate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Void an invoice
func (r *V1InvoiceService) Void(ctx context.Context, body V1InvoiceVoidParams, opts ...option.RequestOption) (res *V1InvoiceVoidResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/invoices/void"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1InvoiceRegenerateResponse struct {
	Data V1InvoiceRegenerateResponseData `json:"data"`
	JSON v1InvoiceRegenerateResponseJSON `json:"-"`
}

// v1InvoiceRegenerateResponseJSON contains the JSON metadata for the struct
// [V1InvoiceRegenerateResponse]
type v1InvoiceRegenerateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1InvoiceRegenerateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1InvoiceRegenerateResponseJSON) RawJSON() string {
	return r.raw
}

type V1InvoiceRegenerateResponseData struct {
	// The new invoice id
	ID   string                              `json:"id,required" format:"uuid"`
	JSON v1InvoiceRegenerateResponseDataJSON `json:"-"`
}

// v1InvoiceRegenerateResponseDataJSON contains the JSON metadata for the struct
// [V1InvoiceRegenerateResponseData]
type v1InvoiceRegenerateResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1InvoiceRegenerateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1InvoiceRegenerateResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1InvoiceVoidResponse struct {
	Data V1InvoiceVoidResponseData `json:"data"`
	JSON v1InvoiceVoidResponseJSON `json:"-"`
}

// v1InvoiceVoidResponseJSON contains the JSON metadata for the struct
// [V1InvoiceVoidResponse]
type v1InvoiceVoidResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1InvoiceVoidResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1InvoiceVoidResponseJSON) RawJSON() string {
	return r.raw
}

type V1InvoiceVoidResponseData struct {
	ID   string                        `json:"id,required" format:"uuid"`
	JSON v1InvoiceVoidResponseDataJSON `json:"-"`
}

// v1InvoiceVoidResponseDataJSON contains the JSON metadata for the struct
// [V1InvoiceVoidResponseData]
type v1InvoiceVoidResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1InvoiceVoidResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1InvoiceVoidResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1InvoiceRegenerateParams struct {
	// The invoice id to regenerate
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V1InvoiceRegenerateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1InvoiceVoidParams struct {
	// The invoice id to void
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V1InvoiceVoidParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
