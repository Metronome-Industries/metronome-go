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

// ContractProductService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContractProductService] method instead.
type ContractProductService struct {
	Options []option.RequestOption
}

// NewContractProductService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractProductService(opts ...option.RequestOption) (r *ContractProductService) {
	r = &ContractProductService{}
	r.Options = opts
	return
}

// Create a new product
func (r *ContractProductService) New(ctx context.Context, body ContractProductNewParams, opts ...option.RequestOption) (res *ContractProductNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/products/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific product
func (r *ContractProductService) Get(ctx context.Context, body ContractProductGetParams, opts ...option.RequestOption) (res *ContractProductGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/products/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a product
func (r *ContractProductService) Update(ctx context.Context, body ContractProductUpdateParams, opts ...option.RequestOption) (res *ContractProductUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/products/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List products
func (r *ContractProductService) List(ctx context.Context, params ContractProductListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ContractProductListResponse], err error) {
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
func (r *ContractProductService) ListAutoPaging(ctx context.Context, params ContractProductListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ContractProductListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Archive a product
func (r *ContractProductService) Archive(ctx context.Context, body ContractProductArchiveParams, opts ...option.RequestOption) (res *ContractProductArchiveResponse, err error) {
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

type ContractProductNewResponse struct {
	Data shared.ID                      `json:"data,required"`
	JSON contractProductNewResponseJSON `json:"-"`
}

// contractProductNewResponseJSON contains the JSON metadata for the struct
// [ContractProductNewResponse]
type contractProductNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractProductNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductNewResponseJSON) RawJSON() string {
	return r.raw
}

type ContractProductGetResponse struct {
	Data ContractProductGetResponseData `json:"data,required"`
	JSON contractProductGetResponseJSON `json:"-"`
}

// contractProductGetResponseJSON contains the JSON metadata for the struct
// [ContractProductGetResponse]
type contractProductGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractProductGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductGetResponseJSON) RawJSON() string {
	return r.raw
}

type ContractProductGetResponseData struct {
	ID           string                                 `json:"id,required" format:"uuid"`
	Current      ProductListItemState                   `json:"current,required"`
	Initial      ProductListItemState                   `json:"initial,required"`
	Type         ContractProductGetResponseDataType     `json:"type,required"`
	Updates      []ContractProductGetResponseDataUpdate `json:"updates,required"`
	ArchivedAt   time.Time                              `json:"archived_at,nullable" format:"date-time"`
	CustomFields map[string]string                      `json:"custom_fields"`
	JSON         contractProductGetResponseDataJSON     `json:"-"`
}

// contractProductGetResponseDataJSON contains the JSON metadata for the struct
// [ContractProductGetResponseData]
type contractProductGetResponseDataJSON struct {
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

func (r *ContractProductGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ContractProductGetResponseDataType string

const (
	ContractProductGetResponseDataTypeUsage        ContractProductGetResponseDataType = "USAGE"
	ContractProductGetResponseDataTypeSubscription ContractProductGetResponseDataType = "SUBSCRIPTION"
	ContractProductGetResponseDataTypeComposite    ContractProductGetResponseDataType = "COMPOSITE"
	ContractProductGetResponseDataTypeFixed        ContractProductGetResponseDataType = "FIXED"
	ContractProductGetResponseDataTypeProService   ContractProductGetResponseDataType = "PRO_SERVICE"
)

func (r ContractProductGetResponseDataType) IsKnown() bool {
	switch r {
	case ContractProductGetResponseDataTypeUsage, ContractProductGetResponseDataTypeSubscription, ContractProductGetResponseDataTypeComposite, ContractProductGetResponseDataTypeFixed, ContractProductGetResponseDataTypeProService:
		return true
	}
	return false
}

type ContractProductGetResponseDataUpdate struct {
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
	QuantityRounding QuantityRounding                         `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                `json:"starting_at" format:"date-time"`
	Tags             []string                                 `json:"tags"`
	JSON             contractProductGetResponseDataUpdateJSON `json:"-"`
}

// contractProductGetResponseDataUpdateJSON contains the JSON metadata for the
// struct [ContractProductGetResponseDataUpdate]
type contractProductGetResponseDataUpdateJSON struct {
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

func (r *ContractProductGetResponseDataUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductGetResponseDataUpdateJSON) RawJSON() string {
	return r.raw
}

type ContractProductUpdateResponse struct {
	Data shared.ID                         `json:"data,required"`
	JSON contractProductUpdateResponseJSON `json:"-"`
}

// contractProductUpdateResponseJSON contains the JSON metadata for the struct
// [ContractProductUpdateResponse]
type contractProductUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractProductUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ContractProductListResponse struct {
	ID           string                              `json:"id,required" format:"uuid"`
	Current      ProductListItemState                `json:"current,required"`
	Initial      ProductListItemState                `json:"initial,required"`
	Type         ContractProductListResponseType     `json:"type,required"`
	Updates      []ContractProductListResponseUpdate `json:"updates,required"`
	ArchivedAt   time.Time                           `json:"archived_at,nullable" format:"date-time"`
	CustomFields map[string]string                   `json:"custom_fields"`
	JSON         contractProductListResponseJSON     `json:"-"`
}

// contractProductListResponseJSON contains the JSON metadata for the struct
// [ContractProductListResponse]
type contractProductListResponseJSON struct {
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

func (r *ContractProductListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductListResponseJSON) RawJSON() string {
	return r.raw
}

type ContractProductListResponseType string

const (
	ContractProductListResponseTypeUsage        ContractProductListResponseType = "USAGE"
	ContractProductListResponseTypeSubscription ContractProductListResponseType = "SUBSCRIPTION"
	ContractProductListResponseTypeComposite    ContractProductListResponseType = "COMPOSITE"
	ContractProductListResponseTypeFixed        ContractProductListResponseType = "FIXED"
	ContractProductListResponseTypeProService   ContractProductListResponseType = "PRO_SERVICE"
)

func (r ContractProductListResponseType) IsKnown() bool {
	switch r {
	case ContractProductListResponseTypeUsage, ContractProductListResponseTypeSubscription, ContractProductListResponseTypeComposite, ContractProductListResponseTypeFixed, ContractProductListResponseTypeProService:
		return true
	}
	return false
}

type ContractProductListResponseUpdate struct {
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
	QuantityRounding QuantityRounding                      `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                             `json:"starting_at" format:"date-time"`
	Tags             []string                              `json:"tags"`
	JSON             contractProductListResponseUpdateJSON `json:"-"`
}

// contractProductListResponseUpdateJSON contains the JSON metadata for the struct
// [ContractProductListResponseUpdate]
type contractProductListResponseUpdateJSON struct {
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

func (r *ContractProductListResponseUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductListResponseUpdateJSON) RawJSON() string {
	return r.raw
}

type ContractProductArchiveResponse struct {
	Data shared.ID                          `json:"data,required"`
	JSON contractProductArchiveResponseJSON `json:"-"`
}

// contractProductArchiveResponseJSON contains the JSON metadata for the struct
// [ContractProductArchiveResponse]
type contractProductArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractProductArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type ContractProductNewParams struct {
	// displayed on invoices
	Name param.Field[string]                       `json:"name,required"`
	Type param.Field[ContractProductNewParamsType] `json:"type,required"`
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

func (r ContractProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractProductNewParamsType string

const (
	ContractProductNewParamsTypeFixed               ContractProductNewParamsType = "FIXED"
	ContractProductNewParamsTypeUsage               ContractProductNewParamsType = "USAGE"
	ContractProductNewParamsTypeComposite           ContractProductNewParamsType = "COMPOSITE"
	ContractProductNewParamsTypeSubscription        ContractProductNewParamsType = "SUBSCRIPTION"
	ContractProductNewParamsTypeProfessionalService ContractProductNewParamsType = "PROFESSIONAL_SERVICE"
	ContractProductNewParamsTypeProService          ContractProductNewParamsType = "PRO_SERVICE"
)

func (r ContractProductNewParamsType) IsKnown() bool {
	switch r {
	case ContractProductNewParamsTypeFixed, ContractProductNewParamsTypeUsage, ContractProductNewParamsTypeComposite, ContractProductNewParamsTypeSubscription, ContractProductNewParamsTypeProfessionalService, ContractProductNewParamsTypeProService:
		return true
	}
	return false
}

type ContractProductGetParams struct {
	ID shared.IDParam `json:"id,required"`
}

func (r ContractProductGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ID)
}

type ContractProductUpdateParams struct {
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

func (r ContractProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractProductListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Filter options for the product list
	ArchiveFilter param.Field[ContractProductListParamsArchiveFilter] `json:"archive_filter"`
}

func (r ContractProductListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ContractProductListParams]'s query parameters as
// `url.Values`.
func (r ContractProductListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter options for the product list
type ContractProductListParamsArchiveFilter string

const (
	ContractProductListParamsArchiveFilterArchived    ContractProductListParamsArchiveFilter = "ARCHIVED"
	ContractProductListParamsArchiveFilterNotArchived ContractProductListParamsArchiveFilter = "NOT_ARCHIVED"
	ContractProductListParamsArchiveFilterAll         ContractProductListParamsArchiveFilter = "ALL"
)

func (r ContractProductListParamsArchiveFilter) IsKnown() bool {
	switch r {
	case ContractProductListParamsArchiveFilterArchived, ContractProductListParamsArchiveFilterNotArchived, ContractProductListParamsArchiveFilterAll:
		return true
	}
	return false
}

type ContractProductArchiveParams struct {
	// ID of the product to be archived
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
}

func (r ContractProductArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
