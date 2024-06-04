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
)

// CustomerInvoiceService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerInvoiceService] method instead.
type CustomerInvoiceService struct {
	Options []option.RequestOption
}

// NewCustomerInvoiceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerInvoiceService(opts ...option.RequestOption) (r *CustomerInvoiceService) {
	r = &CustomerInvoiceService{}
	r.Options = opts
	return
}

// Fetch a specific invoice for a given customer.
func (r *CustomerInvoiceService) Get(ctx context.Context, customerID string, invoiceID string, query CustomerInvoiceGetParams, opts ...option.RequestOption) (res *CustomerInvoiceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/invoices/%s", customerID, invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List all invoices for a given customer, optionally filtered by status, date
// range, and/or credit type.
func (r *CustomerInvoiceService) List(ctx context.Context, customerID string, query CustomerInvoiceListParams, opts ...option.RequestOption) (res *CustomerInvoiceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/invoices", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CustomerInvoiceGetResponse struct {
	Data CustomerInvoiceGetResponseData `json:"data,required"`
	JSON customerInvoiceGetResponseJSON `json:"-"`
}

// customerInvoiceGetResponseJSON contains the JSON metadata for the struct
// [CustomerInvoiceGetResponse]
type customerInvoiceGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseData struct {
	ID                   string                                         `json:"id,required" format:"uuid"`
	BillableStatus       CustomerInvoiceGetResponseDataBillableStatus   `json:"billable_status,required"`
	CreditType           CustomerInvoiceGetResponseDataCreditType       `json:"credit_type,required"`
	CustomerID           string                                         `json:"customer_id,required" format:"uuid"`
	LineItems            []CustomerInvoiceGetResponseDataLineItem       `json:"line_items,required"`
	Status               string                                         `json:"status,required"`
	Total                float64                                        `json:"total,required"`
	Type                 string                                         `json:"type,required"`
	AmendmentID          string                                         `json:"amendment_id" format:"uuid"`
	ContractCustomFields map[string]string                              `json:"contract_custom_fields"`
	ContractID           string                                         `json:"contract_id" format:"uuid"`
	CorrectionRecord     CustomerInvoiceGetResponseDataCorrectionRecord `json:"correction_record"`
	// When the invoice was created (UTC). This field is present for correction
	// invoices only.
	CreatedAt            time.Time              `json:"created_at" format:"date-time"`
	CustomFields         map[string]interface{} `json:"custom_fields"`
	CustomerCustomFields map[string]string      `json:"customer_custom_fields"`
	// End of the usage period this invoice covers (UTC)
	EndTimestamp       time.Time                                         `json:"end_timestamp" format:"date-time"`
	ExternalInvoice    CustomerInvoiceGetResponseDataExternalInvoice     `json:"external_invoice,nullable"`
	InvoiceAdjustments []CustomerInvoiceGetResponseDataInvoiceAdjustment `json:"invoice_adjustments"`
	// When the invoice was issued (UTC)
	IssuedAt            time.Time `json:"issued_at" format:"date-time"`
	NetPaymentTermsDays float64   `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string            `json:"netsuite_sales_order_id"`
	PlanCustomFields     map[string]string `json:"plan_custom_fields"`
	PlanID               string            `json:"plan_id" format:"uuid"`
	PlanName             string            `json:"plan_name"`
	// only present for beta contract invoices with reseller royalties
	ResellerRoyalty CustomerInvoiceGetResponseDataResellerRoyalty `json:"reseller_royalty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Beginning of the usage period this invoice covers (UTC)
	StartTimestamp time.Time                          `json:"start_timestamp" format:"date-time"`
	Subtotal       float64                            `json:"subtotal"`
	JSON           customerInvoiceGetResponseDataJSON `json:"-"`
}

// customerInvoiceGetResponseDataJSON contains the JSON metadata for the struct
// [CustomerInvoiceGetResponseData]
type customerInvoiceGetResponseDataJSON struct {
	ID                      apijson.Field
	BillableStatus          apijson.Field
	CreditType              apijson.Field
	CustomerID              apijson.Field
	LineItems               apijson.Field
	Status                  apijson.Field
	Total                   apijson.Field
	Type                    apijson.Field
	AmendmentID             apijson.Field
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

func (r *CustomerInvoiceGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataBillableStatus string

const (
	CustomerInvoiceGetResponseDataBillableStatusBillable   CustomerInvoiceGetResponseDataBillableStatus = "billable"
	CustomerInvoiceGetResponseDataBillableStatusUnbillable CustomerInvoiceGetResponseDataBillableStatus = "unbillable"
)

func (r CustomerInvoiceGetResponseDataBillableStatus) IsKnown() bool {
	switch r {
	case CustomerInvoiceGetResponseDataBillableStatusBillable, CustomerInvoiceGetResponseDataBillableStatusUnbillable:
		return true
	}
	return false
}

type CustomerInvoiceGetResponseDataCreditType struct {
	ID   string                                       `json:"id,required" format:"uuid"`
	Name string                                       `json:"name,required"`
	JSON customerInvoiceGetResponseDataCreditTypeJSON `json:"-"`
}

// customerInvoiceGetResponseDataCreditTypeJSON contains the JSON metadata for the
// struct [CustomerInvoiceGetResponseDataCreditType]
type customerInvoiceGetResponseDataCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataCreditTypeJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataLineItem struct {
	CreditType         CustomerInvoiceGetResponseDataLineItemsCreditType `json:"credit_type,required"`
	Name               string                                            `json:"name,required"`
	Total              float64                                           `json:"total,required"`
	CommitCustomFields map[string]string                                 `json:"commit_custom_fields"`
	// only present for beta contract invoices
	CommitID string `json:"commit_id" format:"uuid"`
	// only present for beta contract invoices. This field's availability is dependent
	// on your client's configuration.
	CommitNetsuiteItemID string `json:"commit_netsuite_item_id"`
	// only present for beta contract invoices. This field's availability is dependent
	// on your client's configuration.
	CommitNetsuiteSalesOrderID string `json:"commit_netsuite_sales_order_id"`
	// only present for beta contract invoices
	CommitSegmentID string `json:"commit_segment_id" format:"uuid"`
	// only present for beta contract invoices
	CommitType   string            `json:"commit_type"`
	CustomFields map[string]string `json:"custom_fields"`
	// only present for beta contract invoices
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	GroupKey     string    `json:"group_key"`
	GroupValue   string    `json:"group_value,nullable"`
	// only present for beta contract invoices
	IsProrated bool   `json:"is_prorated"`
	Metadata   string `json:"metadata"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd time.Time `json:"netsuite_invoice_billing_end" format:"date-time"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart time.Time `json:"netsuite_invoice_billing_start" format:"date-time"`
	// only present for beta contract invoices. This field's availability is dependent
	// on your client's configuration.
	NetsuiteItemID string `json:"netsuite_item_id"`
	// only present for beta contract invoices
	PostpaidCommit CustomerInvoiceGetResponseDataLineItemsPostpaidCommit `json:"postpaid_commit"`
	// if presentation groups are used, this will contain the values used to break down
	// the line item
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	// if pricing groups are used, this will contain the values used to calculate the
	// price
	PricingGroupValues              map[string]string `json:"pricing_group_values"`
	ProductCustomFields             map[string]string `json:"product_custom_fields"`
	ProductID                       string            `json:"product_id" format:"uuid"`
	ProductType                     string            `json:"product_type"`
	ProfessionalServiceCustomFields map[string]string `json:"professional_service_custom_fields"`
	// only present for beta contract invoices
	ProfessionalServiceID       string                                              `json:"professional_service_id" format:"uuid"`
	Quantity                    float64                                             `json:"quantity"`
	ResellerType                CustomerInvoiceGetResponseDataLineItemsResellerType `json:"reseller_type"`
	ScheduledChargeCustomFields map[string]string                                   `json:"scheduled_charge_custom_fields"`
	// only present for beta contract invoices
	ScheduledChargeID string `json:"scheduled_charge_id" format:"uuid"`
	// only present for beta contract invoices
	StartingAt   time.Time                                            `json:"starting_at" format:"date-time"`
	SubLineItems []CustomerInvoiceGetResponseDataLineItemsSubLineItem `json:"sub_line_items"`
	// only present for beta contract invoices
	UnitPrice float64                                    `json:"unit_price"`
	JSON      customerInvoiceGetResponseDataLineItemJSON `json:"-"`
}

// customerInvoiceGetResponseDataLineItemJSON contains the JSON metadata for the
// struct [CustomerInvoiceGetResponseDataLineItem]
type customerInvoiceGetResponseDataLineItemJSON struct {
	CreditType                      apijson.Field
	Name                            apijson.Field
	Total                           apijson.Field
	CommitCustomFields              apijson.Field
	CommitID                        apijson.Field
	CommitNetsuiteItemID            apijson.Field
	CommitNetsuiteSalesOrderID      apijson.Field
	CommitSegmentID                 apijson.Field
	CommitType                      apijson.Field
	CustomFields                    apijson.Field
	EndingBefore                    apijson.Field
	GroupKey                        apijson.Field
	GroupValue                      apijson.Field
	IsProrated                      apijson.Field
	Metadata                        apijson.Field
	NetsuiteInvoiceBillingEnd       apijson.Field
	NetsuiteInvoiceBillingStart     apijson.Field
	NetsuiteItemID                  apijson.Field
	PostpaidCommit                  apijson.Field
	PresentationGroupValues         apijson.Field
	PricingGroupValues              apijson.Field
	ProductCustomFields             apijson.Field
	ProductID                       apijson.Field
	ProductType                     apijson.Field
	ProfessionalServiceCustomFields apijson.Field
	ProfessionalServiceID           apijson.Field
	Quantity                        apijson.Field
	ResellerType                    apijson.Field
	ScheduledChargeCustomFields     apijson.Field
	ScheduledChargeID               apijson.Field
	StartingAt                      apijson.Field
	SubLineItems                    apijson.Field
	UnitPrice                       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataLineItemJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataLineItemsCreditType struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	Name string                                                `json:"name,required"`
	JSON customerInvoiceGetResponseDataLineItemsCreditTypeJSON `json:"-"`
}

// customerInvoiceGetResponseDataLineItemsCreditTypeJSON contains the JSON metadata
// for the struct [CustomerInvoiceGetResponseDataLineItemsCreditType]
type customerInvoiceGetResponseDataLineItemsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataLineItemsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataLineItemsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// only present for beta contract invoices
type CustomerInvoiceGetResponseDataLineItemsPostpaidCommit struct {
	ID   string                                                    `json:"id,required" format:"uuid"`
	JSON customerInvoiceGetResponseDataLineItemsPostpaidCommitJSON `json:"-"`
}

// customerInvoiceGetResponseDataLineItemsPostpaidCommitJSON contains the JSON
// metadata for the struct [CustomerInvoiceGetResponseDataLineItemsPostpaidCommit]
type customerInvoiceGetResponseDataLineItemsPostpaidCommitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataLineItemsPostpaidCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataLineItemsPostpaidCommitJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataLineItemsResellerType string

const (
	CustomerInvoiceGetResponseDataLineItemsResellerTypeAws           CustomerInvoiceGetResponseDataLineItemsResellerType = "AWS"
	CustomerInvoiceGetResponseDataLineItemsResellerTypeAwsProService CustomerInvoiceGetResponseDataLineItemsResellerType = "AWS_PRO_SERVICE"
	CustomerInvoiceGetResponseDataLineItemsResellerTypeGcp           CustomerInvoiceGetResponseDataLineItemsResellerType = "GCP"
	CustomerInvoiceGetResponseDataLineItemsResellerTypeGcpProService CustomerInvoiceGetResponseDataLineItemsResellerType = "GCP_PRO_SERVICE"
)

func (r CustomerInvoiceGetResponseDataLineItemsResellerType) IsKnown() bool {
	switch r {
	case CustomerInvoiceGetResponseDataLineItemsResellerTypeAws, CustomerInvoiceGetResponseDataLineItemsResellerTypeAwsProService, CustomerInvoiceGetResponseDataLineItemsResellerTypeGcp, CustomerInvoiceGetResponseDataLineItemsResellerTypeGcpProService:
		return true
	}
	return false
}

type CustomerInvoiceGetResponseDataLineItemsSubLineItem struct {
	CustomFields  map[string]string `json:"custom_fields,required"`
	Name          string            `json:"name,required"`
	Quantity      float64           `json:"quantity,required"`
	Subtotal      float64           `json:"subtotal,required"`
	ChargeID      string            `json:"charge_id" format:"uuid"`
	CreditGrantID string            `json:"credit_grant_id" format:"uuid"`
	// the unit price for this charge, present only if the charge is not tiered and the
	// quantity is nonzero
	Price float64                                                   `json:"price"`
	Tiers []CustomerInvoiceGetResponseDataLineItemsSubLineItemsTier `json:"tiers"`
	JSON  customerInvoiceGetResponseDataLineItemsSubLineItemJSON    `json:"-"`
}

// customerInvoiceGetResponseDataLineItemsSubLineItemJSON contains the JSON
// metadata for the struct [CustomerInvoiceGetResponseDataLineItemsSubLineItem]
type customerInvoiceGetResponseDataLineItemsSubLineItemJSON struct {
	CustomFields  apijson.Field
	Name          apijson.Field
	Quantity      apijson.Field
	Subtotal      apijson.Field
	ChargeID      apijson.Field
	CreditGrantID apijson.Field
	Price         apijson.Field
	Tiers         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataLineItemsSubLineItemJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataLineItemsSubLineItemsTier struct {
	Price    float64 `json:"price,required"`
	Quantity float64 `json:"quantity,required"`
	// at what metric amount this tier begins
	StartingAt float64                                                     `json:"starting_at,required"`
	Subtotal   float64                                                     `json:"subtotal,required"`
	JSON       customerInvoiceGetResponseDataLineItemsSubLineItemsTierJSON `json:"-"`
}

// customerInvoiceGetResponseDataLineItemsSubLineItemsTierJSON contains the JSON
// metadata for the struct
// [CustomerInvoiceGetResponseDataLineItemsSubLineItemsTier]
type customerInvoiceGetResponseDataLineItemsSubLineItemsTierJSON struct {
	Price       apijson.Field
	Quantity    apijson.Field
	StartingAt  apijson.Field
	Subtotal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataLineItemsSubLineItemsTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataLineItemsSubLineItemsTierJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataCorrectionRecord struct {
	CorrectedInvoiceID       string                                                                 `json:"corrected_invoice_id,required" format:"uuid"`
	Memo                     string                                                                 `json:"memo,required"`
	Reason                   string                                                                 `json:"reason,required"`
	CorrectedExternalInvoice CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoice `json:"corrected_external_invoice"`
	JSON                     customerInvoiceGetResponseDataCorrectionRecordJSON                     `json:"-"`
}

// customerInvoiceGetResponseDataCorrectionRecordJSON contains the JSON metadata
// for the struct [CustomerInvoiceGetResponseDataCorrectionRecord]
type customerInvoiceGetResponseDataCorrectionRecordJSON struct {
	CorrectedInvoiceID       apijson.Field
	Memo                     apijson.Field
	Reason                   apijson.Field
	CorrectedExternalInvoice apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataCorrectionRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataCorrectionRecordJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoice struct {
	BillingProviderType CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                                                    `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                                                 `json:"issued_at_timestamp" format:"date-time"`
	JSON                customerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceJSON                `json:"-"`
}

// customerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceJSON
// contains the JSON metadata for the struct
// [CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoice]
type customerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType string

const (
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace   CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "aws_marketplace"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe           CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "stripe"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite         CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "netsuite"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom           CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "custom"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "azure_marketplace"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "quickbooks_online"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday          CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "workday"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace   CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus string

const (
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft               CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DRAFT"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized           CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "FINALIZED"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid                CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAID"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible       CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid                CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "VOID"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted             CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DELETED"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed       CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped             CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SKIPPED"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSent                CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SENT"
	CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued              CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "QUEUED"
)

func (r CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSent, CustomerInvoiceGetResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type CustomerInvoiceGetResponseDataExternalInvoice struct {
	BillingProviderType CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                           `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                        `json:"issued_at_timestamp" format:"date-time"`
	JSON                customerInvoiceGetResponseDataExternalInvoiceJSON                `json:"-"`
}

// customerInvoiceGetResponseDataExternalInvoiceJSON contains the JSON metadata for
// the struct [CustomerInvoiceGetResponseDataExternalInvoice]
type customerInvoiceGetResponseDataExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType string

const (
	CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeAwsMarketplace   CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "aws_marketplace"
	CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeStripe           CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "stripe"
	CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeNetsuite         CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "netsuite"
	CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeCustom           CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "custom"
	CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeAzureMarketplace CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "azure_marketplace"
	CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeQuickbooksOnline CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "quickbooks_online"
	CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeWorkday          CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "workday"
	CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeGcpMarketplace   CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeAwsMarketplace, CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeStripe, CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeNetsuite, CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeCustom, CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeAzureMarketplace, CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeQuickbooksOnline, CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeWorkday, CustomerInvoiceGetResponseDataExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus string

const (
	CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusDraft               CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "DRAFT"
	CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusFinalized           CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "FINALIZED"
	CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusPaid                CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "PAID"
	CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusUncollectible       CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusVoid                CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "VOID"
	CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusDeleted             CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "DELETED"
	CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusPaymentFailed       CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusInvalidRequestError CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusSkipped             CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "SKIPPED"
	CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusSent                CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "SENT"
	CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusQueued              CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus = "QUEUED"
)

func (r CustomerInvoiceGetResponseDataExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusDraft, CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusFinalized, CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusPaid, CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusUncollectible, CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusVoid, CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusDeleted, CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusPaymentFailed, CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusInvalidRequestError, CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusSkipped, CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusSent, CustomerInvoiceGetResponseDataExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type CustomerInvoiceGetResponseDataInvoiceAdjustment struct {
	CreditType    CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditType `json:"credit_type,required"`
	Name          string                                                     `json:"name,required"`
	Total         float64                                                    `json:"total,required"`
	CreditGrantID string                                                     `json:"credit_grant_id"`
	JSON          customerInvoiceGetResponseDataInvoiceAdjustmentJSON        `json:"-"`
}

// customerInvoiceGetResponseDataInvoiceAdjustmentJSON contains the JSON metadata
// for the struct [CustomerInvoiceGetResponseDataInvoiceAdjustment]
type customerInvoiceGetResponseDataInvoiceAdjustmentJSON struct {
	CreditType    apijson.Field
	Name          apijson.Field
	Total         apijson.Field
	CreditGrantID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataInvoiceAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataInvoiceAdjustmentJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditType struct {
	ID   string                                                         `json:"id,required" format:"uuid"`
	Name string                                                         `json:"name,required"`
	JSON customerInvoiceGetResponseDataInvoiceAdjustmentsCreditTypeJSON `json:"-"`
}

// customerInvoiceGetResponseDataInvoiceAdjustmentsCreditTypeJSON contains the JSON
// metadata for the struct
// [CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditType]
type customerInvoiceGetResponseDataInvoiceAdjustmentsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataInvoiceAdjustmentsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataInvoiceAdjustmentsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// only present for beta contract invoices with reseller royalties
type CustomerInvoiceGetResponseDataResellerRoyalty struct {
	Fraction           string                                                    `json:"fraction,required"`
	NetsuiteResellerID string                                                    `json:"netsuite_reseller_id,required"`
	ResellerType       CustomerInvoiceGetResponseDataResellerRoyaltyResellerType `json:"reseller_type,required"`
	AwsOptions         CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptions   `json:"aws_options"`
	GcpOptions         CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptions   `json:"gcp_options"`
	JSON               customerInvoiceGetResponseDataResellerRoyaltyJSON         `json:"-"`
}

// customerInvoiceGetResponseDataResellerRoyaltyJSON contains the JSON metadata for
// the struct [CustomerInvoiceGetResponseDataResellerRoyalty]
type customerInvoiceGetResponseDataResellerRoyaltyJSON struct {
	Fraction           apijson.Field
	NetsuiteResellerID apijson.Field
	ResellerType       apijson.Field
	AwsOptions         apijson.Field
	GcpOptions         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataResellerRoyaltyResellerType string

const (
	CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeAws           CustomerInvoiceGetResponseDataResellerRoyaltyResellerType = "AWS"
	CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeAwsProService CustomerInvoiceGetResponseDataResellerRoyaltyResellerType = "AWS_PRO_SERVICE"
	CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeGcp           CustomerInvoiceGetResponseDataResellerRoyaltyResellerType = "GCP"
	CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeGcpProService CustomerInvoiceGetResponseDataResellerRoyaltyResellerType = "GCP_PRO_SERVICE"
)

func (r CustomerInvoiceGetResponseDataResellerRoyaltyResellerType) IsKnown() bool {
	switch r {
	case CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeAws, CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeAwsProService, CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeGcp, CustomerInvoiceGetResponseDataResellerRoyaltyResellerTypeGcpProService:
		return true
	}
	return false
}

type CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    string                                                      `json:"aws_account_number"`
	AwsOfferID          string                                                      `json:"aws_offer_id"`
	AwsPayerReferenceID string                                                      `json:"aws_payer_reference_id"`
	JSON                customerInvoiceGetResponseDataResellerRoyaltyAwsOptionsJSON `json:"-"`
}

// customerInvoiceGetResponseDataResellerRoyaltyAwsOptionsJSON contains the JSON
// metadata for the struct
// [CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptions]
type customerInvoiceGetResponseDataResellerRoyaltyAwsOptionsJSON struct {
	AwsAccountNumber    apijson.Field
	AwsOfferID          apijson.Field
	AwsPayerReferenceID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataResellerRoyaltyAwsOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataResellerRoyaltyAwsOptionsJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptions struct {
	GcpAccountID string                                                      `json:"gcp_account_id"`
	GcpOfferID   string                                                      `json:"gcp_offer_id"`
	JSON         customerInvoiceGetResponseDataResellerRoyaltyGcpOptionsJSON `json:"-"`
}

// customerInvoiceGetResponseDataResellerRoyaltyGcpOptionsJSON contains the JSON
// metadata for the struct
// [CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptions]
type customerInvoiceGetResponseDataResellerRoyaltyGcpOptionsJSON struct {
	GcpAccountID apijson.Field
	GcpOfferID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponseDataResellerRoyaltyGcpOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseDataResellerRoyaltyGcpOptionsJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponse struct {
	Data     []CustomerInvoiceListResponseData `json:"data,required"`
	NextPage string                            `json:"next_page,required,nullable"`
	JSON     customerInvoiceListResponseJSON   `json:"-"`
}

// customerInvoiceListResponseJSON contains the JSON metadata for the struct
// [CustomerInvoiceListResponse]
type customerInvoiceListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseData struct {
	ID                   string                                          `json:"id,required" format:"uuid"`
	BillableStatus       CustomerInvoiceListResponseDataBillableStatus   `json:"billable_status,required"`
	CreditType           CustomerInvoiceListResponseDataCreditType       `json:"credit_type,required"`
	CustomerID           string                                          `json:"customer_id,required" format:"uuid"`
	LineItems            []CustomerInvoiceListResponseDataLineItem       `json:"line_items,required"`
	Status               string                                          `json:"status,required"`
	Total                float64                                         `json:"total,required"`
	Type                 string                                          `json:"type,required"`
	AmendmentID          string                                          `json:"amendment_id" format:"uuid"`
	ContractCustomFields map[string]string                               `json:"contract_custom_fields"`
	ContractID           string                                          `json:"contract_id" format:"uuid"`
	CorrectionRecord     CustomerInvoiceListResponseDataCorrectionRecord `json:"correction_record"`
	// When the invoice was created (UTC). This field is present for correction
	// invoices only.
	CreatedAt            time.Time              `json:"created_at" format:"date-time"`
	CustomFields         map[string]interface{} `json:"custom_fields"`
	CustomerCustomFields map[string]string      `json:"customer_custom_fields"`
	// End of the usage period this invoice covers (UTC)
	EndTimestamp       time.Time                                          `json:"end_timestamp" format:"date-time"`
	ExternalInvoice    CustomerInvoiceListResponseDataExternalInvoice     `json:"external_invoice,nullable"`
	InvoiceAdjustments []CustomerInvoiceListResponseDataInvoiceAdjustment `json:"invoice_adjustments"`
	// When the invoice was issued (UTC)
	IssuedAt            time.Time `json:"issued_at" format:"date-time"`
	NetPaymentTermsDays float64   `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string            `json:"netsuite_sales_order_id"`
	PlanCustomFields     map[string]string `json:"plan_custom_fields"`
	PlanID               string            `json:"plan_id" format:"uuid"`
	PlanName             string            `json:"plan_name"`
	// only present for beta contract invoices with reseller royalties
	ResellerRoyalty CustomerInvoiceListResponseDataResellerRoyalty `json:"reseller_royalty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Beginning of the usage period this invoice covers (UTC)
	StartTimestamp time.Time                           `json:"start_timestamp" format:"date-time"`
	Subtotal       float64                             `json:"subtotal"`
	JSON           customerInvoiceListResponseDataJSON `json:"-"`
}

// customerInvoiceListResponseDataJSON contains the JSON metadata for the struct
// [CustomerInvoiceListResponseData]
type customerInvoiceListResponseDataJSON struct {
	ID                      apijson.Field
	BillableStatus          apijson.Field
	CreditType              apijson.Field
	CustomerID              apijson.Field
	LineItems               apijson.Field
	Status                  apijson.Field
	Total                   apijson.Field
	Type                    apijson.Field
	AmendmentID             apijson.Field
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

func (r *CustomerInvoiceListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataBillableStatus string

const (
	CustomerInvoiceListResponseDataBillableStatusBillable   CustomerInvoiceListResponseDataBillableStatus = "billable"
	CustomerInvoiceListResponseDataBillableStatusUnbillable CustomerInvoiceListResponseDataBillableStatus = "unbillable"
)

func (r CustomerInvoiceListResponseDataBillableStatus) IsKnown() bool {
	switch r {
	case CustomerInvoiceListResponseDataBillableStatusBillable, CustomerInvoiceListResponseDataBillableStatusUnbillable:
		return true
	}
	return false
}

type CustomerInvoiceListResponseDataCreditType struct {
	ID   string                                        `json:"id,required" format:"uuid"`
	Name string                                        `json:"name,required"`
	JSON customerInvoiceListResponseDataCreditTypeJSON `json:"-"`
}

// customerInvoiceListResponseDataCreditTypeJSON contains the JSON metadata for the
// struct [CustomerInvoiceListResponseDataCreditType]
type customerInvoiceListResponseDataCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataCreditTypeJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataLineItem struct {
	CreditType         CustomerInvoiceListResponseDataLineItemsCreditType `json:"credit_type,required"`
	Name               string                                             `json:"name,required"`
	Total              float64                                            `json:"total,required"`
	CommitCustomFields map[string]string                                  `json:"commit_custom_fields"`
	// only present for beta contract invoices
	CommitID string `json:"commit_id" format:"uuid"`
	// only present for beta contract invoices. This field's availability is dependent
	// on your client's configuration.
	CommitNetsuiteItemID string `json:"commit_netsuite_item_id"`
	// only present for beta contract invoices. This field's availability is dependent
	// on your client's configuration.
	CommitNetsuiteSalesOrderID string `json:"commit_netsuite_sales_order_id"`
	// only present for beta contract invoices
	CommitSegmentID string `json:"commit_segment_id" format:"uuid"`
	// only present for beta contract invoices
	CommitType   string            `json:"commit_type"`
	CustomFields map[string]string `json:"custom_fields"`
	// only present for beta contract invoices
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	GroupKey     string    `json:"group_key"`
	GroupValue   string    `json:"group_value,nullable"`
	// only present for beta contract invoices
	IsProrated bool   `json:"is_prorated"`
	Metadata   string `json:"metadata"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd time.Time `json:"netsuite_invoice_billing_end" format:"date-time"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart time.Time `json:"netsuite_invoice_billing_start" format:"date-time"`
	// only present for beta contract invoices. This field's availability is dependent
	// on your client's configuration.
	NetsuiteItemID string `json:"netsuite_item_id"`
	// only present for beta contract invoices
	PostpaidCommit CustomerInvoiceListResponseDataLineItemsPostpaidCommit `json:"postpaid_commit"`
	// if presentation groups are used, this will contain the values used to break down
	// the line item
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	// if pricing groups are used, this will contain the values used to calculate the
	// price
	PricingGroupValues              map[string]string `json:"pricing_group_values"`
	ProductCustomFields             map[string]string `json:"product_custom_fields"`
	ProductID                       string            `json:"product_id" format:"uuid"`
	ProductType                     string            `json:"product_type"`
	ProfessionalServiceCustomFields map[string]string `json:"professional_service_custom_fields"`
	// only present for beta contract invoices
	ProfessionalServiceID       string                                               `json:"professional_service_id" format:"uuid"`
	Quantity                    float64                                              `json:"quantity"`
	ResellerType                CustomerInvoiceListResponseDataLineItemsResellerType `json:"reseller_type"`
	ScheduledChargeCustomFields map[string]string                                    `json:"scheduled_charge_custom_fields"`
	// only present for beta contract invoices
	ScheduledChargeID string `json:"scheduled_charge_id" format:"uuid"`
	// only present for beta contract invoices
	StartingAt   time.Time                                             `json:"starting_at" format:"date-time"`
	SubLineItems []CustomerInvoiceListResponseDataLineItemsSubLineItem `json:"sub_line_items"`
	// only present for beta contract invoices
	UnitPrice float64                                     `json:"unit_price"`
	JSON      customerInvoiceListResponseDataLineItemJSON `json:"-"`
}

// customerInvoiceListResponseDataLineItemJSON contains the JSON metadata for the
// struct [CustomerInvoiceListResponseDataLineItem]
type customerInvoiceListResponseDataLineItemJSON struct {
	CreditType                      apijson.Field
	Name                            apijson.Field
	Total                           apijson.Field
	CommitCustomFields              apijson.Field
	CommitID                        apijson.Field
	CommitNetsuiteItemID            apijson.Field
	CommitNetsuiteSalesOrderID      apijson.Field
	CommitSegmentID                 apijson.Field
	CommitType                      apijson.Field
	CustomFields                    apijson.Field
	EndingBefore                    apijson.Field
	GroupKey                        apijson.Field
	GroupValue                      apijson.Field
	IsProrated                      apijson.Field
	Metadata                        apijson.Field
	NetsuiteInvoiceBillingEnd       apijson.Field
	NetsuiteInvoiceBillingStart     apijson.Field
	NetsuiteItemID                  apijson.Field
	PostpaidCommit                  apijson.Field
	PresentationGroupValues         apijson.Field
	PricingGroupValues              apijson.Field
	ProductCustomFields             apijson.Field
	ProductID                       apijson.Field
	ProductType                     apijson.Field
	ProfessionalServiceCustomFields apijson.Field
	ProfessionalServiceID           apijson.Field
	Quantity                        apijson.Field
	ResellerType                    apijson.Field
	ScheduledChargeCustomFields     apijson.Field
	ScheduledChargeID               apijson.Field
	StartingAt                      apijson.Field
	SubLineItems                    apijson.Field
	UnitPrice                       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataLineItemJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataLineItemsCreditType struct {
	ID   string                                                 `json:"id,required" format:"uuid"`
	Name string                                                 `json:"name,required"`
	JSON customerInvoiceListResponseDataLineItemsCreditTypeJSON `json:"-"`
}

// customerInvoiceListResponseDataLineItemsCreditTypeJSON contains the JSON
// metadata for the struct [CustomerInvoiceListResponseDataLineItemsCreditType]
type customerInvoiceListResponseDataLineItemsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataLineItemsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataLineItemsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// only present for beta contract invoices
type CustomerInvoiceListResponseDataLineItemsPostpaidCommit struct {
	ID   string                                                     `json:"id,required" format:"uuid"`
	JSON customerInvoiceListResponseDataLineItemsPostpaidCommitJSON `json:"-"`
}

// customerInvoiceListResponseDataLineItemsPostpaidCommitJSON contains the JSON
// metadata for the struct [CustomerInvoiceListResponseDataLineItemsPostpaidCommit]
type customerInvoiceListResponseDataLineItemsPostpaidCommitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataLineItemsPostpaidCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataLineItemsPostpaidCommitJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataLineItemsResellerType string

const (
	CustomerInvoiceListResponseDataLineItemsResellerTypeAws           CustomerInvoiceListResponseDataLineItemsResellerType = "AWS"
	CustomerInvoiceListResponseDataLineItemsResellerTypeAwsProService CustomerInvoiceListResponseDataLineItemsResellerType = "AWS_PRO_SERVICE"
	CustomerInvoiceListResponseDataLineItemsResellerTypeGcp           CustomerInvoiceListResponseDataLineItemsResellerType = "GCP"
	CustomerInvoiceListResponseDataLineItemsResellerTypeGcpProService CustomerInvoiceListResponseDataLineItemsResellerType = "GCP_PRO_SERVICE"
)

func (r CustomerInvoiceListResponseDataLineItemsResellerType) IsKnown() bool {
	switch r {
	case CustomerInvoiceListResponseDataLineItemsResellerTypeAws, CustomerInvoiceListResponseDataLineItemsResellerTypeAwsProService, CustomerInvoiceListResponseDataLineItemsResellerTypeGcp, CustomerInvoiceListResponseDataLineItemsResellerTypeGcpProService:
		return true
	}
	return false
}

type CustomerInvoiceListResponseDataLineItemsSubLineItem struct {
	CustomFields  map[string]string `json:"custom_fields,required"`
	Name          string            `json:"name,required"`
	Quantity      float64           `json:"quantity,required"`
	Subtotal      float64           `json:"subtotal,required"`
	ChargeID      string            `json:"charge_id" format:"uuid"`
	CreditGrantID string            `json:"credit_grant_id" format:"uuid"`
	// the unit price for this charge, present only if the charge is not tiered and the
	// quantity is nonzero
	Price float64                                                    `json:"price"`
	Tiers []CustomerInvoiceListResponseDataLineItemsSubLineItemsTier `json:"tiers"`
	JSON  customerInvoiceListResponseDataLineItemsSubLineItemJSON    `json:"-"`
}

// customerInvoiceListResponseDataLineItemsSubLineItemJSON contains the JSON
// metadata for the struct [CustomerInvoiceListResponseDataLineItemsSubLineItem]
type customerInvoiceListResponseDataLineItemsSubLineItemJSON struct {
	CustomFields  apijson.Field
	Name          apijson.Field
	Quantity      apijson.Field
	Subtotal      apijson.Field
	ChargeID      apijson.Field
	CreditGrantID apijson.Field
	Price         apijson.Field
	Tiers         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataLineItemsSubLineItemJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataLineItemsSubLineItemsTier struct {
	Price    float64 `json:"price,required"`
	Quantity float64 `json:"quantity,required"`
	// at what metric amount this tier begins
	StartingAt float64                                                      `json:"starting_at,required"`
	Subtotal   float64                                                      `json:"subtotal,required"`
	JSON       customerInvoiceListResponseDataLineItemsSubLineItemsTierJSON `json:"-"`
}

// customerInvoiceListResponseDataLineItemsSubLineItemsTierJSON contains the JSON
// metadata for the struct
// [CustomerInvoiceListResponseDataLineItemsSubLineItemsTier]
type customerInvoiceListResponseDataLineItemsSubLineItemsTierJSON struct {
	Price       apijson.Field
	Quantity    apijson.Field
	StartingAt  apijson.Field
	Subtotal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataLineItemsSubLineItemsTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataLineItemsSubLineItemsTierJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataCorrectionRecord struct {
	CorrectedInvoiceID       string                                                                  `json:"corrected_invoice_id,required" format:"uuid"`
	Memo                     string                                                                  `json:"memo,required"`
	Reason                   string                                                                  `json:"reason,required"`
	CorrectedExternalInvoice CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoice `json:"corrected_external_invoice"`
	JSON                     customerInvoiceListResponseDataCorrectionRecordJSON                     `json:"-"`
}

// customerInvoiceListResponseDataCorrectionRecordJSON contains the JSON metadata
// for the struct [CustomerInvoiceListResponseDataCorrectionRecord]
type customerInvoiceListResponseDataCorrectionRecordJSON struct {
	CorrectedInvoiceID       apijson.Field
	Memo                     apijson.Field
	Reason                   apijson.Field
	CorrectedExternalInvoice apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataCorrectionRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataCorrectionRecordJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoice struct {
	BillingProviderType CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                                                     `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                                                  `json:"issued_at_timestamp" format:"date-time"`
	JSON                customerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceJSON                `json:"-"`
}

// customerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceJSON
// contains the JSON metadata for the struct
// [CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoice]
type customerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType string

const (
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace   CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "aws_marketplace"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe           CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "stripe"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite         CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "netsuite"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom           CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "custom"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "azure_marketplace"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "quickbooks_online"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday          CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "workday"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace   CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus string

const (
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft               CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DRAFT"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized           CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "FINALIZED"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid                CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAID"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible       CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid                CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "VOID"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted             CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DELETED"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed       CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped             CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SKIPPED"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSent                CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SENT"
	CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued              CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus = "QUEUED"
)

func (r CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusSent, CustomerInvoiceListResponseDataCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type CustomerInvoiceListResponseDataExternalInvoice struct {
	BillingProviderType CustomerInvoiceListResponseDataExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      CustomerInvoiceListResponseDataExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                            `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                         `json:"issued_at_timestamp" format:"date-time"`
	JSON                customerInvoiceListResponseDataExternalInvoiceJSON                `json:"-"`
}

// customerInvoiceListResponseDataExternalInvoiceJSON contains the JSON metadata
// for the struct [CustomerInvoiceListResponseDataExternalInvoice]
type customerInvoiceListResponseDataExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataExternalInvoiceBillingProviderType string

const (
	CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeAwsMarketplace   CustomerInvoiceListResponseDataExternalInvoiceBillingProviderType = "aws_marketplace"
	CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeStripe           CustomerInvoiceListResponseDataExternalInvoiceBillingProviderType = "stripe"
	CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeNetsuite         CustomerInvoiceListResponseDataExternalInvoiceBillingProviderType = "netsuite"
	CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeCustom           CustomerInvoiceListResponseDataExternalInvoiceBillingProviderType = "custom"
	CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeAzureMarketplace CustomerInvoiceListResponseDataExternalInvoiceBillingProviderType = "azure_marketplace"
	CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeQuickbooksOnline CustomerInvoiceListResponseDataExternalInvoiceBillingProviderType = "quickbooks_online"
	CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeWorkday          CustomerInvoiceListResponseDataExternalInvoiceBillingProviderType = "workday"
	CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeGcpMarketplace   CustomerInvoiceListResponseDataExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r CustomerInvoiceListResponseDataExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeAwsMarketplace, CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeStripe, CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeNetsuite, CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeCustom, CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeAzureMarketplace, CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeQuickbooksOnline, CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeWorkday, CustomerInvoiceListResponseDataExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type CustomerInvoiceListResponseDataExternalInvoiceExternalStatus string

const (
	CustomerInvoiceListResponseDataExternalInvoiceExternalStatusDraft               CustomerInvoiceListResponseDataExternalInvoiceExternalStatus = "DRAFT"
	CustomerInvoiceListResponseDataExternalInvoiceExternalStatusFinalized           CustomerInvoiceListResponseDataExternalInvoiceExternalStatus = "FINALIZED"
	CustomerInvoiceListResponseDataExternalInvoiceExternalStatusPaid                CustomerInvoiceListResponseDataExternalInvoiceExternalStatus = "PAID"
	CustomerInvoiceListResponseDataExternalInvoiceExternalStatusUncollectible       CustomerInvoiceListResponseDataExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	CustomerInvoiceListResponseDataExternalInvoiceExternalStatusVoid                CustomerInvoiceListResponseDataExternalInvoiceExternalStatus = "VOID"
	CustomerInvoiceListResponseDataExternalInvoiceExternalStatusDeleted             CustomerInvoiceListResponseDataExternalInvoiceExternalStatus = "DELETED"
	CustomerInvoiceListResponseDataExternalInvoiceExternalStatusPaymentFailed       CustomerInvoiceListResponseDataExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	CustomerInvoiceListResponseDataExternalInvoiceExternalStatusInvalidRequestError CustomerInvoiceListResponseDataExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	CustomerInvoiceListResponseDataExternalInvoiceExternalStatusSkipped             CustomerInvoiceListResponseDataExternalInvoiceExternalStatus = "SKIPPED"
	CustomerInvoiceListResponseDataExternalInvoiceExternalStatusSent                CustomerInvoiceListResponseDataExternalInvoiceExternalStatus = "SENT"
	CustomerInvoiceListResponseDataExternalInvoiceExternalStatusQueued              CustomerInvoiceListResponseDataExternalInvoiceExternalStatus = "QUEUED"
)

func (r CustomerInvoiceListResponseDataExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case CustomerInvoiceListResponseDataExternalInvoiceExternalStatusDraft, CustomerInvoiceListResponseDataExternalInvoiceExternalStatusFinalized, CustomerInvoiceListResponseDataExternalInvoiceExternalStatusPaid, CustomerInvoiceListResponseDataExternalInvoiceExternalStatusUncollectible, CustomerInvoiceListResponseDataExternalInvoiceExternalStatusVoid, CustomerInvoiceListResponseDataExternalInvoiceExternalStatusDeleted, CustomerInvoiceListResponseDataExternalInvoiceExternalStatusPaymentFailed, CustomerInvoiceListResponseDataExternalInvoiceExternalStatusInvalidRequestError, CustomerInvoiceListResponseDataExternalInvoiceExternalStatusSkipped, CustomerInvoiceListResponseDataExternalInvoiceExternalStatusSent, CustomerInvoiceListResponseDataExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type CustomerInvoiceListResponseDataInvoiceAdjustment struct {
	CreditType    CustomerInvoiceListResponseDataInvoiceAdjustmentsCreditType `json:"credit_type,required"`
	Name          string                                                      `json:"name,required"`
	Total         float64                                                     `json:"total,required"`
	CreditGrantID string                                                      `json:"credit_grant_id"`
	JSON          customerInvoiceListResponseDataInvoiceAdjustmentJSON        `json:"-"`
}

// customerInvoiceListResponseDataInvoiceAdjustmentJSON contains the JSON metadata
// for the struct [CustomerInvoiceListResponseDataInvoiceAdjustment]
type customerInvoiceListResponseDataInvoiceAdjustmentJSON struct {
	CreditType    apijson.Field
	Name          apijson.Field
	Total         apijson.Field
	CreditGrantID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataInvoiceAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataInvoiceAdjustmentJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataInvoiceAdjustmentsCreditType struct {
	ID   string                                                          `json:"id,required" format:"uuid"`
	Name string                                                          `json:"name,required"`
	JSON customerInvoiceListResponseDataInvoiceAdjustmentsCreditTypeJSON `json:"-"`
}

// customerInvoiceListResponseDataInvoiceAdjustmentsCreditTypeJSON contains the
// JSON metadata for the struct
// [CustomerInvoiceListResponseDataInvoiceAdjustmentsCreditType]
type customerInvoiceListResponseDataInvoiceAdjustmentsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataInvoiceAdjustmentsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataInvoiceAdjustmentsCreditTypeJSON) RawJSON() string {
	return r.raw
}

// only present for beta contract invoices with reseller royalties
type CustomerInvoiceListResponseDataResellerRoyalty struct {
	Fraction           string                                                     `json:"fraction,required"`
	NetsuiteResellerID string                                                     `json:"netsuite_reseller_id,required"`
	ResellerType       CustomerInvoiceListResponseDataResellerRoyaltyResellerType `json:"reseller_type,required"`
	AwsOptions         CustomerInvoiceListResponseDataResellerRoyaltyAwsOptions   `json:"aws_options"`
	GcpOptions         CustomerInvoiceListResponseDataResellerRoyaltyGcpOptions   `json:"gcp_options"`
	JSON               customerInvoiceListResponseDataResellerRoyaltyJSON         `json:"-"`
}

// customerInvoiceListResponseDataResellerRoyaltyJSON contains the JSON metadata
// for the struct [CustomerInvoiceListResponseDataResellerRoyalty]
type customerInvoiceListResponseDataResellerRoyaltyJSON struct {
	Fraction           apijson.Field
	NetsuiteResellerID apijson.Field
	ResellerType       apijson.Field
	AwsOptions         apijson.Field
	GcpOptions         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataResellerRoyaltyResellerType string

const (
	CustomerInvoiceListResponseDataResellerRoyaltyResellerTypeAws           CustomerInvoiceListResponseDataResellerRoyaltyResellerType = "AWS"
	CustomerInvoiceListResponseDataResellerRoyaltyResellerTypeAwsProService CustomerInvoiceListResponseDataResellerRoyaltyResellerType = "AWS_PRO_SERVICE"
	CustomerInvoiceListResponseDataResellerRoyaltyResellerTypeGcp           CustomerInvoiceListResponseDataResellerRoyaltyResellerType = "GCP"
	CustomerInvoiceListResponseDataResellerRoyaltyResellerTypeGcpProService CustomerInvoiceListResponseDataResellerRoyaltyResellerType = "GCP_PRO_SERVICE"
)

func (r CustomerInvoiceListResponseDataResellerRoyaltyResellerType) IsKnown() bool {
	switch r {
	case CustomerInvoiceListResponseDataResellerRoyaltyResellerTypeAws, CustomerInvoiceListResponseDataResellerRoyaltyResellerTypeAwsProService, CustomerInvoiceListResponseDataResellerRoyaltyResellerTypeGcp, CustomerInvoiceListResponseDataResellerRoyaltyResellerTypeGcpProService:
		return true
	}
	return false
}

type CustomerInvoiceListResponseDataResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    string                                                       `json:"aws_account_number"`
	AwsOfferID          string                                                       `json:"aws_offer_id"`
	AwsPayerReferenceID string                                                       `json:"aws_payer_reference_id"`
	JSON                customerInvoiceListResponseDataResellerRoyaltyAwsOptionsJSON `json:"-"`
}

// customerInvoiceListResponseDataResellerRoyaltyAwsOptionsJSON contains the JSON
// metadata for the struct
// [CustomerInvoiceListResponseDataResellerRoyaltyAwsOptions]
type customerInvoiceListResponseDataResellerRoyaltyAwsOptionsJSON struct {
	AwsAccountNumber    apijson.Field
	AwsOfferID          apijson.Field
	AwsPayerReferenceID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataResellerRoyaltyAwsOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataResellerRoyaltyAwsOptionsJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceListResponseDataResellerRoyaltyGcpOptions struct {
	GcpAccountID string                                                       `json:"gcp_account_id"`
	GcpOfferID   string                                                       `json:"gcp_offer_id"`
	JSON         customerInvoiceListResponseDataResellerRoyaltyGcpOptionsJSON `json:"-"`
}

// customerInvoiceListResponseDataResellerRoyaltyGcpOptionsJSON contains the JSON
// metadata for the struct
// [CustomerInvoiceListResponseDataResellerRoyaltyGcpOptions]
type customerInvoiceListResponseDataResellerRoyaltyGcpOptionsJSON struct {
	GcpAccountID apijson.Field
	GcpOfferID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomerInvoiceListResponseDataResellerRoyaltyGcpOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceListResponseDataResellerRoyaltyGcpOptionsJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetParams struct {
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems param.Field[bool] `query:"skip_zero_qty_line_items"`
}

// URLQuery serializes [CustomerInvoiceGetParams]'s query parameters as
// `url.Values`.
func (r CustomerInvoiceGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerInvoiceListParams struct {
	// Only return invoices for the specified credit type
	CreditTypeID param.Field[string] `query:"credit_type_id"`
	// RFC 3339 timestamp (exclusive). Invoices will only be returned for billing
	// periods that end before this time.
	EndingBefore param.Field[time.Time] `query:"ending_before" format:"date-time"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems param.Field[bool] `query:"skip_zero_qty_line_items"`
	// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
	// date_asc.
	Sort param.Field[CustomerInvoiceListParamsSort] `query:"sort"`
	// RFC 3339 timestamp (inclusive). Invoices will only be returned for billing
	// periods that start at or after this time.
	StartingOn param.Field[time.Time] `query:"starting_on" format:"date-time"`
	// Invoice status, e.g. DRAFT, FINALIZED, or VOID
	Status param.Field[string] `query:"status"`
}

// URLQuery serializes [CustomerInvoiceListParams]'s query parameters as
// `url.Values`.
func (r CustomerInvoiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
// date_asc.
type CustomerInvoiceListParamsSort string

const (
	CustomerInvoiceListParamsSortDateAsc  CustomerInvoiceListParamsSort = "date_asc"
	CustomerInvoiceListParamsSortDateDesc CustomerInvoiceListParamsSort = "date_desc"
)

func (r CustomerInvoiceListParamsSort) IsKnown() bool {
	switch r {
	case CustomerInvoiceListParamsSortDateAsc, CustomerInvoiceListParamsSortDateDesc:
		return true
	}
	return false
}
