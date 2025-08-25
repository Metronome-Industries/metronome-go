// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"errors"
	"fmt"
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
func NewV1CustomerPlanService(opts ...option.RequestOption) (r *V1CustomerPlanService) {
	r = &V1CustomerPlanService{}
	r.Options = opts
	return
}

// List the given customer's plans in reverse-chronological order.
func (r *V1CustomerPlanService) List(ctx context.Context, params V1CustomerPlanListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerPlanListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID.Value == "" {
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

// List the given customer's plans in reverse-chronological order.
func (r *V1CustomerPlanService) ListAutoPaging(ctx context.Context, params V1CustomerPlanListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CustomerPlanListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Associate an existing customer with a plan for a specified date range. See the
// [price adjustments documentation](https://plans-docs.metronome.com/pricing/managing-plans/#price-adjustments)
// for details on the price adjustments.
func (r *V1CustomerPlanService) Add(ctx context.Context, params V1CustomerPlanAddParams, opts ...option.RequestOption) (res *V1CustomerPlanAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/plans/add", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Change the end date of a customer's plan.
func (r *V1CustomerPlanService) End(ctx context.Context, params V1CustomerPlanEndParams, opts ...option.RequestOption) (res *V1CustomerPlanEndResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if params.CustomerPlanID.Value == "" {
		err = errors.New("missing required customer_plan_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/plans/%s/end", params.CustomerID, params.CustomerPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Lists a customer plans adjustments. See the
// [price adjustments documentation](https://plans-docs.metronome.com/pricing/managing-plans/#price-adjustments)
// for details.
func (r *V1CustomerPlanService) ListPriceAdjustments(ctx context.Context, params V1CustomerPlanListPriceAdjustmentsParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerPlanListPriceAdjustmentsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if params.CustomerPlanID.Value == "" {
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
// for details.
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
	JSON                v1CustomerPlanListResponseJSON      `json:"-"`
}

// v1CustomerPlanListResponseJSON contains the JSON metadata for the struct
// [V1CustomerPlanListResponse]
type v1CustomerPlanListResponseJSON struct {
	ID                  apijson.Field
	CustomFields        apijson.Field
	PlanDescription     apijson.Field
	PlanID              apijson.Field
	PlanName            apijson.Field
	StartingOn          apijson.Field
	EndingBefore        apijson.Field
	NetPaymentTermsDays apijson.Field
	TrialInfo           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerPlanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPlanListResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPlanListResponseTrialInfo struct {
	EndingBefore time.Time                                        `json:"ending_before,required" format:"date-time"`
	SpendingCaps []V1CustomerPlanListResponseTrialInfoSpendingCap `json:"spending_caps,required"`
	JSON         v1CustomerPlanListResponseTrialInfoJSON          `json:"-"`
}

// v1CustomerPlanListResponseTrialInfoJSON contains the JSON metadata for the
// struct [V1CustomerPlanListResponseTrialInfo]
type v1CustomerPlanListResponseTrialInfoJSON struct {
	EndingBefore apijson.Field
	SpendingCaps apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerPlanListResponseTrialInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPlanListResponseTrialInfoJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPlanListResponseTrialInfoSpendingCap struct {
	Amount          float64                                            `json:"amount,required"`
	AmountRemaining float64                                            `json:"amount_remaining,required"`
	CreditType      shared.CreditTypeData                              `json:"credit_type,required"`
	JSON            v1CustomerPlanListResponseTrialInfoSpendingCapJSON `json:"-"`
}

// v1CustomerPlanListResponseTrialInfoSpendingCapJSON contains the JSON metadata
// for the struct [V1CustomerPlanListResponseTrialInfoSpendingCap]
type v1CustomerPlanListResponseTrialInfoSpendingCapJSON struct {
	Amount          apijson.Field
	AmountRemaining apijson.Field
	CreditType      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V1CustomerPlanListResponseTrialInfoSpendingCap) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPlanListResponseTrialInfoSpendingCapJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPlanAddResponse struct {
	Data shared.ID                     `json:"data,required"`
	JSON v1CustomerPlanAddResponseJSON `json:"-"`
}

// v1CustomerPlanAddResponseJSON contains the JSON metadata for the struct
// [V1CustomerPlanAddResponse]
type v1CustomerPlanAddResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPlanAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPlanAddResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPlanEndResponse struct {
	JSON v1CustomerPlanEndResponseJSON `json:"-"`
}

// v1CustomerPlanEndResponseJSON contains the JSON metadata for the struct
// [V1CustomerPlanEndResponse]
type v1CustomerPlanEndResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPlanEndResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPlanEndResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPlanListPriceAdjustmentsResponse struct {
	ChargeID    string                                               `json:"charge_id,required" format:"uuid"`
	ChargeType  V1CustomerPlanListPriceAdjustmentsResponseChargeType `json:"charge_type,required"`
	Prices      []V1CustomerPlanListPriceAdjustmentsResponsePrice    `json:"prices,required"`
	StartPeriod float64                                              `json:"start_period,required"`
	Quantity    float64                                              `json:"quantity"`
	JSON        v1CustomerPlanListPriceAdjustmentsResponseJSON       `json:"-"`
}

// v1CustomerPlanListPriceAdjustmentsResponseJSON contains the JSON metadata for
// the struct [V1CustomerPlanListPriceAdjustmentsResponse]
type v1CustomerPlanListPriceAdjustmentsResponseJSON struct {
	ChargeID    apijson.Field
	ChargeType  apijson.Field
	Prices      apijson.Field
	StartPeriod apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPlanListPriceAdjustmentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPlanListPriceAdjustmentsResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPlanListPriceAdjustmentsResponseChargeType string

const (
	V1CustomerPlanListPriceAdjustmentsResponseChargeTypeUsage     V1CustomerPlanListPriceAdjustmentsResponseChargeType = "usage"
	V1CustomerPlanListPriceAdjustmentsResponseChargeTypeFixed     V1CustomerPlanListPriceAdjustmentsResponseChargeType = "fixed"
	V1CustomerPlanListPriceAdjustmentsResponseChargeTypeComposite V1CustomerPlanListPriceAdjustmentsResponseChargeType = "composite"
	V1CustomerPlanListPriceAdjustmentsResponseChargeTypeMinimum   V1CustomerPlanListPriceAdjustmentsResponseChargeType = "minimum"
	V1CustomerPlanListPriceAdjustmentsResponseChargeTypeSeat      V1CustomerPlanListPriceAdjustmentsResponseChargeType = "seat"
)

func (r V1CustomerPlanListPriceAdjustmentsResponseChargeType) IsKnown() bool {
	switch r {
	case V1CustomerPlanListPriceAdjustmentsResponseChargeTypeUsage, V1CustomerPlanListPriceAdjustmentsResponseChargeTypeFixed, V1CustomerPlanListPriceAdjustmentsResponseChargeTypeComposite, V1CustomerPlanListPriceAdjustmentsResponseChargeTypeMinimum, V1CustomerPlanListPriceAdjustmentsResponseChargeTypeSeat:
		return true
	}
	return false
}

type V1CustomerPlanListPriceAdjustmentsResponsePrice struct {
	// Determines how the value will be applied.
	AdjustmentType V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType `json:"adjustment_type,required"`
	// Used in pricing tiers. Indicates at what metric value the price applies.
	Tier  float64                                             `json:"tier"`
	Value float64                                             `json:"value"`
	JSON  v1CustomerPlanListPriceAdjustmentsResponsePriceJSON `json:"-"`
}

// v1CustomerPlanListPriceAdjustmentsResponsePriceJSON contains the JSON metadata
// for the struct [V1CustomerPlanListPriceAdjustmentsResponsePrice]
type v1CustomerPlanListPriceAdjustmentsResponsePriceJSON struct {
	AdjustmentType apijson.Field
	Tier           apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1CustomerPlanListPriceAdjustmentsResponsePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPlanListPriceAdjustmentsResponsePriceJSON) RawJSON() string {
	return r.raw
}

// Determines how the value will be applied.
type V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType string

const (
	V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypeFixed      V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType = "fixed"
	V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypeQuantity   V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType = "quantity"
	V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypePercentage V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType = "percentage"
	V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypeOverride   V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType = "override"
)

func (r V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType) IsKnown() bool {
	switch r {
	case V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypeFixed, V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypeQuantity, V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypePercentage, V1CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypeOverride:
		return true
	}
	return false
}

type V1CustomerPlanListParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [V1CustomerPlanListParams]'s query parameters as
// `url.Values`.
func (r V1CustomerPlanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerPlanAddParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	PlanID     param.Field[string] `json:"plan_id,required" format:"uuid"`
	// RFC 3339 timestamp for when the plan becomes active for this customer. Must be
	// at 0:00 UTC (midnight).
	StartingOn param.Field[time.Time] `json:"starting_on,required" format:"date-time"`
	// RFC 3339 timestamp for when the plan ends (exclusive) for this customer. Must be
	// at 0:00 UTC (midnight).
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// Number of days after issuance of invoice after which the invoice is due (e.g.
	// Net 30).
	NetPaymentTermsDays param.Field[float64] `json:"net_payment_terms_days"`
	// An optional list of overage rates that override the rates of the original plan
	// configuration. These new rates will apply to all pricing ramps.
	OverageRateAdjustments param.Field[[]V1CustomerPlanAddParamsOverageRateAdjustment] `json:"overage_rate_adjustments"`
	// A list of price adjustments can be applied on top of the pricing in the plans.
	// See the
	// [price adjustments documentation](https://plans-docs.metronome.com/pricing/managing-plans/#price-adjustments)
	// for details.
	PriceAdjustments param.Field[[]V1CustomerPlanAddParamsPriceAdjustment] `json:"price_adjustments"`
	// A custom trial can be set for the customer's plan. See the
	// [trial configuration documentation](https://docs.metronome.com/provisioning/configure-trials/)
	// for details.
	TrialSpec param.Field[V1CustomerPlanAddParamsTrialSpec] `json:"trial_spec"`
}

func (r V1CustomerPlanAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerPlanAddParamsOverageRateAdjustment struct {
	CustomCreditTypeID       param.Field[string] `json:"custom_credit_type_id,required" format:"uuid"`
	FiatCurrencyCreditTypeID param.Field[string] `json:"fiat_currency_credit_type_id,required" format:"uuid"`
	// The overage cost in fiat currency for each credit of the custom credit type.
	ToFiatConversionFactor param.Field[float64] `json:"to_fiat_conversion_factor,required"`
}

func (r V1CustomerPlanAddParamsOverageRateAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerPlanAddParamsPriceAdjustment struct {
	AdjustmentType param.Field[V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentType] `json:"adjustment_type,required"`
	ChargeID       param.Field[string]                                                `json:"charge_id,required" format:"uuid"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod param.Field[float64] `json:"start_period,required"`
	// the overridden quantity for a fixed charge
	Quantity param.Field[float64] `json:"quantity"`
	// Used in pricing tiers. Indicates at what metric value the price applies.
	Tier param.Field[float64] `json:"tier"`
	// The amount of change to a price. Percentage and fixed adjustments can be
	// positive or negative. Percentage-based adjustments should be decimals, e.g.
	// -0.05 for a 5% discount.
	Value param.Field[float64] `json:"value"`
}

func (r V1CustomerPlanAddParamsPriceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentType string

const (
	V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypePercentage V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentType = "percentage"
	V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypeFixed      V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentType = "fixed"
	V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypeOverride   V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentType = "override"
	V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypeQuantity   V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentType = "quantity"
)

func (r V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypePercentage, V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypeFixed, V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypeOverride, V1CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypeQuantity:
		return true
	}
	return false
}

// A custom trial can be set for the customer's plan. See the
// [trial configuration documentation](https://docs.metronome.com/provisioning/configure-trials/)
// for details.
type V1CustomerPlanAddParamsTrialSpec struct {
	// Length of the trial period in days.
	LengthInDays param.Field[float64]                                     `json:"length_in_days,required"`
	SpendingCap  param.Field[V1CustomerPlanAddParamsTrialSpecSpendingCap] `json:"spending_cap"`
}

func (r V1CustomerPlanAddParamsTrialSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerPlanAddParamsTrialSpecSpendingCap struct {
	// The credit amount in the given denomination based on the credit type, e.g. US
	// cents.
	Amount param.Field[float64] `json:"amount,required"`
	// The credit type ID for the spending cap.
	CreditTypeID param.Field[string] `json:"credit_type_id,required"`
}

func (r V1CustomerPlanAddParamsTrialSpecSpendingCap) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerPlanEndParams struct {
	CustomerID     param.Field[string] `path:"customer_id,required" format:"uuid"`
	CustomerPlanID param.Field[string] `path:"customer_plan_id,required" format:"uuid"`
	// RFC 3339 timestamp for when the plan ends (exclusive) for this customer. Must be
	// at 0:00 UTC (midnight). If not provided, the plan end date will be cleared.
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// If true, plan end date can be before the last finalized invoice date. Any
	// invoices generated after the plan end date will be voided.
	VoidInvoices param.Field[bool] `json:"void_invoices"`
	// Only applicable when void_invoices is set to true. If true, for every invoice
	// that is voided we will also attempt to void/delete the stripe invoice (if any).
	// Stripe invoices will be voided if finalized or deleted if still in draft state.
	VoidStripeInvoices param.Field[bool] `json:"void_stripe_invoices"`
}

func (r V1CustomerPlanEndParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerPlanListPriceAdjustmentsParams struct {
	CustomerID     param.Field[string] `path:"customer_id,required" format:"uuid"`
	CustomerPlanID param.Field[string] `path:"customer_plan_id,required" format:"uuid"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [V1CustomerPlanListPriceAdjustmentsParams]'s query
// parameters as `url.Values`.
func (r V1CustomerPlanListPriceAdjustmentsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
