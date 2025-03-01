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

// Get all rates for a rate card at a point in time
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

// Get all rates for a rate card at a point in time
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
	Entitled            bool              `json:"entitled,required"`
	ProductCustomFields map[string]string `json:"product_custom_fields,required"`
	ProductID           string            `json:"product_id,required" format:"uuid"`
	ProductName         string            `json:"product_name,required"`
	ProductTags         []string          `json:"product_tags,required"`
	Rate                shared.Rate       `json:"rate,required"`
	StartingAt          time.Time         `json:"starting_at,required" format:"date-time"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate         ContractRateCardRateListResponseCommitRate `json:"commit_rate"`
	EndingBefore       time.Time                                  `json:"ending_before" format:"date-time"`
	PricingGroupValues map[string]string                          `json:"pricing_group_values"`
	JSON               contractRateCardRateListResponseJSON       `json:"-"`
}

// contractRateCardRateListResponseJSON contains the JSON metadata for the struct
// [ContractRateCardRateListResponse]
type contractRateCardRateListResponseJSON struct {
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

func (r *ContractRateCardRateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardRateListResponseJSON) RawJSON() string {
	return r.raw
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type ContractRateCardRateListResponseCommitRate struct {
	RateType ContractRateCardRateListResponseCommitRateRateType `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price float64 `json:"price"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier                                  `json:"tiers"`
	JSON  contractRateCardRateListResponseCommitRateJSON `json:"-"`
}

// contractRateCardRateListResponseCommitRateJSON contains the JSON metadata for
// the struct [ContractRateCardRateListResponseCommitRate]
type contractRateCardRateListResponseCommitRateJSON struct {
	RateType    apijson.Field
	Price       apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractRateCardRateListResponseCommitRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardRateListResponseCommitRateJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardRateListResponseCommitRateRateType string

const (
	ContractRateCardRateListResponseCommitRateRateTypeFlat         ContractRateCardRateListResponseCommitRateRateType = "FLAT"
	ContractRateCardRateListResponseCommitRateRateTypePercentage   ContractRateCardRateListResponseCommitRateRateType = "PERCENTAGE"
	ContractRateCardRateListResponseCommitRateRateTypeSubscription ContractRateCardRateListResponseCommitRateRateType = "SUBSCRIPTION"
	ContractRateCardRateListResponseCommitRateRateTypeTiered       ContractRateCardRateListResponseCommitRateRateType = "TIERED"
	ContractRateCardRateListResponseCommitRateRateTypeCustom       ContractRateCardRateListResponseCommitRateRateType = "CUSTOM"
)

func (r ContractRateCardRateListResponseCommitRateRateType) IsKnown() bool {
	switch r {
	case ContractRateCardRateListResponseCommitRateRateTypeFlat, ContractRateCardRateListResponseCommitRateRateTypePercentage, ContractRateCardRateListResponseCommitRateRateTypeSubscription, ContractRateCardRateListResponseCommitRateRateTypeTiered, ContractRateCardRateListResponseCommitRateRateTypeCustom:
		return true
	}
	return false
}

type ContractRateCardRateAddResponse struct {
	Data ContractRateCardRateAddResponseData `json:"data,required"`
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

type ContractRateCardRateAddResponseData struct {
	RateType ContractRateCardRateAddResponseDataRateType `json:"rate_type,required"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate ContractRateCardRateAddResponseDataCommitRate `json:"commit_rate"`
	CreditType shared.CreditTypeData                         `json:"credit_type"`
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
	UseListPrices bool                                    `json:"use_list_prices"`
	JSON          contractRateCardRateAddResponseDataJSON `json:"-"`
}

// contractRateCardRateAddResponseDataJSON contains the JSON metadata for the
// struct [ContractRateCardRateAddResponseData]
type contractRateCardRateAddResponseDataJSON struct {
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

func (r *ContractRateCardRateAddResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardRateAddResponseDataJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardRateAddResponseDataRateType string

const (
	ContractRateCardRateAddResponseDataRateTypeFlat         ContractRateCardRateAddResponseDataRateType = "FLAT"
	ContractRateCardRateAddResponseDataRateTypePercentage   ContractRateCardRateAddResponseDataRateType = "PERCENTAGE"
	ContractRateCardRateAddResponseDataRateTypeSubscription ContractRateCardRateAddResponseDataRateType = "SUBSCRIPTION"
	ContractRateCardRateAddResponseDataRateTypeCustom       ContractRateCardRateAddResponseDataRateType = "CUSTOM"
	ContractRateCardRateAddResponseDataRateTypeTiered       ContractRateCardRateAddResponseDataRateType = "TIERED"
)

func (r ContractRateCardRateAddResponseDataRateType) IsKnown() bool {
	switch r {
	case ContractRateCardRateAddResponseDataRateTypeFlat, ContractRateCardRateAddResponseDataRateTypePercentage, ContractRateCardRateAddResponseDataRateTypeSubscription, ContractRateCardRateAddResponseDataRateTypeCustom, ContractRateCardRateAddResponseDataRateTypeTiered:
		return true
	}
	return false
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type ContractRateCardRateAddResponseDataCommitRate struct {
	RateType ContractRateCardRateAddResponseDataCommitRateRateType `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price float64 `json:"price"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier                                     `json:"tiers"`
	JSON  contractRateCardRateAddResponseDataCommitRateJSON `json:"-"`
}

// contractRateCardRateAddResponseDataCommitRateJSON contains the JSON metadata for
// the struct [ContractRateCardRateAddResponseDataCommitRate]
type contractRateCardRateAddResponseDataCommitRateJSON struct {
	RateType    apijson.Field
	Price       apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractRateCardRateAddResponseDataCommitRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardRateAddResponseDataCommitRateJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardRateAddResponseDataCommitRateRateType string

const (
	ContractRateCardRateAddResponseDataCommitRateRateTypeFlat         ContractRateCardRateAddResponseDataCommitRateRateType = "FLAT"
	ContractRateCardRateAddResponseDataCommitRateRateTypePercentage   ContractRateCardRateAddResponseDataCommitRateRateType = "PERCENTAGE"
	ContractRateCardRateAddResponseDataCommitRateRateTypeSubscription ContractRateCardRateAddResponseDataCommitRateRateType = "SUBSCRIPTION"
	ContractRateCardRateAddResponseDataCommitRateRateTypeTiered       ContractRateCardRateAddResponseDataCommitRateRateType = "TIERED"
	ContractRateCardRateAddResponseDataCommitRateRateTypeCustom       ContractRateCardRateAddResponseDataCommitRateRateType = "CUSTOM"
)

func (r ContractRateCardRateAddResponseDataCommitRateRateType) IsKnown() bool {
	switch r {
	case ContractRateCardRateAddResponseDataCommitRateRateTypeFlat, ContractRateCardRateAddResponseDataCommitRateRateTypePercentage, ContractRateCardRateAddResponseDataCommitRateRateTypeSubscription, ContractRateCardRateAddResponseDataCommitRateRateTypeTiered, ContractRateCardRateAddResponseDataCommitRateRateTypeCustom:
		return true
	}
	return false
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
	// Optional. Frequency to bill subscriptions with. Required for subscription type
	// products with Flat rate.
	BillingFrequency param.Field[ContractRateCardRateAddParamsBillingFrequency] `json:"billing_frequency"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate param.Field[ContractRateCardRateAddParamsCommitRate] `json:"commit_rate"`
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

// Optional. Frequency to bill subscriptions with. Required for subscription type
// products with Flat rate.
type ContractRateCardRateAddParamsBillingFrequency string

const (
	ContractRateCardRateAddParamsBillingFrequencyMonthly   ContractRateCardRateAddParamsBillingFrequency = "MONTHLY"
	ContractRateCardRateAddParamsBillingFrequencyQuarterly ContractRateCardRateAddParamsBillingFrequency = "QUARTERLY"
	ContractRateCardRateAddParamsBillingFrequencyAnnual    ContractRateCardRateAddParamsBillingFrequency = "ANNUAL"
)

func (r ContractRateCardRateAddParamsBillingFrequency) IsKnown() bool {
	switch r {
	case ContractRateCardRateAddParamsBillingFrequencyMonthly, ContractRateCardRateAddParamsBillingFrequencyQuarterly, ContractRateCardRateAddParamsBillingFrequencyAnnual:
		return true
	}
	return false
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type ContractRateCardRateAddParamsCommitRate struct {
	RateType param.Field[ContractRateCardRateAddParamsCommitRateRateType] `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price param.Field[float64] `json:"price"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]shared.TierParam] `json:"tiers"`
}

func (r ContractRateCardRateAddParamsCommitRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardRateAddParamsCommitRateRateType string

const (
	ContractRateCardRateAddParamsCommitRateRateTypeFlat         ContractRateCardRateAddParamsCommitRateRateType = "FLAT"
	ContractRateCardRateAddParamsCommitRateRateTypePercentage   ContractRateCardRateAddParamsCommitRateRateType = "PERCENTAGE"
	ContractRateCardRateAddParamsCommitRateRateTypeSubscription ContractRateCardRateAddParamsCommitRateRateType = "SUBSCRIPTION"
	ContractRateCardRateAddParamsCommitRateRateTypeTiered       ContractRateCardRateAddParamsCommitRateRateType = "TIERED"
	ContractRateCardRateAddParamsCommitRateRateTypeCustom       ContractRateCardRateAddParamsCommitRateRateType = "CUSTOM"
)

func (r ContractRateCardRateAddParamsCommitRateRateType) IsKnown() bool {
	switch r {
	case ContractRateCardRateAddParamsCommitRateRateTypeFlat, ContractRateCardRateAddParamsCommitRateRateTypePercentage, ContractRateCardRateAddParamsCommitRateRateTypeSubscription, ContractRateCardRateAddParamsCommitRateRateTypeTiered, ContractRateCardRateAddParamsCommitRateRateTypeCustom:
		return true
	}
	return false
}

type ContractRateCardRateAddManyParams struct {
	RateCardID param.Field[string]                                  `json:"rate_card_id,required" format:"uuid"`
	Rates      param.Field[[]ContractRateCardRateAddManyParamsRate] `json:"rates,required"`
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
	// Optional. Frequency to bill subscriptions with. Required for subscription type
	// products with Flat rate.
	BillingFrequency param.Field[ContractRateCardRateAddManyParamsRatesBillingFrequency] `json:"billing_frequency"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate param.Field[ContractRateCardRateAddManyParamsRatesCommitRate] `json:"commit_rate"`
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

// Optional. Frequency to bill subscriptions with. Required for subscription type
// products with Flat rate.
type ContractRateCardRateAddManyParamsRatesBillingFrequency string

const (
	ContractRateCardRateAddManyParamsRatesBillingFrequencyMonthly   ContractRateCardRateAddManyParamsRatesBillingFrequency = "MONTHLY"
	ContractRateCardRateAddManyParamsRatesBillingFrequencyQuarterly ContractRateCardRateAddManyParamsRatesBillingFrequency = "QUARTERLY"
	ContractRateCardRateAddManyParamsRatesBillingFrequencyAnnual    ContractRateCardRateAddManyParamsRatesBillingFrequency = "ANNUAL"
)

func (r ContractRateCardRateAddManyParamsRatesBillingFrequency) IsKnown() bool {
	switch r {
	case ContractRateCardRateAddManyParamsRatesBillingFrequencyMonthly, ContractRateCardRateAddManyParamsRatesBillingFrequencyQuarterly, ContractRateCardRateAddManyParamsRatesBillingFrequencyAnnual:
		return true
	}
	return false
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type ContractRateCardRateAddManyParamsRatesCommitRate struct {
	RateType param.Field[ContractRateCardRateAddManyParamsRatesCommitRateRateType] `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price param.Field[float64] `json:"price"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]shared.TierParam] `json:"tiers"`
}

func (r ContractRateCardRateAddManyParamsRatesCommitRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardRateAddManyParamsRatesCommitRateRateType string

const (
	ContractRateCardRateAddManyParamsRatesCommitRateRateTypeFlat         ContractRateCardRateAddManyParamsRatesCommitRateRateType = "FLAT"
	ContractRateCardRateAddManyParamsRatesCommitRateRateTypePercentage   ContractRateCardRateAddManyParamsRatesCommitRateRateType = "PERCENTAGE"
	ContractRateCardRateAddManyParamsRatesCommitRateRateTypeSubscription ContractRateCardRateAddManyParamsRatesCommitRateRateType = "SUBSCRIPTION"
	ContractRateCardRateAddManyParamsRatesCommitRateRateTypeTiered       ContractRateCardRateAddManyParamsRatesCommitRateRateType = "TIERED"
	ContractRateCardRateAddManyParamsRatesCommitRateRateTypeCustom       ContractRateCardRateAddManyParamsRatesCommitRateRateType = "CUSTOM"
)

func (r ContractRateCardRateAddManyParamsRatesCommitRateRateType) IsKnown() bool {
	switch r {
	case ContractRateCardRateAddManyParamsRatesCommitRateRateTypeFlat, ContractRateCardRateAddManyParamsRatesCommitRateRateTypePercentage, ContractRateCardRateAddManyParamsRatesCommitRateRateTypeSubscription, ContractRateCardRateAddManyParamsRatesCommitRateRateTypeTiered, ContractRateCardRateAddManyParamsRatesCommitRateRateTypeCustom:
		return true
	}
	return false
}
