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
func NewV1PlanService(opts ...option.RequestOption) (r V1PlanService) {
	r = V1PlanService{}
	r.Options = opts
	return
}

// List all available plans.
func (r *V1PlanService) List(ctx context.Context, query V1PlanListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1PlanListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	if query.PlanID == "" {
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.PlanID == "" {
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.PlanID == "" {
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
	ID string `json:"id,required" format:"uuid"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string       `json:"custom_fields,required"`
	Name         string                  `json:"name,required"`
	CreditGrants []PlanDetailCreditGrant `json:"credit_grants"`
	Description  string                  `json:"description"`
	Minimums     []PlanDetailMinimum     `json:"minimums"`
	OverageRates []PlanDetailOverageRate `json:"overage_rates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CustomFields respjson.Field
		Name         respjson.Field
		CreditGrants respjson.Field
		Description  respjson.Field
		Minimums     respjson.Field
		OverageRates respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanDetail) RawJSON() string { return r.JSON.raw }
func (r *PlanDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlanDetailCreditGrant struct {
	AmountGranted           float64               `json:"amount_granted,required"`
	AmountGrantedCreditType shared.CreditTypeData `json:"amount_granted_credit_type,required"`
	AmountPaid              float64               `json:"amount_paid,required"`
	AmountPaidCreditType    shared.CreditTypeData `json:"amount_paid_credit_type,required"`
	EffectiveDuration       float64               `json:"effective_duration,required"`
	Name                    string                `json:"name,required"`
	Priority                string                `json:"priority,required"`
	SendInvoice             bool                  `json:"send_invoice,required"`
	Reason                  string                `json:"reason"`
	RecurrenceDuration      float64               `json:"recurrence_duration"`
	RecurrenceInterval      float64               `json:"recurrence_interval"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountGranted           respjson.Field
		AmountGrantedCreditType respjson.Field
		AmountPaid              respjson.Field
		AmountPaidCreditType    respjson.Field
		EffectiveDuration       respjson.Field
		Name                    respjson.Field
		Priority                respjson.Field
		SendInvoice             respjson.Field
		Reason                  respjson.Field
		RecurrenceDuration      respjson.Field
		RecurrenceInterval      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanDetailCreditGrant) RawJSON() string { return r.JSON.raw }
func (r *PlanDetailCreditGrant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlanDetailMinimum struct {
	CreditType shared.CreditTypeData `json:"credit_type,required"`
	Name       string                `json:"name,required"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod float64 `json:"start_period,required"`
	Value       float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType  respjson.Field
		Name        respjson.Field
		StartPeriod respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanDetailMinimum) RawJSON() string { return r.JSON.raw }
func (r *PlanDetailMinimum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlanDetailOverageRate struct {
	CreditType     shared.CreditTypeData `json:"credit_type,required"`
	FiatCreditType shared.CreditTypeData `json:"fiat_credit_type,required"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod            float64 `json:"start_period,required"`
	ToFiatConversionFactor float64 `json:"to_fiat_conversion_factor,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType             respjson.Field
		FiatCreditType         respjson.Field
		StartPeriod            respjson.Field
		ToFiatConversionFactor respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanDetailOverageRate) RawJSON() string { return r.JSON.raw }
func (r *PlanDetailOverageRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanListResponse struct {
	ID          string `json:"id,required" format:"uuid"`
	Description string `json:"description,required"`
	Name        string `json:"name,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Description  respjson.Field
		Name         respjson.Field
		CustomFields respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanGetDetailsResponse struct {
	Data PlanDetail `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanGetDetailsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PlanGetDetailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanListChargesResponse struct {
	ID string `json:"id,required" format:"uuid"`
	// Any of "usage", "fixed", "composite", "minimum", "seat".
	ChargeType V1PlanListChargesResponseChargeType `json:"charge_type,required"`
	CreditType shared.CreditTypeData               `json:"credit_type,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string                `json:"custom_fields,required"`
	Name         string                           `json:"name,required"`
	Prices       []V1PlanListChargesResponsePrice `json:"prices,required"`
	ProductID    string                           `json:"product_id,required"`
	ProductName  string                           `json:"product_name,required"`
	Quantity     float64                          `json:"quantity"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod float64 `json:"start_period"`
	// Used in pricing tiers. Indicates how often the tier resets. Default is 1 - the
	// tier count resets every billing period.
	TierResetFrequency float64 `json:"tier_reset_frequency"`
	// Specifies how quantities for usage based charges will be converted.
	UnitConversion V1PlanListChargesResponseUnitConversion `json:"unit_conversion"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		ChargeType         respjson.Field
		CreditType         respjson.Field
		CustomFields       respjson.Field
		Name               respjson.Field
		Prices             respjson.Field
		ProductID          respjson.Field
		ProductName        respjson.Field
		Quantity           respjson.Field
		StartPeriod        respjson.Field
		TierResetFrequency respjson.Field
		UnitConversion     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListChargesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListChargesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanListChargesResponseChargeType string

const (
	V1PlanListChargesResponseChargeTypeUsage     V1PlanListChargesResponseChargeType = "usage"
	V1PlanListChargesResponseChargeTypeFixed     V1PlanListChargesResponseChargeType = "fixed"
	V1PlanListChargesResponseChargeTypeComposite V1PlanListChargesResponseChargeType = "composite"
	V1PlanListChargesResponseChargeTypeMinimum   V1PlanListChargesResponseChargeType = "minimum"
	V1PlanListChargesResponseChargeTypeSeat      V1PlanListChargesResponseChargeType = "seat"
)

type V1PlanListChargesResponsePrice struct {
	// Used in pricing tiers. Indicates at what metric value the price applies.
	Tier               float64 `json:"tier,required"`
	Value              float64 `json:"value,required"`
	CollectionInterval float64 `json:"collection_interval"`
	CollectionSchedule string  `json:"collection_schedule"`
	Quantity           float64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tier               respjson.Field
		Value              respjson.Field
		CollectionInterval respjson.Field
		CollectionSchedule respjson.Field
		Quantity           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListChargesResponsePrice) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListChargesResponsePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies how quantities for usage based charges will be converted.
type V1PlanListChargesResponseUnitConversion struct {
	// The conversion factor
	DivisionFactor float64 `json:"division_factor,required"`
	// Whether usage should be rounded down or up to the nearest whole number. If null,
	// quantity will be rounded to 20 decimal places.
	//
	// Any of "floor", "ceiling".
	RoundingBehavior string `json:"rounding_behavior"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DivisionFactor   respjson.Field
		RoundingBehavior respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListChargesResponseUnitConversion) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListChargesResponseUnitConversion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanListCustomersResponse struct {
	CustomerDetails CustomerDetail                         `json:"customer_details,required"`
	PlanDetails     V1PlanListCustomersResponsePlanDetails `json:"plan_details,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerDetails respjson.Field
		PlanDetails     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListCustomersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListCustomersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanListCustomersResponsePlanDetails struct {
	ID string `json:"id,required" format:"uuid"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields   map[string]string `json:"custom_fields,required"`
	CustomerPlanID string            `json:"customer_plan_id,required" format:"uuid"`
	Name           string            `json:"name,required"`
	// The start date of the plan
	StartingOn time.Time `json:"starting_on,required" format:"date-time"`
	// The end date of the plan
	EndingBefore time.Time `json:"ending_before,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CustomFields   respjson.Field
		CustomerPlanID respjson.Field
		Name           respjson.Field
		StartingOn     respjson.Field
		EndingBefore   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListCustomersResponsePlanDetails) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListCustomersResponsePlanDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanListParams struct {
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1PlanListParams]'s query parameters as `url.Values`.
func (r V1PlanListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PlanGetDetailsParams struct {
	PlanID string `path:"plan_id,required" format:"uuid" json:"-"`
	paramObj
}

type V1PlanListChargesParams struct {
	PlanID string `path:"plan_id,required" format:"uuid" json:"-"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1PlanListChargesParams]'s query parameters as
// `url.Values`.
func (r V1PlanListChargesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PlanListCustomersParams struct {
	PlanID string `path:"plan_id,required" format:"uuid" json:"-"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// Status of customers on a given plan. Defaults to `active`.
	//
	// - `all` - Return current, past, and upcoming customers of the plan.
	// - `active` - Return current customers of the plan.
	// - `ended` - Return past customers of the plan.
	// - `upcoming` - Return upcoming customers of the plan.
	//
	// Multiple statuses can be OR'd together using commas, e.g. `active,ended`.
	// **Note:** `ended,upcoming` combination is not yet supported.
	//
	// Any of "all", "active", "ended", "upcoming".
	Status V1PlanListCustomersParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1PlanListCustomersParams]'s query parameters as
// `url.Values`.
func (r V1PlanListCustomersParams) URLQuery() (v url.Values, err error) {
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
