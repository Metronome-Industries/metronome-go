// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
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
func NewV1InvoiceService(opts ...option.RequestOption) (r V1InvoiceService) {
	r = V1InvoiceService{}
	r.Options = opts
	return
}

// This endpoint regenerates a voided invoice and recalculates the invoice based on
// up-to-date rates, available balances, and other fees regardless of the billing
// period.
//
// ### Use this endpoint to:
//
// Recalculate an invoice with updated rate terms, available balance, and fees to
// correct billing disputes or discrepancies
//
// ### Key response fields:
//
// The regenerated invoice id, which is distinct from the previously voided
// invoice.
//
// ### Usage guidelines:
//
// If an invoice is attached to a contract with a billing provider on it, the
// regenerated invoice will be distributed based on the configuration.
func (r *V1InvoiceService) Regenerate(ctx context.Context, body V1InvoiceRegenerateParams, opts ...option.RequestOption) (res *V1InvoiceRegenerateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/invoices/regenerate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Permanently cancels an invoice by setting its status to voided, preventing
// collection and removing it from customer billing. Use this to correct billing
// errors, cancel incorrect charges, or handle disputed invoices that should not be
// collected. Returns the voided invoice ID with the status change applied
// immediately to stop any payment processing.
func (r *V1InvoiceService) Void(ctx context.Context, body V1InvoiceVoidParams, opts ...option.RequestOption) (res *V1InvoiceVoidResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/invoices/void"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1InvoiceRegenerateResponse struct {
	Data V1InvoiceRegenerateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1InvoiceRegenerateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InvoiceRegenerateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InvoiceRegenerateResponseData struct {
	// The new invoice id
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1InvoiceRegenerateResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1InvoiceRegenerateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InvoiceVoidResponse struct {
	Data V1InvoiceVoidResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1InvoiceVoidResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InvoiceVoidResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InvoiceVoidResponseData struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1InvoiceVoidResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1InvoiceVoidResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InvoiceRegenerateParams struct {
	// The invoice id to regenerate
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r V1InvoiceRegenerateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1InvoiceRegenerateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1InvoiceRegenerateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InvoiceVoidParams struct {
	// The invoice id to void
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r V1InvoiceVoidParams) MarshalJSON() (data []byte, err error) {
	type shadow V1InvoiceVoidParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1InvoiceVoidParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
