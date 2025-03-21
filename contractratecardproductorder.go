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

// ContractRateCardProductOrderService contains methods and other services that
// help with interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContractRateCardProductOrderService] method instead.
type ContractRateCardProductOrderService struct {
	Options []option.RequestOption
}

// NewContractRateCardProductOrderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewContractRateCardProductOrderService(opts ...option.RequestOption) (r *ContractRateCardProductOrderService) {
	r = &ContractRateCardProductOrderService{}
	r.Options = opts
	return
}

// Updates ordering of specified products
func (r *ContractRateCardProductOrderService) Update(ctx context.Context, body ContractRateCardProductOrderUpdateParams, opts ...option.RequestOption) (res *ContractRateCardProductOrderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/rate-cards/moveRateCardProducts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Sets the ordering of products within a rate card
func (r *ContractRateCardProductOrderService) Set(ctx context.Context, body ContractRateCardProductOrderSetParams, opts ...option.RequestOption) (res *ContractRateCardProductOrderSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/rate-cards/setRateCardProductsOrder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractRateCardProductOrderUpdateResponse struct {
	Data shared.ID                                      `json:"data,required"`
	JSON contractRateCardProductOrderUpdateResponseJSON `json:"-"`
}

// contractRateCardProductOrderUpdateResponseJSON contains the JSON metadata for
// the struct [ContractRateCardProductOrderUpdateResponse]
type contractRateCardProductOrderUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractRateCardProductOrderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardProductOrderUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardProductOrderSetResponse struct {
	Data shared.ID                                   `json:"data,required"`
	JSON contractRateCardProductOrderSetResponseJSON `json:"-"`
}

// contractRateCardProductOrderSetResponseJSON contains the JSON metadata for the
// struct [ContractRateCardProductOrderSetResponse]
type contractRateCardProductOrderSetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractRateCardProductOrderSetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardProductOrderSetResponseJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardProductOrderUpdateParams struct {
	ProductMoves param.Field[[]ContractRateCardProductOrderUpdateParamsProductMove] `json:"product_moves,required"`
	// ID of the rate card to update
	RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
}

func (r ContractRateCardProductOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardProductOrderUpdateParamsProductMove struct {
	// 0-based index of the new position of the product
	Position param.Field[float64] `json:"position,required"`
	// ID of the product to move
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
}

func (r ContractRateCardProductOrderUpdateParamsProductMove) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardProductOrderSetParams struct {
	ProductOrder param.Field[[]string] `json:"product_order,required" format:"uuid"`
	// ID of the rate card to update
	RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
}

func (r ContractRateCardProductOrderSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
