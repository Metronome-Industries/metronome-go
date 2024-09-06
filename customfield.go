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
)

// CustomFieldService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomFieldService] method instead.
type CustomFieldService struct {
	Options []option.RequestOption
}

// NewCustomFieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomFieldService(opts ...option.RequestOption) (r *CustomFieldService) {
	r = &CustomFieldService{}
	r.Options = opts
	return
}

// Add a key to the allow list for a given entity. There is a 100 character limit
// on custom field keys.
func (r *CustomFieldService) AddKey(ctx context.Context, body CustomFieldAddKeyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "customFields/addKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Deletes one or more custom fields on an instance of a Metronome entity.
func (r *CustomFieldService) DeleteValues(ctx context.Context, body CustomFieldDeleteValuesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "customFields/deleteValues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// List all active custom field keys, optionally filtered by entity type.
func (r *CustomFieldService) ListKeys(ctx context.Context, params CustomFieldListKeysParams, opts ...option.RequestOption) (res *CustomFieldListKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customFields/listKeys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Remove a key from the allow list for a given entity.
func (r *CustomFieldService) RemoveKey(ctx context.Context, body CustomFieldRemoveKeyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "customFields/removeKey"
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
func (r *CustomFieldService) SetValues(ctx context.Context, body CustomFieldSetValuesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "customFields/setValues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CustomFieldListKeysResponse struct {
	Data     []CustomFieldListKeysResponseData `json:"data,required"`
	NextPage string                            `json:"next_page,required,nullable"`
	JSON     customFieldListKeysResponseJSON   `json:"-"`
}

// customFieldListKeysResponseJSON contains the JSON metadata for the struct
// [CustomFieldListKeysResponse]
type customFieldListKeysResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomFieldListKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customFieldListKeysResponseJSON) RawJSON() string {
	return r.raw
}

type CustomFieldListKeysResponseData struct {
	EnforceUniqueness bool                                  `json:"enforce_uniqueness,required"`
	Entity            CustomFieldListKeysResponseDataEntity `json:"entity,required"`
	Key               string                                `json:"key,required"`
	JSON              customFieldListKeysResponseDataJSON   `json:"-"`
}

// customFieldListKeysResponseDataJSON contains the JSON metadata for the struct
// [CustomFieldListKeysResponseData]
type customFieldListKeysResponseDataJSON struct {
	EnforceUniqueness apijson.Field
	Entity            apijson.Field
	Key               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomFieldListKeysResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customFieldListKeysResponseDataJSON) RawJSON() string {
	return r.raw
}

type CustomFieldListKeysResponseDataEntity string

const (
	CustomFieldListKeysResponseDataEntityAlert               CustomFieldListKeysResponseDataEntity = "alert"
	CustomFieldListKeysResponseDataEntityBillableMetric      CustomFieldListKeysResponseDataEntity = "billable_metric"
	CustomFieldListKeysResponseDataEntityCharge              CustomFieldListKeysResponseDataEntity = "charge"
	CustomFieldListKeysResponseDataEntityCommit              CustomFieldListKeysResponseDataEntity = "commit"
	CustomFieldListKeysResponseDataEntityContractCredit      CustomFieldListKeysResponseDataEntity = "contract_credit"
	CustomFieldListKeysResponseDataEntityContractProduct     CustomFieldListKeysResponseDataEntity = "contract_product"
	CustomFieldListKeysResponseDataEntityContract            CustomFieldListKeysResponseDataEntity = "contract"
	CustomFieldListKeysResponseDataEntityCreditGrant         CustomFieldListKeysResponseDataEntity = "credit_grant"
	CustomFieldListKeysResponseDataEntityCustomerPlan        CustomFieldListKeysResponseDataEntity = "customer_plan"
	CustomFieldListKeysResponseDataEntityCustomer            CustomFieldListKeysResponseDataEntity = "customer"
	CustomFieldListKeysResponseDataEntityInvoice             CustomFieldListKeysResponseDataEntity = "invoice"
	CustomFieldListKeysResponseDataEntityPlan                CustomFieldListKeysResponseDataEntity = "plan"
	CustomFieldListKeysResponseDataEntityProfessionalService CustomFieldListKeysResponseDataEntity = "professional_service"
	CustomFieldListKeysResponseDataEntityProduct             CustomFieldListKeysResponseDataEntity = "product"
	CustomFieldListKeysResponseDataEntityRateCard            CustomFieldListKeysResponseDataEntity = "rate_card"
	CustomFieldListKeysResponseDataEntityScheduledCharge     CustomFieldListKeysResponseDataEntity = "scheduled_charge"
)

func (r CustomFieldListKeysResponseDataEntity) IsKnown() bool {
	switch r {
	case CustomFieldListKeysResponseDataEntityAlert, CustomFieldListKeysResponseDataEntityBillableMetric, CustomFieldListKeysResponseDataEntityCharge, CustomFieldListKeysResponseDataEntityCommit, CustomFieldListKeysResponseDataEntityContractCredit, CustomFieldListKeysResponseDataEntityContractProduct, CustomFieldListKeysResponseDataEntityContract, CustomFieldListKeysResponseDataEntityCreditGrant, CustomFieldListKeysResponseDataEntityCustomerPlan, CustomFieldListKeysResponseDataEntityCustomer, CustomFieldListKeysResponseDataEntityInvoice, CustomFieldListKeysResponseDataEntityPlan, CustomFieldListKeysResponseDataEntityProfessionalService, CustomFieldListKeysResponseDataEntityProduct, CustomFieldListKeysResponseDataEntityRateCard, CustomFieldListKeysResponseDataEntityScheduledCharge:
		return true
	}
	return false
}

type CustomFieldAddKeyParams struct {
	EnforceUniqueness param.Field[bool]                          `json:"enforce_uniqueness,required"`
	Entity            param.Field[CustomFieldAddKeyParamsEntity] `json:"entity,required"`
	Key               param.Field[string]                        `json:"key,required"`
}

func (r CustomFieldAddKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomFieldAddKeyParamsEntity string

const (
	CustomFieldAddKeyParamsEntityAlert               CustomFieldAddKeyParamsEntity = "alert"
	CustomFieldAddKeyParamsEntityBillableMetric      CustomFieldAddKeyParamsEntity = "billable_metric"
	CustomFieldAddKeyParamsEntityCharge              CustomFieldAddKeyParamsEntity = "charge"
	CustomFieldAddKeyParamsEntityCommit              CustomFieldAddKeyParamsEntity = "commit"
	CustomFieldAddKeyParamsEntityContractCredit      CustomFieldAddKeyParamsEntity = "contract_credit"
	CustomFieldAddKeyParamsEntityContractProduct     CustomFieldAddKeyParamsEntity = "contract_product"
	CustomFieldAddKeyParamsEntityContract            CustomFieldAddKeyParamsEntity = "contract"
	CustomFieldAddKeyParamsEntityCreditGrant         CustomFieldAddKeyParamsEntity = "credit_grant"
	CustomFieldAddKeyParamsEntityCustomerPlan        CustomFieldAddKeyParamsEntity = "customer_plan"
	CustomFieldAddKeyParamsEntityCustomer            CustomFieldAddKeyParamsEntity = "customer"
	CustomFieldAddKeyParamsEntityInvoice             CustomFieldAddKeyParamsEntity = "invoice"
	CustomFieldAddKeyParamsEntityPlan                CustomFieldAddKeyParamsEntity = "plan"
	CustomFieldAddKeyParamsEntityProfessionalService CustomFieldAddKeyParamsEntity = "professional_service"
	CustomFieldAddKeyParamsEntityProduct             CustomFieldAddKeyParamsEntity = "product"
	CustomFieldAddKeyParamsEntityRateCard            CustomFieldAddKeyParamsEntity = "rate_card"
	CustomFieldAddKeyParamsEntityScheduledCharge     CustomFieldAddKeyParamsEntity = "scheduled_charge"
)

func (r CustomFieldAddKeyParamsEntity) IsKnown() bool {
	switch r {
	case CustomFieldAddKeyParamsEntityAlert, CustomFieldAddKeyParamsEntityBillableMetric, CustomFieldAddKeyParamsEntityCharge, CustomFieldAddKeyParamsEntityCommit, CustomFieldAddKeyParamsEntityContractCredit, CustomFieldAddKeyParamsEntityContractProduct, CustomFieldAddKeyParamsEntityContract, CustomFieldAddKeyParamsEntityCreditGrant, CustomFieldAddKeyParamsEntityCustomerPlan, CustomFieldAddKeyParamsEntityCustomer, CustomFieldAddKeyParamsEntityInvoice, CustomFieldAddKeyParamsEntityPlan, CustomFieldAddKeyParamsEntityProfessionalService, CustomFieldAddKeyParamsEntityProduct, CustomFieldAddKeyParamsEntityRateCard, CustomFieldAddKeyParamsEntityScheduledCharge:
		return true
	}
	return false
}

type CustomFieldDeleteValuesParams struct {
	Entity   param.Field[CustomFieldDeleteValuesParamsEntity] `json:"entity,required"`
	EntityID param.Field[string]                              `json:"entity_id,required" format:"uuid"`
	Keys     param.Field[[]string]                            `json:"keys,required"`
}

func (r CustomFieldDeleteValuesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomFieldDeleteValuesParamsEntity string

const (
	CustomFieldDeleteValuesParamsEntityAlert               CustomFieldDeleteValuesParamsEntity = "alert"
	CustomFieldDeleteValuesParamsEntityBillableMetric      CustomFieldDeleteValuesParamsEntity = "billable_metric"
	CustomFieldDeleteValuesParamsEntityCharge              CustomFieldDeleteValuesParamsEntity = "charge"
	CustomFieldDeleteValuesParamsEntityCommit              CustomFieldDeleteValuesParamsEntity = "commit"
	CustomFieldDeleteValuesParamsEntityContractCredit      CustomFieldDeleteValuesParamsEntity = "contract_credit"
	CustomFieldDeleteValuesParamsEntityContractProduct     CustomFieldDeleteValuesParamsEntity = "contract_product"
	CustomFieldDeleteValuesParamsEntityContract            CustomFieldDeleteValuesParamsEntity = "contract"
	CustomFieldDeleteValuesParamsEntityCreditGrant         CustomFieldDeleteValuesParamsEntity = "credit_grant"
	CustomFieldDeleteValuesParamsEntityCustomerPlan        CustomFieldDeleteValuesParamsEntity = "customer_plan"
	CustomFieldDeleteValuesParamsEntityCustomer            CustomFieldDeleteValuesParamsEntity = "customer"
	CustomFieldDeleteValuesParamsEntityInvoice             CustomFieldDeleteValuesParamsEntity = "invoice"
	CustomFieldDeleteValuesParamsEntityPlan                CustomFieldDeleteValuesParamsEntity = "plan"
	CustomFieldDeleteValuesParamsEntityProfessionalService CustomFieldDeleteValuesParamsEntity = "professional_service"
	CustomFieldDeleteValuesParamsEntityProduct             CustomFieldDeleteValuesParamsEntity = "product"
	CustomFieldDeleteValuesParamsEntityRateCard            CustomFieldDeleteValuesParamsEntity = "rate_card"
	CustomFieldDeleteValuesParamsEntityScheduledCharge     CustomFieldDeleteValuesParamsEntity = "scheduled_charge"
)

func (r CustomFieldDeleteValuesParamsEntity) IsKnown() bool {
	switch r {
	case CustomFieldDeleteValuesParamsEntityAlert, CustomFieldDeleteValuesParamsEntityBillableMetric, CustomFieldDeleteValuesParamsEntityCharge, CustomFieldDeleteValuesParamsEntityCommit, CustomFieldDeleteValuesParamsEntityContractCredit, CustomFieldDeleteValuesParamsEntityContractProduct, CustomFieldDeleteValuesParamsEntityContract, CustomFieldDeleteValuesParamsEntityCreditGrant, CustomFieldDeleteValuesParamsEntityCustomerPlan, CustomFieldDeleteValuesParamsEntityCustomer, CustomFieldDeleteValuesParamsEntityInvoice, CustomFieldDeleteValuesParamsEntityPlan, CustomFieldDeleteValuesParamsEntityProfessionalService, CustomFieldDeleteValuesParamsEntityProduct, CustomFieldDeleteValuesParamsEntityRateCard, CustomFieldDeleteValuesParamsEntityScheduledCharge:
		return true
	}
	return false
}

type CustomFieldListKeysParams struct {
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Optional list of entity types to return keys for
	Entities param.Field[[]CustomFieldListKeysParamsEntity] `json:"entities"`
}

func (r CustomFieldListKeysParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [CustomFieldListKeysParams]'s query parameters as
// `url.Values`.
func (r CustomFieldListKeysParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomFieldListKeysParamsEntity string

const (
	CustomFieldListKeysParamsEntityAlert               CustomFieldListKeysParamsEntity = "alert"
	CustomFieldListKeysParamsEntityBillableMetric      CustomFieldListKeysParamsEntity = "billable_metric"
	CustomFieldListKeysParamsEntityCharge              CustomFieldListKeysParamsEntity = "charge"
	CustomFieldListKeysParamsEntityCommit              CustomFieldListKeysParamsEntity = "commit"
	CustomFieldListKeysParamsEntityContractCredit      CustomFieldListKeysParamsEntity = "contract_credit"
	CustomFieldListKeysParamsEntityContractProduct     CustomFieldListKeysParamsEntity = "contract_product"
	CustomFieldListKeysParamsEntityContract            CustomFieldListKeysParamsEntity = "contract"
	CustomFieldListKeysParamsEntityCreditGrant         CustomFieldListKeysParamsEntity = "credit_grant"
	CustomFieldListKeysParamsEntityCustomerPlan        CustomFieldListKeysParamsEntity = "customer_plan"
	CustomFieldListKeysParamsEntityCustomer            CustomFieldListKeysParamsEntity = "customer"
	CustomFieldListKeysParamsEntityInvoice             CustomFieldListKeysParamsEntity = "invoice"
	CustomFieldListKeysParamsEntityPlan                CustomFieldListKeysParamsEntity = "plan"
	CustomFieldListKeysParamsEntityProfessionalService CustomFieldListKeysParamsEntity = "professional_service"
	CustomFieldListKeysParamsEntityProduct             CustomFieldListKeysParamsEntity = "product"
	CustomFieldListKeysParamsEntityRateCard            CustomFieldListKeysParamsEntity = "rate_card"
	CustomFieldListKeysParamsEntityScheduledCharge     CustomFieldListKeysParamsEntity = "scheduled_charge"
)

func (r CustomFieldListKeysParamsEntity) IsKnown() bool {
	switch r {
	case CustomFieldListKeysParamsEntityAlert, CustomFieldListKeysParamsEntityBillableMetric, CustomFieldListKeysParamsEntityCharge, CustomFieldListKeysParamsEntityCommit, CustomFieldListKeysParamsEntityContractCredit, CustomFieldListKeysParamsEntityContractProduct, CustomFieldListKeysParamsEntityContract, CustomFieldListKeysParamsEntityCreditGrant, CustomFieldListKeysParamsEntityCustomerPlan, CustomFieldListKeysParamsEntityCustomer, CustomFieldListKeysParamsEntityInvoice, CustomFieldListKeysParamsEntityPlan, CustomFieldListKeysParamsEntityProfessionalService, CustomFieldListKeysParamsEntityProduct, CustomFieldListKeysParamsEntityRateCard, CustomFieldListKeysParamsEntityScheduledCharge:
		return true
	}
	return false
}

type CustomFieldRemoveKeyParams struct {
	Entity param.Field[CustomFieldRemoveKeyParamsEntity] `json:"entity,required"`
	Key    param.Field[string]                           `json:"key,required"`
}

func (r CustomFieldRemoveKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomFieldRemoveKeyParamsEntity string

const (
	CustomFieldRemoveKeyParamsEntityAlert               CustomFieldRemoveKeyParamsEntity = "alert"
	CustomFieldRemoveKeyParamsEntityBillableMetric      CustomFieldRemoveKeyParamsEntity = "billable_metric"
	CustomFieldRemoveKeyParamsEntityCharge              CustomFieldRemoveKeyParamsEntity = "charge"
	CustomFieldRemoveKeyParamsEntityCommit              CustomFieldRemoveKeyParamsEntity = "commit"
	CustomFieldRemoveKeyParamsEntityContractCredit      CustomFieldRemoveKeyParamsEntity = "contract_credit"
	CustomFieldRemoveKeyParamsEntityContractProduct     CustomFieldRemoveKeyParamsEntity = "contract_product"
	CustomFieldRemoveKeyParamsEntityContract            CustomFieldRemoveKeyParamsEntity = "contract"
	CustomFieldRemoveKeyParamsEntityCreditGrant         CustomFieldRemoveKeyParamsEntity = "credit_grant"
	CustomFieldRemoveKeyParamsEntityCustomerPlan        CustomFieldRemoveKeyParamsEntity = "customer_plan"
	CustomFieldRemoveKeyParamsEntityCustomer            CustomFieldRemoveKeyParamsEntity = "customer"
	CustomFieldRemoveKeyParamsEntityInvoice             CustomFieldRemoveKeyParamsEntity = "invoice"
	CustomFieldRemoveKeyParamsEntityPlan                CustomFieldRemoveKeyParamsEntity = "plan"
	CustomFieldRemoveKeyParamsEntityProfessionalService CustomFieldRemoveKeyParamsEntity = "professional_service"
	CustomFieldRemoveKeyParamsEntityProduct             CustomFieldRemoveKeyParamsEntity = "product"
	CustomFieldRemoveKeyParamsEntityRateCard            CustomFieldRemoveKeyParamsEntity = "rate_card"
	CustomFieldRemoveKeyParamsEntityScheduledCharge     CustomFieldRemoveKeyParamsEntity = "scheduled_charge"
)

func (r CustomFieldRemoveKeyParamsEntity) IsKnown() bool {
	switch r {
	case CustomFieldRemoveKeyParamsEntityAlert, CustomFieldRemoveKeyParamsEntityBillableMetric, CustomFieldRemoveKeyParamsEntityCharge, CustomFieldRemoveKeyParamsEntityCommit, CustomFieldRemoveKeyParamsEntityContractCredit, CustomFieldRemoveKeyParamsEntityContractProduct, CustomFieldRemoveKeyParamsEntityContract, CustomFieldRemoveKeyParamsEntityCreditGrant, CustomFieldRemoveKeyParamsEntityCustomerPlan, CustomFieldRemoveKeyParamsEntityCustomer, CustomFieldRemoveKeyParamsEntityInvoice, CustomFieldRemoveKeyParamsEntityPlan, CustomFieldRemoveKeyParamsEntityProfessionalService, CustomFieldRemoveKeyParamsEntityProduct, CustomFieldRemoveKeyParamsEntityRateCard, CustomFieldRemoveKeyParamsEntityScheduledCharge:
		return true
	}
	return false
}

type CustomFieldSetValuesParams struct {
	CustomFields param.Field[map[string]string]                `json:"custom_fields,required"`
	Entity       param.Field[CustomFieldSetValuesParamsEntity] `json:"entity,required"`
	EntityID     param.Field[string]                           `json:"entity_id,required" format:"uuid"`
}

func (r CustomFieldSetValuesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomFieldSetValuesParamsEntity string

const (
	CustomFieldSetValuesParamsEntityAlert               CustomFieldSetValuesParamsEntity = "alert"
	CustomFieldSetValuesParamsEntityBillableMetric      CustomFieldSetValuesParamsEntity = "billable_metric"
	CustomFieldSetValuesParamsEntityCharge              CustomFieldSetValuesParamsEntity = "charge"
	CustomFieldSetValuesParamsEntityCommit              CustomFieldSetValuesParamsEntity = "commit"
	CustomFieldSetValuesParamsEntityContractCredit      CustomFieldSetValuesParamsEntity = "contract_credit"
	CustomFieldSetValuesParamsEntityContractProduct     CustomFieldSetValuesParamsEntity = "contract_product"
	CustomFieldSetValuesParamsEntityContract            CustomFieldSetValuesParamsEntity = "contract"
	CustomFieldSetValuesParamsEntityCreditGrant         CustomFieldSetValuesParamsEntity = "credit_grant"
	CustomFieldSetValuesParamsEntityCustomerPlan        CustomFieldSetValuesParamsEntity = "customer_plan"
	CustomFieldSetValuesParamsEntityCustomer            CustomFieldSetValuesParamsEntity = "customer"
	CustomFieldSetValuesParamsEntityInvoice             CustomFieldSetValuesParamsEntity = "invoice"
	CustomFieldSetValuesParamsEntityPlan                CustomFieldSetValuesParamsEntity = "plan"
	CustomFieldSetValuesParamsEntityProfessionalService CustomFieldSetValuesParamsEntity = "professional_service"
	CustomFieldSetValuesParamsEntityProduct             CustomFieldSetValuesParamsEntity = "product"
	CustomFieldSetValuesParamsEntityRateCard            CustomFieldSetValuesParamsEntity = "rate_card"
	CustomFieldSetValuesParamsEntityScheduledCharge     CustomFieldSetValuesParamsEntity = "scheduled_charge"
)

func (r CustomFieldSetValuesParamsEntity) IsKnown() bool {
	switch r {
	case CustomFieldSetValuesParamsEntityAlert, CustomFieldSetValuesParamsEntityBillableMetric, CustomFieldSetValuesParamsEntityCharge, CustomFieldSetValuesParamsEntityCommit, CustomFieldSetValuesParamsEntityContractCredit, CustomFieldSetValuesParamsEntityContractProduct, CustomFieldSetValuesParamsEntityContract, CustomFieldSetValuesParamsEntityCreditGrant, CustomFieldSetValuesParamsEntityCustomerPlan, CustomFieldSetValuesParamsEntityCustomer, CustomFieldSetValuesParamsEntityInvoice, CustomFieldSetValuesParamsEntityPlan, CustomFieldSetValuesParamsEntityProfessionalService, CustomFieldSetValuesParamsEntityProduct, CustomFieldSetValuesParamsEntityRateCard, CustomFieldSetValuesParamsEntityScheduledCharge:
		return true
	}
	return false
}
