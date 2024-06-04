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
)

// PlanService contains methods and other services that help with interacting with
// the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanService] method instead.
type PlanService struct {
	Options []option.RequestOption
}

// NewPlanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPlanService(opts ...option.RequestOption) (r *PlanService) {
	r = &PlanService{}
	r.Options = opts
	return
}

// List all available plans.
func (r *PlanService) List(ctx context.Context, query PlanListParams, opts ...option.RequestOption) (res *PlanListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "plans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch high level details of a specific plan.
func (r *PlanService) GetDetails(ctx context.Context, planID string, opts ...option.RequestOption) (res *PlanGetDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("planDetails/%s", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a list of charges of a specific plan.
func (r *PlanService) ListCharges(ctx context.Context, planID string, query PlanListChargesParams, opts ...option.RequestOption) (res *PlanListChargesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("planDetails/%s/charges", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetches a list of customers on a specific plan (by default, only currently
// active plans are included)
func (r *PlanService) ListCustomers(ctx context.Context, planID string, query PlanListCustomersParams, opts ...option.RequestOption) (res *PlanListCustomersResponse, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("planDetails/%s/customers", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PlanListResponse struct {
	Data     []PlanListResponseData `json:"data,required"`
	NextPage string                 `json:"next_page,required,nullable"`
	JSON     planListResponseJSON   `json:"-"`
}

// planListResponseJSON contains the JSON metadata for the struct
// [PlanListResponse]
type planListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListResponseJSON) RawJSON() string {
	return r.raw
}

type PlanListResponseData struct {
	ID           string                   `json:"id,required" format:"uuid"`
	Description  string                   `json:"description,required"`
	Name         string                   `json:"name,required"`
	CustomFields map[string]string        `json:"custom_fields"`
	JSON         planListResponseDataJSON `json:"-"`
}

// planListResponseDataJSON contains the JSON metadata for the struct
// [PlanListResponseData]
type planListResponseDataJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	CustomFields apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PlanListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListResponseDataJSON) RawJSON() string {
	return r.raw
}

type PlanGetDetailsResponse struct {
	Data PlanGetDetailsResponseData `json:"data,required"`
	JSON planGetDetailsResponseJSON `json:"-"`
}

// planGetDetailsResponseJSON contains the JSON metadata for the struct
// [PlanGetDetailsResponse]
type planGetDetailsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type PlanGetDetailsResponseData struct {
	ID           string                                  `json:"id,required" format:"uuid"`
	CustomFields map[string]string                       `json:"custom_fields,required"`
	Name         string                                  `json:"name,required"`
	CreditGrants []PlanGetDetailsResponseDataCreditGrant `json:"credit_grants"`
	Description  string                                  `json:"description"`
	Minimums     []PlanGetDetailsResponseDataMinimum     `json:"minimums"`
	OverageRates []PlanGetDetailsResponseDataOverageRate `json:"overage_rates"`
	JSON         planGetDetailsResponseDataJSON          `json:"-"`
}

// planGetDetailsResponseDataJSON contains the JSON metadata for the struct
// [PlanGetDetailsResponseData]
type planGetDetailsResponseDataJSON struct {
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

func (r *PlanGetDetailsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetDetailsResponseDataJSON) RawJSON() string {
	return r.raw
}

type PlanGetDetailsResponseDataCreditGrant struct {
	AmountGranted           float64                                                       `json:"amount_granted,required"`
	AmountGrantedCreditType PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType `json:"amount_granted_credit_type,required"`
	AmountPaid              float64                                                       `json:"amount_paid,required"`
	AmountPaidCreditType    PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType    `json:"amount_paid_credit_type,required"`
	EffectiveDuration       float64                                                       `json:"effective_duration,required"`
	Name                    string                                                        `json:"name,required"`
	Priority                string                                                        `json:"priority,required"`
	SendInvoice             bool                                                          `json:"send_invoice,required"`
	Reason                  string                                                        `json:"reason"`
	RecurrenceDuration      float64                                                       `json:"recurrence_duration"`
	RecurrenceInterval      float64                                                       `json:"recurrence_interval"`
	JSON                    planGetDetailsResponseDataCreditGrantJSON                     `json:"-"`
}

// planGetDetailsResponseDataCreditGrantJSON contains the JSON metadata for the
// struct [PlanGetDetailsResponseDataCreditGrant]
type planGetDetailsResponseDataCreditGrantJSON struct {
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

func (r *PlanGetDetailsResponseDataCreditGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetDetailsResponseDataCreditGrantJSON) RawJSON() string {
	return r.raw
}

type PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType struct {
	ID   string                                                            `json:"id,required" format:"uuid"`
	Name string                                                            `json:"name,required"`
	JSON planGetDetailsResponseDataCreditGrantsAmountGrantedCreditTypeJSON `json:"-"`
}

// planGetDetailsResponseDataCreditGrantsAmountGrantedCreditTypeJSON contains the
// JSON metadata for the struct
// [PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType]
type planGetDetailsResponseDataCreditGrantsAmountGrantedCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetDetailsResponseDataCreditGrantsAmountGrantedCreditTypeJSON) RawJSON() string {
	return r.raw
}

type PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType struct {
	ID   string                                                         `json:"id,required" format:"uuid"`
	Name string                                                         `json:"name,required"`
	JSON planGetDetailsResponseDataCreditGrantsAmountPaidCreditTypeJSON `json:"-"`
}

// planGetDetailsResponseDataCreditGrantsAmountPaidCreditTypeJSON contains the JSON
// metadata for the struct
// [PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType]
type planGetDetailsResponseDataCreditGrantsAmountPaidCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetDetailsResponseDataCreditGrantsAmountPaidCreditTypeJSON) RawJSON() string {
	return r.raw
}

type PlanGetDetailsResponseDataMinimum struct {
	CreditType PlanGetDetailsResponseDataMinimumsCreditType `json:"credit_type,required"`
	Name       string                                       `json:"name,required"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod float64                               `json:"start_period,required"`
	Value       float64                               `json:"value,required"`
	JSON        planGetDetailsResponseDataMinimumJSON `json:"-"`
}

// planGetDetailsResponseDataMinimumJSON contains the JSON metadata for the struct
// [PlanGetDetailsResponseDataMinimum]
type planGetDetailsResponseDataMinimumJSON struct {
	CreditType  apijson.Field
	Name        apijson.Field
	StartPeriod apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetDetailsResponseDataMinimumJSON) RawJSON() string {
	return r.raw
}

type PlanGetDetailsResponseDataMinimumsCreditType struct {
	ID   string                                           `json:"id,required" format:"uuid"`
	Name string                                           `json:"name,required"`
	JSON planGetDetailsResponseDataMinimumsCreditTypeJSON `json:"-"`
}

// planGetDetailsResponseDataMinimumsCreditTypeJSON contains the JSON metadata for
// the struct [PlanGetDetailsResponseDataMinimumsCreditType]
type planGetDetailsResponseDataMinimumsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataMinimumsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetDetailsResponseDataMinimumsCreditTypeJSON) RawJSON() string {
	return r.raw
}

type PlanGetDetailsResponseDataOverageRate struct {
	CreditType     PlanGetDetailsResponseDataOverageRatesCreditType     `json:"credit_type,required"`
	FiatCreditType PlanGetDetailsResponseDataOverageRatesFiatCreditType `json:"fiat_credit_type,required"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod            float64                                   `json:"start_period,required"`
	ToFiatConversionFactor float64                                   `json:"to_fiat_conversion_factor,required"`
	JSON                   planGetDetailsResponseDataOverageRateJSON `json:"-"`
}

// planGetDetailsResponseDataOverageRateJSON contains the JSON metadata for the
// struct [PlanGetDetailsResponseDataOverageRate]
type planGetDetailsResponseDataOverageRateJSON struct {
	CreditType             apijson.Field
	FiatCreditType         apijson.Field
	StartPeriod            apijson.Field
	ToFiatConversionFactor apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataOverageRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetDetailsResponseDataOverageRateJSON) RawJSON() string {
	return r.raw
}

type PlanGetDetailsResponseDataOverageRatesCreditType struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	Name string                                               `json:"name,required"`
	JSON planGetDetailsResponseDataOverageRatesCreditTypeJSON `json:"-"`
}

// planGetDetailsResponseDataOverageRatesCreditTypeJSON contains the JSON metadata
// for the struct [PlanGetDetailsResponseDataOverageRatesCreditType]
type planGetDetailsResponseDataOverageRatesCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataOverageRatesCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetDetailsResponseDataOverageRatesCreditTypeJSON) RawJSON() string {
	return r.raw
}

type PlanGetDetailsResponseDataOverageRatesFiatCreditType struct {
	ID   string                                                   `json:"id,required" format:"uuid"`
	Name string                                                   `json:"name,required"`
	JSON planGetDetailsResponseDataOverageRatesFiatCreditTypeJSON `json:"-"`
}

// planGetDetailsResponseDataOverageRatesFiatCreditTypeJSON contains the JSON
// metadata for the struct [PlanGetDetailsResponseDataOverageRatesFiatCreditType]
type planGetDetailsResponseDataOverageRatesFiatCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataOverageRatesFiatCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetDetailsResponseDataOverageRatesFiatCreditTypeJSON) RawJSON() string {
	return r.raw
}

type PlanListChargesResponse struct {
	Data     []PlanListChargesResponseData `json:"data,required"`
	NextPage string                        `json:"next_page,required,nullable"`
	JSON     planListChargesResponseJSON   `json:"-"`
}

// planListChargesResponseJSON contains the JSON metadata for the struct
// [PlanListChargesResponse]
type planListChargesResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListChargesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListChargesResponseJSON) RawJSON() string {
	return r.raw
}

type PlanListChargesResponseData struct {
	ID           string                                `json:"id,required" format:"uuid"`
	ChargeType   PlanListChargesResponseDataChargeType `json:"charge_type,required"`
	CreditType   PlanListChargesResponseDataCreditType `json:"credit_type,required"`
	CustomFields map[string]string                     `json:"custom_fields,required"`
	Name         string                                `json:"name,required"`
	Prices       []PlanListChargesResponseDataPrice    `json:"prices,required"`
	ProductID    string                                `json:"product_id,required"`
	ProductName  string                                `json:"product_name,required"`
	Quantity     float64                               `json:"quantity"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod float64 `json:"start_period"`
	// Specifies how quantities for usage based charges will be converted.
	UnitConversion PlanListChargesResponseDataUnitConversion `json:"unit_conversion"`
	JSON           planListChargesResponseDataJSON           `json:"-"`
}

// planListChargesResponseDataJSON contains the JSON metadata for the struct
// [PlanListChargesResponseData]
type planListChargesResponseDataJSON struct {
	ID             apijson.Field
	ChargeType     apijson.Field
	CreditType     apijson.Field
	CustomFields   apijson.Field
	Name           apijson.Field
	Prices         apijson.Field
	ProductID      apijson.Field
	ProductName    apijson.Field
	Quantity       apijson.Field
	StartPeriod    apijson.Field
	UnitConversion apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlanListChargesResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListChargesResponseDataJSON) RawJSON() string {
	return r.raw
}

type PlanListChargesResponseDataChargeType string

const (
	PlanListChargesResponseDataChargeTypeUsage     PlanListChargesResponseDataChargeType = "usage"
	PlanListChargesResponseDataChargeTypeFixed     PlanListChargesResponseDataChargeType = "fixed"
	PlanListChargesResponseDataChargeTypeComposite PlanListChargesResponseDataChargeType = "composite"
	PlanListChargesResponseDataChargeTypeMinimum   PlanListChargesResponseDataChargeType = "minimum"
	PlanListChargesResponseDataChargeTypeSeat      PlanListChargesResponseDataChargeType = "seat"
)

func (r PlanListChargesResponseDataChargeType) IsKnown() bool {
	switch r {
	case PlanListChargesResponseDataChargeTypeUsage, PlanListChargesResponseDataChargeTypeFixed, PlanListChargesResponseDataChargeTypeComposite, PlanListChargesResponseDataChargeTypeMinimum, PlanListChargesResponseDataChargeTypeSeat:
		return true
	}
	return false
}

type PlanListChargesResponseDataCreditType struct {
	ID   string                                    `json:"id,required" format:"uuid"`
	Name string                                    `json:"name,required"`
	JSON planListChargesResponseDataCreditTypeJSON `json:"-"`
}

// planListChargesResponseDataCreditTypeJSON contains the JSON metadata for the
// struct [PlanListChargesResponseDataCreditType]
type planListChargesResponseDataCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListChargesResponseDataCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListChargesResponseDataCreditTypeJSON) RawJSON() string {
	return r.raw
}

type PlanListChargesResponseDataPrice struct {
	// Used in pricing tiers. Indicates at what metric value the price applies.
	Tier               float64                              `json:"tier,required"`
	Value              float64                              `json:"value,required"`
	CollectionInterval float64                              `json:"collection_interval"`
	CollectionSchedule string                               `json:"collection_schedule"`
	Quantity           float64                              `json:"quantity"`
	JSON               planListChargesResponseDataPriceJSON `json:"-"`
}

// planListChargesResponseDataPriceJSON contains the JSON metadata for the struct
// [PlanListChargesResponseDataPrice]
type planListChargesResponseDataPriceJSON struct {
	Tier               apijson.Field
	Value              apijson.Field
	CollectionInterval apijson.Field
	CollectionSchedule apijson.Field
	Quantity           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PlanListChargesResponseDataPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListChargesResponseDataPriceJSON) RawJSON() string {
	return r.raw
}

// Specifies how quantities for usage based charges will be converted.
type PlanListChargesResponseDataUnitConversion struct {
	// The conversion factor
	DivisionFactor float64 `json:"division_factor,required"`
	// Whether usage should be rounded down or up to the nearest whole number. If null,
	// quantity will be rounded to 20 decimal places.
	RoundingBehavior PlanListChargesResponseDataUnitConversionRoundingBehavior `json:"rounding_behavior"`
	JSON             planListChargesResponseDataUnitConversionJSON             `json:"-"`
}

// planListChargesResponseDataUnitConversionJSON contains the JSON metadata for the
// struct [PlanListChargesResponseDataUnitConversion]
type planListChargesResponseDataUnitConversionJSON struct {
	DivisionFactor   apijson.Field
	RoundingBehavior apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PlanListChargesResponseDataUnitConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListChargesResponseDataUnitConversionJSON) RawJSON() string {
	return r.raw
}

// Whether usage should be rounded down or up to the nearest whole number. If null,
// quantity will be rounded to 20 decimal places.
type PlanListChargesResponseDataUnitConversionRoundingBehavior string

const (
	PlanListChargesResponseDataUnitConversionRoundingBehaviorFloor   PlanListChargesResponseDataUnitConversionRoundingBehavior = "floor"
	PlanListChargesResponseDataUnitConversionRoundingBehaviorCeiling PlanListChargesResponseDataUnitConversionRoundingBehavior = "ceiling"
)

func (r PlanListChargesResponseDataUnitConversionRoundingBehavior) IsKnown() bool {
	switch r {
	case PlanListChargesResponseDataUnitConversionRoundingBehaviorFloor, PlanListChargesResponseDataUnitConversionRoundingBehaviorCeiling:
		return true
	}
	return false
}

type PlanListCustomersResponse struct {
	Data     []PlanListCustomersResponseData `json:"data,required"`
	NextPage string                          `json:"next_page,required,nullable"`
	JSON     planListCustomersResponseJSON   `json:"-"`
}

// planListCustomersResponseJSON contains the JSON metadata for the struct
// [PlanListCustomersResponse]
type planListCustomersResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListCustomersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListCustomersResponseJSON) RawJSON() string {
	return r.raw
}

type PlanListCustomersResponseData struct {
	CustomerDetails PlanListCustomersResponseDataCustomerDetails `json:"customer_details,required"`
	PlanDetails     PlanListCustomersResponseDataPlanDetails     `json:"plan_details,required"`
	JSON            planListCustomersResponseDataJSON            `json:"-"`
}

// planListCustomersResponseDataJSON contains the JSON metadata for the struct
// [PlanListCustomersResponseData]
type planListCustomersResponseDataJSON struct {
	CustomerDetails apijson.Field
	PlanDetails     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PlanListCustomersResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListCustomersResponseDataJSON) RawJSON() string {
	return r.raw
}

type PlanListCustomersResponseDataCustomerDetails struct {
	// the Metronome ID of the customer
	ID                    string                                                            `json:"id,required" format:"uuid"`
	CurrentBillableStatus PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatus `json:"current_billable_status,required"`
	CustomFields          map[string]string                                                 `json:"custom_fields,required"`
	CustomerConfig        PlanListCustomersResponseDataCustomerDetailsCustomerConfig        `json:"customer_config,required"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string                                         `json:"ingest_aliases,required"`
	Name          string                                           `json:"name,required"`
	JSON          planListCustomersResponseDataCustomerDetailsJSON `json:"-"`
}

// planListCustomersResponseDataCustomerDetailsJSON contains the JSON metadata for
// the struct [PlanListCustomersResponseDataCustomerDetails]
type planListCustomersResponseDataCustomerDetailsJSON struct {
	ID                    apijson.Field
	CurrentBillableStatus apijson.Field
	CustomFields          apijson.Field
	CustomerConfig        apijson.Field
	ExternalID            apijson.Field
	IngestAliases         apijson.Field
	Name                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PlanListCustomersResponseDataCustomerDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListCustomersResponseDataCustomerDetailsJSON) RawJSON() string {
	return r.raw
}

type PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatus struct {
	Value       PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatusValue `json:"value,required"`
	EffectiveAt time.Time                                                              `json:"effective_at,nullable" format:"date-time"`
	JSON        planListCustomersResponseDataCustomerDetailsCurrentBillableStatusJSON  `json:"-"`
}

// planListCustomersResponseDataCustomerDetailsCurrentBillableStatusJSON contains
// the JSON metadata for the struct
// [PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatus]
type planListCustomersResponseDataCustomerDetailsCurrentBillableStatusJSON struct {
	Value       apijson.Field
	EffectiveAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListCustomersResponseDataCustomerDetailsCurrentBillableStatusJSON) RawJSON() string {
	return r.raw
}

type PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatusValue string

const (
	PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatusValueBillable   PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatusValue = "billable"
	PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatusValueUnbillable PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatusValue = "unbillable"
)

func (r PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatusValue) IsKnown() bool {
	switch r {
	case PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatusValueBillable, PlanListCustomersResponseDataCustomerDetailsCurrentBillableStatusValueUnbillable:
		return true
	}
	return false
}

type PlanListCustomersResponseDataCustomerDetailsCustomerConfig struct {
	// The Salesforce account ID for the customer
	SalesforceAccountID string                                                         `json:"salesforce_account_id,required,nullable"`
	JSON                planListCustomersResponseDataCustomerDetailsCustomerConfigJSON `json:"-"`
}

// planListCustomersResponseDataCustomerDetailsCustomerConfigJSON contains the JSON
// metadata for the struct
// [PlanListCustomersResponseDataCustomerDetailsCustomerConfig]
type planListCustomersResponseDataCustomerDetailsCustomerConfigJSON struct {
	SalesforceAccountID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PlanListCustomersResponseDataCustomerDetailsCustomerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListCustomersResponseDataCustomerDetailsCustomerConfigJSON) RawJSON() string {
	return r.raw
}

type PlanListCustomersResponseDataPlanDetails struct {
	ID             string            `json:"id,required" format:"uuid"`
	CustomFields   map[string]string `json:"custom_fields,required"`
	CustomerPlanID string            `json:"customer_plan_id,required" format:"uuid"`
	Name           string            `json:"name,required"`
	// The start date of the plan
	StartingOn time.Time `json:"starting_on,required" format:"date-time"`
	// The end date of the plan
	EndingBefore time.Time                                    `json:"ending_before,nullable" format:"date-time"`
	JSON         planListCustomersResponseDataPlanDetailsJSON `json:"-"`
}

// planListCustomersResponseDataPlanDetailsJSON contains the JSON metadata for the
// struct [PlanListCustomersResponseDataPlanDetails]
type planListCustomersResponseDataPlanDetailsJSON struct {
	ID             apijson.Field
	CustomFields   apijson.Field
	CustomerPlanID apijson.Field
	Name           apijson.Field
	StartingOn     apijson.Field
	EndingBefore   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlanListCustomersResponseDataPlanDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListCustomersResponseDataPlanDetailsJSON) RawJSON() string {
	return r.raw
}

type PlanListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [PlanListParams]'s query parameters as `url.Values`.
func (r PlanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlanListChargesParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [PlanListChargesParams]'s query parameters as `url.Values`.
func (r PlanListChargesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlanListCustomersParams struct {
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
	Status param.Field[PlanListCustomersParamsStatus] `query:"status"`
}

// URLQuery serializes [PlanListCustomersParams]'s query parameters as
// `url.Values`.
func (r PlanListCustomersParams) URLQuery() (v url.Values) {
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
type PlanListCustomersParamsStatus string

const (
	PlanListCustomersParamsStatusAll      PlanListCustomersParamsStatus = "all"
	PlanListCustomersParamsStatusActive   PlanListCustomersParamsStatus = "active"
	PlanListCustomersParamsStatusEnded    PlanListCustomersParamsStatus = "ended"
	PlanListCustomersParamsStatusUpcoming PlanListCustomersParamsStatus = "upcoming"
)

func (r PlanListCustomersParamsStatus) IsKnown() bool {
	switch r {
	case PlanListCustomersParamsStatusAll, PlanListCustomersParamsStatusActive, PlanListCustomersParamsStatusEnded, PlanListCustomersParamsStatusUpcoming:
		return true
	}
	return false
}
