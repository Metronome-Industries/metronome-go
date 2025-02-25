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

// ContractRateCardService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContractRateCardService] method instead.
type ContractRateCardService struct {
	Options        []option.RequestOption
	ProductOrders  *ContractRateCardProductOrderService
	Rates          *ContractRateCardRateService
	NamedSchedules *ContractRateCardNamedScheduleService
}

// NewContractRateCardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContractRateCardService(opts ...option.RequestOption) (r *ContractRateCardService) {
	r = &ContractRateCardService{}
	r.Options = opts
	r.ProductOrders = NewContractRateCardProductOrderService(opts...)
	r.Rates = NewContractRateCardRateService(opts...)
	r.NamedSchedules = NewContractRateCardNamedScheduleService(opts...)
	return
}

// Create a new rate card
func (r *ContractRateCardService) New(ctx context.Context, body ContractRateCardNewParams, opts ...option.RequestOption) (res *ContractRateCardNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/rate-cards/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific rate card NOTE: Use `/contract-pricing/rate-cards/getRates` to
// retrieve rate card rates.
func (r *ContractRateCardService) Get(ctx context.Context, body ContractRateCardGetParams, opts ...option.RequestOption) (res *ContractRateCardGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/rate-cards/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a rate card
func (r *ContractRateCardService) Update(ctx context.Context, body ContractRateCardUpdateParams, opts ...option.RequestOption) (res *ContractRateCardUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/rate-cards/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List rate cards NOTE: Use `/contract-pricing/rate-cards/getRates` to retrieve
// rate card rates.
func (r *ContractRateCardService) List(ctx context.Context, params ContractRateCardListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ContractRateCardListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "contract-pricing/rate-cards/list"
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

// List rate cards NOTE: Use `/contract-pricing/rate-cards/getRates` to retrieve
// rate card rates.
func (r *ContractRateCardService) ListAutoPaging(ctx context.Context, params ContractRateCardListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ContractRateCardListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Get all rates for a rate card from starting_at (either in perpetuity or until
// ending_before, if provided)
func (r *ContractRateCardService) GetRateSchedule(ctx context.Context, params ContractRateCardGetRateScheduleParams, opts ...option.RequestOption) (res *ContractRateCardGetRateScheduleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/rate-cards/getRateSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ContractRateCardNewResponse struct {
	Data shared.ID                       `json:"data,required"`
	JSON contractRateCardNewResponseJSON `json:"-"`
}

// contractRateCardNewResponseJSON contains the JSON metadata for the struct
// [ContractRateCardNewResponse]
type contractRateCardNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractRateCardNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardNewResponseJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardGetResponse struct {
	Data ContractRateCardGetResponseData `json:"data,required"`
	JSON contractRateCardGetResponseJSON `json:"-"`
}

// contractRateCardGetResponseJSON contains the JSON metadata for the struct
// [ContractRateCardGetResponse]
type contractRateCardGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractRateCardGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardGetResponseJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardGetResponseData struct {
	ID                    string                                                `json:"id,required" format:"uuid"`
	CreatedAt             time.Time                                             `json:"created_at,required" format:"date-time"`
	CreatedBy             string                                                `json:"created_by,required"`
	Name                  string                                                `json:"name,required"`
	Aliases               []ContractRateCardGetResponseDataAlias                `json:"aliases"`
	CreditTypeConversions []ContractRateCardGetResponseDataCreditTypeConversion `json:"credit_type_conversions"`
	CustomFields          map[string]string                                     `json:"custom_fields"`
	Description           string                                                `json:"description"`
	FiatCreditType        shared.CreditTypeData                                 `json:"fiat_credit_type"`
	JSON                  contractRateCardGetResponseDataJSON                   `json:"-"`
}

// contractRateCardGetResponseDataJSON contains the JSON metadata for the struct
// [ContractRateCardGetResponseData]
type contractRateCardGetResponseDataJSON struct {
	ID                    apijson.Field
	CreatedAt             apijson.Field
	CreatedBy             apijson.Field
	Name                  apijson.Field
	Aliases               apijson.Field
	CreditTypeConversions apijson.Field
	CustomFields          apijson.Field
	Description           apijson.Field
	FiatCreditType        apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractRateCardGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardGetResponseDataAlias struct {
	Name         string                                   `json:"name,required"`
	EndingBefore time.Time                                `json:"ending_before" format:"date-time"`
	StartingAt   time.Time                                `json:"starting_at" format:"date-time"`
	JSON         contractRateCardGetResponseDataAliasJSON `json:"-"`
}

// contractRateCardGetResponseDataAliasJSON contains the JSON metadata for the
// struct [ContractRateCardGetResponseDataAlias]
type contractRateCardGetResponseDataAliasJSON struct {
	Name         apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractRateCardGetResponseDataAlias) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardGetResponseDataAliasJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardGetResponseDataCreditTypeConversion struct {
	CustomCreditType    shared.CreditTypeData                                   `json:"custom_credit_type,required"`
	FiatPerCustomCredit string                                                  `json:"fiat_per_custom_credit,required"`
	JSON                contractRateCardGetResponseDataCreditTypeConversionJSON `json:"-"`
}

// contractRateCardGetResponseDataCreditTypeConversionJSON contains the JSON
// metadata for the struct [ContractRateCardGetResponseDataCreditTypeConversion]
type contractRateCardGetResponseDataCreditTypeConversionJSON struct {
	CustomCreditType    apijson.Field
	FiatPerCustomCredit apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ContractRateCardGetResponseDataCreditTypeConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardGetResponseDataCreditTypeConversionJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardUpdateResponse struct {
	Data shared.ID                          `json:"data,required"`
	JSON contractRateCardUpdateResponseJSON `json:"-"`
}

// contractRateCardUpdateResponseJSON contains the JSON metadata for the struct
// [ContractRateCardUpdateResponse]
type contractRateCardUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractRateCardUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardListResponse struct {
	ID                    string                                             `json:"id,required" format:"uuid"`
	CreatedAt             time.Time                                          `json:"created_at,required" format:"date-time"`
	CreatedBy             string                                             `json:"created_by,required"`
	Name                  string                                             `json:"name,required"`
	Aliases               []ContractRateCardListResponseAlias                `json:"aliases"`
	CreditTypeConversions []ContractRateCardListResponseCreditTypeConversion `json:"credit_type_conversions"`
	CustomFields          map[string]string                                  `json:"custom_fields"`
	Description           string                                             `json:"description"`
	FiatCreditType        shared.CreditTypeData                              `json:"fiat_credit_type"`
	JSON                  contractRateCardListResponseJSON                   `json:"-"`
}

// contractRateCardListResponseJSON contains the JSON metadata for the struct
// [ContractRateCardListResponse]
type contractRateCardListResponseJSON struct {
	ID                    apijson.Field
	CreatedAt             apijson.Field
	CreatedBy             apijson.Field
	Name                  apijson.Field
	Aliases               apijson.Field
	CreditTypeConversions apijson.Field
	CustomFields          apijson.Field
	Description           apijson.Field
	FiatCreditType        apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractRateCardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardListResponseJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardListResponseAlias struct {
	Name         string                                `json:"name,required"`
	EndingBefore time.Time                             `json:"ending_before" format:"date-time"`
	StartingAt   time.Time                             `json:"starting_at" format:"date-time"`
	JSON         contractRateCardListResponseAliasJSON `json:"-"`
}

// contractRateCardListResponseAliasJSON contains the JSON metadata for the struct
// [ContractRateCardListResponseAlias]
type contractRateCardListResponseAliasJSON struct {
	Name         apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractRateCardListResponseAlias) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardListResponseAliasJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardListResponseCreditTypeConversion struct {
	CustomCreditType    shared.CreditTypeData                                `json:"custom_credit_type,required"`
	FiatPerCustomCredit string                                               `json:"fiat_per_custom_credit,required"`
	JSON                contractRateCardListResponseCreditTypeConversionJSON `json:"-"`
}

// contractRateCardListResponseCreditTypeConversionJSON contains the JSON metadata
// for the struct [ContractRateCardListResponseCreditTypeConversion]
type contractRateCardListResponseCreditTypeConversionJSON struct {
	CustomCreditType    apijson.Field
	FiatPerCustomCredit apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ContractRateCardListResponseCreditTypeConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardListResponseCreditTypeConversionJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardGetRateScheduleResponse struct {
	Data     []ContractRateCardGetRateScheduleResponseData `json:"data,required"`
	NextPage string                                        `json:"next_page,nullable"`
	JSON     contractRateCardGetRateScheduleResponseJSON   `json:"-"`
}

// contractRateCardGetRateScheduleResponseJSON contains the JSON metadata for the
// struct [ContractRateCardGetRateScheduleResponse]
type contractRateCardGetRateScheduleResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractRateCardGetRateScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardGetRateScheduleResponseJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardGetRateScheduleResponseData struct {
	Entitled            bool              `json:"entitled,required"`
	ProductCustomFields map[string]string `json:"product_custom_fields,required"`
	ProductID           string            `json:"product_id,required" format:"uuid"`
	ProductName         string            `json:"product_name,required"`
	ProductTags         []string          `json:"product_tags,required"`
	Rate                shared.Rate       `json:"rate,required"`
	StartingAt          time.Time         `json:"starting_at,required" format:"date-time"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate         ContractRateCardGetRateScheduleResponseDataCommitRate `json:"commit_rate"`
	EndingBefore       time.Time                                             `json:"ending_before" format:"date-time"`
	PricingGroupValues map[string]string                                     `json:"pricing_group_values"`
	JSON               contractRateCardGetRateScheduleResponseDataJSON       `json:"-"`
}

// contractRateCardGetRateScheduleResponseDataJSON contains the JSON metadata for
// the struct [ContractRateCardGetRateScheduleResponseData]
type contractRateCardGetRateScheduleResponseDataJSON struct {
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

func (r *ContractRateCardGetRateScheduleResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardGetRateScheduleResponseDataJSON) RawJSON() string {
	return r.raw
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type ContractRateCardGetRateScheduleResponseDataCommitRate struct {
	RateType ContractRateCardGetRateScheduleResponseDataCommitRateRateType `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price float64 `json:"price"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier                                             `json:"tiers"`
	JSON  contractRateCardGetRateScheduleResponseDataCommitRateJSON `json:"-"`
}

// contractRateCardGetRateScheduleResponseDataCommitRateJSON contains the JSON
// metadata for the struct [ContractRateCardGetRateScheduleResponseDataCommitRate]
type contractRateCardGetRateScheduleResponseDataCommitRateJSON struct {
	RateType    apijson.Field
	Price       apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractRateCardGetRateScheduleResponseDataCommitRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractRateCardGetRateScheduleResponseDataCommitRateJSON) RawJSON() string {
	return r.raw
}

type ContractRateCardGetRateScheduleResponseDataCommitRateRateType string

const (
	ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeFlat         ContractRateCardGetRateScheduleResponseDataCommitRateRateType = "FLAT"
	ContractRateCardGetRateScheduleResponseDataCommitRateRateTypePercentage   ContractRateCardGetRateScheduleResponseDataCommitRateRateType = "PERCENTAGE"
	ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeSubscription ContractRateCardGetRateScheduleResponseDataCommitRateRateType = "SUBSCRIPTION"
	ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeTiered       ContractRateCardGetRateScheduleResponseDataCommitRateRateType = "TIERED"
	ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeCustom       ContractRateCardGetRateScheduleResponseDataCommitRateRateType = "CUSTOM"
)

func (r ContractRateCardGetRateScheduleResponseDataCommitRateRateType) IsKnown() bool {
	switch r {
	case ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeFlat, ContractRateCardGetRateScheduleResponseDataCommitRateRateTypePercentage, ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeSubscription, ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeTiered, ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeCustom:
		return true
	}
	return false
}

type ContractRateCardNewParams struct {
	// Used only in UI/API. It is not exposed to end customers.
	Name param.Field[string] `json:"name,required"`
	// Reference this alias when creating a contract. If the same alias is assigned to
	// multiple rate cards, it will reference the rate card to which it was most
	// recently assigned. It is not exposed to end customers.
	Aliases param.Field[[]ContractRateCardNewParamsAlias] `json:"aliases"`
	// Required when using custom pricing units in rates.
	CreditTypeConversions param.Field[[]ContractRateCardNewParamsCreditTypeConversion] `json:"credit_type_conversions"`
	CustomFields          param.Field[map[string]string]                               `json:"custom_fields"`
	Description           param.Field[string]                                          `json:"description"`
	// The Metronome ID of the credit type to associate with the rate card, defaults to
	// USD (cents) if not passed.
	FiatCreditTypeID param.Field[string] `json:"fiat_credit_type_id" format:"uuid"`
}

func (r ContractRateCardNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardNewParamsAlias struct {
	Name         param.Field[string]    `json:"name,required"`
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r ContractRateCardNewParamsAlias) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardNewParamsCreditTypeConversion struct {
	CustomCreditTypeID  param.Field[string]  `json:"custom_credit_type_id,required" format:"uuid"`
	FiatPerCustomCredit param.Field[float64] `json:"fiat_per_custom_credit,required"`
}

func (r ContractRateCardNewParamsCreditTypeConversion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardGetParams struct {
	ID shared.IDParam `json:"id,required"`
}

func (r ContractRateCardGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ID)
}

type ContractRateCardUpdateParams struct {
	// ID of the rate card to update
	RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
	// Reference this alias when creating a contract. If the same alias is assigned to
	// multiple rate cards, it will reference the rate card to which it was most
	// recently assigned. It is not exposed to end customers.
	Aliases     param.Field[[]ContractRateCardUpdateParamsAlias] `json:"aliases"`
	Description param.Field[string]                              `json:"description"`
	// Used only in UI/API. It is not exposed to end customers.
	Name param.Field[string] `json:"name"`
}

func (r ContractRateCardUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardUpdateParamsAlias struct {
	Name         param.Field[string]    `json:"name,required"`
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r ContractRateCardUpdateParamsAlias) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractRateCardListParams struct {
	Body interface{} `json:"body,required"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

func (r ContractRateCardListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [ContractRateCardListParams]'s query parameters as
// `url.Values`.
func (r ContractRateCardListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContractRateCardGetRateScheduleParams struct {
	// ID of the rate card to get the schedule for
	RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
	// inclusive starting point for the rates schedule
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// optional exclusive end date for the rates schedule. When not specified rates
	// will show all future schedule segments.
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// List of rate selectors, rates matching ANY of the selector will be included in
	// the response Passing no selectors will result in all rates being returned.
	Selectors param.Field[[]ContractRateCardGetRateScheduleParamsSelector] `json:"selectors"`
}

func (r ContractRateCardGetRateScheduleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ContractRateCardGetRateScheduleParams]'s query parameters
// as `url.Values`.
func (r ContractRateCardGetRateScheduleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContractRateCardGetRateScheduleParamsSelector struct {
	// List of pricing group key value pairs, rates containing the matching key / value
	// pairs will be included in the response.
	PartialPricingGroupValues param.Field[map[string]string] `json:"partial_pricing_group_values"`
	// List of pricing group key value pairs, rates matching all of the key / value
	// pairs will be included in the response.
	PricingGroupValues param.Field[map[string]string] `json:"pricing_group_values"`
	// Rates matching the product id will be included in the response.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
}

func (r ContractRateCardGetRateScheduleParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
