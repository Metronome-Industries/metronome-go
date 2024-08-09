// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
	"github.com/tidwall/gjson"
)

// CustomerCreditService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerCreditService] method instead.
type CustomerCreditService struct {
	Options []option.RequestOption
}

// NewCustomerCreditService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerCreditService(opts ...option.RequestOption) (r *CustomerCreditService) {
	r = &CustomerCreditService{}
	r.Options = opts
	return
}

// Create a new credit at the customer level.
func (r *CustomerCreditService) New(ctx context.Context, body CustomerCreditNewParams, opts ...option.RequestOption) (res *CustomerCreditNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/customerCredits/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List credits.
func (r *CustomerCreditService) List(ctx context.Context, body CustomerCreditListParams, opts ...option.RequestOption) (res *CustomerCreditListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/customerCredits/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the end date of a credit
func (r *CustomerCreditService) UpdateEndDate(ctx context.Context, body CustomerCreditUpdateEndDateParams, opts ...option.RequestOption) (res *CustomerCreditUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/customerCredits/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CustomerCreditNewResponse struct {
	Data shared.ID                     `json:"data,required"`
	JSON customerCreditNewResponseJSON `json:"-"`
}

// customerCreditNewResponseJSON contains the JSON metadata for the struct
// [CustomerCreditNewResponse]
type customerCreditNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditNewResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListResponse struct {
	Data     []CustomerCreditListResponseData `json:"data,required"`
	NextPage string                           `json:"next_page,required,nullable"`
	JSON     customerCreditListResponseJSON   `json:"-"`
}

// customerCreditListResponseJSON contains the JSON metadata for the struct
// [CustomerCreditListResponse]
type customerCreditListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListResponseData struct {
	ID      string                                `json:"id,required" format:"uuid"`
	Product CustomerCreditListResponseDataProduct `json:"product,required"`
	Type    CustomerCreditListResponseDataType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        CustomerCreditListResponseDataAccessSchedule `json:"access_schedule"`
	ApplicableContractIDs []string                                     `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string                                     `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                                     `json:"applicable_product_tags"`
	Contract              CustomerCreditListResponseDataContract       `json:"contract"`
	CustomFields          map[string]string                            `json:"custom_fields"`
	Description           string                                       `json:"description"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []CustomerCreditListResponseDataLedger `json:"ledger"`
	Name   string                                 `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                             `json:"salesforce_opportunity_id"`
	JSON                    customerCreditListResponseDataJSON `json:"-"`
}

// customerCreditListResponseDataJSON contains the JSON metadata for the struct
// [CustomerCreditListResponseData]
type customerCreditListResponseDataJSON struct {
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

func (r *CustomerCreditListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseDataJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListResponseDataProduct struct {
	ID   string                                    `json:"id,required" format:"uuid"`
	Name string                                    `json:"name,required"`
	JSON customerCreditListResponseDataProductJSON `json:"-"`
}

// customerCreditListResponseDataProductJSON contains the JSON metadata for the
// struct [CustomerCreditListResponseDataProduct]
type customerCreditListResponseDataProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponseDataProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseDataProductJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListResponseDataType string

const (
	CustomerCreditListResponseDataTypeCredit CustomerCreditListResponseDataType = "CREDIT"
)

func (r CustomerCreditListResponseDataType) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseDataTypeCredit:
		return true
	}
	return false
}

// The schedule that the customer will gain access to the credits.
type CustomerCreditListResponseDataAccessSchedule struct {
	ScheduleItems []CustomerCreditListResponseDataAccessScheduleScheduleItem `json:"schedule_items,required"`
	CreditType    shared.CreditType                                          `json:"credit_type"`
	JSON          customerCreditListResponseDataAccessScheduleJSON           `json:"-"`
}

// customerCreditListResponseDataAccessScheduleJSON contains the JSON metadata for
// the struct [CustomerCreditListResponseDataAccessSchedule]
type customerCreditListResponseDataAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomerCreditListResponseDataAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseDataAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListResponseDataAccessScheduleScheduleItem struct {
	ID           string                                                       `json:"id,required" format:"uuid"`
	Amount       float64                                                      `json:"amount,required"`
	EndingBefore time.Time                                                    `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                                                    `json:"starting_at,required" format:"date-time"`
	JSON         customerCreditListResponseDataAccessScheduleScheduleItemJSON `json:"-"`
}

// customerCreditListResponseDataAccessScheduleScheduleItemJSON contains the JSON
// metadata for the struct
// [CustomerCreditListResponseDataAccessScheduleScheduleItem]
type customerCreditListResponseDataAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomerCreditListResponseDataAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseDataAccessScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListResponseDataContract struct {
	ID   string                                     `json:"id,required" format:"uuid"`
	JSON customerCreditListResponseDataContractJSON `json:"-"`
}

// customerCreditListResponseDataContractJSON contains the JSON metadata for the
// struct [CustomerCreditListResponseDataContract]
type customerCreditListResponseDataContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponseDataContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseDataContractJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListResponseDataLedger struct {
	Type      CustomerCreditListResponseDataLedgerType `json:"type,required"`
	Timestamp time.Time                                `json:"timestamp,required" format:"date-time"`
	Amount    float64                                  `json:"amount,required"`
	SegmentID string                                   `json:"segment_id" format:"uuid"`
	InvoiceID string                                   `json:"invoice_id" format:"uuid"`
	Reason    string                                   `json:"reason"`
	JSON      customerCreditListResponseDataLedgerJSON `json:"-"`
	union     CustomerCreditListResponseDataLedgerUnion
}

// customerCreditListResponseDataLedgerJSON contains the JSON metadata for the
// struct [CustomerCreditListResponseDataLedger]
type customerCreditListResponseDataLedgerJSON struct {
	Type        apijson.Field
	Timestamp   apijson.Field
	Amount      apijson.Field
	SegmentID   apijson.Field
	InvoiceID   apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r customerCreditListResponseDataLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *CustomerCreditListResponseDataLedger) UnmarshalJSON(data []byte) (err error) {
	*r = CustomerCreditListResponseDataLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CustomerCreditListResponseDataLedgerUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntry],
// [CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntry],
// [CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntry],
// [CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntry],
// [CustomerCreditListResponseDataLedgerCreditManualLedgerEntry].
func (r CustomerCreditListResponseDataLedger) AsUnion() CustomerCreditListResponseDataLedgerUnion {
	return r.union
}

// Union satisfied by
// [CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntry],
// [CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntry],
// [CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntry],
// [CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntry] or
// [CustomerCreditListResponseDataLedgerCreditManualLedgerEntry].
type CustomerCreditListResponseDataLedgerUnion interface {
	implementsCustomerCreditListResponseDataLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerCreditListResponseDataLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerCreditListResponseDataLedgerCreditManualLedgerEntry{}),
		},
	)
}

type CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64                                                               `json:"amount,required"`
	SegmentID string                                                                `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                             `json:"timestamp,required" format:"date-time"`
	Type      CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntryType `json:"type,required"`
	JSON      customerCreditListResponseDataLedgerCreditSegmentStartLedgerEntryJSON `json:"-"`
}

// customerCreditListResponseDataLedgerCreditSegmentStartLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntry]
type customerCreditListResponseDataLedgerCreditSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseDataLedgerCreditSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntry) implementsCustomerCreditListResponseDataLedger() {
}

type CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntryType string

const (
	CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntryType = "CREDIT_SEGMENT_START"
)

func (r CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseDataLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart:
		return true
	}
	return false
}

type CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                            `json:"amount,required"`
	InvoiceID string                                                                             `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                             `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                          `json:"timestamp,required" format:"date-time"`
	Type      CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      customerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// customerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntry]
type customerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntry) implementsCustomerCreditListResponseDataLedger() {
}

type CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntryType string

const (
	CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseDataLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntry struct {
	Amount    float64                                                             `json:"amount,required"`
	SegmentID string                                                              `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                           `json:"timestamp,required" format:"date-time"`
	Type      CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntryType `json:"type,required"`
	JSON      customerCreditListResponseDataLedgerCreditExpirationLedgerEntryJSON `json:"-"`
}

// customerCreditListResponseDataLedgerCreditExpirationLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntry]
type customerCreditListResponseDataLedgerCreditExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseDataLedgerCreditExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntry) implementsCustomerCreditListResponseDataLedger() {
}

type CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntryType string

const (
	CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntryTypeCreditExpiration CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntryType = "CREDIT_EXPIRATION"
)

func (r CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseDataLedgerCreditExpirationLedgerEntryTypeCreditExpiration:
		return true
	}
	return false
}

type CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntry struct {
	Amount    float64                                                           `json:"amount,required"`
	InvoiceID string                                                            `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                         `json:"timestamp,required" format:"date-time"`
	Type      CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntryType `json:"type,required"`
	JSON      customerCreditListResponseDataLedgerCreditCanceledLedgerEntryJSON `json:"-"`
}

// customerCreditListResponseDataLedgerCreditCanceledLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntry]
type customerCreditListResponseDataLedgerCreditCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseDataLedgerCreditCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntry) implementsCustomerCreditListResponseDataLedger() {
}

type CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntryType string

const (
	CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntryTypeCreditCanceled CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntryType = "CREDIT_CANCELED"
)

func (r CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseDataLedgerCreditCanceledLedgerEntryTypeCreditCanceled:
		return true
	}
	return false
}

type CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntry struct {
	Amount    float64                                                           `json:"amount,required"`
	InvoiceID string                                                            `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                         `json:"timestamp,required" format:"date-time"`
	Type      CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntryType `json:"type,required"`
	JSON      customerCreditListResponseDataLedgerCreditCreditedLedgerEntryJSON `json:"-"`
}

// customerCreditListResponseDataLedgerCreditCreditedLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntry]
type customerCreditListResponseDataLedgerCreditCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseDataLedgerCreditCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntry) implementsCustomerCreditListResponseDataLedger() {
}

type CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntryType string

const (
	CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntryTypeCreditCredited CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntryType = "CREDIT_CREDITED"
)

func (r CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseDataLedgerCreditCreditedLedgerEntryTypeCreditCredited:
		return true
	}
	return false
}

type CustomerCreditListResponseDataLedgerCreditManualLedgerEntry struct {
	Amount    float64                                                         `json:"amount,required"`
	Reason    string                                                          `json:"reason,required"`
	Timestamp time.Time                                                       `json:"timestamp,required" format:"date-time"`
	Type      CustomerCreditListResponseDataLedgerCreditManualLedgerEntryType `json:"type,required"`
	JSON      customerCreditListResponseDataLedgerCreditManualLedgerEntryJSON `json:"-"`
}

// customerCreditListResponseDataLedgerCreditManualLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditListResponseDataLedgerCreditManualLedgerEntry]
type customerCreditListResponseDataLedgerCreditManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponseDataLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseDataLedgerCreditManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditListResponseDataLedgerCreditManualLedgerEntry) implementsCustomerCreditListResponseDataLedger() {
}

type CustomerCreditListResponseDataLedgerCreditManualLedgerEntryType string

const (
	CustomerCreditListResponseDataLedgerCreditManualLedgerEntryTypeCreditManual CustomerCreditListResponseDataLedgerCreditManualLedgerEntryType = "CREDIT_MANUAL"
)

func (r CustomerCreditListResponseDataLedgerCreditManualLedgerEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseDataLedgerCreditManualLedgerEntryTypeCreditManual:
		return true
	}
	return false
}

type CustomerCreditListResponseDataLedgerType string

const (
	CustomerCreditListResponseDataLedgerTypeCreditSegmentStart              CustomerCreditListResponseDataLedgerType = "CREDIT_SEGMENT_START"
	CustomerCreditListResponseDataLedgerTypeCreditAutomatedInvoiceDeduction CustomerCreditListResponseDataLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	CustomerCreditListResponseDataLedgerTypeCreditExpiration                CustomerCreditListResponseDataLedgerType = "CREDIT_EXPIRATION"
	CustomerCreditListResponseDataLedgerTypeCreditCanceled                  CustomerCreditListResponseDataLedgerType = "CREDIT_CANCELED"
	CustomerCreditListResponseDataLedgerTypeCreditCredited                  CustomerCreditListResponseDataLedgerType = "CREDIT_CREDITED"
	CustomerCreditListResponseDataLedgerTypeCreditManual                    CustomerCreditListResponseDataLedgerType = "CREDIT_MANUAL"
)

func (r CustomerCreditListResponseDataLedgerType) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseDataLedgerTypeCreditSegmentStart, CustomerCreditListResponseDataLedgerTypeCreditAutomatedInvoiceDeduction, CustomerCreditListResponseDataLedgerTypeCreditExpiration, CustomerCreditListResponseDataLedgerTypeCreditCanceled, CustomerCreditListResponseDataLedgerTypeCreditCredited, CustomerCreditListResponseDataLedgerTypeCreditManual:
		return true
	}
	return false
}

type CustomerCreditUpdateEndDateResponse struct {
	Data shared.ID                               `json:"data,required"`
	JSON customerCreditUpdateEndDateResponseJSON `json:"-"`
}

// customerCreditUpdateEndDateResponseJSON contains the JSON metadata for the
// struct [CustomerCreditUpdateEndDateResponse]
type customerCreditUpdateEndDateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditUpdateEndDateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditUpdateEndDateResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditNewParams struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule param.Field[CustomerCreditNewParamsAccessSchedule] `json:"access_schedule,required"`
	CustomerID     param.Field[string]                                `json:"customer_id,required" format:"uuid"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority  param.Field[float64] `json:"priority,required"`
	ProductID param.Field[string]  `json:"product_id,required" format:"uuid"`
	// Which contract the credit applies to. If not provided, the credit applies to all
	// contracts.
	ApplicableContractIDs param.Field[[]string] `json:"applicable_contract_ids"`
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
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string] `json:"salesforce_opportunity_id"`
}

func (r CustomerCreditNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the credit to the customer.
type CustomerCreditNewParamsAccessSchedule struct {
	ScheduleItems param.Field[[]CustomerCreditNewParamsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	CreditTypeID  param.Field[string]                                              `json:"credit_type_id" format:"uuid"`
}

func (r CustomerCreditNewParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCreditNewParamsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r CustomerCreditNewParamsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCreditListParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Return only credits that have access schedules that "cover" the provided date
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	CreditID     param.Field[string]    `json:"credit_id" format:"uuid"`
	// Include only credits that have any access before the provided date (exclusive)
	EffectiveBefore param.Field[time.Time] `json:"effective_before" format:"date-time"`
	// Include credits from archived contracts.
	IncludeArchived param.Field[bool] `json:"include_archived"`
	// Include credits on the contract level.
	IncludeContractCredits param.Field[bool] `json:"include_contract_credits"`
	// Include credit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
	// The next page token from a previous response.
	NextPage param.Field[string] `json:"next_page"`
	// Include only credits that have any access on or after the provided date
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r CustomerCreditListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCreditUpdateEndDateParams struct {
	// RFC 3339 timestamp indicating when access to the credit will end and it will no
	// longer be possible to draw it down (exclusive).
	AccessEndingBefore param.Field[time.Time] `json:"access_ending_before,required" format:"date-time"`
	// ID of the commit to update
	CreditID param.Field[string] `json:"credit_id,required" format:"uuid"`
	// ID of the customer whose credit is to be updated
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r CustomerCreditUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
