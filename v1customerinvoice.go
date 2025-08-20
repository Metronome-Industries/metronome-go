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

// V1CustomerInvoiceService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerInvoiceService] method instead.
type V1CustomerInvoiceService struct {
	Options []option.RequestOption
}

// NewV1CustomerInvoiceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerInvoiceService(opts ...option.RequestOption) (r *V1CustomerInvoiceService) {
	r = &V1CustomerInvoiceService{}
	r.Options = opts
	return
}

// Fetch a specific invoice for a given customer.
func (r *V1CustomerInvoiceService) Get(ctx context.Context, params V1CustomerInvoiceGetParams, opts ...option.RequestOption) (res *V1CustomerInvoiceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if params.InvoiceID.Value == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/invoices/%s", params.CustomerID, params.InvoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// List all invoices for a given customer, optionally filtered by status, date
// range, and/or credit type.
func (r *V1CustomerInvoiceService) List(ctx context.Context, params V1CustomerInvoiceListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerInvoiceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/invoices", params.CustomerID)
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

// List all invoices for a given customer, optionally filtered by status, date
// range, and/or credit type.
func (r *V1CustomerInvoiceService) ListAutoPaging(ctx context.Context, params V1CustomerInvoiceListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CustomerInvoiceListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Add a one time charge to the specified invoice
func (r *V1CustomerInvoiceService) AddCharge(ctx context.Context, params V1CustomerInvoiceAddChargeParams, opts ...option.RequestOption) (res *V1CustomerInvoiceAddChargeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/addCharge", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List daily or hourly invoice breakdowns for a given customer, optionally
// filtered by status, date range, and/or credit type. Important considerations:
//
//   - If we receive backdated usage after an invoice has been finalized, the
//     backdated usage will be included in the response and usage numbers may differ.
func (r *V1CustomerInvoiceService) ListBreakdowns(ctx context.Context, params V1CustomerInvoiceListBreakdownsParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerInvoiceListBreakdownsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/invoices/breakdowns", params.CustomerID)
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

// List daily or hourly invoice breakdowns for a given customer, optionally
// filtered by status, date range, and/or credit type. Important considerations:
//
//   - If we receive backdated usage after an invoice has been finalized, the
//     backdated usage will be included in the response and usage numbers may differ.
func (r *V1CustomerInvoiceService) ListBreakdownsAutoPaging(ctx context.Context, params V1CustomerInvoiceListBreakdownsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CustomerInvoiceListBreakdownsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListBreakdowns(ctx, params, opts...))
}

type V1CustomerInvoiceGetResponse struct {
	Data V1CustomerInvoiceGetResponseData `json:"data,required"`
	JSON v1CustomerInvoiceGetResponseJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseJSON contains the JSON metadata for the struct
// [V1CustomerInvoiceGetResponse]
type v1CustomerInvoiceGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseData struct {
	ID          string                                     `json:"id,required" format:"uuid"`
	CreditType  V1CustomerInvoiceGetResponseDataCreditType `json:"credit_type,required"`
	CustomerID  string                                     `json:"customer_id,required" format:"uuid"`
	LineItems   []V1CustomerInvoiceGetResponseDataLineItem `json:"line_items,required"`
	Status      string                                     `json:"status,required"`
	Total       float64                                    `json:"total,required"`
	Type        string                                     `json:"type,required"`
	AmendmentID string                                     `json:"amendment_id" format:"uuid"`
	// This field's availability is dependent on your client's configuration.
	BillableStatus       V1CustomerInvoiceGetResponseDataBillableStatus   `json:"billable_status"`
	ContractCustomFields map[string]string                                `json:"contract_custom_fields"`
	ContractID           string                                           `json:"contract_id" format:"uuid"`
	CorrectionRecord     V1CustomerInvoiceGetResponseDataCorrectionRecord `json:"correction_record"`
	// When the invoice was created (UTC). This field is present for correction
	// invoices only.
	CreatedAt            time.Time              `json:"created_at" format:"date-time"`
	CustomFields         map[string]interface{} `json:"custom_fields"`
	CustomerCustomFields map[string]string      `json:"customer_custom_fields"`
	// End of the usage period this invoice covers (UTC)
	EndTimestamp       time.Time                                           `json:"end_timestamp" format:"date-time"`
	ExternalInvoice    V1CustomerInvoiceGetResponseDataExternalInvoice     `json:"external_invoice,nullable"`
	InvoiceAdjustments []V1CustomerInvoiceGetResponseDataInvoiceAdjustment `json:"invoice_adjustments"`
	// When the invoice was issued (UTC)
	IssuedAt            time.Time `json:"issued_at" format:"date-time"`
	NetPaymentTermsDays float64   `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string            `json:"netsuite_sales_order_id"`
	PlanCustomFields     map[string]string `json:"plan_custom_fields"`
	PlanID               string            `json:"plan_id" format:"uuid"`
	PlanName             string            `json:"plan_name"`
	// Only present for contract invoices with reseller royalties.
	ResellerRoyalty V1CustomerInvoiceGetResponseDataResellerRoyalty `json:"reseller_royalty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Beginning of the usage period this invoice covers (UTC)
	StartTimestamp time.Time                            `json:"start_timestamp" format:"date-time"`
	Subtotal       float64                              `json:"subtotal"`
	JSON           v1CustomerInvoiceGetResponseDataJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerInvoiceGetResponseData]
type v1CustomerInvoiceGetResponseDataJSON struct {
	ID                      apijson.Field
	CreditType              apijson.Field
	CustomerID              apijson.Field
	LineItems               apijson.Field
	Status                  apijson.Field
	Total                   apijson.Field
	Type                    apijson.Field
	AmendmentID             apijson.Field
	BillableStatus          apijson.Field
	ContractCustomFields    apijson.Field
	ContractID              apijson.Field
	CorrectionRecord        apijson.Field
	CreatedAt               apijson.Field
	CustomFields            apijson.Field
	CustomerCustomFields    apijson.Field
	EndTimestamp            apijson.Field
	ExternalInvoice         apijson.Field
	InvoiceAdjustments      apijson.Field
	IssuedAt                apijson.Field
	NetPaymentTermsDays     apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	PlanCustomFields        apijson.Field
	PlanID                  apijson.Field
	PlanName                apijson.Field
	ResellerRoyalty         apijson.Field
	SalesforceOpportunityID apijson.Field
	StartTimestamp          apijson.Field
	Subtotal                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataCreditType struct {
	ID   string                                         `json:"id,required" format:"uuid"`
	Name string                                         `json:"name,required"`
	JSON v1CustomerInvoiceGetResponseDataCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataCreditTypeJSON contains the JSON metadata for
// the struct [V1CustomerInvoiceGetResponseDataCreditType]
type v1CustomerInvoiceGetResponseDataCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataLineItem struct {
	CreditType V1CustomerInvoiceGetResponseDataLineItemsCreditType `json:"credit_type,required"`
	Name       string                                              `json:"name,required"`
	Total      float64                                             `json:"total,required"`
	// The type of line item. scheduled - Line item is associated with a scheduled
	// charge. View the scheduled_charge_id on the line item. commit_purchase - Line
	// item is associated with a payment for a prepaid commit. View the commit_id on
	// the line item. usage - Line item is associated with a usage product or composite
	// product. View the product_id on the line item to determine which product.
	// subscription - Line item is associated with a subscription. e.g. monthly
	// recurring payment for an in-advance subscription. applied_commit_or_credit - On
	// metronome invoices, applied commits and credits are associated with their own
	// line items. These line items have negative totals. Use the
	// applied_commit_or_credit object on the line item to understand the id of the
	// applied commit or credit, and its type. Note that the application of a postpaid
	// commit is associated with a line item, but the total on the line item is not
	// included in the invoice's total as postpaid commits are paid in-arrears.
	// postpaid_trueup - Line item is associated with the true up amount for a postpaid
	// commit. This line item type will only appear on invoices with type TRUEUP .
	// cpu_conversion - Line item converting between a custom pricing unit and fiat
	// currency, using the conversion rate set on the rate card. This line item will
	// appear when there are products priced in custom pricing units, and there is
	// insufficient prepaid commit/credit in that custom pricing unit to fully cover
	// the spend. Then, the outstanding spend in custom pricing units will be converted
	// to fiat currency using a cpu_conversion line item.
	Type string `json:"type,required"`
	// Details about the credit or commit that was applied to this line item. Only
	// present on line items with product of `USAGE`, `SUBSCRIPTION` or `COMPOSITE`
	// types.
	AppliedCommitOrCredit V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCredit `json:"applied_commit_or_credit"`
	CommitCustomFields    map[string]string                                              `json:"commit_custom_fields"`
	// For line items with product of `USAGE`, `SUBSCRIPTION`, or `COMPOSITE` types,
	// the ID of the credit or commit that was applied to this line item. For line
	// items with product type of `FIXED`, the ID of the prepaid or postpaid commit
	// that is being paid for.
	CommitID                   string `json:"commit_id" format:"uuid"`
	CommitNetsuiteItemID       string `json:"commit_netsuite_item_id"`
	CommitNetsuiteSalesOrderID string `json:"commit_netsuite_sales_order_id"`
	CommitSegmentID            string `json:"commit_segment_id" format:"uuid"`
	// `PrepaidCommit` (for commit types `PREPAID` and `CREDIT`) or `PostpaidCommit`
	// (for commit type `POSTPAID`).
	CommitType           string            `json:"commit_type"`
	CustomFields         map[string]string `json:"custom_fields"`
	DiscountCustomFields map[string]string `json:"discount_custom_fields"`
	// ID of the discount applied to this line item.
	DiscountID string `json:"discount_id" format:"uuid"`
	// The line item's end date (exclusive).
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	GroupKey     string    `json:"group_key"`
	GroupValue   string    `json:"group_value,nullable"`
	// Indicates whether the line item is prorated for `SUBSCRIPTION` type product.
	IsProrated bool `json:"is_prorated"`
	// Only present for contract invoices and when the `include_list_prices` query
	// parameter is set to true. This will include the list rate for the charge if
	// applicable. Only present for usage and subscription line items.
	ListPrice V1CustomerInvoiceGetResponseDataLineItemsListPrice `json:"list_price"`
	Metadata  string                                             `json:"metadata"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd time.Time `json:"netsuite_invoice_billing_end" format:"date-time"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart time.Time `json:"netsuite_invoice_billing_start" format:"date-time"`
	NetsuiteItemID              string    `json:"netsuite_item_id"`
	// Only present for line items paying for a postpaid commit true-up.
	PostpaidCommit V1CustomerInvoiceGetResponseDataLineItemsPostpaidCommit `json:"postpaid_commit"`
	// Includes the presentation group values associated with this line item if
	// presentation group keys are used.
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	// Includes the pricing group values associated with this line item if dimensional
	// pricing is used.
	PricingGroupValues  map[string]string `json:"pricing_group_values"`
	ProductCustomFields map[string]string `json:"product_custom_fields"`
	// ID of the product associated with the line item.
	ProductID string `json:"product_id" format:"uuid"`
	// The current product tags associated with the line item's `product_id`.
	ProductTags []string `json:"product_tags"`
	// The type of the line item's product. Possible values are `FixedProductListItem`
	// (for `FIXED` type products), `UsageProductListItem` (for `USAGE` type products),
	// `SubscriptionProductListItem` (for `SUBSCRIPTION` type products) or
	// `CompositeProductListItem` (for `COMPOSITE` type products). For scheduled
	// charges, commit and credit payments, the value is `FixedProductListItem`.
	ProductType                     string            `json:"product_type"`
	ProfessionalServiceCustomFields map[string]string `json:"professional_service_custom_fields"`
	ProfessionalServiceID           string            `json:"professional_service_id" format:"uuid"`
	// The quantity associated with the line item.
	Quantity                    float64                                               `json:"quantity"`
	ResellerType                V1CustomerInvoiceGetResponseDataLineItemsResellerType `json:"reseller_type"`
	ScheduledChargeCustomFields map[string]string                                     `json:"scheduled_charge_custom_fields"`
	// ID of scheduled charge.
	ScheduledChargeID string `json:"scheduled_charge_id" format:"uuid"`
	// The line item's start date (inclusive).
	StartingAt               time.Time                                              `json:"starting_at" format:"date-time"`
	SubLineItems             []V1CustomerInvoiceGetResponseDataLineItemsSubLineItem `json:"sub_line_items"`
	SubscriptionCustomFields map[string]string                                      `json:"subscription_custom_fields"`
	// Populated if the line item has a tiered price.
	Tier V1CustomerInvoiceGetResponseDataLineItemsTier `json:"tier"`
	// The unit price associated with the line item.
	UnitPrice float64                                      `json:"unit_price"`
	JSON      v1CustomerInvoiceGetResponseDataLineItemJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataLineItemJSON contains the JSON metadata for the
// struct [V1CustomerInvoiceGetResponseDataLineItem]
type v1CustomerInvoiceGetResponseDataLineItemJSON struct {
	CreditType                      apijson.Field
	Name                            apijson.Field
	Total                           apijson.Field
	Type                            apijson.Field
	AppliedCommitOrCredit           apijson.Field
	CommitCustomFields              apijson.Field
	CommitID                        apijson.Field
	CommitNetsuiteItemID            apijson.Field
	CommitNetsuiteSalesOrderID      apijson.Field
	CommitSegmentID                 apijson.Field
	CommitType                      apijson.Field
	CustomFields                    apijson.Field
	DiscountCustomFields            apijson.Field
	DiscountID                      apijson.Field
	EndingBefore                    apijson.Field
	GroupKey                        apijson.Field
	GroupValue                      apijson.Field
	IsProrated                      apijson.Field
	ListPrice                       apijson.Field
	Metadata                        apijson.Field
	NetsuiteInvoiceBillingEnd       apijson.Field
	NetsuiteInvoiceBillingStart     apijson.Field
	NetsuiteItemID                  apijson.Field
	PostpaidCommit                  apijson.Field
	PresentationGroupValues         apijson.Field
	PricingGroupValues              apijson.Field
	ProductCustomFields             apijson.Field
	ProductID                       apijson.Field
	ProductTags                     apijson.Field
	ProductType                     apijson.Field
	ProfessionalServiceCustomFields apijson.Field
	ProfessionalServiceID           apijson.Field
	Quantity                        apijson.Field
	ResellerType                    apijson.Field
	ScheduledChargeCustomFields     apijson.Field
	ScheduledChargeID               apijson.Field
	StartingAt                      apijson.Field
	SubLineItems                    apijson.Field
	SubscriptionCustomFields        apijson.Field
	Tier                            apijson.Field
	UnitPrice                       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataLineItemJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataLineItemsCreditType struct {
	ID   string                                                  `json:"id,required" format:"uuid"`
	Name string                                                  `json:"name,required"`
	JSON v1CustomerInvoiceGetResponseDataLineItemsCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataLineItemsCreditTypeJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceGetResponseDataLineItemsCreditType]
type v1CustomerInvoiceGetResponseDataLineItemsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataLineItemsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataLineItemsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// Details about the credit or commit that was applied to this line item. Only
// present on line items with product of `USAGE`, `SUBSCRIPTION` or `COMPOSITE`
// types.
type V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCredit struct {
	ID   string                                                             `json:"id,required" format:"uuid"`
	Type V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditType `json:"type,required"`
	JSON v1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditJSON contains the
// JSON metadata for the struct
// [V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCredit]
type v1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditType string

const (
	V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditTypePrepaid  V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditType = "PREPAID"
	V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditTypePostpaid V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditType = "POSTPAID"
	V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditTypeCredit   V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditType = "CREDIT"
)

func (r V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditTypePrepaid, V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditTypePostpaid, V1CustomerInvoiceGetResponseDataLineItemsAppliedCommitOrCreditTypeCredit:
		return true
	}
	return false
}

// Only present for contract invoices and when the `include_list_prices` query
// parameter is set to true. This will include the list rate for the charge if
// applicable. Only present for usage and subscription line items.
type V1CustomerInvoiceGetResponseDataLineItemsListPrice struct {
	RateType   V1CustomerInvoiceGetResponseDataLineItemsListPriceRateType   `json:"rate_type,required"`
	CreditType V1CustomerInvoiceGetResponseDataLineItemsListPriceCreditType `json:"credit_type"`
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
	Tiers []V1CustomerInvoiceGetResponseDataLineItemsListPriceTier `json:"tiers"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool                                                   `json:"use_list_prices"`
	JSON          v1CustomerInvoiceGetResponseDataLineItemsListPriceJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataLineItemsListPriceJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceGetResponseDataLineItemsListPrice]
type v1CustomerInvoiceGetResponseDataLineItemsListPriceJSON struct {
	RateType           apijson.Field
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

func (r *V1CustomerInvoiceGetResponseDataLineItemsListPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataLineItemsListPriceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataLineItemsListPriceRateType string

const (
	V1CustomerInvoiceGetResponseDataLineItemsListPriceRateTypeFlat         V1CustomerInvoiceGetResponseDataLineItemsListPriceRateType = "FLAT"
	V1CustomerInvoiceGetResponseDataLineItemsListPriceRateTypePercentage   V1CustomerInvoiceGetResponseDataLineItemsListPriceRateType = "PERCENTAGE"
	V1CustomerInvoiceGetResponseDataLineItemsListPriceRateTypeSubscription V1CustomerInvoiceGetResponseDataLineItemsListPriceRateType = "SUBSCRIPTION"
	V1CustomerInvoiceGetResponseDataLineItemsListPriceRateTypeCustom       V1CustomerInvoiceGetResponseDataLineItemsListPriceRateType = "CUSTOM"
	V1CustomerInvoiceGetResponseDataLineItemsListPriceRateTypeTiered       V1CustomerInvoiceGetResponseDataLineItemsListPriceRateType = "TIERED"
)

func (r V1CustomerInvoiceGetResponseDataLineItemsListPriceRateType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceGetResponseDataLineItemsListPriceRateTypeFlat, V1CustomerInvoiceGetResponseDataLineItemsListPriceRateTypePercentage, V1CustomerInvoiceGetResponseDataLineItemsListPriceRateTypeSubscription, V1CustomerInvoiceGetResponseDataLineItemsListPriceRateTypeCustom, V1CustomerInvoiceGetResponseDataLineItemsListPriceRateTypeTiered:
		return true
	}
	return false
}

type V1CustomerInvoiceGetResponseDataLineItemsListPriceCreditType struct {
	ID   string                                                           `json:"id,required" format:"uuid"`
	Name string                                                           `json:"name,required"`
	JSON v1CustomerInvoiceGetResponseDataLineItemsListPriceCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataLineItemsListPriceCreditTypeJSON contains the
// JSON metadata for the struct
// [V1CustomerInvoiceGetResponseDataLineItemsListPriceCreditType]
type v1CustomerInvoiceGetResponseDataLineItemsListPriceCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataLineItemsListPriceCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataLineItemsListPriceCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataLineItemsListPriceTier struct {
	Price float64                                                    `json:"price,required"`
	Size  float64                                                    `json:"size"`
	JSON  v1CustomerInvoiceGetResponseDataLineItemsListPriceTierJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataLineItemsListPriceTierJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceGetResponseDataLineItemsListPriceTier]
type v1CustomerInvoiceGetResponseDataLineItemsListPriceTierJSON struct {
	Price       apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataLineItemsListPriceTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataLineItemsListPriceTierJSON) RawJSON() string {
	return r.raw
}

// Only present for line items paying for a postpaid commit true-up.
type V1CustomerInvoiceGetResponseDataLineItemsPostpaidCommit struct {
	ID   string                                                      `json:"id,required" format:"uuid"`
	JSON v1CustomerInvoiceGetResponseDataLineItemsPostpaidCommitJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataLineItemsPostpaidCommitJSON contains the JSON
// metadata for the struct
// [V1CustomerInvoiceGetResponseDataLineItemsPostpaidCommit]
type v1CustomerInvoiceGetResponseDataLineItemsPostpaidCommitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataLineItemsPostpaidCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataLineItemsPostpaidCommitJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataLineItemsResellerType string

const (
	V1CustomerInvoiceGetResponseDataLineItemsResellerTypeAws           V1CustomerInvoiceGetResponseDataLineItemsResellerType = "AWS"
	V1CustomerInvoiceGetResponseDataLineItemsResellerTypeAwsProService V1CustomerInvoiceGetResponseDataLineItemsResellerType = "AWS_PRO_SERVICE"
	V1CustomerInvoiceGetResponseDataLineItemsResellerTypeGcp           V1CustomerInvoiceGetResponseDataLineItemsResellerType = "GCP"
	V1CustomerInvoiceGetResponseDataLineItemsResellerTypeGcpProService V1CustomerInvoiceGetResponseDataLineItemsResellerType = "GCP_PRO_SERVICE"
)

func (r V1CustomerInvoiceGetResponseDataLineItemsResellerType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceGetResponseDataLineItemsResellerTypeAws, V1CustomerInvoiceGetResponseDataLineItemsResellerTypeAwsProService, V1CustomerInvoiceGetResponseDataLineItemsResellerTypeGcp, V1CustomerInvoiceGetResponseDataLineItemsResellerTypeGcpProService:
		return true
	}
	return false
}

type V1CustomerInvoiceGetResponseDataLineItemsSubLineItem struct {
	CustomFields  map[string]string `json:"custom_fields,required"`
	Name          string            `json:"name,required"`
	Quantity      float64           `json:"quantity,required"`
	Subtotal      float64           `json:"subtotal,required"`
	ChargeID      string            `json:"charge_id" format:"uuid"`
	CreditGrantID string            `json:"credit_grant_id" format:"uuid"`
	// The end date for the charge (for seats charges only).
	EndDate time.Time `json:"end_date" format:"date-time"`
	// the unit price for this charge, present only if the charge is not tiered and the
	// quantity is nonzero
	Price float64 `json:"price"`
	// The start date for the charge (for seats charges only).
	StartDate time.Time `json:"start_date" format:"date-time"`
	// when the current tier started and ends (for tiered charges only)
	TierPeriod V1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierPeriod `json:"tier_period"`
	Tiers      []V1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTier     `json:"tiers"`
	JSON       v1CustomerInvoiceGetResponseDataLineItemsSubLineItemJSON        `json:"-"`
}

// v1CustomerInvoiceGetResponseDataLineItemsSubLineItemJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceGetResponseDataLineItemsSubLineItem]
type v1CustomerInvoiceGetResponseDataLineItemsSubLineItemJSON struct {
	CustomFields  apijson.Field
	Name          apijson.Field
	Quantity      apijson.Field
	Subtotal      apijson.Field
	ChargeID      apijson.Field
	CreditGrantID apijson.Field
	EndDate       apijson.Field
	Price         apijson.Field
	StartDate     apijson.Field
	TierPeriod    apijson.Field
	Tiers         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataLineItemsSubLineItemJSON) RawJSON() string {
	return r.raw
}

// when the current tier started and ends (for tiered charges only)
type V1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierPeriod struct {
	StartingAt   time.Time                                                           `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time                                                           `json:"ending_before" format:"date-time"`
	JSON         v1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierPeriodJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierPeriodJSON contains the
// JSON metadata for the struct
// [V1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierPeriod]
type v1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierPeriodJSON struct {
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierPeriodJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTier struct {
	Price    float64 `json:"price,required"`
	Quantity float64 `json:"quantity,required"`
	// at what metric amount this tier begins
	StartingAt float64                                                       `json:"starting_at,required"`
	Subtotal   float64                                                       `json:"subtotal,required"`
	JSON       v1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierJSON contains the JSON
// metadata for the struct
// [V1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTier]
type v1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierJSON struct {
	Price       apijson.Field
	Quantity    apijson.Field
	StartingAt  apijson.Field
	Subtotal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataLineItemsSubLineItemsTierJSON) RawJSON() string {
	return r.raw
}

// Populated if the line item has a tiered price.
type V1CustomerInvoiceGetResponseDataLineItemsTier struct {
	Level      float64                                           `json:"level,required"`
	StartingAt string                                            `json:"starting_at,required"`
	Size       string                                            `json:"size,nullable"`
	JSON       v1CustomerInvoiceGetResponseDataLineItemsTierJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataLineItemsTierJSON contains the JSON metadata for
// the struct [V1CustomerInvoiceGetResponseDataLineItemsTier]
type v1CustomerInvoiceGetResponseDataLineItemsTierJSON struct {
	Level       apijson.Field
	StartingAt  apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataLineItemsTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataLineItemsTierJSON) RawJSON() string {
	return r.raw
}

// This field's availability is dependent on your client's configuration.
type V1CustomerInvoiceGetResponseDataBillableStatus string

const (
	V1CustomerInvoiceGetResponseDataBillableStatusBillable   V1CustomerInvoiceGetResponseDataBillableStatus = "billable"
	V1CustomerInvoiceGetResponseDataBillableStatusUnbillable V1CustomerInvoiceGetResponseDataBillableStatus = "unbillable"
)

func (r V1CustomerInvoiceGetResponseDataBillableStatus) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceGetResponseDataBillableStatusBillable, V1CustomerInvoiceGetResponseDataBillableStatusUnbillable:
		return true
	}
	return false
}

type V1CustomerInvoiceGetResponseDataCorrectionRecord struct {
	CorrectedInvoiceID       string                                                                   `json:"corrected_invoice_id,required" format:"uuid"`
	Memo                     string                                                                   `json:"memo,required"`
	Reason                   string                                                                   `json:"reason,required"`
	CorrectedExternalInvoice V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoice `json:"corrected_external_invoice"`
	JSON                     v1CustomerInvoiceGetResponseDataCorrectionRecordJSON                     `json:"-"`
}

// v1CustomerInvoiceGetResponseDataCorrectionRecordJSON contains the JSON metadata
// for the struct [V1CustomerInvoiceGetResponseDataCorrectionRecord]
type v1CustomerInvoiceGetResponseDataCorrectionRecordJSON struct {
	CorrectedInvoiceID       apijson.Field
	Memo                     apijson.Field
	Reason                   apijson.Field
	CorrectedExternalInvoice apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataCorrectionRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataCorrectionRecordJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoice struct {
	BillingProviderType V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                                                      `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                                                   `json:"issued_at_timestamp" format:"date-time"`
	JSON                v1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceJSON                `json:"-"`
}

// v1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceJSON
// contains the JSON metadata for the struct
// [V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoice]
type v1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType string

const (
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace   V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "aws_marketplace"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe           V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "stripe"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite         V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "netsuite"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom           V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "custom"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "azure_marketplace"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "quickbooks_online"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday          V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "workday"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace   V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus string

const (
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft               V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DRAFT"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized           V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "FINALIZED"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid                V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAID"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible       V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid                V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "VOID"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted             V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DELETED"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed       V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped             V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SKIPPED"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSent                V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SENT"
	V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued              V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "QUEUED"
)

func (r V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSent, V1CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type V1CustomerInvoiceGetResponseDataExternalInvoice struct {
	BillingProviderType V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                             `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                          `json:"issued_at_timestamp" format:"date-time"`
	JSON                v1CustomerInvoiceGetResponseDataExternalInvoiceJSON                `json:"-"`
}

// v1CustomerInvoiceGetResponseDataExternalInvoiceJSON contains the JSON metadata
// for the struct [V1CustomerInvoiceGetResponseDataExternalInvoice]
type v1CustomerInvoiceGetResponseDataExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType string

const (
	V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeAwsMarketplace   V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "aws_marketplace"
	V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeStripe           V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "stripe"
	V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeNetsuite         V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "netsuite"
	V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeCustom           V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "custom"
	V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeAzureMarketplace V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "azure_marketplace"
	V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeQuickbooksOnline V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "quickbooks_online"
	V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeWorkday          V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "workday"
	V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeGcpMarketplace   V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeAwsMarketplace, V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeStripe, V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeNetsuite, V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeCustom, V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeAzureMarketplace, V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeQuickbooksOnline, V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeWorkday, V1CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus string

const (
	V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusDraft               V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "DRAFT"
	V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusFinalized           V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "FINALIZED"
	V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusPaid                V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "PAID"
	V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusUncollectible       V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusVoid                V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "VOID"
	V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusDeleted             V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "DELETED"
	V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusPaymentFailed       V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusInvalidRequestError V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusSkipped             V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "SKIPPED"
	V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusSent                V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "SENT"
	V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusQueued              V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "QUEUED"
)

func (r V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusDraft, V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusFinalized, V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusPaid, V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusUncollectible, V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusVoid, V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusDeleted, V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusPaymentFailed, V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusInvalidRequestError, V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusSkipped, V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusSent, V1CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type V1CustomerInvoiceGetResponseDataInvoiceAdjustment struct {
	CreditType              V1CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditType `json:"credit_type,required"`
	Name                    string                                                       `json:"name,required"`
	Total                   float64                                                      `json:"total,required"`
	CreditGrantCustomFields map[string]string                                            `json:"credit_grant_custom_fields"`
	CreditGrantID           string                                                       `json:"credit_grant_id"`
	JSON                    v1CustomerInvoiceGetResponseDataInvoiceAdjustmentJSON        `json:"-"`
}

// v1CustomerInvoiceGetResponseDataInvoiceAdjustmentJSON contains the JSON metadata
// for the struct [V1CustomerInvoiceGetResponseDataInvoiceAdjustment]
type v1CustomerInvoiceGetResponseDataInvoiceAdjustmentJSON struct {
	CreditType              apijson.Field
	Name                    apijson.Field
	Total                   apijson.Field
	CreditGrantCustomFields apijson.Field
	CreditGrantID           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataInvoiceAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataInvoiceAdjustmentJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditType struct {
	ID   string                                                           `json:"id,required" format:"uuid"`
	Name string                                                           `json:"name,required"`
	JSON v1CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditTypeJSON contains the
// JSON metadata for the struct
// [V1CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditType]
type v1CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// Only present for contract invoices with reseller royalties.
type V1CustomerInvoiceGetResponseDataResellerRoyalty struct {
	Fraction           string                                                      `json:"fraction,required"`
	NetsuiteResellerID string                                                      `json:"netsuite_reseller_id,required"`
	ResellerType       V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerType `json:"reseller_type,required"`
	AwsOptions         V1CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptions   `json:"aws_options"`
	GcpOptions         V1CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptions   `json:"gcp_options"`
	JSON               v1CustomerInvoiceGetResponseDataResellerRoyaltyJSON         `json:"-"`
}

// v1CustomerInvoiceGetResponseDataResellerRoyaltyJSON contains the JSON metadata
// for the struct [V1CustomerInvoiceGetResponseDataResellerRoyalty]
type v1CustomerInvoiceGetResponseDataResellerRoyaltyJSON struct {
	Fraction           apijson.Field
	NetsuiteResellerID apijson.Field
	ResellerType       apijson.Field
	AwsOptions         apijson.Field
	GcpOptions         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerType string

const (
	V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeAws           V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerType = "AWS"
	V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeAwsProService V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerType = "AWS_PRO_SERVICE"
	V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeGcp           V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerType = "GCP"
	V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeGcpProService V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerType = "GCP_PRO_SERVICE"
)

func (r V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeAws, V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeAwsProService, V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeGcp, V1CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeGcpProService:
		return true
	}
	return false
}

type V1CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    string                                                        `json:"aws_account_number"`
	AwsOfferID          string                                                        `json:"aws_offer_id"`
	AwsPayerReferenceID string                                                        `json:"aws_payer_reference_id"`
	JSON                v1CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptionsJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptionsJSON contains the JSON
// metadata for the struct
// [V1CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptions]
type v1CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptionsJSON struct {
	AwsAccountNumber    apijson.Field
	AwsOfferID          apijson.Field
	AwsPayerReferenceID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptionsJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptions struct {
	GcpAccountID string                                                        `json:"gcp_account_id"`
	GcpOfferID   string                                                        `json:"gcp_offer_id"`
	JSON         v1CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptionsJSON `json:"-"`
}

// v1CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptionsJSON contains the JSON
// metadata for the struct
// [V1CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptions]
type v1CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptionsJSON struct {
	GcpAccountID apijson.Field
	GcpOfferID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptionsJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponse struct {
	ID          string                                  `json:"id,required" format:"uuid"`
	CreditType  V1CustomerInvoiceListResponseCreditType `json:"credit_type,required"`
	CustomerID  string                                  `json:"customer_id,required" format:"uuid"`
	LineItems   []V1CustomerInvoiceListResponseLineItem `json:"line_items,required"`
	Status      string                                  `json:"status,required"`
	Total       float64                                 `json:"total,required"`
	Type        string                                  `json:"type,required"`
	AmendmentID string                                  `json:"amendment_id" format:"uuid"`
	// This field's availability is dependent on your client's configuration.
	BillableStatus       V1CustomerInvoiceListResponseBillableStatus   `json:"billable_status"`
	ContractCustomFields map[string]string                             `json:"contract_custom_fields"`
	ContractID           string                                        `json:"contract_id" format:"uuid"`
	CorrectionRecord     V1CustomerInvoiceListResponseCorrectionRecord `json:"correction_record"`
	// When the invoice was created (UTC). This field is present for correction
	// invoices only.
	CreatedAt            time.Time              `json:"created_at" format:"date-time"`
	CustomFields         map[string]interface{} `json:"custom_fields"`
	CustomerCustomFields map[string]string      `json:"customer_custom_fields"`
	// End of the usage period this invoice covers (UTC)
	EndTimestamp       time.Time                                        `json:"end_timestamp" format:"date-time"`
	ExternalInvoice    V1CustomerInvoiceListResponseExternalInvoice     `json:"external_invoice,nullable"`
	InvoiceAdjustments []V1CustomerInvoiceListResponseInvoiceAdjustment `json:"invoice_adjustments"`
	// When the invoice was issued (UTC)
	IssuedAt            time.Time `json:"issued_at" format:"date-time"`
	NetPaymentTermsDays float64   `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string            `json:"netsuite_sales_order_id"`
	PlanCustomFields     map[string]string `json:"plan_custom_fields"`
	PlanID               string            `json:"plan_id" format:"uuid"`
	PlanName             string            `json:"plan_name"`
	// Only present for contract invoices with reseller royalties.
	ResellerRoyalty V1CustomerInvoiceListResponseResellerRoyalty `json:"reseller_royalty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Beginning of the usage period this invoice covers (UTC)
	StartTimestamp time.Time                         `json:"start_timestamp" format:"date-time"`
	Subtotal       float64                           `json:"subtotal"`
	JSON           v1CustomerInvoiceListResponseJSON `json:"-"`
}

// v1CustomerInvoiceListResponseJSON contains the JSON metadata for the struct
// [V1CustomerInvoiceListResponse]
type v1CustomerInvoiceListResponseJSON struct {
	ID                      apijson.Field
	CreditType              apijson.Field
	CustomerID              apijson.Field
	LineItems               apijson.Field
	Status                  apijson.Field
	Total                   apijson.Field
	Type                    apijson.Field
	AmendmentID             apijson.Field
	BillableStatus          apijson.Field
	ContractCustomFields    apijson.Field
	ContractID              apijson.Field
	CorrectionRecord        apijson.Field
	CreatedAt               apijson.Field
	CustomFields            apijson.Field
	CustomerCustomFields    apijson.Field
	EndTimestamp            apijson.Field
	ExternalInvoice         apijson.Field
	InvoiceAdjustments      apijson.Field
	IssuedAt                apijson.Field
	NetPaymentTermsDays     apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	PlanCustomFields        apijson.Field
	PlanID                  apijson.Field
	PlanName                apijson.Field
	ResellerRoyalty         apijson.Field
	SalesforceOpportunityID apijson.Field
	StartTimestamp          apijson.Field
	Subtotal                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseCreditType struct {
	ID   string                                      `json:"id,required" format:"uuid"`
	Name string                                      `json:"name,required"`
	JSON v1CustomerInvoiceListResponseCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceListResponseCreditTypeJSON contains the JSON metadata for the
// struct [V1CustomerInvoiceListResponseCreditType]
type v1CustomerInvoiceListResponseCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseLineItem struct {
	CreditType V1CustomerInvoiceListResponseLineItemsCreditType `json:"credit_type,required"`
	Name       string                                           `json:"name,required"`
	Total      float64                                          `json:"total,required"`
	// The type of line item. scheduled - Line item is associated with a scheduled
	// charge. View the scheduled_charge_id on the line item. commit_purchase - Line
	// item is associated with a payment for a prepaid commit. View the commit_id on
	// the line item. usage - Line item is associated with a usage product or composite
	// product. View the product_id on the line item to determine which product.
	// subscription - Line item is associated with a subscription. e.g. monthly
	// recurring payment for an in-advance subscription. applied_commit_or_credit - On
	// metronome invoices, applied commits and credits are associated with their own
	// line items. These line items have negative totals. Use the
	// applied_commit_or_credit object on the line item to understand the id of the
	// applied commit or credit, and its type. Note that the application of a postpaid
	// commit is associated with a line item, but the total on the line item is not
	// included in the invoice's total as postpaid commits are paid in-arrears.
	// postpaid_trueup - Line item is associated with the true up amount for a postpaid
	// commit. This line item type will only appear on invoices with type TRUEUP .
	// cpu_conversion - Line item converting between a custom pricing unit and fiat
	// currency, using the conversion rate set on the rate card. This line item will
	// appear when there are products priced in custom pricing units, and there is
	// insufficient prepaid commit/credit in that custom pricing unit to fully cover
	// the spend. Then, the outstanding spend in custom pricing units will be converted
	// to fiat currency using a cpu_conversion line item.
	Type string `json:"type,required"`
	// Details about the credit or commit that was applied to this line item. Only
	// present on line items with product of `USAGE`, `SUBSCRIPTION` or `COMPOSITE`
	// types.
	AppliedCommitOrCredit V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCredit `json:"applied_commit_or_credit"`
	CommitCustomFields    map[string]string                                           `json:"commit_custom_fields"`
	// For line items with product of `USAGE`, `SUBSCRIPTION`, or `COMPOSITE` types,
	// the ID of the credit or commit that was applied to this line item. For line
	// items with product type of `FIXED`, the ID of the prepaid or postpaid commit
	// that is being paid for.
	CommitID                   string `json:"commit_id" format:"uuid"`
	CommitNetsuiteItemID       string `json:"commit_netsuite_item_id"`
	CommitNetsuiteSalesOrderID string `json:"commit_netsuite_sales_order_id"`
	CommitSegmentID            string `json:"commit_segment_id" format:"uuid"`
	// `PrepaidCommit` (for commit types `PREPAID` and `CREDIT`) or `PostpaidCommit`
	// (for commit type `POSTPAID`).
	CommitType           string            `json:"commit_type"`
	CustomFields         map[string]string `json:"custom_fields"`
	DiscountCustomFields map[string]string `json:"discount_custom_fields"`
	// ID of the discount applied to this line item.
	DiscountID string `json:"discount_id" format:"uuid"`
	// The line item's end date (exclusive).
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	GroupKey     string    `json:"group_key"`
	GroupValue   string    `json:"group_value,nullable"`
	// Indicates whether the line item is prorated for `SUBSCRIPTION` type product.
	IsProrated bool `json:"is_prorated"`
	// Only present for contract invoices and when the `include_list_prices` query
	// parameter is set to true. This will include the list rate for the charge if
	// applicable. Only present for usage and subscription line items.
	ListPrice V1CustomerInvoiceListResponseLineItemsListPrice `json:"list_price"`
	Metadata  string                                          `json:"metadata"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd time.Time `json:"netsuite_invoice_billing_end" format:"date-time"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart time.Time `json:"netsuite_invoice_billing_start" format:"date-time"`
	NetsuiteItemID              string    `json:"netsuite_item_id"`
	// Only present for line items paying for a postpaid commit true-up.
	PostpaidCommit V1CustomerInvoiceListResponseLineItemsPostpaidCommit `json:"postpaid_commit"`
	// Includes the presentation group values associated with this line item if
	// presentation group keys are used.
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	// Includes the pricing group values associated with this line item if dimensional
	// pricing is used.
	PricingGroupValues  map[string]string `json:"pricing_group_values"`
	ProductCustomFields map[string]string `json:"product_custom_fields"`
	// ID of the product associated with the line item.
	ProductID string `json:"product_id" format:"uuid"`
	// The current product tags associated with the line item's `product_id`.
	ProductTags []string `json:"product_tags"`
	// The type of the line item's product. Possible values are `FixedProductListItem`
	// (for `FIXED` type products), `UsageProductListItem` (for `USAGE` type products),
	// `SubscriptionProductListItem` (for `SUBSCRIPTION` type products) or
	// `CompositeProductListItem` (for `COMPOSITE` type products). For scheduled
	// charges, commit and credit payments, the value is `FixedProductListItem`.
	ProductType                     string            `json:"product_type"`
	ProfessionalServiceCustomFields map[string]string `json:"professional_service_custom_fields"`
	ProfessionalServiceID           string            `json:"professional_service_id" format:"uuid"`
	// The quantity associated with the line item.
	Quantity                    float64                                            `json:"quantity"`
	ResellerType                V1CustomerInvoiceListResponseLineItemsResellerType `json:"reseller_type"`
	ScheduledChargeCustomFields map[string]string                                  `json:"scheduled_charge_custom_fields"`
	// ID of scheduled charge.
	ScheduledChargeID string `json:"scheduled_charge_id" format:"uuid"`
	// The line item's start date (inclusive).
	StartingAt               time.Time                                           `json:"starting_at" format:"date-time"`
	SubLineItems             []V1CustomerInvoiceListResponseLineItemsSubLineItem `json:"sub_line_items"`
	SubscriptionCustomFields map[string]string                                   `json:"subscription_custom_fields"`
	// Populated if the line item has a tiered price.
	Tier V1CustomerInvoiceListResponseLineItemsTier `json:"tier"`
	// The unit price associated with the line item.
	UnitPrice float64                                   `json:"unit_price"`
	JSON      v1CustomerInvoiceListResponseLineItemJSON `json:"-"`
}

// v1CustomerInvoiceListResponseLineItemJSON contains the JSON metadata for the
// struct [V1CustomerInvoiceListResponseLineItem]
type v1CustomerInvoiceListResponseLineItemJSON struct {
	CreditType                      apijson.Field
	Name                            apijson.Field
	Total                           apijson.Field
	Type                            apijson.Field
	AppliedCommitOrCredit           apijson.Field
	CommitCustomFields              apijson.Field
	CommitID                        apijson.Field
	CommitNetsuiteItemID            apijson.Field
	CommitNetsuiteSalesOrderID      apijson.Field
	CommitSegmentID                 apijson.Field
	CommitType                      apijson.Field
	CustomFields                    apijson.Field
	DiscountCustomFields            apijson.Field
	DiscountID                      apijson.Field
	EndingBefore                    apijson.Field
	GroupKey                        apijson.Field
	GroupValue                      apijson.Field
	IsProrated                      apijson.Field
	ListPrice                       apijson.Field
	Metadata                        apijson.Field
	NetsuiteInvoiceBillingEnd       apijson.Field
	NetsuiteInvoiceBillingStart     apijson.Field
	NetsuiteItemID                  apijson.Field
	PostpaidCommit                  apijson.Field
	PresentationGroupValues         apijson.Field
	PricingGroupValues              apijson.Field
	ProductCustomFields             apijson.Field
	ProductID                       apijson.Field
	ProductTags                     apijson.Field
	ProductType                     apijson.Field
	ProfessionalServiceCustomFields apijson.Field
	ProfessionalServiceID           apijson.Field
	Quantity                        apijson.Field
	ResellerType                    apijson.Field
	ScheduledChargeCustomFields     apijson.Field
	ScheduledChargeID               apijson.Field
	StartingAt                      apijson.Field
	SubLineItems                    apijson.Field
	SubscriptionCustomFields        apijson.Field
	Tier                            apijson.Field
	UnitPrice                       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseLineItemJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseLineItemsCreditType struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	Name string                                               `json:"name,required"`
	JSON v1CustomerInvoiceListResponseLineItemsCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceListResponseLineItemsCreditTypeJSON contains the JSON metadata
// for the struct [V1CustomerInvoiceListResponseLineItemsCreditType]
type v1CustomerInvoiceListResponseLineItemsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseLineItemsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseLineItemsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// Details about the credit or commit that was applied to this line item. Only
// present on line items with product of `USAGE`, `SUBSCRIPTION` or `COMPOSITE`
// types.
type V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCredit struct {
	ID   string                                                          `json:"id,required" format:"uuid"`
	Type V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditType `json:"type,required"`
	JSON v1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditJSON `json:"-"`
}

// v1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditJSON contains the
// JSON metadata for the struct
// [V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCredit]
type v1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditType string

const (
	V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditTypePrepaid  V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditType = "PREPAID"
	V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditTypePostpaid V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditType = "POSTPAID"
	V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditTypeCredit   V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditType = "CREDIT"
)

func (r V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditTypePrepaid, V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditTypePostpaid, V1CustomerInvoiceListResponseLineItemsAppliedCommitOrCreditTypeCredit:
		return true
	}
	return false
}

// Only present for contract invoices and when the `include_list_prices` query
// parameter is set to true. This will include the list rate for the charge if
// applicable. Only present for usage and subscription line items.
type V1CustomerInvoiceListResponseLineItemsListPrice struct {
	RateType   V1CustomerInvoiceListResponseLineItemsListPriceRateType   `json:"rate_type,required"`
	CreditType V1CustomerInvoiceListResponseLineItemsListPriceCreditType `json:"credit_type"`
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
	Tiers []V1CustomerInvoiceListResponseLineItemsListPriceTier `json:"tiers"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool                                                `json:"use_list_prices"`
	JSON          v1CustomerInvoiceListResponseLineItemsListPriceJSON `json:"-"`
}

// v1CustomerInvoiceListResponseLineItemsListPriceJSON contains the JSON metadata
// for the struct [V1CustomerInvoiceListResponseLineItemsListPrice]
type v1CustomerInvoiceListResponseLineItemsListPriceJSON struct {
	RateType           apijson.Field
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

func (r *V1CustomerInvoiceListResponseLineItemsListPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseLineItemsListPriceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseLineItemsListPriceRateType string

const (
	V1CustomerInvoiceListResponseLineItemsListPriceRateTypeFlat         V1CustomerInvoiceListResponseLineItemsListPriceRateType = "FLAT"
	V1CustomerInvoiceListResponseLineItemsListPriceRateTypePercentage   V1CustomerInvoiceListResponseLineItemsListPriceRateType = "PERCENTAGE"
	V1CustomerInvoiceListResponseLineItemsListPriceRateTypeSubscription V1CustomerInvoiceListResponseLineItemsListPriceRateType = "SUBSCRIPTION"
	V1CustomerInvoiceListResponseLineItemsListPriceRateTypeCustom       V1CustomerInvoiceListResponseLineItemsListPriceRateType = "CUSTOM"
	V1CustomerInvoiceListResponseLineItemsListPriceRateTypeTiered       V1CustomerInvoiceListResponseLineItemsListPriceRateType = "TIERED"
)

func (r V1CustomerInvoiceListResponseLineItemsListPriceRateType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListResponseLineItemsListPriceRateTypeFlat, V1CustomerInvoiceListResponseLineItemsListPriceRateTypePercentage, V1CustomerInvoiceListResponseLineItemsListPriceRateTypeSubscription, V1CustomerInvoiceListResponseLineItemsListPriceRateTypeCustom, V1CustomerInvoiceListResponseLineItemsListPriceRateTypeTiered:
		return true
	}
	return false
}

type V1CustomerInvoiceListResponseLineItemsListPriceCreditType struct {
	ID   string                                                        `json:"id,required" format:"uuid"`
	Name string                                                        `json:"name,required"`
	JSON v1CustomerInvoiceListResponseLineItemsListPriceCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceListResponseLineItemsListPriceCreditTypeJSON contains the JSON
// metadata for the struct
// [V1CustomerInvoiceListResponseLineItemsListPriceCreditType]
type v1CustomerInvoiceListResponseLineItemsListPriceCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseLineItemsListPriceCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseLineItemsListPriceCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseLineItemsListPriceTier struct {
	Price float64                                                 `json:"price,required"`
	Size  float64                                                 `json:"size"`
	JSON  v1CustomerInvoiceListResponseLineItemsListPriceTierJSON `json:"-"`
}

// v1CustomerInvoiceListResponseLineItemsListPriceTierJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceListResponseLineItemsListPriceTier]
type v1CustomerInvoiceListResponseLineItemsListPriceTierJSON struct {
	Price       apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseLineItemsListPriceTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseLineItemsListPriceTierJSON) RawJSON() string {
	return r.raw
}

// Only present for line items paying for a postpaid commit true-up.
type V1CustomerInvoiceListResponseLineItemsPostpaidCommit struct {
	ID   string                                                   `json:"id,required" format:"uuid"`
	JSON v1CustomerInvoiceListResponseLineItemsPostpaidCommitJSON `json:"-"`
}

// v1CustomerInvoiceListResponseLineItemsPostpaidCommitJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceListResponseLineItemsPostpaidCommit]
type v1CustomerInvoiceListResponseLineItemsPostpaidCommitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseLineItemsPostpaidCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseLineItemsPostpaidCommitJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseLineItemsResellerType string

const (
	V1CustomerInvoiceListResponseLineItemsResellerTypeAws           V1CustomerInvoiceListResponseLineItemsResellerType = "AWS"
	V1CustomerInvoiceListResponseLineItemsResellerTypeAwsProService V1CustomerInvoiceListResponseLineItemsResellerType = "AWS_PRO_SERVICE"
	V1CustomerInvoiceListResponseLineItemsResellerTypeGcp           V1CustomerInvoiceListResponseLineItemsResellerType = "GCP"
	V1CustomerInvoiceListResponseLineItemsResellerTypeGcpProService V1CustomerInvoiceListResponseLineItemsResellerType = "GCP_PRO_SERVICE"
)

func (r V1CustomerInvoiceListResponseLineItemsResellerType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListResponseLineItemsResellerTypeAws, V1CustomerInvoiceListResponseLineItemsResellerTypeAwsProService, V1CustomerInvoiceListResponseLineItemsResellerTypeGcp, V1CustomerInvoiceListResponseLineItemsResellerTypeGcpProService:
		return true
	}
	return false
}

type V1CustomerInvoiceListResponseLineItemsSubLineItem struct {
	CustomFields  map[string]string `json:"custom_fields,required"`
	Name          string            `json:"name,required"`
	Quantity      float64           `json:"quantity,required"`
	Subtotal      float64           `json:"subtotal,required"`
	ChargeID      string            `json:"charge_id" format:"uuid"`
	CreditGrantID string            `json:"credit_grant_id" format:"uuid"`
	// The end date for the charge (for seats charges only).
	EndDate time.Time `json:"end_date" format:"date-time"`
	// the unit price for this charge, present only if the charge is not tiered and the
	// quantity is nonzero
	Price float64 `json:"price"`
	// The start date for the charge (for seats charges only).
	StartDate time.Time `json:"start_date" format:"date-time"`
	// when the current tier started and ends (for tiered charges only)
	TierPeriod V1CustomerInvoiceListResponseLineItemsSubLineItemsTierPeriod `json:"tier_period"`
	Tiers      []V1CustomerInvoiceListResponseLineItemsSubLineItemsTier     `json:"tiers"`
	JSON       v1CustomerInvoiceListResponseLineItemsSubLineItemJSON        `json:"-"`
}

// v1CustomerInvoiceListResponseLineItemsSubLineItemJSON contains the JSON metadata
// for the struct [V1CustomerInvoiceListResponseLineItemsSubLineItem]
type v1CustomerInvoiceListResponseLineItemsSubLineItemJSON struct {
	CustomFields  apijson.Field
	Name          apijson.Field
	Quantity      apijson.Field
	Subtotal      apijson.Field
	ChargeID      apijson.Field
	CreditGrantID apijson.Field
	EndDate       apijson.Field
	Price         apijson.Field
	StartDate     apijson.Field
	TierPeriod    apijson.Field
	Tiers         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseLineItemsSubLineItemJSON) RawJSON() string {
	return r.raw
}

// when the current tier started and ends (for tiered charges only)
type V1CustomerInvoiceListResponseLineItemsSubLineItemsTierPeriod struct {
	StartingAt   time.Time                                                        `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time                                                        `json:"ending_before" format:"date-time"`
	JSON         v1CustomerInvoiceListResponseLineItemsSubLineItemsTierPeriodJSON `json:"-"`
}

// v1CustomerInvoiceListResponseLineItemsSubLineItemsTierPeriodJSON contains the
// JSON metadata for the struct
// [V1CustomerInvoiceListResponseLineItemsSubLineItemsTierPeriod]
type v1CustomerInvoiceListResponseLineItemsSubLineItemsTierPeriodJSON struct {
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseLineItemsSubLineItemsTierPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseLineItemsSubLineItemsTierPeriodJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseLineItemsSubLineItemsTier struct {
	Price    float64 `json:"price,required"`
	Quantity float64 `json:"quantity,required"`
	// at what metric amount this tier begins
	StartingAt float64                                                    `json:"starting_at,required"`
	Subtotal   float64                                                    `json:"subtotal,required"`
	JSON       v1CustomerInvoiceListResponseLineItemsSubLineItemsTierJSON `json:"-"`
}

// v1CustomerInvoiceListResponseLineItemsSubLineItemsTierJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceListResponseLineItemsSubLineItemsTier]
type v1CustomerInvoiceListResponseLineItemsSubLineItemsTierJSON struct {
	Price       apijson.Field
	Quantity    apijson.Field
	StartingAt  apijson.Field
	Subtotal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseLineItemsSubLineItemsTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseLineItemsSubLineItemsTierJSON) RawJSON() string {
	return r.raw
}

// Populated if the line item has a tiered price.
type V1CustomerInvoiceListResponseLineItemsTier struct {
	Level      float64                                        `json:"level,required"`
	StartingAt string                                         `json:"starting_at,required"`
	Size       string                                         `json:"size,nullable"`
	JSON       v1CustomerInvoiceListResponseLineItemsTierJSON `json:"-"`
}

// v1CustomerInvoiceListResponseLineItemsTierJSON contains the JSON metadata for
// the struct [V1CustomerInvoiceListResponseLineItemsTier]
type v1CustomerInvoiceListResponseLineItemsTierJSON struct {
	Level       apijson.Field
	StartingAt  apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseLineItemsTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseLineItemsTierJSON) RawJSON() string {
	return r.raw
}

// This field's availability is dependent on your client's configuration.
type V1CustomerInvoiceListResponseBillableStatus string

const (
	V1CustomerInvoiceListResponseBillableStatusBillable   V1CustomerInvoiceListResponseBillableStatus = "billable"
	V1CustomerInvoiceListResponseBillableStatusUnbillable V1CustomerInvoiceListResponseBillableStatus = "unbillable"
)

func (r V1CustomerInvoiceListResponseBillableStatus) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListResponseBillableStatusBillable, V1CustomerInvoiceListResponseBillableStatusUnbillable:
		return true
	}
	return false
}

type V1CustomerInvoiceListResponseCorrectionRecord struct {
	CorrectedInvoiceID       string                                                                `json:"corrected_invoice_id,required" format:"uuid"`
	Memo                     string                                                                `json:"memo,required"`
	Reason                   string                                                                `json:"reason,required"`
	CorrectedExternalInvoice V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoice `json:"corrected_external_invoice"`
	JSON                     v1CustomerInvoiceListResponseCorrectionRecordJSON                     `json:"-"`
}

// v1CustomerInvoiceListResponseCorrectionRecordJSON contains the JSON metadata for
// the struct [V1CustomerInvoiceListResponseCorrectionRecord]
type v1CustomerInvoiceListResponseCorrectionRecordJSON struct {
	CorrectedInvoiceID       apijson.Field
	Memo                     apijson.Field
	Reason                   apijson.Field
	CorrectedExternalInvoice apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseCorrectionRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseCorrectionRecordJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoice struct {
	BillingProviderType V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                                                   `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                                                `json:"issued_at_timestamp" format:"date-time"`
	JSON                v1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceJSON                `json:"-"`
}

// v1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceJSON
// contains the JSON metadata for the struct
// [V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoice]
type v1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType string

const (
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace   V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "aws_marketplace"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe           V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "stripe"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite         V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "netsuite"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom           V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "custom"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "azure_marketplace"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "quickbooks_online"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday          V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "workday"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace   V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus string

const (
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft               V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DRAFT"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized           V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "FINALIZED"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid                V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAID"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible       V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid                V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "VOID"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted             V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DELETED"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed       V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped             V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SKIPPED"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusSent                V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SENT"
	V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued              V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "QUEUED"
)

func (r V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusSent, V1CustomerInvoiceListResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type V1CustomerInvoiceListResponseExternalInvoice struct {
	BillingProviderType V1CustomerInvoiceListResponseExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      V1CustomerInvoiceListResponseExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                          `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                       `json:"issued_at_timestamp" format:"date-time"`
	JSON                v1CustomerInvoiceListResponseExternalInvoiceJSON                `json:"-"`
}

// v1CustomerInvoiceListResponseExternalInvoiceJSON contains the JSON metadata for
// the struct [V1CustomerInvoiceListResponseExternalInvoice]
type v1CustomerInvoiceListResponseExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseExternalInvoiceBillingProviderType string

const (
	V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeAwsMarketplace   V1CustomerInvoiceListResponseExternalInvoiceBillingProviderType = "aws_marketplace"
	V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeStripe           V1CustomerInvoiceListResponseExternalInvoiceBillingProviderType = "stripe"
	V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeNetsuite         V1CustomerInvoiceListResponseExternalInvoiceBillingProviderType = "netsuite"
	V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeCustom           V1CustomerInvoiceListResponseExternalInvoiceBillingProviderType = "custom"
	V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeAzureMarketplace V1CustomerInvoiceListResponseExternalInvoiceBillingProviderType = "azure_marketplace"
	V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeQuickbooksOnline V1CustomerInvoiceListResponseExternalInvoiceBillingProviderType = "quickbooks_online"
	V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeWorkday          V1CustomerInvoiceListResponseExternalInvoiceBillingProviderType = "workday"
	V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeGcpMarketplace   V1CustomerInvoiceListResponseExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r V1CustomerInvoiceListResponseExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeAwsMarketplace, V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeStripe, V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeNetsuite, V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeCustom, V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeAzureMarketplace, V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeQuickbooksOnline, V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeWorkday, V1CustomerInvoiceListResponseExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type V1CustomerInvoiceListResponseExternalInvoiceExternalStatus string

const (
	V1CustomerInvoiceListResponseExternalInvoiceExternalStatusDraft               V1CustomerInvoiceListResponseExternalInvoiceExternalStatus = "DRAFT"
	V1CustomerInvoiceListResponseExternalInvoiceExternalStatusFinalized           V1CustomerInvoiceListResponseExternalInvoiceExternalStatus = "FINALIZED"
	V1CustomerInvoiceListResponseExternalInvoiceExternalStatusPaid                V1CustomerInvoiceListResponseExternalInvoiceExternalStatus = "PAID"
	V1CustomerInvoiceListResponseExternalInvoiceExternalStatusUncollectible       V1CustomerInvoiceListResponseExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	V1CustomerInvoiceListResponseExternalInvoiceExternalStatusVoid                V1CustomerInvoiceListResponseExternalInvoiceExternalStatus = "VOID"
	V1CustomerInvoiceListResponseExternalInvoiceExternalStatusDeleted             V1CustomerInvoiceListResponseExternalInvoiceExternalStatus = "DELETED"
	V1CustomerInvoiceListResponseExternalInvoiceExternalStatusPaymentFailed       V1CustomerInvoiceListResponseExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	V1CustomerInvoiceListResponseExternalInvoiceExternalStatusInvalidRequestError V1CustomerInvoiceListResponseExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	V1CustomerInvoiceListResponseExternalInvoiceExternalStatusSkipped             V1CustomerInvoiceListResponseExternalInvoiceExternalStatus = "SKIPPED"
	V1CustomerInvoiceListResponseExternalInvoiceExternalStatusSent                V1CustomerInvoiceListResponseExternalInvoiceExternalStatus = "SENT"
	V1CustomerInvoiceListResponseExternalInvoiceExternalStatusQueued              V1CustomerInvoiceListResponseExternalInvoiceExternalStatus = "QUEUED"
)

func (r V1CustomerInvoiceListResponseExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListResponseExternalInvoiceExternalStatusDraft, V1CustomerInvoiceListResponseExternalInvoiceExternalStatusFinalized, V1CustomerInvoiceListResponseExternalInvoiceExternalStatusPaid, V1CustomerInvoiceListResponseExternalInvoiceExternalStatusUncollectible, V1CustomerInvoiceListResponseExternalInvoiceExternalStatusVoid, V1CustomerInvoiceListResponseExternalInvoiceExternalStatusDeleted, V1CustomerInvoiceListResponseExternalInvoiceExternalStatusPaymentFailed, V1CustomerInvoiceListResponseExternalInvoiceExternalStatusInvalidRequestError, V1CustomerInvoiceListResponseExternalInvoiceExternalStatusSkipped, V1CustomerInvoiceListResponseExternalInvoiceExternalStatusSent, V1CustomerInvoiceListResponseExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type V1CustomerInvoiceListResponseInvoiceAdjustment struct {
	CreditType              V1CustomerInvoiceListResponseInvoiceAdjustmentsCreditType `json:"credit_type,required"`
	Name                    string                                                    `json:"name,required"`
	Total                   float64                                                   `json:"total,required"`
	CreditGrantCustomFields map[string]string                                         `json:"credit_grant_custom_fields"`
	CreditGrantID           string                                                    `json:"credit_grant_id"`
	JSON                    v1CustomerInvoiceListResponseInvoiceAdjustmentJSON        `json:"-"`
}

// v1CustomerInvoiceListResponseInvoiceAdjustmentJSON contains the JSON metadata
// for the struct [V1CustomerInvoiceListResponseInvoiceAdjustment]
type v1CustomerInvoiceListResponseInvoiceAdjustmentJSON struct {
	CreditType              apijson.Field
	Name                    apijson.Field
	Total                   apijson.Field
	CreditGrantCustomFields apijson.Field
	CreditGrantID           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseInvoiceAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseInvoiceAdjustmentJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseInvoiceAdjustmentsCreditType struct {
	ID   string                                                        `json:"id,required" format:"uuid"`
	Name string                                                        `json:"name,required"`
	JSON v1CustomerInvoiceListResponseInvoiceAdjustmentsCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceListResponseInvoiceAdjustmentsCreditTypeJSON contains the JSON
// metadata for the struct
// [V1CustomerInvoiceListResponseInvoiceAdjustmentsCreditType]
type v1CustomerInvoiceListResponseInvoiceAdjustmentsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseInvoiceAdjustmentsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseInvoiceAdjustmentsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// Only present for contract invoices with reseller royalties.
type V1CustomerInvoiceListResponseResellerRoyalty struct {
	Fraction           string                                                   `json:"fraction,required"`
	NetsuiteResellerID string                                                   `json:"netsuite_reseller_id,required"`
	ResellerType       V1CustomerInvoiceListResponseResellerRoyaltyResellerType `json:"reseller_type,required"`
	AwsOptions         V1CustomerInvoiceListResponseResellerRoyaltyAwsOptions   `json:"aws_options"`
	GcpOptions         V1CustomerInvoiceListResponseResellerRoyaltyGcpOptions   `json:"gcp_options"`
	JSON               v1CustomerInvoiceListResponseResellerRoyaltyJSON         `json:"-"`
}

// v1CustomerInvoiceListResponseResellerRoyaltyJSON contains the JSON metadata for
// the struct [V1CustomerInvoiceListResponseResellerRoyalty]
type v1CustomerInvoiceListResponseResellerRoyaltyJSON struct {
	Fraction           apijson.Field
	NetsuiteResellerID apijson.Field
	ResellerType       apijson.Field
	AwsOptions         apijson.Field
	GcpOptions         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseResellerRoyaltyResellerType string

const (
	V1CustomerInvoiceListResponseResellerRoyaltyResellerTypeAws           V1CustomerInvoiceListResponseResellerRoyaltyResellerType = "AWS"
	V1CustomerInvoiceListResponseResellerRoyaltyResellerTypeAwsProService V1CustomerInvoiceListResponseResellerRoyaltyResellerType = "AWS_PRO_SERVICE"
	V1CustomerInvoiceListResponseResellerRoyaltyResellerTypeGcp           V1CustomerInvoiceListResponseResellerRoyaltyResellerType = "GCP"
	V1CustomerInvoiceListResponseResellerRoyaltyResellerTypeGcpProService V1CustomerInvoiceListResponseResellerRoyaltyResellerType = "GCP_PRO_SERVICE"
)

func (r V1CustomerInvoiceListResponseResellerRoyaltyResellerType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListResponseResellerRoyaltyResellerTypeAws, V1CustomerInvoiceListResponseResellerRoyaltyResellerTypeAwsProService, V1CustomerInvoiceListResponseResellerRoyaltyResellerTypeGcp, V1CustomerInvoiceListResponseResellerRoyaltyResellerTypeGcpProService:
		return true
	}
	return false
}

type V1CustomerInvoiceListResponseResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    string                                                     `json:"aws_account_number"`
	AwsOfferID          string                                                     `json:"aws_offer_id"`
	AwsPayerReferenceID string                                                     `json:"aws_payer_reference_id"`
	JSON                v1CustomerInvoiceListResponseResellerRoyaltyAwsOptionsJSON `json:"-"`
}

// v1CustomerInvoiceListResponseResellerRoyaltyAwsOptionsJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceListResponseResellerRoyaltyAwsOptions]
type v1CustomerInvoiceListResponseResellerRoyaltyAwsOptionsJSON struct {
	AwsAccountNumber    apijson.Field
	AwsOfferID          apijson.Field
	AwsPayerReferenceID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseResellerRoyaltyAwsOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseResellerRoyaltyAwsOptionsJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListResponseResellerRoyaltyGcpOptions struct {
	GcpAccountID string                                                     `json:"gcp_account_id"`
	GcpOfferID   string                                                     `json:"gcp_offer_id"`
	JSON         v1CustomerInvoiceListResponseResellerRoyaltyGcpOptionsJSON `json:"-"`
}

// v1CustomerInvoiceListResponseResellerRoyaltyGcpOptionsJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceListResponseResellerRoyaltyGcpOptions]
type v1CustomerInvoiceListResponseResellerRoyaltyGcpOptionsJSON struct {
	GcpAccountID apijson.Field
	GcpOfferID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerInvoiceListResponseResellerRoyaltyGcpOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListResponseResellerRoyaltyGcpOptionsJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceAddChargeResponse struct {
	JSON v1CustomerInvoiceAddChargeResponseJSON `json:"-"`
}

// v1CustomerInvoiceAddChargeResponseJSON contains the JSON metadata for the struct
// [V1CustomerInvoiceAddChargeResponse]
type v1CustomerInvoiceAddChargeResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceAddChargeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceAddChargeResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponse struct {
	ID                      string                                            `json:"id,required" format:"uuid"`
	BreakdownEndTimestamp   time.Time                                         `json:"breakdown_end_timestamp,required" format:"date-time"`
	BreakdownStartTimestamp time.Time                                         `json:"breakdown_start_timestamp,required" format:"date-time"`
	CreditType              V1CustomerInvoiceListBreakdownsResponseCreditType `json:"credit_type,required"`
	CustomerID              string                                            `json:"customer_id,required" format:"uuid"`
	LineItems               []V1CustomerInvoiceListBreakdownsResponseLineItem `json:"line_items,required"`
	Status                  string                                            `json:"status,required"`
	Total                   float64                                           `json:"total,required"`
	Type                    string                                            `json:"type,required"`
	AmendmentID             string                                            `json:"amendment_id" format:"uuid"`
	// This field's availability is dependent on your client's configuration.
	BillableStatus       V1CustomerInvoiceListBreakdownsResponseBillableStatus   `json:"billable_status"`
	ContractCustomFields map[string]string                                       `json:"contract_custom_fields"`
	ContractID           string                                                  `json:"contract_id" format:"uuid"`
	CorrectionRecord     V1CustomerInvoiceListBreakdownsResponseCorrectionRecord `json:"correction_record"`
	// When the invoice was created (UTC). This field is present for correction
	// invoices only.
	CreatedAt            time.Time              `json:"created_at" format:"date-time"`
	CustomFields         map[string]interface{} `json:"custom_fields"`
	CustomerCustomFields map[string]string      `json:"customer_custom_fields"`
	// End of the usage period this invoice covers (UTC)
	EndTimestamp       time.Time                                                  `json:"end_timestamp" format:"date-time"`
	ExternalInvoice    V1CustomerInvoiceListBreakdownsResponseExternalInvoice     `json:"external_invoice,nullable"`
	InvoiceAdjustments []V1CustomerInvoiceListBreakdownsResponseInvoiceAdjustment `json:"invoice_adjustments"`
	// When the invoice was issued (UTC)
	IssuedAt            time.Time `json:"issued_at" format:"date-time"`
	NetPaymentTermsDays float64   `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string            `json:"netsuite_sales_order_id"`
	PlanCustomFields     map[string]string `json:"plan_custom_fields"`
	PlanID               string            `json:"plan_id" format:"uuid"`
	PlanName             string            `json:"plan_name"`
	// Only present for contract invoices with reseller royalties.
	ResellerRoyalty V1CustomerInvoiceListBreakdownsResponseResellerRoyalty `json:"reseller_royalty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Beginning of the usage period this invoice covers (UTC)
	StartTimestamp time.Time                                   `json:"start_timestamp" format:"date-time"`
	Subtotal       float64                                     `json:"subtotal"`
	JSON           v1CustomerInvoiceListBreakdownsResponseJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseJSON contains the JSON metadata for the
// struct [V1CustomerInvoiceListBreakdownsResponse]
type v1CustomerInvoiceListBreakdownsResponseJSON struct {
	ID                      apijson.Field
	BreakdownEndTimestamp   apijson.Field
	BreakdownStartTimestamp apijson.Field
	CreditType              apijson.Field
	CustomerID              apijson.Field
	LineItems               apijson.Field
	Status                  apijson.Field
	Total                   apijson.Field
	Type                    apijson.Field
	AmendmentID             apijson.Field
	BillableStatus          apijson.Field
	ContractCustomFields    apijson.Field
	ContractID              apijson.Field
	CorrectionRecord        apijson.Field
	CreatedAt               apijson.Field
	CustomFields            apijson.Field
	CustomerCustomFields    apijson.Field
	EndTimestamp            apijson.Field
	ExternalInvoice         apijson.Field
	InvoiceAdjustments      apijson.Field
	IssuedAt                apijson.Field
	NetPaymentTermsDays     apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	PlanCustomFields        apijson.Field
	PlanID                  apijson.Field
	PlanName                apijson.Field
	ResellerRoyalty         apijson.Field
	SalesforceOpportunityID apijson.Field
	StartTimestamp          apijson.Field
	Subtotal                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseCreditType struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	Name string                                                `json:"name,required"`
	JSON v1CustomerInvoiceListBreakdownsResponseCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseCreditTypeJSON contains the JSON metadata
// for the struct [V1CustomerInvoiceListBreakdownsResponseCreditType]
type v1CustomerInvoiceListBreakdownsResponseCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseLineItem struct {
	CreditType V1CustomerInvoiceListBreakdownsResponseLineItemsCreditType `json:"credit_type,required"`
	Name       string                                                     `json:"name,required"`
	Total      float64                                                    `json:"total,required"`
	// The type of line item. scheduled - Line item is associated with a scheduled
	// charge. View the scheduled_charge_id on the line item. commit_purchase - Line
	// item is associated with a payment for a prepaid commit. View the commit_id on
	// the line item. usage - Line item is associated with a usage product or composite
	// product. View the product_id on the line item to determine which product.
	// subscription - Line item is associated with a subscription. e.g. monthly
	// recurring payment for an in-advance subscription. applied_commit_or_credit - On
	// metronome invoices, applied commits and credits are associated with their own
	// line items. These line items have negative totals. Use the
	// applied_commit_or_credit object on the line item to understand the id of the
	// applied commit or credit, and its type. Note that the application of a postpaid
	// commit is associated with a line item, but the total on the line item is not
	// included in the invoice's total as postpaid commits are paid in-arrears.
	// postpaid_trueup - Line item is associated with the true up amount for a postpaid
	// commit. This line item type will only appear on invoices with type TRUEUP .
	// cpu_conversion - Line item converting between a custom pricing unit and fiat
	// currency, using the conversion rate set on the rate card. This line item will
	// appear when there are products priced in custom pricing units, and there is
	// insufficient prepaid commit/credit in that custom pricing unit to fully cover
	// the spend. Then, the outstanding spend in custom pricing units will be converted
	// to fiat currency using a cpu_conversion line item.
	Type string `json:"type,required"`
	// Details about the credit or commit that was applied to this line item. Only
	// present on line items with product of `USAGE`, `SUBSCRIPTION` or `COMPOSITE`
	// types.
	AppliedCommitOrCredit V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCredit `json:"applied_commit_or_credit"`
	CommitCustomFields    map[string]string                                                     `json:"commit_custom_fields"`
	// For line items with product of `USAGE`, `SUBSCRIPTION`, or `COMPOSITE` types,
	// the ID of the credit or commit that was applied to this line item. For line
	// items with product type of `FIXED`, the ID of the prepaid or postpaid commit
	// that is being paid for.
	CommitID                   string `json:"commit_id" format:"uuid"`
	CommitNetsuiteItemID       string `json:"commit_netsuite_item_id"`
	CommitNetsuiteSalesOrderID string `json:"commit_netsuite_sales_order_id"`
	CommitSegmentID            string `json:"commit_segment_id" format:"uuid"`
	// `PrepaidCommit` (for commit types `PREPAID` and `CREDIT`) or `PostpaidCommit`
	// (for commit type `POSTPAID`).
	CommitType           string            `json:"commit_type"`
	CustomFields         map[string]string `json:"custom_fields"`
	DiscountCustomFields map[string]string `json:"discount_custom_fields"`
	// ID of the discount applied to this line item.
	DiscountID string `json:"discount_id" format:"uuid"`
	// The line item's end date (exclusive).
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	GroupKey     string    `json:"group_key"`
	GroupValue   string    `json:"group_value,nullable"`
	// Indicates whether the line item is prorated for `SUBSCRIPTION` type product.
	IsProrated bool `json:"is_prorated"`
	// Only present for contract invoices and when the `include_list_prices` query
	// parameter is set to true. This will include the list rate for the charge if
	// applicable. Only present for usage and subscription line items.
	ListPrice V1CustomerInvoiceListBreakdownsResponseLineItemsListPrice `json:"list_price"`
	Metadata  string                                                    `json:"metadata"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd time.Time `json:"netsuite_invoice_billing_end" format:"date-time"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart time.Time `json:"netsuite_invoice_billing_start" format:"date-time"`
	NetsuiteItemID              string    `json:"netsuite_item_id"`
	// Only present for line items paying for a postpaid commit true-up.
	PostpaidCommit V1CustomerInvoiceListBreakdownsResponseLineItemsPostpaidCommit `json:"postpaid_commit"`
	// Includes the presentation group values associated with this line item if
	// presentation group keys are used.
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	// Includes the pricing group values associated with this line item if dimensional
	// pricing is used.
	PricingGroupValues  map[string]string `json:"pricing_group_values"`
	ProductCustomFields map[string]string `json:"product_custom_fields"`
	// ID of the product associated with the line item.
	ProductID string `json:"product_id" format:"uuid"`
	// The current product tags associated with the line item's `product_id`.
	ProductTags []string `json:"product_tags"`
	// The type of the line item's product. Possible values are `FixedProductListItem`
	// (for `FIXED` type products), `UsageProductListItem` (for `USAGE` type products),
	// `SubscriptionProductListItem` (for `SUBSCRIPTION` type products) or
	// `CompositeProductListItem` (for `COMPOSITE` type products). For scheduled
	// charges, commit and credit payments, the value is `FixedProductListItem`.
	ProductType                     string            `json:"product_type"`
	ProfessionalServiceCustomFields map[string]string `json:"professional_service_custom_fields"`
	ProfessionalServiceID           string            `json:"professional_service_id" format:"uuid"`
	// The quantity associated with the line item.
	Quantity                    float64                                                      `json:"quantity"`
	ResellerType                V1CustomerInvoiceListBreakdownsResponseLineItemsResellerType `json:"reseller_type"`
	ScheduledChargeCustomFields map[string]string                                            `json:"scheduled_charge_custom_fields"`
	// ID of scheduled charge.
	ScheduledChargeID string `json:"scheduled_charge_id" format:"uuid"`
	// The line item's start date (inclusive).
	StartingAt               time.Time                                                     `json:"starting_at" format:"date-time"`
	SubLineItems             []V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItem `json:"sub_line_items"`
	SubscriptionCustomFields map[string]string                                             `json:"subscription_custom_fields"`
	// Populated if the line item has a tiered price.
	Tier V1CustomerInvoiceListBreakdownsResponseLineItemsTier `json:"tier"`
	// The unit price associated with the line item.
	UnitPrice float64                                             `json:"unit_price"`
	JSON      v1CustomerInvoiceListBreakdownsResponseLineItemJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseLineItemJSON contains the JSON metadata
// for the struct [V1CustomerInvoiceListBreakdownsResponseLineItem]
type v1CustomerInvoiceListBreakdownsResponseLineItemJSON struct {
	CreditType                      apijson.Field
	Name                            apijson.Field
	Total                           apijson.Field
	Type                            apijson.Field
	AppliedCommitOrCredit           apijson.Field
	CommitCustomFields              apijson.Field
	CommitID                        apijson.Field
	CommitNetsuiteItemID            apijson.Field
	CommitNetsuiteSalesOrderID      apijson.Field
	CommitSegmentID                 apijson.Field
	CommitType                      apijson.Field
	CustomFields                    apijson.Field
	DiscountCustomFields            apijson.Field
	DiscountID                      apijson.Field
	EndingBefore                    apijson.Field
	GroupKey                        apijson.Field
	GroupValue                      apijson.Field
	IsProrated                      apijson.Field
	ListPrice                       apijson.Field
	Metadata                        apijson.Field
	NetsuiteInvoiceBillingEnd       apijson.Field
	NetsuiteInvoiceBillingStart     apijson.Field
	NetsuiteItemID                  apijson.Field
	PostpaidCommit                  apijson.Field
	PresentationGroupValues         apijson.Field
	PricingGroupValues              apijson.Field
	ProductCustomFields             apijson.Field
	ProductID                       apijson.Field
	ProductTags                     apijson.Field
	ProductType                     apijson.Field
	ProfessionalServiceCustomFields apijson.Field
	ProfessionalServiceID           apijson.Field
	Quantity                        apijson.Field
	ResellerType                    apijson.Field
	ScheduledChargeCustomFields     apijson.Field
	ScheduledChargeID               apijson.Field
	StartingAt                      apijson.Field
	SubLineItems                    apijson.Field
	SubscriptionCustomFields        apijson.Field
	Tier                            apijson.Field
	UnitPrice                       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseLineItemJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseLineItemsCreditType struct {
	ID   string                                                         `json:"id,required" format:"uuid"`
	Name string                                                         `json:"name,required"`
	JSON v1CustomerInvoiceListBreakdownsResponseLineItemsCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseLineItemsCreditTypeJSON contains the JSON
// metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseLineItemsCreditType]
type v1CustomerInvoiceListBreakdownsResponseLineItemsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseLineItemsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseLineItemsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// Details about the credit or commit that was applied to this line item. Only
// present on line items with product of `USAGE`, `SUBSCRIPTION` or `COMPOSITE`
// types.
type V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCredit struct {
	ID   string                                                                    `json:"id,required" format:"uuid"`
	Type V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditType `json:"type,required"`
	JSON v1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditJSON
// contains the JSON metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCredit]
type v1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditType string

const (
	V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditTypePrepaid  V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditType = "PREPAID"
	V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditTypePostpaid V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditType = "POSTPAID"
	V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditTypeCredit   V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditType = "CREDIT"
)

func (r V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditTypePrepaid, V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditTypePostpaid, V1CustomerInvoiceListBreakdownsResponseLineItemsAppliedCommitOrCreditTypeCredit:
		return true
	}
	return false
}

// Only present for contract invoices and when the `include_list_prices` query
// parameter is set to true. This will include the list rate for the charge if
// applicable. Only present for usage and subscription line items.
type V1CustomerInvoiceListBreakdownsResponseLineItemsListPrice struct {
	RateType   V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateType   `json:"rate_type,required"`
	CreditType V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceCreditType `json:"credit_type"`
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
	Tiers []V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceTier `json:"tiers"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool                                                          `json:"use_list_prices"`
	JSON          v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceJSON contains the JSON
// metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseLineItemsListPrice]
type v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceJSON struct {
	RateType           apijson.Field
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

func (r *V1CustomerInvoiceListBreakdownsResponseLineItemsListPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateType string

const (
	V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateTypeFlat         V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateType = "FLAT"
	V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateTypePercentage   V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateType = "PERCENTAGE"
	V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateTypeSubscription V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateType = "SUBSCRIPTION"
	V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateTypeCustom       V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateType = "CUSTOM"
	V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateTypeTiered       V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateType = "TIERED"
)

func (r V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateTypeFlat, V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateTypePercentage, V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateTypeSubscription, V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateTypeCustom, V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceRateTypeTiered:
		return true
	}
	return false
}

type V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceCreditType struct {
	ID   string                                                                  `json:"id,required" format:"uuid"`
	Name string                                                                  `json:"name,required"`
	JSON v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceCreditTypeJSON contains
// the JSON metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceCreditType]
type v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceTier struct {
	Price float64                                                           `json:"price,required"`
	Size  float64                                                           `json:"size"`
	JSON  v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceTierJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceTierJSON contains the
// JSON metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceTier]
type v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceTierJSON struct {
	Price       apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseLineItemsListPriceTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseLineItemsListPriceTierJSON) RawJSON() string {
	return r.raw
}

// Only present for line items paying for a postpaid commit true-up.
type V1CustomerInvoiceListBreakdownsResponseLineItemsPostpaidCommit struct {
	ID   string                                                             `json:"id,required" format:"uuid"`
	JSON v1CustomerInvoiceListBreakdownsResponseLineItemsPostpaidCommitJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseLineItemsPostpaidCommitJSON contains the
// JSON metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseLineItemsPostpaidCommit]
type v1CustomerInvoiceListBreakdownsResponseLineItemsPostpaidCommitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseLineItemsPostpaidCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseLineItemsPostpaidCommitJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseLineItemsResellerType string

const (
	V1CustomerInvoiceListBreakdownsResponseLineItemsResellerTypeAws           V1CustomerInvoiceListBreakdownsResponseLineItemsResellerType = "AWS"
	V1CustomerInvoiceListBreakdownsResponseLineItemsResellerTypeAwsProService V1CustomerInvoiceListBreakdownsResponseLineItemsResellerType = "AWS_PRO_SERVICE"
	V1CustomerInvoiceListBreakdownsResponseLineItemsResellerTypeGcp           V1CustomerInvoiceListBreakdownsResponseLineItemsResellerType = "GCP"
	V1CustomerInvoiceListBreakdownsResponseLineItemsResellerTypeGcpProService V1CustomerInvoiceListBreakdownsResponseLineItemsResellerType = "GCP_PRO_SERVICE"
)

func (r V1CustomerInvoiceListBreakdownsResponseLineItemsResellerType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListBreakdownsResponseLineItemsResellerTypeAws, V1CustomerInvoiceListBreakdownsResponseLineItemsResellerTypeAwsProService, V1CustomerInvoiceListBreakdownsResponseLineItemsResellerTypeGcp, V1CustomerInvoiceListBreakdownsResponseLineItemsResellerTypeGcpProService:
		return true
	}
	return false
}

type V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItem struct {
	CustomFields  map[string]string `json:"custom_fields,required"`
	Name          string            `json:"name,required"`
	Quantity      float64           `json:"quantity,required"`
	Subtotal      float64           `json:"subtotal,required"`
	ChargeID      string            `json:"charge_id" format:"uuid"`
	CreditGrantID string            `json:"credit_grant_id" format:"uuid"`
	// The end date for the charge (for seats charges only).
	EndDate time.Time `json:"end_date" format:"date-time"`
	// the unit price for this charge, present only if the charge is not tiered and the
	// quantity is nonzero
	Price float64 `json:"price"`
	// The start date for the charge (for seats charges only).
	StartDate time.Time `json:"start_date" format:"date-time"`
	// when the current tier started and ends (for tiered charges only)
	TierPeriod V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierPeriod `json:"tier_period"`
	Tiers      []V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTier     `json:"tiers"`
	JSON       v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemJSON        `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemJSON contains the
// JSON metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItem]
type v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemJSON struct {
	CustomFields  apijson.Field
	Name          apijson.Field
	Quantity      apijson.Field
	Subtotal      apijson.Field
	ChargeID      apijson.Field
	CreditGrantID apijson.Field
	EndDate       apijson.Field
	Price         apijson.Field
	StartDate     apijson.Field
	TierPeriod    apijson.Field
	Tiers         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemJSON) RawJSON() string {
	return r.raw
}

// when the current tier started and ends (for tiered charges only)
type V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierPeriod struct {
	StartingAt   time.Time                                                                  `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time                                                                  `json:"ending_before" format:"date-time"`
	JSON         v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierPeriodJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierPeriodJSON
// contains the JSON metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierPeriod]
type v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierPeriodJSON struct {
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierPeriodJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTier struct {
	Price    float64 `json:"price,required"`
	Quantity float64 `json:"quantity,required"`
	// at what metric amount this tier begins
	StartingAt float64                                                              `json:"starting_at,required"`
	Subtotal   float64                                                              `json:"subtotal,required"`
	JSON       v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierJSON contains
// the JSON metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTier]
type v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierJSON struct {
	Price       apijson.Field
	Quantity    apijson.Field
	StartingAt  apijson.Field
	Subtotal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseLineItemsSubLineItemsTierJSON) RawJSON() string {
	return r.raw
}

// Populated if the line item has a tiered price.
type V1CustomerInvoiceListBreakdownsResponseLineItemsTier struct {
	Level      float64                                                  `json:"level,required"`
	StartingAt string                                                   `json:"starting_at,required"`
	Size       string                                                   `json:"size,nullable"`
	JSON       v1CustomerInvoiceListBreakdownsResponseLineItemsTierJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseLineItemsTierJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceListBreakdownsResponseLineItemsTier]
type v1CustomerInvoiceListBreakdownsResponseLineItemsTierJSON struct {
	Level       apijson.Field
	StartingAt  apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseLineItemsTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseLineItemsTierJSON) RawJSON() string {
	return r.raw
}

// This field's availability is dependent on your client's configuration.
type V1CustomerInvoiceListBreakdownsResponseBillableStatus string

const (
	V1CustomerInvoiceListBreakdownsResponseBillableStatusBillable   V1CustomerInvoiceListBreakdownsResponseBillableStatus = "billable"
	V1CustomerInvoiceListBreakdownsResponseBillableStatusUnbillable V1CustomerInvoiceListBreakdownsResponseBillableStatus = "unbillable"
)

func (r V1CustomerInvoiceListBreakdownsResponseBillableStatus) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListBreakdownsResponseBillableStatusBillable, V1CustomerInvoiceListBreakdownsResponseBillableStatusUnbillable:
		return true
	}
	return false
}

type V1CustomerInvoiceListBreakdownsResponseCorrectionRecord struct {
	CorrectedInvoiceID       string                                                                          `json:"corrected_invoice_id,required" format:"uuid"`
	Memo                     string                                                                          `json:"memo,required"`
	Reason                   string                                                                          `json:"reason,required"`
	CorrectedExternalInvoice V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoice `json:"corrected_external_invoice"`
	JSON                     v1CustomerInvoiceListBreakdownsResponseCorrectionRecordJSON                     `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseCorrectionRecordJSON contains the JSON
// metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseCorrectionRecord]
type v1CustomerInvoiceListBreakdownsResponseCorrectionRecordJSON struct {
	CorrectedInvoiceID       apijson.Field
	Memo                     apijson.Field
	Reason                   apijson.Field
	CorrectedExternalInvoice apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseCorrectionRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseCorrectionRecordJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoice struct {
	BillingProviderType V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                                                             `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                                                          `json:"issued_at_timestamp" format:"date-time"`
	JSON                v1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceJSON                `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceJSON
// contains the JSON metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoice]
type v1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType string

const (
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace   V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "aws_marketplace"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe           V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "stripe"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite         V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "netsuite"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom           V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "custom"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "azure_marketplace"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "quickbooks_online"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday          V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "workday"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace   V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus string

const (
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft               V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DRAFT"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized           V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "FINALIZED"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid                V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAID"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible       V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid                V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "VOID"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted             V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DELETED"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed       V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped             V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SKIPPED"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusSent                V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SENT"
	V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued              V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus = "QUEUED"
)

func (r V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusSent, V1CustomerInvoiceListBreakdownsResponseCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type V1CustomerInvoiceListBreakdownsResponseExternalInvoice struct {
	BillingProviderType V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                                    `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                                 `json:"issued_at_timestamp" format:"date-time"`
	JSON                v1CustomerInvoiceListBreakdownsResponseExternalInvoiceJSON                `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseExternalInvoiceJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceListBreakdownsResponseExternalInvoice]
type v1CustomerInvoiceListBreakdownsResponseExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderType string

const (
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeAwsMarketplace   V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderType = "aws_marketplace"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeStripe           V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderType = "stripe"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeNetsuite         V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderType = "netsuite"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeCustom           V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderType = "custom"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeAzureMarketplace V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderType = "azure_marketplace"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeQuickbooksOnline V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderType = "quickbooks_online"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeWorkday          V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderType = "workday"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeGcpMarketplace   V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeAwsMarketplace, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeStripe, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeNetsuite, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeCustom, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeAzureMarketplace, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeQuickbooksOnline, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeWorkday, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus string

const (
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusDraft               V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus = "DRAFT"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusFinalized           V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus = "FINALIZED"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusPaid                V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus = "PAID"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusUncollectible       V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusVoid                V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus = "VOID"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusDeleted             V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus = "DELETED"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusPaymentFailed       V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusInvalidRequestError V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusSkipped             V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus = "SKIPPED"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusSent                V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus = "SENT"
	V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusQueued              V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus = "QUEUED"
)

func (r V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusDraft, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusFinalized, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusPaid, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusUncollectible, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusVoid, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusDeleted, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusPaymentFailed, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusInvalidRequestError, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusSkipped, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusSent, V1CustomerInvoiceListBreakdownsResponseExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type V1CustomerInvoiceListBreakdownsResponseInvoiceAdjustment struct {
	CreditType              V1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentsCreditType `json:"credit_type,required"`
	Name                    string                                                              `json:"name,required"`
	Total                   float64                                                             `json:"total,required"`
	CreditGrantCustomFields map[string]string                                                   `json:"credit_grant_custom_fields"`
	CreditGrantID           string                                                              `json:"credit_grant_id"`
	JSON                    v1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentJSON        `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentJSON contains the JSON
// metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseInvoiceAdjustment]
type v1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentJSON struct {
	CreditType              apijson.Field
	Name                    apijson.Field
	Total                   apijson.Field
	CreditGrantCustomFields apijson.Field
	CreditGrantID           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseInvoiceAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentsCreditType struct {
	ID   string                                                                  `json:"id,required" format:"uuid"`
	Name string                                                                  `json:"name,required"`
	JSON v1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentsCreditTypeJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentsCreditTypeJSON contains
// the JSON metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentsCreditType]
type v1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseInvoiceAdjustmentsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// Only present for contract invoices with reseller royalties.
type V1CustomerInvoiceListBreakdownsResponseResellerRoyalty struct {
	Fraction           string                                                             `json:"fraction,required"`
	NetsuiteResellerID string                                                             `json:"netsuite_reseller_id,required"`
	ResellerType       V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerType `json:"reseller_type,required"`
	AwsOptions         V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyAwsOptions   `json:"aws_options"`
	GcpOptions         V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyGcpOptions   `json:"gcp_options"`
	JSON               v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyJSON         `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyJSON contains the JSON
// metadata for the struct [V1CustomerInvoiceListBreakdownsResponseResellerRoyalty]
type v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyJSON struct {
	Fraction           apijson.Field
	NetsuiteResellerID apijson.Field
	ResellerType       apijson.Field
	AwsOptions         apijson.Field
	GcpOptions         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerType string

const (
	V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerTypeAws           V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerType = "AWS"
	V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerTypeAwsProService V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerType = "AWS_PRO_SERVICE"
	V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerTypeGcp           V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerType = "GCP"
	V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerTypeGcpProService V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerType = "GCP_PRO_SERVICE"
)

func (r V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerType) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerTypeAws, V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerTypeAwsProService, V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerTypeGcp, V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyResellerTypeGcpProService:
		return true
	}
	return false
}

type V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    string                                                               `json:"aws_account_number"`
	AwsOfferID          string                                                               `json:"aws_offer_id"`
	AwsPayerReferenceID string                                                               `json:"aws_payer_reference_id"`
	JSON                v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyAwsOptionsJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyAwsOptionsJSON contains
// the JSON metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyAwsOptions]
type v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyAwsOptionsJSON struct {
	AwsAccountNumber    apijson.Field
	AwsOfferID          apijson.Field
	AwsPayerReferenceID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyAwsOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyAwsOptionsJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyGcpOptions struct {
	GcpAccountID string                                                               `json:"gcp_account_id"`
	GcpOfferID   string                                                               `json:"gcp_offer_id"`
	JSON         v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyGcpOptionsJSON `json:"-"`
}

// v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyGcpOptionsJSON contains
// the JSON metadata for the struct
// [V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyGcpOptions]
type v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyGcpOptionsJSON struct {
	GcpAccountID apijson.Field
	GcpOfferID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerInvoiceListBreakdownsResponseResellerRoyaltyGcpOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerInvoiceListBreakdownsResponseResellerRoyaltyGcpOptionsJSON) RawJSON() string {
	return r.raw
}

type V1CustomerInvoiceGetParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	InvoiceID  param.Field[string] `path:"invoice_id,required" format:"uuid"`
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems param.Field[bool] `query:"skip_zero_qty_line_items"`
}

// URLQuery serializes [V1CustomerInvoiceGetParams]'s query parameters as
// `url.Values`.
func (r V1CustomerInvoiceGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerInvoiceListParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// Only return invoices for the specified credit type
	CreditTypeID param.Field[string] `query:"credit_type_id"`
	// RFC 3339 timestamp (exclusive). Invoices will only be returned for billing
	// periods that end before this time.
	EndingBefore param.Field[time.Time] `query:"ending_before" format:"date-time"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems param.Field[bool] `query:"skip_zero_qty_line_items"`
	// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
	// date_asc.
	Sort param.Field[V1CustomerInvoiceListParamsSort] `query:"sort"`
	// RFC 3339 timestamp (inclusive). Invoices will only be returned for billing
	// periods that start at or after this time.
	StartingOn param.Field[time.Time] `query:"starting_on" format:"date-time"`
	// Invoice status, e.g. DRAFT, FINALIZED, or VOID
	Status param.Field[string] `query:"status"`
}

// URLQuery serializes [V1CustomerInvoiceListParams]'s query parameters as
// `url.Values`.
func (r V1CustomerInvoiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
// date_asc.
type V1CustomerInvoiceListParamsSort string

const (
	V1CustomerInvoiceListParamsSortDateAsc  V1CustomerInvoiceListParamsSort = "date_asc"
	V1CustomerInvoiceListParamsSortDateDesc V1CustomerInvoiceListParamsSort = "date_desc"
)

func (r V1CustomerInvoiceListParamsSort) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListParamsSortDateAsc, V1CustomerInvoiceListParamsSortDateDesc:
		return true
	}
	return false
}

type V1CustomerInvoiceAddChargeParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// The Metronome ID of the charge to add to the invoice. Note that the charge must
	// be on a product that is not on the current plan, and the product must have only
	// fixed charges.
	ChargeID param.Field[string] `json:"charge_id,required" format:"uuid"`
	// The Metronome ID of the customer plan to add the charge to.
	CustomerPlanID param.Field[string] `json:"customer_plan_id,required" format:"uuid"`
	Description    param.Field[string] `json:"description,required"`
	// The start_timestamp of the invoice to add the charge to.
	InvoiceStartTimestamp param.Field[time.Time] `json:"invoice_start_timestamp,required" format:"date-time"`
	// The price of the charge. This price will match the currency on the invoice, e.g.
	// USD cents.
	Price    param.Field[float64] `json:"price,required"`
	Quantity param.Field[float64] `json:"quantity,required"`
}

func (r V1CustomerInvoiceAddChargeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerInvoiceListBreakdownsParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// RFC 3339 timestamp. Breakdowns will only be returned for time windows that end
	// on or before this time.
	EndingBefore param.Field[time.Time] `query:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp. Breakdowns will only be returned for time windows that start
	// on or after this time.
	StartingOn param.Field[time.Time] `query:"starting_on,required" format:"date-time"`
	// Only return invoices for the specified credit type
	CreditTypeID param.Field[string] `query:"credit_type_id"`
	// Max number of results that should be returned. For daily breakdowns, the
	// response can return up to 35 days worth of breakdowns. For hourly breakdowns,
	// the response can return up to 24 hours. If there are more results, a cursor to
	// the next page is returned.
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems param.Field[bool] `query:"skip_zero_qty_line_items"`
	// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
	// date_asc.
	Sort param.Field[V1CustomerInvoiceListBreakdownsParamsSort] `query:"sort"`
	// Invoice status, e.g. DRAFT or FINALIZED
	Status param.Field[string] `query:"status"`
	// The granularity of the breakdowns to return. Defaults to day.
	WindowSize param.Field[V1CustomerInvoiceListBreakdownsParamsWindowSize] `query:"window_size"`
}

// URLQuery serializes [V1CustomerInvoiceListBreakdownsParams]'s query parameters
// as `url.Values`.
func (r V1CustomerInvoiceListBreakdownsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
// date_asc.
type V1CustomerInvoiceListBreakdownsParamsSort string

const (
	V1CustomerInvoiceListBreakdownsParamsSortDateAsc  V1CustomerInvoiceListBreakdownsParamsSort = "date_asc"
	V1CustomerInvoiceListBreakdownsParamsSortDateDesc V1CustomerInvoiceListBreakdownsParamsSort = "date_desc"
)

func (r V1CustomerInvoiceListBreakdownsParamsSort) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListBreakdownsParamsSortDateAsc, V1CustomerInvoiceListBreakdownsParamsSortDateDesc:
		return true
	}
	return false
}

// The granularity of the breakdowns to return. Defaults to day.
type V1CustomerInvoiceListBreakdownsParamsWindowSize string

const (
	V1CustomerInvoiceListBreakdownsParamsWindowSizeHour V1CustomerInvoiceListBreakdownsParamsWindowSize = "HOUR"
	V1CustomerInvoiceListBreakdownsParamsWindowSizeDay  V1CustomerInvoiceListBreakdownsParamsWindowSize = "DAY"
)

func (r V1CustomerInvoiceListBreakdownsParamsWindowSize) IsKnown() bool {
	switch r {
	case V1CustomerInvoiceListBreakdownsParamsWindowSizeHour, V1CustomerInvoiceListBreakdownsParamsWindowSizeDay:
		return true
	}
	return false
}
