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
	Data V1PlanGetDetailsResponseData `json:"data,required"`
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

type V1PlanGetDetailsResponseData struct {
	ID           string                                    `json:"id,required" format:"uuid"`
	CustomFields map[string]string                         `json:"custom_fields,required"`
	Name         string                                    `json:"name,required"`
	CreditGrants []V1PlanGetDetailsResponseDataCreditGrant `json:"credit_grants"`
	Description  string                                    `json:"description"`
	Minimums     []V1PlanGetDetailsResponseDataMinimum     `json:"minimums"`
	OverageRates []V1PlanGetDetailsResponseDataOverageRate `json:"overage_rates"`
	JSON         v1PlanGetDetailsResponseDataJSON          `json:"-"`
}

// v1PlanGetDetailsResponseDataJSON contains the JSON metadata for the struct
// [V1PlanGetDetailsResponseData]
type v1PlanGetDetailsResponseDataJSON struct {
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

func (r *V1PlanGetDetailsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanGetDetailsResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1PlanGetDetailsResponseDataCreditGrant struct {
	AmountGranted           float64                                                         `json:"amount_granted,required"`
	AmountGrantedCreditType V1PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType `json:"amount_granted_credit_type,required"`
	AmountPaid              float64                                                         `json:"amount_paid,required"`
	AmountPaidCreditType    V1PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType    `json:"amount_paid_credit_type,required"`
	EffectiveDuration       float64                                                         `json:"effective_duration,required"`
	Name                    string                                                          `json:"name,required"`
	Priority                string                                                          `json:"priority,required"`
	SendInvoice             bool                                                            `json:"send_invoice,required"`
	Reason                  string                                                          `json:"reason"`
	RecurrenceDuration      float64                                                         `json:"recurrence_duration"`
	RecurrenceInterval      float64                                                         `json:"recurrence_interval"`
	JSON                    v1PlanGetDetailsResponseDataCreditGrantJSON                     `json:"-"`
}

// v1PlanGetDetailsResponseDataCreditGrantJSON contains the JSON metadata for the
// struct [V1PlanGetDetailsResponseDataCreditGrant]
type v1PlanGetDetailsResponseDataCreditGrantJSON struct {
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

func (r *V1PlanGetDetailsResponseDataCreditGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanGetDetailsResponseDataCreditGrantJSON) RawJSON() string {
	return r.raw
}

type V1PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType struct {
	ID   string                                                              `json:"id,required" format:"uuid"`
	Name string                                                              `json:"name,required"`
	JSON v1PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditTypeJSON `json:"-"`
}

// v1PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditTypeJSON contains the
// JSON metadata for the struct
// [V1PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType]
type v1PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType struct {
	ID   string                                                           `json:"id,required" format:"uuid"`
	Name string                                                           `json:"name,required"`
	JSON v1PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditTypeJSON `json:"-"`
}

// v1PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditTypeJSON contains the
// JSON metadata for the struct
// [V1PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType]
type v1PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1PlanGetDetailsResponseDataMinimum struct {
	CreditType V1PlanGetDetailsResponseDataMinimumsCreditType `json:"credit_type,required"`
	Name       string                                         `json:"name,required"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod float64                                 `json:"start_period,required"`
	Value       float64                                 `json:"value,required"`
	JSON        v1PlanGetDetailsResponseDataMinimumJSON `json:"-"`
}

// v1PlanGetDetailsResponseDataMinimumJSON contains the JSON metadata for the
// struct [V1PlanGetDetailsResponseDataMinimum]
type v1PlanGetDetailsResponseDataMinimumJSON struct {
	CreditType  apijson.Field
	Name        apijson.Field
	StartPeriod apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1PlanGetDetailsResponseDataMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanGetDetailsResponseDataMinimumJSON) RawJSON() string {
	return r.raw
}

type V1PlanGetDetailsResponseDataMinimumsCreditType struct {
	ID   string                                             `json:"id,required" format:"uuid"`
	Name string                                             `json:"name,required"`
	JSON v1PlanGetDetailsResponseDataMinimumsCreditTypeJSON `json:"-"`
}

// v1PlanGetDetailsResponseDataMinimumsCreditTypeJSON contains the JSON metadata
// for the struct [V1PlanGetDetailsResponseDataMinimumsCreditType]
type v1PlanGetDetailsResponseDataMinimumsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1PlanGetDetailsResponseDataMinimumsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanGetDetailsResponseDataMinimumsCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1PlanGetDetailsResponseDataOverageRate struct {
	CreditType     V1PlanGetDetailsResponseDataOverageRatesCreditType     `json:"credit_type,required"`
	FiatCreditType V1PlanGetDetailsResponseDataOverageRatesFiatCreditType `json:"fiat_credit_type,required"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod            float64                                     `json:"start_period,required"`
	ToFiatConversionFactor float64                                     `json:"to_fiat_conversion_factor,required"`
	JSON                   v1PlanGetDetailsResponseDataOverageRateJSON `json:"-"`
}

// v1PlanGetDetailsResponseDataOverageRateJSON contains the JSON metadata for the
// struct [V1PlanGetDetailsResponseDataOverageRate]
type v1PlanGetDetailsResponseDataOverageRateJSON struct {
	CreditType             apijson.Field
	FiatCreditType         apijson.Field
	StartPeriod            apijson.Field
	ToFiatConversionFactor apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1PlanGetDetailsResponseDataOverageRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanGetDetailsResponseDataOverageRateJSON) RawJSON() string {
	return r.raw
}

type V1PlanGetDetailsResponseDataOverageRatesCreditType struct {
	ID   string                                                 `json:"id,required" format:"uuid"`
	Name string                                                 `json:"name,required"`
	JSON v1PlanGetDetailsResponseDataOverageRatesCreditTypeJSON `json:"-"`
}

// v1PlanGetDetailsResponseDataOverageRatesCreditTypeJSON contains the JSON
// metadata for the struct [V1PlanGetDetailsResponseDataOverageRatesCreditType]
type v1PlanGetDetailsResponseDataOverageRatesCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1PlanGetDetailsResponseDataOverageRatesCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanGetDetailsResponseDataOverageRatesCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1PlanGetDetailsResponseDataOverageRatesFiatCreditType struct {
	ID   string                                                     `json:"id,required" format:"uuid"`
	Name string                                                     `json:"name,required"`
	JSON v1PlanGetDetailsResponseDataOverageRatesFiatCreditTypeJSON `json:"-"`
}

// v1PlanGetDetailsResponseDataOverageRatesFiatCreditTypeJSON contains the JSON
// metadata for the struct [V1PlanGetDetailsResponseDataOverageRatesFiatCreditType]
type v1PlanGetDetailsResponseDataOverageRatesFiatCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1PlanGetDetailsResponseDataOverageRatesFiatCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanGetDetailsResponseDataOverageRatesFiatCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1PlanListChargesResponse struct {
	ID           string                              `json:"id,required" format:"uuid"`
	ChargeType   V1PlanListChargesResponseChargeType `json:"charge_type,required"`
	CreditType   V1PlanListChargesResponseCreditType `json:"credit_type,required"`
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

type V1PlanListChargesResponseCreditType struct {
	ID   string                                  `json:"id,required" format:"uuid"`
	Name string                                  `json:"name,required"`
	JSON v1PlanListChargesResponseCreditTypeJSON `json:"-"`
}

// v1PlanListChargesResponseCreditTypeJSON contains the JSON metadata for the
// struct [V1PlanListChargesResponseCreditType]
type v1PlanListChargesResponseCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1PlanListChargesResponseCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanListChargesResponseCreditTypeJSON) RawJSON() string {
	return r.raw
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
	CustomerDetails V1PlanListCustomersResponseCustomerDetails `json:"customer_details,required"`
	PlanDetails     V1PlanListCustomersResponsePlanDetails     `json:"plan_details,required"`
	JSON            v1PlanListCustomersResponseJSON            `json:"-"`
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

type V1PlanListCustomersResponseCustomerDetails struct {
	// the Metronome ID of the customer
	ID             string                                                   `json:"id,required" format:"uuid"`
	CustomFields   map[string]string                                        `json:"custom_fields,required"`
	CustomerConfig V1PlanListCustomersResponseCustomerDetailsCustomerConfig `json:"customer_config,required"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string `json:"ingest_aliases,required"`
	Name          string   `json:"name,required"`
	// RFC 3339 timestamp indicating when the customer was archived. Null if the
	// customer is active.
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// This field's availability is dependent on your client's configuration.
	CurrentBillableStatus V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatus `json:"current_billable_status"`
	JSON                  v1PlanListCustomersResponseCustomerDetailsJSON                  `json:"-"`
}

// v1PlanListCustomersResponseCustomerDetailsJSON contains the JSON metadata for
// the struct [V1PlanListCustomersResponseCustomerDetails]
type v1PlanListCustomersResponseCustomerDetailsJSON struct {
	ID                    apijson.Field
	CustomFields          apijson.Field
	CustomerConfig        apijson.Field
	ExternalID            apijson.Field
	IngestAliases         apijson.Field
	Name                  apijson.Field
	ArchivedAt            apijson.Field
	CurrentBillableStatus apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V1PlanListCustomersResponseCustomerDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanListCustomersResponseCustomerDetailsJSON) RawJSON() string {
	return r.raw
}

type V1PlanListCustomersResponseCustomerDetailsCustomerConfig struct {
	// The Salesforce account ID for the customer
	SalesforceAccountID string                                                       `json:"salesforce_account_id,required,nullable"`
	JSON                v1PlanListCustomersResponseCustomerDetailsCustomerConfigJSON `json:"-"`
}

// v1PlanListCustomersResponseCustomerDetailsCustomerConfigJSON contains the JSON
// metadata for the struct
// [V1PlanListCustomersResponseCustomerDetailsCustomerConfig]
type v1PlanListCustomersResponseCustomerDetailsCustomerConfigJSON struct {
	SalesforceAccountID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1PlanListCustomersResponseCustomerDetailsCustomerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanListCustomersResponseCustomerDetailsCustomerConfigJSON) RawJSON() string {
	return r.raw
}

// This field's availability is dependent on your client's configuration.
type V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatus struct {
	Value       V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValue `json:"value,required"`
	EffectiveAt time.Time                                                            `json:"effective_at,nullable" format:"date-time"`
	JSON        v1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusJSON  `json:"-"`
}

// v1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusJSON contains the
// JSON metadata for the struct
// [V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatus]
type v1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusJSON struct {
	Value       apijson.Field
	EffectiveAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusJSON) RawJSON() string {
	return r.raw
}

type V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValue string

const (
	V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValueBillable   V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValue = "billable"
	V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValueUnbillable V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValue = "unbillable"
)

func (r V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValue) IsKnown() bool {
	switch r {
	case V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValueBillable, V1PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValueUnbillable:
		return true
	}
	return false
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
