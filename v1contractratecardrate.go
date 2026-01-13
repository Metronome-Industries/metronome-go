// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v2/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v2/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/v2/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/v2/packages/param"
	"github.com/Metronome-Industries/metronome-go/v2/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/v2/shared"
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
func NewV1ContractRateCardRateService(opts ...option.RequestOption) (r V1ContractRateCardRateService) {
	r = V1ContractRateCardRateService{}
	r.Options = opts
	return
}

// Understand the rate schedule at a given timestamp, optionally filtering the list
// of rates returned based on properties such as `product_id` and
// `pricing_group_values`. For example, you may want to display the current price
// for a given product in your product experience - use this endpoint to fetch that
// information from its source of truth in Metronome.
//
// If you want to understand the rates for a specific customer's contract,
// inclusive of contract-level overrides, use the `getContractRateSchedule`
// endpoint.
func (r *V1ContractRateCardRateService) List(ctx context.Context, params V1ContractRateCardRateListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1ContractRateCardRateListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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

// Understand the rate schedule at a given timestamp, optionally filtering the list
// of rates returned based on properties such as `product_id` and
// `pricing_group_values`. For example, you may want to display the current price
// for a given product in your product experience - use this endpoint to fetch that
// information from its source of truth in Metronome.
//
// If you want to understand the rates for a specific customer's contract,
// inclusive of contract-level overrides, use the `getContractRateSchedule`
// endpoint.
func (r *V1ContractRateCardRateService) ListAutoPaging(ctx context.Context, params V1ContractRateCardRateListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1ContractRateCardRateListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Add a new rate
func (r *V1ContractRateCardRateService) Add(ctx context.Context, body V1ContractRateCardRateAddParams, opts ...option.RequestOption) (res *V1ContractRateCardRateAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contract-pricing/rate-cards/addRate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Add new rates
func (r *V1ContractRateCardRateService) AddMany(ctx context.Context, body V1ContractRateCardRateAddManyParams, opts ...option.RequestOption) (res *V1ContractRateCardRateAddManyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contract-pricing/rate-cards/addRates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1ContractRateCardRateListResponse struct {
	Entitled bool `json:"entitled,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	ProductCustomFields map[string]string `json:"product_custom_fields,required"`
	ProductID           string            `json:"product_id,required" format:"uuid"`
	ProductName         string            `json:"product_name,required"`
	ProductTags         []string          `json:"product_tags,required"`
	Rate                shared.Rate       `json:"rate,required"`
	StartingAt          time.Time         `json:"starting_at,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency V1ContractRateCardRateListResponseBillingFrequency `json:"billing_frequency"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate         shared.CommitRate `json:"commit_rate"`
	EndingBefore       time.Time         `json:"ending_before" format:"date-time"`
	PricingGroupValues map[string]string `json:"pricing_group_values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entitled            respjson.Field
		ProductCustomFields respjson.Field
		ProductID           respjson.Field
		ProductName         respjson.Field
		ProductTags         respjson.Field
		Rate                respjson.Field
		StartingAt          respjson.Field
		BillingFrequency    respjson.Field
		CommitRate          respjson.Field
		EndingBefore        respjson.Field
		PricingGroupValues  respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardRateListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardRateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardRateListResponseBillingFrequency string

const (
	V1ContractRateCardRateListResponseBillingFrequencyMonthly   V1ContractRateCardRateListResponseBillingFrequency = "MONTHLY"
	V1ContractRateCardRateListResponseBillingFrequencyQuarterly V1ContractRateCardRateListResponseBillingFrequency = "QUARTERLY"
	V1ContractRateCardRateListResponseBillingFrequencyAnnual    V1ContractRateCardRateListResponseBillingFrequency = "ANNUAL"
	V1ContractRateCardRateListResponseBillingFrequencyWeekly    V1ContractRateCardRateListResponseBillingFrequency = "WEEKLY"
)

type V1ContractRateCardRateAddResponse struct {
	Data V1ContractRateCardRateAddResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardRateAddResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardRateAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardRateAddResponseData struct {
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "CUSTOM", "TIERED".
	RateType string `json:"rate_type,required"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate shared.CommitRate     `json:"commit_rate"`
	CreditType shared.CreditTypeData `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]any `json:"custom_rate"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RateType           respjson.Field
		CommitRate         respjson.Field
		CreditType         respjson.Field
		CustomRate         respjson.Field
		IsProrated         respjson.Field
		Price              respjson.Field
		PricingGroupValues respjson.Field
		Quantity           respjson.Field
		Tiers              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardRateAddResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardRateAddResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardRateAddManyResponse struct {
	// The ID of the rate card to which the rates were added.
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardRateAddManyResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardRateAddManyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardRateListParams struct {
	// inclusive starting point for the rates schedule
	At time.Time `json:"at,required" format:"date-time"`
	// ID of the rate card to get the schedule for
	RateCardID string `json:"rate_card_id,required" format:"uuid"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// List of rate selectors, rates matching ANY of the selector will be included in
	// the response Passing no selectors will result in all rates being returned.
	Selectors []V1ContractRateCardRateListParamsSelector `json:"selectors,omitzero"`
	paramObj
}

func (r V1ContractRateCardRateListParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardRateListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardRateListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1ContractRateCardRateListParams]'s query parameters as
// `url.Values`.
func (r V1ContractRateCardRateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ContractRateCardRateListParamsSelector struct {
	// Rates matching the product id will be included in the response.
	ProductID param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	// Subscription rates matching the billing frequency will be included in the
	// response.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero"`
	// List of pricing group key value pairs, rates containing the matching key / value
	// pairs will be included in the response.
	PartialPricingGroupValues map[string]string `json:"partial_pricing_group_values,omitzero"`
	// List of pricing group key value pairs, rates matching all of the key / value
	// pairs will be included in the response.
	PricingGroupValues map[string]string `json:"pricing_group_values,omitzero"`
	// List of product tags, rates matching any of the tags will be included in the
	// response.
	ProductTags []string `json:"product_tags,omitzero"`
	paramObj
}

func (r V1ContractRateCardRateListParamsSelector) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardRateListParamsSelector
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardRateListParamsSelector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractRateCardRateListParamsSelector](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

type V1ContractRateCardRateAddParams struct {
	Entitled bool `json:"entitled,required"`
	// ID of the product to add a rate for
	ProductID string `json:"product_id,required" format:"uuid"`
	// ID of the rate card to update
	RateCardID string `json:"rate_card_id,required" format:"uuid"`
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM".
	RateType V1ContractRateCardRateAddParamsRateType `json:"rate_type,omitzero,required"`
	// inclusive effective date
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// The Metronome ID of the credit type to associate with price, defaults to USD
	// (cents) if not passed. Used by all rate_types except type PERCENTAGE. PERCENTAGE
	// rates use the credit type of associated rates.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// exclusive end date
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated param.Opt[bool] `json:"is_prorated,omitzero"`
	// Default price. For FLAT and SUBSCRIPTION rate_type, this must be >=0. For
	// PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this
	// must be >=0 and <=1.
	Price param.Opt[float64] `json:"price,omitzero"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Optional. Frequency to bill subscriptions with. Required for subscription type
	// products with Flat rate.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency V1ContractRateCardRateAddParamsBillingFrequency `json:"billing_frequency,omitzero"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate shared.CommitRateParam `json:"commit_rate,omitzero"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]any `json:"custom_rate,omitzero"`
	// Optional. List of pricing group key value pairs which will be used to calculate
	// the price.
	PricingGroupValues map[string]string `json:"pricing_group_values,omitzero"`
	// Only set for TIERED rate_type.
	Tiers []shared.TierParam `json:"tiers,omitzero"`
	paramObj
}

func (r V1ContractRateCardRateAddParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardRateAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardRateAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardRateAddParamsRateType string

const (
	V1ContractRateCardRateAddParamsRateTypeFlat         V1ContractRateCardRateAddParamsRateType = "FLAT"
	V1ContractRateCardRateAddParamsRateTypePercentage   V1ContractRateCardRateAddParamsRateType = "PERCENTAGE"
	V1ContractRateCardRateAddParamsRateTypeSubscription V1ContractRateCardRateAddParamsRateType = "SUBSCRIPTION"
	V1ContractRateCardRateAddParamsRateTypeTiered       V1ContractRateCardRateAddParamsRateType = "TIERED"
	V1ContractRateCardRateAddParamsRateTypeCustom       V1ContractRateCardRateAddParamsRateType = "CUSTOM"
)

// Optional. Frequency to bill subscriptions with. Required for subscription type
// products with Flat rate.
type V1ContractRateCardRateAddParamsBillingFrequency string

const (
	V1ContractRateCardRateAddParamsBillingFrequencyMonthly   V1ContractRateCardRateAddParamsBillingFrequency = "MONTHLY"
	V1ContractRateCardRateAddParamsBillingFrequencyQuarterly V1ContractRateCardRateAddParamsBillingFrequency = "QUARTERLY"
	V1ContractRateCardRateAddParamsBillingFrequencyAnnual    V1ContractRateCardRateAddParamsBillingFrequency = "ANNUAL"
	V1ContractRateCardRateAddParamsBillingFrequencyWeekly    V1ContractRateCardRateAddParamsBillingFrequency = "WEEKLY"
)

type V1ContractRateCardRateAddManyParams struct {
	RateCardID string                                    `json:"rate_card_id,required" format:"uuid"`
	Rates      []V1ContractRateCardRateAddManyParamsRate `json:"rates,omitzero,required"`
	paramObj
}

func (r V1ContractRateCardRateAddManyParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardRateAddManyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardRateAddManyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Entitled, ProductID, RateType, StartingAt are required.
type V1ContractRateCardRateAddManyParamsRate struct {
	Entitled bool `json:"entitled,required"`
	// ID of the product to add a rate for
	ProductID string `json:"product_id,required" format:"uuid"`
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM".
	RateType string `json:"rate_type,omitzero,required"`
	// inclusive effective date
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// "The Metronome ID of the credit type to associate with price, defaults to USD
	// (cents) if not passed. Used by all rate_types except type PERCENTAGE. PERCENTAGE
	// rates use the credit type of associated rates."
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// exclusive end date
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated param.Opt[bool] `json:"is_prorated,omitzero"`
	// Default price. For FLAT and SUBSCRIPTION rate_type, this must be >=0. For
	// PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this
	// must be >=0 and <=1.
	Price param.Opt[float64] `json:"price,omitzero"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Optional. Frequency to bill subscriptions with. Required for subscription type
	// products with Flat rate.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate shared.CommitRateParam `json:"commit_rate,omitzero"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]any `json:"custom_rate,omitzero"`
	// Optional. List of pricing group key value pairs which will be used to calculate
	// the price.
	PricingGroupValues map[string]string `json:"pricing_group_values,omitzero"`
	// Only set for TIERED rate_type.
	Tiers []shared.TierParam `json:"tiers,omitzero"`
	paramObj
}

func (r V1ContractRateCardRateAddManyParamsRate) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardRateAddManyParamsRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardRateAddManyParamsRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractRateCardRateAddManyParamsRate](
		"rate_type", "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM",
	)
	apijson.RegisterFieldValidator[V1ContractRateCardRateAddManyParamsRate](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}
