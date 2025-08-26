// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// V1ContractService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContractService] method instead.
type V1ContractService struct {
	Options        []option.RequestOption
	Products       V1ContractProductService
	RateCards      V1ContractRateCardService
	NamedSchedules V1ContractNamedScheduleService
}

// NewV1ContractService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ContractService(opts ...option.RequestOption) (r V1ContractService) {
	r = V1ContractService{}
	r.Options = opts
	r.Products = NewV1ContractProductService(opts...)
	r.RateCards = NewV1ContractRateCardService(opts...)
	r.NamedSchedules = NewV1ContractNamedScheduleService(opts...)
	return
}

// Contracts define a customer's products, pricing, discounts, access duration, and
// billing configuration. Contracts serve as the central billing agreement for both
// PLG and Enterprise customers, you can automatically customers access to your
// products and services directly from your product or CRM.
//
// Common Use Cases:
//
//   - PLG onboarding: Automatically provision new self-serve customers with
//     contracts when they sign up.
//   - Enterprise sales: Push negotiated contracts from Salesforce with custom
//     pricing and commitments
//   - Promotional pricing: Implement time-limited discounts and free trials through
//     overrides
//
// Key Components:
//
//   - Contract Term and Billing Schedule
//   - Set contract duration using starting_at and ending_before fields. PLG
//     contracts typically use perpetual agreements (no end date), while Enterprise
//     contracts have fixed end dates which can be edited over time in the case of
//     co-term upsells.
//
// Rate Card\
// If you are offering usage based pricing, you can set a rate card for the contract
// to reference through rate_card_id or rate_card_alias. The rate card is a store of
// all of your usage based products and their centralized pricing. Any new products
// or price changes on the rate card can be set to automatically propagate to all associated
// contracts - this ensures consistent pricing and product launches flow to contracts
// without manual updates and migrations. The usage_statement_schedule determines the
// cadence on which Metronome will finalize a usage invoice for the customer. This defaults
// to monthly on the 1st, with options for custom dates, quarterly, or annual cadences.
// Note: Most usage based billing companies align usage statements to be evaluated aligned
// to the first of the month. Read more about [Create and Manage Rate Cards](https://docs.metronome.com/pricing-packaging/create-manage-rate-cards/).
//
// Overrides and discounts\
// Customize pricing on the contract through time-bounded overrides that can target
// specific products, product families, or complex usage scenarios. Overrides enable
// two key capabilities:
//
//   - Discounts: Apply percentage discounts, fixed rate reductions, or
//     quantity-based pricing tiers
//   - Entitlements: Provide special pricing or access to specific products for
//     negotiated deals
//
// Read more about
// [Add Contract Overrides](https://docs.metronome.com/manage-product-access/add-contract-override/).
//
// Commits and Credits\
// Using commits, configure prepaid or postpaid spending commitments where customers
// promise to spend a certain amount over the contract period paid in advance or in
// arrears. Use credits to provide free spending allowances. Under the hood these are
// the same mechanisms, however, credits are typically offered for free (SLA or promotional)
// or as a part of an allotment associated with a Subscription.
//
// In Metronome, you can set commits and credits to only be applicable for a subset
// of usage. Use applicable_product_ids or applicable_product_tags to create
// product or product-family specific commits or credits, or you can build complex
// boolean logic specifiers to target usage based on pricing and presentation group
// values using override_specifiers.
//
// These objects can also also be configured to have a recurrence schedule to
// easily model customer packaging which includes recurring monthly or quarterly
// allotments.
//
// Commits support rollover settings (rollover_fraction) to transfer unused
// balances between contract periods, either entirely or as a percentage.
//
// Read more about
// [Apply Credits and Commits](https://docs.metronome.com/pricing-packaging/apply-credits-commits/).
//
// Subscriptions\
// You can add a fixed recurring charge to a contract, like monthly licenses or seat-based
// fees, using the subscription charge. Subscription charges are defined on your rate
// card and you can select which subscription is applicable to add to each contract.
// When you add a subscription to a contract you need to:
//
//   - Define whether the subscription is paid for in-advance or in-arrears
//     (collection_schedule)
//   - Define the proration behavior (proration)
//   - Specify an initial quantity (initial_quantity)
//   - Define which subscription rate on the rate card should be used
//     (subscription_rate)
//
// Read more about
// [Create Subscriptions](https://docs.metronome.com/manage-product-access/create-subscription/).
//
// Scheduled Charges\
// Set up one-time, recurring, or entirely custom charges that occur on specific dates,
// separate from usage-based billing or commitments. These can be used to model non-recurring
// platform charges or professional services.
//
// Threshold Billing\
// Metronome allows you to configure automatic billing triggers when customers reach
// spending thresholds to prevent fraud and manage risk. You can use spend_threshold_configuration
// to trigger an invoice to cover current charges whenever the threshold is reached
// or you can ensure the customer maintains a minimum prepaid balance using the prepaid_balance_configuration
// .
//
// Read more about
// [Spend Threshold](https://docs.metronome.com/manage-product-access/spend-thresholds/)
// and
// [Prepaid Balance Thresholds](https://docs.metronome.com/manage-product-access/prepaid-balance-thresholds/).
//
// Usage guidelines:
//
//   - You can always
//     [Edit Contracts](https://docs.metronome.com/manage-product-access/edit-contract/)
//     after it has been created, using the editContract endpoint. Metronome keeps
//     track of all edits, both in the audit log and over the getEditHistory
//     endpoint.
//   - Customers in Metronome can have multiple concurrent contracts at one time. Use
//     usage_filters to route the correct usage to each contract.
//     [Read more about usage filters](https://docs.metronome.com/manage-product-access/provision-customer/#create-a-usage-filter).
func (r *V1ContractService) New(ctx context.Context, body V1ContractNewParams, opts ...option.RequestOption) (res *V1ContractNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This is the v1 endpoint to get a contract. New clients should implement using
// the v2 endpoint.
func (r *V1ContractService) Get(ctx context.Context, body V1ContractGetParams, opts ...option.RequestOption) (res *V1ContractGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves all contracts for a specific customer, including pricing, terms,
// credits, and commitments. Use this to view a customer's contract history and
// current agreements for billing management. Returns contract details with
// optional ledgers and balance information.
//
// ⚠️ Note: This is the legacy v1 endpoint - new integrations should use the v2
// endpoint for enhanced features.
func (r *V1ContractService) List(ctx context.Context, body V1ContractListParams, opts ...option.RequestOption) (res *V1ContractListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Manually adjust the available balance on a commit or credit. This entry is
// appended to the commit ledger as a new event. Optionally include a description
// that provides the reasoning for the entry.
//
// Use this endpoint to:
//
//   - Address incorrect usage burn-down caused by malformed usage or invalid config
//   - Decrease available balance to account for outages where usage may have not
//     been tracked or sent to Metronome
//   - Issue credits to customers in the form of increased balance on existing commit
//     or credit
//
// Usage guidelines:\
// Manual ledger entries can be extremely useful for resolving discrepancies in Metronome.
// However, most corrections to inaccurate billings can be modified upstream of the
// commit, whether that is via contract editing, rate editing, or other actions that
// cause an invoice to be recalculated.
func (r *V1ContractService) AddManualBalanceEntry(ctx context.Context, body V1ContractAddManualBalanceEntryParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/contracts/addManualBalanceLedgerEntry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Amendments will be replaced by Contract editing. New clients should implement
// using the editContract endpoint. Read more about the migration to contract
// editing [here](https://docs.metronome.com/migrate-amendments-to-edits/) and
// reach out to your Metronome representative for more details. Once contract
// editing is enabled, access to this endpoint will be removed.
func (r *V1ContractService) Amend(ctx context.Context, body V1ContractAmendParams, opts ...option.RequestOption) (res *V1ContractAmendResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/amend"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Permanently end and archive a contract along with all its terms. Any draft
// invoices will be canceled, and all upcoming scheduled invoices will be
// voided–also all finalized invoices can optionally be voided. Use this in the
// event a contract was incorrectly created and needed to be removed from a
// customer.
//
// Impact on commits and credits:
//
// When archiving a contract, all associated commits and credits are also archived.
// For prepaid commits with active segments, Metronome automatically generates
// expiration ledger entries to close out any remaining balances, ensuring accurate
// accounting of unused prepaid amounts. These ledger entries will appear in the
// commit's transaction history with type PREPAID_COMMIT_EXPIRATION.
//
// Archived contract visibility:
//
// Archived contracts remain accessible for historical reporting and audit
// purposes. They can be retrieved using the ListContracts endpoint by setting the
// include_archived parameter to true or in the Metronome UI when the "Show
// archived" option is enabled.
func (r *V1ContractService) Archive(ctx context.Context, body V1ContractArchiveParams, opts ...option.RequestOption) (res *V1ContractArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create historical usage invoices for past billing periods on specific contracts.
// Use this endpoint to generate retroactive invoices with custom usage line items,
// quantities, and date ranges. Supports preview mode to validate invoice data
// before creation. Ideal for billing migrations or correcting past billing
// periods.
func (r *V1ContractService) NewHistoricalInvoices(ctx context.Context, body V1ContractNewHistoricalInvoicesParams, opts ...option.RequestOption) (res *V1ContractNewHistoricalInvoicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/createHistoricalInvoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a comprehensive view of all available balances (commits and credits)
// for a customer. This endpoint provides real-time visibility into prepaid funds,
// postpaid commitments, promotional credits, and other balance types that can
// offset usage charges, helping you build transparent billing experiences.
//
// Use this endpoint to:
//
// - Display current available balances in customer dashboards
// - Verify available funds before approving high-usage operations
// - Generate balance reports for finance teams
// - Filter balances by contract or date ranges
//
// Key response fields: An array of balance objects (all credits and commits)
// containing:
//
//   - Balance details: Current available amount for each commit or credit
//   - Metadata: Product associations, priorities, applicable date ranges
//   - Optional ledger entries: Detailed transaction history (if
//     include_ledgers=true)
//   - Balance calculations: Including pending transactions and future-dated entries
//   - Custom fields: Any additional metadata attached to balances
//
// Usage guidelines:
//
//   - Date filtering: Use effective_before to include only balances with access
//     before a specific date (exclusive)
//   - Set include_balance=true for calculated balance amounts on each commit or
//     credit
//   - Set include_ledgers=true for full transaction history
//   - Set include_contract_balances = true to see contract level balances
//
//   - Balance logic: Reflects currently accessible amounts, excluding expired/future
//     segments
//   - Manual adjustments: Includes all manual ledger entries, even future-dated ones
func (r *V1ContractService) ListBalances(ctx context.Context, body V1ContractListBalancesParams, opts ...option.RequestOption) (res *pagination.BodyCursorPage[V1ContractListBalancesResponseUnion], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/contracts/customerBalances/list"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
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

// Retrieve a comprehensive view of all available balances (commits and credits)
// for a customer. This endpoint provides real-time visibility into prepaid funds,
// postpaid commitments, promotional credits, and other balance types that can
// offset usage charges, helping you build transparent billing experiences.
//
// Use this endpoint to:
//
// - Display current available balances in customer dashboards
// - Verify available funds before approving high-usage operations
// - Generate balance reports for finance teams
// - Filter balances by contract or date ranges
//
// Key response fields: An array of balance objects (all credits and commits)
// containing:
//
//   - Balance details: Current available amount for each commit or credit
//   - Metadata: Product associations, priorities, applicable date ranges
//   - Optional ledger entries: Detailed transaction history (if
//     include_ledgers=true)
//   - Balance calculations: Including pending transactions and future-dated entries
//   - Custom fields: Any additional metadata attached to balances
//
// Usage guidelines:
//
//   - Date filtering: Use effective_before to include only balances with access
//     before a specific date (exclusive)
//   - Set include_balance=true for calculated balance amounts on each commit or
//     credit
//   - Set include_ledgers=true for full transaction history
//   - Set include_contract_balances = true to see contract level balances
//
//   - Balance logic: Reflects currently accessible amounts, excluding expired/future
//     segments
//   - Manual adjustments: Includes all manual ledger entries, even future-dated ones
func (r *V1ContractService) ListBalancesAutoPaging(ctx context.Context, body V1ContractListBalancesParams, opts ...option.RequestOption) *pagination.BodyCursorPageAutoPager[V1ContractListBalancesResponseUnion] {
	return pagination.NewBodyCursorPageAutoPager(r.ListBalances(ctx, body, opts...))
}

// For a specific customer and contract, get the rates at a specific point in time.
// This endpoint takes the contract's rate card into consideration, including
// scheduled changes. It also takes into account overrides on the contract. For
// example, if you want to show your customer a summary of the prices they are
// paying, inclusive of any negotiated discounts or promotions, use this endpoint.
// This endpoint only returns rates that are entitled.
func (r *V1ContractService) GetRateSchedule(ctx context.Context, params V1ContractGetRateScheduleParams, opts ...option.RequestOption) (res *V1ContractGetRateScheduleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/getContractRateSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get the history of subscription quantities and prices over time for a given
// subscription_id. This endpoint can be used to power an in-product experience
// where you show a customer their historical changes to seat count. Future changes
// are not included in this endpoint - use the getContract endpoint to view the
// future scheduled changes to a subscription's quantity.
//
// Subscriptions are used to model fixed recurring fees as well as seat-based
// recurring fees. To model changes to the number of seats in Metronome, you can
// increment or decrement the quantity on a subscription at any point in the past
// or future.
func (r *V1ContractService) GetSubscriptionQuantityHistory(ctx context.Context, body V1ContractGetSubscriptionQuantityHistoryParams, opts ...option.RequestOption) (res *V1ContractGetSubscriptionQuantityHistoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/getSubscriptionQuantityHistory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new scheduled invoice for Professional Services terms on a contract.
// This endpoint's availability is dependent on your client's configuration.
func (r *V1ContractService) ScheduleProServicesInvoice(ctx context.Context, body V1ContractScheduleProServicesInvoiceParams, opts ...option.RequestOption) (res *V1ContractScheduleProServicesInvoiceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/scheduleProServicesInvoice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// If a customer has multiple contracts with overlapping rates, the usage filter
// routes usage to the appropriate contract based on a predefined group key.
//
// As an example, imagine you have a customer associated with two projects. Each
// project is associated with its own contract. You can create a usage filter with
// group key project_id on each contract, and route usage for project_1 to the
// first contract and project_2 to the second contract.
//
// Use this endpoint to:
//
//   - Support enterprise contracting scenarios where multiple contracts are
//     associated to the same customer with the same rates.
//   - Update the usage filter associated with the contract over time.
//
// Usage guidelines:\
// To use usage filters, the group_key must be defined on the billable metrics underlying
// the rate card on the contracts.
func (r *V1ContractService) SetUsageFilter(ctx context.Context, body V1ContractSetUsageFilterParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/contracts/setUsageFilter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Update or and an end date to a contract. Ending a contract early will impact
// draft usage statements, truncate any terms, and remove upcoming scheduled
// invoices. Moving the date into the future will only extend the contract length.
// Terms and scheduled invoices are not extended. Use this if a contract's end date
// has changed or if a perpetual contract ends.
func (r *V1ContractService) UpdateEndDate(ctx context.Context, body V1ContractUpdateEndDateParams, opts ...option.RequestOption) (res *V1ContractUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1ContractNewResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractGetResponse struct {
	Data shared.Contract `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractListResponse struct {
	Data []shared.Contract `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractAmendResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractAmendResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractAmendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractArchiveResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractNewHistoricalInvoicesResponse struct {
	Data []Invoice `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractNewHistoricalInvoicesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractNewHistoricalInvoicesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1ContractListBalancesResponseUnion contains all possible properties and values
// from [shared.Commit], [shared.Credit].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1ContractListBalancesResponseUnion struct {
	ID string `json:"id"`
	// This field is a union of [shared.CommitProduct], [shared.CreditProduct]
	Product V1ContractListBalancesResponseUnionProduct `json:"product"`
	Type    string                                     `json:"type"`
	// This field is from variant [shared.Commit].
	AccessSchedule shared.ScheduleDuration `json:"access_schedule"`
	// This field is from variant [shared.Commit].
	Amount                float64  `json:"amount"`
	ApplicableContractIDs []string `json:"applicable_contract_ids"`
	ApplicableProductIDs  []string `json:"applicable_product_ids"`
	ApplicableProductTags []string `json:"applicable_product_tags"`
	// This field is from variant [shared.Commit].
	ArchivedAt time.Time `json:"archived_at"`
	Balance    float64   `json:"balance"`
	// This field is a union of [shared.CommitContract], [shared.CreditContract]
	Contract     V1ContractListBalancesResponseUnionContract `json:"contract"`
	CustomFields string                                      `json:"custom_fields"`
	Description  string                                      `json:"description"`
	// This field is from variant [shared.Commit].
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// This field is from variant [shared.Commit].
	InvoiceContract shared.CommitInvoiceContract `json:"invoice_contract"`
	// This field is from variant [shared.Commit].
	InvoiceSchedule shared.SchedulePointInTime `json:"invoice_schedule"`
	// This field is a union of [[]shared.CommitLedgerUnion],
	// [[]shared.CreditLedgerUnion]
	Ledger               V1ContractListBalancesResponseUnionLedger `json:"ledger"`
	Name                 string                                    `json:"name"`
	NetsuiteSalesOrderID string                                    `json:"netsuite_sales_order_id"`
	Priority             float64                                   `json:"priority"`
	RateType             string                                    `json:"rate_type"`
	// This field is from variant [shared.Commit].
	RolledOverFrom shared.CommitRolledOverFrom `json:"rolled_over_from"`
	// This field is from variant [shared.Commit].
	RolloverFraction        float64                  `json:"rollover_fraction"`
	SalesforceOpportunityID string                   `json:"salesforce_opportunity_id"`
	Specifiers              []shared.CommitSpecifier `json:"specifiers"`
	UniquenessKey           string                   `json:"uniqueness_key"`
	JSON                    struct {
		ID                      respjson.Field
		Product                 respjson.Field
		Type                    respjson.Field
		AccessSchedule          respjson.Field
		Amount                  respjson.Field
		ApplicableContractIDs   respjson.Field
		ApplicableProductIDs    respjson.Field
		ApplicableProductTags   respjson.Field
		ArchivedAt              respjson.Field
		Balance                 respjson.Field
		Contract                respjson.Field
		CustomFields            respjson.Field
		Description             respjson.Field
		HierarchyConfiguration  respjson.Field
		InvoiceContract         respjson.Field
		InvoiceSchedule         respjson.Field
		Ledger                  respjson.Field
		Name                    respjson.Field
		NetsuiteSalesOrderID    respjson.Field
		Priority                respjson.Field
		RateType                respjson.Field
		RolledOverFrom          respjson.Field
		RolloverFraction        respjson.Field
		SalesforceOpportunityID respjson.Field
		Specifiers              respjson.Field
		UniquenessKey           respjson.Field
		raw                     string
	} `json:"-"`
}

func (u V1ContractListBalancesResponseUnion) AsCommit() (v shared.Commit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1ContractListBalancesResponseUnion) AsCredit() (v shared.Credit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1ContractListBalancesResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1ContractListBalancesResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1ContractListBalancesResponseUnionProduct is an implicit subunion of
// [V1ContractListBalancesResponseUnion].
// V1ContractListBalancesResponseUnionProduct provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1ContractListBalancesResponseUnion].
type V1ContractListBalancesResponseUnionProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	JSON struct {
		ID   respjson.Field
		Name respjson.Field
		raw  string
	} `json:"-"`
}

func (r *V1ContractListBalancesResponseUnionProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1ContractListBalancesResponseUnionContract is an implicit subunion of
// [V1ContractListBalancesResponseUnion].
// V1ContractListBalancesResponseUnionContract provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1ContractListBalancesResponseUnion].
type V1ContractListBalancesResponseUnionContract struct {
	ID   string `json:"id"`
	JSON struct {
		ID  respjson.Field
		raw string
	} `json:"-"`
}

func (r *V1ContractListBalancesResponseUnionContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1ContractListBalancesResponseUnionLedger is an implicit subunion of
// [V1ContractListBalancesResponseUnion]. V1ContractListBalancesResponseUnionLedger
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1ContractListBalancesResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCommitLedgerArray OfCreditLedgerArray]
type V1ContractListBalancesResponseUnionLedger struct {
	// This field will be present if the value is a [[]shared.CommitLedgerUnion]
	// instead of an object.
	OfCommitLedgerArray []shared.CommitLedgerUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.CreditLedgerUnion]
	// instead of an object.
	OfCreditLedgerArray []shared.CreditLedgerUnion `json:",inline"`
	JSON                struct {
		OfCommitLedgerArray respjson.Field
		OfCreditLedgerArray respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *V1ContractListBalancesResponseUnionLedger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractGetRateScheduleResponse struct {
	Data     []V1ContractGetRateScheduleResponseData `json:"data,required"`
	NextPage string                                  `json:"next_page,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextPage    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetRateScheduleResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractGetRateScheduleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractGetRateScheduleResponseData struct {
	Entitled bool        `json:"entitled,required"`
	ListRate shared.Rate `json:"list_rate,required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	ProductCustomFields map[string]string `json:"product_custom_fields,required"`
	ProductID           string            `json:"product_id,required" format:"uuid"`
	ProductName         string            `json:"product_name,required"`
	ProductTags         []string          `json:"product_tags,required"`
	RateCardID          string            `json:"rate_card_id,required" format:"uuid"`
	StartingAt          time.Time         `json:"starting_at,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate         shared.CommitRate `json:"commit_rate"`
	EndingBefore       time.Time         `json:"ending_before" format:"date-time"`
	OverrideRate       shared.Rate       `json:"override_rate"`
	PricingGroupValues map[string]string `json:"pricing_group_values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entitled            respjson.Field
		ListRate            respjson.Field
		ProductCustomFields respjson.Field
		ProductID           respjson.Field
		ProductName         respjson.Field
		ProductTags         respjson.Field
		RateCardID          respjson.Field
		StartingAt          respjson.Field
		BillingFrequency    respjson.Field
		CommitRate          respjson.Field
		EndingBefore        respjson.Field
		OverrideRate        respjson.Field
		PricingGroupValues  respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetRateScheduleResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ContractGetRateScheduleResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractGetSubscriptionQuantityHistoryResponse struct {
	Data V1ContractGetSubscriptionQuantityHistoryResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetSubscriptionQuantityHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractGetSubscriptionQuantityHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractGetSubscriptionQuantityHistoryResponseData struct {
	FiatCreditTypeID string                                                        `json:"fiat_credit_type_id" format:"uuid"`
	History          []V1ContractGetSubscriptionQuantityHistoryResponseDataHistory `json:"history"`
	SubscriptionID   string                                                        `json:"subscription_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FiatCreditTypeID respjson.Field
		History          respjson.Field
		SubscriptionID   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetSubscriptionQuantityHistoryResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ContractGetSubscriptionQuantityHistoryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractGetSubscriptionQuantityHistoryResponseDataHistory struct {
	Data       []V1ContractGetSubscriptionQuantityHistoryResponseDataHistoryData `json:"data,required"`
	StartingAt time.Time                                                         `json:"starting_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		StartingAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetSubscriptionQuantityHistoryResponseDataHistory) RawJSON() string {
	return r.JSON.raw
}
func (r *V1ContractGetSubscriptionQuantityHistoryResponseDataHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractGetSubscriptionQuantityHistoryResponseDataHistoryData struct {
	Quantity  float64 `json:"quantity,required"`
	Total     float64 `json:"total,required"`
	UnitPrice float64 `json:"unit_price,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		Total       respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetSubscriptionQuantityHistoryResponseDataHistoryData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1ContractGetSubscriptionQuantityHistoryResponseDataHistoryData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractScheduleProServicesInvoiceResponse struct {
	Data []Invoice `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractScheduleProServicesInvoiceResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractScheduleProServicesInvoiceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractUpdateEndDateResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractUpdateEndDateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractUpdateEndDateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractNewParams struct {
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// inclusive contract start time
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// exclusive contract end time
	EndingBefore        param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	Name                param.Opt[string]    `json:"name,omitzero"`
	NetPaymentTermsDays param.Opt[float64]   `json:"net_payment_terms_days,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Priority of the contract.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Selects the rate card linked to the specified alias as of the contract's start
	// date.
	RateCardAlias param.Opt[string] `json:"rate_card_alias,omitzero"`
	RateCardID    param.Opt[string] `json:"rate_card_id,omitzero" format:"uuid"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Opt[string] `json:"salesforce_opportunity_id,omitzero"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue param.Opt[float64] `json:"total_contract_value,omitzero"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey param.Opt[string] `json:"uniqueness_key,omitzero"`
	// The billing provider configuration associated with a contract. Provide either an
	// ID or the provider and delivery method.
	BillingProviderConfiguration V1ContractNewParamsBillingProviderConfiguration `json:"billing_provider_configuration,omitzero"`
	Commits                      []V1ContractNewParamsCommit                     `json:"commits,omitzero"`
	Credits                      []V1ContractNewParamsCredit                     `json:"credits,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// This field's availability is dependent on your client's configuration.
	Discounts              []V1ContractNewParamsDiscount             `json:"discounts,omitzero"`
	HierarchyConfiguration V1ContractNewParamsHierarchyConfiguration `json:"hierarchy_configuration,omitzero"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
	// prices automatically. EXPLICIT prioritization requires specifying priorities for
	// each multiplier; the one with the lowest priority value will be prioritized
	// first. If tiered overrides are used, prioritization must be explicit.
	//
	// Any of "LOWEST_MULTIPLIER", "EXPLICIT".
	MultiplierOverridePrioritization     V1ContractNewParamsMultiplierOverridePrioritization `json:"multiplier_override_prioritization,omitzero"`
	ContractOverrides                    []V1ContractNewParamsOverride                       `json:"overrides,omitzero"`
	PrepaidBalanceThresholdConfiguration shared.PrepaidBalanceThresholdConfigurationParam    `json:"prepaid_balance_threshold_configuration,omitzero"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []V1ContractNewParamsProfessionalService `json:"professional_services,omitzero"`
	RecurringCommits     []V1ContractNewParamsRecurringCommit     `json:"recurring_commits,omitzero"`
	RecurringCredits     []V1ContractNewParamsRecurringCredit     `json:"recurring_credits,omitzero"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []V1ContractNewParamsResellerRoyalty `json:"reseller_royalties,omitzero"`
	ScheduledCharges  []V1ContractNewParamsScheduledCharge `json:"scheduled_charges,omitzero"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	//
	// Any of "ALL".
	ScheduledChargesOnUsageInvoices V1ContractNewParamsScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices,omitzero"`
	SpendThresholdConfiguration     shared.SpendThresholdConfigurationParam            `json:"spend_threshold_configuration,omitzero"`
	// Optional list of
	// [subscriptions](https://docs.metronome.com/manage-product-access/create-subscription/)
	// to add to the contract.
	Subscriptions          []V1ContractNewParamsSubscription         `json:"subscriptions,omitzero"`
	Transition             V1ContractNewParamsTransition             `json:"transition,omitzero"`
	UsageFilter            shared.BaseUsageFilterParam               `json:"usage_filter,omitzero"`
	UsageStatementSchedule V1ContractNewParamsUsageStatementSchedule `json:"usage_statement_schedule,omitzero"`
	paramObj
}

func (r V1ContractNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The billing provider configuration associated with a contract. Provide either an
// ID or the provider and delivery method.
type V1ContractNewParamsBillingProviderConfiguration struct {
	// The Metronome ID of the billing provider configuration. Use when a customer has
	// multiple configurations with the same billing provider and delivery method.
	// Otherwise, specify the billing_provider and delivery_method.
	BillingProviderConfigurationID param.Opt[string] `json:"billing_provider_configuration_id,omitzero" format:"uuid"`
	// Do not specify if using billing_provider_configuration_id.
	//
	// Any of "aws_marketplace", "azure_marketplace", "gcp_marketplace", "stripe",
	// "netsuite".
	BillingProvider string `json:"billing_provider,omitzero"`
	// Do not specify if using billing_provider_configuration_id.
	//
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod string `json:"delivery_method,omitzero"`
	paramObj
}

func (r V1ContractNewParamsBillingProviderConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsBillingProviderConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsBillingProviderConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsBillingProviderConfiguration](
		"billing_provider", "aws_marketplace", "azure_marketplace", "gcp_marketplace", "stripe", "netsuite",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsBillingProviderConfiguration](
		"delivery_method", "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns",
	)
}

// The properties ProductID, Type are required.
type V1ContractNewParamsCommit struct {
	ProductID string `json:"product_id,required" format:"uuid"`
	// Any of "PREPAID", "POSTPAID".
	Type string `json:"type,omitzero,required"`
	// (DEPRECATED) Use access_schedule and invoice_schedule instead.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Opt[string] `json:"description,omitzero"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Opt[float64] `json:"rollover_fraction,omitzero"`
	// A temporary ID for the commit that can be used to reference the commit for
	// commit specific overrides.
	TemporaryID param.Opt[string] `json:"temporary_id,omitzero"`
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
	// commits only one schedule item is allowed and amount must match invoice_schedule
	// total.
	AccessSchedule V1ContractNewParamsCommitAccessSchedule `json:"access_schedule,omitzero"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam `json:"hierarchy_configuration,omitzero"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this
	// time and only one schedule item is allowed; the total must match access_schedule
	// amount. Optional for "PREPAID" commits: if not provided, this will be a
	// "complimentary" commit with no invoice.
	InvoiceSchedule V1ContractNewParamsCommitInvoiceSchedule `json:"invoice_schedule,omitzero"`
	// optionally payment gate this commit
	PaymentGateConfig V1ContractNewParamsCommitPaymentGateConfig `json:"payment_gate_config,omitzero"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	paramObj
}

func (r V1ContractNewParamsCommit) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsCommit](
		"type", "PREPAID", "POSTPAID",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsCommit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
//
// The property ScheduleItems is required.
type V1ContractNewParamsCommitAccessSchedule struct {
	ScheduleItems []V1ContractNewParamsCommitAccessScheduleScheduleItem `json:"schedule_items,omitzero,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r V1ContractNewParamsCommitAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCommitAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCommitAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, EndingBefore, StartingAt are required.
type V1ContractNewParamsCommitAccessScheduleScheduleItem struct {
	Amount float64 `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	paramObj
}

func (r V1ContractNewParamsCommitAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCommitAccessScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCommitAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match access_schedule
// amount. Optional for "PREPAID" commits: if not provided, this will be a
// "complimentary" commit with no invoice.
type V1ContractNewParamsCommitInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Opt[bool] `json:"do_not_invoice,omitzero"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V1ContractNewParamsCommitInvoiceScheduleRecurringSchedule `json:"recurring_schedule,omitzero"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V1ContractNewParamsCommitInvoiceScheduleScheduleItem `json:"schedule_items,omitzero"`
	paramObj
}

func (r V1ContractNewParamsCommitInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCommitInvoiceSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
//
// The properties AmountDistribution, EndingBefore, Frequency, StartingAt are
// required.
type V1ContractNewParamsCommitInvoiceScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution,omitzero,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL".
	Frequency string `json:"frequency,omitzero,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractNewParamsCommitInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCommitInvoiceScheduleRecurringSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCommitInvoiceScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsCommitInvoiceScheduleRecurringSchedule](
		"amount_distribution", "DIVIDED", "DIVIDED_ROUNDED", "EACH",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsCommitInvoiceScheduleRecurringSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL",
	)
}

// The property Timestamp is required.
type V1ContractNewParamsCommitInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractNewParamsCommitInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCommitInvoiceScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCommitInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// optionally payment gate this commit
//
// The property PaymentGateType is required.
type V1ContractNewParamsCommitPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	//
	// Any of "NONE", "STRIPE", "EXTERNAL".
	PaymentGateType string `json:"payment_gate_type,omitzero,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig V1ContractNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config,omitzero"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig V1ContractNewParamsCommitPaymentGateConfigStripeConfig `json:"stripe_config,omitzero"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	//
	// Any of "NONE", "STRIPE", "ANROK", "PRECALCULATED".
	TaxType string `json:"tax_type,omitzero"`
	paramObj
}

func (r V1ContractNewParamsCommitPaymentGateConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCommitPaymentGateConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCommitPaymentGateConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsCommitPaymentGateConfig](
		"payment_gate_type", "NONE", "STRIPE", "EXTERNAL",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsCommitPaymentGateConfig](
		"tax_type", "NONE", "STRIPE", "ANROK", "PRECALCULATED",
	)
}

// Only applicable if using PRECALCULATED as your tax type.
//
// The property TaxAmount is required.
type V1ContractNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Opt[string] `json:"tax_name,omitzero"`
	paramObj
}

func (r V1ContractNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only applicable if using STRIPE as your payment gate type.
//
// The property PaymentType is required.
type V1ContractNewParamsCommitPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	//
	// Any of "INVOICE", "PAYMENT_INTENT".
	PaymentType string `json:"payment_type,omitzero,required"`
	// If true, the payment will be made assuming the customer is present (i.e. on
	// session).
	//
	// If false, the payment will be made assuming the customer is not present (i.e.
	// off session). For cardholders from a country with an e-mandate requirement (e.g.
	// India), the payment may be declined.
	//
	// If left blank, will default to false.
	OnSessionPayment param.Opt[bool] `json:"on_session_payment,omitzero"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string `json:"invoice_metadata,omitzero"`
	paramObj
}

func (r V1ContractNewParamsCommitPaymentGateConfigStripeConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCommitPaymentGateConfigStripeConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCommitPaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsCommitPaymentGateConfigStripeConfig](
		"payment_type", "INVOICE", "PAYMENT_INTENT",
	)
}

// The properties AccessSchedule, ProductID are required.
type V1ContractNewParamsCredit struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule V1ContractNewParamsCreditAccessSchedule `json:"access_schedule,omitzero,required"`
	ProductID      string                                  `json:"product_id,required" format:"uuid"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Opt[string] `json:"description,omitzero"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Which products the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam `json:"hierarchy_configuration,omitzero"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	paramObj
}

func (r V1ContractNewParamsCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsCredit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
}

// Schedule for distributing the credit to the customer.
//
// The property ScheduleItems is required.
type V1ContractNewParamsCreditAccessSchedule struct {
	ScheduleItems []V1ContractNewParamsCreditAccessScheduleScheduleItem `json:"schedule_items,omitzero,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r V1ContractNewParamsCreditAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCreditAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCreditAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, EndingBefore, StartingAt are required.
type V1ContractNewParamsCreditAccessScheduleScheduleItem struct {
	Amount float64 `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	paramObj
}

func (r V1ContractNewParamsCreditAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsCreditAccessScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsCreditAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ProductID, Schedule are required.
type V1ContractNewParamsDiscount struct {
	ProductID string `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule V1ContractNewParamsDiscountSchedule `json:"schedule,omitzero,required"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V1ContractNewParamsDiscount) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsDiscount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must provide either schedule_items or recurring_schedule.
type V1ContractNewParamsDiscountSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Opt[bool] `json:"do_not_invoice,omitzero"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V1ContractNewParamsDiscountScheduleRecurringSchedule `json:"recurring_schedule,omitzero"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V1ContractNewParamsDiscountScheduleScheduleItem `json:"schedule_items,omitzero"`
	paramObj
}

func (r V1ContractNewParamsDiscountSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsDiscountSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsDiscountSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
//
// The properties AmountDistribution, EndingBefore, Frequency, StartingAt are
// required.
type V1ContractNewParamsDiscountScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution,omitzero,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL".
	Frequency string `json:"frequency,omitzero,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractNewParamsDiscountScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsDiscountScheduleRecurringSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsDiscountScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsDiscountScheduleRecurringSchedule](
		"amount_distribution", "DIVIDED", "DIVIDED_ROUNDED", "EACH",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsDiscountScheduleRecurringSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL",
	)
}

// The property Timestamp is required.
type V1ContractNewParamsDiscountScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractNewParamsDiscountScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsDiscountScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsDiscountScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Parent is required.
type V1ContractNewParamsHierarchyConfiguration struct {
	Parent V1ContractNewParamsHierarchyConfigurationParent `json:"parent,omitzero,required"`
	paramObj
}

func (r V1ContractNewParamsHierarchyConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsHierarchyConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsHierarchyConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContractID, CustomerID are required.
type V1ContractNewParamsHierarchyConfigurationParent struct {
	ContractID string `json:"contract_id,required" format:"uuid"`
	CustomerID string `json:"customer_id,required" format:"uuid"`
	paramObj
}

func (r V1ContractNewParamsHierarchyConfigurationParent) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsHierarchyConfigurationParent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsHierarchyConfigurationParent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
// prices automatically. EXPLICIT prioritization requires specifying priorities for
// each multiplier; the one with the lowest priority value will be prioritized
// first. If tiered overrides are used, prioritization must be explicit.
type V1ContractNewParamsMultiplierOverridePrioritization string

const (
	V1ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier V1ContractNewParamsMultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	V1ContractNewParamsMultiplierOverridePrioritizationExplicit         V1ContractNewParamsMultiplierOverridePrioritization = "EXPLICIT"
)

// The property StartingAt is required.
type V1ContractNewParamsOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	Entitled     param.Opt[bool]      `json:"entitled,omitzero"`
	// Indicates whether the override should only apply to commits. Defaults to
	// `false`. If `true`, you can specify relevant commits in `override_specifiers` by
	// passing `commit_ids`. if you do not specify `commit_ids`, then the override will
	// apply when consuming any prepaid or postpaid commit.
	IsCommitSpecific param.Opt[bool] `json:"is_commit_specific,omitzero"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Opt[float64] `json:"multiplier,omitzero"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides.
	// Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered
	// and multiplier overrides are prioritized by their priority value (lowest first).
	// Must be > 0.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// ID of the product whose rate is being overridden. Cannot be used in conjunction
	// with override_specifiers.
	ProductID param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	// tags identifying products whose rates are being overridden. Cannot be used in
	// conjunction with override_specifiers.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If
	// provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers []V1ContractNewParamsOverrideOverrideSpecifier `json:"override_specifiers,omitzero"`
	// Required for OVERWRITE type.
	OverwriteRate V1ContractNewParamsOverrideOverwriteRate `json:"overwrite_rate,omitzero"`
	// Indicates whether the override applies to commit rates or list rates. Can only
	// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
	// `"LIST_RATE"`.
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	Target string `json:"target,omitzero"`
	// Required for TIERED type. Must have at least one tier.
	Tiers []V1ContractNewParamsOverrideTier `json:"tiers,omitzero"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	//
	// Any of "OVERWRITE", "MULTIPLIER", "TIERED".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r V1ContractNewParamsOverride) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsOverride](
		"target", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsOverride](
		"type", "OVERWRITE", "MULTIPLIER", "TIERED",
	)
}

type V1ContractNewParamsOverrideOverrideSpecifier struct {
	// If provided, the override will only apply to the product with the specified ID.
	ProductID param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to the
	// specified commits. If not provided, the override will apply to all commits.
	CommitIDs []string `json:"commit_ids,omitzero"`
	// A map of group names to values. The override will only apply to line items with
	// the specified presentation group values.
	PresentationGroupValues map[string]string `json:"presentation_group_values,omitzero"`
	// A map of pricing group names to values. The override will only apply to products
	// with the specified pricing group values.
	PricingGroupValues map[string]string `json:"pricing_group_values,omitzero"`
	// If provided, the override will only apply to products with all the specified
	// tags.
	ProductTags []string `json:"product_tags,omitzero"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to commits
	// created by the specified recurring commit ids.
	RecurringCommitIDs []string `json:"recurring_commit_ids,omitzero"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to credits
	// created by the specified recurring credit ids.
	RecurringCreditIDs []string `json:"recurring_credit_ids,omitzero"`
	paramObj
}

func (r V1ContractNewParamsOverrideOverrideSpecifier) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsOverrideOverrideSpecifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsOverrideOverrideSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsOverrideOverrideSpecifier](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// Required for OVERWRITE type.
//
// The property RateType is required.
type V1ContractNewParamsOverrideOverwriteRate struct {
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM".
	RateType     string            `json:"rate_type,omitzero,required"`
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated param.Opt[bool] `json:"is_prorated,omitzero"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price param.Opt[float64] `json:"price,omitzero"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]any `json:"custom_rate,omitzero"`
	// Only set for TIERED rate_type.
	Tiers []shared.TierParam `json:"tiers,omitzero"`
	paramObj
}

func (r V1ContractNewParamsOverrideOverwriteRate) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsOverrideOverwriteRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsOverrideOverwriteRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsOverrideOverwriteRate](
		"rate_type", "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM",
	)
}

// The property Multiplier is required.
type V1ContractNewParamsOverrideTier struct {
	Multiplier float64            `json:"multiplier,required"`
	Size       param.Opt[float64] `json:"size,omitzero"`
	paramObj
}

func (r V1ContractNewParamsOverrideTier) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsOverrideTier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsOverrideTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxAmount, ProductID, Quantity, UnitPrice are required.
type V1ContractNewParamsProfessionalService struct {
	// Maximum amount for the term.
	MaxAmount float64 `json:"max_amount,required"`
	ProductID string  `json:"product_id,required" format:"uuid"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity float64 `json:"quantity,required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified.
	UnitPrice   float64           `json:"unit_price,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V1ContractNewParamsProfessionalService) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsProfessionalService
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsProfessionalService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessAmount, CommitDuration, Priority, ProductID, StartingAt are
// required.
type V1ContractNewParamsRecurringCommit struct {
	// The amount of commit to grant.
	AccessAmount V1ContractNewParamsRecurringCommitAccessAmount `json:"access_amount,omitzero,required"`
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
	CommitDuration V1ContractNewParamsRecurringCommitCommitDuration `json:"commit_duration,omitzero,required"`
	// Will be passed down to the individual commits
	Priority  float64 `json:"priority,required"`
	ProductID string  `json:"product_id,required" format:"uuid"`
	// determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	Description param.Opt[string] `json:"description,omitzero"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// displayed on invoices. will be passed through to the individual commits
	Name param.Opt[string] `json:"name,omitzero"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction param.Opt[float64] `json:"rollover_fraction,omitzero"`
	// A temporary ID that can be used to reference the recurring commit for commit
	// specific overrides.
	TemporaryID param.Opt[string] `json:"temporary_id,omitzero"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam `json:"hierarchy_configuration,omitzero"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount V1ContractNewParamsRecurringCommitInvoiceAmount `json:"invoice_amount,omitzero"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration,omitzero"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	RecurrenceFrequency string `json:"recurrence_frequency,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig V1ContractNewParamsRecurringCommitSubscriptionConfig `json:"subscription_config,omitzero"`
	paramObj
}

func (r V1ContractNewParamsRecurringCommit) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsRecurringCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsRecurringCommit](
		"proration", "NONE", "FIRST", "LAST", "FIRST_AND_LAST",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsRecurringCommit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsRecurringCommit](
		"recurrence_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// The amount of commit to grant.
//
// The properties CreditTypeID, UnitPrice are required.
type V1ContractNewParamsRecurringCommitAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
	// This field is required unless a subscription is attached via
	// `subscription_config`.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	paramObj
}

func (r V1ContractNewParamsRecurringCommitAccessAmount) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsRecurringCommitAccessAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the length of the access schedule for each created commit/credit. The
// value represents the number of units. Unit defaults to "PERIODS", where the
// length of a period is determined by the recurrence_frequency.
//
// The property Value is required.
type V1ContractNewParamsRecurringCommitCommitDuration struct {
	Value float64 `json:"value,required"`
	// Any of "PERIODS".
	Unit string `json:"unit,omitzero"`
	paramObj
}

func (r V1ContractNewParamsRecurringCommitCommitDuration) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsRecurringCommitCommitDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsRecurringCommitCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsRecurringCommitCommitDuration](
		"unit", "PERIODS",
	)
}

// The amount the customer should be billed for the commit. Not required.
//
// The properties CreditTypeID, Quantity, UnitPrice are required.
type V1ContractNewParamsRecurringCommitInvoiceAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64 `json:"quantity,required"`
	UnitPrice    float64 `json:"unit_price,required"`
	paramObj
}

func (r V1ContractNewParamsRecurringCommitInvoiceAmount) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsRecurringCommitInvoiceAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attach a subscription to the recurring commit/credit.
//
// The properties ApplySeatIncreaseConfig, SubscriptionID are required.
type V1ContractNewParamsRecurringCommitSubscriptionConfig struct {
	ApplySeatIncreaseConfig V1ContractNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,omitzero,required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID string `json:"subscription_id,required"`
	// If set to POOLED, allocation added per seat is pooled across the account.
	//
	// Any of "POOLED".
	Allocation string `json:"allocation,omitzero"`
	paramObj
}

func (r V1ContractNewParamsRecurringCommitSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsRecurringCommitSubscriptionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsRecurringCommitSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsRecurringCommitSubscriptionConfig](
		"allocation", "POOLED",
	)
}

// The property IsProrated is required.
type V1ContractNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated,required"`
	paramObj
}

func (r V1ContractNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessAmount, CommitDuration, Priority, ProductID, StartingAt are
// required.
type V1ContractNewParamsRecurringCredit struct {
	// The amount of commit to grant.
	AccessAmount V1ContractNewParamsRecurringCreditAccessAmount `json:"access_amount,omitzero,required"`
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
	CommitDuration V1ContractNewParamsRecurringCreditCommitDuration `json:"commit_duration,omitzero,required"`
	// Will be passed down to the individual commits
	Priority  float64 `json:"priority,required"`
	ProductID string  `json:"product_id,required" format:"uuid"`
	// determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	Description param.Opt[string] `json:"description,omitzero"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// displayed on invoices. will be passed through to the individual commits
	Name param.Opt[string] `json:"name,omitzero"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction param.Opt[float64] `json:"rollover_fraction,omitzero"`
	// A temporary ID that can be used to reference the recurring commit for commit
	// specific overrides.
	TemporaryID param.Opt[string] `json:"temporary_id,omitzero"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam `json:"hierarchy_configuration,omitzero"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration,omitzero"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	RecurrenceFrequency string `json:"recurrence_frequency,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig V1ContractNewParamsRecurringCreditSubscriptionConfig `json:"subscription_config,omitzero"`
	paramObj
}

func (r V1ContractNewParamsRecurringCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsRecurringCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsRecurringCredit](
		"proration", "NONE", "FIRST", "LAST", "FIRST_AND_LAST",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsRecurringCredit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsRecurringCredit](
		"recurrence_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// The amount of commit to grant.
//
// The properties CreditTypeID, UnitPrice are required.
type V1ContractNewParamsRecurringCreditAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
	// This field is required unless a subscription is attached via
	// `subscription_config`.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	paramObj
}

func (r V1ContractNewParamsRecurringCreditAccessAmount) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsRecurringCreditAccessAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the length of the access schedule for each created commit/credit. The
// value represents the number of units. Unit defaults to "PERIODS", where the
// length of a period is determined by the recurrence_frequency.
//
// The property Value is required.
type V1ContractNewParamsRecurringCreditCommitDuration struct {
	Value float64 `json:"value,required"`
	// Any of "PERIODS".
	Unit string `json:"unit,omitzero"`
	paramObj
}

func (r V1ContractNewParamsRecurringCreditCommitDuration) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsRecurringCreditCommitDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsRecurringCreditCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsRecurringCreditCommitDuration](
		"unit", "PERIODS",
	)
}

// Attach a subscription to the recurring commit/credit.
//
// The properties ApplySeatIncreaseConfig, SubscriptionID are required.
type V1ContractNewParamsRecurringCreditSubscriptionConfig struct {
	ApplySeatIncreaseConfig V1ContractNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,omitzero,required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID string `json:"subscription_id,required"`
	// If set to POOLED, allocation added per seat is pooled across the account.
	//
	// Any of "POOLED".
	Allocation string `json:"allocation,omitzero"`
	paramObj
}

func (r V1ContractNewParamsRecurringCreditSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsRecurringCreditSubscriptionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsRecurringCreditSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsRecurringCreditSubscriptionConfig](
		"allocation", "POOLED",
	)
}

// The property IsProrated is required.
type V1ContractNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated,required"`
	paramObj
}

func (r V1ContractNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Fraction, NetsuiteResellerID, ResellerType, StartingAt are
// required.
type V1ContractNewParamsResellerRoyalty struct {
	Fraction           float64 `json:"fraction,required"`
	NetsuiteResellerID string  `json:"netsuite_reseller_id,required"`
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType          string               `json:"reseller_type,omitzero,required"`
	StartingAt            time.Time            `json:"starting_at,required" format:"date-time"`
	EndingBefore          param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	ResellerContractValue param.Opt[float64]   `json:"reseller_contract_value,omitzero"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags []string                                     `json:"applicable_product_tags,omitzero"`
	AwsOptions            V1ContractNewParamsResellerRoyaltyAwsOptions `json:"aws_options,omitzero"`
	GcpOptions            V1ContractNewParamsResellerRoyaltyGcpOptions `json:"gcp_options,omitzero"`
	paramObj
}

func (r V1ContractNewParamsResellerRoyalty) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsResellerRoyalty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsResellerRoyalty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsResellerRoyalty](
		"reseller_type", "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE",
	)
}

type V1ContractNewParamsResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    param.Opt[string] `json:"aws_account_number,omitzero"`
	AwsOfferID          param.Opt[string] `json:"aws_offer_id,omitzero"`
	AwsPayerReferenceID param.Opt[string] `json:"aws_payer_reference_id,omitzero"`
	paramObj
}

func (r V1ContractNewParamsResellerRoyaltyAwsOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsResellerRoyaltyAwsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsResellerRoyaltyAwsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractNewParamsResellerRoyaltyGcpOptions struct {
	GcpAccountID param.Opt[string] `json:"gcp_account_id,omitzero"`
	GcpOfferID   param.Opt[string] `json:"gcp_offer_id,omitzero"`
	paramObj
}

func (r V1ContractNewParamsResellerRoyaltyGcpOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsResellerRoyaltyGcpOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsResellerRoyaltyGcpOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ProductID, Schedule are required.
type V1ContractNewParamsScheduledCharge struct {
	ProductID string `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule V1ContractNewParamsScheduledChargeSchedule `json:"schedule,omitzero,required"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V1ContractNewParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsScheduledCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must provide either schedule_items or recurring_schedule.
type V1ContractNewParamsScheduledChargeSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Opt[bool] `json:"do_not_invoice,omitzero"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V1ContractNewParamsScheduledChargeScheduleRecurringSchedule `json:"recurring_schedule,omitzero"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V1ContractNewParamsScheduledChargeScheduleScheduleItem `json:"schedule_items,omitzero"`
	paramObj
}

func (r V1ContractNewParamsScheduledChargeSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsScheduledChargeSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsScheduledChargeSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
//
// The properties AmountDistribution, EndingBefore, Frequency, StartingAt are
// required.
type V1ContractNewParamsScheduledChargeScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution,omitzero,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL".
	Frequency string `json:"frequency,omitzero,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractNewParamsScheduledChargeScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsScheduledChargeScheduleRecurringSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsScheduledChargeScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsScheduledChargeScheduleRecurringSchedule](
		"amount_distribution", "DIVIDED", "DIVIDED_ROUNDED", "EACH",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsScheduledChargeScheduleRecurringSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL",
	)
}

// The property Timestamp is required.
type V1ContractNewParamsScheduledChargeScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractNewParamsScheduledChargeScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsScheduledChargeScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsScheduledChargeScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type V1ContractNewParamsScheduledChargesOnUsageInvoices string

const (
	V1ContractNewParamsScheduledChargesOnUsageInvoicesAll V1ContractNewParamsScheduledChargesOnUsageInvoices = "ALL"
)

// The properties CollectionSchedule, Proration, SubscriptionRate are required.
type V1ContractNewParamsSubscription struct {
	// Any of "ADVANCE", "ARREARS".
	CollectionSchedule string                                          `json:"collection_schedule,omitzero,required"`
	Proration          V1ContractNewParamsSubscriptionProration        `json:"proration,omitzero,required"`
	SubscriptionRate   V1ContractNewParamsSubscriptionSubscriptionRate `json:"subscription_rate,omitzero,required"`
	Description        param.Opt[string]                               `json:"description,omitzero"`
	// Exclusive end time for the subscription. If not provided, subscription inherits
	// contract end date.
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// The initial quantity for the subscription. It must be non-negative value.
	// Required if quantity_management_mode is QUANTITY_ONLY.
	InitialQuantity param.Opt[float64] `json:"initial_quantity,omitzero"`
	Name            param.Opt[string]  `json:"name,omitzero"`
	// Inclusive start time for the subscription. If not provided, defaults to contract
	// start date
	StartingAt param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	// A temporary ID used to reference the subscription in recurring commit/credit
	// subscription configs created within the same payload.
	TemporaryID param.Opt[string] `json:"temporary_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// Determines how the subscription's quantity is controlled. Defaults to
	// QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified
	// directly on the subscription. `initial_quantity` must be provided with this
	// option. Compatible with recurring commits/credits that use POOLED allocation.
	//
	// Any of "SEAT_BASED", "QUANTITY_ONLY".
	QuantityManagementMode string `json:"quantity_management_mode,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscription) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscription](
		"collection_schedule", "ADVANCE", "ARREARS",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscription](
		"quantity_management_mode", "SEAT_BASED", "QUANTITY_ONLY",
	)
}

type V1ContractNewParamsSubscriptionProration struct {
	// Indicates if the partial period will be prorated or charged a full amount.
	IsProrated param.Opt[bool] `json:"is_prorated,omitzero"`
	// Indicates how mid-period quantity adjustments are invoiced.
	// **BILL_IMMEDIATELY**: Only available when collection schedule is `ADVANCE`. The
	// quantity increase will be billed immediately on the scheduled date.
	// **BILL_ON_NEXT_COLLECTION_DATE**: The quantity increase will be billed for
	// in-arrears at the end of the period.
	//
	// Any of "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE".
	InvoiceBehavior string `json:"invoice_behavior,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionProration) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionProration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionProration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionProration](
		"invoice_behavior", "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE",
	)
}

// The properties BillingFrequency, ProductID are required.
type V1ContractNewParamsSubscriptionSubscriptionRate struct {
	// Frequency to bill subscription with. Together with product_id, must match
	// existing rate on the rate card.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero,required"`
	// Must be subscription type product
	ProductID string `json:"product_id,required" format:"uuid"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionSubscriptionRate) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionSubscriptionRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionSubscriptionRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionSubscriptionRate](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// The properties FromContractID, Type are required.
type V1ContractNewParamsTransition struct {
	FromContractID string `json:"from_contract_id,required" format:"uuid"`
	// This field's available values may vary based on your client's configuration.
	//
	// Any of "SUPERSEDE", "RENEWAL".
	Type                  string                                             `json:"type,omitzero,required"`
	FutureInvoiceBehavior V1ContractNewParamsTransitionFutureInvoiceBehavior `json:"future_invoice_behavior,omitzero"`
	paramObj
}

func (r V1ContractNewParamsTransition) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsTransition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsTransition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsTransition](
		"type", "SUPERSEDE", "RENEWAL",
	)
}

type V1ContractNewParamsTransitionFutureInvoiceBehavior struct {
	// Controls whether future trueup invoices are billed or removed. Default behavior
	// is AS_IS if not specified.
	//
	// Any of "REMOVE", "AS_IS".
	Trueup string `json:"trueup,omitzero"`
	paramObj
}

func (r V1ContractNewParamsTransitionFutureInvoiceBehavior) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsTransitionFutureInvoiceBehavior
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsTransitionFutureInvoiceBehavior) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsTransitionFutureInvoiceBehavior](
		"trueup", "REMOVE", "AS_IS",
	)
}

// The property Frequency is required.
type V1ContractNewParamsUsageStatementSchedule struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency,omitzero,required"`
	// Required when using CUSTOM_DATE. This option lets you set a historical billing
	// anchor date, aligning future billing cycles with a chosen cadence. For example,
	// if a contract starts on 2024-09-15 and you set the anchor date to 2024-09-10
	// with a MONTHLY frequency, the first usage statement will cover 09-15 to 10-10.
	// Subsequent statements will follow the 10th of each month.
	BillingAnchorDate param.Opt[time.Time] `json:"billing_anchor_date,omitzero" format:"date-time"`
	// The date Metronome should start generating usage invoices. If unspecified,
	// contract start date will be used. This is useful to set if you want to import
	// historical invoices via our 'Create Historical Invoices' API rather than having
	// Metronome automatically generate them.
	InvoiceGenerationStartingAt param.Opt[time.Time] `json:"invoice_generation_starting_at,omitzero" format:"date-time"`
	// If not provided, defaults to the first day of the month.
	//
	// Any of "FIRST_OF_MONTH", "CONTRACT_START", "CUSTOM_DATE".
	Day string `json:"day,omitzero"`
	paramObj
}

func (r V1ContractNewParamsUsageStatementSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsUsageStatementSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsUsageStatementSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsUsageStatementSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsUsageStatementSchedule](
		"day", "FIRST_OF_MONTH", "CONTRACT_START", "CUSTOM_DATE",
	)
}

type V1ContractGetParams struct {
	ContractID string `json:"contract_id,required" format:"uuid"`
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the query to be slower.
	IncludeBalance param.Opt[bool] `json:"include_balance,omitzero"`
	// Include commit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Opt[bool] `json:"include_ledgers,omitzero"`
	paramObj
}

func (r V1ContractGetParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractListParams struct {
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// Optional RFC 3339 timestamp. If provided, the response will include only
	// contracts effective on the provided date. This cannot be provided if the
	// starting_at filter is provided.
	CoveringDate param.Opt[time.Time] `json:"covering_date,omitzero" format:"date-time"`
	// Include archived contracts in the response
	IncludeArchived param.Opt[bool] `json:"include_archived,omitzero"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the query to be slower.
	IncludeBalance param.Opt[bool] `json:"include_balance,omitzero"`
	// Include commit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Opt[bool] `json:"include_ledgers,omitzero"`
	// Optional RFC 3339 timestamp. If provided, the response will include only
	// contracts where effective_at is on or after the provided date. This cannot be
	// provided if the covering_date filter is provided.
	StartingAt param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V1ContractListParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractAddManualBalanceEntryParams struct {
	// ID of the balance (commit or credit) to update.
	ID string `json:"id,required" format:"uuid"`
	// Amount to add to the segment. A negative number will draw down from the balance.
	Amount float64 `json:"amount,required"`
	// ID of the customer whose balance is to be updated.
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// Reason for the manual adjustment. This will be displayed in the ledger.
	Reason string `json:"reason,required"`
	// ID of the segment to update.
	SegmentID string `json:"segment_id,required" format:"uuid"`
	// ID of the contract to update. Leave blank to update a customer level balance.
	ContractID param.Opt[string] `json:"contract_id,omitzero" format:"uuid"`
	// RFC 3339 timestamp indicating when the manual adjustment takes place. If not
	// provided, it will default to the start of the segment.
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	paramObj
}

func (r V1ContractAddManualBalanceEntryParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAddManualBalanceEntryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAddManualBalanceEntryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractAmendParams struct {
	// ID of the contract to amend
	ContractID string `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is to be amended
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// inclusive start time for the amendment
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Opt[string] `json:"salesforce_opportunity_id,omitzero"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue param.Opt[float64]            `json:"total_contract_value,omitzero"`
	Commits            []V1ContractAmendParamsCommit `json:"commits,omitzero"`
	Credits            []V1ContractAmendParamsCredit `json:"credits,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// This field's availability is dependent on your client's configuration.
	Discounts         []V1ContractAmendParamsDiscount `json:"discounts,omitzero"`
	ContractOverrides []V1ContractAmendParamsOverride `json:"overrides,omitzero"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []V1ContractAmendParamsProfessionalService `json:"professional_services,omitzero"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []V1ContractAmendParamsResellerRoyalty `json:"reseller_royalties,omitzero"`
	ScheduledCharges  []V1ContractAmendParamsScheduledCharge `json:"scheduled_charges,omitzero"`
	paramObj
}

func (r V1ContractAmendParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ProductID, Type are required.
type V1ContractAmendParamsCommit struct {
	ProductID string `json:"product_id,required" format:"uuid"`
	// Any of "PREPAID", "POSTPAID".
	Type string `json:"type,omitzero,required"`
	// (DEPRECATED) Use access_schedule and invoice_schedule instead.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Opt[string] `json:"description,omitzero"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Opt[float64] `json:"rollover_fraction,omitzero"`
	// A temporary ID for the commit that can be used to reference the commit for
	// commit specific overrides.
	TemporaryID param.Opt[string] `json:"temporary_id,omitzero"`
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
	// commits only one schedule item is allowed and amount must match invoice_schedule
	// total.
	AccessSchedule V1ContractAmendParamsCommitAccessSchedule `json:"access_schedule,omitzero"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam `json:"hierarchy_configuration,omitzero"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this
	// time and only one schedule item is allowed; the total must match access_schedule
	// amount. Optional for "PREPAID" commits: if not provided, this will be a
	// "complimentary" commit with no invoice.
	InvoiceSchedule V1ContractAmendParamsCommitInvoiceSchedule `json:"invoice_schedule,omitzero"`
	// optionally payment gate this commit
	PaymentGateConfig V1ContractAmendParamsCommitPaymentGateConfig `json:"payment_gate_config,omitzero"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsCommit) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractAmendParamsCommit](
		"type", "PREPAID", "POSTPAID",
	)
	apijson.RegisterFieldValidator[V1ContractAmendParamsCommit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
//
// The property ScheduleItems is required.
type V1ContractAmendParamsCommitAccessSchedule struct {
	ScheduleItems []V1ContractAmendParamsCommitAccessScheduleScheduleItem `json:"schedule_items,omitzero,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r V1ContractAmendParamsCommitAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCommitAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCommitAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, EndingBefore, StartingAt are required.
type V1ContractAmendParamsCommitAccessScheduleScheduleItem struct {
	Amount float64 `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	paramObj
}

func (r V1ContractAmendParamsCommitAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCommitAccessScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCommitAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match access_schedule
// amount. Optional for "PREPAID" commits: if not provided, this will be a
// "complimentary" commit with no invoice.
type V1ContractAmendParamsCommitInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Opt[bool] `json:"do_not_invoice,omitzero"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V1ContractAmendParamsCommitInvoiceScheduleRecurringSchedule `json:"recurring_schedule,omitzero"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V1ContractAmendParamsCommitInvoiceScheduleScheduleItem `json:"schedule_items,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsCommitInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCommitInvoiceSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
//
// The properties AmountDistribution, EndingBefore, Frequency, StartingAt are
// required.
type V1ContractAmendParamsCommitInvoiceScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution,omitzero,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL".
	Frequency string `json:"frequency,omitzero,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsCommitInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCommitInvoiceScheduleRecurringSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCommitInvoiceScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractAmendParamsCommitInvoiceScheduleRecurringSchedule](
		"amount_distribution", "DIVIDED", "DIVIDED_ROUNDED", "EACH",
	)
	apijson.RegisterFieldValidator[V1ContractAmendParamsCommitInvoiceScheduleRecurringSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL",
	)
}

// The property Timestamp is required.
type V1ContractAmendParamsCommitInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsCommitInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCommitInvoiceScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCommitInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// optionally payment gate this commit
//
// The property PaymentGateType is required.
type V1ContractAmendParamsCommitPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	//
	// Any of "NONE", "STRIPE", "EXTERNAL".
	PaymentGateType string `json:"payment_gate_type,omitzero,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig V1ContractAmendParamsCommitPaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config,omitzero"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig V1ContractAmendParamsCommitPaymentGateConfigStripeConfig `json:"stripe_config,omitzero"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	//
	// Any of "NONE", "STRIPE", "ANROK", "PRECALCULATED".
	TaxType string `json:"tax_type,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsCommitPaymentGateConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCommitPaymentGateConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCommitPaymentGateConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractAmendParamsCommitPaymentGateConfig](
		"payment_gate_type", "NONE", "STRIPE", "EXTERNAL",
	)
	apijson.RegisterFieldValidator[V1ContractAmendParamsCommitPaymentGateConfig](
		"tax_type", "NONE", "STRIPE", "ANROK", "PRECALCULATED",
	)
}

// Only applicable if using PRECALCULATED as your tax type.
//
// The property TaxAmount is required.
type V1ContractAmendParamsCommitPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Opt[string] `json:"tax_name,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsCommitPaymentGateConfigPrecalculatedTaxConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCommitPaymentGateConfigPrecalculatedTaxConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCommitPaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only applicable if using STRIPE as your payment gate type.
//
// The property PaymentType is required.
type V1ContractAmendParamsCommitPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	//
	// Any of "INVOICE", "PAYMENT_INTENT".
	PaymentType string `json:"payment_type,omitzero,required"`
	// If true, the payment will be made assuming the customer is present (i.e. on
	// session).
	//
	// If false, the payment will be made assuming the customer is not present (i.e.
	// off session). For cardholders from a country with an e-mandate requirement (e.g.
	// India), the payment may be declined.
	//
	// If left blank, will default to false.
	OnSessionPayment param.Opt[bool] `json:"on_session_payment,omitzero"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string `json:"invoice_metadata,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsCommitPaymentGateConfigStripeConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCommitPaymentGateConfigStripeConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCommitPaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractAmendParamsCommitPaymentGateConfigStripeConfig](
		"payment_type", "INVOICE", "PAYMENT_INTENT",
	)
}

// The properties AccessSchedule, ProductID are required.
type V1ContractAmendParamsCredit struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule V1ContractAmendParamsCreditAccessSchedule `json:"access_schedule,omitzero,required"`
	ProductID      string                                    `json:"product_id,required" format:"uuid"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Opt[string] `json:"description,omitzero"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Which products the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam `json:"hierarchy_configuration,omitzero"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractAmendParamsCredit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
}

// Schedule for distributing the credit to the customer.
//
// The property ScheduleItems is required.
type V1ContractAmendParamsCreditAccessSchedule struct {
	ScheduleItems []V1ContractAmendParamsCreditAccessScheduleScheduleItem `json:"schedule_items,omitzero,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r V1ContractAmendParamsCreditAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCreditAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCreditAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, EndingBefore, StartingAt are required.
type V1ContractAmendParamsCreditAccessScheduleScheduleItem struct {
	Amount float64 `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	paramObj
}

func (r V1ContractAmendParamsCreditAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsCreditAccessScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsCreditAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ProductID, Schedule are required.
type V1ContractAmendParamsDiscount struct {
	ProductID string `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule V1ContractAmendParamsDiscountSchedule `json:"schedule,omitzero,required"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsDiscount) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsDiscount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must provide either schedule_items or recurring_schedule.
type V1ContractAmendParamsDiscountSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Opt[bool] `json:"do_not_invoice,omitzero"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V1ContractAmendParamsDiscountScheduleRecurringSchedule `json:"recurring_schedule,omitzero"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V1ContractAmendParamsDiscountScheduleScheduleItem `json:"schedule_items,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsDiscountSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsDiscountSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsDiscountSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
//
// The properties AmountDistribution, EndingBefore, Frequency, StartingAt are
// required.
type V1ContractAmendParamsDiscountScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution,omitzero,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL".
	Frequency string `json:"frequency,omitzero,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsDiscountScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsDiscountScheduleRecurringSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsDiscountScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractAmendParamsDiscountScheduleRecurringSchedule](
		"amount_distribution", "DIVIDED", "DIVIDED_ROUNDED", "EACH",
	)
	apijson.RegisterFieldValidator[V1ContractAmendParamsDiscountScheduleRecurringSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL",
	)
}

// The property Timestamp is required.
type V1ContractAmendParamsDiscountScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsDiscountScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsDiscountScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsDiscountScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property StartingAt is required.
type V1ContractAmendParamsOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	Entitled     param.Opt[bool]      `json:"entitled,omitzero"`
	// Indicates whether the override should only apply to commits. Defaults to
	// `false`. If `true`, you can specify relevant commits in `override_specifiers` by
	// passing `commit_ids`. if you do not specify `commit_ids`, then the override will
	// apply when consuming any prepaid or postpaid commit.
	IsCommitSpecific param.Opt[bool] `json:"is_commit_specific,omitzero"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Opt[float64] `json:"multiplier,omitzero"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides.
	// Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered
	// and multiplier overrides are prioritized by their priority value (lowest first).
	// Must be > 0.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// ID of the product whose rate is being overridden. Cannot be used in conjunction
	// with override_specifiers.
	ProductID param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	// tags identifying products whose rates are being overridden. Cannot be used in
	// conjunction with override_specifiers.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If
	// provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers []V1ContractAmendParamsOverrideOverrideSpecifier `json:"override_specifiers,omitzero"`
	// Required for OVERWRITE type.
	OverwriteRate V1ContractAmendParamsOverrideOverwriteRate `json:"overwrite_rate,omitzero"`
	// Indicates whether the override applies to commit rates or list rates. Can only
	// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
	// `"LIST_RATE"`.
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	Target string `json:"target,omitzero"`
	// Required for TIERED type. Must have at least one tier.
	Tiers []V1ContractAmendParamsOverrideTier `json:"tiers,omitzero"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	//
	// Any of "OVERWRITE", "MULTIPLIER", "TIERED".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsOverride) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractAmendParamsOverride](
		"target", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[V1ContractAmendParamsOverride](
		"type", "OVERWRITE", "MULTIPLIER", "TIERED",
	)
}

type V1ContractAmendParamsOverrideOverrideSpecifier struct {
	// If provided, the override will only apply to the product with the specified ID.
	ProductID param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to the
	// specified commits. If not provided, the override will apply to all commits.
	CommitIDs []string `json:"commit_ids,omitzero"`
	// A map of group names to values. The override will only apply to line items with
	// the specified presentation group values.
	PresentationGroupValues map[string]string `json:"presentation_group_values,omitzero"`
	// A map of pricing group names to values. The override will only apply to products
	// with the specified pricing group values.
	PricingGroupValues map[string]string `json:"pricing_group_values,omitzero"`
	// If provided, the override will only apply to products with all the specified
	// tags.
	ProductTags []string `json:"product_tags,omitzero"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to commits
	// created by the specified recurring commit ids.
	RecurringCommitIDs []string `json:"recurring_commit_ids,omitzero"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to credits
	// created by the specified recurring credit ids.
	RecurringCreditIDs []string `json:"recurring_credit_ids,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsOverrideOverrideSpecifier) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsOverrideOverrideSpecifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsOverrideOverrideSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractAmendParamsOverrideOverrideSpecifier](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// Required for OVERWRITE type.
//
// The property RateType is required.
type V1ContractAmendParamsOverrideOverwriteRate struct {
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM".
	RateType     string            `json:"rate_type,omitzero,required"`
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated param.Opt[bool] `json:"is_prorated,omitzero"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price param.Opt[float64] `json:"price,omitzero"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]any `json:"custom_rate,omitzero"`
	// Only set for TIERED rate_type.
	Tiers []shared.TierParam `json:"tiers,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsOverrideOverwriteRate) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsOverrideOverwriteRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsOverrideOverwriteRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractAmendParamsOverrideOverwriteRate](
		"rate_type", "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM",
	)
}

// The property Multiplier is required.
type V1ContractAmendParamsOverrideTier struct {
	Multiplier float64            `json:"multiplier,required"`
	Size       param.Opt[float64] `json:"size,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsOverrideTier) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsOverrideTier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsOverrideTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxAmount, ProductID, Quantity, UnitPrice are required.
type V1ContractAmendParamsProfessionalService struct {
	// Maximum amount for the term.
	MaxAmount float64 `json:"max_amount,required"`
	ProductID string  `json:"product_id,required" format:"uuid"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity float64 `json:"quantity,required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified.
	UnitPrice   float64           `json:"unit_price,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsProfessionalService) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsProfessionalService
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsProfessionalService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ResellerType is required.
type V1ContractAmendParamsResellerRoyalty struct {
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType string `json:"reseller_type,omitzero,required"`
	// Use null to indicate that the existing end timestamp should be removed.
	EndingBefore          param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	Fraction              param.Opt[float64]   `json:"fraction,omitzero"`
	NetsuiteResellerID    param.Opt[string]    `json:"netsuite_reseller_id,omitzero"`
	ResellerContractValue param.Opt[float64]   `json:"reseller_contract_value,omitzero"`
	StartingAt            param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags []string                                       `json:"applicable_product_tags,omitzero"`
	AwsOptions            V1ContractAmendParamsResellerRoyaltyAwsOptions `json:"aws_options,omitzero"`
	GcpOptions            V1ContractAmendParamsResellerRoyaltyGcpOptions `json:"gcp_options,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsResellerRoyalty) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsResellerRoyalty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsResellerRoyalty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractAmendParamsResellerRoyalty](
		"reseller_type", "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE",
	)
}

type V1ContractAmendParamsResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    param.Opt[string] `json:"aws_account_number,omitzero"`
	AwsOfferID          param.Opt[string] `json:"aws_offer_id,omitzero"`
	AwsPayerReferenceID param.Opt[string] `json:"aws_payer_reference_id,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsResellerRoyaltyAwsOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsResellerRoyaltyAwsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsResellerRoyaltyAwsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractAmendParamsResellerRoyaltyGcpOptions struct {
	GcpAccountID param.Opt[string] `json:"gcp_account_id,omitzero"`
	GcpOfferID   param.Opt[string] `json:"gcp_offer_id,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsResellerRoyaltyGcpOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsResellerRoyaltyGcpOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsResellerRoyaltyGcpOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ProductID, Schedule are required.
type V1ContractAmendParamsScheduledCharge struct {
	ProductID string `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule V1ContractAmendParamsScheduledChargeSchedule `json:"schedule,omitzero,required"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsScheduledCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must provide either schedule_items or recurring_schedule.
type V1ContractAmendParamsScheduledChargeSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Opt[bool] `json:"do_not_invoice,omitzero"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V1ContractAmendParamsScheduledChargeScheduleRecurringSchedule `json:"recurring_schedule,omitzero"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V1ContractAmendParamsScheduledChargeScheduleScheduleItem `json:"schedule_items,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsScheduledChargeSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsScheduledChargeSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsScheduledChargeSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
//
// The properties AmountDistribution, EndingBefore, Frequency, StartingAt are
// required.
type V1ContractAmendParamsScheduledChargeScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution,omitzero,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL".
	Frequency string `json:"frequency,omitzero,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsScheduledChargeScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsScheduledChargeScheduleRecurringSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsScheduledChargeScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractAmendParamsScheduledChargeScheduleRecurringSchedule](
		"amount_distribution", "DIVIDED", "DIVIDED_ROUNDED", "EACH",
	)
	apijson.RegisterFieldValidator[V1ContractAmendParamsScheduledChargeScheduleRecurringSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL",
	)
}

// The property Timestamp is required.
type V1ContractAmendParamsScheduledChargeScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractAmendParamsScheduledChargeScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractAmendParamsScheduledChargeScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractAmendParamsScheduledChargeScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractArchiveParams struct {
	// ID of the contract to archive
	ContractID string `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is to be archived
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// If false, the existing finalized invoices will remain after the contract is
	// archived.
	VoidInvoices bool `json:"void_invoices,required"`
	paramObj
}

func (r V1ContractArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractNewHistoricalInvoicesParams struct {
	Invoices []V1ContractNewHistoricalInvoicesParamsInvoice `json:"invoices,omitzero,required"`
	Preview  bool                                           `json:"preview,required"`
	paramObj
}

func (r V1ContractNewHistoricalInvoicesParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewHistoricalInvoicesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewHistoricalInvoicesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContractID, CreditTypeID, CustomerID, ExclusiveEndDate,
// InclusiveStartDate, IssueDate, UsageLineItems are required.
type V1ContractNewHistoricalInvoicesParamsInvoice struct {
	ContractID         string                                                      `json:"contract_id,required" format:"uuid"`
	CreditTypeID       string                                                      `json:"credit_type_id,required" format:"uuid"`
	CustomerID         string                                                      `json:"customer_id,required" format:"uuid"`
	ExclusiveEndDate   time.Time                                                   `json:"exclusive_end_date,required" format:"date-time"`
	InclusiveStartDate time.Time                                                   `json:"inclusive_start_date,required" format:"date-time"`
	IssueDate          time.Time                                                   `json:"issue_date,required" format:"date-time"`
	UsageLineItems     []V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItem `json:"usage_line_items,omitzero,required"`
	// This field's availability is dependent on your client's configuration.
	//
	// Any of "billable", "unbillable".
	BillableStatus string `json:"billable_status,omitzero"`
	// Any of "HOUR", "DAY".
	BreakdownGranularity string `json:"breakdown_granularity,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V1ContractNewHistoricalInvoicesParamsInvoice) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewHistoricalInvoicesParamsInvoice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewHistoricalInvoicesParamsInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewHistoricalInvoicesParamsInvoice](
		"billable_status", "billable", "unbillable",
	)
	apijson.RegisterFieldValidator[V1ContractNewHistoricalInvoicesParamsInvoice](
		"breakdown_granularity", "HOUR", "DAY",
	)
}

// The properties ExclusiveEndDate, InclusiveStartDate, ProductID are required.
type V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItem struct {
	ExclusiveEndDate        time.Time                                                                        `json:"exclusive_end_date,required" format:"date-time"`
	InclusiveStartDate      time.Time                                                                        `json:"inclusive_start_date,required" format:"date-time"`
	ProductID               string                                                                           `json:"product_id,required" format:"uuid"`
	Quantity                param.Opt[float64]                                                               `json:"quantity,omitzero"`
	PresentationGroupValues map[string]string                                                                `json:"presentation_group_values,omitzero"`
	PricingGroupValues      map[string]string                                                                `json:"pricing_group_values,omitzero"`
	SubtotalsWithQuantity   []V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItemSubtotalsWithQuantity `json:"subtotals_with_quantity,omitzero"`
	paramObj
}

func (r V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ExclusiveEndDate, InclusiveStartDate, Quantity are required.
type V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItemSubtotalsWithQuantity struct {
	ExclusiveEndDate   time.Time `json:"exclusive_end_date,required" format:"date-time"`
	InclusiveStartDate time.Time `json:"inclusive_start_date,required" format:"date-time"`
	Quantity           float64   `json:"quantity,required"`
	paramObj
}

func (r V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItemSubtotalsWithQuantity) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItemSubtotalsWithQuantity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewHistoricalInvoicesParamsInvoiceUsageLineItemSubtotalsWithQuantity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractListBalancesParams struct {
	CustomerID string            `json:"customer_id,required" format:"uuid"`
	ID         param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// Return only balances that have access schedules that "cover" the provided date
	CoveringDate param.Opt[time.Time] `json:"covering_date,omitzero" format:"date-time"`
	// Include only balances that have any access before the provided date (exclusive)
	EffectiveBefore param.Opt[time.Time] `json:"effective_before,omitzero" format:"date-time"`
	// Include archived credits and credits from archived contracts.
	IncludeArchived param.Opt[bool] `json:"include_archived,omitzero"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the query to be slower.
	IncludeBalance param.Opt[bool] `json:"include_balance,omitzero"`
	// Include balances on the contract level.
	IncludeContractBalances param.Opt[bool] `json:"include_contract_balances,omitzero"`
	// Include ledgers in the response. Setting this flag may cause the query to be
	// slower.
	IncludeLedgers param.Opt[bool] `json:"include_ledgers,omitzero"`
	// The maximum number of commits to return. Defaults to 25.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// The next page token from a previous response.
	NextPage param.Opt[string] `json:"next_page,omitzero"`
	// Include only balances that have any access on or after the provided date
	StartingAt param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V1ContractListBalancesParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractListBalancesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractListBalancesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractGetRateScheduleParams struct {
	// ID of the contract to get the rate schedule for.
	ContractID string `json:"contract_id,required" format:"uuid"`
	// ID of the customer for whose contract to get the rate schedule for.
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// optional timestamp which overlaps with the returned rate schedule segments. When
	// not specified, the current timestamp will be used.
	At param.Opt[time.Time] `json:"at,omitzero" format:"date-time"`
	// List of rate selectors, rates matching ANY of the selectors will be included in
	// the response. Passing no selectors will result in all rates being returned.
	Selectors []V1ContractGetRateScheduleParamsSelector `json:"selectors,omitzero"`
	paramObj
}

func (r V1ContractGetRateScheduleParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractGetRateScheduleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractGetRateScheduleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1ContractGetRateScheduleParams]'s query parameters as
// `url.Values`.
func (r V1ContractGetRateScheduleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ContractGetRateScheduleParamsSelector struct {
	// Rates matching the product id will be included in the response.
	ProductID param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	// Subscription rates matching the billing frequency will be included in the
	// response.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero"`
	// List of pricing group key value pairs, rates containing the matching key / value
	// pairs will be included in the response.
	PartialPricingGroupValues map[string]string `json:"partial_pricing_group_values,omitzero"`
	// List of pricing group key value pairs, rates matching all of the key / value
	// pairs will be included in the response.
	PricingGroupValues map[string]string `json:"pricing_group_values,omitzero"`
	// List of product tags, rates matching any of the tags will be included in the
	// response.
	ProductTags []string `json:"product_tags,omitzero"`
	paramObj
}

func (r V1ContractGetRateScheduleParamsSelector) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractGetRateScheduleParamsSelector
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractGetRateScheduleParamsSelector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractGetRateScheduleParamsSelector](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

type V1ContractGetSubscriptionQuantityHistoryParams struct {
	ContractID     string `json:"contract_id,required" format:"uuid"`
	CustomerID     string `json:"customer_id,required" format:"uuid"`
	SubscriptionID string `json:"subscription_id,required" format:"uuid"`
	paramObj
}

func (r V1ContractGetSubscriptionQuantityHistoryParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractGetSubscriptionQuantityHistoryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractGetSubscriptionQuantityHistoryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractScheduleProServicesInvoiceParams struct {
	ContractID string `json:"contract_id,required" format:"uuid"`
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// The date the invoice is issued
	IssuedAt time.Time `json:"issued_at,required" format:"date-time"`
	// Each line requires an amount or both unit_price and quantity.
	LineItems []V1ContractScheduleProServicesInvoiceParamsLineItem `json:"line_items,omitzero,required"`
	// The end date of the invoice header in Netsuite
	NetsuiteInvoiceHeaderEnd param.Opt[time.Time] `json:"netsuite_invoice_header_end,omitzero" format:"date-time"`
	// The start date of the invoice header in Netsuite
	NetsuiteInvoiceHeaderStart param.Opt[time.Time] `json:"netsuite_invoice_header_start,omitzero" format:"date-time"`
	paramObj
}

func (r V1ContractScheduleProServicesInvoiceParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractScheduleProServicesInvoiceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractScheduleProServicesInvoiceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes the line item for a professional service charge on an invoice.
//
// The property ProfessionalServiceID is required.
type V1ContractScheduleProServicesInvoiceParamsLineItem struct {
	ProfessionalServiceID string `json:"professional_service_id,required" format:"uuid"`
	// If the professional_service_id was added on an amendment, this is required.
	AmendmentID param.Opt[string] `json:"amendment_id,omitzero" format:"uuid"`
	// Amount for the term on the new invoice.
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// For client use.
	Metadata param.Opt[string] `json:"metadata,omitzero"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd param.Opt[time.Time] `json:"netsuite_invoice_billing_end,omitzero" format:"date-time"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart param.Opt[time.Time] `json:"netsuite_invoice_billing_start,omitzero" format:"date-time"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// If specified, this overrides the unit price on the pro service term. Must also
	// provide quantity (but not amount) if providing unit_price.
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V1ContractScheduleProServicesInvoiceParamsLineItem) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractScheduleProServicesInvoiceParamsLineItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractScheduleProServicesInvoiceParamsLineItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractSetUsageFilterParams struct {
	ContractID  string    `json:"contract_id,required" format:"uuid"`
	CustomerID  string    `json:"customer_id,required" format:"uuid"`
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,omitzero,required"`
	StartingAt  time.Time `json:"starting_at,required" format:"date-time"`
	paramObj
}

func (r V1ContractSetUsageFilterParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractSetUsageFilterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractSetUsageFilterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractUpdateEndDateParams struct {
	// ID of the contract to update
	ContractID string `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is to be updated
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// If true, allows setting the contract end date earlier than the end_timestamp of
	// existing finalized invoices. Finalized invoices will be unchanged; if you want
	// to incorporate the new end date, you can void and regenerate finalized usage
	// invoices. Defaults to true.
	AllowEndingBeforeFinalizedInvoice param.Opt[bool] `json:"allow_ending_before_finalized_invoice,omitzero"`
	// RFC 3339 timestamp indicating when the contract will end (exclusive). If not
	// provided, the contract will be updated to be open-ended.
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	paramObj
}

func (r V1ContractUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractUpdateEndDateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractUpdateEndDateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
