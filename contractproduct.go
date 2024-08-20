// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/pagination"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
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
	path := "contract-pricing/products/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific product
func (r *ContractProductService) Get(ctx context.Context, body ContractProductGetParams, opts ...option.RequestOption) (res *ContractProductGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/products/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a product
func (r *ContractProductService) Update(ctx context.Context, body ContractProductUpdateParams, opts ...option.RequestOption) (res *ContractProductUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/products/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List products
func (r *ContractProductService) List(ctx context.Context, params ContractProductListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ContractProductListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "contract-pricing/products/list"
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
	path := "contract-pricing/products/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
	Current      ContractProductGetResponseDataCurrent  `json:"current,required"`
	Initial      ContractProductGetResponseDataInitial  `json:"initial,required"`
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

type ContractProductGetResponseDataCurrent struct {
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
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion ContractProductGetResponseDataCurrentQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding ContractProductGetResponseDataCurrentQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                             `json:"starting_at" format:"date-time"`
	Tags             []string                                              `json:"tags"`
	JSON             contractProductGetResponseDataCurrentJSON             `json:"-"`
}

// contractProductGetResponseDataCurrentJSON contains the JSON metadata for the
// struct [ContractProductGetResponseDataCurrent]
type contractProductGetResponseDataCurrentJSON struct {
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

func (r *ContractProductGetResponseDataCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductGetResponseDataCurrentJSON) RawJSON() string {
	return r.raw
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type ContractProductGetResponseDataCurrentQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation ContractProductGetResponseDataCurrentQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                      `json:"name"`
	JSON contractProductGetResponseDataCurrentQuantityConversionJSON `json:"-"`
}

// contractProductGetResponseDataCurrentQuantityConversionJSON contains the JSON
// metadata for the struct
// [ContractProductGetResponseDataCurrentQuantityConversion]
type contractProductGetResponseDataCurrentQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ContractProductGetResponseDataCurrentQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductGetResponseDataCurrentQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type ContractProductGetResponseDataCurrentQuantityConversionOperation string

const (
	ContractProductGetResponseDataCurrentQuantityConversionOperationMultiply ContractProductGetResponseDataCurrentQuantityConversionOperation = "multiply"
	ContractProductGetResponseDataCurrentQuantityConversionOperationDivide   ContractProductGetResponseDataCurrentQuantityConversionOperation = "divide"
	ContractProductGetResponseDataCurrentQuantityConversionOperationMultiply ContractProductGetResponseDataCurrentQuantityConversionOperation = "MULTIPLY"
	ContractProductGetResponseDataCurrentQuantityConversionOperationDivide   ContractProductGetResponseDataCurrentQuantityConversionOperation = "DIVIDE"
)

func (r ContractProductGetResponseDataCurrentQuantityConversionOperation) IsKnown() bool {
	switch r {
	case ContractProductGetResponseDataCurrentQuantityConversionOperationMultiply, ContractProductGetResponseDataCurrentQuantityConversionOperationDivide, ContractProductGetResponseDataCurrentQuantityConversionOperationMultiply, ContractProductGetResponseDataCurrentQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type ContractProductGetResponseDataCurrentQuantityRounding struct {
	DecimalPlaces  float64                                                             `json:"decimal_places,required"`
	RoundingMethod ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           contractProductGetResponseDataCurrentQuantityRoundingJSON           `json:"-"`
}

// contractProductGetResponseDataCurrentQuantityRoundingJSON contains the JSON
// metadata for the struct [ContractProductGetResponseDataCurrentQuantityRounding]
type contractProductGetResponseDataCurrentQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractProductGetResponseDataCurrentQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductGetResponseDataCurrentQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod string

const (
	ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundUp     ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod = "round_up"
	ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundDown   ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod = "round_down"
	ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundHalfUp ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod = "round_half_up"
	ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundUp     ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod = "ROUND_UP"
	ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundDown   ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod = "ROUND_DOWN"
	ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundHalfUp ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundUp, ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundDown, ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundHalfUp, ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundUp, ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundDown, ContractProductGetResponseDataCurrentQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
}

type ContractProductGetResponseDataInitial struct {
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
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion ContractProductGetResponseDataInitialQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding ContractProductGetResponseDataInitialQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                             `json:"starting_at" format:"date-time"`
	Tags             []string                                              `json:"tags"`
	JSON             contractProductGetResponseDataInitialJSON             `json:"-"`
}

// contractProductGetResponseDataInitialJSON contains the JSON metadata for the
// struct [ContractProductGetResponseDataInitial]
type contractProductGetResponseDataInitialJSON struct {
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

func (r *ContractProductGetResponseDataInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductGetResponseDataInitialJSON) RawJSON() string {
	return r.raw
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type ContractProductGetResponseDataInitialQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation ContractProductGetResponseDataInitialQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                      `json:"name"`
	JSON contractProductGetResponseDataInitialQuantityConversionJSON `json:"-"`
}

// contractProductGetResponseDataInitialQuantityConversionJSON contains the JSON
// metadata for the struct
// [ContractProductGetResponseDataInitialQuantityConversion]
type contractProductGetResponseDataInitialQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ContractProductGetResponseDataInitialQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductGetResponseDataInitialQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type ContractProductGetResponseDataInitialQuantityConversionOperation string

const (
	ContractProductGetResponseDataInitialQuantityConversionOperationMultiply ContractProductGetResponseDataInitialQuantityConversionOperation = "multiply"
	ContractProductGetResponseDataInitialQuantityConversionOperationDivide   ContractProductGetResponseDataInitialQuantityConversionOperation = "divide"
	ContractProductGetResponseDataInitialQuantityConversionOperationMultiply ContractProductGetResponseDataInitialQuantityConversionOperation = "MULTIPLY"
	ContractProductGetResponseDataInitialQuantityConversionOperationDivide   ContractProductGetResponseDataInitialQuantityConversionOperation = "DIVIDE"
)

func (r ContractProductGetResponseDataInitialQuantityConversionOperation) IsKnown() bool {
	switch r {
	case ContractProductGetResponseDataInitialQuantityConversionOperationMultiply, ContractProductGetResponseDataInitialQuantityConversionOperationDivide, ContractProductGetResponseDataInitialQuantityConversionOperationMultiply, ContractProductGetResponseDataInitialQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type ContractProductGetResponseDataInitialQuantityRounding struct {
	DecimalPlaces  float64                                                             `json:"decimal_places,required"`
	RoundingMethod ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           contractProductGetResponseDataInitialQuantityRoundingJSON           `json:"-"`
}

// contractProductGetResponseDataInitialQuantityRoundingJSON contains the JSON
// metadata for the struct [ContractProductGetResponseDataInitialQuantityRounding]
type contractProductGetResponseDataInitialQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractProductGetResponseDataInitialQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductGetResponseDataInitialQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod string

const (
	ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundUp     ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod = "round_up"
	ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundDown   ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod = "round_down"
	ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundHalfUp ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod = "round_half_up"
	ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundUp     ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod = "ROUND_UP"
	ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundDown   ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod = "ROUND_DOWN"
	ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundHalfUp ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r ContractProductGetResponseDataInitialQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundUp, ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundDown, ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundHalfUp, ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundUp, ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundDown, ContractProductGetResponseDataInitialQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
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
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion ContractProductGetResponseDataUpdatesQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding ContractProductGetResponseDataUpdatesQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                             `json:"starting_at" format:"date-time"`
	Tags             []string                                              `json:"tags"`
	JSON             contractProductGetResponseDataUpdateJSON              `json:"-"`
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

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type ContractProductGetResponseDataUpdatesQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation ContractProductGetResponseDataUpdatesQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                      `json:"name"`
	JSON contractProductGetResponseDataUpdatesQuantityConversionJSON `json:"-"`
}

// contractProductGetResponseDataUpdatesQuantityConversionJSON contains the JSON
// metadata for the struct
// [ContractProductGetResponseDataUpdatesQuantityConversion]
type contractProductGetResponseDataUpdatesQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ContractProductGetResponseDataUpdatesQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductGetResponseDataUpdatesQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type ContractProductGetResponseDataUpdatesQuantityConversionOperation string

const (
	ContractProductGetResponseDataUpdatesQuantityConversionOperationMultiply ContractProductGetResponseDataUpdatesQuantityConversionOperation = "multiply"
	ContractProductGetResponseDataUpdatesQuantityConversionOperationDivide   ContractProductGetResponseDataUpdatesQuantityConversionOperation = "divide"
	ContractProductGetResponseDataUpdatesQuantityConversionOperationMultiply ContractProductGetResponseDataUpdatesQuantityConversionOperation = "MULTIPLY"
	ContractProductGetResponseDataUpdatesQuantityConversionOperationDivide   ContractProductGetResponseDataUpdatesQuantityConversionOperation = "DIVIDE"
)

func (r ContractProductGetResponseDataUpdatesQuantityConversionOperation) IsKnown() bool {
	switch r {
	case ContractProductGetResponseDataUpdatesQuantityConversionOperationMultiply, ContractProductGetResponseDataUpdatesQuantityConversionOperationDivide, ContractProductGetResponseDataUpdatesQuantityConversionOperationMultiply, ContractProductGetResponseDataUpdatesQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type ContractProductGetResponseDataUpdatesQuantityRounding struct {
	DecimalPlaces  float64                                                             `json:"decimal_places,required"`
	RoundingMethod ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           contractProductGetResponseDataUpdatesQuantityRoundingJSON           `json:"-"`
}

// contractProductGetResponseDataUpdatesQuantityRoundingJSON contains the JSON
// metadata for the struct [ContractProductGetResponseDataUpdatesQuantityRounding]
type contractProductGetResponseDataUpdatesQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractProductGetResponseDataUpdatesQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductGetResponseDataUpdatesQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod string

const (
	ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundUp     ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod = "round_up"
	ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundDown   ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod = "round_down"
	ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundHalfUp ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod = "round_half_up"
	ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundUp     ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod = "ROUND_UP"
	ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundDown   ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod = "ROUND_DOWN"
	ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundHalfUp ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundUp, ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundDown, ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundHalfUp, ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundUp, ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundDown, ContractProductGetResponseDataUpdatesQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
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
	Current      ContractProductListResponseCurrent  `json:"current,required"`
	Initial      ContractProductListResponseInitial  `json:"initial,required"`
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

type ContractProductListResponseCurrent struct {
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
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion ContractProductListResponseCurrentQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding ContractProductListResponseCurrentQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                          `json:"starting_at" format:"date-time"`
	Tags             []string                                           `json:"tags"`
	JSON             contractProductListResponseCurrentJSON             `json:"-"`
}

// contractProductListResponseCurrentJSON contains the JSON metadata for the struct
// [ContractProductListResponseCurrent]
type contractProductListResponseCurrentJSON struct {
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

func (r *ContractProductListResponseCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductListResponseCurrentJSON) RawJSON() string {
	return r.raw
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type ContractProductListResponseCurrentQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation ContractProductListResponseCurrentQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                   `json:"name"`
	JSON contractProductListResponseCurrentQuantityConversionJSON `json:"-"`
}

// contractProductListResponseCurrentQuantityConversionJSON contains the JSON
// metadata for the struct [ContractProductListResponseCurrentQuantityConversion]
type contractProductListResponseCurrentQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ContractProductListResponseCurrentQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductListResponseCurrentQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type ContractProductListResponseCurrentQuantityConversionOperation string

const (
	ContractProductListResponseCurrentQuantityConversionOperationMultiply ContractProductListResponseCurrentQuantityConversionOperation = "multiply"
	ContractProductListResponseCurrentQuantityConversionOperationDivide   ContractProductListResponseCurrentQuantityConversionOperation = "divide"
	ContractProductListResponseCurrentQuantityConversionOperationMultiply ContractProductListResponseCurrentQuantityConversionOperation = "MULTIPLY"
	ContractProductListResponseCurrentQuantityConversionOperationDivide   ContractProductListResponseCurrentQuantityConversionOperation = "DIVIDE"
)

func (r ContractProductListResponseCurrentQuantityConversionOperation) IsKnown() bool {
	switch r {
	case ContractProductListResponseCurrentQuantityConversionOperationMultiply, ContractProductListResponseCurrentQuantityConversionOperationDivide, ContractProductListResponseCurrentQuantityConversionOperationMultiply, ContractProductListResponseCurrentQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type ContractProductListResponseCurrentQuantityRounding struct {
	DecimalPlaces  float64                                                          `json:"decimal_places,required"`
	RoundingMethod ContractProductListResponseCurrentQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           contractProductListResponseCurrentQuantityRoundingJSON           `json:"-"`
}

// contractProductListResponseCurrentQuantityRoundingJSON contains the JSON
// metadata for the struct [ContractProductListResponseCurrentQuantityRounding]
type contractProductListResponseCurrentQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractProductListResponseCurrentQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductListResponseCurrentQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type ContractProductListResponseCurrentQuantityRoundingRoundingMethod string

const (
	ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundUp     ContractProductListResponseCurrentQuantityRoundingRoundingMethod = "round_up"
	ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundDown   ContractProductListResponseCurrentQuantityRoundingRoundingMethod = "round_down"
	ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundHalfUp ContractProductListResponseCurrentQuantityRoundingRoundingMethod = "round_half_up"
	ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundUp     ContractProductListResponseCurrentQuantityRoundingRoundingMethod = "ROUND_UP"
	ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundDown   ContractProductListResponseCurrentQuantityRoundingRoundingMethod = "ROUND_DOWN"
	ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundHalfUp ContractProductListResponseCurrentQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r ContractProductListResponseCurrentQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundUp, ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundDown, ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundHalfUp, ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundUp, ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundDown, ContractProductListResponseCurrentQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
}

type ContractProductListResponseInitial struct {
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
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion ContractProductListResponseInitialQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding ContractProductListResponseInitialQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                          `json:"starting_at" format:"date-time"`
	Tags             []string                                           `json:"tags"`
	JSON             contractProductListResponseInitialJSON             `json:"-"`
}

// contractProductListResponseInitialJSON contains the JSON metadata for the struct
// [ContractProductListResponseInitial]
type contractProductListResponseInitialJSON struct {
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

func (r *ContractProductListResponseInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductListResponseInitialJSON) RawJSON() string {
	return r.raw
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type ContractProductListResponseInitialQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation ContractProductListResponseInitialQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                   `json:"name"`
	JSON contractProductListResponseInitialQuantityConversionJSON `json:"-"`
}

// contractProductListResponseInitialQuantityConversionJSON contains the JSON
// metadata for the struct [ContractProductListResponseInitialQuantityConversion]
type contractProductListResponseInitialQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ContractProductListResponseInitialQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductListResponseInitialQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type ContractProductListResponseInitialQuantityConversionOperation string

const (
	ContractProductListResponseInitialQuantityConversionOperationMultiply ContractProductListResponseInitialQuantityConversionOperation = "multiply"
	ContractProductListResponseInitialQuantityConversionOperationDivide   ContractProductListResponseInitialQuantityConversionOperation = "divide"
	ContractProductListResponseInitialQuantityConversionOperationMultiply ContractProductListResponseInitialQuantityConversionOperation = "MULTIPLY"
	ContractProductListResponseInitialQuantityConversionOperationDivide   ContractProductListResponseInitialQuantityConversionOperation = "DIVIDE"
)

func (r ContractProductListResponseInitialQuantityConversionOperation) IsKnown() bool {
	switch r {
	case ContractProductListResponseInitialQuantityConversionOperationMultiply, ContractProductListResponseInitialQuantityConversionOperationDivide, ContractProductListResponseInitialQuantityConversionOperationMultiply, ContractProductListResponseInitialQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type ContractProductListResponseInitialQuantityRounding struct {
	DecimalPlaces  float64                                                          `json:"decimal_places,required"`
	RoundingMethod ContractProductListResponseInitialQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           contractProductListResponseInitialQuantityRoundingJSON           `json:"-"`
}

// contractProductListResponseInitialQuantityRoundingJSON contains the JSON
// metadata for the struct [ContractProductListResponseInitialQuantityRounding]
type contractProductListResponseInitialQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractProductListResponseInitialQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductListResponseInitialQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type ContractProductListResponseInitialQuantityRoundingRoundingMethod string

const (
	ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundUp     ContractProductListResponseInitialQuantityRoundingRoundingMethod = "round_up"
	ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundDown   ContractProductListResponseInitialQuantityRoundingRoundingMethod = "round_down"
	ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundHalfUp ContractProductListResponseInitialQuantityRoundingRoundingMethod = "round_half_up"
	ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundUp     ContractProductListResponseInitialQuantityRoundingRoundingMethod = "ROUND_UP"
	ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundDown   ContractProductListResponseInitialQuantityRoundingRoundingMethod = "ROUND_DOWN"
	ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundHalfUp ContractProductListResponseInitialQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r ContractProductListResponseInitialQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundUp, ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundDown, ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundHalfUp, ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundUp, ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundDown, ContractProductListResponseInitialQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
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
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion ContractProductListResponseUpdatesQuantityConversion `json:"quantity_conversion,nullable"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding ContractProductListResponseUpdatesQuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time                                          `json:"starting_at" format:"date-time"`
	Tags             []string                                           `json:"tags"`
	JSON             contractProductListResponseUpdateJSON              `json:"-"`
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

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type ContractProductListResponseUpdatesQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation ContractProductListResponseUpdatesQuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string                                                   `json:"name"`
	JSON contractProductListResponseUpdatesQuantityConversionJSON `json:"-"`
}

// contractProductListResponseUpdatesQuantityConversionJSON contains the JSON
// metadata for the struct [ContractProductListResponseUpdatesQuantityConversion]
type contractProductListResponseUpdatesQuantityConversionJSON struct {
	ConversionFactor apijson.Field
	Operation        apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ContractProductListResponseUpdatesQuantityConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductListResponseUpdatesQuantityConversionJSON) RawJSON() string {
	return r.raw
}

// The operation to perform on the quantity
type ContractProductListResponseUpdatesQuantityConversionOperation string

const (
	ContractProductListResponseUpdatesQuantityConversionOperationMultiply ContractProductListResponseUpdatesQuantityConversionOperation = "multiply"
	ContractProductListResponseUpdatesQuantityConversionOperationDivide   ContractProductListResponseUpdatesQuantityConversionOperation = "divide"
	ContractProductListResponseUpdatesQuantityConversionOperationMultiply ContractProductListResponseUpdatesQuantityConversionOperation = "MULTIPLY"
	ContractProductListResponseUpdatesQuantityConversionOperationDivide   ContractProductListResponseUpdatesQuantityConversionOperation = "DIVIDE"
)

func (r ContractProductListResponseUpdatesQuantityConversionOperation) IsKnown() bool {
	switch r {
	case ContractProductListResponseUpdatesQuantityConversionOperationMultiply, ContractProductListResponseUpdatesQuantityConversionOperationDivide, ContractProductListResponseUpdatesQuantityConversionOperationMultiply, ContractProductListResponseUpdatesQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type ContractProductListResponseUpdatesQuantityRounding struct {
	DecimalPlaces  float64                                                          `json:"decimal_places,required"`
	RoundingMethod ContractProductListResponseUpdatesQuantityRoundingRoundingMethod `json:"rounding_method,required"`
	JSON           contractProductListResponseUpdatesQuantityRoundingJSON           `json:"-"`
}

// contractProductListResponseUpdatesQuantityRoundingJSON contains the JSON
// metadata for the struct [ContractProductListResponseUpdatesQuantityRounding]
type contractProductListResponseUpdatesQuantityRoundingJSON struct {
	DecimalPlaces  apijson.Field
	RoundingMethod apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractProductListResponseUpdatesQuantityRounding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractProductListResponseUpdatesQuantityRoundingJSON) RawJSON() string {
	return r.raw
}

type ContractProductListResponseUpdatesQuantityRoundingRoundingMethod string

const (
	ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundUp     ContractProductListResponseUpdatesQuantityRoundingRoundingMethod = "round_up"
	ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundDown   ContractProductListResponseUpdatesQuantityRoundingRoundingMethod = "round_down"
	ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundHalfUp ContractProductListResponseUpdatesQuantityRoundingRoundingMethod = "round_half_up"
	ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundUp     ContractProductListResponseUpdatesQuantityRoundingRoundingMethod = "ROUND_UP"
	ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundDown   ContractProductListResponseUpdatesQuantityRoundingRoundingMethod = "ROUND_DOWN"
	ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundHalfUp ContractProductListResponseUpdatesQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r ContractProductListResponseUpdatesQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundUp, ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundDown, ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundHalfUp, ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundUp, ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundDown, ContractProductListResponseUpdatesQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
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
	// to true
	IsRefundable param.Field[bool] `json:"is_refundable"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID param.Field[string] `json:"netsuite_internal_item_id"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID param.Field[string] `json:"netsuite_overage_item_id"`
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey param.Field[[]string] `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey param.Field[[]string] `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion param.Field[ContractProductNewParamsQuantityConversion] `json:"quantity_conversion"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding param.Field[ContractProductNewParamsQuantityRounding] `json:"quantity_rounding"`
	Tags             param.Field[[]string]                                 `json:"tags"`
}

func (r ContractProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractProductNewParamsType string

const (
	ContractProductNewParamsTypeFixed               ContractProductNewParamsType = "FIXED"
	ContractProductNewParamsTypeFixed               ContractProductNewParamsType = "fixed"
	ContractProductNewParamsTypeUsage               ContractProductNewParamsType = "USAGE"
	ContractProductNewParamsTypeUsage               ContractProductNewParamsType = "usage"
	ContractProductNewParamsTypeComposite           ContractProductNewParamsType = "COMPOSITE"
	ContractProductNewParamsTypeComposite           ContractProductNewParamsType = "composite"
	ContractProductNewParamsTypeSubscription        ContractProductNewParamsType = "SUBSCRIPTION"
	ContractProductNewParamsTypeSubscription        ContractProductNewParamsType = "subscription"
	ContractProductNewParamsTypeProfessionalService ContractProductNewParamsType = "PROFESSIONAL_SERVICE"
	ContractProductNewParamsTypeProfessionalService ContractProductNewParamsType = "professional_service"
	ContractProductNewParamsTypeProService          ContractProductNewParamsType = "PRO_SERVICE"
	ContractProductNewParamsTypeProService          ContractProductNewParamsType = "pro_service"
)

func (r ContractProductNewParamsType) IsKnown() bool {
	switch r {
	case ContractProductNewParamsTypeFixed, ContractProductNewParamsTypeFixed, ContractProductNewParamsTypeUsage, ContractProductNewParamsTypeUsage, ContractProductNewParamsTypeComposite, ContractProductNewParamsTypeComposite, ContractProductNewParamsTypeSubscription, ContractProductNewParamsTypeSubscription, ContractProductNewParamsTypeProfessionalService, ContractProductNewParamsTypeProfessionalService, ContractProductNewParamsTypeProService, ContractProductNewParamsTypeProService:
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
type ContractProductNewParamsQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor param.Field[float64] `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation param.Field[ContractProductNewParamsQuantityConversionOperation] `json:"operation,required"`
	// Optional name for this conversion.
	Name param.Field[string] `json:"name"`
}

func (r ContractProductNewParamsQuantityConversion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operation to perform on the quantity
type ContractProductNewParamsQuantityConversionOperation string

const (
	ContractProductNewParamsQuantityConversionOperationMultiply ContractProductNewParamsQuantityConversionOperation = "multiply"
	ContractProductNewParamsQuantityConversionOperationDivide   ContractProductNewParamsQuantityConversionOperation = "divide"
	ContractProductNewParamsQuantityConversionOperationMultiply ContractProductNewParamsQuantityConversionOperation = "MULTIPLY"
	ContractProductNewParamsQuantityConversionOperationDivide   ContractProductNewParamsQuantityConversionOperation = "DIVIDE"
)

func (r ContractProductNewParamsQuantityConversionOperation) IsKnown() bool {
	switch r {
	case ContractProductNewParamsQuantityConversionOperationMultiply, ContractProductNewParamsQuantityConversionOperationDivide, ContractProductNewParamsQuantityConversionOperationMultiply, ContractProductNewParamsQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type ContractProductNewParamsQuantityRounding struct {
	DecimalPlaces  param.Field[float64]                                                `json:"decimal_places,required"`
	RoundingMethod param.Field[ContractProductNewParamsQuantityRoundingRoundingMethod] `json:"rounding_method,required"`
}

func (r ContractProductNewParamsQuantityRounding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractProductNewParamsQuantityRoundingRoundingMethod string

const (
	ContractProductNewParamsQuantityRoundingRoundingMethodRoundUp     ContractProductNewParamsQuantityRoundingRoundingMethod = "round_up"
	ContractProductNewParamsQuantityRoundingRoundingMethodRoundDown   ContractProductNewParamsQuantityRoundingRoundingMethod = "round_down"
	ContractProductNewParamsQuantityRoundingRoundingMethodRoundHalfUp ContractProductNewParamsQuantityRoundingRoundingMethod = "round_half_up"
	ContractProductNewParamsQuantityRoundingRoundingMethodRoundUp     ContractProductNewParamsQuantityRoundingRoundingMethod = "ROUND_UP"
	ContractProductNewParamsQuantityRoundingRoundingMethodRoundDown   ContractProductNewParamsQuantityRoundingRoundingMethod = "ROUND_DOWN"
	ContractProductNewParamsQuantityRoundingRoundingMethodRoundHalfUp ContractProductNewParamsQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r ContractProductNewParamsQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case ContractProductNewParamsQuantityRoundingRoundingMethodRoundUp, ContractProductNewParamsQuantityRoundingRoundingMethodRoundDown, ContractProductNewParamsQuantityRoundingRoundingMethodRoundHalfUp, ContractProductNewParamsQuantityRoundingRoundingMethodRoundUp, ContractProductNewParamsQuantityRoundingRoundingMethodRoundDown, ContractProductNewParamsQuantityRoundingRoundingMethodRoundHalfUp:
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
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey param.Field[[]string] `json:"presentation_group_key"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey param.Field[[]string] `json:"pricing_group_key"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion param.Field[ContractProductUpdateParamsQuantityConversion] `json:"quantity_conversion"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding param.Field[ContractProductUpdateParamsQuantityRounding] `json:"quantity_rounding"`
	// If not provided, defaults to product's current tags
	Tags param.Field[[]string] `json:"tags"`
}

func (r ContractProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
type ContractProductUpdateParamsQuantityConversion struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor param.Field[float64] `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	Operation param.Field[ContractProductUpdateParamsQuantityConversionOperation] `json:"operation,required"`
	// Optional name for this conversion.
	Name param.Field[string] `json:"name"`
}

func (r ContractProductUpdateParamsQuantityConversion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operation to perform on the quantity
type ContractProductUpdateParamsQuantityConversionOperation string

const (
	ContractProductUpdateParamsQuantityConversionOperationMultiply ContractProductUpdateParamsQuantityConversionOperation = "multiply"
	ContractProductUpdateParamsQuantityConversionOperationDivide   ContractProductUpdateParamsQuantityConversionOperation = "divide"
	ContractProductUpdateParamsQuantityConversionOperationMultiply ContractProductUpdateParamsQuantityConversionOperation = "MULTIPLY"
	ContractProductUpdateParamsQuantityConversionOperationDivide   ContractProductUpdateParamsQuantityConversionOperation = "DIVIDE"
)

func (r ContractProductUpdateParamsQuantityConversionOperation) IsKnown() bool {
	switch r {
	case ContractProductUpdateParamsQuantityConversionOperationMultiply, ContractProductUpdateParamsQuantityConversionOperationDivide, ContractProductUpdateParamsQuantityConversionOperationMultiply, ContractProductUpdateParamsQuantityConversionOperationDivide:
		return true
	}
	return false
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type ContractProductUpdateParamsQuantityRounding struct {
	DecimalPlaces  param.Field[float64]                                                   `json:"decimal_places,required"`
	RoundingMethod param.Field[ContractProductUpdateParamsQuantityRoundingRoundingMethod] `json:"rounding_method,required"`
}

func (r ContractProductUpdateParamsQuantityRounding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractProductUpdateParamsQuantityRoundingRoundingMethod string

const (
	ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundUp     ContractProductUpdateParamsQuantityRoundingRoundingMethod = "round_up"
	ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundDown   ContractProductUpdateParamsQuantityRoundingRoundingMethod = "round_down"
	ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundHalfUp ContractProductUpdateParamsQuantityRoundingRoundingMethod = "round_half_up"
	ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundUp     ContractProductUpdateParamsQuantityRoundingRoundingMethod = "ROUND_UP"
	ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundDown   ContractProductUpdateParamsQuantityRoundingRoundingMethod = "ROUND_DOWN"
	ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundHalfUp ContractProductUpdateParamsQuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

func (r ContractProductUpdateParamsQuantityRoundingRoundingMethod) IsKnown() bool {
	switch r {
	case ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundUp, ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundDown, ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundHalfUp, ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundUp, ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundDown, ContractProductUpdateParamsQuantityRoundingRoundingMethodRoundHalfUp:
		return true
	}
	return false
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
