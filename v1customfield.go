// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/Metronome-Industries/metronome-go/v2/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v2/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/v2/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/v2/packages/param"
	"github.com/Metronome-Industries/metronome-go/v2/packages/respjson"
)

// V1CustomFieldService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomFieldService] method instead.
type V1CustomFieldService struct {
	Options []option.RequestOption
}

// NewV1CustomFieldService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomFieldService(opts ...option.RequestOption) (r V1CustomFieldService) {
	r = V1CustomFieldService{}
	r.Options = opts
	return
}

// Creates a new custom field key for a given entity (e.g. billable metric,
// contract, alert).
//
// Custom fields are properties that you can add to Metronome objects to store
// metadata like foreign keys or other descriptors. This metadata can get
// transferred to or accessed by other systems to contextualize Metronome data and
// power business processes. For example, to service workflows like revenue
// recognition, reconciliation, and invoicing, custom fields help Metronome know
// the relationship between entities in the platform and third-party systems.
//
// ### Use this endpoint to:
//
//   - Create a new custom field key for Customer objects in Metronome. You can then
//     use the Set Custom Field Values endpoint to set the value of this key for a
//     specific customer.
//   - Specify whether the key should enforce uniqueness. If the key is set to
//     enforce uniqueness and you attempt to set a custom field value for the key
//     that already exists, it will fail.
//
// ### Usage guidelines:
//
//   - Custom fields set on commits, credits, and contracts can be used to scope
//     alert evaluation. For example, you can create a spend threshold alert that
//     only considers spend associated with contracts with custom field key
//     `contract_type` and value `paygo`
//   - Custom fields set on products can be used in the Stripe integration to set
//     metadata on invoices.
//   - Custom fields for customers, contracts, invoices, products, commits, scheduled
//     charges, and subscriptions are passed down to the invoice.
func (r *V1CustomFieldService) AddKey(ctx context.Context, body V1CustomFieldAddKeyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/customFields/addKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Remove specific custom field values from a Metronome entity instance by
// specifying the field keys to delete. Use this endpoint to clean up unwanted
// custom field data while preserving other fields on the same entity. Requires the
// entity type, entity ID, and array of keys to remove.
func (r *V1CustomFieldService) DeleteValues(ctx context.Context, body V1CustomFieldDeleteValuesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/customFields/deleteValues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve all your active custom field keys, with optional filtering by entity
// type (customer, contract, product, etc.). Use this endpoint to discover what
// custom field keys are available before setting values on entities or to audit
// your custom field configuration across different entity types.
func (r *V1CustomFieldService) ListKeys(ctx context.Context, params V1CustomFieldListKeysParams, opts ...option.RequestOption) (res *pagination.CursorPageWithoutLimit[V1CustomFieldListKeysResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/customFields/listKeys"
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

// Retrieve all your active custom field keys, with optional filtering by entity
// type (customer, contract, product, etc.). Use this endpoint to discover what
// custom field keys are available before setting values on entities or to audit
// your custom field configuration across different entity types.
func (r *V1CustomFieldService) ListKeysAutoPaging(ctx context.Context, params V1CustomFieldListKeysParams, opts ...option.RequestOption) *pagination.CursorPageWithoutLimitAutoPager[V1CustomFieldListKeysResponse] {
	return pagination.NewCursorPageWithoutLimitAutoPager(r.ListKeys(ctx, params, opts...))
}

// Removes a custom field key from the allowlist for a specific entity type,
// preventing future use of that key across all instances of the entity. Existing
// values for this key on entity instances will no longer be accessible once the
// key is removed.
func (r *V1CustomFieldService) RemoveKey(ctx context.Context, body V1CustomFieldRemoveKeyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/customFields/removeKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Sets custom field values on a specific Metronome entity instance. Overwrites
// existing values for matching keys while preserving other fields. All updates are
// transactional—either all values are set or none are. Custom field values are
// limited to 200 characters each.
func (r *V1CustomFieldService) SetValues(ctx context.Context, body V1CustomFieldSetValuesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/customFields/setValues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type V1CustomFieldListKeysResponse struct {
	EnforceUniqueness bool `json:"enforce_uniqueness,required"`
	// Any of "alert", "billable_metric", "charge", "commit", "contract_credit",
	// "contract_product", "contract", "credit_grant", "customer_plan", "customer",
	// "discount", "invoice", "plan", "professional_service", "product", "rate_card",
	// "scheduled_charge", "subscription".
	Entity V1CustomFieldListKeysResponseEntity `json:"entity,required"`
	Key    string                              `json:"key,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnforceUniqueness respjson.Field
		Entity            respjson.Field
		Key               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomFieldListKeysResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomFieldListKeysResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomFieldListKeysResponseEntity string

const (
	V1CustomFieldListKeysResponseEntityAlert               V1CustomFieldListKeysResponseEntity = "alert"
	V1CustomFieldListKeysResponseEntityBillableMetric      V1CustomFieldListKeysResponseEntity = "billable_metric"
	V1CustomFieldListKeysResponseEntityCharge              V1CustomFieldListKeysResponseEntity = "charge"
	V1CustomFieldListKeysResponseEntityCommit              V1CustomFieldListKeysResponseEntity = "commit"
	V1CustomFieldListKeysResponseEntityContractCredit      V1CustomFieldListKeysResponseEntity = "contract_credit"
	V1CustomFieldListKeysResponseEntityContractProduct     V1CustomFieldListKeysResponseEntity = "contract_product"
	V1CustomFieldListKeysResponseEntityContract            V1CustomFieldListKeysResponseEntity = "contract"
	V1CustomFieldListKeysResponseEntityCreditGrant         V1CustomFieldListKeysResponseEntity = "credit_grant"
	V1CustomFieldListKeysResponseEntityCustomerPlan        V1CustomFieldListKeysResponseEntity = "customer_plan"
	V1CustomFieldListKeysResponseEntityCustomer            V1CustomFieldListKeysResponseEntity = "customer"
	V1CustomFieldListKeysResponseEntityDiscount            V1CustomFieldListKeysResponseEntity = "discount"
	V1CustomFieldListKeysResponseEntityInvoice             V1CustomFieldListKeysResponseEntity = "invoice"
	V1CustomFieldListKeysResponseEntityPlan                V1CustomFieldListKeysResponseEntity = "plan"
	V1CustomFieldListKeysResponseEntityProfessionalService V1CustomFieldListKeysResponseEntity = "professional_service"
	V1CustomFieldListKeysResponseEntityProduct             V1CustomFieldListKeysResponseEntity = "product"
	V1CustomFieldListKeysResponseEntityRateCard            V1CustomFieldListKeysResponseEntity = "rate_card"
	V1CustomFieldListKeysResponseEntityScheduledCharge     V1CustomFieldListKeysResponseEntity = "scheduled_charge"
	V1CustomFieldListKeysResponseEntitySubscription        V1CustomFieldListKeysResponseEntity = "subscription"
)

type V1CustomFieldAddKeyParams struct {
	EnforceUniqueness bool `json:"enforce_uniqueness,required"`
	// Any of "alert", "billable_metric", "charge", "commit", "contract_credit",
	// "contract_product", "contract", "credit_grant", "customer_plan", "customer",
	// "discount", "invoice", "plan", "professional_service", "product", "rate_card",
	// "scheduled_charge", "subscription".
	Entity V1CustomFieldAddKeyParamsEntity `json:"entity,omitzero,required"`
	Key    string                          `json:"key,required"`
	paramObj
}

func (r V1CustomFieldAddKeyParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomFieldAddKeyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomFieldAddKeyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomFieldAddKeyParamsEntity string

const (
	V1CustomFieldAddKeyParamsEntityAlert               V1CustomFieldAddKeyParamsEntity = "alert"
	V1CustomFieldAddKeyParamsEntityBillableMetric      V1CustomFieldAddKeyParamsEntity = "billable_metric"
	V1CustomFieldAddKeyParamsEntityCharge              V1CustomFieldAddKeyParamsEntity = "charge"
	V1CustomFieldAddKeyParamsEntityCommit              V1CustomFieldAddKeyParamsEntity = "commit"
	V1CustomFieldAddKeyParamsEntityContractCredit      V1CustomFieldAddKeyParamsEntity = "contract_credit"
	V1CustomFieldAddKeyParamsEntityContractProduct     V1CustomFieldAddKeyParamsEntity = "contract_product"
	V1CustomFieldAddKeyParamsEntityContract            V1CustomFieldAddKeyParamsEntity = "contract"
	V1CustomFieldAddKeyParamsEntityCreditGrant         V1CustomFieldAddKeyParamsEntity = "credit_grant"
	V1CustomFieldAddKeyParamsEntityCustomerPlan        V1CustomFieldAddKeyParamsEntity = "customer_plan"
	V1CustomFieldAddKeyParamsEntityCustomer            V1CustomFieldAddKeyParamsEntity = "customer"
	V1CustomFieldAddKeyParamsEntityDiscount            V1CustomFieldAddKeyParamsEntity = "discount"
	V1CustomFieldAddKeyParamsEntityInvoice             V1CustomFieldAddKeyParamsEntity = "invoice"
	V1CustomFieldAddKeyParamsEntityPlan                V1CustomFieldAddKeyParamsEntity = "plan"
	V1CustomFieldAddKeyParamsEntityProfessionalService V1CustomFieldAddKeyParamsEntity = "professional_service"
	V1CustomFieldAddKeyParamsEntityProduct             V1CustomFieldAddKeyParamsEntity = "product"
	V1CustomFieldAddKeyParamsEntityRateCard            V1CustomFieldAddKeyParamsEntity = "rate_card"
	V1CustomFieldAddKeyParamsEntityScheduledCharge     V1CustomFieldAddKeyParamsEntity = "scheduled_charge"
	V1CustomFieldAddKeyParamsEntitySubscription        V1CustomFieldAddKeyParamsEntity = "subscription"
)

type V1CustomFieldDeleteValuesParams struct {
	// Any of "alert", "billable_metric", "charge", "commit", "contract_credit",
	// "contract_product", "contract", "credit_grant", "customer_plan", "customer",
	// "discount", "invoice", "plan", "professional_service", "product", "rate_card",
	// "scheduled_charge", "subscription".
	Entity   V1CustomFieldDeleteValuesParamsEntity `json:"entity,omitzero,required"`
	EntityID string                                `json:"entity_id,required" format:"uuid"`
	Keys     []string                              `json:"keys,omitzero,required"`
	paramObj
}

func (r V1CustomFieldDeleteValuesParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomFieldDeleteValuesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomFieldDeleteValuesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomFieldDeleteValuesParamsEntity string

const (
	V1CustomFieldDeleteValuesParamsEntityAlert               V1CustomFieldDeleteValuesParamsEntity = "alert"
	V1CustomFieldDeleteValuesParamsEntityBillableMetric      V1CustomFieldDeleteValuesParamsEntity = "billable_metric"
	V1CustomFieldDeleteValuesParamsEntityCharge              V1CustomFieldDeleteValuesParamsEntity = "charge"
	V1CustomFieldDeleteValuesParamsEntityCommit              V1CustomFieldDeleteValuesParamsEntity = "commit"
	V1CustomFieldDeleteValuesParamsEntityContractCredit      V1CustomFieldDeleteValuesParamsEntity = "contract_credit"
	V1CustomFieldDeleteValuesParamsEntityContractProduct     V1CustomFieldDeleteValuesParamsEntity = "contract_product"
	V1CustomFieldDeleteValuesParamsEntityContract            V1CustomFieldDeleteValuesParamsEntity = "contract"
	V1CustomFieldDeleteValuesParamsEntityCreditGrant         V1CustomFieldDeleteValuesParamsEntity = "credit_grant"
	V1CustomFieldDeleteValuesParamsEntityCustomerPlan        V1CustomFieldDeleteValuesParamsEntity = "customer_plan"
	V1CustomFieldDeleteValuesParamsEntityCustomer            V1CustomFieldDeleteValuesParamsEntity = "customer"
	V1CustomFieldDeleteValuesParamsEntityDiscount            V1CustomFieldDeleteValuesParamsEntity = "discount"
	V1CustomFieldDeleteValuesParamsEntityInvoice             V1CustomFieldDeleteValuesParamsEntity = "invoice"
	V1CustomFieldDeleteValuesParamsEntityPlan                V1CustomFieldDeleteValuesParamsEntity = "plan"
	V1CustomFieldDeleteValuesParamsEntityProfessionalService V1CustomFieldDeleteValuesParamsEntity = "professional_service"
	V1CustomFieldDeleteValuesParamsEntityProduct             V1CustomFieldDeleteValuesParamsEntity = "product"
	V1CustomFieldDeleteValuesParamsEntityRateCard            V1CustomFieldDeleteValuesParamsEntity = "rate_card"
	V1CustomFieldDeleteValuesParamsEntityScheduledCharge     V1CustomFieldDeleteValuesParamsEntity = "scheduled_charge"
	V1CustomFieldDeleteValuesParamsEntitySubscription        V1CustomFieldDeleteValuesParamsEntity = "subscription"
)

type V1CustomFieldListKeysParams struct {
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// Optional list of entity types to return keys for
	//
	// Any of "alert", "billable_metric", "charge", "commit", "contract_credit",
	// "contract_product", "contract", "credit_grant", "customer_plan", "customer",
	// "discount", "invoice", "plan", "professional_service", "product", "rate_card",
	// "scheduled_charge", "subscription".
	Entities []string `json:"entities,omitzero"`
	paramObj
}

func (r V1CustomFieldListKeysParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomFieldListKeysParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomFieldListKeysParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1CustomFieldListKeysParams]'s query parameters as
// `url.Values`.
func (r V1CustomFieldListKeysParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomFieldRemoveKeyParams struct {
	// Any of "alert", "billable_metric", "charge", "commit", "contract_credit",
	// "contract_product", "contract", "credit_grant", "customer_plan", "customer",
	// "discount", "invoice", "plan", "professional_service", "product", "rate_card",
	// "scheduled_charge", "subscription".
	Entity V1CustomFieldRemoveKeyParamsEntity `json:"entity,omitzero,required"`
	Key    string                             `json:"key,required"`
	paramObj
}

func (r V1CustomFieldRemoveKeyParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomFieldRemoveKeyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomFieldRemoveKeyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomFieldRemoveKeyParamsEntity string

const (
	V1CustomFieldRemoveKeyParamsEntityAlert               V1CustomFieldRemoveKeyParamsEntity = "alert"
	V1CustomFieldRemoveKeyParamsEntityBillableMetric      V1CustomFieldRemoveKeyParamsEntity = "billable_metric"
	V1CustomFieldRemoveKeyParamsEntityCharge              V1CustomFieldRemoveKeyParamsEntity = "charge"
	V1CustomFieldRemoveKeyParamsEntityCommit              V1CustomFieldRemoveKeyParamsEntity = "commit"
	V1CustomFieldRemoveKeyParamsEntityContractCredit      V1CustomFieldRemoveKeyParamsEntity = "contract_credit"
	V1CustomFieldRemoveKeyParamsEntityContractProduct     V1CustomFieldRemoveKeyParamsEntity = "contract_product"
	V1CustomFieldRemoveKeyParamsEntityContract            V1CustomFieldRemoveKeyParamsEntity = "contract"
	V1CustomFieldRemoveKeyParamsEntityCreditGrant         V1CustomFieldRemoveKeyParamsEntity = "credit_grant"
	V1CustomFieldRemoveKeyParamsEntityCustomerPlan        V1CustomFieldRemoveKeyParamsEntity = "customer_plan"
	V1CustomFieldRemoveKeyParamsEntityCustomer            V1CustomFieldRemoveKeyParamsEntity = "customer"
	V1CustomFieldRemoveKeyParamsEntityDiscount            V1CustomFieldRemoveKeyParamsEntity = "discount"
	V1CustomFieldRemoveKeyParamsEntityInvoice             V1CustomFieldRemoveKeyParamsEntity = "invoice"
	V1CustomFieldRemoveKeyParamsEntityPlan                V1CustomFieldRemoveKeyParamsEntity = "plan"
	V1CustomFieldRemoveKeyParamsEntityProfessionalService V1CustomFieldRemoveKeyParamsEntity = "professional_service"
	V1CustomFieldRemoveKeyParamsEntityProduct             V1CustomFieldRemoveKeyParamsEntity = "product"
	V1CustomFieldRemoveKeyParamsEntityRateCard            V1CustomFieldRemoveKeyParamsEntity = "rate_card"
	V1CustomFieldRemoveKeyParamsEntityScheduledCharge     V1CustomFieldRemoveKeyParamsEntity = "scheduled_charge"
	V1CustomFieldRemoveKeyParamsEntitySubscription        V1CustomFieldRemoveKeyParamsEntity = "subscription"
)

type V1CustomFieldSetValuesParams struct {
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero,required"`
	// Any of "alert", "billable_metric", "charge", "commit", "contract_credit",
	// "contract_product", "contract", "credit_grant", "customer_plan", "customer",
	// "discount", "invoice", "plan", "professional_service", "product", "rate_card",
	// "scheduled_charge", "subscription".
	Entity   V1CustomFieldSetValuesParamsEntity `json:"entity,omitzero,required"`
	EntityID string                             `json:"entity_id,required" format:"uuid"`
	paramObj
}

func (r V1CustomFieldSetValuesParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomFieldSetValuesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomFieldSetValuesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomFieldSetValuesParamsEntity string

const (
	V1CustomFieldSetValuesParamsEntityAlert               V1CustomFieldSetValuesParamsEntity = "alert"
	V1CustomFieldSetValuesParamsEntityBillableMetric      V1CustomFieldSetValuesParamsEntity = "billable_metric"
	V1CustomFieldSetValuesParamsEntityCharge              V1CustomFieldSetValuesParamsEntity = "charge"
	V1CustomFieldSetValuesParamsEntityCommit              V1CustomFieldSetValuesParamsEntity = "commit"
	V1CustomFieldSetValuesParamsEntityContractCredit      V1CustomFieldSetValuesParamsEntity = "contract_credit"
	V1CustomFieldSetValuesParamsEntityContractProduct     V1CustomFieldSetValuesParamsEntity = "contract_product"
	V1CustomFieldSetValuesParamsEntityContract            V1CustomFieldSetValuesParamsEntity = "contract"
	V1CustomFieldSetValuesParamsEntityCreditGrant         V1CustomFieldSetValuesParamsEntity = "credit_grant"
	V1CustomFieldSetValuesParamsEntityCustomerPlan        V1CustomFieldSetValuesParamsEntity = "customer_plan"
	V1CustomFieldSetValuesParamsEntityCustomer            V1CustomFieldSetValuesParamsEntity = "customer"
	V1CustomFieldSetValuesParamsEntityDiscount            V1CustomFieldSetValuesParamsEntity = "discount"
	V1CustomFieldSetValuesParamsEntityInvoice             V1CustomFieldSetValuesParamsEntity = "invoice"
	V1CustomFieldSetValuesParamsEntityPlan                V1CustomFieldSetValuesParamsEntity = "plan"
	V1CustomFieldSetValuesParamsEntityProfessionalService V1CustomFieldSetValuesParamsEntity = "professional_service"
	V1CustomFieldSetValuesParamsEntityProduct             V1CustomFieldSetValuesParamsEntity = "product"
	V1CustomFieldSetValuesParamsEntityRateCard            V1CustomFieldSetValuesParamsEntity = "rate_card"
	V1CustomFieldSetValuesParamsEntityScheduledCharge     V1CustomFieldSetValuesParamsEntity = "scheduled_charge"
	V1CustomFieldSetValuesParamsEntitySubscription        V1CustomFieldSetValuesParamsEntity = "subscription"
)
