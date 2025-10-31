// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v2/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v2/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/packages/param"
	"github.com/Metronome-Industries/metronome-go/v2/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/v2/shared"
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
	return
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
	return
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
//   - Contract editing must be enabled to use this endpoint. Reach out to your
//     Metronome representative to learn more.
func (r *V2ContractService) Edit(ctx context.Context, body V2ContractEditParams, opts ...option.RequestOption) (res *V2ContractEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/contracts/edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
	return
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
	return
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
	return
}

type V2ContractGetResponse struct {
	Data shared.ContractV2 `json:"data,required"`
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
	Data []shared.ContractV2 `json:"data,required"`
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
	Data shared.ID `json:"data,required"`
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

type V2ContractEditCommitResponse struct {
	Data shared.ID `json:"data,required"`
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
	Data shared.ID `json:"data,required"`
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
	Data []V2ContractGetEditHistoryResponseData `json:"data,required"`
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
	ID                                      string                                                   `json:"id,required" format:"uuid"`
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
	AddSubscriptions        []shared.Subscription                                        `json:"add_subscriptions"`
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
	UpdateContractName                         string                                                                         `json:"update_contract_name,nullable"`
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
	ID      string                                               `json:"id,required" format:"uuid"`
	Product V2ContractGetEditHistoryResponseDataAddCommitProduct `json:"product,required"`
	// Any of "PREPAID", "POSTPAID".
	Type string `json:"type,required"`
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
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
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
	ID        string    `json:"id,required" format:"uuid"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	Amount    float64   `json:"amount"`
	InvoiceID string    `json:"invoice_id,nullable" format:"uuid"`
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
	ID      string                                               `json:"id,required" format:"uuid"`
	Product V2ContractGetEditHistoryResponseDataAddCreditProduct `json:"product,required"`
	// Any of "CREDIT".
	Type string `json:"type,required"`
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
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
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
	ID                    string                                                             `json:"id,required" format:"uuid"`
	StartingAt            time.Time                                                          `json:"starting_at,required" format:"date-time"`
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
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency        string            `json:"billing_frequency"`
	CommitIDs               []string          `json:"commit_ids"`
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	ProductID               string            `json:"product_id" format:"uuid"`
	ProductTags             []string          `json:"product_tags"`
	RecurringCommitIDs      []string          `json:"recurring_commit_ids"`
	RecurringCreditIDs      []string          `json:"recurring_credit_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFrequency        respjson.Field
		CommitIDs               respjson.Field
		PresentationGroupValues respjson.Field
		PricingGroupValues      respjson.Field
		ProductID               respjson.Field
		ProductTags             respjson.Field
		RecurringCommitIDs      respjson.Field
		RecurringCreditIDs      respjson.Field
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
	// Any of "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM".
	RateType   string                `json:"rate_type,required"`
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
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
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
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractGetEditHistoryResponseDataAddRecurringCommitAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractGetEditHistoryResponseDataAddRecurringCommitCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                                       `json:"priority,required"`
	Product  V2ContractGetEditHistoryResponseDataAddRecurringCommitProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
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
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
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
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
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
	Value float64 `json:"value,required"`
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
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
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
	ID string `json:"id,required" format:"uuid"`
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
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64 `json:"quantity,required"`
	UnitPrice    float64 `json:"unit_price,required"`
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

type V2ContractGetEditHistoryResponseDataAddRecurringCredit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractGetEditHistoryResponseDataAddRecurringCreditAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractGetEditHistoryResponseDataAddRecurringCreditCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                                       `json:"priority,required"`
	Product  V2ContractGetEditHistoryResponseDataAddRecurringCreditProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
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
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
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
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
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
	Value float64 `json:"value,required"`
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
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
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
	ID string `json:"id,required" format:"uuid"`
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

type V2ContractGetEditHistoryResponseDataAddResellerRoyalty struct {
	// Any of "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".
	ResellerType          string    `json:"reseller_type,required"`
	ApplicableProductIDs  []string  `json:"applicable_product_ids"`
	ApplicableProductTags []string  `json:"applicable_product_tags"`
	AwsAccountNumber      string    `json:"aws_account_number"`
	AwsOfferID            string    `json:"aws_offer_id"`
	AwsPayerReferenceID   string    `json:"aws_payer_reference_id"`
	EndingBefore          time.Time `json:"ending_before,nullable" format:"date-time"`
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
	ID       string                                                        `json:"id,required" format:"uuid"`
	Product  V2ContractGetEditHistoryResponseDataAddScheduledChargeProduct `json:"product,required"`
	Schedule shared.SchedulePointInTime                                    `json:"schedule,required"`
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
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
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

type V2ContractGetEditHistoryResponseDataAddUsageFilter struct {
	GroupKey    string   `json:"group_key,required"`
	GroupValues []string `json:"group_values,required"`
	// This will match contract starting_at value if usage filter is active from the
	// beginning of the contract.
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID             string                                                         `json:"id,required" format:"uuid"`
	AccessSchedule V2ContractGetEditHistoryResponseDataUpdateCommitAccessSchedule `json:"access_schedule"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,nullable" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags,nullable"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration                             `json:"hierarchy_configuration"`
	InvoiceSchedule        V2ContractGetEditHistoryResponseDataUpdateCommitInvoiceSchedule `json:"invoice_schedule"`
	Name                   string                                                          `json:"name"`
	NetsuiteSalesOrderID   string                                                          `json:"netsuite_sales_order_id,nullable"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority  float64 `json:"priority,nullable"`
	ProductID string  `json:"product_id" format:"uuid"`
	// If set, the commit's rate type was updated to the specified value.
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction,nullable"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessSchedule         respjson.Field
		ApplicableProductIDs   respjson.Field
		ApplicableProductTags  respjson.Field
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
	Amount float64 `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID     string  `json:"id,required" format:"uuid"`
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
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID        string    `json:"id,required" format:"uuid"`
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
	ID             string                                                         `json:"id,required" format:"uuid"`
	AccessSchedule V2ContractGetEditHistoryResponseDataUpdateCreditAccessSchedule `json:"access_schedule"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfiguration `json:"hierarchy_configuration"`
	Name                   string                              `json:"name"`
	NetsuiteSalesOrderID   string                              `json:"netsuite_sales_order_id,nullable"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority float64 `json:"priority,nullable"`
	// If set, the credit's rate type was updated to the specified value.
	//
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccessSchedule         respjson.Field
		HierarchyConfiguration respjson.Field
		Name                   respjson.Field
		NetsuiteSalesOrderID   respjson.Field
		Priority               respjson.Field
		RateType               respjson.Field
		RolloverFraction       respjson.Field
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
	Amount float64 `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID     string  `json:"id,required" format:"uuid"`
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
	ID string `json:"id,required" format:"uuid"`
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
	AmountDistribution string `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
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
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
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
	CustomCreditTypeID string `json:"custom_credit_type_id,nullable" format:"uuid"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit             respjson.Field
		CustomCreditTypeID respjson.Field
		IsEnabled          respjson.Field
		PaymentGateConfig  respjson.Field
		RechargeToAmount   respjson.Field
		ThresholdAmount    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
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
	ApplicableProductIDs []string `json:"applicable_product_ids,nullable" format:"uuid"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags,nullable"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []shared.CommitSpecifierInput `json:"specifiers,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
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

type V2ContractGetEditHistoryResponseDataUpdateRecurringCommit struct {
	ID            string                                                                 `json:"id,required" format:"uuid"`
	AccessAmount  V2ContractGetEditHistoryResponseDataUpdateRecurringCommitAccessAmount  `json:"access_amount"`
	EndingBefore  time.Time                                                              `json:"ending_before" format:"date-time"`
	InvoiceAmount V2ContractGetEditHistoryResponseDataUpdateRecurringCommitInvoiceAmount `json:"invoice_amount"`
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType string `json:"rate_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccessAmount  respjson.Field
		EndingBefore  respjson.Field
		InvoiceAmount respjson.Field
		RateType      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
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

type V2ContractGetEditHistoryResponseDataUpdateRecurringCredit struct {
	ID           string                                                                `json:"id,required" format:"uuid"`
	AccessAmount V2ContractGetEditHistoryResponseDataUpdateRecurringCreditAccessAmount `json:"access_amount"`
	EndingBefore time.Time                                                             `json:"ending_before" format:"date-time"`
	// Any of "LIST_RATE", "COMMIT_RATE".
	RateType string `json:"rate_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccessAmount respjson.Field
		EndingBefore respjson.Field
		RateType     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
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

type V2ContractGetEditHistoryResponseDataUpdateRefundInvoice struct {
	Date      time.Time `json:"date,required" format:"date-time"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
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
	ID                   string                                                                   `json:"id,required" format:"uuid"`
	InvoiceSchedule      V2ContractGetEditHistoryResponseDataUpdateScheduledChargeInvoiceSchedule `json:"invoice_schedule"`
	Name                 string                                                                   `json:"name"`
	NetsuiteSalesOrderID string                                                                   `json:"netsuite_sales_order_id,nullable"`
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
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID        string    `json:"id,required" format:"uuid"`
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
	Commit shared.UpdateBaseThresholdCommit `json:"commit"`
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
		Commit            respjson.Field
		IsEnabled         respjson.Field
		PaymentGateConfig respjson.Field
		ThresholdAmount   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ContractGetEditHistoryResponseDataUpdateSubscription struct {
	ID              string                                                                 `json:"id,required" format:"uuid"`
	EndingBefore    time.Time                                                              `json:"ending_before" format:"date-time"`
	QuantityUpdates []V2ContractGetEditHistoryResponseDataUpdateSubscriptionQuantityUpdate `json:"quantity_updates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		EndingBefore    respjson.Field
		QuantityUpdates respjson.Field
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
	StartingAt    time.Time `json:"starting_at,required" format:"date-time"`
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

type V2ContractGetParams struct {
	ContractID string `json:"contract_id,required" format:"uuid"`
	CustomerID string `json:"customer_id,required" format:"uuid"`
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
	CustomerID string `json:"customer_id,required" format:"uuid"`
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
	ContractID string `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is being edited
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// RFC 3339 timestamp indicating when the contract will end (exclusive).
	UpdateContractEndDate param.Opt[time.Time] `json:"update_contract_end_date,omitzero" format:"date-time"`
	// Value to update the contract name to. If not provided, the contract name will
	// remain unchanged.
	UpdateContractName param.Opt[string] `json:"update_contract_name,omitzero"`
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
	AddProfessionalServices        []V2ContractEditParamsAddProfessionalService `json:"add_professional_services,omitzero"`
	AddRecurringCommits            []V2ContractEditParamsAddRecurringCommit     `json:"add_recurring_commits,omitzero"`
	AddRecurringCredits            []V2ContractEditParamsAddRecurringCredit     `json:"add_recurring_credits,omitzero"`
	AddResellerRoyalties           []V2ContractEditParamsAddResellerRoyalty     `json:"add_reseller_royalties,omitzero"`
	AddScheduledCharges            []V2ContractEditParamsAddScheduledCharge     `json:"add_scheduled_charges,omitzero"`
	AddSpendThresholdConfiguration shared.SpendThresholdConfigurationV2Param    `json:"add_spend_threshold_configuration,omitzero"`
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
	BillingProviderConfiguration V2ContractEditParamsAddBillingProviderConfigurationUpdateBillingProviderConfiguration `json:"billing_provider_configuration,omitzero,required"`
	// Indicates when the billing provider will be active on the contract. Any charges
	// accrued during the schedule will be billed to the indicated billing provider.
	Schedule V2ContractEditParamsAddBillingProviderConfigurationUpdateSchedule `json:"schedule,omitzero,required"`
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
	// Any of "START_OF_CURRENT_PERIOD".
	EffectiveAt string `json:"effective_at,omitzero,required"`
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
		"effective_at", "START_OF_CURRENT_PERIOD",
	)
}

// The properties ProductID, Type are required.
type V2ContractEditParamsAddCommit struct {
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
	ScheduleItems []V2ContractEditParamsAddCommitAccessScheduleScheduleItem `json:"schedule_items,omitzero,required"`
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
	Amount float64 `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
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
	AmountDistribution string `json:"amount_distribution,omitzero,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".
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
	PaymentGateType string `json:"payment_gate_type,omitzero,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig V2ContractEditParamsAddCommitPaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config,omitzero"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig V2ContractEditParamsAddCommitPaymentGateConfigStripeConfig `json:"stripe_config,omitzero"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	//
	// Any of "NONE", "STRIPE", "ANROK", "AVALARA", "PRECALCULATED".
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
		"tax_type", "NONE", "STRIPE", "ANROK", "AVALARA", "PRECALCULATED",
	)
}

// Only applicable if using PRECALCULATED as your tax type.
//
// The property TaxAmount is required.
type V2ContractEditParamsAddCommitPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
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

// The properties AccessSchedule, ProductID are required.
type V2ContractEditParamsAddCredit struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule V2ContractEditParamsAddCreditAccessSchedule `json:"access_schedule,omitzero,required"`
	ProductID      string                                      `json:"product_id,required" format:"uuid"`
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
	ScheduleItems []V2ContractEditParamsAddCreditAccessScheduleScheduleItem `json:"schedule_items,omitzero,required"`
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
	Amount float64 `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
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
	ProductID string `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule V2ContractEditParamsAddDiscountSchedule `json:"schedule,omitzero,required"`
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
	AmountDistribution string `json:"amount_distribution,omitzero,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".
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
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
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
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to commits
	// created by the specified recurring credit ids.
	RecurringCreditIDs []string `json:"recurring_credit_ids,omitzero"`
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

func (r V2ContractEditParamsAddOverrideOverwriteRate) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddOverrideOverwriteRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddOverrideOverwriteRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2ContractEditParamsAddOverrideOverwriteRate](
		"rate_type", "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM",
	)
}

// The property Multiplier is required.
type V2ContractEditParamsAddOverrideTier struct {
	Multiplier float64            `json:"multiplier,required"`
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
	AccessAmount V2ContractEditParamsAddRecurringCommitAccessAmount `json:"access_amount,omitzero,required"`
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
	CommitDuration V2ContractEditParamsAddRecurringCommitCommitDuration `json:"commit_duration,omitzero,required"`
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
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration shared.CommitHierarchyConfigurationParam `json:"hierarchy_configuration,omitzero"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount V2ContractEditParamsAddRecurringCommitInvoiceAmount `json:"invoice_amount,omitzero"`
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
		"recurrence_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// The amount of commit to grant.
//
// The properties CreditTypeID, UnitPrice are required.
type V2ContractEditParamsAddRecurringCommitAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
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
	Value float64 `json:"value,required"`
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
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64 `json:"quantity,required"`
	UnitPrice    float64 `json:"unit_price,required"`
	paramObj
}

func (r V2ContractEditParamsAddRecurringCommitInvoiceAmount) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddRecurringCommitInvoiceAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attach a subscription to the recurring commit/credit.
//
// The properties ApplySeatIncreaseConfig, SubscriptionID are required.
type V2ContractEditParamsAddRecurringCommitSubscriptionConfig struct {
	ApplySeatIncreaseConfig V2ContractEditParamsAddRecurringCommitSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,omitzero,required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID string `json:"subscription_id,required"`
	// If set to POOLED, allocation added per seat is pooled across the account. (BETA)
	// If set to INDIVIDUAL, each seat in the subscription will have its own
	// allocation.
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
	IsProrated bool `json:"is_prorated,required"`
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
	AccessAmount V2ContractEditParamsAddRecurringCreditAccessAmount `json:"access_amount,omitzero,required"`
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
	CommitDuration V2ContractEditParamsAddRecurringCreditCommitDuration `json:"commit_duration,omitzero,required"`
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
	// Optional configuration for recurring credit hierarchy access control
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
		"recurrence_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// The amount of commit to grant.
//
// The properties CreditTypeID, UnitPrice are required.
type V2ContractEditParamsAddRecurringCreditAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
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
	Value float64 `json:"value,required"`
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

// Attach a subscription to the recurring commit/credit.
//
// The properties ApplySeatIncreaseConfig, SubscriptionID are required.
type V2ContractEditParamsAddRecurringCreditSubscriptionConfig struct {
	ApplySeatIncreaseConfig V2ContractEditParamsAddRecurringCreditSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,omitzero,required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID string `json:"subscription_id,required"`
	// If set to POOLED, allocation added per seat is pooled across the account. (BETA)
	// If set to INDIVIDUAL, each seat in the subscription will have its own
	// allocation.
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
	IsProrated bool `json:"is_prorated,required"`
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

// The properties ProductID, Schedule are required.
type V2ContractEditParamsAddScheduledCharge struct {
	ProductID string `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule V2ContractEditParamsAddScheduledChargeSchedule `json:"schedule,omitzero,required"`
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
	AmountDistribution string `json:"amount_distribution,omitzero,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// Any of "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".
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

func (r V2ContractEditParamsAddScheduledChargeScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsAddScheduledChargeScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsAddScheduledChargeScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CollectionSchedule, Proration, SubscriptionRate are required.
type V2ContractEditParamsAddSubscription struct {
	// Any of "ADVANCE", "ARREARS".
	CollectionSchedule string                                              `json:"collection_schedule,omitzero,required"`
	Proration          V2ContractEditParamsAddSubscriptionProration        `json:"proration,omitzero,required"`
	SubscriptionRate   V2ContractEditParamsAddSubscriptionSubscriptionRate `json:"subscription_rate,omitzero,required"`
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
	TemporaryID param.Opt[string] `json:"temporary_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// Determines how the subscription's quantity is controlled. Defaults to
	// QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified
	// directly on the subscription. `initial_quantity` must be provided with this
	// option. Compatible with recurring commits/credits that use POOLED allocation.
	// **SEAT_BASED**: (BETA) Use when you want to pass specific seat identifiers (e.g.
	// add user_123) to increment and decrement a subscription quantity, rather than
	// directly providing the quantity. You must use a **SEAT_BASED** subscription to
	// use a linked recurring credit with an allocation per seat. `seat_config` must be
	// provided with this option.
	//
	// Any of "SEAT_BASED", "QUANTITY_ONLY".
	QuantityManagementMode string `json:"quantity_management_mode,omitzero"`
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
	InvoiceBehavior string `json:"invoice_behavior,omitzero"`
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

// The properties BillingFrequency, ProductID are required.
type V2ContractEditParamsAddSubscriptionSubscriptionRate struct {
	// Frequency to bill subscription with. Together with product_id, must match
	// existing rate on the rate card.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero,required"`
	// Must be subscription type product
	ProductID string `json:"product_id,required" format:"uuid"`
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

// The property ID is required.
type V2ContractEditParamsArchiveCommit struct {
	ID string `json:"id,required" format:"uuid"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID string `json:"id,required" format:"uuid"`
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
	CommitID             string             `json:"commit_id,required" format:"uuid"`
	NetsuiteSalesOrderID param.Opt[string]  `json:"netsuite_sales_order_id,omitzero"`
	Priority             param.Opt[float64] `json:"priority,omitzero"`
	RolloverFraction     param.Opt[float64] `json:"rollover_fraction,omitzero"`
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
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID           string               `json:"id,required" format:"uuid"`
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
	Timestamp time.Time          `json:"timestamp,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID        string               `json:"id,required" format:"uuid"`
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
	CreditID             string             `json:"credit_id,required" format:"uuid"`
	NetsuiteSalesOrderID param.Opt[string]  `json:"netsuite_sales_order_id,omitzero"`
	Priority             param.Opt[float64] `json:"priority,omitzero"`
	ProductID            param.Opt[string]  `json:"product_id,omitzero" format:"uuid"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitzero" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
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
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID           string               `json:"id,required" format:"uuid"`
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
	ThresholdAmount   param.Opt[float64]                                                   `json:"threshold_amount,omitzero"`
	Commit            V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommit `json:"commit,omitzero"`
	PaymentGateConfig shared.PaymentGateConfigV2Param                                      `json:"payment_gate_config,omitzero"`
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
	type shadow V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommit
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property RecurringCommitID is required.
type V2ContractEditParamsUpdateRecurringCommit struct {
	RecurringCommitID string                                                 `json:"recurring_commit_id,required" format:"uuid"`
	EndingBefore      param.Opt[time.Time]                                   `json:"ending_before,omitzero" format:"date-time"`
	AccessAmount      V2ContractEditParamsUpdateRecurringCommitAccessAmount  `json:"access_amount,omitzero"`
	InvoiceAmount     V2ContractEditParamsUpdateRecurringCommitInvoiceAmount `json:"invoice_amount,omitzero"`
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

// The property RecurringCreditID is required.
type V2ContractEditParamsUpdateRecurringCredit struct {
	RecurringCreditID string                                                `json:"recurring_credit_id,required" format:"uuid"`
	EndingBefore      param.Opt[time.Time]                                  `json:"ending_before,omitzero" format:"date-time"`
	AccessAmount      V2ContractEditParamsUpdateRecurringCreditAccessAmount `json:"access_amount,omitzero"`
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

// The property ScheduledChargeID is required.
type V2ContractEditParamsUpdateScheduledCharge struct {
	ScheduledChargeID    string                                                   `json:"scheduled_charge_id,required" format:"uuid"`
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
	Timestamp time.Time          `json:"timestamp,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID        string               `json:"id,required" format:"uuid"`
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
	ThresholdAmount   param.Opt[float64]                    `json:"threshold_amount,omitzero"`
	Commit            shared.UpdateBaseThresholdCommitParam `json:"commit,omitzero"`
	PaymentGateConfig shared.PaymentGateConfigV2Param       `json:"payment_gate_config,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateSpendThresholdConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSpendThresholdConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSpendThresholdConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SubscriptionID is required.
type V2ContractEditParamsUpdateSubscription struct {
	SubscriptionID string               `json:"subscription_id,required" format:"uuid"`
	EndingBefore   param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	// Quantity changes are applied on the effective date based on the order which they
	// are sent. For example, if I scheduled the quantity to be 12 on May 21 and then
	// scheduled a quantity delta change of -1, the result from that day would be 11.
	QuantityUpdates []V2ContractEditParamsUpdateSubscriptionQuantityUpdate `json:"quantity_updates,omitzero"`
	paramObj
}

func (r V2ContractEditParamsUpdateSubscription) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractEditParamsUpdateSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractEditParamsUpdateSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property StartingAt is required.
type V2ContractEditParamsUpdateSubscriptionQuantityUpdate struct {
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
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

type V2ContractEditCommitParams struct {
	// ID of the commit to edit
	CommitID string `json:"commit_id,required" format:"uuid"`
	// ID of the customer whose commit is being edited
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// ID of contract to use for invoicing
	InvoiceContractID param.Opt[string] `json:"invoice_contract_id,omitzero" format:"uuid"`
	ProductID         param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
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
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID           string               `json:"id,required" format:"uuid"`
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
	Timestamp time.Time          `json:"timestamp,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID        string               `json:"id,required" format:"uuid"`
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
	CreditID string `json:"credit_id,required" format:"uuid"`
	// ID of the customer whose credit is being edited
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority  param.Opt[float64] `json:"priority,omitzero"`
	ProductID param.Opt[string]  `json:"product_id,omitzero" format:"uuid"`
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
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
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
	ID string `json:"id,required" format:"uuid"`
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
	ID           string               `json:"id,required" format:"uuid"`
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
	ContractID string `json:"contract_id,required" format:"uuid"`
	CustomerID string `json:"customer_id,required" format:"uuid"`
	paramObj
}

func (r V2ContractGetEditHistoryParams) MarshalJSON() (data []byte, err error) {
	type shadow V2ContractGetEditHistoryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2ContractGetEditHistoryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
