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
func NewV2ContractService(opts ...option.RequestOption) (r *V2ContractService) {
	r = &V2ContractService{}
	r.Options = opts
	return
}

// Get a specific contract. New clients should use this endpoint rather than the v1
// endpoint.
func (r *V2ContractService) Get(ctx context.Context, body V2ContractGetParams, opts ...option.RequestOption) (res *V2ContractGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/contracts/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all contracts for a customer. New clients should use this endpoint rather
// than the v1 endpoint.
func (r *V2ContractService) List(ctx context.Context, body V2ContractListParams, opts ...option.RequestOption) (res *V2ContractListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/contracts/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Edit a contract. Contract editing must be enabled to use this endpoint.
func (r *V2ContractService) Edit(ctx context.Context, body V2ContractEditParams, opts ...option.RequestOption) (res *V2ContractEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/contracts/edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Edit a customer or contract commit. Contract commits can only be edited using
// this endpoint if contract editing is enabled.
func (r *V2ContractService) EditCommit(ctx context.Context, body V2ContractEditCommitParams, opts ...option.RequestOption) (res *V2ContractEditCommitResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/contracts/commits/edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Edit a customer or contract credit. Contract credits can only be edited using
// this endpoint if contract editing is enabled.
func (r *V2ContractService) EditCredit(ctx context.Context, body V2ContractEditCreditParams, opts ...option.RequestOption) (res *V2ContractEditCreditResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/contracts/credits/edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the edit history of a specific contract. Contract editing must be enabled to
// use this endpoint.
func (r *V2ContractService) GetEditHistory(ctx context.Context, body V2ContractGetEditHistoryParams, opts ...option.RequestOption) (res *V2ContractGetEditHistoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/contracts/getEditHistory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V2ContractGetResponse struct {
	Data V2ContractGetResponseData `json:"data,required"`
	JSON v2ContractGetResponseJSON `json:"-"`
}

// v2ContractGetResponseJSON contains the JSON metadata for the struct
// [V2ContractGetResponse]
type v2ContractGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseData struct {
	ID                     string                                          `json:"id,required" format:"uuid"`
	Commits                []V2ContractGetResponseDataCommit               `json:"commits,required"`
	CreatedAt              time.Time                                       `json:"created_at,required" format:"date-time"`
	CreatedBy              string                                          `json:"created_by,required"`
	CustomerID             string                                          `json:"customer_id,required" format:"uuid"`
	Overrides              []V2ContractGetResponseDataOverride             `json:"overrides,required"`
	ScheduledCharges       []shared.ScheduledCharge                        `json:"scheduled_charges,required"`
	StartingAt             time.Time                                       `json:"starting_at,required" format:"date-time"`
	Transitions            []V2ContractGetResponseDataTransition           `json:"transitions,required"`
	UsageFilter            []V2ContractGetResponseDataUsageFilter          `json:"usage_filter,required"`
	UsageStatementSchedule V2ContractGetResponseDataUsageStatementSchedule `json:"usage_statement_schedule,required"`
	ArchivedAt             time.Time                                       `json:"archived_at" format:"date-time"`
	Credits                []V2ContractGetResponseDataCredit               `json:"credits"`
	CustomFields           map[string]string                               `json:"custom_fields"`
	// This field's availability is dependent on your client's configuration.
	CustomerBillingProviderConfiguration V2ContractGetResponseDataCustomerBillingProviderConfiguration `json:"customer_billing_provider_configuration"`
	// This field's availability is dependent on your client's configuration.
	Discounts    []shared.Discount `json:"discounts"`
	EndingBefore time.Time         `json:"ending_before" format:"date-time"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
	// prices automatically. EXPLICIT prioritization requires specifying priorities for
	// each multiplier; the one with the lowest priority value will be prioritized
	// first.
	MultiplierOverridePrioritization V2ContractGetResponseDataMultiplierOverridePrioritization `json:"multiplier_override_prioritization"`
	Name                             string                                                    `json:"name"`
	NetPaymentTermsDays              float64                                                   `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []shared.ProService                        `json:"professional_services"`
	RateCardID           string                                     `json:"rate_card_id" format:"uuid"`
	RecurringCommits     []V2ContractGetResponseDataRecurringCommit `json:"recurring_commits"`
	RecurringCredits     []V2ContractGetResponseDataRecurringCredit `json:"recurring_credits"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []V2ContractGetResponseDataResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices V2ContractGetResponseDataScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	ThresholdBillingConfiguration   V2ContractGetResponseDataThresholdBillingConfiguration   `json:"threshold_billing_configuration"`
	TotalContractValue              float64                                                  `json:"total_contract_value"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                        `json:"uniqueness_key"`
	JSON          v2ContractGetResponseDataJSON `json:"-"`
}

// v2ContractGetResponseDataJSON contains the JSON metadata for the struct
// [V2ContractGetResponseData]
type v2ContractGetResponseDataJSON struct {
	ID                                   apijson.Field
	Commits                              apijson.Field
	CreatedAt                            apijson.Field
	CreatedBy                            apijson.Field
	CustomerID                           apijson.Field
	Overrides                            apijson.Field
	ScheduledCharges                     apijson.Field
	StartingAt                           apijson.Field
	Transitions                          apijson.Field
	UsageFilter                          apijson.Field
	UsageStatementSchedule               apijson.Field
	ArchivedAt                           apijson.Field
	Credits                              apijson.Field
	CustomFields                         apijson.Field
	CustomerBillingProviderConfiguration apijson.Field
	Discounts                            apijson.Field
	EndingBefore                         apijson.Field
	MultiplierOverridePrioritization     apijson.Field
	Name                                 apijson.Field
	NetPaymentTermsDays                  apijson.Field
	NetsuiteSalesOrderID                 apijson.Field
	ProfessionalServices                 apijson.Field
	RateCardID                           apijson.Field
	RecurringCommits                     apijson.Field
	RecurringCredits                     apijson.Field
	ResellerRoyalties                    apijson.Field
	SalesforceOpportunityID              apijson.Field
	ScheduledChargesOnUsageInvoices      apijson.Field
	ThresholdBillingConfiguration        apijson.Field
	TotalContractValue                   apijson.Field
	UniquenessKey                        apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *V2ContractGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataCommit struct {
	ID      string                                  `json:"id,required" format:"uuid"`
	Product V2ContractGetResponseDataCommitsProduct `json:"product,required"`
	Type    V2ContractGetResponseDataCommitsType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule        shared.ScheduleDuration `json:"access_schedule"`
	ApplicableContractIDs []string                `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string                `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64                                  `json:"balance"`
	Contract     V2ContractGetResponseDataCommitsContract `json:"contract"`
	CustomFields map[string]string                        `json:"custom_fields"`
	Description  string                                   `json:"description"`
	// The contract that this commit will be billed on.
	InvoiceContract V2ContractGetResponseDataCommitsInvoiceContract `json:"invoice_contract"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule shared.SchedulePointInTime `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger []V2ContractGetResponseDataCommitsLedger `json:"ledger"`
	Name   string                                   `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority         float64                                        `json:"priority"`
	RateType         V2ContractGetResponseDataCommitsRateType       `json:"rate_type"`
	RolledOverFrom   V2ContractGetResponseDataCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction float64                                        `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                              `json:"salesforce_opportunity_id"`
	JSON                    v2ContractGetResponseDataCommitJSON `json:"-"`
}

// v2ContractGetResponseDataCommitJSON contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommit]
type v2ContractGetResponseDataCommitJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
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
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataCommitsProduct struct {
	ID   string                                      `json:"id,required" format:"uuid"`
	Name string                                      `json:"name,required"`
	JSON v2ContractGetResponseDataCommitsProductJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsProductJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataCommitsProduct]
type v2ContractGetResponseDataCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsProductJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataCommitsType string

const (
	V2ContractGetResponseDataCommitsTypePrepaid  V2ContractGetResponseDataCommitsType = "PREPAID"
	V2ContractGetResponseDataCommitsTypePostpaid V2ContractGetResponseDataCommitsType = "POSTPAID"
)

func (r V2ContractGetResponseDataCommitsType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsTypePrepaid, V2ContractGetResponseDataCommitsTypePostpaid:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsContract struct {
	ID   string                                       `json:"id,required" format:"uuid"`
	JSON v2ContractGetResponseDataCommitsContractJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsContractJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataCommitsContract]
type v2ContractGetResponseDataCommitsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsContractJSON) RawJSON() string {
	return r.raw
}

// The contract that this commit will be billed on.
type V2ContractGetResponseDataCommitsInvoiceContract struct {
	ID   string                                              `json:"id,required" format:"uuid"`
	JSON v2ContractGetResponseDataCommitsInvoiceContractJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsInvoiceContractJSON contains the JSON metadata
// for the struct [V2ContractGetResponseDataCommitsInvoiceContract]
type v2ContractGetResponseDataCommitsInvoiceContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsInvoiceContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsInvoiceContractJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataCommitsLedger struct {
	Amount        float64                                    `json:"amount,required"`
	Timestamp     time.Time                                  `json:"timestamp,required" format:"date-time"`
	Type          V2ContractGetResponseDataCommitsLedgerType `json:"type,required"`
	InvoiceID     string                                     `json:"invoice_id" format:"uuid"`
	NewContractID string                                     `json:"new_contract_id" format:"uuid"`
	Reason        string                                     `json:"reason"`
	SegmentID     string                                     `json:"segment_id" format:"uuid"`
	JSON          v2ContractGetResponseDataCommitsLedgerJSON `json:"-"`
	union         V2ContractGetResponseDataCommitsLedgerUnion
}

// v2ContractGetResponseDataCommitsLedgerJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataCommitsLedger]
type v2ContractGetResponseDataCommitsLedgerJSON struct {
	Amount        apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	InvoiceID     apijson.Field
	NewContractID apijson.Field
	Reason        apijson.Field
	SegmentID     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r v2ContractGetResponseDataCommitsLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *V2ContractGetResponseDataCommitsLedger) UnmarshalJSON(data []byte) (err error) {
	*r = V2ContractGetResponseDataCommitsLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V2ContractGetResponseDataCommitsLedgerUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry].
func (r V2ContractGetResponseDataCommitsLedger) AsUnion() V2ContractGetResponseDataCommitsLedgerUnion {
	return r.union
}

// Union satisfied by
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry],
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry] or
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry].
type V2ContractGetResponseDataCommitsLedgerUnion interface {
	implementsV2ContractGetResponseDataCommitsLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ContractGetResponseDataCommitsLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry{}),
		},
	)
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                        `json:"amount,required"`
	SegmentID string                                                                         `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                      `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                     `json:"amount,required"`
	InvoiceID string                                                                                      `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                   `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                    `json:"amount,required"`
	NewContractID string                                                                     `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                  `json:"timestamp,required" format:"date-time"`
	Type          V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          v2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                      `json:"amount,required"`
	SegmentID string                                                                       `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                    `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                    `json:"amount,required"`
	InvoiceID string                                                                     `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                  `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                    `json:"amount,required"`
	InvoiceID string                                                                     `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                  `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                           `json:"amount,required"`
	Timestamp time.Time                                                                         `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                      `json:"amount,required"`
	InvoiceID string                                                                                       `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                       `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                    `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                     `json:"amount,required"`
	NewContractID string                                                                      `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                   `json:"timestamp,required" format:"date-time"`
	Type          V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          v2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryType = "POSTPAID_COMMIT_ROLLOVER"
)

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                   `json:"amount,required"`
	InvoiceID string                                                                    `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                 `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry struct {
	Amount    float64                                                                  `json:"amount,required"`
	Reason    string                                                                   `json:"reason,required"`
	Timestamp time.Time                                                                `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryType = "PREPAID_COMMIT_MANUAL"
)

func (r V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry struct {
	Amount    float64                                                                   `json:"amount,required"`
	Reason    string                                                                    `json:"reason,required"`
	Timestamp time.Time                                                                 `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryType = "POSTPAID_COMMIT_MANUAL"
)

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                       `json:"amount,required"`
	Timestamp time.Time                                                                     `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry]
type v2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry) implementsV2ContractGetResponseDataCommitsLedger() {
}

type V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryType string

const (
	V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsLedgerType string

const (
	V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitSegmentStart               V2ContractGetResponseDataCommitsLedgerType = "PREPAID_COMMIT_SEGMENT_START"
	V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitAutomatedInvoiceDeduction  V2ContractGetResponseDataCommitsLedgerType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitRollover                   V2ContractGetResponseDataCommitsLedgerType = "PREPAID_COMMIT_ROLLOVER"
	V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitExpiration                 V2ContractGetResponseDataCommitsLedgerType = "PREPAID_COMMIT_EXPIRATION"
	V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitCanceled                   V2ContractGetResponseDataCommitsLedgerType = "PREPAID_COMMIT_CANCELED"
	V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitCredited                   V2ContractGetResponseDataCommitsLedgerType = "PREPAID_COMMIT_CREDITED"
	V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitInitialBalance            V2ContractGetResponseDataCommitsLedgerType = "POSTPAID_COMMIT_INITIAL_BALANCE"
	V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitAutomatedInvoiceDeduction V2ContractGetResponseDataCommitsLedgerType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitRollover                  V2ContractGetResponseDataCommitsLedgerType = "POSTPAID_COMMIT_ROLLOVER"
	V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitTrueup                    V2ContractGetResponseDataCommitsLedgerType = "POSTPAID_COMMIT_TRUEUP"
	V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitManual                     V2ContractGetResponseDataCommitsLedgerType = "PREPAID_COMMIT_MANUAL"
	V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitManual                    V2ContractGetResponseDataCommitsLedgerType = "POSTPAID_COMMIT_MANUAL"
	V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitExpiration                V2ContractGetResponseDataCommitsLedgerType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r V2ContractGetResponseDataCommitsLedgerType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitSegmentStart, V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitAutomatedInvoiceDeduction, V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitRollover, V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitExpiration, V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitCanceled, V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitCredited, V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitInitialBalance, V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitAutomatedInvoiceDeduction, V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitRollover, V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitTrueup, V2ContractGetResponseDataCommitsLedgerTypePrepaidCommitManual, V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitManual, V2ContractGetResponseDataCommitsLedgerTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsRateType string

const (
	V2ContractGetResponseDataCommitsRateTypeCommitRate V2ContractGetResponseDataCommitsRateType = "COMMIT_RATE"
	V2ContractGetResponseDataCommitsRateTypeListRate   V2ContractGetResponseDataCommitsRateType = "LIST_RATE"
)

func (r V2ContractGetResponseDataCommitsRateType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCommitsRateTypeCommitRate, V2ContractGetResponseDataCommitsRateTypeListRate:
		return true
	}
	return false
}

type V2ContractGetResponseDataCommitsRolledOverFrom struct {
	CommitID   string                                             `json:"commit_id,required" format:"uuid"`
	ContractID string                                             `json:"contract_id,required" format:"uuid"`
	JSON       v2ContractGetResponseDataCommitsRolledOverFromJSON `json:"-"`
}

// v2ContractGetResponseDataCommitsRolledOverFromJSON contains the JSON metadata
// for the struct [V2ContractGetResponseDataCommitsRolledOverFrom]
type v2ContractGetResponseDataCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCommitsRolledOverFromJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataOverride struct {
	ID                    string                                                `json:"id,required" format:"uuid"`
	StartingAt            time.Time                                             `json:"starting_at,required" format:"date-time"`
	ApplicableProductTags []string                                              `json:"applicable_product_tags"`
	EndingBefore          time.Time                                             `json:"ending_before" format:"date-time"`
	Entitled              bool                                                  `json:"entitled"`
	IsCommitSpecific      bool                                                  `json:"is_commit_specific"`
	Multiplier            float64                                               `json:"multiplier"`
	OverrideSpecifiers    []V2ContractGetResponseDataOverridesOverrideSpecifier `json:"override_specifiers"`
	OverrideTiers         []V2ContractGetResponseDataOverridesOverrideTier      `json:"override_tiers"`
	OverwriteRate         V2ContractGetResponseDataOverridesOverwriteRate       `json:"overwrite_rate"`
	Priority              float64                                               `json:"priority"`
	Product               V2ContractGetResponseDataOverridesProduct             `json:"product"`
	Target                V2ContractGetResponseDataOverridesTarget              `json:"target"`
	Type                  V2ContractGetResponseDataOverridesType                `json:"type"`
	JSON                  v2ContractGetResponseDataOverrideJSON                 `json:"-"`
}

// v2ContractGetResponseDataOverrideJSON contains the JSON metadata for the struct
// [V2ContractGetResponseDataOverride]
type v2ContractGetResponseDataOverrideJSON struct {
	ID                    apijson.Field
	StartingAt            apijson.Field
	ApplicableProductTags apijson.Field
	EndingBefore          apijson.Field
	Entitled              apijson.Field
	IsCommitSpecific      apijson.Field
	Multiplier            apijson.Field
	OverrideSpecifiers    apijson.Field
	OverrideTiers         apijson.Field
	OverwriteRate         apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	Target                apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractGetResponseDataOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataOverrideJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataOverridesOverrideSpecifier struct {
	BillingFrequency        V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequency `json:"billing_frequency"`
	CommitIDs               []string                                                             `json:"commit_ids"`
	PresentationGroupValues map[string]string                                                    `json:"presentation_group_values"`
	PricingGroupValues      map[string]string                                                    `json:"pricing_group_values"`
	ProductID               string                                                               `json:"product_id" format:"uuid"`
	ProductTags             []string                                                             `json:"product_tags"`
	RecurringCommitIDs      []string                                                             `json:"recurring_commit_ids"`
	RecurringCreditIDs      []string                                                             `json:"recurring_credit_ids"`
	JSON                    v2ContractGetResponseDataOverridesOverrideSpecifierJSON              `json:"-"`
}

// v2ContractGetResponseDataOverridesOverrideSpecifierJSON contains the JSON
// metadata for the struct [V2ContractGetResponseDataOverridesOverrideSpecifier]
type v2ContractGetResponseDataOverridesOverrideSpecifierJSON struct {
	BillingFrequency        apijson.Field
	CommitIDs               apijson.Field
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	RecurringCommitIDs      apijson.Field
	RecurringCreditIDs      apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetResponseDataOverridesOverrideSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataOverridesOverrideSpecifierJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequency string

const (
	V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequencyMonthly   V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequency = "MONTHLY"
	V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequencyQuarterly V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequency = "QUARTERLY"
	V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequencyAnnual    V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequency = "ANNUAL"
	V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequencyWeekly    V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequency = "WEEKLY"
)

func (r V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequency) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequencyMonthly, V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequencyQuarterly, V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequencyAnnual, V2ContractGetResponseDataOverridesOverrideSpecifiersBillingFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractGetResponseDataOverridesOverrideTier struct {
	Multiplier float64                                            `json:"multiplier,required"`
	Size       float64                                            `json:"size"`
	JSON       v2ContractGetResponseDataOverridesOverrideTierJSON `json:"-"`
}

// v2ContractGetResponseDataOverridesOverrideTierJSON contains the JSON metadata
// for the struct [V2ContractGetResponseDataOverridesOverrideTier]
type v2ContractGetResponseDataOverridesOverrideTierJSON struct {
	Multiplier  apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataOverridesOverrideTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataOverridesOverrideTierJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataOverridesOverwriteRate struct {
	RateType   V2ContractGetResponseDataOverridesOverwriteRateRateType `json:"rate_type,required"`
	CreditType shared.CreditTypeData                                   `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]interface{} `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier                                       `json:"tiers"`
	JSON  v2ContractGetResponseDataOverridesOverwriteRateJSON `json:"-"`
}

// v2ContractGetResponseDataOverridesOverwriteRateJSON contains the JSON metadata
// for the struct [V2ContractGetResponseDataOverridesOverwriteRate]
type v2ContractGetResponseDataOverridesOverwriteRateJSON struct {
	RateType    apijson.Field
	CreditType  apijson.Field
	CustomRate  apijson.Field
	IsProrated  apijson.Field
	Price       apijson.Field
	Quantity    apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataOverridesOverwriteRateJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataOverridesOverwriteRateRateType string

const (
	V2ContractGetResponseDataOverridesOverwriteRateRateTypeFlat         V2ContractGetResponseDataOverridesOverwriteRateRateType = "FLAT"
	V2ContractGetResponseDataOverridesOverwriteRateRateTypePercentage   V2ContractGetResponseDataOverridesOverwriteRateRateType = "PERCENTAGE"
	V2ContractGetResponseDataOverridesOverwriteRateRateTypeSubscription V2ContractGetResponseDataOverridesOverwriteRateRateType = "SUBSCRIPTION"
	V2ContractGetResponseDataOverridesOverwriteRateRateTypeTiered       V2ContractGetResponseDataOverridesOverwriteRateRateType = "TIERED"
	V2ContractGetResponseDataOverridesOverwriteRateRateTypeCustom       V2ContractGetResponseDataOverridesOverwriteRateRateType = "CUSTOM"
)

func (r V2ContractGetResponseDataOverridesOverwriteRateRateType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataOverridesOverwriteRateRateTypeFlat, V2ContractGetResponseDataOverridesOverwriteRateRateTypePercentage, V2ContractGetResponseDataOverridesOverwriteRateRateTypeSubscription, V2ContractGetResponseDataOverridesOverwriteRateRateTypeTiered, V2ContractGetResponseDataOverridesOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

type V2ContractGetResponseDataOverridesProduct struct {
	ID   string                                        `json:"id,required" format:"uuid"`
	Name string                                        `json:"name,required"`
	JSON v2ContractGetResponseDataOverridesProductJSON `json:"-"`
}

// v2ContractGetResponseDataOverridesProductJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataOverridesProduct]
type v2ContractGetResponseDataOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataOverridesProductJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataOverridesTarget string

const (
	V2ContractGetResponseDataOverridesTargetCommitRate V2ContractGetResponseDataOverridesTarget = "COMMIT_RATE"
	V2ContractGetResponseDataOverridesTargetListRate   V2ContractGetResponseDataOverridesTarget = "LIST_RATE"
)

func (r V2ContractGetResponseDataOverridesTarget) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataOverridesTargetCommitRate, V2ContractGetResponseDataOverridesTargetListRate:
		return true
	}
	return false
}

type V2ContractGetResponseDataOverridesType string

const (
	V2ContractGetResponseDataOverridesTypeOverwrite  V2ContractGetResponseDataOverridesType = "OVERWRITE"
	V2ContractGetResponseDataOverridesTypeMultiplier V2ContractGetResponseDataOverridesType = "MULTIPLIER"
	V2ContractGetResponseDataOverridesTypeTiered     V2ContractGetResponseDataOverridesType = "TIERED"
)

func (r V2ContractGetResponseDataOverridesType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataOverridesTypeOverwrite, V2ContractGetResponseDataOverridesTypeMultiplier, V2ContractGetResponseDataOverridesTypeTiered:
		return true
	}
	return false
}

type V2ContractGetResponseDataTransition struct {
	FromContractID string                                   `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                   `json:"to_contract_id,required" format:"uuid"`
	Type           V2ContractGetResponseDataTransitionsType `json:"type,required"`
	JSON           v2ContractGetResponseDataTransitionJSON  `json:"-"`
}

// v2ContractGetResponseDataTransitionJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataTransition]
type v2ContractGetResponseDataTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V2ContractGetResponseDataTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataTransitionJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataTransitionsType string

const (
	V2ContractGetResponseDataTransitionsTypeSupersede V2ContractGetResponseDataTransitionsType = "SUPERSEDE"
	V2ContractGetResponseDataTransitionsTypeRenewal   V2ContractGetResponseDataTransitionsType = "RENEWAL"
)

func (r V2ContractGetResponseDataTransitionsType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataTransitionsTypeSupersede, V2ContractGetResponseDataTransitionsTypeRenewal:
		return true
	}
	return false
}

type V2ContractGetResponseDataUsageFilter struct {
	GroupKey    string   `json:"group_key,required"`
	GroupValues []string `json:"group_values,required"`
	// This will match contract starting_at value if usage filter is active from the
	// beginning of the contract.
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// This will match contract ending_before value if usage filter is active until the
	// end of the contract. It will be undefined if the contract is open-ended.
	EndingBefore time.Time                                `json:"ending_before" format:"date-time"`
	JSON         v2ContractGetResponseDataUsageFilterJSON `json:"-"`
}

// v2ContractGetResponseDataUsageFilterJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataUsageFilter]
type v2ContractGetResponseDataUsageFilterJSON struct {
	GroupKey     apijson.Field
	GroupValues  apijson.Field
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetResponseDataUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataUsageFilterJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataUsageStatementSchedule struct {
	// Contract usage statements follow a selected cadence based on this date.
	BillingAnchorDate time.Time                                                `json:"billing_anchor_date,required" format:"date-time"`
	Frequency         V2ContractGetResponseDataUsageStatementScheduleFrequency `json:"frequency,required"`
	JSON              v2ContractGetResponseDataUsageStatementScheduleJSON      `json:"-"`
}

// v2ContractGetResponseDataUsageStatementScheduleJSON contains the JSON metadata
// for the struct [V2ContractGetResponseDataUsageStatementSchedule]
type v2ContractGetResponseDataUsageStatementScheduleJSON struct {
	BillingAnchorDate apijson.Field
	Frequency         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *V2ContractGetResponseDataUsageStatementSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataUsageStatementScheduleJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataUsageStatementScheduleFrequency string

const (
	V2ContractGetResponseDataUsageStatementScheduleFrequencyMonthly   V2ContractGetResponseDataUsageStatementScheduleFrequency = "MONTHLY"
	V2ContractGetResponseDataUsageStatementScheduleFrequencyQuarterly V2ContractGetResponseDataUsageStatementScheduleFrequency = "QUARTERLY"
	V2ContractGetResponseDataUsageStatementScheduleFrequencyAnnual    V2ContractGetResponseDataUsageStatementScheduleFrequency = "ANNUAL"
	V2ContractGetResponseDataUsageStatementScheduleFrequencyWeekly    V2ContractGetResponseDataUsageStatementScheduleFrequency = "WEEKLY"
)

func (r V2ContractGetResponseDataUsageStatementScheduleFrequency) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataUsageStatementScheduleFrequencyMonthly, V2ContractGetResponseDataUsageStatementScheduleFrequencyQuarterly, V2ContractGetResponseDataUsageStatementScheduleFrequencyAnnual, V2ContractGetResponseDataUsageStatementScheduleFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractGetResponseDataCredit struct {
	ID      string                                  `json:"id,required" format:"uuid"`
	Product V2ContractGetResponseDataCreditsProduct `json:"product,required"`
	Type    V2ContractGetResponseDataCreditsType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        shared.ScheduleDuration `json:"access_schedule"`
	ApplicableContractIDs []string                `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string                `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64                                  `json:"balance"`
	Contract     V2ContractGetResponseDataCreditsContract `json:"contract"`
	CustomFields map[string]string                        `json:"custom_fields"`
	Description  string                                   `json:"description"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []V2ContractGetResponseDataCreditsLedger `json:"ledger"`
	Name   string                                   `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                              `json:"salesforce_opportunity_id"`
	JSON                    v2ContractGetResponseDataCreditJSON `json:"-"`
}

// v2ContractGetResponseDataCreditJSON contains the JSON metadata for the struct
// [V2ContractGetResponseDataCredit]
type v2ContractGetResponseDataCreditJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Balance                 apijson.Field
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

func (r *V2ContractGetResponseDataCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCreditJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataCreditsProduct struct {
	ID   string                                      `json:"id,required" format:"uuid"`
	Name string                                      `json:"name,required"`
	JSON v2ContractGetResponseDataCreditsProductJSON `json:"-"`
}

// v2ContractGetResponseDataCreditsProductJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataCreditsProduct]
type v2ContractGetResponseDataCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCreditsProductJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataCreditsType string

const (
	V2ContractGetResponseDataCreditsTypeCredit V2ContractGetResponseDataCreditsType = "CREDIT"
)

func (r V2ContractGetResponseDataCreditsType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCreditsTypeCredit:
		return true
	}
	return false
}

type V2ContractGetResponseDataCreditsContract struct {
	ID   string                                       `json:"id,required" format:"uuid"`
	JSON v2ContractGetResponseDataCreditsContractJSON `json:"-"`
}

// v2ContractGetResponseDataCreditsContractJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataCreditsContract]
type v2ContractGetResponseDataCreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCreditsContractJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataCreditsLedger struct {
	Amount    float64                                    `json:"amount,required"`
	Timestamp time.Time                                  `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCreditsLedgerType `json:"type,required"`
	InvoiceID string                                     `json:"invoice_id" format:"uuid"`
	Reason    string                                     `json:"reason"`
	SegmentID string                                     `json:"segment_id" format:"uuid"`
	JSON      v2ContractGetResponseDataCreditsLedgerJSON `json:"-"`
	union     V2ContractGetResponseDataCreditsLedgerUnion
}

// v2ContractGetResponseDataCreditsLedgerJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataCreditsLedger]
type v2ContractGetResponseDataCreditsLedgerJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	InvoiceID   apijson.Field
	Reason      apijson.Field
	SegmentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2ContractGetResponseDataCreditsLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *V2ContractGetResponseDataCreditsLedger) UnmarshalJSON(data []byte) (err error) {
	*r = V2ContractGetResponseDataCreditsLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V2ContractGetResponseDataCreditsLedgerUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntry],
// [V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntry],
// [V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntry],
// [V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntry],
// [V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntry].
func (r V2ContractGetResponseDataCreditsLedger) AsUnion() V2ContractGetResponseDataCreditsLedgerUnion {
	return r.union
}

// Union satisfied by
// [V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntry],
// [V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntry],
// [V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntry],
// [V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntry] or
// [V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntry].
type V2ContractGetResponseDataCreditsLedgerUnion interface {
	implementsV2ContractGetResponseDataCreditsLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ContractGetResponseDataCreditsLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntry{}),
		},
	)
}

type V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64                                                                 `json:"amount,required"`
	SegmentID string                                                                  `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                               `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntryJSON contains
// the JSON metadata for the struct
// [V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntry]
type v2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntry) implementsV2ContractGetResponseDataCreditsLedger() {
}

type V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntryType string

const (
	V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntryType = "CREDIT_SEGMENT_START"
)

func (r V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart:
		return true
	}
	return false
}

type V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                              `json:"amount,required"`
	InvoiceID string                                                                               `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                            `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry]
type v2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) implementsV2ContractGetResponseDataCreditsLedger() {
}

type V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType string

const (
	V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntry struct {
	Amount    float64                                                               `json:"amount,required"`
	SegmentID string                                                                `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                             `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntryJSON contains
// the JSON metadata for the struct
// [V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntry]
type v2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntry) implementsV2ContractGetResponseDataCreditsLedger() {
}

type V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntryType string

const (
	V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntryType = "CREDIT_EXPIRATION"
)

func (r V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration:
		return true
	}
	return false
}

type V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntry struct {
	Amount    float64                                                             `json:"amount,required"`
	InvoiceID string                                                              `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                              `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                           `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntryJSON contains the
// JSON metadata for the struct
// [V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntry]
type v2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntry) implementsV2ContractGetResponseDataCreditsLedger() {
}

type V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntryType string

const (
	V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntryType = "CREDIT_CANCELED"
)

func (r V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled:
		return true
	}
	return false
}

type V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntry struct {
	Amount    float64                                                             `json:"amount,required"`
	InvoiceID string                                                              `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                              `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                           `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntryJSON contains the
// JSON metadata for the struct
// [V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntry]
type v2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntry) implementsV2ContractGetResponseDataCreditsLedger() {
}

type V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntryType string

const (
	V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntryType = "CREDIT_CREDITED"
)

func (r V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited:
		return true
	}
	return false
}

type V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntry struct {
	Amount    float64                                                           `json:"amount,required"`
	Reason    string                                                            `json:"reason,required"`
	Timestamp time.Time                                                         `json:"timestamp,required" format:"date-time"`
	Type      V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntryType `json:"type,required"`
	JSON      v2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntryJSON `json:"-"`
}

// v2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntryJSON contains the
// JSON metadata for the struct
// [V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntry]
type v2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntry) implementsV2ContractGetResponseDataCreditsLedger() {
}

type V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntryType string

const (
	V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntryTypeCreditManual V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntryType = "CREDIT_MANUAL"
)

func (r V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCreditsLedgerCreditManualLedgerEntryTypeCreditManual:
		return true
	}
	return false
}

type V2ContractGetResponseDataCreditsLedgerType string

const (
	V2ContractGetResponseDataCreditsLedgerTypeCreditSegmentStart              V2ContractGetResponseDataCreditsLedgerType = "CREDIT_SEGMENT_START"
	V2ContractGetResponseDataCreditsLedgerTypeCreditAutomatedInvoiceDeduction V2ContractGetResponseDataCreditsLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	V2ContractGetResponseDataCreditsLedgerTypeCreditExpiration                V2ContractGetResponseDataCreditsLedgerType = "CREDIT_EXPIRATION"
	V2ContractGetResponseDataCreditsLedgerTypeCreditCanceled                  V2ContractGetResponseDataCreditsLedgerType = "CREDIT_CANCELED"
	V2ContractGetResponseDataCreditsLedgerTypeCreditCredited                  V2ContractGetResponseDataCreditsLedgerType = "CREDIT_CREDITED"
	V2ContractGetResponseDataCreditsLedgerTypeCreditManual                    V2ContractGetResponseDataCreditsLedgerType = "CREDIT_MANUAL"
)

func (r V2ContractGetResponseDataCreditsLedgerType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCreditsLedgerTypeCreditSegmentStart, V2ContractGetResponseDataCreditsLedgerTypeCreditAutomatedInvoiceDeduction, V2ContractGetResponseDataCreditsLedgerTypeCreditExpiration, V2ContractGetResponseDataCreditsLedgerTypeCreditCanceled, V2ContractGetResponseDataCreditsLedgerTypeCreditCredited, V2ContractGetResponseDataCreditsLedgerTypeCreditManual:
		return true
	}
	return false
}

// This field's availability is dependent on your client's configuration.
type V2ContractGetResponseDataCustomerBillingProviderConfiguration struct {
	BillingProvider V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider `json:"billing_provider,required"`
	DeliveryMethod  V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod  `json:"delivery_method,required"`
	JSON            v2ContractGetResponseDataCustomerBillingProviderConfigurationJSON            `json:"-"`
}

// v2ContractGetResponseDataCustomerBillingProviderConfigurationJSON contains the
// JSON metadata for the struct
// [V2ContractGetResponseDataCustomerBillingProviderConfiguration]
type v2ContractGetResponseDataCustomerBillingProviderConfigurationJSON struct {
	BillingProvider apijson.Field
	DeliveryMethod  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2ContractGetResponseDataCustomerBillingProviderConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataCustomerBillingProviderConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider string

const (
	V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace   V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "aws_marketplace"
	V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderStripe           V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "stripe"
	V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite         V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "netsuite"
	V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderCustom           V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "custom"
	V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "azure_marketplace"
	V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "quickbooks_online"
	V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday          V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "workday"
	V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace   V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "gcp_marketplace"
)

func (r V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace, V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderStripe, V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite, V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderCustom, V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace, V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline, V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday, V2ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace:
		return true
	}
	return false
}

type V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod string

const (
	V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "direct_to_billing_provider"
	V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs                  V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sqs"
	V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle                  V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "tackle"
	V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns                  V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sns"
)

func (r V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider, V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs, V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle, V2ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns:
		return true
	}
	return false
}

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
// prices automatically. EXPLICIT prioritization requires specifying priorities for
// each multiplier; the one with the lowest priority value will be prioritized
// first.
type V2ContractGetResponseDataMultiplierOverridePrioritization string

const (
	V2ContractGetResponseDataMultiplierOverridePrioritizationLowestMultiplier V2ContractGetResponseDataMultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	V2ContractGetResponseDataMultiplierOverridePrioritizationExplicit         V2ContractGetResponseDataMultiplierOverridePrioritization = "EXPLICIT"
)

func (r V2ContractGetResponseDataMultiplierOverridePrioritization) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataMultiplierOverridePrioritizationLowestMultiplier, V2ContractGetResponseDataMultiplierOverridePrioritizationExplicit:
		return true
	}
	return false
}

type V2ContractGetResponseDataRecurringCommit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractGetResponseDataRecurringCommitsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractGetResponseDataRecurringCommitsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                          `json:"priority,required"`
	Product  V2ContractGetResponseDataRecurringCommitsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType V2ContractGetResponseDataRecurringCommitsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                          `json:"applicable_product_tags"`
	Contract              V2ContractGetResponseDataRecurringCommitsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount V2ContractGetResponseDataRecurringCommitsInvoiceAmount `json:"invoice_amount"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration V2ContractGetResponseDataRecurringCommitsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency V2ContractGetResponseDataRecurringCommitsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64                                      `json:"rollover_fraction"`
	JSON             v2ContractGetResponseDataRecurringCommitJSON `json:"-"`
}

// v2ContractGetResponseDataRecurringCommitJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataRecurringCommit]
type v2ContractGetResponseDataRecurringCommitJSON struct {
	ID                    apijson.Field
	AccessAmount          apijson.Field
	CommitDuration        apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	RateType              apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Contract              apijson.Field
	Description           apijson.Field
	EndingBefore          apijson.Field
	InvoiceAmount         apijson.Field
	Name                  apijson.Field
	NetsuiteSalesOrderID  apijson.Field
	Proration             apijson.Field
	RecurrenceFrequency   apijson.Field
	RolloverFraction      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractGetResponseDataRecurringCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataRecurringCommitJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type V2ContractGetResponseDataRecurringCommitsAccessAmount struct {
	CreditTypeID string                                                    `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                   `json:"quantity,required"`
	UnitPrice    float64                                                   `json:"unit_price,required"`
	JSON         v2ContractGetResponseDataRecurringCommitsAccessAmountJSON `json:"-"`
}

// v2ContractGetResponseDataRecurringCommitsAccessAmountJSON contains the JSON
// metadata for the struct [V2ContractGetResponseDataRecurringCommitsAccessAmount]
type v2ContractGetResponseDataRecurringCommitsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetResponseDataRecurringCommitsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataRecurringCommitsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type V2ContractGetResponseDataRecurringCommitsCommitDuration struct {
	Value float64                                                     `json:"value,required"`
	Unit  V2ContractGetResponseDataRecurringCommitsCommitDurationUnit `json:"unit"`
	JSON  v2ContractGetResponseDataRecurringCommitsCommitDurationJSON `json:"-"`
}

// v2ContractGetResponseDataRecurringCommitsCommitDurationJSON contains the JSON
// metadata for the struct
// [V2ContractGetResponseDataRecurringCommitsCommitDuration]
type v2ContractGetResponseDataRecurringCommitsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataRecurringCommitsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataRecurringCommitsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataRecurringCommitsCommitDurationUnit string

const (
	V2ContractGetResponseDataRecurringCommitsCommitDurationUnitPeriods V2ContractGetResponseDataRecurringCommitsCommitDurationUnit = "PERIODS"
)

func (r V2ContractGetResponseDataRecurringCommitsCommitDurationUnit) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataRecurringCommitsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type V2ContractGetResponseDataRecurringCommitsProduct struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	Name string                                               `json:"name,required"`
	JSON v2ContractGetResponseDataRecurringCommitsProductJSON `json:"-"`
}

// v2ContractGetResponseDataRecurringCommitsProductJSON contains the JSON metadata
// for the struct [V2ContractGetResponseDataRecurringCommitsProduct]
type v2ContractGetResponseDataRecurringCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataRecurringCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataRecurringCommitsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type V2ContractGetResponseDataRecurringCommitsRateType string

const (
	V2ContractGetResponseDataRecurringCommitsRateTypeCommitRate V2ContractGetResponseDataRecurringCommitsRateType = "COMMIT_RATE"
	V2ContractGetResponseDataRecurringCommitsRateTypeListRate   V2ContractGetResponseDataRecurringCommitsRateType = "LIST_RATE"
)

func (r V2ContractGetResponseDataRecurringCommitsRateType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataRecurringCommitsRateTypeCommitRate, V2ContractGetResponseDataRecurringCommitsRateTypeListRate:
		return true
	}
	return false
}

type V2ContractGetResponseDataRecurringCommitsContract struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	JSON v2ContractGetResponseDataRecurringCommitsContractJSON `json:"-"`
}

// v2ContractGetResponseDataRecurringCommitsContractJSON contains the JSON metadata
// for the struct [V2ContractGetResponseDataRecurringCommitsContract]
type v2ContractGetResponseDataRecurringCommitsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataRecurringCommitsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataRecurringCommitsContractJSON) RawJSON() string {
	return r.raw
}

// The amount the customer should be billed for the commit. Not required.
type V2ContractGetResponseDataRecurringCommitsInvoiceAmount struct {
	CreditTypeID string                                                     `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                    `json:"quantity,required"`
	UnitPrice    float64                                                    `json:"unit_price,required"`
	JSON         v2ContractGetResponseDataRecurringCommitsInvoiceAmountJSON `json:"-"`
}

// v2ContractGetResponseDataRecurringCommitsInvoiceAmountJSON contains the JSON
// metadata for the struct [V2ContractGetResponseDataRecurringCommitsInvoiceAmount]
type v2ContractGetResponseDataRecurringCommitsInvoiceAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetResponseDataRecurringCommitsInvoiceAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataRecurringCommitsInvoiceAmountJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type V2ContractGetResponseDataRecurringCommitsProration string

const (
	V2ContractGetResponseDataRecurringCommitsProrationNone         V2ContractGetResponseDataRecurringCommitsProration = "NONE"
	V2ContractGetResponseDataRecurringCommitsProrationFirst        V2ContractGetResponseDataRecurringCommitsProration = "FIRST"
	V2ContractGetResponseDataRecurringCommitsProrationLast         V2ContractGetResponseDataRecurringCommitsProration = "LAST"
	V2ContractGetResponseDataRecurringCommitsProrationFirstAndLast V2ContractGetResponseDataRecurringCommitsProration = "FIRST_AND_LAST"
)

func (r V2ContractGetResponseDataRecurringCommitsProration) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataRecurringCommitsProrationNone, V2ContractGetResponseDataRecurringCommitsProrationFirst, V2ContractGetResponseDataRecurringCommitsProrationLast, V2ContractGetResponseDataRecurringCommitsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type V2ContractGetResponseDataRecurringCommitsRecurrenceFrequency string

const (
	V2ContractGetResponseDataRecurringCommitsRecurrenceFrequencyMonthly   V2ContractGetResponseDataRecurringCommitsRecurrenceFrequency = "MONTHLY"
	V2ContractGetResponseDataRecurringCommitsRecurrenceFrequencyQuarterly V2ContractGetResponseDataRecurringCommitsRecurrenceFrequency = "QUARTERLY"
	V2ContractGetResponseDataRecurringCommitsRecurrenceFrequencyAnnual    V2ContractGetResponseDataRecurringCommitsRecurrenceFrequency = "ANNUAL"
	V2ContractGetResponseDataRecurringCommitsRecurrenceFrequencyWeekly    V2ContractGetResponseDataRecurringCommitsRecurrenceFrequency = "WEEKLY"
)

func (r V2ContractGetResponseDataRecurringCommitsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataRecurringCommitsRecurrenceFrequencyMonthly, V2ContractGetResponseDataRecurringCommitsRecurrenceFrequencyQuarterly, V2ContractGetResponseDataRecurringCommitsRecurrenceFrequencyAnnual, V2ContractGetResponseDataRecurringCommitsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractGetResponseDataRecurringCredit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractGetResponseDataRecurringCreditsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractGetResponseDataRecurringCreditsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                          `json:"priority,required"`
	Product  V2ContractGetResponseDataRecurringCreditsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType V2ContractGetResponseDataRecurringCreditsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                          `json:"applicable_product_tags"`
	Contract              V2ContractGetResponseDataRecurringCreditsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration V2ContractGetResponseDataRecurringCreditsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency V2ContractGetResponseDataRecurringCreditsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64                                      `json:"rollover_fraction"`
	JSON             v2ContractGetResponseDataRecurringCreditJSON `json:"-"`
}

// v2ContractGetResponseDataRecurringCreditJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataRecurringCredit]
type v2ContractGetResponseDataRecurringCreditJSON struct {
	ID                    apijson.Field
	AccessAmount          apijson.Field
	CommitDuration        apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	RateType              apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Contract              apijson.Field
	Description           apijson.Field
	EndingBefore          apijson.Field
	Name                  apijson.Field
	NetsuiteSalesOrderID  apijson.Field
	Proration             apijson.Field
	RecurrenceFrequency   apijson.Field
	RolloverFraction      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractGetResponseDataRecurringCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataRecurringCreditJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type V2ContractGetResponseDataRecurringCreditsAccessAmount struct {
	CreditTypeID string                                                    `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                   `json:"quantity,required"`
	UnitPrice    float64                                                   `json:"unit_price,required"`
	JSON         v2ContractGetResponseDataRecurringCreditsAccessAmountJSON `json:"-"`
}

// v2ContractGetResponseDataRecurringCreditsAccessAmountJSON contains the JSON
// metadata for the struct [V2ContractGetResponseDataRecurringCreditsAccessAmount]
type v2ContractGetResponseDataRecurringCreditsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetResponseDataRecurringCreditsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataRecurringCreditsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type V2ContractGetResponseDataRecurringCreditsCommitDuration struct {
	Value float64                                                     `json:"value,required"`
	Unit  V2ContractGetResponseDataRecurringCreditsCommitDurationUnit `json:"unit"`
	JSON  v2ContractGetResponseDataRecurringCreditsCommitDurationJSON `json:"-"`
}

// v2ContractGetResponseDataRecurringCreditsCommitDurationJSON contains the JSON
// metadata for the struct
// [V2ContractGetResponseDataRecurringCreditsCommitDuration]
type v2ContractGetResponseDataRecurringCreditsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataRecurringCreditsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataRecurringCreditsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataRecurringCreditsCommitDurationUnit string

const (
	V2ContractGetResponseDataRecurringCreditsCommitDurationUnitPeriods V2ContractGetResponseDataRecurringCreditsCommitDurationUnit = "PERIODS"
)

func (r V2ContractGetResponseDataRecurringCreditsCommitDurationUnit) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataRecurringCreditsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type V2ContractGetResponseDataRecurringCreditsProduct struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	Name string                                               `json:"name,required"`
	JSON v2ContractGetResponseDataRecurringCreditsProductJSON `json:"-"`
}

// v2ContractGetResponseDataRecurringCreditsProductJSON contains the JSON metadata
// for the struct [V2ContractGetResponseDataRecurringCreditsProduct]
type v2ContractGetResponseDataRecurringCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataRecurringCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataRecurringCreditsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type V2ContractGetResponseDataRecurringCreditsRateType string

const (
	V2ContractGetResponseDataRecurringCreditsRateTypeCommitRate V2ContractGetResponseDataRecurringCreditsRateType = "COMMIT_RATE"
	V2ContractGetResponseDataRecurringCreditsRateTypeListRate   V2ContractGetResponseDataRecurringCreditsRateType = "LIST_RATE"
)

func (r V2ContractGetResponseDataRecurringCreditsRateType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataRecurringCreditsRateTypeCommitRate, V2ContractGetResponseDataRecurringCreditsRateTypeListRate:
		return true
	}
	return false
}

type V2ContractGetResponseDataRecurringCreditsContract struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	JSON v2ContractGetResponseDataRecurringCreditsContractJSON `json:"-"`
}

// v2ContractGetResponseDataRecurringCreditsContractJSON contains the JSON metadata
// for the struct [V2ContractGetResponseDataRecurringCreditsContract]
type v2ContractGetResponseDataRecurringCreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetResponseDataRecurringCreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataRecurringCreditsContractJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type V2ContractGetResponseDataRecurringCreditsProration string

const (
	V2ContractGetResponseDataRecurringCreditsProrationNone         V2ContractGetResponseDataRecurringCreditsProration = "NONE"
	V2ContractGetResponseDataRecurringCreditsProrationFirst        V2ContractGetResponseDataRecurringCreditsProration = "FIRST"
	V2ContractGetResponseDataRecurringCreditsProrationLast         V2ContractGetResponseDataRecurringCreditsProration = "LAST"
	V2ContractGetResponseDataRecurringCreditsProrationFirstAndLast V2ContractGetResponseDataRecurringCreditsProration = "FIRST_AND_LAST"
)

func (r V2ContractGetResponseDataRecurringCreditsProration) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataRecurringCreditsProrationNone, V2ContractGetResponseDataRecurringCreditsProrationFirst, V2ContractGetResponseDataRecurringCreditsProrationLast, V2ContractGetResponseDataRecurringCreditsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type V2ContractGetResponseDataRecurringCreditsRecurrenceFrequency string

const (
	V2ContractGetResponseDataRecurringCreditsRecurrenceFrequencyMonthly   V2ContractGetResponseDataRecurringCreditsRecurrenceFrequency = "MONTHLY"
	V2ContractGetResponseDataRecurringCreditsRecurrenceFrequencyQuarterly V2ContractGetResponseDataRecurringCreditsRecurrenceFrequency = "QUARTERLY"
	V2ContractGetResponseDataRecurringCreditsRecurrenceFrequencyAnnual    V2ContractGetResponseDataRecurringCreditsRecurrenceFrequency = "ANNUAL"
	V2ContractGetResponseDataRecurringCreditsRecurrenceFrequencyWeekly    V2ContractGetResponseDataRecurringCreditsRecurrenceFrequency = "WEEKLY"
)

func (r V2ContractGetResponseDataRecurringCreditsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataRecurringCreditsRecurrenceFrequencyMonthly, V2ContractGetResponseDataRecurringCreditsRecurrenceFrequencyQuarterly, V2ContractGetResponseDataRecurringCreditsRecurrenceFrequencyAnnual, V2ContractGetResponseDataRecurringCreditsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractGetResponseDataResellerRoyalty struct {
	ResellerType V2ContractGetResponseDataResellerRoyaltiesResellerType `json:"reseller_type,required"`
	Segments     []V2ContractGetResponseDataResellerRoyaltiesSegment    `json:"segments,required"`
	JSON         v2ContractGetResponseDataResellerRoyaltyJSON           `json:"-"`
}

// v2ContractGetResponseDataResellerRoyaltyJSON contains the JSON metadata for the
// struct [V2ContractGetResponseDataResellerRoyalty]
type v2ContractGetResponseDataResellerRoyaltyJSON struct {
	ResellerType apijson.Field
	Segments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetResponseDataResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataResellerRoyaltiesResellerType string

const (
	V2ContractGetResponseDataResellerRoyaltiesResellerTypeAws           V2ContractGetResponseDataResellerRoyaltiesResellerType = "AWS"
	V2ContractGetResponseDataResellerRoyaltiesResellerTypeAwsProService V2ContractGetResponseDataResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	V2ContractGetResponseDataResellerRoyaltiesResellerTypeGcp           V2ContractGetResponseDataResellerRoyaltiesResellerType = "GCP"
	V2ContractGetResponseDataResellerRoyaltiesResellerTypeGcpProService V2ContractGetResponseDataResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r V2ContractGetResponseDataResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataResellerRoyaltiesResellerTypeAws, V2ContractGetResponseDataResellerRoyaltiesResellerTypeAwsProService, V2ContractGetResponseDataResellerRoyaltiesResellerTypeGcp, V2ContractGetResponseDataResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

type V2ContractGetResponseDataResellerRoyaltiesSegment struct {
	Fraction              float64                                                        `json:"fraction,required"`
	NetsuiteResellerID    string                                                         `json:"netsuite_reseller_id,required"`
	ResellerType          V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                                      `json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  []string                                                       `json:"applicable_product_ids"`
	ApplicableProductTags []string                                                       `json:"applicable_product_tags"`
	AwsAccountNumber      string                                                         `json:"aws_account_number"`
	AwsOfferID            string                                                         `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                         `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                      `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                         `json:"gcp_account_id"`
	GcpOfferID            string                                                         `json:"gcp_offer_id"`
	ResellerContractValue float64                                                        `json:"reseller_contract_value"`
	JSON                  v2ContractGetResponseDataResellerRoyaltiesSegmentJSON          `json:"-"`
}

// v2ContractGetResponseDataResellerRoyaltiesSegmentJSON contains the JSON metadata
// for the struct [V2ContractGetResponseDataResellerRoyaltiesSegment]
type v2ContractGetResponseDataResellerRoyaltiesSegmentJSON struct {
	Fraction              apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerType          apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	ResellerContractValue apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractGetResponseDataResellerRoyaltiesSegment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataResellerRoyaltiesSegmentJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerType string

const (
	V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerTypeAws           V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerType = "AWS"
	V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerTypeAwsProService V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerType = "AWS_PRO_SERVICE"
	V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerTypeGcp           V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerType = "GCP"
	V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerTypeGcpProService V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerType = "GCP_PRO_SERVICE"
)

func (r V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerType) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerTypeAws, V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerTypeAwsProService, V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerTypeGcp, V2ContractGetResponseDataResellerRoyaltiesSegmentsResellerTypeGcpProService:
		return true
	}
	return false
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type V2ContractGetResponseDataScheduledChargesOnUsageInvoices string

const (
	V2ContractGetResponseDataScheduledChargesOnUsageInvoicesAll V2ContractGetResponseDataScheduledChargesOnUsageInvoices = "ALL"
)

func (r V2ContractGetResponseDataScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case V2ContractGetResponseDataScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
}

type V2ContractGetResponseDataThresholdBillingConfiguration struct {
	Commit V2ContractGetResponseDataThresholdBillingConfigurationCommit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state
	IsEnabled bool `json:"is_enabled,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64                                                    `json:"threshold_amount,required"`
	JSON            v2ContractGetResponseDataThresholdBillingConfigurationJSON `json:"-"`
}

// v2ContractGetResponseDataThresholdBillingConfigurationJSON contains the JSON
// metadata for the struct [V2ContractGetResponseDataThresholdBillingConfiguration]
type v2ContractGetResponseDataThresholdBillingConfigurationJSON struct {
	Commit          apijson.Field
	IsEnabled       apijson.Field
	ThresholdAmount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2ContractGetResponseDataThresholdBillingConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataThresholdBillingConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetResponseDataThresholdBillingConfigurationCommit struct {
	ProductID string `json:"product_id,required"`
	// Which products the threshold commit applies to. If both applicable_product_ids
	// and applicable_product_tags are not provided, the commit applies to all
	// products.
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags"`
	Description           string   `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string                                                           `json:"name"`
	JSON v2ContractGetResponseDataThresholdBillingConfigurationCommitJSON `json:"-"`
}

// v2ContractGetResponseDataThresholdBillingConfigurationCommitJSON contains the
// JSON metadata for the struct
// [V2ContractGetResponseDataThresholdBillingConfigurationCommit]
type v2ContractGetResponseDataThresholdBillingConfigurationCommitJSON struct {
	ProductID             apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Description           apijson.Field
	Name                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractGetResponseDataThresholdBillingConfigurationCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetResponseDataThresholdBillingConfigurationCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponse struct {
	Data []V2ContractListResponseData `json:"data,required"`
	JSON v2ContractListResponseJSON   `json:"-"`
}

// v2ContractListResponseJSON contains the JSON metadata for the struct
// [V2ContractListResponse]
type v2ContractListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseData struct {
	ID                     string                                           `json:"id,required" format:"uuid"`
	Commits                []V2ContractListResponseDataCommit               `json:"commits,required"`
	CreatedAt              time.Time                                        `json:"created_at,required" format:"date-time"`
	CreatedBy              string                                           `json:"created_by,required"`
	CustomerID             string                                           `json:"customer_id,required" format:"uuid"`
	Overrides              []V2ContractListResponseDataOverride             `json:"overrides,required"`
	ScheduledCharges       []shared.ScheduledCharge                         `json:"scheduled_charges,required"`
	StartingAt             time.Time                                        `json:"starting_at,required" format:"date-time"`
	Transitions            []V2ContractListResponseDataTransition           `json:"transitions,required"`
	UsageFilter            []V2ContractListResponseDataUsageFilter          `json:"usage_filter,required"`
	UsageStatementSchedule V2ContractListResponseDataUsageStatementSchedule `json:"usage_statement_schedule,required"`
	ArchivedAt             time.Time                                        `json:"archived_at" format:"date-time"`
	Credits                []V2ContractListResponseDataCredit               `json:"credits"`
	CustomFields           map[string]string                                `json:"custom_fields"`
	// This field's availability is dependent on your client's configuration.
	CustomerBillingProviderConfiguration V2ContractListResponseDataCustomerBillingProviderConfiguration `json:"customer_billing_provider_configuration"`
	// This field's availability is dependent on your client's configuration.
	Discounts    []shared.Discount `json:"discounts"`
	EndingBefore time.Time         `json:"ending_before" format:"date-time"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
	// prices automatically. EXPLICIT prioritization requires specifying priorities for
	// each multiplier; the one with the lowest priority value will be prioritized
	// first.
	MultiplierOverridePrioritization V2ContractListResponseDataMultiplierOverridePrioritization `json:"multiplier_override_prioritization"`
	Name                             string                                                     `json:"name"`
	NetPaymentTermsDays              float64                                                    `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []shared.ProService                         `json:"professional_services"`
	RateCardID           string                                      `json:"rate_card_id" format:"uuid"`
	RecurringCommits     []V2ContractListResponseDataRecurringCommit `json:"recurring_commits"`
	RecurringCredits     []V2ContractListResponseDataRecurringCredit `json:"recurring_credits"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []V2ContractListResponseDataResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices V2ContractListResponseDataScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	ThresholdBillingConfiguration   V2ContractListResponseDataThresholdBillingConfiguration   `json:"threshold_billing_configuration"`
	TotalContractValue              float64                                                   `json:"total_contract_value"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                         `json:"uniqueness_key"`
	JSON          v2ContractListResponseDataJSON `json:"-"`
}

// v2ContractListResponseDataJSON contains the JSON metadata for the struct
// [V2ContractListResponseData]
type v2ContractListResponseDataJSON struct {
	ID                                   apijson.Field
	Commits                              apijson.Field
	CreatedAt                            apijson.Field
	CreatedBy                            apijson.Field
	CustomerID                           apijson.Field
	Overrides                            apijson.Field
	ScheduledCharges                     apijson.Field
	StartingAt                           apijson.Field
	Transitions                          apijson.Field
	UsageFilter                          apijson.Field
	UsageStatementSchedule               apijson.Field
	ArchivedAt                           apijson.Field
	Credits                              apijson.Field
	CustomFields                         apijson.Field
	CustomerBillingProviderConfiguration apijson.Field
	Discounts                            apijson.Field
	EndingBefore                         apijson.Field
	MultiplierOverridePrioritization     apijson.Field
	Name                                 apijson.Field
	NetPaymentTermsDays                  apijson.Field
	NetsuiteSalesOrderID                 apijson.Field
	ProfessionalServices                 apijson.Field
	RateCardID                           apijson.Field
	RecurringCommits                     apijson.Field
	RecurringCredits                     apijson.Field
	ResellerRoyalties                    apijson.Field
	SalesforceOpportunityID              apijson.Field
	ScheduledChargesOnUsageInvoices      apijson.Field
	ThresholdBillingConfiguration        apijson.Field
	TotalContractValue                   apijson.Field
	UniquenessKey                        apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *V2ContractListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataCommit struct {
	ID      string                                   `json:"id,required" format:"uuid"`
	Product V2ContractListResponseDataCommitsProduct `json:"product,required"`
	Type    V2ContractListResponseDataCommitsType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule        shared.ScheduleDuration `json:"access_schedule"`
	ApplicableContractIDs []string                `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string                `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64                                   `json:"balance"`
	Contract     V2ContractListResponseDataCommitsContract `json:"contract"`
	CustomFields map[string]string                         `json:"custom_fields"`
	Description  string                                    `json:"description"`
	// The contract that this commit will be billed on.
	InvoiceContract V2ContractListResponseDataCommitsInvoiceContract `json:"invoice_contract"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule shared.SchedulePointInTime `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger []V2ContractListResponseDataCommitsLedger `json:"ledger"`
	Name   string                                    `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority         float64                                         `json:"priority"`
	RateType         V2ContractListResponseDataCommitsRateType       `json:"rate_type"`
	RolledOverFrom   V2ContractListResponseDataCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction float64                                         `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                               `json:"salesforce_opportunity_id"`
	JSON                    v2ContractListResponseDataCommitJSON `json:"-"`
}

// v2ContractListResponseDataCommitJSON contains the JSON metadata for the struct
// [V2ContractListResponseDataCommit]
type v2ContractListResponseDataCommitJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
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
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataCommitsProduct struct {
	ID   string                                       `json:"id,required" format:"uuid"`
	Name string                                       `json:"name,required"`
	JSON v2ContractListResponseDataCommitsProductJSON `json:"-"`
}

// v2ContractListResponseDataCommitsProductJSON contains the JSON metadata for the
// struct [V2ContractListResponseDataCommitsProduct]
type v2ContractListResponseDataCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsProductJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataCommitsType string

const (
	V2ContractListResponseDataCommitsTypePrepaid  V2ContractListResponseDataCommitsType = "PREPAID"
	V2ContractListResponseDataCommitsTypePostpaid V2ContractListResponseDataCommitsType = "POSTPAID"
)

func (r V2ContractListResponseDataCommitsType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsTypePrepaid, V2ContractListResponseDataCommitsTypePostpaid:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsContract struct {
	ID   string                                        `json:"id,required" format:"uuid"`
	JSON v2ContractListResponseDataCommitsContractJSON `json:"-"`
}

// v2ContractListResponseDataCommitsContractJSON contains the JSON metadata for the
// struct [V2ContractListResponseDataCommitsContract]
type v2ContractListResponseDataCommitsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsContractJSON) RawJSON() string {
	return r.raw
}

// The contract that this commit will be billed on.
type V2ContractListResponseDataCommitsInvoiceContract struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	JSON v2ContractListResponseDataCommitsInvoiceContractJSON `json:"-"`
}

// v2ContractListResponseDataCommitsInvoiceContractJSON contains the JSON metadata
// for the struct [V2ContractListResponseDataCommitsInvoiceContract]
type v2ContractListResponseDataCommitsInvoiceContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsInvoiceContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsInvoiceContractJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataCommitsLedger struct {
	Amount        float64                                     `json:"amount,required"`
	Timestamp     time.Time                                   `json:"timestamp,required" format:"date-time"`
	Type          V2ContractListResponseDataCommitsLedgerType `json:"type,required"`
	InvoiceID     string                                      `json:"invoice_id" format:"uuid"`
	NewContractID string                                      `json:"new_contract_id" format:"uuid"`
	Reason        string                                      `json:"reason"`
	SegmentID     string                                      `json:"segment_id" format:"uuid"`
	JSON          v2ContractListResponseDataCommitsLedgerJSON `json:"-"`
	union         V2ContractListResponseDataCommitsLedgerUnion
}

// v2ContractListResponseDataCommitsLedgerJSON contains the JSON metadata for the
// struct [V2ContractListResponseDataCommitsLedger]
type v2ContractListResponseDataCommitsLedgerJSON struct {
	Amount        apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	InvoiceID     apijson.Field
	NewContractID apijson.Field
	Reason        apijson.Field
	SegmentID     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r v2ContractListResponseDataCommitsLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *V2ContractListResponseDataCommitsLedger) UnmarshalJSON(data []byte) (err error) {
	*r = V2ContractListResponseDataCommitsLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V2ContractListResponseDataCommitsLedgerUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry].
func (r V2ContractListResponseDataCommitsLedger) AsUnion() V2ContractListResponseDataCommitsLedgerUnion {
	return r.union
}

// Union satisfied by
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry],
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry] or
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry].
type V2ContractListResponseDataCommitsLedgerUnion interface {
	implementsV2ContractListResponseDataCommitsLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ContractListResponseDataCommitsLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry{}),
		},
	)
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                         `json:"amount,required"`
	SegmentID string                                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                       `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                      `json:"amount,required"`
	InvoiceID string                                                                                       `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                       `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                    `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                     `json:"amount,required"`
	NewContractID string                                                                      `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                   `json:"timestamp,required" format:"date-time"`
	Type          V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          v2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                       `json:"amount,required"`
	SegmentID string                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                     `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                     `json:"amount,required"`
	InvoiceID string                                                                      `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                   `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                     `json:"amount,required"`
	InvoiceID string                                                                      `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                   `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                            `json:"amount,required"`
	Timestamp time.Time                                                                          `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                       `json:"amount,required"`
	InvoiceID string                                                                                        `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                     `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                      `json:"amount,required"`
	NewContractID string                                                                       `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                       `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                    `json:"timestamp,required" format:"date-time"`
	Type          V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          v2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryType = "POSTPAID_COMMIT_ROLLOVER"
)

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                    `json:"amount,required"`
	InvoiceID string                                                                     `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                  `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry struct {
	Amount    float64                                                                   `json:"amount,required"`
	Reason    string                                                                    `json:"reason,required"`
	Timestamp time.Time                                                                 `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryType = "PREPAID_COMMIT_MANUAL"
)

func (r V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry struct {
	Amount    float64                                                                    `json:"amount,required"`
	Reason    string                                                                     `json:"reason,required"`
	Timestamp time.Time                                                                  `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryType = "POSTPAID_COMMIT_MANUAL"
)

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                        `json:"amount,required"`
	Timestamp time.Time                                                                      `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry]
type v2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntry) implementsV2ContractListResponseDataCommitsLedger() {
}

type V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryType string

const (
	V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsLedgerType string

const (
	V2ContractListResponseDataCommitsLedgerTypePrepaidCommitSegmentStart               V2ContractListResponseDataCommitsLedgerType = "PREPAID_COMMIT_SEGMENT_START"
	V2ContractListResponseDataCommitsLedgerTypePrepaidCommitAutomatedInvoiceDeduction  V2ContractListResponseDataCommitsLedgerType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	V2ContractListResponseDataCommitsLedgerTypePrepaidCommitRollover                   V2ContractListResponseDataCommitsLedgerType = "PREPAID_COMMIT_ROLLOVER"
	V2ContractListResponseDataCommitsLedgerTypePrepaidCommitExpiration                 V2ContractListResponseDataCommitsLedgerType = "PREPAID_COMMIT_EXPIRATION"
	V2ContractListResponseDataCommitsLedgerTypePrepaidCommitCanceled                   V2ContractListResponseDataCommitsLedgerType = "PREPAID_COMMIT_CANCELED"
	V2ContractListResponseDataCommitsLedgerTypePrepaidCommitCredited                   V2ContractListResponseDataCommitsLedgerType = "PREPAID_COMMIT_CREDITED"
	V2ContractListResponseDataCommitsLedgerTypePostpaidCommitInitialBalance            V2ContractListResponseDataCommitsLedgerType = "POSTPAID_COMMIT_INITIAL_BALANCE"
	V2ContractListResponseDataCommitsLedgerTypePostpaidCommitAutomatedInvoiceDeduction V2ContractListResponseDataCommitsLedgerType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	V2ContractListResponseDataCommitsLedgerTypePostpaidCommitRollover                  V2ContractListResponseDataCommitsLedgerType = "POSTPAID_COMMIT_ROLLOVER"
	V2ContractListResponseDataCommitsLedgerTypePostpaidCommitTrueup                    V2ContractListResponseDataCommitsLedgerType = "POSTPAID_COMMIT_TRUEUP"
	V2ContractListResponseDataCommitsLedgerTypePrepaidCommitManual                     V2ContractListResponseDataCommitsLedgerType = "PREPAID_COMMIT_MANUAL"
	V2ContractListResponseDataCommitsLedgerTypePostpaidCommitManual                    V2ContractListResponseDataCommitsLedgerType = "POSTPAID_COMMIT_MANUAL"
	V2ContractListResponseDataCommitsLedgerTypePostpaidCommitExpiration                V2ContractListResponseDataCommitsLedgerType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r V2ContractListResponseDataCommitsLedgerType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsLedgerTypePrepaidCommitSegmentStart, V2ContractListResponseDataCommitsLedgerTypePrepaidCommitAutomatedInvoiceDeduction, V2ContractListResponseDataCommitsLedgerTypePrepaidCommitRollover, V2ContractListResponseDataCommitsLedgerTypePrepaidCommitExpiration, V2ContractListResponseDataCommitsLedgerTypePrepaidCommitCanceled, V2ContractListResponseDataCommitsLedgerTypePrepaidCommitCredited, V2ContractListResponseDataCommitsLedgerTypePostpaidCommitInitialBalance, V2ContractListResponseDataCommitsLedgerTypePostpaidCommitAutomatedInvoiceDeduction, V2ContractListResponseDataCommitsLedgerTypePostpaidCommitRollover, V2ContractListResponseDataCommitsLedgerTypePostpaidCommitTrueup, V2ContractListResponseDataCommitsLedgerTypePrepaidCommitManual, V2ContractListResponseDataCommitsLedgerTypePostpaidCommitManual, V2ContractListResponseDataCommitsLedgerTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsRateType string

const (
	V2ContractListResponseDataCommitsRateTypeCommitRate V2ContractListResponseDataCommitsRateType = "COMMIT_RATE"
	V2ContractListResponseDataCommitsRateTypeListRate   V2ContractListResponseDataCommitsRateType = "LIST_RATE"
)

func (r V2ContractListResponseDataCommitsRateType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCommitsRateTypeCommitRate, V2ContractListResponseDataCommitsRateTypeListRate:
		return true
	}
	return false
}

type V2ContractListResponseDataCommitsRolledOverFrom struct {
	CommitID   string                                              `json:"commit_id,required" format:"uuid"`
	ContractID string                                              `json:"contract_id,required" format:"uuid"`
	JSON       v2ContractListResponseDataCommitsRolledOverFromJSON `json:"-"`
}

// v2ContractListResponseDataCommitsRolledOverFromJSON contains the JSON metadata
// for the struct [V2ContractListResponseDataCommitsRolledOverFrom]
type v2ContractListResponseDataCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCommitsRolledOverFromJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataOverride struct {
	ID                    string                                                 `json:"id,required" format:"uuid"`
	StartingAt            time.Time                                              `json:"starting_at,required" format:"date-time"`
	ApplicableProductTags []string                                               `json:"applicable_product_tags"`
	EndingBefore          time.Time                                              `json:"ending_before" format:"date-time"`
	Entitled              bool                                                   `json:"entitled"`
	IsCommitSpecific      bool                                                   `json:"is_commit_specific"`
	Multiplier            float64                                                `json:"multiplier"`
	OverrideSpecifiers    []V2ContractListResponseDataOverridesOverrideSpecifier `json:"override_specifiers"`
	OverrideTiers         []V2ContractListResponseDataOverridesOverrideTier      `json:"override_tiers"`
	OverwriteRate         V2ContractListResponseDataOverridesOverwriteRate       `json:"overwrite_rate"`
	Priority              float64                                                `json:"priority"`
	Product               V2ContractListResponseDataOverridesProduct             `json:"product"`
	Target                V2ContractListResponseDataOverridesTarget              `json:"target"`
	Type                  V2ContractListResponseDataOverridesType                `json:"type"`
	JSON                  v2ContractListResponseDataOverrideJSON                 `json:"-"`
}

// v2ContractListResponseDataOverrideJSON contains the JSON metadata for the struct
// [V2ContractListResponseDataOverride]
type v2ContractListResponseDataOverrideJSON struct {
	ID                    apijson.Field
	StartingAt            apijson.Field
	ApplicableProductTags apijson.Field
	EndingBefore          apijson.Field
	Entitled              apijson.Field
	IsCommitSpecific      apijson.Field
	Multiplier            apijson.Field
	OverrideSpecifiers    apijson.Field
	OverrideTiers         apijson.Field
	OverwriteRate         apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	Target                apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractListResponseDataOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataOverrideJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataOverridesOverrideSpecifier struct {
	BillingFrequency        V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequency `json:"billing_frequency"`
	CommitIDs               []string                                                              `json:"commit_ids"`
	PresentationGroupValues map[string]string                                                     `json:"presentation_group_values"`
	PricingGroupValues      map[string]string                                                     `json:"pricing_group_values"`
	ProductID               string                                                                `json:"product_id" format:"uuid"`
	ProductTags             []string                                                              `json:"product_tags"`
	RecurringCommitIDs      []string                                                              `json:"recurring_commit_ids"`
	RecurringCreditIDs      []string                                                              `json:"recurring_credit_ids"`
	JSON                    v2ContractListResponseDataOverridesOverrideSpecifierJSON              `json:"-"`
}

// v2ContractListResponseDataOverridesOverrideSpecifierJSON contains the JSON
// metadata for the struct [V2ContractListResponseDataOverridesOverrideSpecifier]
type v2ContractListResponseDataOverridesOverrideSpecifierJSON struct {
	BillingFrequency        apijson.Field
	CommitIDs               apijson.Field
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	RecurringCommitIDs      apijson.Field
	RecurringCreditIDs      apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractListResponseDataOverridesOverrideSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataOverridesOverrideSpecifierJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequency string

const (
	V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequencyMonthly   V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequency = "MONTHLY"
	V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequencyQuarterly V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequency = "QUARTERLY"
	V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequencyAnnual    V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequency = "ANNUAL"
	V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequencyWeekly    V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequency = "WEEKLY"
)

func (r V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequency) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequencyMonthly, V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequencyQuarterly, V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequencyAnnual, V2ContractListResponseDataOverridesOverrideSpecifiersBillingFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractListResponseDataOverridesOverrideTier struct {
	Multiplier float64                                             `json:"multiplier,required"`
	Size       float64                                             `json:"size"`
	JSON       v2ContractListResponseDataOverridesOverrideTierJSON `json:"-"`
}

// v2ContractListResponseDataOverridesOverrideTierJSON contains the JSON metadata
// for the struct [V2ContractListResponseDataOverridesOverrideTier]
type v2ContractListResponseDataOverridesOverrideTierJSON struct {
	Multiplier  apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataOverridesOverrideTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataOverridesOverrideTierJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataOverridesOverwriteRate struct {
	RateType   V2ContractListResponseDataOverridesOverwriteRateRateType `json:"rate_type,required"`
	CreditType shared.CreditTypeData                                    `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]interface{} `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier                                        `json:"tiers"`
	JSON  v2ContractListResponseDataOverridesOverwriteRateJSON `json:"-"`
}

// v2ContractListResponseDataOverridesOverwriteRateJSON contains the JSON metadata
// for the struct [V2ContractListResponseDataOverridesOverwriteRate]
type v2ContractListResponseDataOverridesOverwriteRateJSON struct {
	RateType    apijson.Field
	CreditType  apijson.Field
	CustomRate  apijson.Field
	IsProrated  apijson.Field
	Price       apijson.Field
	Quantity    apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataOverridesOverwriteRateJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataOverridesOverwriteRateRateType string

const (
	V2ContractListResponseDataOverridesOverwriteRateRateTypeFlat         V2ContractListResponseDataOverridesOverwriteRateRateType = "FLAT"
	V2ContractListResponseDataOverridesOverwriteRateRateTypePercentage   V2ContractListResponseDataOverridesOverwriteRateRateType = "PERCENTAGE"
	V2ContractListResponseDataOverridesOverwriteRateRateTypeSubscription V2ContractListResponseDataOverridesOverwriteRateRateType = "SUBSCRIPTION"
	V2ContractListResponseDataOverridesOverwriteRateRateTypeTiered       V2ContractListResponseDataOverridesOverwriteRateRateType = "TIERED"
	V2ContractListResponseDataOverridesOverwriteRateRateTypeCustom       V2ContractListResponseDataOverridesOverwriteRateRateType = "CUSTOM"
)

func (r V2ContractListResponseDataOverridesOverwriteRateRateType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataOverridesOverwriteRateRateTypeFlat, V2ContractListResponseDataOverridesOverwriteRateRateTypePercentage, V2ContractListResponseDataOverridesOverwriteRateRateTypeSubscription, V2ContractListResponseDataOverridesOverwriteRateRateTypeTiered, V2ContractListResponseDataOverridesOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

type V2ContractListResponseDataOverridesProduct struct {
	ID   string                                         `json:"id,required" format:"uuid"`
	Name string                                         `json:"name,required"`
	JSON v2ContractListResponseDataOverridesProductJSON `json:"-"`
}

// v2ContractListResponseDataOverridesProductJSON contains the JSON metadata for
// the struct [V2ContractListResponseDataOverridesProduct]
type v2ContractListResponseDataOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataOverridesProductJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataOverridesTarget string

const (
	V2ContractListResponseDataOverridesTargetCommitRate V2ContractListResponseDataOverridesTarget = "COMMIT_RATE"
	V2ContractListResponseDataOverridesTargetListRate   V2ContractListResponseDataOverridesTarget = "LIST_RATE"
)

func (r V2ContractListResponseDataOverridesTarget) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataOverridesTargetCommitRate, V2ContractListResponseDataOverridesTargetListRate:
		return true
	}
	return false
}

type V2ContractListResponseDataOverridesType string

const (
	V2ContractListResponseDataOverridesTypeOverwrite  V2ContractListResponseDataOverridesType = "OVERWRITE"
	V2ContractListResponseDataOverridesTypeMultiplier V2ContractListResponseDataOverridesType = "MULTIPLIER"
	V2ContractListResponseDataOverridesTypeTiered     V2ContractListResponseDataOverridesType = "TIERED"
)

func (r V2ContractListResponseDataOverridesType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataOverridesTypeOverwrite, V2ContractListResponseDataOverridesTypeMultiplier, V2ContractListResponseDataOverridesTypeTiered:
		return true
	}
	return false
}

type V2ContractListResponseDataTransition struct {
	FromContractID string                                    `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                    `json:"to_contract_id,required" format:"uuid"`
	Type           V2ContractListResponseDataTransitionsType `json:"type,required"`
	JSON           v2ContractListResponseDataTransitionJSON  `json:"-"`
}

// v2ContractListResponseDataTransitionJSON contains the JSON metadata for the
// struct [V2ContractListResponseDataTransition]
type v2ContractListResponseDataTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V2ContractListResponseDataTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataTransitionJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataTransitionsType string

const (
	V2ContractListResponseDataTransitionsTypeSupersede V2ContractListResponseDataTransitionsType = "SUPERSEDE"
	V2ContractListResponseDataTransitionsTypeRenewal   V2ContractListResponseDataTransitionsType = "RENEWAL"
)

func (r V2ContractListResponseDataTransitionsType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataTransitionsTypeSupersede, V2ContractListResponseDataTransitionsTypeRenewal:
		return true
	}
	return false
}

type V2ContractListResponseDataUsageFilter struct {
	GroupKey    string   `json:"group_key,required"`
	GroupValues []string `json:"group_values,required"`
	// This will match contract starting_at value if usage filter is active from the
	// beginning of the contract.
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// This will match contract ending_before value if usage filter is active until the
	// end of the contract. It will be undefined if the contract is open-ended.
	EndingBefore time.Time                                 `json:"ending_before" format:"date-time"`
	JSON         v2ContractListResponseDataUsageFilterJSON `json:"-"`
}

// v2ContractListResponseDataUsageFilterJSON contains the JSON metadata for the
// struct [V2ContractListResponseDataUsageFilter]
type v2ContractListResponseDataUsageFilterJSON struct {
	GroupKey     apijson.Field
	GroupValues  apijson.Field
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractListResponseDataUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataUsageFilterJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataUsageStatementSchedule struct {
	// Contract usage statements follow a selected cadence based on this date.
	BillingAnchorDate time.Time                                                 `json:"billing_anchor_date,required" format:"date-time"`
	Frequency         V2ContractListResponseDataUsageStatementScheduleFrequency `json:"frequency,required"`
	JSON              v2ContractListResponseDataUsageStatementScheduleJSON      `json:"-"`
}

// v2ContractListResponseDataUsageStatementScheduleJSON contains the JSON metadata
// for the struct [V2ContractListResponseDataUsageStatementSchedule]
type v2ContractListResponseDataUsageStatementScheduleJSON struct {
	BillingAnchorDate apijson.Field
	Frequency         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *V2ContractListResponseDataUsageStatementSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataUsageStatementScheduleJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataUsageStatementScheduleFrequency string

const (
	V2ContractListResponseDataUsageStatementScheduleFrequencyMonthly   V2ContractListResponseDataUsageStatementScheduleFrequency = "MONTHLY"
	V2ContractListResponseDataUsageStatementScheduleFrequencyQuarterly V2ContractListResponseDataUsageStatementScheduleFrequency = "QUARTERLY"
	V2ContractListResponseDataUsageStatementScheduleFrequencyAnnual    V2ContractListResponseDataUsageStatementScheduleFrequency = "ANNUAL"
	V2ContractListResponseDataUsageStatementScheduleFrequencyWeekly    V2ContractListResponseDataUsageStatementScheduleFrequency = "WEEKLY"
)

func (r V2ContractListResponseDataUsageStatementScheduleFrequency) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataUsageStatementScheduleFrequencyMonthly, V2ContractListResponseDataUsageStatementScheduleFrequencyQuarterly, V2ContractListResponseDataUsageStatementScheduleFrequencyAnnual, V2ContractListResponseDataUsageStatementScheduleFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractListResponseDataCredit struct {
	ID      string                                   `json:"id,required" format:"uuid"`
	Product V2ContractListResponseDataCreditsProduct `json:"product,required"`
	Type    V2ContractListResponseDataCreditsType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        shared.ScheduleDuration `json:"access_schedule"`
	ApplicableContractIDs []string                `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string                `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64                                   `json:"balance"`
	Contract     V2ContractListResponseDataCreditsContract `json:"contract"`
	CustomFields map[string]string                         `json:"custom_fields"`
	Description  string                                    `json:"description"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []V2ContractListResponseDataCreditsLedger `json:"ledger"`
	Name   string                                    `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                               `json:"salesforce_opportunity_id"`
	JSON                    v2ContractListResponseDataCreditJSON `json:"-"`
}

// v2ContractListResponseDataCreditJSON contains the JSON metadata for the struct
// [V2ContractListResponseDataCredit]
type v2ContractListResponseDataCreditJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Balance                 apijson.Field
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

func (r *V2ContractListResponseDataCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCreditJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataCreditsProduct struct {
	ID   string                                       `json:"id,required" format:"uuid"`
	Name string                                       `json:"name,required"`
	JSON v2ContractListResponseDataCreditsProductJSON `json:"-"`
}

// v2ContractListResponseDataCreditsProductJSON contains the JSON metadata for the
// struct [V2ContractListResponseDataCreditsProduct]
type v2ContractListResponseDataCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCreditsProductJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataCreditsType string

const (
	V2ContractListResponseDataCreditsTypeCredit V2ContractListResponseDataCreditsType = "CREDIT"
)

func (r V2ContractListResponseDataCreditsType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCreditsTypeCredit:
		return true
	}
	return false
}

type V2ContractListResponseDataCreditsContract struct {
	ID   string                                        `json:"id,required" format:"uuid"`
	JSON v2ContractListResponseDataCreditsContractJSON `json:"-"`
}

// v2ContractListResponseDataCreditsContractJSON contains the JSON metadata for the
// struct [V2ContractListResponseDataCreditsContract]
type v2ContractListResponseDataCreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCreditsContractJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataCreditsLedger struct {
	Amount    float64                                     `json:"amount,required"`
	Timestamp time.Time                                   `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCreditsLedgerType `json:"type,required"`
	InvoiceID string                                      `json:"invoice_id" format:"uuid"`
	Reason    string                                      `json:"reason"`
	SegmentID string                                      `json:"segment_id" format:"uuid"`
	JSON      v2ContractListResponseDataCreditsLedgerJSON `json:"-"`
	union     V2ContractListResponseDataCreditsLedgerUnion
}

// v2ContractListResponseDataCreditsLedgerJSON contains the JSON metadata for the
// struct [V2ContractListResponseDataCreditsLedger]
type v2ContractListResponseDataCreditsLedgerJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	InvoiceID   apijson.Field
	Reason      apijson.Field
	SegmentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2ContractListResponseDataCreditsLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *V2ContractListResponseDataCreditsLedger) UnmarshalJSON(data []byte) (err error) {
	*r = V2ContractListResponseDataCreditsLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V2ContractListResponseDataCreditsLedgerUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntry],
// [V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntry],
// [V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntry],
// [V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntry],
// [V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntry].
func (r V2ContractListResponseDataCreditsLedger) AsUnion() V2ContractListResponseDataCreditsLedgerUnion {
	return r.union
}

// Union satisfied by
// [V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntry],
// [V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntry],
// [V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntry],
// [V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntry] or
// [V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntry].
type V2ContractListResponseDataCreditsLedgerUnion interface {
	implementsV2ContractListResponseDataCreditsLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ContractListResponseDataCreditsLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntry{}),
		},
	)
}

type V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64                                                                  `json:"amount,required"`
	SegmentID string                                                                   `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntry]
type v2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntry) implementsV2ContractListResponseDataCreditsLedger() {
}

type V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntryType string

const (
	V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntryType = "CREDIT_SEGMENT_START"
)

func (r V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCreditsLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart:
		return true
	}
	return false
}

type V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                               `json:"amount,required"`
	InvoiceID string                                                                                `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                             `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry]
type v2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntry) implementsV2ContractListResponseDataCreditsLedger() {
}

type V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType string

const (
	V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCreditsLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntry struct {
	Amount    float64                                                                `json:"amount,required"`
	SegmentID string                                                                 `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                              `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntryJSON contains
// the JSON metadata for the struct
// [V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntry]
type v2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntry) implementsV2ContractListResponseDataCreditsLedger() {
}

type V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntryType string

const (
	V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntryType = "CREDIT_EXPIRATION"
)

func (r V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCreditsLedgerCreditExpirationLedgerEntryTypeCreditExpiration:
		return true
	}
	return false
}

type V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntry struct {
	Amount    float64                                                              `json:"amount,required"`
	InvoiceID string                                                               `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                            `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntryJSON contains
// the JSON metadata for the struct
// [V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntry]
type v2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntry) implementsV2ContractListResponseDataCreditsLedger() {
}

type V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntryType string

const (
	V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntryType = "CREDIT_CANCELED"
)

func (r V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCreditsLedgerCreditCanceledLedgerEntryTypeCreditCanceled:
		return true
	}
	return false
}

type V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntry struct {
	Amount    float64                                                              `json:"amount,required"`
	InvoiceID string                                                               `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                            `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntryJSON contains
// the JSON metadata for the struct
// [V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntry]
type v2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntry) implementsV2ContractListResponseDataCreditsLedger() {
}

type V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntryType string

const (
	V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntryType = "CREDIT_CREDITED"
)

func (r V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCreditsLedgerCreditCreditedLedgerEntryTypeCreditCredited:
		return true
	}
	return false
}

type V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntry struct {
	Amount    float64                                                            `json:"amount,required"`
	Reason    string                                                             `json:"reason,required"`
	Timestamp time.Time                                                          `json:"timestamp,required" format:"date-time"`
	Type      V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntryType `json:"type,required"`
	JSON      v2ContractListResponseDataCreditsLedgerCreditManualLedgerEntryJSON `json:"-"`
}

// v2ContractListResponseDataCreditsLedgerCreditManualLedgerEntryJSON contains the
// JSON metadata for the struct
// [V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntry]
type v2ContractListResponseDataCreditsLedgerCreditManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCreditsLedgerCreditManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntry) implementsV2ContractListResponseDataCreditsLedger() {
}

type V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntryType string

const (
	V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntryTypeCreditManual V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntryType = "CREDIT_MANUAL"
)

func (r V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntryType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCreditsLedgerCreditManualLedgerEntryTypeCreditManual:
		return true
	}
	return false
}

type V2ContractListResponseDataCreditsLedgerType string

const (
	V2ContractListResponseDataCreditsLedgerTypeCreditSegmentStart              V2ContractListResponseDataCreditsLedgerType = "CREDIT_SEGMENT_START"
	V2ContractListResponseDataCreditsLedgerTypeCreditAutomatedInvoiceDeduction V2ContractListResponseDataCreditsLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	V2ContractListResponseDataCreditsLedgerTypeCreditExpiration                V2ContractListResponseDataCreditsLedgerType = "CREDIT_EXPIRATION"
	V2ContractListResponseDataCreditsLedgerTypeCreditCanceled                  V2ContractListResponseDataCreditsLedgerType = "CREDIT_CANCELED"
	V2ContractListResponseDataCreditsLedgerTypeCreditCredited                  V2ContractListResponseDataCreditsLedgerType = "CREDIT_CREDITED"
	V2ContractListResponseDataCreditsLedgerTypeCreditManual                    V2ContractListResponseDataCreditsLedgerType = "CREDIT_MANUAL"
)

func (r V2ContractListResponseDataCreditsLedgerType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCreditsLedgerTypeCreditSegmentStart, V2ContractListResponseDataCreditsLedgerTypeCreditAutomatedInvoiceDeduction, V2ContractListResponseDataCreditsLedgerTypeCreditExpiration, V2ContractListResponseDataCreditsLedgerTypeCreditCanceled, V2ContractListResponseDataCreditsLedgerTypeCreditCredited, V2ContractListResponseDataCreditsLedgerTypeCreditManual:
		return true
	}
	return false
}

// This field's availability is dependent on your client's configuration.
type V2ContractListResponseDataCustomerBillingProviderConfiguration struct {
	BillingProvider V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider `json:"billing_provider,required"`
	DeliveryMethod  V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod  `json:"delivery_method,required"`
	JSON            v2ContractListResponseDataCustomerBillingProviderConfigurationJSON            `json:"-"`
}

// v2ContractListResponseDataCustomerBillingProviderConfigurationJSON contains the
// JSON metadata for the struct
// [V2ContractListResponseDataCustomerBillingProviderConfiguration]
type v2ContractListResponseDataCustomerBillingProviderConfigurationJSON struct {
	BillingProvider apijson.Field
	DeliveryMethod  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2ContractListResponseDataCustomerBillingProviderConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataCustomerBillingProviderConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider string

const (
	V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace   V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "aws_marketplace"
	V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderStripe           V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "stripe"
	V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite         V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "netsuite"
	V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderCustom           V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "custom"
	V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "azure_marketplace"
	V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "quickbooks_online"
	V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday          V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "workday"
	V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace   V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "gcp_marketplace"
)

func (r V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace, V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderStripe, V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite, V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderCustom, V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace, V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline, V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday, V2ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace:
		return true
	}
	return false
}

type V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod string

const (
	V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "direct_to_billing_provider"
	V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs                  V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sqs"
	V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle                  V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "tackle"
	V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns                  V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sns"
)

func (r V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider, V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs, V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle, V2ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns:
		return true
	}
	return false
}

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
// prices automatically. EXPLICIT prioritization requires specifying priorities for
// each multiplier; the one with the lowest priority value will be prioritized
// first.
type V2ContractListResponseDataMultiplierOverridePrioritization string

const (
	V2ContractListResponseDataMultiplierOverridePrioritizationLowestMultiplier V2ContractListResponseDataMultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	V2ContractListResponseDataMultiplierOverridePrioritizationExplicit         V2ContractListResponseDataMultiplierOverridePrioritization = "EXPLICIT"
)

func (r V2ContractListResponseDataMultiplierOverridePrioritization) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataMultiplierOverridePrioritizationLowestMultiplier, V2ContractListResponseDataMultiplierOverridePrioritizationExplicit:
		return true
	}
	return false
}

type V2ContractListResponseDataRecurringCommit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractListResponseDataRecurringCommitsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractListResponseDataRecurringCommitsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                           `json:"priority,required"`
	Product  V2ContractListResponseDataRecurringCommitsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType V2ContractListResponseDataRecurringCommitsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                           `json:"applicable_product_tags"`
	Contract              V2ContractListResponseDataRecurringCommitsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount V2ContractListResponseDataRecurringCommitsInvoiceAmount `json:"invoice_amount"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration V2ContractListResponseDataRecurringCommitsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency V2ContractListResponseDataRecurringCommitsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64                                       `json:"rollover_fraction"`
	JSON             v2ContractListResponseDataRecurringCommitJSON `json:"-"`
}

// v2ContractListResponseDataRecurringCommitJSON contains the JSON metadata for the
// struct [V2ContractListResponseDataRecurringCommit]
type v2ContractListResponseDataRecurringCommitJSON struct {
	ID                    apijson.Field
	AccessAmount          apijson.Field
	CommitDuration        apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	RateType              apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Contract              apijson.Field
	Description           apijson.Field
	EndingBefore          apijson.Field
	InvoiceAmount         apijson.Field
	Name                  apijson.Field
	NetsuiteSalesOrderID  apijson.Field
	Proration             apijson.Field
	RecurrenceFrequency   apijson.Field
	RolloverFraction      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractListResponseDataRecurringCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataRecurringCommitJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type V2ContractListResponseDataRecurringCommitsAccessAmount struct {
	CreditTypeID string                                                     `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                    `json:"quantity,required"`
	UnitPrice    float64                                                    `json:"unit_price,required"`
	JSON         v2ContractListResponseDataRecurringCommitsAccessAmountJSON `json:"-"`
}

// v2ContractListResponseDataRecurringCommitsAccessAmountJSON contains the JSON
// metadata for the struct [V2ContractListResponseDataRecurringCommitsAccessAmount]
type v2ContractListResponseDataRecurringCommitsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractListResponseDataRecurringCommitsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataRecurringCommitsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type V2ContractListResponseDataRecurringCommitsCommitDuration struct {
	Value float64                                                      `json:"value,required"`
	Unit  V2ContractListResponseDataRecurringCommitsCommitDurationUnit `json:"unit"`
	JSON  v2ContractListResponseDataRecurringCommitsCommitDurationJSON `json:"-"`
}

// v2ContractListResponseDataRecurringCommitsCommitDurationJSON contains the JSON
// metadata for the struct
// [V2ContractListResponseDataRecurringCommitsCommitDuration]
type v2ContractListResponseDataRecurringCommitsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataRecurringCommitsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataRecurringCommitsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataRecurringCommitsCommitDurationUnit string

const (
	V2ContractListResponseDataRecurringCommitsCommitDurationUnitPeriods V2ContractListResponseDataRecurringCommitsCommitDurationUnit = "PERIODS"
)

func (r V2ContractListResponseDataRecurringCommitsCommitDurationUnit) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataRecurringCommitsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type V2ContractListResponseDataRecurringCommitsProduct struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	Name string                                                `json:"name,required"`
	JSON v2ContractListResponseDataRecurringCommitsProductJSON `json:"-"`
}

// v2ContractListResponseDataRecurringCommitsProductJSON contains the JSON metadata
// for the struct [V2ContractListResponseDataRecurringCommitsProduct]
type v2ContractListResponseDataRecurringCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataRecurringCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataRecurringCommitsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type V2ContractListResponseDataRecurringCommitsRateType string

const (
	V2ContractListResponseDataRecurringCommitsRateTypeCommitRate V2ContractListResponseDataRecurringCommitsRateType = "COMMIT_RATE"
	V2ContractListResponseDataRecurringCommitsRateTypeListRate   V2ContractListResponseDataRecurringCommitsRateType = "LIST_RATE"
)

func (r V2ContractListResponseDataRecurringCommitsRateType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataRecurringCommitsRateTypeCommitRate, V2ContractListResponseDataRecurringCommitsRateTypeListRate:
		return true
	}
	return false
}

type V2ContractListResponseDataRecurringCommitsContract struct {
	ID   string                                                 `json:"id,required" format:"uuid"`
	JSON v2ContractListResponseDataRecurringCommitsContractJSON `json:"-"`
}

// v2ContractListResponseDataRecurringCommitsContractJSON contains the JSON
// metadata for the struct [V2ContractListResponseDataRecurringCommitsContract]
type v2ContractListResponseDataRecurringCommitsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataRecurringCommitsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataRecurringCommitsContractJSON) RawJSON() string {
	return r.raw
}

// The amount the customer should be billed for the commit. Not required.
type V2ContractListResponseDataRecurringCommitsInvoiceAmount struct {
	CreditTypeID string                                                      `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                     `json:"quantity,required"`
	UnitPrice    float64                                                     `json:"unit_price,required"`
	JSON         v2ContractListResponseDataRecurringCommitsInvoiceAmountJSON `json:"-"`
}

// v2ContractListResponseDataRecurringCommitsInvoiceAmountJSON contains the JSON
// metadata for the struct
// [V2ContractListResponseDataRecurringCommitsInvoiceAmount]
type v2ContractListResponseDataRecurringCommitsInvoiceAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractListResponseDataRecurringCommitsInvoiceAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataRecurringCommitsInvoiceAmountJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type V2ContractListResponseDataRecurringCommitsProration string

const (
	V2ContractListResponseDataRecurringCommitsProrationNone         V2ContractListResponseDataRecurringCommitsProration = "NONE"
	V2ContractListResponseDataRecurringCommitsProrationFirst        V2ContractListResponseDataRecurringCommitsProration = "FIRST"
	V2ContractListResponseDataRecurringCommitsProrationLast         V2ContractListResponseDataRecurringCommitsProration = "LAST"
	V2ContractListResponseDataRecurringCommitsProrationFirstAndLast V2ContractListResponseDataRecurringCommitsProration = "FIRST_AND_LAST"
)

func (r V2ContractListResponseDataRecurringCommitsProration) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataRecurringCommitsProrationNone, V2ContractListResponseDataRecurringCommitsProrationFirst, V2ContractListResponseDataRecurringCommitsProrationLast, V2ContractListResponseDataRecurringCommitsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type V2ContractListResponseDataRecurringCommitsRecurrenceFrequency string

const (
	V2ContractListResponseDataRecurringCommitsRecurrenceFrequencyMonthly   V2ContractListResponseDataRecurringCommitsRecurrenceFrequency = "MONTHLY"
	V2ContractListResponseDataRecurringCommitsRecurrenceFrequencyQuarterly V2ContractListResponseDataRecurringCommitsRecurrenceFrequency = "QUARTERLY"
	V2ContractListResponseDataRecurringCommitsRecurrenceFrequencyAnnual    V2ContractListResponseDataRecurringCommitsRecurrenceFrequency = "ANNUAL"
	V2ContractListResponseDataRecurringCommitsRecurrenceFrequencyWeekly    V2ContractListResponseDataRecurringCommitsRecurrenceFrequency = "WEEKLY"
)

func (r V2ContractListResponseDataRecurringCommitsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataRecurringCommitsRecurrenceFrequencyMonthly, V2ContractListResponseDataRecurringCommitsRecurrenceFrequencyQuarterly, V2ContractListResponseDataRecurringCommitsRecurrenceFrequencyAnnual, V2ContractListResponseDataRecurringCommitsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractListResponseDataRecurringCredit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractListResponseDataRecurringCreditsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractListResponseDataRecurringCreditsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                           `json:"priority,required"`
	Product  V2ContractListResponseDataRecurringCreditsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType V2ContractListResponseDataRecurringCreditsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                           `json:"applicable_product_tags"`
	Contract              V2ContractListResponseDataRecurringCreditsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration V2ContractListResponseDataRecurringCreditsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency V2ContractListResponseDataRecurringCreditsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64                                       `json:"rollover_fraction"`
	JSON             v2ContractListResponseDataRecurringCreditJSON `json:"-"`
}

// v2ContractListResponseDataRecurringCreditJSON contains the JSON metadata for the
// struct [V2ContractListResponseDataRecurringCredit]
type v2ContractListResponseDataRecurringCreditJSON struct {
	ID                    apijson.Field
	AccessAmount          apijson.Field
	CommitDuration        apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	RateType              apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Contract              apijson.Field
	Description           apijson.Field
	EndingBefore          apijson.Field
	Name                  apijson.Field
	NetsuiteSalesOrderID  apijson.Field
	Proration             apijson.Field
	RecurrenceFrequency   apijson.Field
	RolloverFraction      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractListResponseDataRecurringCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataRecurringCreditJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type V2ContractListResponseDataRecurringCreditsAccessAmount struct {
	CreditTypeID string                                                     `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                    `json:"quantity,required"`
	UnitPrice    float64                                                    `json:"unit_price,required"`
	JSON         v2ContractListResponseDataRecurringCreditsAccessAmountJSON `json:"-"`
}

// v2ContractListResponseDataRecurringCreditsAccessAmountJSON contains the JSON
// metadata for the struct [V2ContractListResponseDataRecurringCreditsAccessAmount]
type v2ContractListResponseDataRecurringCreditsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractListResponseDataRecurringCreditsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataRecurringCreditsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type V2ContractListResponseDataRecurringCreditsCommitDuration struct {
	Value float64                                                      `json:"value,required"`
	Unit  V2ContractListResponseDataRecurringCreditsCommitDurationUnit `json:"unit"`
	JSON  v2ContractListResponseDataRecurringCreditsCommitDurationJSON `json:"-"`
}

// v2ContractListResponseDataRecurringCreditsCommitDurationJSON contains the JSON
// metadata for the struct
// [V2ContractListResponseDataRecurringCreditsCommitDuration]
type v2ContractListResponseDataRecurringCreditsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataRecurringCreditsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataRecurringCreditsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataRecurringCreditsCommitDurationUnit string

const (
	V2ContractListResponseDataRecurringCreditsCommitDurationUnitPeriods V2ContractListResponseDataRecurringCreditsCommitDurationUnit = "PERIODS"
)

func (r V2ContractListResponseDataRecurringCreditsCommitDurationUnit) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataRecurringCreditsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type V2ContractListResponseDataRecurringCreditsProduct struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	Name string                                                `json:"name,required"`
	JSON v2ContractListResponseDataRecurringCreditsProductJSON `json:"-"`
}

// v2ContractListResponseDataRecurringCreditsProductJSON contains the JSON metadata
// for the struct [V2ContractListResponseDataRecurringCreditsProduct]
type v2ContractListResponseDataRecurringCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataRecurringCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataRecurringCreditsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type V2ContractListResponseDataRecurringCreditsRateType string

const (
	V2ContractListResponseDataRecurringCreditsRateTypeCommitRate V2ContractListResponseDataRecurringCreditsRateType = "COMMIT_RATE"
	V2ContractListResponseDataRecurringCreditsRateTypeListRate   V2ContractListResponseDataRecurringCreditsRateType = "LIST_RATE"
)

func (r V2ContractListResponseDataRecurringCreditsRateType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataRecurringCreditsRateTypeCommitRate, V2ContractListResponseDataRecurringCreditsRateTypeListRate:
		return true
	}
	return false
}

type V2ContractListResponseDataRecurringCreditsContract struct {
	ID   string                                                 `json:"id,required" format:"uuid"`
	JSON v2ContractListResponseDataRecurringCreditsContractJSON `json:"-"`
}

// v2ContractListResponseDataRecurringCreditsContractJSON contains the JSON
// metadata for the struct [V2ContractListResponseDataRecurringCreditsContract]
type v2ContractListResponseDataRecurringCreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractListResponseDataRecurringCreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataRecurringCreditsContractJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type V2ContractListResponseDataRecurringCreditsProration string

const (
	V2ContractListResponseDataRecurringCreditsProrationNone         V2ContractListResponseDataRecurringCreditsProration = "NONE"
	V2ContractListResponseDataRecurringCreditsProrationFirst        V2ContractListResponseDataRecurringCreditsProration = "FIRST"
	V2ContractListResponseDataRecurringCreditsProrationLast         V2ContractListResponseDataRecurringCreditsProration = "LAST"
	V2ContractListResponseDataRecurringCreditsProrationFirstAndLast V2ContractListResponseDataRecurringCreditsProration = "FIRST_AND_LAST"
)

func (r V2ContractListResponseDataRecurringCreditsProration) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataRecurringCreditsProrationNone, V2ContractListResponseDataRecurringCreditsProrationFirst, V2ContractListResponseDataRecurringCreditsProrationLast, V2ContractListResponseDataRecurringCreditsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type V2ContractListResponseDataRecurringCreditsRecurrenceFrequency string

const (
	V2ContractListResponseDataRecurringCreditsRecurrenceFrequencyMonthly   V2ContractListResponseDataRecurringCreditsRecurrenceFrequency = "MONTHLY"
	V2ContractListResponseDataRecurringCreditsRecurrenceFrequencyQuarterly V2ContractListResponseDataRecurringCreditsRecurrenceFrequency = "QUARTERLY"
	V2ContractListResponseDataRecurringCreditsRecurrenceFrequencyAnnual    V2ContractListResponseDataRecurringCreditsRecurrenceFrequency = "ANNUAL"
	V2ContractListResponseDataRecurringCreditsRecurrenceFrequencyWeekly    V2ContractListResponseDataRecurringCreditsRecurrenceFrequency = "WEEKLY"
)

func (r V2ContractListResponseDataRecurringCreditsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataRecurringCreditsRecurrenceFrequencyMonthly, V2ContractListResponseDataRecurringCreditsRecurrenceFrequencyQuarterly, V2ContractListResponseDataRecurringCreditsRecurrenceFrequencyAnnual, V2ContractListResponseDataRecurringCreditsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractListResponseDataResellerRoyalty struct {
	ResellerType V2ContractListResponseDataResellerRoyaltiesResellerType `json:"reseller_type,required"`
	Segments     []V2ContractListResponseDataResellerRoyaltiesSegment    `json:"segments,required"`
	JSON         v2ContractListResponseDataResellerRoyaltyJSON           `json:"-"`
}

// v2ContractListResponseDataResellerRoyaltyJSON contains the JSON metadata for the
// struct [V2ContractListResponseDataResellerRoyalty]
type v2ContractListResponseDataResellerRoyaltyJSON struct {
	ResellerType apijson.Field
	Segments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractListResponseDataResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataResellerRoyaltiesResellerType string

const (
	V2ContractListResponseDataResellerRoyaltiesResellerTypeAws           V2ContractListResponseDataResellerRoyaltiesResellerType = "AWS"
	V2ContractListResponseDataResellerRoyaltiesResellerTypeAwsProService V2ContractListResponseDataResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	V2ContractListResponseDataResellerRoyaltiesResellerTypeGcp           V2ContractListResponseDataResellerRoyaltiesResellerType = "GCP"
	V2ContractListResponseDataResellerRoyaltiesResellerTypeGcpProService V2ContractListResponseDataResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r V2ContractListResponseDataResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataResellerRoyaltiesResellerTypeAws, V2ContractListResponseDataResellerRoyaltiesResellerTypeAwsProService, V2ContractListResponseDataResellerRoyaltiesResellerTypeGcp, V2ContractListResponseDataResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

type V2ContractListResponseDataResellerRoyaltiesSegment struct {
	Fraction              float64                                                         `json:"fraction,required"`
	NetsuiteResellerID    string                                                          `json:"netsuite_reseller_id,required"`
	ResellerType          V2ContractListResponseDataResellerRoyaltiesSegmentsResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                                       `json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  []string                                                        `json:"applicable_product_ids"`
	ApplicableProductTags []string                                                        `json:"applicable_product_tags"`
	AwsAccountNumber      string                                                          `json:"aws_account_number"`
	AwsOfferID            string                                                          `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                          `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                       `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                          `json:"gcp_account_id"`
	GcpOfferID            string                                                          `json:"gcp_offer_id"`
	ResellerContractValue float64                                                         `json:"reseller_contract_value"`
	JSON                  v2ContractListResponseDataResellerRoyaltiesSegmentJSON          `json:"-"`
}

// v2ContractListResponseDataResellerRoyaltiesSegmentJSON contains the JSON
// metadata for the struct [V2ContractListResponseDataResellerRoyaltiesSegment]
type v2ContractListResponseDataResellerRoyaltiesSegmentJSON struct {
	Fraction              apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerType          apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	ResellerContractValue apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractListResponseDataResellerRoyaltiesSegment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataResellerRoyaltiesSegmentJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataResellerRoyaltiesSegmentsResellerType string

const (
	V2ContractListResponseDataResellerRoyaltiesSegmentsResellerTypeAws           V2ContractListResponseDataResellerRoyaltiesSegmentsResellerType = "AWS"
	V2ContractListResponseDataResellerRoyaltiesSegmentsResellerTypeAwsProService V2ContractListResponseDataResellerRoyaltiesSegmentsResellerType = "AWS_PRO_SERVICE"
	V2ContractListResponseDataResellerRoyaltiesSegmentsResellerTypeGcp           V2ContractListResponseDataResellerRoyaltiesSegmentsResellerType = "GCP"
	V2ContractListResponseDataResellerRoyaltiesSegmentsResellerTypeGcpProService V2ContractListResponseDataResellerRoyaltiesSegmentsResellerType = "GCP_PRO_SERVICE"
)

func (r V2ContractListResponseDataResellerRoyaltiesSegmentsResellerType) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataResellerRoyaltiesSegmentsResellerTypeAws, V2ContractListResponseDataResellerRoyaltiesSegmentsResellerTypeAwsProService, V2ContractListResponseDataResellerRoyaltiesSegmentsResellerTypeGcp, V2ContractListResponseDataResellerRoyaltiesSegmentsResellerTypeGcpProService:
		return true
	}
	return false
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type V2ContractListResponseDataScheduledChargesOnUsageInvoices string

const (
	V2ContractListResponseDataScheduledChargesOnUsageInvoicesAll V2ContractListResponseDataScheduledChargesOnUsageInvoices = "ALL"
)

func (r V2ContractListResponseDataScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case V2ContractListResponseDataScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
}

type V2ContractListResponseDataThresholdBillingConfiguration struct {
	Commit V2ContractListResponseDataThresholdBillingConfigurationCommit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state
	IsEnabled bool `json:"is_enabled,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64                                                     `json:"threshold_amount,required"`
	JSON            v2ContractListResponseDataThresholdBillingConfigurationJSON `json:"-"`
}

// v2ContractListResponseDataThresholdBillingConfigurationJSON contains the JSON
// metadata for the struct
// [V2ContractListResponseDataThresholdBillingConfiguration]
type v2ContractListResponseDataThresholdBillingConfigurationJSON struct {
	Commit          apijson.Field
	IsEnabled       apijson.Field
	ThresholdAmount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2ContractListResponseDataThresholdBillingConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataThresholdBillingConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractListResponseDataThresholdBillingConfigurationCommit struct {
	ProductID string `json:"product_id,required"`
	// Which products the threshold commit applies to. If both applicable_product_ids
	// and applicable_product_tags are not provided, the commit applies to all
	// products.
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags"`
	Description           string   `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string                                                            `json:"name"`
	JSON v2ContractListResponseDataThresholdBillingConfigurationCommitJSON `json:"-"`
}

// v2ContractListResponseDataThresholdBillingConfigurationCommitJSON contains the
// JSON metadata for the struct
// [V2ContractListResponseDataThresholdBillingConfigurationCommit]
type v2ContractListResponseDataThresholdBillingConfigurationCommitJSON struct {
	ProductID             apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Description           apijson.Field
	Name                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractListResponseDataThresholdBillingConfigurationCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractListResponseDataThresholdBillingConfigurationCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractEditResponse struct {
	Data shared.ID                  `json:"data,required"`
	JSON v2ContractEditResponseJSON `json:"-"`
}

// v2ContractEditResponseJSON contains the JSON metadata for the struct
// [V2ContractEditResponse]
type v2ContractEditResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractEditResponseJSON) RawJSON() string {
	return r.raw
}

type V2ContractEditCommitResponse struct {
	Data shared.ID                        `json:"data,required"`
	JSON v2ContractEditCommitResponseJSON `json:"-"`
}

// v2ContractEditCommitResponseJSON contains the JSON metadata for the struct
// [V2ContractEditCommitResponse]
type v2ContractEditCommitResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractEditCommitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractEditCommitResponseJSON) RawJSON() string {
	return r.raw
}

type V2ContractEditCreditResponse struct {
	Data shared.ID                        `json:"data,required"`
	JSON v2ContractEditCreditResponseJSON `json:"-"`
}

// v2ContractEditCreditResponseJSON contains the JSON metadata for the struct
// [V2ContractEditCreditResponse]
type v2ContractEditCreditResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractEditCreditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractEditCreditResponseJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponse struct {
	Data []V2ContractGetEditHistoryResponseData `json:"data,required"`
	JSON v2ContractGetEditHistoryResponseJSON   `json:"-"`
}

// v2ContractGetEditHistoryResponseJSON contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponse]
type v2ContractGetEditHistoryResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseData struct {
	ID                     string                                                      `json:"id,required" format:"uuid"`
	AddCommits             []V2ContractGetEditHistoryResponseDataAddCommit             `json:"add_commits"`
	AddCredits             []V2ContractGetEditHistoryResponseDataAddCredit             `json:"add_credits"`
	AddDiscounts           []shared.Discount                                           `json:"add_discounts"`
	AddOverrides           []V2ContractGetEditHistoryResponseDataAddOverride           `json:"add_overrides"`
	AddProServices         []shared.ProService                                         `json:"add_pro_services"`
	AddRecurringCommits    []V2ContractGetEditHistoryResponseDataAddRecurringCommit    `json:"add_recurring_commits"`
	AddRecurringCredits    []V2ContractGetEditHistoryResponseDataAddRecurringCredit    `json:"add_recurring_credits"`
	AddResellerRoyalties   []V2ContractGetEditHistoryResponseDataAddResellerRoyalty    `json:"add_reseller_royalties"`
	AddScheduledCharges    []V2ContractGetEditHistoryResponseDataAddScheduledCharge    `json:"add_scheduled_charges"`
	AddUsageFilters        []V2ContractGetEditHistoryResponseDataAddUsageFilter        `json:"add_usage_filters"`
	Timestamp              time.Time                                                   `json:"timestamp" format:"date-time"`
	UpdateCommits          []V2ContractGetEditHistoryResponseDataUpdateCommit          `json:"update_commits"`
	UpdateContractEndDate  time.Time                                                   `json:"update_contract_end_date" format:"date-time"`
	UpdateCredits          []V2ContractGetEditHistoryResponseDataUpdateCredit          `json:"update_credits"`
	UpdateDiscounts        []V2ContractGetEditHistoryResponseDataUpdateDiscount        `json:"update_discounts"`
	UpdateRefundInvoices   []V2ContractGetEditHistoryResponseDataUpdateRefundInvoice   `json:"update_refund_invoices"`
	UpdateScheduledCharges []V2ContractGetEditHistoryResponseDataUpdateScheduledCharge `json:"update_scheduled_charges"`
	JSON                   v2ContractGetEditHistoryResponseDataJSON                    `json:"-"`
}

// v2ContractGetEditHistoryResponseDataJSON contains the JSON metadata for the
// struct [V2ContractGetEditHistoryResponseData]
type v2ContractGetEditHistoryResponseDataJSON struct {
	ID                     apijson.Field
	AddCommits             apijson.Field
	AddCredits             apijson.Field
	AddDiscounts           apijson.Field
	AddOverrides           apijson.Field
	AddProServices         apijson.Field
	AddRecurringCommits    apijson.Field
	AddRecurringCredits    apijson.Field
	AddResellerRoyalties   apijson.Field
	AddScheduledCharges    apijson.Field
	AddUsageFilters        apijson.Field
	Timestamp              apijson.Field
	UpdateCommits          apijson.Field
	UpdateContractEndDate  apijson.Field
	UpdateCredits          apijson.Field
	UpdateDiscounts        apijson.Field
	UpdateRefundInvoices   apijson.Field
	UpdateScheduledCharges apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddCommit struct {
	ID      string                                                `json:"id,required" format:"uuid"`
	Product V2ContractGetEditHistoryResponseDataAddCommitsProduct `json:"product,required"`
	Type    V2ContractGetEditHistoryResponseDataAddCommitsType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule        shared.ScheduleDuration `json:"access_schedule"`
	ApplicableProductIDs  []string                `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                `json:"applicable_product_tags"`
	Description           string                  `json:"description"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule shared.SchedulePointInTime `json:"invoice_schedule"`
	Name            string                     `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority         float64                                                `json:"priority"`
	RateType         V2ContractGetEditHistoryResponseDataAddCommitsRateType `json:"rate_type"`
	RolloverFraction float64                                                `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                                            `json:"salesforce_opportunity_id"`
	JSON                    v2ContractGetEditHistoryResponseDataAddCommitJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCommitJSON contains the JSON metadata for
// the struct [V2ContractGetEditHistoryResponseDataAddCommit]
type v2ContractGetEditHistoryResponseDataAddCommitJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Description             apijson.Field
	InvoiceSchedule         apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	RateType                apijson.Field
	RolloverFraction        apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddCommitsProduct struct {
	ID   string                                                    `json:"id,required" format:"uuid"`
	Name string                                                    `json:"name,required"`
	JSON v2ContractGetEditHistoryResponseDataAddCommitsProductJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCommitsProductJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataAddCommitsProduct]
type v2ContractGetEditHistoryResponseDataAddCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCommitsProductJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddCommitsType string

const (
	V2ContractGetEditHistoryResponseDataAddCommitsTypePrepaid  V2ContractGetEditHistoryResponseDataAddCommitsType = "PREPAID"
	V2ContractGetEditHistoryResponseDataAddCommitsTypePostpaid V2ContractGetEditHistoryResponseDataAddCommitsType = "POSTPAID"
)

func (r V2ContractGetEditHistoryResponseDataAddCommitsType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddCommitsTypePrepaid, V2ContractGetEditHistoryResponseDataAddCommitsTypePostpaid:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddCommitsRateType string

const (
	V2ContractGetEditHistoryResponseDataAddCommitsRateTypeCommitRate V2ContractGetEditHistoryResponseDataAddCommitsRateType = "COMMIT_RATE"
	V2ContractGetEditHistoryResponseDataAddCommitsRateTypeListRate   V2ContractGetEditHistoryResponseDataAddCommitsRateType = "LIST_RATE"
)

func (r V2ContractGetEditHistoryResponseDataAddCommitsRateType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddCommitsRateTypeCommitRate, V2ContractGetEditHistoryResponseDataAddCommitsRateTypeListRate:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddCredit struct {
	ID      string                                                `json:"id,required" format:"uuid"`
	Product V2ContractGetEditHistoryResponseDataAddCreditsProduct `json:"product,required"`
	Type    V2ContractGetEditHistoryResponseDataAddCreditsType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        shared.ScheduleDuration `json:"access_schedule"`
	ApplicableProductIDs  []string                `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                `json:"applicable_product_tags"`
	Description           string                  `json:"description"`
	Name                  string                  `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64 `json:"priority"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                                            `json:"salesforce_opportunity_id"`
	JSON                    v2ContractGetEditHistoryResponseDataAddCreditJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCreditJSON contains the JSON metadata for
// the struct [V2ContractGetEditHistoryResponseDataAddCredit]
type v2ContractGetEditHistoryResponseDataAddCreditJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Description             apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCreditJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddCreditsProduct struct {
	ID   string                                                    `json:"id,required" format:"uuid"`
	Name string                                                    `json:"name,required"`
	JSON v2ContractGetEditHistoryResponseDataAddCreditsProductJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCreditsProductJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataAddCreditsProduct]
type v2ContractGetEditHistoryResponseDataAddCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCreditsProductJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddCreditsType string

const (
	V2ContractGetEditHistoryResponseDataAddCreditsTypeCredit V2ContractGetEditHistoryResponseDataAddCreditsType = "CREDIT"
)

func (r V2ContractGetEditHistoryResponseDataAddCreditsType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddCreditsTypeCredit:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddOverride struct {
	ID                    string                                                              `json:"id,required" format:"uuid"`
	StartingAt            time.Time                                                           `json:"starting_at,required" format:"date-time"`
	ApplicableProductTags []string                                                            `json:"applicable_product_tags"`
	EndingBefore          time.Time                                                           `json:"ending_before" format:"date-time"`
	Entitled              bool                                                                `json:"entitled"`
	IsCommitSpecific      bool                                                                `json:"is_commit_specific"`
	Multiplier            float64                                                             `json:"multiplier"`
	OverrideSpecifiers    []V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifier `json:"override_specifiers"`
	OverrideTiers         []V2ContractGetEditHistoryResponseDataAddOverridesOverrideTier      `json:"override_tiers"`
	OverwriteRate         V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRate       `json:"overwrite_rate"`
	Priority              float64                                                             `json:"priority"`
	Product               V2ContractGetEditHistoryResponseDataAddOverridesProduct             `json:"product"`
	Target                V2ContractGetEditHistoryResponseDataAddOverridesTarget              `json:"target"`
	Type                  V2ContractGetEditHistoryResponseDataAddOverridesType                `json:"type"`
	JSON                  v2ContractGetEditHistoryResponseDataAddOverrideJSON                 `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddOverrideJSON contains the JSON metadata
// for the struct [V2ContractGetEditHistoryResponseDataAddOverride]
type v2ContractGetEditHistoryResponseDataAddOverrideJSON struct {
	ID                    apijson.Field
	StartingAt            apijson.Field
	ApplicableProductTags apijson.Field
	EndingBefore          apijson.Field
	Entitled              apijson.Field
	IsCommitSpecific      apijson.Field
	Multiplier            apijson.Field
	OverrideSpecifiers    apijson.Field
	OverrideTiers         apijson.Field
	OverwriteRate         apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	Target                apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddOverrideJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifier struct {
	BillingFrequency        V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequency `json:"billing_frequency"`
	CommitIDs               []string                                                                           `json:"commit_ids"`
	PresentationGroupValues map[string]string                                                                  `json:"presentation_group_values"`
	PricingGroupValues      map[string]string                                                                  `json:"pricing_group_values"`
	ProductID               string                                                                             `json:"product_id" format:"uuid"`
	ProductTags             []string                                                                           `json:"product_tags"`
	RecurringCommitIDs      []string                                                                           `json:"recurring_commit_ids"`
	RecurringCreditIDs      []string                                                                           `json:"recurring_credit_ids"`
	JSON                    v2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifierJSON              `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifierJSON contains
// the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifier]
type v2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifierJSON struct {
	BillingFrequency        apijson.Field
	CommitIDs               apijson.Field
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	RecurringCommitIDs      apijson.Field
	RecurringCreditIDs      apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifierJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequency string

const (
	V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequencyMonthly   V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequency = "MONTHLY"
	V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequencyQuarterly V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequency = "QUARTERLY"
	V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequencyAnnual    V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequency = "ANNUAL"
	V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequencyWeekly    V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequency = "WEEKLY"
)

func (r V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequency) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequencyMonthly, V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequencyQuarterly, V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequencyAnnual, V2ContractGetEditHistoryResponseDataAddOverridesOverrideSpecifiersBillingFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddOverridesOverrideTier struct {
	Multiplier float64                                                          `json:"multiplier,required"`
	Size       float64                                                          `json:"size"`
	JSON       v2ContractGetEditHistoryResponseDataAddOverridesOverrideTierJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddOverridesOverrideTierJSON contains the
// JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddOverridesOverrideTier]
type v2ContractGetEditHistoryResponseDataAddOverridesOverrideTierJSON struct {
	Multiplier  apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddOverridesOverrideTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddOverridesOverrideTierJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRate struct {
	RateType   V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateType `json:"rate_type,required"`
	CreditType shared.CreditTypeData                                                 `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]interface{} `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier                                                     `json:"tiers"`
	JSON  v2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateJSON contains the
// JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRate]
type v2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateJSON struct {
	RateType    apijson.Field
	CreditType  apijson.Field
	CustomRate  apijson.Field
	IsProrated  apijson.Field
	Price       apijson.Field
	Quantity    apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateType string

const (
	V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateTypeFlat         V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateType = "FLAT"
	V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateTypePercentage   V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateType = "PERCENTAGE"
	V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateTypeSubscription V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateType = "SUBSCRIPTION"
	V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateTypeTiered       V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateType = "TIERED"
	V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateTypeCustom       V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateType = "CUSTOM"
)

func (r V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateTypeFlat, V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateTypePercentage, V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateTypeSubscription, V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateTypeTiered, V2ContractGetEditHistoryResponseDataAddOverridesOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddOverridesProduct struct {
	ID   string                                                      `json:"id,required" format:"uuid"`
	Name string                                                      `json:"name,required"`
	JSON v2ContractGetEditHistoryResponseDataAddOverridesProductJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddOverridesProductJSON contains the JSON
// metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddOverridesProduct]
type v2ContractGetEditHistoryResponseDataAddOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddOverridesProductJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddOverridesTarget string

const (
	V2ContractGetEditHistoryResponseDataAddOverridesTargetCommitRate V2ContractGetEditHistoryResponseDataAddOverridesTarget = "COMMIT_RATE"
	V2ContractGetEditHistoryResponseDataAddOverridesTargetListRate   V2ContractGetEditHistoryResponseDataAddOverridesTarget = "LIST_RATE"
)

func (r V2ContractGetEditHistoryResponseDataAddOverridesTarget) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddOverridesTargetCommitRate, V2ContractGetEditHistoryResponseDataAddOverridesTargetListRate:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddOverridesType string

const (
	V2ContractGetEditHistoryResponseDataAddOverridesTypeOverwrite  V2ContractGetEditHistoryResponseDataAddOverridesType = "OVERWRITE"
	V2ContractGetEditHistoryResponseDataAddOverridesTypeMultiplier V2ContractGetEditHistoryResponseDataAddOverridesType = "MULTIPLIER"
	V2ContractGetEditHistoryResponseDataAddOverridesTypeTiered     V2ContractGetEditHistoryResponseDataAddOverridesType = "TIERED"
)

func (r V2ContractGetEditHistoryResponseDataAddOverridesType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddOverridesTypeOverwrite, V2ContractGetEditHistoryResponseDataAddOverridesTypeMultiplier, V2ContractGetEditHistoryResponseDataAddOverridesTypeTiered:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                                        `json:"priority,required"`
	Product  V2ContractGetEditHistoryResponseDataAddRecurringCommitsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType V2ContractGetEditHistoryResponseDataAddRecurringCommitsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                                        `json:"applicable_product_tags"`
	Contract              V2ContractGetEditHistoryResponseDataAddRecurringCommitsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount V2ContractGetEditHistoryResponseDataAddRecurringCommitsInvoiceAmount `json:"invoice_amount"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration V2ContractGetEditHistoryResponseDataAddRecurringCommitsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64                                                    `json:"rollover_fraction"`
	JSON             v2ContractGetEditHistoryResponseDataAddRecurringCommitJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataAddRecurringCommit]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitJSON struct {
	ID                    apijson.Field
	AccessAmount          apijson.Field
	CommitDuration        apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	RateType              apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Contract              apijson.Field
	Description           apijson.Field
	EndingBefore          apijson.Field
	InvoiceAmount         apijson.Field
	Name                  apijson.Field
	NetsuiteSalesOrderID  apijson.Field
	Proration             apijson.Field
	RecurrenceFrequency   apijson.Field
	RolloverFraction      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type V2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmount struct {
	CreditTypeID string                                                                  `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                                 `json:"quantity,required"`
	UnitPrice    float64                                                                 `json:"unit_price,required"`
	JSON         v2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmountJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmountJSON contains
// the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmount]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type V2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDuration struct {
	Value float64                                                                   `json:"value,required"`
	Unit  V2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDurationUnit `json:"unit"`
	JSON  v2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDurationJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDurationJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDuration]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDurationUnit string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDurationUnitPeriods V2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDurationUnit = "PERIODS"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDurationUnit) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCommitsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsProduct struct {
	ID   string                                                             `json:"id,required" format:"uuid"`
	Name string                                                             `json:"name,required"`
	JSON v2ContractGetEditHistoryResponseDataAddRecurringCommitsProductJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsProductJSON contains the
// JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsProduct]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type V2ContractGetEditHistoryResponseDataAddRecurringCommitsRateType string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsRateTypeCommitRate V2ContractGetEditHistoryResponseDataAddRecurringCommitsRateType = "COMMIT_RATE"
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsRateTypeListRate   V2ContractGetEditHistoryResponseDataAddRecurringCommitsRateType = "LIST_RATE"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsRateType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCommitsRateTypeCommitRate, V2ContractGetEditHistoryResponseDataAddRecurringCommitsRateTypeListRate:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsContract struct {
	ID   string                                                              `json:"id,required" format:"uuid"`
	JSON v2ContractGetEditHistoryResponseDataAddRecurringCommitsContractJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsContractJSON contains the
// JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsContract]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsContractJSON) RawJSON() string {
	return r.raw
}

// The amount the customer should be billed for the commit. Not required.
type V2ContractGetEditHistoryResponseDataAddRecurringCommitsInvoiceAmount struct {
	CreditTypeID string                                                                   `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                                  `json:"quantity,required"`
	UnitPrice    float64                                                                  `json:"unit_price,required"`
	JSON         v2ContractGetEditHistoryResponseDataAddRecurringCommitsInvoiceAmountJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsInvoiceAmountJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsInvoiceAmount]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsInvoiceAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsInvoiceAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsInvoiceAmountJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type V2ContractGetEditHistoryResponseDataAddRecurringCommitsProration string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsProrationNone         V2ContractGetEditHistoryResponseDataAddRecurringCommitsProration = "NONE"
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsProrationFirst        V2ContractGetEditHistoryResponseDataAddRecurringCommitsProration = "FIRST"
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsProrationLast         V2ContractGetEditHistoryResponseDataAddRecurringCommitsProration = "LAST"
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsProrationFirstAndLast V2ContractGetEditHistoryResponseDataAddRecurringCommitsProration = "FIRST_AND_LAST"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsProration) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCommitsProrationNone, V2ContractGetEditHistoryResponseDataAddRecurringCommitsProrationFirst, V2ContractGetEditHistoryResponseDataAddRecurringCommitsProrationLast, V2ContractGetEditHistoryResponseDataAddRecurringCommitsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequency string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequencyMonthly   V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequency = "MONTHLY"
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequencyQuarterly V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequency = "QUARTERLY"
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequencyAnnual    V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequency = "ANNUAL"
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequencyWeekly    V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequency = "WEEKLY"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequencyMonthly, V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequencyQuarterly, V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequencyAnnual, V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCredit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount V2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration V2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                                        `json:"priority,required"`
	Product  V2ContractGetEditHistoryResponseDataAddRecurringCreditsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType V2ContractGetEditHistoryResponseDataAddRecurringCreditsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                                        `json:"applicable_product_tags"`
	Contract              V2ContractGetEditHistoryResponseDataAddRecurringCreditsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration V2ContractGetEditHistoryResponseDataAddRecurringCreditsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64                                                    `json:"rollover_fraction"`
	JSON             v2ContractGetEditHistoryResponseDataAddRecurringCreditJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataAddRecurringCredit]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditJSON struct {
	ID                    apijson.Field
	AccessAmount          apijson.Field
	CommitDuration        apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	RateType              apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Contract              apijson.Field
	Description           apijson.Field
	EndingBefore          apijson.Field
	Name                  apijson.Field
	NetsuiteSalesOrderID  apijson.Field
	Proration             apijson.Field
	RecurrenceFrequency   apijson.Field
	RolloverFraction      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type V2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmount struct {
	CreditTypeID string                                                                  `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                                 `json:"quantity,required"`
	UnitPrice    float64                                                                 `json:"unit_price,required"`
	JSON         v2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmountJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmountJSON contains
// the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmount]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type V2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDuration struct {
	Value float64                                                                   `json:"value,required"`
	Unit  V2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDurationUnit `json:"unit"`
	JSON  v2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDurationJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDurationJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDuration]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDurationUnit string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDurationUnitPeriods V2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDurationUnit = "PERIODS"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDurationUnit) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCreditsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsProduct struct {
	ID   string                                                             `json:"id,required" format:"uuid"`
	Name string                                                             `json:"name,required"`
	JSON v2ContractGetEditHistoryResponseDataAddRecurringCreditsProductJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsProductJSON contains the
// JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsProduct]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type V2ContractGetEditHistoryResponseDataAddRecurringCreditsRateType string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsRateTypeCommitRate V2ContractGetEditHistoryResponseDataAddRecurringCreditsRateType = "COMMIT_RATE"
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsRateTypeListRate   V2ContractGetEditHistoryResponseDataAddRecurringCreditsRateType = "LIST_RATE"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsRateType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCreditsRateTypeCommitRate, V2ContractGetEditHistoryResponseDataAddRecurringCreditsRateTypeListRate:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsContract struct {
	ID   string                                                              `json:"id,required" format:"uuid"`
	JSON v2ContractGetEditHistoryResponseDataAddRecurringCreditsContractJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsContractJSON contains the
// JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsContract]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsContractJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type V2ContractGetEditHistoryResponseDataAddRecurringCreditsProration string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsProrationNone         V2ContractGetEditHistoryResponseDataAddRecurringCreditsProration = "NONE"
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsProrationFirst        V2ContractGetEditHistoryResponseDataAddRecurringCreditsProration = "FIRST"
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsProrationLast         V2ContractGetEditHistoryResponseDataAddRecurringCreditsProration = "LAST"
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsProrationFirstAndLast V2ContractGetEditHistoryResponseDataAddRecurringCreditsProration = "FIRST_AND_LAST"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsProration) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCreditsProrationNone, V2ContractGetEditHistoryResponseDataAddRecurringCreditsProrationFirst, V2ContractGetEditHistoryResponseDataAddRecurringCreditsProrationLast, V2ContractGetEditHistoryResponseDataAddRecurringCreditsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequency string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequencyMonthly   V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequency = "MONTHLY"
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequencyQuarterly V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequency = "QUARTERLY"
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequencyAnnual    V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequency = "ANNUAL"
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequencyWeekly    V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequency = "WEEKLY"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequencyMonthly, V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequencyQuarterly, V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequencyAnnual, V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddResellerRoyalty struct {
	ResellerType          V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerType `json:"reseller_type,required"`
	ApplicableProductIDs  []string                                                             `json:"applicable_product_ids"`
	ApplicableProductTags []string                                                             `json:"applicable_product_tags"`
	AwsAccountNumber      string                                                               `json:"aws_account_number"`
	AwsOfferID            string                                                               `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                               `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                            `json:"ending_before,nullable" format:"date-time"`
	Fraction              float64                                                              `json:"fraction"`
	GcpAccountID          string                                                               `json:"gcp_account_id"`
	GcpOfferID            string                                                               `json:"gcp_offer_id"`
	NetsuiteResellerID    string                                                               `json:"netsuite_reseller_id"`
	ResellerContractValue float64                                                              `json:"reseller_contract_value"`
	StartingAt            time.Time                                                            `json:"starting_at" format:"date-time"`
	JSON                  v2ContractGetEditHistoryResponseDataAddResellerRoyaltyJSON           `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddResellerRoyaltyJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataAddResellerRoyalty]
type v2ContractGetEditHistoryResponseDataAddResellerRoyaltyJSON struct {
	ResellerType          apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
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

func (r *V2ContractGetEditHistoryResponseDataAddResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerType string

const (
	V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerTypeAws           V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerType = "AWS"
	V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerTypeAwsProService V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerTypeGcp           V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerType = "GCP"
	V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerTypeGcpProService V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerTypeAws, V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerTypeAwsProService, V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerTypeGcp, V2ContractGetEditHistoryResponseDataAddResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddScheduledCharge struct {
	ID       string                                                         `json:"id,required" format:"uuid"`
	Product  V2ContractGetEditHistoryResponseDataAddScheduledChargesProduct `json:"product,required"`
	Schedule shared.SchedulePointInTime                                     `json:"schedule,required"`
	// displayed on invoices
	Name string `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string                                                     `json:"netsuite_sales_order_id"`
	JSON                 v2ContractGetEditHistoryResponseDataAddScheduledChargeJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddScheduledChargeJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataAddScheduledCharge]
type v2ContractGetEditHistoryResponseDataAddScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddScheduledChargeJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddScheduledChargesProduct struct {
	ID   string                                                             `json:"id,required" format:"uuid"`
	Name string                                                             `json:"name,required"`
	JSON v2ContractGetEditHistoryResponseDataAddScheduledChargesProductJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddScheduledChargesProductJSON contains the
// JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddScheduledChargesProduct]
type v2ContractGetEditHistoryResponseDataAddScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddScheduledChargesProductJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddUsageFilter struct {
	GroupKey    string   `json:"group_key,required"`
	GroupValues []string `json:"group_values,required"`
	// This will match contract starting_at value if usage filter is active from the
	// beginning of the contract.
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// This will match contract ending_before value if usage filter is active until the
	// end of the contract. It will be undefined if the contract is open-ended.
	EndingBefore time.Time                                              `json:"ending_before" format:"date-time"`
	JSON         v2ContractGetEditHistoryResponseDataAddUsageFilterJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddUsageFilterJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataAddUsageFilter]
type v2ContractGetEditHistoryResponseDataAddUsageFilterJSON struct {
	GroupKey     apijson.Field
	GroupValues  apijson.Field
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddUsageFilterJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCommit struct {
	ID             string                                                          `json:"id,required" format:"uuid"`
	AccessSchedule V2ContractGetEditHistoryResponseDataUpdateCommitsAccessSchedule `json:"access_schedule"`
	// Which products the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,nullable" format:"uuid"`
	// Which tags the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []string                                                         `json:"applicable_product_tags,nullable"`
	InvoiceSchedule       V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceSchedule `json:"invoice_schedule"`
	Name                  string                                                           `json:"name"`
	NetsuiteSalesOrderID  string                                                           `json:"netsuite_sales_order_id,nullable"`
	ProductID             string                                                           `json:"product_id" format:"uuid"`
	RolloverFraction      float64                                                          `json:"rollover_fraction,nullable"`
	JSON                  v2ContractGetEditHistoryResponseDataUpdateCommitJSON             `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitJSON contains the JSON metadata
// for the struct [V2ContractGetEditHistoryResponseDataUpdateCommit]
type v2ContractGetEditHistoryResponseDataUpdateCommitJSON struct {
	ID                    apijson.Field
	AccessSchedule        apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	InvoiceSchedule       apijson.Field
	Name                  apijson.Field
	NetsuiteSalesOrderID  apijson.Field
	ProductID             apijson.Field
	RolloverFraction      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsAccessSchedule struct {
	AddScheduleItems    []V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleUpdateScheduleItem `json:"update_schedule_items"`
	JSON                v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleJSON                 `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleJSON contains the
// JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsAccessSchedule]
type v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleJSON struct {
	AddScheduleItems    apijson.Field
	RemoveScheduleItems apijson.Field
	UpdateScheduleItems apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleAddScheduleItem struct {
	Amount float64 `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time                                                                          `json:"starting_at,required" format:"date-time"`
	JSON       v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleAddScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleAddScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleAddScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleAddScheduleItemJSON struct {
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleAddScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleAddScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleRemoveScheduleItem struct {
	ID   string                                                                                `json:"id,required" format:"uuid"`
	JSON v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleRemoveScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleRemoveScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleRemoveScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleRemoveScheduleItemJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleRemoveScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleUpdateScheduleItem struct {
	ID     string  `json:"id,required" format:"uuid"`
	Amount float64 `json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time                                                                             `json:"starting_at" format:"date-time"`
	JSON       v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleUpdateScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleUpdateScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleUpdateScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleUpdateScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsAccessScheduleUpdateScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceSchedule struct {
	AddScheduleItems    []V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleUpdateScheduleItem `json:"update_schedule_items"`
	JSON                v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleJSON                 `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleJSON contains
// the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceSchedule]
type v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleJSON struct {
	AddScheduleItems    apijson.Field
	RemoveScheduleItems apijson.Field
	UpdateScheduleItems apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleAddScheduleItem struct {
	Timestamp time.Time                                                                           `json:"timestamp,required" format:"date-time"`
	Amount    float64                                                                             `json:"amount"`
	Quantity  float64                                                                             `json:"quantity"`
	UnitPrice float64                                                                             `json:"unit_price"`
	JSON      v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleAddScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleAddScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleAddScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleAddScheduleItemJSON struct {
	Timestamp   apijson.Field
	Amount      apijson.Field
	Quantity    apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleAddScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleAddScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleRemoveScheduleItem struct {
	ID   string                                                                                 `json:"id,required" format:"uuid"`
	JSON v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleRemoveScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleRemoveScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleRemoveScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleRemoveScheduleItemJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleRemoveScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleUpdateScheduleItem struct {
	ID        string                                                                                 `json:"id,required" format:"uuid"`
	Amount    float64                                                                                `json:"amount"`
	Quantity  float64                                                                                `json:"quantity"`
	Timestamp time.Time                                                                              `json:"timestamp" format:"date-time"`
	UnitPrice float64                                                                                `json:"unit_price"`
	JSON      v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleUpdateScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleUpdateScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleUpdateScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleUpdateScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceScheduleUpdateScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCredit struct {
	ID                   string                                                          `json:"id,required" format:"uuid"`
	AccessSchedule       V2ContractGetEditHistoryResponseDataUpdateCreditsAccessSchedule `json:"access_schedule"`
	Name                 string                                                          `json:"name"`
	NetsuiteSalesOrderID string                                                          `json:"netsuite_sales_order_id,nullable"`
	RolloverFraction     float64                                                         `json:"rollover_fraction,nullable"`
	JSON                 v2ContractGetEditHistoryResponseDataUpdateCreditJSON            `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCreditJSON contains the JSON metadata
// for the struct [V2ContractGetEditHistoryResponseDataUpdateCredit]
type v2ContractGetEditHistoryResponseDataUpdateCreditJSON struct {
	ID                   apijson.Field
	AccessSchedule       apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	RolloverFraction     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCreditJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsAccessSchedule struct {
	AddScheduleItems    []V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleUpdateScheduleItem `json:"update_schedule_items"`
	JSON                v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleJSON                 `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleJSON contains the
// JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCreditsAccessSchedule]
type v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleJSON struct {
	AddScheduleItems    apijson.Field
	RemoveScheduleItems apijson.Field
	UpdateScheduleItems apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCreditsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleAddScheduleItem struct {
	Amount float64 `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time                                                                          `json:"starting_at,required" format:"date-time"`
	JSON       v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleAddScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleAddScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleAddScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleAddScheduleItemJSON struct {
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleAddScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleAddScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleRemoveScheduleItem struct {
	ID   string                                                                                `json:"id,required" format:"uuid"`
	JSON v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleRemoveScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleRemoveScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleRemoveScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleRemoveScheduleItemJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleRemoveScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleUpdateScheduleItem struct {
	ID     string  `json:"id,required" format:"uuid"`
	Amount float64 `json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time                                                                             `json:"starting_at" format:"date-time"`
	JSON       v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleUpdateScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleUpdateScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleUpdateScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleUpdateScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCreditsAccessScheduleUpdateScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateDiscount struct {
	ID                   string            `json:"id,required" format:"uuid"`
	CustomFields         map[string]string `json:"custom_fields"`
	Name                 string            `json:"name"`
	NetsuiteSalesOrderID string            `json:"netsuite_sales_order_id"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule V2ContractGetEditHistoryResponseDataUpdateDiscountsSchedule `json:"schedule"`
	JSON     v2ContractGetEditHistoryResponseDataUpdateDiscountJSON      `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateDiscountJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataUpdateDiscount]
type v2ContractGetEditHistoryResponseDataUpdateDiscountJSON struct {
	ID                   apijson.Field
	CustomFields         apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	Schedule             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateDiscountJSON) RawJSON() string {
	return r.raw
}

// Must provide either schedule_items or recurring_schedule.
type V2ContractGetEditHistoryResponseDataUpdateDiscountsSchedule struct {
	// Defaults to USD if not passed. Only USD is supported at this time.
	CreditTypeID string `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringSchedule `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleJSON           `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleJSON contains the
// JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateDiscountsSchedule]
type v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleJSON struct {
	CreditTypeID      apijson.Field
	RecurringSchedule apijson.Field
	ScheduleItems     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleJSON) RawJSON() string {
	return r.raw
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringSchedule struct {
	AmountDistribution V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistribution `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time                                                                             `json:"ending_before,required" format:"date-time"`
	Frequency    V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequency `json:"frequency,required"`
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
	UnitPrice float64                                                                          `json:"unit_price"`
	JSON      v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringSchedule]
type v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleJSON struct {
	AmountDistribution apijson.Field
	EndingBefore       apijson.Field
	Frequency          apijson.Field
	StartingAt         apijson.Field
	Amount             apijson.Field
	Quantity           apijson.Field
	UnitPrice          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistribution string

const (
	V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistributionDivided        V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistributionEach           V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistributionDivided, V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded, V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequency string

const (
	V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequencyMonthly    V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequencyQuarterly  V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequencySemiAnnual V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequencyAnnual     V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
	V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequencyWeekly     V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequency = "WEEKLY"
)

func (r V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequencyMonthly, V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequencyQuarterly, V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequencySemiAnnual, V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequencyAnnual, V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleRecurringScheduleFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleScheduleItem struct {
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
	UnitPrice float64                                                                     `json:"unit_price"`
	JSON      v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleScheduleItemJSON struct {
	Timestamp   apijson.Field
	Amount      apijson.Field
	Quantity    apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateDiscountsScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateRefundInvoice struct {
	Date      time.Time                                                   `json:"date,required" format:"date-time"`
	InvoiceID string                                                      `json:"invoice_id,required" format:"uuid"`
	JSON      v2ContractGetEditHistoryResponseDataUpdateRefundInvoiceJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateRefundInvoiceJSON contains the JSON
// metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateRefundInvoice]
type v2ContractGetEditHistoryResponseDataUpdateRefundInvoiceJSON struct {
	Date        apijson.Field
	InvoiceID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateRefundInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateRefundInvoiceJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateScheduledCharge struct {
	ID                   string                                                                    `json:"id,required" format:"uuid"`
	InvoiceSchedule      V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceSchedule `json:"invoice_schedule"`
	Name                 string                                                                    `json:"name"`
	NetsuiteSalesOrderID string                                                                    `json:"netsuite_sales_order_id,nullable"`
	JSON                 v2ContractGetEditHistoryResponseDataUpdateScheduledChargeJSON             `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateScheduledChargeJSON contains the JSON
// metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateScheduledCharge]
type v2ContractGetEditHistoryResponseDataUpdateScheduledChargeJSON struct {
	ID                   apijson.Field
	InvoiceSchedule      apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateScheduledChargeJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceSchedule struct {
	AddScheduleItems    []V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleAddScheduleItem    `json:"add_schedule_items"`
	RemoveScheduleItems []V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleRemoveScheduleItem `json:"remove_schedule_items"`
	UpdateScheduleItems []V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleUpdateScheduleItem `json:"update_schedule_items"`
	JSON                v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleJSON                 `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceSchedule]
type v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleJSON struct {
	AddScheduleItems    apijson.Field
	RemoveScheduleItems apijson.Field
	UpdateScheduleItems apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleAddScheduleItem struct {
	Timestamp time.Time                                                                                    `json:"timestamp,required" format:"date-time"`
	Amount    float64                                                                                      `json:"amount"`
	Quantity  float64                                                                                      `json:"quantity"`
	UnitPrice float64                                                                                      `json:"unit_price"`
	JSON      v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleAddScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleAddScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleAddScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleAddScheduleItemJSON struct {
	Timestamp   apijson.Field
	Amount      apijson.Field
	Quantity    apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleAddScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleAddScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleRemoveScheduleItem struct {
	ID   string                                                                                          `json:"id,required" format:"uuid"`
	JSON v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleRemoveScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleRemoveScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleRemoveScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleRemoveScheduleItemJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleRemoveScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleRemoveScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleUpdateScheduleItem struct {
	ID        string                                                                                          `json:"id,required" format:"uuid"`
	Amount    float64                                                                                         `json:"amount"`
	Quantity  float64                                                                                         `json:"quantity"`
	Timestamp time.Time                                                                                       `json:"timestamp" format:"date-time"`
	UnitPrice float64                                                                                         `json:"unit_price"`
	JSON      v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleUpdateScheduleItemJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleUpdateScheduleItemJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleUpdateScheduleItem]
type v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleUpdateScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleUpdateScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateScheduledChargesInvoiceScheduleUpdateScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetParams struct {
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Optional RFC 3339 timestamp. Return the contract as of this date. Cannot be used
	// with include_ledgers parameter.
	AsOfDate param.Field[time.Time] `json:"as_of_date" format:"date-time"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the query to be slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include commit/credit ledgers in the response. Setting this flag may cause the
	// query to be slower. Cannot be used with as_of_date parameter.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
}

func (r V2ContractGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractListParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Optional RFC 3339 timestamp. Only include contracts active on the provided date.
	// This cannot be provided if starting_at filter is provided.
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	// Include archived contracts in the response.
	IncludeArchived param.Field[bool] `json:"include_archived"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the response to be slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include commit/credit ledgers in the response. Setting this flag may cause the
	// response to be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
	// Optional RFC 3339 timestamp. Only include contracts that started on or after
	// this date. This cannot be provided if covering_date filter is provided.
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V2ContractListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParams struct {
	// ID of the contract being edited
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is being edited
	CustomerID   param.Field[string]                            `json:"customer_id,required" format:"uuid"`
	AddCommits   param.Field[[]V2ContractEditParamsAddCommit]   `json:"add_commits"`
	AddCredits   param.Field[[]V2ContractEditParamsAddCredit]   `json:"add_credits"`
	AddDiscounts param.Field[[]V2ContractEditParamsAddDiscount] `json:"add_discounts"`
	AddOverrides param.Field[[]V2ContractEditParamsAddOverride] `json:"add_overrides"`
	// This field's availability is dependent on your client's configuration.
	AddProfessionalServices param.Field[[]V2ContractEditParamsAddProfessionalService] `json:"add_professional_services"`
	AddRecurringCommits     param.Field[[]V2ContractEditParamsAddRecurringCommit]     `json:"add_recurring_commits"`
	AddRecurringCredits     param.Field[[]V2ContractEditParamsAddRecurringCredit]     `json:"add_recurring_credits"`
	AddResellerRoyalties    param.Field[[]V2ContractEditParamsAddResellerRoyalty]     `json:"add_reseller_royalties"`
	AddScheduledCharges     param.Field[[]V2ContractEditParamsAddScheduledCharge]     `json:"add_scheduled_charges"`
	UpdateCommits           param.Field[[]V2ContractEditParamsUpdateCommit]           `json:"update_commits"`
	UpdateCredits           param.Field[[]V2ContractEditParamsUpdateCredit]           `json:"update_credits"`
	UpdateScheduledCharges  param.Field[[]V2ContractEditParamsUpdateScheduledCharge]  `json:"update_scheduled_charges"`
}

func (r V2ContractEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddCommit struct {
	ProductID param.Field[string]                             `json:"product_id,required" format:"uuid"`
	Type      param.Field[V2ContractEditParamsAddCommitsType] `json:"type,required"`
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
	// commits only one schedule item is allowed and amount must match invoice_schedule
	// total.
	AccessSchedule param.Field[V2ContractEditParamsAddCommitsAccessSchedule] `json:"access_schedule"`
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
	InvoiceSchedule param.Field[V2ContractEditParamsAddCommitsInvoiceSchedule] `json:"invoice_schedule"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Field[float64]                                `json:"priority"`
	RateType param.Field[V2ContractEditParamsAddCommitsRateType] `json:"rate_type"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// A temporary ID for the commit that can be used to reference the commit for
	// commit specific overrides.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r V2ContractEditParamsAddCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddCommitsType string

const (
	V2ContractEditParamsAddCommitsTypePrepaid  V2ContractEditParamsAddCommitsType = "PREPAID"
	V2ContractEditParamsAddCommitsTypePostpaid V2ContractEditParamsAddCommitsType = "POSTPAID"
)

func (r V2ContractEditParamsAddCommitsType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCommitsTypePrepaid, V2ContractEditParamsAddCommitsTypePostpaid:
		return true
	}
	return false
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
type V2ContractEditParamsAddCommitsAccessSchedule struct {
	ScheduleItems param.Field[[]V2ContractEditParamsAddCommitsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	CreditTypeID  param.Field[string]                                                     `json:"credit_type_id" format:"uuid"`
}

func (r V2ContractEditParamsAddCommitsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddCommitsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V2ContractEditParamsAddCommitsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match access_schedule
// amount. Optional for "PREPAID" commits: if not provided, this will be a
// "complimentary" commit with no invoice.
type V2ContractEditParamsAddCommitsInvoiceSchedule struct {
	// Defaults to USD if not passed. Only USD is supported at this time.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[V2ContractEditParamsAddCommitsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]V2ContractEditParamsAddCommitsInvoiceScheduleScheduleItem] `json:"schedule_items"`
}

func (r V2ContractEditParamsAddCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V2ContractEditParamsAddCommitsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                               `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r V2ContractEditParamsAddCommitsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
	V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided, V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
	V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
	V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequencyWeekly     V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequency = "WEEKLY"
)

func (r V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly, V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly, V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual, V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractEditParamsAddCommitsInvoiceScheduleScheduleItem struct {
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

func (r V2ContractEditParamsAddCommitsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddCommitsRateType string

const (
	V2ContractEditParamsAddCommitsRateTypeCommitRate V2ContractEditParamsAddCommitsRateType = "COMMIT_RATE"
	V2ContractEditParamsAddCommitsRateTypeListRate   V2ContractEditParamsAddCommitsRateType = "LIST_RATE"
)

func (r V2ContractEditParamsAddCommitsRateType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCommitsRateTypeCommitRate, V2ContractEditParamsAddCommitsRateTypeListRate:
		return true
	}
	return false
}

type V2ContractEditParamsAddCredit struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule param.Field[V2ContractEditParamsAddCreditsAccessSchedule] `json:"access_schedule,required"`
	ProductID      param.Field[string]                                       `json:"product_id,required" format:"uuid"`
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
	Priority param.Field[float64]                                `json:"priority"`
	RateType param.Field[V2ContractEditParamsAddCreditsRateType] `json:"rate_type"`
}

func (r V2ContractEditParamsAddCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the credit to the customer.
type V2ContractEditParamsAddCreditsAccessSchedule struct {
	ScheduleItems param.Field[[]V2ContractEditParamsAddCreditsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	CreditTypeID  param.Field[string]                                                     `json:"credit_type_id" format:"uuid"`
}

func (r V2ContractEditParamsAddCreditsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddCreditsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V2ContractEditParamsAddCreditsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddCreditsRateType string

const (
	V2ContractEditParamsAddCreditsRateTypeCommitRate V2ContractEditParamsAddCreditsRateType = "COMMIT_RATE"
	V2ContractEditParamsAddCreditsRateTypeListRate   V2ContractEditParamsAddCreditsRateType = "LIST_RATE"
)

func (r V2ContractEditParamsAddCreditsRateType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCreditsRateTypeCommitRate, V2ContractEditParamsAddCreditsRateTypeListRate:
		return true
	}
	return false
}

type V2ContractEditParamsAddDiscount struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule     param.Field[V2ContractEditParamsAddDiscountsSchedule] `json:"schedule,required"`
	CustomFields param.Field[map[string]string]                        `json:"custom_fields"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r V2ContractEditParamsAddDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule.
type V2ContractEditParamsAddDiscountsSchedule struct {
	// Defaults to USD if not passed. Only USD is supported at this time.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[V2ContractEditParamsAddDiscountsScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]V2ContractEditParamsAddDiscountsScheduleScheduleItem] `json:"schedule_items"`
}

func (r V2ContractEditParamsAddDiscountsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V2ContractEditParamsAddDiscountsScheduleRecurringSchedule struct {
	AmountDistribution param.Field[V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                          `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r V2ContractEditParamsAddDiscountsScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistribution string

const (
	V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistributionDivided        V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistributionEach           V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistributionDivided, V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded, V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequency string

const (
	V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequencyMonthly    V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequencyQuarterly  V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequencySemiAnnual V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequencyAnnual     V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
	V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequencyWeekly     V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequency = "WEEKLY"
)

func (r V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequencyMonthly, V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequencyQuarterly, V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequencySemiAnnual, V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequencyAnnual, V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractEditParamsAddDiscountsScheduleScheduleItem struct {
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

func (r V2ContractEditParamsAddDiscountsScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// tags identifying products whose rates are being overridden
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Entitled     param.Field[bool]      `json:"entitled"`
	// Indicates whether the override should only apply to commits. Defaults to
	// `false`. If `true`, you can specify relevant commits in `override_specifiers` by
	// passing `commit_ids`.
	IsCommitSpecific param.Field[bool] `json:"is_commit_specific"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Field[float64] `json:"multiplier"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If
	// provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers param.Field[[]V2ContractEditParamsAddOverridesOverrideSpecifier] `json:"override_specifiers"`
	// Required for OVERWRITE type.
	OverwriteRate param.Field[V2ContractEditParamsAddOverridesOverwriteRate] `json:"overwrite_rate"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides.
	// Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered
	// and multiplier overrides are prioritized by their priority value (lowest first).
	// Must be > 0.
	Priority param.Field[float64] `json:"priority"`
	// ID of the product whose rate is being overridden
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// Indicates whether the override applies to commit rates or list rates. Can only
	// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
	// `"LIST_RATE"`.
	Target param.Field[V2ContractEditParamsAddOverridesTarget] `json:"target"`
	// Required for TIERED type. Must have at least one tier.
	Tiers param.Field[[]V2ContractEditParamsAddOverridesTier] `json:"tiers"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type param.Field[V2ContractEditParamsAddOverridesType] `json:"type"`
}

func (r V2ContractEditParamsAddOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddOverridesOverrideSpecifier struct {
	// If provided, the override will only apply to the specified commits. Can only be
	// used for commit specific overrides. If not provided, the override will apply to
	// all commits.
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
	// presentation_group_values. If provided, the override will only apply to commits
	// created by the specified recurring credit ids.
	RecurringCreditIDs param.Field[[]string] `json:"recurring_credit_ids"`
}

func (r V2ContractEditParamsAddOverridesOverrideSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required for OVERWRITE type.
type V2ContractEditParamsAddOverridesOverwriteRate struct {
	RateType     param.Field[V2ContractEditParamsAddOverridesOverwriteRateRateType] `json:"rate_type,required"`
	CreditTypeID param.Field[string]                                                `json:"credit_type_id" format:"uuid"`
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

func (r V2ContractEditParamsAddOverridesOverwriteRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddOverridesOverwriteRateRateType string

const (
	V2ContractEditParamsAddOverridesOverwriteRateRateTypeFlat         V2ContractEditParamsAddOverridesOverwriteRateRateType = "FLAT"
	V2ContractEditParamsAddOverridesOverwriteRateRateTypePercentage   V2ContractEditParamsAddOverridesOverwriteRateRateType = "PERCENTAGE"
	V2ContractEditParamsAddOverridesOverwriteRateRateTypeSubscription V2ContractEditParamsAddOverridesOverwriteRateRateType = "SUBSCRIPTION"
	V2ContractEditParamsAddOverridesOverwriteRateRateTypeTiered       V2ContractEditParamsAddOverridesOverwriteRateRateType = "TIERED"
	V2ContractEditParamsAddOverridesOverwriteRateRateTypeCustom       V2ContractEditParamsAddOverridesOverwriteRateRateType = "CUSTOM"
)

func (r V2ContractEditParamsAddOverridesOverwriteRateRateType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddOverridesOverwriteRateRateTypeFlat, V2ContractEditParamsAddOverridesOverwriteRateRateTypePercentage, V2ContractEditParamsAddOverridesOverwriteRateRateTypeSubscription, V2ContractEditParamsAddOverridesOverwriteRateRateTypeTiered, V2ContractEditParamsAddOverridesOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

// Indicates whether the override applies to commit rates or list rates. Can only
// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
// `"LIST_RATE"`.
type V2ContractEditParamsAddOverridesTarget string

const (
	V2ContractEditParamsAddOverridesTargetCommitRate V2ContractEditParamsAddOverridesTarget = "COMMIT_RATE"
	V2ContractEditParamsAddOverridesTargetListRate   V2ContractEditParamsAddOverridesTarget = "LIST_RATE"
)

func (r V2ContractEditParamsAddOverridesTarget) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddOverridesTargetCommitRate, V2ContractEditParamsAddOverridesTargetListRate:
		return true
	}
	return false
}

type V2ContractEditParamsAddOverridesTier struct {
	Multiplier param.Field[float64] `json:"multiplier,required"`
	Size       param.Field[float64] `json:"size"`
}

func (r V2ContractEditParamsAddOverridesTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Overwrites are prioritized over multipliers and tiered overrides.
type V2ContractEditParamsAddOverridesType string

const (
	V2ContractEditParamsAddOverridesTypeOverwrite  V2ContractEditParamsAddOverridesType = "OVERWRITE"
	V2ContractEditParamsAddOverridesTypeMultiplier V2ContractEditParamsAddOverridesType = "MULTIPLIER"
	V2ContractEditParamsAddOverridesTypeTiered     V2ContractEditParamsAddOverridesType = "TIERED"
)

func (r V2ContractEditParamsAddOverridesType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddOverridesTypeOverwrite, V2ContractEditParamsAddOverridesTypeMultiplier, V2ContractEditParamsAddOverridesTypeTiered:
		return true
	}
	return false
}

type V2ContractEditParamsAddProfessionalService struct {
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

func (r V2ContractEditParamsAddProfessionalService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddRecurringCommit struct {
	// The amount of commit to grant.
	AccessAmount param.Field[V2ContractEditParamsAddRecurringCommitsAccessAmount] `json:"access_amount,required"`
	// The amount of time the created commits will be valid for.
	CommitDuration param.Field[V2ContractEditParamsAddRecurringCommitsCommitDuration] `json:"commit_duration,required"`
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
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount param.Field[V2ContractEditParamsAddRecurringCommitsInvoiceAmount] `json:"invoice_amount"`
	// displayed on invoices. will be passed through to the individual commits
	Name param.Field[string] `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration param.Field[V2ContractEditParamsAddRecurringCommitsProration] `json:"proration"`
	// Whether the created commits will use the commit rate or list rate
	RateType param.Field[V2ContractEditParamsAddRecurringCommitsRateType] `json:"rate_type"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency param.Field[V2ContractEditParamsAddRecurringCommitsRecurrenceFrequency] `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// A temporary ID that can be used to reference the recurring commit for commit
	// specific overrides.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r V2ContractEditParamsAddRecurringCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of commit to grant.
type V2ContractEditParamsAddRecurringCommitsAccessAmount struct {
	CreditTypeID param.Field[string]  `json:"credit_type_id,required" format:"uuid"`
	Quantity     param.Field[float64] `json:"quantity,required"`
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
}

func (r V2ContractEditParamsAddRecurringCommitsAccessAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of time the created commits will be valid for.
type V2ContractEditParamsAddRecurringCommitsCommitDuration struct {
	Unit  param.Field[V2ContractEditParamsAddRecurringCommitsCommitDurationUnit] `json:"unit,required"`
	Value param.Field[float64]                                                   `json:"value,required"`
}

func (r V2ContractEditParamsAddRecurringCommitsCommitDuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddRecurringCommitsCommitDurationUnit string

const (
	V2ContractEditParamsAddRecurringCommitsCommitDurationUnitPeriods V2ContractEditParamsAddRecurringCommitsCommitDurationUnit = "PERIODS"
)

func (r V2ContractEditParamsAddRecurringCommitsCommitDurationUnit) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCommitsCommitDurationUnitPeriods:
		return true
	}
	return false
}

// The amount the customer should be billed for the commit. Not required.
type V2ContractEditParamsAddRecurringCommitsInvoiceAmount struct {
	CreditTypeID param.Field[string]  `json:"credit_type_id,required" format:"uuid"`
	Quantity     param.Field[float64] `json:"quantity,required"`
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
}

func (r V2ContractEditParamsAddRecurringCommitsInvoiceAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type V2ContractEditParamsAddRecurringCommitsProration string

const (
	V2ContractEditParamsAddRecurringCommitsProrationNone         V2ContractEditParamsAddRecurringCommitsProration = "NONE"
	V2ContractEditParamsAddRecurringCommitsProrationFirst        V2ContractEditParamsAddRecurringCommitsProration = "FIRST"
	V2ContractEditParamsAddRecurringCommitsProrationLast         V2ContractEditParamsAddRecurringCommitsProration = "LAST"
	V2ContractEditParamsAddRecurringCommitsProrationFirstAndLast V2ContractEditParamsAddRecurringCommitsProration = "FIRST_AND_LAST"
)

func (r V2ContractEditParamsAddRecurringCommitsProration) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCommitsProrationNone, V2ContractEditParamsAddRecurringCommitsProrationFirst, V2ContractEditParamsAddRecurringCommitsProrationLast, V2ContractEditParamsAddRecurringCommitsProrationFirstAndLast:
		return true
	}
	return false
}

// Whether the created commits will use the commit rate or list rate
type V2ContractEditParamsAddRecurringCommitsRateType string

const (
	V2ContractEditParamsAddRecurringCommitsRateTypeCommitRate V2ContractEditParamsAddRecurringCommitsRateType = "COMMIT_RATE"
	V2ContractEditParamsAddRecurringCommitsRateTypeListRate   V2ContractEditParamsAddRecurringCommitsRateType = "LIST_RATE"
)

func (r V2ContractEditParamsAddRecurringCommitsRateType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCommitsRateTypeCommitRate, V2ContractEditParamsAddRecurringCommitsRateTypeListRate:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type V2ContractEditParamsAddRecurringCommitsRecurrenceFrequency string

const (
	V2ContractEditParamsAddRecurringCommitsRecurrenceFrequencyMonthly   V2ContractEditParamsAddRecurringCommitsRecurrenceFrequency = "MONTHLY"
	V2ContractEditParamsAddRecurringCommitsRecurrenceFrequencyQuarterly V2ContractEditParamsAddRecurringCommitsRecurrenceFrequency = "QUARTERLY"
	V2ContractEditParamsAddRecurringCommitsRecurrenceFrequencyAnnual    V2ContractEditParamsAddRecurringCommitsRecurrenceFrequency = "ANNUAL"
	V2ContractEditParamsAddRecurringCommitsRecurrenceFrequencyWeekly    V2ContractEditParamsAddRecurringCommitsRecurrenceFrequency = "WEEKLY"
)

func (r V2ContractEditParamsAddRecurringCommitsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCommitsRecurrenceFrequencyMonthly, V2ContractEditParamsAddRecurringCommitsRecurrenceFrequencyQuarterly, V2ContractEditParamsAddRecurringCommitsRecurrenceFrequencyAnnual, V2ContractEditParamsAddRecurringCommitsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractEditParamsAddRecurringCredit struct {
	// The amount of commit to grant.
	AccessAmount param.Field[V2ContractEditParamsAddRecurringCreditsAccessAmount] `json:"access_amount,required"`
	// The amount of time the created commits will be valid for.
	CommitDuration param.Field[V2ContractEditParamsAddRecurringCreditsCommitDuration] `json:"commit_duration,required"`
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
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// displayed on invoices. will be passed through to the individual commits
	Name param.Field[string] `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration param.Field[V2ContractEditParamsAddRecurringCreditsProration] `json:"proration"`
	// Whether the created commits will use the commit rate or list rate
	RateType param.Field[V2ContractEditParamsAddRecurringCreditsRateType] `json:"rate_type"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency param.Field[V2ContractEditParamsAddRecurringCreditsRecurrenceFrequency] `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// A temporary ID that can be used to reference the recurring commit for commit
	// specific overrides.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r V2ContractEditParamsAddRecurringCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of commit to grant.
type V2ContractEditParamsAddRecurringCreditsAccessAmount struct {
	CreditTypeID param.Field[string]  `json:"credit_type_id,required" format:"uuid"`
	Quantity     param.Field[float64] `json:"quantity,required"`
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
}

func (r V2ContractEditParamsAddRecurringCreditsAccessAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of time the created commits will be valid for.
type V2ContractEditParamsAddRecurringCreditsCommitDuration struct {
	Unit  param.Field[V2ContractEditParamsAddRecurringCreditsCommitDurationUnit] `json:"unit,required"`
	Value param.Field[float64]                                                   `json:"value,required"`
}

func (r V2ContractEditParamsAddRecurringCreditsCommitDuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddRecurringCreditsCommitDurationUnit string

const (
	V2ContractEditParamsAddRecurringCreditsCommitDurationUnitPeriods V2ContractEditParamsAddRecurringCreditsCommitDurationUnit = "PERIODS"
)

func (r V2ContractEditParamsAddRecurringCreditsCommitDurationUnit) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCreditsCommitDurationUnitPeriods:
		return true
	}
	return false
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type V2ContractEditParamsAddRecurringCreditsProration string

const (
	V2ContractEditParamsAddRecurringCreditsProrationNone         V2ContractEditParamsAddRecurringCreditsProration = "NONE"
	V2ContractEditParamsAddRecurringCreditsProrationFirst        V2ContractEditParamsAddRecurringCreditsProration = "FIRST"
	V2ContractEditParamsAddRecurringCreditsProrationLast         V2ContractEditParamsAddRecurringCreditsProration = "LAST"
	V2ContractEditParamsAddRecurringCreditsProrationFirstAndLast V2ContractEditParamsAddRecurringCreditsProration = "FIRST_AND_LAST"
)

func (r V2ContractEditParamsAddRecurringCreditsProration) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCreditsProrationNone, V2ContractEditParamsAddRecurringCreditsProrationFirst, V2ContractEditParamsAddRecurringCreditsProrationLast, V2ContractEditParamsAddRecurringCreditsProrationFirstAndLast:
		return true
	}
	return false
}

// Whether the created commits will use the commit rate or list rate
type V2ContractEditParamsAddRecurringCreditsRateType string

const (
	V2ContractEditParamsAddRecurringCreditsRateTypeCommitRate V2ContractEditParamsAddRecurringCreditsRateType = "COMMIT_RATE"
	V2ContractEditParamsAddRecurringCreditsRateTypeListRate   V2ContractEditParamsAddRecurringCreditsRateType = "LIST_RATE"
)

func (r V2ContractEditParamsAddRecurringCreditsRateType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCreditsRateTypeCommitRate, V2ContractEditParamsAddRecurringCreditsRateTypeListRate:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type V2ContractEditParamsAddRecurringCreditsRecurrenceFrequency string

const (
	V2ContractEditParamsAddRecurringCreditsRecurrenceFrequencyMonthly   V2ContractEditParamsAddRecurringCreditsRecurrenceFrequency = "MONTHLY"
	V2ContractEditParamsAddRecurringCreditsRecurrenceFrequencyQuarterly V2ContractEditParamsAddRecurringCreditsRecurrenceFrequency = "QUARTERLY"
	V2ContractEditParamsAddRecurringCreditsRecurrenceFrequencyAnnual    V2ContractEditParamsAddRecurringCreditsRecurrenceFrequency = "ANNUAL"
	V2ContractEditParamsAddRecurringCreditsRecurrenceFrequencyWeekly    V2ContractEditParamsAddRecurringCreditsRecurrenceFrequency = "WEEKLY"
)

func (r V2ContractEditParamsAddRecurringCreditsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCreditsRecurrenceFrequencyMonthly, V2ContractEditParamsAddRecurringCreditsRecurrenceFrequencyQuarterly, V2ContractEditParamsAddRecurringCreditsRecurrenceFrequencyAnnual, V2ContractEditParamsAddRecurringCreditsRecurrenceFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractEditParamsAddResellerRoyalty struct {
	ResellerType param.Field[V2ContractEditParamsAddResellerRoyaltiesResellerType] `json:"reseller_type,required"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags param.Field[[]string]                                           `json:"applicable_product_tags"`
	AwsOptions            param.Field[V2ContractEditParamsAddResellerRoyaltiesAwsOptions] `json:"aws_options"`
	// Use null to indicate that the existing end timestamp should be removed.
	EndingBefore          param.Field[time.Time]                                          `json:"ending_before" format:"date-time"`
	Fraction              param.Field[float64]                                            `json:"fraction"`
	GcpOptions            param.Field[V2ContractEditParamsAddResellerRoyaltiesGcpOptions] `json:"gcp_options"`
	NetsuiteResellerID    param.Field[string]                                             `json:"netsuite_reseller_id"`
	ResellerContractValue param.Field[float64]                                            `json:"reseller_contract_value"`
	StartingAt            param.Field[time.Time]                                          `json:"starting_at" format:"date-time"`
}

func (r V2ContractEditParamsAddResellerRoyalty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddResellerRoyaltiesResellerType string

const (
	V2ContractEditParamsAddResellerRoyaltiesResellerTypeAws           V2ContractEditParamsAddResellerRoyaltiesResellerType = "AWS"
	V2ContractEditParamsAddResellerRoyaltiesResellerTypeAwsProService V2ContractEditParamsAddResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	V2ContractEditParamsAddResellerRoyaltiesResellerTypeGcp           V2ContractEditParamsAddResellerRoyaltiesResellerType = "GCP"
	V2ContractEditParamsAddResellerRoyaltiesResellerTypeGcpProService V2ContractEditParamsAddResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r V2ContractEditParamsAddResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddResellerRoyaltiesResellerTypeAws, V2ContractEditParamsAddResellerRoyaltiesResellerTypeAwsProService, V2ContractEditParamsAddResellerRoyaltiesResellerTypeGcp, V2ContractEditParamsAddResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

type V2ContractEditParamsAddResellerRoyaltiesAwsOptions struct {
	AwsAccountNumber    param.Field[string] `json:"aws_account_number"`
	AwsOfferID          param.Field[string] `json:"aws_offer_id"`
	AwsPayerReferenceID param.Field[string] `json:"aws_payer_reference_id"`
}

func (r V2ContractEditParamsAddResellerRoyaltiesAwsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddResellerRoyaltiesGcpOptions struct {
	GcpAccountID param.Field[string] `json:"gcp_account_id"`
	GcpOfferID   param.Field[string] `json:"gcp_offer_id"`
}

func (r V2ContractEditParamsAddResellerRoyaltiesGcpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddScheduledCharge struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule param.Field[V2ContractEditParamsAddScheduledChargesSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r V2ContractEditParamsAddScheduledCharge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule.
type V2ContractEditParamsAddScheduledChargesSchedule struct {
	// Defaults to USD if not passed. Only USD is supported at this time.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[V2ContractEditParamsAddScheduledChargesScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]V2ContractEditParamsAddScheduledChargesScheduleScheduleItem] `json:"schedule_items"`
}

func (r V2ContractEditParamsAddScheduledChargesSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V2ContractEditParamsAddScheduledChargesScheduleRecurringSchedule struct {
	AmountDistribution param.Field[V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                                 `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r V2ContractEditParamsAddScheduledChargesScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistribution string

const (
	V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistributionEach           V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistributionDivided, V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded, V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequency string

const (
	V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequencyMonthly    V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
	V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
	V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequencyAnnual     V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
	V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequencyWeekly     V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequency = "WEEKLY"
)

func (r V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequencyMonthly, V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequencyQuarterly, V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual, V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequencyAnnual, V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractEditParamsAddScheduledChargesScheduleScheduleItem struct {
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

func (r V2ContractEditParamsAddScheduledChargesScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCommit struct {
	CommitID       param.Field[string]                                          `json:"commit_id,required" format:"uuid"`
	AccessSchedule param.Field[V2ContractEditParamsUpdateCommitsAccessSchedule] `json:"access_schedule"`
	// Which products the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags param.Field[[]string]                                         `json:"applicable_product_tags"`
	InvoiceSchedule       param.Field[V2ContractEditParamsUpdateCommitsInvoiceSchedule] `json:"invoice_schedule"`
	NetsuiteSalesOrderID  param.Field[string]                                           `json:"netsuite_sales_order_id"`
	ProductID             param.Field[string]                                           `json:"product_id" format:"uuid"`
	RolloverFraction      param.Field[float64]                                          `json:"rollover_fraction"`
}

func (r V2ContractEditParamsUpdateCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCommitsAccessSchedule struct {
	AddScheduleItems    param.Field[[]V2ContractEditParamsUpdateCommitsAccessScheduleAddScheduleItem]    `json:"add_schedule_items"`
	RemoveScheduleItems param.Field[[]V2ContractEditParamsUpdateCommitsAccessScheduleRemoveScheduleItem] `json:"remove_schedule_items"`
	UpdateScheduleItems param.Field[[]V2ContractEditParamsUpdateCommitsAccessScheduleUpdateScheduleItem] `json:"update_schedule_items"`
}

func (r V2ContractEditParamsUpdateCommitsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCommitsAccessScheduleAddScheduleItem struct {
	Amount       param.Field[float64]   `json:"amount,required"`
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V2ContractEditParamsUpdateCommitsAccessScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCommitsAccessScheduleRemoveScheduleItem struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V2ContractEditParamsUpdateCommitsAccessScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCommitsAccessScheduleUpdateScheduleItem struct {
	ID           param.Field[string]    `json:"id,required" format:"uuid"`
	Amount       param.Field[float64]   `json:"amount"`
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V2ContractEditParamsUpdateCommitsAccessScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCommitsInvoiceSchedule struct {
	AddScheduleItems    param.Field[[]V2ContractEditParamsUpdateCommitsInvoiceScheduleAddScheduleItem]    `json:"add_schedule_items"`
	RemoveScheduleItems param.Field[[]V2ContractEditParamsUpdateCommitsInvoiceScheduleRemoveScheduleItem] `json:"remove_schedule_items"`
	UpdateScheduleItems param.Field[[]V2ContractEditParamsUpdateCommitsInvoiceScheduleUpdateScheduleItem] `json:"update_schedule_items"`
}

func (r V2ContractEditParamsUpdateCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCommitsInvoiceScheduleAddScheduleItem struct {
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r V2ContractEditParamsUpdateCommitsInvoiceScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCommitsInvoiceScheduleRemoveScheduleItem struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V2ContractEditParamsUpdateCommitsInvoiceScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCommitsInvoiceScheduleUpdateScheduleItem struct {
	ID        param.Field[string]    `json:"id,required" format:"uuid"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	Timestamp param.Field[time.Time] `json:"timestamp" format:"date-time"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r V2ContractEditParamsUpdateCommitsInvoiceScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCredit struct {
	CreditID       param.Field[string]                                          `json:"credit_id,required" format:"uuid"`
	AccessSchedule param.Field[V2ContractEditParamsUpdateCreditsAccessSchedule] `json:"access_schedule"`
	// Which products the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	NetsuiteSalesOrderID  param.Field[string]   `json:"netsuite_sales_order_id"`
	ProductID             param.Field[string]   `json:"product_id" format:"uuid"`
}

func (r V2ContractEditParamsUpdateCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCreditsAccessSchedule struct {
	AddScheduleItems    param.Field[[]V2ContractEditParamsUpdateCreditsAccessScheduleAddScheduleItem]    `json:"add_schedule_items"`
	RemoveScheduleItems param.Field[[]V2ContractEditParamsUpdateCreditsAccessScheduleRemoveScheduleItem] `json:"remove_schedule_items"`
	UpdateScheduleItems param.Field[[]V2ContractEditParamsUpdateCreditsAccessScheduleUpdateScheduleItem] `json:"update_schedule_items"`
}

func (r V2ContractEditParamsUpdateCreditsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCreditsAccessScheduleAddScheduleItem struct {
	Amount       param.Field[float64]   `json:"amount,required"`
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V2ContractEditParamsUpdateCreditsAccessScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCreditsAccessScheduleRemoveScheduleItem struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V2ContractEditParamsUpdateCreditsAccessScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCreditsAccessScheduleUpdateScheduleItem struct {
	ID           param.Field[string]    `json:"id,required" format:"uuid"`
	Amount       param.Field[float64]   `json:"amount"`
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V2ContractEditParamsUpdateCreditsAccessScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateScheduledCharge struct {
	ScheduledChargeID    param.Field[string]                                                    `json:"scheduled_charge_id,required" format:"uuid"`
	InvoiceSchedule      param.Field[V2ContractEditParamsUpdateScheduledChargesInvoiceSchedule] `json:"invoice_schedule"`
	NetsuiteSalesOrderID param.Field[string]                                                    `json:"netsuite_sales_order_id"`
}

func (r V2ContractEditParamsUpdateScheduledCharge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateScheduledChargesInvoiceSchedule struct {
	AddScheduleItems    param.Field[[]V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleAddScheduleItem]    `json:"add_schedule_items"`
	RemoveScheduleItems param.Field[[]V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleRemoveScheduleItem] `json:"remove_schedule_items"`
	UpdateScheduleItems param.Field[[]V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleUpdateScheduleItem] `json:"update_schedule_items"`
}

func (r V2ContractEditParamsUpdateScheduledChargesInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleAddScheduleItem struct {
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleRemoveScheduleItem struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleUpdateScheduleItem struct {
	ID        param.Field[string]    `json:"id,required" format:"uuid"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	Timestamp param.Field[time.Time] `json:"timestamp" format:"date-time"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCommitParams struct {
	// ID of the commit to edit
	CommitID param.Field[string] `json:"commit_id,required" format:"uuid"`
	// ID of the customer whose commit is being edited
	CustomerID     param.Field[string]                                   `json:"customer_id,required" format:"uuid"`
	AccessSchedule param.Field[V2ContractEditCommitParamsAccessSchedule] `json:"access_schedule"`
	// Which products the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// ID of contract to use for invoicing
	InvoiceContractID param.Field[string]                                    `json:"invoice_contract_id" format:"uuid"`
	InvoiceSchedule   param.Field[V2ContractEditCommitParamsInvoiceSchedule] `json:"invoice_schedule"`
	ProductID         param.Field[string]                                    `json:"product_id" format:"uuid"`
}

func (r V2ContractEditCommitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCommitParamsAccessSchedule struct {
	AddScheduleItems    param.Field[[]V2ContractEditCommitParamsAccessScheduleAddScheduleItem]    `json:"add_schedule_items"`
	RemoveScheduleItems param.Field[[]V2ContractEditCommitParamsAccessScheduleRemoveScheduleItem] `json:"remove_schedule_items"`
	UpdateScheduleItems param.Field[[]V2ContractEditCommitParamsAccessScheduleUpdateScheduleItem] `json:"update_schedule_items"`
}

func (r V2ContractEditCommitParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCommitParamsAccessScheduleAddScheduleItem struct {
	Amount       param.Field[float64]   `json:"amount,required"`
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V2ContractEditCommitParamsAccessScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCommitParamsAccessScheduleRemoveScheduleItem struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V2ContractEditCommitParamsAccessScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCommitParamsAccessScheduleUpdateScheduleItem struct {
	ID           param.Field[string]    `json:"id,required" format:"uuid"`
	Amount       param.Field[float64]   `json:"amount"`
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V2ContractEditCommitParamsAccessScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCommitParamsInvoiceSchedule struct {
	AddScheduleItems    param.Field[[]V2ContractEditCommitParamsInvoiceScheduleAddScheduleItem]    `json:"add_schedule_items"`
	RemoveScheduleItems param.Field[[]V2ContractEditCommitParamsInvoiceScheduleRemoveScheduleItem] `json:"remove_schedule_items"`
	UpdateScheduleItems param.Field[[]V2ContractEditCommitParamsInvoiceScheduleUpdateScheduleItem] `json:"update_schedule_items"`
}

func (r V2ContractEditCommitParamsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCommitParamsInvoiceScheduleAddScheduleItem struct {
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r V2ContractEditCommitParamsInvoiceScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCommitParamsInvoiceScheduleRemoveScheduleItem struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V2ContractEditCommitParamsInvoiceScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCommitParamsInvoiceScheduleUpdateScheduleItem struct {
	ID        param.Field[string]    `json:"id,required" format:"uuid"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	Timestamp param.Field[time.Time] `json:"timestamp" format:"date-time"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r V2ContractEditCommitParamsInvoiceScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCreditParams struct {
	// ID of the credit to edit
	CreditID param.Field[string] `json:"credit_id,required" format:"uuid"`
	// ID of the customer whose credit is being edited
	CustomerID     param.Field[string]                                   `json:"customer_id,required" format:"uuid"`
	AccessSchedule param.Field[V2ContractEditCreditParamsAccessSchedule] `json:"access_schedule"`
	// Which products the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	ProductID             param.Field[string]   `json:"product_id" format:"uuid"`
}

func (r V2ContractEditCreditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCreditParamsAccessSchedule struct {
	AddScheduleItems    param.Field[[]V2ContractEditCreditParamsAccessScheduleAddScheduleItem]    `json:"add_schedule_items"`
	RemoveScheduleItems param.Field[[]V2ContractEditCreditParamsAccessScheduleRemoveScheduleItem] `json:"remove_schedule_items"`
	UpdateScheduleItems param.Field[[]V2ContractEditCreditParamsAccessScheduleUpdateScheduleItem] `json:"update_schedule_items"`
}

func (r V2ContractEditCreditParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCreditParamsAccessScheduleAddScheduleItem struct {
	Amount       param.Field[float64]   `json:"amount,required"`
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V2ContractEditCreditParamsAccessScheduleAddScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCreditParamsAccessScheduleRemoveScheduleItem struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V2ContractEditCreditParamsAccessScheduleRemoveScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCreditParamsAccessScheduleUpdateScheduleItem struct {
	ID           param.Field[string]    `json:"id,required" format:"uuid"`
	Amount       param.Field[float64]   `json:"amount"`
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V2ContractEditCreditParamsAccessScheduleUpdateScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractGetEditHistoryParams struct {
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r V2ContractGetEditHistoryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
