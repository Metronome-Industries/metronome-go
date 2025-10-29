// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v2/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v2/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/v2/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/v2/packages/param"
	"github.com/Metronome-Industries/metronome-go/v2/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/v2/shared"
)

// V1CustomerInvoiceService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerInvoiceService] method instead.
type V1CustomerInvoiceService struct {
	Options []option.RequestOption
}

// NewV1CustomerInvoiceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerInvoiceService(opts ...option.RequestOption) (r V1CustomerInvoiceService) {
	r = V1CustomerInvoiceService{}
	r.Options = opts
	return
}

// Retrieve detailed information for a specific invoice by its unique identifier.
// This endpoint returns comprehensive invoice data including line items, applied
// credits, totals, and billing period details for both finalized and draft
// invoices.
//
// ### Use this endpoint to:
//
//   - Display historical invoice details in customer-facing dashboards or billing
//     portals.
//   - Retrieve current month draft invoices to show customers their month-to-date
//     spend.
//   - Access finalized invoices for historical billing records and payment
//     reconciliation.
//   - Validate customer pricing and credit applications for customer support
//     queries.
//
// ### Key response fields:
//
// Invoice status (DRAFT, FINALIZED, VOID) Billing period start and end dates Total
// amount and amount due after credits Detailed line items broken down by:
//
// - Customer and contract information
// - Invoice line item type
// - Product/service name and ID
// - Quantity consumed
// - Unit and total price
// - Time period for usage-based charges
// - Applied credits or prepaid commitments
//
// ### Usage guidelines:
//
//   - Draft invoices update in real-time as usage is reported and may change before
//     finalization
//   - The response includes both usage-based line items (e.g., API calls, data
//     processed) and scheduled charges (e.g., monthly subscriptions, commitment
//     fees)
//   - Credit and commitment applications are shown as separate line items with
//     negative amounts
//   - For voided invoices, the response will indicate VOID status but retain all
//     original line item details
func (r *V1CustomerInvoiceService) Get(ctx context.Context, params V1CustomerInvoiceGetParams, opts ...option.RequestOption) (res *V1CustomerInvoiceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if params.InvoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/invoices/%s", params.CustomerID, params.InvoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Retrieves a paginated list of invoices for a specific customer, with flexible
// filtering options to narrow results by status, date range, credit type, and
// more. This endpoint provides a comprehensive view of a customer's billing
// history and current charges, supporting both real-time billing dashboards and
// historical reporting needs.
//
// ### Use this endpoint to:
//
//   - Display historical invoice details in customer-facing dashboards or billing
//     portals.
//   - Retrieve current month draft invoices to show customers their month-to-date
//     spend.
//   - Access finalized invoices for historical billing records and payment
//     reconciliation.
//   - Validate customer pricing and credit applications for customer support
//     queries.
//   - Generate financial reports by filtering invoices within specific date ranges
//
// ### Key response fields:
//
// Array of invoice objects containing:
//
// - Invoice ID and status (DRAFT, FINALIZED, VOID)
// - Invoice type (USAGE, SCHEDULED)
// - Billing period start and end dates
// - Issue date and due date
// - Total amount, subtotal, and amount due
// - Applied credits summary
// - Contract ID reference
// - External billing provider status (if integrated with Stripe, etc.)
// - Pagination metadata `next_page` cursor
//
// ### Usage guidelines:
//
//   - The endpoint returns invoice summaries; use the Get Invoice endpoint for
//     detailed line items
//   - Draft invoices are continuously updated as new usage is reported and will show
//     real-time spend
//   - Results are ordered by creation date descending by default (newest first)
//   - When filtering by date range, the filter applies to the billing period, not
//     the issue date
//   - For customers with many invoices, implement pagination to ensure all results
//     are retrieved External billing provider statuses (like Stripe payment status)
//     are included when applicable
//   - Voided invoices are included in results by default unless filtered out by
//     status
func (r *V1CustomerInvoiceService) List(ctx context.Context, params V1CustomerInvoiceListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Invoice], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/invoices", params.CustomerID)
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

// Retrieves a paginated list of invoices for a specific customer, with flexible
// filtering options to narrow results by status, date range, credit type, and
// more. This endpoint provides a comprehensive view of a customer's billing
// history and current charges, supporting both real-time billing dashboards and
// historical reporting needs.
//
// ### Use this endpoint to:
//
//   - Display historical invoice details in customer-facing dashboards or billing
//     portals.
//   - Retrieve current month draft invoices to show customers their month-to-date
//     spend.
//   - Access finalized invoices for historical billing records and payment
//     reconciliation.
//   - Validate customer pricing and credit applications for customer support
//     queries.
//   - Generate financial reports by filtering invoices within specific date ranges
//
// ### Key response fields:
//
// Array of invoice objects containing:
//
// - Invoice ID and status (DRAFT, FINALIZED, VOID)
// - Invoice type (USAGE, SCHEDULED)
// - Billing period start and end dates
// - Issue date and due date
// - Total amount, subtotal, and amount due
// - Applied credits summary
// - Contract ID reference
// - External billing provider status (if integrated with Stripe, etc.)
// - Pagination metadata `next_page` cursor
//
// ### Usage guidelines:
//
//   - The endpoint returns invoice summaries; use the Get Invoice endpoint for
//     detailed line items
//   - Draft invoices are continuously updated as new usage is reported and will show
//     real-time spend
//   - Results are ordered by creation date descending by default (newest first)
//   - When filtering by date range, the filter applies to the billing period, not
//     the issue date
//   - For customers with many invoices, implement pagination to ensure all results
//     are retrieved External billing provider statuses (like Stripe payment status)
//     are included when applicable
//   - Voided invoices are included in results by default unless filtered out by
//     status
func (r *V1CustomerInvoiceService) ListAutoPaging(ctx context.Context, params V1CustomerInvoiceListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Invoice] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Add a one time charge to the specified invoice. This is a Plans (deprecated)
// endpoint. New clients should implement using Contracts.
func (r *V1CustomerInvoiceService) AddCharge(ctx context.Context, params V1CustomerInvoiceAddChargeParams, opts ...option.RequestOption) (res *V1CustomerInvoiceAddChargeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/addCharge", params.CustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve granular time-series breakdowns of invoice data at hourly or daily
// intervals. This endpoint transforms standard invoices into detailed timelines,
// enabling you to track usage patterns, identify consumption spikes, and provide
// customers with transparency into their billing details throughout the billing
// period.
//
// ### Use this endpoint to:
//
// - Build usage analytics dashboards showing daily or hourly consumption trends
// - Identify peak usage periods for capacity planning and cost optimization
// - Generate detailed billing reports for finance teams and customer success
// - Troubleshoot billing disputes by examining usage patterns at specific times
// - Power real-time cost monitoring and alerting systems
//
// ### Key response fields:
//
// An array of BreakdownInvoice objects, each containing:
//
// - All standard invoice fields (ID, customer, commit, line items, totals, status)
// - Line items with quantities and costs for that specific period
// - `breakdown_start_timestamp`: Start of the specific time window
// - `breakdown_end_timestamp`: End of the specific time window
// - `next_page`: Pagination cursor for large result sets
//
// ### Usage guidelines:
//
//   - Time granularity: Set `window_size` to hour or day based on your analysis
//     needs
//   - Response limits: Daily breakdowns return up to 35 days; hourly breakdowns
//     return up to 24 hours per request
//   - Date filtering: Use `starting_on` and `ending_before` to focus on specific
//     periods
//   - Performance: For large date ranges, use pagination to retrieve all data
//     efficiently
//   - Backdated usage: If usage events arrive after invoice finalization, breakdowns
//     will reflect the updated usage
//   - Zero quantity filtering: Use `skip_zero_qty_line_items=true` to exclude
//     periods with no usage
func (r *V1CustomerInvoiceService) ListBreakdowns(ctx context.Context, params V1CustomerInvoiceListBreakdownsParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1CustomerInvoiceListBreakdownsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/invoices/breakdowns", params.CustomerID)
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

// Retrieve granular time-series breakdowns of invoice data at hourly or daily
// intervals. This endpoint transforms standard invoices into detailed timelines,
// enabling you to track usage patterns, identify consumption spikes, and provide
// customers with transparency into their billing details throughout the billing
// period.
//
// ### Use this endpoint to:
//
// - Build usage analytics dashboards showing daily or hourly consumption trends
// - Identify peak usage periods for capacity planning and cost optimization
// - Generate detailed billing reports for finance teams and customer success
// - Troubleshoot billing disputes by examining usage patterns at specific times
// - Power real-time cost monitoring and alerting systems
//
// ### Key response fields:
//
// An array of BreakdownInvoice objects, each containing:
//
// - All standard invoice fields (ID, customer, commit, line items, totals, status)
// - Line items with quantities and costs for that specific period
// - `breakdown_start_timestamp`: Start of the specific time window
// - `breakdown_end_timestamp`: End of the specific time window
// - `next_page`: Pagination cursor for large result sets
//
// ### Usage guidelines:
//
//   - Time granularity: Set `window_size` to hour or day based on your analysis
//     needs
//   - Response limits: Daily breakdowns return up to 35 days; hourly breakdowns
//     return up to 24 hours per request
//   - Date filtering: Use `starting_on` and `ending_before` to focus on specific
//     periods
//   - Performance: For large date ranges, use pagination to retrieve all data
//     efficiently
//   - Backdated usage: If usage events arrive after invoice finalization, breakdowns
//     will reflect the updated usage
//   - Zero quantity filtering: Use `skip_zero_qty_line_items=true` to exclude
//     periods with no usage
func (r *V1CustomerInvoiceService) ListBreakdownsAutoPaging(ctx context.Context, params V1CustomerInvoiceListBreakdownsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1CustomerInvoiceListBreakdownsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListBreakdowns(ctx, params, opts...))
}

// Retrieve a PDF version of a specific invoice by its unique identifier. This
// endpoint generates a professionally formatted invoice document suitable for
// sharing with customers, accounting teams, or for record-keeping purposes.
//
// ### Use this endpoint to:
//
// - Provide customers with downloadable or emailable copies of their invoices
// - Support accounting and finance teams with official billing documents
// - Maintain accurate records of billing transactions for audits and compliance
//
// ### Key response details:
//
//   - The response is a binary PDF file representing the full invoice
//   - The PDF includes all standard invoice information such as line items, totals,
//     billing period, and customer details
//   - The document is formatted for clarity and professionalism, suitable for
//     official use
//
// ### Usage guidelines:
//
//   - Ensure the `invoice_id` corresponds to an existing invoice for the specified
//     `customer_id`
//   - The PDF is generated on-demand; frequent requests for the same invoice may
//     impact performance
//   - Use appropriate headers to handle the binary response in your application
//     (e.g., setting `Content-Type: application/pdf`)
func (r *V1CustomerInvoiceService) GetPdf(ctx context.Context, query V1CustomerInvoiceGetPdfParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if query.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if query.InvoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/invoices/%s/pdf", query.CustomerID, query.InvoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Invoice struct {
	ID          string                `json:"id,required" format:"uuid"`
	CreditType  shared.CreditTypeData `json:"credit_type,required"`
	CustomerID  string                `json:"customer_id,required" format:"uuid"`
	LineItems   []InvoiceLineItem     `json:"line_items,required"`
	Status      string                `json:"status,required"`
	Total       float64               `json:"total,required"`
	Type        string                `json:"type,required"`
	AmendmentID string                `json:"amendment_id" format:"uuid"`
	// This field's availability is dependent on your client's configuration.
	BillableStatus any `json:"billable_status"`
	// Account hierarchy M3 - Required on invoices with type USAGE_CONSOLIDATED. List
	// of constituent invoices that were consolidated to create this invoice.
	ConstituentInvoices []InvoiceConstituentInvoice `json:"constituent_invoices"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	ContractCustomFields map[string]string       `json:"contract_custom_fields"`
	ContractID           string                  `json:"contract_id" format:"uuid"`
	CorrectionRecord     InvoiceCorrectionRecord `json:"correction_record"`
	// When the invoice was created (UTC). This field is present for correction
	// invoices only.
	CreatedAt    time.Time      `json:"created_at" format:"date-time"`
	CustomFields map[string]any `json:"custom_fields"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomerCustomFields map[string]string `json:"customer_custom_fields"`
	// End of the usage period this invoice covers (UTC)
	EndTimestamp       time.Time                  `json:"end_timestamp" format:"date-time"`
	ExternalInvoice    InvoiceExternalInvoice     `json:"external_invoice,nullable"`
	InvoiceAdjustments []InvoiceInvoiceAdjustment `json:"invoice_adjustments"`
	// When the invoice was issued (UTC)
	IssuedAt            time.Time `json:"issued_at" format:"date-time"`
	NetPaymentTermsDays float64   `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Account hierarchy M3 - Required for account hierarchy usage invoices. An object
	// containing the contract and customer UUIDs that pay for this invoice.
	Payer InvoicePayer `json:"payer"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	PlanCustomFields map[string]string `json:"plan_custom_fields"`
	PlanID           string            `json:"plan_id" format:"uuid"`
	PlanName         string            `json:"plan_name"`
	// Only present for contract invoices with reseller royalties.
	ResellerRoyalty InvoiceResellerRoyalty `json:"reseller_royalty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Beginning of the usage period this invoice covers (UTC)
	StartTimestamp time.Time `json:"start_timestamp" format:"date-time"`
	Subtotal       float64   `json:"subtotal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CreditType              respjson.Field
		CustomerID              respjson.Field
		LineItems               respjson.Field
		Status                  respjson.Field
		Total                   respjson.Field
		Type                    respjson.Field
		AmendmentID             respjson.Field
		BillableStatus          respjson.Field
		ConstituentInvoices     respjson.Field
		ContractCustomFields    respjson.Field
		ContractID              respjson.Field
		CorrectionRecord        respjson.Field
		CreatedAt               respjson.Field
		CustomFields            respjson.Field
		CustomerCustomFields    respjson.Field
		EndTimestamp            respjson.Field
		ExternalInvoice         respjson.Field
		InvoiceAdjustments      respjson.Field
		IssuedAt                respjson.Field
		NetPaymentTermsDays     respjson.Field
		NetsuiteSalesOrderID    respjson.Field
		Payer                   respjson.Field
		PlanCustomFields        respjson.Field
		PlanID                  respjson.Field
		PlanName                respjson.Field
		ResellerRoyalty         respjson.Field
		SalesforceOpportunityID respjson.Field
		StartTimestamp          respjson.Field
		Subtotal                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Invoice) RawJSON() string { return r.JSON.raw }
func (r *Invoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItem struct {
	CreditType shared.CreditTypeData `json:"credit_type,required"`
	Name       string                `json:"name,required"`
	Total      float64               `json:"total,required"`
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
	AppliedCommitOrCredit InvoiceLineItemAppliedCommitOrCredit `json:"applied_commit_or_credit"`
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
	ListPrice shared.Rate `json:"list_price"`
	Metadata  string      `json:"metadata"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd time.Time `json:"netsuite_invoice_billing_end" format:"date-time"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart time.Time `json:"netsuite_invoice_billing_start" format:"date-time"`
	NetsuiteItemID              string    `json:"netsuite_item_id"`
	// Account hierarchy M3 - Present on line items from invoices with type
	// USAGE_CONSOLIDATED. Indicates the original customer, contract, invoice and line
	// item from which this line item was copied.
	Origin InvoiceLineItemOrigin `json:"origin"`
	// Only present for line items paying for a postpaid commit true-up.
	PostpaidCommit InvoiceLineItemPostpaidCommit `json:"postpaid_commit"`
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
	Quantity float64 `json:"quantity"`
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType string `json:"reseller_type"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	ScheduledChargeCustomFields map[string]string `json:"scheduled_charge_custom_fields"`
	// ID of scheduled charge.
	ScheduledChargeID string `json:"scheduled_charge_id" format:"uuid"`
	// The line item's start date (inclusive).
	StartingAt   time.Time                    `json:"starting_at" format:"date-time"`
	SubLineItems []InvoiceLineItemSubLineItem `json:"sub_line_items"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	SubscriptionCustomFields map[string]string `json:"subscription_custom_fields"`
	// Populated if the line item has a tiered price.
	Tier InvoiceLineItemTier `json:"tier"`
	// The unit price associated with the line item.
	UnitPrice float64 `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType                      respjson.Field
		Name                            respjson.Field
		Total                           respjson.Field
		Type                            respjson.Field
		AppliedCommitOrCredit           respjson.Field
		CommitCustomFields              respjson.Field
		CommitID                        respjson.Field
		CommitNetsuiteItemID            respjson.Field
		CommitNetsuiteSalesOrderID      respjson.Field
		CommitSegmentID                 respjson.Field
		CommitType                      respjson.Field
		CustomFields                    respjson.Field
		DiscountCustomFields            respjson.Field
		DiscountID                      respjson.Field
		EndingBefore                    respjson.Field
		GroupKey                        respjson.Field
		GroupValue                      respjson.Field
		IsProrated                      respjson.Field
		ListPrice                       respjson.Field
		Metadata                        respjson.Field
		NetsuiteInvoiceBillingEnd       respjson.Field
		NetsuiteInvoiceBillingStart     respjson.Field
		NetsuiteItemID                  respjson.Field
		Origin                          respjson.Field
		PostpaidCommit                  respjson.Field
		PresentationGroupValues         respjson.Field
		PricingGroupValues              respjson.Field
		ProductCustomFields             respjson.Field
		ProductID                       respjson.Field
		ProductTags                     respjson.Field
		ProductType                     respjson.Field
		ProfessionalServiceCustomFields respjson.Field
		ProfessionalServiceID           respjson.Field
		Quantity                        respjson.Field
		ResellerType                    respjson.Field
		ScheduledChargeCustomFields     respjson.Field
		ScheduledChargeID               respjson.Field
		StartingAt                      respjson.Field
		SubLineItems                    respjson.Field
		SubscriptionCustomFields        respjson.Field
		Tier                            respjson.Field
		UnitPrice                       respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceLineItem) RawJSON() string { return r.JSON.raw }
func (r *InvoiceLineItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about the credit or commit that was applied to this line item. Only
// present on line items with product of `USAGE`, `SUBSCRIPTION` or `COMPOSITE`
// types.
type InvoiceLineItemAppliedCommitOrCredit struct {
	ID string `json:"id,required" format:"uuid"`
	// Any of "PREPAID", "POSTPAID", "CREDIT".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceLineItemAppliedCommitOrCredit) RawJSON() string { return r.JSON.raw }
func (r *InvoiceLineItemAppliedCommitOrCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account hierarchy M3 - Present on line items from invoices with type
// USAGE_CONSOLIDATED. Indicates the original customer, contract, invoice and line
// item from which this line item was copied.
type InvoiceLineItemOrigin struct {
	ContractID string `json:"contract_id,required" format:"uuid"`
	CustomerID string `json:"customer_id,required" format:"uuid"`
	InvoiceID  string `json:"invoice_id,required" format:"uuid"`
	LineItemID string `json:"line_item_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContractID  respjson.Field
		CustomerID  respjson.Field
		InvoiceID   respjson.Field
		LineItemID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceLineItemOrigin) RawJSON() string { return r.JSON.raw }
func (r *InvoiceLineItemOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only present for line items paying for a postpaid commit true-up.
type InvoiceLineItemPostpaidCommit struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceLineItemPostpaidCommit) RawJSON() string { return r.JSON.raw }
func (r *InvoiceLineItemPostpaidCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemSubLineItem struct {
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
	TierPeriod InvoiceLineItemSubLineItemTierPeriod `json:"tier_period"`
	Tiers      []InvoiceLineItemSubLineItemTier     `json:"tiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomFields  respjson.Field
		Name          respjson.Field
		Quantity      respjson.Field
		Subtotal      respjson.Field
		ChargeID      respjson.Field
		CreditGrantID respjson.Field
		EndDate       respjson.Field
		Price         respjson.Field
		StartDate     respjson.Field
		TierPeriod    respjson.Field
		Tiers         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceLineItemSubLineItem) RawJSON() string { return r.JSON.raw }
func (r *InvoiceLineItemSubLineItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// when the current tier started and ends (for tiered charges only)
type InvoiceLineItemSubLineItemTierPeriod struct {
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StartingAt   respjson.Field
		EndingBefore respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceLineItemSubLineItemTierPeriod) RawJSON() string { return r.JSON.raw }
func (r *InvoiceLineItemSubLineItemTierPeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemSubLineItemTier struct {
	Price    float64 `json:"price,required"`
	Quantity float64 `json:"quantity,required"`
	// at what metric amount this tier begins
	StartingAt float64 `json:"starting_at,required"`
	Subtotal   float64 `json:"subtotal,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Price       respjson.Field
		Quantity    respjson.Field
		StartingAt  respjson.Field
		Subtotal    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceLineItemSubLineItemTier) RawJSON() string { return r.JSON.raw }
func (r *InvoiceLineItemSubLineItemTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Populated if the line item has a tiered price.
type InvoiceLineItemTier struct {
	Level      float64 `json:"level,required"`
	StartingAt string  `json:"starting_at,required"`
	Size       string  `json:"size,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		StartingAt  respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceLineItemTier) RawJSON() string { return r.JSON.raw }
func (r *InvoiceLineItemTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceConstituentInvoice struct {
	ContractID string `json:"contract_id,required" format:"uuid"`
	CustomerID string `json:"customer_id,required" format:"uuid"`
	InvoiceID  string `json:"invoice_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContractID  respjson.Field
		CustomerID  respjson.Field
		InvoiceID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceConstituentInvoice) RawJSON() string { return r.JSON.raw }
func (r *InvoiceConstituentInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceCorrectionRecord struct {
	CorrectedInvoiceID       string                                          `json:"corrected_invoice_id,required" format:"uuid"`
	Memo                     string                                          `json:"memo,required"`
	Reason                   string                                          `json:"reason,required"`
	CorrectedExternalInvoice InvoiceCorrectionRecordCorrectedExternalInvoice `json:"corrected_external_invoice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CorrectedInvoiceID       respjson.Field
		Memo                     respjson.Field
		Reason                   respjson.Field
		CorrectedExternalInvoice respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceCorrectionRecord) RawJSON() string { return r.JSON.raw }
func (r *InvoiceCorrectionRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceCorrectionRecordCorrectedExternalInvoice struct {
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace", "metronome".
	BillingProviderType string `json:"billing_provider_type,required"`
	// Any of "DRAFT", "FINALIZED", "PAID", "UNCOLLECTIBLE", "VOID", "DELETED",
	// "PAYMENT_FAILED", "INVALID_REQUEST_ERROR", "SKIPPED", "SENT", "QUEUED".
	ExternalStatus string `json:"external_status"`
	InvoiceID      string `json:"invoice_id"`
	// The subtotal amount invoiced, if available from the billing provider.
	InvoicedSubTotal float64 `json:"invoiced_sub_total"`
	// The total amount invoiced, if available from the billing provider.
	InvoicedTotal     float64   `json:"invoiced_total"`
	IssuedAtTimestamp time.Time `json:"issued_at_timestamp" format:"date-time"`
	// A URL to the PDF of the invoice, if available from the billing provider.
	PdfURL string `json:"pdf_url" format:"uri"`
	// Tax details for the invoice, if available from the billing provider.
	Tax InvoiceCorrectionRecordCorrectedExternalInvoiceTax `json:"tax"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingProviderType respjson.Field
		ExternalStatus      respjson.Field
		InvoiceID           respjson.Field
		InvoicedSubTotal    respjson.Field
		InvoicedTotal       respjson.Field
		IssuedAtTimestamp   respjson.Field
		PdfURL              respjson.Field
		Tax                 respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceCorrectionRecordCorrectedExternalInvoice) RawJSON() string { return r.JSON.raw }
func (r *InvoiceCorrectionRecordCorrectedExternalInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tax details for the invoice, if available from the billing provider.
type InvoiceCorrectionRecordCorrectedExternalInvoiceTax struct {
	// The total tax amount applied to the invoice.
	TotalTaxAmount float64 `json:"total_tax_amount"`
	// The total taxable amount of the invoice.
	TotalTaxableAmount float64 `json:"total_taxable_amount"`
	// The transaction ID associated with the tax calculation.
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalTaxAmount     respjson.Field
		TotalTaxableAmount respjson.Field
		TransactionID      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceCorrectionRecordCorrectedExternalInvoiceTax) RawJSON() string { return r.JSON.raw }
func (r *InvoiceCorrectionRecordCorrectedExternalInvoiceTax) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceExternalInvoice struct {
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace", "metronome".
	BillingProviderType string `json:"billing_provider_type,required"`
	// Any of "DRAFT", "FINALIZED", "PAID", "UNCOLLECTIBLE", "VOID", "DELETED",
	// "PAYMENT_FAILED", "INVALID_REQUEST_ERROR", "SKIPPED", "SENT", "QUEUED".
	ExternalStatus string `json:"external_status"`
	InvoiceID      string `json:"invoice_id"`
	// The subtotal amount invoiced, if available from the billing provider.
	InvoicedSubTotal float64 `json:"invoiced_sub_total"`
	// The total amount invoiced, if available from the billing provider.
	InvoicedTotal     float64   `json:"invoiced_total"`
	IssuedAtTimestamp time.Time `json:"issued_at_timestamp" format:"date-time"`
	// A URL to the PDF of the invoice, if available from the billing provider.
	PdfURL string `json:"pdf_url" format:"uri"`
	// Tax details for the invoice, if available from the billing provider.
	Tax InvoiceExternalInvoiceTax `json:"tax"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingProviderType respjson.Field
		ExternalStatus      respjson.Field
		InvoiceID           respjson.Field
		InvoicedSubTotal    respjson.Field
		InvoicedTotal       respjson.Field
		IssuedAtTimestamp   respjson.Field
		PdfURL              respjson.Field
		Tax                 respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceExternalInvoice) RawJSON() string { return r.JSON.raw }
func (r *InvoiceExternalInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tax details for the invoice, if available from the billing provider.
type InvoiceExternalInvoiceTax struct {
	// The total tax amount applied to the invoice.
	TotalTaxAmount float64 `json:"total_tax_amount"`
	// The total taxable amount of the invoice.
	TotalTaxableAmount float64 `json:"total_taxable_amount"`
	// The transaction ID associated with the tax calculation.
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalTaxAmount     respjson.Field
		TotalTaxableAmount respjson.Field
		TransactionID      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceExternalInvoiceTax) RawJSON() string { return r.JSON.raw }
func (r *InvoiceExternalInvoiceTax) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceInvoiceAdjustment struct {
	CreditType shared.CreditTypeData `json:"credit_type,required"`
	Name       string                `json:"name,required"`
	Total      float64               `json:"total,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CreditGrantCustomFields map[string]string `json:"credit_grant_custom_fields"`
	CreditGrantID           string            `json:"credit_grant_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType              respjson.Field
		Name                    respjson.Field
		Total                   respjson.Field
		CreditGrantCustomFields respjson.Field
		CreditGrantID           respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceInvoiceAdjustment) RawJSON() string { return r.JSON.raw }
func (r *InvoiceInvoiceAdjustment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account hierarchy M3 - Required for account hierarchy usage invoices. An object
// containing the contract and customer UUIDs that pay for this invoice.
type InvoicePayer struct {
	ContractID string `json:"contract_id,required" format:"uuid"`
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContractID  respjson.Field
		CustomerID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoicePayer) RawJSON() string { return r.JSON.raw }
func (r *InvoicePayer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only present for contract invoices with reseller royalties.
type InvoiceResellerRoyalty struct {
	Fraction           string `json:"fraction,required"`
	NetsuiteResellerID string `json:"netsuite_reseller_id,required"`
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType string                           `json:"reseller_type,required"`
	AwsOptions   InvoiceResellerRoyaltyAwsOptions `json:"aws_options"`
	GcpOptions   InvoiceResellerRoyaltyGcpOptions `json:"gcp_options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fraction           respjson.Field
		NetsuiteResellerID respjson.Field
		ResellerType       respjson.Field
		AwsOptions         respjson.Field
		GcpOptions         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceResellerRoyalty) RawJSON() string { return r.JSON.raw }
func (r *InvoiceResellerRoyalty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    string `json:"aws_account_number"`
	AwsOfferID          string `json:"aws_offer_id"`
	AwsPayerReferenceID string `json:"aws_payer_reference_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwsAccountNumber    respjson.Field
		AwsOfferID          respjson.Field
		AwsPayerReferenceID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceResellerRoyaltyAwsOptions) RawJSON() string { return r.JSON.raw }
func (r *InvoiceResellerRoyaltyAwsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceResellerRoyaltyGcpOptions struct {
	GcpAccountID string `json:"gcp_account_id"`
	GcpOfferID   string `json:"gcp_offer_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GcpAccountID respjson.Field
		GcpOfferID   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceResellerRoyaltyGcpOptions) RawJSON() string { return r.JSON.raw }
func (r *InvoiceResellerRoyaltyGcpOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerInvoiceGetResponse struct {
	Data Invoice `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerInvoiceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerInvoiceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerInvoiceAddChargeResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerInvoiceAddChargeResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerInvoiceAddChargeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerInvoiceListBreakdownsResponse struct {
	BreakdownEndTimestamp   time.Time `json:"breakdown_end_timestamp,required" format:"date-time"`
	BreakdownStartTimestamp time.Time `json:"breakdown_start_timestamp,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BreakdownEndTimestamp   respjson.Field
		BreakdownStartTimestamp respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
	Invoice
}

// Returns the unmodified JSON received from the API
func (r V1CustomerInvoiceListBreakdownsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerInvoiceListBreakdownsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerInvoiceGetParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	InvoiceID  string `path:"invoice_id,required" format:"uuid" json:"-"`
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems param.Opt[bool] `query:"skip_zero_qty_line_items,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerInvoiceGetParams]'s query parameters as
// `url.Values`.
func (r V1CustomerInvoiceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerInvoiceListParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	// Only return invoices for the specified credit type
	CreditTypeID param.Opt[string] `query:"credit_type_id,omitzero" json:"-"`
	// RFC 3339 timestamp (exclusive). Invoices will only be returned for billing
	// periods that end before this time.
	EndingBefore param.Opt[time.Time] `query:"ending_before,omitzero" format:"date-time" json:"-"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems param.Opt[bool] `query:"skip_zero_qty_line_items,omitzero" json:"-"`
	// RFC 3339 timestamp (inclusive). Invoices will only be returned for billing
	// periods that start at or after this time.
	StartingOn param.Opt[time.Time] `query:"starting_on,omitzero" format:"date-time" json:"-"`
	// Invoice status, e.g. DRAFT, FINALIZED, or VOID
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
	// date_asc.
	//
	// Any of "date_asc", "date_desc".
	Sort V1CustomerInvoiceListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerInvoiceListParams]'s query parameters as
// `url.Values`.
func (r V1CustomerInvoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
// date_asc.
type V1CustomerInvoiceListParamsSort string

const (
	V1CustomerInvoiceListParamsSortDateAsc  V1CustomerInvoiceListParamsSort = "date_asc"
	V1CustomerInvoiceListParamsSortDateDesc V1CustomerInvoiceListParamsSort = "date_desc"
)

type V1CustomerInvoiceAddChargeParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	// The Metronome ID of the charge to add to the invoice. Note that the charge must
	// be on a product that is not on the current plan, and the product must have only
	// fixed charges.
	ChargeID string `json:"charge_id,required" format:"uuid"`
	// The Metronome ID of the customer plan to add the charge to.
	CustomerPlanID string `json:"customer_plan_id,required" format:"uuid"`
	Description    string `json:"description,required"`
	// The start_timestamp of the invoice to add the charge to.
	InvoiceStartTimestamp time.Time `json:"invoice_start_timestamp,required" format:"date-time"`
	// The price of the charge. This price will match the currency on the invoice, e.g.
	// USD cents.
	Price    float64 `json:"price,required"`
	Quantity float64 `json:"quantity,required"`
	paramObj
}

func (r V1CustomerInvoiceAddChargeParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerInvoiceAddChargeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerInvoiceAddChargeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerInvoiceListBreakdownsParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	// RFC 3339 timestamp. Breakdowns will only be returned for time windows that end
	// on or before this time.
	EndingBefore time.Time `query:"ending_before,required" format:"date-time" json:"-"`
	// RFC 3339 timestamp. Breakdowns will only be returned for time windows that start
	// on or after this time.
	StartingOn time.Time `query:"starting_on,required" format:"date-time" json:"-"`
	// Only return invoices for the specified credit type
	CreditTypeID param.Opt[string] `query:"credit_type_id,omitzero" json:"-"`
	// Max number of results that should be returned. For daily breakdowns, the
	// response can return up to 35 days worth of breakdowns. For hourly breakdowns,
	// the response can return up to 24 hours. If there are more results, a cursor to
	// the next page is returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems param.Opt[bool] `query:"skip_zero_qty_line_items,omitzero" json:"-"`
	// Invoice status, e.g. DRAFT or FINALIZED
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
	// date_asc.
	//
	// Any of "date_asc", "date_desc".
	Sort V1CustomerInvoiceListBreakdownsParamsSort `query:"sort,omitzero" json:"-"`
	// The granularity of the breakdowns to return. Defaults to day.
	//
	// Any of "HOUR", "DAY".
	WindowSize V1CustomerInvoiceListBreakdownsParamsWindowSize `query:"window_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerInvoiceListBreakdownsParams]'s query parameters
// as `url.Values`.
func (r V1CustomerInvoiceListBreakdownsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
// date_asc.
type V1CustomerInvoiceListBreakdownsParamsSort string

const (
	V1CustomerInvoiceListBreakdownsParamsSortDateAsc  V1CustomerInvoiceListBreakdownsParamsSort = "date_asc"
	V1CustomerInvoiceListBreakdownsParamsSortDateDesc V1CustomerInvoiceListBreakdownsParamsSort = "date_desc"
)

// The granularity of the breakdowns to return. Defaults to day.
type V1CustomerInvoiceListBreakdownsParamsWindowSize string

const (
	V1CustomerInvoiceListBreakdownsParamsWindowSizeHour V1CustomerInvoiceListBreakdownsParamsWindowSize = "HOUR"
	V1CustomerInvoiceListBreakdownsParamsWindowSizeDay  V1CustomerInvoiceListBreakdownsParamsWindowSize = "DAY"
)

type V1CustomerInvoiceGetPdfParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	InvoiceID  string `path:"invoice_id,required" format:"uuid" json:"-"`
	paramObj
}
