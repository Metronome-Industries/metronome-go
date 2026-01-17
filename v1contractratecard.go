// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v3/internal/apiquery"
	shimjson "github.com/Metronome-Industries/metronome-go/v3/internal/encoding/json"
	"github.com/Metronome-Industries/metronome-go/v3/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"github.com/Metronome-Industries/metronome-go/v3/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/v3/packages/param"
	"github.com/Metronome-Industries/metronome-go/v3/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/v3/shared"
)

// V1ContractRateCardService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContractRateCardService] method instead.
type V1ContractRateCardService struct {
	Options        []option.RequestOption
	ProductOrders  V1ContractRateCardProductOrderService
	Rates          V1ContractRateCardRateService
	NamedSchedules V1ContractRateCardNamedScheduleService
}

// NewV1ContractRateCardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1ContractRateCardService(opts ...option.RequestOption) (r V1ContractRateCardService) {
	r = V1ContractRateCardService{}
	r.Options = opts
	r.ProductOrders = NewV1ContractRateCardProductOrderService(opts...)
	r.Rates = NewV1ContractRateCardRateService(opts...)
	r.NamedSchedules = NewV1ContractRateCardNamedScheduleService(opts...)
	return
}

// In Metronome, the rate card is the central location for your pricing. Rate cards
// were built with new product launches and pricing changes in mind - you can
// update your products and pricing in one place, and that change will be
// automatically propagated across your customer cohorts. Most clients need only
// maintain one or a few rate cards within Metronome.
//
// ### Use this endpoint to:
//
//   - Create a rate card with a name and description
//   - Define the rate card's single underlying fiat currency, and any number of
//     conversion rates between that fiat currency and custom pricing units. You can
//     then add products and associated rates in the fiat currency or custom pricing
//     unit for which you have defined a conversion rate.
//   - Set aliases for the rate card. Aliases are human-readable names that you can
//     use in the place of the id of the rate card when provisioning a customer's
//     contract. By using an alias, you can easily create a contract and provision a
//     customer by choosing the paygo rate card, without storing the rate card id in
//     your internal systems. This is helpful when launching a new rate card for
//     paygo customers, you can update the alias for paygo to be scheduled to be
//     assigned to the new rate card without updating your code.
//
// ### Key response fields:
//
// - The ID of the rate card you just created
//
// ### Usage guidelines:
//
//   - After creating a rate card, you can now use the addRate or addRates endpoints
//     to add products and their prices to it
//   - A rate card alias can only be used by one rate card at a time. If you create a
//     contract with a rate card alias that is already in use by another rate card,
//     the original rate card's alias schedule will be updated. The alias will
//     reference the rate card to which it was most recently assigned.
func (r *V1ContractRateCardService) New(ctx context.Context, body V1ContractRateCardNewParams, opts ...option.RequestOption) (res *V1ContractRateCardNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contract-pricing/rate-cards/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return details for a specific rate card including name, description, and
// aliases. This endpoint does not return rates - use the dedicated getRates or
// getRateSchedule endpoints to understand the rates on a rate card.
func (r *V1ContractRateCardService) Get(ctx context.Context, body V1ContractRateCardGetParams, opts ...option.RequestOption) (res *V1ContractRateCardGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contract-pricing/rate-cards/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the metadata properties of an existing rate card, including its name,
// description, and aliases. This endpoint is designed for managing rate card
// identity and reference aliases rather than modifying pricing rates.
//
// Modifies the descriptive properties and alias configuration of a rate card
// without affecting the underlying pricing rates or schedules. This allows you to
// update how a rate card is identified and referenced throughout your system.
//
// ### Use this endpoint to:
//
//   - Rate card renaming: Update display names or descriptions for organizational
//     clarity
//   - Alias management: Add, modify, or schedule alias transitions for seamless rate
//     card migrations
//   - Documentation updates: Keep rate card descriptions current with business
//     context
//   - Self-serve provisioning setup: Configure aliases to enable code-free rate card
//     transitions
//
// #### Active contract impact:
//
//   - Alias changes: Already-created contracts continue using their originally
//     assigned rate cards.
//   - Other changes made using this endpoint will only impact the Metronome UI.
//
// #### Grandfathering existing PLG customer pricing:
//
//   - Rate card aliases support scheduled transitions, enabling seamless rate card
//     migrations for new customers, allowing existing customers to be grandfathered
//     into their existing prices without code. Note that there are multiple
//     mechanisms to support grandfathering in Metronome.
//
// #### How scheduled aliases work for PLG grandfathering:
//
// Initial setup:
//
//   - Add alias to current rate card: Assign a stable alias (e.g.,
//     "standard-pricing") to your active rate card
//   - Reference alias during contract creation: Configure your self-serve workflow
//     to create contracts using `rate_card_alias` instead of direct `rate_card_id`
//   - Automatic resolution: New contracts referencing the alias automatically
//     resolve to the rate card associated with the alias at the point in time of
//     provisioning
//
// #### Grandfathering process:
//
//   - Create new rate card: Build your new rate card with updated pricing structure
//   - Schedule alias transition: Add the same alias to the new rate card with a
//     `starting_at` timestamp
//   - Automatic cutover: Starting at the scheduled time, new contracts created in
//     your PLG workflow using that alias will automatically reference the new rate
//     card
func (r *V1ContractRateCardService) Update(ctx context.Context, body V1ContractRateCardUpdateParams, opts ...option.RequestOption) (res *V1ContractRateCardUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contract-pricing/rate-cards/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all rate cards. Returns rate card IDs, names, descriptions, aliases, and
// other details. To view the rates associated with a given rate card, use the
// getRates or getRateSchedule endpoints.
func (r *V1ContractRateCardService) List(ctx context.Context, params V1ContractRateCardListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1ContractRateCardListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/contract-pricing/rate-cards/list"
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

// List all rate cards. Returns rate card IDs, names, descriptions, aliases, and
// other details. To view the rates associated with a given rate card, use the
// getRates or getRateSchedule endpoints.
func (r *V1ContractRateCardService) ListAutoPaging(ctx context.Context, params V1ContractRateCardListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1ContractRateCardListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Permanently disable a rate card by archiving it, preventing use in new contracts
// while preserving existing contract pricing. Use this when retiring old pricing
// models, consolidating rate cards, or removing outdated pricing structures.
// Returns the archived rate card ID and stops the rate card from appearing in
// contract creation workflows.
func (r *V1ContractRateCardService) Archive(ctx context.Context, body V1ContractRateCardArchiveParams, opts ...option.RequestOption) (res *V1ContractRateCardArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contract-pricing/rate-cards/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A rate card defines the prices that you charge for your products. Rate cards
// support scheduled changes over time, to allow you to easily roll out pricing
// changes and new product launches across your customer base. Use this endpoint to
// understand the rate schedule `starting_at` a given date, optionally filtering
// the list of rates returned based on product id or pricing group values. For
// example, you may want to display a schedule of upcoming price changes for a
// given product in your product experience - use this endpoint to fetch that
// information from its source of truth in Metronome.
//
// If you want to understand the rates for a specific customer's contract,
// inclusive of contract-level overrides, use the `getContractRateSchedule`
// endpoint.
func (r *V1ContractRateCardService) GetRateSchedule(ctx context.Context, params V1ContractRateCardGetRateScheduleParams, opts ...option.RequestOption) (res *V1ContractRateCardGetRateScheduleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contract-pricing/rate-cards/getRateSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type V1ContractRateCardNewResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardGetResponse struct {
	Data V1ContractRateCardGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardGetResponseData struct {
	ID                    string                                                  `json:"id,required" format:"uuid"`
	CreatedAt             time.Time                                               `json:"created_at,required" format:"date-time"`
	CreatedBy             string                                                  `json:"created_by,required"`
	Name                  string                                                  `json:"name,required"`
	Aliases               []V1ContractRateCardGetResponseDataAlias                `json:"aliases"`
	CreditTypeConversions []V1ContractRateCardGetResponseDataCreditTypeConversion `json:"credit_type_conversions"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields   map[string]string     `json:"custom_fields"`
	Description    string                `json:"description"`
	FiatCreditType shared.CreditTypeData `json:"fiat_credit_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Name                  respjson.Field
		Aliases               respjson.Field
		CreditTypeConversions respjson.Field
		CustomFields          respjson.Field
		Description           respjson.Field
		FiatCreditType        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardGetResponseDataAlias struct {
	Name         string    `json:"name,required"`
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardGetResponseDataAlias) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardGetResponseDataAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardGetResponseDataCreditTypeConversion struct {
	CustomCreditType    shared.CreditTypeData `json:"custom_credit_type,required"`
	FiatPerCustomCredit string                `json:"fiat_per_custom_credit,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomCreditType    respjson.Field
		FiatPerCustomCredit respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardGetResponseDataCreditTypeConversion) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardGetResponseDataCreditTypeConversion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardUpdateResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardListResponse struct {
	ID                    string                                               `json:"id,required" format:"uuid"`
	CreatedAt             time.Time                                            `json:"created_at,required" format:"date-time"`
	CreatedBy             string                                               `json:"created_by,required"`
	Name                  string                                               `json:"name,required"`
	Aliases               []V1ContractRateCardListResponseAlias                `json:"aliases"`
	CreditTypeConversions []V1ContractRateCardListResponseCreditTypeConversion `json:"credit_type_conversions"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields   map[string]string     `json:"custom_fields"`
	Description    string                `json:"description"`
	FiatCreditType shared.CreditTypeData `json:"fiat_credit_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Name                  respjson.Field
		Aliases               respjson.Field
		CreditTypeConversions respjson.Field
		CustomFields          respjson.Field
		Description           respjson.Field
		FiatCreditType        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardListResponseAlias struct {
	Name         string    `json:"name,required"`
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardListResponseAlias) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardListResponseAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardListResponseCreditTypeConversion struct {
	CustomCreditType    shared.CreditTypeData `json:"custom_credit_type,required"`
	FiatPerCustomCredit string                `json:"fiat_per_custom_credit,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomCreditType    respjson.Field
		FiatPerCustomCredit respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardListResponseCreditTypeConversion) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardListResponseCreditTypeConversion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardArchiveResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardGetRateScheduleResponse struct {
	Data     []V1ContractRateCardGetRateScheduleResponseData `json:"data,required"`
	NextPage string                                          `json:"next_page,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextPage    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractRateCardGetRateScheduleResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardGetRateScheduleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardGetRateScheduleResponseData struct {
	Entitled bool `json:"entitled,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	ProductCustomFields map[string]string `json:"product_custom_fields,required"`
	ProductID           string            `json:"product_id,required" format:"uuid"`
	ProductName         string            `json:"product_name,required"`
	ProductTags         []string          `json:"product_tags,required"`
	Rate                shared.Rate       `json:"rate,required"`
	StartingAt          time.Time         `json:"starting_at,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency"`
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
func (r V1ContractRateCardGetRateScheduleResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ContractRateCardGetRateScheduleResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardNewParams struct {
	// Used only in UI/API. It is not exposed to end customers.
	Name        string            `json:"name,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	// The Metronome ID of the credit type to associate with the rate card, defaults to
	// USD (cents) if not passed.
	FiatCreditTypeID param.Opt[string] `json:"fiat_credit_type_id,omitzero" format:"uuid"`
	// Reference this alias when creating a contract. If the same alias is assigned to
	// multiple rate cards, it will reference the rate card to which it was most
	// recently assigned. It is not exposed to end customers.
	Aliases []V1ContractRateCardNewParamsAlias `json:"aliases,omitzero"`
	// Required when using custom pricing units in rates.
	CreditTypeConversions []V1ContractRateCardNewParamsCreditTypeConversion `json:"credit_type_conversions,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V1ContractRateCardNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type V1ContractRateCardNewParamsAlias struct {
	Name         string               `json:"name,required"`
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	StartingAt   param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V1ContractRateCardNewParamsAlias) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardNewParamsAlias
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardNewParamsAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CustomCreditTypeID, FiatPerCustomCredit are required.
type V1ContractRateCardNewParamsCreditTypeConversion struct {
	CustomCreditTypeID  string  `json:"custom_credit_type_id,required" format:"uuid"`
	FiatPerCustomCredit float64 `json:"fiat_per_custom_credit,required"`
	paramObj
}

func (r V1ContractRateCardNewParamsCreditTypeConversion) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardNewParamsCreditTypeConversion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardNewParamsCreditTypeConversion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardGetParams struct {
	ID shared.IDParam
	paramObj
}

func (r V1ContractRateCardGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ID)
}
func (r *V1ContractRateCardGetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ID)
}

type V1ContractRateCardUpdateParams struct {
	// ID of the rate card to update
	RateCardID  string            `json:"rate_card_id,required" format:"uuid"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Used only in UI/API. It is not exposed to end customers.
	Name param.Opt[string] `json:"name,omitzero"`
	// Reference this alias when creating a contract. If the same alias is assigned to
	// multiple rate cards, it will reference the rate card to which it was most
	// recently assigned. It is not exposed to end customers.
	Aliases []V1ContractRateCardUpdateParamsAlias `json:"aliases,omitzero"`
	paramObj
}

func (r V1ContractRateCardUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type V1ContractRateCardUpdateParamsAlias struct {
	Name         string               `json:"name,required"`
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	StartingAt   param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V1ContractRateCardUpdateParamsAlias) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardUpdateParamsAlias
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardUpdateParamsAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractRateCardListParams struct {
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	Body     any
	paramObj
}

func (r V1ContractRateCardListParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *V1ContractRateCardListParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// URLQuery serializes [V1ContractRateCardListParams]'s query parameters as
// `url.Values`.
func (r V1ContractRateCardListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ContractRateCardArchiveParams struct {
	ID shared.IDParam
	paramObj
}

func (r V1ContractRateCardArchiveParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ID)
}
func (r *V1ContractRateCardArchiveParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ID)
}

type V1ContractRateCardGetRateScheduleParams struct {
	// ID of the rate card to get the schedule for
	RateCardID string `json:"rate_card_id,required" format:"uuid"`
	// inclusive starting point for the rates schedule
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// optional exclusive end date for the rates schedule. When not specified rates
	// will show all future schedule segments.
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// List of rate selectors, rates matching ANY of the selector will be included in
	// the response Passing no selectors will result in all rates being returned.
	Selectors []V1ContractRateCardGetRateScheduleParamsSelector `json:"selectors,omitzero"`
	paramObj
}

func (r V1ContractRateCardGetRateScheduleParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardGetRateScheduleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardGetRateScheduleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1ContractRateCardGetRateScheduleParams]'s query parameters
// as `url.Values`.
func (r V1ContractRateCardGetRateScheduleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ContractRateCardGetRateScheduleParamsSelector struct {
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
	paramObj
}

func (r V1ContractRateCardGetRateScheduleParamsSelector) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractRateCardGetRateScheduleParamsSelector
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractRateCardGetRateScheduleParamsSelector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractRateCardGetRateScheduleParamsSelector](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}
