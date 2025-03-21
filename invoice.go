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

// InvoiceService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceService] method instead.
type InvoiceService struct {
	Options []option.RequestOption
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r *InvoiceService) {
	r = &InvoiceService{}
	r.Options = opts
	return
}

// Regenerate a voided contract invoice
func (r *InvoiceService) Regenerate(ctx context.Context, body InvoiceRegenerateParams, opts ...option.RequestOption) (res *InvoiceRegenerateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/invoices/regenerate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Void an invoice
func (r *InvoiceService) Void(ctx context.Context, body InvoiceVoidParams, opts ...option.RequestOption) (res *InvoiceVoidResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/invoices/void"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type InvoiceRegenerateResponse struct {
	Data InvoiceRegenerateResponseData `json:"data"`
	JSON invoiceRegenerateResponseJSON `json:"-"`
}

// invoiceRegenerateResponseJSON contains the JSON metadata for the struct
// [InvoiceRegenerateResponse]
type invoiceRegenerateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceRegenerateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceRegenerateResponseJSON) RawJSON() string {
	return r.raw
}

type InvoiceRegenerateResponseData struct {
	// The new invoice id
	ID   string                            `json:"id,required" format:"uuid"`
	JSON invoiceRegenerateResponseDataJSON `json:"-"`
}

// invoiceRegenerateResponseDataJSON contains the JSON metadata for the struct
// [InvoiceRegenerateResponseData]
type invoiceRegenerateResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceRegenerateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceRegenerateResponseDataJSON) RawJSON() string {
	return r.raw
}

type InvoiceVoidResponse struct {
	Data InvoiceVoidResponseData `json:"data"`
	JSON invoiceVoidResponseJSON `json:"-"`
}

// invoiceVoidResponseJSON contains the JSON metadata for the struct
// [InvoiceVoidResponse]
type invoiceVoidResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceVoidResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceVoidResponseJSON) RawJSON() string {
	return r.raw
}

type InvoiceVoidResponseData struct {
	ID   string                      `json:"id,required" format:"uuid"`
	JSON invoiceVoidResponseDataJSON `json:"-"`
}

// invoiceVoidResponseDataJSON contains the JSON metadata for the struct
// [InvoiceVoidResponseData]
type invoiceVoidResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceVoidResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceVoidResponseDataJSON) RawJSON() string {
	return r.raw
}

type InvoiceRegenerateParams struct {
	// The invoice id to regenerate
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r InvoiceRegenerateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceVoidParams struct {
	// The invoice id to void
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r InvoiceVoidParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
