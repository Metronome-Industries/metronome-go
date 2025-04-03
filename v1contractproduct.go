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

// V1ContractProductService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContractProductService] method instead.
type V1ContractProductService struct {
	Options []option.RequestOption
}

// NewV1ContractProductService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1ContractProductService(opts ...option.RequestOption) (r *V1ContractProductService) {
	r = &V1ContractProductService{}
	r.Options = opts
	return
}

// Create a new product
func (r *V1ContractProductService) New(ctx context.Context, body V1ContractProductNewParams, opts ...option.RequestOption) (res *V1ContractProductNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/products/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific product
func (r *V1ContractProductService) Get(ctx context.Context, body V1ContractProductGetParams, opts ...option.RequestOption) (res *V1ContractProductGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/products/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a product
func (r *V1ContractProductService) Update(ctx context.Context, body V1ContractProductUpdateParams, opts ...option.RequestOption) (res *V1ContractProductUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/products/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List products
func (r *V1ContractProductService) List(ctx context.Context, params V1ContractProductListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1ContractProductListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/contract-pricing/products/list"
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

// List products
func (r *V1ContractProductService) ListAutoPaging(ctx context.Context, params V1ContractProductListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1ContractProductListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Archive a product
func (r *V1ContractProductService) Archive(ctx context.Context, body V1ContractProductArchiveParams, opts ...option.RequestOption) (res *V1ContractProductArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/products/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProductListItemState struct {
	CreatedAt           time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy           string    `json:"created_by,required"`
	Name                string    `json:"name,required"`
	BillableMetricID    string    `json:"billable_metric_id"`
	CompositeProductIDs []string  `json:"composite_product_ids" format:"uuid"`
	CompositeTags       []string  `json:"composite_tags"`
	ExcludeFreeUsage    bool      `json:"exclude_free_usage"`
	// This field's availability is dependent on your client's configuration.
	IsRefundable bool `json:"is_refundable"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID string `json:"netsuite_internal_item_id"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID string `json:"netsuite_overage_item_id"`
	// For USAGE products only. Groups usage line items on invoices. The superset of
	// values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PresentationGroupKey []string `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole. The superset
	// of values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PricingGroupKey []string `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion QuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding QuantityRounding         `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                `json:"starting_at" format:"date-time"`
	Tags             []string                 `json:"tags"`
	JSON             productListItemStateJSON `json:"-"`
}

// productListItemStateJSON contains the JSON metadata for the struct
// [ProductListItemState]
type productListItemStateJSON struct {
	CreatedAt              apijson.Field
	CreatedBy              apijson.Field
	Name                   apijson.Field
	BillableMetricID       apijson.Field
	CompositeProductIDs    apijson.Field
	CompositeTags          apijson.Field
	ExcludeFreeUsage       apijson.Field
	IsRefundable           apijson.Field
	NetsuiteInternalItemID apijson.Field
	NetsuiteOverageItemID  apijson.Field
	PresentationGroupKey   apijson.Field
	PricingGroupKey        apijson.Field
	QuantityConversion     apijson.Field
	QuantityRounding       apijson.Field
	StartingAt             apijson.Field
	Tags                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProductListItemState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListItemStateJSON) RawJSON() string {
	return r.raw
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type QuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation QuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                 `json:"name"`
	JSON quantityConversionJSON `json:"-"`
}

// quantityConversionJSON contains the JSON metadata for the struct
// [QuantityConversion]
type quantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *QuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type QuantityConversionOperation string

const (
	QuantityConversionOperationMultiply QuantityConversionOperation = "MULTIPLY"
	QuantityConversionOperationDivide   QuantityConversionOperation = "DIVIDE"
)

func (r QuantityConversionOperation) IsKnown() bool {
	switch r {
	case QuantityConversionOperationMultiply, QuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type QuantityConversionParam struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor param.Field[float64] `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation param.Field[QuantityConversionOperation] `json:"operation,required"`
	// Optional name for this conversion.
	Name param.Field[string] `json:"name"`
}

func (r QuantityConversionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type QuantityRounding struct {
	DecimalPlaces  float64                        `json:"decimal_places,required"`
	RoundingMethod QuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           quantityRoundingJSON           `json:"-"`
}

// quantityRoundingJSON contains the JSON metadata for the struct
// [QuantityRounding]
type quantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *QuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quantityRoundingJSON) RawJSON() string {
	return r.raw
}

type QuantityRoundingRoundingMethod string

const (
	QuantityRoundingRoundingMethodRoundUp     QuantityRoundingRoundingMethod = "ROUND_UP"
	QuantityRoundingRoundingMethodRoundDown   QuantityRoundingRoundingMethod = "ROUND_DOWN"
	QuantityRoundingRoundingMethodRoundHalfUp QuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r QuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case QuantityRoundingRoundingMethodRoundUp, QuantityRoundingRoundingMethodRoundDown, QuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type QuantityRoundingParam struct {
	DecimalPlaces  param.Field[float64]                        `json:"decimal_places,required"`
	RoundingMethod param.Field[QuantityRoundingRoundingMethod] `json:"rounding_method,required"`
}

func (r QuantityRoundingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractProductNewResponse struct {
	Data shared.ID                        `json:"data,required"`
	JSON v1ContractProductNewResponseJSON `json:"-"`
}

// v1ContractProductNewResponseJSON contains the JSON metadata for the struct
// [V1ContractProductNewResponse]
type v1ContractProductNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractProductNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductGetResponse struct {
	Data V1ContractProductGetResponseData `json:"data,required"`
	JSON v1ContractProductGetResponseJSON `json:"-"`
}

// v1ContractProductGetResponseJSON contains the JSON metadata for the struct
// [V1ContractProductGetResponse]
type v1ContractProductGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractProductGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductGetResponseData struct {
	ID           string                                   `json:"id,required" format:"uuid"`
	Current      ProductListItemState                     `json:"current,required"`
	Initial      ProductListItemState                     `json:"initial,required"`
	Type         V1ContractProductGetResponseDataType     `json:"type,required"`
	Updates      []V1ContractProductGetResponseDataUpdate `json:"updates,required"`
	ArchivedAt   time.Time                                `json:"archived_at,nullable" format:"date-time"`
	CustomFields map[string]string                        `json:"custom_fields"`
	JSON         v1ContractProductGetResponseDataJSON     `json:"-"`
}

// v1ContractProductGetResponseDataJSON contains the JSON metadata for the struct
// [V1ContractProductGetResponseData]
type v1ContractProductGetResponseDataJSON struct {
	ID           apijson.Field
	Current      apijson.Field
	Initial      apijson.Field
	Type         apijson.Field
	Updates      apijson.Field
	ArchivedAt   apijson.Field
	CustomFields apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1ContractProductGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductGetResponseDataType string

const (
	V1ContractProductGetResponseDataTypeUsage        V1ContractProductGetResponseDataType = "USAGE"
	V1ContractProductGetResponseDataTypeSubscription V1ContractProductGetResponseDataType = "SUBSCRIPTION"
	V1ContractProductGetResponseDataTypeComposite    V1ContractProductGetResponseDataType = "COMPOSITE"
	V1ContractProductGetResponseDataTypeFixed        V1ContractProductGetResponseDataType = "FIXED"
	V1ContractProductGetResponseDataTypeProService   V1ContractProductGetResponseDataType = "PRO_SERVICE"
)

func (r V1ContractProductGetResponseDataType) IsKnown() bool {
	switch r {
	case V1ContractProductGetResponseDataTypeUsage, V1ContractProductGetResponseDataTypeSubscription, V1ContractProductGetResponseDataTypeComposite, V1ContractProductGetResponseDataTypeFixed, V1ContractProductGetResponseDataTypeProService:
		return true
	}
	return false
}

type V1ContractProductGetResponseDataUpdate struct {
	CreatedAt           time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy           string    `json:"created_by,required"`
	BillableMetricID    string    `json:"billable_metric_id" format:"uuid"`
	CompositeProductIDs []string  `json:"composite_product_ids" format:"uuid"`
	CompositeTags       []string  `json:"composite_tags"`
	ExcludeFreeUsage    bool      `json:"exclude_free_usage"`
	IsRefundable        bool      `json:"is_refundable"`
	Name                string    `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID string `json:"netsuite_internal_item_id"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID string `json:"netsuite_overage_item_id"`
	// For USAGE products only. Groups usage line items on invoices. The superset of
	// values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PresentationGroupKey []string `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole. The superset
	// of values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PricingGroupKey []string `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion QuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding QuantityRounding                           `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                  `json:"starting_at" format:"date-time"`
	Tags             []string                                   `json:"tags"`
	JSON             v1ContractProductGetResponseDataUpdateJSON `json:"-"`
}

// v1ContractProductGetResponseDataUpdateJSON contains the JSON metadata for the
// struct [V1ContractProductGetResponseDataUpdate]
type v1ContractProductGetResponseDataUpdateJSON struct {
	CreatedAt              apijson.Field
	CreatedBy              apijson.Field
	BillableMetricID       apijson.Field
	CompositeProductIDs    apijson.Field
	CompositeTags          apijson.Field
	ExcludeFreeUsage       apijson.Field
	IsRefundable           apijson.Field
	Name                   apijson.Field
	NetsuiteInternalItemID apijson.Field
	NetsuiteOverageItemID  apijson.Field
	PresentationGroupKey   apijson.Field
	PricingGroupKey        apijson.Field
	QuantityConversion     apijson.Field
	QuantityRounding       apijson.Field
	StartingAt             apijson.Field
	Tags                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1ContractProductGetResponseDataUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductGetResponseDataUpdateJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductUpdateResponse struct {
	Data shared.ID                           `json:"data,required"`
	JSON v1ContractProductUpdateResponseJSON `json:"-"`
}

// v1ContractProductUpdateResponseJSON contains the JSON metadata for the struct
// [V1ContractProductUpdateResponse]
type v1ContractProductUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractProductUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductListResponse struct {
	ID           string                                `json:"id,required" format:"uuid"`
	Current      ProductListItemState                  `json:"current,required"`
	Initial      ProductListItemState                  `json:"initial,required"`
	Type         V1ContractProductListResponseType     `json:"type,required"`
	Updates      []V1ContractProductListResponseUpdate `json:"updates,required"`
	ArchivedAt   time.Time                             `json:"archived_at,nullable" format:"date-time"`
	CustomFields map[string]string                     `json:"custom_fields"`
	JSON         v1ContractProductListResponseJSON     `json:"-"`
}

// v1ContractProductListResponseJSON contains the JSON metadata for the struct
// [V1ContractProductListResponse]
type v1ContractProductListResponseJSON struct {
	ID           apijson.Field
	Current      apijson.Field
	Initial      apijson.Field
	Type         apijson.Field
	Updates      apijson.Field
	ArchivedAt   apijson.Field
	CustomFields apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1ContractProductListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductListResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductListResponseType string

const (
	V1ContractProductListResponseTypeUsage        V1ContractProductListResponseType = "USAGE"
	V1ContractProductListResponseTypeSubscription V1ContractProductListResponseType = "SUBSCRIPTION"
	V1ContractProductListResponseTypeComposite    V1ContractProductListResponseType = "COMPOSITE"
	V1ContractProductListResponseTypeFixed        V1ContractProductListResponseType = "FIXED"
	V1ContractProductListResponseTypeProService   V1ContractProductListResponseType = "PRO_SERVICE"
)

func (r V1ContractProductListResponseType) IsKnown() bool {
	switch r {
	case V1ContractProductListResponseTypeUsage, V1ContractProductListResponseTypeSubscription, V1ContractProductListResponseTypeComposite, V1ContractProductListResponseTypeFixed, V1ContractProductListResponseTypeProService:
		return true
	}
	return false
}

type V1ContractProductListResponseUpdate struct {
	CreatedAt           time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy           string    `json:"created_by,required"`
	BillableMetricID    string    `json:"billable_metric_id" format:"uuid"`
	CompositeProductIDs []string  `json:"composite_product_ids" format:"uuid"`
	CompositeTags       []string  `json:"composite_tags"`
	ExcludeFreeUsage    bool      `json:"exclude_free_usage"`
	IsRefundable        bool      `json:"is_refundable"`
	Name                string    `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID string `json:"netsuite_internal_item_id"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID string `json:"netsuite_overage_item_id"`
	// For USAGE products only. Groups usage line items on invoices. The superset of
	// values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PresentationGroupKey []string `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole. The superset
	// of values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PricingGroupKey []string `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion QuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding QuantityRounding                        `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                               `json:"starting_at" format:"date-time"`
	Tags             []string                                `json:"tags"`
	JSON             v1ContractProductListResponseUpdateJSON `json:"-"`
}

// v1ContractProductListResponseUpdateJSON contains the JSON metadata for the
// struct [V1ContractProductListResponseUpdate]
type v1ContractProductListResponseUpdateJSON struct {
	CreatedAt              apijson.Field
	CreatedBy              apijson.Field
	BillableMetricID       apijson.Field
	CompositeProductIDs    apijson.Field
	CompositeTags          apijson.Field
	ExcludeFreeUsage       apijson.Field
	IsRefundable           apijson.Field
	Name                   apijson.Field
	NetsuiteInternalItemID apijson.Field
	NetsuiteOverageItemID  apijson.Field
	PresentationGroupKey   apijson.Field
	PricingGroupKey        apijson.Field
	QuantityConversion     apijson.Field
	QuantityRounding       apijson.Field
	StartingAt             apijson.Field
	Tags                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1ContractProductListResponseUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductListResponseUpdateJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductArchiveResponse struct {
	Data shared.ID                            `json:"data,required"`
	JSON v1ContractProductArchiveResponseJSON `json:"-"`
}

// v1ContractProductArchiveResponseJSON contains the JSON metadata for the struct
// [V1ContractProductArchiveResponse]
type v1ContractProductArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractProductArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductNewParams struct {
	// displayed on invoices
	Name param.Field[string]                         `json:"name,required"`
	Type param.Field[V1ContractProductNewParamsType] `json:"type,required"`
	// Required for USAGE products
	BillableMetricID param.Field[string] `json:"billable_metric_id" format:"uuid"`
	// Required for COMPOSITE products
	CompositeProductIDs param.Field[[]string] `json:"composite_product_ids" format:"uuid"`
	// Required for COMPOSITE products
	CompositeTags param.Field[[]string] `json:"composite_tags"`
	// Beta feature only available for composite products. If true, products with $0
	// will not be included when computing composite usage. Defaults to false
	ExcludeFreeUsage param.Field[bool] `json:"exclude_free_usage"`
	// This field's availability is dependent on your client's configuration. Defaults
	// to true.
	IsRefundable param.Field[bool] `json:"is_refundable"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID param.Field[string] `json:"netsuite_internal_item_id"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID param.Field[string] `json:"netsuite_overage_item_id"`
	// For USAGE products only. Groups usage line items on invoices. The superset of
	// values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PresentationGroupKey param.Field[[]string] `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole. The superset
	// of values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PricingGroupKey param.Field[[]string] `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion param.Field[QuantityConversionParam] `json:"quantity_conversion"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding param.Field[QuantityRoundingParam] `json:"quantity_rounding"`
	Tags             param.Field[[]string]              `json:"tags"`
}

func (r V1ContractProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractProductNewParamsType string

const (
	V1ContractProductNewParamsTypeFixed               V1ContractProductNewParamsType = "FIXED"
	V1ContractProductNewParamsTypeUsage               V1ContractProductNewParamsType = "USAGE"
	V1ContractProductNewParamsTypeComposite           V1ContractProductNewParamsType = "COMPOSITE"
	V1ContractProductNewParamsTypeSubscription        V1ContractProductNewParamsType = "SUBSCRIPTION"
	V1ContractProductNewParamsTypeProfessionalService V1ContractProductNewParamsType = "PROFESSIONAL_SERVICE"
	V1ContractProductNewParamsTypeProService          V1ContractProductNewParamsType = "PRO_SERVICE"
)

func (r V1ContractProductNewParamsType) IsKnown() bool {
	switch r {
	case V1ContractProductNewParamsTypeFixed, V1ContractProductNewParamsTypeUsage, V1ContractProductNewParamsTypeComposite, V1ContractProductNewParamsTypeSubscription, V1ContractProductNewParamsTypeProfessionalService, V1ContractProductNewParamsTypeProService:
		return true
	}
	return false
}

type V1ContractProductGetParams struct {
	ID shared.IDParam `json:"id,required"`
}

func (r V1ContractProductGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ID)
}

type V1ContractProductUpdateParams struct {
	// ID of the product to update
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Timestamp representing when the update should go into effect. It must be on an
	// hour boundary (e.g. 1:00, not 1:30).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Available for USAGE products only. If not provided, defaults to product's
	// current billable metric.
	BillableMetricID param.Field[string] `json:"billable_metric_id" format:"uuid"`
	// Available for COMPOSITE products only. If not provided, defaults to product's
	// current composite_product_ids.
	CompositeProductIDs param.Field[[]string] `json:"composite_product_ids" format:"uuid"`
	// Available for COMPOSITE products only. If not provided, defaults to product's
	// current composite_tags.
	CompositeTags param.Field[[]string] `json:"composite_tags"`
	// Beta feature only available for composite products. If true, products with $0
	// will not be included when computing composite usage. Defaults to false
	ExcludeFreeUsage param.Field[bool] `json:"exclude_free_usage"`
	// Defaults to product's current refundability status. This field's availability is
	// dependent on your client's configuration.
	IsRefundable param.Field[bool] `json:"is_refundable"`
	// displayed on invoices. If not provided, defaults to product's current name.
	Name param.Field[string] `json:"name"`
	// If not provided, defaults to product's current netsuite_internal_item_id. This
	// field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID param.Field[string] `json:"netsuite_internal_item_id"`
	// Available for USAGE and COMPOSITE products only. If not provided, defaults to
	// product's current netsuite_overage_item_id. This field's availability is
	// dependent on your client's configuration.
	NetsuiteOverageItemID param.Field[string] `json:"netsuite_overage_item_id"`
	// For USAGE products only. Groups usage line items on invoices. The superset of
	// values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PresentationGroupKey param.Field[[]string] `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole. The superset
	// of values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PricingGroupKey param.Field[[]string] `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion param.Field[QuantityConversionParam] `json:"quantity_conversion"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding param.Field[QuantityRoundingParam] `json:"quantity_rounding"`
	// If not provided, defaults to product's current tags
	Tags param.Field[[]string] `json:"tags"`
}

func (r V1ContractProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractProductListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Filter options for the product list
	ArchiveFilter param.Field[V1ContractProductListParamsArchiveFilter] `json:"archive_filter"`
}

func (r V1ContractProductListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V1ContractProductListParams]'s query parameters as
// `url.Values`.
func (r V1ContractProductListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter options for the product list
type V1ContractProductListParamsArchiveFilter string

const (
	V1ContractProductListParamsArchiveFilterArchived    V1ContractProductListParamsArchiveFilter = "ARCHIVED"
	V1ContractProductListParamsArchiveFilterNotArchived V1ContractProductListParamsArchiveFilter = "NOT_ARCHIVED"
	V1ContractProductListParamsArchiveFilterAll         V1ContractProductListParamsArchiveFilter = "ALL"
)

func (r V1ContractProductListParamsArchiveFilter) IsKnown() bool {
	switch r {
	case V1ContractProductListParamsArchiveFilterArchived, V1ContractProductListParamsArchiveFilterNotArchived, V1ContractProductListParamsArchiveFilterAll:
		return true
	}
	return false
}

type V1ContractProductArchiveParams struct {
	// ID of the product to be archived
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
}

func (r V1ContractProductArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
