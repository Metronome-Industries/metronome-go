// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	shimjson "github.com/Metronome-Industries/metronome-go/internal/encoding/json"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
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
func NewV1ContractProductService(opts ...option.RequestOption) (r V1ContractProductService) {
	r = V1ContractProductService{}
	r.Options = opts
	return
}

// Create a new product object. Products in Metronome represent your company's
// individual product or service offerings. A Product can be thought of as the
// basic unit of a line item on the invoice. This is analogous to SKUs or items in
// an ERP system. Give the product a meaningful name as they will appear on
// customer invoices.
func (r *V1ContractProductService) New(ctx context.Context, body V1ContractProductNewParams, opts ...option.RequestOption) (res *V1ContractProductNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/products/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a product by its ID, including all metadata and historical changes.
func (r *V1ContractProductService) Get(ctx context.Context, body V1ContractProductGetParams, opts ...option.RequestOption) (res *V1ContractProductGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/products/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates a product's configuration while maintaining billing continuity for
// active customers. Use this endpoint to modify product names, metrics, pricing
// rules, and composite settings without disrupting ongoing billing cycles. Changes
// are scheduled using the starting_at timestamp, which must be on an hour
// boundary—set future dates to schedule updates ahead of time, or past dates for
// retroactive changes. Returns the updated product ID upon success.
//
// ### Usage guidance:
//
//   - Product type cannot be changed after creation. For incorrect product types,
//     create a new product and archive the original instead.
func (r *V1ContractProductService) Update(ctx context.Context, body V1ContractProductUpdateParams, opts ...option.RequestOption) (res *V1ContractProductUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/products/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a paginated list of all products in your organization with their complete
// configuration, version history, and metadata. By default excludes archived
// products unless explicitly requested via the `archive_filter` parameter.
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

// Get a paginated list of all products in your organization with their complete
// configuration, version history, and metadata. By default excludes archived
// products unless explicitly requested via the `archive_filter` parameter.
func (r *V1ContractProductService) ListAutoPaging(ctx context.Context, params V1ContractProductListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1ContractProductListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Archive a product. Any current rate cards associated with this product will
// continue to function as normal. However, it will no longer be available as an
// option for newly created rates. Once you archive a product, you can still
// retrieve it in the UI and API, but you cannot unarchive it.
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
	QuantityRounding QuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time        `json:"starting_at" format:"date-time"`
	Tags             []string         `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		Name                   respjson.Field
		BillableMetricID       respjson.Field
		CompositeProductIDs    respjson.Field
		CompositeTags          respjson.Field
		ExcludeFreeUsage       respjson.Field
		IsRefundable           respjson.Field
		NetsuiteInternalItemID respjson.Field
		NetsuiteOverageItemID  respjson.Field
		PresentationGroupKey   respjson.Field
		PricingGroupKey        respjson.Field
		QuantityConversion     respjson.Field
		QuantityRounding       respjson.Field
		StartingAt             respjson.Field
		Tags                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductListItemState) RawJSON() string { return r.JSON.raw }
func (r *ProductListItemState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	//
	// Any of "MULTIPLY", "DIVIDE".
	Operation QuantityConversionOperation `json:"operation,required"`
	// Optional name for this conversion.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversionFactor respjson.Field
		Operation        respjson.Field
		Name             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuantityConversion) RawJSON() string { return r.JSON.raw }
func (r *QuantityConversion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this QuantityConversion to a QuantityConversionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// QuantityConversionParam.Overrides()
func (r QuantityConversion) ToParam() QuantityConversionParam {
	return param.Override[QuantityConversionParam](json.RawMessage(r.RawJSON()))
}

// The operation to perform on the quantity
type QuantityConversionOperation string

const (
	QuantityConversionOperationMultiply QuantityConversionOperation = "MULTIPLY"
	QuantityConversionOperationDivide   QuantityConversionOperation = "DIVIDE"
)

// Optional. Only valid for USAGE products. If provided, the quantity will be
// converted using the provided conversion factor and operation. For example, if
// the operation is "multiply" and the conversion factor is 100, then the quantity
// will be multiplied by 100. This can be used in cases where data is sent in one
// unit and priced in another. For example, data could be sent in MB and priced in
// GB. In this case, the conversion factor would be 1024 and the operation would be
// "divide".
//
// The properties ConversionFactor, Operation are required.
type QuantityConversionParam struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor,required"`
	// The operation to perform on the quantity
	//
	// Any of "MULTIPLY", "DIVIDE".
	Operation QuantityConversionOperation `json:"operation,omitzero,required"`
	// Optional name for this conversion.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r QuantityConversionParam) MarshalJSON() (data []byte, err error) {
	type shadow QuantityConversionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuantityConversionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
type QuantityRounding struct {
	DecimalPlaces float64 `json:"decimal_places,required"`
	// Any of "ROUND_UP", "ROUND_DOWN", "ROUND_HALF_UP".
	RoundingMethod QuantityRoundingRoundingMethod `json:"rounding_method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuantityRounding) RawJSON() string { return r.JSON.raw }
func (r *QuantityRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this QuantityRounding to a QuantityRoundingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// QuantityRoundingParam.Overrides()
func (r QuantityRounding) ToParam() QuantityRoundingParam {
	return param.Override[QuantityRoundingParam](json.RawMessage(r.RawJSON()))
}

type QuantityRoundingRoundingMethod string

const (
	QuantityRoundingRoundingMethodRoundUp     QuantityRoundingRoundingMethod = "ROUND_UP"
	QuantityRoundingRoundingMethodRoundDown   QuantityRoundingRoundingMethod = "ROUND_DOWN"
	QuantityRoundingRoundingMethodRoundHalfUp QuantityRoundingRoundingMethod = "ROUND_HALF_UP"
)

// Optional. Only valid for USAGE products. If provided, the quantity will be
// rounded using the provided rounding method and decimal places. For example, if
// the method is "round up" and the decimal places is 0, then the quantity will be
// rounded up to the nearest integer.
//
// The properties DecimalPlaces, RoundingMethod are required.
type QuantityRoundingParam struct {
	DecimalPlaces float64 `json:"decimal_places,required"`
	// Any of "ROUND_UP", "ROUND_DOWN", "ROUND_HALF_UP".
	RoundingMethod QuantityRoundingRoundingMethod `json:"rounding_method,omitzero,required"`
	paramObj
}

func (r QuantityRoundingParam) MarshalJSON() (data []byte, err error) {
	type shadow QuantityRoundingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuantityRoundingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractProductNewResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractProductNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractProductNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractProductGetResponse struct {
	Data V1ContractProductGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractProductGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractProductGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractProductGetResponseData struct {
	ID      string               `json:"id,required" format:"uuid"`
	Current ProductListItemState `json:"current,required"`
	Initial ProductListItemState `json:"initial,required"`
	// Any of "USAGE", "SUBSCRIPTION", "COMPOSITE", "FIXED", "PRO_SERVICE".
	Type       string                                   `json:"type,required"`
	Updates    []V1ContractProductGetResponseDataUpdate `json:"updates,required"`
	ArchivedAt time.Time                                `json:"archived_at,nullable" format:"date-time"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Current      respjson.Field
		Initial      respjson.Field
		Type         respjson.Field
		Updates      respjson.Field
		ArchivedAt   respjson.Field
		CustomFields respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractProductGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ContractProductGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	QuantityRounding QuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time        `json:"starting_at" format:"date-time"`
	Tags             []string         `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		BillableMetricID       respjson.Field
		CompositeProductIDs    respjson.Field
		CompositeTags          respjson.Field
		ExcludeFreeUsage       respjson.Field
		IsRefundable           respjson.Field
		Name                   respjson.Field
		NetsuiteInternalItemID respjson.Field
		NetsuiteOverageItemID  respjson.Field
		PresentationGroupKey   respjson.Field
		PricingGroupKey        respjson.Field
		QuantityConversion     respjson.Field
		QuantityRounding       respjson.Field
		StartingAt             respjson.Field
		Tags                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractProductGetResponseDataUpdate) RawJSON() string { return r.JSON.raw }
func (r *V1ContractProductGetResponseDataUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractProductUpdateResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractProductUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractProductUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractProductListResponse struct {
	ID      string               `json:"id,required" format:"uuid"`
	Current ProductListItemState `json:"current,required"`
	Initial ProductListItemState `json:"initial,required"`
	// Any of "USAGE", "SUBSCRIPTION", "COMPOSITE", "FIXED", "PRO_SERVICE".
	Type       V1ContractProductListResponseType     `json:"type,required"`
	Updates    []V1ContractProductListResponseUpdate `json:"updates,required"`
	ArchivedAt time.Time                             `json:"archived_at,nullable" format:"date-time"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Current      respjson.Field
		Initial      respjson.Field
		Type         respjson.Field
		Updates      respjson.Field
		ArchivedAt   respjson.Field
		CustomFields respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractProductListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractProductListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractProductListResponseType string

const (
	V1ContractProductListResponseTypeUsage        V1ContractProductListResponseType = "USAGE"
	V1ContractProductListResponseTypeSubscription V1ContractProductListResponseType = "SUBSCRIPTION"
	V1ContractProductListResponseTypeComposite    V1ContractProductListResponseType = "COMPOSITE"
	V1ContractProductListResponseTypeFixed        V1ContractProductListResponseType = "FIXED"
	V1ContractProductListResponseTypeProService   V1ContractProductListResponseType = "PRO_SERVICE"
)

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
	QuantityRounding QuantityRounding `json:"quantity_rounding,nullable"`
	StartingAt       time.Time        `json:"starting_at" format:"date-time"`
	Tags             []string         `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		BillableMetricID       respjson.Field
		CompositeProductIDs    respjson.Field
		CompositeTags          respjson.Field
		ExcludeFreeUsage       respjson.Field
		IsRefundable           respjson.Field
		Name                   respjson.Field
		NetsuiteInternalItemID respjson.Field
		NetsuiteOverageItemID  respjson.Field
		PresentationGroupKey   respjson.Field
		PricingGroupKey        respjson.Field
		QuantityConversion     respjson.Field
		QuantityRounding       respjson.Field
		StartingAt             respjson.Field
		Tags                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractProductListResponseUpdate) RawJSON() string { return r.JSON.raw }
func (r *V1ContractProductListResponseUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractProductArchiveResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractProductArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractProductArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractProductNewParams struct {
	// displayed on invoices
	Name string `json:"name,required"`
	// Any of "FIXED", "USAGE", "COMPOSITE", "SUBSCRIPTION", "PROFESSIONAL_SERVICE",
	// "PRO_SERVICE".
	Type V1ContractProductNewParamsType `json:"type,omitzero,required"`
	// Required for USAGE products
	BillableMetricID param.Opt[string] `json:"billable_metric_id,omitzero" format:"uuid"`
	// Beta feature only available for composite products. If true, products with $0
	// will not be included when computing composite usage. Defaults to false
	ExcludeFreeUsage param.Opt[bool] `json:"exclude_free_usage,omitzero"`
	// This field's availability is dependent on your client's configuration. Defaults
	// to true.
	IsRefundable param.Opt[bool] `json:"is_refundable,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID param.Opt[string] `json:"netsuite_internal_item_id,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID param.Opt[string] `json:"netsuite_overage_item_id,omitzero"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion QuantityConversionParam `json:"quantity_conversion,omitzero"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding QuantityRoundingParam `json:"quantity_rounding,omitzero"`
	// Required for COMPOSITE products
	CompositeProductIDs []string `json:"composite_product_ids,omitzero" format:"uuid"`
	// Required for COMPOSITE products
	CompositeTags []string `json:"composite_tags,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// For USAGE products only. Groups usage line items on invoices. The superset of
	// values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PresentationGroupKey []string `json:"presentation_group_key,omitzero"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole. The superset
	// of values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PricingGroupKey []string `json:"pricing_group_key,omitzero"`
	Tags            []string `json:"tags,omitzero"`
	paramObj
}

func (r V1ContractProductNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractProductNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractProductNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type V1ContractProductGetParams struct {
	ID shared.IDParam
	paramObj
}

func (r V1ContractProductGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ID)
}
func (r *V1ContractProductGetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ID)
}

type V1ContractProductUpdateParams struct {
	// ID of the product to update
	ProductID string `json:"product_id,required" format:"uuid"`
	// Timestamp representing when the update should go into effect. It must be on an
	// hour boundary (e.g. 1:00, not 1:30).
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Available for USAGE products only. If not provided, defaults to product's
	// current billable metric.
	BillableMetricID param.Opt[string] `json:"billable_metric_id,omitzero" format:"uuid"`
	// Beta feature only available for composite products. If true, products with $0
	// will not be included when computing composite usage. Defaults to false
	ExcludeFreeUsage param.Opt[bool] `json:"exclude_free_usage,omitzero"`
	// Defaults to product's current refundability status. This field's availability is
	// dependent on your client's configuration.
	IsRefundable param.Opt[bool] `json:"is_refundable,omitzero"`
	// displayed on invoices. If not provided, defaults to product's current name.
	Name param.Opt[string] `json:"name,omitzero"`
	// If not provided, defaults to product's current netsuite_internal_item_id. This
	// field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID param.Opt[string] `json:"netsuite_internal_item_id,omitzero"`
	// Available for USAGE and COMPOSITE products only. If not provided, defaults to
	// product's current netsuite_overage_item_id. This field's availability is
	// dependent on your client's configuration.
	NetsuiteOverageItemID param.Opt[string] `json:"netsuite_overage_item_id,omitzero"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// converted using the provided conversion factor and operation. For example, if
	// the operation is "multiply" and the conversion factor is 100, then the quantity
	// will be multiplied by 100. This can be used in cases where data is sent in one
	// unit and priced in another. For example, data could be sent in MB and priced in
	// GB. In this case, the conversion factor would be 1024 and the operation would be
	// "divide".
	QuantityConversion QuantityConversionParam `json:"quantity_conversion,omitzero"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be
	// rounded using the provided rounding method and decimal places. For example, if
	// the method is "round up" and the decimal places is 0, then the quantity will be
	// rounded up to the nearest integer.
	QuantityRounding QuantityRoundingParam `json:"quantity_rounding,omitzero"`
	// Available for COMPOSITE products only. If not provided, defaults to product's
	// current composite_product_ids.
	CompositeProductIDs []string `json:"composite_product_ids,omitzero" format:"uuid"`
	// Available for COMPOSITE products only. If not provided, defaults to product's
	// current composite_tags.
	CompositeTags []string `json:"composite_tags,omitzero"`
	// For USAGE products only. Groups usage line items on invoices. The superset of
	// values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PresentationGroupKey []string `json:"presentation_group_key,omitzero"`
	// For USAGE products only. If set, pricing for this product will be determined for
	// each pricing_group_key value, as opposed to the product as a whole. The superset
	// of values in the pricing group key and presentation group key must be set as one
	// compound group key on the billable metric.
	PricingGroupKey []string `json:"pricing_group_key,omitzero"`
	// If not provided, defaults to product's current tags
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r V1ContractProductUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractProductUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractProductUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractProductListParams struct {
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// Filter options for the product list. If not provided, defaults to not archived.
	//
	// Any of "ARCHIVED", "NOT_ARCHIVED", "ALL".
	ArchiveFilter V1ContractProductListParamsArchiveFilter `json:"archive_filter,omitzero"`
	paramObj
}

func (r V1ContractProductListParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractProductListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractProductListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1ContractProductListParams]'s query parameters as
// `url.Values`.
func (r V1ContractProductListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter options for the product list. If not provided, defaults to not archived.
type V1ContractProductListParamsArchiveFilter string

const (
	V1ContractProductListParamsArchiveFilterArchived    V1ContractProductListParamsArchiveFilter = "ARCHIVED"
	V1ContractProductListParamsArchiveFilterNotArchived V1ContractProductListParamsArchiveFilter = "NOT_ARCHIVED"
	V1ContractProductListParamsArchiveFilterAll         V1ContractProductListParamsArchiveFilter = "ALL"
)

type V1ContractProductArchiveParams struct {
	// ID of the product to be archived
	ProductID string `json:"product_id,required" format:"uuid"`
	paramObj
}

func (r V1ContractProductArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractProductArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractProductArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
