// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// V1ContractRateCardRateService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContractRateCardRateService] method instead.
type V1ContractRateCardRateService struct {
	Options []option.RequestOption
}

// NewV1ContractRateCardRateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1ContractRateCardRateService(opts ...option.RequestOption) (r *V1ContractRateCardRateService) {
	r = &V1ContractRateCardRateService{}
	r.Options = opts
	return
}

// Get all rates for a rate card at a point in time
func (r *V1ContractRateCardRateService) List(ctx context.Context, params V1ContractRateCardRateListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1ContractRateCardRateListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/contract-pricing/rate-cards/getRates"
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

// Get all rates for a rate card at a point in time
func (r *V1ContractRateCardRateService) ListAutoPaging(ctx context.Context, params V1ContractRateCardRateListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1ContractRateCardRateListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Add a new rate
func (r *V1ContractRateCardRateService) Add(ctx context.Context, body V1ContractRateCardRateAddParams, opts ...option.RequestOption) (res *V1ContractRateCardRateAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/rate-cards/addRate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Add new rates
func (r *V1ContractRateCardRateService) AddMany(ctx context.Context, body V1ContractRateCardRateAddManyParams, opts ...option.RequestOption) (res *V1ContractRateCardRateAddManyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/rate-cards/addRates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1ContractRateCardRateListResponse struct {
	Entitled            bool              `json:"entitled,required"`
	ProductCustomFields map[string]string `json:"product_custom_fields,required"`
	ProductID           string            `json:"product_id,required" format:"uuid"`
	ProductName         string            `json:"product_name,required"`
	ProductTags         []string          `json:"product_tags,required"`
	Rate                shared.Rate       `json:"rate,required"`
	StartingAt          time.Time         `json:"starting_at,required" format:"date-time"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate         V1ContractRateCardRateListResponseCommitRate `json:"commit_rate"`
	EndingBefore       time.Time                                    `json:"ending_before" format:"date-time"`
	PricingGroupValues map[string]string                            `json:"pricing_group_values"`
	JSON               v1ContractRateCardRateListResponseJSON       `json:"-"`
}

// v1ContractRateCardRateListResponseJSON contains the JSON metadata for the struct
// [V1ContractRateCardRateListResponse]
type v1ContractRateCardRateListResponseJSON struct {
	Entitled            apijson.Field
	ProductCustomFields apijson.Field
	ProductID           apijson.Field
	ProductName         apijson.Field
	ProductTags         apijson.Field
	Rate                apijson.Field
	StartingAt          apijson.Field
	CommitRate          apijson.Field
	EndingBefore        apijson.Field
	PricingGroupValues  apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1ContractRateCardRateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardRateListResponseJSON) RawJSON() string {
	return r.raw
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type V1ContractRateCardRateListResponseCommitRate struct {
	RateType V1ContractRateCardRateListResponseCommitRateRateType `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price float64 `json:"price"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier                                    `json:"tiers"`
	JSON  v1ContractRateCardRateListResponseCommitRateJSON `json:"-"`
}

// v1ContractRateCardRateListResponseCommitRateJSON contains the JSON metadata for
// the struct [V1ContractRateCardRateListResponseCommitRate]
type v1ContractRateCardRateListResponseCommitRateJSON struct {
	RateType    apijson.Field
	Price       apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractRateCardRateListResponseCommitRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardRateListResponseCommitRateJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardRateListResponseCommitRateRateType string

const (
	V1ContractRateCardRateListResponseCommitRateRateTypeFlat         V1ContractRateCardRateListResponseCommitRateRateType = "FLAT"
	V1ContractRateCardRateListResponseCommitRateRateTypePercentage   V1ContractRateCardRateListResponseCommitRateRateType = "PERCENTAGE"
	V1ContractRateCardRateListResponseCommitRateRateTypeSubscription V1ContractRateCardRateListResponseCommitRateRateType = "SUBSCRIPTION"
	V1ContractRateCardRateListResponseCommitRateRateTypeTiered       V1ContractRateCardRateListResponseCommitRateRateType = "TIERED"
	V1ContractRateCardRateListResponseCommitRateRateTypeCustom       V1ContractRateCardRateListResponseCommitRateRateType = "CUSTOM"
)

func (r V1ContractRateCardRateListResponseCommitRateRateType) IsKnown() bool {
	switch r {
	case V1ContractRateCardRateListResponseCommitRateRateTypeFlat, V1ContractRateCardRateListResponseCommitRateRateTypePercentage, V1ContractRateCardRateListResponseCommitRateRateTypeSubscription, V1ContractRateCardRateListResponseCommitRateRateTypeTiered, V1ContractRateCardRateListResponseCommitRateRateTypeCustom:
		return true
	}
	return false
}

type V1ContractRateCardRateAddResponse struct {
	Data V1ContractRateCardRateAddResponseData `json:"data,required"`
	JSON v1ContractRateCardRateAddResponseJSON `json:"-"`
}

// v1ContractRateCardRateAddResponseJSON contains the JSON metadata for the struct
// [V1ContractRateCardRateAddResponse]
type v1ContractRateCardRateAddResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractRateCardRateAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardRateAddResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardRateAddResponseData struct {
	RateType V1ContractRateCardRateAddResponseDataRateType `json:"rate_type,required"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate V1ContractRateCardRateAddResponseDataCommitRate `json:"commit_rate"`
	CreditType shared.CreditTypeData                           `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]interface{} `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// if pricing groups are used, this will contain the values used to calculate the
	// price
	PricingGroupValues map[string]string `json:"pricing_group_values"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier `json:"tiers"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool                                      `json:"use_list_prices"`
	JSON          v1ContractRateCardRateAddResponseDataJSON `json:"-"`
}

// v1ContractRateCardRateAddResponseDataJSON contains the JSON metadata for the
// struct [V1ContractRateCardRateAddResponseData]
type v1ContractRateCardRateAddResponseDataJSON struct {
	RateType           apijson.Field
	CommitRate         apijson.Field
	CreditType         apijson.Field
	CustomRate         apijson.Field
	IsProrated         apijson.Field
	Price              apijson.Field
	PricingGroupValues apijson.Field
	Quantity           apijson.Field
	Tiers              apijson.Field
	UseListPrices      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V1ContractRateCardRateAddResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardRateAddResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardRateAddResponseDataRateType string

const (
	V1ContractRateCardRateAddResponseDataRateTypeFlat         V1ContractRateCardRateAddResponseDataRateType = "FLAT"
	V1ContractRateCardRateAddResponseDataRateTypePercentage   V1ContractRateCardRateAddResponseDataRateType = "PERCENTAGE"
	V1ContractRateCardRateAddResponseDataRateTypeSubscription V1ContractRateCardRateAddResponseDataRateType = "SUBSCRIPTION"
	V1ContractRateCardRateAddResponseDataRateTypeCustom       V1ContractRateCardRateAddResponseDataRateType = "CUSTOM"
	V1ContractRateCardRateAddResponseDataRateTypeTiered       V1ContractRateCardRateAddResponseDataRateType = "TIERED"
)

func (r V1ContractRateCardRateAddResponseDataRateType) IsKnown() bool {
	switch r {
	case V1ContractRateCardRateAddResponseDataRateTypeFlat, V1ContractRateCardRateAddResponseDataRateTypePercentage, V1ContractRateCardRateAddResponseDataRateTypeSubscription, V1ContractRateCardRateAddResponseDataRateTypeCustom, V1ContractRateCardRateAddResponseDataRateTypeTiered:
		return true
	}
	return false
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type V1ContractRateCardRateAddResponseDataCommitRate struct {
	RateType V1ContractRateCardRateAddResponseDataCommitRateRateType `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price float64 `json:"price"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier                                       `json:"tiers"`
	JSON  v1ContractRateCardRateAddResponseDataCommitRateJSON `json:"-"`
}

// v1ContractRateCardRateAddResponseDataCommitRateJSON contains the JSON metadata
// for the struct [V1ContractRateCardRateAddResponseDataCommitRate]
type v1ContractRateCardRateAddResponseDataCommitRateJSON struct {
	RateType    apijson.Field
	Price       apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractRateCardRateAddResponseDataCommitRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardRateAddResponseDataCommitRateJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardRateAddResponseDataCommitRateRateType string

const (
	V1ContractRateCardRateAddResponseDataCommitRateRateTypeFlat         V1ContractRateCardRateAddResponseDataCommitRateRateType = "FLAT"
	V1ContractRateCardRateAddResponseDataCommitRateRateTypePercentage   V1ContractRateCardRateAddResponseDataCommitRateRateType = "PERCENTAGE"
	V1ContractRateCardRateAddResponseDataCommitRateRateTypeSubscription V1ContractRateCardRateAddResponseDataCommitRateRateType = "SUBSCRIPTION"
	V1ContractRateCardRateAddResponseDataCommitRateRateTypeTiered       V1ContractRateCardRateAddResponseDataCommitRateRateType = "TIERED"
	V1ContractRateCardRateAddResponseDataCommitRateRateTypeCustom       V1ContractRateCardRateAddResponseDataCommitRateRateType = "CUSTOM"
)

func (r V1ContractRateCardRateAddResponseDataCommitRateRateType) IsKnown() bool {
	switch r {
	case V1ContractRateCardRateAddResponseDataCommitRateRateTypeFlat, V1ContractRateCardRateAddResponseDataCommitRateRateTypePercentage, V1ContractRateCardRateAddResponseDataCommitRateRateTypeSubscription, V1ContractRateCardRateAddResponseDataCommitRateRateTypeTiered, V1ContractRateCardRateAddResponseDataCommitRateRateTypeCustom:
		return true
	}
	return false
}

type V1ContractRateCardRateAddManyResponse struct {
	// The ID of the rate card to which the rates were added.
	Data shared.ID                                 `json:"data,required"`
	JSON v1ContractRateCardRateAddManyResponseJSON `json:"-"`
}

// v1ContractRateCardRateAddManyResponseJSON contains the JSON metadata for the
// struct [V1ContractRateCardRateAddManyResponse]
type v1ContractRateCardRateAddManyResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractRateCardRateAddManyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardRateAddManyResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardRateListParams struct {
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
	Selectors param.Field[[]V1ContractRateCardRateListParamsSelector] `json:"selectors"`
}

func (r V1ContractRateCardRateListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V1ContractRateCardRateListParams]'s query parameters as
// `url.Values`.
func (r V1ContractRateCardRateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ContractRateCardRateListParamsSelector struct {
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

func (r V1ContractRateCardRateListParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardRateAddParams struct {
	Entitled param.Field[bool] `json:"entitled,required"`
	// ID of the product to add a rate for
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// ID of the rate card to update
	RateCardID param.Field[string]                                  `json:"rate_card_id,required" format:"uuid"`
	RateType   param.Field[V1ContractRateCardRateAddParamsRateType] `json:"rate_type,required"`
	// inclusive effective date
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate param.Field[V1ContractRateCardRateAddParamsCommitRate] `json:"commit_rate"`
	// The Metronome ID of the credit type to associate with price, defaults to USD
	// (cents) if not passed. Used by all rate_types except type PERCENTAGE. PERCENTAGE
	// rates use the credit type of associated rates.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate param.Field[map[string]interface{}] `json:"custom_rate"`
	// exclusive end date
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
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
	Tiers param.Field[[]shared.TierParam] `json:"tiers"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices param.Field[bool] `json:"use_list_prices"`
}

func (r V1ContractRateCardRateAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardRateAddParamsRateType string

const (
	V1ContractRateCardRateAddParamsRateTypeFlat         V1ContractRateCardRateAddParamsRateType = "FLAT"
	V1ContractRateCardRateAddParamsRateTypePercentage   V1ContractRateCardRateAddParamsRateType = "PERCENTAGE"
	V1ContractRateCardRateAddParamsRateTypeSubscription V1ContractRateCardRateAddParamsRateType = "SUBSCRIPTION"
	V1ContractRateCardRateAddParamsRateTypeTiered       V1ContractRateCardRateAddParamsRateType = "TIERED"
	V1ContractRateCardRateAddParamsRateTypeCustom       V1ContractRateCardRateAddParamsRateType = "CUSTOM"
)

func (r V1ContractRateCardRateAddParamsRateType) IsKnown() bool {
	switch r {
	case V1ContractRateCardRateAddParamsRateTypeFlat, V1ContractRateCardRateAddParamsRateTypePercentage, V1ContractRateCardRateAddParamsRateTypeSubscription, V1ContractRateCardRateAddParamsRateTypeTiered, V1ContractRateCardRateAddParamsRateTypeCustom:
		return true
	}
	return false
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type V1ContractRateCardRateAddParamsCommitRate struct {
	RateType param.Field[V1ContractRateCardRateAddParamsCommitRateRateType] `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price param.Field[float64] `json:"price"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]shared.TierParam] `json:"tiers"`
}

func (r V1ContractRateCardRateAddParamsCommitRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardRateAddParamsCommitRateRateType string

const (
	V1ContractRateCardRateAddParamsCommitRateRateTypeFlat         V1ContractRateCardRateAddParamsCommitRateRateType = "FLAT"
	V1ContractRateCardRateAddParamsCommitRateRateTypePercentage   V1ContractRateCardRateAddParamsCommitRateRateType = "PERCENTAGE"
	V1ContractRateCardRateAddParamsCommitRateRateTypeSubscription V1ContractRateCardRateAddParamsCommitRateRateType = "SUBSCRIPTION"
	V1ContractRateCardRateAddParamsCommitRateRateTypeTiered       V1ContractRateCardRateAddParamsCommitRateRateType = "TIERED"
	V1ContractRateCardRateAddParamsCommitRateRateTypeCustom       V1ContractRateCardRateAddParamsCommitRateRateType = "CUSTOM"
)

func (r V1ContractRateCardRateAddParamsCommitRateRateType) IsKnown() bool {
	switch r {
	case V1ContractRateCardRateAddParamsCommitRateRateTypeFlat, V1ContractRateCardRateAddParamsCommitRateRateTypePercentage, V1ContractRateCardRateAddParamsCommitRateRateTypeSubscription, V1ContractRateCardRateAddParamsCommitRateRateTypeTiered, V1ContractRateCardRateAddParamsCommitRateRateTypeCustom:
		return true
	}
	return false
}

type V1ContractRateCardRateAddManyParams struct {
	RateCardID param.Field[string]                                    `json:"rate_card_id,required" format:"uuid"`
	Rates      param.Field[[]V1ContractRateCardRateAddManyParamsRate] `json:"rates,required"`
}

func (r V1ContractRateCardRateAddManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardRateAddManyParamsRate struct {
	Entitled param.Field[bool] `json:"entitled,required"`
	// ID of the product to add a rate for
	ProductID param.Field[string]                                           `json:"product_id,required" format:"uuid"`
	RateType  param.Field[V1ContractRateCardRateAddManyParamsRatesRateType] `json:"rate_type,required"`
	// inclusive effective date
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate param.Field[V1ContractRateCardRateAddManyParamsRatesCommitRate] `json:"commit_rate"`
	// "The Metronome ID of the credit type to associate with price, defaults to USD
	// (cents) if not passed. Used by all rate_types except type PERCENTAGE. PERCENTAGE
	// rates use the credit type of associated rates."
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate param.Field[map[string]interface{}] `json:"custom_rate"`
	// exclusive end date
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
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
	Tiers param.Field[[]shared.TierParam] `json:"tiers"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices param.Field[bool] `json:"use_list_prices"`
}

func (r V1ContractRateCardRateAddManyParamsRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardRateAddManyParamsRatesRateType string

const (
	V1ContractRateCardRateAddManyParamsRatesRateTypeFlat         V1ContractRateCardRateAddManyParamsRatesRateType = "FLAT"
	V1ContractRateCardRateAddManyParamsRatesRateTypePercentage   V1ContractRateCardRateAddManyParamsRatesRateType = "PERCENTAGE"
	V1ContractRateCardRateAddManyParamsRatesRateTypeSubscription V1ContractRateCardRateAddManyParamsRatesRateType = "SUBSCRIPTION"
	V1ContractRateCardRateAddManyParamsRatesRateTypeTiered       V1ContractRateCardRateAddManyParamsRatesRateType = "TIERED"
	V1ContractRateCardRateAddManyParamsRatesRateTypeCustom       V1ContractRateCardRateAddManyParamsRatesRateType = "CUSTOM"
)

func (r V1ContractRateCardRateAddManyParamsRatesRateType) IsKnown() bool {
	switch r {
	case V1ContractRateCardRateAddManyParamsRatesRateTypeFlat, V1ContractRateCardRateAddManyParamsRatesRateTypePercentage, V1ContractRateCardRateAddManyParamsRatesRateTypeSubscription, V1ContractRateCardRateAddManyParamsRatesRateTypeTiered, V1ContractRateCardRateAddManyParamsRatesRateTypeCustom:
		return true
	}
	return false
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type V1ContractRateCardRateAddManyParamsRatesCommitRate struct {
	RateType param.Field[V1ContractRateCardRateAddManyParamsRatesCommitRateRateType] `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price param.Field[float64] `json:"price"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]shared.TierParam] `json:"tiers"`
}

func (r V1ContractRateCardRateAddManyParamsRatesCommitRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardRateAddManyParamsRatesCommitRateRateType string

const (
	V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypeFlat         V1ContractRateCardRateAddManyParamsRatesCommitRateRateType = "FLAT"
	V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypePercentage   V1ContractRateCardRateAddManyParamsRatesCommitRateRateType = "PERCENTAGE"
	V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypeSubscription V1ContractRateCardRateAddManyParamsRatesCommitRateRateType = "SUBSCRIPTION"
	V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypeTiered       V1ContractRateCardRateAddManyParamsRatesCommitRateRateType = "TIERED"
	V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypeCustom       V1ContractRateCardRateAddManyParamsRatesCommitRateRateType = "CUSTOM"
)

func (r V1ContractRateCardRateAddManyParamsRatesCommitRateRateType) IsKnown() bool {
	switch r {
	case V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypeFlat, V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypePercentage, V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypeSubscription, V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypeTiered, V1ContractRateCardRateAddManyParamsRatesCommitRateRateTypeCustom:
		return true
	}
	return false
}
