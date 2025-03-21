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

// V1PlanService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PlanService] method instead.
type V1PlanService struct {
	Options []option.RequestOption
}

// NewV1PlanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1PlanService(opts ...option.RequestOption) (r *V1PlanService) {
	r = &V1PlanService{}
	r.Options = opts
	return
}

// List all available plans.
func (r *V1PlanService) List(ctx context.Context, query V1PlanListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1PlanListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/plans"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List all available plans.
func (r *V1PlanService) ListAutoPaging(ctx context.Context, query V1PlanListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1PlanListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Fetch high level details of a specific plan.
func (r *V1PlanService) GetDetails(ctx context.Context, query V1PlanGetDetailsParams, opts ...option.RequestOption) (res *V1PlanGetDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.PlanID.Value == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("v1/planDetails/%s", query.PlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a list of charges of a specific plan.
func (r *V1PlanService) ListCharges(ctx context.Context, params V1PlanListChargesParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1PlanListChargesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.PlanID.Value == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("v1/planDetails/%s/charges", params.PlanID)
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

// Fetches a list of charges of a specific plan.
func (r *V1PlanService) ListChargesAutoPaging(ctx context.Context, params V1PlanListChargesParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1PlanListChargesResponse] {
	return pagination.NewCursorPageAutoPager(r.ListCharges(ctx, params, opts...))
}

// Fetches a list of customers on a specific plan (by default, only currently
// active plans are included)
func (r *V1PlanService) ListCustomers(ctx context.Context, params V1PlanListCustomersParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1PlanListCustomersResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.PlanID.Value == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("v1/planDetails/%s/customers", params.PlanID)
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

// Fetches a list of customers on a specific plan (by default, only currently
// active plans are included)
func (r *V1PlanService) ListCustomersAutoPaging(ctx context.Context, params V1PlanListCustomersParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1PlanListCustomersResponse] {
	return pagination.NewCursorPageAutoPager(r.ListCustomers(ctx, params, opts...))
}

type PlanDetail struct {
	ID           string                  `json:"id,required" format:"uuid"`
	CustomFields map[string]string       `json:"custom_fields,required"`
	Name         string                  `json:"name,required"`
	CreditGrants []PlanDetailCreditGrant `json:"credit_grants"`
	Description  string                  `json:"description"`
	Minimums     []PlanDetailMinimum     `json:"minimums"`
	OverageRates []PlanDetailOverageRate `json:"overage_rates"`
	JSON         planDetailJSON          `json:"-"`
}

// planDetailJSON contains the JSON metadata for the struct [PlanDetail]
type planDetailJSON struct {
	ID           apijson.Field
	CustomFields apijson.Field
	Name         apijson.Field
	CreditGrants apijson.Field
	Description  apijson.Field
	Minimums     apijson.Field
	OverageRates apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PlanDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planDetailJSON) RawJSON() string {
	return r.raw
}

type PlanDetailCreditGrant struct {
	AmountGranted           float64                   `json:"amount_granted,required"`
	AmountGrantedCreditType shared.CreditTypeData     `json:"amount_granted_credit_type,required"`
	AmountPaid              float64                   `json:"amount_paid,required"`
	AmountPaidCreditType    shared.CreditTypeData     `json:"amount_paid_credit_type,required"`
	EffectiveDuration       float64                   `json:"effective_duration,required"`
	Name                    string                    `json:"name,required"`
	Priority                string                    `json:"priority,required"`
	SendInvoice             bool                      `json:"send_invoice,required"`
	Reason                  string                    `json:"reason"`
	RecurrenceDuration      float64                   `json:"recurrence_duration"`
	RecurrenceInterval      float64                   `json:"recurrence_interval"`
	JSON                    planDetailCreditGrantJSON `json:"-"`
}

// planDetailCreditGrantJSON contains the JSON metadata for the struct
// [PlanDetailCreditGrant]
type planDetailCreditGrantJSON struct {
	AmountGranted           apijson.Field
	AmountGrantedCreditType apijson.Field
	AmountPaid              apijson.Field
	AmountPaidCreditType    apijson.Field
	EffectiveDuration       apijson.Field
	Name                    apijson.Field
	Priority                apijson.Field
	SendInvoice             apijson.Field
	Reason                  apijson.Field
	RecurrenceDuration      apijson.Field
	RecurrenceInterval      apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PlanDetailCreditGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planDetailCreditGrantJSON) RawJSON() string {
	return r.raw
}

type PlanDetailMinimum struct {
	CreditType shared.CreditTypeData `json:"credit_type,required"`
	Name       string                `json:"name,required"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod float64               `json:"start_period,required"`
	Value       float64               `json:"value,required"`
	JSON        planDetailMinimumJSON `json:"-"`
}

// planDetailMinimumJSON contains the JSON metadata for the struct
// [PlanDetailMinimum]
type planDetailMinimumJSON struct {
	CreditType  apijson.Field
	Name        apijson.Field
	StartPeriod apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanDetailMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planDetailMinimumJSON) RawJSON() string {
	return r.raw
}

type PlanDetailOverageRate struct {
	CreditType     shared.CreditTypeData `json:"credit_type,required"`
	FiatCreditType shared.CreditTypeData `json:"fiat_credit_type,required"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod            float64                   `json:"start_period,required"`
	ToFiatConversionFactor float64                   `json:"to_fiat_conversion_factor,required"`
	JSON                   planDetailOverageRateJSON `json:"-"`
}

// planDetailOverageRateJSON contains the JSON metadata for the struct
// [PlanDetailOverageRate]
type planDetailOverageRateJSON struct {
	CreditType             apijson.Field
	FiatCreditType         apijson.Field
	StartPeriod            apijson.Field
	ToFiatConversionFactor apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PlanDetailOverageRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planDetailOverageRateJSON) RawJSON() string {
	return r.raw
}

type V1PlanListResponse struct {
	ID           string                 `json:"id,required" format:"uuid"`
	Description  string                 `json:"description,required"`
	Name         string                 `json:"name,required"`
	CustomFields map[string]string      `json:"custom_fields"`
	JSON         v1PlanListResponseJSON `json:"-"`
}

// v1PlanListResponseJSON contains the JSON metadata for the struct
// [V1PlanListResponse]
type v1PlanListResponseJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	CustomFields apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1PlanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanListResponseJSON) RawJSON() string {
	return r.raw
}

type V1PlanGetDetailsResponse struct {
	Data PlanDetail                   `json:"data,required"`
	JSON v1PlanGetDetailsResponseJSON `json:"-"`
}

// v1PlanGetDetailsResponseJSON contains the JSON metadata for the struct
// [V1PlanGetDetailsResponse]
type v1PlanGetDetailsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1PlanGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type V1PlanListChargesResponse struct {
	ID           string                              `json:"id,required" format:"uuid"`
	ChargeType   V1PlanListChargesResponseChargeType `json:"charge_type,required"`
	CreditType   shared.CreditTypeData               `json:"credit_type,required"`
	CustomFields map[string]string                   `json:"custom_fields,required"`
	Name         string                              `json:"name,required"`
	Prices       []V1PlanListChargesResponsePrice    `json:"prices,required"`
	ProductID    string                              `json:"product_id,required"`
	ProductName  string                              `json:"product_name,required"`
	Quantity     float64                             `json:"quantity"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod float64 `json:"start_period"`
	// Used in pricing tiers. Indicates how often the tier resets. Default is 1 - the
	// tier count resets every billing period.
	TierResetFrequency float64 `json:"tier_reset_frequency"`
	// Specifies how quantities for usage based charges will be converted.
	UnitConversion V1PlanListChargesResponseUnitConversion `json:"unit_conversion"`
	JSON           v1PlanListChargesResponseJSON           `json:"-"`
}

// v1PlanListChargesResponseJSON contains the JSON metadata for the struct
// [V1PlanListChargesResponse]
type v1PlanListChargesResponseJSON struct {
	ID                 apijson.Field
	ChargeType         apijson.Field
	CreditType         apijson.Field
	CustomFields       apijson.Field
	Name               apijson.Field
	Prices             apijson.Field
	ProductID          apijson.Field
	ProductName        apijson.Field
	Quantity           apijson.Field
	StartPeriod        apijson.Field
	TierResetFrequency apijson.Field
	UnitConversion     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V1PlanListChargesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanListChargesResponseJSON) RawJSON() string {
	return r.raw
}

type V1PlanListChargesResponseChargeType string

const (
	V1PlanListChargesResponseChargeTypeUsage     V1PlanListChargesResponseChargeType = "usage"
	V1PlanListChargesResponseChargeTypeFixed     V1PlanListChargesResponseChargeType = "fixed"
	V1PlanListChargesResponseChargeTypeComposite V1PlanListChargesResponseChargeType = "composite"
	V1PlanListChargesResponseChargeTypeMinimum   V1PlanListChargesResponseChargeType = "minimum"
	V1PlanListChargesResponseChargeTypeSeat      V1PlanListChargesResponseChargeType = "seat"
)

func (r V1PlanListChargesResponseChargeType) IsKnown() bool {
	switch r {
	case V1PlanListChargesResponseChargeTypeUsage, V1PlanListChargesResponseChargeTypeFixed, V1PlanListChargesResponseChargeTypeComposite, V1PlanListChargesResponseChargeTypeMinimum, V1PlanListChargesResponseChargeTypeSeat:
		return true
	}
	return false
}

type V1PlanListChargesResponsePrice struct {
	// Used in pricing tiers. Indicates at what metric value the price applies.
	Tier               float64                            `json:"tier,required"`
	Value              float64                            `json:"value,required"`
	CollectionInterval float64                            `json:"collection_interval"`
	CollectionSchedule string                             `json:"collection_schedule"`
	Quantity           float64                            `json:"quantity"`
	JSON               v1PlanListChargesResponsePriceJSON `json:"-"`
}

// v1PlanListChargesResponsePriceJSON contains the JSON metadata for the struct
// [V1PlanListChargesResponsePrice]
type v1PlanListChargesResponsePriceJSON struct {
	Tier               apijson.Field
	Value              apijson.Field
	CollectionInterval apijson.Field
	CollectionSchedule apijson.Field
	Quantity           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V1PlanListChargesResponsePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanListChargesResponsePriceJSON) RawJSON() string {
	return r.raw
}

// Specifies how quantities for usage based charges will be converted.
type V1PlanListChargesResponseUnitConversion struct {
	// The conversion factor
	DivisionFactor float64 `json:"division_factor,required"`
	// Whether usage should be rounded down or up to the nearest whole number. If null,
	// quantity will be rounded to 20 decimal places.
	RoundingBehavior V1PlanListChargesResponseUnitConversionRoundingBehavior `json:"rounding_behavior"`
	JSON             v1PlanListChargesResponseUnitConversionJSON             `json:"-"`
}

// v1PlanListChargesResponseUnitConversionJSON contains the JSON metadata for the
// struct [V1PlanListChargesResponseUnitConversion]
type v1PlanListChargesResponseUnitConversionJSON struct {
	DivisionFactor   apijson.Field
	RoundingBehavior apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1PlanListChargesResponseUnitConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanListChargesResponseUnitConversionJSON) RawJSON() string {
	return r.raw
}

// Whether usage should be rounded down or up to the nearest whole number. If null,
// quantity will be rounded to 20 decimal places.
type V1PlanListChargesResponseUnitConversionRoundingBehavior string

const (
	V1PlanListChargesResponseUnitConversionRoundingBehaviorFloor   V1PlanListChargesResponseUnitConversionRoundingBehavior = "floor"
	V1PlanListChargesResponseUnitConversionRoundingBehaviorCeiling V1PlanListChargesResponseUnitConversionRoundingBehavior = "ceiling"
)

func (r V1PlanListChargesResponseUnitConversionRoundingBehavior) IsKnown() bool {
	switch r {
	case V1PlanListChargesResponseUnitConversionRoundingBehaviorFloor, V1PlanListChargesResponseUnitConversionRoundingBehaviorCeiling:
		return true
	}
	return false
}

type V1PlanListCustomersResponse struct {
	CustomerDetails CustomerDetail                         `json:"customer_details,required"`
	PlanDetails     V1PlanListCustomersResponsePlanDetails `json:"plan_details,required"`
	JSON            v1PlanListCustomersResponseJSON        `json:"-"`
}

// v1PlanListCustomersResponseJSON contains the JSON metadata for the struct
// [V1PlanListCustomersResponse]
type v1PlanListCustomersResponseJSON struct {
	CustomerDetails apijson.Field
	PlanDetails     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V1PlanListCustomersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanListCustomersResponseJSON) RawJSON() string {
	return r.raw
}

type V1PlanListCustomersResponsePlanDetails struct {
	ID             string            `json:"id,required" format:"uuid"`
	CustomFields   map[string]string `json:"custom_fields,required"`
	CustomerPlanID string            `json:"customer_plan_id,required" format:"uuid"`
	Name           string            `json:"name,required"`
	// The start date of the plan
	StartingOn time.Time `json:"starting_on,required" format:"date-time"`
	// The end date of the plan
	EndingBefore time.Time                                  `json:"ending_before,nullable" format:"date-time"`
	JSON         v1PlanListCustomersResponsePlanDetailsJSON `json:"-"`
}

// v1PlanListCustomersResponsePlanDetailsJSON contains the JSON metadata for the
// struct [V1PlanListCustomersResponsePlanDetails]
type v1PlanListCustomersResponsePlanDetailsJSON struct {
	ID             apijson.Field
	CustomFields   apijson.Field
	CustomerPlanID apijson.Field
	Name           apijson.Field
	StartingOn     apijson.Field
	EndingBefore   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1PlanListCustomersResponsePlanDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanListCustomersResponsePlanDetailsJSON) RawJSON() string {
	return r.raw
}

type V1PlanListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [V1PlanListParams]'s query parameters as `url.Values`.
func (r V1PlanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PlanGetDetailsParams struct {
	PlanID param.Field[string] `path:"plan_id,required" format:"uuid"`
}

type V1PlanListChargesParams struct {
	PlanID param.Field[string] `path:"plan_id,required" format:"uuid"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [V1PlanListChargesParams]'s query parameters as
// `url.Values`.
func (r V1PlanListChargesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PlanListCustomersParams struct {
	PlanID param.Field[string] `path:"plan_id,required" format:"uuid"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Status of customers on a given plan. Defaults to `active`.
	//
	// - `all` - Return current, past, and upcoming customers of the plan.
	// - `active` - Return current customers of the plan.
	// - `ended` - Return past customers of the plan.
	// - `upcoming` - Return upcoming customers of the plan.
	//
	// Multiple statuses can be OR'd together using commas, e.g. `active,ended`.
	// **Note:** `ended,upcoming` combination is not yet supported.
	Status param.Field[V1PlanListCustomersParamsStatus] `query:"status"`
}

// URLQuery serializes [V1PlanListCustomersParams]'s query parameters as
// `url.Values`.
func (r V1PlanListCustomersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Status of customers on a given plan. Defaults to `active`.
//
// - `all` - Return current, past, and upcoming customers of the plan.
// - `active` - Return current customers of the plan.
// - `ended` - Return past customers of the plan.
// - `upcoming` - Return upcoming customers of the plan.
//
// Multiple statuses can be OR'd together using commas, e.g. `active,ended`.
// **Note:** `ended,upcoming` combination is not yet supported.
type V1PlanListCustomersParamsStatus string

const (
	V1PlanListCustomersParamsStatusAll      V1PlanListCustomersParamsStatus = "all"
	V1PlanListCustomersParamsStatusActive   V1PlanListCustomersParamsStatus = "active"
	V1PlanListCustomersParamsStatusEnded    V1PlanListCustomersParamsStatus = "ended"
	V1PlanListCustomersParamsStatusUpcoming V1PlanListCustomersParamsStatus = "upcoming"
)

func (r V1PlanListCustomersParamsStatus) IsKnown() bool {
	switch r {
	case V1PlanListCustomersParamsStatusAll, V1PlanListCustomersParamsStatusActive, V1PlanListCustomersParamsStatusEnded, V1PlanListCustomersParamsStatusUpcoming:
		return true
	}
	return false
}
