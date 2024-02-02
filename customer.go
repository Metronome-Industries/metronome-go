// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
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

// CustomerService contains methods and other services that help with interacting
// with the metronome API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewCustomerService] method instead.
type CustomerService struct {
	Options       []option.RequestOption
	Plans         *CustomerPlanService
	Invoices      *CustomerInvoiceService
	BillingConfig *CustomerBillingConfigService
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r *CustomerService) {
	r = &CustomerService{}
	r.Options = opts
	r.Plans = NewCustomerPlanService(opts...)
	r.Invoices = NewCustomerInvoiceService(opts...)
	r.BillingConfig = NewCustomerBillingConfigService(opts...)
	return
}

// Create a new customer
func (r *CustomerService) New(ctx context.Context, body CustomerNewParams, opts ...option.RequestOption) (res *CustomerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a customer by Metronome ID.
func (r *CustomerService) Get(ctx context.Context, customerID string, opts ...option.RequestOption) (res *CustomerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all customers.
func (r *CustomerService) List(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) (res *CustomerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Archive a customer
func (r *CustomerService) Archive(ctx context.Context, body CustomerArchiveParams, opts ...option.RequestOption) (res *CustomerArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customers/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all billable metrics.
func (r *CustomerService) ListBillableMetrics(ctx context.Context, customerID string, query CustomerListBillableMetricsParams, opts ...option.RequestOption) (res *CustomerListBillableMetricsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s/billable-metrics", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch daily pending costs for the specified customer, broken down by credit type
// and line items. Note: this is not supported for customers whose plan includes a
// UNIQUE-type billable metric.
func (r *CustomerService) ListCosts(ctx context.Context, customerID string, query CustomerListCostsParams, opts ...option.RequestOption) (res *CustomerListCostsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s/costs", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Sets the ingest aliases for a customer. Ingest aliases can be used in the
// `customer_id` field when sending usage events to Metronome. This call is
// idempotent. It fully replaces the set of ingest aliases for the given customer.
func (r *CustomerService) SetIngestAliases(ctx context.Context, customerID string, body CustomerSetIngestAliasesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("customers/%s/setIngestAliases", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Updates the specified customer's name.
func (r *CustomerService) SetName(ctx context.Context, customerID string, body CustomerSetNameParams, opts ...option.RequestOption) (res *CustomerSetNameResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s/setName", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates the specified customer's config.
func (r *CustomerService) UpdateConfig(ctx context.Context, customerID string, body CustomerUpdateConfigParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("customers/%s/updateConfig", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CustomerNewResponse struct {
	Data CustomerNewResponseData `json:"data,required"`
	JSON customerNewResponseJSON `json:"-"`
}

// customerNewResponseJSON contains the JSON metadata for the struct
// [CustomerNewResponse]
type customerNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerNewResponseData struct {
	// the Metronome ID of the customer
	ID string `json:"id,required" format:"uuid"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string                    `json:"ingest_aliases,required"`
	Name          string                      `json:"name,required"`
	CustomFields  map[string]string           `json:"custom_fields"`
	JSON          customerNewResponseDataJSON `json:"-"`
}

// customerNewResponseDataJSON contains the JSON metadata for the struct
// [CustomerNewResponseData]
type customerNewResponseDataJSON struct {
	ID            apijson.Field
	ExternalID    apijson.Field
	IngestAliases apijson.Field
	Name          apijson.Field
	CustomFields  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomerNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerGetResponse struct {
	Data CustomerGetResponseData `json:"data,required"`
	JSON customerGetResponseJSON `json:"-"`
}

// customerGetResponseJSON contains the JSON metadata for the struct
// [CustomerGetResponse]
type customerGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerGetResponseData struct {
	// the Metronome ID of the customer
	ID                    string                                       `json:"id,required" format:"uuid"`
	CurrentBillableStatus CustomerGetResponseDataCurrentBillableStatus `json:"current_billable_status,required"`
	CustomFields          map[string]string                            `json:"custom_fields,required"`
	CustomerConfig        CustomerGetResponseDataCustomerConfig        `json:"customer_config,required"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string                    `json:"ingest_aliases,required"`
	Name          string                      `json:"name,required"`
	JSON          customerGetResponseDataJSON `json:"-"`
}

// customerGetResponseDataJSON contains the JSON metadata for the struct
// [CustomerGetResponseData]
type customerGetResponseDataJSON struct {
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

func (r *CustomerGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerGetResponseDataCurrentBillableStatus struct {
	Value       CustomerGetResponseDataCurrentBillableStatusValue `json:"value,required"`
	EffectiveAt time.Time                                         `json:"effective_at,nullable" format:"date-time"`
	JSON        customerGetResponseDataCurrentBillableStatusJSON  `json:"-"`
}

// customerGetResponseDataCurrentBillableStatusJSON contains the JSON metadata for
// the struct [CustomerGetResponseDataCurrentBillableStatus]
type customerGetResponseDataCurrentBillableStatusJSON struct {
	Value       apijson.Field
	EffectiveAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerGetResponseDataCurrentBillableStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerGetResponseDataCurrentBillableStatusValue string

const (
	CustomerGetResponseDataCurrentBillableStatusValueBillable   CustomerGetResponseDataCurrentBillableStatusValue = "billable"
	CustomerGetResponseDataCurrentBillableStatusValueUnbillable CustomerGetResponseDataCurrentBillableStatusValue = "unbillable"
)

type CustomerGetResponseDataCustomerConfig struct {
	// The Salesforce account ID for the customer
	SalesforceAccountID string                                    `json:"salesforce_account_id,required,nullable"`
	JSON                customerGetResponseDataCustomerConfigJSON `json:"-"`
}

// customerGetResponseDataCustomerConfigJSON contains the JSON metadata for the
// struct [CustomerGetResponseDataCustomerConfig]
type customerGetResponseDataCustomerConfigJSON struct {
	SalesforceAccountID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerGetResponseDataCustomerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListResponse struct {
	Data     []CustomerListResponseData `json:"data,required"`
	NextPage string                     `json:"next_page,required,nullable"`
	JSON     customerListResponseJSON   `json:"-"`
}

// customerListResponseJSON contains the JSON metadata for the struct
// [CustomerListResponse]
type customerListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListResponseData struct {
	// the Metronome ID of the customer
	ID                    string                                        `json:"id,required" format:"uuid"`
	CurrentBillableStatus CustomerListResponseDataCurrentBillableStatus `json:"current_billable_status,required"`
	CustomFields          map[string]string                             `json:"custom_fields,required"`
	CustomerConfig        CustomerListResponseDataCustomerConfig        `json:"customer_config,required"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string                     `json:"ingest_aliases,required"`
	Name          string                       `json:"name,required"`
	JSON          customerListResponseDataJSON `json:"-"`
}

// customerListResponseDataJSON contains the JSON metadata for the struct
// [CustomerListResponseData]
type customerListResponseDataJSON struct {
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

func (r *CustomerListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListResponseDataCurrentBillableStatus struct {
	Value       CustomerListResponseDataCurrentBillableStatusValue `json:"value,required"`
	EffectiveAt time.Time                                          `json:"effective_at,nullable" format:"date-time"`
	JSON        customerListResponseDataCurrentBillableStatusJSON  `json:"-"`
}

// customerListResponseDataCurrentBillableStatusJSON contains the JSON metadata for
// the struct [CustomerListResponseDataCurrentBillableStatus]
type customerListResponseDataCurrentBillableStatusJSON struct {
	Value       apijson.Field
	EffectiveAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerListResponseDataCurrentBillableStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListResponseDataCurrentBillableStatusValue string

const (
	CustomerListResponseDataCurrentBillableStatusValueBillable   CustomerListResponseDataCurrentBillableStatusValue = "billable"
	CustomerListResponseDataCurrentBillableStatusValueUnbillable CustomerListResponseDataCurrentBillableStatusValue = "unbillable"
)

type CustomerListResponseDataCustomerConfig struct {
	// The Salesforce account ID for the customer
	SalesforceAccountID string                                     `json:"salesforce_account_id,required,nullable"`
	JSON                customerListResponseDataCustomerConfigJSON `json:"-"`
}

// customerListResponseDataCustomerConfigJSON contains the JSON metadata for the
// struct [CustomerListResponseDataCustomerConfig]
type customerListResponseDataCustomerConfigJSON struct {
	SalesforceAccountID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerListResponseDataCustomerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerArchiveResponse struct {
	Data CustomerArchiveResponseData `json:"data,required"`
	JSON customerArchiveResponseJSON `json:"-"`
}

// customerArchiveResponseJSON contains the JSON metadata for the struct
// [CustomerArchiveResponse]
type customerArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerArchiveResponseData struct {
	ID   string                          `json:"id,required" format:"uuid"`
	JSON customerArchiveResponseDataJSON `json:"-"`
}

// customerArchiveResponseDataJSON contains the JSON metadata for the struct
// [CustomerArchiveResponseData]
type customerArchiveResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerArchiveResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListBillableMetricsResponse struct {
	Data     []CustomerListBillableMetricsResponseData `json:"data,required"`
	NextPage string                                    `json:"next_page,required,nullable"`
	JSON     customerListBillableMetricsResponseJSON   `json:"-"`
}

// customerListBillableMetricsResponseJSON contains the JSON metadata for the
// struct [CustomerListBillableMetricsResponse]
type customerListBillableMetricsResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerListBillableMetricsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListBillableMetricsResponseData struct {
	ID      string                                      `json:"id,required" format:"uuid"`
	Name    string                                      `json:"name,required"`
	GroupBy []string                                    `json:"group_by"`
	JSON    customerListBillableMetricsResponseDataJSON `json:"-"`
}

// customerListBillableMetricsResponseDataJSON contains the JSON metadata for the
// struct [CustomerListBillableMetricsResponseData]
type customerListBillableMetricsResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	GroupBy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerListBillableMetricsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListCostsResponse struct {
	Data     []CustomerListCostsResponseData `json:"data,required"`
	NextPage string                          `json:"next_page,required,nullable"`
	JSON     customerListCostsResponseJSON   `json:"-"`
}

// customerListCostsResponseJSON contains the JSON metadata for the struct
// [CustomerListCostsResponse]
type customerListCostsResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerListCostsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListCostsResponseData struct {
	CreditTypes    map[string]CustomerListCostsResponseDataCreditType `json:"credit_types,required"`
	EndTimestamp   time.Time                                          `json:"end_timestamp,required" format:"date-time"`
	StartTimestamp time.Time                                          `json:"start_timestamp,required" format:"date-time"`
	JSON           customerListCostsResponseDataJSON                  `json:"-"`
}

// customerListCostsResponseDataJSON contains the JSON metadata for the struct
// [CustomerListCostsResponseData]
type customerListCostsResponseDataJSON struct {
	CreditTypes    apijson.Field
	EndTimestamp   apijson.Field
	StartTimestamp apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerListCostsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListCostsResponseDataCreditType struct {
	Cost              float64                                                     `json:"cost"`
	LineItemBreakdown []CustomerListCostsResponseDataCreditTypesLineItemBreakdown `json:"line_item_breakdown"`
	Name              string                                                      `json:"name"`
	JSON              customerListCostsResponseDataCreditTypeJSON                 `json:"-"`
}

// customerListCostsResponseDataCreditTypeJSON contains the JSON metadata for the
// struct [CustomerListCostsResponseDataCreditType]
type customerListCostsResponseDataCreditTypeJSON struct {
	Cost              apijson.Field
	LineItemBreakdown apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerListCostsResponseDataCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListCostsResponseDataCreditTypesLineItemBreakdown struct {
	Cost       float64                                                       `json:"cost,required"`
	Name       string                                                        `json:"name,required"`
	GroupKey   string                                                        `json:"group_key"`
	GroupValue string                                                        `json:"group_value,nullable"`
	JSON       customerListCostsResponseDataCreditTypesLineItemBreakdownJSON `json:"-"`
}

// customerListCostsResponseDataCreditTypesLineItemBreakdownJSON contains the JSON
// metadata for the struct
// [CustomerListCostsResponseDataCreditTypesLineItemBreakdown]
type customerListCostsResponseDataCreditTypesLineItemBreakdownJSON struct {
	Cost        apijson.Field
	Name        apijson.Field
	GroupKey    apijson.Field
	GroupValue  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerListCostsResponseDataCreditTypesLineItemBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerSetNameResponse struct {
	Data CustomerSetNameResponseData `json:"data,required"`
	JSON customerSetNameResponseJSON `json:"-"`
}

// customerSetNameResponseJSON contains the JSON metadata for the struct
// [CustomerSetNameResponse]
type customerSetNameResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerSetNameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerSetNameResponseData struct {
	// the Metronome ID of the customer
	ID string `json:"id,required" format:"uuid"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string                        `json:"ingest_aliases,required"`
	Name          string                          `json:"name,required"`
	CustomFields  map[string]string               `json:"custom_fields"`
	JSON          customerSetNameResponseDataJSON `json:"-"`
}

// customerSetNameResponseDataJSON contains the JSON metadata for the struct
// [CustomerSetNameResponseData]
type customerSetNameResponseDataJSON struct {
	ID            apijson.Field
	ExternalID    apijson.Field
	IngestAliases apijson.Field
	Name          apijson.Field
	CustomFields  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomerSetNameResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerNewParams struct {
	Name          param.Field[string]                         `json:"name,required"`
	BillingConfig param.Field[CustomerNewParamsBillingConfig] `json:"billing_config"`
	CustomFields  param.Field[map[string]string]              `json:"custom_fields"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome ID or ingest
	// alias) that can be used in usage events
	ExternalID param.Field[string] `json:"external_id"`
	// Aliases that can be used to refer to this customer in usage events
	IngestAliases param.Field[[]string] `json:"ingest_aliases"`
}

func (r CustomerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParamsBillingConfig struct {
	BillingProviderCustomerID param.Field[string]                                               `json:"billing_provider_customer_id,required"`
	BillingProviderType       param.Field[CustomerNewParamsBillingConfigBillingProviderType]    `json:"billing_provider_type,required"`
	AwsProductCode            param.Field[string]                                               `json:"aws_product_code"`
	AwsRegion                 param.Field[CustomerNewParamsBillingConfigAwsRegion]              `json:"aws_region"`
	StripeCollectionMethod    param.Field[CustomerNewParamsBillingConfigStripeCollectionMethod] `json:"stripe_collection_method"`
}

func (r CustomerNewParamsBillingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParamsBillingConfigBillingProviderType string

const (
	CustomerNewParamsBillingConfigBillingProviderTypeAwsMarketplace   CustomerNewParamsBillingConfigBillingProviderType = "aws_marketplace"
	CustomerNewParamsBillingConfigBillingProviderTypeStripe           CustomerNewParamsBillingConfigBillingProviderType = "stripe"
	CustomerNewParamsBillingConfigBillingProviderTypeNetsuite         CustomerNewParamsBillingConfigBillingProviderType = "netsuite"
	CustomerNewParamsBillingConfigBillingProviderTypeCustom           CustomerNewParamsBillingConfigBillingProviderType = "custom"
	CustomerNewParamsBillingConfigBillingProviderTypeAzureMarketplace CustomerNewParamsBillingConfigBillingProviderType = "azure_marketplace"
	CustomerNewParamsBillingConfigBillingProviderTypeQuickbooksOnline CustomerNewParamsBillingConfigBillingProviderType = "quickbooks_online"
)

type CustomerNewParamsBillingConfigAwsRegion string

const (
	CustomerNewParamsBillingConfigAwsRegionAfSouth1     CustomerNewParamsBillingConfigAwsRegion = "af-south-1"
	CustomerNewParamsBillingConfigAwsRegionApEast1      CustomerNewParamsBillingConfigAwsRegion = "ap-east-1"
	CustomerNewParamsBillingConfigAwsRegionApNortheast1 CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-1"
	CustomerNewParamsBillingConfigAwsRegionApNortheast2 CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-2"
	CustomerNewParamsBillingConfigAwsRegionApNortheast3 CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-3"
	CustomerNewParamsBillingConfigAwsRegionApSouth1     CustomerNewParamsBillingConfigAwsRegion = "ap-south-1"
	CustomerNewParamsBillingConfigAwsRegionApSoutheast1 CustomerNewParamsBillingConfigAwsRegion = "ap-southeast-1"
	CustomerNewParamsBillingConfigAwsRegionApSoutheast2 CustomerNewParamsBillingConfigAwsRegion = "ap-southeast-2"
	CustomerNewParamsBillingConfigAwsRegionCaCentral1   CustomerNewParamsBillingConfigAwsRegion = "ca-central-1"
	CustomerNewParamsBillingConfigAwsRegionCnNorth1     CustomerNewParamsBillingConfigAwsRegion = "cn-north-1"
	CustomerNewParamsBillingConfigAwsRegionCnNorthwest1 CustomerNewParamsBillingConfigAwsRegion = "cn-northwest-1"
	CustomerNewParamsBillingConfigAwsRegionEuCentral1   CustomerNewParamsBillingConfigAwsRegion = "eu-central-1"
	CustomerNewParamsBillingConfigAwsRegionEuNorth1     CustomerNewParamsBillingConfigAwsRegion = "eu-north-1"
	CustomerNewParamsBillingConfigAwsRegionEuSouth1     CustomerNewParamsBillingConfigAwsRegion = "eu-south-1"
	CustomerNewParamsBillingConfigAwsRegionEuWest1      CustomerNewParamsBillingConfigAwsRegion = "eu-west-1"
	CustomerNewParamsBillingConfigAwsRegionEuWest2      CustomerNewParamsBillingConfigAwsRegion = "eu-west-2"
	CustomerNewParamsBillingConfigAwsRegionEuWest3      CustomerNewParamsBillingConfigAwsRegion = "eu-west-3"
	CustomerNewParamsBillingConfigAwsRegionMeSouth1     CustomerNewParamsBillingConfigAwsRegion = "me-south-1"
	CustomerNewParamsBillingConfigAwsRegionSaEast1      CustomerNewParamsBillingConfigAwsRegion = "sa-east-1"
	CustomerNewParamsBillingConfigAwsRegionUsEast1      CustomerNewParamsBillingConfigAwsRegion = "us-east-1"
	CustomerNewParamsBillingConfigAwsRegionUsEast2      CustomerNewParamsBillingConfigAwsRegion = "us-east-2"
	CustomerNewParamsBillingConfigAwsRegionUsGovEast1   CustomerNewParamsBillingConfigAwsRegion = "us-gov-east-1"
	CustomerNewParamsBillingConfigAwsRegionUsGovWest1   CustomerNewParamsBillingConfigAwsRegion = "us-gov-west-1"
	CustomerNewParamsBillingConfigAwsRegionUsWest1      CustomerNewParamsBillingConfigAwsRegion = "us-west-1"
	CustomerNewParamsBillingConfigAwsRegionUsWest2      CustomerNewParamsBillingConfigAwsRegion = "us-west-2"
)

type CustomerNewParamsBillingConfigStripeCollectionMethod string

const (
	CustomerNewParamsBillingConfigStripeCollectionMethodChargeAutomatically CustomerNewParamsBillingConfigStripeCollectionMethod = "charge_automatically"
	CustomerNewParamsBillingConfigStripeCollectionMethodSendInvoice         CustomerNewParamsBillingConfigStripeCollectionMethod = "send_invoice"
)

type CustomerListParams struct {
	// Filter the customer list by customer_id. Up to 100 ids can be provided.
	CustomerIDs param.Field[[]string] `query:"customer_ids"`
	// Filter the customer list by ingest_alias
	IngestAlias param.Field[string] `query:"ingest_alias"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Filter the customer list by only archived customers.
	OnlyArchived param.Field[bool] `query:"only_archived"`
	// Filter the customer list by salesforce_account_id. Up to 100 ids can be
	// provided.
	SalesforceAccountIDs param.Field[[]string] `query:"salesforce_account_ids"`
}

// URLQuery serializes [CustomerListParams]'s query parameters as `url.Values`.
func (r CustomerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerArchiveParams struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r CustomerArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerListBillableMetricsParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// If true, the list of metrics will be filtered to just ones that are on the
	// customer's current plan
	OnCurrentPlan param.Field[bool] `query:"on_current_plan"`
}

// URLQuery serializes [CustomerListBillableMetricsParams]'s query parameters as
// `url.Values`.
func (r CustomerListBillableMetricsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerListCostsParams struct {
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `query:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingOn param.Field[time.Time] `query:"starting_on,required" format:"date-time"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [CustomerListCostsParams]'s query parameters as
// `url.Values`.
func (r CustomerListCostsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerSetIngestAliasesParams struct {
	IngestAliases param.Field[[]string] `json:"ingest_aliases,required"`
}

func (r CustomerSetIngestAliasesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerSetNameParams struct {
	// The new name for the customer
	Name param.Field[string] `json:"name,required"`
}

func (r CustomerSetNameParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateConfigParams struct {
	// Leave in draft or set to auto-advance on invoices sent to Stripe. Falls back to
	// the client-level config if unset, which defaults to true if unset.
	LeaveStripeInvoicesInDraft param.Field[bool] `json:"leave_stripe_invoices_in_draft"`
	// The Salesforce account ID for the customer
	SalesforceAccountID param.Field[string] `json:"salesforce_account_id"`
}

func (r CustomerUpdateConfigParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
