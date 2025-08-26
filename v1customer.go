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

// V1CustomerService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerService] method instead.
type V1CustomerService struct {
	Options        []option.RequestOption
	Alerts         *V1CustomerAlertService
	Plans          *V1CustomerPlanService
	Invoices       *V1CustomerInvoiceService
	BillingConfig  *V1CustomerBillingConfigService
	Commits        *V1CustomerCommitService
	Credits        *V1CustomerCreditService
	NamedSchedules *V1CustomerNamedScheduleService
}

// NewV1CustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomerService(opts ...option.RequestOption) (r *V1CustomerService) {
	r = &V1CustomerService{}
	r.Options = opts
	r.Alerts = NewV1CustomerAlertService(opts...)
	r.Plans = NewV1CustomerPlanService(opts...)
	r.Invoices = NewV1CustomerInvoiceService(opts...)
	r.BillingConfig = NewV1CustomerBillingConfigService(opts...)
	r.Commits = NewV1CustomerCommitService(opts...)
	r.Credits = NewV1CustomerCreditService(opts...)
	r.NamedSchedules = NewV1CustomerNamedScheduleService(opts...)
	return
}

// Create a new customer in Metronome and optionally the billing configuration
// (recommended) which dictates where invoices for the customer will be sent or
// where payment will be collected.
//
// Use this endpoint to:\
// Execute your customer provisioning workflows for either PLG motions, where customers
// originate in your platform, or SLG motions, where customers originate in your sales
// system.
//
//   - Key response fields: This end-point returns the customer_id created by the
//     request. This id can be used to fetch relevant billing configurations and
//     create contracts.
//
// Example workflow:
//
//   - Generally, Metronome recommends first creating the customer in the downstream
//     payment / ERP system when payment method is collected and then creating the
//     customer in Metronome using the response (i.e. customer_id) from the
//     downstream system. If you do not create a billing configuration on customer
//     creation, you can add it later.
//   - Once a customer is created, you can then create a contract for the customer.
//     In the contract creation process, you will need to add the customer billing
//     configuration to the contract to ensure Metronome invoices the customer
//     correctly. This is because a customer can have multiple configurations.
//   - As part of the customer creation process, set the ingest alias for the
//     customer which will ensure usage is accurately mapped to the customer. Ingest
//     aliases can be added or changed after the creation process as well.
//
// Usage guidelines:\
// For details on different billing configurations for different systems, review the
// /setCustomerBillingConfiguration end-point.
func (r *V1CustomerService) New(ctx context.Context, body V1CustomerNewParams, opts ...option.RequestOption) (res *V1CustomerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get detailed information for a specific customer by their Metronome ID. Returns
// customer profile data including name, creation date, ingest aliases,
// configuration settings, and custom fields. Use this endpoint to fetch complete
// customer details for billing operations or account management.
//
// Note: If searching for a customer billing configuration, use the
// /getCustomerBillingConfigurations end-point.
func (r *V1CustomerService) Get(ctx context.Context, query V1CustomerGetParams, opts ...option.RequestOption) (res *V1CustomerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s", query.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets a paginated list of all customers in your Metronome account. Use this
// endpoint to browse your customer base, implement customer search functionality,
// or sync customer data with external systems. Returns customer details including
// IDs, names, and configuration settings. Supports filtering and pagination
// parameters for efficient data retrieval.
func (r *V1CustomerService) List(ctx context.Context, query V1CustomerListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/customers"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Gets a paginated list of all customers in your Metronome account. Use this
// endpoint to browse your customer base, implement customer search functionality,
// or sync customer data with external systems. Returns customer details including
// IDs, names, and configuration settings. Supports filtering and pagination
// parameters for efficient data retrieval.
func (r *V1CustomerService) ListAutoPaging(ctx context.Context, query V1CustomerListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CustomerListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Use this endpoint to archive a customer while preserving auditability. Archiving
// a customer will automatically archive all contracts as of the current date and
// void all corresponding invoices. Use this endpoint if a customer is onboarded by
// mistake.
//
// Usage guidelines:
//
//   - Once a customer is archived, it cannot be unarchived.
//   - Archived customers can still be viewed through the API or the UI for audit
//     purposes.
//   - Ingest aliases remain idempotent for archived customers. In order to reuse an
//     ingest alias, first remove the ingest alias from the customer prior to
//     archiving.
//   - Any alerts associated with the customer will no longer be triggered.
func (r *V1CustomerService) Archive(ctx context.Context, body V1CustomerArchiveParams, opts ...option.RequestOption) (res *V1CustomerArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/customers/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get all billable metrics available for a specific customer. Supports pagination
// and filtering by current plan status or archived metrics. Use this endpoint to
// see which metrics are being tracked for billing calculations for a given
// customer.
func (r *V1CustomerService) ListBillableMetrics(ctx context.Context, params V1CustomerListBillableMetricsParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerListBillableMetricsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/billable-metrics", params.CustomerID)
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

// Get all billable metrics available for a specific customer. Supports pagination
// and filtering by current plan status or archived metrics. Use this endpoint to
// see which metrics are being tracked for billing calculations for a given
// customer.
func (r *V1CustomerService) ListBillableMetricsAutoPaging(ctx context.Context, params V1CustomerListBillableMetricsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CustomerListBillableMetricsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListBillableMetrics(ctx, params, opts...))
}

// Fetch daily pending costs for the specified customer, broken down by credit type
// and line items. Note: this is not supported for customers whose plan includes a
// UNIQUE-type billable metric.
func (r *V1CustomerService) ListCosts(ctx context.Context, params V1CustomerListCostsParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerListCostsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/costs", params.CustomerID)
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

// Fetch daily pending costs for the specified customer, broken down by credit type
// and line items. Note: this is not supported for customers whose plan includes a
// UNIQUE-type billable metric.
func (r *V1CustomerService) ListCostsAutoPaging(ctx context.Context, params V1CustomerListCostsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CustomerListCostsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListCosts(ctx, params, opts...))
}

// Preview how a set of events will affect a customer's invoice. Generates a draft
// invoice for a customer using their current contract configuration and the
// provided events. This is useful for testing how new events will affect the
// customer's invoice before they are actually processed.
func (r *V1CustomerService) PreviewEvents(ctx context.Context, params V1CustomerPreviewEventsParams, opts ...option.RequestOption) (res *V1CustomerPreviewEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/previewEvents", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns all billing configurations previously set for the customer. Use during
// the contract provisioning process to fetch the
// `billing_provider_configuration_id` needed to set the contract billing
// configuration.
func (r *V1CustomerService) GetBillingConfigurations(ctx context.Context, body V1CustomerGetBillingConfigurationsParams, opts ...option.RequestOption) (res *V1CustomerGetBillingConfigurationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/getCustomerBillingProviderConfigurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a billing configuration for a customer. Once created, these
// configurations are available to associate to a contract and dictates which
// downstream system to collect payment in or send the invoice to. You can create
// multiple configurations per customer. The configuration formats are distinct for
// each downstream provider.
//
// Use this endpoint to:
//
//   - Add the initial configuration to an existing customer. Once created, the
//     billing configuration can then be associated to the customer's contract.
//   - Add a new configuration to an existing customer. This might be used as part of
//     an upgrade or downgrade workflow where the customer was previously billed
//     through system A (e.g. Stripe) but will now be billed through system B (e.g.
//     AWS). Once created, the new configuration can then be associated to the
//     customer's contract.
//
// Delivery Method Options:
//
//   - direct_to_billing_provider: Use when Metronome should send invoices directly
//     to the billing provider's API (e.g., Stripe, NetSuite). This is the most
//     common method for automated billing workflows.
//   - tackle: Use specifically for AWS Marketplace transactions that require
//     Tackle's co-selling platform for partner attribution and commission tracking.
//   - aws_sqs: Use when you want invoice data delivered to an AWS SQS queue for
//     custom processing before sending to your billing system.
//   - aws_sns: Use when you want invoice notifications published to an AWS SNS topic
//     for event-driven billing workflows.
//
// Key response fields: The id for the customer billing configuration. This id can
// be used to associate the billing configuration to a contract.
//
// Usage guidelines:\
// Must use the delivery_method_id if you have multiple Stripe accounts connected to
// Metronome.
func (r *V1CustomerService) SetBillingConfigurations(ctx context.Context, body V1CustomerSetBillingConfigurationsParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/setCustomerBillingProviderConfigurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Sets the ingest aliases for a customer. Use this endpoint to associate a
// Metronome customer with an internal ID for easier tracking between systems.
// Ingest aliases can be used in the customer_id field when sending usage events to
// Metronome.
//
// Usage guidelines:
//
//   - This call is idempotent and fully replaces the set of ingest aliases for the
//     given customer.
//   - Switching an ingest alias from one customer to another will associate all
//     corresponding usage to the new customer.
//   - Use multiple ingest aliases to model child organizations within a single
//     Metronome customer.
func (r *V1CustomerService) SetIngestAliases(ctx context.Context, params V1CustomerSetIngestAliasesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/setIngestAliases", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Updates the display name for a customer record. Use this to correct customer
// names, update business names after rebranding, or maintain accurate customer
// information for invoicing and reporting. Returns the updated customer object
// with the new name applied immediately across all billing documents and
// interfaces.
func (r *V1CustomerService) SetName(ctx context.Context, params V1CustomerSetNameParams, opts ...option.RequestOption) (res *V1CustomerSetNameResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/setName", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update configuration settings for a specific customer, such as external system
// integrations (e.g., Salesforce account ID) and other customer-specific billing
// parameters. Use this endpoint to modify customer configurations without
// affecting core customer data like name or ingest aliases.
func (r *V1CustomerService) UpdateConfig(ctx context.Context, params V1CustomerUpdateConfigParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/updateConfig", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type V1CustomerNewResponse struct {
	Data V1CustomerNewResponseData `json:"data,required"`
	JSON v1CustomerNewResponseJSON `json:"-"`
}

// v1CustomerNewResponseJSON contains the JSON metadata for the struct
// [V1CustomerNewResponse]
type v1CustomerNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerNewResponseData struct {
	// the Metronome ID of the customer
	ID string `json:"id,required" format:"uuid"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string `json:"ingest_aliases,required"`
	Name          string   `json:"name,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string             `json:"custom_fields"`
	JSON         v1CustomerNewResponseDataJSON `json:"-"`
}

// v1CustomerNewResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerNewResponseData]
type v1CustomerNewResponseDataJSON struct {
	ID            apijson.Field
	ExternalID    apijson.Field
	IngestAliases apijson.Field
	Name          apijson.Field
	CustomFields  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1CustomerNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerGetResponse struct {
	Data V1CustomerGetResponseData `json:"data,required"`
	JSON v1CustomerGetResponseJSON `json:"-"`
}

// v1CustomerGetResponseJSON contains the JSON metadata for the struct
// [V1CustomerGetResponse]
type v1CustomerGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerGetResponseData struct {
	// the Metronome ID of the customer
	ID string `json:"id,required" format:"uuid"`
	// RFC 3339 timestamp indicating when the customer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields   map[string]string                       `json:"custom_fields,required"`
	CustomerConfig V1CustomerGetResponseDataCustomerConfig `json:"customer_config,required"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string `json:"ingest_aliases,required"`
	Name          string   `json:"name,required"`
	// RFC 3339 timestamp indicating when the customer was archived. Null if the
	// customer is active.
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// This field's availability is dependent on your client's configuration.
	CurrentBillableStatus V1CustomerGetResponseDataCurrentBillableStatus `json:"current_billable_status"`
	JSON                  v1CustomerGetResponseDataJSON                  `json:"-"`
}

// v1CustomerGetResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerGetResponseData]
type v1CustomerGetResponseDataJSON struct {
	ID                    apijson.Field
	CreatedAt             apijson.Field
	CustomFields          apijson.Field
	CustomerConfig        apijson.Field
	ExternalID            apijson.Field
	IngestAliases         apijson.Field
	Name                  apijson.Field
	ArchivedAt            apijson.Field
	CurrentBillableStatus apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V1CustomerGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerGetResponseDataCustomerConfig struct {
	// The Salesforce account ID for the customer
	SalesforceAccountID string                                      `json:"salesforce_account_id,required,nullable"`
	JSON                v1CustomerGetResponseDataCustomerConfigJSON `json:"-"`
}

// v1CustomerGetResponseDataCustomerConfigJSON contains the JSON metadata for the
// struct [V1CustomerGetResponseDataCustomerConfig]
type v1CustomerGetResponseDataCustomerConfigJSON struct {
	SalesforceAccountID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerGetResponseDataCustomerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerGetResponseDataCustomerConfigJSON) RawJSON() string {
	return r.raw
}

// This field's availability is dependent on your client's configuration.
type V1CustomerGetResponseDataCurrentBillableStatus struct {
	Value       V1CustomerGetResponseDataCurrentBillableStatusValue `json:"value,required"`
	EffectiveAt time.Time                                           `json:"effective_at,nullable" format:"date-time"`
	JSON        v1CustomerGetResponseDataCurrentBillableStatusJSON  `json:"-"`
}

// v1CustomerGetResponseDataCurrentBillableStatusJSON contains the JSON metadata
// for the struct [V1CustomerGetResponseDataCurrentBillableStatus]
type v1CustomerGetResponseDataCurrentBillableStatusJSON struct {
	Value       apijson.Field
	EffectiveAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerGetResponseDataCurrentBillableStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerGetResponseDataCurrentBillableStatusJSON) RawJSON() string {
	return r.raw
}

type V1CustomerGetResponseDataCurrentBillableStatusValue string

const (
	V1CustomerGetResponseDataCurrentBillableStatusValueBillable   V1CustomerGetResponseDataCurrentBillableStatusValue = "billable"
	V1CustomerGetResponseDataCurrentBillableStatusValueUnbillable V1CustomerGetResponseDataCurrentBillableStatusValue = "unbillable"
)

func (r V1CustomerGetResponseDataCurrentBillableStatusValue) IsKnown() bool {
	switch r {
	case V1CustomerGetResponseDataCurrentBillableStatusValueBillable, V1CustomerGetResponseDataCurrentBillableStatusValueUnbillable:
		return true
	}
	return false
}

type V1CustomerListResponse struct {
	// the Metronome ID of the customer
	ID string `json:"id,required" format:"uuid"`
	// RFC 3339 timestamp indicating when the customer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields   map[string]string                    `json:"custom_fields,required"`
	CustomerConfig V1CustomerListResponseCustomerConfig `json:"customer_config,required"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string `json:"ingest_aliases,required"`
	Name          string   `json:"name,required"`
	// RFC 3339 timestamp indicating when the customer was archived. Null if the
	// customer is active.
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// This field's availability is dependent on your client's configuration.
	CurrentBillableStatus V1CustomerListResponseCurrentBillableStatus `json:"current_billable_status"`
	JSON                  v1CustomerListResponseJSON                  `json:"-"`
}

// v1CustomerListResponseJSON contains the JSON metadata for the struct
// [V1CustomerListResponse]
type v1CustomerListResponseJSON struct {
	ID                    apijson.Field
	CreatedAt             apijson.Field
	CustomFields          apijson.Field
	CustomerConfig        apijson.Field
	ExternalID            apijson.Field
	IngestAliases         apijson.Field
	Name                  apijson.Field
	ArchivedAt            apijson.Field
	CurrentBillableStatus apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V1CustomerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerListResponseCustomerConfig struct {
	// The Salesforce account ID for the customer
	SalesforceAccountID string                                   `json:"salesforce_account_id,required,nullable"`
	JSON                v1CustomerListResponseCustomerConfigJSON `json:"-"`
}

// v1CustomerListResponseCustomerConfigJSON contains the JSON metadata for the
// struct [V1CustomerListResponseCustomerConfig]
type v1CustomerListResponseCustomerConfigJSON struct {
	SalesforceAccountID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerListResponseCustomerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListResponseCustomerConfigJSON) RawJSON() string {
	return r.raw
}

// This field's availability is dependent on your client's configuration.
type V1CustomerListResponseCurrentBillableStatus struct {
	Value       V1CustomerListResponseCurrentBillableStatusValue `json:"value,required"`
	EffectiveAt time.Time                                        `json:"effective_at,nullable" format:"date-time"`
	JSON        v1CustomerListResponseCurrentBillableStatusJSON  `json:"-"`
}

// v1CustomerListResponseCurrentBillableStatusJSON contains the JSON metadata for
// the struct [V1CustomerListResponseCurrentBillableStatus]
type v1CustomerListResponseCurrentBillableStatusJSON struct {
	Value       apijson.Field
	EffectiveAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerListResponseCurrentBillableStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListResponseCurrentBillableStatusJSON) RawJSON() string {
	return r.raw
}

type V1CustomerListResponseCurrentBillableStatusValue string

const (
	V1CustomerListResponseCurrentBillableStatusValueBillable   V1CustomerListResponseCurrentBillableStatusValue = "billable"
	V1CustomerListResponseCurrentBillableStatusValueUnbillable V1CustomerListResponseCurrentBillableStatusValue = "unbillable"
)

func (r V1CustomerListResponseCurrentBillableStatusValue) IsKnown() bool {
	switch r {
	case V1CustomerListResponseCurrentBillableStatusValueBillable, V1CustomerListResponseCurrentBillableStatusValueUnbillable:
		return true
	}
	return false
}

type V1CustomerArchiveResponse struct {
	Data V1CustomerArchiveResponseData `json:"data,required"`
	JSON v1CustomerArchiveResponseJSON `json:"-"`
}

// v1CustomerArchiveResponseJSON contains the JSON metadata for the struct
// [V1CustomerArchiveResponse]
type v1CustomerArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerArchiveResponseData struct {
	ID   string                            `json:"id,required" format:"uuid"`
	JSON v1CustomerArchiveResponseDataJSON `json:"-"`
}

// v1CustomerArchiveResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerArchiveResponseData]
type v1CustomerArchiveResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerArchiveResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerArchiveResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerListBillableMetricsResponse struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// (DEPRECATED) use aggregation_type instead
	Aggregate string `json:"aggregate"`
	// (DEPRECATED) use aggregation_key instead
	AggregateKeys []string `json:"aggregate_keys"`
	// A key that specifies which property of the event is used to aggregate data. This
	// key must be one of the property filter names and is not applicable when the
	// aggregation type is 'count'.
	AggregationKey string `json:"aggregation_key"`
	// Specifies the type of aggregation performed on matching events.
	AggregationType V1CustomerListBillableMetricsResponseAggregationType `json:"aggregation_type"`
	// RFC 3339 timestamp indicating when the billable metric was archived. If not
	// provided, the billable metric is not archived.
	ArchivedAt time.Time `json:"archived_at" format:"date-time"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	// An optional filtering rule to match the 'event_type' property of an event.
	EventTypeFilter V1CustomerListBillableMetricsResponseEventTypeFilter `json:"event_type_filter"`
	// (DEPRECATED) use property_filters & event_type_filter instead
	Filter map[string]interface{} `json:"filter"`
	// (DEPRECATED) use group_keys instead
	GroupBy []string `json:"group_by"`
	// Property names that are used to group usage costs on an invoice. Each entry
	// represents a set of properties used to slice events into distinct buckets.
	GroupKeys [][]string `json:"group_keys"`
	// A list of filters to match events to this billable metric. Each filter defines a
	// rule on an event property. All rules must pass for the event to match the
	// billable metric.
	PropertyFilters []V1CustomerListBillableMetricsResponsePropertyFilter `json:"property_filters"`
	// The SQL query associated with the billable metric
	Sql  string                                    `json:"sql"`
	JSON v1CustomerListBillableMetricsResponseJSON `json:"-"`
}

// v1CustomerListBillableMetricsResponseJSON contains the JSON metadata for the
// struct [V1CustomerListBillableMetricsResponse]
type v1CustomerListBillableMetricsResponseJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	Aggregate       apijson.Field
	AggregateKeys   apijson.Field
	AggregationKey  apijson.Field
	AggregationType apijson.Field
	ArchivedAt      apijson.Field
	CustomFields    apijson.Field
	EventTypeFilter apijson.Field
	Filter          apijson.Field
	GroupBy         apijson.Field
	GroupKeys       apijson.Field
	PropertyFilters apijson.Field
	Sql             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V1CustomerListBillableMetricsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListBillableMetricsResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of aggregation performed on matching events.
type V1CustomerListBillableMetricsResponseAggregationType string

const (
	V1CustomerListBillableMetricsResponseAggregationTypeCount  V1CustomerListBillableMetricsResponseAggregationType = "COUNT"
	V1CustomerListBillableMetricsResponseAggregationTypeLatest V1CustomerListBillableMetricsResponseAggregationType = "LATEST"
	V1CustomerListBillableMetricsResponseAggregationTypeMax    V1CustomerListBillableMetricsResponseAggregationType = "MAX"
	V1CustomerListBillableMetricsResponseAggregationTypeSum    V1CustomerListBillableMetricsResponseAggregationType = "SUM"
	V1CustomerListBillableMetricsResponseAggregationTypeUnique V1CustomerListBillableMetricsResponseAggregationType = "UNIQUE"
)

func (r V1CustomerListBillableMetricsResponseAggregationType) IsKnown() bool {
	switch r {
	case V1CustomerListBillableMetricsResponseAggregationTypeCount, V1CustomerListBillableMetricsResponseAggregationTypeLatest, V1CustomerListBillableMetricsResponseAggregationTypeMax, V1CustomerListBillableMetricsResponseAggregationTypeSum, V1CustomerListBillableMetricsResponseAggregationTypeUnique:
		return true
	}
	return false
}

// An optional filtering rule to match the 'event_type' property of an event.
type V1CustomerListBillableMetricsResponseEventTypeFilter struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues []string `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues []string                                                 `json:"not_in_values"`
	JSON        v1CustomerListBillableMetricsResponseEventTypeFilterJSON `json:"-"`
}

// v1CustomerListBillableMetricsResponseEventTypeFilterJSON contains the JSON
// metadata for the struct [V1CustomerListBillableMetricsResponseEventTypeFilter]
type v1CustomerListBillableMetricsResponseEventTypeFilterJSON struct {
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerListBillableMetricsResponseEventTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListBillableMetricsResponseEventTypeFilterJSON) RawJSON() string {
	return r.raw
}

type V1CustomerListBillableMetricsResponsePropertyFilter struct {
	// The name of the event property.
	Name string `json:"name,required"`
	// Determines whether the property must exist in the event. If true, only events
	// with this property will pass the filter. If false, only events without this
	// property will pass the filter. If null or omitted, the existence of the property
	// is optional.
	Exists bool `json:"exists"`
	// Specifies the allowed values for the property to match an event. An event will
	// pass the filter only if its property value is included in this list. If
	// undefined, all property values will pass the filter. Must be non-empty if
	// present.
	InValues []string `json:"in_values"`
	// Specifies the values that prevent an event from matching the filter. An event
	// will not pass the filter if its property value is included in this list. If null
	// or empty, all property values will pass the filter. Must be non-empty if
	// present.
	NotInValues []string                                                `json:"not_in_values"`
	JSON        v1CustomerListBillableMetricsResponsePropertyFilterJSON `json:"-"`
}

// v1CustomerListBillableMetricsResponsePropertyFilterJSON contains the JSON
// metadata for the struct [V1CustomerListBillableMetricsResponsePropertyFilter]
type v1CustomerListBillableMetricsResponsePropertyFilterJSON struct {
	Name        apijson.Field
	Exists      apijson.Field
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerListBillableMetricsResponsePropertyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListBillableMetricsResponsePropertyFilterJSON) RawJSON() string {
	return r.raw
}

type V1CustomerListCostsResponse struct {
	CreditTypes    map[string]V1CustomerListCostsResponseCreditType `json:"credit_types,required"`
	EndTimestamp   time.Time                                        `json:"end_timestamp,required" format:"date-time"`
	StartTimestamp time.Time                                        `json:"start_timestamp,required" format:"date-time"`
	JSON           v1CustomerListCostsResponseJSON                  `json:"-"`
}

// v1CustomerListCostsResponseJSON contains the JSON metadata for the struct
// [V1CustomerListCostsResponse]
type v1CustomerListCostsResponseJSON struct {
	CreditTypes    apijson.Field
	EndTimestamp   apijson.Field
	StartTimestamp apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1CustomerListCostsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListCostsResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerListCostsResponseCreditType struct {
	Cost              float64                                                   `json:"cost"`
	LineItemBreakdown []V1CustomerListCostsResponseCreditTypesLineItemBreakdown `json:"line_item_breakdown"`
	Name              string                                                    `json:"name"`
	JSON              v1CustomerListCostsResponseCreditTypeJSON                 `json:"-"`
}

// v1CustomerListCostsResponseCreditTypeJSON contains the JSON metadata for the
// struct [V1CustomerListCostsResponseCreditType]
type v1CustomerListCostsResponseCreditTypeJSON struct {
	Cost              apijson.Field
	LineItemBreakdown apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *V1CustomerListCostsResponseCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListCostsResponseCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerListCostsResponseCreditTypesLineItemBreakdown struct {
	Cost       float64                                                     `json:"cost,required"`
	Name       string                                                      `json:"name,required"`
	GroupKey   string                                                      `json:"group_key"`
	GroupValue string                                                      `json:"group_value,nullable"`
	JSON       v1CustomerListCostsResponseCreditTypesLineItemBreakdownJSON `json:"-"`
}

// v1CustomerListCostsResponseCreditTypesLineItemBreakdownJSON contains the JSON
// metadata for the struct
// [V1CustomerListCostsResponseCreditTypesLineItemBreakdown]
type v1CustomerListCostsResponseCreditTypesLineItemBreakdownJSON struct {
	Cost        apijson.Field
	Name        apijson.Field
	GroupKey    apijson.Field
	GroupValue  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerListCostsResponseCreditTypesLineItemBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerListCostsResponseCreditTypesLineItemBreakdownJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponse struct {
	Data V1CustomerPreviewEventsResponseData `json:"data,required"`
	JSON v1CustomerPreviewEventsResponseJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseJSON contains the JSON metadata for the struct
// [V1CustomerPreviewEventsResponse]
type v1CustomerPreviewEventsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseData struct {
	ID          string                                        `json:"id,required" format:"uuid"`
	CreditType  V1CustomerPreviewEventsResponseDataCreditType `json:"credit_type,required"`
	CustomerID  string                                        `json:"customer_id,required" format:"uuid"`
	LineItems   []V1CustomerPreviewEventsResponseDataLineItem `json:"line_items,required"`
	Status      string                                        `json:"status,required"`
	Total       float64                                       `json:"total,required"`
	Type        string                                        `json:"type,required"`
	AmendmentID string                                        `json:"amendment_id" format:"uuid"`
	// This field's availability is dependent on your client's configuration.
	BillableStatus V1CustomerPreviewEventsResponseDataBillableStatus `json:"billable_status"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	ContractCustomFields map[string]string                                   `json:"contract_custom_fields"`
	ContractID           string                                              `json:"contract_id" format:"uuid"`
	CorrectionRecord     V1CustomerPreviewEventsResponseDataCorrectionRecord `json:"correction_record"`
	// When the invoice was created (UTC). This field is present for correction
	// invoices only.
	CreatedAt    time.Time              `json:"created_at" format:"date-time"`
	CustomFields map[string]interface{} `json:"custom_fields"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomerCustomFields map[string]string `json:"customer_custom_fields"`
	// End of the usage period this invoice covers (UTC)
	EndTimestamp       time.Time                                              `json:"end_timestamp" format:"date-time"`
	ExternalInvoice    V1CustomerPreviewEventsResponseDataExternalInvoice     `json:"external_invoice,nullable"`
	InvoiceAdjustments []V1CustomerPreviewEventsResponseDataInvoiceAdjustment `json:"invoice_adjustments"`
	// When the invoice was issued (UTC)
	IssuedAt            time.Time `json:"issued_at" format:"date-time"`
	NetPaymentTermsDays float64   `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	PlanCustomFields map[string]string `json:"plan_custom_fields"`
	PlanID           string            `json:"plan_id" format:"uuid"`
	PlanName         string            `json:"plan_name"`
	// Only present for contract invoices with reseller royalties.
	ResellerRoyalty V1CustomerPreviewEventsResponseDataResellerRoyalty `json:"reseller_royalty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Beginning of the usage period this invoice covers (UTC)
	StartTimestamp time.Time                               `json:"start_timestamp" format:"date-time"`
	Subtotal       float64                                 `json:"subtotal"`
	JSON           v1CustomerPreviewEventsResponseDataJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataJSON contains the JSON metadata for the
// struct [V1CustomerPreviewEventsResponseData]
type v1CustomerPreviewEventsResponseDataJSON struct {
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

func (r *V1CustomerPreviewEventsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataCreditType struct {
	ID   string                                            `json:"id,required" format:"uuid"`
	Name string                                            `json:"name,required"`
	JSON v1CustomerPreviewEventsResponseDataCreditTypeJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataCreditTypeJSON contains the JSON metadata for
// the struct [V1CustomerPreviewEventsResponseDataCreditType]
type v1CustomerPreviewEventsResponseDataCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataLineItem struct {
	CreditType V1CustomerPreviewEventsResponseDataLineItemsCreditType `json:"credit_type,required"`
	Name       string                                                 `json:"name,required"`
	Total      float64                                                `json:"total,required"`
	// The type of line item.
	//
	//   - `scheduled`: Line item is associated with a scheduled charge. View the
	//     scheduled_charge_id on the line item.
	//   - `commit_purchase`: Line item is associated with a payment for a prepaid
	//     commit. View the commit_id on the line item.
	//   - `usage`: Line item is associated with a usage product or composite product.
	//     View the product_id on the line item to determine which product.
	//   - `subscription`: Line item is associated with a subscription. e.g. monthly
	//     recurring payment for an in-advance subscription.
	//   - `applied_commit_or_credit`: On metronome invoices, applied commits and credits
	//     are associated with their own line items. These line items have negative
	//     totals. Use the applied_commit_or_credit object on the line item to understand
	//     the id of the applied commit or credit, and its type. Note that the
	//     application of a postpaid commit is associated with a line item, but the total
	//     on the line item is not included in the invoice's total as postpaid commits
	//     are paid in-arrears.
	//   - `cpu_conversion`: Line item converting between a custom pricing unit and fiat
	//     currency, using the conversion rate set on the rate card. This line item will
	//     appear when there are products priced in custom pricing units, and there is
	//     insufficient prepaid commit/credit in that custom pricing unit to fully cover
	//     the spend. Then, the outstanding spend in custom pricing units will be
	//     converted to fiat currency using a cpu_conversion line item.
	Type string `json:"type,required"`
	// Details about the credit or commit that was applied to this line item. Only
	// present on line items with product of `USAGE`, `SUBSCRIPTION` or `COMPOSITE`
	// types.
	AppliedCommitOrCredit V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCredit `json:"applied_commit_or_credit"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CommitCustomFields map[string]string `json:"commit_custom_fields"`
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
	CommitType string `json:"commit_type"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
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
	ListPrice V1CustomerPreviewEventsResponseDataLineItemsListPrice `json:"list_price"`
	Metadata  string                                                `json:"metadata"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd time.Time `json:"netsuite_invoice_billing_end" format:"date-time"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart time.Time `json:"netsuite_invoice_billing_start" format:"date-time"`
	NetsuiteItemID              string    `json:"netsuite_item_id"`
	// Only present for line items paying for a postpaid commit true-up.
	PostpaidCommit V1CustomerPreviewEventsResponseDataLineItemsPostpaidCommit `json:"postpaid_commit"`
	// Includes the presentation group values associated with this line item if
	// presentation group keys are used.
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	// Includes the pricing group values associated with this line item if dimensional
	// pricing is used.
	PricingGroupValues map[string]string `json:"pricing_group_values"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
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
	ProductType string `json:"product_type"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	ProfessionalServiceCustomFields map[string]string `json:"professional_service_custom_fields"`
	ProfessionalServiceID           string            `json:"professional_service_id" format:"uuid"`
	// The quantity associated with the line item.
	Quantity     float64                                                  `json:"quantity"`
	ResellerType V1CustomerPreviewEventsResponseDataLineItemsResellerType `json:"reseller_type"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	ScheduledChargeCustomFields map[string]string `json:"scheduled_charge_custom_fields"`
	// ID of scheduled charge.
	ScheduledChargeID string `json:"scheduled_charge_id" format:"uuid"`
	// The line item's start date (inclusive).
	StartingAt   time.Time                                                 `json:"starting_at" format:"date-time"`
	SubLineItems []V1CustomerPreviewEventsResponseDataLineItemsSubLineItem `json:"sub_line_items"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	SubscriptionCustomFields map[string]string `json:"subscription_custom_fields"`
	// Populated if the line item has a tiered price.
	Tier V1CustomerPreviewEventsResponseDataLineItemsTier `json:"tier"`
	// The unit price associated with the line item.
	UnitPrice float64                                         `json:"unit_price"`
	JSON      v1CustomerPreviewEventsResponseDataLineItemJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataLineItemJSON contains the JSON metadata for
// the struct [V1CustomerPreviewEventsResponseDataLineItem]
type v1CustomerPreviewEventsResponseDataLineItemJSON struct {
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

func (r *V1CustomerPreviewEventsResponseDataLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataLineItemJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataLineItemsCreditType struct {
	ID   string                                                     `json:"id,required" format:"uuid"`
	Name string                                                     `json:"name,required"`
	JSON v1CustomerPreviewEventsResponseDataLineItemsCreditTypeJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataLineItemsCreditTypeJSON contains the JSON
// metadata for the struct [V1CustomerPreviewEventsResponseDataLineItemsCreditType]
type v1CustomerPreviewEventsResponseDataLineItemsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataLineItemsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataLineItemsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// Details about the credit or commit that was applied to this line item. Only
// present on line items with product of `USAGE`, `SUBSCRIPTION` or `COMPOSITE`
// types.
type V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCredit struct {
	ID   string                                                                `json:"id,required" format:"uuid"`
	Type V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditType `json:"type,required"`
	JSON v1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditJSON contains
// the JSON metadata for the struct
// [V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCredit]
type v1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditType string

const (
	V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditTypePrepaid  V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditType = "PREPAID"
	V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditTypePostpaid V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditType = "POSTPAID"
	V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditTypeCredit   V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditType = "CREDIT"
)

func (r V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditType) IsKnown() bool {
	switch r {
	case V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditTypePrepaid, V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditTypePostpaid, V1CustomerPreviewEventsResponseDataLineItemsAppliedCommitOrCreditTypeCredit:
		return true
	}
	return false
}

// Only present for contract invoices and when the `include_list_prices` query
// parameter is set to true. This will include the list rate for the charge if
// applicable. Only present for usage and subscription line items.
type V1CustomerPreviewEventsResponseDataLineItemsListPrice struct {
	RateType   V1CustomerPreviewEventsResponseDataLineItemsListPriceRateType   `json:"rate_type,required"`
	CreditType V1CustomerPreviewEventsResponseDataLineItemsListPriceCreditType `json:"credit_type"`
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
	Tiers []V1CustomerPreviewEventsResponseDataLineItemsListPriceTier `json:"tiers"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool                                                      `json:"use_list_prices"`
	JSON          v1CustomerPreviewEventsResponseDataLineItemsListPriceJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataLineItemsListPriceJSON contains the JSON
// metadata for the struct [V1CustomerPreviewEventsResponseDataLineItemsListPrice]
type v1CustomerPreviewEventsResponseDataLineItemsListPriceJSON struct {
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

func (r *V1CustomerPreviewEventsResponseDataLineItemsListPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataLineItemsListPriceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataLineItemsListPriceRateType string

const (
	V1CustomerPreviewEventsResponseDataLineItemsListPriceRateTypeFlat         V1CustomerPreviewEventsResponseDataLineItemsListPriceRateType = "FLAT"
	V1CustomerPreviewEventsResponseDataLineItemsListPriceRateTypePercentage   V1CustomerPreviewEventsResponseDataLineItemsListPriceRateType = "PERCENTAGE"
	V1CustomerPreviewEventsResponseDataLineItemsListPriceRateTypeSubscription V1CustomerPreviewEventsResponseDataLineItemsListPriceRateType = "SUBSCRIPTION"
	V1CustomerPreviewEventsResponseDataLineItemsListPriceRateTypeCustom       V1CustomerPreviewEventsResponseDataLineItemsListPriceRateType = "CUSTOM"
	V1CustomerPreviewEventsResponseDataLineItemsListPriceRateTypeTiered       V1CustomerPreviewEventsResponseDataLineItemsListPriceRateType = "TIERED"
)

func (r V1CustomerPreviewEventsResponseDataLineItemsListPriceRateType) IsKnown() bool {
	switch r {
	case V1CustomerPreviewEventsResponseDataLineItemsListPriceRateTypeFlat, V1CustomerPreviewEventsResponseDataLineItemsListPriceRateTypePercentage, V1CustomerPreviewEventsResponseDataLineItemsListPriceRateTypeSubscription, V1CustomerPreviewEventsResponseDataLineItemsListPriceRateTypeCustom, V1CustomerPreviewEventsResponseDataLineItemsListPriceRateTypeTiered:
		return true
	}
	return false
}

type V1CustomerPreviewEventsResponseDataLineItemsListPriceCreditType struct {
	ID   string                                                              `json:"id,required" format:"uuid"`
	Name string                                                              `json:"name,required"`
	JSON v1CustomerPreviewEventsResponseDataLineItemsListPriceCreditTypeJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataLineItemsListPriceCreditTypeJSON contains the
// JSON metadata for the struct
// [V1CustomerPreviewEventsResponseDataLineItemsListPriceCreditType]
type v1CustomerPreviewEventsResponseDataLineItemsListPriceCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataLineItemsListPriceCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataLineItemsListPriceCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataLineItemsListPriceTier struct {
	Price float64                                                       `json:"price,required"`
	Size  float64                                                       `json:"size"`
	JSON  v1CustomerPreviewEventsResponseDataLineItemsListPriceTierJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataLineItemsListPriceTierJSON contains the JSON
// metadata for the struct
// [V1CustomerPreviewEventsResponseDataLineItemsListPriceTier]
type v1CustomerPreviewEventsResponseDataLineItemsListPriceTierJSON struct {
	Price       apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataLineItemsListPriceTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataLineItemsListPriceTierJSON) RawJSON() string {
	return r.raw
}

// Only present for line items paying for a postpaid commit true-up.
type V1CustomerPreviewEventsResponseDataLineItemsPostpaidCommit struct {
	ID   string                                                         `json:"id,required" format:"uuid"`
	JSON v1CustomerPreviewEventsResponseDataLineItemsPostpaidCommitJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataLineItemsPostpaidCommitJSON contains the JSON
// metadata for the struct
// [V1CustomerPreviewEventsResponseDataLineItemsPostpaidCommit]
type v1CustomerPreviewEventsResponseDataLineItemsPostpaidCommitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataLineItemsPostpaidCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataLineItemsPostpaidCommitJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataLineItemsResellerType string

const (
	V1CustomerPreviewEventsResponseDataLineItemsResellerTypeAws           V1CustomerPreviewEventsResponseDataLineItemsResellerType = "AWS"
	V1CustomerPreviewEventsResponseDataLineItemsResellerTypeAwsProService V1CustomerPreviewEventsResponseDataLineItemsResellerType = "AWS_PRO_SERVICE"
	V1CustomerPreviewEventsResponseDataLineItemsResellerTypeGcp           V1CustomerPreviewEventsResponseDataLineItemsResellerType = "GCP"
	V1CustomerPreviewEventsResponseDataLineItemsResellerTypeGcpProService V1CustomerPreviewEventsResponseDataLineItemsResellerType = "GCP_PRO_SERVICE"
)

func (r V1CustomerPreviewEventsResponseDataLineItemsResellerType) IsKnown() bool {
	switch r {
	case V1CustomerPreviewEventsResponseDataLineItemsResellerTypeAws, V1CustomerPreviewEventsResponseDataLineItemsResellerTypeAwsProService, V1CustomerPreviewEventsResponseDataLineItemsResellerTypeGcp, V1CustomerPreviewEventsResponseDataLineItemsResellerTypeGcpProService:
		return true
	}
	return false
}

type V1CustomerPreviewEventsResponseDataLineItemsSubLineItem struct {
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
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
	TierPeriod V1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierPeriod `json:"tier_period"`
	Tiers      []V1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTier     `json:"tiers"`
	JSON       v1CustomerPreviewEventsResponseDataLineItemsSubLineItemJSON        `json:"-"`
}

// v1CustomerPreviewEventsResponseDataLineItemsSubLineItemJSON contains the JSON
// metadata for the struct
// [V1CustomerPreviewEventsResponseDataLineItemsSubLineItem]
type v1CustomerPreviewEventsResponseDataLineItemsSubLineItemJSON struct {
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

func (r *V1CustomerPreviewEventsResponseDataLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataLineItemsSubLineItemJSON) RawJSON() string {
	return r.raw
}

// when the current tier started and ends (for tiered charges only)
type V1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierPeriod struct {
	StartingAt   time.Time                                                              `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time                                                              `json:"ending_before" format:"date-time"`
	JSON         v1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierPeriodJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierPeriodJSON contains
// the JSON metadata for the struct
// [V1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierPeriod]
type v1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierPeriodJSON struct {
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierPeriodJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTier struct {
	Price    float64 `json:"price,required"`
	Quantity float64 `json:"quantity,required"`
	// at what metric amount this tier begins
	StartingAt float64                                                          `json:"starting_at,required"`
	Subtotal   float64                                                          `json:"subtotal,required"`
	JSON       v1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierJSON contains the
// JSON metadata for the struct
// [V1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTier]
type v1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierJSON struct {
	Price       apijson.Field
	Quantity    apijson.Field
	StartingAt  apijson.Field
	Subtotal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataLineItemsSubLineItemsTierJSON) RawJSON() string {
	return r.raw
}

// Populated if the line item has a tiered price.
type V1CustomerPreviewEventsResponseDataLineItemsTier struct {
	Level      float64                                              `json:"level,required"`
	StartingAt string                                               `json:"starting_at,required"`
	Size       string                                               `json:"size,nullable"`
	JSON       v1CustomerPreviewEventsResponseDataLineItemsTierJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataLineItemsTierJSON contains the JSON metadata
// for the struct [V1CustomerPreviewEventsResponseDataLineItemsTier]
type v1CustomerPreviewEventsResponseDataLineItemsTierJSON struct {
	Level       apijson.Field
	StartingAt  apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataLineItemsTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataLineItemsTierJSON) RawJSON() string {
	return r.raw
}

// This field's availability is dependent on your client's configuration.
type V1CustomerPreviewEventsResponseDataBillableStatus string

const (
	V1CustomerPreviewEventsResponseDataBillableStatusBillable   V1CustomerPreviewEventsResponseDataBillableStatus = "billable"
	V1CustomerPreviewEventsResponseDataBillableStatusUnbillable V1CustomerPreviewEventsResponseDataBillableStatus = "unbillable"
)

func (r V1CustomerPreviewEventsResponseDataBillableStatus) IsKnown() bool {
	switch r {
	case V1CustomerPreviewEventsResponseDataBillableStatusBillable, V1CustomerPreviewEventsResponseDataBillableStatusUnbillable:
		return true
	}
	return false
}

type V1CustomerPreviewEventsResponseDataCorrectionRecord struct {
	CorrectedInvoiceID       string                                                                      `json:"corrected_invoice_id,required" format:"uuid"`
	Memo                     string                                                                      `json:"memo,required"`
	Reason                   string                                                                      `json:"reason,required"`
	CorrectedExternalInvoice V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoice `json:"corrected_external_invoice"`
	JSON                     v1CustomerPreviewEventsResponseDataCorrectionRecordJSON                     `json:"-"`
}

// v1CustomerPreviewEventsResponseDataCorrectionRecordJSON contains the JSON
// metadata for the struct [V1CustomerPreviewEventsResponseDataCorrectionRecord]
type v1CustomerPreviewEventsResponseDataCorrectionRecordJSON struct {
	CorrectedInvoiceID       apijson.Field
	Memo                     apijson.Field
	Reason                   apijson.Field
	CorrectedExternalInvoice apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataCorrectionRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataCorrectionRecordJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoice struct {
	BillingProviderType V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                                                         `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                                                      `json:"issued_at_timestamp" format:"date-time"`
	JSON                v1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceJSON                `json:"-"`
}

// v1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceJSON
// contains the JSON metadata for the struct
// [V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoice]
type v1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType string

const (
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace   V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "aws_marketplace"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe           V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "stripe"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite         V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "netsuite"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom           V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "custom"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "azure_marketplace"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "quickbooks_online"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday          V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "workday"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace   V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus string

const (
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft               V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DRAFT"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized           V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "FINALIZED"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid                V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAID"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible       V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid                V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "VOID"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted             V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DELETED"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed       V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped             V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SKIPPED"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSent                V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SENT"
	V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued              V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "QUEUED"
)

func (r V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSent, V1CustomerPreviewEventsResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type V1CustomerPreviewEventsResponseDataExternalInvoice struct {
	BillingProviderType V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                                `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                             `json:"issued_at_timestamp" format:"date-time"`
	JSON                v1CustomerPreviewEventsResponseDataExternalInvoiceJSON                `json:"-"`
}

// v1CustomerPreviewEventsResponseDataExternalInvoiceJSON contains the JSON
// metadata for the struct [V1CustomerPreviewEventsResponseDataExternalInvoice]
type v1CustomerPreviewEventsResponseDataExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderType string

const (
	V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeAwsMarketplace   V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderType = "aws_marketplace"
	V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeStripe           V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderType = "stripe"
	V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeNetsuite         V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderType = "netsuite"
	V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeCustom           V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderType = "custom"
	V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeAzureMarketplace V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderType = "azure_marketplace"
	V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeQuickbooksOnline V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderType = "quickbooks_online"
	V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeWorkday          V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderType = "workday"
	V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeGcpMarketplace   V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeAwsMarketplace, V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeStripe, V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeNetsuite, V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeCustom, V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeAzureMarketplace, V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeQuickbooksOnline, V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeWorkday, V1CustomerPreviewEventsResponseDataExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus string

const (
	V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusDraft               V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus = "DRAFT"
	V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusFinalized           V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus = "FINALIZED"
	V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusPaid                V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus = "PAID"
	V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusUncollectible       V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusVoid                V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus = "VOID"
	V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusDeleted             V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus = "DELETED"
	V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusPaymentFailed       V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusInvalidRequestError V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusSkipped             V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus = "SKIPPED"
	V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusSent                V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus = "SENT"
	V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusQueued              V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus = "QUEUED"
)

func (r V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusDraft, V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusFinalized, V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusPaid, V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusUncollectible, V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusVoid, V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusDeleted, V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusPaymentFailed, V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusInvalidRequestError, V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusSkipped, V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusSent, V1CustomerPreviewEventsResponseDataExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type V1CustomerPreviewEventsResponseDataInvoiceAdjustment struct {
	CreditType V1CustomerPreviewEventsResponseDataInvoiceAdjustmentsCreditType `json:"credit_type,required"`
	Name       string                                                          `json:"name,required"`
	Total      float64                                                         `json:"total,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CreditGrantCustomFields map[string]string                                        `json:"credit_grant_custom_fields"`
	CreditGrantID           string                                                   `json:"credit_grant_id"`
	JSON                    v1CustomerPreviewEventsResponseDataInvoiceAdjustmentJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataInvoiceAdjustmentJSON contains the JSON
// metadata for the struct [V1CustomerPreviewEventsResponseDataInvoiceAdjustment]
type v1CustomerPreviewEventsResponseDataInvoiceAdjustmentJSON struct {
	CreditType              apijson.Field
	Name                    apijson.Field
	Total                   apijson.Field
	CreditGrantCustomFields apijson.Field
	CreditGrantID           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataInvoiceAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataInvoiceAdjustmentJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataInvoiceAdjustmentsCreditType struct {
	ID   string                                                              `json:"id,required" format:"uuid"`
	Name string                                                              `json:"name,required"`
	JSON v1CustomerPreviewEventsResponseDataInvoiceAdjustmentsCreditTypeJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataInvoiceAdjustmentsCreditTypeJSON contains the
// JSON metadata for the struct
// [V1CustomerPreviewEventsResponseDataInvoiceAdjustmentsCreditType]
type v1CustomerPreviewEventsResponseDataInvoiceAdjustmentsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataInvoiceAdjustmentsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataInvoiceAdjustmentsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// Only present for contract invoices with reseller royalties.
type V1CustomerPreviewEventsResponseDataResellerRoyalty struct {
	Fraction           string                                                         `json:"fraction,required"`
	NetsuiteResellerID string                                                         `json:"netsuite_reseller_id,required"`
	ResellerType       V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerType `json:"reseller_type,required"`
	AwsOptions         V1CustomerPreviewEventsResponseDataResellerRoyaltyAwsOptions   `json:"aws_options"`
	GcpOptions         V1CustomerPreviewEventsResponseDataResellerRoyaltyGcpOptions   `json:"gcp_options"`
	JSON               v1CustomerPreviewEventsResponseDataResellerRoyaltyJSON         `json:"-"`
}

// v1CustomerPreviewEventsResponseDataResellerRoyaltyJSON contains the JSON
// metadata for the struct [V1CustomerPreviewEventsResponseDataResellerRoyalty]
type v1CustomerPreviewEventsResponseDataResellerRoyaltyJSON struct {
	Fraction           apijson.Field
	NetsuiteResellerID apijson.Field
	ResellerType       apijson.Field
	AwsOptions         apijson.Field
	GcpOptions         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerType string

const (
	V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerTypeAws           V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerType = "AWS"
	V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerTypeAwsProService V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerType = "AWS_PRO_SERVICE"
	V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerTypeGcp           V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerType = "GCP"
	V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerTypeGcpProService V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerType = "GCP_PRO_SERVICE"
)

func (r V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerType) IsKnown() bool {
	switch r {
	case V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerTypeAws, V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerTypeAwsProService, V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerTypeGcp, V1CustomerPreviewEventsResponseDataResellerRoyaltyResellerTypeGcpProService:
		return true
	}
	return false
}

type V1CustomerPreviewEventsResponseDataResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    string                                                           `json:"aws_account_number"`
	AwsOfferID          string                                                           `json:"aws_offer_id"`
	AwsPayerReferenceID string                                                           `json:"aws_payer_reference_id"`
	JSON                v1CustomerPreviewEventsResponseDataResellerRoyaltyAwsOptionsJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataResellerRoyaltyAwsOptionsJSON contains the
// JSON metadata for the struct
// [V1CustomerPreviewEventsResponseDataResellerRoyaltyAwsOptions]
type v1CustomerPreviewEventsResponseDataResellerRoyaltyAwsOptionsJSON struct {
	AwsAccountNumber    apijson.Field
	AwsOfferID          apijson.Field
	AwsPayerReferenceID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataResellerRoyaltyAwsOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataResellerRoyaltyAwsOptionsJSON) RawJSON() string {
	return r.raw
}

type V1CustomerPreviewEventsResponseDataResellerRoyaltyGcpOptions struct {
	GcpAccountID string                                                           `json:"gcp_account_id"`
	GcpOfferID   string                                                           `json:"gcp_offer_id"`
	JSON         v1CustomerPreviewEventsResponseDataResellerRoyaltyGcpOptionsJSON `json:"-"`
}

// v1CustomerPreviewEventsResponseDataResellerRoyaltyGcpOptionsJSON contains the
// JSON metadata for the struct
// [V1CustomerPreviewEventsResponseDataResellerRoyaltyGcpOptions]
type v1CustomerPreviewEventsResponseDataResellerRoyaltyGcpOptionsJSON struct {
	GcpAccountID apijson.Field
	GcpOfferID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerPreviewEventsResponseDataResellerRoyaltyGcpOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerPreviewEventsResponseDataResellerRoyaltyGcpOptionsJSON) RawJSON() string {
	return r.raw
}

type V1CustomerGetBillingConfigurationsResponse struct {
	Data []V1CustomerGetBillingConfigurationsResponseData `json:"data,required"`
	JSON v1CustomerGetBillingConfigurationsResponseJSON   `json:"-"`
}

// v1CustomerGetBillingConfigurationsResponseJSON contains the JSON metadata for
// the struct [V1CustomerGetBillingConfigurationsResponse]
type v1CustomerGetBillingConfigurationsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerGetBillingConfigurationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerGetBillingConfigurationsResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerGetBillingConfigurationsResponseData struct {
	// ID of this configuration; can be provided as the
	// billing_provider_configuration_id when creating a contract.
	ID string `json:"id,required" format:"uuid"`
	// The billing provider set for this configuration.
	BillingProvider V1CustomerGetBillingConfigurationsResponseDataBillingProvider `json:"billing_provider,required"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider.
	Configuration map[string]interface{} `json:"configuration,required"`
	CustomerID    string                 `json:"customer_id,required" format:"uuid"`
	// The method to use for delivering invoices to this customer.
	DeliveryMethod V1CustomerGetBillingConfigurationsResponseDataDeliveryMethod `json:"delivery_method,required"`
	// Configuration for the delivery method. The structure of this object is specific
	// to the delivery method.
	DeliveryMethodConfiguration map[string]interface{} `json:"delivery_method_configuration,required"`
	// ID of the delivery method to use for this customer.
	DeliveryMethodID string                                             `json:"delivery_method_id,required" format:"uuid"`
	JSON             v1CustomerGetBillingConfigurationsResponseDataJSON `json:"-"`
}

// v1CustomerGetBillingConfigurationsResponseDataJSON contains the JSON metadata
// for the struct [V1CustomerGetBillingConfigurationsResponseData]
type v1CustomerGetBillingConfigurationsResponseDataJSON struct {
	ID                          apijson.Field
	BillingProvider             apijson.Field
	Configuration               apijson.Field
	CustomerID                  apijson.Field
	DeliveryMethod              apijson.Field
	DeliveryMethodConfiguration apijson.Field
	DeliveryMethodID            apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *V1CustomerGetBillingConfigurationsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerGetBillingConfigurationsResponseDataJSON) RawJSON() string {
	return r.raw
}

// The billing provider set for this configuration.
type V1CustomerGetBillingConfigurationsResponseDataBillingProvider string

const (
	V1CustomerGetBillingConfigurationsResponseDataBillingProviderAwsMarketplace   V1CustomerGetBillingConfigurationsResponseDataBillingProvider = "aws_marketplace"
	V1CustomerGetBillingConfigurationsResponseDataBillingProviderStripe           V1CustomerGetBillingConfigurationsResponseDataBillingProvider = "stripe"
	V1CustomerGetBillingConfigurationsResponseDataBillingProviderNetsuite         V1CustomerGetBillingConfigurationsResponseDataBillingProvider = "netsuite"
	V1CustomerGetBillingConfigurationsResponseDataBillingProviderCustom           V1CustomerGetBillingConfigurationsResponseDataBillingProvider = "custom"
	V1CustomerGetBillingConfigurationsResponseDataBillingProviderAzureMarketplace V1CustomerGetBillingConfigurationsResponseDataBillingProvider = "azure_marketplace"
	V1CustomerGetBillingConfigurationsResponseDataBillingProviderQuickbooksOnline V1CustomerGetBillingConfigurationsResponseDataBillingProvider = "quickbooks_online"
	V1CustomerGetBillingConfigurationsResponseDataBillingProviderWorkday          V1CustomerGetBillingConfigurationsResponseDataBillingProvider = "workday"
	V1CustomerGetBillingConfigurationsResponseDataBillingProviderGcpMarketplace   V1CustomerGetBillingConfigurationsResponseDataBillingProvider = "gcp_marketplace"
)

func (r V1CustomerGetBillingConfigurationsResponseDataBillingProvider) IsKnown() bool {
	switch r {
	case V1CustomerGetBillingConfigurationsResponseDataBillingProviderAwsMarketplace, V1CustomerGetBillingConfigurationsResponseDataBillingProviderStripe, V1CustomerGetBillingConfigurationsResponseDataBillingProviderNetsuite, V1CustomerGetBillingConfigurationsResponseDataBillingProviderCustom, V1CustomerGetBillingConfigurationsResponseDataBillingProviderAzureMarketplace, V1CustomerGetBillingConfigurationsResponseDataBillingProviderQuickbooksOnline, V1CustomerGetBillingConfigurationsResponseDataBillingProviderWorkday, V1CustomerGetBillingConfigurationsResponseDataBillingProviderGcpMarketplace:
		return true
	}
	return false
}

// The method to use for delivering invoices to this customer.
type V1CustomerGetBillingConfigurationsResponseDataDeliveryMethod string

const (
	V1CustomerGetBillingConfigurationsResponseDataDeliveryMethodDirectToBillingProvider V1CustomerGetBillingConfigurationsResponseDataDeliveryMethod = "direct_to_billing_provider"
	V1CustomerGetBillingConfigurationsResponseDataDeliveryMethodAwsSqs                  V1CustomerGetBillingConfigurationsResponseDataDeliveryMethod = "aws_sqs"
	V1CustomerGetBillingConfigurationsResponseDataDeliveryMethodTackle                  V1CustomerGetBillingConfigurationsResponseDataDeliveryMethod = "tackle"
	V1CustomerGetBillingConfigurationsResponseDataDeliveryMethodAwsSns                  V1CustomerGetBillingConfigurationsResponseDataDeliveryMethod = "aws_sns"
)

func (r V1CustomerGetBillingConfigurationsResponseDataDeliveryMethod) IsKnown() bool {
	switch r {
	case V1CustomerGetBillingConfigurationsResponseDataDeliveryMethodDirectToBillingProvider, V1CustomerGetBillingConfigurationsResponseDataDeliveryMethodAwsSqs, V1CustomerGetBillingConfigurationsResponseDataDeliveryMethodTackle, V1CustomerGetBillingConfigurationsResponseDataDeliveryMethodAwsSns:
		return true
	}
	return false
}

type V1CustomerSetNameResponse struct {
	Data V1CustomerSetNameResponseData `json:"data,required"`
	JSON v1CustomerSetNameResponseJSON `json:"-"`
}

// v1CustomerSetNameResponseJSON contains the JSON metadata for the struct
// [V1CustomerSetNameResponse]
type v1CustomerSetNameResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerSetNameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerSetNameResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerSetNameResponseData struct {
	// the Metronome ID of the customer
	ID string `json:"id,required" format:"uuid"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string `json:"ingest_aliases,required"`
	Name          string   `json:"name,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string                 `json:"custom_fields"`
	JSON         v1CustomerSetNameResponseDataJSON `json:"-"`
}

// v1CustomerSetNameResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerSetNameResponseData]
type v1CustomerSetNameResponseDataJSON struct {
	ID            apijson.Field
	ExternalID    apijson.Field
	IngestAliases apijson.Field
	Name          apijson.Field
	CustomFields  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1CustomerSetNameResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerSetNameResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerNewParams struct {
	// This will be truncated to 160 characters if the provided name is longer.
	Name          param.Field[string]                           `json:"name,required"`
	BillingConfig param.Field[V1CustomerNewParamsBillingConfig] `json:"billing_config"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields                          param.Field[map[string]string]                                         `json:"custom_fields"`
	CustomerBillingProviderConfigurations param.Field[[]V1CustomerNewParamsCustomerBillingProviderConfiguration] `json:"customer_billing_provider_configurations"`
	// (deprecated, use ingest_aliases instead) an alias that can be used to refer to
	// this customer in usage events
	ExternalID param.Field[string] `json:"external_id"`
	// Aliases that can be used to refer to this customer in usage events
	IngestAliases param.Field[[]string] `json:"ingest_aliases"`
}

func (r V1CustomerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerNewParamsBillingConfig struct {
	BillingProviderCustomerID param.Field[string]                                              `json:"billing_provider_customer_id,required"`
	BillingProviderType       param.Field[V1CustomerNewParamsBillingConfigBillingProviderType] `json:"billing_provider_type,required"`
	// True if the aws_product_code is a SAAS subscription product, false otherwise.
	AwsIsSubscriptionProduct param.Field[bool]                                                   `json:"aws_is_subscription_product"`
	AwsProductCode           param.Field[string]                                                 `json:"aws_product_code"`
	AwsRegion                param.Field[V1CustomerNewParamsBillingConfigAwsRegion]              `json:"aws_region"`
	StripeCollectionMethod   param.Field[V1CustomerNewParamsBillingConfigStripeCollectionMethod] `json:"stripe_collection_method"`
}

func (r V1CustomerNewParamsBillingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerNewParamsBillingConfigBillingProviderType string

const (
	V1CustomerNewParamsBillingConfigBillingProviderTypeAwsMarketplace   V1CustomerNewParamsBillingConfigBillingProviderType = "aws_marketplace"
	V1CustomerNewParamsBillingConfigBillingProviderTypeStripe           V1CustomerNewParamsBillingConfigBillingProviderType = "stripe"
	V1CustomerNewParamsBillingConfigBillingProviderTypeNetsuite         V1CustomerNewParamsBillingConfigBillingProviderType = "netsuite"
	V1CustomerNewParamsBillingConfigBillingProviderTypeCustom           V1CustomerNewParamsBillingConfigBillingProviderType = "custom"
	V1CustomerNewParamsBillingConfigBillingProviderTypeAzureMarketplace V1CustomerNewParamsBillingConfigBillingProviderType = "azure_marketplace"
	V1CustomerNewParamsBillingConfigBillingProviderTypeQuickbooksOnline V1CustomerNewParamsBillingConfigBillingProviderType = "quickbooks_online"
	V1CustomerNewParamsBillingConfigBillingProviderTypeWorkday          V1CustomerNewParamsBillingConfigBillingProviderType = "workday"
	V1CustomerNewParamsBillingConfigBillingProviderTypeGcpMarketplace   V1CustomerNewParamsBillingConfigBillingProviderType = "gcp_marketplace"
)

func (r V1CustomerNewParamsBillingConfigBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerNewParamsBillingConfigBillingProviderTypeAwsMarketplace, V1CustomerNewParamsBillingConfigBillingProviderTypeStripe, V1CustomerNewParamsBillingConfigBillingProviderTypeNetsuite, V1CustomerNewParamsBillingConfigBillingProviderTypeCustom, V1CustomerNewParamsBillingConfigBillingProviderTypeAzureMarketplace, V1CustomerNewParamsBillingConfigBillingProviderTypeQuickbooksOnline, V1CustomerNewParamsBillingConfigBillingProviderTypeWorkday, V1CustomerNewParamsBillingConfigBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type V1CustomerNewParamsBillingConfigAwsRegion string

const (
	V1CustomerNewParamsBillingConfigAwsRegionAfSouth1     V1CustomerNewParamsBillingConfigAwsRegion = "af-south-1"
	V1CustomerNewParamsBillingConfigAwsRegionApEast1      V1CustomerNewParamsBillingConfigAwsRegion = "ap-east-1"
	V1CustomerNewParamsBillingConfigAwsRegionApNortheast1 V1CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-1"
	V1CustomerNewParamsBillingConfigAwsRegionApNortheast2 V1CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-2"
	V1CustomerNewParamsBillingConfigAwsRegionApNortheast3 V1CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-3"
	V1CustomerNewParamsBillingConfigAwsRegionApSouth1     V1CustomerNewParamsBillingConfigAwsRegion = "ap-south-1"
	V1CustomerNewParamsBillingConfigAwsRegionApSoutheast1 V1CustomerNewParamsBillingConfigAwsRegion = "ap-southeast-1"
	V1CustomerNewParamsBillingConfigAwsRegionApSoutheast2 V1CustomerNewParamsBillingConfigAwsRegion = "ap-southeast-2"
	V1CustomerNewParamsBillingConfigAwsRegionCaCentral1   V1CustomerNewParamsBillingConfigAwsRegion = "ca-central-1"
	V1CustomerNewParamsBillingConfigAwsRegionCnNorth1     V1CustomerNewParamsBillingConfigAwsRegion = "cn-north-1"
	V1CustomerNewParamsBillingConfigAwsRegionCnNorthwest1 V1CustomerNewParamsBillingConfigAwsRegion = "cn-northwest-1"
	V1CustomerNewParamsBillingConfigAwsRegionEuCentral1   V1CustomerNewParamsBillingConfigAwsRegion = "eu-central-1"
	V1CustomerNewParamsBillingConfigAwsRegionEuNorth1     V1CustomerNewParamsBillingConfigAwsRegion = "eu-north-1"
	V1CustomerNewParamsBillingConfigAwsRegionEuSouth1     V1CustomerNewParamsBillingConfigAwsRegion = "eu-south-1"
	V1CustomerNewParamsBillingConfigAwsRegionEuWest1      V1CustomerNewParamsBillingConfigAwsRegion = "eu-west-1"
	V1CustomerNewParamsBillingConfigAwsRegionEuWest2      V1CustomerNewParamsBillingConfigAwsRegion = "eu-west-2"
	V1CustomerNewParamsBillingConfigAwsRegionEuWest3      V1CustomerNewParamsBillingConfigAwsRegion = "eu-west-3"
	V1CustomerNewParamsBillingConfigAwsRegionMeSouth1     V1CustomerNewParamsBillingConfigAwsRegion = "me-south-1"
	V1CustomerNewParamsBillingConfigAwsRegionSaEast1      V1CustomerNewParamsBillingConfigAwsRegion = "sa-east-1"
	V1CustomerNewParamsBillingConfigAwsRegionUsEast1      V1CustomerNewParamsBillingConfigAwsRegion = "us-east-1"
	V1CustomerNewParamsBillingConfigAwsRegionUsEast2      V1CustomerNewParamsBillingConfigAwsRegion = "us-east-2"
	V1CustomerNewParamsBillingConfigAwsRegionUsGovEast1   V1CustomerNewParamsBillingConfigAwsRegion = "us-gov-east-1"
	V1CustomerNewParamsBillingConfigAwsRegionUsGovWest1   V1CustomerNewParamsBillingConfigAwsRegion = "us-gov-west-1"
	V1CustomerNewParamsBillingConfigAwsRegionUsWest1      V1CustomerNewParamsBillingConfigAwsRegion = "us-west-1"
	V1CustomerNewParamsBillingConfigAwsRegionUsWest2      V1CustomerNewParamsBillingConfigAwsRegion = "us-west-2"
)

func (r V1CustomerNewParamsBillingConfigAwsRegion) IsKnown() bool {
	switch r {
	case V1CustomerNewParamsBillingConfigAwsRegionAfSouth1, V1CustomerNewParamsBillingConfigAwsRegionApEast1, V1CustomerNewParamsBillingConfigAwsRegionApNortheast1, V1CustomerNewParamsBillingConfigAwsRegionApNortheast2, V1CustomerNewParamsBillingConfigAwsRegionApNortheast3, V1CustomerNewParamsBillingConfigAwsRegionApSouth1, V1CustomerNewParamsBillingConfigAwsRegionApSoutheast1, V1CustomerNewParamsBillingConfigAwsRegionApSoutheast2, V1CustomerNewParamsBillingConfigAwsRegionCaCentral1, V1CustomerNewParamsBillingConfigAwsRegionCnNorth1, V1CustomerNewParamsBillingConfigAwsRegionCnNorthwest1, V1CustomerNewParamsBillingConfigAwsRegionEuCentral1, V1CustomerNewParamsBillingConfigAwsRegionEuNorth1, V1CustomerNewParamsBillingConfigAwsRegionEuSouth1, V1CustomerNewParamsBillingConfigAwsRegionEuWest1, V1CustomerNewParamsBillingConfigAwsRegionEuWest2, V1CustomerNewParamsBillingConfigAwsRegionEuWest3, V1CustomerNewParamsBillingConfigAwsRegionMeSouth1, V1CustomerNewParamsBillingConfigAwsRegionSaEast1, V1CustomerNewParamsBillingConfigAwsRegionUsEast1, V1CustomerNewParamsBillingConfigAwsRegionUsEast2, V1CustomerNewParamsBillingConfigAwsRegionUsGovEast1, V1CustomerNewParamsBillingConfigAwsRegionUsGovWest1, V1CustomerNewParamsBillingConfigAwsRegionUsWest1, V1CustomerNewParamsBillingConfigAwsRegionUsWest2:
		return true
	}
	return false
}

type V1CustomerNewParamsBillingConfigStripeCollectionMethod string

const (
	V1CustomerNewParamsBillingConfigStripeCollectionMethodChargeAutomatically V1CustomerNewParamsBillingConfigStripeCollectionMethod = "charge_automatically"
	V1CustomerNewParamsBillingConfigStripeCollectionMethodSendInvoice         V1CustomerNewParamsBillingConfigStripeCollectionMethod = "send_invoice"
)

func (r V1CustomerNewParamsBillingConfigStripeCollectionMethod) IsKnown() bool {
	switch r {
	case V1CustomerNewParamsBillingConfigStripeCollectionMethodChargeAutomatically, V1CustomerNewParamsBillingConfigStripeCollectionMethodSendInvoice:
		return true
	}
	return false
}

type V1CustomerNewParamsCustomerBillingProviderConfiguration struct {
	// The billing provider set for this configuration.
	BillingProvider param.Field[V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider] `json:"billing_provider,required"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider and delivery provider combination. Defaults to an empty
	// object, however, for most billing provider + delivery method combinations, it
	// will not be a valid configuration.
	Configuration param.Field[map[string]interface{}] `json:"configuration"`
	// The method to use for delivering invoices to this customer. If not provided, the
	// `delivery_method_id` must be provided.
	DeliveryMethod param.Field[V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod] `json:"delivery_method"`
	// ID of the delivery method to use for this customer. If not provided, the
	// `delivery_method` must be provided.
	DeliveryMethodID param.Field[string] `json:"delivery_method_id" format:"uuid"`
}

func (r V1CustomerNewParamsCustomerBillingProviderConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The billing provider set for this configuration.
type V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider string

const (
	V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderAwsMarketplace   V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider = "aws_marketplace"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderAzureMarketplace V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider = "azure_marketplace"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderGcpMarketplace   V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider = "gcp_marketplace"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderStripe           V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider = "stripe"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderNetsuite         V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider = "netsuite"
)

func (r V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProvider) IsKnown() bool {
	switch r {
	case V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderAwsMarketplace, V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderAzureMarketplace, V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderGcpMarketplace, V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderStripe, V1CustomerNewParamsCustomerBillingProviderConfigurationsBillingProviderNetsuite:
		return true
	}
	return false
}

// The method to use for delivering invoices to this customer. If not provided, the
// `delivery_method_id` must be provided.
type V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod string

const (
	V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodDirectToBillingProvider V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod = "direct_to_billing_provider"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodAwsSqs                  V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod = "aws_sqs"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodTackle                  V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod = "tackle"
	V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodAwsSns                  V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod = "aws_sns"
)

func (r V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethod) IsKnown() bool {
	switch r {
	case V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodDirectToBillingProvider, V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodAwsSqs, V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodTackle, V1CustomerNewParamsCustomerBillingProviderConfigurationsDeliveryMethodAwsSns:
		return true
	}
	return false
}

type V1CustomerGetParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
}

type V1CustomerListParams struct {
	// Filter the customer list by customer_id. Up to 100 ids can be provided.
	CustomerIDs param.Field[[]string] `query:"customer_ids"`
	// Filter the customer list by ingest_alias
	IngestAlias param.Field[string] `query:"ingest_alias"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Filter the customer list to only return archived customers. By default, only
	// active customers are returned.
	OnlyArchived param.Field[bool] `query:"only_archived"`
	// Filter the customer list by salesforce_account_id. Up to 100 ids can be
	// provided.
	SalesforceAccountIDs param.Field[[]string] `query:"salesforce_account_ids"`
}

// URLQuery serializes [V1CustomerListParams]'s query parameters as `url.Values`.
func (r V1CustomerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerArchiveParams struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V1CustomerArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerListBillableMetricsParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// If true, the list of returned metrics will include archived metrics
	IncludeArchived param.Field[bool] `query:"include_archived"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// If true, the list of metrics will be filtered to just ones that are on the
	// customer's current plan
	OnCurrentPlan param.Field[bool] `query:"on_current_plan"`
}

// URLQuery serializes [V1CustomerListBillableMetricsParams]'s query parameters as
// `url.Values`.
func (r V1CustomerListBillableMetricsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerListCostsParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `query:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingOn param.Field[time.Time] `query:"starting_on,required" format:"date-time"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [V1CustomerListCostsParams]'s query parameters as
// `url.Values`.
func (r V1CustomerListCostsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerPreviewEventsParams struct {
	CustomerID param.Field[string]                               `path:"customer_id,required" format:"uuid"`
	Events     param.Field[[]V1CustomerPreviewEventsParamsEvent] `json:"events,required"`
	// If set to "replace", the preview will be generated as if those were the only
	// events for the specified customer. If set to "merge", the events will be merged
	// with any existing events for the specified customer. Defaults to "replace".
	Mode param.Field[V1CustomerPreviewEventsParamsMode] `json:"mode"`
	// If set, all zero quantity line items will be filtered out of the response.
	SkipZeroQtyLineItems param.Field[bool] `json:"skip_zero_qty_line_items"`
}

func (r V1CustomerPreviewEventsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerPreviewEventsParamsEvent struct {
	EventType param.Field[string] `json:"event_type,required"`
	// This has no effect for preview events, but may be set for consistency with Event
	// objects. They will be processed even if they do not match the customer's ID or
	// ingest aliases.
	CustomerID param.Field[string]                 `json:"customer_id"`
	Properties param.Field[map[string]interface{}] `json:"properties"`
	// RFC 3339 formatted. If not provided, the current time will be used.
	Timestamp param.Field[string] `json:"timestamp"`
	// This has no effect for preview events, but may be set for consistency with Event
	// objects. Duplicate transaction_ids are NOT filtered out, even within the same
	// request.
	TransactionID param.Field[string] `json:"transaction_id"`
}

func (r V1CustomerPreviewEventsParamsEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set to "replace", the preview will be generated as if those were the only
// events for the specified customer. If set to "merge", the events will be merged
// with any existing events for the specified customer. Defaults to "replace".
type V1CustomerPreviewEventsParamsMode string

const (
	V1CustomerPreviewEventsParamsModeReplace V1CustomerPreviewEventsParamsMode = "replace"
	V1CustomerPreviewEventsParamsModeMerge   V1CustomerPreviewEventsParamsMode = "merge"
)

func (r V1CustomerPreviewEventsParamsMode) IsKnown() bool {
	switch r {
	case V1CustomerPreviewEventsParamsModeReplace, V1CustomerPreviewEventsParamsModeMerge:
		return true
	}
	return false
}

type V1CustomerGetBillingConfigurationsParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r V1CustomerGetBillingConfigurationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerSetBillingConfigurationsParams struct {
	Data param.Field[[]V1CustomerSetBillingConfigurationsParamsData] `json:"data,required"`
}

func (r V1CustomerSetBillingConfigurationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerSetBillingConfigurationsParamsData struct {
	// The billing provider set for this configuration.
	BillingProvider param.Field[V1CustomerSetBillingConfigurationsParamsDataBillingProvider] `json:"billing_provider,required"`
	CustomerID      param.Field[string]                                                      `json:"customer_id,required" format:"uuid"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider and delivery method combination. Defaults to an empty
	// object, however, for most billing provider + delivery method combinations, it
	// will not be a valid configuration. For AWS marketplace configurations, the
	// aws_is_subscription_product flag can be used to indicate a product with
	// usage-based pricing. More information can be found
	// [here](https://docs.metronome.com/invoice-customers/solutions/marketplaces/invoice-aws/#provision-aws-marketplace-customers-in-metronome).
	Configuration param.Field[map[string]interface{}] `json:"configuration"`
	// The method to use for delivering invoices to this customer. If not provided, the
	// `delivery_method_id` must be provided.
	DeliveryMethod param.Field[V1CustomerSetBillingConfigurationsParamsDataDeliveryMethod] `json:"delivery_method"`
	// ID of the delivery method to use for this customer. If not provided, the
	// `delivery_method` must be provided.
	DeliveryMethodID param.Field[string] `json:"delivery_method_id" format:"uuid"`
}

func (r V1CustomerSetBillingConfigurationsParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The billing provider set for this configuration.
type V1CustomerSetBillingConfigurationsParamsDataBillingProvider string

const (
	V1CustomerSetBillingConfigurationsParamsDataBillingProviderAwsMarketplace   V1CustomerSetBillingConfigurationsParamsDataBillingProvider = "aws_marketplace"
	V1CustomerSetBillingConfigurationsParamsDataBillingProviderStripe           V1CustomerSetBillingConfigurationsParamsDataBillingProvider = "stripe"
	V1CustomerSetBillingConfigurationsParamsDataBillingProviderNetsuite         V1CustomerSetBillingConfigurationsParamsDataBillingProvider = "netsuite"
	V1CustomerSetBillingConfigurationsParamsDataBillingProviderCustom           V1CustomerSetBillingConfigurationsParamsDataBillingProvider = "custom"
	V1CustomerSetBillingConfigurationsParamsDataBillingProviderAzureMarketplace V1CustomerSetBillingConfigurationsParamsDataBillingProvider = "azure_marketplace"
	V1CustomerSetBillingConfigurationsParamsDataBillingProviderQuickbooksOnline V1CustomerSetBillingConfigurationsParamsDataBillingProvider = "quickbooks_online"
	V1CustomerSetBillingConfigurationsParamsDataBillingProviderWorkday          V1CustomerSetBillingConfigurationsParamsDataBillingProvider = "workday"
	V1CustomerSetBillingConfigurationsParamsDataBillingProviderGcpMarketplace   V1CustomerSetBillingConfigurationsParamsDataBillingProvider = "gcp_marketplace"
)

func (r V1CustomerSetBillingConfigurationsParamsDataBillingProvider) IsKnown() bool {
	switch r {
	case V1CustomerSetBillingConfigurationsParamsDataBillingProviderAwsMarketplace, V1CustomerSetBillingConfigurationsParamsDataBillingProviderStripe, V1CustomerSetBillingConfigurationsParamsDataBillingProviderNetsuite, V1CustomerSetBillingConfigurationsParamsDataBillingProviderCustom, V1CustomerSetBillingConfigurationsParamsDataBillingProviderAzureMarketplace, V1CustomerSetBillingConfigurationsParamsDataBillingProviderQuickbooksOnline, V1CustomerSetBillingConfigurationsParamsDataBillingProviderWorkday, V1CustomerSetBillingConfigurationsParamsDataBillingProviderGcpMarketplace:
		return true
	}
	return false
}

// The method to use for delivering invoices to this customer. If not provided, the
// `delivery_method_id` must be provided.
type V1CustomerSetBillingConfigurationsParamsDataDeliveryMethod string

const (
	V1CustomerSetBillingConfigurationsParamsDataDeliveryMethodDirectToBillingProvider V1CustomerSetBillingConfigurationsParamsDataDeliveryMethod = "direct_to_billing_provider"
	V1CustomerSetBillingConfigurationsParamsDataDeliveryMethodAwsSqs                  V1CustomerSetBillingConfigurationsParamsDataDeliveryMethod = "aws_sqs"
	V1CustomerSetBillingConfigurationsParamsDataDeliveryMethodTackle                  V1CustomerSetBillingConfigurationsParamsDataDeliveryMethod = "tackle"
	V1CustomerSetBillingConfigurationsParamsDataDeliveryMethodAwsSns                  V1CustomerSetBillingConfigurationsParamsDataDeliveryMethod = "aws_sns"
)

func (r V1CustomerSetBillingConfigurationsParamsDataDeliveryMethod) IsKnown() bool {
	switch r {
	case V1CustomerSetBillingConfigurationsParamsDataDeliveryMethodDirectToBillingProvider, V1CustomerSetBillingConfigurationsParamsDataDeliveryMethodAwsSqs, V1CustomerSetBillingConfigurationsParamsDataDeliveryMethodTackle, V1CustomerSetBillingConfigurationsParamsDataDeliveryMethodAwsSns:
		return true
	}
	return false
}

type V1CustomerSetIngestAliasesParams struct {
	CustomerID    param.Field[string]   `path:"customer_id,required" format:"uuid"`
	IngestAliases param.Field[[]string] `json:"ingest_aliases,required"`
}

func (r V1CustomerSetIngestAliasesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerSetNameParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// The new name for the customer. This will be truncated to 160 characters if the
	// provided name is longer.
	Name param.Field[string] `json:"name,required"`
}

func (r V1CustomerSetNameParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerUpdateConfigParams struct {
	CustomerID param.Field[string] `path:"customer_id,required" format:"uuid"`
	// Leave in draft or set to auto-advance on invoices sent to Stripe. Falls back to
	// the client-level config if unset, which defaults to true if unset.
	LeaveStripeInvoicesInDraft param.Field[bool] `json:"leave_stripe_invoices_in_draft"`
	// The Salesforce account ID for the customer
	SalesforceAccountID param.Field[string] `json:"salesforce_account_id"`
}

func (r V1CustomerUpdateConfigParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
