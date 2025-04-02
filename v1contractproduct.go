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

type V1ContractProductNewResponse struct {
	Data V1ContractProductNewResponseData `json:"data,required"`
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

type V1ContractProductNewResponseData struct {
	ID   string                               `json:"id,required" format:"uuid"`
	JSON v1ContractProductNewResponseDataJSON `json:"-"`
}

// v1ContractProductNewResponseDataJSON contains the JSON metadata for the struct
// [V1ContractProductNewResponseData]
type v1ContractProductNewResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractProductNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductNewResponseDataJSON) RawJSON() string {
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
	Current      V1ContractProductGetResponseDataCurrent  `json:"current,required"`
	Initial      V1ContractProductGetResponseDataInitial  `json:"initial,required"`
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

type V1ContractProductGetResponseDataCurrent struct {
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
	QuantityConversion V1ContractProductGetResponseDataCurrentQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding V1ContractProductGetResponseDataCurrentQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                               `json:"starting_at" format:"date-time"`
	Tags             []string                                                `json:"tags"`
	JSON             v1ContractProductGetResponseDataCurrentJSON             `json:"-"`
}

// v1ContractProductGetResponseDataCurrentJSON contains the JSON metadata for the
// struct [V1ContractProductGetResponseDataCurrent]
type v1ContractProductGetResponseDataCurrentJSON struct {
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

func (r *V1ContractProductGetResponseDataCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductGetResponseDataCurrentJSON) RawJSON() string {
	return r.raw
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type V1ContractProductGetResponseDataCurrentQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation V1ContractProductGetResponseDataCurrentQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                        `json:"name"`
	JSON v1ContractProductGetResponseDataCurrentQuantityConversionJSON `json:"-"`
}

// v1ContractProductGetResponseDataCurrentQuantityConversionJSON contains the JSON
// metadata for the struct
// [V1ContractProductGetResponseDataCurrentQuantityConversion]
type v1ContractProductGetResponseDataCurrentQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1ContractProductGetResponseDataCurrentQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductGetResponseDataCurrentQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type V1ContractProductGetResponseDataCurrentQuantityConversionOperation string

const (
	V1ContractProductGetResponseDataCurrentQuantityConversionOperationMultiply V1ContractProductGetResponseDataCurrentQuantityConversionOperation = "MULTIPLY"
	V1ContractProductGetResponseDataCurrentQuantityConversionOperationDivide   V1ContractProductGetResponseDataCurrentQuantityConversionOperation = "DIVIDE"
)

func (r V1ContractProductGetResponseDataCurrentQuantityConversionOperation) IsKnown() bool {
	switch r {
	case V1ContractProductGetResponseDataCurrentQuantityConversionOperationMultiply, V1ContractProductGetResponseDataCurrentQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type V1ContractProductGetResponseDataCurrentQuantityRounding struct {
	DecimalPlaces  float64                                                               `json:"decimal_places,required"`
	RoundingMethod V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           v1ContractProductGetResponseDataCurrentQuantityRoundingJSON           `json:"-"`
}

// v1ContractProductGetResponseDataCurrentQuantityRoundingJSON contains the JSON
// metadata for the struct
// [V1ContractProductGetResponseDataCurrentQuantityRounding]
type v1ContractProductGetResponseDataCurrentQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1ContractProductGetResponseDataCurrentQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductGetResponseDataCurrentQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod string

const (
	V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundUp     V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod = "ROUND_UP"
	V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundDown   V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod = "ROUND_DOWN"
	V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundHalfUp V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundUp, V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundDown, V1ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
}

type V1ContractProductGetResponseDataInitial struct {
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
	QuantityConversion V1ContractProductGetResponseDataInitialQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding V1ContractProductGetResponseDataInitialQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                               `json:"starting_at" format:"date-time"`
	Tags             []string                                                `json:"tags"`
	JSON             v1ContractProductGetResponseDataInitialJSON             `json:"-"`
}

// v1ContractProductGetResponseDataInitialJSON contains the JSON metadata for the
// struct [V1ContractProductGetResponseDataInitial]
type v1ContractProductGetResponseDataInitialJSON struct {
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

func (r *V1ContractProductGetResponseDataInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductGetResponseDataInitialJSON) RawJSON() string {
	return r.raw
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type V1ContractProductGetResponseDataInitialQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation V1ContractProductGetResponseDataInitialQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                        `json:"name"`
	JSON v1ContractProductGetResponseDataInitialQuantityConversionJSON `json:"-"`
}

// v1ContractProductGetResponseDataInitialQuantityConversionJSON contains the JSON
// metadata for the struct
// [V1ContractProductGetResponseDataInitialQuantityConversion]
type v1ContractProductGetResponseDataInitialQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1ContractProductGetResponseDataInitialQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductGetResponseDataInitialQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type V1ContractProductGetResponseDataInitialQuantityConversionOperation string

const (
	V1ContractProductGetResponseDataInitialQuantityConversionOperationMultiply V1ContractProductGetResponseDataInitialQuantityConversionOperation = "MULTIPLY"
	V1ContractProductGetResponseDataInitialQuantityConversionOperationDivide   V1ContractProductGetResponseDataInitialQuantityConversionOperation = "DIVIDE"
)

func (r V1ContractProductGetResponseDataInitialQuantityConversionOperation) IsKnown() bool {
	switch r {
	case V1ContractProductGetResponseDataInitialQuantityConversionOperationMultiply, V1ContractProductGetResponseDataInitialQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type V1ContractProductGetResponseDataInitialQuantityRounding struct {
	DecimalPlaces  float64                                                               `json:"decimal_places,required"`
	RoundingMethod V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           v1ContractProductGetResponseDataInitialQuantityRoundingJSON           `json:"-"`
}

// v1ContractProductGetResponseDataInitialQuantityRoundingJSON contains the JSON
// metadata for the struct
// [V1ContractProductGetResponseDataInitialQuantityRounding]
type v1ContractProductGetResponseDataInitialQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1ContractProductGetResponseDataInitialQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductGetResponseDataInitialQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod string

const (
	V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundUp     V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod = "ROUND_UP"
	V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundDown   V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod = "ROUND_DOWN"
	V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundHalfUp V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundUp, V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundDown, V1ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
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
	QuantityConversion V1ContractProductGetResponseDataUpdatesQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding V1ContractProductGetResponseDataUpdatesQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                               `json:"starting_at" format:"date-time"`
	Tags             []string                                                `json:"tags"`
	JSON             v1ContractProductGetResponseDataUpdateJSON              `json:"-"`
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

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type V1ContractProductGetResponseDataUpdatesQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation V1ContractProductGetResponseDataUpdatesQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                        `json:"name"`
	JSON v1ContractProductGetResponseDataUpdatesQuantityConversionJSON `json:"-"`
}

// v1ContractProductGetResponseDataUpdatesQuantityConversionJSON contains the JSON
// metadata for the struct
// [V1ContractProductGetResponseDataUpdatesQuantityConversion]
type v1ContractProductGetResponseDataUpdatesQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1ContractProductGetResponseDataUpdatesQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductGetResponseDataUpdatesQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type V1ContractProductGetResponseDataUpdatesQuantityConversionOperation string

const (
	V1ContractProductGetResponseDataUpdatesQuantityConversionOperationMultiply V1ContractProductGetResponseDataUpdatesQuantityConversionOperation = "MULTIPLY"
	V1ContractProductGetResponseDataUpdatesQuantityConversionOperationDivide   V1ContractProductGetResponseDataUpdatesQuantityConversionOperation = "DIVIDE"
)

func (r V1ContractProductGetResponseDataUpdatesQuantityConversionOperation) IsKnown() bool {
	switch r {
	case V1ContractProductGetResponseDataUpdatesQuantityConversionOperationMultiply, V1ContractProductGetResponseDataUpdatesQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type V1ContractProductGetResponseDataUpdatesQuantityRounding struct {
	DecimalPlaces  float64                                                               `json:"decimal_places,required"`
	RoundingMethod V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           v1ContractProductGetResponseDataUpdatesQuantityRoundingJSON           `json:"-"`
}

// v1ContractProductGetResponseDataUpdatesQuantityRoundingJSON contains the JSON
// metadata for the struct
// [V1ContractProductGetResponseDataUpdatesQuantityRounding]
type v1ContractProductGetResponseDataUpdatesQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1ContractProductGetResponseDataUpdatesQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductGetResponseDataUpdatesQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod string

const (
	V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundUp     V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod = "ROUND_UP"
	V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundDown   V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod = "ROUND_DOWN"
	V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundHalfUp V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundUp, V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundDown, V1ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
}

type V1ContractProductUpdateResponse struct {
	Data V1ContractProductUpdateResponseData `json:"data,required"`
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

type V1ContractProductUpdateResponseData struct {
	ID   string                                  `json:"id,required" format:"uuid"`
	JSON v1ContractProductUpdateResponseDataJSON `json:"-"`
}

// v1ContractProductUpdateResponseDataJSON contains the JSON metadata for the
// struct [V1ContractProductUpdateResponseData]
type v1ContractProductUpdateResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractProductUpdateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductUpdateResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductListResponse struct {
	ID           string                                `json:"id,required" format:"uuid"`
	Current      V1ContractProductListResponseCurrent  `json:"current,required"`
	Initial      V1ContractProductListResponseInitial  `json:"initial,required"`
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

type V1ContractProductListResponseCurrent struct {
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
	QuantityConversion V1ContractProductListResponseCurrentQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding V1ContractProductListResponseCurrentQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                            `json:"starting_at" format:"date-time"`
	Tags             []string                                             `json:"tags"`
	JSON             v1ContractProductListResponseCurrentJSON             `json:"-"`
}

// v1ContractProductListResponseCurrentJSON contains the JSON metadata for the
// struct [V1ContractProductListResponseCurrent]
type v1ContractProductListResponseCurrentJSON struct {
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

func (r *V1ContractProductListResponseCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductListResponseCurrentJSON) RawJSON() string {
	return r.raw
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type V1ContractProductListResponseCurrentQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation V1ContractProductListResponseCurrentQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                     `json:"name"`
	JSON v1ContractProductListResponseCurrentQuantityConversionJSON `json:"-"`
}

// v1ContractProductListResponseCurrentQuantityConversionJSON contains the JSON
// metadata for the struct [V1ContractProductListResponseCurrentQuantityConversion]
type v1ContractProductListResponseCurrentQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1ContractProductListResponseCurrentQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductListResponseCurrentQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type V1ContractProductListResponseCurrentQuantityConversionOperation string

const (
	V1ContractProductListResponseCurrentQuantityConversionOperationMultiply V1ContractProductListResponseCurrentQuantityConversionOperation = "MULTIPLY"
	V1ContractProductListResponseCurrentQuantityConversionOperationDivide   V1ContractProductListResponseCurrentQuantityConversionOperation = "DIVIDE"
)

func (r V1ContractProductListResponseCurrentQuantityConversionOperation) IsKnown() bool {
	switch r {
	case V1ContractProductListResponseCurrentQuantityConversionOperationMultiply, V1ContractProductListResponseCurrentQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type V1ContractProductListResponseCurrentQuantityRounding struct {
	DecimalPlaces  float64                                                            `json:"decimal_places,required"`
	RoundingMethod V1ContractProductListResponseCurrentQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           v1ContractProductListResponseCurrentQuantityRoundingJSON           `json:"-"`
}

// v1ContractProductListResponseCurrentQuantityRoundingJSON contains the JSON
// metadata for the struct [V1ContractProductListResponseCurrentQuantityRounding]
type v1ContractProductListResponseCurrentQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1ContractProductListResponseCurrentQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductListResponseCurrentQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductListResponseCurrentQuantityRoundingRoundingMethod string

const (
	V1ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundUp     V1ContractProductListResponseCurrentQuantityRoundingRoundingMethod = "ROUND_UP"
	V1ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundDown   V1ContractProductListResponseCurrentQuantityRoundingRoundingMethod = "ROUND_DOWN"
	V1ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundHalfUp V1ContractProductListResponseCurrentQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r V1ContractProductListResponseCurrentQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case V1ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundUp, V1ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundDown, V1ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
}

type V1ContractProductListResponseInitial struct {
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
	QuantityConversion V1ContractProductListResponseInitialQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding V1ContractProductListResponseInitialQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                            `json:"starting_at" format:"date-time"`
	Tags             []string                                             `json:"tags"`
	JSON             v1ContractProductListResponseInitialJSON             `json:"-"`
}

// v1ContractProductListResponseInitialJSON contains the JSON metadata for the
// struct [V1ContractProductListResponseInitial]
type v1ContractProductListResponseInitialJSON struct {
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

func (r *V1ContractProductListResponseInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductListResponseInitialJSON) RawJSON() string {
	return r.raw
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type V1ContractProductListResponseInitialQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation V1ContractProductListResponseInitialQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                     `json:"name"`
	JSON v1ContractProductListResponseInitialQuantityConversionJSON `json:"-"`
}

// v1ContractProductListResponseInitialQuantityConversionJSON contains the JSON
// metadata for the struct [V1ContractProductListResponseInitialQuantityConversion]
type v1ContractProductListResponseInitialQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1ContractProductListResponseInitialQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductListResponseInitialQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type V1ContractProductListResponseInitialQuantityConversionOperation string

const (
	V1ContractProductListResponseInitialQuantityConversionOperationMultiply V1ContractProductListResponseInitialQuantityConversionOperation = "MULTIPLY"
	V1ContractProductListResponseInitialQuantityConversionOperationDivide   V1ContractProductListResponseInitialQuantityConversionOperation = "DIVIDE"
)

func (r V1ContractProductListResponseInitialQuantityConversionOperation) IsKnown() bool {
	switch r {
	case V1ContractProductListResponseInitialQuantityConversionOperationMultiply, V1ContractProductListResponseInitialQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type V1ContractProductListResponseInitialQuantityRounding struct {
	DecimalPlaces  float64                                                            `json:"decimal_places,required"`
	RoundingMethod V1ContractProductListResponseInitialQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           v1ContractProductListResponseInitialQuantityRoundingJSON           `json:"-"`
}

// v1ContractProductListResponseInitialQuantityRoundingJSON contains the JSON
// metadata for the struct [V1ContractProductListResponseInitialQuantityRounding]
type v1ContractProductListResponseInitialQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1ContractProductListResponseInitialQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductListResponseInitialQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductListResponseInitialQuantityRoundingRoundingMethod string

const (
	V1ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundUp     V1ContractProductListResponseInitialQuantityRoundingRoundingMethod = "ROUND_UP"
	V1ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundDown   V1ContractProductListResponseInitialQuantityRoundingRoundingMethod = "ROUND_DOWN"
	V1ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundHalfUp V1ContractProductListResponseInitialQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r V1ContractProductListResponseInitialQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case V1ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundUp, V1ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundDown, V1ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
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
	QuantityConversion V1ContractProductListResponseUpdatesQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding V1ContractProductListResponseUpdatesQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                            `json:"starting_at" format:"date-time"`
	Tags             []string                                             `json:"tags"`
	JSON             v1ContractProductListResponseUpdateJSON              `json:"-"`
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

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type V1ContractProductListResponseUpdatesQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation V1ContractProductListResponseUpdatesQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                     `json:"name"`
	JSON v1ContractProductListResponseUpdatesQuantityConversionJSON `json:"-"`
}

// v1ContractProductListResponseUpdatesQuantityConversionJSON contains the JSON
// metadata for the struct [V1ContractProductListResponseUpdatesQuantityConversion]
type v1ContractProductListResponseUpdatesQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1ContractProductListResponseUpdatesQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductListResponseUpdatesQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type V1ContractProductListResponseUpdatesQuantityConversionOperation string

const (
	V1ContractProductListResponseUpdatesQuantityConversionOperationMultiply V1ContractProductListResponseUpdatesQuantityConversionOperation = "MULTIPLY"
	V1ContractProductListResponseUpdatesQuantityConversionOperationDivide   V1ContractProductListResponseUpdatesQuantityConversionOperation = "DIVIDE"
)

func (r V1ContractProductListResponseUpdatesQuantityConversionOperation) IsKnown() bool {
	switch r {
	case V1ContractProductListResponseUpdatesQuantityConversionOperationMultiply, V1ContractProductListResponseUpdatesQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type V1ContractProductListResponseUpdatesQuantityRounding struct {
	DecimalPlaces  float64                                                            `json:"decimal_places,required"`
	RoundingMethod V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           v1ContractProductListResponseUpdatesQuantityRoundingJSON           `json:"-"`
}

// v1ContractProductListResponseUpdatesQuantityRoundingJSON contains the JSON
// metadata for the struct [V1ContractProductListResponseUpdatesQuantityRounding]
type v1ContractProductListResponseUpdatesQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1ContractProductListResponseUpdatesQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductListResponseUpdatesQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethod string

const (
	V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundUp     V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethod = "ROUND_UP"
	V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundDown   V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethod = "ROUND_DOWN"
	V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundHalfUp V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundUp, V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundDown, V1ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
}

type V1ContractProductArchiveResponse struct {
	Data V1ContractProductArchiveResponseData `json:"data,required"`
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

type V1ContractProductArchiveResponseData struct {
	ID   string                                   `json:"id,required" format:"uuid"`
	JSON v1ContractProductArchiveResponseDataJSON `json:"-"`
}

// v1ContractProductArchiveResponseDataJSON contains the JSON metadata for the
// struct [V1ContractProductArchiveResponseData]
type v1ContractProductArchiveResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractProductArchiveResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractProductArchiveResponseDataJSON) RawJSON() string {
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
	QuantityConversion param.Field[V1ContractProductNewParamsQuantityConversion] `json:"quantity_conversion"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding param.Field[V1ContractProductNewParamsQuantityRounding] `json:"quantity_rounding"`
	Tags             param.Field[[]string]                                   `json:"tags"`
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

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type V1ContractProductNewParamsQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor param.Field[float64] `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation param.Field[V1ContractProductNewParamsQuantityConversionOperation] `json:"operation,required"`
	// Optional name for this conversion.
	Name param.Field[string] `json:"name"`
}

func (r V1ContractProductNewParamsQuantityConversion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operation to perform on the quantity
type V1ContractProductNewParamsQuantityConversionOperation string

const (
	V1ContractProductNewParamsQuantityConversionOperationMultiply V1ContractProductNewParamsQuantityConversionOperation = "MULTIPLY"
	V1ContractProductNewParamsQuantityConversionOperationDivide   V1ContractProductNewParamsQuantityConversionOperation = "DIVIDE"
)

func (r V1ContractProductNewParamsQuantityConversionOperation) IsKnown() bool {
	switch r {
	case V1ContractProductNewParamsQuantityConversionOperationMultiply, V1ContractProductNewParamsQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type V1ContractProductNewParamsQuantityRounding struct {
	DecimalPlaces  param.Field[float64]                                                  `json:"decimal_places,required"`
	RoundingMethod param.Field[V1ContractProductNewParamsQuantityRoundingRoundingMethod] `json:"rounding_method,required"`
}

func (r V1ContractProductNewParamsQuantityRounding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractProductNewParamsQuantityRoundingRoundingMethod string

const (
	V1ContractProductNewParamsQuantityRoundingRoundingMethodRoundUp     V1ContractProductNewParamsQuantityRoundingRoundingMethod = "ROUND_UP"
	V1ContractProductNewParamsQuantityRoundingRoundingMethodRoundDown   V1ContractProductNewParamsQuantityRoundingRoundingMethod = "ROUND_DOWN"
	V1ContractProductNewParamsQuantityRoundingRoundingMethodRoundHalfUp V1ContractProductNewParamsQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r V1ContractProductNewParamsQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case V1ContractProductNewParamsQuantityRoundingRoundingMethodRoundUp, V1ContractProductNewParamsQuantityRoundingRoundingMethodRoundDown, V1ContractProductNewParamsQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
}

type V1ContractProductGetParams struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V1ContractProductGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	QuantityConversion param.Field[V1ContractProductUpdateParamsQuantityConversion] `json:"quantity_conversion"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding param.Field[V1ContractProductUpdateParamsQuantityRounding] `json:"quantity_rounding"`
	// If not provided, defaults to product's current tags
	Tags param.Field[[]string] `json:"tags"`
}

func (r V1ContractProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type V1ContractProductUpdateParamsQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor param.Field[float64] `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation param.Field[V1ContractProductUpdateParamsQuantityConversionOperation] `json:"operation,required"`
	// Optional name for this conversion.
	Name param.Field[string] `json:"name"`
}

func (r V1ContractProductUpdateParamsQuantityConversion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operation to perform on the quantity
type V1ContractProductUpdateParamsQuantityConversionOperation string

const (
	V1ContractProductUpdateParamsQuantityConversionOperationMultiply V1ContractProductUpdateParamsQuantityConversionOperation = "MULTIPLY"
	V1ContractProductUpdateParamsQuantityConversionOperationDivide   V1ContractProductUpdateParamsQuantityConversionOperation = "DIVIDE"
)

func (r V1ContractProductUpdateParamsQuantityConversionOperation) IsKnown() bool {
	switch r {
	case V1ContractProductUpdateParamsQuantityConversionOperationMultiply, V1ContractProductUpdateParamsQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type V1ContractProductUpdateParamsQuantityRounding struct {
	DecimalPlaces  param.Field[float64]                                                     `json:"decimal_places,required"`
	RoundingMethod param.Field[V1ContractProductUpdateParamsQuantityRoundingRoundingMethod] `json:"rounding_method,required"`
}

func (r V1ContractProductUpdateParamsQuantityRounding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractProductUpdateParamsQuantityRoundingRoundingMethod string

const (
	V1ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundUp     V1ContractProductUpdateParamsQuantityRoundingRoundingMethod = "ROUND_UP"
	V1ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundDown   V1ContractProductUpdateParamsQuantityRoundingRoundingMethod = "ROUND_DOWN"
	V1ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundHalfUp V1ContractProductUpdateParamsQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r V1ContractProductUpdateParamsQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case V1ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundUp, V1ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundDown, V1ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
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
