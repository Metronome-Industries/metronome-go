// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v3/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"github.com/Metronome-Industries/metronome-go/v3/packages/param"
	"github.com/Metronome-Industries/metronome-go/v3/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/v3/shared"
)

// V2ContractService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2ContractService] method instead.
type V2ContractService struct {
	Options []option.RequestOption
}

// NewV2ContractService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2ContractService(opts ...option.RequestOption) (r V2ContractService) {
	r = V2ContractService{}
	r.Options = opts
	return
}

// Gets the details for a specific contract, including contract term, rate card
// information, credits and commits, and more.
//
// ### Use this endpoint to:
//
//   - Check the duration of a customer's current contract
//   - Get details on contract terms, including access schedule amounts for
//     commitments and credits
//   - Understand the state of a contract at a past time. As you can evolve the terms
//     of a contract over time through editing, use the `as_of_date` parameter to
//     view the full contract configuration as of that point in time.
//
// ### Usage guidelines:
//
//   - Optionally, use the `include_balance` and `include_ledger` fields to include
//     balances and ledgers in the credit and commit responses. Using these fields
//     will cause the query to be slower.
func (r *V2ContractService) Get(ctx context.Context, body V2ContractGetParams, opts ...option.RequestOption) (res *V2ContractGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/contracts/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// For a given customer, lists all of their contracts in chronological order.
//
// ### Use this endpoint to:
//
//   - Check if a customer is provisioned with any contract, and at which tier
//   - Check the duration and terms of a customer's current contract
//   - Power a page in your end customer experience that shows the customer's history
//     of tiers (e.g. this customer started out on the Pro Plan, then downgraded to
//     the Starter plan).
//
// ### Usage guidelines:
//
// Use the `starting_at`, `covering_date`, and `include_archived` parameters to
// filter the list of returned contracts. For example, to list only currently
// active contracts, pass `covering_date` equal to the current time.
func (r *V2ContractService) List(ctx context.Context, body V2ContractListParams, opts ...option.RequestOption) (res *V2ContractListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/contracts/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The ability to edit a contract helps you react quickly to the needs of your
// customers and your business.
//
// ### Use this endpoint to:
//
// - Encode mid-term commitment and discount changes
// - Fix configuration mistakes and easily roll back packaging changes
//
// ### Key response fields:
//
//   - The `id` of the edit
//   - Complete edit details. For example, if you edited the contract to add new
//     overrides and credits, you will receive the IDs of those overrides and credits
//     in the response.
//
// ### Usage guidelines:
//
//   - When you edit a contract, any draft invoices update immediately to reflect
//     that edit. Finalized invoices remain unchanged - you must void and regenerate
//     them in the UI or API to reflect the edit.
//   - Contract editing must be enabled to use this endpoint. Contact us via the
//     [Metronome support portal](https://support.metronome.com/) to learn more.
func (r *V2ContractService) Edit(ctx context.Context, body V2ContractEditParams, opts ...option.RequestOption) (res *V2ContractEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/contracts/edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Edit specific details for a contract-level or customer-level commit. Use this
// endpoint to modify individual commit access schedules, invoice schedules,
// applicable products, invoicing contracts, or other fields.
//
// ### Usage guidelines:
//
//   - As with all edits in Metronome, draft invoices will reflect the edit
//     immediately, while finalized invoices are untouched unless voided and
//     regenerated.
//   - If a commit's invoice schedule item is associated with a finalized invoice,
//     you cannot remove or update the invoice schedule item.
//   - If a commit's invoice schedule item is associated with a voided invoice, you
//     cannot remove the invoice schedule item.
//   - You cannot remove an commit access schedule segment that was applied to a
//     finalized invoice. You can void the invoice beforehand and then remove the
//     access schedule segment.
func (r *V2ContractService) EditCommit(ctx context.Context, body V2ContractEditCommitParams, opts ...option.RequestOption) (res *V2ContractEditCommitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/contracts/commits/edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Edit details for a contract-level or customer-level credit.
//
// ### Use this endpoint to:
//
//   - Extend the duration or the amount of an existing free credit like a trial
//   - Modify individual credit access schedules, applicable products, priority, or
//     other fields.
//
// ### Usage guidelines:
//
//   - As with all edits in Metronome, draft invoices will reflect the edit
//     immediately, while finalized invoices are untouched unless voided and
//     regenerated.
//   - You cannot remove an access schedule segment that was applied to a finalized
//     invoice. You can void the invoice beforehand and then remove the access
//     schedule segment.
func (r *V2ContractService) EditCredit(ctx context.Context, body V2ContractEditCreditParams, opts ...option.RequestOption) (res *V2ContractEditCreditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/contracts/credits/edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List all the edits made to a contract over time. In Metronome, you can edit a
// contract at any point after it's created to fix mistakes or reflect changes in
// terms. Metronome stores a full history of all edits that were ever made to a
// contract, whether through the UI, `editContract` endpoint, or other endpoints
// like `updateContractEndDate`.
//
// ### Use this endpoint to:
//
// - Understand what changes were made to a contract, when, and by who
//
// ### Key response fields:
//
//   - An array of every edit ever made to the contract
//   - Details on each individual edit - for example showing that in one edit, a user
//     added two discounts and incremented a subscription quantity.
func (r *V2ContractService) GetEditHistory(ctx context.Context, body V2ContractGetEditHistoryParams, opts ...option.RequestOption) (res *V2ContractGetEditHistoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/contracts/getEditHistory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type V2ContractGetResponse struct {
	Data shared.ContractV2 `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractListResponse struct {
	Data []shared.ContractV2 `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractListResponse) RawJSON() string { return r.JSON.raw }
func (r *V2ContractListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponse struct {
	Data V2ContractEditResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponse) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseData struct {
	ID   string                         `json:"id" api:"required" format:"uuid"`
	Edit V2ContractEditResponseDataEdit `json:"edit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Edit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseData) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEdit struct {
	ID                                      string                                             `json:"id" api:"required" format:"uuid"`
	AddCommits                              []V2ContractEditResponseDataEditAddCommit          `json:"add_commits"`
	AddCredits                              []V2ContractEditResponseDataEditAddCredit          `json:"add_credits"`
	AddDiscounts                            []shared.Discount                                  `json:"add_discounts"`
	AddOverrides                            []V2ContractEditResponseDataEditAddOverride        `json:"add_overrides"`
	AddPrepaidBalanceThresholdConfiguration shared.PrepaidBalanceThresholdConfigurationV2      `json:"add_prepaid_balance_threshold_configuration"`
	AddProServices                          []shared.ProService                                `json:"add_pro_services"`
	AddRecurringCommits                     []V2ContractEditResponseDataEditAddRecurringCommit `json:"add_recurring_commits"`
	AddRecurringCredits                     []V2ContractEditResponseDataEditAddRecurringCredit `json:"add_recurring_credits"`
	AddResellerRoyalties                    []V2ContractEditResponseDataEditAddResellerRoyalty `json:"add_reseller_royalties"`
	AddScheduledCharges                     []V2ContractEditResponseDataEditAddScheduledCharge `json:"add_scheduled_charges"`
	AddSpendThresholdConfiguration          shared.SpendThresholdConfigurationV2               `json:"add_spend_threshold_configuration"`
	// List of subscriptions on the contract.
	AddSubscriptions        []V2ContractEditResponseDataEditAddSubscription        `json:"add_subscriptions"`
	AddUsageFilters         []V2ContractEditResponseDataEditAddUsageFilter         `json:"add_usage_filters"`
	ArchiveCommits          []V2ContractEditResponseDataEditArchiveCommit          `json:"archive_commits"`
	ArchiveCredits          []V2ContractEditResponseDataEditArchiveCredit          `json:"archive_credits"`
	ArchiveScheduledCharges []V2ContractEditResponseDataEditArchiveScheduledCharge `json:"archive_scheduled_charges"`
	RemoveOverrides         []V2ContractEditResponseDataEditRemoveOverride         `json:"remove_overrides"`
	Timestamp               time.Time                                              `json:"timestamp" format:"date-time"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey         string                                       `json:"uniqueness_key"`
	UpdateCommits         []V2ContractEditResponseDataEditUpdateCommit `json:"update_commits"`
	UpdateContractEndDate time.Time                                    `json:"update_contract_end_date" format:"date-time"`
	// Value to update the contract name to. If not provided, the contract name will
	// remain unchanged.
	UpdateContractName                         string                                                                   `json:"update_contract_name" api:"nullable"`
	UpdateCredits                              []V2ContractEditResponseDataEditUpdateCredit                             `json:"update_credits"`
	UpdateDiscounts                            []V2ContractEditResponseDataEditUpdateDiscount                           `json:"update_discounts"`
	UpdatePrepaidBalanceThresholdConfiguration V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfiguration `json:"update_prepaid_balance_threshold_configuration"`
	UpdateRecurringCommits                     []V2ContractEditResponseDataEditUpdateRecurringCommit                    `json:"update_recurring_commits"`
	UpdateRecurringCredits                     []V2ContractEditResponseDataEditUpdateRecurringCredit                    `json:"update_recurring_credits"`
	UpdateRefundInvoices                       []V2ContractEditResponseDataEditUpdateRefundInvoice                      `json:"update_refund_invoices"`
	UpdateScheduledCharges                     []V2ContractEditResponseDataEditUpdateScheduledCharge                    `json:"update_scheduled_charges"`
	UpdateSpendThresholdConfiguration          V2ContractEditResponseDataEditUpdateSpendThresholdConfiguration          `json:"update_spend_threshold_configuration"`
	// Optional list of subscriptions to update.
	UpdateSubscriptions []V2ContractEditResponseDataEditUpdateSubscription `json:"update_subscriptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                         respjson.Field
		AddCommits                                 respjson.Field
		AddCredits                                 respjson.Field
		AddDiscounts                               respjson.Field
		AddOverrides                               respjson.Field
		AddPrepaidBalanceThresholdConfiguration    respjson.Field
		AddProServices                             respjson.Field
		AddRecurringCommits                        respjson.Field
		AddRecurringCredits                        respjson.Field
		AddResellerRoyalties                       respjson.Field
		AddScheduledCharges                        respjson.Field
		AddSpendThresholdConfiguration             respjson.Field
		AddSubscriptions                           respjson.Field
		AddUsageFilters                            respjson.Field
		ArchiveCommits                             respjson.Field
		ArchiveCredits                             respjson.Field
		ArchiveScheduledCharges                    respjson.Field
		RemoveOverrides                            respjson.Field
		Timestamp                                  respjson.Field
		UniquenessKey                              respjson.Field
		UpdateCommits                              respjson.Field
		UpdateContractEndDate                      respjson.Field
		UpdateContractName                         respjson.Field
		UpdateCredits                              respjson.Field
		UpdateDiscounts                            respjson.Field
		UpdatePrepaidBalanceThresholdConfiguration respjson.Field
		UpdateRecurringCommits                     respjson.Field
		UpdateRecurringCredits                     respjson.Field
		UpdateRefundInvoices                       respjson.Field
		UpdateScheduledCharges                     respjson.Field
		UpdateSpendThresholdConfiguration          respjson.Field
		UpdateSubscriptions                        respjson.Field
		ExtraFields                                map[string]respjson.Field
		raw                                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEdit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEdit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddCommit struct {
	ID      string                                         `json:"id" api:"required" format:"uuid"`
	Product V2ContractEditResponseDataEditAddCommitProduct `json:"product" api:"required"`
	// Any of "PREPAID", "POSTPAID".
	Type string `json:"type" api:"required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule        shared.ScheduleDuration `json:"access_schedule"`
	ApplicableProductIDs  []string                `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                `json:"applicable_product_tags"`
	Description           string                  `json:"description"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule V2ContractEditResponseDataEditAddCommitInvoiceSchedule `json:"invoice_schedule"`
	Name            string                                                 `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Product                 respjson.Field
		Type                    respjson.Field
		AccessSchedule          respjson.Field
		ApplicableProductIDs    respjson.Field
		ApplicableProductTags   respjson.Field
		Description             respjson.Field
		HierarchyConfiguration  respjson.Field
		InvoiceSchedule         respjson.Field
		Name                    respjson.Field
		NetsuiteSalesOrderID    respjson.Field
		Priority                respjson.Field
		RateType                respjson.Field
		RolloverFraction        respjson.Field
		SalesforceOpportunityID respjson.Field
		Specifiers              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddCommit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddCommitProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddCommitProduct) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddCommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The schedule that the customer will be invoiced for this commit.
type V2ContractEditResponseDataEditAddCommitInvoiceSchedule struct {
	CreditType shared.CreditTypeData `json:"credit_type"`
	// If true, this schedule will not generate an invoice.
	DoNotInvoice  bool                                                                 `json:"do_not_invoice"`
	ScheduleItems []V2ContractEditResponseDataEditAddCommitInvoiceScheduleScheduleItem `json:"schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType    respjson.Field
		DoNotInvoice  respjson.Field
		ScheduleItems respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddCommitInvoiceSchedule) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddCommitInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	Amount    float64   `json:"amount"`
	InvoiceID string    `json:"invoice_id" api:"nullable" format:"uuid"`
	Quantity  float64   `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Timestamp   respjson.Field
		Amount      respjson.Field
		InvoiceID   respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddCommitInvoiceScheduleScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddCommitInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddCredit struct {
	ID      string                                         `json:"id" api:"required" format:"uuid"`
	Product V2ContractEditResponseDataEditAddCreditProduct `json:"product" api:"required"`
	// Any of "CREDIT".
	Type string `json:"type" api:"required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        shared.ScheduleDuration `json:"access_schedule"`
	ApplicableProductIDs  []string                `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                `json:"applicable_product_tags"`
	Description           string                  `json:"description"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	Name                   string                              `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Product                 respjson.Field
		Type                    respjson.Field
		AccessSchedule          respjson.Field
		ApplicableProductIDs    respjson.Field
		ApplicableProductTags   respjson.Field
		Description             respjson.Field
		HierarchyConfiguration  respjson.Field
		Name                    respjson.Field
		NetsuiteSalesOrderID    respjson.Field
		Priority                respjson.Field
		RateType                respjson.Field
		RolloverFraction        respjson.Field
		SalesforceOpportunityID respjson.Field
		Specifiers              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddCredit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddCreditProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddCreditProduct) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddCreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddOverride struct {
	ID                    string                                                       `json:"id" api:"required" format:"uuid"`
	CreatedAt             time.Time                                                    `json:"created_at" api:"required" format:"date-time"`
	StartingAt            time.Time                                                    `json:"starting_at" api:"required" format:"date-time"`
	ApplicableProductTags []string                                                     `json:"applicable_product_tags"`
	EndingBefore          time.Time                                                    `json:"ending_before" format:"date-time"`
	Entitled              bool                                                         `json:"entitled"`
	IsCommitSpecific      bool                                                         `json:"is_commit_specific"`
	Multiplier            float64                                                      `json:"multiplier"`
	OverrideSpecifiers    []V2ContractEditResponseDataEditAddOverrideOverrideSpecifier `json:"override_specifiers"`
	OverrideTiers         []shared.OverrideTier                                        `json:"override_tiers"`
	OverwriteRate         V2ContractEditResponseDataEditAddOverrideOverwriteRate       `json:"overwrite_rate"`
	Priority              float64                                                      `json:"priority"`
	Product               V2ContractEditResponseDataEditAddOverrideProduct             `json:"product"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	Target string `json:"target"`
	// Any of "OVERWRITE", "MULTIPLIER", "TIERED".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		StartingAt            respjson.Field
		ApplicableProductTags respjson.Field
		EndingBefore          respjson.Field
		Entitled              respjson.Field
		IsCommitSpecific      respjson.Field
		Multiplier            respjson.Field
		OverrideSpecifiers    respjson.Field
		OverrideTiers         respjson.Field
		OverwriteRate         respjson.Field
		Priority              respjson.Field
		Product               respjson.Field
		Target                respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddOverride) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddOverrideOverrideSpecifier struct {
	AnyCommitOrCreditIDs []string `json:"any_commit_or_credit_ids"`
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency        string            `json:"billing_frequency"`
	CommitIDs               []string          `json:"commit_ids"`
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	ProductID               string            `json:"product_id" format:"uuid"`
	ProductTags             []string          `json:"product_tags"`
	RecurringCommitIDs      []string          `json:"recurring_commit_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnyCommitOrCreditIDs    respjson.Field
		BillingFrequency        respjson.Field
		CommitIDs               respjson.Field
		PresentationGroupValues respjson.Field
		PricingGroupValues      respjson.Field
		ProductID               respjson.Field
		ProductTags             respjson.Field
		RecurringCommitIDs      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddOverrideOverrideSpecifier) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddOverrideOverrideSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddOverrideOverwriteRate struct {
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "TIERED_PERCENTAGE",
	// "CUSTOM".
	RateType   string                `json:"rate_type" api:"required"`
	CreditType shared.CreditTypeData `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]any `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier `json:"tiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RateType    respjson.Field
		CreditType  respjson.Field
		CustomRate  respjson.Field
		IsProrated  respjson.Field
		Price       respjson.Field
		Quantity    respjson.Field
		Tiers       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddOverrideOverwriteRate) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddOverrideOverwriteRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddOverrideProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddOverrideProduct) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddOverrideProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddRecurringCommit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractEditResponseDataEditAddRecurringCommitAccessAmount `json:"access_amount" api:"required"`
	// The date this recurring commit's billing periods are anchored to.
	AnchorDate time.Time `json:"anchor_date" api:"required" format:"date-time"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractEditResponseDataEditAddRecurringCommitCommitDuration `json:"commit_duration" api:"required"`
	// Will be passed down to the individual commits
	Priority float64                                                 `json:"priority" api:"required"`
	Product  V2ContractEditResponseDataEditAddRecurringCommitProduct `json:"product" api:"required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type" api:"required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                                 `json:"applicable_product_tags"`
	Contract              V2ContractEditResponseDataEditAddRecurringCommitContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount V2ContractEditResponseDataEditAddRecurringCommitInvoiceAmount `json:"invoice_amount"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration"`
	// Rounding configuration for prorated recurring commit amounts.
	ProrationRounding V2ContractEditResponseDataEditAddRecurringCommitProrationRounding `json:"proration_rounding" api:"nullable"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates. - Daily recurring commits have a limit of one per contract, and
	// are unable to be created with seat-based subscriptions
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY", "DAILY".
	RecurrenceFrequency string `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []shared.CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig shared.RecurringCommitSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessAmount           respjson.Field
		AnchorDate             respjson.Field
		CommitDuration         respjson.Field
		Priority               respjson.Field
		Product                respjson.Field
		RateType               respjson.Field
		StartingAt             respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Contract               respjson.Field
		Description            respjson.Field
		EndingBefore           respjson.Field
		HierarchyConfiguration respjson.Field
		InvoiceAmount          respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Proration              respjson.Field
		ProrationRounding      respjson.Field
		RecurrenceFrequency    respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		SubscriptionConfig     respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCommit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type V2ContractEditResponseDataEditAddRecurringCommitAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	Quantity     float64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		UnitPrice    respjson.Field
		Quantity     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCommitAccessAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time the created commits will be valid for
type V2ContractEditResponseDataEditAddRecurringCommitCommitDuration struct {
	Value float64 `json:"value" api:"required"`
	// Any of "PERIODS".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCommitCommitDuration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddRecurringCommitCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddRecurringCommitProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCommitProduct) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddRecurringCommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddRecurringCommitContract struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCommitContract) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddRecurringCommitContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount the customer should be billed for the commit. Not required.
type V2ContractEditResponseDataEditAddRecurringCommitInvoiceAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	Quantity     float64 `json:"quantity" api:"required"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		Quantity     respjson.Field
		UnitPrice    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCommitInvoiceAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rounding configuration for prorated recurring commit amounts.
type V2ContractEditResponseDataEditAddRecurringCommitProrationRounding struct {
	Access  V2ContractEditResponseDataEditAddRecurringCommitProrationRoundingAccess  `json:"access"`
	Invoice V2ContractEditResponseDataEditAddRecurringCommitProrationRoundingInvoice `json:"invoice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Access      respjson.Field
		Invoice     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCommitProrationRounding) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddRecurringCommitProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddRecurringCommitProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCommitProrationRoundingAccess) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddRecurringCommitProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddRecurringCommitProrationRoundingInvoice struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCommitProrationRoundingInvoice) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddRecurringCommitProrationRoundingInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddRecurringCredit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractEditResponseDataEditAddRecurringCreditAccessAmount `json:"access_amount" api:"required"`
	// The date this recurring commit's billing periods are anchored to.
	AnchorDate time.Time `json:"anchor_date" api:"required" format:"date-time"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractEditResponseDataEditAddRecurringCreditCommitDuration `json:"commit_duration" api:"required"`
	// Will be passed down to the individual commits
	Priority float64                                                 `json:"priority" api:"required"`
	Product  V2ContractEditResponseDataEditAddRecurringCreditProduct `json:"product" api:"required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type" api:"required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                                 `json:"applicable_product_tags"`
	Contract              V2ContractEditResponseDataEditAddRecurringCreditContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration"`
	// Rounding configuration for prorated recurring credit amounts.
	ProrationRounding V2ContractEditResponseDataEditAddRecurringCreditProrationRounding `json:"proration_rounding" api:"nullable"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates. - Daily recurring commits have a limit of one per contract, and
	// are unable to be created with seat-based subscriptions
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY", "DAILY".
	RecurrenceFrequency string `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []shared.CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig shared.RecurringCommitSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessAmount           respjson.Field
		AnchorDate             respjson.Field
		CommitDuration         respjson.Field
		Priority               respjson.Field
		Product                respjson.Field
		RateType               respjson.Field
		StartingAt             respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Contract               respjson.Field
		Description            respjson.Field
		EndingBefore           respjson.Field
		HierarchyConfiguration respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Proration              respjson.Field
		ProrationRounding      respjson.Field
		RecurrenceFrequency    respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		SubscriptionConfig     respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCredit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type V2ContractEditResponseDataEditAddRecurringCreditAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	Quantity     float64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		UnitPrice    respjson.Field
		Quantity     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCreditAccessAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time the created commits will be valid for
type V2ContractEditResponseDataEditAddRecurringCreditCommitDuration struct {
	Value float64 `json:"value" api:"required"`
	// Any of "PERIODS".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCreditCommitDuration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddRecurringCreditCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddRecurringCreditProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCreditProduct) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddRecurringCreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddRecurringCreditContract struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCreditContract) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddRecurringCreditContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rounding configuration for prorated recurring credit amounts.
type V2ContractEditResponseDataEditAddRecurringCreditProrationRounding struct {
	Access V2ContractEditResponseDataEditAddRecurringCreditProrationRoundingAccess `json:"access"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Access      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCreditProrationRounding) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddRecurringCreditProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddRecurringCreditProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddRecurringCreditProrationRoundingAccess) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddRecurringCreditProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddResellerRoyalty struct {
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType          string    `json:"reseller_type" api:"required"`
	ApplicableProductIDs  []string  `json:"applicable_product_ids"`
	ApplicableProductTags []string  `json:"applicable_product_tags"`
	AwsAccountNumber      string    `json:"aws_account_number"`
	AwsOfferID            string    `json:"aws_offer_id"`
	AwsPayerReferenceID   string    `json:"aws_payer_reference_id"`
	EndingBefore          time.Time `json:"ending_before" api:"nullable" format:"date-time"`
	Fraction              float64   `json:"fraction"`
	GcpAccountID          string    `json:"gcp_account_id"`
	GcpOfferID            string    `json:"gcp_offer_id"`
	NetsuiteResellerID    string    `json:"netsuite_reseller_id"`
	ResellerContractValue float64   `json:"reseller_contract_value"`
	StartingAt            time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResellerType          respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		AwsAccountNumber      respjson.Field
		AwsOfferID            respjson.Field
		AwsPayerReferenceID   respjson.Field
		EndingBefore          respjson.Field
		Fraction              respjson.Field
		GcpAccountID          respjson.Field
		GcpOfferID            respjson.Field
		NetsuiteResellerID    respjson.Field
		ResellerContractValue respjson.Field
		StartingAt            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddResellerRoyalty) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddResellerRoyalty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddScheduledCharge struct {
	ID       string                                                  `json:"id" api:"required" format:"uuid"`
	Product  V2ContractEditResponseDataEditAddScheduledChargeProduct `json:"product" api:"required"`
	Schedule shared.SchedulePointInTime                              `json:"schedule" api:"required"`
	// displayed on invoices
	Name string `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Product              respjson.Field
		Schedule             respjson.Field
		Name                 respjson.Field
		NetsuiteSalesOrderID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddScheduledCharge) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddScheduledChargeProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddScheduledChargeProduct) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddScheduledChargeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddSubscription struct {
	// Previous, current, and next billing periods for the subscription.
	BillingPeriods V2ContractEditResponseDataEditAddSubscriptionBillingPeriods `json:"billing_periods" api:"required"`
	// Any of "ADVANCE", "ARREARS".
	CollectionSchedule string                                                 `json:"collection_schedule" api:"required"`
	Proration          V2ContractEditResponseDataEditAddSubscriptionProration `json:"proration" api:"required"`
	// Determines how the subscription's quantity is controlled. Defaults to
	// QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified
	// directly on the subscription. `initial_quantity` must be provided with this
	// option. Compatible with recurring commits/credits that use POOLED allocation.
	// **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add
	// user_123) to increment and decrement a subscription quantity, rather than
	// directly providing the quantity. You must use a **SEAT_BASED** subscription to
	// use a linked recurring credit with an allocation per seat. `seat_config` must be
	// provided with this option.
	//
	// Any of "SEAT_BASED", "QUANTITY_ONLY".
	QuantityManagementMode string `json:"quantity_management_mode" api:"required"`
	// List of quantity schedule items for the subscription. Only includes the current
	// quantity and future quantity changes.
	QuantitySchedule   []V2ContractEditResponseDataEditAddSubscriptionQuantitySchedule `json:"quantity_schedule" api:"required"`
	StartingAt         time.Time                                                       `json:"starting_at" api:"required" format:"date-time"`
	SubscriptionRate   V2ContractEditResponseDataEditAddSubscriptionSubscriptionRate   `json:"subscription_rate" api:"required"`
	ID                 string                                                          `json:"id" format:"uuid"`
	BillingCycleConfig V2ContractEditResponseDataEditAddSubscriptionBillingCycleConfig `json:"billing_cycle_config"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields     map[string]string `json:"custom_fields"`
	Description      string            `json:"description"`
	EndingBefore     time.Time         `json:"ending_before" format:"date-time"`
	FiatCreditTypeID string            `json:"fiat_credit_type_id" format:"uuid"`
	Name             string            `json:"name"`
	// Custom fields from the subscription product referenced by
	// `subscription_rate.product`. These are distinct from the subscription instance's
	// `custom_fields`.
	ProductCustomFields map[string]string                                       `json:"product_custom_fields"`
	SeatConfig          V2ContractEditResponseDataEditAddSubscriptionSeatConfig `json:"seat_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingPeriods         respjson.Field
		CollectionSchedule     respjson.Field
		Proration              respjson.Field
		QuantityManagementMode respjson.Field
		QuantitySchedule       respjson.Field
		StartingAt             respjson.Field
		SubscriptionRate       respjson.Field
		ID                     respjson.Field
		BillingCycleConfig     respjson.Field
		CustomFields           respjson.Field
		Description            respjson.Field
		EndingBefore           respjson.Field
		FiatCreditTypeID       respjson.Field
		Name                   respjson.Field
		ProductCustomFields    respjson.Field
		SeatConfig             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscription) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Previous, current, and next billing periods for the subscription.
type V2ContractEditResponseDataEditAddSubscriptionBillingPeriods struct {
	Current  V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsCurrent  `json:"current"`
	Next     V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsNext     `json:"next"`
	Previous V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsPrevious `json:"previous"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscriptionBillingPeriods) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddSubscriptionBillingPeriods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsCurrent struct {
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsCurrent) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsCurrent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsNext struct {
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsNext) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsNext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsPrevious struct {
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsPrevious) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddSubscriptionBillingPeriodsPrevious) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddSubscriptionProration struct {
	// Any of "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE".
	InvoiceBehavior string                                                         `json:"invoice_behavior" api:"required"`
	IsProrated      bool                                                           `json:"is_prorated" api:"required"`
	Rounding        V2ContractEditResponseDataEditAddSubscriptionProrationRounding `json:"rounding"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InvoiceBehavior respjson.Field
		IsProrated      respjson.Field
		Rounding        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscriptionProration) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddSubscriptionProration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddSubscriptionProrationRounding struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscriptionProrationRounding) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddSubscriptionProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddSubscriptionQuantitySchedule struct {
	Quantity     float64   `json:"quantity" api:"required"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity     respjson.Field
		StartingAt   respjson.Field
		EndingBefore respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscriptionQuantitySchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddSubscriptionQuantitySchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddSubscriptionSubscriptionRate struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string                                                               `json:"billing_frequency" api:"required"`
	Product          V2ContractEditResponseDataEditAddSubscriptionSubscriptionRateProduct `json:"product" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFrequency respjson.Field
		Product          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscriptionSubscriptionRate) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddSubscriptionSubscriptionRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddSubscriptionSubscriptionRateProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscriptionSubscriptionRateProduct) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddSubscriptionSubscriptionRateProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddSubscriptionBillingCycleConfig struct {
	// The date this subscription's billing cycle is anchored to.
	AnchorDate time.Time `json:"anchor_date" api:"required" format:"date-time"`
	// Controls whether this subscription consolidates onto usage invoices or gets its
	// own scheduled invoice.
	//
	// Any of "ON_SCHEDULED_INVOICE", "ON_USAGE_INVOICE".
	InvoicePlacement string `json:"invoice_placement" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnchorDate       respjson.Field
		InvoicePlacement respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscriptionBillingCycleConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditAddSubscriptionBillingCycleConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddSubscriptionSeatConfig struct {
	// The property name, sent on usage events, that identifies the seat ID associated
	// with the usage event. For example, the property name might be seat_id or
	// user_id. The property must be set as a group key on billable metrics and a
	// presentation/pricing group key on contract products. This allows linked
	// recurring credits with an allocation per seat to be consumed by only one seat's
	// usage.
	SeatGroupKey string `json:"seat_group_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SeatGroupKey respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddSubscriptionSeatConfig) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddSubscriptionSeatConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditAddUsageFilter struct {
	GroupKey    string   `json:"group_key" api:"required"`
	GroupValues []string `json:"group_values" api:"required"`
	// This will match contract starting_at value if usage filter is active from the
	// beginning of the contract.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// This will match contract ending_before value if usage filter is active until the
	// end of the contract. It will be undefined if the contract is open-ended.
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupKey     respjson.Field
		GroupValues  respjson.Field
		StartingAt   respjson.Field
		EndingBefore respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditAddUsageFilter) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditAddUsageFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditArchiveCommit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditArchiveCommit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditArchiveCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditArchiveCredit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditArchiveCredit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditArchiveCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditArchiveScheduledCharge struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditArchiveScheduledCharge) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditArchiveScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditRemoveOverride struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditRemoveOverride) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditRemoveOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCommit struct {
	ID             string                                                   `json:"id" api:"required" format:"uuid"`
	AccessSchedule V2ContractEditResponseDataEditUpdateCommitAccessSchedule `json:"access_schedule"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids" api:"nullable" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags" api:"nullable"`
	Description           string   `json:"description"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration                       `json:"hierarchy_configuration"`
	InvoiceSchedule        V2ContractEditResponseDataEditUpdateCommitInvoiceSchedule `json:"invoice_schedule"`
	Name                   string                                                    `json:"name"`
	NetsuiteSalesOrderID   string                                                    `json:"netsuite_sales_order_id" api:"nullable"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority  float64 `json:"priority" api:"nullable"`
	ProductID string  `json:"product_id" format:"uuid"`
	// If set, the commit's rate type was updated to the specified value.
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction" api:"nullable"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessSchedule         respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Description            respjson.Field
		HierarchyConfiguration respjson.Field
		InvoiceSchedule        respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Priority               respjson.Field
		ProductID              respjson.Field
		RateType               respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCommit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditUpdateCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCommitAccessSchedule struct {
	AddScheduleItems    []V2ContractEditResponseDataEditUpdateCommitAccessScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractEditResponseDataEditUpdateCommitAccessScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractEditResponseDataEditUpdateCommitAccessScheduleUpdateScheduleItem `json:"update_schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddScheduleItems    respjson.Field
		RemoveScheduleItems respjson.Field
		UpdateScheduleItems respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCommitAccessSchedule) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditUpdateCommitAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCommitAccessScheduleAddScheduleItem struct {
	Amount float64 `json:"amount" api:"required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount       respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCommitAccessScheduleAddScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateCommitAccessScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCommitAccessScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCommitAccessScheduleRemoveScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateCommitAccessScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCommitAccessScheduleUpdateScheduleItem struct {
	ID     string  `json:"id" api:"required" format:"uuid"`
	Amount float64 `json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Amount       respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCommitAccessScheduleUpdateScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateCommitAccessScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCommitInvoiceSchedule struct {
	AddScheduleItems    []V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleUpdateScheduleItem `json:"update_schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddScheduleItems    respjson.Field
		RemoveScheduleItems respjson.Field
		UpdateScheduleItems respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCommitInvoiceSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleAddScheduleItem struct {
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	Amount    float64   `json:"amount"`
	Quantity  float64   `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Amount      respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleAddScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleRemoveScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleUpdateScheduleItem struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	Amount    float64   `json:"amount"`
	Quantity  float64   `json:"quantity"`
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	UnitPrice float64   `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Quantity    respjson.Field
		Timestamp   respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleUpdateScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateCommitInvoiceScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCredit struct {
	ID             string                                                   `json:"id" api:"required" format:"uuid"`
	AccessSchedule V2ContractEditResponseDataEditUpdateCreditAccessSchedule `json:"access_schedule"`
	// Which products the credit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the credit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids" api:"nullable" format:"uuid"`
	// Which tags the credit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the credit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags" api:"nullable"`
	Description           string   `json:"description"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	Name                   string                              `json:"name"`
	NetsuiteSalesOrderID   string                              `json:"netsuite_sales_order_id" api:"nullable"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority  float64 `json:"priority" api:"nullable"`
	ProductID string  `json:"product_id" format:"uuid"`
	// If set, the credit's rate type was updated to the specified value.
	//
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction" api:"nullable"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessSchedule         respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Description            respjson.Field
		HierarchyConfiguration respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Priority               respjson.Field
		ProductID              respjson.Field
		RateType               respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCredit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditUpdateCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCreditAccessSchedule struct {
	AddScheduleItems    []V2ContractEditResponseDataEditUpdateCreditAccessScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractEditResponseDataEditUpdateCreditAccessScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractEditResponseDataEditUpdateCreditAccessScheduleUpdateScheduleItem `json:"update_schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddScheduleItems    respjson.Field
		RemoveScheduleItems respjson.Field
		UpdateScheduleItems respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCreditAccessSchedule) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditUpdateCreditAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCreditAccessScheduleAddScheduleItem struct {
	Amount float64 `json:"amount" api:"required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount       respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCreditAccessScheduleAddScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateCreditAccessScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCreditAccessScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCreditAccessScheduleRemoveScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateCreditAccessScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateCreditAccessScheduleUpdateScheduleItem struct {
	ID     string  `json:"id" api:"required" format:"uuid"`
	Amount float64 `json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Amount       respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateCreditAccessScheduleUpdateScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateCreditAccessScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateDiscount struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields         map[string]string `json:"custom_fields"`
	Name                 string            `json:"name"`
	NetsuiteSalesOrderID string            `json:"netsuite_sales_order_id"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule V2ContractEditResponseDataEditUpdateDiscountSchedule `json:"schedule"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CustomFields         respjson.Field
		Name                 respjson.Field
		NetsuiteSalesOrderID respjson.Field
		Schedule             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateDiscount) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditUpdateDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must provide either schedule_items or recurring_schedule.
type V2ContractEditResponseDataEditUpdateDiscountSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID string `json:"credit_type_id" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice bool `json:"do_not_invoice"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V2ContractEditResponseDataEditUpdateDiscountScheduleRecurringSchedule `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V2ContractEditResponseDataEditUpdateDiscountScheduleScheduleItem `json:"schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID      respjson.Field
		DoNotInvoice      respjson.Field
		RecurringSchedule respjson.Field
		ScheduleItems     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateDiscountSchedule) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditUpdateDiscountSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V2ContractEditResponseDataEditUpdateDiscountScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution" api:"required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency" api:"required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount float64 `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity float64 `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice float64 `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountDistribution respjson.Field
		EndingBefore       respjson.Field
		Frequency          respjson.Field
		StartingAt         respjson.Field
		Amount             respjson.Field
		Quantity           respjson.Field
		UnitPrice          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateDiscountScheduleRecurringSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateDiscountScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateDiscountScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount float64 `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity float64 `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice float64 `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Amount      respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateDiscountScheduleScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateDiscountScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfiguration struct {
	Commit V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationCommit `json:"commit"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID    string                                                                                        `json:"custom_credit_type_id" api:"nullable" format:"uuid"`
	DiscountConfiguration V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration `json:"discount_configuration" api:"nullable"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                       `json:"is_enabled"`
	PaymentGateConfig shared.PaymentGateConfigV2 `json:"payment_gate_config"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount float64 `json:"recharge_to_amount"`
	// Specify the threshold amount for the contract. Each time the contract's balance
	// lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount"`
	// Determines which balances are excluded from remaining balance calculation for
	// threshold billing.
	ThresholdBalanceSpecifiers []V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier `json:"threshold_balance_specifiers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit                     respjson.Field
		CustomCreditTypeID         respjson.Field
		DiscountConfiguration      respjson.Field
		IsEnabled                  respjson.Field
		PaymentGateConfig          respjson.Field
		RechargeToAmount           respjson.Field
		ThresholdAmount            respjson.Field
		ThresholdBalanceSpecifiers respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationCommit struct {
	// Which products the threshold commit applies to. If both applicable_product_ids
	// and applicable_product_tags are not provided, the commit applies to all
	// products.
	ApplicableProductIDs []string `json:"applicable_product_ids" api:"nullable" format:"uuid"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags" api:"nullable"`
	// The length of time the created commit will be valid, starting from the end of
	// the invoice's service period. Set to null to clear a previously configured
	// duration.
	Duration V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationCommitDuration `json:"duration" api:"nullable"`
	// Whether the created commits will be charged at commit rate or list rate. Set to
	// null to clear a previously configured rate type.
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type" api:"nullable"`
	// Fraction of the created commit's unused balance that will roll over. Must be
	// between 0 and 1. Set to null to clear a previously configured rollover fraction.
	RolloverFraction float64 `json:"rollover_fraction" api:"nullable"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		Duration              respjson.Field
		RateType              respjson.Field
		RolloverFraction      respjson.Field
		Specifiers            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
	shared.UpdateBaseThresholdCommit
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationCommit) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The length of time the created commit will be valid, starting from the end of
// the invoice's service period. Set to null to clear a previously configured
// duration.
type V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationCommitDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit" api:"required"`
	Value int64  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationCommitDuration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration struct {
	// Update the discount cap. Set to null to remove an existing cap.
	Cap V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap `json:"cap" api:"nullable"`
	// The fraction of the original amount that the customer pays after applying the
	// discount. Set to null to remove the discount fraction. For example, 0.85 means
	// the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction float64 `json:"payment_fraction" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cap             respjson.Field
		PaymentFraction respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update the discount cap. Set to null to remove an existing cap.
type V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount float64 `json:"amount" api:"required"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias string `json:"spend_tracker_alias" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount            respjson.Field
		SpendTrackerAlias respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier struct {
	Exclude []V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude `json:"exclude" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exclude     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude struct {
	CustomFieldFilters []V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter `json:"custom_field_filters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomFieldFilters respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter struct {
	// Any of "Commit", "ContractCredit", "ContractCreditOrCommit".
	Entity string `json:"entity" api:"required"`
	Key    string `json:"key" api:"required"`
	Value  string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entity      respjson.Field
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateRecurringCommit struct {
	ID            string                                                           `json:"id" api:"required" format:"uuid"`
	AccessAmount  V2ContractEditResponseDataEditUpdateRecurringCommitAccessAmount  `json:"access_amount"`
	EndingBefore  time.Time                                                        `json:"ending_before" format:"date-time"`
	InvoiceAmount V2ContractEditResponseDataEditUpdateRecurringCommitInvoiceAmount `json:"invoice_amount"`
	// Rounding configuration for prorated recurring commit amounts.
	ProrationRounding V2ContractEditResponseDataEditUpdateRecurringCommitProrationRounding `json:"proration_rounding" api:"nullable"`
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType string `json:"rate_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AccessAmount      respjson.Field
		EndingBefore      respjson.Field
		InvoiceAmount     respjson.Field
		ProrationRounding respjson.Field
		RateType          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateRecurringCommit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditUpdateRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateRecurringCommitAccessAmount struct {
	Quantity  float64 `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateRecurringCommitAccessAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateRecurringCommitInvoiceAmount struct {
	Quantity  float64 `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateRecurringCommitInvoiceAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rounding configuration for prorated recurring commit amounts.
type V2ContractEditResponseDataEditUpdateRecurringCommitProrationRounding struct {
	Access  V2ContractEditResponseDataEditUpdateRecurringCommitProrationRoundingAccess  `json:"access"`
	Invoice V2ContractEditResponseDataEditUpdateRecurringCommitProrationRoundingInvoice `json:"invoice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Access      respjson.Field
		Invoice     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateRecurringCommitProrationRounding) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateRecurringCommitProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateRecurringCommitProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateRecurringCommitProrationRoundingAccess) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateRecurringCommitProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateRecurringCommitProrationRoundingInvoice struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateRecurringCommitProrationRoundingInvoice) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateRecurringCommitProrationRoundingInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateRecurringCredit struct {
	ID           string                                                          `json:"id" api:"required" format:"uuid"`
	AccessAmount V2ContractEditResponseDataEditUpdateRecurringCreditAccessAmount `json:"access_amount"`
	EndingBefore time.Time                                                       `json:"ending_before" format:"date-time"`
	// Rounding configuration for prorated recurring credit amounts.
	ProrationRounding V2ContractEditResponseDataEditUpdateRecurringCreditProrationRounding `json:"proration_rounding" api:"nullable"`
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType string `json:"rate_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AccessAmount      respjson.Field
		EndingBefore      respjson.Field
		ProrationRounding respjson.Field
		RateType          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateRecurringCredit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditUpdateRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateRecurringCreditAccessAmount struct {
	Quantity  float64 `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateRecurringCreditAccessAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rounding configuration for prorated recurring credit amounts.
type V2ContractEditResponseDataEditUpdateRecurringCreditProrationRounding struct {
	Access V2ContractEditResponseDataEditUpdateRecurringCreditProrationRoundingAccess `json:"access"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Access      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateRecurringCreditProrationRounding) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateRecurringCreditProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateRecurringCreditProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateRecurringCreditProrationRoundingAccess) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateRecurringCreditProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateRefundInvoice struct {
	Date      time.Time `json:"date" api:"required" format:"date-time"`
	InvoiceID string    `json:"invoice_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		InvoiceID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateRefundInvoice) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditUpdateRefundInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateScheduledCharge struct {
	ID                   string                                                             `json:"id" api:"required" format:"uuid"`
	InvoiceSchedule      V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceSchedule `json:"invoice_schedule"`
	Name                 string                                                             `json:"name"`
	NetsuiteSalesOrderID string                                                             `json:"netsuite_sales_order_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		InvoiceSchedule      respjson.Field
		Name                 respjson.Field
		NetsuiteSalesOrderID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateScheduledCharge) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditUpdateScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceSchedule struct {
	AddScheduleItems    []V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem `json:"update_schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddScheduleItems    respjson.Field
		RemoveScheduleItems respjson.Field
		UpdateScheduleItems respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleAddScheduleItem struct {
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	Amount    float64   `json:"amount"`
	Quantity  float64   `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Amount      respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleAddScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	Amount    float64   `json:"amount"`
	Quantity  float64   `json:"quantity"`
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	UnitPrice float64   `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Quantity    respjson.Field
		Timestamp   respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateSpendThresholdConfiguration struct {
	Commit                shared.UpdateBaseThresholdCommit                                                     `json:"commit"`
	DiscountConfiguration V2ContractEditResponseDataEditUpdateSpendThresholdConfigurationDiscountConfiguration `json:"discount_configuration" api:"nullable"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                       `json:"is_enabled"`
	PaymentGateConfig shared.PaymentGateConfigV2 `json:"payment_gate_config"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit                respjson.Field
		DiscountConfiguration respjson.Field
		IsEnabled             respjson.Field
		PaymentGateConfig     respjson.Field
		ThresholdAmount       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateSpendThresholdConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateSpendThresholdConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateSpendThresholdConfigurationDiscountConfiguration struct {
	// Update the discount cap. Set to null to remove an existing cap.
	Cap V2ContractEditResponseDataEditUpdateSpendThresholdConfigurationDiscountConfigurationCap `json:"cap" api:"nullable"`
	// The fraction of the original amount that the customer pays after applying the
	// discount. Set to null to remove the discount fraction. For example, 0.85 means
	// the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction float64 `json:"payment_fraction" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cap             respjson.Field
		PaymentFraction respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateSpendThresholdConfigurationDiscountConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateSpendThresholdConfigurationDiscountConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update the discount cap. Set to null to remove an existing cap.
type V2ContractEditResponseDataEditUpdateSpendThresholdConfigurationDiscountConfigurationCap struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount float64 `json:"amount" api:"required"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias string `json:"spend_tracker_alias" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount            respjson.Field
		SpendTrackerAlias respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateSpendThresholdConfigurationDiscountConfigurationCap) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateSpendThresholdConfigurationDiscountConfigurationCap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateSubscription struct {
	ID              string                                                           `json:"id" api:"required" format:"uuid"`
	EndingBefore    time.Time                                                        `json:"ending_before" format:"date-time"`
	Name            string                                                           `json:"name"`
	QuantityUpdates []V2ContractEditResponseDataEditUpdateSubscriptionQuantityUpdate `json:"quantity_updates"`
	// Manage subscription seats for subscriptions in SEAT_BASED mode.
	SeatUpdates V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdates `json:"seat_updates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		EndingBefore    respjson.Field
		Name            respjson.Field
		QuantityUpdates respjson.Field
		SeatUpdates     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateSubscription) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditResponseDataEditUpdateSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateSubscriptionQuantityUpdate struct {
	StartingAt    time.Time `json:"starting_at" api:"required" format:"date-time"`
	Quantity      float64   `json:"quantity"`
	QuantityDelta float64   `json:"quantity_delta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StartingAt    respjson.Field
		Quantity      respjson.Field
		QuantityDelta respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateSubscriptionQuantityUpdate) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateSubscriptionQuantityUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Manage subscription seats for subscriptions in SEAT_BASED mode.
type V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdates struct {
	// Adds seat IDs to the subscription. If there are unassigned seats, the new seat
	// IDs will fill these unassigned seats and not increase the total subscription
	// quantity. Otherwise, if there are more new seat IDs than unassigned seats, the
	// total subscription quantity will increase.
	AddSeatIDs []V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesAddSeatID `json:"add_seat_ids"`
	// Adds unassigned seats to the subscription. This will increase the total
	// subscription quantity.
	AddUnassignedSeats []V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesAddUnassignedSeat `json:"add_unassigned_seats"`
	// Removes seat IDs from the subscription, if possible. If a seat ID is removed,
	// the total subscription quantity will decrease. Otherwise, if the seat ID is not
	// found on the subscription, this is a no-op.
	RemoveSeatIDs []V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesRemoveSeatID `json:"remove_seat_ids"`
	// Removes unassigned seats from the subscription. This will decrease the total
	// subscription quantity if there are are unassigned seats.
	RemoveUnassignedSeats []V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat `json:"remove_unassigned_seats"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddSeatIDs            respjson.Field
		AddUnassignedSeats    respjson.Field
		RemoveSeatIDs         respjson.Field
		RemoveUnassignedSeats respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdates) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesAddSeatID struct {
	SeatIDs []string `json:"seat_ids" api:"required"`
	// Assigned seats will be added/removed starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SeatIDs     respjson.Field
		StartingAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesAddSeatID) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesAddSeatID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesAddUnassignedSeat struct {
	// The number of unassigned seats on the subscription will increase/decrease by
	// this delta. Must be greater than 0.
	Quantity float64 `json:"quantity" api:"required"`
	// Unassigned seats will be updated starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		StartingAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesAddUnassignedSeat) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesAddUnassignedSeat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesRemoveSeatID struct {
	SeatIDs []string `json:"seat_ids" api:"required"`
	// Assigned seats will be added/removed starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SeatIDs     respjson.Field
		StartingAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesRemoveSeatID) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesRemoveSeatID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat struct {
	// The number of unassigned seats on the subscription will increase/decrease by
	// this delta. Must be greater than 0.
	Quantity float64 `json:"quantity" api:"required"`
	// Unassigned seats will be updated starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		StartingAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractEditResponseDataEditUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditCommitResponse struct {
	Data shared.ID `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditCommitResponse) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditCommitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditCreditResponse struct {
	Data shared.ID `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractEditCreditResponse) RawJSON() string { return r.JSON.raw }
func (r *V2ContractEditCreditResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponse struct {
	Data []V2ContractGetEditHistoryResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseData struct {
	ID                                      string                                                   `json:"id" api:"required" format:"uuid"`
	AddCommits                              []V2ContractGetEditHistoryResponseDataAddCommit          `json:"add_commits"`
	AddCredits                              []V2ContractGetEditHistoryResponseDataAddCredit          `json:"add_credits"`
	AddDiscounts                            []shared.Discount                                        `json:"add_discounts"`
	AddOverrides                            []V2ContractGetEditHistoryResponseDataAddOverride        `json:"add_overrides"`
	AddPrepaidBalanceThresholdConfiguration shared.PrepaidBalanceThresholdConfigurationV2            `json:"add_prepaid_balance_threshold_configuration"`
	AddProServices                          []shared.ProService                                      `json:"add_pro_services"`
	AddRecurringCommits                     []V2ContractGetEditHistoryResponseDataAddRecurringCommit `json:"add_recurring_commits"`
	AddRecurringCredits                     []V2ContractGetEditHistoryResponseDataAddRecurringCredit `json:"add_recurring_credits"`
	AddResellerRoyalties                    []V2ContractGetEditHistoryResponseDataAddResellerRoyalty `json:"add_reseller_royalties"`
	AddScheduledCharges                     []V2ContractGetEditHistoryResponseDataAddScheduledCharge `json:"add_scheduled_charges"`
	AddSpendThresholdConfiguration          shared.SpendThresholdConfigurationV2                     `json:"add_spend_threshold_configuration"`
	// List of subscriptions on the contract.
	AddSubscriptions        []V2ContractGetEditHistoryResponseDataAddSubscription        `json:"add_subscriptions"`
	AddUsageFilters         []V2ContractGetEditHistoryResponseDataAddUsageFilter         `json:"add_usage_filters"`
	ArchiveCommits          []V2ContractGetEditHistoryResponseDataArchiveCommit          `json:"archive_commits"`
	ArchiveCredits          []V2ContractGetEditHistoryResponseDataArchiveCredit          `json:"archive_credits"`
	ArchiveScheduledCharges []V2ContractGetEditHistoryResponseDataArchiveScheduledCharge `json:"archive_scheduled_charges"`
	RemoveOverrides         []V2ContractGetEditHistoryResponseDataRemoveOverride         `json:"remove_overrides"`
	Timestamp               time.Time                                                    `json:"timestamp" format:"date-time"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey         string                                             `json:"uniqueness_key"`
	UpdateCommits         []V2ContractGetEditHistoryResponseDataUpdateCommit `json:"update_commits"`
	UpdateContractEndDate time.Time                                          `json:"update_contract_end_date" format:"date-time"`
	// Value to update the contract name to. If not provided, the contract name will
	// remain unchanged.
	UpdateContractName                         string                                                                         `json:"update_contract_name" api:"nullable"`
	UpdateCredits                              []V2ContractGetEditHistoryResponseDataUpdateCredit                             `json:"update_credits"`
	UpdateDiscounts                            []V2ContractGetEditHistoryResponseDataUpdateDiscount                           `json:"update_discounts"`
	UpdatePrepaidBalanceThresholdConfiguration V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfiguration `json:"update_prepaid_balance_threshold_configuration"`
	UpdateRecurringCommits                     []V2ContractGetEditHistoryResponseDataUpdateRecurringCommit                    `json:"update_recurring_commits"`
	UpdateRecurringCredits                     []V2ContractGetEditHistoryResponseDataUpdateRecurringCredit                    `json:"update_recurring_credits"`
	UpdateRefundInvoices                       []V2ContractGetEditHistoryResponseDataUpdateRefundInvoice                      `json:"update_refund_invoices"`
	UpdateScheduledCharges                     []V2ContractGetEditHistoryResponseDataUpdateScheduledCharge                    `json:"update_scheduled_charges"`
	UpdateSpendThresholdConfiguration          V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfiguration          `json:"update_spend_threshold_configuration"`
	// Optional list of subscriptions to update.
	UpdateSubscriptions []V2ContractGetEditHistoryResponseDataUpdateSubscription `json:"update_subscriptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                         respjson.Field
		AddCommits                                 respjson.Field
		AddCredits                                 respjson.Field
		AddDiscounts                               respjson.Field
		AddOverrides                               respjson.Field
		AddPrepaidBalanceThresholdConfiguration    respjson.Field
		AddProServices                             respjson.Field
		AddRecurringCommits                        respjson.Field
		AddRecurringCredits                        respjson.Field
		AddResellerRoyalties                       respjson.Field
		AddScheduledCharges                        respjson.Field
		AddSpendThresholdConfiguration             respjson.Field
		AddSubscriptions                           respjson.Field
		AddUsageFilters                            respjson.Field
		ArchiveCommits                             respjson.Field
		ArchiveCredits                             respjson.Field
		ArchiveScheduledCharges                    respjson.Field
		RemoveOverrides                            respjson.Field
		Timestamp                                  respjson.Field
		UniquenessKey                              respjson.Field
		UpdateCommits                              respjson.Field
		UpdateContractEndDate                      respjson.Field
		UpdateContractName                         respjson.Field
		UpdateCredits                              respjson.Field
		UpdateDiscounts                            respjson.Field
		UpdatePrepaidBalanceThresholdConfiguration respjson.Field
		UpdateRecurringCommits                     respjson.Field
		UpdateRecurringCredits                     respjson.Field
		UpdateRefundInvoices                       respjson.Field
		UpdateScheduledCharges                     respjson.Field
		UpdateSpendThresholdConfiguration          respjson.Field
		UpdateSubscriptions                        respjson.Field
		ExtraFields                                map[string]respjson.Field
		raw                                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseData) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddCommit struct {
	ID      string                                               `json:"id" api:"required" format:"uuid"`
	Product V2ContractGetEditHistoryResponseDataAddCommitProduct `json:"product" api:"required"`
	// Any of "PREPAID", "POSTPAID".
	Type string `json:"type" api:"required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule        shared.ScheduleDuration `json:"access_schedule"`
	ApplicableProductIDs  []string                `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                `json:"applicable_product_tags"`
	Description           string                  `json:"description"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule V2ContractGetEditHistoryResponseDataAddCommitInvoiceSchedule `json:"invoice_schedule"`
	Name            string                                                       `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Product                 respjson.Field
		Type                    respjson.Field
		AccessSchedule          respjson.Field
		ApplicableProductIDs    respjson.Field
		ApplicableProductTags   respjson.Field
		Description             respjson.Field
		HierarchyConfiguration  respjson.Field
		InvoiceSchedule         respjson.Field
		Name                    respjson.Field
		NetsuiteSalesOrderID    respjson.Field
		Priority                respjson.Field
		RateType                respjson.Field
		RolloverFraction        respjson.Field
		SalesforceOpportunityID respjson.Field
		Specifiers              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddCommit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddCommitProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddCommitProduct) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddCommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The schedule that the customer will be invoiced for this commit.
type V2ContractGetEditHistoryResponseDataAddCommitInvoiceSchedule struct {
	CreditType shared.CreditTypeData `json:"credit_type"`
	// If true, this schedule will not generate an invoice.
	DoNotInvoice  bool                                                                       `json:"do_not_invoice"`
	ScheduleItems []V2ContractGetEditHistoryResponseDataAddCommitInvoiceScheduleScheduleItem `json:"schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType    respjson.Field
		DoNotInvoice  respjson.Field
		ScheduleItems respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddCommitInvoiceSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddCommitInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	Amount    float64   `json:"amount"`
	InvoiceID string    `json:"invoice_id" api:"nullable" format:"uuid"`
	Quantity  float64   `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Timestamp   respjson.Field
		Amount      respjson.Field
		InvoiceID   respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddCommitInvoiceScheduleScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddCommitInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddCredit struct {
	ID      string                                               `json:"id" api:"required" format:"uuid"`
	Product V2ContractGetEditHistoryResponseDataAddCreditProduct `json:"product" api:"required"`
	// Any of "CREDIT".
	Type string `json:"type" api:"required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        shared.ScheduleDuration `json:"access_schedule"`
	ApplicableProductIDs  []string                `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                `json:"applicable_product_tags"`
	Description           string                  `json:"description"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	Name                   string                              `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Product                 respjson.Field
		Type                    respjson.Field
		AccessSchedule          respjson.Field
		ApplicableProductIDs    respjson.Field
		ApplicableProductTags   respjson.Field
		Description             respjson.Field
		HierarchyConfiguration  respjson.Field
		Name                    respjson.Field
		NetsuiteSalesOrderID    respjson.Field
		Priority                respjson.Field
		RateType                respjson.Field
		RolloverFraction        respjson.Field
		SalesforceOpportunityID respjson.Field
		Specifiers              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddCredit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddCreditProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddCreditProduct) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddCreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddOverride struct {
	ID                    string                                                             `json:"id" api:"required" format:"uuid"`
	CreatedAt             time.Time                                                          `json:"created_at" api:"required" format:"date-time"`
	StartingAt            time.Time                                                          `json:"starting_at" api:"required" format:"date-time"`
	ApplicableProductTags []string                                                           `json:"applicable_product_tags"`
	EndingBefore          time.Time                                                          `json:"ending_before" format:"date-time"`
	Entitled              bool                                                               `json:"entitled"`
	IsCommitSpecific      bool                                                               `json:"is_commit_specific"`
	Multiplier            float64                                                            `json:"multiplier"`
	OverrideSpecifiers    []V2ContractGetEditHistoryResponseDataAddOverrideOverrideSpecifier `json:"override_specifiers"`
	OverrideTiers         []shared.OverrideTier                                              `json:"override_tiers"`
	OverwriteRate         V2ContractGetEditHistoryResponseDataAddOverrideOverwriteRate       `json:"overwrite_rate"`
	Priority              float64                                                            `json:"priority"`
	Product               V2ContractGetEditHistoryResponseDataAddOverrideProduct             `json:"product"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	Target string `json:"target"`
	// Any of "OVERWRITE", "MULTIPLIER", "TIERED".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		StartingAt            respjson.Field
		ApplicableProductTags respjson.Field
		EndingBefore          respjson.Field
		Entitled              respjson.Field
		IsCommitSpecific      respjson.Field
		Multiplier            respjson.Field
		OverrideSpecifiers    respjson.Field
		OverrideTiers         respjson.Field
		OverwriteRate         respjson.Field
		Priority              respjson.Field
		Product               respjson.Field
		Target                respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddOverride) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddOverrideOverrideSpecifier struct {
	AnyCommitOrCreditIDs []string `json:"any_commit_or_credit_ids"`
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency        string            `json:"billing_frequency"`
	CommitIDs               []string          `json:"commit_ids"`
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	ProductID               string            `json:"product_id" format:"uuid"`
	ProductTags             []string          `json:"product_tags"`
	RecurringCommitIDs      []string          `json:"recurring_commit_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnyCommitOrCreditIDs    respjson.Field
		BillingFrequency        respjson.Field
		CommitIDs               respjson.Field
		PresentationGroupValues respjson.Field
		PricingGroupValues      respjson.Field
		ProductID               respjson.Field
		ProductTags             respjson.Field
		RecurringCommitIDs      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddOverrideOverrideSpecifier) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddOverrideOverrideSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddOverrideOverwriteRate struct {
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "TIERED_PERCENTAGE",
	// "CUSTOM".
	RateType   string                `json:"rate_type" api:"required"`
	CreditType shared.CreditTypeData `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]any `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier `json:"tiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RateType    respjson.Field
		CreditType  respjson.Field
		CustomRate  respjson.Field
		IsProrated  respjson.Field
		Price       respjson.Field
		Quantity    respjson.Field
		Tiers       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddOverrideOverwriteRate) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddOverrideOverwriteRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddOverrideProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddOverrideProduct) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddOverrideProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractGetEditHistoryResponseDataAddRecurringCommitAccessAmount `json:"access_amount" api:"required"`
	// The date this recurring commit's billing periods are anchored to.
	AnchorDate time.Time `json:"anchor_date" api:"required" format:"date-time"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractGetEditHistoryResponseDataAddRecurringCommitCommitDuration `json:"commit_duration" api:"required"`
	// Will be passed down to the individual commits
	Priority float64                                                       `json:"priority" api:"required"`
	Product  V2ContractGetEditHistoryResponseDataAddRecurringCommitProduct `json:"product" api:"required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type" api:"required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                                       `json:"applicable_product_tags"`
	Contract              V2ContractGetEditHistoryResponseDataAddRecurringCommitContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount V2ContractGetEditHistoryResponseDataAddRecurringCommitInvoiceAmount `json:"invoice_amount"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration"`
	// Rounding configuration for prorated recurring commit amounts.
	ProrationRounding V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRounding `json:"proration_rounding" api:"nullable"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates. - Daily recurring commits have a limit of one per contract, and
	// are unable to be created with seat-based subscriptions
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY", "DAILY".
	RecurrenceFrequency string `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []shared.CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig shared.RecurringCommitSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessAmount           respjson.Field
		AnchorDate             respjson.Field
		CommitDuration         respjson.Field
		Priority               respjson.Field
		Product                respjson.Field
		RateType               respjson.Field
		StartingAt             respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Contract               respjson.Field
		Description            respjson.Field
		EndingBefore           respjson.Field
		HierarchyConfiguration respjson.Field
		InvoiceAmount          respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Proration              respjson.Field
		ProrationRounding      respjson.Field
		RecurrenceFrequency    respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		SubscriptionConfig     respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCommit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type V2ContractGetEditHistoryResponseDataAddRecurringCommitAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	Quantity     float64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		UnitPrice    respjson.Field
		Quantity     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitAccessAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time the created commits will be valid for
type V2ContractGetEditHistoryResponseDataAddRecurringCommitCommitDuration struct {
	Value float64 `json:"value" api:"required"`
	// Any of "PERIODS".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitCommitDuration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitProduct) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitContract struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitContract) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount the customer should be billed for the commit. Not required.
type V2ContractGetEditHistoryResponseDataAddRecurringCommitInvoiceAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	Quantity     float64 `json:"quantity" api:"required"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		Quantity     respjson.Field
		UnitPrice    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitInvoiceAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rounding configuration for prorated recurring commit amounts.
type V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRounding struct {
	Access  V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRoundingAccess  `json:"access"`
	Invoice V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRoundingInvoice `json:"invoice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Access      respjson.Field
		Invoice     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRounding) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRoundingAccess) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRoundingInvoice struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRoundingInvoice) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitProrationRoundingInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddRecurringCredit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractGetEditHistoryResponseDataAddRecurringCreditAccessAmount `json:"access_amount" api:"required"`
	// The date this recurring commit's billing periods are anchored to.
	AnchorDate time.Time `json:"anchor_date" api:"required" format:"date-time"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractGetEditHistoryResponseDataAddRecurringCreditCommitDuration `json:"commit_duration" api:"required"`
	// Will be passed down to the individual commits
	Priority float64                                                       `json:"priority" api:"required"`
	Product  V2ContractGetEditHistoryResponseDataAddRecurringCreditProduct `json:"product" api:"required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type" api:"required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                                       `json:"applicable_product_tags"`
	Contract              V2ContractGetEditHistoryResponseDataAddRecurringCreditContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration"`
	// Rounding configuration for prorated recurring credit amounts.
	ProrationRounding V2ContractGetEditHistoryResponseDataAddRecurringCreditProrationRounding `json:"proration_rounding" api:"nullable"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates. - Daily recurring commits have a limit of one per contract, and
	// are unable to be created with seat-based subscriptions
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY", "DAILY".
	RecurrenceFrequency string `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []shared.CommitSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig shared.RecurringCommitSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessAmount           respjson.Field
		AnchorDate             respjson.Field
		CommitDuration         respjson.Field
		Priority               respjson.Field
		Product                respjson.Field
		RateType               respjson.Field
		StartingAt             respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Contract               respjson.Field
		Description            respjson.Field
		EndingBefore           respjson.Field
		HierarchyConfiguration respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Proration              respjson.Field
		ProrationRounding      respjson.Field
		RecurrenceFrequency    respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		SubscriptionConfig     respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCredit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type V2ContractGetEditHistoryResponseDataAddRecurringCreditAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	Quantity     float64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID respjson.Field
		UnitPrice    respjson.Field
		Quantity     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditAccessAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time the created commits will be valid for
type V2ContractGetEditHistoryResponseDataAddRecurringCreditCommitDuration struct {
	Value float64 `json:"value" api:"required"`
	// Any of "PERIODS".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditCommitDuration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditProduct) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditContract struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditContract) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rounding configuration for prorated recurring credit amounts.
type V2ContractGetEditHistoryResponseDataAddRecurringCreditProrationRounding struct {
	Access V2ContractGetEditHistoryResponseDataAddRecurringCreditProrationRoundingAccess `json:"access"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Access      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditProrationRounding) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditProrationRoundingAccess) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddResellerRoyalty struct {
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType          string    `json:"reseller_type" api:"required"`
	ApplicableProductIDs  []string  `json:"applicable_product_ids"`
	ApplicableProductTags []string  `json:"applicable_product_tags"`
	AwsAccountNumber      string    `json:"aws_account_number"`
	AwsOfferID            string    `json:"aws_offer_id"`
	AwsPayerReferenceID   string    `json:"aws_payer_reference_id"`
	EndingBefore          time.Time `json:"ending_before" api:"nullable" format:"date-time"`
	Fraction              float64   `json:"fraction"`
	GcpAccountID          string    `json:"gcp_account_id"`
	GcpOfferID            string    `json:"gcp_offer_id"`
	NetsuiteResellerID    string    `json:"netsuite_reseller_id"`
	ResellerContractValue float64   `json:"reseller_contract_value"`
	StartingAt            time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResellerType          respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		AwsAccountNumber      respjson.Field
		AwsOfferID            respjson.Field
		AwsPayerReferenceID   respjson.Field
		EndingBefore          respjson.Field
		Fraction              respjson.Field
		GcpAccountID          respjson.Field
		GcpOfferID            respjson.Field
		NetsuiteResellerID    respjson.Field
		ResellerContractValue respjson.Field
		StartingAt            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddResellerRoyalty) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddResellerRoyalty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddScheduledCharge struct {
	ID       string                                                        `json:"id" api:"required" format:"uuid"`
	Product  V2ContractGetEditHistoryResponseDataAddScheduledChargeProduct `json:"product" api:"required"`
	Schedule shared.SchedulePointInTime                                    `json:"schedule" api:"required"`
	// displayed on invoices
	Name string `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Product              respjson.Field
		Schedule             respjson.Field
		Name                 respjson.Field
		NetsuiteSalesOrderID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddScheduledCharge) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddScheduledChargeProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddScheduledChargeProduct) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddScheduledChargeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddSubscription struct {
	// Previous, current, and next billing periods for the subscription.
	BillingPeriods V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriods `json:"billing_periods" api:"required"`
	// Any of "ADVANCE", "ARREARS".
	CollectionSchedule string                                                       `json:"collection_schedule" api:"required"`
	Proration          V2ContractGetEditHistoryResponseDataAddSubscriptionProration `json:"proration" api:"required"`
	// Determines how the subscription's quantity is controlled. Defaults to
	// QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified
	// directly on the subscription. `initial_quantity` must be provided with this
	// option. Compatible with recurring commits/credits that use POOLED allocation.
	// **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add
	// user_123) to increment and decrement a subscription quantity, rather than
	// directly providing the quantity. You must use a **SEAT_BASED** subscription to
	// use a linked recurring credit with an allocation per seat. `seat_config` must be
	// provided with this option.
	//
	// Any of "SEAT_BASED", "QUANTITY_ONLY".
	QuantityManagementMode string `json:"quantity_management_mode" api:"required"`
	// List of quantity schedule items for the subscription. Only includes the current
	// quantity and future quantity changes.
	QuantitySchedule   []V2ContractGetEditHistoryResponseDataAddSubscriptionQuantitySchedule `json:"quantity_schedule" api:"required"`
	StartingAt         time.Time                                                             `json:"starting_at" api:"required" format:"date-time"`
	SubscriptionRate   V2ContractGetEditHistoryResponseDataAddSubscriptionSubscriptionRate   `json:"subscription_rate" api:"required"`
	ID                 string                                                                `json:"id" format:"uuid"`
	BillingCycleConfig V2ContractGetEditHistoryResponseDataAddSubscriptionBillingCycleConfig `json:"billing_cycle_config"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields     map[string]string `json:"custom_fields"`
	Description      string            `json:"description"`
	EndingBefore     time.Time         `json:"ending_before" format:"date-time"`
	FiatCreditTypeID string            `json:"fiat_credit_type_id" format:"uuid"`
	Name             string            `json:"name"`
	// Custom fields from the subscription product referenced by
	// `subscription_rate.product`. These are distinct from the subscription instance's
	// `custom_fields`.
	ProductCustomFields map[string]string                                             `json:"product_custom_fields"`
	SeatConfig          V2ContractGetEditHistoryResponseDataAddSubscriptionSeatConfig `json:"seat_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingPeriods         respjson.Field
		CollectionSchedule     respjson.Field
		Proration              respjson.Field
		QuantityManagementMode respjson.Field
		QuantitySchedule       respjson.Field
		StartingAt             respjson.Field
		SubscriptionRate       respjson.Field
		ID                     respjson.Field
		BillingCycleConfig     respjson.Field
		CustomFields           respjson.Field
		Description            respjson.Field
		EndingBefore           respjson.Field
		FiatCreditTypeID       respjson.Field
		Name                   respjson.Field
		ProductCustomFields    respjson.Field
		SeatConfig             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscription) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Previous, current, and next billing periods for the subscription.
type V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriods struct {
	Current  V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsCurrent  `json:"current"`
	Next     V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsNext     `json:"next"`
	Previous V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsPrevious `json:"previous"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriods) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsCurrent struct {
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsCurrent) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsCurrent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsNext struct {
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsNext) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsNext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsPrevious struct {
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsPrevious) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionBillingPeriodsPrevious) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionProration struct {
	// Any of "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE".
	InvoiceBehavior string                                                               `json:"invoice_behavior" api:"required"`
	IsProrated      bool                                                                 `json:"is_prorated" api:"required"`
	Rounding        V2ContractGetEditHistoryResponseDataAddSubscriptionProrationRounding `json:"rounding"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InvoiceBehavior respjson.Field
		IsProrated      respjson.Field
		Rounding        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscriptionProration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionProration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionProrationRounding struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscriptionProrationRounding) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionQuantitySchedule struct {
	Quantity     float64   `json:"quantity" api:"required"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity     respjson.Field
		StartingAt   respjson.Field
		EndingBefore respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscriptionQuantitySchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionQuantitySchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionSubscriptionRate struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string                                                                     `json:"billing_frequency" api:"required"`
	Product          V2ContractGetEditHistoryResponseDataAddSubscriptionSubscriptionRateProduct `json:"product" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFrequency respjson.Field
		Product          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscriptionSubscriptionRate) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionSubscriptionRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionSubscriptionRateProduct struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscriptionSubscriptionRateProduct) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionSubscriptionRateProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionBillingCycleConfig struct {
	// The date this subscription's billing cycle is anchored to.
	AnchorDate time.Time `json:"anchor_date" api:"required" format:"date-time"`
	// Controls whether this subscription consolidates onto usage invoices or gets its
	// own scheduled invoice.
	//
	// Any of "ON_SCHEDULED_INVOICE", "ON_USAGE_INVOICE".
	InvoicePlacement string `json:"invoice_placement" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnchorDate       respjson.Field
		InvoicePlacement respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscriptionBillingCycleConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionBillingCycleConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionSeatConfig struct {
	// The property name, sent on usage events, that identifies the seat ID associated
	// with the usage event. For example, the property name might be seat_id or
	// user_id. The property must be set as a group key on billable metrics and a
	// presentation/pricing group key on contract products. This allows linked
	// recurring credits with an allocation per seat to be consumed by only one seat's
	// usage.
	SeatGroupKey string `json:"seat_group_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SeatGroupKey respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddSubscriptionSeatConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionSeatConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataAddUsageFilter struct {
	GroupKey    string   `json:"group_key" api:"required"`
	GroupValues []string `json:"group_values" api:"required"`
	// This will match contract starting_at value if usage filter is active from the
	// beginning of the contract.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// This will match contract ending_before value if usage filter is active until the
	// end of the contract. It will be undefined if the contract is open-ended.
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupKey     respjson.Field
		GroupValues  respjson.Field
		StartingAt   respjson.Field
		EndingBefore respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataAddUsageFilter) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataAddUsageFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataArchiveCommit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataArchiveCommit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataArchiveCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataArchiveCredit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataArchiveCredit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataArchiveCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataArchiveScheduledCharge struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataArchiveScheduledCharge) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataArchiveScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataRemoveOverride struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataRemoveOverride) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataRemoveOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCommit struct {
	ID             string                                                         `json:"id" api:"required" format:"uuid"`
	AccessSchedule V2ContractGetEditHistoryResponseDataUpdateCommitAccessSchedule `json:"access_schedule"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids" api:"nullable" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags" api:"nullable"`
	Description           string   `json:"description"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration                             `json:"hierarchy_configuration"`
	InvoiceSchedule        V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceSchedule `json:"invoice_schedule"`
	Name                   string                                                          `json:"name"`
	NetsuiteSalesOrderID   string                                                          `json:"netsuite_sales_order_id" api:"nullable"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority  float64 `json:"priority" api:"nullable"`
	ProductID string  `json:"product_id" format:"uuid"`
	// If set, the commit's rate type was updated to the specified value.
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction" api:"nullable"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessSchedule         respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Description            respjson.Field
		HierarchyConfiguration respjson.Field
		InvoiceSchedule        respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Priority               respjson.Field
		ProductID              respjson.Field
		RateType               respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCommit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataUpdateCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCommitAccessSchedule struct {
	AddScheduleItems    []V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleUpdateScheduleItem `json:"update_schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddScheduleItems    respjson.Field
		RemoveScheduleItems respjson.Field
		UpdateScheduleItems respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCommitAccessSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCommitAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleAddScheduleItem struct {
	Amount float64 `json:"amount" api:"required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount       respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleAddScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleRemoveScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleUpdateScheduleItem struct {
	ID     string  `json:"id" api:"required" format:"uuid"`
	Amount float64 `json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Amount       respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleUpdateScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCommitAccessScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceSchedule struct {
	AddScheduleItems    []V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleUpdateScheduleItem `json:"update_schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddScheduleItems    respjson.Field
		RemoveScheduleItems respjson.Field
		UpdateScheduleItems respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleAddScheduleItem struct {
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	Amount    float64   `json:"amount"`
	Quantity  float64   `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Amount      respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleAddScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleRemoveScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleUpdateScheduleItem struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	Amount    float64   `json:"amount"`
	Quantity  float64   `json:"quantity"`
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	UnitPrice float64   `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Quantity    respjson.Field
		Timestamp   respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleUpdateScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCredit struct {
	ID             string                                                         `json:"id" api:"required" format:"uuid"`
	AccessSchedule V2ContractGetEditHistoryResponseDataUpdateCreditAccessSchedule `json:"access_schedule"`
	// Which products the credit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the credit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids" api:"nullable" format:"uuid"`
	// Which tags the credit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the credit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags" api:"nullable"`
	Description           string   `json:"description"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	Name                   string                              `json:"name"`
	NetsuiteSalesOrderID   string                              `json:"netsuite_sales_order_id" api:"nullable"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority  float64 `json:"priority" api:"nullable"`
	ProductID string  `json:"product_id" format:"uuid"`
	// If set, the credit's rate type was updated to the specified value.
	//
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction" api:"nullable"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessSchedule         respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
		Description            respjson.Field
		HierarchyConfiguration respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Priority               respjson.Field
		ProductID              respjson.Field
		RateType               respjson.Field
		RolloverFraction       respjson.Field
		Specifiers             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCredit) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataUpdateCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCreditAccessSchedule struct {
	AddScheduleItems    []V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleUpdateScheduleItem `json:"update_schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddScheduleItems    respjson.Field
		RemoveScheduleItems respjson.Field
		UpdateScheduleItems respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCreditAccessSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCreditAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleAddScheduleItem struct {
	Amount float64 `json:"amount" api:"required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount       respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleAddScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleRemoveScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleUpdateScheduleItem struct {
	ID     string  `json:"id" api:"required" format:"uuid"`
	Amount float64 `json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Amount       respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleUpdateScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateCreditAccessScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateDiscount struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields         map[string]string `json:"custom_fields"`
	Name                 string            `json:"name"`
	NetsuiteSalesOrderID string            `json:"netsuite_sales_order_id"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule V2ContractGetEditHistoryResponseDataUpdateDiscountSchedule `json:"schedule"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CustomFields         respjson.Field
		Name                 respjson.Field
		NetsuiteSalesOrderID respjson.Field
		Schedule             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateDiscount) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataUpdateDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must provide either schedule_items or recurring_schedule.
type V2ContractGetEditHistoryResponseDataUpdateDiscountSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID string `json:"credit_type_id" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice bool `json:"do_not_invoice"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V2ContractGetEditHistoryResponseDataUpdateDiscountScheduleRecurringSchedule `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V2ContractGetEditHistoryResponseDataUpdateDiscountScheduleScheduleItem `json:"schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTypeID      respjson.Field
		DoNotInvoice      respjson.Field
		RecurringSchedule respjson.Field
		ScheduleItems     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateDiscountSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateDiscountSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V2ContractGetEditHistoryResponseDataUpdateDiscountScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution" api:"required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency" api:"required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount float64 `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity float64 `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice float64 `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountDistribution respjson.Field
		EndingBefore       respjson.Field
		Frequency          respjson.Field
		StartingAt         respjson.Field
		Amount             respjson.Field
		Quantity           respjson.Field
		UnitPrice          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateDiscountScheduleRecurringSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateDiscountScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateDiscountScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount float64 `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity float64 `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice float64 `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Amount      respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateDiscountScheduleScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateDiscountScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfiguration struct {
	Commit V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommit `json:"commit"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID    string                                                                                              `json:"custom_credit_type_id" api:"nullable" format:"uuid"`
	DiscountConfiguration V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration `json:"discount_configuration" api:"nullable"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                       `json:"is_enabled"`
	PaymentGateConfig shared.PaymentGateConfigV2 `json:"payment_gate_config"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount float64 `json:"recharge_to_amount"`
	// Specify the threshold amount for the contract. Each time the contract's balance
	// lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount"`
	// Determines which balances are excluded from remaining balance calculation for
	// threshold billing.
	ThresholdBalanceSpecifiers []V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier `json:"threshold_balance_specifiers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit                     respjson.Field
		CustomCreditTypeID         respjson.Field
		DiscountConfiguration      respjson.Field
		IsEnabled                  respjson.Field
		PaymentGateConfig          respjson.Field
		RechargeToAmount           respjson.Field
		ThresholdAmount            respjson.Field
		ThresholdBalanceSpecifiers respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommit struct {
	// Which products the threshold commit applies to. If both applicable_product_ids
	// and applicable_product_tags are not provided, the commit applies to all
	// products.
	ApplicableProductIDs []string `json:"applicable_product_ids" api:"nullable" format:"uuid"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags" api:"nullable"`
	// The length of time the created commit will be valid, starting from the end of
	// the invoice's service period. Set to null to clear a previously configured
	// duration.
	Duration V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitDuration `json:"duration" api:"nullable"`
	// Whether the created commits will be charged at commit rate or list rate. Set to
	// null to clear a previously configured rate type.
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type" api:"nullable"`
	// Fraction of the created commit's unused balance that will roll over. Must be
	// between 0 and 1. Set to null to clear a previously configured rollover fraction.
	RolloverFraction float64 `json:"rollover_fraction" api:"nullable"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		Duration              respjson.Field
		RateType              respjson.Field
		RolloverFraction      respjson.Field
		Specifiers            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
	shared.UpdateBaseThresholdCommit
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommit) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The length of time the created commit will be valid, starting from the end of
// the invoice's service period. Set to null to clear a previously configured
// duration.
type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit" api:"required"`
	Value int64  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitDuration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration struct {
	// Update the discount cap. Set to null to remove an existing cap.
	Cap V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap `json:"cap" api:"nullable"`
	// The fraction of the original amount that the customer pays after applying the
	// discount. Set to null to remove the discount fraction. For example, 0.85 means
	// the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction float64 `json:"payment_fraction" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cap             respjson.Field
		PaymentFraction respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update the discount cap. Set to null to remove an existing cap.
type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount float64 `json:"amount" api:"required"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias string `json:"spend_tracker_alias" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount            respjson.Field
		SpendTrackerAlias respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier struct {
	Exclude []V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude `json:"exclude" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exclude     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude struct {
	CustomFieldFilters []V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter `json:"custom_field_filters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomFieldFilters respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter struct {
	// Any of "Commit", "ContractCredit", "ContractCreditOrCommit".
	Entity string `json:"entity" api:"required"`
	Key    string `json:"key" api:"required"`
	Value  string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entity      respjson.Field
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCommit struct {
	ID            string                                                                 `json:"id" api:"required" format:"uuid"`
	AccessAmount  V2ContractGetEditHistoryResponseDataUpdateRecurringCommitAccessAmount  `json:"access_amount"`
	EndingBefore  time.Time                                                              `json:"ending_before" format:"date-time"`
	InvoiceAmount V2ContractGetEditHistoryResponseDataUpdateRecurringCommitInvoiceAmount `json:"invoice_amount"`
	// Rounding configuration for prorated recurring commit amounts.
	ProrationRounding V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRounding `json:"proration_rounding" api:"nullable"`
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType string `json:"rate_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AccessAmount      respjson.Field
		EndingBefore      respjson.Field
		InvoiceAmount     respjson.Field
		ProrationRounding respjson.Field
		RateType          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateRecurringCommit) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCommitAccessAmount struct {
	Quantity  float64 `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateRecurringCommitAccessAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCommitInvoiceAmount struct {
	Quantity  float64 `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateRecurringCommitInvoiceAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rounding configuration for prorated recurring commit amounts.
type V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRounding struct {
	Access  V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRoundingAccess  `json:"access"`
	Invoice V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRoundingInvoice `json:"invoice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Access      respjson.Field
		Invoice     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRounding) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRoundingAccess) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRoundingInvoice struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRoundingInvoice) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCommitProrationRoundingInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCredit struct {
	ID           string                                                                `json:"id" api:"required" format:"uuid"`
	AccessAmount V2ContractGetEditHistoryResponseDataUpdateRecurringCreditAccessAmount `json:"access_amount"`
	EndingBefore time.Time                                                             `json:"ending_before" format:"date-time"`
	// Rounding configuration for prorated recurring credit amounts.
	ProrationRounding V2ContractGetEditHistoryResponseDataUpdateRecurringCreditProrationRounding `json:"proration_rounding" api:"nullable"`
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType string `json:"rate_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AccessAmount      respjson.Field
		EndingBefore      respjson.Field
		ProrationRounding respjson.Field
		RateType          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateRecurringCredit) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCreditAccessAmount struct {
	Quantity  float64 `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateRecurringCreditAccessAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rounding configuration for prorated recurring credit amounts.
type V2ContractGetEditHistoryResponseDataUpdateRecurringCreditProrationRounding struct {
	Access V2ContractGetEditHistoryResponseDataUpdateRecurringCreditProrationRoundingAccess `json:"access"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Access      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateRecurringCreditProrationRounding) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCreditProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCreditProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecimalPlaces  respjson.Field
		RoundingMethod respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateRecurringCreditProrationRoundingAccess) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCreditProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateRefundInvoice struct {
	Date      time.Time `json:"date" api:"required" format:"date-time"`
	InvoiceID string    `json:"invoice_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		InvoiceID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateRefundInvoice) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataUpdateRefundInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateScheduledCharge struct {
	ID                   string                                                                   `json:"id" api:"required" format:"uuid"`
	InvoiceSchedule      V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceSchedule `json:"invoice_schedule"`
	Name                 string                                                                   `json:"name"`
	NetsuiteSalesOrderID string                                                                   `json:"netsuite_sales_order_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		InvoiceSchedule      respjson.Field
		Name                 respjson.Field
		NetsuiteSalesOrderID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateScheduledCharge) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceSchedule struct {
	AddScheduleItems    []V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem `json:"update_schedule_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddScheduleItems    respjson.Field
		RemoveScheduleItems respjson.Field
		UpdateScheduleItems respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleAddScheduleItem struct {
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	Amount    float64   `json:"amount"`
	Quantity  float64   `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Amount      respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleAddScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	Amount    float64   `json:"amount"`
	Quantity  float64   `json:"quantity"`
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	UnitPrice float64   `json:"unit_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Quantity    respjson.Field
		Timestamp   respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfiguration struct {
	Commit                shared.UpdateBaseThresholdCommit                                                           `json:"commit"`
	DiscountConfiguration V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationDiscountConfiguration `json:"discount_configuration" api:"nullable"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                       `json:"is_enabled"`
	PaymentGateConfig shared.PaymentGateConfigV2 `json:"payment_gate_config"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit                respjson.Field
		DiscountConfiguration respjson.Field
		IsEnabled             respjson.Field
		PaymentGateConfig     respjson.Field
		ThresholdAmount       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationDiscountConfiguration struct {
	// Update the discount cap. Set to null to remove an existing cap.
	Cap V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationDiscountConfigurationCap `json:"cap" api:"nullable"`
	// The fraction of the original amount that the customer pays after applying the
	// discount. Set to null to remove the discount fraction. For example, 0.85 means
	// the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction float64 `json:"payment_fraction" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cap             respjson.Field
		PaymentFraction respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationDiscountConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationDiscountConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update the discount cap. Set to null to remove an existing cap.
type V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationDiscountConfigurationCap struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount float64 `json:"amount" api:"required"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias string `json:"spend_tracker_alias" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount            respjson.Field
		SpendTrackerAlias respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationDiscountConfigurationCap) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationDiscountConfigurationCap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateSubscription struct {
	ID              string                                                                 `json:"id" api:"required" format:"uuid"`
	EndingBefore    time.Time                                                              `json:"ending_before" format:"date-time"`
	Name            string                                                                 `json:"name"`
	QuantityUpdates []V2ContractGetEditHistoryResponseDataUpdateSubscriptionQuantityUpdate `json:"quantity_updates"`
	// Manage subscription seats for subscriptions in SEAT_BASED mode.
	SeatUpdates V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdates `json:"seat_updates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		EndingBefore    respjson.Field
		Name            respjson.Field
		QuantityUpdates respjson.Field
		SeatUpdates     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateSubscription) RawJSON() string { return r.JSON.raw }
func (r *V2ContractGetEditHistoryResponseDataUpdateSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateSubscriptionQuantityUpdate struct {
	StartingAt    time.Time `json:"starting_at" api:"required" format:"date-time"`
	Quantity      float64   `json:"quantity"`
	QuantityDelta float64   `json:"quantity_delta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StartingAt    respjson.Field
		Quantity      respjson.Field
		QuantityDelta respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateSubscriptionQuantityUpdate) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateSubscriptionQuantityUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Manage subscription seats for subscriptions in SEAT_BASED mode.
type V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdates struct {
	// Adds seat IDs to the subscription. If there are unassigned seats, the new seat
	// IDs will fill these unassigned seats and not increase the total subscription
	// quantity. Otherwise, if there are more new seat IDs than unassigned seats, the
	// total subscription quantity will increase.
	AddSeatIDs []V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesAddSeatID `json:"add_seat_ids"`
	// Adds unassigned seats to the subscription. This will increase the total
	// subscription quantity.
	AddUnassignedSeats []V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesAddUnassignedSeat `json:"add_unassigned_seats"`
	// Removes seat IDs from the subscription, if possible. If a seat ID is removed,
	// the total subscription quantity will decrease. Otherwise, if the seat ID is not
	// found on the subscription, this is a no-op.
	RemoveSeatIDs []V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesRemoveSeatID `json:"remove_seat_ids"`
	// Removes unassigned seats from the subscription. This will decrease the total
	// subscription quantity if there are are unassigned seats.
	RemoveUnassignedSeats []V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat `json:"remove_unassigned_seats"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddSeatIDs            respjson.Field
		AddUnassignedSeats    respjson.Field
		RemoveSeatIDs         respjson.Field
		RemoveUnassignedSeats respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdates) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesAddSeatID struct {
	SeatIDs []string `json:"seat_ids" api:"required"`
	// Assigned seats will be added/removed starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SeatIDs     respjson.Field
		StartingAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesAddSeatID) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesAddSeatID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesAddUnassignedSeat struct {
	// The number of unassigned seats on the subscription will increase/decrease by
	// this delta. Must be greater than 0.
	Quantity float64 `json:"quantity" api:"required"`
	// Unassigned seats will be updated starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		StartingAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesAddUnassignedSeat) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesAddUnassignedSeat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesRemoveSeatID struct {
	SeatIDs []string `json:"seat_ids" api:"required"`
	// Assigned seats will be added/removed starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SeatIDs     respjson.Field
		StartingAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesRemoveSeatID) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesRemoveSeatID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat struct {
	// The number of unassigned seats on the subscription will increase/decrease by
	// this delta. Must be greater than 0.
	Quantity float64 `json:"quantity" api:"required"`
	// Unassigned seats will be updated starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		StartingAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetParams struct {
	ContractID string `json:"contract_id" api:"required" format:"uuid"`
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Optional RFC 3339 timestamp. Return the contract as of this date. Cannot be used
	// with include_ledgers parameter.
	AsOfDate param.Opt[time.Time] `json:"as_of_date,omitzero" format:"date-time"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the query to be slower.
	IncludeBalance param.Opt[bool] `json:"include_balance,omitzero"`
	// Include commit/credit ledgers in the response. Setting this flag may cause the
	// query to be slower. Cannot be used with as_of_date parameter.
	IncludeLedgers param.Opt[bool] `json:"include_ledgers,omitzero"`
	paramObj
}

func (r V2ContractGetParams) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractListParams struct {
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Optional RFC 3339 timestamp. Only include contracts active on the provided date.
	// This cannot be provided if starting_at filter is provided.
	CoveringDate param.Opt[time.Time] `json:"covering_date,omitzero" format:"date-time"`
	// Include archived contracts in the response.
	IncludeArchived param.Opt[bool] `json:"include_archived,omitzero"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the response to be slower.
	IncludeBalance param.Opt[bool] `json:"include_balance,omitzero"`
	// Include commit/credit ledgers in the response. Setting this flag may cause the
	// response to be slower.
	IncludeLedgers param.Opt[bool] `json:"include_ledgers,omitzero"`
	// Optional RFC 3339 timestamp. Only include contracts that started on or after
	// this date. This cannot be provided if covering_date filter is provided.
	StartingAt param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V2ContractListParams) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParams struct {
	// ID of the contract being edited
	ContractID string `json:"contract_id" api:"required" format:"uuid"`
	// ID of the customer whose contract is being edited
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// RFC 3339 timestamp indicating when the contract will end (exclusive).
	UpdateContractEndDate param.Opt[time.Time] `json:"update_contract_end_date,omitzero" format:"date-time"`
	// Value to update the contract name to. If not provided, the contract name will
	// remain unchanged.
	UpdateContractName param.Opt[string] `json:"update_contract_name,omitzero"`
	// Number of days after issuance of invoice after which the invoice is due (e.g.
	// Net 30).
	UpdateNetPaymentTermsDays param.Opt[float64] `json:"update_net_payment_terms_days,omitzero"`
	// If true, allows setting the contract end date earlier than the end_timestamp of
	// existing finalized invoices. Finalized invoices will be unchanged; if you want
	// to incorporate the new end date, you can void and regenerate finalized usage
	// invoices. Defaults to true.
	AllowContractEndingBeforeFinalizedInvoice param.Opt[bool] `json:"allow_contract_ending_before_finalized_invoice,omitzero"`
	// Optional uniqueness key to prevent duplicate contract edits.
	UniquenessKey param.Opt[string] `json:"uniqueness_key,omitzero"`
	// Update the billing provider configuration on the contract. Currently only
	// supports adding a billing provider configuration to a contract that does not
	// already have one.
	AddBillingProviderConfigurationUpdate   V2ContractEditParamsAddBillingProviderConfigurationUpdate `json:"add_billing_provider_configuration_update,omitzero"`
	AddCommits                              []V2ContractEditParamsAddCommit                           `json:"add_commits,omitzero"`
	AddCredits                              []V2ContractEditParamsAddCredit                           `json:"add_credits,omitzero"`
	AddDiscounts                            []V2ContractEditParamsAddDiscount                         `json:"add_discounts,omitzero"`
	AddOverrides                            []V2ContractEditParamsAddOverride                         `json:"add_overrides,omitzero"`
	AddPrepaidBalanceThresholdConfiguration shared.PrepaidBalanceThresholdConfigurationV2Param        `json:"add_prepaid_balance_threshold_configuration,omitzero"`
	// This field's availability is dependent on your client's configuration.
	AddProfessionalServices []V2ContractEditParamsAddProfessionalService `json:"add_professional_services,omitzero"`
	AddRecurringCommits     []V2ContractEditParamsAddRecurringCommit     `json:"add_recurring_commits,omitzero"`
	AddRecurringCredits     []V2ContractEditParamsAddRecurringCredit     `json:"add_recurring_credits,omitzero"`
	AddResellerRoyalties    []V2ContractEditParamsAddResellerRoyalty     `json:"add_reseller_royalties,omitzero"`
	// Update the revenue system configuration on the contract. Currently only supports
	// adding a revenue system configuration to a contract that does not already have
	// one.
	AddRevenueSystemConfigurationUpdate V2ContractEditParamsAddRevenueSystemConfigurationUpdate `json:"add_revenue_system_configuration_update,omitzero"`
	AddScheduledCharges                 []V2ContractEditParamsAddScheduledCharge                `json:"add_scheduled_charges,omitzero"`
	AddSpendThresholdConfiguration      shared.SpendThresholdConfigurationV2Param               `json:"add_spend_threshold_configuration,omitzero"`
	// Spend trackers to add to this contract. Aliases must be unique within a
	// contract.
	AddSpendTrackers []V2ContractEditParamsAddSpendTracker `json:"add_spend_trackers,omitzero"`
	// Optional list of
	// [subscriptions](https://docs.metronome.com/manage-product-access/create-subscription/)
	// to add to the contract.
	AddSubscriptions []V2ContractEditParamsAddSubscription `json:"add_subscriptions,omitzero"`
	// IDs of commits to archive
	ArchiveCommits []V2ContractEditParamsArchiveCommit `json:"archive_commits,omitzero"`
	// IDs of credits to archive
	ArchiveCredits []V2ContractEditParamsArchiveCredit `json:"archive_credits,omitzero"`
	// IDs of scheduled charges to archive
	ArchiveScheduledCharges []V2ContractEditParamsArchiveScheduledCharge `json:"archive_scheduled_charges,omitzero"`
	// Aliases of spend trackers to archive.
	ArchiveSpendTrackers []string `json:"archive_spend_trackers,omitzero"`
	// IDs of overrides to remove
	RemoveOverrides                            []V2ContractEditParamsRemoveOverride                           `json:"remove_overrides,omitzero"`
	UpdateCommits                              []V2ContractEditParamsUpdateCommit                             `json:"update_commits,omitzero"`
	UpdateCredits                              []V2ContractEditParamsUpdateCredit                             `json:"update_credits,omitzero"`
	UpdatePrepaidBalanceThresholdConfiguration V2ContractEditParamsUpdatePrepaidBalanceThresholdConfiguration `json:"update_prepaid_balance_threshold_configuration,omitzero"`
	// Edits to these recurring commits will only affect commits whose access schedules
	// has not started. Expired commits, and commits with an active access schedule
	// will remain unchanged.
	UpdateRecurringCommits []V2ContractEditParamsUpdateRecurringCommit `json:"update_recurring_commits,omitzero"`
	// Edits to these recurring credits will only affect credits whose access schedules
	// has not started. Expired credits, and credits with an active access schedule
	// will remain unchanged.
	UpdateRecurringCredits            []V2ContractEditParamsUpdateRecurringCredit           `json:"update_recurring_credits,omitzero"`
	UpdateScheduledCharges            []V2ContractEditParamsUpdateScheduledCharge           `json:"update_scheduled_charges,omitzero"`
	UpdateSpendThresholdConfiguration V2ContractEditParamsUpdateSpendThresholdConfiguration `json:"update_spend_threshold_configuration,omitzero"`
	// Optional list of subscriptions to update.
	UpdateSubscriptions []V2ContractEditParamsUpdateSubscription `json:"update_subscriptions,omitzero"`
	paramObj
}

func (r V2ContractEditParams) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update the billing provider configuration on the contract. Currently only
// supports adding a billing provider configuration to a contract that does not
// already have one.
//
// The properties BillingProviderConfiguration, Schedule are required.
type V2ContractEditParamsAddBillingProviderConfigurationUpdate struct {
	BillingProviderConfiguration V2ContractEditParamsAddBillingProviderConfigurationUpdateBillingProviderConfiguration `json:"billing_provider_configuration,omitzero" api:"required"`
	// Indicates when the billing provider will be active on the contract. Any charges
	// accrued during the schedule will be billed to the indicated billing provider.
	Schedule V2ContractEditParamsAddBillingProviderConfigurationUpdateSchedule `json:"schedule,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddBillingProviderConfigurationUpdate) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddBillingProviderConfigurationUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddBillingProviderConfigurationUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParamsAddBillingProviderConfigurationUpdateBillingProviderConfiguration struct {
	BillingProviderConfigurationID param.Opt[string] `json:"billing_provider_configuration_id,omitzero" format:"uuid"`
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace", "metronome".
	BillingProvider string `json:"billing_provider,omitzero"`
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod string `json:"delivery_method,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddBillingProviderConfigurationUpdateBillingProviderConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddBillingProviderConfigurationUpdateBillingProviderConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddBillingProviderConfigurationUpdateBillingProviderConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddBillingProviderConfigurationUpdateBillingProviderConfiguration](
		"billing_provider", "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace", "quickbooks_online", "workday", "gcp_marketplace", "metronome",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddBillingProviderConfigurationUpdateBillingProviderConfiguration](
		"delivery_method", "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns",
	)
}

// Indicates when the billing provider will be active on the contract. Any charges
// accrued during the schedule will be billed to the indicated billing provider.
//
// The property EffectiveAt is required.
type V2ContractEditParamsAddBillingProviderConfigurationUpdateSchedule struct {
	// When the billing provider update will take effect.
	//
	// Any of "START_OF_CURRENT_PERIOD", "START_OF_NEXT_PERIOD".
	EffectiveAt string `json:"effective_at,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddBillingProviderConfigurationUpdateSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddBillingProviderConfigurationUpdateSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddBillingProviderConfigurationUpdateSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddBillingProviderConfigurationUpdateSchedule](
		"effective_at", "START_OF_CURRENT_PERIOD", "START_OF_NEXT_PERIOD",
	)
}

// The properties ProductID, Type are required.
type V2ContractEditParamsAddCommit struct {
	ProductID string `json:"product_id" api:"required" format:"uuid"`
	// Any of "PREPAID", "POSTPAID".
	Type string `json:"type,omitzero" api:"required"`
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
	AccessSchedule V2ContractEditParamsAddCommitAccessSchedule `json:"access_schedule,omitzero"`
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
	InvoiceSchedule V2ContractEditParamsAddCommitInvoiceSchedule `json:"invoice_schedule,omitzero"`
	// optionally payment gate this commit
	PaymentGateConfig V2ContractEditParamsAddCommitPaymentGateConfig `json:"payment_gate_config,omitzero"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	// Optional attributes for spend tracker integration. Immutable after creation.
	SpendTrackerAttributes V2ContractEditParamsAddCommitSpendTrackerAttributes `json:"spend_tracker_attributes,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddCommit) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddCommit](
		"type", "PREPAID", "POSTPAID",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddCommit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
//
// The property ScheduleItems is required.
type V2ContractEditParamsAddCommitAccessSchedule struct {
	ScheduleItems []V2ContractEditParamsAddCommitAccessScheduleScheduleItem `json:"schedule_items,omitzero" api:"required"`
	CreditTypeID  param.Opt[string]                                         `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r V2ContractEditParamsAddCommitAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCommitAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCommitAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, EndingBefore, StartingAt are required.
type V2ContractEditParamsAddCommitAccessScheduleScheduleItem struct {
	Amount float64 `json:"amount" api:"required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	paramObj
}

func (r V2ContractEditParamsAddCommitAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCommitAccessScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCommitAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match access_schedule
// amount. Optional for "PREPAID" commits: if not provided, this will be a
// "complimentary" commit with no invoice.
type V2ContractEditParamsAddCommitInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Opt[bool] `json:"do_not_invoice,omitzero"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V2ContractEditParamsAddCommitInvoiceScheduleRecurringSchedule `json:"recurring_schedule,omitzero"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V2ContractEditParamsAddCommitInvoiceScheduleScheduleItem `json:"schedule_items,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddCommitInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCommitInvoiceSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
//
// The properties AmountDistribution, EndingBefore, Frequency, StartingAt are
// required.
type V2ContractEditParamsAddCommitInvoiceScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution,omitzero" api:"required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency,omitzero" api:"required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
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

func (r V2ContractEditParamsAddCommitInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCommitInvoiceScheduleRecurringSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCommitInvoiceScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddCommitInvoiceScheduleRecurringSchedule](
		"amount_distribution", "DIVIDED", "DIVIDED_ROUNDED", "EACH",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddCommitInvoiceScheduleRecurringSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY",
	)
}

// The property Timestamp is required.
type V2ContractEditParamsAddCommitInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
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

func (r V2ContractEditParamsAddCommitInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCommitInvoiceScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCommitInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// optionally payment gate this commit
//
// The property PaymentGateType is required.
type V2ContractEditParamsAddCommitPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	//
	// Any of "NONE", "STRIPE", "EXTERNAL".
	PaymentGateType string `json:"payment_gate_type,omitzero" api:"required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig V2ContractEditParamsAddCommitPaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config,omitzero"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig V2ContractEditParamsAddCommitPaymentGateConfigStripeConfig `json:"stripe_config,omitzero"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	//
	// Any of "NONE", "STRIPE", "ANROK", "PRECALCULATED".
	TaxType string `json:"tax_type,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddCommitPaymentGateConfig) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCommitPaymentGateConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCommitPaymentGateConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddCommitPaymentGateConfig](
		"payment_gate_type", "NONE", "STRIPE", "EXTERNAL",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddCommitPaymentGateConfig](
		"tax_type", "NONE", "STRIPE", "ANROK", "PRECALCULATED",
	)
}

// Only applicable if using PRECALCULATED as your tax type.
//
// The property TaxAmount is required.
type V2ContractEditParamsAddCommitPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount" api:"required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Opt[string] `json:"tax_name,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddCommitPaymentGateConfigPrecalculatedTaxConfig) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCommitPaymentGateConfigPrecalculatedTaxConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCommitPaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only applicable if using STRIPE as your payment gateway type.
//
// The property PaymentType is required.
type V2ContractEditParamsAddCommitPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	//
	// Any of "INVOICE", "PAYMENT_INTENT".
	PaymentType string `json:"payment_type,omitzero" api:"required"`
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

func (r V2ContractEditParamsAddCommitPaymentGateConfigStripeConfig) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCommitPaymentGateConfigStripeConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCommitPaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddCommitPaymentGateConfigStripeConfig](
		"payment_type", "INVOICE", "PAYMENT_INTENT",
	)
}

// Optional attributes for spend tracker integration. Immutable after creation.
//
// The property CountsAsDiscounted is required.
type V2ContractEditParamsAddCommitSpendTrackerAttributes struct {
	// If true, this commit will be included in spend trackers with discounted set to
	// DISCOUNTED_ONLY
	CountsAsDiscounted bool `json:"counts_as_discounted" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddCommitSpendTrackerAttributes) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCommitSpendTrackerAttributes
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCommitSpendTrackerAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessSchedule, ProductID are required.
type V2ContractEditParamsAddCredit struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule V2ContractEditParamsAddCreditAccessSchedule `json:"access_schedule,omitzero" api:"required"`
	ProductID      string                                      `json:"product_id" api:"required" format:"uuid"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Opt[string] `json:"description,omitzero"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Opt[float64] `json:"rollover_fraction,omitzero"`
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
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddCredit) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddCredit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
}

// Schedule for distributing the credit to the customer.
//
// The property ScheduleItems is required.
type V2ContractEditParamsAddCreditAccessSchedule struct {
	ScheduleItems []V2ContractEditParamsAddCreditAccessScheduleScheduleItem `json:"schedule_items,omitzero" api:"required"`
	CreditTypeID  param.Opt[string]                                         `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r V2ContractEditParamsAddCreditAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCreditAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCreditAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, EndingBefore, StartingAt are required.
type V2ContractEditParamsAddCreditAccessScheduleScheduleItem struct {
	Amount float64 `json:"amount" api:"required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	paramObj
}

func (r V2ContractEditParamsAddCreditAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddCreditAccessScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddCreditAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ProductID, Schedule are required.
type V2ContractEditParamsAddDiscount struct {
	ProductID string `json:"product_id" api:"required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule V2ContractEditParamsAddDiscountSchedule `json:"schedule,omitzero" api:"required"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddDiscount) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddDiscount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must provide either schedule_items or recurring_schedule.
type V2ContractEditParamsAddDiscountSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Opt[bool] `json:"do_not_invoice,omitzero"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V2ContractEditParamsAddDiscountScheduleRecurringSchedule `json:"recurring_schedule,omitzero"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V2ContractEditParamsAddDiscountScheduleScheduleItem `json:"schedule_items,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddDiscountSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddDiscountSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddDiscountSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
//
// The properties AmountDistribution, EndingBefore, Frequency, StartingAt are
// required.
type V2ContractEditParamsAddDiscountScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution,omitzero" api:"required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency,omitzero" api:"required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
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

func (r V2ContractEditParamsAddDiscountScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddDiscountScheduleRecurringSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddDiscountScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddDiscountScheduleRecurringSchedule](
		"amount_distribution", "DIVIDED", "DIVIDED_ROUNDED", "EACH",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddDiscountScheduleRecurringSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY",
	)
}

// The property Timestamp is required.
type V2ContractEditParamsAddDiscountScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
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

func (r V2ContractEditParamsAddDiscountScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddDiscountScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddDiscountScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property StartingAt is required.
type V2ContractEditParamsAddOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	Entitled     param.Opt[bool]      `json:"entitled,omitzero"`
	// Indicates whether the override should only apply to commits. Defaults to
	// `false`. If `true`, you can specify relevant commits in `override_specifiers` by
	// passing `commit_ids`.
	IsCommitSpecific param.Opt[bool] `json:"is_commit_specific,omitzero"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Opt[float64] `json:"multiplier,omitzero"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides.
	// Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered
	// and multiplier overrides are prioritized by their priority value (lowest first).
	// Must be > 0.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// ID of the product whose rate is being overridden
	ProductID param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	// tags identifying products whose rates are being overridden
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If
	// provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers []V2ContractEditParamsAddOverrideOverrideSpecifier `json:"override_specifiers,omitzero"`
	// Required for OVERWRITE type.
	OverwriteRate V2ContractEditParamsAddOverrideOverwriteRate `json:"overwrite_rate,omitzero"`
	// Indicates whether the override applies to commit rates or list rates. Can only
	// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
	// `"LIST_RATE"`.
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	Target string `json:"target,omitzero"`
	// Required for TIERED type. Must have at least one tier.
	Tiers []V2ContractEditParamsAddOverrideTier `json:"tiers,omitzero"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	//
	// Any of "OVERWRITE", "MULTIPLIER", "TIERED".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddOverride) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddOverride](
		"target", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddOverride](
		"type", "OVERWRITE", "MULTIPLIER", "TIERED",
	)
}

type V2ContractEditParamsAddOverrideOverrideSpecifier struct {
	// If provided, the override will only apply to the product with the specified ID.
	ProductID param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of `product_id`, `product_tags`, `pricing_group_values`, or
	// `presentation_group_values`. Must be used instead of both `commit_ids` and
	// `recurring_commit_ids` If provided, the override will apply to any specified
	// commit, credit, recurring commit or recurring credit IDs.
	AnyCommitOrCreditIDs []string `json:"any_commit_or_credit_ids,omitzero"`
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero"`
	// If provided, the override will only apply to the specified commits. Can only be
	// used for commit specific overrides. If not provided, the override will apply to
	// all commits.
	CommitIDs []string `json:"commit_ids,omitzero"`
	// A map of group names to values. The override will only apply to line items with
	// the specified presentation group values. Can only be used for multiplier
	// overrides.
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
	paramObj
}

func (r V2ContractEditParamsAddOverrideOverrideSpecifier) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddOverrideOverrideSpecifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddOverrideOverrideSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddOverrideOverrideSpecifier](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// Required for OVERWRITE type.
//
// The property RateType is required.
type V2ContractEditParamsAddOverrideOverwriteRate struct {
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "TIERED_PERCENTAGE",
	// "CUSTOM".
	RateType     string            `json:"rate_type,omitzero" api:"required"`
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

func (r V2ContractEditParamsAddOverrideOverwriteRate) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddOverrideOverwriteRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddOverrideOverwriteRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddOverrideOverwriteRate](
		"rate_type", "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "TIERED_PERCENTAGE", "CUSTOM",
	)
}

// The property Multiplier is required.
type V2ContractEditParamsAddOverrideTier struct {
	Multiplier float64            `json:"multiplier" api:"required"`
	Size       param.Opt[float64] `json:"size,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddOverrideTier) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddOverrideTier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddOverrideTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxAmount, ProductID, Quantity, UnitPrice are required.
type V2ContractEditParamsAddProfessionalService struct {
	// Maximum amount for the term.
	MaxAmount float64 `json:"max_amount" api:"required"`
	ProductID string  `json:"product_id" api:"required" format:"uuid"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity float64 `json:"quantity" api:"required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified.
	UnitPrice   float64           `json:"unit_price" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddProfessionalService) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddProfessionalService
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddProfessionalService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessAmount, CommitDuration, Priority, ProductID, StartingAt are
// required.
type V2ContractEditParamsAddRecurringCommit struct {
	// The amount of commit to grant.
	AccessAmount V2ContractEditParamsAddRecurringCommitAccessAmount `json:"access_amount,omitzero" api:"required"`
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
	CommitDuration V2ContractEditParamsAddRecurringCommitCommitDuration `json:"commit_duration,omitzero" api:"required"`
	// Will be passed down to the individual commits
	Priority  float64 `json:"priority" api:"required"`
	ProductID string  `json:"product_id" api:"required" format:"uuid"`
	// determines the start time for the first commit
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
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
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam `json:"hierarchy_configuration,omitzero"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount V2ContractEditParamsAddRecurringCommitInvoiceAmount `json:"invoice_amount,omitzero"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration,omitzero"`
	// Optional rounding configuration for prorated recurring commit amounts.
	ProrationRounding V2ContractEditParamsAddRecurringCommitProrationRounding `json:"proration_rounding,omitzero"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates. - Daily recurring commits have a limit of one per contract, and
	// are unable to be created with seat-based subscriptions
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY", "DAILY".
	RecurrenceFrequency string `json:"recurrence_frequency,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig V2ContractEditParamsAddRecurringCommitSubscriptionConfig `json:"subscription_config,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCommit) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCommit](
		"proration", "NONE", "FIRST", "LAST", "FIRST_AND_LAST",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCommit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCommit](
		"recurrence_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY", "DAILY",
	)
}

// The amount of commit to grant.
//
// The properties CreditTypeID, UnitPrice are required.
type V2ContractEditParamsAddRecurringCommitAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	// This field is required unless a subscription is attached via
	// `subscription_config`.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCommitAccessAmount) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCommitAccessAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the length of the access schedule for each created commit/credit. The
// value represents the number of units. Unit defaults to "PERIODS", where the
// length of a period is determined by the recurrence_frequency.
//
// The property Value is required.
type V2ContractEditParamsAddRecurringCommitCommitDuration struct {
	Value float64 `json:"value" api:"required"`
	// Any of "PERIODS".
	Unit string `json:"unit,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCommitCommitDuration) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCommitCommitDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCommitCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCommitCommitDuration](
		"unit", "PERIODS",
	)
}

// The amount the customer should be billed for the commit. Not required.
//
// The properties CreditTypeID, Quantity, UnitPrice are required.
type V2ContractEditParamsAddRecurringCommitInvoiceAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	Quantity     float64 `json:"quantity" api:"required"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCommitInvoiceAmount) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCommitInvoiceAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional rounding configuration for prorated recurring commit amounts.
type V2ContractEditParamsAddRecurringCommitProrationRounding struct {
	Access  V2ContractEditParamsAddRecurringCommitProrationRoundingAccess  `json:"access,omitzero"`
	Invoice V2ContractEditParamsAddRecurringCommitProrationRoundingInvoice `json:"invoice,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCommitProrationRounding) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCommitProrationRounding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCommitProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DecimalPlaces, RoundingMethod are required.
type V2ContractEditParamsAddRecurringCommitProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCommitProrationRoundingAccess) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCommitProrationRoundingAccess
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCommitProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCommitProrationRoundingAccess](
		"rounding_method", "HALF_UP", "FLOOR", "CEILING",
	)
}

// The properties DecimalPlaces, RoundingMethod are required.
type V2ContractEditParamsAddRecurringCommitProrationRoundingInvoice struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCommitProrationRoundingInvoice) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCommitProrationRoundingInvoice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCommitProrationRoundingInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCommitProrationRoundingInvoice](
		"rounding_method", "HALF_UP", "FLOOR", "CEILING",
	)
}

// Attach a subscription to the recurring commit/credit.
//
// The properties ApplySeatIncreaseConfig, SubscriptionID are required.
type V2ContractEditParamsAddRecurringCommitSubscriptionConfig struct {
	ApplySeatIncreaseConfig V2ContractEditParamsAddRecurringCommitSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,omitzero" api:"required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID string `json:"subscription_id" api:"required"`
	// If set to POOLED, allocation added per seat is pooled across the account. If set
	// to INDIVIDUAL, each seat in the subscription will have its own allocation.
	//
	// Any of "POOLED", "INDIVIDUAL".
	Allocation string `json:"allocation,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCommitSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCommitSubscriptionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCommitSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCommitSubscriptionConfig](
		"allocation", "POOLED", "INDIVIDUAL",
	)
}

// The property IsProrated is required.
type V2ContractEditParamsAddRecurringCommitSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCommitSubscriptionConfigApplySeatIncreaseConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessAmount, CommitDuration, Priority, ProductID, StartingAt are
// required.
type V2ContractEditParamsAddRecurringCredit struct {
	// The amount of commit to grant.
	AccessAmount V2ContractEditParamsAddRecurringCreditAccessAmount `json:"access_amount,omitzero" api:"required"`
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
	CommitDuration V2ContractEditParamsAddRecurringCreditCommitDuration `json:"commit_duration,omitzero" api:"required"`
	// Will be passed down to the individual commits
	Priority  float64 `json:"priority" api:"required"`
	ProductID string  `json:"product_id" api:"required" format:"uuid"`
	// determines the start time for the first commit
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
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
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam `json:"hierarchy_configuration,omitzero"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	//
	// Any of "NONE", "FIRST", "LAST", "FIRST_AND_LAST".
	Proration string `json:"proration,omitzero"`
	// Optional rounding configuration for prorated recurring credit amounts.
	ProrationRounding V2ContractEditParamsAddRecurringCreditProrationRounding `json:"proration_rounding,omitzero"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates. - Daily recurring commits have a limit of one per contract, and
	// are unable to be created with seat-based subscriptions
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY", "DAILY".
	RecurrenceFrequency string `json:"recurrence_frequency,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig V2ContractEditParamsAddRecurringCreditSubscriptionConfig `json:"subscription_config,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCredit) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCredit](
		"proration", "NONE", "FIRST", "LAST", "FIRST_AND_LAST",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCredit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCredit](
		"recurrence_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY", "DAILY",
	)
}

// The amount of commit to grant.
//
// The properties CreditTypeID, UnitPrice are required.
type V2ContractEditParamsAddRecurringCreditAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	// This field is required unless a subscription is attached via
	// `subscription_config`.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCreditAccessAmount) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCreditAccessAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the length of the access schedule for each created commit/credit. The
// value represents the number of units. Unit defaults to "PERIODS", where the
// length of a period is determined by the recurrence_frequency.
//
// The property Value is required.
type V2ContractEditParamsAddRecurringCreditCommitDuration struct {
	Value float64 `json:"value" api:"required"`
	// Any of "PERIODS".
	Unit string `json:"unit,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCreditCommitDuration) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCreditCommitDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCreditCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCreditCommitDuration](
		"unit", "PERIODS",
	)
}

// Optional rounding configuration for prorated recurring credit amounts.
type V2ContractEditParamsAddRecurringCreditProrationRounding struct {
	Access V2ContractEditParamsAddRecurringCreditProrationRoundingAccess `json:"access,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCreditProrationRounding) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCreditProrationRounding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCreditProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DecimalPlaces, RoundingMethod are required.
type V2ContractEditParamsAddRecurringCreditProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCreditProrationRoundingAccess) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCreditProrationRoundingAccess
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCreditProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCreditProrationRoundingAccess](
		"rounding_method", "HALF_UP", "FLOOR", "CEILING",
	)
}

// Attach a subscription to the recurring commit/credit.
//
// The properties ApplySeatIncreaseConfig, SubscriptionID are required.
type V2ContractEditParamsAddRecurringCreditSubscriptionConfig struct {
	ApplySeatIncreaseConfig V2ContractEditParamsAddRecurringCreditSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,omitzero" api:"required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID string `json:"subscription_id" api:"required"`
	// If set to POOLED, allocation added per seat is pooled across the account. If set
	// to INDIVIDUAL, each seat in the subscription will have its own allocation.
	//
	// Any of "POOLED", "INDIVIDUAL".
	Allocation string `json:"allocation,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCreditSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCreditSubscriptionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCreditSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRecurringCreditSubscriptionConfig](
		"allocation", "POOLED", "INDIVIDUAL",
	)
}

// The property IsProrated is required.
type V2ContractEditParamsAddRecurringCreditSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCreditSubscriptionConfigApplySeatIncreaseConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ResellerType is required.
type V2ContractEditParamsAddResellerRoyalty struct {
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType string `json:"reseller_type,omitzero" api:"required"`
	// Use null to indicate that the existing end timestamp should be removed.
	EndingBefore          param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	Fraction              param.Opt[float64]   `json:"fraction,omitzero"`
	NetsuiteResellerID    param.Opt[string]    `json:"netsuite_reseller_id,omitzero"`
	ResellerContractValue param.Opt[float64]   `json:"reseller_contract_value,omitzero"`
	StartingAt            param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags []string                                         `json:"applicable_product_tags,omitzero"`
	AwsOptions            V2ContractEditParamsAddResellerRoyaltyAwsOptions `json:"aws_options,omitzero"`
	GcpOptions            V2ContractEditParamsAddResellerRoyaltyGcpOptions `json:"gcp_options,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddResellerRoyalty) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddResellerRoyalty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddResellerRoyalty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddResellerRoyalty](
		"reseller_type", "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE",
	)
}

type V2ContractEditParamsAddResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    param.Opt[string] `json:"aws_account_number,omitzero"`
	AwsOfferID          param.Opt[string] `json:"aws_offer_id,omitzero"`
	AwsPayerReferenceID param.Opt[string] `json:"aws_payer_reference_id,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddResellerRoyaltyAwsOptions) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddResellerRoyaltyAwsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddResellerRoyaltyAwsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParamsAddResellerRoyaltyGcpOptions struct {
	GcpAccountID param.Opt[string] `json:"gcp_account_id,omitzero"`
	GcpOfferID   param.Opt[string] `json:"gcp_offer_id,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddResellerRoyaltyGcpOptions) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddResellerRoyaltyGcpOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddResellerRoyaltyGcpOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update the revenue system configuration on the contract. Currently only supports
// adding a revenue system configuration to a contract that does not already have
// one.
//
// The properties RevenueSystemConfiguration, Schedule are required.
type V2ContractEditParamsAddRevenueSystemConfigurationUpdate struct {
	RevenueSystemConfiguration V2ContractEditParamsAddRevenueSystemConfigurationUpdateRevenueSystemConfiguration `json:"revenue_system_configuration,omitzero" api:"required"`
	Schedule                   V2ContractEditParamsAddRevenueSystemConfigurationUpdateSchedule                   `json:"schedule,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddRevenueSystemConfigurationUpdate) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRevenueSystemConfigurationUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRevenueSystemConfigurationUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParamsAddRevenueSystemConfigurationUpdateRevenueSystemConfiguration struct {
	RevenueSystemConfigurationID param.Opt[string] `json:"revenue_system_configuration_id,omitzero" format:"uuid"`
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod string `json:"delivery_method,omitzero"`
	// The revenue system provider type.
	//
	// Any of "netsuite".
	Provider string `json:"provider,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddRevenueSystemConfigurationUpdateRevenueSystemConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRevenueSystemConfigurationUpdateRevenueSystemConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRevenueSystemConfigurationUpdateRevenueSystemConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRevenueSystemConfigurationUpdateRevenueSystemConfiguration](
		"delivery_method", "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRevenueSystemConfigurationUpdateRevenueSystemConfiguration](
		"provider", "netsuite",
	)
}

// The property EffectiveAt is required.
type V2ContractEditParamsAddRevenueSystemConfigurationUpdateSchedule struct {
	// When the revenue system configuration update will take effect.
	//
	// Any of "START_OF_CURRENT_PERIOD", "START_OF_NEXT_PERIOD".
	EffectiveAt string `json:"effective_at,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddRevenueSystemConfigurationUpdateSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRevenueSystemConfigurationUpdateSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRevenueSystemConfigurationUpdateSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddRevenueSystemConfigurationUpdateSchedule](
		"effective_at", "START_OF_CURRENT_PERIOD", "START_OF_NEXT_PERIOD",
	)
}

// The properties ProductID, Schedule are required.
type V2ContractEditParamsAddScheduledCharge struct {
	ProductID string `json:"product_id" api:"required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule V2ContractEditParamsAddScheduledChargeSchedule `json:"schedule,omitzero" api:"required"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Opt[string] `json:"netsuite_sales_order_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddScheduledCharge) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddScheduledCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must provide either schedule_items or recurring_schedule.
type V2ContractEditParamsAddScheduledChargeSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Opt[bool] `json:"do_not_invoice,omitzero"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V2ContractEditParamsAddScheduledChargeScheduleRecurringSchedule `json:"recurring_schedule,omitzero"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V2ContractEditParamsAddScheduledChargeScheduleScheduleItem `json:"schedule_items,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddScheduledChargeSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddScheduledChargeSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddScheduledChargeSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
//
// The properties AmountDistribution, EndingBefore, Frequency, StartingAt are
// required.
type V2ContractEditParamsAddScheduledChargeScheduleRecurringSchedule struct {
	// Any of "DIVIDED", "DIVIDED_ROUNDED", "EACH".
	AmountDistribution string `json:"amount_distribution,omitzero" api:"required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency,omitzero" api:"required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
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

func (r V2ContractEditParamsAddScheduledChargeScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddScheduledChargeScheduleRecurringSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddScheduledChargeScheduleRecurringSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddScheduledChargeScheduleRecurringSchedule](
		"amount_distribution", "DIVIDED", "DIVIDED_ROUNDED", "EACH",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddScheduledChargeScheduleRecurringSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY",
	)
}

// The property Timestamp is required.
type V2ContractEditParamsAddScheduledChargeScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
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

func (r V2ContractEditParamsAddScheduledChargeScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddScheduledChargeScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddScheduledChargeScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Alias, ApplicableSpendSpecifiers, CreditTypeID, ResetFrequency
// are required.
type V2ContractEditParamsAddSpendTracker struct {
	// Human-readable identifier, unique per contract.
	Alias                     string                                                        `json:"alias" api:"required"`
	ApplicableSpendSpecifiers []V2ContractEditParamsAddSpendTrackerApplicableSpendSpecifier `json:"applicable_spend_specifiers,omitzero" api:"required"`
	CreditTypeID              string                                                        `json:"credit_type_id" api:"required" format:"uuid"`
	// Any of "BILLING_PERIOD".
	ResetFrequency string `json:"reset_frequency,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddSpendTracker) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddSpendTracker
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddSpendTracker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddSpendTracker](
		"reset_frequency", "BILLING_PERIOD",
	)
}

// The properties Sources, SpendType are required.
type V2ContractEditParamsAddSpendTrackerApplicableSpendSpecifier struct {
	// Any of "THRESHOLD_RECHARGE", "MANUAL".
	Sources []string `json:"sources,omitzero" api:"required"`
	// Any of "COMMIT_PURCHASE".
	SpendType string `json:"spend_type,omitzero" api:"required"`
	// Filter by whether the spend was discounted. Defaults to ANY if omitted.
	//
	// Any of "ANY", "DISCOUNTED_ONLY", "UNDISCOUNTED_ONLY".
	Discounted string `json:"discounted,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddSpendTrackerApplicableSpendSpecifier) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddSpendTrackerApplicableSpendSpecifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddSpendTrackerApplicableSpendSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddSpendTrackerApplicableSpendSpecifier](
		"spend_type", "COMMIT_PURCHASE",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddSpendTrackerApplicableSpendSpecifier](
		"discounted", "ANY", "DISCOUNTED_ONLY", "UNDISCOUNTED_ONLY",
	)
}

// The properties CollectionSchedule, Proration, SubscriptionRate are required.
type V2ContractEditParamsAddSubscription struct {
	// Any of "ADVANCE", "ARREARS".
	CollectionSchedule string                                              `json:"collection_schedule,omitzero" api:"required"`
	Proration          V2ContractEditParamsAddSubscriptionProration        `json:"proration,omitzero" api:"required"`
	SubscriptionRate   V2ContractEditParamsAddSubscriptionSubscriptionRate `json:"subscription_rate,omitzero" api:"required"`
	Description        param.Opt[string]                                   `json:"description,omitzero"`
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
	TemporaryID        param.Opt[string]                                     `json:"temporary_id,omitzero"`
	BillingCycleConfig V2ContractEditParamsAddSubscriptionBillingCycleConfig `json:"billing_cycle_config,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// Determines how the subscription's quantity is controlled. Defaults to
	// QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified
	// directly on the subscription. `initial_quantity` must be provided with this
	// option. Compatible with recurring commits/credits that use POOLED allocation.
	// **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add
	// user_123) to increment and decrement a subscription quantity, rather than
	// directly providing the quantity. You must use a **SEAT_BASED** subscription to
	// use a linked recurring credit with an allocation per seat. `seat_config` must be
	// provided with this option.
	//
	// Any of "SEAT_BASED", "QUANTITY_ONLY".
	QuantityManagementMode string                                        `json:"quantity_management_mode,omitzero"`
	SeatConfig             V2ContractEditParamsAddSubscriptionSeatConfig `json:"seat_config,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddSubscription) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddSubscription](
		"collection_schedule", "ADVANCE", "ARREARS",
	)
	apijson.RegisterFieldValidator[V2ContractEditParamsAddSubscription](
		"quantity_management_mode", "SEAT_BASED", "QUANTITY_ONLY",
	)
}

type V2ContractEditParamsAddSubscriptionProration struct {
	// Indicates if the partial period will be prorated or charged a full amount.
	IsProrated param.Opt[bool] `json:"is_prorated,omitzero"`
	// Indicates how mid-period quantity adjustments are invoiced.
	// **BILL_IMMEDIATELY**: Only available when collection schedule is `ADVANCE`. The
	// quantity increase will be billed immediately on the scheduled date.
	// **BILL_ON_NEXT_COLLECTION_DATE**: The quantity increase will be billed for
	// in-arrears at the end of the period.
	//
	// Any of "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE".
	InvoiceBehavior string                                               `json:"invoice_behavior,omitzero"`
	Rounding        V2ContractEditParamsAddSubscriptionProrationRounding `json:"rounding,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddSubscriptionProration) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddSubscriptionProration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddSubscriptionProration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddSubscriptionProration](
		"invoice_behavior", "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE",
	)
}

// The properties DecimalPlaces, RoundingMethod are required.
type V2ContractEditParamsAddSubscriptionProrationRounding struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsAddSubscriptionProrationRounding) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddSubscriptionProrationRounding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddSubscriptionProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddSubscriptionProrationRounding](
		"rounding_method", "HALF_UP", "FLOOR", "CEILING",
	)
}

// The properties BillingFrequency, ProductID are required.
type V2ContractEditParamsAddSubscriptionSubscriptionRate struct {
	// Frequency to bill subscription with. Together with product_id, must match
	// existing rate on the rate card.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero" api:"required"`
	// Must be subscription type product
	ProductID string `json:"product_id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditParamsAddSubscriptionSubscriptionRate) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddSubscriptionSubscriptionRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddSubscriptionSubscriptionRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddSubscriptionSubscriptionRate](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

type V2ContractEditParamsAddSubscriptionBillingCycleConfig struct {
	// The date to anchor the billing cycle to. If omitted, defaults to the contract's
	// usage invoice billing cycle anchor date.
	AnchorDate param.Opt[time.Time] `json:"anchor_date,omitzero" format:"date-time"`
	// Controls whether this subscription consolidates onto usage invoices or gets its
	// own scheduled invoice. Defaults to ON_USAGE_INVOICE if omitted.
	//
	// Any of "ON_SCHEDULED_INVOICE", "ON_USAGE_INVOICE".
	InvoicePlacement string `json:"invoice_placement,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddSubscriptionBillingCycleConfig) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddSubscriptionBillingCycleConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddSubscriptionBillingCycleConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddSubscriptionBillingCycleConfig](
		"invoice_placement", "ON_SCHEDULED_INVOICE", "ON_USAGE_INVOICE",
	)
}

// The properties InitialSeatIDs, SeatGroupKey are required.
type V2ContractEditParamsAddSubscriptionSeatConfig struct {
	// The initial assigned seats on this subscription.
	InitialSeatIDs []string `json:"initial_seat_ids,omitzero" api:"required"`
	// The property name, sent on usage events, that identifies the seat ID associated
	// with the usage event. For example, the property name might be seat_id or
	// user_id. The property must be set as a group key on billable metrics and a
	// presentation/pricing group key on contract products. This allows linked
	// recurring credits with an allocation per seat to be consumed by only one seat's
	// usage.
	SeatGroupKey string `json:"seat_group_key" api:"required"`
	// The initial amount of unassigned seats on this subscription.
	InitialUnassignedSeats param.Opt[float64] `json:"initial_unassigned_seats,omitzero"`
	paramObj
}

func (r V2ContractEditParamsAddSubscriptionSeatConfig) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddSubscriptionSeatConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddSubscriptionSeatConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsArchiveCommit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditParamsArchiveCommit) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsArchiveCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsArchiveCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsArchiveCredit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditParamsArchiveCredit) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsArchiveCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsArchiveCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsArchiveScheduledCharge struct {
	ID string `json:"id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditParamsArchiveScheduledCharge) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsArchiveScheduledCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsArchiveScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsRemoveOverride struct {
	ID string `json:"id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditParamsRemoveOverride) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsRemoveOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsRemoveOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CommitID is required.
type V2ContractEditParamsUpdateCommit struct {
	CommitID             string             `json:"commit_id" api:"required" format:"uuid"`
	NetsuiteSalesOrderID param.Opt[string]  `json:"netsuite_sales_order_id,omitzero"`
	Priority             param.Opt[float64] `json:"priority,omitzero"`
	RolloverFraction     param.Opt[float64] `json:"rollover_fraction,omitzero"`
	Description          param.Opt[string]  `json:"description,omitzero"`
	Name                 param.Opt[string]  `json:"name,omitzero"`
	ProductID            param.Opt[string]  `json:"product_id,omitzero" format:"uuid"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string                                       `json:"applicable_product_tags,omitzero"`
	AccessSchedule        V2ContractEditParamsUpdateCommitAccessSchedule `json:"access_schedule,omitzero"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam        `json:"hierarchy_configuration,omitzero"`
	InvoiceSchedule        V2ContractEditParamsUpdateCommitInvoiceSchedule `json:"invoice_schedule,omitzero"`
	// If provided, updates the commit to use the specified rate type for current and
	// future invoices. Previously finalized invoices will need to be voided and
	// regenerated to reflect the rate type change.
	//
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType string `json:"rate_type,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateCommit) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsUpdateCommit](
		"rate_type", "LIST_RATE", "COMMIT_RATE",
	)
}

type V2ContractEditParamsUpdateCommitAccessSchedule struct {
	AddScheduleItems    []V2ContractEditParamsUpdateCommitAccessScheduleAddScheduleItem    `json:"add_schedule_items,omitzero"`
	RemoveScheduleItems []V2ContractEditParamsUpdateCommitAccessScheduleRemoveScheduleItem `json:"remove_schedule_items,omitzero"`
	UpdateScheduleItems []V2ContractEditParamsUpdateCommitAccessScheduleUpdateScheduleItem `json:"update_schedule_items,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateCommitAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCommitAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCommitAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, EndingBefore, StartingAt are required.
type V2ContractEditParamsUpdateCommitAccessScheduleAddScheduleItem struct {
	Amount       float64   `json:"amount" api:"required"`
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	paramObj
}

func (r V2ContractEditParamsUpdateCommitAccessScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCommitAccessScheduleAddScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCommitAccessScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsUpdateCommitAccessScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditParamsUpdateCommitAccessScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCommitAccessScheduleRemoveScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCommitAccessScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsUpdateCommitAccessScheduleUpdateScheduleItem struct {
	ID           string               `json:"id" api:"required" format:"uuid"`
	Amount       param.Opt[float64]   `json:"amount,omitzero"`
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	StartingAt   param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V2ContractEditParamsUpdateCommitAccessScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCommitAccessScheduleUpdateScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCommitAccessScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParamsUpdateCommitInvoiceSchedule struct {
	AddScheduleItems    []V2ContractEditParamsUpdateCommitInvoiceScheduleAddScheduleItem    `json:"add_schedule_items,omitzero"`
	RemoveScheduleItems []V2ContractEditParamsUpdateCommitInvoiceScheduleRemoveScheduleItem `json:"remove_schedule_items,omitzero"`
	UpdateScheduleItems []V2ContractEditParamsUpdateCommitInvoiceScheduleUpdateScheduleItem `json:"update_schedule_items,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateCommitInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCommitInvoiceSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Timestamp is required.
type V2ContractEditParamsUpdateCommitInvoiceScheduleAddScheduleItem struct {
	Timestamp time.Time          `json:"timestamp" api:"required" format:"date-time"`
	Amount    param.Opt[float64] `json:"amount,omitzero"`
	Quantity  param.Opt[float64] `json:"quantity,omitzero"`
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateCommitInvoiceScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCommitInvoiceScheduleAddScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCommitInvoiceScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsUpdateCommitInvoiceScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditParamsUpdateCommitInvoiceScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCommitInvoiceScheduleRemoveScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCommitInvoiceScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsUpdateCommitInvoiceScheduleUpdateScheduleItem struct {
	ID        string               `json:"id" api:"required" format:"uuid"`
	Amount    param.Opt[float64]   `json:"amount,omitzero"`
	Quantity  param.Opt[float64]   `json:"quantity,omitzero"`
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	UnitPrice param.Opt[float64]   `json:"unit_price,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateCommitInvoiceScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCommitInvoiceScheduleUpdateScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCommitInvoiceScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CreditID is required.
type V2ContractEditParamsUpdateCredit struct {
	CreditID             string             `json:"credit_id" api:"required" format:"uuid"`
	NetsuiteSalesOrderID param.Opt[string]  `json:"netsuite_sales_order_id,omitzero"`
	Priority             param.Opt[float64] `json:"priority,omitzero"`
	RolloverFraction     param.Opt[float64] `json:"rollover_fraction,omitzero"`
	Description          param.Opt[string]  `json:"description,omitzero"`
	Name                 param.Opt[string]  `json:"name,omitzero"`
	ProductID            param.Opt[string]  `json:"product_id,omitzero" format:"uuid"`
	// Which products the credit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the credit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the credit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the credit applies to
	// all products.
	ApplicableProductTags []string                                       `json:"applicable_product_tags,omitzero"`
	AccessSchedule        V2ContractEditParamsUpdateCreditAccessSchedule `json:"access_schedule,omitzero"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam `json:"hierarchy_configuration,omitzero"`
	// If provided, updates the credit to use the specified rate type for current and
	// future invoices. Previously finalized invoices will need to be voided and
	// regenerated to reflect the rate type change.
	//
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType string `json:"rate_type,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateCredit) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsUpdateCredit](
		"rate_type", "LIST_RATE", "COMMIT_RATE",
	)
}

type V2ContractEditParamsUpdateCreditAccessSchedule struct {
	AddScheduleItems    []V2ContractEditParamsUpdateCreditAccessScheduleAddScheduleItem    `json:"add_schedule_items,omitzero"`
	RemoveScheduleItems []V2ContractEditParamsUpdateCreditAccessScheduleRemoveScheduleItem `json:"remove_schedule_items,omitzero"`
	UpdateScheduleItems []V2ContractEditParamsUpdateCreditAccessScheduleUpdateScheduleItem `json:"update_schedule_items,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateCreditAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCreditAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCreditAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, EndingBefore, StartingAt are required.
type V2ContractEditParamsUpdateCreditAccessScheduleAddScheduleItem struct {
	Amount       float64   `json:"amount" api:"required"`
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	paramObj
}

func (r V2ContractEditParamsUpdateCreditAccessScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCreditAccessScheduleAddScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCreditAccessScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsUpdateCreditAccessScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditParamsUpdateCreditAccessScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCreditAccessScheduleRemoveScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCreditAccessScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsUpdateCreditAccessScheduleUpdateScheduleItem struct {
	ID           string               `json:"id" api:"required" format:"uuid"`
	Amount       param.Opt[float64]   `json:"amount,omitzero"`
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	StartingAt   param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V2ContractEditParamsUpdateCreditAccessScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateCreditAccessScheduleUpdateScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateCreditAccessScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfiguration struct {
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID param.Opt[string] `json:"custom_credit_type_id,omitzero" format:"uuid"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled param.Opt[bool] `json:"is_enabled,omitzero"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount param.Opt[float64] `json:"recharge_to_amount,omitzero"`
	// Specify the threshold amount for the contract. Each time the contract's balance
	// lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount       param.Opt[float64]                                                                  `json:"threshold_amount,omitzero"`
	DiscountConfiguration V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration `json:"discount_configuration,omitzero"`
	// Determines which balances are excluded from remaining balance calculation for
	// threshold billing.
	ThresholdBalanceSpecifiers []V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier `json:"threshold_balance_specifiers,omitzero"`
	Commit                     V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommit                      `json:"commit,omitzero"`
	PaymentGateConfig          shared.PaymentGateConfigV2Param                                                           `json:"payment_gate_config,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdatePrepaidBalanceThresholdConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdatePrepaidBalanceThresholdConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommit struct {
	// Which products the threshold commit applies to. If both applicable_product_ids
	// and applicable_product_tags are not provided, the commit applies to all
	// products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// The length of time the created commit will be valid, starting from the end of
	// the invoice's service period. Set to null to clear a previously configured
	// duration.
	Duration V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommitDuration `json:"duration,omitzero"`
	// Whether the created commits will be charged at commit rate or list rate. Set to
	// null to clear a previously configured rate type.
	RateType string `json:"rate_type,omitzero"`
	// Fraction of the created commit's unused balance that will roll over. Must be
	// between 0 and 1. Set to null to clear a previously configured rollover fraction.
	RolloverFraction param.Opt[float64] `json:"rollover_fraction,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	shared.UpdateBaseThresholdCommitParam
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommit) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommit
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// The length of time the created commit will be valid, starting from the end of
// the invoice's service period. Set to null to clear a previously configured
// duration.
//
// The properties Unit, Value are required.
type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommitDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommitDuration) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommitDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommitDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration struct {
	// The fraction of the original amount that the customer pays after applying the
	// discount. Set to null to remove the discount fraction. For example, 0.85 means
	// the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction param.Opt[float64] `json:"payment_fraction,omitzero"`
	// Update the discount cap. Set to null to remove an existing cap.
	Cap V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap `json:"cap,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationDiscountConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update the discount cap. Set to null to remove an existing cap.
//
// The properties Amount, SpendTrackerAlias are required.
type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount float64 `json:"amount" api:"required"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias string `json:"spend_tracker_alias" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Exclude is required.
type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier struct {
	Exclude []V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude `json:"exclude,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CustomFieldFilters is required.
type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude struct {
	CustomFieldFilters []V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter `json:"custom_field_filters,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExclude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Entity, Key, Value are required.
type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter struct {
	// Any of "Commit", "ContractCredit", "ContractCreditOrCommit".
	Entity string `json:"entity,omitzero" api:"required"`
	Key    string `json:"key" api:"required"`
	Value  string `json:"value" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilter](
		"entity", "Commit", "ContractCredit", "ContractCreditOrCommit",
	)
}

// The property RecurringCommitID is required.
type V2ContractEditParamsUpdateRecurringCommit struct {
	RecurringCommitID string               `json:"recurring_commit_id" api:"required" format:"uuid"`
	EndingBefore      param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// If provided, updates the rounding config on the recurring commit. Set to null to
	// clear rounding. Omit to leave unchanged.
	ProrationRounding V2ContractEditParamsUpdateRecurringCommitProrationRounding `json:"proration_rounding,omitzero"`
	AccessAmount      V2ContractEditParamsUpdateRecurringCommitAccessAmount      `json:"access_amount,omitzero"`
	InvoiceAmount     V2ContractEditParamsUpdateRecurringCommitInvoiceAmount     `json:"invoice_amount,omitzero"`
	// If provided, updates the recurring commit to use the specified rate type when
	// generating future commits.
	//
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType string `json:"rate_type,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateRecurringCommit) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateRecurringCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsUpdateRecurringCommit](
		"rate_type", "LIST_RATE", "COMMIT_RATE",
	)
}

type V2ContractEditParamsUpdateRecurringCommitAccessAmount struct {
	Quantity  param.Opt[float64] `json:"quantity,omitzero"`
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateRecurringCommitAccessAmount) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateRecurringCommitAccessAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParamsUpdateRecurringCommitInvoiceAmount struct {
	Quantity  param.Opt[float64] `json:"quantity,omitzero"`
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateRecurringCommitInvoiceAmount) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateRecurringCommitInvoiceAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// If provided, updates the rounding config on the recurring commit. Set to null to
// clear rounding. Omit to leave unchanged.
type V2ContractEditParamsUpdateRecurringCommitProrationRounding struct {
	Access  V2ContractEditParamsUpdateRecurringCommitProrationRoundingAccess  `json:"access,omitzero"`
	Invoice V2ContractEditParamsUpdateRecurringCommitProrationRoundingInvoice `json:"invoice,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateRecurringCommitProrationRounding) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateRecurringCommitProrationRounding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateRecurringCommitProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DecimalPlaces, RoundingMethod are required.
type V2ContractEditParamsUpdateRecurringCommitProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdateRecurringCommitProrationRoundingAccess) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateRecurringCommitProrationRoundingAccess
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateRecurringCommitProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsUpdateRecurringCommitProrationRoundingAccess](
		"rounding_method", "HALF_UP", "FLOOR", "CEILING",
	)
}

// The properties DecimalPlaces, RoundingMethod are required.
type V2ContractEditParamsUpdateRecurringCommitProrationRoundingInvoice struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdateRecurringCommitProrationRoundingInvoice) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateRecurringCommitProrationRoundingInvoice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateRecurringCommitProrationRoundingInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsUpdateRecurringCommitProrationRoundingInvoice](
		"rounding_method", "HALF_UP", "FLOOR", "CEILING",
	)
}

// The property RecurringCreditID is required.
type V2ContractEditParamsUpdateRecurringCredit struct {
	RecurringCreditID string               `json:"recurring_credit_id" api:"required" format:"uuid"`
	EndingBefore      param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// If provided, updates the rounding config on the recurring credit. Set to null to
	// clear rounding. Omit to leave unchanged.
	ProrationRounding V2ContractEditParamsUpdateRecurringCreditProrationRounding `json:"proration_rounding,omitzero"`
	AccessAmount      V2ContractEditParamsUpdateRecurringCreditAccessAmount      `json:"access_amount,omitzero"`
	// If provided, updates the recurring credit to use the specified rate type when
	// generating future credits.
	//
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType string `json:"rate_type,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateRecurringCredit) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateRecurringCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsUpdateRecurringCredit](
		"rate_type", "LIST_RATE", "COMMIT_RATE",
	)
}

type V2ContractEditParamsUpdateRecurringCreditAccessAmount struct {
	Quantity  param.Opt[float64] `json:"quantity,omitzero"`
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateRecurringCreditAccessAmount) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateRecurringCreditAccessAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// If provided, updates the rounding config on the recurring credit. Set to null to
// clear rounding. Omit to leave unchanged.
type V2ContractEditParamsUpdateRecurringCreditProrationRounding struct {
	Access V2ContractEditParamsUpdateRecurringCreditProrationRoundingAccess `json:"access,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateRecurringCreditProrationRounding) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateRecurringCreditProrationRounding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateRecurringCreditProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DecimalPlaces, RoundingMethod are required.
type V2ContractEditParamsUpdateRecurringCreditProrationRoundingAccess struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdateRecurringCreditProrationRoundingAccess) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateRecurringCreditProrationRoundingAccess
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateRecurringCreditProrationRoundingAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsUpdateRecurringCreditProrationRoundingAccess](
		"rounding_method", "HALF_UP", "FLOOR", "CEILING",
	)
}

// The property ScheduledChargeID is required.
type V2ContractEditParamsUpdateScheduledCharge struct {
	ScheduledChargeID    string                                                   `json:"scheduled_charge_id" api:"required" format:"uuid"`
	NetsuiteSalesOrderID param.Opt[string]                                        `json:"netsuite_sales_order_id,omitzero"`
	InvoiceSchedule      V2ContractEditParamsUpdateScheduledChargeInvoiceSchedule `json:"invoice_schedule,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateScheduledCharge) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateScheduledCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParamsUpdateScheduledChargeInvoiceSchedule struct {
	AddScheduleItems    []V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleAddScheduleItem    `json:"add_schedule_items,omitzero"`
	RemoveScheduleItems []V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem `json:"remove_schedule_items,omitzero"`
	UpdateScheduleItems []V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem `json:"update_schedule_items,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateScheduledChargeInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateScheduledChargeInvoiceSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateScheduledChargeInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Timestamp is required.
type V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleAddScheduleItem struct {
	Timestamp time.Time          `json:"timestamp" api:"required" format:"date-time"`
	Amount    param.Opt[float64] `json:"amount,omitzero"`
	Quantity  param.Opt[float64] `json:"quantity,omitzero"`
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleAddScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem struct {
	ID        string               `json:"id" api:"required" format:"uuid"`
	Amount    param.Opt[float64]   `json:"amount,omitzero"`
	Quantity  param.Opt[float64]   `json:"quantity,omitzero"`
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	UnitPrice param.Opt[float64]   `json:"unit_price,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateScheduledChargeInvoiceScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParamsUpdateSpendThresholdConfiguration struct {
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled param.Opt[bool] `json:"is_enabled,omitzero"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount       param.Opt[float64]                                                         `json:"threshold_amount,omitzero"`
	DiscountConfiguration V2ContractEditParamsUpdateSpendThresholdConfigurationDiscountConfiguration `json:"discount_configuration,omitzero"`
	Commit                shared.UpdateBaseThresholdCommitParam                                      `json:"commit,omitzero"`
	PaymentGateConfig     shared.PaymentGateConfigV2Param                                            `json:"payment_gate_config,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateSpendThresholdConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSpendThresholdConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSpendThresholdConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParamsUpdateSpendThresholdConfigurationDiscountConfiguration struct {
	// The fraction of the original amount that the customer pays after applying the
	// discount. Set to null to remove the discount fraction. For example, 0.85 means
	// the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction param.Opt[float64] `json:"payment_fraction,omitzero"`
	// Update the discount cap. Set to null to remove an existing cap.
	Cap V2ContractEditParamsUpdateSpendThresholdConfigurationDiscountConfigurationCap `json:"cap,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateSpendThresholdConfigurationDiscountConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSpendThresholdConfigurationDiscountConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSpendThresholdConfigurationDiscountConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update the discount cap. Set to null to remove an existing cap.
//
// The properties Amount, SpendTrackerAlias are required.
type V2ContractEditParamsUpdateSpendThresholdConfigurationDiscountConfigurationCap struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount float64 `json:"amount" api:"required"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias string `json:"spend_tracker_alias" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdateSpendThresholdConfigurationDiscountConfigurationCap) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSpendThresholdConfigurationDiscountConfigurationCap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSpendThresholdConfigurationDiscountConfigurationCap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SubscriptionID is required.
type V2ContractEditParamsUpdateSubscription struct {
	SubscriptionID    string                                                  `json:"subscription_id" api:"required" format:"uuid"`
	EndingBefore      param.Opt[time.Time]                                    `json:"ending_before,omitzero" format:"date-time"`
	Name              param.Opt[string]                                       `json:"name,omitzero"`
	ProrationRounding V2ContractEditParamsUpdateSubscriptionProrationRounding `json:"proration_rounding,omitzero"`
	// Update the subscription's quantity management mode from QUANTITY_ONLY to
	// SEAT_BASED with the provided seat_group_key.
	QuantityManagementModeUpdate V2ContractEditParamsUpdateSubscriptionQuantityManagementModeUpdate `json:"quantity_management_mode_update,omitzero"`
	// Quantity changes are applied on the effective date based on the order which they
	// are sent. For example, if I scheduled the quantity to be 12 on May 21 and then
	// scheduled a quantity delta change of -1, the result from that day would be 11.
	QuantityUpdates []V2ContractEditParamsUpdateSubscriptionQuantityUpdate `json:"quantity_updates,omitzero"`
	SeatUpdates     V2ContractEditParamsUpdateSubscriptionSeatUpdates      `json:"seat_updates,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateSubscription) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DecimalPlaces, RoundingMethod are required.
type V2ContractEditParamsUpdateSubscriptionProrationRounding struct {
	// Number of decimal places to round to. Applied directly to the stored monetary
	// representation. Negative values round to powers of 10 (e.g., -2 rounds to
	// nearest 100 in the stored unit).
	DecimalPlaces float64 `json:"decimal_places" api:"required"`
	// Any of "HALF_UP", "FLOOR", "CEILING".
	RoundingMethod string `json:"rounding_method,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdateSubscriptionProrationRounding) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSubscriptionProrationRounding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSubscriptionProrationRounding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsUpdateSubscriptionProrationRounding](
		"rounding_method", "HALF_UP", "FLOOR", "CEILING",
	)
}

// Update the subscription's quantity management mode from QUANTITY_ONLY to
// SEAT_BASED with the provided seat_group_key.
//
// The properties QuantityManagementMode, SeatConfig are required.
type V2ContractEditParamsUpdateSubscriptionQuantityManagementModeUpdate struct {
	// Any of "SEAT_BASED".
	QuantityManagementMode string                                                                       `json:"quantity_management_mode,omitzero" api:"required"`
	SeatConfig             V2ContractEditParamsUpdateSubscriptionQuantityManagementModeUpdateSeatConfig `json:"seat_config,omitzero" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdateSubscriptionQuantityManagementModeUpdate) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSubscriptionQuantityManagementModeUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSubscriptionQuantityManagementModeUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsUpdateSubscriptionQuantityManagementModeUpdate](
		"quantity_management_mode", "SEAT_BASED",
	)
}

// The property SeatGroupKey is required.
type V2ContractEditParamsUpdateSubscriptionQuantityManagementModeUpdateSeatConfig struct {
	SeatGroupKey string `json:"seat_group_key" api:"required"`
	paramObj
}

func (r V2ContractEditParamsUpdateSubscriptionQuantityManagementModeUpdateSeatConfig) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSubscriptionQuantityManagementModeUpdateSeatConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSubscriptionQuantityManagementModeUpdateSeatConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property StartingAt is required.
type V2ContractEditParamsUpdateSubscriptionQuantityUpdate struct {
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	// The new quantity for the subscription. Must be provided if quantity_delta is not
	// provided. Must be non-negative.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// The delta to add to the subscription's quantity. Must be provided if quantity is
	// not provided. Can't be zero. It also can't result in a negative quantity on the
	// subscription.
	QuantityDelta param.Opt[float64] `json:"quantity_delta,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateSubscriptionQuantityUpdate) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSubscriptionQuantityUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSubscriptionQuantityUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditParamsUpdateSubscriptionSeatUpdates struct {
	// Adds seat IDs to the subscription. If there are unassigned seats, the new seat
	// IDs will fill these unassigned seats and not increase the total subscription
	// quantity. Otherwise, if there are more new seat IDs than unassigned seats, the
	// total subscription quantity will increase.
	AddSeatIDs []V2ContractEditParamsUpdateSubscriptionSeatUpdatesAddSeatID `json:"add_seat_ids,omitzero"`
	// Adds unassigned seats to the subscription. This will increase the total
	// subscription quantity.
	AddUnassignedSeats []V2ContractEditParamsUpdateSubscriptionSeatUpdatesAddUnassignedSeat `json:"add_unassigned_seats,omitzero"`
	// Removes seat IDs from the subscription, if possible. If a seat ID is removed,
	// the total subscription quantity will decrease. Otherwise, if the seat ID is not
	// found on the subscription, this is a no-op.
	RemoveSeatIDs []V2ContractEditParamsUpdateSubscriptionSeatUpdatesRemoveSeatID `json:"remove_seat_ids,omitzero"`
	// Removes unassigned seats from the subscription. This will decrease the total
	// subscription quantity if there are are unassigned seats.
	RemoveUnassignedSeats []V2ContractEditParamsUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat `json:"remove_unassigned_seats,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateSubscriptionSeatUpdates) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSubscriptionSeatUpdates
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSubscriptionSeatUpdates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SeatIDs, StartingAt are required.
type V2ContractEditParamsUpdateSubscriptionSeatUpdatesAddSeatID struct {
	SeatIDs []string `json:"seat_ids,omitzero" api:"required"`
	// Assigned seats will be added/removed starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	paramObj
}

func (r V2ContractEditParamsUpdateSubscriptionSeatUpdatesAddSeatID) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSubscriptionSeatUpdatesAddSeatID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSubscriptionSeatUpdatesAddSeatID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Quantity, StartingAt are required.
type V2ContractEditParamsUpdateSubscriptionSeatUpdatesAddUnassignedSeat struct {
	// The number of unassigned seats on the subscription will increase/decrease by
	// this delta. Must be greater than 0.
	Quantity float64 `json:"quantity" api:"required"`
	// Unassigned seats will be updated starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	paramObj
}

func (r V2ContractEditParamsUpdateSubscriptionSeatUpdatesAddUnassignedSeat) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSubscriptionSeatUpdatesAddUnassignedSeat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSubscriptionSeatUpdatesAddUnassignedSeat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SeatIDs, StartingAt are required.
type V2ContractEditParamsUpdateSubscriptionSeatUpdatesRemoveSeatID struct {
	SeatIDs []string `json:"seat_ids,omitzero" api:"required"`
	// Assigned seats will be added/removed starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	paramObj
}

func (r V2ContractEditParamsUpdateSubscriptionSeatUpdatesRemoveSeatID) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSubscriptionSeatUpdatesRemoveSeatID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSubscriptionSeatUpdatesRemoveSeatID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Quantity, StartingAt are required.
type V2ContractEditParamsUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat struct {
	// The number of unassigned seats on the subscription will increase/decrease by
	// this delta. Must be greater than 0.
	Quantity float64 `json:"quantity" api:"required"`
	// Unassigned seats will be updated starting at this date.
	StartingAt time.Time `json:"starting_at" api:"required" format:"date-time"`
	paramObj
}

func (r V2ContractEditParamsUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSubscriptionSeatUpdatesRemoveUnassignedSeat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditCommitParams struct {
	// ID of the commit to edit
	CommitID string `json:"commit_id" api:"required" format:"uuid"`
	// ID of the customer whose commit is being edited
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Updated description for the commit
	Description param.Opt[string] `json:"description,omitzero"`
	// ID of contract to use for invoicing
	InvoiceContractID param.Opt[string] `json:"invoice_contract_id,omitzero" format:"uuid"`
	// Updated name for the commit
	Name      param.Opt[string] `json:"name,omitzero"`
	ProductID param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	// Which contracts the customer-level commit applies to. If set to null, the commit
	// applies to all of the customer's contracts. This field cannot be edited for
	// POSTPAID commits or contract-level commits.
	ApplicableContractIDs []string `json:"applicable_contract_ids,omitzero" format:"uuid"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers     []shared.CommitSpecifierInputParam       `json:"specifiers,omitzero"`
	AccessSchedule V2ContractEditCommitParamsAccessSchedule `json:"access_schedule,omitzero"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam  `json:"hierarchy_configuration,omitzero"`
	InvoiceSchedule        V2ContractEditCommitParamsInvoiceSchedule `json:"invoice_schedule,omitzero"`
	// If provided, updates the commit to use the specified rate type for current and
	// future invoices. Previously finalized invoices will need to be voided and
	// regenerated to reflect the rate type change.
	//
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType V2ContractEditCommitParamsRateType `json:"rate_type,omitzero"`
	paramObj
}

func (r V2ContractEditCommitParams) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCommitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCommitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditCommitParamsAccessSchedule struct {
	AddScheduleItems    []V2ContractEditCommitParamsAccessScheduleAddScheduleItem    `json:"add_schedule_items,omitzero"`
	RemoveScheduleItems []V2ContractEditCommitParamsAccessScheduleRemoveScheduleItem `json:"remove_schedule_items,omitzero"`
	UpdateScheduleItems []V2ContractEditCommitParamsAccessScheduleUpdateScheduleItem `json:"update_schedule_items,omitzero"`
	paramObj
}

func (r V2ContractEditCommitParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCommitParamsAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCommitParamsAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, EndingBefore, StartingAt are required.
type V2ContractEditCommitParamsAccessScheduleAddScheduleItem struct {
	Amount       float64   `json:"amount" api:"required"`
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	paramObj
}

func (r V2ContractEditCommitParamsAccessScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCommitParamsAccessScheduleAddScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCommitParamsAccessScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditCommitParamsAccessScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditCommitParamsAccessScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCommitParamsAccessScheduleRemoveScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCommitParamsAccessScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditCommitParamsAccessScheduleUpdateScheduleItem struct {
	ID           string               `json:"id" api:"required" format:"uuid"`
	Amount       param.Opt[float64]   `json:"amount,omitzero"`
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	StartingAt   param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V2ContractEditCommitParamsAccessScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCommitParamsAccessScheduleUpdateScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCommitParamsAccessScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditCommitParamsInvoiceSchedule struct {
	AddScheduleItems    []V2ContractEditCommitParamsInvoiceScheduleAddScheduleItem    `json:"add_schedule_items,omitzero"`
	RemoveScheduleItems []V2ContractEditCommitParamsInvoiceScheduleRemoveScheduleItem `json:"remove_schedule_items,omitzero"`
	UpdateScheduleItems []V2ContractEditCommitParamsInvoiceScheduleUpdateScheduleItem `json:"update_schedule_items,omitzero"`
	paramObj
}

func (r V2ContractEditCommitParamsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCommitParamsInvoiceSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCommitParamsInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Timestamp is required.
type V2ContractEditCommitParamsInvoiceScheduleAddScheduleItem struct {
	Timestamp time.Time          `json:"timestamp" api:"required" format:"date-time"`
	Amount    param.Opt[float64] `json:"amount,omitzero"`
	Quantity  param.Opt[float64] `json:"quantity,omitzero"`
	UnitPrice param.Opt[float64] `json:"unit_price,omitzero"`
	paramObj
}

func (r V2ContractEditCommitParamsInvoiceScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCommitParamsInvoiceScheduleAddScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCommitParamsInvoiceScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditCommitParamsInvoiceScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditCommitParamsInvoiceScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCommitParamsInvoiceScheduleRemoveScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCommitParamsInvoiceScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditCommitParamsInvoiceScheduleUpdateScheduleItem struct {
	ID        string               `json:"id" api:"required" format:"uuid"`
	Amount    param.Opt[float64]   `json:"amount,omitzero"`
	Quantity  param.Opt[float64]   `json:"quantity,omitzero"`
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	UnitPrice param.Opt[float64]   `json:"unit_price,omitzero"`
	paramObj
}

func (r V2ContractEditCommitParamsInvoiceScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCommitParamsInvoiceScheduleUpdateScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCommitParamsInvoiceScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// If provided, updates the commit to use the specified rate type for current and
// future invoices. Previously finalized invoices will need to be voided and
// regenerated to reflect the rate type change.
type V2ContractEditCommitParamsRateType string

const (
	V2ContractEditCommitParamsRateTypeListRate   V2ContractEditCommitParamsRateType = "LIST_RATE"
	V2ContractEditCommitParamsRateTypeCommitRate V2ContractEditCommitParamsRateType = "COMMIT_RATE"
)

type V2ContractEditCreditParams struct {
	// ID of the credit to edit
	CreditID string `json:"credit_id" api:"required" format:"uuid"`
	// ID of the customer whose credit is being edited
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Updated description for the credit
	Description param.Opt[string] `json:"description,omitzero"`
	// Updated name for the credit
	Name      param.Opt[string] `json:"name,omitzero"`
	ProductID param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	// Which contracts the customer-level credit applies to. If set to null, the credit
	// applies to all of the customer's contracts. This field cannot be set on a
	// contract-level credit.
	ApplicableContractIDs []string `json:"applicable_contract_ids,omitzero" format:"uuid"`
	// Which products the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers     []shared.CommitSpecifierInputParam       `json:"specifiers,omitzero"`
	AccessSchedule V2ContractEditCreditParamsAccessSchedule `json:"access_schedule,omitzero"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam `json:"hierarchy_configuration,omitzero"`
	// If provided, updates the credit to use the specified rate type for current and
	// future invoices. Previously finalized invoices will need to be voided and
	// regenerated to reflect the rate type change.
	//
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType V2ContractEditCreditParamsRateType `json:"rate_type,omitzero"`
	paramObj
}

func (r V2ContractEditCreditParams) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCreditParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCreditParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractEditCreditParamsAccessSchedule struct {
	AddScheduleItems    []V2ContractEditCreditParamsAccessScheduleAddScheduleItem    `json:"add_schedule_items,omitzero"`
	RemoveScheduleItems []V2ContractEditCreditParamsAccessScheduleRemoveScheduleItem `json:"remove_schedule_items,omitzero"`
	UpdateScheduleItems []V2ContractEditCreditParamsAccessScheduleUpdateScheduleItem `json:"update_schedule_items,omitzero"`
	paramObj
}

func (r V2ContractEditCreditParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCreditParamsAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCreditParamsAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, EndingBefore, StartingAt are required.
type V2ContractEditCreditParamsAccessScheduleAddScheduleItem struct {
	Amount       float64   `json:"amount" api:"required"`
	EndingBefore time.Time `json:"ending_before" api:"required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	paramObj
}

func (r V2ContractEditCreditParamsAccessScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCreditParamsAccessScheduleAddScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCreditParamsAccessScheduleAddScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditCreditParamsAccessScheduleRemoveScheduleItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractEditCreditParamsAccessScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCreditParamsAccessScheduleRemoveScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCreditParamsAccessScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type V2ContractEditCreditParamsAccessScheduleUpdateScheduleItem struct {
	ID           string               `json:"id" api:"required" format:"uuid"`
	Amount       param.Opt[float64]   `json:"amount,omitzero"`
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	StartingAt   param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V2ContractEditCreditParamsAccessScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditCreditParamsAccessScheduleUpdateScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditCreditParamsAccessScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// If provided, updates the credit to use the specified rate type for current and
// future invoices. Previously finalized invoices will need to be voided and
// regenerated to reflect the rate type change.
type V2ContractEditCreditParamsRateType string

const (
	V2ContractEditCreditParamsRateTypeListRate   V2ContractEditCreditParamsRateType = "LIST_RATE"
	V2ContractEditCreditParamsRateTypeCommitRate V2ContractEditCreditParamsRateType = "COMMIT_RATE"
)

type V2ContractGetEditHistoryParams struct {
	ContractID string `json:"contract_id" api:"required" format:"uuid"`
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	paramObj
}

func (r V2ContractGetEditHistoryParams) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractGetEditHistoryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractGetEditHistoryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
