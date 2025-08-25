// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
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
func NewV1ContractRateCardProductOrderService(opts ...option.RequestOption) (r *V1ContractRateCardProductOrderService) {
	r = &V1ContractRateCardProductOrderService{}
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
	Data shared.ID                                        `json:"data,required"`
	JSON v1ContractRateCardProductOrderUpdateResponseJSON `json:"-"`
}

// v1ContractRateCardProductOrderUpdateResponseJSON contains the JSON metadata for
// the struct [V1ContractRateCardProductOrderUpdateResponse]
type v1ContractRateCardProductOrderUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractRateCardProductOrderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardProductOrderUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardProductOrderSetResponse struct {
	Data shared.ID                                     `json:"data,required"`
	JSON v1ContractRateCardProductOrderSetResponseJSON `json:"-"`
}

// v1ContractRateCardProductOrderSetResponseJSON contains the JSON metadata for the
// struct [V1ContractRateCardProductOrderSetResponse]
type v1ContractRateCardProductOrderSetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractRateCardProductOrderSetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardProductOrderSetResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardProductOrderUpdateParams struct {
	ProductMoves param.Field[[]V1ContractRateCardProductOrderUpdateParamsProductMove] `json:"product_moves,required"`
	// ID of the rate card to update
	RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
}

func (r V1ContractRateCardProductOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardProductOrderUpdateParamsProductMove struct {
	// 0-based index of the new position of the product
	Position param.Field[float64] `json:"position,required"`
	// ID of the product to move
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
}

func (r V1ContractRateCardProductOrderUpdateParamsProductMove) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardProductOrderSetParams struct {
	ProductOrder param.Field[[]string] `json:"product_order,required" format:"uuid"`
	// ID of the rate card to update
	RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
}

func (r V1ContractRateCardProductOrderSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
