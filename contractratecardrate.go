// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/pagination"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// ContractRateCardRateService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContractRateCardRateService] method instead.
type ContractRateCardRateService struct {
	Options []option.RequestOption
}

// NewContractRateCardRateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContractRateCardRateService(opts ...option.RequestOption) (r *ContractRateCardRateService) {
	r = &ContractRateCardRateService{}
	r.Options = opts
	return
}

// Get rate card rates for a specific time.
func (r *ContractRateCardRateService) List(ctx context.Context, params ContractRateCardRateListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ContractRateCardRateListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "contract-pricing/rate-cards/getRates"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get rate card rates for a specific time.
func (r *ContractRateCardRateService) ListAutoPaging(ctx context.Context, params ContractRateCardRateListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ContractRateCardRateListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Add a new rate
func (r *ContractRateCardRateService) Add(ctx context.Context, body ContractRateCardRateAddParams, opts ...option.RequestOption) (res *ContractRateCardRateAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/rate-cards/addRate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Add new rates
func (r *ContractRateCardRateService) AddMany(ctx context.Context, body ContractRateCardRateAddManyParams, opts ...option.RequestOption) (res *ContractRateCardRateAddManyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/rate-cards/addRates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractRateCardRateListResponse struct {
	Entitled           bool                                 `json:"entitled,required"`
	ProductID          string                               `json:"product_id,required" format:"uuid"`
	ProductName        string                               `json:"product_name,required"`
	ProductTags        []string                             `json:"product_tags,required"`
	Rate               shared.Rate                          `json:"rate,required"`
	StartingAt         time.Time                            `json:"starting_at,required" format:"date-time"`
	EndingBefore       time.Time                            `json:"ending_before" format:"date-time"`
	PricingGroupValues map[string]string                    `json:"pricing_group_values"`
	JSON               contractRateCardRateListResponseJSON `json:"-"`
}

// contractRateCardRateListResponseJSON contains the JSON metadata for the struct
// [ContractRateCardRateListResponse]
type contractRateCardRateListResponseJSON struct {
	Entitled           apijson.Field
	ProductID          apijson.Field
	ProductName        apijson.Field
	ProductTags        apijson.Field
	Rate               apijson.Field
	StartingAt         apijson.Field
	EndingBefore       apijson.Field
	PricingGroupValues apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ContractRateCardRateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardRateListResponseJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardRateAddResponse struct {
	Data shared.Rate                         `json:"data,required"`
	JSON contractRateCardRateAddResponseJSON `json:"-"`
}

// contractRateCardRateAddResponseJSON contains the JSON metadata for the struct
// [ContractRateCardRateAddResponse]
type contractRateCardRateAddResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractRateCardRateAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardRateAddResponseJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardRateAddManyResponse struct {
	// The ID of the rate card to which the rates were added.
	Data shared.ID                               `json:"data,required"`
	JSON contractRateCardRateAddManyResponseJSON `json:"-"`
}

// contractRateCardRateAddManyResponseJSON contains the JSON metadata for the
// struct [ContractRateCardRateAddManyResponse]
type contractRateCardRateAddManyResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractRateCardRateAddManyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardRateAddManyResponseJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardRateListParams struct {
	// inclusive starting point for the rates schedule
	At param.Field[time.Time] `json:"at,required" format:"date-time"`
	// ID of the rate card to get the schedule for
	RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// List of rate selectors, rates matching ANY of the selector will be included in
	// the response Passing no selectors will result in all rates being returned.
	Selectors param.Field[[]ContractRateCardRateListParamsSelector] `json:"selectors"`
}

func (r ContractRateCardRateListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ContractRateCardRateListParams]'s query parameters as
// `url.Values`.
func (r ContractRateCardRateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContractRateCardRateListParamsSelector struct {
	// List of pricing group key value pairs, rates containing the matching key / value
	// pairs will be included in the response.
	PartialPricingGroupValues param.Field[map[string]string] `json:"partial_pricing_group_values"`
	// List of pricing group key value pairs, rates matching all of the key / value
	// pairs will be included in the response.
	PricingGroupValues param.Field[map[string]string] `json:"pricing_group_values"`
	// Rates matching the product id will be included in the response.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// List of product tags, rates matching any of the tags will be included in the
	// response.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r ContractRateCardRateListParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardRateAddParams struct {
	Entitled param.Field[bool] `json:"entitled,required"`
	// ID of the product to add a rate for
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// ID of the rate card to update
	RateCardID param.Field[string]                                `json:"rate_card_id,required" format:"uuid"`
	RateType   param.Field[ContractRateCardRateAddParamsRateType] `json:"rate_type,required"`
	// inclusive effective date
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// "The Metronome ID of the credit type to associate with price, defaults to USD
	// (cents) if not passed. Used by all rate_types except type PERCENTAGE. PERCENTAGE
	// rates use the credit type of associated rates."
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate param.Field[map[string]interface{}] `json:"custom_rate"`
	// exclusive end date
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type.
	IsProrated param.Field[bool] `json:"is_prorated"`
	// Default price. For FLAT and SUBSCRIPTION rate_type, this must be >=0. For
	// PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this
	// must be >=0 and <=1.
	Price param.Field[float64] `json:"price"`
	// Optional. List of pricing group key value pairs which will be used to calculate
	// the price.
	PricingGroupValues param.Field[map[string]string] `json:"pricing_group_values"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Field[float64] `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]ContractRateCardRateAddParamsTier] `json:"tiers"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices param.Field[bool] `json:"use_list_prices"`
}

func (r ContractRateCardRateAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardRateAddParamsRateType string

const (
	ContractRateCardRateAddParamsRateTypeFlat         ContractRateCardRateAddParamsRateType = "FLAT"
	ContractRateCardRateAddParamsRateTypePercentage   ContractRateCardRateAddParamsRateType = "PERCENTAGE"
	ContractRateCardRateAddParamsRateTypeSubscription ContractRateCardRateAddParamsRateType = "SUBSCRIPTION"
	ContractRateCardRateAddParamsRateTypeTiered       ContractRateCardRateAddParamsRateType = "TIERED"
	ContractRateCardRateAddParamsRateTypeCustom       ContractRateCardRateAddParamsRateType = "CUSTOM"
)

func (r ContractRateCardRateAddParamsRateType) IsKnown() bool {
	switch r {
	case ContractRateCardRateAddParamsRateTypeFlat, ContractRateCardRateAddParamsRateTypePercentage, ContractRateCardRateAddParamsRateTypeSubscription, ContractRateCardRateAddParamsRateTypeTiered, ContractRateCardRateAddParamsRateTypeCustom:
		return true
	}
	return false
}

type ContractRateCardRateAddParamsTier struct {
	Price param.Field[float64] `json:"price,required"`
	Size  param.Field[float64] `json:"size"`
}

func (r ContractRateCardRateAddParamsTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardRateAddManyParams struct {
	RateCardID param.Field[string]                                  `json:"rate_card_id" format:"uuid"`
	Rates      param.Field[[]ContractRateCardRateAddManyParamsRate] `json:"rates"`
}

func (r ContractRateCardRateAddManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardRateAddManyParamsRate struct {
	Entitled param.Field[bool] `json:"entitled,required"`
	// ID of the product to add a rate for
	ProductID param.Field[string]                                         `json:"product_id,required" format:"uuid"`
	RateType  param.Field[ContractRateCardRateAddManyParamsRatesRateType] `json:"rate_type,required"`
	// inclusive effective date
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// "The Metronome ID of the credit type to associate with price, defaults to USD
	// (cents) if not passed. Used by all rate_types except type PERCENTAGE. PERCENTAGE
	// rates use the credit type of associated rates."
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate param.Field[map[string]interface{}] `json:"custom_rate"`
	// exclusive end date
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type.
	IsProrated param.Field[bool] `json:"is_prorated"`
	// Default price. For FLAT and SUBSCRIPTION rate_type, this must be >=0. For
	// PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this
	// must be >=0 and <=1.
	Price param.Field[float64] `json:"price"`
	// Optional. List of pricing group key value pairs which will be used to calculate
	// the price.
	PricingGroupValues param.Field[map[string]string] `json:"pricing_group_values"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Field[float64] `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]ContractRateCardRateAddManyParamsRatesTier] `json:"tiers"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices param.Field[bool] `json:"use_list_prices"`
}

func (r ContractRateCardRateAddManyParamsRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardRateAddManyParamsRatesRateType string

const (
	ContractRateCardRateAddManyParamsRatesRateTypeFlat         ContractRateCardRateAddManyParamsRatesRateType = "FLAT"
	ContractRateCardRateAddManyParamsRatesRateTypePercentage   ContractRateCardRateAddManyParamsRatesRateType = "PERCENTAGE"
	ContractRateCardRateAddManyParamsRatesRateTypeSubscription ContractRateCardRateAddManyParamsRatesRateType = "SUBSCRIPTION"
	ContractRateCardRateAddManyParamsRatesRateTypeTiered       ContractRateCardRateAddManyParamsRatesRateType = "TIERED"
	ContractRateCardRateAddManyParamsRatesRateTypeCustom       ContractRateCardRateAddManyParamsRatesRateType = "CUSTOM"
)

func (r ContractRateCardRateAddManyParamsRatesRateType) IsKnown() bool {
	switch r {
	case ContractRateCardRateAddManyParamsRatesRateTypeFlat, ContractRateCardRateAddManyParamsRatesRateTypePercentage, ContractRateCardRateAddManyParamsRatesRateTypeSubscription, ContractRateCardRateAddManyParamsRatesRateTypeTiered, ContractRateCardRateAddManyParamsRatesRateTypeCustom:
		return true
	}
	return false
}

type ContractRateCardRateAddManyParamsRatesTier struct {
	Price param.Field[float64] `json:"price,required"`
	Size  param.Field[float64] `json:"size"`
}

func (r ContractRateCardRateAddManyParamsRatesTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
