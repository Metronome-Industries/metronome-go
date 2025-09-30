// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// PackageService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPackageService] method instead.
type PackageService struct {
	Options []option.RequestOption
}

// NewPackageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPackageService(opts ...option.RequestOption) (r PackageService) {
	r = PackageService{}
	r.Options = opts
	return
}

// Create a new package
func (r *PackageService) New(ctx context.Context, body PackageNewParams, opts ...option.RequestOption) (res *PackageNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/packages/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PackageNewResponse struct {
	Data shared.ID `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PackageNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PackageNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PackageNewParams struct {
	Name                string             `json:"name,required"`
	ContractName        param.Opt[string]  `json:"contract_name,omitzero"`
	NetPaymentTermsDays param.Opt[float64] `json:"net_payment_terms_days,omitzero"`
	// Priority of the generated contract.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Selects the rate card linked to the specified alias as of the contract's start
	// date.
	RateCardAlias param.Opt[string] `json:"rate_card_alias,omitzero"`
	RateCardID    param.Opt[string] `json:"rate_card_id,omitzero" format:"uuid"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey param.Opt[string] `json:"uniqueness_key,omitzero"`
	// Any of "contract_start_date", "first_billing_period".
	BillingAnchorDate PackageNewParamsBillingAnchorDate `json:"billing_anchor_date,omitzero"`
	// Any of "aws_marketplace", "azure_marketplace", "gcp_marketplace", "stripe",
	// "netsuite".
	BillingProvider PackageNewParamsBillingProvider `json:"billing_provider,omitzero"`
	Commits         []PackageNewParamsCommit        `json:"commits,omitzero"`
	Credits         []PackageNewParamsCredit        `json:"credits,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod PackageNewParamsDeliveryMethod `json:"delivery_method,omitzero"`
	Duration       PackageNewParamsDuration       `json:"duration,omitzero"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
	// prices automatically. EXPLICIT prioritization requires specifying priorities for
	// each multiplier; the one with the lowest priority value will be prioritized
	// first. If tiered overrides are used, prioritization must be explicit.
	//
	// Any of "LOWEST_MULTIPLIER", "EXPLICIT".
	MultiplierOverridePrioritization     PackageNewParamsMultiplierOverridePrioritization `json:"multiplier_override_prioritization,omitzero"`
	Overrides                            []PackageNewParamsOverride                       `json:"overrides,omitzero"`
	PrepaidBalanceThresholdConfiguration shared.PrepaidBalanceThresholdConfigurationParam `json:"prepaid_balance_threshold_configuration,omitzero"`
	RecurringCommits                     []PackageNewParamsRecurringCommit                `json:"recurring_commits,omitzero"`
	RecurringCredits                     []PackageNewParamsRecurringCredit                `json:"recurring_credits,omitzero"`
	ScheduledCharges                     []PackageNewParamsScheduledCharge                `json:"scheduled_charges,omitzero"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	//
	// Any of "ALL".
	ScheduledChargesOnUsageInvoices PackageNewParamsScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices,omitzero"`
	SpendThresholdConfiguration     shared.SpendThresholdConfigurationParam         `json:"spend_threshold_configuration,omitzero"`
	Subscriptions                   []PackageNewParamsSubscription                  `json:"subscriptions,omitzero"`
	UsageStatementSchedule          PackageNewParamsUsageStatementSchedule          `json:"usage_statement_schedule,omitzero"`
	paramObj
}

func (r PackageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PackageNewParamsBillingAnchorDate string

const (
	PackageNewParamsBillingAnchorDateContractStartDate  PackageNewParamsBillingAnchorDate = "contract_start_date"
	PackageNewParamsBillingAnchorDateFirstBillingPeriod PackageNewParamsBillingAnchorDate = "first_billing_period"
)

type PackageNewParamsBillingProvider string

const (
	PackageNewParamsBillingProviderAwsMarketplace   PackageNewParamsBillingProvider = "aws_marketplace"
	PackageNewParamsBillingProviderAzureMarketplace PackageNewParamsBillingProvider = "azure_marketplace"
	PackageNewParamsBillingProviderGcpMarketplace   PackageNewParamsBillingProvider = "gcp_marketplace"
	PackageNewParamsBillingProviderStripe           PackageNewParamsBillingProvider = "stripe"
	PackageNewParamsBillingProviderNetsuite         PackageNewParamsBillingProvider = "netsuite"
)

// The properties ProductID, Type are required.
type PackageNewParamsCommit struct {
	ProductID string `json:"product_id,required" format:"uuid"`
	// Any of "PREPAID", "POSTPAID".
	Type string `json:"type,omitzero,required"`
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
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
	// commits only one schedule item is allowed and amount must match invoice_schedule
	// total.
	AccessSchedule PackageNewParamsCommitAccessSchedule `json:"access_schedule,omitzero"`
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
	InvoiceSchedule PackageNewParamsCommitInvoiceSchedule `json:"invoice_schedule,omitzero"`
	// optionally payment gate this commit
	PaymentGateConfig PackageNewParamsCommitPaymentGateConfig `json:"payment_gate_config,omitzero"`
	// Any of "COMMIT_RATE", "LIST_RATE".
	RateType string `json:"rate_type,omitzero"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []shared.CommitSpecifierInputParam `json:"specifiers,omitzero"`
	paramObj
}

func (r PackageNewParamsCommit) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsCommit](
		"type", "PREPAID", "POSTPAID",
	)
	apijson.RegisterFieldValidator[PackageNewParamsCommit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
//
// The property ScheduleItems is required.
type PackageNewParamsCommitAccessSchedule struct {
	ScheduleItems []PackageNewParamsCommitAccessScheduleScheduleItem `json:"schedule_items,omitzero,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r PackageNewParamsCommitAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCommitAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCommitAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Duration, StartingAtOffset are required.
type PackageNewParamsCommitAccessScheduleScheduleItem struct {
	Amount float64 `json:"amount,required"`
	// Offset relative to the start of this segment indicating when it should end.
	Duration PackageNewParamsCommitAccessScheduleScheduleItemDuration `json:"duration,omitzero,required"`
	// Date relative to the contract start date indicating the start of this schedule
	// segment.
	StartingAtOffset PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset `json:"starting_at_offset,omitzero,required"`
	paramObj
}

func (r PackageNewParamsCommitAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCommitAccessScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCommitAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the start of this segment indicating when it should end.
//
// The properties Unit, Value are required.
type PackageNewParamsCommitAccessScheduleScheduleItemDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsCommitAccessScheduleScheduleItemDuration) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCommitAccessScheduleScheduleItemDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCommitAccessScheduleScheduleItemDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsCommitAccessScheduleScheduleItemDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// Date relative to the contract start date indicating the start of this schedule
// segment.
//
// The properties Unit, Value are required.
type PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsCommitAccessScheduleScheduleItemStartingAtOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match access_schedule
// amount. Optional for "PREPAID" commits: if not provided, this will be a
// "complimentary" commit with no invoice.
//
// The property ScheduleItems is required.
type PackageNewParamsCommitInvoiceSchedule struct {
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []PackageNewParamsCommitInvoiceScheduleScheduleItem `json:"schedule_items,omitzero,required"`
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r PackageNewParamsCommitInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCommitInvoiceSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCommitInvoiceSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property DateOffset is required.
type PackageNewParamsCommitInvoiceScheduleScheduleItem struct {
	// Date relative to the contract start date.
	DateOffset PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset `json:"date_offset,omitzero,required"`
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

func (r PackageNewParamsCommitInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCommitInvoiceScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCommitInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Date relative to the contract start date.
//
// The properties Unit, Value are required.
type PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsCommitInvoiceScheduleScheduleItemDateOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// optionally payment gate this commit
//
// The property PaymentGateType is required.
type PackageNewParamsCommitPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	//
	// Any of "NONE", "STRIPE", "EXTERNAL".
	PaymentGateType string `json:"payment_gate_type,omitzero,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig PackageNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config,omitzero"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig PackageNewParamsCommitPaymentGateConfigStripeConfig `json:"stripe_config,omitzero"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	//
	// Any of "NONE", "STRIPE", "ANROK", "AVALARA", "PRECALCULATED".
	TaxType string `json:"tax_type,omitzero"`
	paramObj
}

func (r PackageNewParamsCommitPaymentGateConfig) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCommitPaymentGateConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCommitPaymentGateConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsCommitPaymentGateConfig](
		"payment_gate_type", "NONE", "STRIPE", "EXTERNAL",
	)
	apijson.RegisterFieldValidator[PackageNewParamsCommitPaymentGateConfig](
		"tax_type", "NONE", "STRIPE", "ANROK", "AVALARA", "PRECALCULATED",
	)
}

// Only applicable if using PRECALCULATED as your tax type.
//
// The property TaxAmount is required.
type PackageNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Opt[string] `json:"tax_name,omitzero"`
	paramObj
}

func (r PackageNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCommitPaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only applicable if using STRIPE as your payment gate type.
//
// The property PaymentType is required.
type PackageNewParamsCommitPaymentGateConfigStripeConfig struct {
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

func (r PackageNewParamsCommitPaymentGateConfigStripeConfig) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCommitPaymentGateConfigStripeConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCommitPaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsCommitPaymentGateConfigStripeConfig](
		"payment_type", "INVOICE", "PAYMENT_INTENT",
	)
}

// The properties AccessSchedule, ProductID are required.
type PackageNewParamsCredit struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule PackageNewParamsCreditAccessSchedule `json:"access_schedule,omitzero,required"`
	ProductID      string                               `json:"product_id,required" format:"uuid"`
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

func (r PackageNewParamsCredit) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsCredit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
}

// Schedule for distributing the credit to the customer.
//
// The property ScheduleItems is required.
type PackageNewParamsCreditAccessSchedule struct {
	ScheduleItems []PackageNewParamsCreditAccessScheduleScheduleItem `json:"schedule_items,omitzero,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r PackageNewParamsCreditAccessSchedule) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCreditAccessSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCreditAccessSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Duration, StartingAtOffset are required.
type PackageNewParamsCreditAccessScheduleScheduleItem struct {
	Amount float64 `json:"amount,required"`
	// Offset relative to the start of this segment indicating when it should end.
	Duration PackageNewParamsCreditAccessScheduleScheduleItemDuration `json:"duration,omitzero,required"`
	// Date relative to the contract start date indicating the start of this schedule
	// segment.
	StartingAtOffset PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset `json:"starting_at_offset,omitzero,required"`
	paramObj
}

func (r PackageNewParamsCreditAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCreditAccessScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCreditAccessScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Offset relative to the start of this segment indicating when it should end.
//
// The properties Unit, Value are required.
type PackageNewParamsCreditAccessScheduleScheduleItemDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsCreditAccessScheduleScheduleItemDuration) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCreditAccessScheduleScheduleItemDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCreditAccessScheduleScheduleItemDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsCreditAccessScheduleScheduleItemDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// Date relative to the contract start date indicating the start of this schedule
// segment.
//
// The properties Unit, Value are required.
type PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsCreditAccessScheduleScheduleItemStartingAtOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

type PackageNewParamsDeliveryMethod string

const (
	PackageNewParamsDeliveryMethodDirectToBillingProvider PackageNewParamsDeliveryMethod = "direct_to_billing_provider"
	PackageNewParamsDeliveryMethodAwsSqs                  PackageNewParamsDeliveryMethod = "aws_sqs"
	PackageNewParamsDeliveryMethodTackle                  PackageNewParamsDeliveryMethod = "tackle"
	PackageNewParamsDeliveryMethodAwsSns                  PackageNewParamsDeliveryMethod = "aws_sns"
)

// The properties Unit, Value are required.
type PackageNewParamsDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsDuration) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
// prices automatically. EXPLICIT prioritization requires specifying priorities for
// each multiplier; the one with the lowest priority value will be prioritized
// first. If tiered overrides are used, prioritization must be explicit.
type PackageNewParamsMultiplierOverridePrioritization string

const (
	PackageNewParamsMultiplierOverridePrioritizationLowestMultiplier PackageNewParamsMultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	PackageNewParamsMultiplierOverridePrioritizationExplicit         PackageNewParamsMultiplierOverridePrioritization = "EXPLICIT"
)

// The properties OverrideSpecifiers, StartingOffset are required.
type PackageNewParamsOverride struct {
	// Specifies which products the override will apply to.
	OverrideSpecifiers []PackageNewParamsOverrideOverrideSpecifier `json:"override_specifiers,omitzero,required"`
	// Offset relative to contract start date indicating when the override will start
	// applying (inclusive)
	StartingOffset PackageNewParamsOverrideStartingOffset `json:"starting_offset,omitzero,required"`
	Entitled       param.Opt[bool]                        `json:"entitled,omitzero"`
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
	Duration PackageNewParamsOverrideDuration `json:"duration,omitzero"`
	// Required for OVERWRITE type.
	OverwriteRate PackageNewParamsOverrideOverwriteRate `json:"overwrite_rate,omitzero"`
	// Indicates whether the override applies to commit rates or list rates. Can only
	// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
	// `"LIST_RATE"`.
	//
	// Any of "COMMIT_RATE", "LIST_RATE".
	Target string `json:"target,omitzero"`
	// Required for TIERED type. Must have at least one tier.
	Tiers []PackageNewParamsOverrideTier `json:"tiers,omitzero"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	//
	// Any of "OVERWRITE", "MULTIPLIER", "TIERED".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r PackageNewParamsOverride) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsOverride](
		"target", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[PackageNewParamsOverride](
		"type", "OVERWRITE", "MULTIPLIER", "TIERED",
	)
}

type PackageNewParamsOverrideOverrideSpecifier struct {
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

func (r PackageNewParamsOverrideOverrideSpecifier) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsOverrideOverrideSpecifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsOverrideOverrideSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsOverrideOverrideSpecifier](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// Offset relative to contract start date indicating when the override will start
// applying (inclusive)
//
// The properties Unit, Value are required.
type PackageNewParamsOverrideStartingOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsOverrideStartingOffset) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsOverrideStartingOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsOverrideStartingOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsOverrideStartingOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// Offset relative to override start indicating when the override will stop
// applying (exclusive)
//
// The properties Unit, Value are required.
type PackageNewParamsOverrideDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsOverrideDuration) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsOverrideDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsOverrideDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsOverrideDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// Required for OVERWRITE type.
//
// The property RateType is required.
type PackageNewParamsOverrideOverwriteRate struct {
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

func (r PackageNewParamsOverrideOverwriteRate) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsOverrideOverwriteRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsOverrideOverwriteRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsOverrideOverwriteRate](
		"rate_type", "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM",
	)
}

// The property Multiplier is required.
type PackageNewParamsOverrideTier struct {
	Multiplier float64            `json:"multiplier,required"`
	Size       param.Opt[float64] `json:"size,omitzero"`
	paramObj
}

func (r PackageNewParamsOverrideTier) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsOverrideTier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsOverrideTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessAmount, CommitDuration, Priority, ProductID, StartingOffset
// are required.
type PackageNewParamsRecurringCommit struct {
	// The amount of commit to grant.
	AccessAmount PackageNewParamsRecurringCommitAccessAmount `json:"access_amount,omitzero,required"`
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
	CommitDuration PackageNewParamsRecurringCommitCommitDuration `json:"commit_duration,omitzero,required"`
	// Will be passed down to the individual commits
	Priority  float64 `json:"priority,required"`
	ProductID string  `json:"product_id,required" format:"uuid"`
	// Determines the start time for the first commit
	StartingOffset PackageNewParamsRecurringCommitStartingOffset `json:"starting_offset,omitzero,required"`
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
	// Determines when the contract will stop creating recurring commits. optional
	EndingOffset PackageNewParamsRecurringCommitEndingOffset `json:"ending_offset,omitzero"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount PackageNewParamsRecurringCommitInvoiceAmount `json:"invoice_amount,omitzero"`
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
	SubscriptionConfig PackageNewParamsRecurringCommitSubscriptionConfig `json:"subscription_config,omitzero"`
	paramObj
}

func (r PackageNewParamsRecurringCommit) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCommit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCommit](
		"proration", "NONE", "FIRST", "LAST", "FIRST_AND_LAST",
	)
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCommit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCommit](
		"recurrence_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// The amount of commit to grant.
//
// The properties CreditTypeID, UnitPrice are required.
type PackageNewParamsRecurringCommitAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
	// This field is required unless a subscription is attached via
	// `subscription_config`.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	paramObj
}

func (r PackageNewParamsRecurringCommitAccessAmount) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCommitAccessAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCommitAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the length of the access schedule for each created commit/credit. The
// value represents the number of units. Unit defaults to "PERIODS", where the
// length of a period is determined by the recurrence_frequency.
//
// The property Value is required.
type PackageNewParamsRecurringCommitCommitDuration struct {
	Value float64 `json:"value,required"`
	// Any of "PERIODS".
	Unit string `json:"unit,omitzero"`
	paramObj
}

func (r PackageNewParamsRecurringCommitCommitDuration) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCommitCommitDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCommitCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCommitCommitDuration](
		"unit", "PERIODS",
	)
}

// Determines the start time for the first commit
//
// The properties Unit, Value are required.
type PackageNewParamsRecurringCommitStartingOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsRecurringCommitStartingOffset) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCommitStartingOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCommitStartingOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCommitStartingOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// Determines when the contract will stop creating recurring commits. optional
//
// The properties Unit, Value are required.
type PackageNewParamsRecurringCommitEndingOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsRecurringCommitEndingOffset) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCommitEndingOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCommitEndingOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCommitEndingOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// The amount the customer should be billed for the commit. Not required.
//
// The properties CreditTypeID, Quantity, UnitPrice are required.
type PackageNewParamsRecurringCommitInvoiceAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64 `json:"quantity,required"`
	UnitPrice    float64 `json:"unit_price,required"`
	paramObj
}

func (r PackageNewParamsRecurringCommitInvoiceAmount) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCommitInvoiceAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCommitInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attach a subscription to the recurring commit/credit.
//
// The properties ApplySeatIncreaseConfig, SubscriptionID are required.
type PackageNewParamsRecurringCommitSubscriptionConfig struct {
	ApplySeatIncreaseConfig PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,omitzero,required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID string `json:"subscription_id,required"`
	// If set to POOLED, allocation added per seat is pooled across the account.
	//
	// Any of "INDIVIDUAL", "POOLED".
	Allocation string `json:"allocation,omitzero"`
	paramObj
}

func (r PackageNewParamsRecurringCommitSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCommitSubscriptionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCommitSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCommitSubscriptionConfig](
		"allocation", "INDIVIDUAL", "POOLED",
	)
}

// The property IsProrated is required.
type PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated,required"`
	paramObj
}

func (r PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCommitSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessAmount, CommitDuration, Priority, ProductID, StartingOffset
// are required.
type PackageNewParamsRecurringCredit struct {
	// The amount of commit to grant.
	AccessAmount PackageNewParamsRecurringCreditAccessAmount `json:"access_amount,omitzero,required"`
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
	CommitDuration PackageNewParamsRecurringCreditCommitDuration `json:"commit_duration,omitzero,required"`
	// Will be passed down to the individual commits
	Priority  float64 `json:"priority,required"`
	ProductID string  `json:"product_id,required" format:"uuid"`
	// Determines the start time for the first commit
	StartingOffset PackageNewParamsRecurringCreditStartingOffset `json:"starting_offset,omitzero,required"`
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
	// Determines when the contract will stop creating recurring commits. optional
	EndingOffset PackageNewParamsRecurringCreditEndingOffset `json:"ending_offset,omitzero"`
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
	SubscriptionConfig PackageNewParamsRecurringCreditSubscriptionConfig `json:"subscription_config,omitzero"`
	paramObj
}

func (r PackageNewParamsRecurringCredit) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCredit](
		"proration", "NONE", "FIRST", "LAST", "FIRST_AND_LAST",
	)
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCredit](
		"rate_type", "COMMIT_RATE", "LIST_RATE",
	)
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCredit](
		"recurrence_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// The amount of commit to grant.
//
// The properties CreditTypeID, UnitPrice are required.
type PackageNewParamsRecurringCreditAccessAmount struct {
	CreditTypeID string  `json:"credit_type_id,required" format:"uuid"`
	UnitPrice    float64 `json:"unit_price,required"`
	// This field is required unless a subscription is attached via
	// `subscription_config`.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	paramObj
}

func (r PackageNewParamsRecurringCreditAccessAmount) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCreditAccessAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCreditAccessAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the length of the access schedule for each created commit/credit. The
// value represents the number of units. Unit defaults to "PERIODS", where the
// length of a period is determined by the recurrence_frequency.
//
// The property Value is required.
type PackageNewParamsRecurringCreditCommitDuration struct {
	Value float64 `json:"value,required"`
	// Any of "PERIODS".
	Unit string `json:"unit,omitzero"`
	paramObj
}

func (r PackageNewParamsRecurringCreditCommitDuration) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCreditCommitDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCreditCommitDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCreditCommitDuration](
		"unit", "PERIODS",
	)
}

// Determines the start time for the first commit
//
// The properties Unit, Value are required.
type PackageNewParamsRecurringCreditStartingOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsRecurringCreditStartingOffset) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCreditStartingOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCreditStartingOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCreditStartingOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// Determines when the contract will stop creating recurring commits. optional
//
// The properties Unit, Value are required.
type PackageNewParamsRecurringCreditEndingOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsRecurringCreditEndingOffset) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCreditEndingOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCreditEndingOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCreditEndingOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// Attach a subscription to the recurring commit/credit.
//
// The properties ApplySeatIncreaseConfig, SubscriptionID are required.
type PackageNewParamsRecurringCreditSubscriptionConfig struct {
	ApplySeatIncreaseConfig PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,omitzero,required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID string `json:"subscription_id,required"`
	// If set to POOLED, allocation added per seat is pooled across the account.
	//
	// Any of "INDIVIDUAL", "POOLED".
	Allocation string `json:"allocation,omitzero"`
	paramObj
}

func (r PackageNewParamsRecurringCreditSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCreditSubscriptionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCreditSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsRecurringCreditSubscriptionConfig](
		"allocation", "INDIVIDUAL", "POOLED",
	)
}

// The property IsProrated is required.
type PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool `json:"is_prorated,required"`
	paramObj
}

func (r PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsRecurringCreditSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ProductID, Schedule are required.
type PackageNewParamsScheduledCharge struct {
	ProductID string `json:"product_id,required" format:"uuid"`
	// Must provide schedule_items.
	Schedule PackageNewParamsScheduledChargeSchedule `json:"schedule,omitzero,required"`
	// displayed on invoices
	Name param.Opt[string] `json:"name,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	paramObj
}

func (r PackageNewParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsScheduledCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsScheduledCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must provide schedule_items.
//
// The property ScheduleItems is required.
type PackageNewParamsScheduledChargeSchedule struct {
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []PackageNewParamsScheduledChargeScheduleScheduleItem `json:"schedule_items,omitzero,required"`
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Opt[string] `json:"credit_type_id,omitzero" format:"uuid"`
	paramObj
}

func (r PackageNewParamsScheduledChargeSchedule) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsScheduledChargeSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsScheduledChargeSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property DateOffset is required.
type PackageNewParamsScheduledChargeScheduleScheduleItem struct {
	// Date relative to the contract start date.
	DateOffset PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset `json:"date_offset,omitzero,required"`
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

func (r PackageNewParamsScheduledChargeScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsScheduledChargeScheduleScheduleItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsScheduledChargeScheduleScheduleItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Date relative to the contract start date.
//
// The properties Unit, Value are required.
type PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsScheduledChargeScheduleScheduleItemDateOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type PackageNewParamsScheduledChargesOnUsageInvoices string

const (
	PackageNewParamsScheduledChargesOnUsageInvoicesAll PackageNewParamsScheduledChargesOnUsageInvoices = "ALL"
)

// The properties CollectionSchedule, InitialQuantity, Proration, SubscriptionRate
// are required.
type PackageNewParamsSubscription struct {
	// Any of "ADVANCE", "ARREARS".
	CollectionSchedule string `json:"collection_schedule,omitzero,required"`
	// The initial quantity for the subscription. It must be non-negative value.
	InitialQuantity  float64                                      `json:"initial_quantity,required"`
	Proration        PackageNewParamsSubscriptionProration        `json:"proration,omitzero,required"`
	SubscriptionRate PackageNewParamsSubscriptionSubscriptionRate `json:"subscription_rate,omitzero,required"`
	Description      param.Opt[string]                            `json:"description,omitzero"`
	Name             param.Opt[string]                            `json:"name,omitzero"`
	// A temporary ID used to reference the subscription in recurring commit/credit
	// subscription configs created within the same payload.
	TemporaryID param.Opt[string] `json:"temporary_id,omitzero"`
	// Custom fields to be added eg. { "key1": "value1", "key2": "value2" }
	CustomFields map[string]string `json:"custom_fields,omitzero"`
	// Lifetime of the subscription from its start. If not provided, subscription
	// inherits contract end date.
	Duration PackageNewParamsSubscriptionDuration `json:"duration,omitzero"`
	// Relative date from contract start date corresponding to the inclusive start time
	// for the subscription. If not provided, defaults to contract start date
	StartingOffset PackageNewParamsSubscriptionStartingOffset `json:"starting_offset,omitzero"`
	paramObj
}

func (r PackageNewParamsSubscription) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsSubscription](
		"collection_schedule", "ADVANCE", "ARREARS",
	)
}

type PackageNewParamsSubscriptionProration struct {
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

func (r PackageNewParamsSubscriptionProration) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsSubscriptionProration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsSubscriptionProration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsSubscriptionProration](
		"invoice_behavior", "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE",
	)
}

// The properties BillingFrequency, ProductID are required.
type PackageNewParamsSubscriptionSubscriptionRate struct {
	// Frequency to bill subscription with. Together with product_id, must match
	// existing rate on the rate card.
	//
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	BillingFrequency string `json:"billing_frequency,omitzero,required"`
	// Must be subscription type product
	ProductID string `json:"product_id,required" format:"uuid"`
	paramObj
}

func (r PackageNewParamsSubscriptionSubscriptionRate) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsSubscriptionSubscriptionRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsSubscriptionSubscriptionRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsSubscriptionSubscriptionRate](
		"billing_frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
}

// Lifetime of the subscription from its start. If not provided, subscription
// inherits contract end date.
//
// The properties Unit, Value are required.
type PackageNewParamsSubscriptionDuration struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsSubscriptionDuration) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsSubscriptionDuration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsSubscriptionDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsSubscriptionDuration](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// Relative date from contract start date corresponding to the inclusive start time
// for the subscription. If not provided, defaults to contract start date
//
// The properties Unit, Value are required.
type PackageNewParamsSubscriptionStartingOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsSubscriptionStartingOffset) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsSubscriptionStartingOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsSubscriptionStartingOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsSubscriptionStartingOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}

// The property Frequency is required.
type PackageNewParamsUsageStatementSchedule struct {
	// Any of "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".
	Frequency string `json:"frequency,omitzero,required"`
	// If not provided, defaults to the first day of the month.
	//
	// Any of "FIRST_OF_MONTH", "CONTRACT_START".
	Day string `json:"day,omitzero"`
	// The offset at which Metronome should start generating usage invoices, relative
	// to the contract start date. If unspecified, contract start date will be used.
	// This is useful to set if you want to import historical invoices via our 'Create
	// Historical Invoices' API rather than having Metronome automatically generate
	// them.
	InvoiceGenerationStartingAtOffset PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset `json:"invoice_generation_starting_at_offset,omitzero"`
	paramObj
}

func (r PackageNewParamsUsageStatementSchedule) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsUsageStatementSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsUsageStatementSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsUsageStatementSchedule](
		"frequency", "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY",
	)
	apijson.RegisterFieldValidator[PackageNewParamsUsageStatementSchedule](
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
type PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset struct {
	// Any of "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS".
	Unit  string `json:"unit,omitzero,required"`
	Value int64  `json:"value,required"`
	paramObj
}

func (r PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset) MarshalJSON() (data []byte, err error) {
	type shadow PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PackageNewParamsUsageStatementScheduleInvoiceGenerationStartingAtOffset](
		"unit", "DAYS", "WEEKS", "MONTHS", "QUARTERS", "YEARS",
	)
}
