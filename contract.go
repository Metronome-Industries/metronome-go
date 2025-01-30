// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
	"github.com/tidwall/gjson"
)

// ContractService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContractService] method instead.
type ContractService struct {
	Options        []option.RequestOption
	Products       *ContractProductService
	RateCards      *ContractRateCardService
	NamedSchedules *ContractNamedScheduleService
}

// NewContractService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractService(opts ...option.RequestOption) (r *ContractService) {
	r = &ContractService{}
	r.Options = opts
	r.Products = NewContractProductService(opts...)
	r.RateCards = NewContractRateCardService(opts...)
	r.NamedSchedules = NewContractNamedScheduleService(opts...)
	return
}

// Create a new contract
func (r *ContractService) New(ctx context.Context, body ContractNewParams, opts ...option.RequestOption) (res *ContractNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific contract
func (r *ContractService) Get(ctx context.Context, body ContractGetParams, opts ...option.RequestOption) (res *ContractGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all contracts for a customer
func (r *ContractService) List(ctx context.Context, body ContractListParams, opts ...option.RequestOption) (res *ContractListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Add a manual balance entry
func (r *ContractService) AddManualBalanceEntry(ctx context.Context, body ContractAddManualBalanceEntryParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "contracts/addManualBalanceLedgerEntry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Amend a contract
func (r *ContractService) Amend(ctx context.Context, body ContractAmendParams, opts ...option.RequestOption) (res *ContractAmendResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/amend"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Archive a contract
func (r *ContractService) Archive(ctx context.Context, body ContractArchiveParams, opts ...option.RequestOption) (res *ContractArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates historical usage invoices for a contract
func (r *ContractService) NewHistoricalInvoices(ctx context.Context, body ContractNewHistoricalInvoicesParams, opts ...option.RequestOption) (res *ContractNewHistoricalInvoicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/createHistoricalInvoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List balances (commits and credits).
func (r *ContractService) ListBalances(ctx context.Context, body ContractListBalancesParams, opts ...option.RequestOption) (res *ContractListBalancesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/customerBalances/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the rate schedule for the rate card on a given contract.
func (r *ContractService) GetRateSchedule(ctx context.Context, params ContractGetRateScheduleParams, opts ...option.RequestOption) (res *ContractGetRateScheduleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/getContractRateSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create a new scheduled invoice for Professional Services terms on a contract.
// This endpoint's availability is dependent on your client's configuration.
func (r *ContractService) ScheduleProServicesInvoice(ctx context.Context, body ContractScheduleProServicesInvoiceParams, opts ...option.RequestOption) (res *ContractScheduleProServicesInvoiceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/scheduleProServicesInvoice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Set usage filter for a contract
func (r *ContractService) SetUsageFilter(ctx context.Context, body ContractSetUsageFilterParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "contracts/setUsageFilter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Update the end date of a contract
func (r *ContractService) UpdateEndDate(ctx context.Context, body ContractUpdateEndDateParams, opts ...option.RequestOption) (res *ContractUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractNewResponse struct {
	Data shared.ID               `json:"data,required"`
	JSON contractNewResponseJSON `json:"-"`
}

// contractNewResponseJSON contains the JSON metadata for the struct
// [ContractNewResponse]
type contractNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractNewResponseJSON) RawJSON() string {
	return r.raw
}

type ContractGetResponse struct {
	Data ContractGetResponseData `json:"data,required"`
	JSON contractGetResponseJSON `json:"-"`
}

// contractGetResponseJSON contains the JSON metadata for the struct
// [ContractGetResponse]
type contractGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseJSON) RawJSON() string {
	return r.raw
}

type ContractGetResponseData struct {
	ID         string                             `json:"id,required" format:"uuid"`
	Amendments []ContractGetResponseDataAmendment `json:"amendments,required"`
	Current    shared.ContractWithoutAmendments   `json:"current,required"`
	CustomerID string                             `json:"customer_id,required" format:"uuid"`
	Initial    shared.ContractWithoutAmendments   `json:"initial,required"`
	// RFC 3339 timestamp indicating when the contract was archived. If not returned,
	// the contract is not archived.
	ArchivedAt   time.Time         `json:"archived_at" format:"date-time"`
	CustomFields map[string]string `json:"custom_fields"`
	// The billing provider configuration associated with a contract.
	CustomerBillingProviderConfiguration ContractGetResponseDataCustomerBillingProviderConfiguration `json:"customer_billing_provider_configuration"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices ContractGetResponseDataScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                      `json:"uniqueness_key"`
	JSON          contractGetResponseDataJSON `json:"-"`
}

// contractGetResponseDataJSON contains the JSON metadata for the struct
// [ContractGetResponseData]
type contractGetResponseDataJSON struct {
	ID                                   apijson.Field
	Amendments                           apijson.Field
	Current                              apijson.Field
	CustomerID                           apijson.Field
	Initial                              apijson.Field
	ArchivedAt                           apijson.Field
	CustomFields                         apijson.Field
	CustomerBillingProviderConfiguration apijson.Field
	ScheduledChargesOnUsageInvoices      apijson.Field
	UniquenessKey                        apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *ContractGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ContractGetResponseDataAmendment struct {
	ID               string                   `json:"id,required" format:"uuid"`
	Commits          []shared.Commit          `json:"commits,required"`
	CreatedAt        time.Time                `json:"created_at,required" format:"date-time"`
	CreatedBy        string                   `json:"created_by,required"`
	Overrides        []shared.Override        `json:"overrides,required"`
	ScheduledCharges []shared.ScheduledCharge `json:"scheduled_charges,required"`
	StartingAt       time.Time                `json:"starting_at,required" format:"date-time"`
	Credits          []shared.Credit          `json:"credits"`
	// This field's availability is dependent on your client's configuration.
	Discounts []shared.Discount `json:"discounts"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []shared.ProService `json:"professional_services"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []ContractGetResponseDataAmendmentsResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                               `json:"salesforce_opportunity_id"`
	JSON                    contractGetResponseDataAmendmentJSON `json:"-"`
}

// contractGetResponseDataAmendmentJSON contains the JSON metadata for the struct
// [ContractGetResponseDataAmendment]
type contractGetResponseDataAmendmentJSON struct {
	ID                      apijson.Field
	Commits                 apijson.Field
	CreatedAt               apijson.Field
	CreatedBy               apijson.Field
	Overrides               apijson.Field
	ScheduledCharges        apijson.Field
	StartingAt              apijson.Field
	Credits                 apijson.Field
	Discounts               apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	ProfessionalServices    apijson.Field
	ResellerRoyalties       apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentJSON) RawJSON() string {
	return r.raw
}

type ContractGetResponseDataAmendmentsResellerRoyalty struct {
	ResellerType          ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
	AwsAccountNumber      string                                                         `json:"aws_account_number"`
	AwsOfferID            string                                                         `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                         `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                      `json:"ending_before,nullable" format:"date-time"`
	Fraction              float64                                                        `json:"fraction"`
	GcpAccountID          string                                                         `json:"gcp_account_id"`
	GcpOfferID            string                                                         `json:"gcp_offer_id"`
	NetsuiteResellerID    string                                                         `json:"netsuite_reseller_id"`
	ResellerContractValue float64                                                        `json:"reseller_contract_value"`
	StartingAt            time.Time                                                      `json:"starting_at" format:"date-time"`
	JSON                  contractGetResponseDataAmendmentsResellerRoyaltyJSON           `json:"-"`
}

// contractGetResponseDataAmendmentsResellerRoyaltyJSON contains the JSON metadata
// for the struct [ContractGetResponseDataAmendmentsResellerRoyalty]
type contractGetResponseDataAmendmentsResellerRoyaltyJSON struct {
	ResellerType          apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	Fraction              apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerContractValue apijson.Field
	StartingAt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType string

const (
	ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeAws           ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS"
	ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeAwsProService ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp           ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP"
	ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeGcpProService ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeAws, ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeAwsProService, ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp, ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

// The billing provider configuration associated with a contract.
type ContractGetResponseDataCustomerBillingProviderConfiguration struct {
	BillingProvider ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider `json:"billing_provider,required"`
	DeliveryMethod  ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod  `json:"delivery_method,required"`
	JSON            contractGetResponseDataCustomerBillingProviderConfigurationJSON            `json:"-"`
}

// contractGetResponseDataCustomerBillingProviderConfigurationJSON contains the
// JSON metadata for the struct
// [ContractGetResponseDataCustomerBillingProviderConfiguration]
type contractGetResponseDataCustomerBillingProviderConfigurationJSON struct {
	BillingProvider apijson.Field
	DeliveryMethod  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ContractGetResponseDataCustomerBillingProviderConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataCustomerBillingProviderConfigurationJSON) RawJSON() string {
	return r.raw
}

type ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider string

const (
	ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace   ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "aws_marketplace"
	ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderStripe           ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "stripe"
	ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite         ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "netsuite"
	ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderCustom           ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "custom"
	ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "azure_marketplace"
	ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "quickbooks_online"
	ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday          ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "workday"
	ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace   ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "gcp_marketplace"
)

func (r ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider) IsKnown() bool {
	switch r {
	case ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace, ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderStripe, ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite, ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderCustom, ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace, ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline, ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday, ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace:
		return true
	}
	return false
}

type ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod string

const (
	ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "direct_to_billing_provider"
	ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs                  ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sqs"
	ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle                  ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "tackle"
	ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns                  ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sns"
)

func (r ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod) IsKnown() bool {
	switch r {
	case ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider, ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs, ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle, ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns:
		return true
	}
	return false
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type ContractGetResponseDataScheduledChargesOnUsageInvoices string

const (
	ContractGetResponseDataScheduledChargesOnUsageInvoicesAll ContractGetResponseDataScheduledChargesOnUsageInvoices = "ALL"
)

func (r ContractGetResponseDataScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case ContractGetResponseDataScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
}

type ContractListResponse struct {
	Data []ContractListResponseData `json:"data,required"`
	JSON contractListResponseJSON   `json:"-"`
}

// contractListResponseJSON contains the JSON metadata for the struct
// [ContractListResponse]
type contractListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseJSON) RawJSON() string {
	return r.raw
}

type ContractListResponseData struct {
	ID         string                              `json:"id,required" format:"uuid"`
	Amendments []ContractListResponseDataAmendment `json:"amendments,required"`
	Current    shared.ContractWithoutAmendments    `json:"current,required"`
	CustomerID string                              `json:"customer_id,required" format:"uuid"`
	Initial    shared.ContractWithoutAmendments    `json:"initial,required"`
	// RFC 3339 timestamp indicating when the contract was archived. If not returned,
	// the contract is not archived.
	ArchivedAt   time.Time         `json:"archived_at" format:"date-time"`
	CustomFields map[string]string `json:"custom_fields"`
	// The billing provider configuration associated with a contract.
	CustomerBillingProviderConfiguration ContractListResponseDataCustomerBillingProviderConfiguration `json:"customer_billing_provider_configuration"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices ContractListResponseDataScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                       `json:"uniqueness_key"`
	JSON          contractListResponseDataJSON `json:"-"`
}

// contractListResponseDataJSON contains the JSON metadata for the struct
// [ContractListResponseData]
type contractListResponseDataJSON struct {
	ID                                   apijson.Field
	Amendments                           apijson.Field
	Current                              apijson.Field
	CustomerID                           apijson.Field
	Initial                              apijson.Field
	ArchivedAt                           apijson.Field
	CustomFields                         apijson.Field
	CustomerBillingProviderConfiguration apijson.Field
	ScheduledChargesOnUsageInvoices      apijson.Field
	UniquenessKey                        apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *ContractListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataJSON) RawJSON() string {
	return r.raw
}

type ContractListResponseDataAmendment struct {
	ID               string                   `json:"id,required" format:"uuid"`
	Commits          []shared.Commit          `json:"commits,required"`
	CreatedAt        time.Time                `json:"created_at,required" format:"date-time"`
	CreatedBy        string                   `json:"created_by,required"`
	Overrides        []shared.Override        `json:"overrides,required"`
	ScheduledCharges []shared.ScheduledCharge `json:"scheduled_charges,required"`
	StartingAt       time.Time                `json:"starting_at,required" format:"date-time"`
	Credits          []shared.Credit          `json:"credits"`
	// This field's availability is dependent on your client's configuration.
	Discounts []shared.Discount `json:"discounts"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []shared.ProService `json:"professional_services"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []ContractListResponseDataAmendmentsResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                                `json:"salesforce_opportunity_id"`
	JSON                    contractListResponseDataAmendmentJSON `json:"-"`
}

// contractListResponseDataAmendmentJSON contains the JSON metadata for the struct
// [ContractListResponseDataAmendment]
type contractListResponseDataAmendmentJSON struct {
	ID                      apijson.Field
	Commits                 apijson.Field
	CreatedAt               apijson.Field
	CreatedBy               apijson.Field
	Overrides               apijson.Field
	ScheduledCharges        apijson.Field
	StartingAt              apijson.Field
	Credits                 apijson.Field
	Discounts               apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	ProfessionalServices    apijson.Field
	ResellerRoyalties       apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractListResponseDataAmendment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentJSON) RawJSON() string {
	return r.raw
}

type ContractListResponseDataAmendmentsResellerRoyalty struct {
	ResellerType          ContractListResponseDataAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
	AwsAccountNumber      string                                                          `json:"aws_account_number"`
	AwsOfferID            string                                                          `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                          `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                       `json:"ending_before,nullable" format:"date-time"`
	Fraction              float64                                                         `json:"fraction"`
	GcpAccountID          string                                                          `json:"gcp_account_id"`
	GcpOfferID            string                                                          `json:"gcp_offer_id"`
	NetsuiteResellerID    string                                                          `json:"netsuite_reseller_id"`
	ResellerContractValue float64                                                         `json:"reseller_contract_value"`
	StartingAt            time.Time                                                       `json:"starting_at" format:"date-time"`
	JSON                  contractListResponseDataAmendmentsResellerRoyaltyJSON           `json:"-"`
}

// contractListResponseDataAmendmentsResellerRoyaltyJSON contains the JSON metadata
// for the struct [ContractListResponseDataAmendmentsResellerRoyalty]
type contractListResponseDataAmendmentsResellerRoyaltyJSON struct {
	ResellerType          apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	Fraction              apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerContractValue apijson.Field
	StartingAt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type ContractListResponseDataAmendmentsResellerRoyaltiesResellerType string

const (
	ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeAws           ContractListResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS"
	ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeAwsProService ContractListResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp           ContractListResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP"
	ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeGcpProService ContractListResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r ContractListResponseDataAmendmentsResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeAws, ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeAwsProService, ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp, ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

// The billing provider configuration associated with a contract.
type ContractListResponseDataCustomerBillingProviderConfiguration struct {
	BillingProvider ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider `json:"billing_provider,required"`
	DeliveryMethod  ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod  `json:"delivery_method,required"`
	JSON            contractListResponseDataCustomerBillingProviderConfigurationJSON            `json:"-"`
}

// contractListResponseDataCustomerBillingProviderConfigurationJSON contains the
// JSON metadata for the struct
// [ContractListResponseDataCustomerBillingProviderConfiguration]
type contractListResponseDataCustomerBillingProviderConfigurationJSON struct {
	BillingProvider apijson.Field
	DeliveryMethod  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ContractListResponseDataCustomerBillingProviderConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataCustomerBillingProviderConfigurationJSON) RawJSON() string {
	return r.raw
}

type ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider string

const (
	ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace   ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "aws_marketplace"
	ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderStripe           ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "stripe"
	ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite         ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "netsuite"
	ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderCustom           ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "custom"
	ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "azure_marketplace"
	ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "quickbooks_online"
	ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday          ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "workday"
	ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace   ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "gcp_marketplace"
)

func (r ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider) IsKnown() bool {
	switch r {
	case ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace, ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderStripe, ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite, ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderCustom, ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace, ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline, ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday, ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace:
		return true
	}
	return false
}

type ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod string

const (
	ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "direct_to_billing_provider"
	ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs                  ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sqs"
	ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle                  ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "tackle"
	ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns                  ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sns"
)

func (r ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod) IsKnown() bool {
	switch r {
	case ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider, ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs, ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle, ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns:
		return true
	}
	return false
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type ContractListResponseDataScheduledChargesOnUsageInvoices string

const (
	ContractListResponseDataScheduledChargesOnUsageInvoicesAll ContractListResponseDataScheduledChargesOnUsageInvoices = "ALL"
)

func (r ContractListResponseDataScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case ContractListResponseDataScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
}

type ContractAmendResponse struct {
	Data shared.ID                 `json:"data,required"`
	JSON contractAmendResponseJSON `json:"-"`
}

// contractAmendResponseJSON contains the JSON metadata for the struct
// [ContractAmendResponse]
type contractAmendResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractAmendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractAmendResponseJSON) RawJSON() string {
	return r.raw
}

type ContractArchiveResponse struct {
	Data shared.ID                   `json:"data,required"`
	JSON contractArchiveResponseJSON `json:"-"`
}

// contractArchiveResponseJSON contains the JSON metadata for the struct
// [ContractArchiveResponse]
type contractArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type ContractNewHistoricalInvoicesResponse struct {
	Data []Invoice                                 `json:"data,required"`
	JSON contractNewHistoricalInvoicesResponseJSON `json:"-"`
}

// contractNewHistoricalInvoicesResponseJSON contains the JSON metadata for the
// struct [ContractNewHistoricalInvoicesResponse]
type contractNewHistoricalInvoicesResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractNewHistoricalInvoicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractNewHistoricalInvoicesResponseJSON) RawJSON() string {
	return r.raw
}

type ContractListBalancesResponse struct {
	Data     []ContractListBalancesResponseData `json:"data,required"`
	NextPage string                             `json:"next_page,required,nullable"`
	JSON     contractListBalancesResponseJSON   `json:"-"`
}

// contractListBalancesResponseJSON contains the JSON metadata for the struct
// [ContractListBalancesResponse]
type contractListBalancesResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListBalancesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseJSON) RawJSON() string {
	return r.raw
}

type ContractListBalancesResponseData struct {
	ID string `json:"id,required" format:"uuid"`
	// This field can have the runtime type of [shared.CommitProduct],
	// [shared.CreditProduct].
	Product interface{}                          `json:"product,required"`
	Type    ContractListBalancesResponseDataType `json:"type,required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule shared.ScheduleDuration `json:"access_schedule"`
	// (DEPRECATED) Use access_schedule + invoice_schedule instead.
	Amount float64 `json:"amount"`
	// This field can have the runtime type of [[]string].
	ApplicableContractIDs interface{} `json:"applicable_contract_ids"`
	// This field can have the runtime type of [[]string].
	ApplicableProductIDs interface{} `json:"applicable_product_ids"`
	// This field can have the runtime type of [[]string].
	ApplicableProductTags interface{} `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance float64 `json:"balance"`
	// This field can have the runtime type of [shared.CommitContract],
	// [shared.CreditContract].
	Contract interface{} `json:"contract"`
	// This field can have the runtime type of [map[string]string].
	CustomFields interface{} `json:"custom_fields"`
	Description  string      `json:"description"`
	// This field can have the runtime type of [shared.CommitInvoiceContract].
	InvoiceContract interface{} `json:"invoice_contract"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule shared.SchedulePointInTime `json:"invoice_schedule"`
	// This field can have the runtime type of [[]shared.CommitLedger],
	// [[]shared.CreditLedger].
	Ledger interface{} `json:"ledger"`
	Name   string      `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64                                  `json:"priority"`
	RateType ContractListBalancesResponseDataRateType `json:"rate_type"`
	// This field can have the runtime type of [shared.CommitRolledOverFrom].
	RolledOverFrom   interface{} `json:"rolled_over_from"`
	RolloverFraction float64     `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey string                               `json:"uniqueness_key"`
	JSON          contractListBalancesResponseDataJSON `json:"-"`
	union         ContractListBalancesResponseDataUnion
}

// contractListBalancesResponseDataJSON contains the JSON metadata for the struct
// [ContractListBalancesResponseData]
type contractListBalancesResponseDataJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	Amount                  apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Balance                 apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	InvoiceContract         apijson.Field
	InvoiceSchedule         apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	RateType                apijson.Field
	RolledOverFrom          apijson.Field
	RolloverFraction        apijson.Field
	SalesforceOpportunityID apijson.Field
	UniquenessKey           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r contractListBalancesResponseDataJSON) RawJSON() string {
	return r.raw
}

func (r *ContractListBalancesResponseData) UnmarshalJSON(data []byte) (err error) {
	*r = ContractListBalancesResponseData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContractListBalancesResponseDataUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.Commit], [shared.Credit].
func (r ContractListBalancesResponseData) AsUnion() ContractListBalancesResponseDataUnion {
	return r.union
}

// Union satisfied by [shared.Commit] or [shared.Credit].
type ContractListBalancesResponseDataUnion interface {
	ImplementsContractListBalancesResponseData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContractListBalancesResponseDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Commit{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Credit{}),
		},
	)
}

type ContractListBalancesResponseDataType string

const (
	ContractListBalancesResponseDataTypePrepaid  ContractListBalancesResponseDataType = "PREPAID"
	ContractListBalancesResponseDataTypePostpaid ContractListBalancesResponseDataType = "POSTPAID"
	ContractListBalancesResponseDataTypeCredit   ContractListBalancesResponseDataType = "CREDIT"
)

func (r ContractListBalancesResponseDataType) IsKnown() bool {
	switch r {
	case ContractListBalancesResponseDataTypePrepaid, ContractListBalancesResponseDataTypePostpaid, ContractListBalancesResponseDataTypeCredit:
		return true
	}
	return false
}

type ContractListBalancesResponseDataRateType string

const (
	ContractListBalancesResponseDataRateTypeCommitRate ContractListBalancesResponseDataRateType = "COMMIT_RATE"
	ContractListBalancesResponseDataRateTypeListRate   ContractListBalancesResponseDataRateType = "LIST_RATE"
)

func (r ContractListBalancesResponseDataRateType) IsKnown() bool {
	switch r {
	case ContractListBalancesResponseDataRateTypeCommitRate, ContractListBalancesResponseDataRateTypeListRate:
		return true
	}
	return false
}

type ContractGetRateScheduleResponse struct {
	Data     []ContractGetRateScheduleResponseData `json:"data,required"`
	NextPage string                                `json:"next_page,nullable"`
	JSON     contractGetRateScheduleResponseJSON   `json:"-"`
}

// contractGetRateScheduleResponseJSON contains the JSON metadata for the struct
// [ContractGetRateScheduleResponse]
type contractGetRateScheduleResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetRateScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetRateScheduleResponseJSON) RawJSON() string {
	return r.raw
}

type ContractGetRateScheduleResponseData struct {
	Entitled            bool              `json:"entitled,required"`
	ListRate            shared.Rate       `json:"list_rate,required"`
	ProductCustomFields map[string]string `json:"product_custom_fields,required"`
	ProductID           string            `json:"product_id,required" format:"uuid"`
	ProductName         string            `json:"product_name,required"`
	ProductTags         []string          `json:"product_tags,required"`
	RateCardID          string            `json:"rate_card_id,required" format:"uuid"`
	StartingAt          time.Time         `json:"starting_at,required" format:"date-time"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate         ContractGetRateScheduleResponseDataCommitRate `json:"commit_rate"`
	EndingBefore       time.Time                                     `json:"ending_before" format:"date-time"`
	OverrideRate       shared.Rate                                   `json:"override_rate"`
	PricingGroupValues map[string]string                             `json:"pricing_group_values"`
	JSON               contractGetRateScheduleResponseDataJSON       `json:"-"`
}

// contractGetRateScheduleResponseDataJSON contains the JSON metadata for the
// struct [ContractGetRateScheduleResponseData]
type contractGetRateScheduleResponseDataJSON struct {
	Entitled            apijson.Field
	ListRate            apijson.Field
	ProductCustomFields apijson.Field
	ProductID           apijson.Field
	ProductName         apijson.Field
	ProductTags         apijson.Field
	RateCardID          apijson.Field
	StartingAt          apijson.Field
	CommitRate          apijson.Field
	EndingBefore        apijson.Field
	OverrideRate        apijson.Field
	PricingGroupValues  apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ContractGetRateScheduleResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetRateScheduleResponseDataJSON) RawJSON() string {
	return r.raw
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type ContractGetRateScheduleResponseDataCommitRate struct {
	RateType ContractGetRateScheduleResponseDataCommitRateRateType `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price float64 `json:"price"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier                                     `json:"tiers"`
	JSON  contractGetRateScheduleResponseDataCommitRateJSON `json:"-"`
}

// contractGetRateScheduleResponseDataCommitRateJSON contains the JSON metadata for
// the struct [ContractGetRateScheduleResponseDataCommitRate]
type contractGetRateScheduleResponseDataCommitRateJSON struct {
	RateType    apijson.Field
	Price       apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetRateScheduleResponseDataCommitRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetRateScheduleResponseDataCommitRateJSON) RawJSON() string {
	return r.raw
}

type ContractGetRateScheduleResponseDataCommitRateRateType string

const (
	ContractGetRateScheduleResponseDataCommitRateRateTypeFlat         ContractGetRateScheduleResponseDataCommitRateRateType = "FLAT"
	ContractGetRateScheduleResponseDataCommitRateRateTypePercentage   ContractGetRateScheduleResponseDataCommitRateRateType = "PERCENTAGE"
	ContractGetRateScheduleResponseDataCommitRateRateTypeSubscription ContractGetRateScheduleResponseDataCommitRateRateType = "SUBSCRIPTION"
	ContractGetRateScheduleResponseDataCommitRateRateTypeTiered       ContractGetRateScheduleResponseDataCommitRateRateType = "TIERED"
	ContractGetRateScheduleResponseDataCommitRateRateTypeCustom       ContractGetRateScheduleResponseDataCommitRateRateType = "CUSTOM"
)

func (r ContractGetRateScheduleResponseDataCommitRateRateType) IsKnown() bool {
	switch r {
	case ContractGetRateScheduleResponseDataCommitRateRateTypeFlat, ContractGetRateScheduleResponseDataCommitRateRateTypePercentage, ContractGetRateScheduleResponseDataCommitRateRateTypeSubscription, ContractGetRateScheduleResponseDataCommitRateRateTypeTiered, ContractGetRateScheduleResponseDataCommitRateRateTypeCustom:
		return true
	}
	return false
}

type ContractScheduleProServicesInvoiceResponse struct {
	Data []Invoice                                      `json:"data,required"`
	JSON contractScheduleProServicesInvoiceResponseJSON `json:"-"`
}

// contractScheduleProServicesInvoiceResponseJSON contains the JSON metadata for
// the struct [ContractScheduleProServicesInvoiceResponse]
type contractScheduleProServicesInvoiceResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractScheduleProServicesInvoiceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractScheduleProServicesInvoiceResponseJSON) RawJSON() string {
	return r.raw
}

type ContractUpdateEndDateResponse struct {
	Data shared.ID                         `json:"data,required"`
	JSON contractUpdateEndDateResponseJSON `json:"-"`
}

// contractUpdateEndDateResponseJSON contains the JSON metadata for the struct
// [ContractUpdateEndDateResponse]
type contractUpdateEndDateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractUpdateEndDateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractUpdateEndDateResponseJSON) RawJSON() string {
	return r.raw
}

type ContractNewParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// inclusive contract start time
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// The billing provider configuration associated with a contract.
	BillingProviderConfiguration param.Field[ContractNewParamsBillingProviderConfiguration] `json:"billing_provider_configuration"`
	Commits                      param.Field[[]ContractNewParamsCommit]                     `json:"commits"`
	Credits                      param.Field[[]ContractNewParamsCredit]                     `json:"credits"`
	CustomFields                 param.Field[map[string]string]                             `json:"custom_fields"`
	// This field's availability is dependent on your client's configuration.
	Discounts param.Field[[]ContractNewParamsDiscount] `json:"discounts"`
	// exclusive contract end time
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
	// prices automatically. EXPLICIT prioritization requires specifying priorities for
	// each multiplier; the one with the lowest priority value will be prioritized
	// first. If tiered overrides are used, prioritization must be explicit.
	MultiplierOverridePrioritization param.Field[ContractNewParamsMultiplierOverridePrioritization] `json:"multiplier_override_prioritization"`
	Name                             param.Field[string]                                            `json:"name"`
	NetPaymentTermsDays              param.Field[float64]                                           `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string]                      `json:"netsuite_sales_order_id"`
	Overrides            param.Field[[]ContractNewParamsOverride] `json:"overrides"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices param.Field[[]ContractNewParamsProfessionalService] `json:"professional_services"`
	// Selects the rate card linked to the specified alias as of the contract's start
	// date.
	RateCardAlias    param.Field[string]                             `json:"rate_card_alias"`
	RateCardID       param.Field[string]                             `json:"rate_card_id" format:"uuid"`
	RecurringCommits param.Field[[]ContractNewParamsRecurringCommit] `json:"recurring_commits"`
	RecurringCredits param.Field[[]ContractNewParamsRecurringCredit] `json:"recurring_credits"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties param.Field[[]ContractNewParamsResellerRoyalty] `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string]                             `json:"salesforce_opportunity_id"`
	ScheduledCharges        param.Field[[]ContractNewParamsScheduledCharge] `json:"scheduled_charges"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices param.Field[ContractNewParamsScheduledChargesOnUsageInvoices] `json:"scheduled_charges_on_usage_invoices"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue param.Field[float64]                     `json:"total_contract_value"`
	Transition         param.Field[ContractNewParamsTransition] `json:"transition"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey          param.Field[string]                                  `json:"uniqueness_key"`
	UsageFilter            param.Field[shared.BaseUsageFilterParam]             `json:"usage_filter"`
	UsageStatementSchedule param.Field[ContractNewParamsUsageStatementSchedule] `json:"usage_statement_schedule"`
}

func (r ContractNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The billing provider configuration associated with a contract.
type ContractNewParamsBillingProviderConfiguration struct {
	BillingProvider param.Field[ContractNewParamsBillingProviderConfigurationBillingProvider] `json:"billing_provider"`
	// The Metronome ID of the billing provider configuration
	BillingProviderConfigurationID param.Field[string]                                                      `json:"billing_provider_configuration_id" format:"uuid"`
	DeliveryMethod                 param.Field[ContractNewParamsBillingProviderConfigurationDeliveryMethod] `json:"delivery_method"`
}

func (r ContractNewParamsBillingProviderConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsBillingProviderConfigurationBillingProvider string

const (
	ContractNewParamsBillingProviderConfigurationBillingProviderAwsMarketplace   ContractNewParamsBillingProviderConfigurationBillingProvider = "aws_marketplace"
	ContractNewParamsBillingProviderConfigurationBillingProviderAzureMarketplace ContractNewParamsBillingProviderConfigurationBillingProvider = "azure_marketplace"
	ContractNewParamsBillingProviderConfigurationBillingProviderGcpMarketplace   ContractNewParamsBillingProviderConfigurationBillingProvider = "gcp_marketplace"
	ContractNewParamsBillingProviderConfigurationBillingProviderStripe           ContractNewParamsBillingProviderConfigurationBillingProvider = "stripe"
	ContractNewParamsBillingProviderConfigurationBillingProviderNetsuite         ContractNewParamsBillingProviderConfigurationBillingProvider = "netsuite"
)

func (r ContractNewParamsBillingProviderConfigurationBillingProvider) IsKnown() bool {
	switch r {
	case ContractNewParamsBillingProviderConfigurationBillingProviderAwsMarketplace, ContractNewParamsBillingProviderConfigurationBillingProviderAzureMarketplace, ContractNewParamsBillingProviderConfigurationBillingProviderGcpMarketplace, ContractNewParamsBillingProviderConfigurationBillingProviderStripe, ContractNewParamsBillingProviderConfigurationBillingProviderNetsuite:
		return true
	}
	return false
}

type ContractNewParamsBillingProviderConfigurationDeliveryMethod string

const (
	ContractNewParamsBillingProviderConfigurationDeliveryMethodDirectToBillingProvider ContractNewParamsBillingProviderConfigurationDeliveryMethod = "direct_to_billing_provider"
	ContractNewParamsBillingProviderConfigurationDeliveryMethodAwsSqs                  ContractNewParamsBillingProviderConfigurationDeliveryMethod = "aws_sqs"
	ContractNewParamsBillingProviderConfigurationDeliveryMethodTackle                  ContractNewParamsBillingProviderConfigurationDeliveryMethod = "tackle"
	ContractNewParamsBillingProviderConfigurationDeliveryMethodAwsSns                  ContractNewParamsBillingProviderConfigurationDeliveryMethod = "aws_sns"
)

func (r ContractNewParamsBillingProviderConfigurationDeliveryMethod) IsKnown() bool {
	switch r {
	case ContractNewParamsBillingProviderConfigurationDeliveryMethodDirectToBillingProvider, ContractNewParamsBillingProviderConfigurationDeliveryMethodAwsSqs, ContractNewParamsBillingProviderConfigurationDeliveryMethodTackle, ContractNewParamsBillingProviderConfigurationDeliveryMethodAwsSns:
		return true
	}
	return false
}

type ContractNewParamsCommit struct {
	ProductID param.Field[string]                       `json:"product_id,required" format:"uuid"`
	Type      param.Field[ContractNewParamsCommitsType] `json:"type,required"`
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
	// commits only one schedule item is allowed and amount must match invoice_schedule
	// total.
	AccessSchedule param.Field[ContractNewParamsCommitsAccessSchedule] `json:"access_schedule"`
	// (DEPRECATED) Use access_schedule and invoice_schedule instead.
	Amount param.Field[float64] `json:"amount"`
	// Which products the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags param.Field[[]string]          `json:"applicable_product_tags"`
	CustomFields          param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this
	// time and only one schedule item is allowed; the total must match access_schedule
	// amount. Optional for "PREPAID" commits: if not provided, this will be a
	// "complimentary" commit with no invoice.
	InvoiceSchedule param.Field[ContractNewParamsCommitsInvoiceSchedule] `json:"invoice_schedule"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Field[float64]                          `json:"priority"`
	RateType param.Field[ContractNewParamsCommitsRateType] `json:"rate_type"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// A temporary ID for the commit that can be used to reference the commit for
	// commit specific overrides.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r ContractNewParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsCommitsType string

const (
	ContractNewParamsCommitsTypePrepaid  ContractNewParamsCommitsType = "PREPAID"
	ContractNewParamsCommitsTypePostpaid ContractNewParamsCommitsType = "POSTPAID"
)

func (r ContractNewParamsCommitsType) IsKnown() bool {
	switch r {
	case ContractNewParamsCommitsTypePrepaid, ContractNewParamsCommitsTypePostpaid:
		return true
	}
	return false
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
type ContractNewParamsCommitsAccessSchedule struct {
	ScheduleItems param.Field[[]ContractNewParamsCommitsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r ContractNewParamsCommitsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsCommitsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r ContractNewParamsCommitsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match access_schedule
// amount. Optional for "PREPAID" commits: if not provided, this will be a
// "complimentary" commit with no invoice.
type ContractNewParamsCommitsInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[ContractNewParamsCommitsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]ContractNewParamsCommitsInvoiceScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractNewParamsCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type ContractNewParamsCommitsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                         `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractNewParamsCommitsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type ContractNewParamsCommitsInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractNewParamsCommitsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsCommitsRateType string

const (
	ContractNewParamsCommitsRateTypeCommitRate ContractNewParamsCommitsRateType = "COMMIT_RATE"
	ContractNewParamsCommitsRateTypeListRate   ContractNewParamsCommitsRateType = "LIST_RATE"
)

func (r ContractNewParamsCommitsRateType) IsKnown() bool {
	switch r {
	case ContractNewParamsCommitsRateTypeCommitRate, ContractNewParamsCommitsRateTypeListRate:
		return true
	}
	return false
}

type ContractNewParamsCredit struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule param.Field[ContractNewParamsCreditsAccessSchedule] `json:"access_schedule,required"`
	ProductID      param.Field[string]                                 `json:"product_id,required" format:"uuid"`
	// Which products the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags param.Field[[]string]          `json:"applicable_product_tags"`
	CustomFields          param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Field[float64]                          `json:"priority"`
	RateType param.Field[ContractNewParamsCreditsRateType] `json:"rate_type"`
}

func (r ContractNewParamsCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the credit to the customer.
type ContractNewParamsCreditsAccessSchedule struct {
	ScheduleItems param.Field[[]ContractNewParamsCreditsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r ContractNewParamsCreditsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsCreditsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r ContractNewParamsCreditsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsCreditsRateType string

const (
	ContractNewParamsCreditsRateTypeCommitRate ContractNewParamsCreditsRateType = "COMMIT_RATE"
	ContractNewParamsCreditsRateTypeListRate   ContractNewParamsCreditsRateType = "LIST_RATE"
)

func (r ContractNewParamsCreditsRateType) IsKnown() bool {
	switch r {
	case ContractNewParamsCreditsRateTypeCommitRate, ContractNewParamsCreditsRateTypeListRate:
		return true
	}
	return false
}

type ContractNewParamsDiscount struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule     param.Field[ContractNewParamsDiscountsSchedule] `json:"schedule,required"`
	CustomFields param.Field[map[string]string]                  `json:"custom_fields"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractNewParamsDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule.
type ContractNewParamsDiscountsSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[ContractNewParamsDiscountsScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]ContractNewParamsDiscountsScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractNewParamsDiscountsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type ContractNewParamsDiscountsScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                    `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractNewParamsDiscountsScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractNewParamsDiscountsScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution string

const (
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided, ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded, ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractNewParamsDiscountsScheduleRecurringScheduleFrequency string

const (
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r ContractNewParamsDiscountsScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly, ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly, ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual, ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type ContractNewParamsDiscountsScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractNewParamsDiscountsScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
// prices automatically. EXPLICIT prioritization requires specifying priorities for
// each multiplier; the one with the lowest priority value will be prioritized
// first. If tiered overrides are used, prioritization must be explicit.
type ContractNewParamsMultiplierOverridePrioritization string

const (
	ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier ContractNewParamsMultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	ContractNewParamsMultiplierOverridePrioritizationExplicit         ContractNewParamsMultiplierOverridePrioritization = "EXPLICIT"
)

func (r ContractNewParamsMultiplierOverridePrioritization) IsKnown() bool {
	switch r {
	case ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier, ContractNewParamsMultiplierOverridePrioritizationExplicit:
		return true
	}
	return false
}

type ContractNewParamsOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// tags identifying products whose rates are being overridden. Cannot be used in
	// conjunction with override_specifiers.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Entitled     param.Field[bool]      `json:"entitled"`
	// Indicates whether the override should only apply to commits. Defaults to
	// `false`. If `true`, you can specify relevant commits in `override_specifiers` by
	// passing `commit_ids`. if you do not specify `commit_ids`, then the override will
	// apply when consuming any prepaid or postpaid commit.
	IsCommitSpecific param.Field[bool] `json:"is_commit_specific"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Field[float64] `json:"multiplier"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If
	// provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers param.Field[[]ContractNewParamsOverridesOverrideSpecifier] `json:"override_specifiers"`
	// Required for OVERWRITE type.
	OverwriteRate param.Field[ContractNewParamsOverridesOverwriteRate] `json:"overwrite_rate"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides.
	// Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered
	// and multiplier overrides are prioritized by their priority value (lowest first).
	// Must be > 0.
	Priority param.Field[float64] `json:"priority"`
	// ID of the product whose rate is being overridden. Cannot be used in conjunction
	// with override_specifiers.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// Indicates whether the override applies to commit rates or list rates. Can only
	// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
	// `"LIST_RATE"`.
	Target param.Field[ContractNewParamsOverridesTarget] `json:"target"`
	// Required for TIERED type. Must have at least one tier.
	Tiers param.Field[[]ContractNewParamsOverridesTier] `json:"tiers"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type param.Field[ContractNewParamsOverridesType] `json:"type"`
}

func (r ContractNewParamsOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsOverridesOverrideSpecifier struct {
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to the
	// specified commits. If not provided, the override will apply to all commits.
	CommitIDs param.Field[[]string] `json:"commit_ids"`
	// A map of group names to values. The override will only apply to line items with
	// the specified presentation group values. Can only be used for multiplier
	// overrides.
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	// A map of pricing group names to values. The override will only apply to products
	// with the specified pricing group values.
	PricingGroupValues param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the override will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the override will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to commits
	// created by the specified recurring commit ids.
	RecurringCommitIDs param.Field[[]string] `json:"recurring_commit_ids"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to credits
	// created by the specified recurring credit ids.
	RecurringCreditIDs param.Field[[]string] `json:"recurring_credit_ids"`
}

func (r ContractNewParamsOverridesOverrideSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required for OVERWRITE type.
type ContractNewParamsOverridesOverwriteRate struct {
	RateType     param.Field[ContractNewParamsOverridesOverwriteRateRateType] `json:"rate_type,required"`
	CreditTypeID param.Field[string]                                          `json:"credit_type_id" format:"uuid"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate param.Field[map[string]interface{}] `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated param.Field[bool] `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price param.Field[float64] `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Field[float64] `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]shared.TierParam] `json:"tiers"`
}

func (r ContractNewParamsOverridesOverwriteRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsOverridesOverwriteRateRateType string

const (
	ContractNewParamsOverridesOverwriteRateRateTypeFlat         ContractNewParamsOverridesOverwriteRateRateType = "FLAT"
	ContractNewParamsOverridesOverwriteRateRateTypePercentage   ContractNewParamsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractNewParamsOverridesOverwriteRateRateTypeSubscription ContractNewParamsOverridesOverwriteRateRateType = "SUBSCRIPTION"
	ContractNewParamsOverridesOverwriteRateRateTypeTiered       ContractNewParamsOverridesOverwriteRateRateType = "TIERED"
	ContractNewParamsOverridesOverwriteRateRateTypeCustom       ContractNewParamsOverridesOverwriteRateRateType = "CUSTOM"
)

func (r ContractNewParamsOverridesOverwriteRateRateType) IsKnown() bool {
	switch r {
	case ContractNewParamsOverridesOverwriteRateRateTypeFlat, ContractNewParamsOverridesOverwriteRateRateTypePercentage, ContractNewParamsOverridesOverwriteRateRateTypeSubscription, ContractNewParamsOverridesOverwriteRateRateTypeTiered, ContractNewParamsOverridesOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

// Indicates whether the override applies to commit rates or list rates. Can only
// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
// `"LIST_RATE"`.
type ContractNewParamsOverridesTarget string

const (
	ContractNewParamsOverridesTargetCommitRate ContractNewParamsOverridesTarget = "COMMIT_RATE"
	ContractNewParamsOverridesTargetListRate   ContractNewParamsOverridesTarget = "LIST_RATE"
)

func (r ContractNewParamsOverridesTarget) IsKnown() bool {
	switch r {
	case ContractNewParamsOverridesTargetCommitRate, ContractNewParamsOverridesTargetListRate:
		return true
	}
	return false
}

type ContractNewParamsOverridesTier struct {
	Multiplier param.Field[float64] `json:"multiplier,required"`
	Size       param.Field[float64] `json:"size"`
}

func (r ContractNewParamsOverridesTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Overwrites are prioritized over multipliers and tiered overrides.
type ContractNewParamsOverridesType string

const (
	ContractNewParamsOverridesTypeOverwrite  ContractNewParamsOverridesType = "OVERWRITE"
	ContractNewParamsOverridesTypeMultiplier ContractNewParamsOverridesType = "MULTIPLIER"
	ContractNewParamsOverridesTypeTiered     ContractNewParamsOverridesType = "TIERED"
)

func (r ContractNewParamsOverridesType) IsKnown() bool {
	switch r {
	case ContractNewParamsOverridesTypeOverwrite, ContractNewParamsOverridesTypeMultiplier, ContractNewParamsOverridesTypeTiered:
		return true
	}
	return false
}

type ContractNewParamsProfessionalService struct {
	// Maximum amount for the term.
	MaxAmount param.Field[float64] `json:"max_amount,required"`
	ProductID param.Field[string]  `json:"product_id,required" format:"uuid"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity param.Field[float64] `json:"quantity,required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified.
	UnitPrice    param.Field[float64]           `json:"unit_price,required"`
	CustomFields param.Field[map[string]string] `json:"custom_fields"`
	Description  param.Field[string]            `json:"description"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractNewParamsProfessionalService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsRecurringCommit struct {
	// The amount of commit to grant.
	AccessAmount param.Field[ContractNewParamsRecurringCommitsAccessAmount] `json:"access_amount,required"`
	// The amount of time the created commits will be valid for.
	CommitDuration param.Field[ContractNewParamsRecurringCommitsCommitDuration] `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority  param.Field[float64] `json:"priority,required"`
	ProductID param.Field[string]  `json:"product_id,required" format:"uuid"`
	// determines the start time for the first commit
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// Will be passed down to the individual commits
	Description param.Field[string] `json:"description"`
	// determines when the contract will stop creating recurring commits. optional
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount param.Field[ContractNewParamsRecurringCommitsInvoiceAmount] `json:"invoice_amount"`
	// displayed on invoices. will be passed through to the individual commits
	Name param.Field[string] `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// Whether the created commits will use the commit rate or list rate
	RateType param.Field[ContractNewParamsRecurringCommitsRateType] `json:"rate_type"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// A temporary ID that can be used to reference the recurring commit for commit
	// specific overrides.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r ContractNewParamsRecurringCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of commit to grant.
type ContractNewParamsRecurringCommitsAccessAmount struct {
	CreditTypeID param.Field[string]  `json:"credit_type_id,required" format:"uuid"`
	Quantity     param.Field[float64] `json:"quantity,required"`
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
}

func (r ContractNewParamsRecurringCommitsAccessAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of time the created commits will be valid for.
type ContractNewParamsRecurringCommitsCommitDuration struct {
	Unit  param.Field[ContractNewParamsRecurringCommitsCommitDurationUnit] `json:"unit,required"`
	Value param.Field[float64]                                             `json:"value,required"`
}

func (r ContractNewParamsRecurringCommitsCommitDuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsRecurringCommitsCommitDurationUnit string

const (
	ContractNewParamsRecurringCommitsCommitDurationUnitPeriods ContractNewParamsRecurringCommitsCommitDurationUnit = "PERIODS"
)

func (r ContractNewParamsRecurringCommitsCommitDurationUnit) IsKnown() bool {
	switch r {
	case ContractNewParamsRecurringCommitsCommitDurationUnitPeriods:
		return true
	}
	return false
}

// The amount the customer should be billed for the commit. Not required.
type ContractNewParamsRecurringCommitsInvoiceAmount struct {
	CreditTypeID param.Field[string]  `json:"credit_type_id,required" format:"uuid"`
	Quantity     param.Field[float64] `json:"quantity,required"`
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
}

func (r ContractNewParamsRecurringCommitsInvoiceAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the created commits will use the commit rate or list rate
type ContractNewParamsRecurringCommitsRateType string

const (
	ContractNewParamsRecurringCommitsRateTypeCommitRate ContractNewParamsRecurringCommitsRateType = "COMMIT_RATE"
	ContractNewParamsRecurringCommitsRateTypeListRate   ContractNewParamsRecurringCommitsRateType = "LIST_RATE"
)

func (r ContractNewParamsRecurringCommitsRateType) IsKnown() bool {
	switch r {
	case ContractNewParamsRecurringCommitsRateTypeCommitRate, ContractNewParamsRecurringCommitsRateTypeListRate:
		return true
	}
	return false
}

type ContractNewParamsRecurringCredit struct {
	// The amount of commit to grant.
	AccessAmount param.Field[ContractNewParamsRecurringCreditsAccessAmount] `json:"access_amount,required"`
	// The amount of time the created commits will be valid for.
	CommitDuration param.Field[ContractNewParamsRecurringCreditsCommitDuration] `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority  param.Field[float64] `json:"priority,required"`
	ProductID param.Field[string]  `json:"product_id,required" format:"uuid"`
	// determines the start time for the first commit
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// Will be passed down to the individual commits
	Description param.Field[string] `json:"description"`
	// determines when the contract will stop creating recurring commits. optional
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount param.Field[ContractNewParamsRecurringCreditsInvoiceAmount] `json:"invoice_amount"`
	// displayed on invoices. will be passed through to the individual commits
	Name param.Field[string] `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// Whether the created commits will use the commit rate or list rate
	RateType param.Field[ContractNewParamsRecurringCreditsRateType] `json:"rate_type"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// A temporary ID that can be used to reference the recurring commit for commit
	// specific overrides.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r ContractNewParamsRecurringCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of commit to grant.
type ContractNewParamsRecurringCreditsAccessAmount struct {
	CreditTypeID param.Field[string]  `json:"credit_type_id,required" format:"uuid"`
	Quantity     param.Field[float64] `json:"quantity,required"`
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
}

func (r ContractNewParamsRecurringCreditsAccessAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of time the created commits will be valid for.
type ContractNewParamsRecurringCreditsCommitDuration struct {
	Unit  param.Field[ContractNewParamsRecurringCreditsCommitDurationUnit] `json:"unit,required"`
	Value param.Field[float64]                                             `json:"value,required"`
}

func (r ContractNewParamsRecurringCreditsCommitDuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsRecurringCreditsCommitDurationUnit string

const (
	ContractNewParamsRecurringCreditsCommitDurationUnitPeriods ContractNewParamsRecurringCreditsCommitDurationUnit = "PERIODS"
)

func (r ContractNewParamsRecurringCreditsCommitDurationUnit) IsKnown() bool {
	switch r {
	case ContractNewParamsRecurringCreditsCommitDurationUnitPeriods:
		return true
	}
	return false
}

// The amount the customer should be billed for the commit. Not required.
type ContractNewParamsRecurringCreditsInvoiceAmount struct {
	CreditTypeID param.Field[string]  `json:"credit_type_id,required" format:"uuid"`
	Quantity     param.Field[float64] `json:"quantity,required"`
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
}

func (r ContractNewParamsRecurringCreditsInvoiceAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the created commits will use the commit rate or list rate
type ContractNewParamsRecurringCreditsRateType string

const (
	ContractNewParamsRecurringCreditsRateTypeCommitRate ContractNewParamsRecurringCreditsRateType = "COMMIT_RATE"
	ContractNewParamsRecurringCreditsRateTypeListRate   ContractNewParamsRecurringCreditsRateType = "LIST_RATE"
)

func (r ContractNewParamsRecurringCreditsRateType) IsKnown() bool {
	switch r {
	case ContractNewParamsRecurringCreditsRateTypeCommitRate, ContractNewParamsRecurringCreditsRateTypeListRate:
		return true
	}
	return false
}

type ContractNewParamsResellerRoyalty struct {
	Fraction           param.Field[float64]                                        `json:"fraction,required"`
	NetsuiteResellerID param.Field[string]                                         `json:"netsuite_reseller_id,required"`
	ResellerType       param.Field[ContractNewParamsResellerRoyaltiesResellerType] `json:"reseller_type,required"`
	StartingAt         param.Field[time.Time]                                      `json:"starting_at,required" format:"date-time"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags param.Field[[]string]                                     `json:"applicable_product_tags"`
	AwsOptions            param.Field[ContractNewParamsResellerRoyaltiesAwsOptions] `json:"aws_options"`
	EndingBefore          param.Field[time.Time]                                    `json:"ending_before" format:"date-time"`
	GcpOptions            param.Field[ContractNewParamsResellerRoyaltiesGcpOptions] `json:"gcp_options"`
	ResellerContractValue param.Field[float64]                                      `json:"reseller_contract_value"`
}

func (r ContractNewParamsResellerRoyalty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsResellerRoyaltiesResellerType string

const (
	ContractNewParamsResellerRoyaltiesResellerTypeAws           ContractNewParamsResellerRoyaltiesResellerType = "AWS"
	ContractNewParamsResellerRoyaltiesResellerTypeAwsProService ContractNewParamsResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	ContractNewParamsResellerRoyaltiesResellerTypeGcp           ContractNewParamsResellerRoyaltiesResellerType = "GCP"
	ContractNewParamsResellerRoyaltiesResellerTypeGcpProService ContractNewParamsResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r ContractNewParamsResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case ContractNewParamsResellerRoyaltiesResellerTypeAws, ContractNewParamsResellerRoyaltiesResellerTypeAwsProService, ContractNewParamsResellerRoyaltiesResellerTypeGcp, ContractNewParamsResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

type ContractNewParamsResellerRoyaltiesAwsOptions struct {
	AwsAccountNumber    param.Field[string] `json:"aws_account_number"`
	AwsOfferID          param.Field[string] `json:"aws_offer_id"`
	AwsPayerReferenceID param.Field[string] `json:"aws_payer_reference_id"`
}

func (r ContractNewParamsResellerRoyaltiesAwsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsResellerRoyaltiesGcpOptions struct {
	GcpAccountID param.Field[string] `json:"gcp_account_id"`
	GcpOfferID   param.Field[string] `json:"gcp_offer_id"`
}

func (r ContractNewParamsResellerRoyaltiesGcpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsScheduledCharge struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule param.Field[ContractNewParamsScheduledChargesSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractNewParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule.
type ContractNewParamsScheduledChargesSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[ContractNewParamsScheduledChargesScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]ContractNewParamsScheduledChargesScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractNewParamsScheduledChargesSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type ContractNewParamsScheduledChargesScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                           `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractNewParamsScheduledChargesScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution string

const (
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided, ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded, ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly, ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly, ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual, ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type ContractNewParamsScheduledChargesScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractNewParamsScheduledChargesScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type ContractNewParamsScheduledChargesOnUsageInvoices string

const (
	ContractNewParamsScheduledChargesOnUsageInvoicesAll ContractNewParamsScheduledChargesOnUsageInvoices = "ALL"
)

func (r ContractNewParamsScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case ContractNewParamsScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
}

type ContractNewParamsTransition struct {
	FromContractID param.Field[string] `json:"from_contract_id,required" format:"uuid"`
	// This field's available values may vary based on your client's configuration.
	Type                  param.Field[ContractNewParamsTransitionType]                  `json:"type,required"`
	FutureInvoiceBehavior param.Field[ContractNewParamsTransitionFutureInvoiceBehavior] `json:"future_invoice_behavior"`
}

func (r ContractNewParamsTransition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This field's available values may vary based on your client's configuration.
type ContractNewParamsTransitionType string

const (
	ContractNewParamsTransitionTypeSupersede ContractNewParamsTransitionType = "SUPERSEDE"
	ContractNewParamsTransitionTypeRenewal   ContractNewParamsTransitionType = "RENEWAL"
)

func (r ContractNewParamsTransitionType) IsKnown() bool {
	switch r {
	case ContractNewParamsTransitionTypeSupersede, ContractNewParamsTransitionTypeRenewal:
		return true
	}
	return false
}

type ContractNewParamsTransitionFutureInvoiceBehavior struct {
	// Controls whether future trueup invoices are billed or removed. Default behavior
	// is AS_IS if not specified.
	Trueup param.Field[ContractNewParamsTransitionFutureInvoiceBehaviorTrueup] `json:"trueup"`
}

func (r ContractNewParamsTransitionFutureInvoiceBehavior) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls whether future trueup invoices are billed or removed. Default behavior
// is AS_IS if not specified.
type ContractNewParamsTransitionFutureInvoiceBehaviorTrueup string

const (
	ContractNewParamsTransitionFutureInvoiceBehaviorTrueupRemove ContractNewParamsTransitionFutureInvoiceBehaviorTrueup = "REMOVE"
	ContractNewParamsTransitionFutureInvoiceBehaviorTrueupAsIs   ContractNewParamsTransitionFutureInvoiceBehaviorTrueup = "AS_IS"
)

func (r ContractNewParamsTransitionFutureInvoiceBehaviorTrueup) IsKnown() bool {
	switch r {
	case ContractNewParamsTransitionFutureInvoiceBehaviorTrueupRemove, ContractNewParamsTransitionFutureInvoiceBehaviorTrueupAsIs:
		return true
	}
	return false
}

type ContractNewParamsUsageStatementSchedule struct {
	Frequency param.Field[ContractNewParamsUsageStatementScheduleFrequency] `json:"frequency,required"`
	// Required when using CUSTOM_DATE. This option lets you set a historical billing
	// anchor date, aligning future billing cycles with a chosen cadence. For example,
	// if a contract starts on 2024-09-15 and you set the anchor date to 2024-09-10
	// with a MONTHLY frequency, the first usage statement will cover 09-15 to 10-10.
	// Subsequent statements will follow the 10th of each month.
	BillingAnchorDate param.Field[time.Time] `json:"billing_anchor_date" format:"date-time"`
	// If not provided, defaults to the first day of the month.
	Day param.Field[ContractNewParamsUsageStatementScheduleDay] `json:"day"`
	// The date Metronome should start generating usage invoices. If unspecified,
	// contract start date will be used. This is useful to set if you want to import
	// historical invoices via our 'Create Historical Invoices' API rather than having
	// Metronome automatically generate them.
	InvoiceGenerationStartingAt param.Field[time.Time] `json:"invoice_generation_starting_at" format:"date-time"`
}

func (r ContractNewParamsUsageStatementSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsUsageStatementScheduleFrequency string

const (
	ContractNewParamsUsageStatementScheduleFrequencyMonthly   ContractNewParamsUsageStatementScheduleFrequency = "MONTHLY"
	ContractNewParamsUsageStatementScheduleFrequencyQuarterly ContractNewParamsUsageStatementScheduleFrequency = "QUARTERLY"
	ContractNewParamsUsageStatementScheduleFrequencyAnnual    ContractNewParamsUsageStatementScheduleFrequency = "ANNUAL"
)

func (r ContractNewParamsUsageStatementScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractNewParamsUsageStatementScheduleFrequencyMonthly, ContractNewParamsUsageStatementScheduleFrequencyQuarterly, ContractNewParamsUsageStatementScheduleFrequencyAnnual:
		return true
	}
	return false
}

// If not provided, defaults to the first day of the month.
type ContractNewParamsUsageStatementScheduleDay string

const (
	ContractNewParamsUsageStatementScheduleDayFirstOfMonth  ContractNewParamsUsageStatementScheduleDay = "FIRST_OF_MONTH"
	ContractNewParamsUsageStatementScheduleDayContractStart ContractNewParamsUsageStatementScheduleDay = "CONTRACT_START"
	ContractNewParamsUsageStatementScheduleDayCustomDate    ContractNewParamsUsageStatementScheduleDay = "CUSTOM_DATE"
)

func (r ContractNewParamsUsageStatementScheduleDay) IsKnown() bool {
	switch r {
	case ContractNewParamsUsageStatementScheduleDayFirstOfMonth, ContractNewParamsUsageStatementScheduleDayContractStart, ContractNewParamsUsageStatementScheduleDayCustomDate:
		return true
	}
	return false
}

type ContractGetParams struct {
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the query to be slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include commit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
}

func (r ContractGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractListParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Optional RFC 3339 timestamp. If provided, the response will include only
	// contracts effective on the provided date. This cannot be provided if the
	// starting_at filter is provided.
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	// Include archived contracts in the response
	IncludeArchived param.Field[bool] `json:"include_archived"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the query to be slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include commit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
	// Optional RFC 3339 timestamp. If provided, the response will include only
	// contracts where effective_at is on or after the provided date. This cannot be
	// provided if the covering_date filter is provided.
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r ContractListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAddManualBalanceEntryParams struct {
	// ID of the balance (commit or credit) to update.
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// Amount to add to the segment. A negative number will draw down from the balance.
	Amount param.Field[float64] `json:"amount,required"`
	// ID of the customer whose balance is to be updated.
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Reason for the manual adjustment. This will be displayed in the ledger.
	Reason param.Field[string] `json:"reason,required"`
	// ID of the segment to update.
	SegmentID param.Field[string] `json:"segment_id,required" format:"uuid"`
	// ID of the contract to update. Leave blank to update a customer level balance.
	ContractID param.Field[string] `json:"contract_id" format:"uuid"`
	// RFC 3339 timestamp indicating when the manual adjustment takes place. If not
	// provided, it will default to the start of the segment.
	Timestamp param.Field[time.Time] `json:"timestamp" format:"date-time"`
}

func (r ContractAddManualBalanceEntryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParams struct {
	// ID of the contract to amend
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is to be amended
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// inclusive start time for the amendment
	StartingAt   param.Field[time.Time]                   `json:"starting_at,required" format:"date-time"`
	Commits      param.Field[[]ContractAmendParamsCommit] `json:"commits"`
	Credits      param.Field[[]ContractAmendParamsCredit] `json:"credits"`
	CustomFields param.Field[map[string]string]           `json:"custom_fields"`
	// This field's availability is dependent on your client's configuration.
	Discounts param.Field[[]ContractAmendParamsDiscount] `json:"discounts"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string]                        `json:"netsuite_sales_order_id"`
	Overrides            param.Field[[]ContractAmendParamsOverride] `json:"overrides"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices param.Field[[]ContractAmendParamsProfessionalService] `json:"professional_services"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties param.Field[[]ContractAmendParamsResellerRoyalty] `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string]                               `json:"salesforce_opportunity_id"`
	ScheduledCharges        param.Field[[]ContractAmendParamsScheduledCharge] `json:"scheduled_charges"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue param.Field[float64] `json:"total_contract_value"`
}

func (r ContractAmendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommit struct {
	ProductID param.Field[string]                         `json:"product_id,required" format:"uuid"`
	Type      param.Field[ContractAmendParamsCommitsType] `json:"type,required"`
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
	// commits only one schedule item is allowed and amount must match invoice_schedule
	// total.
	AccessSchedule param.Field[ContractAmendParamsCommitsAccessSchedule] `json:"access_schedule"`
	// (DEPRECATED) Use access_schedule and invoice_schedule instead.
	Amount param.Field[float64] `json:"amount"`
	// Which products the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags param.Field[[]string]          `json:"applicable_product_tags"`
	CustomFields          param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this
	// time and only one schedule item is allowed; the total must match access_schedule
	// amount. Optional for "PREPAID" commits: if not provided, this will be a
	// "complimentary" commit with no invoice.
	InvoiceSchedule param.Field[ContractAmendParamsCommitsInvoiceSchedule] `json:"invoice_schedule"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Field[float64]                            `json:"priority"`
	RateType param.Field[ContractAmendParamsCommitsRateType] `json:"rate_type"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// A temporary ID for the commit that can be used to reference the commit for
	// commit specific overrides.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r ContractAmendParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommitsType string

const (
	ContractAmendParamsCommitsTypePrepaid  ContractAmendParamsCommitsType = "PREPAID"
	ContractAmendParamsCommitsTypePostpaid ContractAmendParamsCommitsType = "POSTPAID"
)

func (r ContractAmendParamsCommitsType) IsKnown() bool {
	switch r {
	case ContractAmendParamsCommitsTypePrepaid, ContractAmendParamsCommitsTypePostpaid:
		return true
	}
	return false
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
type ContractAmendParamsCommitsAccessSchedule struct {
	ScheduleItems param.Field[[]ContractAmendParamsCommitsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r ContractAmendParamsCommitsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommitsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r ContractAmendParamsCommitsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match access_schedule
// amount. Optional for "PREPAID" commits: if not provided, this will be a
// "complimentary" commit with no invoice.
type ContractAmendParamsCommitsInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]ContractAmendParamsCommitsInvoiceScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractAmendParamsCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                           `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type ContractAmendParamsCommitsInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractAmendParamsCommitsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommitsRateType string

const (
	ContractAmendParamsCommitsRateTypeCommitRate ContractAmendParamsCommitsRateType = "COMMIT_RATE"
	ContractAmendParamsCommitsRateTypeListRate   ContractAmendParamsCommitsRateType = "LIST_RATE"
)

func (r ContractAmendParamsCommitsRateType) IsKnown() bool {
	switch r {
	case ContractAmendParamsCommitsRateTypeCommitRate, ContractAmendParamsCommitsRateTypeListRate:
		return true
	}
	return false
}

type ContractAmendParamsCredit struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule param.Field[ContractAmendParamsCreditsAccessSchedule] `json:"access_schedule,required"`
	ProductID      param.Field[string]                                   `json:"product_id,required" format:"uuid"`
	// Which products the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags param.Field[[]string]          `json:"applicable_product_tags"`
	CustomFields          param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Field[float64]                            `json:"priority"`
	RateType param.Field[ContractAmendParamsCreditsRateType] `json:"rate_type"`
}

func (r ContractAmendParamsCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the credit to the customer.
type ContractAmendParamsCreditsAccessSchedule struct {
	ScheduleItems param.Field[[]ContractAmendParamsCreditsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r ContractAmendParamsCreditsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsCreditsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r ContractAmendParamsCreditsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsCreditsRateType string

const (
	ContractAmendParamsCreditsRateTypeCommitRate ContractAmendParamsCreditsRateType = "COMMIT_RATE"
	ContractAmendParamsCreditsRateTypeListRate   ContractAmendParamsCreditsRateType = "LIST_RATE"
)

func (r ContractAmendParamsCreditsRateType) IsKnown() bool {
	switch r {
	case ContractAmendParamsCreditsRateTypeCommitRate, ContractAmendParamsCreditsRateTypeListRate:
		return true
	}
	return false
}

type ContractAmendParamsDiscount struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule     param.Field[ContractAmendParamsDiscountsSchedule] `json:"schedule,required"`
	CustomFields param.Field[map[string]string]                    `json:"custom_fields"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractAmendParamsDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule.
type ContractAmendParamsDiscountsSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[ContractAmendParamsDiscountsScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]ContractAmendParamsDiscountsScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractAmendParamsDiscountsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type ContractAmendParamsDiscountsScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                      `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractAmendParamsDiscountsScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution string

const (
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided, ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded, ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency string

const (
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly, ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly, ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual, ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type ContractAmendParamsDiscountsScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractAmendParamsDiscountsScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// tags identifying products whose rates are being overridden. Cannot be used in
	// conjunction with override_specifiers.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Entitled     param.Field[bool]      `json:"entitled"`
	// Indicates whether the override should only apply to commits. Defaults to
	// `false`. If `true`, you can specify relevant commits in `override_specifiers` by
	// passing `commit_ids`. if you do not specify `commit_ids`, then the override will
	// apply when consuming any prepaid or postpaid commit.
	IsCommitSpecific param.Field[bool] `json:"is_commit_specific"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Field[float64] `json:"multiplier"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If
	// provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers param.Field[[]ContractAmendParamsOverridesOverrideSpecifier] `json:"override_specifiers"`
	// Required for OVERWRITE type.
	OverwriteRate param.Field[ContractAmendParamsOverridesOverwriteRate] `json:"overwrite_rate"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides.
	// Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered
	// and multiplier overrides are prioritized by their priority value (lowest first).
	// Must be > 0.
	Priority param.Field[float64] `json:"priority"`
	// ID of the product whose rate is being overridden. Cannot be used in conjunction
	// with override_specifiers.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// Indicates whether the override applies to commit rates or list rates. Can only
	// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
	// `"LIST_RATE"`.
	Target param.Field[ContractAmendParamsOverridesTarget] `json:"target"`
	// Required for TIERED type. Must have at least one tier.
	Tiers param.Field[[]ContractAmendParamsOverridesTier] `json:"tiers"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type param.Field[ContractAmendParamsOverridesType] `json:"type"`
}

func (r ContractAmendParamsOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsOverridesOverrideSpecifier struct {
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to the
	// specified commits. If not provided, the override will apply to all commits.
	CommitIDs param.Field[[]string] `json:"commit_ids"`
	// A map of group names to values. The override will only apply to line items with
	// the specified presentation group values. Can only be used for multiplier
	// overrides.
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	// A map of pricing group names to values. The override will only apply to products
	// with the specified pricing group values.
	PricingGroupValues param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the override will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the override will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to commits
	// created by the specified recurring commit ids.
	RecurringCommitIDs param.Field[[]string] `json:"recurring_commit_ids"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to credits
	// created by the specified recurring credit ids.
	RecurringCreditIDs param.Field[[]string] `json:"recurring_credit_ids"`
}

func (r ContractAmendParamsOverridesOverrideSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required for OVERWRITE type.
type ContractAmendParamsOverridesOverwriteRate struct {
	RateType     param.Field[ContractAmendParamsOverridesOverwriteRateRateType] `json:"rate_type,required"`
	CreditTypeID param.Field[string]                                            `json:"credit_type_id" format:"uuid"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate param.Field[map[string]interface{}] `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated param.Field[bool] `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price param.Field[float64] `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Field[float64] `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]shared.TierParam] `json:"tiers"`
}

func (r ContractAmendParamsOverridesOverwriteRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsOverridesOverwriteRateRateType string

const (
	ContractAmendParamsOverridesOverwriteRateRateTypeFlat         ContractAmendParamsOverridesOverwriteRateRateType = "FLAT"
	ContractAmendParamsOverridesOverwriteRateRateTypePercentage   ContractAmendParamsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractAmendParamsOverridesOverwriteRateRateTypeSubscription ContractAmendParamsOverridesOverwriteRateRateType = "SUBSCRIPTION"
	ContractAmendParamsOverridesOverwriteRateRateTypeTiered       ContractAmendParamsOverridesOverwriteRateRateType = "TIERED"
	ContractAmendParamsOverridesOverwriteRateRateTypeCustom       ContractAmendParamsOverridesOverwriteRateRateType = "CUSTOM"
)

func (r ContractAmendParamsOverridesOverwriteRateRateType) IsKnown() bool {
	switch r {
	case ContractAmendParamsOverridesOverwriteRateRateTypeFlat, ContractAmendParamsOverridesOverwriteRateRateTypePercentage, ContractAmendParamsOverridesOverwriteRateRateTypeSubscription, ContractAmendParamsOverridesOverwriteRateRateTypeTiered, ContractAmendParamsOverridesOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

// Indicates whether the override applies to commit rates or list rates. Can only
// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
// `"LIST_RATE"`.
type ContractAmendParamsOverridesTarget string

const (
	ContractAmendParamsOverridesTargetCommitRate ContractAmendParamsOverridesTarget = "COMMIT_RATE"
	ContractAmendParamsOverridesTargetListRate   ContractAmendParamsOverridesTarget = "LIST_RATE"
)

func (r ContractAmendParamsOverridesTarget) IsKnown() bool {
	switch r {
	case ContractAmendParamsOverridesTargetCommitRate, ContractAmendParamsOverridesTargetListRate:
		return true
	}
	return false
}

type ContractAmendParamsOverridesTier struct {
	Multiplier param.Field[float64] `json:"multiplier,required"`
	Size       param.Field[float64] `json:"size"`
}

func (r ContractAmendParamsOverridesTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Overwrites are prioritized over multipliers and tiered overrides.
type ContractAmendParamsOverridesType string

const (
	ContractAmendParamsOverridesTypeOverwrite  ContractAmendParamsOverridesType = "OVERWRITE"
	ContractAmendParamsOverridesTypeMultiplier ContractAmendParamsOverridesType = "MULTIPLIER"
	ContractAmendParamsOverridesTypeTiered     ContractAmendParamsOverridesType = "TIERED"
)

func (r ContractAmendParamsOverridesType) IsKnown() bool {
	switch r {
	case ContractAmendParamsOverridesTypeOverwrite, ContractAmendParamsOverridesTypeMultiplier, ContractAmendParamsOverridesTypeTiered:
		return true
	}
	return false
}

type ContractAmendParamsProfessionalService struct {
	// Maximum amount for the term.
	MaxAmount param.Field[float64] `json:"max_amount,required"`
	ProductID param.Field[string]  `json:"product_id,required" format:"uuid"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity param.Field[float64] `json:"quantity,required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified.
	UnitPrice    param.Field[float64]           `json:"unit_price,required"`
	CustomFields param.Field[map[string]string] `json:"custom_fields"`
	Description  param.Field[string]            `json:"description"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractAmendParamsProfessionalService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsResellerRoyalty struct {
	ResellerType param.Field[ContractAmendParamsResellerRoyaltiesResellerType] `json:"reseller_type,required"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags param.Field[[]string]                                       `json:"applicable_product_tags"`
	AwsOptions            param.Field[ContractAmendParamsResellerRoyaltiesAwsOptions] `json:"aws_options"`
	// Use null to indicate that the existing end timestamp should be removed.
	EndingBefore          param.Field[time.Time]                                      `json:"ending_before" format:"date-time"`
	Fraction              param.Field[float64]                                        `json:"fraction"`
	GcpOptions            param.Field[ContractAmendParamsResellerRoyaltiesGcpOptions] `json:"gcp_options"`
	NetsuiteResellerID    param.Field[string]                                         `json:"netsuite_reseller_id"`
	ResellerContractValue param.Field[float64]                                        `json:"reseller_contract_value"`
	StartingAt            param.Field[time.Time]                                      `json:"starting_at" format:"date-time"`
}

func (r ContractAmendParamsResellerRoyalty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsResellerRoyaltiesResellerType string

const (
	ContractAmendParamsResellerRoyaltiesResellerTypeAws           ContractAmendParamsResellerRoyaltiesResellerType = "AWS"
	ContractAmendParamsResellerRoyaltiesResellerTypeAwsProService ContractAmendParamsResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	ContractAmendParamsResellerRoyaltiesResellerTypeGcp           ContractAmendParamsResellerRoyaltiesResellerType = "GCP"
	ContractAmendParamsResellerRoyaltiesResellerTypeGcpProService ContractAmendParamsResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r ContractAmendParamsResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case ContractAmendParamsResellerRoyaltiesResellerTypeAws, ContractAmendParamsResellerRoyaltiesResellerTypeAwsProService, ContractAmendParamsResellerRoyaltiesResellerTypeGcp, ContractAmendParamsResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

type ContractAmendParamsResellerRoyaltiesAwsOptions struct {
	AwsAccountNumber    param.Field[string] `json:"aws_account_number"`
	AwsOfferID          param.Field[string] `json:"aws_offer_id"`
	AwsPayerReferenceID param.Field[string] `json:"aws_payer_reference_id"`
}

func (r ContractAmendParamsResellerRoyaltiesAwsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsResellerRoyaltiesGcpOptions struct {
	GcpAccountID param.Field[string] `json:"gcp_account_id"`
	GcpOfferID   param.Field[string] `json:"gcp_offer_id"`
}

func (r ContractAmendParamsResellerRoyaltiesGcpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsScheduledCharge struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule param.Field[ContractAmendParamsScheduledChargesSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractAmendParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule.
type ContractAmendParamsScheduledChargesSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[ContractAmendParamsScheduledChargesScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]ContractAmendParamsScheduledChargesScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractAmendParamsScheduledChargesSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type ContractAmendParamsScheduledChargesScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                             `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractAmendParamsScheduledChargesScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution string

const (
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided, ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded, ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly, ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly, ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual, ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type ContractAmendParamsScheduledChargesScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractAmendParamsScheduledChargesScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractArchiveParams struct {
	// ID of the contract to archive
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is to be archived
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// If false, the existing finalized invoices will remain after the contract is
	// archived.
	VoidInvoices param.Field[bool] `json:"void_invoices,required"`
}

func (r ContractArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewHistoricalInvoicesParams struct {
	Invoices param.Field[[]ContractNewHistoricalInvoicesParamsInvoice] `json:"invoices,required"`
	Preview  param.Field[bool]                                         `json:"preview,required"`
}

func (r ContractNewHistoricalInvoicesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewHistoricalInvoicesParamsInvoice struct {
	ContractID         param.Field[string]                                                     `json:"contract_id,required" format:"uuid"`
	CreditTypeID       param.Field[string]                                                     `json:"credit_type_id,required" format:"uuid"`
	CustomerID         param.Field[string]                                                     `json:"customer_id,required" format:"uuid"`
	ExclusiveEndDate   param.Field[time.Time]                                                  `json:"exclusive_end_date,required" format:"date-time"`
	InclusiveStartDate param.Field[time.Time]                                                  `json:"inclusive_start_date,required" format:"date-time"`
	IssueDate          param.Field[time.Time]                                                  `json:"issue_date,required" format:"date-time"`
	UsageLineItems     param.Field[[]ContractNewHistoricalInvoicesParamsInvoicesUsageLineItem] `json:"usage_line_items,required"`
	// This field's availability is dependent on your client's configuration.
	BillableStatus       param.Field[ContractNewHistoricalInvoicesParamsInvoicesBillableStatus]       `json:"billable_status"`
	BreakdownGranularity param.Field[ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularity] `json:"breakdown_granularity"`
	CustomFields         param.Field[map[string]string]                                               `json:"custom_fields"`
}

func (r ContractNewHistoricalInvoicesParamsInvoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewHistoricalInvoicesParamsInvoicesUsageLineItem struct {
	ExclusiveEndDate        param.Field[time.Time]                                                                        `json:"exclusive_end_date,required" format:"date-time"`
	InclusiveStartDate      param.Field[time.Time]                                                                        `json:"inclusive_start_date,required" format:"date-time"`
	ProductID               param.Field[string]                                                                           `json:"product_id,required" format:"uuid"`
	PresentationGroupValues param.Field[map[string]string]                                                                `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string]                                                                `json:"pricing_group_values"`
	Quantity                param.Field[float64]                                                                          `json:"quantity"`
	SubtotalsWithQuantity   param.Field[[]ContractNewHistoricalInvoicesParamsInvoicesUsageLineItemsSubtotalsWithQuantity] `json:"subtotals_with_quantity"`
}

func (r ContractNewHistoricalInvoicesParamsInvoicesUsageLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewHistoricalInvoicesParamsInvoicesUsageLineItemsSubtotalsWithQuantity struct {
	ExclusiveEndDate   param.Field[time.Time] `json:"exclusive_end_date,required" format:"date-time"`
	InclusiveStartDate param.Field[time.Time] `json:"inclusive_start_date,required" format:"date-time"`
	Quantity           param.Field[float64]   `json:"quantity,required"`
}

func (r ContractNewHistoricalInvoicesParamsInvoicesUsageLineItemsSubtotalsWithQuantity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This field's availability is dependent on your client's configuration.
type ContractNewHistoricalInvoicesParamsInvoicesBillableStatus string

const (
	ContractNewHistoricalInvoicesParamsInvoicesBillableStatusBillable   ContractNewHistoricalInvoicesParamsInvoicesBillableStatus = "billable"
	ContractNewHistoricalInvoicesParamsInvoicesBillableStatusUnbillable ContractNewHistoricalInvoicesParamsInvoicesBillableStatus = "unbillable"
)

func (r ContractNewHistoricalInvoicesParamsInvoicesBillableStatus) IsKnown() bool {
	switch r {
	case ContractNewHistoricalInvoicesParamsInvoicesBillableStatusBillable, ContractNewHistoricalInvoicesParamsInvoicesBillableStatusUnbillable:
		return true
	}
	return false
}

type ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularity string

const (
	ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularityHour ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularity = "HOUR"
	ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularityDay  ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularity = "DAY"
)

func (r ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularity) IsKnown() bool {
	switch r {
	case ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularityHour, ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularityDay:
		return true
	}
	return false
}

type ContractListBalancesParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	ID         param.Field[string] `json:"id" format:"uuid"`
	// Return only balances that have access schedules that "cover" the provided date
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	// Include only balances that have any access before the provided date (exclusive)
	EffectiveBefore param.Field[time.Time] `json:"effective_before" format:"date-time"`
	// Include credits from archived contracts.
	IncludeArchived param.Field[bool] `json:"include_archived"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the query to be slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include balances on the contract level.
	IncludeContractBalances param.Field[bool] `json:"include_contract_balances"`
	// Include ledgers in the response. Setting this flag may cause the query to be
	// slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
	// The next page token from a previous response.
	NextPage param.Field[string] `json:"next_page"`
	// Include only balances that have any access on or after the provided date
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r ContractListBalancesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractGetRateScheduleParams struct {
	// ID of the contract to get the rate schedule for.
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	// ID of the customer for whose contract to get the rate schedule for.
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// optional timestamp which overlaps with the returned rate schedule segments. When
	// not specified, the current timestamp will be used.
	At param.Field[time.Time] `json:"at" format:"date-time"`
	// List of rate selectors, rates matching ANY of the selectors will be included in
	// the response. Passing no selectors will result in all rates being returned.
	Selectors param.Field[[]ContractGetRateScheduleParamsSelector] `json:"selectors"`
}

func (r ContractGetRateScheduleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ContractGetRateScheduleParams]'s query parameters as
// `url.Values`.
func (r ContractGetRateScheduleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContractGetRateScheduleParamsSelector struct {
	// List of pricing group key value pairs, rates containing the matching key / value
	// pairs will be included in the response.
	PartialPricingGroupValues param.Field[map[string]string] `json:"partial_pricing_group_values"`
	// List of pricing group key value pairs, rates matching all of the key / value
	// pairs will be included in the response.
	PricingGroupValues param.Field[map[string]string] `json:"pricing_group_values"`
	// Rates matching the product id will be included in the response.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// List of product tags, rates matching any of the tags will be included in the
	// response.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r ContractGetRateScheduleParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractScheduleProServicesInvoiceParams struct {
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// The date the invoice is issued
	IssuedAt param.Field[time.Time] `json:"issued_at,required" format:"date-time"`
	// Each line requires an amount or both unit_price and quantity.
	LineItems param.Field[[]ContractScheduleProServicesInvoiceParamsLineItem] `json:"line_items,required"`
	// The end date of the invoice header in Netsuite
	NetsuiteInvoiceHeaderEnd param.Field[time.Time] `json:"netsuite_invoice_header_end" format:"date-time"`
	// The start date of the invoice header in Netsuite
	NetsuiteInvoiceHeaderStart param.Field[time.Time] `json:"netsuite_invoice_header_start" format:"date-time"`
}

func (r ContractScheduleProServicesInvoiceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Describes the line item for a professional service charge on an invoice.
type ContractScheduleProServicesInvoiceParamsLineItem struct {
	ProfessionalServiceID param.Field[string] `json:"professional_service_id,required" format:"uuid"`
	// If the professional_service_id was added on an amendment, this is required.
	AmendmentID param.Field[string] `json:"amendment_id" format:"uuid"`
	// Amount for the term on the new invoice.
	Amount param.Field[float64] `json:"amount"`
	// For client use.
	Metadata param.Field[string] `json:"metadata"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd param.Field[time.Time] `json:"netsuite_invoice_billing_end" format:"date-time"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart param.Field[time.Time] `json:"netsuite_invoice_billing_start" format:"date-time"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity param.Field[float64] `json:"quantity"`
	// If specified, this overrides the unit price on the pro service term. Must also
	// provide quantity (but not amount) if providing unit_price.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractScheduleProServicesInvoiceParamsLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractSetUsageFilterParams struct {
	ContractID  param.Field[string]    `json:"contract_id,required" format:"uuid"`
	CustomerID  param.Field[string]    `json:"customer_id,required" format:"uuid"`
	GroupKey    param.Field[string]    `json:"group_key,required"`
	GroupValues param.Field[[]string]  `json:"group_values,required"`
	StartingAt  param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r ContractSetUsageFilterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractUpdateEndDateParams struct {
	// ID of the contract to update
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is to be updated
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// If true, allows setting the contract end date earlier than the end_timestamp of
	// existing finalized invoices. Finalized invoices will be unchanged; if you want
	// to incorporate the new end date, you can void and regenerate finalized usage
	// invoices. Defaults to true.
	AllowEndingBeforeFinalizedInvoice param.Field[bool] `json:"allow_ending_before_finalized_invoice"`
	// RFC 3339 timestamp indicating when the contract will end (exclusive). If not
	// provided, the contract will be updated to be open-ended.
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
}

func (r ContractUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
