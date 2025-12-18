// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"errors"
	"fmt"
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

// V1CustomerPlanService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerPlanService] method instead.
type V1CustomerPlanService struct {
	Options []option.RequestOption
}

// NewV1CustomerPlanService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomerPlanService(opts ...option.RequestOption) (r V1CustomerPlanService) {
	r = V1CustomerPlanService{}
	r.Options = opts
	return
}

// List the given customer's plans in reverse-chronological order. This is a Plans
// (deprecated) endpoint. New clients should implement using Contracts.
func (r *V1CustomerPlanService) List(ctx context.Context, params V1CustomerPlanListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerPlanListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/plans", params.CustomerID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List the given customer's plans in reverse-chronological order. This is a Plans
// (deprecated) endpoint. New clients should implement using Contracts.
func (r *V1CustomerPlanService) ListAutoPaging(ctx context.Context, params V1CustomerPlanListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CustomerPlanListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Associate an existing customer with a plan for a specified date range. See the
// [price adjustments documentation](https://plans-docs.metronome.com/pricing/managing-plans/#price-adjustments)
// for details on the price adjustments. This is a Plans (deprecated) endpoint. New
// clients should implement using Contracts.
func (r *V1CustomerPlanService) Add(ctx context.Context, params V1CustomerPlanAddParams, opts ...option.RequestOption) (res *V1CustomerPlanAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/plans/add", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Change the end date of a customer's plan. This is a Plans (deprecated) endpoint.
// New clients should implement using Contracts.
func (r *V1CustomerPlanService) End(ctx context.Context, params V1CustomerPlanEndParams, opts ...option.RequestOption) (res *V1CustomerPlanEndResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if params.CustomerPlanID == "" {
		err = errors.New("missing required customer_plan_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/plans/%s/end", params.CustomerID, params.CustomerPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Lists a customer plans adjustments. See the
// [price adjustments documentation](https://plans-docs.metronome.com/pricing/managing-plans/#price-adjustments)
// for details. This is a Plans (deprecated) endpoint. New clients should implement
// using Contracts.
func (r *V1CustomerPlanService) ListPriceAdjustments(ctx context.Context, params V1CustomerPlanListPriceAdjustmentsParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerPlanListPriceAdjustmentsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if params.CustomerPlanID == "" {
		err = errors.New("missing required customer_plan_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/plans/%s/priceAdjustments", params.CustomerID, params.CustomerPlanID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Lists a customer plans adjustments. See the
// [price adjustments documentation](https://plans-docs.metronome.com/pricing/managing-plans/#price-adjustments)
// for details. This is a Plans (deprecated) endpoint. New clients should implement
// using Contracts.
func (r *V1CustomerPlanService) ListPriceAdjustmentsAutoPaging(ctx context.Context, params V1CustomerPlanListPriceAdjustmentsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CustomerPlanListPriceAdjustmentsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListPriceAdjustments(ctx, params, opts...))
}

type V1CustomerPlanListResponse struct {
	// the ID of the customer plan
	ID string `json:"id,required" format:"uuid"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields    map[string]string `json:"custom_fields,required"`
	PlanDescription string            `json:"plan_description,required"`
	// the ID of the plan
	PlanID              string                              `json:"plan_id,required" format:"uuid"`
	PlanName            string                              `json:"plan_name,required"`
	StartingOn          time.Time                           `json:"starting_on,required" format:"date-time"`
	EndingBefore        time.Time                           `json:"ending_before" format:"date-time"`
	NetPaymentTermsDays float64                             `json:"net_payment_terms_days"`
	TrialInfo           V1CustomerPlanListResponseTrialInfo `json:"trial_info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CustomFields        respjson.Field
		PlanDescription     respjson.Field
		PlanID              respjson.Field
		PlanName            respjson.Field
		StartingOn          respjson.Field
		EndingBefore        respjson.Field
		NetPaymentTermsDays respjson.Field
		TrialInfo           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPlanListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPlanListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPlanListResponseTrialInfo struct {
	EndingBefore time.Time                                        `json:"ending_before,required" format:"date-time"`
	SpendingCaps []V1CustomerPlanListResponseTrialInfoSpendingCap `json:"spending_caps,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndingBefore respjson.Field
		SpendingCaps respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPlanListResponseTrialInfo) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPlanListResponseTrialInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPlanListResponseTrialInfoSpendingCap struct {
	Amount          float64               `json:"amount,required"`
	AmountRemaining float64               `json:"amount_remaining,required"`
	CreditType      shared.CreditTypeData `json:"credit_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount          respjson.Field
		AmountRemaining respjson.Field
		CreditType      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPlanListResponseTrialInfoSpendingCap) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPlanListResponseTrialInfoSpendingCap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPlanAddResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPlanAddResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPlanAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPlanEndResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPlanEndResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPlanEndResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPlanListPriceAdjustmentsResponse struct {
	ChargeID string `json:"charge_id,required" format:"uuid"`
	// Any of "usage", "fixed", "composite", "minimum", "seat".
	ChargeType  V1CustomerPlanListPriceAdjustmentsResponseChargeType `json:"charge_type,required"`
	Prices      []V1CustomerPlanListPriceAdjustmentsResponsePrice    `json:"prices,required"`
	StartPeriod float64                                              `json:"start_period,required"`
	Quantity    float64                                              `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChargeID    respjson.Field
		ChargeType  respjson.Field
		Prices      respjson.Field
		StartPeriod respjson.Field
		Quantity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPlanListPriceAdjustmentsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPlanListPriceAdjustmentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPlanListPriceAdjustmentsResponseChargeType string

const (
	V1CustomerPlanListPriceAdjustmentsResponseChargeTypeUsage     V1CustomerPlanListPriceAdjustmentsResponseChargeType = "usage"
	V1CustomerPlanListPriceAdjustmentsResponseChargeTypeFixed     V1CustomerPlanListPriceAdjustmentsResponseChargeType = "fixed"
	V1CustomerPlanListPriceAdjustmentsResponseChargeTypeComposite V1CustomerPlanListPriceAdjustmentsResponseChargeType = "composite"
	V1CustomerPlanListPriceAdjustmentsResponseChargeTypeMinimum   V1CustomerPlanListPriceAdjustmentsResponseChargeType = "minimum"
	V1CustomerPlanListPriceAdjustmentsResponseChargeTypeSeat      V1CustomerPlanListPriceAdjustmentsResponseChargeType = "seat"
)

type V1CustomerPlanListPriceAdjustmentsResponsePrice struct {
	// Determines how the value will be applied.
	//
	// Any of "fixed", "quantity", "percentage", "override".
	AdjustmentType string  `json:"adjustment_type,required"`
	Quantity       float64 `json:"quantity"`
	// Used in pricing tiers. Indicates at what metric value the price applies.
	Tier  float64 `json:"tier"`
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdjustmentType respjson.Field
		Quantity       respjson.Field
		Tier           respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPlanListPriceAdjustmentsResponsePrice) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPlanListPriceAdjustmentsResponsePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPlanListParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerPlanListParams]'s query parameters as
// `url.Values`.
func (r V1CustomerPlanListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerPlanAddParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	PlanID     string `json:"plan_id,required" format:"uuid"`
	// RFC 3339 timestamp for when the plan becomes active for this customer. Must be
	// at 0:00 UTC (midnight).
	StartingOn time.Time `json:"starting_on,required" format:"date-time"`
	// RFC 3339 timestamp for when the plan ends (exclusive) for this customer. Must be
	// at 0:00 UTC (midnight).
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// Number of days after issuance of invoice after which the invoice is due (e.g.
	// Net 30).
	NetPaymentTermsDays param.Opt[float64] `json:"net_payment_terms_days,omitzero"`
	// An optional list of overage rates that override the rates of the original plan
	// configuration. These new rates will apply to all pricing ramps.
	OverageRateAdjustments []V1CustomerPlanAddParamsOverageRateAdjustment `json:"overage_rate_adjustments,omitzero"`
	// A list of price adjustments can be applied on top of the pricing in the plans.
	// See the
	// [price adjustments documentation](https://plans-docs.metronome.com/pricing/managing-plans/#price-adjustments)
	// for details.
	PriceAdjustments []V1CustomerPlanAddParamsPriceAdjustment `json:"price_adjustments,omitzero"`
	// A custom trial can be set for the customer's plan. See the
	// [trial configuration documentation](https://docs.metronome.com/provisioning/configure-trials/)
	// for details.
	TrialSpec V1CustomerPlanAddParamsTrialSpec `json:"trial_spec,omitzero"`
	paramObj
}

func (r V1CustomerPlanAddParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPlanAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPlanAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CustomCreditTypeID, FiatCurrencyCreditTypeID,
// ToFiatConversionFactor are required.
type V1CustomerPlanAddParamsOverageRateAdjustment struct {
	CustomCreditTypeID       string `json:"custom_credit_type_id,required" format:"uuid"`
	FiatCurrencyCreditTypeID string `json:"fiat_currency_credit_type_id,required" format:"uuid"`
	// The overage cost in fiat currency for each credit of the custom credit type.
	ToFiatConversionFactor float64 `json:"to_fiat_conversion_factor,required"`
	paramObj
}

func (r V1CustomerPlanAddParamsOverageRateAdjustment) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPlanAddParamsOverageRateAdjustment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPlanAddParamsOverageRateAdjustment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AdjustmentType, ChargeID, StartPeriod are required.
type V1CustomerPlanAddParamsPriceAdjustment struct {
	// Any of "percentage", "fixed", "override", "quantity".
	AdjustmentType string `json:"adjustment_type,omitzero,required"`
	ChargeID       string `json:"charge_id,required" format:"uuid"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod float64 `json:"start_period,required"`
	// the overridden quantity for a fixed charge
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Used in pricing tiers. Indicates at what metric value the price applies.
	Tier param.Opt[float64] `json:"tier,omitzero"`
	// The amount of change to a price. Percentage and fixed adjustments can be
	// positive or negative. Percentage-based adjustments should be decimals, e.g.
	// -0.05 for a 5% discount.
	Value param.Opt[float64] `json:"value,omitzero"`
	paramObj
}

func (r V1CustomerPlanAddParamsPriceAdjustment) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPlanAddParamsPriceAdjustment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPlanAddParamsPriceAdjustment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPlanAddParamsPriceAdjustment](
		"adjustment_type", "percentage", "fixed", "override", "quantity",
	)
}

// A custom trial can be set for the customer's plan. See the
// [trial configuration documentation](https://docs.metronome.com/provisioning/configure-trials/)
// for details.
//
// The property LengthInDays is required.
type V1CustomerPlanAddParamsTrialSpec struct {
	// Length of the trial period in days.
	LengthInDays float64                                     `json:"length_in_days,required"`
	SpendingCap  V1CustomerPlanAddParamsTrialSpecSpendingCap `json:"spending_cap,omitzero"`
	paramObj
}

func (r V1CustomerPlanAddParamsTrialSpec) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPlanAddParamsTrialSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPlanAddParamsTrialSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, CreditTypeID are required.
type V1CustomerPlanAddParamsTrialSpecSpendingCap struct {
	// The credit amount in the given denomination based on the credit type, e.g. US
	// cents.
	Amount float64 `json:"amount,required"`
	// The credit type ID for the spending cap.
	CreditTypeID string `json:"credit_type_id,required"`
	paramObj
}

func (r V1CustomerPlanAddParamsTrialSpecSpendingCap) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPlanAddParamsTrialSpecSpendingCap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPlanAddParamsTrialSpecSpendingCap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPlanEndParams struct {
	CustomerID     string `path:"customer_id,required" format:"uuid" json:"-"`
	CustomerPlanID string `path:"customer_plan_id,required" format:"uuid" json:"-"`
	// RFC 3339 timestamp for when the plan ends (exclusive) for this customer. Must be
	// at 0:00 UTC (midnight). If not provided, the plan end date will be cleared.
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// If true, plan end date can be before the last finalized invoice date. Any
	// invoices generated after the plan end date will be voided.
	VoidInvoices param.Opt[bool] `json:"void_invoices,omitzero"`
	// Only applicable when void_invoices is set to true. If true, for every invoice
	// that is voided we will also attempt to void/delete the stripe invoice (if any).
	// Stripe invoices will be voided if finalized or deleted if still in draft state.
	VoidStripeInvoices param.Opt[bool] `json:"void_stripe_invoices,omitzero"`
	paramObj
}

func (r V1CustomerPlanEndParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPlanEndParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPlanEndParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPlanListPriceAdjustmentsParams struct {
	CustomerID     string `path:"customer_id,required" format:"uuid" json:"-"`
	CustomerPlanID string `path:"customer_plan_id,required" format:"uuid" json:"-"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerPlanListPriceAdjustmentsParams]'s query
// parameters as `url.Values`.
func (r V1CustomerPlanListPriceAdjustmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
