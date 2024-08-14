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

// Create a new, scheduled invoice for Professional Services terms on a contract.
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
	ID           string                             `json:"id,required" format:"uuid"`
	Amendments   []ContractGetResponseDataAmendment `json:"amendments,required"`
	Current      shared.ContractWithoutAmendments   `json:"current,required"`
	CustomerID   string                             `json:"customer_id,required" format:"uuid"`
	Initial      shared.ContractWithoutAmendments   `json:"initial,required"`
	CustomFields map[string]string                  `json:"custom_fields"`
	// This field's availability is dependent on your client's configuration.
	CustomerBillingProviderConfiguration ContractGetResponseDataCustomerBillingProviderConfiguration `json:"customer_billing_provider_configuration"`
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
	CustomFields                         apijson.Field
	CustomerBillingProviderConfiguration apijson.Field
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
	ID               string                                    `json:"id,required" format:"uuid"`
	Commits          []shared.Commit                           `json:"commits,required"`
	CreatedAt        time.Time                                 `json:"created_at,required" format:"date-time"`
	CreatedBy        string                                    `json:"created_by,required"`
	Overrides        []shared.Override                         `json:"overrides,required"`
	ScheduledCharges []shared.ScheduledCharge                  `json:"scheduled_charges,required"`
	StartingAt       time.Time                                 `json:"starting_at,required" format:"date-time"`
	Credits          []ContractGetResponseDataAmendmentsCredit `json:"credits"`
	// This field's availability is dependent on your client's configuration.
	Discounts []shared.Discount `json:"discounts"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []ContractGetResponseDataAmendmentsProfessionalService `json:"professional_services"`
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

type ContractGetResponseDataAmendmentsCredit struct {
	ID      string                                          `json:"id,required" format:"uuid"`
	Product ContractGetResponseDataAmendmentsCreditsProduct `json:"product,required"`
	Type    ContractGetResponseDataAmendmentsCreditsType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        ContractGetResponseDataAmendmentsCreditsAccessSchedule `json:"access_schedule"`
	ApplicableContractIDs []string                                               `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string                                               `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                                               `json:"applicable_product_tags"`
	Contract              ContractGetResponseDataAmendmentsCreditsContract       `json:"contract"`
	CustomFields          map[string]string                                      `json:"custom_fields"`
	Description           string                                                 `json:"description"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []ContractGetResponseDataAmendmentsCreditsLedger `json:"ledger"`
	Name   string                                           `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                                      `json:"salesforce_opportunity_id"`
	JSON                    contractGetResponseDataAmendmentsCreditJSON `json:"-"`
}

// contractGetResponseDataAmendmentsCreditJSON contains the JSON metadata for the
// struct [ContractGetResponseDataAmendmentsCredit]
type contractGetResponseDataAmendmentsCreditJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsCreditJSON) RawJSON() string {
	return r.raw
}

type ContractGetResponseDataAmendmentsCreditsProduct struct {
	ID   string                                              `json:"id,required" format:"uuid"`
	Name string                                              `json:"name,required"`
	JSON contractGetResponseDataAmendmentsCreditsProductJSON `json:"-"`
}

// contractGetResponseDataAmendmentsCreditsProductJSON contains the JSON metadata
// for the struct [ContractGetResponseDataAmendmentsCreditsProduct]
type contractGetResponseDataAmendmentsCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsCreditsProductJSON) RawJSON() string {
	return r.raw
}

type ContractGetResponseDataAmendmentsCreditsType string

const (
	ContractGetResponseDataAmendmentsCreditsTypeCredit ContractGetResponseDataAmendmentsCreditsType = "CREDIT"
)

func (r ContractGetResponseDataAmendmentsCreditsType) IsKnown() bool {
	switch r {
	case ContractGetResponseDataAmendmentsCreditsTypeCredit:
		return true
	}
	return false
}

// The schedule that the customer will gain access to the credits.
type ContractGetResponseDataAmendmentsCreditsAccessSchedule struct {
	ScheduleItems []ContractGetResponseDataAmendmentsCreditsAccessScheduleScheduleItem `json:"schedule_items,required"`
	CreditType    shared.CreditType                                                    `json:"credit_type"`
	JSON          contractGetResponseDataAmendmentsCreditsAccessScheduleJSON           `json:"-"`
}

// contractGetResponseDataAmendmentsCreditsAccessScheduleJSON contains the JSON
// metadata for the struct [ContractGetResponseDataAmendmentsCreditsAccessSchedule]
type contractGetResponseDataAmendmentsCreditsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCreditsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsCreditsAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type ContractGetResponseDataAmendmentsCreditsAccessScheduleScheduleItem struct {
	ID           string                                                                 `json:"id,required" format:"uuid"`
	Amount       float64                                                                `json:"amount,required"`
	EndingBefore time.Time                                                              `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                                                              `json:"starting_at,required" format:"date-time"`
	JSON         contractGetResponseDataAmendmentsCreditsAccessScheduleScheduleItemJSON `json:"-"`
}

// contractGetResponseDataAmendmentsCreditsAccessScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCreditsAccessScheduleScheduleItem]
type contractGetResponseDataAmendmentsCreditsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCreditsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsCreditsAccessScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type ContractGetResponseDataAmendmentsCreditsContract struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	JSON contractGetResponseDataAmendmentsCreditsContractJSON `json:"-"`
}

// contractGetResponseDataAmendmentsCreditsContractJSON contains the JSON metadata
// for the struct [ContractGetResponseDataAmendmentsCreditsContract]
type contractGetResponseDataAmendmentsCreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsCreditsContractJSON) RawJSON() string {
	return r.raw
}

type ContractGetResponseDataAmendmentsCreditsLedger struct {
	Type      ContractGetResponseDataAmendmentsCreditsLedgerType `json:"type,required"`
	Timestamp time.Time                                          `json:"timestamp,required" format:"date-time"`
	Amount    float64                                            `json:"amount,required"`
	SegmentID string                                             `json:"segment_id" format:"uuid"`
	InvoiceID string                                             `json:"invoice_id" format:"uuid"`
	Reason    string                                             `json:"reason"`
	JSON      contractGetResponseDataAmendmentsCreditsLedgerJSON `json:"-"`
	union     ContractGetResponseDataAmendmentsCreditsLedgerUnion
}

// contractGetResponseDataAmendmentsCreditsLedgerJSON contains the JSON metadata
// for the struct [ContractGetResponseDataAmendmentsCreditsLedger]
type contractGetResponseDataAmendmentsCreditsLedgerJSON struct {
	Type        apijson.Field
	Timestamp   apijson.Field
	Amount      apijson.Field
	SegmentID   apijson.Field
	InvoiceID   apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r contractGetResponseDataAmendmentsCreditsLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *ContractGetResponseDataAmendmentsCreditsLedger) UnmarshalJSON(data []byte) (err error) {
	*r = ContractGetResponseDataAmendmentsCreditsLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContractGetResponseDataAmendmentsCreditsLedgerUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry],
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry],
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry],
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry],
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry].
func (r ContractGetResponseDataAmendmentsCreditsLedger) AsUnion() ContractGetResponseDataAmendmentsCreditsLedgerUnion {
	return r.union
}

// Union satisfied by
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry],
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry],
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry],
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry] or
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry].
type ContractGetResponseDataAmendmentsCreditsLedgerUnion interface {
	implementsContractGetResponseDataAmendmentsCreditsLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContractGetResponseDataAmendmentsCreditsLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry{}),
		},
	)
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64                                                                         `json:"amount,required"`
	SegmentID string                                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON `json:"-"`
}

// contractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry]
type contractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry) implementsContractGetResponseDataAmendmentsCreditsLedger() {
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType = "CREDIT_SEGMENT_START"
)

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractGetResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart:
		return true
	}
	return false
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                      `json:"amount,required"`
	InvoiceID string                                                                                       `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                       `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                    `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// contractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry]
type contractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) implementsContractGetResponseDataAmendmentsCreditsLedger() {
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractGetResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry struct {
	Amount    float64                                                                       `json:"amount,required"`
	SegmentID string                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                     `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON `json:"-"`
}

// contractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry]
type contractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry) implementsContractGetResponseDataAmendmentsCreditsLedger() {
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryType = "CREDIT_EXPIRATION"
)

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractGetResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration:
		return true
	}
	return false
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry struct {
	Amount    float64                                                                     `json:"amount,required"`
	InvoiceID string                                                                      `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                   `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON `json:"-"`
}

// contractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry]
type contractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry) implementsContractGetResponseDataAmendmentsCreditsLedger() {
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryType = "CREDIT_CANCELED"
)

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractGetResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled:
		return true
	}
	return false
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry struct {
	Amount    float64                                                                     `json:"amount,required"`
	InvoiceID string                                                                      `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                   `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON `json:"-"`
}

// contractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry]
type contractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry) implementsContractGetResponseDataAmendmentsCreditsLedger() {
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryType = "CREDIT_CREDITED"
)

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractGetResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited:
		return true
	}
	return false
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry struct {
	Amount    float64                                                                   `json:"amount,required"`
	Reason    string                                                                    `json:"reason,required"`
	Timestamp time.Time                                                                 `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryJSON `json:"-"`
}

// contractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry]
type contractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry) implementsContractGetResponseDataAmendmentsCreditsLedger() {
}

type ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryTypeCreditManual ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryType = "CREDIT_MANUAL"
)

func (r ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractGetResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryTypeCreditManual:
		return true
	}
	return false
}

type ContractGetResponseDataAmendmentsCreditsLedgerType string

const (
	ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditSegmentStart              ContractGetResponseDataAmendmentsCreditsLedgerType = "CREDIT_SEGMENT_START"
	ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditAutomatedInvoiceDeduction ContractGetResponseDataAmendmentsCreditsLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditExpiration                ContractGetResponseDataAmendmentsCreditsLedgerType = "CREDIT_EXPIRATION"
	ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditCanceled                  ContractGetResponseDataAmendmentsCreditsLedgerType = "CREDIT_CANCELED"
	ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditCredited                  ContractGetResponseDataAmendmentsCreditsLedgerType = "CREDIT_CREDITED"
	ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditManual                    ContractGetResponseDataAmendmentsCreditsLedgerType = "CREDIT_MANUAL"
)

func (r ContractGetResponseDataAmendmentsCreditsLedgerType) IsKnown() bool {
	switch r {
	case ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditSegmentStart, ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditAutomatedInvoiceDeduction, ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditExpiration, ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditCanceled, ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditCredited, ContractGetResponseDataAmendmentsCreditsLedgerTypeCreditManual:
		return true
	}
	return false
}

type ContractGetResponseDataAmendmentsProfessionalService struct {
	ID string `json:"id,required" format:"uuid"`
	// Maximum amount for the term.
	MaxAmount float64 `json:"max_amount,required"`
	ProductID string  `json:"product_id,required" format:"uuid"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity float64 `json:"quantity,required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified.
	UnitPrice    float64           `json:"unit_price,required"`
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string                                                   `json:"netsuite_sales_order_id"`
	JSON                 contractGetResponseDataAmendmentsProfessionalServiceJSON `json:"-"`
}

// contractGetResponseDataAmendmentsProfessionalServiceJSON contains the JSON
// metadata for the struct [ContractGetResponseDataAmendmentsProfessionalService]
type contractGetResponseDataAmendmentsProfessionalServiceJSON struct {
	ID                   apijson.Field
	MaxAmount            apijson.Field
	ProductID            apijson.Field
	Quantity             apijson.Field
	UnitPrice            apijson.Field
	CustomFields         apijson.Field
	Description          apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsProfessionalService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractGetResponseDataAmendmentsProfessionalServiceJSON) RawJSON() string {
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

// This field's availability is dependent on your client's configuration.
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
	ID           string                              `json:"id,required" format:"uuid"`
	Amendments   []ContractListResponseDataAmendment `json:"amendments,required"`
	Current      shared.ContractWithoutAmendments    `json:"current,required"`
	CustomerID   string                              `json:"customer_id,required" format:"uuid"`
	Initial      shared.ContractWithoutAmendments    `json:"initial,required"`
	CustomFields map[string]string                   `json:"custom_fields"`
	// This field's availability is dependent on your client's configuration.
	CustomerBillingProviderConfiguration ContractListResponseDataCustomerBillingProviderConfiguration `json:"customer_billing_provider_configuration"`
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
	CustomFields                         apijson.Field
	CustomerBillingProviderConfiguration apijson.Field
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
	ID               string                                     `json:"id,required" format:"uuid"`
	Commits          []shared.Commit                            `json:"commits,required"`
	CreatedAt        time.Time                                  `json:"created_at,required" format:"date-time"`
	CreatedBy        string                                     `json:"created_by,required"`
	Overrides        []shared.Override                          `json:"overrides,required"`
	ScheduledCharges []shared.ScheduledCharge                   `json:"scheduled_charges,required"`
	StartingAt       time.Time                                  `json:"starting_at,required" format:"date-time"`
	Credits          []ContractListResponseDataAmendmentsCredit `json:"credits"`
	// This field's availability is dependent on your client's configuration.
	Discounts []shared.Discount `json:"discounts"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []ContractListResponseDataAmendmentsProfessionalService `json:"professional_services"`
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

type ContractListResponseDataAmendmentsCredit struct {
	ID      string                                           `json:"id,required" format:"uuid"`
	Product ContractListResponseDataAmendmentsCreditsProduct `json:"product,required"`
	Type    ContractListResponseDataAmendmentsCreditsType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        ContractListResponseDataAmendmentsCreditsAccessSchedule `json:"access_schedule"`
	ApplicableContractIDs []string                                                `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string                                                `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                                                `json:"applicable_product_tags"`
	Contract              ContractListResponseDataAmendmentsCreditsContract       `json:"contract"`
	CustomFields          map[string]string                                       `json:"custom_fields"`
	Description           string                                                  `json:"description"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []ContractListResponseDataAmendmentsCreditsLedger `json:"ledger"`
	Name   string                                            `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                                       `json:"salesforce_opportunity_id"`
	JSON                    contractListResponseDataAmendmentsCreditJSON `json:"-"`
}

// contractListResponseDataAmendmentsCreditJSON contains the JSON metadata for the
// struct [ContractListResponseDataAmendmentsCredit]
type contractListResponseDataAmendmentsCreditJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsCreditJSON) RawJSON() string {
	return r.raw
}

type ContractListResponseDataAmendmentsCreditsProduct struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	Name string                                               `json:"name,required"`
	JSON contractListResponseDataAmendmentsCreditsProductJSON `json:"-"`
}

// contractListResponseDataAmendmentsCreditsProductJSON contains the JSON metadata
// for the struct [ContractListResponseDataAmendmentsCreditsProduct]
type contractListResponseDataAmendmentsCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsCreditsProductJSON) RawJSON() string {
	return r.raw
}

type ContractListResponseDataAmendmentsCreditsType string

const (
	ContractListResponseDataAmendmentsCreditsTypeCredit ContractListResponseDataAmendmentsCreditsType = "CREDIT"
)

func (r ContractListResponseDataAmendmentsCreditsType) IsKnown() bool {
	switch r {
	case ContractListResponseDataAmendmentsCreditsTypeCredit:
		return true
	}
	return false
}

// The schedule that the customer will gain access to the credits.
type ContractListResponseDataAmendmentsCreditsAccessSchedule struct {
	ScheduleItems []ContractListResponseDataAmendmentsCreditsAccessScheduleScheduleItem `json:"schedule_items,required"`
	CreditType    shared.CreditType                                                     `json:"credit_type"`
	JSON          contractListResponseDataAmendmentsCreditsAccessScheduleJSON           `json:"-"`
}

// contractListResponseDataAmendmentsCreditsAccessScheduleJSON contains the JSON
// metadata for the struct
// [ContractListResponseDataAmendmentsCreditsAccessSchedule]
type contractListResponseDataAmendmentsCreditsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCreditsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsCreditsAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type ContractListResponseDataAmendmentsCreditsAccessScheduleScheduleItem struct {
	ID           string                                                                  `json:"id,required" format:"uuid"`
	Amount       float64                                                                 `json:"amount,required"`
	EndingBefore time.Time                                                               `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                                                               `json:"starting_at,required" format:"date-time"`
	JSON         contractListResponseDataAmendmentsCreditsAccessScheduleScheduleItemJSON `json:"-"`
}

// contractListResponseDataAmendmentsCreditsAccessScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCreditsAccessScheduleScheduleItem]
type contractListResponseDataAmendmentsCreditsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCreditsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsCreditsAccessScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type ContractListResponseDataAmendmentsCreditsContract struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	JSON contractListResponseDataAmendmentsCreditsContractJSON `json:"-"`
}

// contractListResponseDataAmendmentsCreditsContractJSON contains the JSON metadata
// for the struct [ContractListResponseDataAmendmentsCreditsContract]
type contractListResponseDataAmendmentsCreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsCreditsContractJSON) RawJSON() string {
	return r.raw
}

type ContractListResponseDataAmendmentsCreditsLedger struct {
	Type      ContractListResponseDataAmendmentsCreditsLedgerType `json:"type,required"`
	Timestamp time.Time                                           `json:"timestamp,required" format:"date-time"`
	Amount    float64                                             `json:"amount,required"`
	SegmentID string                                              `json:"segment_id" format:"uuid"`
	InvoiceID string                                              `json:"invoice_id" format:"uuid"`
	Reason    string                                              `json:"reason"`
	JSON      contractListResponseDataAmendmentsCreditsLedgerJSON `json:"-"`
	union     ContractListResponseDataAmendmentsCreditsLedgerUnion
}

// contractListResponseDataAmendmentsCreditsLedgerJSON contains the JSON metadata
// for the struct [ContractListResponseDataAmendmentsCreditsLedger]
type contractListResponseDataAmendmentsCreditsLedgerJSON struct {
	Type        apijson.Field
	Timestamp   apijson.Field
	Amount      apijson.Field
	SegmentID   apijson.Field
	InvoiceID   apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r contractListResponseDataAmendmentsCreditsLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *ContractListResponseDataAmendmentsCreditsLedger) UnmarshalJSON(data []byte) (err error) {
	*r = ContractListResponseDataAmendmentsCreditsLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContractListResponseDataAmendmentsCreditsLedgerUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry],
// [ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry],
// [ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry],
// [ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry],
// [ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry].
func (r ContractListResponseDataAmendmentsCreditsLedger) AsUnion() ContractListResponseDataAmendmentsCreditsLedgerUnion {
	return r.union
}

// Union satisfied by
// [ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry],
// [ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry],
// [ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry],
// [ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry] or
// [ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry].
type ContractListResponseDataAmendmentsCreditsLedgerUnion interface {
	implementsContractListResponseDataAmendmentsCreditsLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContractListResponseDataAmendmentsCreditsLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry{}),
		},
	)
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64                                                                          `json:"amount,required"`
	SegmentID string                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                        `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON `json:"-"`
}

// contractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry]
type contractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntry) implementsContractListResponseDataAmendmentsCreditsLedger() {
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType = "CREDIT_SEGMENT_START"
)

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListResponseDataAmendmentsCreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart:
		return true
	}
	return false
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                       `json:"amount,required"`
	InvoiceID string                                                                                        `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                     `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// contractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry]
type contractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) implementsContractListResponseDataAmendmentsCreditsLedger() {
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListResponseDataAmendmentsCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry struct {
	Amount    float64                                                                        `json:"amount,required"`
	SegmentID string                                                                         `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                      `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON `json:"-"`
}

// contractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry]
type contractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntry) implementsContractListResponseDataAmendmentsCreditsLedger() {
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryType = "CREDIT_EXPIRATION"
)

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListResponseDataAmendmentsCreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration:
		return true
	}
	return false
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry struct {
	Amount    float64                                                                      `json:"amount,required"`
	InvoiceID string                                                                       `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                       `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                    `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON `json:"-"`
}

// contractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry]
type contractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntry) implementsContractListResponseDataAmendmentsCreditsLedger() {
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryType = "CREDIT_CANCELED"
)

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListResponseDataAmendmentsCreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled:
		return true
	}
	return false
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry struct {
	Amount    float64                                                                      `json:"amount,required"`
	InvoiceID string                                                                       `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                       `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                    `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON `json:"-"`
}

// contractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry]
type contractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntry) implementsContractListResponseDataAmendmentsCreditsLedger() {
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryType = "CREDIT_CREDITED"
)

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListResponseDataAmendmentsCreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited:
		return true
	}
	return false
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry struct {
	Amount    float64                                                                    `json:"amount,required"`
	Reason    string                                                                     `json:"reason,required"`
	Timestamp time.Time                                                                  `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryJSON `json:"-"`
}

// contractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry]
type contractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntry) implementsContractListResponseDataAmendmentsCreditsLedger() {
}

type ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryTypeCreditManual ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryType = "CREDIT_MANUAL"
)

func (r ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListResponseDataAmendmentsCreditsLedgerCreditManualLedgerEntryTypeCreditManual:
		return true
	}
	return false
}

type ContractListResponseDataAmendmentsCreditsLedgerType string

const (
	ContractListResponseDataAmendmentsCreditsLedgerTypeCreditSegmentStart              ContractListResponseDataAmendmentsCreditsLedgerType = "CREDIT_SEGMENT_START"
	ContractListResponseDataAmendmentsCreditsLedgerTypeCreditAutomatedInvoiceDeduction ContractListResponseDataAmendmentsCreditsLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	ContractListResponseDataAmendmentsCreditsLedgerTypeCreditExpiration                ContractListResponseDataAmendmentsCreditsLedgerType = "CREDIT_EXPIRATION"
	ContractListResponseDataAmendmentsCreditsLedgerTypeCreditCanceled                  ContractListResponseDataAmendmentsCreditsLedgerType = "CREDIT_CANCELED"
	ContractListResponseDataAmendmentsCreditsLedgerTypeCreditCredited                  ContractListResponseDataAmendmentsCreditsLedgerType = "CREDIT_CREDITED"
	ContractListResponseDataAmendmentsCreditsLedgerTypeCreditManual                    ContractListResponseDataAmendmentsCreditsLedgerType = "CREDIT_MANUAL"
)

func (r ContractListResponseDataAmendmentsCreditsLedgerType) IsKnown() bool {
	switch r {
	case ContractListResponseDataAmendmentsCreditsLedgerTypeCreditSegmentStart, ContractListResponseDataAmendmentsCreditsLedgerTypeCreditAutomatedInvoiceDeduction, ContractListResponseDataAmendmentsCreditsLedgerTypeCreditExpiration, ContractListResponseDataAmendmentsCreditsLedgerTypeCreditCanceled, ContractListResponseDataAmendmentsCreditsLedgerTypeCreditCredited, ContractListResponseDataAmendmentsCreditsLedgerTypeCreditManual:
		return true
	}
	return false
}

type ContractListResponseDataAmendmentsProfessionalService struct {
	ID string `json:"id,required" format:"uuid"`
	// Maximum amount for the term.
	MaxAmount float64 `json:"max_amount,required"`
	ProductID string  `json:"product_id,required" format:"uuid"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity float64 `json:"quantity,required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified.
	UnitPrice    float64           `json:"unit_price,required"`
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string                                                    `json:"netsuite_sales_order_id"`
	JSON                 contractListResponseDataAmendmentsProfessionalServiceJSON `json:"-"`
}

// contractListResponseDataAmendmentsProfessionalServiceJSON contains the JSON
// metadata for the struct [ContractListResponseDataAmendmentsProfessionalService]
type contractListResponseDataAmendmentsProfessionalServiceJSON struct {
	ID                   apijson.Field
	MaxAmount            apijson.Field
	ProductID            apijson.Field
	Quantity             apijson.Field
	UnitPrice            apijson.Field
	CustomFields         apijson.Field
	Description          apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsProfessionalService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListResponseDataAmendmentsProfessionalServiceJSON) RawJSON() string {
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

// This field's availability is dependent on your client's configuration.
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
	// This field can have the runtime type of [shared.CommitContract],
	// [ContractListBalancesResponseDataCreditContract].
	Contract interface{}                          `json:"contract,required"`
	Type     ContractListBalancesResponseDataType `json:"type,required"`
	Name     string                               `json:"name"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// This field can have the runtime type of [shared.CommitProduct],
	// [ContractListBalancesResponseDataCreditProduct].
	Product interface{} `json:"product"`
	// This field can have the runtime type of [shared.CommitAccessSchedule],
	// [ContractListBalancesResponseDataCreditAccessSchedule].
	AccessSchedule interface{} `json:"access_schedule,required"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule shared.SchedulePointInTime `json:"invoice_schedule"`
	// This field can have the runtime type of [shared.CommitInvoiceContract].
	InvoiceContract interface{} `json:"invoice_contract,required"`
	// This field can have the runtime type of [shared.CommitRolledOverFrom].
	RolledOverFrom   interface{} `json:"rolled_over_from,required"`
	Description      string      `json:"description"`
	RolloverFraction float64     `json:"rollover_fraction"`
	// This field can have the runtime type of [[]string].
	ApplicableProductIDs interface{} `json:"applicable_product_ids,required"`
	// This field can have the runtime type of [[]string].
	ApplicableProductTags interface{} `json:"applicable_product_tags,required"`
	// This field can have the runtime type of [[]string].
	ApplicableContractIDs interface{} `json:"applicable_contract_ids,required"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// (DEPRECATED) Use access_schedule + invoice_schedule instead.
	Amount float64 `json:"amount"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// This field can have the runtime type of [[]shared.CommitLedger],
	// [[]ContractListBalancesResponseDataCreditLedger].
	Ledger interface{} `json:"ledger,required"`
	// This field can have the runtime type of [map[string]string].
	CustomFields interface{}                          `json:"custom_fields,required"`
	JSON         contractListBalancesResponseDataJSON `json:"-"`
	union        ContractListBalancesResponseDataUnion
}

// contractListBalancesResponseDataJSON contains the JSON metadata for the struct
// [ContractListBalancesResponseData]
type contractListBalancesResponseDataJSON struct {
	ID                      apijson.Field
	Contract                apijson.Field
	Type                    apijson.Field
	Name                    apijson.Field
	Priority                apijson.Field
	Product                 apijson.Field
	AccessSchedule          apijson.Field
	InvoiceSchedule         apijson.Field
	InvoiceContract         apijson.Field
	RolledOverFrom          apijson.Field
	Description             apijson.Field
	RolloverFraction        apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	ApplicableContractIDs   apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Amount                  apijson.Field
	SalesforceOpportunityID apijson.Field
	Ledger                  apijson.Field
	CustomFields            apijson.Field
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
// Possible runtime types of the union are [shared.Commit],
// [ContractListBalancesResponseDataCredit].
func (r ContractListBalancesResponseData) AsUnion() ContractListBalancesResponseDataUnion {
	return r.union
}

// Union satisfied by [shared.Commit] or [ContractListBalancesResponseDataCredit].
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
			Type:       reflect.TypeOf(ContractListBalancesResponseDataCredit{}),
		},
	)
}

type ContractListBalancesResponseDataCredit struct {
	ID      string                                        `json:"id,required" format:"uuid"`
	Product ContractListBalancesResponseDataCreditProduct `json:"product,required"`
	Type    ContractListBalancesResponseDataCreditType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        ContractListBalancesResponseDataCreditAccessSchedule `json:"access_schedule"`
	ApplicableContractIDs []string                                             `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string                                             `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                                             `json:"applicable_product_tags"`
	Contract              ContractListBalancesResponseDataCreditContract       `json:"contract"`
	CustomFields          map[string]string                                    `json:"custom_fields"`
	Description           string                                               `json:"description"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []ContractListBalancesResponseDataCreditLedger `json:"ledger"`
	Name   string                                         `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                                     `json:"salesforce_opportunity_id"`
	JSON                    contractListBalancesResponseDataCreditJSON `json:"-"`
}

// contractListBalancesResponseDataCreditJSON contains the JSON metadata for the
// struct [ContractListBalancesResponseDataCredit]
type contractListBalancesResponseDataCreditJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractListBalancesResponseDataCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseDataCreditJSON) RawJSON() string {
	return r.raw
}

func (r ContractListBalancesResponseDataCredit) ImplementsContractListBalancesResponseData() {}

type ContractListBalancesResponseDataCreditProduct struct {
	ID   string                                            `json:"id,required" format:"uuid"`
	Name string                                            `json:"name,required"`
	JSON contractListBalancesResponseDataCreditProductJSON `json:"-"`
}

// contractListBalancesResponseDataCreditProductJSON contains the JSON metadata for
// the struct [ContractListBalancesResponseDataCreditProduct]
type contractListBalancesResponseDataCreditProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListBalancesResponseDataCreditProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseDataCreditProductJSON) RawJSON() string {
	return r.raw
}

type ContractListBalancesResponseDataCreditType string

const (
	ContractListBalancesResponseDataCreditTypeCredit ContractListBalancesResponseDataCreditType = "CREDIT"
)

func (r ContractListBalancesResponseDataCreditType) IsKnown() bool {
	switch r {
	case ContractListBalancesResponseDataCreditTypeCredit:
		return true
	}
	return false
}

// The schedule that the customer will gain access to the credits.
type ContractListBalancesResponseDataCreditAccessSchedule struct {
	ScheduleItems []ContractListBalancesResponseDataCreditAccessScheduleScheduleItem `json:"schedule_items,required"`
	CreditType    shared.CreditType                                                  `json:"credit_type"`
	JSON          contractListBalancesResponseDataCreditAccessScheduleJSON           `json:"-"`
}

// contractListBalancesResponseDataCreditAccessScheduleJSON contains the JSON
// metadata for the struct [ContractListBalancesResponseDataCreditAccessSchedule]
type contractListBalancesResponseDataCreditAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListBalancesResponseDataCreditAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseDataCreditAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type ContractListBalancesResponseDataCreditAccessScheduleScheduleItem struct {
	ID           string                                                               `json:"id,required" format:"uuid"`
	Amount       float64                                                              `json:"amount,required"`
	EndingBefore time.Time                                                            `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                                                            `json:"starting_at,required" format:"date-time"`
	JSON         contractListBalancesResponseDataCreditAccessScheduleScheduleItemJSON `json:"-"`
}

// contractListBalancesResponseDataCreditAccessScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractListBalancesResponseDataCreditAccessScheduleScheduleItem]
type contractListBalancesResponseDataCreditAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractListBalancesResponseDataCreditAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseDataCreditAccessScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type ContractListBalancesResponseDataCreditContract struct {
	ID   string                                             `json:"id,required" format:"uuid"`
	JSON contractListBalancesResponseDataCreditContractJSON `json:"-"`
}

// contractListBalancesResponseDataCreditContractJSON contains the JSON metadata
// for the struct [ContractListBalancesResponseDataCreditContract]
type contractListBalancesResponseDataCreditContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListBalancesResponseDataCreditContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseDataCreditContractJSON) RawJSON() string {
	return r.raw
}

type ContractListBalancesResponseDataCreditLedger struct {
	Type      ContractListBalancesResponseDataCreditLedgerType `json:"type,required"`
	Timestamp time.Time                                        `json:"timestamp,required" format:"date-time"`
	Amount    float64                                          `json:"amount,required"`
	SegmentID string                                           `json:"segment_id" format:"uuid"`
	InvoiceID string                                           `json:"invoice_id" format:"uuid"`
	Reason    string                                           `json:"reason"`
	JSON      contractListBalancesResponseDataCreditLedgerJSON `json:"-"`
	union     ContractListBalancesResponseDataCreditLedgerUnion
}

// contractListBalancesResponseDataCreditLedgerJSON contains the JSON metadata for
// the struct [ContractListBalancesResponseDataCreditLedger]
type contractListBalancesResponseDataCreditLedgerJSON struct {
	Type        apijson.Field
	Timestamp   apijson.Field
	Amount      apijson.Field
	SegmentID   apijson.Field
	InvoiceID   apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r contractListBalancesResponseDataCreditLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *ContractListBalancesResponseDataCreditLedger) UnmarshalJSON(data []byte) (err error) {
	*r = ContractListBalancesResponseDataCreditLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContractListBalancesResponseDataCreditLedgerUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntry],
// [ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntry],
// [ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntry],
// [ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntry],
// [ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntry].
func (r ContractListBalancesResponseDataCreditLedger) AsUnion() ContractListBalancesResponseDataCreditLedgerUnion {
	return r.union
}

// Union satisfied by
// [ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntry],
// [ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntry],
// [ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntry],
// [ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntry] or
// [ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntry].
type ContractListBalancesResponseDataCreditLedgerUnion interface {
	implementsContractListBalancesResponseDataCreditLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContractListBalancesResponseDataCreditLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntry{}),
		},
	)
}

type ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64                                                                       `json:"amount,required"`
	SegmentID string                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                     `json:"timestamp,required" format:"date-time"`
	Type      ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntryJSON `json:"-"`
}

// contractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntry]
type contractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntry) implementsContractListBalancesResponseDataCreditLedger() {
}

type ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntryType string

const (
	ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntryType = "CREDIT_SEGMENT_START"
)

func (r ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListBalancesResponseDataCreditLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart:
		return true
	}
	return false
}

type ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                    `json:"amount,required"`
	InvoiceID string                                                                                     `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                  `json:"timestamp,required" format:"date-time"`
	Type      ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// contractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry]
type contractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) implementsContractListBalancesResponseDataCreditLedger() {
}

type ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListBalancesResponseDataCreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntry struct {
	Amount    float64                                                                     `json:"amount,required"`
	SegmentID string                                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                   `json:"timestamp,required" format:"date-time"`
	Type      ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntryType `json:"type,required"`
	JSON      contractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntryJSON `json:"-"`
}

// contractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntry]
type contractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntry) implementsContractListBalancesResponseDataCreditLedger() {
}

type ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntryType string

const (
	ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntryTypeCreditExpiration ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntryType = "CREDIT_EXPIRATION"
)

func (r ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListBalancesResponseDataCreditLedgerCreditExpirationLedgerEntryTypeCreditExpiration:
		return true
	}
	return false
}

type ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntry struct {
	Amount    float64                                                                   `json:"amount,required"`
	InvoiceID string                                                                    `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                 `json:"timestamp,required" format:"date-time"`
	Type      ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntryType `json:"type,required"`
	JSON      contractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntryJSON `json:"-"`
}

// contractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntry]
type contractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntry) implementsContractListBalancesResponseDataCreditLedger() {
}

type ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntryType string

const (
	ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntryTypeCreditCanceled ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntryType = "CREDIT_CANCELED"
)

func (r ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListBalancesResponseDataCreditLedgerCreditCanceledLedgerEntryTypeCreditCanceled:
		return true
	}
	return false
}

type ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntry struct {
	Amount    float64                                                                   `json:"amount,required"`
	InvoiceID string                                                                    `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                 `json:"timestamp,required" format:"date-time"`
	Type      ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntryType `json:"type,required"`
	JSON      contractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntryJSON `json:"-"`
}

// contractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntry]
type contractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntry) implementsContractListBalancesResponseDataCreditLedger() {
}

type ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntryType string

const (
	ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntryTypeCreditCredited ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntryType = "CREDIT_CREDITED"
)

func (r ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListBalancesResponseDataCreditLedgerCreditCreditedLedgerEntryTypeCreditCredited:
		return true
	}
	return false
}

type ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntry struct {
	Amount    float64                                                                 `json:"amount,required"`
	Reason    string                                                                  `json:"reason,required"`
	Timestamp time.Time                                                               `json:"timestamp,required" format:"date-time"`
	Type      ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntryType `json:"type,required"`
	JSON      contractListBalancesResponseDataCreditLedgerCreditManualLedgerEntryJSON `json:"-"`
}

// contractListBalancesResponseDataCreditLedgerCreditManualLedgerEntryJSON contains
// the JSON metadata for the struct
// [ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntry]
type contractListBalancesResponseDataCreditLedgerCreditManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractListBalancesResponseDataCreditLedgerCreditManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntry) implementsContractListBalancesResponseDataCreditLedger() {
}

type ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntryType string

const (
	ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntryTypeCreditManual ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntryType = "CREDIT_MANUAL"
)

func (r ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntryType) IsKnown() bool {
	switch r {
	case ContractListBalancesResponseDataCreditLedgerCreditManualLedgerEntryTypeCreditManual:
		return true
	}
	return false
}

type ContractListBalancesResponseDataCreditLedgerType string

const (
	ContractListBalancesResponseDataCreditLedgerTypeCreditSegmentStart              ContractListBalancesResponseDataCreditLedgerType = "CREDIT_SEGMENT_START"
	ContractListBalancesResponseDataCreditLedgerTypeCreditAutomatedInvoiceDeduction ContractListBalancesResponseDataCreditLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	ContractListBalancesResponseDataCreditLedgerTypeCreditExpiration                ContractListBalancesResponseDataCreditLedgerType = "CREDIT_EXPIRATION"
	ContractListBalancesResponseDataCreditLedgerTypeCreditCanceled                  ContractListBalancesResponseDataCreditLedgerType = "CREDIT_CANCELED"
	ContractListBalancesResponseDataCreditLedgerTypeCreditCredited                  ContractListBalancesResponseDataCreditLedgerType = "CREDIT_CREDITED"
	ContractListBalancesResponseDataCreditLedgerTypeCreditManual                    ContractListBalancesResponseDataCreditLedgerType = "CREDIT_MANUAL"
)

func (r ContractListBalancesResponseDataCreditLedgerType) IsKnown() bool {
	switch r {
	case ContractListBalancesResponseDataCreditLedgerTypeCreditSegmentStart, ContractListBalancesResponseDataCreditLedgerTypeCreditAutomatedInvoiceDeduction, ContractListBalancesResponseDataCreditLedgerTypeCreditExpiration, ContractListBalancesResponseDataCreditLedgerTypeCreditCanceled, ContractListBalancesResponseDataCreditLedgerTypeCreditCredited, ContractListBalancesResponseDataCreditLedgerTypeCreditManual:
		return true
	}
	return false
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
	Entitled            bool                                    `json:"entitled,required"`
	ListRate            shared.Rate                             `json:"list_rate,required"`
	ProductCustomFields map[string]string                       `json:"product_custom_fields,required"`
	ProductID           string                                  `json:"product_id,required" format:"uuid"`
	ProductName         string                                  `json:"product_name,required"`
	ProductTags         []string                                `json:"product_tags,required"`
	RateCardID          string                                  `json:"rate_card_id,required" format:"uuid"`
	StartingAt          time.Time                               `json:"starting_at,required" format:"date-time"`
	EndingBefore        time.Time                               `json:"ending_before" format:"date-time"`
	OverrideRate        shared.Rate                             `json:"override_rate"`
	PricingGroupValues  map[string]string                       `json:"pricing_group_values"`
	JSON                contractGetRateScheduleResponseDataJSON `json:"-"`
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
	// This field's availability is dependent on your client's configuration.
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
	RateCardAlias param.Field[string] `json:"rate_card_alias"`
	RateCardID    param.Field[string] `json:"rate_card_id" format:"uuid"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties param.Field[[]ContractNewParamsResellerRoyalty] `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string]                             `json:"salesforce_opportunity_id"`
	ScheduledCharges        param.Field[[]ContractNewParamsScheduledCharge] `json:"scheduled_charges"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue param.Field[float64]                     `json:"total_contract_value"`
	Transition         param.Field[ContractNewParamsTransition] `json:"transition"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey          param.Field[string]                                  `json:"uniqueness_key"`
	UsageFilter            param.Field[ContractNewParamsUsageFilter]            `json:"usage_filter"`
	UsageStatementSchedule param.Field[ContractNewParamsUsageStatementSchedule] `json:"usage_statement_schedule"`
}

func (r ContractNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This field's availability is dependent on your client's configuration.
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
	Priority param.Field[float64] `json:"priority"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
}

func (r ContractNewParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsCommitsType string

const (
	ContractNewParamsCommitsTypePrepaid  ContractNewParamsCommitsType = "PREPAID"
	ContractNewParamsCommitsTypePrepaid  ContractNewParamsCommitsType = "prepaid"
	ContractNewParamsCommitsTypePostpaid ContractNewParamsCommitsType = "POSTPAID"
	ContractNewParamsCommitsTypePostpaid ContractNewParamsCommitsType = "postpaid"
)

func (r ContractNewParamsCommitsType) IsKnown() bool {
	switch r {
	case ContractNewParamsCommitsTypePrepaid, ContractNewParamsCommitsTypePrepaid, ContractNewParamsCommitsTypePostpaid, ContractNewParamsCommitsTypePostpaid:
		return true
	}
	return false
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
type ContractNewParamsCommitsAccessSchedule struct {
	ScheduleItems param.Field[[]ContractNewParamsCommitsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	CreditTypeID  param.Field[string]                                               `json:"credit_type_id" format:"uuid"`
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
	// Defaults to USD if not passed. Only USD is supported at this time.
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
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "each"
)

func (r ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "monthly"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "quarterly"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "semi_annual"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "annual"
)

func (r ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual, ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual:
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
	Priority param.Field[float64] `json:"priority"`
}

func (r ContractNewParamsCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the credit to the customer.
type ContractNewParamsCreditsAccessSchedule struct {
	ScheduleItems param.Field[[]ContractNewParamsCreditsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	CreditTypeID  param.Field[string]                                               `json:"credit_type_id" format:"uuid"`
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

type ContractNewParamsDiscount struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule param.Field[ContractNewParamsDiscountsSchedule] `json:"schedule,required"`
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
	// Defaults to USD if not passed. Only USD is supported at this time.
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
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided"
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "each"
)

func (r ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided, ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided, ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded, ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded, ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach, ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractNewParamsDiscountsScheduleRecurringScheduleFrequency string

const (
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "monthly"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "quarterly"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "semi_annual"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "annual"
)

func (r ContractNewParamsDiscountsScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly, ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly, ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly, ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly, ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual, ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual, ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual, ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual:
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
	ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier ContractNewParamsMultiplierOverridePrioritization = "lowest_multiplier"
	ContractNewParamsMultiplierOverridePrioritizationExplicit         ContractNewParamsMultiplierOverridePrioritization = "EXPLICIT"
	ContractNewParamsMultiplierOverridePrioritizationExplicit         ContractNewParamsMultiplierOverridePrioritization = "explicit"
)

func (r ContractNewParamsMultiplierOverridePrioritization) IsKnown() bool {
	switch r {
	case ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier, ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier, ContractNewParamsMultiplierOverridePrioritizationExplicit, ContractNewParamsMultiplierOverridePrioritizationExplicit:
		return true
	}
	return false
}

type ContractNewParamsOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// tags identifying products whose rates are being overridden
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Entitled     param.Field[bool]      `json:"entitled"`
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
	// ID of the product whose rate is being overridden
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// Required for TIERED type. Must have at least one tier.
	Tiers param.Field[[]ContractNewParamsOverridesTier] `json:"tiers"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type param.Field[ContractNewParamsOverridesType] `json:"type"`
}

func (r ContractNewParamsOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsOverridesOverrideSpecifier struct {
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
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type.
	IsProrated param.Field[bool] `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price param.Field[float64] `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Field[float64] `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]ContractNewParamsOverridesOverwriteRateTier] `json:"tiers"`
}

func (r ContractNewParamsOverridesOverwriteRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsOverridesOverwriteRateRateType string

const (
	ContractNewParamsOverridesOverwriteRateRateTypeFlat         ContractNewParamsOverridesOverwriteRateRateType = "FLAT"
	ContractNewParamsOverridesOverwriteRateRateTypeFlat         ContractNewParamsOverridesOverwriteRateRateType = "flat"
	ContractNewParamsOverridesOverwriteRateRateTypePercentage   ContractNewParamsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractNewParamsOverridesOverwriteRateRateTypePercentage   ContractNewParamsOverridesOverwriteRateRateType = "percentage"
	ContractNewParamsOverridesOverwriteRateRateTypeSubscription ContractNewParamsOverridesOverwriteRateRateType = "SUBSCRIPTION"
	ContractNewParamsOverridesOverwriteRateRateTypeSubscription ContractNewParamsOverridesOverwriteRateRateType = "subscription"
	ContractNewParamsOverridesOverwriteRateRateTypeTiered       ContractNewParamsOverridesOverwriteRateRateType = "TIERED"
	ContractNewParamsOverridesOverwriteRateRateTypeTiered       ContractNewParamsOverridesOverwriteRateRateType = "tiered"
	ContractNewParamsOverridesOverwriteRateRateTypeCustom       ContractNewParamsOverridesOverwriteRateRateType = "CUSTOM"
	ContractNewParamsOverridesOverwriteRateRateTypeCustom       ContractNewParamsOverridesOverwriteRateRateType = "custom"
)

func (r ContractNewParamsOverridesOverwriteRateRateType) IsKnown() bool {
	switch r {
	case ContractNewParamsOverridesOverwriteRateRateTypeFlat, ContractNewParamsOverridesOverwriteRateRateTypeFlat, ContractNewParamsOverridesOverwriteRateRateTypePercentage, ContractNewParamsOverridesOverwriteRateRateTypePercentage, ContractNewParamsOverridesOverwriteRateRateTypeSubscription, ContractNewParamsOverridesOverwriteRateRateTypeSubscription, ContractNewParamsOverridesOverwriteRateRateTypeTiered, ContractNewParamsOverridesOverwriteRateRateTypeTiered, ContractNewParamsOverridesOverwriteRateRateTypeCustom, ContractNewParamsOverridesOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

type ContractNewParamsOverridesOverwriteRateTier struct {
	Price param.Field[float64] `json:"price,required"`
	Size  param.Field[float64] `json:"size"`
}

func (r ContractNewParamsOverridesOverwriteRateTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	ContractNewParamsOverridesTypeOverwrite  ContractNewParamsOverridesType = "overwrite"
	ContractNewParamsOverridesTypeMultiplier ContractNewParamsOverridesType = "MULTIPLIER"
	ContractNewParamsOverridesTypeMultiplier ContractNewParamsOverridesType = "multiplier"
	ContractNewParamsOverridesTypeTiered     ContractNewParamsOverridesType = "TIERED"
	ContractNewParamsOverridesTypeTiered     ContractNewParamsOverridesType = "tiered"
)

func (r ContractNewParamsOverridesType) IsKnown() bool {
	switch r {
	case ContractNewParamsOverridesTypeOverwrite, ContractNewParamsOverridesTypeOverwrite, ContractNewParamsOverridesTypeMultiplier, ContractNewParamsOverridesTypeMultiplier, ContractNewParamsOverridesTypeTiered, ContractNewParamsOverridesTypeTiered:
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
	// Defaults to USD if not passed. Only USD is supported at this time.
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
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "each"
)

func (r ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided, ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided, ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded, ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded, ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach, ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "monthly"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "quarterly"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "semi_annual"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "annual"
)

func (r ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly, ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly, ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly, ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly, ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual, ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual, ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual, ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual:
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
	ContractNewParamsTransitionTypeSupersede ContractNewParamsTransitionType = "supersede"
	ContractNewParamsTransitionTypeRenewal   ContractNewParamsTransitionType = "renewal"
)

func (r ContractNewParamsTransitionType) IsKnown() bool {
	switch r {
	case ContractNewParamsTransitionTypeSupersede, ContractNewParamsTransitionTypeRenewal, ContractNewParamsTransitionTypeSupersede, ContractNewParamsTransitionTypeRenewal:
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
	ContractNewParamsTransitionFutureInvoiceBehaviorTrueupRemove ContractNewParamsTransitionFutureInvoiceBehaviorTrueup = "remove"
	ContractNewParamsTransitionFutureInvoiceBehaviorTrueupAsIs   ContractNewParamsTransitionFutureInvoiceBehaviorTrueup = "as_is"
	ContractNewParamsTransitionFutureInvoiceBehaviorTrueupRemove ContractNewParamsTransitionFutureInvoiceBehaviorTrueup = "REMOVE"
	ContractNewParamsTransitionFutureInvoiceBehaviorTrueupAsIs   ContractNewParamsTransitionFutureInvoiceBehaviorTrueup = "AS_IS"
)

func (r ContractNewParamsTransitionFutureInvoiceBehaviorTrueup) IsKnown() bool {
	switch r {
	case ContractNewParamsTransitionFutureInvoiceBehaviorTrueupRemove, ContractNewParamsTransitionFutureInvoiceBehaviorTrueupAsIs, ContractNewParamsTransitionFutureInvoiceBehaviorTrueupRemove, ContractNewParamsTransitionFutureInvoiceBehaviorTrueupAsIs:
		return true
	}
	return false
}

type ContractNewParamsUsageFilter struct {
	GroupKey    param.Field[string]    `json:"group_key,required"`
	GroupValues param.Field[[]string]  `json:"group_values,required"`
	StartingAt  param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r ContractNewParamsUsageFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsUsageStatementSchedule struct {
	Frequency param.Field[ContractNewParamsUsageStatementScheduleFrequency] `json:"frequency,required"`
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
	ContractNewParamsUsageStatementScheduleFrequencyMonthly   ContractNewParamsUsageStatementScheduleFrequency = "monthly"
	ContractNewParamsUsageStatementScheduleFrequencyQuarterly ContractNewParamsUsageStatementScheduleFrequency = "QUARTERLY"
	ContractNewParamsUsageStatementScheduleFrequencyQuarterly ContractNewParamsUsageStatementScheduleFrequency = "quarterly"
)

func (r ContractNewParamsUsageStatementScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractNewParamsUsageStatementScheduleFrequencyMonthly, ContractNewParamsUsageStatementScheduleFrequencyMonthly, ContractNewParamsUsageStatementScheduleFrequencyQuarterly, ContractNewParamsUsageStatementScheduleFrequencyQuarterly:
		return true
	}
	return false
}

// If not provided, defaults to the first day of the month.
type ContractNewParamsUsageStatementScheduleDay string

const (
	ContractNewParamsUsageStatementScheduleDayFirstOfMonth  ContractNewParamsUsageStatementScheduleDay = "FIRST_OF_MONTH"
	ContractNewParamsUsageStatementScheduleDayFirstOfMonth  ContractNewParamsUsageStatementScheduleDay = "first_of_month"
	ContractNewParamsUsageStatementScheduleDayContractStart ContractNewParamsUsageStatementScheduleDay = "CONTRACT_START"
	ContractNewParamsUsageStatementScheduleDayContractStart ContractNewParamsUsageStatementScheduleDay = "contract_start"
)

func (r ContractNewParamsUsageStatementScheduleDay) IsKnown() bool {
	switch r {
	case ContractNewParamsUsageStatementScheduleDayFirstOfMonth, ContractNewParamsUsageStatementScheduleDayFirstOfMonth, ContractNewParamsUsageStatementScheduleDayContractStart, ContractNewParamsUsageStatementScheduleDayContractStart:
		return true
	}
	return false
}

type ContractGetParams struct {
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
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
	Priority param.Field[float64] `json:"priority"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
}

func (r ContractAmendParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommitsType string

const (
	ContractAmendParamsCommitsTypePrepaid  ContractAmendParamsCommitsType = "PREPAID"
	ContractAmendParamsCommitsTypePrepaid  ContractAmendParamsCommitsType = "prepaid"
	ContractAmendParamsCommitsTypePostpaid ContractAmendParamsCommitsType = "POSTPAID"
	ContractAmendParamsCommitsTypePostpaid ContractAmendParamsCommitsType = "postpaid"
)

func (r ContractAmendParamsCommitsType) IsKnown() bool {
	switch r {
	case ContractAmendParamsCommitsTypePrepaid, ContractAmendParamsCommitsTypePrepaid, ContractAmendParamsCommitsTypePostpaid, ContractAmendParamsCommitsTypePostpaid:
		return true
	}
	return false
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
type ContractAmendParamsCommitsAccessSchedule struct {
	ScheduleItems param.Field[[]ContractAmendParamsCommitsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	CreditTypeID  param.Field[string]                                                 `json:"credit_type_id" format:"uuid"`
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
	// Defaults to USD if not passed. Only USD is supported at this time.
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
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "each"
)

func (r ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "monthly"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "quarterly"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "semi_annual"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "annual"
)

func (r ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual, ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual:
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
	Priority param.Field[float64] `json:"priority"`
}

func (r ContractAmendParamsCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the credit to the customer.
type ContractAmendParamsCreditsAccessSchedule struct {
	ScheduleItems param.Field[[]ContractAmendParamsCreditsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	CreditTypeID  param.Field[string]                                                 `json:"credit_type_id" format:"uuid"`
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

type ContractAmendParamsDiscount struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule param.Field[ContractAmendParamsDiscountsSchedule] `json:"schedule,required"`
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
	// Defaults to USD if not passed. Only USD is supported at this time.
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
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided"
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "each"
)

func (r ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided, ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided, ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded, ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded, ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach, ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency string

const (
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "monthly"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "quarterly"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "semi_annual"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "annual"
)

func (r ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly, ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly, ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly, ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly, ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual, ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual, ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual, ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual:
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
	// tags identifying products whose rates are being overridden
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Entitled     param.Field[bool]      `json:"entitled"`
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
	// ID of the product whose rate is being overridden
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// Required for TIERED type. Must have at least one tier.
	Tiers param.Field[[]ContractAmendParamsOverridesTier] `json:"tiers"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type param.Field[ContractAmendParamsOverridesType] `json:"type"`
}

func (r ContractAmendParamsOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsOverridesOverrideSpecifier struct {
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
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type.
	IsProrated param.Field[bool] `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price param.Field[float64] `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Field[float64] `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]ContractAmendParamsOverridesOverwriteRateTier] `json:"tiers"`
}

func (r ContractAmendParamsOverridesOverwriteRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsOverridesOverwriteRateRateType string

const (
	ContractAmendParamsOverridesOverwriteRateRateTypeFlat         ContractAmendParamsOverridesOverwriteRateRateType = "FLAT"
	ContractAmendParamsOverridesOverwriteRateRateTypeFlat         ContractAmendParamsOverridesOverwriteRateRateType = "flat"
	ContractAmendParamsOverridesOverwriteRateRateTypePercentage   ContractAmendParamsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractAmendParamsOverridesOverwriteRateRateTypePercentage   ContractAmendParamsOverridesOverwriteRateRateType = "percentage"
	ContractAmendParamsOverridesOverwriteRateRateTypeSubscription ContractAmendParamsOverridesOverwriteRateRateType = "SUBSCRIPTION"
	ContractAmendParamsOverridesOverwriteRateRateTypeSubscription ContractAmendParamsOverridesOverwriteRateRateType = "subscription"
	ContractAmendParamsOverridesOverwriteRateRateTypeTiered       ContractAmendParamsOverridesOverwriteRateRateType = "TIERED"
	ContractAmendParamsOverridesOverwriteRateRateTypeTiered       ContractAmendParamsOverridesOverwriteRateRateType = "tiered"
	ContractAmendParamsOverridesOverwriteRateRateTypeCustom       ContractAmendParamsOverridesOverwriteRateRateType = "CUSTOM"
	ContractAmendParamsOverridesOverwriteRateRateTypeCustom       ContractAmendParamsOverridesOverwriteRateRateType = "custom"
)

func (r ContractAmendParamsOverridesOverwriteRateRateType) IsKnown() bool {
	switch r {
	case ContractAmendParamsOverridesOverwriteRateRateTypeFlat, ContractAmendParamsOverridesOverwriteRateRateTypeFlat, ContractAmendParamsOverridesOverwriteRateRateTypePercentage, ContractAmendParamsOverridesOverwriteRateRateTypePercentage, ContractAmendParamsOverridesOverwriteRateRateTypeSubscription, ContractAmendParamsOverridesOverwriteRateRateTypeSubscription, ContractAmendParamsOverridesOverwriteRateRateTypeTiered, ContractAmendParamsOverridesOverwriteRateRateTypeTiered, ContractAmendParamsOverridesOverwriteRateRateTypeCustom, ContractAmendParamsOverridesOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

type ContractAmendParamsOverridesOverwriteRateTier struct {
	Price param.Field[float64] `json:"price,required"`
	Size  param.Field[float64] `json:"size"`
}

func (r ContractAmendParamsOverridesOverwriteRateTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	ContractAmendParamsOverridesTypeOverwrite  ContractAmendParamsOverridesType = "overwrite"
	ContractAmendParamsOverridesTypeMultiplier ContractAmendParamsOverridesType = "MULTIPLIER"
	ContractAmendParamsOverridesTypeMultiplier ContractAmendParamsOverridesType = "multiplier"
	ContractAmendParamsOverridesTypeTiered     ContractAmendParamsOverridesType = "TIERED"
	ContractAmendParamsOverridesTypeTiered     ContractAmendParamsOverridesType = "tiered"
)

func (r ContractAmendParamsOverridesType) IsKnown() bool {
	switch r {
	case ContractAmendParamsOverridesTypeOverwrite, ContractAmendParamsOverridesTypeOverwrite, ContractAmendParamsOverridesTypeMultiplier, ContractAmendParamsOverridesTypeMultiplier, ContractAmendParamsOverridesTypeTiered, ContractAmendParamsOverridesTypeTiered:
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
	// Defaults to USD if not passed. Only USD is supported at this time.
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
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "each"
)

func (r ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided, ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided, ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded, ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded, ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach, ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "monthly"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "quarterly"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "semi_annual"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "annual"
)

func (r ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly, ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly, ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly, ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly, ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual, ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual, ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual, ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual:
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

type ContractListBalancesParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	ID         param.Field[string] `json:"id" format:"uuid"`
	// Return only balances that have access schedules that "cover" the provided date
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	// Include only balances that have any access before the provided date (exclusive)
	EffectiveBefore param.Field[time.Time] `json:"effective_before" format:"date-time"`
	// Include credits from archived contracts.
	IncludeArchived param.Field[bool] `json:"include_archived"`
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
	// RFC 3339 timestamp indicating when the contract will end (exclusive). If not
	// provided, the contract will be updated to be open-ended.
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
}

func (r ContractUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
