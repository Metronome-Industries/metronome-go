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
	"github.com/Metronome-Industries/metronome-go/shared"
)

// V1ContractRateCardProductOrderService contains methods and other services that
// help with interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContractRateCardProductOrderService] method instead.
type V1ContractRateCardProductOrderService struct {
	Options []option.RequestOption
}

// NewV1ContractRateCardProductOrderService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1ContractRateCardProductOrderService(opts ...option.RequestOption) (r V1ContractRateCardProductOrderService) {
	r = V1ContractRateCardProductOrderService{}
	r.Options = opts
	return
}

// The ordering of products on a rate card determines the order in which the
// products will appear on customers' invoices. Use this endpoint to set the order
// of specific products on the rate card by moving them relative to their current
// location.
func (r *V1ContractRateCardProductOrderService) Update(ctx context.Context, body V1ContractRateCardProductOrderUpdateParams, opts ...option.RequestOption) (res *V1ContractRateCardProductOrderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/rate-cards/moveRateCardProducts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The ordering of products on a rate card determines the order in which the
// products will appear on customers' invoices. Use this endpoint to set the order
// of products on the rate card.
func (r *V1ContractRateCardProductOrderService) Set(ctx context.Context, body V1ContractRateCardProductOrderSetParams, opts ...option.RequestOption) (res *V1ContractRateCardProductOrderSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/rate-cards/setRateCardProductsOrder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1ContractRateCardProductOrderUpdateResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardProductOrderUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardProductOrderUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardProductOrderSetResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardProductOrderSetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardProductOrderSetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardProductOrderUpdateParams struct {
	ProductMoves []V1ContractRateCardProductOrderUpdateParamsProductMove `json:"product_moves,omitzero,required"`
	// ID of the rate card to update
	RateCardID string `json:"rate_card_id,required" format:"uuid"`
	paramObj
}

func (r V1ContractRateCardProductOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardProductOrderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardProductOrderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Position, ProductID are required.
type V1ContractRateCardProductOrderUpdateParamsProductMove struct {
	// 0-based index of the new position of the product
	Position float64 `json:"position,required"`
	// ID of the product to move
	ProductID string `json:"product_id,required" format:"uuid"`
	paramObj
}

func (r V1ContractRateCardProductOrderUpdateParamsProductMove) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardProductOrderUpdateParamsProductMove
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardProductOrderUpdateParamsProductMove) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardProductOrderSetParams struct {
	ProductOrder []string `json:"product_order,omitzero,required" format:"uuid"`
	// ID of the rate card to update
	RateCardID string `json:"rate_card_id,required" format:"uuid"`
	paramObj
}

func (r V1ContractRateCardProductOrderSetParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardProductOrderSetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardProductOrderSetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
