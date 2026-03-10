// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v3/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/v3/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"github.com/Metronome-Industries/metronome-go/v3/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/v3/packages/param"
	"github.com/Metronome-Industries/metronome-go/v3/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/v3/shared"
)

// V1PackageService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PackageService] method instead.
type V1PackageService struct {
	Options []option.RequestOption
}

// NewV1PackageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1PackageService(opts ...option.RequestOption) (r V1PackageService) {
	r = V1PackageService{}
	r.Options = opts
	return
}

// Create a package that defines a set of reusable, time-relative contract terms
// that can be used across cohorts of customers. Packages provide an abstraction
// layer on top of rate cards to provide an easy way to provision customers with
// standard pricing.
//
// ### **Use this endpoint to:**
//
//   - Model standard pay-as-you-go pricing packages that can be easily reused across
//     customers
//   - Define standardized contract terms and discounting for sales-led motions
//   - Set aliases for the package to facilitate easy package transition. Aliases are
//     human-readable names that you can use in the place of the id of the package
//     when provisioning a customer’s contract. By using an alias, you can easily
//     create a contract and provision a customer by choosing the “Starter Plan”
//     package, without storing the package ID in your internal systems. This is
//     helpful when launching terms for a package, as you can create a new package
//     with the “Starter Plan” alias scheduled to be assigned without updating your
//     provisioning code.
//
// ### Key input fields:
//
//   - `starting_at_offset`: Starting date relative to contract start. Generates the
//     `starting_at` date when a contract is provisioned using a package.
//   - `duration`: Duration starting from `starting_at_offset`. Generates the
//     `ending_before` date when a contract is provisioned using a package.
//   - `date_offset`: Date relative to contract start. Used for point-in-time dates
//     without a duration.
//   - `aliases`: Human-readable name to use when provisioning contracts with a
//     package.
//
// ### Usage guidelines:
//
//   - Use packages for standard self-serve use cases where customers have consistent
//     terms. For customers with negotiated custom contract terms, use the
//     `createContract` endpoint for maximum flexibility.
//   - Billing provider configuration can be set when creating a package by using
//     `billing_provider` and `delivery_method`. To provision a customer successfully
//     with a package, the customer must have one and only one billing provider
//     configuration that matches the billing provider configuration set on the
//     package.
//   - A package alias can only be used by one package at a time. If you create a new
//     package with an alias that is already in use by another package, the original
//     package’s alias schedule will be updated. The alias will reference the package
//     to which it was most recently assigned.
//   - Terms can only be specified using times relative to the contract start date.
//     Supported granularities are: `days`, `weeks`, `months`, `years`
//   - Packages cannot be edited once created. Use the rate card to easily add new
//     rates across all of your customers or make direct edits to a contract after
//     provisioning with a package. Edited contracts will still be associated with
//     the package used during provisioning.
func (r *V1PackageService) New(ctx context.Context, body V1PackageNewParams, opts ...option.RequestOption) (res *V1PackageNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/packages/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets the details for a specific package, including name, aliases, duration, and
// terms. Use this endpoint to understand a package’s alias schedule, or display a
// specific package’s details to end customers.
func (r *V1PackageService) Get(ctx context.Context, body V1PackageGetParams, opts ...option.RequestOption) (res *V1PackageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/packages/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists all packages with details including name, aliases, duration, and terms. To
// view contracts on a specific package, use the `listContractsOnPackage` endpoint.
func (r *V1PackageService) List(ctx context.Context, params V1PackageListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1PackageListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/packages/list"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Lists all packages with details including name, aliases, duration, and terms. To
// view contracts on a specific package, use the `listContractsOnPackage` endpoint.
func (r *V1PackageService) ListAutoPaging(ctx context.Context, params V1PackageListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1PackageListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Archive a package. Archived packages cannot be used to create new contracts.
// However, existing contracts associated with the package will continue to
// function as normal. Once you archive a package, you can still retrieve it in the
// UI and API, but you cannot unarchive it.
func (r *V1PackageService) Archive(ctx context.Context, body V1PackageArchiveParams, opts ...option.RequestOption) (res *V1PackageArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/packages/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// For a given package, returns all contract IDs and customer IDs associated with
// the package over a specific time period.
//
// ### Use this endpoint to:
//
//   - Understand which customers are provisioned on a package at any given time for
//     internal cohort management
//   - Manage customer migrations to a new package. For example, to migrate all
//     active customers to a new package, call this endpoint, end contracts, and
//     provision customers on a new package.
//
// ### **Usage guidelines:**
//
// Use the **`starting_at`**, **`covering_date`**,
// and **`include_archived`** parameters to filter the list of returned contracts.
// For example, to list only currently active contracts,
// pass **`covering_date`** equal to the current time.
func (r *V1PackageService) ListContractsOnPackage(ctx context.Context, params V1PackageListContractsOnPackageParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1PackageListContractsOnPackageResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/packages/listContractsOnPackage"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// For a given package, returns all contract IDs and customer IDs associated with
// the package over a specific time period.
//
// ### Use this endpoint to:
//
//   - Understand which customers are provisioned on a package at any given time for
//     internal cohort management
//   - Manage customer migrations to a new package. For example, to migrate all
//     active customers to a new package, call this endpoint, end contracts, and
//     provision customers on a new package.
//
// ### **Usage guidelines:**
//
// Use the **`starting_at`**, **`covering_date`**,
// and **`include_archived`** parameters to filter the list of returned contracts.
// For example, to list only currently active contracts,
// pass **`covering_date`** equal to the current time.
func (r *V1PackageService) ListContractsOnPackageAutoPaging(ctx context.Context, params V1PackageListContractsOnPackageParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1PackageListContractsOnPackageResponse] {
	return pagination.NewCursorPageAutoPager(r.ListContractsOnPackage(ctx, params, opts...))
}

type V1PackageNewResponse struct {
	Data shared.ID `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PackageNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponse struct {
	Data V1PackageGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseData struct {
	ID                     string                                         `json:"id" api:"required" format:"uuid"`
	Commits                []V1PackageGetResponseDataCommit               `json:"commits" api:"required"`
	CreatedAt              time.Time                                      `json:"created_at" api:"required" format:"date-time"`
	CreatedBy              string                                         `json:"created_by" api:"required"`
	Overrides              []V1PackageGetResponseDataOverride             `json:"overrides" api:"required"`
	ScheduledCharges       []V1PackageGetResponseDataScheduledCharge      `json:"scheduled_charges" api:"required"`
	UsageStatementSchedule V1PackageGetResponseDataUsageStatementSchedule `json:"usage_statement_schedule" api:"required"`
	Aliases                []V1PackageGetResponseDataAlias                `json:"aliases"`
	ArchivedAt             time.Time                                      `json:"archived_at" format:"date-time"`
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace", "metronome".
	BillingProvider string                           `json:"billing_provider"`
	Credits         []V1PackageGetResponseDataCredit `json:"credits"`
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod string                           `json:"delivery_method"`
	Duration       V1PackageGetResponseDataDuration `json:"duration"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
	// prices automatically. EXPLICIT prioritization requires specifying priorities for
	// each multiplier; the one with the lowest priority value will be prioritized
	// first.
	//
	// Any of "LOWEST_MULTIPLIER", "EXPLICIT".
	MultiplierOverridePrioritization     string                                      `json:"multiplier_override_prioritization"`
	Name                                 string                                      `json:"name"`
	NetPaymentTermsDays                  float64                                     `json:"net_payment_terms_days"`
	PrepaidBalanceThresholdConfiguration shared.PrepaidBalanceThresholdConfiguration `json:"prepaid_balance_threshold_configuration"`
	RateCardID                           string                                      `json:"rate_card_id" format:"uuid"`
	RecurringCommits                     []V1PackageGetResponseDataRecurringCommit   `json:"recurring_commits"`
	RecurringCredits                     []V1PackageGetResponseDataRecurringCredit   `json:"recurring_credits"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	//
	// Any of "ALL".
	ScheduledChargesOnUsageInvoices string                                 `json:"scheduled_charges_on_usage_invoices"`
	SpendThresholdConfiguration     shared.SpendThresholdConfiguration     `json:"spend_threshold_configuration"`
	Subscriptions                   []V1PackageGetResponseDataSubscription `json:"subscriptions"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string `json:"uniqueness_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                   respjson.Field
		Commits                              respjson.Field
		CreatedAt                            respjson.Field
		CreatedBy                            respjson.Field
		Overrides                            respjson.Field
		ScheduledCharges                     respjson.Field
		UsageStatementSchedule               respjson.Field
		Aliases                              respjson.Field
		ArchivedAt                           respjson.Field
		BillingProvider                      respjson.Field
		Credits                              respjson.Field
		DeliveryMethod                       respjson.Field
		Duration                             respjson.Field
		MultiplierOverridePrioritization     respjson.Field
		Name                                 respjson.Field
		NetPaymentTermsDays                  respjson.Field
		PrepaidBalanceThresholdConfiguration respjson.Field
		RateCardID                           respjson.Field
		RecurringCommits                     respjson.Field
		RecurringCredits                     respjson.Field
		ScheduledChargesOnUsageInvoices      respjson.Field
		SpendThresholdConfiguration          respjson.Field
		Subscriptions                        respjson.Field
		UniquenessKey                        respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCommit struct {
	ID      string                                `json:"id" api:"required" format:"uuid"`
	Product V1PackageGetResponseDataCommitProduct `json:"product" api:"required"`
	// Any of "PREPAID", "POSTPAID".
	Type string `json:"type" api:"required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule        V1PackageGetResponseDataCommitAccessSchedule `json:"access_schedule"`
	ApplicableProductIDs  []string                                     `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                                     `json:"applicable_product_tags"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule V1PackageGetResponseDataCommitInvoiceSchedule `json:"invoice_schedule"`
	Name            string                                        `json:"name"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []shared.CommitSpecifier `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Product               respjson.Field
		Type                  respjson.Field
		AccessSchedule        respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		CustomFields          respjson.Field
		Description           respjson.Field
		InvoiceSchedule       respjson.Field
		Name                  respjson.Field
		Priority              respjson.Field
		RateType              respjson.Field
		RolloverFraction      respjson.Field
		Specifiers            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataCommit) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCommitProduct struct {
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
func (r V1PackageGetResponseDataCommitProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataCommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The schedule that the customer will gain access to the credits purposed with
// this commit.
type V1PackageGetResponseDataCommitAccessSchedule struct {
	CreditType    shared.CreditTypeData                                      `json:"credit_type" api:"required"`
	ScheduleItems []V1PackageGetResponseDataCommitAccessScheduleScheduleItem `json:"schedule_items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType    respjson.Field
		ScheduleItems respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataCommitAccessSchedule) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataCommitAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCommitAccessScheduleScheduleItem struct {
	ID               string                                                                   `json:"id" api:"required" format:"uuid"`
	Amount           float64                                                                  `json:"amount" api:"required"`
	Duration         V1PackageGetResponseDataCommitAccessScheduleScheduleItemDuration         `json:"duration" api:"required"`
	StartingAtOffset V1PackageGetResponseDataCommitAccessScheduleScheduleItemStartingAtOffset `json:"starting_at_offset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Amount           respjson.Field
		Duration         respjson.Field
		StartingAtOffset respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataCommitAccessScheduleScheduleItem) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataCommitAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCommitAccessScheduleScheduleItemDuration struct {
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
func (r V1PackageGetResponseDataCommitAccessScheduleScheduleItemDuration) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataCommitAccessScheduleScheduleItemDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCommitAccessScheduleScheduleItemStartingAtOffset struct {
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
func (r V1PackageGetResponseDataCommitAccessScheduleScheduleItemStartingAtOffset) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataCommitAccessScheduleScheduleItemStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The schedule that the customer will be invoiced for this commit.
type V1PackageGetResponseDataCommitInvoiceSchedule struct {
	CreditType shared.CreditTypeData `json:"credit_type" api:"required"`
	// If true, this schedule will not generate an invoice.
	DoNotInvoice  bool                                                        `json:"do_not_invoice" api:"required"`
	ScheduleItems []V1PackageGetResponseDataCommitInvoiceScheduleScheduleItem `json:"schedule_items" api:"required"`
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
func (r V1PackageGetResponseDataCommitInvoiceSchedule) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCommitInvoiceScheduleScheduleItem struct {
	ID         string                                                              `json:"id" api:"required" format:"uuid"`
	DateOffset V1PackageGetResponseDataCommitInvoiceScheduleScheduleItemDateOffset `json:"date_offset" api:"required"`
	Quantity   float64                                                             `json:"quantity" api:"required"`
	UnitPrice  float64                                                             `json:"unit_price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DateOffset  respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataCommitInvoiceScheduleScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataCommitInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCommitInvoiceScheduleScheduleItemDateOffset struct {
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
func (r V1PackageGetResponseDataCommitInvoiceScheduleScheduleItemDateOffset) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataCommitInvoiceScheduleScheduleItemDateOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataOverride struct {
	ID                    string                                              `json:"id" api:"required" format:"uuid"`
	OverrideSpecifiers    []V1PackageGetResponseDataOverrideOverrideSpecifier `json:"override_specifiers" api:"required"`
	StartingAtOffset      V1PackageGetResponseDataOverrideStartingAtOffset    `json:"starting_at_offset" api:"required"`
	ApplicableProductTags []string                                            `json:"applicable_product_tags"`
	Duration              V1PackageGetResponseDataOverrideDuration            `json:"duration"`
	Entitled              bool                                                `json:"entitled"`
	IsCommitSpecific      bool                                                `json:"is_commit_specific"`
	Multiplier            float64                                             `json:"multiplier"`
	OverrideTiers         []shared.OverrideTier                               `json:"override_tiers"`
	OverwriteRate         shared.OverwriteRate                                `json:"overwrite_rate"`
	Priority              float64                                             `json:"priority"`
	Product               V1PackageGetResponseDataOverrideProduct             `json:"product"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	Target string `json:"target"`
	// Any of "OVERWRITE", "MULTIPLIER", "TIERED".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		OverrideSpecifiers    respjson.Field
		StartingAtOffset      respjson.Field
		ApplicableProductTags respjson.Field
		Duration              respjson.Field
		Entitled              respjson.Field
		IsCommitSpecific      respjson.Field
		Multiplier            respjson.Field
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
func (r V1PackageGetResponseDataOverride) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataOverrideOverrideSpecifier struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency           string            `json:"billing_frequency"`
	CommitTemplateIDs          []string          `json:"commit_template_ids"`
	PresentationGroupValues    map[string]string `json:"presentation_group_values"`
	PricingGroupValues         map[string]string `json:"pricing_group_values"`
	ProductID                  string            `json:"product_id" format:"uuid"`
	ProductTags                []string          `json:"product_tags"`
	RecurringCommitTemplateIDs []string          `json:"recurring_commit_template_ids"`
	RecurringCreditTemplateIDs []string          `json:"recurring_credit_template_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFrequency           respjson.Field
		CommitTemplateIDs          respjson.Field
		PresentationGroupValues    respjson.Field
		PricingGroupValues         respjson.Field
		ProductID                  respjson.Field
		ProductTags                respjson.Field
		RecurringCommitTemplateIDs respjson.Field
		RecurringCreditTemplateIDs respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataOverrideOverrideSpecifier) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataOverrideOverrideSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataOverrideStartingAtOffset struct {
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
func (r V1PackageGetResponseDataOverrideStartingAtOffset) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataOverrideStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataOverrideDuration struct {
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
func (r V1PackageGetResponseDataOverrideDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataOverrideDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataOverrideProduct struct {
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
func (r V1PackageGetResponseDataOverrideProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataOverrideProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataScheduledCharge struct {
	ID       string                                          `json:"id" api:"required" format:"uuid"`
	Product  V1PackageGetResponseDataScheduledChargeProduct  `json:"product" api:"required"`
	Schedule V1PackageGetResponseDataScheduledChargeSchedule `json:"schedule" api:"required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	Name         string            `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Product      respjson.Field
		Schedule     respjson.Field
		CustomFields respjson.Field
		Description  respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataScheduledCharge) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataScheduledChargeProduct struct {
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
func (r V1PackageGetResponseDataScheduledChargeProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataScheduledChargeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataScheduledChargeSchedule struct {
	CreditType    shared.CreditTypeData                                         `json:"credit_type" api:"required"`
	ScheduleItems []V1PackageGetResponseDataScheduledChargeScheduleScheduleItem `json:"schedule_items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType    respjson.Field
		ScheduleItems respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataScheduledChargeSchedule) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataScheduledChargeSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataScheduledChargeScheduleScheduleItem struct {
	ID         string                                                                `json:"id" api:"required" format:"uuid"`
	DateOffset V1PackageGetResponseDataScheduledChargeScheduleScheduleItemDateOffset `json:"date_offset" api:"required"`
	Quantity   float64                                                               `json:"quantity" api:"required"`
	UnitPrice  float64                                                               `json:"unit_price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DateOffset  respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataScheduledChargeScheduleScheduleItem) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataScheduledChargeScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataScheduledChargeScheduleScheduleItemDateOffset struct {
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
func (r V1PackageGetResponseDataScheduledChargeScheduleScheduleItemDateOffset) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataScheduledChargeScheduleScheduleItemDateOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataUsageStatementSchedule struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Frequency   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataUsageStatementSchedule) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataUsageStatementSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataAlias struct {
	Name         string    `json:"name" api:"required"`
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataAlias) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCredit struct {
	ID                    string                                       `json:"id" api:"required" format:"uuid"`
	Product               V1PackageGetResponseDataCreditProduct        `json:"product" api:"required"`
	AccessSchedule        V1PackageGetResponseDataCreditAccessSchedule `json:"access_schedule"`
	ApplicableProductIDs  []string                                     `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                                     `json:"applicable_product_tags"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	Name         string            `json:"name"`
	Priority     float64           `json:"priority"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []shared.CommitSpecifier `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Product               respjson.Field
		AccessSchedule        respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		CustomFields          respjson.Field
		Description           respjson.Field
		Name                  respjson.Field
		Priority              respjson.Field
		RateType              respjson.Field
		Specifiers            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataCredit) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCreditProduct struct {
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
func (r V1PackageGetResponseDataCreditProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataCreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCreditAccessSchedule struct {
	CreditType    shared.CreditTypeData                                      `json:"credit_type" api:"required"`
	ScheduleItems []V1PackageGetResponseDataCreditAccessScheduleScheduleItem `json:"schedule_items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType    respjson.Field
		ScheduleItems respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataCreditAccessSchedule) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataCreditAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCreditAccessScheduleScheduleItem struct {
	ID               string                                                                   `json:"id" api:"required" format:"uuid"`
	Amount           float64                                                                  `json:"amount" api:"required"`
	Duration         V1PackageGetResponseDataCreditAccessScheduleScheduleItemDuration         `json:"duration" api:"required"`
	StartingAtOffset V1PackageGetResponseDataCreditAccessScheduleScheduleItemStartingAtOffset `json:"starting_at_offset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Amount           respjson.Field
		Duration         respjson.Field
		StartingAtOffset respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataCreditAccessScheduleScheduleItem) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataCreditAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCreditAccessScheduleScheduleItemDuration struct {
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
func (r V1PackageGetResponseDataCreditAccessScheduleScheduleItemDuration) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataCreditAccessScheduleScheduleItemDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataCreditAccessScheduleScheduleItemStartingAtOffset struct {
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
func (r V1PackageGetResponseDataCreditAccessScheduleScheduleItemStartingAtOffset) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataCreditAccessScheduleScheduleItemStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataDuration struct {
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
func (r V1PackageGetResponseDataDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataRecurringCommit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V1PackageGetResponseDataRecurringCommitAccessAmount `json:"access_amount" api:"required"`
	// The amount of time each of the created commits will be valid for
	CommitDuration V1PackageGetResponseDataRecurringCommitCommitDuration `json:"commit_duration" api:"required"`
	Priority       float64                                               `json:"priority" api:"required"`
	Product        V1PackageGetResponseDataRecurringCommitProduct        `json:"product" api:"required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type" api:"required"`
	// Offset relative to the contract start date that determines the start time for
	// the first commit
	StartingAtOffset V1PackageGetResponseDataRecurringCommitStartingAtOffset `json:"starting_at_offset" api:"required"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags"`
	Description           string   `json:"description"`
	// Offset relative to the recurring credit start that determines when the contract
	// will stop creating recurring commits. optional
	Duration V1PackageGetResponseDataRecurringCommitDuration `json:"duration"`
	// The amount the customer should be billed for the commit.
	InvoiceAmount V1PackageGetResponseDataRecurringCommitInvoiceAmount `json:"invoice_amount"`
	Name          string                                               `json:"name"`
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
	SubscriptionConfig V1PackageGetResponseDataRecurringCommitSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccessAmount          respjson.Field
		CommitDuration        respjson.Field
		Priority              respjson.Field
		Product               respjson.Field
		RateType              respjson.Field
		StartingAtOffset      respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		Description           respjson.Field
		Duration              respjson.Field
		InvoiceAmount         respjson.Field
		Name                  respjson.Field
		Proration             respjson.Field
		RecurrenceFrequency   respjson.Field
		RolloverFraction      respjson.Field
		Specifiers            respjson.Field
		SubscriptionConfig    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataRecurringCommit) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type V1PackageGetResponseDataRecurringCommitAccessAmount struct {
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
func (r V1PackageGetResponseDataRecurringCommitAccessAmount) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time each of the created commits will be valid for
type V1PackageGetResponseDataRecurringCommitCommitDuration struct {
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
func (r V1PackageGetResponseDataRecurringCommitCommitDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCommitCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataRecurringCommitProduct struct {
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
func (r V1PackageGetResponseDataRecurringCommitProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the contract start date that determines the start time for
// the first commit
type V1PackageGetResponseDataRecurringCommitStartingAtOffset struct {
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
func (r V1PackageGetResponseDataRecurringCommitStartingAtOffset) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCommitStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the recurring credit start that determines when the contract
// will stop creating recurring commits. optional
type V1PackageGetResponseDataRecurringCommitDuration struct {
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
func (r V1PackageGetResponseDataRecurringCommitDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount the customer should be billed for the commit.
type V1PackageGetResponseDataRecurringCommitInvoiceAmount struct {
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
func (r V1PackageGetResponseDataRecurringCommitInvoiceAmount) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attach a subscription to the recurring commit/credit.
type V1PackageGetResponseDataRecurringCommitSubscriptionConfig struct {
	// Any of "INDIVIDUAL", "POOLED".
	Allocation              string                                                                           `json:"allocation" api:"required"`
	ApplySeatIncreaseConfig V1PackageGetResponseDataRecurringCommitSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config" api:"required"`
	SubscriptionTemplateID  string                                                                           `json:"subscription_template_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allocation              respjson.Field
		ApplySeatIncreaseConfig respjson.Field
		SubscriptionTemplateID  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataRecurringCommitSubscriptionConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataRecurringCommitSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataRecurringCommitSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsProrated  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataRecurringCredit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V1PackageGetResponseDataRecurringCreditAccessAmount `json:"access_amount" api:"required"`
	// The amount of time each of the created commits will be valid for
	CommitDuration V1PackageGetResponseDataRecurringCreditCommitDuration `json:"commit_duration" api:"required"`
	Priority       float64                                               `json:"priority" api:"required"`
	Product        V1PackageGetResponseDataRecurringCreditProduct        `json:"product" api:"required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type" api:"required"`
	// Offset relative to the contract start date that determines the start time for
	// the first commit
	StartingAtOffset V1PackageGetResponseDataRecurringCreditStartingAtOffset `json:"starting_at_offset" api:"required"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags"`
	Description           string   `json:"description"`
	// Offset relative to the recurring credit start that determines when the contract
	// will stop creating recurring commits. optional
	Duration V1PackageGetResponseDataRecurringCreditDuration `json:"duration"`
	Name     string                                          `json:"name"`
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
	SubscriptionConfig V1PackageGetResponseDataRecurringCreditSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccessAmount          respjson.Field
		CommitDuration        respjson.Field
		Priority              respjson.Field
		Product               respjson.Field
		RateType              respjson.Field
		StartingAtOffset      respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		Description           respjson.Field
		Duration              respjson.Field
		Name                  respjson.Field
		Proration             respjson.Field
		RecurrenceFrequency   respjson.Field
		RolloverFraction      respjson.Field
		Specifiers            respjson.Field
		SubscriptionConfig    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataRecurringCredit) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type V1PackageGetResponseDataRecurringCreditAccessAmount struct {
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
func (r V1PackageGetResponseDataRecurringCreditAccessAmount) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time each of the created commits will be valid for
type V1PackageGetResponseDataRecurringCreditCommitDuration struct {
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
func (r V1PackageGetResponseDataRecurringCreditCommitDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCreditCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataRecurringCreditProduct struct {
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
func (r V1PackageGetResponseDataRecurringCreditProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the contract start date that determines the start time for
// the first commit
type V1PackageGetResponseDataRecurringCreditStartingAtOffset struct {
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
func (r V1PackageGetResponseDataRecurringCreditStartingAtOffset) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCreditStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the recurring credit start that determines when the contract
// will stop creating recurring commits. optional
type V1PackageGetResponseDataRecurringCreditDuration struct {
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
func (r V1PackageGetResponseDataRecurringCreditDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataRecurringCreditDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attach a subscription to the recurring commit/credit.
type V1PackageGetResponseDataRecurringCreditSubscriptionConfig struct {
	// Any of "INDIVIDUAL", "POOLED".
	Allocation              string                                                                           `json:"allocation" api:"required"`
	ApplySeatIncreaseConfig V1PackageGetResponseDataRecurringCreditSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config" api:"required"`
	SubscriptionTemplateID  string                                                                           `json:"subscription_template_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allocation              respjson.Field
		ApplySeatIncreaseConfig respjson.Field
		SubscriptionTemplateID  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataRecurringCreditSubscriptionConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataRecurringCreditSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataRecurringCreditSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsProrated  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataSubscription struct {
	// Any of "ADVANCE", "ARREARS".
	CollectionSchedule string                                               `json:"collection_schedule" api:"required"`
	Proration          V1PackageGetResponseDataSubscriptionProration        `json:"proration" api:"required"`
	SubscriptionRate   V1PackageGetResponseDataSubscriptionSubscriptionRate `json:"subscription_rate" api:"required"`
	ID                 string                                               `json:"id" format:"uuid"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields     map[string]string                            `json:"custom_fields"`
	Description      string                                       `json:"description"`
	Duration         V1PackageGetResponseDataSubscriptionDuration `json:"duration"`
	FiatCreditTypeID string                                       `json:"fiat_credit_type_id" format:"uuid"`
	InitialQuantity  float64                                      `json:"initial_quantity"`
	Name             string                                       `json:"name"`
	// Determines how the subscription's quantity is controlled. Defaults to
	// QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified
	// directly on the subscription. `initial_quantity` must be provided with this
	// option. Compatible with recurring commits/credits that use POOLED allocation.
	// **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add
	// user_123) to increment and decrement a subscription quantity, rather than
	// directly providing the quantity. You must use a SEAT_BASED subscription to use a
	// linked recurring credit with an allocation per seat. `seat_config` must be
	// provided with this option.
	//
	// Any of "SEAT_BASED", "QUANTITY_ONLY".
	QuantityManagementMode string                                               `json:"quantity_management_mode"`
	SeatConfig             V1PackageGetResponseDataSubscriptionSeatConfig       `json:"seat_config"`
	StartingAtOffset       V1PackageGetResponseDataSubscriptionStartingAtOffset `json:"starting_at_offset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionSchedule     respjson.Field
		Proration              respjson.Field
		SubscriptionRate       respjson.Field
		ID                     respjson.Field
		CustomFields           respjson.Field
		Description            respjson.Field
		Duration               respjson.Field
		FiatCreditTypeID       respjson.Field
		InitialQuantity        respjson.Field
		Name                   respjson.Field
		QuantityManagementMode respjson.Field
		SeatConfig             respjson.Field
		StartingAtOffset       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataSubscription) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataSubscriptionProration struct {
	// Any of "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE".
	InvoiceBehavior string `json:"invoice_behavior" api:"required"`
	IsProrated      bool   `json:"is_prorated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InvoiceBehavior respjson.Field
		IsProrated      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataSubscriptionProration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataSubscriptionProration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataSubscriptionSubscriptionRate struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string                                                      `json:"billing_frequency" api:"required"`
	Product          V1PackageGetResponseDataSubscriptionSubscriptionRateProduct `json:"product" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFrequency respjson.Field
		Product          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageGetResponseDataSubscriptionSubscriptionRate) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataSubscriptionSubscriptionRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataSubscriptionSubscriptionRateProduct struct {
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
func (r V1PackageGetResponseDataSubscriptionSubscriptionRateProduct) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageGetResponseDataSubscriptionSubscriptionRateProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataSubscriptionDuration struct {
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
func (r V1PackageGetResponseDataSubscriptionDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataSubscriptionDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataSubscriptionSeatConfig struct {
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
func (r V1PackageGetResponseDataSubscriptionSeatConfig) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataSubscriptionSeatConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageGetResponseDataSubscriptionStartingAtOffset struct {
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
func (r V1PackageGetResponseDataSubscriptionStartingAtOffset) RawJSON() string { return r.JSON.raw }
func (r *V1PackageGetResponseDataSubscriptionStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponse struct {
	ID                     string                                      `json:"id" api:"required" format:"uuid"`
	Commits                []V1PackageListResponseCommit               `json:"commits" api:"required"`
	CreatedAt              time.Time                                   `json:"created_at" api:"required" format:"date-time"`
	CreatedBy              string                                      `json:"created_by" api:"required"`
	Overrides              []V1PackageListResponseOverride             `json:"overrides" api:"required"`
	ScheduledCharges       []V1PackageListResponseScheduledCharge      `json:"scheduled_charges" api:"required"`
	UsageStatementSchedule V1PackageListResponseUsageStatementSchedule `json:"usage_statement_schedule" api:"required"`
	Aliases                []V1PackageListResponseAlias                `json:"aliases"`
	ArchivedAt             time.Time                                   `json:"archived_at" format:"date-time"`
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace", "metronome".
	BillingProvider V1PackageListResponseBillingProvider `json:"billing_provider"`
	Credits         []V1PackageListResponseCredit        `json:"credits"`
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod V1PackageListResponseDeliveryMethod `json:"delivery_method"`
	Duration       V1PackageListResponseDuration       `json:"duration"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
	// prices automatically. EXPLICIT prioritization requires specifying priorities for
	// each multiplier; the one with the lowest priority value will be prioritized
	// first.
	//
	// Any of "LOWEST_MULTIPLIER", "EXPLICIT".
	MultiplierOverridePrioritization     V1PackageListResponseMultiplierOverridePrioritization `json:"multiplier_override_prioritization"`
	Name                                 string                                                `json:"name"`
	NetPaymentTermsDays                  float64                                               `json:"net_payment_terms_days"`
	PrepaidBalanceThresholdConfiguration shared.PrepaidBalanceThresholdConfiguration           `json:"prepaid_balance_threshold_configuration"`
	RateCardID                           string                                                `json:"rate_card_id" format:"uuid"`
	RecurringCommits                     []V1PackageListResponseRecurringCommit                `json:"recurring_commits"`
	RecurringCredits                     []V1PackageListResponseRecurringCredit                `json:"recurring_credits"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	//
	// Any of "ALL".
	ScheduledChargesOnUsageInvoices V1PackageListResponseScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	SpendThresholdConfiguration     shared.SpendThresholdConfiguration                   `json:"spend_threshold_configuration"`
	Subscriptions                   []V1PackageListResponseSubscription                  `json:"subscriptions"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string `json:"uniqueness_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                   respjson.Field
		Commits                              respjson.Field
		CreatedAt                            respjson.Field
		CreatedBy                            respjson.Field
		Overrides                            respjson.Field
		ScheduledCharges                     respjson.Field
		UsageStatementSchedule               respjson.Field
		Aliases                              respjson.Field
		ArchivedAt                           respjson.Field
		BillingProvider                      respjson.Field
		Credits                              respjson.Field
		DeliveryMethod                       respjson.Field
		Duration                             respjson.Field
		MultiplierOverridePrioritization     respjson.Field
		Name                                 respjson.Field
		NetPaymentTermsDays                  respjson.Field
		PrepaidBalanceThresholdConfiguration respjson.Field
		RateCardID                           respjson.Field
		RecurringCommits                     respjson.Field
		RecurringCredits                     respjson.Field
		ScheduledChargesOnUsageInvoices      respjson.Field
		SpendThresholdConfiguration          respjson.Field
		Subscriptions                        respjson.Field
		UniquenessKey                        respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCommit struct {
	ID      string                             `json:"id" api:"required" format:"uuid"`
	Product V1PackageListResponseCommitProduct `json:"product" api:"required"`
	// Any of "PREPAID", "POSTPAID".
	Type string `json:"type" api:"required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule        V1PackageListResponseCommitAccessSchedule `json:"access_schedule"`
	ApplicableProductIDs  []string                                  `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                                  `json:"applicable_product_tags"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule V1PackageListResponseCommitInvoiceSchedule `json:"invoice_schedule"`
	Name            string                                     `json:"name"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType         string  `json:"rate_type"`
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []shared.CommitSpecifier `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Product               respjson.Field
		Type                  respjson.Field
		AccessSchedule        respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		CustomFields          respjson.Field
		Description           respjson.Field
		InvoiceSchedule       respjson.Field
		Name                  respjson.Field
		Priority              respjson.Field
		RateType              respjson.Field
		RolloverFraction      respjson.Field
		Specifiers            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseCommit) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCommitProduct struct {
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
func (r V1PackageListResponseCommitProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseCommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The schedule that the customer will gain access to the credits purposed with
// this commit.
type V1PackageListResponseCommitAccessSchedule struct {
	CreditType    shared.CreditTypeData                                   `json:"credit_type" api:"required"`
	ScheduleItems []V1PackageListResponseCommitAccessScheduleScheduleItem `json:"schedule_items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType    respjson.Field
		ScheduleItems respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseCommitAccessSchedule) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseCommitAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCommitAccessScheduleScheduleItem struct {
	ID               string                                                                `json:"id" api:"required" format:"uuid"`
	Amount           float64                                                               `json:"amount" api:"required"`
	Duration         V1PackageListResponseCommitAccessScheduleScheduleItemDuration         `json:"duration" api:"required"`
	StartingAtOffset V1PackageListResponseCommitAccessScheduleScheduleItemStartingAtOffset `json:"starting_at_offset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Amount           respjson.Field
		Duration         respjson.Field
		StartingAtOffset respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseCommitAccessScheduleScheduleItem) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseCommitAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCommitAccessScheduleScheduleItemDuration struct {
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
func (r V1PackageListResponseCommitAccessScheduleScheduleItemDuration) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageListResponseCommitAccessScheduleScheduleItemDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCommitAccessScheduleScheduleItemStartingAtOffset struct {
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
func (r V1PackageListResponseCommitAccessScheduleScheduleItemStartingAtOffset) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageListResponseCommitAccessScheduleScheduleItemStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The schedule that the customer will be invoiced for this commit.
type V1PackageListResponseCommitInvoiceSchedule struct {
	CreditType shared.CreditTypeData `json:"credit_type" api:"required"`
	// If true, this schedule will not generate an invoice.
	DoNotInvoice  bool                                                     `json:"do_not_invoice" api:"required"`
	ScheduleItems []V1PackageListResponseCommitInvoiceScheduleScheduleItem `json:"schedule_items" api:"required"`
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
func (r V1PackageListResponseCommitInvoiceSchedule) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCommitInvoiceScheduleScheduleItem struct {
	ID         string                                                           `json:"id" api:"required" format:"uuid"`
	DateOffset V1PackageListResponseCommitInvoiceScheduleScheduleItemDateOffset `json:"date_offset" api:"required"`
	Quantity   float64                                                          `json:"quantity" api:"required"`
	UnitPrice  float64                                                          `json:"unit_price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DateOffset  respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseCommitInvoiceScheduleScheduleItem) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseCommitInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCommitInvoiceScheduleScheduleItemDateOffset struct {
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
func (r V1PackageListResponseCommitInvoiceScheduleScheduleItemDateOffset) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageListResponseCommitInvoiceScheduleScheduleItemDateOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseOverride struct {
	ID                    string                                           `json:"id" api:"required" format:"uuid"`
	OverrideSpecifiers    []V1PackageListResponseOverrideOverrideSpecifier `json:"override_specifiers" api:"required"`
	StartingAtOffset      V1PackageListResponseOverrideStartingAtOffset    `json:"starting_at_offset" api:"required"`
	ApplicableProductTags []string                                         `json:"applicable_product_tags"`
	Duration              V1PackageListResponseOverrideDuration            `json:"duration"`
	Entitled              bool                                             `json:"entitled"`
	IsCommitSpecific      bool                                             `json:"is_commit_specific"`
	Multiplier            float64                                          `json:"multiplier"`
	OverrideTiers         []shared.OverrideTier                            `json:"override_tiers"`
	OverwriteRate         shared.OverwriteRate                             `json:"overwrite_rate"`
	Priority              float64                                          `json:"priority"`
	Product               V1PackageListResponseOverrideProduct             `json:"product"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	Target string `json:"target"`
	// Any of "OVERWRITE", "MULTIPLIER", "TIERED".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		OverrideSpecifiers    respjson.Field
		StartingAtOffset      respjson.Field
		ApplicableProductTags respjson.Field
		Duration              respjson.Field
		Entitled              respjson.Field
		IsCommitSpecific      respjson.Field
		Multiplier            respjson.Field
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
func (r V1PackageListResponseOverride) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseOverrideOverrideSpecifier struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency           string            `json:"billing_frequency"`
	CommitTemplateIDs          []string          `json:"commit_template_ids"`
	PresentationGroupValues    map[string]string `json:"presentation_group_values"`
	PricingGroupValues         map[string]string `json:"pricing_group_values"`
	ProductID                  string            `json:"product_id" format:"uuid"`
	ProductTags                []string          `json:"product_tags"`
	RecurringCommitTemplateIDs []string          `json:"recurring_commit_template_ids"`
	RecurringCreditTemplateIDs []string          `json:"recurring_credit_template_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFrequency           respjson.Field
		CommitTemplateIDs          respjson.Field
		PresentationGroupValues    respjson.Field
		PricingGroupValues         respjson.Field
		ProductID                  respjson.Field
		ProductTags                respjson.Field
		RecurringCommitTemplateIDs respjson.Field
		RecurringCreditTemplateIDs respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseOverrideOverrideSpecifier) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseOverrideOverrideSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseOverrideStartingAtOffset struct {
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
func (r V1PackageListResponseOverrideStartingAtOffset) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseOverrideStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseOverrideDuration struct {
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
func (r V1PackageListResponseOverrideDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseOverrideDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseOverrideProduct struct {
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
func (r V1PackageListResponseOverrideProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseOverrideProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseScheduledCharge struct {
	ID       string                                       `json:"id" api:"required" format:"uuid"`
	Product  V1PackageListResponseScheduledChargeProduct  `json:"product" api:"required"`
	Schedule V1PackageListResponseScheduledChargeSchedule `json:"schedule" api:"required"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	Name         string            `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Product      respjson.Field
		Schedule     respjson.Field
		CustomFields respjson.Field
		Description  respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseScheduledCharge) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseScheduledChargeProduct struct {
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
func (r V1PackageListResponseScheduledChargeProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseScheduledChargeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseScheduledChargeSchedule struct {
	CreditType    shared.CreditTypeData                                      `json:"credit_type" api:"required"`
	ScheduleItems []V1PackageListResponseScheduledChargeScheduleScheduleItem `json:"schedule_items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType    respjson.Field
		ScheduleItems respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseScheduledChargeSchedule) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseScheduledChargeSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseScheduledChargeScheduleScheduleItem struct {
	ID         string                                                             `json:"id" api:"required" format:"uuid"`
	DateOffset V1PackageListResponseScheduledChargeScheduleScheduleItemDateOffset `json:"date_offset" api:"required"`
	Quantity   float64                                                            `json:"quantity" api:"required"`
	UnitPrice  float64                                                            `json:"unit_price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DateOffset  respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseScheduledChargeScheduleScheduleItem) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseScheduledChargeScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseScheduledChargeScheduleScheduleItemDateOffset struct {
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
func (r V1PackageListResponseScheduledChargeScheduleScheduleItemDateOffset) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageListResponseScheduledChargeScheduleScheduleItemDateOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseUsageStatementSchedule struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Frequency   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseUsageStatementSchedule) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseUsageStatementSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseAlias struct {
	Name         string    `json:"name" api:"required"`
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	StartingAt   time.Time `json:"starting_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		EndingBefore respjson.Field
		StartingAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseAlias) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseBillingProvider string

const (
	V1PackageListResponseBillingProviderAwsMarketplace   V1PackageListResponseBillingProvider = "aws_marketplace"
	V1PackageListResponseBillingProviderStripe           V1PackageListResponseBillingProvider = "stripe"
	V1PackageListResponseBillingProviderNetsuite         V1PackageListResponseBillingProvider = "netsuite"
	V1PackageListResponseBillingProviderCustom           V1PackageListResponseBillingProvider = "custom"
	V1PackageListResponseBillingProviderAzureMarketplace V1PackageListResponseBillingProvider = "azure_marketplace"
	V1PackageListResponseBillingProviderQuickbooksOnline V1PackageListResponseBillingProvider = "quickbooks_online"
	V1PackageListResponseBillingProviderWorkday          V1PackageListResponseBillingProvider = "workday"
	V1PackageListResponseBillingProviderGcpMarketplace   V1PackageListResponseBillingProvider = "gcp_marketplace"
	V1PackageListResponseBillingProviderMetronome        V1PackageListResponseBillingProvider = "metronome"
)

type V1PackageListResponseCredit struct {
	ID                    string                                    `json:"id" api:"required" format:"uuid"`
	Product               V1PackageListResponseCreditProduct        `json:"product" api:"required"`
	AccessSchedule        V1PackageListResponseCreditAccessSchedule `json:"access_schedule"`
	ApplicableProductIDs  []string                                  `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                                  `json:"applicable_product_tags"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	Name         string            `json:"name"`
	Priority     float64           `json:"priority"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []shared.CommitSpecifier `json:"specifiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Product               respjson.Field
		AccessSchedule        respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		CustomFields          respjson.Field
		Description           respjson.Field
		Name                  respjson.Field
		Priority              respjson.Field
		RateType              respjson.Field
		Specifiers            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseCredit) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCreditProduct struct {
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
func (r V1PackageListResponseCreditProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseCreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCreditAccessSchedule struct {
	CreditType    shared.CreditTypeData                                   `json:"credit_type" api:"required"`
	ScheduleItems []V1PackageListResponseCreditAccessScheduleScheduleItem `json:"schedule_items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditType    respjson.Field
		ScheduleItems respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseCreditAccessSchedule) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseCreditAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCreditAccessScheduleScheduleItem struct {
	ID               string                                                                `json:"id" api:"required" format:"uuid"`
	Amount           float64                                                               `json:"amount" api:"required"`
	Duration         V1PackageListResponseCreditAccessScheduleScheduleItemDuration         `json:"duration" api:"required"`
	StartingAtOffset V1PackageListResponseCreditAccessScheduleScheduleItemStartingAtOffset `json:"starting_at_offset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Amount           respjson.Field
		Duration         respjson.Field
		StartingAtOffset respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseCreditAccessScheduleScheduleItem) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseCreditAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCreditAccessScheduleScheduleItemDuration struct {
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
func (r V1PackageListResponseCreditAccessScheduleScheduleItemDuration) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageListResponseCreditAccessScheduleScheduleItemDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseCreditAccessScheduleScheduleItemStartingAtOffset struct {
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
func (r V1PackageListResponseCreditAccessScheduleScheduleItemStartingAtOffset) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageListResponseCreditAccessScheduleScheduleItemStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseDeliveryMethod string

const (
	V1PackageListResponseDeliveryMethodDirectToBillingProvider V1PackageListResponseDeliveryMethod = "direct_to_billing_provider"
	V1PackageListResponseDeliveryMethodAwsSqs                  V1PackageListResponseDeliveryMethod = "aws_sqs"
	V1PackageListResponseDeliveryMethodTackle                  V1PackageListResponseDeliveryMethod = "tackle"
	V1PackageListResponseDeliveryMethodAwsSns                  V1PackageListResponseDeliveryMethod = "aws_sns"
)

type V1PackageListResponseDuration struct {
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
func (r V1PackageListResponseDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
// prices automatically. EXPLICIT prioritization requires specifying priorities for
// each multiplier; the one with the lowest priority value will be prioritized
// first.
type V1PackageListResponseMultiplierOverridePrioritization string

const (
	V1PackageListResponseMultiplierOverridePrioritizationLowestMultiplier V1PackageListResponseMultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	V1PackageListResponseMultiplierOverridePrioritizationExplicit         V1PackageListResponseMultiplierOverridePrioritization = "EXPLICIT"
)

type V1PackageListResponseRecurringCommit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V1PackageListResponseRecurringCommitAccessAmount `json:"access_amount" api:"required"`
	// The amount of time each of the created commits will be valid for
	CommitDuration V1PackageListResponseRecurringCommitCommitDuration `json:"commit_duration" api:"required"`
	Priority       float64                                            `json:"priority" api:"required"`
	Product        V1PackageListResponseRecurringCommitProduct        `json:"product" api:"required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type" api:"required"`
	// Offset relative to the contract start date that determines the start time for
	// the first commit
	StartingAtOffset V1PackageListResponseRecurringCommitStartingAtOffset `json:"starting_at_offset" api:"required"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags"`
	Description           string   `json:"description"`
	// Offset relative to the recurring credit start that determines when the contract
	// will stop creating recurring commits. optional
	Duration V1PackageListResponseRecurringCommitDuration `json:"duration"`
	// The amount the customer should be billed for the commit.
	InvoiceAmount V1PackageListResponseRecurringCommitInvoiceAmount `json:"invoice_amount"`
	Name          string                                            `json:"name"`
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
	SubscriptionConfig V1PackageListResponseRecurringCommitSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccessAmount          respjson.Field
		CommitDuration        respjson.Field
		Priority              respjson.Field
		Product               respjson.Field
		RateType              respjson.Field
		StartingAtOffset      respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		Description           respjson.Field
		Duration              respjson.Field
		InvoiceAmount         respjson.Field
		Name                  respjson.Field
		Proration             respjson.Field
		RecurrenceFrequency   respjson.Field
		RolloverFraction      respjson.Field
		Specifiers            respjson.Field
		SubscriptionConfig    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseRecurringCommit) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type V1PackageListResponseRecurringCommitAccessAmount struct {
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
func (r V1PackageListResponseRecurringCommitAccessAmount) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time each of the created commits will be valid for
type V1PackageListResponseRecurringCommitCommitDuration struct {
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
func (r V1PackageListResponseRecurringCommitCommitDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCommitCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseRecurringCommitProduct struct {
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
func (r V1PackageListResponseRecurringCommitProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCommitProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the contract start date that determines the start time for
// the first commit
type V1PackageListResponseRecurringCommitStartingAtOffset struct {
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
func (r V1PackageListResponseRecurringCommitStartingAtOffset) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCommitStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the recurring credit start that determines when the contract
// will stop creating recurring commits. optional
type V1PackageListResponseRecurringCommitDuration struct {
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
func (r V1PackageListResponseRecurringCommitDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount the customer should be billed for the commit.
type V1PackageListResponseRecurringCommitInvoiceAmount struct {
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
func (r V1PackageListResponseRecurringCommitInvoiceAmount) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attach a subscription to the recurring commit/credit.
type V1PackageListResponseRecurringCommitSubscriptionConfig struct {
	// Any of "INDIVIDUAL", "POOLED".
	Allocation              string                                                                        `json:"allocation" api:"required"`
	ApplySeatIncreaseConfig V1PackageListResponseRecurringCommitSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config" api:"required"`
	SubscriptionTemplateID  string                                                                        `json:"subscription_template_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allocation              respjson.Field
		ApplySeatIncreaseConfig respjson.Field
		SubscriptionTemplateID  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseRecurringCommitSubscriptionConfig) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCommitSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseRecurringCommitSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsProrated  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageListResponseRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseRecurringCredit struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V1PackageListResponseRecurringCreditAccessAmount `json:"access_amount" api:"required"`
	// The amount of time each of the created commits will be valid for
	CommitDuration V1PackageListResponseRecurringCreditCommitDuration `json:"commit_duration" api:"required"`
	Priority       float64                                            `json:"priority" api:"required"`
	Product        V1PackageListResponseRecurringCreditProduct        `json:"product" api:"required"`
	// Whether the created commits will use the commit rate or list rate
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type" api:"required"`
	// Offset relative to the contract start date that determines the start time for
	// the first commit
	StartingAtOffset V1PackageListResponseRecurringCreditStartingAtOffset `json:"starting_at_offset" api:"required"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags"`
	Description           string   `json:"description"`
	// Offset relative to the recurring credit start that determines when the contract
	// will stop creating recurring commits. optional
	Duration V1PackageListResponseRecurringCreditDuration `json:"duration"`
	Name     string                                       `json:"name"`
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
	SubscriptionConfig V1PackageListResponseRecurringCreditSubscriptionConfig `json:"subscription_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccessAmount          respjson.Field
		CommitDuration        respjson.Field
		Priority              respjson.Field
		Product               respjson.Field
		RateType              respjson.Field
		StartingAtOffset      respjson.Field
		ApplicableProductIDs  respjson.Field
		ApplicableProductTags respjson.Field
		Description           respjson.Field
		Duration              respjson.Field
		Name                  respjson.Field
		Proration             respjson.Field
		RecurrenceFrequency   respjson.Field
		RolloverFraction      respjson.Field
		Specifiers            respjson.Field
		SubscriptionConfig    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseRecurringCredit) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of commit to grant.
type V1PackageListResponseRecurringCreditAccessAmount struct {
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
func (r V1PackageListResponseRecurringCreditAccessAmount) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The amount of time each of the created commits will be valid for
type V1PackageListResponseRecurringCreditCommitDuration struct {
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
func (r V1PackageListResponseRecurringCreditCommitDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCreditCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseRecurringCreditProduct struct {
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
func (r V1PackageListResponseRecurringCreditProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCreditProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the contract start date that determines the start time for
// the first commit
type V1PackageListResponseRecurringCreditStartingAtOffset struct {
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
func (r V1PackageListResponseRecurringCreditStartingAtOffset) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCreditStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the recurring credit start that determines when the contract
// will stop creating recurring commits. optional
type V1PackageListResponseRecurringCreditDuration struct {
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
func (r V1PackageListResponseRecurringCreditDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCreditDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attach a subscription to the recurring commit/credit.
type V1PackageListResponseRecurringCreditSubscriptionConfig struct {
	// Any of "INDIVIDUAL", "POOLED".
	Allocation              string                                                                        `json:"allocation" api:"required"`
	ApplySeatIncreaseConfig V1PackageListResponseRecurringCreditSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config" api:"required"`
	SubscriptionTemplateID  string                                                                        `json:"subscription_template_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allocation              respjson.Field
		ApplySeatIncreaseConfig respjson.Field
		SubscriptionTemplateID  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseRecurringCreditSubscriptionConfig) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseRecurringCreditSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseRecurringCreditSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsProrated  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PackageListResponseRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type V1PackageListResponseScheduledChargesOnUsageInvoices string

const (
	V1PackageListResponseScheduledChargesOnUsageInvoicesAll V1PackageListResponseScheduledChargesOnUsageInvoices = "ALL"
)

type V1PackageListResponseSubscription struct {
	// Any of "ADVANCE", "ARREARS".
	CollectionSchedule string                                            `json:"collection_schedule" api:"required"`
	Proration          V1PackageListResponseSubscriptionProration        `json:"proration" api:"required"`
	SubscriptionRate   V1PackageListResponseSubscriptionSubscriptionRate `json:"subscription_rate" api:"required"`
	ID                 string                                            `json:"id" format:"uuid"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields     map[string]string                         `json:"custom_fields"`
	Description      string                                    `json:"description"`
	Duration         V1PackageListResponseSubscriptionDuration `json:"duration"`
	FiatCreditTypeID string                                    `json:"fiat_credit_type_id" format:"uuid"`
	InitialQuantity  float64                                   `json:"initial_quantity"`
	Name             string                                    `json:"name"`
	// Determines how the subscription's quantity is controlled. Defaults to
	// QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified
	// directly on the subscription. `initial_quantity` must be provided with this
	// option. Compatible with recurring commits/credits that use POOLED allocation.
	// **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add
	// user_123) to increment and decrement a subscription quantity, rather than
	// directly providing the quantity. You must use a SEAT_BASED subscription to use a
	// linked recurring credit with an allocation per seat. `seat_config` must be
	// provided with this option.
	//
	// Any of "SEAT_BASED", "QUANTITY_ONLY".
	QuantityManagementMode string                                            `json:"quantity_management_mode"`
	SeatConfig             V1PackageListResponseSubscriptionSeatConfig       `json:"seat_config"`
	StartingAtOffset       V1PackageListResponseSubscriptionStartingAtOffset `json:"starting_at_offset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionSchedule     respjson.Field
		Proration              respjson.Field
		SubscriptionRate       respjson.Field
		ID                     respjson.Field
		CustomFields           respjson.Field
		Description            respjson.Field
		Duration               respjson.Field
		FiatCreditTypeID       respjson.Field
		InitialQuantity        respjson.Field
		Name                   respjson.Field
		QuantityManagementMode respjson.Field
		SeatConfig             respjson.Field
		StartingAtOffset       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseSubscription) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseSubscriptionProration struct {
	// Any of "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE".
	InvoiceBehavior string `json:"invoice_behavior" api:"required"`
	IsProrated      bool   `json:"is_prorated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InvoiceBehavior respjson.Field
		IsProrated      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseSubscriptionProration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseSubscriptionProration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseSubscriptionSubscriptionRate struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string                                                   `json:"billing_frequency" api:"required"`
	Product          V1PackageListResponseSubscriptionSubscriptionRateProduct `json:"product" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFrequency respjson.Field
		Product          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListResponseSubscriptionSubscriptionRate) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseSubscriptionSubscriptionRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseSubscriptionSubscriptionRateProduct struct {
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
func (r V1PackageListResponseSubscriptionSubscriptionRateProduct) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseSubscriptionSubscriptionRateProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseSubscriptionDuration struct {
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
func (r V1PackageListResponseSubscriptionDuration) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseSubscriptionDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseSubscriptionSeatConfig struct {
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
func (r V1PackageListResponseSubscriptionSeatConfig) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseSubscriptionSeatConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListResponseSubscriptionStartingAtOffset struct {
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
func (r V1PackageListResponseSubscriptionStartingAtOffset) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListResponseSubscriptionStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageArchiveResponse struct {
	Data shared.ID `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PackageArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListContractsOnPackageResponse struct {
	ContractID   string    `json:"contract_id" api:"required" format:"uuid"`
	CustomerID   string    `json:"customer_id" api:"required" format:"uuid"`
	StartingAt   time.Time `json:"starting_at" api:"required" format:"date-time"`
	ArchivedAt   time.Time `json:"archived_at" format:"date-time"`
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContractID   respjson.Field
		CustomerID   respjson.Field
		StartingAt   respjson.Field
		ArchivedAt   respjson.Field
		EndingBefore respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PackageListContractsOnPackageResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PackageListContractsOnPackageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageNewParams struct {
	Name                string             `json:"name" api:"required"`
	ContractName        param.Opt[string]  `json:"contract_name,omitzero"`
	NetPaymentTermsDays param.Opt[float64] `json:"net_payment_terms_days,omitzero"`
	// Selects the rate card linked to the specified alias as of the contract's start
	// date.
	RateCardAlias param.Opt[string] `json:"rate_card_alias,omitzero"`
	RateCardID    param.Opt[string] `json:"rate_card_id,omitzero" format:"uuid"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey param.Opt[string] `json:"uniqueness_key,omitzero"`
	// Reference this alias when creating a contract. If the same alias is assigned to
	// multiple packages, it will reference the package to which it was most recently
	// assigned. It is not exposed to end customers.
	Aliases []V1PackageNewParamsAlias `json:"aliases,omitzero"`
	// Any of "contract_start_date", "first_billing_period".
	BillingAnchorDate V1PackageNewParamsBillingAnchorDate `json:"billing_anchor_date,omitzero"`
	// Any of "aws_marketplace", "azure_marketplace", "gcp_marketplace", "stripe",
	// "netsuite".
	BillingProvider V1PackageNewParamsBillingProvider `json:"billing_provider,omitzero"`
	Commits         []V1PackageNewParamsCommit        `json:"commits,omitzero"`
	Credits         []V1PackageNewParamsCredit        `json:"credits,omitzero"`
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod V1PackageNewParamsDeliveryMethod `json:"delivery_method,omitzero"`
	Duration       V1PackageNewParamsDuration       `json:"duration,omitzero"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
	// prices automatically. EXPLICIT prioritization requires specifying priorities for
	// each multiplier; the one with the lowest priority value will be prioritized
	// first. If tiered overrides are used, prioritization must be explicit.
	//
	// Any of "LOWEST_MULTIPLIER", "EXPLICIT".
	MultiplierOverridePrioritization     V1PackageNewParamsMultiplierOverridePrioritization `json:"multiplier_override_prioritization,omitzero"`
	ContractOverrides                    []V1PackageNewParamsOverride                       `json:"overrides,omitzero"`
	PrepaidBalanceThresholdConfiguration shared.PrepaidBalanceThresholdConfigurationParam   `json:"prepaid_balance_threshold_configuration,omitzero"`
	RecurringCommits                     []V1PackageNewParamsRecurringCommit                `json:"recurring_commits,omitzero"`
	RecurringCredits                     []V1PackageNewParamsRecurringCredit                `json:"recurring_credits,omitzero"`
	ScheduledCharges                     []V1PackageNewParamsScheduledCharge                `json:"scheduled_charges,omitzero"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	//
	// Any of "ALL".
	ScheduledChargesOnUsageInvoices V1PackageNewParamsScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices,omitzero"`
	SpendThresholdConfiguration     shared.SpendThresholdConfigurationParam           `json:"spend_threshold_configuration,omitzero"`
	Subscriptions                   []V1PackageNewParamsSubscription                  `json:"subscriptions,omitzero"`
	UsageStatementSchedule          V1PackageNewParamsUsageStatementSchedule          `json:"usage_statement_schedule,omitzero"`
	paramObj
}

func (r V1PackageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type V1PackageNewParamsAlias struct {
	Name         string               `json:"name" api:"required"`
	EndingBefore param.Opt[time.Time] `json:"ending_before,omitzero" format:"date-time"`
	StartingAt   param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V1PackageNewParamsAlias) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsAlias
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageNewParamsBillingAnchorDate string

const (
	V1PackageNewParamsBillingAnchorDateContractStartDate  V1PackageNewParamsBillingAnchorDate = "contract_start_date"
	V1PackageNewParamsBillingAnchorDateFirstBillingPeriod V1PackageNewParamsBillingAnchorDate = "first_billing_period"
)

type V1PackageNewParamsBillingProvider string

const (
	V1PackageNewParamsBillingProviderAwsMarketplace   V1PackageNewParamsBillingProvider = "aws_marketplace"
	V1PackageNewParamsBillingProviderAzureMarketplace V1PackageNewParamsBillingProvider = "azure_marketplace"
	V1PackageNewParamsBillingProviderGcpMarketplace   V1PackageNewParamsBillingProvider = "gcp_marketplace"
	V1PackageNewParamsBillingProviderStripe           V1PackageNewParamsBillingProvider = "stripe"
	V1PackageNewParamsBillingProviderNetsuite         V1PackageNewParamsBillingProvider = "netsuite"
)

// The properties AccessSchedule, ProductID, Type are required.
type V1PackageNewParamsCommit struct {
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
	// commits only one schedule item is allowed and amount must match invoice_schedule
	// total.
	AccessSchedule V1PackageNewParamsCommitAccessSchedule `json:"access_schedule,omitzero" api:"required"`
	ProductID      string                                 `json:"product_id" api:"required" format:"uuid"`
	// Any of "PREPAID", "POSTPAID".
	Type string `json:"type,omitzero" api:"required"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Opt[string] `json:"description,omitzero"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Opt[float64] `json:"rollover_fraction,omitzero"`
	// A temporary ID for the commit that can be used to reference the commit for
	// commit specific overrides.
	TemporaryID param.Opt[string] `json:"temporary_id,omitzero"`
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
	// Required for "POSTPAID" commits: the true up invoice will be generated at this
	// time and only one schedule item is allowed; the total must match access_schedule
	// amount. Optional for "PREPAID" commits: if not provided, this will be a
	// "complimentary" commit with no invoice.
	InvoiceSchedule V1PackageNewParamsCommitInvoiceSchedule `json:"invoice_schedule,omitzero"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	paramObj
}

func (r V1PackageNewParamsCommit) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsCommit](
		"type", "PREPAID", "POSTPAID",
	)
	apijson.RegisterFieldValidator[V1PackageNewParamsCommit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
//
// The property ScheduleItems is required.
type V1PackageNewParamsCommitAccessSchedule struct {
	ScheduleItems []V1PackageNewParamsCommitAccessScheduleScheduleItem `json:"schedule_items,omitzero" api:"required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r V1PackageNewParamsCommitAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCommitAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCommitAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Duration, StartingAtOffset are required.
type V1PackageNewParamsCommitAccessScheduleScheduleItem struct {
	Amount float64 `json:"amount" api:"required"`
	// Offset relative to the start of this segment indicating when it should end.
	Duration V1PackageNewParamsCommitAccessScheduleScheduleItemDuration `json:"duration,omitzero" api:"required"`
	// Date relative to the contract start date indicating the start of this schedule
	// segment.
	StartingAtOffset V1PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset `json:"starting_at_offset,omitzero" api:"required"`
	paramObj
}

func (r V1PackageNewParamsCommitAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCommitAccessScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCommitAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the start of this segment indicating when it should end.
//
// The properties Unit, Value are required.
type V1PackageNewParamsCommitAccessScheduleScheduleItemDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsCommitAccessScheduleScheduleItemDuration) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCommitAccessScheduleScheduleItemDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCommitAccessScheduleScheduleItemDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsCommitAccessScheduleScheduleItemDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// Date relative to the contract start date indicating the start of this schedule
// segment.
//
// The properties Unit, Value are required.
type V1PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match access_schedule
// amount. Optional for "PREPAID" commits: if not provided, this will be a
// "complimentary" commit with no invoice.
//
// The property ScheduleItems is required.
type V1PackageNewParamsCommitInvoiceSchedule struct {
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V1PackageNewParamsCommitInvoiceScheduleScheduleItem `json:"schedule_items,omitzero" api:"required"`
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	// If true, this schedule will not generate an invoice.
	DoNotInvoice param.Opt[bool] `json:"do_not_invoice,omitzero"`
	paramObj
}

func (r V1PackageNewParamsCommitInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCommitInvoiceSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DateOffset, Quantity, UnitPrice are required.
type V1PackageNewParamsCommitInvoiceScheduleScheduleItem struct {
	// Date relative to the contract start date.
	DateOffset V1PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset `json:"date_offset,omitzero" api:"required"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity float64 `json:"quantity" api:"required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount.
	UnitPrice float64 `json:"unit_price" api:"required"`
	paramObj
}

func (r V1PackageNewParamsCommitInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCommitInvoiceScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCommitInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Date relative to the contract start date.
//
// The properties Unit, Value are required.
type V1PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// The properties AccessSchedule, ProductID are required.
type V1PackageNewParamsCredit struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule V1PackageNewParamsCreditAccessSchedule `json:"access_schedule,omitzero" api:"required"`
	ProductID      string                                 `json:"product_id" api:"required" format:"uuid"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Opt[string] `json:"description,omitzero"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
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
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	paramObj
}

func (r V1PackageNewParamsCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsCredit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
}

// Schedule for distributing the credit to the customer.
//
// The property ScheduleItems is required.
type V1PackageNewParamsCreditAccessSchedule struct {
	ScheduleItems []V1PackageNewParamsCreditAccessScheduleScheduleItem `json:"schedule_items,omitzero" api:"required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r V1PackageNewParamsCreditAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCreditAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCreditAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Duration, StartingAtOffset are required.
type V1PackageNewParamsCreditAccessScheduleScheduleItem struct {
	Amount float64 `json:"amount" api:"required"`
	// Offset relative to the start of this segment indicating when it should end.
	Duration V1PackageNewParamsCreditAccessScheduleScheduleItemDuration `json:"duration,omitzero" api:"required"`
	// Date relative to the contract start date indicating the start of this schedule
	// segment.
	StartingAtOffset V1PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset `json:"starting_at_offset,omitzero" api:"required"`
	paramObj
}

func (r V1PackageNewParamsCreditAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCreditAccessScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCreditAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the start of this segment indicating when it should end.
//
// The properties Unit, Value are required.
type V1PackageNewParamsCreditAccessScheduleScheduleItemDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsCreditAccessScheduleScheduleItemDuration) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCreditAccessScheduleScheduleItemDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCreditAccessScheduleScheduleItemDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsCreditAccessScheduleScheduleItemDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// Date relative to the contract start date indicating the start of this schedule
// segment.
//
// The properties Unit, Value are required.
type V1PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

type V1PackageNewParamsDeliveryMethod string

const (
	V1PackageNewParamsDeliveryMethodDirectToBillingProvider V1PackageNewParamsDeliveryMethod = "direct_to_billing_provider"
	V1PackageNewParamsDeliveryMethodAwsSqs                  V1PackageNewParamsDeliveryMethod = "aws_sqs"
	V1PackageNewParamsDeliveryMethodTackle                  V1PackageNewParamsDeliveryMethod = "tackle"
	V1PackageNewParamsDeliveryMethodAwsSns                  V1PackageNewParamsDeliveryMethod = "aws_sns"
)

// The properties Unit, Value are required.
type V1PackageNewParamsDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsDuration) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
// prices automatically. EXPLICIT prioritization requires specifying priorities for
// each multiplier; the one with the lowest priority value will be prioritized
// first. If tiered overrides are used, prioritization must be explicit.
type V1PackageNewParamsMultiplierOverridePrioritization string

const (
	V1PackageNewParamsMultiplierOverridePrioritizationLowestMultiplier V1PackageNewParamsMultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	V1PackageNewParamsMultiplierOverridePrioritizationExplicit         V1PackageNewParamsMultiplierOverridePrioritization = "EXPLICIT"
)

// The properties OverrideSpecifiers, StartingAtOffset are required.
type V1PackageNewParamsOverride struct {
	// Specifies which products the override will apply to.
	OverrideSpecifiers []V1PackageNewParamsOverrideOverrideSpecifier `json:"override_specifiers,omitzero" api:"required"`
	// Offset relative to contract start date indicating when the override will start
	// applying (inclusive)
	StartingAtOffset V1PackageNewParamsOverrideStartingAtOffset `json:"starting_at_offset,omitzero" api:"required"`
	Entitled         param.Opt[bool]                            `json:"entitled,omitzero"`
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
	// Offset relative to override start indicating when the override will stop
	// applying (exclusive)
	Duration V1PackageNewParamsOverrideDuration `json:"duration,omitzero"`
	// Required for OVERWRITE type.
	OverwriteRate V1PackageNewParamsOverrideOverwriteRate `json:"overwrite_rate,omitzero"`
	// Indicates whether the override applies to commit rates or list rates. Can only
	// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
	// `"LIST_RATE"`.
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	Target string `json:"target,omitzero"`
	// Required for TIERED type. Must have at least one tier.
	Tiers []V1PackageNewParamsOverrideTier `json:"tiers,omitzero"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	//
	// Any of "OVERWRITE", "MULTIPLIER", "TIERED".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r V1PackageNewParamsOverride) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsOverride](
		"target", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[V1PackageNewParamsOverride](
		"type", "OVERWRITE", "MULTIPLIER", "TIERED",
	)
}

type V1PackageNewParamsOverrideOverrideSpecifier struct {
	// If provided, the override will only apply to the product with the specified ID.
	ProductID param.Opt[string] `json:"product_id,omitzero" format:"uuid"`
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of `product_id`, `product_tags`, `pricing_group_values`, or
	// `presentation_group_values`. If provided, the override will only apply to the
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
	// one of `product_id`, `product_tags`, `pricing_group_values`, or
	// `presentation_group_values`. If provided, the override will only apply to
	// commits created by the specified recurring commit ids.
	RecurringCommitIDs []string `json:"recurring_commit_ids,omitzero"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of `product_id`, `product_tags`, `pricing_group_values`, or
	// `presentation_group_values`. If provided, the override will only apply to
	// credits created by the specified recurring credit ids.
	RecurringCreditIDs []string `json:"recurring_credit_ids,omitzero"`
	paramObj
}

func (r V1PackageNewParamsOverrideOverrideSpecifier) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsOverrideOverrideSpecifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsOverrideOverrideSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsOverrideOverrideSpecifier](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// Offset relative to contract start date indicating when the override will start
// applying (inclusive)
//
// The properties Unit, Value are required.
type V1PackageNewParamsOverrideStartingAtOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsOverrideStartingAtOffset) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsOverrideStartingAtOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsOverrideStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsOverrideStartingAtOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// Offset relative to override start indicating when the override will stop
// applying (exclusive)
//
// The properties Unit, Value are required.
type V1PackageNewParamsOverrideDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsOverrideDuration) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsOverrideDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsOverrideDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsOverrideDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// Required for OVERWRITE type.
//
// The property RateType is required.
type V1PackageNewParamsOverrideOverwriteRate struct {
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

func (r V1PackageNewParamsOverrideOverwriteRate) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsOverrideOverwriteRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsOverrideOverwriteRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsOverrideOverwriteRate](
		"rate_type", "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "TIERED_PERCENTAGE", "CUSTOM",
	)
}

// The property Multiplier is required.
type V1PackageNewParamsOverrideTier struct {
	Multiplier float64            `json:"multiplier" api:"required"`
	Size       param.Opt[float64] `json:"size,omitzero"`
	paramObj
}

func (r V1PackageNewParamsOverrideTier) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsOverrideTier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsOverrideTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessAmount, CommitDuration, Priority, ProductID,
// StartingAtOffset are required.
type V1PackageNewParamsRecurringCommit struct {
	// The amount of commit to grant.
	AccessAmount V1PackageNewParamsRecurringCommitAccessAmount `json:"access_amount,omitzero" api:"required"`
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
	CommitDuration V1PackageNewParamsRecurringCommitCommitDuration `json:"commit_duration,omitzero" api:"required"`
	// Will be passed down to the individual commits
	Priority  float64 `json:"priority" api:"required"`
	ProductID string  `json:"product_id" api:"required" format:"uuid"`
	// Offset relative to the contract start date that determines the start time for
	// the first commit
	StartingAtOffset V1PackageNewParamsRecurringCommitStartingAtOffset `json:"starting_at_offset,omitzero" api:"required"`
	// Will be passed down to the individual commits
	Description param.Opt[string] `json:"description,omitzero"`
	// displayed on invoices. will be passed through to the individual commits
	Name param.Opt[string] `json:"name,omitzero"`
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
	// Offset relative to the recurring credit start that determines when the contract
	// will stop creating recurring commits. optional
	Duration V1PackageNewParamsRecurringCommitDuration `json:"duration,omitzero"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount V1PackageNewParamsRecurringCommitInvoiceAmount `json:"invoice_amount,omitzero"`
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
	SubscriptionConfig V1PackageNewParamsRecurringCommitSubscriptionConfig `json:"subscription_config,omitzero"`
	paramObj
}

func (r V1PackageNewParamsRecurringCommit) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCommit](
		"proration", "NONE", "FIRST", "LAST", "FIRST_AND_LAST",
	)
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCommit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCommit](
		"recurrence_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// The amount of commit to grant.
//
// The properties CreditTypeID, UnitPrice are required.
type V1PackageNewParamsRecurringCommitAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	// This field is required unless a subscription is attached via
	// `subscription_config`.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	paramObj
}

func (r V1PackageNewParamsRecurringCommitAccessAmount) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCommitAccessAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the length of the access schedule for each created commit/credit. The
// value represents the number of units. Unit defaults to "PERIODS", where the
// length of a period is determined by the recurrence_frequency.
//
// The property Value is required.
type V1PackageNewParamsRecurringCommitCommitDuration struct {
	Value float64 `json:"value" api:"required"`
	// Any of "PERIODS".
	Unit string `json:"unit,omitzero"`
	paramObj
}

func (r V1PackageNewParamsRecurringCommitCommitDuration) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCommitCommitDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCommitCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCommitCommitDuration](
		"unit", "PERIODS",
	)
}

// Offset relative to the contract start date that determines the start time for
// the first commit
//
// The properties Unit, Value are required.
type V1PackageNewParamsRecurringCommitStartingAtOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsRecurringCommitStartingAtOffset) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCommitStartingAtOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCommitStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCommitStartingAtOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// Offset relative to the recurring credit start that determines when the contract
// will stop creating recurring commits. optional
//
// The properties Unit, Value are required.
type V1PackageNewParamsRecurringCommitDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsRecurringCommitDuration) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCommitDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCommitDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// The amount the customer should be billed for the commit. Not required.
//
// The properties CreditTypeID, Quantity, UnitPrice are required.
type V1PackageNewParamsRecurringCommitInvoiceAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	Quantity     float64 `json:"quantity" api:"required"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	paramObj
}

func (r V1PackageNewParamsRecurringCommitInvoiceAmount) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCommitInvoiceAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attach a subscription to the recurring commit/credit.
//
// The properties ApplySeatIncreaseConfig, SubscriptionID are required.
type V1PackageNewParamsRecurringCommitSubscriptionConfig struct {
	ApplySeatIncreaseConfig V1PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,omitzero" api:"required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID string `json:"subscription_id" api:"required"`
	// If set to POOLED, allocation added per seat is pooled across the account. If set
	// to INDIVIDUAL, each seat in the subscription will have its own allocation.
	//
	// Any of "INDIVIDUAL", "POOLED".
	Allocation string `json:"allocation,omitzero"`
	paramObj
}

func (r V1PackageNewParamsRecurringCommitSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCommitSubscriptionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCommitSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCommitSubscriptionConfig](
		"allocation", "INDIVIDUAL", "POOLED",
	)
}

// The property IsProrated is required.
type V1PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated" api:"required"`
	paramObj
}

func (r V1PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessAmount, CommitDuration, Priority, ProductID,
// StartingAtOffset are required.
type V1PackageNewParamsRecurringCredit struct {
	// The amount of commit to grant.
	AccessAmount V1PackageNewParamsRecurringCreditAccessAmount `json:"access_amount,omitzero" api:"required"`
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
	CommitDuration V1PackageNewParamsRecurringCreditCommitDuration `json:"commit_duration,omitzero" api:"required"`
	// Will be passed down to the individual commits
	Priority  float64 `json:"priority" api:"required"`
	ProductID string  `json:"product_id" api:"required" format:"uuid"`
	// Offset relative to the contract start date that determines the start time for
	// the first commit
	StartingAtOffset V1PackageNewParamsRecurringCreditStartingAtOffset `json:"starting_at_offset,omitzero" api:"required"`
	// Will be passed down to the individual commits
	Description param.Opt[string] `json:"description,omitzero"`
	// displayed on invoices. will be passed through to the individual commits
	Name param.Opt[string] `json:"name,omitzero"`
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
	// Offset relative to the recurring credit start that determines when the contract
	// will stop creating recurring commits. optional
	Duration V1PackageNewParamsRecurringCreditDuration `json:"duration,omitzero"`
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
	SubscriptionConfig V1PackageNewParamsRecurringCreditSubscriptionConfig `json:"subscription_config,omitzero"`
	paramObj
}

func (r V1PackageNewParamsRecurringCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCredit](
		"proration", "NONE", "FIRST", "LAST", "FIRST_AND_LAST",
	)
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCredit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCredit](
		"recurrence_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// The amount of commit to grant.
//
// The properties CreditTypeID, UnitPrice are required.
type V1PackageNewParamsRecurringCreditAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id" api:"required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price" api:"required"`
	// This field is required unless a subscription is attached via
	// `subscription_config`.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	paramObj
}

func (r V1PackageNewParamsRecurringCreditAccessAmount) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCreditAccessAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the length of the access schedule for each created commit/credit. The
// value represents the number of units. Unit defaults to "PERIODS", where the
// length of a period is determined by the recurrence_frequency.
//
// The property Value is required.
type V1PackageNewParamsRecurringCreditCommitDuration struct {
	Value float64 `json:"value" api:"required"`
	// Any of "PERIODS".
	Unit string `json:"unit,omitzero"`
	paramObj
}

func (r V1PackageNewParamsRecurringCreditCommitDuration) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCreditCommitDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCreditCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCreditCommitDuration](
		"unit", "PERIODS",
	)
}

// Offset relative to the contract start date that determines the start time for
// the first commit
//
// The properties Unit, Value are required.
type V1PackageNewParamsRecurringCreditStartingAtOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsRecurringCreditStartingAtOffset) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCreditStartingAtOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCreditStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCreditStartingAtOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// Offset relative to the recurring credit start that determines when the contract
// will stop creating recurring commits. optional
//
// The properties Unit, Value are required.
type V1PackageNewParamsRecurringCreditDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsRecurringCreditDuration) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCreditDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCreditDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCreditDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// Attach a subscription to the recurring commit/credit.
//
// The properties ApplySeatIncreaseConfig, SubscriptionID are required.
type V1PackageNewParamsRecurringCreditSubscriptionConfig struct {
	ApplySeatIncreaseConfig V1PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,omitzero" api:"required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID string `json:"subscription_id" api:"required"`
	// If set to POOLED, allocation added per seat is pooled across the account. If set
	// to INDIVIDUAL, each seat in the subscription will have its own allocation.
	//
	// Any of "INDIVIDUAL", "POOLED".
	Allocation string `json:"allocation,omitzero"`
	paramObj
}

func (r V1PackageNewParamsRecurringCreditSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCreditSubscriptionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCreditSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsRecurringCreditSubscriptionConfig](
		"allocation", "INDIVIDUAL", "POOLED",
	)
}

// The property IsProrated is required.
type V1PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated" api:"required"`
	paramObj
}

func (r V1PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ProductID, Schedule are required.
type V1PackageNewParamsScheduledCharge struct {
	ProductID string `json:"product_id" api:"required" format:"uuid"`
	// Must provide schedule_items.
	Schedule V1PackageNewParamsScheduledChargeSchedule `json:"schedule,omitzero" api:"required"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r V1PackageNewParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsScheduledCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must provide schedule_items.
//
// The property ScheduleItems is required.
type V1PackageNewParamsScheduledChargeSchedule struct {
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V1PackageNewParamsScheduledChargeScheduleScheduleItem `json:"schedule_items,omitzero" api:"required"`
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r V1PackageNewParamsScheduledChargeSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsScheduledChargeSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsScheduledChargeSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DateOffset, Quantity, UnitPrice are required.
type V1PackageNewParamsScheduledChargeScheduleScheduleItem struct {
	// Date relative to the contract start date.
	DateOffset V1PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset `json:"date_offset,omitzero" api:"required"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity float64 `json:"quantity" api:"required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount.
	UnitPrice float64 `json:"unit_price" api:"required"`
	paramObj
}

func (r V1PackageNewParamsScheduledChargeScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsScheduledChargeScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsScheduledChargeScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Date relative to the contract start date.
//
// The properties Unit, Value are required.
type V1PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type V1PackageNewParamsScheduledChargesOnUsageInvoices string

const (
	V1PackageNewParamsScheduledChargesOnUsageInvoicesAll V1PackageNewParamsScheduledChargesOnUsageInvoices = "ALL"
)

// The properties CollectionSchedule, Proration, SubscriptionRate are required.
type V1PackageNewParamsSubscription struct {
	// Any of "ADVANCE", "ARREARS".
	CollectionSchedule string                                         `json:"collection_schedule,omitzero" api:"required"`
	Proration          V1PackageNewParamsSubscriptionProration        `json:"proration,omitzero" api:"required"`
	SubscriptionRate   V1PackageNewParamsSubscriptionSubscriptionRate `json:"subscription_rate,omitzero" api:"required"`
	Description        param.Opt[string]                              `json:"description,omitzero"`
	// The initial quantity for the subscription. It must be non-negative value.
	// Required if quantity_management_mode is QUANTITY_ONLY.
	InitialQuantity param.Opt[float64] `json:"initial_quantity,omitzero"`
	Name            param.Opt[string]  `json:"name,omitzero"`
	// A temporary ID used to reference the subscription in recurring commit/credit
	// subscription configs created within the same payload.
	TemporaryID param.Opt[string] `json:"temporary_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// Lifetime of the subscription from its start. If not provided, subscription
	// inherits contract end date.
	Duration V1PackageNewParamsSubscriptionDuration `json:"duration,omitzero"`
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
	QuantityManagementMode string                                   `json:"quantity_management_mode,omitzero"`
	SeatConfig             V1PackageNewParamsSubscriptionSeatConfig `json:"seat_config,omitzero"`
	// Relative date from contract start date corresponding to the inclusive start time
	// for the subscription. If not provided, defaults to contract start date
	StartingAtOffset V1PackageNewParamsSubscriptionStartingAtOffset `json:"starting_at_offset,omitzero"`
	paramObj
}

func (r V1PackageNewParamsSubscription) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsSubscription](
		"collection_schedule", "ADVANCE", "ARREARS",
	)
	apijson.RegisterFieldValidator[V1PackageNewParamsSubscription](
		"quantity_management_mode", "SEAT_BASED", "QUANTITY_ONLY",
	)
}

type V1PackageNewParamsSubscriptionProration struct {
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

func (r V1PackageNewParamsSubscriptionProration) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsSubscriptionProration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsSubscriptionProration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsSubscriptionProration](
		"invoice_behavior", "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE",
	)
}

// The properties BillingFrequency, ProductID are required.
type V1PackageNewParamsSubscriptionSubscriptionRate struct {
	// Frequency to bill subscription with. Together with product_id, must match
	// existing rate on the rate card.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero" api:"required"`
	// Must be subscription type product
	ProductID string `json:"product_id" api:"required" format:"uuid"`
	paramObj
}

func (r V1PackageNewParamsSubscriptionSubscriptionRate) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsSubscriptionSubscriptionRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsSubscriptionSubscriptionRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsSubscriptionSubscriptionRate](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// Lifetime of the subscription from its start. If not provided, subscription
// inherits contract end date.
//
// The properties Unit, Value are required.
type V1PackageNewParamsSubscriptionDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsSubscriptionDuration) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsSubscriptionDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsSubscriptionDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsSubscriptionDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// The property SeatGroupKey is required.
type V1PackageNewParamsSubscriptionSeatConfig struct {
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

func (r V1PackageNewParamsSubscriptionSeatConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsSubscriptionSeatConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsSubscriptionSeatConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Relative date from contract start date corresponding to the inclusive start time
// for the subscription. If not provided, defaults to contract start date
//
// The properties Unit, Value are required.
type V1PackageNewParamsSubscriptionStartingAtOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsSubscriptionStartingAtOffset) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsSubscriptionStartingAtOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsSubscriptionStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsSubscriptionStartingAtOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

// The property Frequency is required.
type V1PackageNewParamsUsageStatementSchedule struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency,omitzero" api:"required"`
	// If not provided, defaults to the first day of the month.
	//
	// Any of "FIRST_OF_MONTH", "CONTRACT_START".
	Day string `json:"day,omitzero"`
	// The offset at which Metronome should start generating usage invoices, relative
	// to the contract start date. If unspecified, contract start date will be used.
	// This is useful to set if you want to import historical invoices via our 'Create
	// Historical Invoices' API rather than having Metronome automatically generate
	// them.
	InvoiceGenerationStartingAtOffset V1PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset `json:"invoice_generation_starting_at_offset,omitzero"`
	paramObj
}

func (r V1PackageNewParamsUsageStatementSchedule) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsUsageStatementSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsUsageStatementSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsUsageStatementSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
	apijson.RegisterFieldValidator[V1PackageNewParamsUsageStatementSchedule](
		"day", "FIRST_OF_MONTH", "CONTRACT_START",
	)
}

// The offset at which Metronome should start generating usage invoices, relative
// to the contract start date. If unspecified, contract start date will be used.
// This is useful to set if you want to import historical invoices via our 'Create
// Historical Invoices' API rather than having Metronome automatically generate
// them.
//
// The properties Unit, Value are required.
type V1PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "YEARS".
	Unit  string `json:"unit,omitzero" api:"required"`
	Value int64  `json:"value" api:"required"`
	paramObj
}

func (r V1PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "YEARS",
	)
}

type V1PackageGetParams struct {
	PackageID string `json:"package_id" api:"required" format:"uuid"`
	paramObj
}

func (r V1PackageGetParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListParams struct {
	// The maximum number of packages to return. Defaults to 10.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// Filter packages by archived status. Defaults to NOT_ARCHIVED.
	//
	// Any of "ARCHIVED", "NOT_ARCHIVED", "ALL".
	ArchiveFilter V1PackageListParamsArchiveFilter `json:"archive_filter,omitzero"`
	paramObj
}

func (r V1PackageListParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1PackageListParams]'s query parameters as `url.Values`.
func (r V1PackageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter packages by archived status. Defaults to NOT_ARCHIVED.
type V1PackageListParamsArchiveFilter string

const (
	V1PackageListParamsArchiveFilterArchived    V1PackageListParamsArchiveFilter = "ARCHIVED"
	V1PackageListParamsArchiveFilterNotArchived V1PackageListParamsArchiveFilter = "NOT_ARCHIVED"
	V1PackageListParamsArchiveFilterAll         V1PackageListParamsArchiveFilter = "ALL"
)

type V1PackageArchiveParams struct {
	// ID of the package to archive
	PackageID string `json:"package_id" api:"required" format:"uuid"`
	paramObj
}

func (r V1PackageArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PackageListContractsOnPackageParams struct {
	PackageID string `json:"package_id" api:"required" format:"uuid"`
	// Max number of results that should be returned
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Opt[string] `query:"next_page,omitzero" json:"-"`
	// Optional RFC 3339 timestamp. Only include contracts active on the provided date.
	// This cannot be provided if starting_at filter is provided.
	CoveringDate param.Opt[time.Time] `json:"covering_date,omitzero" format:"date-time"`
	// Default false. Determines whether to include archived contracts in the results
	IncludeArchived param.Opt[bool] `json:"include_archived,omitzero"`
	// Optional RFC 3339 timestamp. Only include contracts that started on or after
	// this date. This cannot be provided if covering_date filter is provided.
	StartingAt param.Opt[time.Time] `json:"starting_at,omitzero" format:"date-time"`
	paramObj
}

func (r V1PackageListContractsOnPackageParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PackageListContractsOnPackageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PackageListContractsOnPackageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1PackageListContractsOnPackageParams]'s query parameters
// as `url.Values`.
func (r V1PackageListContractsOnPackageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
