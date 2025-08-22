// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
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
func NewV1CustomFieldService(opts ...option.RequestOption) (r *V1CustomFieldService) {
	r = &V1CustomFieldService{}
	r.Options = opts
	return
}

// Add a key to the allow list for a given entity. There is a 100 character limit
// on custom field keys.
func (r *V1CustomFieldService) AddKey(ctx context.Context, body V1CustomFieldAddKeyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/customFields/addKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Deletes one or more custom fields on an instance of a Metronome entity.
func (r *V1CustomFieldService) DeleteValues(ctx context.Context, body V1CustomFieldDeleteValuesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/customFields/deleteValues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// List all active custom field keys, optionally filtered by entity type.
func (r *V1CustomFieldService) ListKeys(ctx context.Context, params V1CustomFieldListKeysParams, opts ...option.RequestOption) (res *pagination.CursorPageWithoutLimit[V1CustomFieldListKeysResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// List all active custom field keys, optionally filtered by entity type.
func (r *V1CustomFieldService) ListKeysAutoPaging(ctx context.Context, params V1CustomFieldListKeysParams, opts ...option.RequestOption) *pagination.CursorPageWithoutLimitAutoPager[V1CustomFieldListKeysResponse] {
	return pagination.NewCursorPageWithoutLimitAutoPager(r.ListKeys(ctx, params, opts...))
}

// Remove a key from the allow list for a given entity.
func (r *V1CustomFieldService) RemoveKey(ctx context.Context, body V1CustomFieldRemoveKeyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/customFields/removeKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Sets one or more custom fields on an instance of a Metronome entity. If a
// key/value pair passed in this request matches one already set on the entity, its
// value will be overwritten. Any key/value pairs that exist on the entity that do
// not match those passed in this request will remain untouched. This endpoint is
// transactional and will update all key/value pairs or no key/value pairs. Partial
// updates are not supported. There is a 200 character limit on custom field
// values.
func (r *V1CustomFieldService) SetValues(ctx context.Context, body V1CustomFieldSetValuesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/customFields/setValues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type V1CustomFieldListKeysResponse struct {
	EnforceUniqueness bool                                `json:"enforce_uniqueness,required"`
	Entity            V1CustomFieldListKeysResponseEntity `json:"entity,required"`
	Key               string                              `json:"key,required"`
	JSON              v1CustomFieldListKeysResponseJSON   `json:"-"`
}

// v1CustomFieldListKeysResponseJSON contains the JSON metadata for the struct
// [V1CustomFieldListKeysResponse]
type v1CustomFieldListKeysResponseJSON struct {
	EnforceUniqueness apijson.Field
	Entity            apijson.Field
	Key               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *V1CustomFieldListKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomFieldListKeysResponseJSON) RawJSON() string {
	return r.raw
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

func (r V1CustomFieldListKeysResponseEntity) IsKnown() bool {
	switch r {
	case V1CustomFieldListKeysResponseEntityAlert, V1CustomFieldListKeysResponseEntityBillableMetric, V1CustomFieldListKeysResponseEntityCharge, V1CustomFieldListKeysResponseEntityCommit, V1CustomFieldListKeysResponseEntityContractCredit, V1CustomFieldListKeysResponseEntityContractProduct, V1CustomFieldListKeysResponseEntityContract, V1CustomFieldListKeysResponseEntityCreditGrant, V1CustomFieldListKeysResponseEntityCustomerPlan, V1CustomFieldListKeysResponseEntityCustomer, V1CustomFieldListKeysResponseEntityDiscount, V1CustomFieldListKeysResponseEntityInvoice, V1CustomFieldListKeysResponseEntityPlan, V1CustomFieldListKeysResponseEntityProfessionalService, V1CustomFieldListKeysResponseEntityProduct, V1CustomFieldListKeysResponseEntityRateCard, V1CustomFieldListKeysResponseEntityScheduledCharge, V1CustomFieldListKeysResponseEntitySubscription:
		return true
	}
	return false
}

type V1CustomFieldAddKeyParams struct {
	EnforceUniqueness param.Field[bool]                            `json:"enforce_uniqueness,required"`
	Entity            param.Field[V1CustomFieldAddKeyParamsEntity] `json:"entity,required"`
	Key               param.Field[string]                          `json:"key,required"`
}

func (r V1CustomFieldAddKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

func (r V1CustomFieldAddKeyParamsEntity) IsKnown() bool {
	switch r {
	case V1CustomFieldAddKeyParamsEntityAlert, V1CustomFieldAddKeyParamsEntityBillableMetric, V1CustomFieldAddKeyParamsEntityCharge, V1CustomFieldAddKeyParamsEntityCommit, V1CustomFieldAddKeyParamsEntityContractCredit, V1CustomFieldAddKeyParamsEntityContractProduct, V1CustomFieldAddKeyParamsEntityContract, V1CustomFieldAddKeyParamsEntityCreditGrant, V1CustomFieldAddKeyParamsEntityCustomerPlan, V1CustomFieldAddKeyParamsEntityCustomer, V1CustomFieldAddKeyParamsEntityDiscount, V1CustomFieldAddKeyParamsEntityInvoice, V1CustomFieldAddKeyParamsEntityPlan, V1CustomFieldAddKeyParamsEntityProfessionalService, V1CustomFieldAddKeyParamsEntityProduct, V1CustomFieldAddKeyParamsEntityRateCard, V1CustomFieldAddKeyParamsEntityScheduledCharge, V1CustomFieldAddKeyParamsEntitySubscription:
		return true
	}
	return false
}

type V1CustomFieldDeleteValuesParams struct {
	Entity   param.Field[V1CustomFieldDeleteValuesParamsEntity] `json:"entity,required"`
	EntityID param.Field[string]                                `json:"entity_id,required" format:"uuid"`
	Keys     param.Field[[]string]                              `json:"keys,required"`
}

func (r V1CustomFieldDeleteValuesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

func (r V1CustomFieldDeleteValuesParamsEntity) IsKnown() bool {
	switch r {
	case V1CustomFieldDeleteValuesParamsEntityAlert, V1CustomFieldDeleteValuesParamsEntityBillableMetric, V1CustomFieldDeleteValuesParamsEntityCharge, V1CustomFieldDeleteValuesParamsEntityCommit, V1CustomFieldDeleteValuesParamsEntityContractCredit, V1CustomFieldDeleteValuesParamsEntityContractProduct, V1CustomFieldDeleteValuesParamsEntityContract, V1CustomFieldDeleteValuesParamsEntityCreditGrant, V1CustomFieldDeleteValuesParamsEntityCustomerPlan, V1CustomFieldDeleteValuesParamsEntityCustomer, V1CustomFieldDeleteValuesParamsEntityDiscount, V1CustomFieldDeleteValuesParamsEntityInvoice, V1CustomFieldDeleteValuesParamsEntityPlan, V1CustomFieldDeleteValuesParamsEntityProfessionalService, V1CustomFieldDeleteValuesParamsEntityProduct, V1CustomFieldDeleteValuesParamsEntityRateCard, V1CustomFieldDeleteValuesParamsEntityScheduledCharge, V1CustomFieldDeleteValuesParamsEntitySubscription:
		return true
	}
	return false
}

type V1CustomFieldListKeysParams struct {
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Optional list of entity types to return keys for
	Entities param.Field[[]V1CustomFieldListKeysParamsEntity] `json:"entities"`
}

func (r V1CustomFieldListKeysParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V1CustomFieldListKeysParams]'s query parameters as
// `url.Values`.
func (r V1CustomFieldListKeysParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomFieldListKeysParamsEntity string

const (
	V1CustomFieldListKeysParamsEntityAlert               V1CustomFieldListKeysParamsEntity = "alert"
	V1CustomFieldListKeysParamsEntityBillableMetric      V1CustomFieldListKeysParamsEntity = "billable_metric"
	V1CustomFieldListKeysParamsEntityCharge              V1CustomFieldListKeysParamsEntity = "charge"
	V1CustomFieldListKeysParamsEntityCommit              V1CustomFieldListKeysParamsEntity = "commit"
	V1CustomFieldListKeysParamsEntityContractCredit      V1CustomFieldListKeysParamsEntity = "contract_credit"
	V1CustomFieldListKeysParamsEntityContractProduct     V1CustomFieldListKeysParamsEntity = "contract_product"
	V1CustomFieldListKeysParamsEntityContract            V1CustomFieldListKeysParamsEntity = "contract"
	V1CustomFieldListKeysParamsEntityCreditGrant         V1CustomFieldListKeysParamsEntity = "credit_grant"
	V1CustomFieldListKeysParamsEntityCustomerPlan        V1CustomFieldListKeysParamsEntity = "customer_plan"
	V1CustomFieldListKeysParamsEntityCustomer            V1CustomFieldListKeysParamsEntity = "customer"
	V1CustomFieldListKeysParamsEntityDiscount            V1CustomFieldListKeysParamsEntity = "discount"
	V1CustomFieldListKeysParamsEntityInvoice             V1CustomFieldListKeysParamsEntity = "invoice"
	V1CustomFieldListKeysParamsEntityPlan                V1CustomFieldListKeysParamsEntity = "plan"
	V1CustomFieldListKeysParamsEntityProfessionalService V1CustomFieldListKeysParamsEntity = "professional_service"
	V1CustomFieldListKeysParamsEntityProduct             V1CustomFieldListKeysParamsEntity = "product"
	V1CustomFieldListKeysParamsEntityRateCard            V1CustomFieldListKeysParamsEntity = "rate_card"
	V1CustomFieldListKeysParamsEntityScheduledCharge     V1CustomFieldListKeysParamsEntity = "scheduled_charge"
	V1CustomFieldListKeysParamsEntitySubscription        V1CustomFieldListKeysParamsEntity = "subscription"
)

func (r V1CustomFieldListKeysParamsEntity) IsKnown() bool {
	switch r {
	case V1CustomFieldListKeysParamsEntityAlert, V1CustomFieldListKeysParamsEntityBillableMetric, V1CustomFieldListKeysParamsEntityCharge, V1CustomFieldListKeysParamsEntityCommit, V1CustomFieldListKeysParamsEntityContractCredit, V1CustomFieldListKeysParamsEntityContractProduct, V1CustomFieldListKeysParamsEntityContract, V1CustomFieldListKeysParamsEntityCreditGrant, V1CustomFieldListKeysParamsEntityCustomerPlan, V1CustomFieldListKeysParamsEntityCustomer, V1CustomFieldListKeysParamsEntityDiscount, V1CustomFieldListKeysParamsEntityInvoice, V1CustomFieldListKeysParamsEntityPlan, V1CustomFieldListKeysParamsEntityProfessionalService, V1CustomFieldListKeysParamsEntityProduct, V1CustomFieldListKeysParamsEntityRateCard, V1CustomFieldListKeysParamsEntityScheduledCharge, V1CustomFieldListKeysParamsEntitySubscription:
		return true
	}
	return false
}

type V1CustomFieldRemoveKeyParams struct {
	Entity param.Field[V1CustomFieldRemoveKeyParamsEntity] `json:"entity,required"`
	Key    param.Field[string]                             `json:"key,required"`
}

func (r V1CustomFieldRemoveKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

func (r V1CustomFieldRemoveKeyParamsEntity) IsKnown() bool {
	switch r {
	case V1CustomFieldRemoveKeyParamsEntityAlert, V1CustomFieldRemoveKeyParamsEntityBillableMetric, V1CustomFieldRemoveKeyParamsEntityCharge, V1CustomFieldRemoveKeyParamsEntityCommit, V1CustomFieldRemoveKeyParamsEntityContractCredit, V1CustomFieldRemoveKeyParamsEntityContractProduct, V1CustomFieldRemoveKeyParamsEntityContract, V1CustomFieldRemoveKeyParamsEntityCreditGrant, V1CustomFieldRemoveKeyParamsEntityCustomerPlan, V1CustomFieldRemoveKeyParamsEntityCustomer, V1CustomFieldRemoveKeyParamsEntityDiscount, V1CustomFieldRemoveKeyParamsEntityInvoice, V1CustomFieldRemoveKeyParamsEntityPlan, V1CustomFieldRemoveKeyParamsEntityProfessionalService, V1CustomFieldRemoveKeyParamsEntityProduct, V1CustomFieldRemoveKeyParamsEntityRateCard, V1CustomFieldRemoveKeyParamsEntityScheduledCharge, V1CustomFieldRemoveKeyParamsEntitySubscription:
		return true
	}
	return false
}

type V1CustomFieldSetValuesParams struct {
	CustomFields param.Field[map[string]string]                  `json:"custom_fields,required"`
	Entity       param.Field[V1CustomFieldSetValuesParamsEntity] `json:"entity,required"`
	EntityID     param.Field[string]                             `json:"entity_id,required" format:"uuid"`
}

func (r V1CustomFieldSetValuesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

func (r V1CustomFieldSetValuesParamsEntity) IsKnown() bool {
	switch r {
	case V1CustomFieldSetValuesParamsEntityAlert, V1CustomFieldSetValuesParamsEntityBillableMetric, V1CustomFieldSetValuesParamsEntityCharge, V1CustomFieldSetValuesParamsEntityCommit, V1CustomFieldSetValuesParamsEntityContractCredit, V1CustomFieldSetValuesParamsEntityContractProduct, V1CustomFieldSetValuesParamsEntityContract, V1CustomFieldSetValuesParamsEntityCreditGrant, V1CustomFieldSetValuesParamsEntityCustomerPlan, V1CustomFieldSetValuesParamsEntityCustomer, V1CustomFieldSetValuesParamsEntityDiscount, V1CustomFieldSetValuesParamsEntityInvoice, V1CustomFieldSetValuesParamsEntityPlan, V1CustomFieldSetValuesParamsEntityProfessionalService, V1CustomFieldSetValuesParamsEntityProduct, V1CustomFieldSetValuesParamsEntityRateCard, V1CustomFieldSetValuesParamsEntityScheduledCharge, V1CustomFieldSetValuesParamsEntitySubscription:
		return true
	}
	return false
}
