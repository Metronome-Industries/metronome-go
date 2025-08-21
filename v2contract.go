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

// List all contracts for a customer in chronological order. New clients should use
// this endpoint rather than the v1 endpoint.
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
	Data shared.ContractV2         `json:"data,required"`
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

type V2ContractListResponse struct {
	Data []shared.ContractV2        `json:"data,required"`
	JSON v2ContractListResponseJSON `json:"-"`
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
	ID                                      string                                                                      `json:"id,required" format:"uuid"`
	AddCommits                              []V2ContractGetEditHistoryResponseDataAddCommit                             `json:"add_commits"`
	AddCredits                              []V2ContractGetEditHistoryResponseDataAddCredit                             `json:"add_credits"`
	AddDiscounts                            []shared.Discount                                                           `json:"add_discounts"`
	AddOverrides                            []V2ContractGetEditHistoryResponseDataAddOverride                           `json:"add_overrides"`
	AddPrepaidBalanceThresholdConfiguration V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfiguration `json:"add_prepaid_balance_threshold_configuration"`
	AddProServices                          []shared.ProService                                                         `json:"add_pro_services"`
	AddRecurringCommits                     []V2ContractGetEditHistoryResponseDataAddRecurringCommit                    `json:"add_recurring_commits"`
	AddRecurringCredits                     []V2ContractGetEditHistoryResponseDataAddRecurringCredit                    `json:"add_recurring_credits"`
	AddResellerRoyalties                    []V2ContractGetEditHistoryResponseDataAddResellerRoyalty                    `json:"add_reseller_royalties"`
	AddScheduledCharges                     []V2ContractGetEditHistoryResponseDataAddScheduledCharge                    `json:"add_scheduled_charges"`
	AddSpendThresholdConfiguration          V2ContractGetEditHistoryResponseDataAddSpendThresholdConfiguration          `json:"add_spend_threshold_configuration"`
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
	JSON                v2ContractGetEditHistoryResponseDataJSON                 `json:"-"`
}

// v2ContractGetEditHistoryResponseDataJSON contains the JSON metadata for the
// struct [V2ContractGetEditHistoryResponseData]
type v2ContractGetEditHistoryResponseDataJSON struct {
	ID                                         apijson.Field
	AddCommits                                 apijson.Field
	AddCredits                                 apijson.Field
	AddDiscounts                               apijson.Field
	AddOverrides                               apijson.Field
	AddPrepaidBalanceThresholdConfiguration    apijson.Field
	AddProServices                             apijson.Field
	AddRecurringCommits                        apijson.Field
	AddRecurringCredits                        apijson.Field
	AddResellerRoyalties                       apijson.Field
	AddScheduledCharges                        apijson.Field
	AddSpendThresholdConfiguration             apijson.Field
	AddSubscriptions                           apijson.Field
	AddUsageFilters                            apijson.Field
	ArchiveCommits                             apijson.Field
	ArchiveCredits                             apijson.Field
	ArchiveScheduledCharges                    apijson.Field
	RemoveOverrides                            apijson.Field
	Timestamp                                  apijson.Field
	UniquenessKey                              apijson.Field
	UpdateCommits                              apijson.Field
	UpdateContractEndDate                      apijson.Field
	UpdateContractName                         apijson.Field
	UpdateCredits                              apijson.Field
	UpdateDiscounts                            apijson.Field
	UpdatePrepaidBalanceThresholdConfiguration apijson.Field
	UpdateRecurringCommits                     apijson.Field
	UpdateRecurringCredits                     apijson.Field
	UpdateRefundInvoices                       apijson.Field
	UpdateScheduledCharges                     apijson.Field
	UpdateSpendThresholdConfiguration          apijson.Field
	UpdateSubscriptions                        apijson.Field
	raw                                        string
	ExtraFields                                map[string]apijson.Field
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
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfiguration `json:"hierarchy_configuration"`
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
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []V2ContractGetEditHistoryResponseDataAddCommitsSpecifier `json:"specifiers"`
	JSON       v2ContractGetEditHistoryResponseDataAddCommitJSON         `json:"-"`
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
	HierarchyConfiguration  apijson.Field
	InvoiceSchedule         apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	RateType                apijson.Field
	RolloverFraction        apijson.Field
	SalesforceOpportunityID apijson.Field
	Specifiers              apijson.Field
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

// Optional configuration for commit hierarchy access control
type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfiguration struct {
	ChildAccess V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationJSON        `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfiguration]
type v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccess struct {
	Type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                                         `json:"contract_ids"`
	JSON        v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessJSON `json:"-"`
	union       V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessUnion
}

// v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccess]
type v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
func (r V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccess) AsUnion() V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
// or
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs{}),
		},
	)
}

type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll]
type v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
type v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs []string                                                                                                                 `json:"contract_ids,required" format:"uuid"`
	Type        V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType `json:"type,required"`
	JSON        v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs]
type v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessType string

const (
	V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessTypeAll         V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessTypeNone        V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessTypeContractIDs V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessTypeAll, V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessTypeNone, V2ContractGetEditHistoryResponseDataAddCommitsHierarchyConfigurationChildAccessTypeContractIDs:
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

type V2ContractGetEditHistoryResponseDataAddCommitsSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                                    `json:"product_tags"`
	JSON        v2ContractGetEditHistoryResponseDataAddCommitsSpecifierJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCommitsSpecifierJSON contains the JSON
// metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCommitsSpecifier]
type v2ContractGetEditHistoryResponseDataAddCommitsSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCommitsSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCommitsSpecifierJSON) RawJSON() string {
	return r.raw
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
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfiguration `json:"hierarchy_configuration"`
	Name                   string                                                               `json:"name"`
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
	Specifiers []V2ContractGetEditHistoryResponseDataAddCreditsSpecifier `json:"specifiers"`
	JSON       v2ContractGetEditHistoryResponseDataAddCreditJSON         `json:"-"`
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
	HierarchyConfiguration  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	SalesforceOpportunityID apijson.Field
	Specifiers              apijson.Field
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

// Optional configuration for recurring credit hierarchy access control
type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfiguration struct {
	ChildAccess V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationJSON        `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfiguration]
type v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccess struct {
	Type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                                         `json:"contract_ids"`
	JSON        v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessJSON `json:"-"`
	union       V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessUnion
}

// v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccess]
type v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
func (r V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccess) AsUnion() V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
// or
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs{}),
		},
	)
}

type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll]
type v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
type v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs []string                                                                                                                 `json:"contract_ids,required" format:"uuid"`
	Type        V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType `json:"type,required"`
	JSON        v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs]
type v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessType string

const (
	V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessTypeAll         V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessTypeNone        V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessTypeContractIDs V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessTypeAll, V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessTypeNone, V2ContractGetEditHistoryResponseDataAddCreditsHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddCreditsSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                                    `json:"product_tags"`
	JSON        v2ContractGetEditHistoryResponseDataAddCreditsSpecifierJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddCreditsSpecifierJSON contains the JSON
// metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddCreditsSpecifier]
type v2ContractGetEditHistoryResponseDataAddCreditsSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddCreditsSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddCreditsSpecifierJSON) RawJSON() string {
	return r.raw
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

type V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfiguration struct {
	Commit V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                                                                                         `json:"is_enabled,required"`
	PaymentGateConfig V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfig `json:"payment_gate_config,required"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount float64 `json:"recharge_to_amount,required"`
	// Specify the threshold amount for the contract. Each time the contract's balance
	// lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount float64 `json:"threshold_amount,required"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID string                                                                          `json:"custom_credit_type_id" format:"uuid"`
	JSON               v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfiguration]
type v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationJSON struct {
	Commit             apijson.Field
	IsEnabled          apijson.Field
	PaymentGateConfig  apijson.Field
	RechargeToAmount   apijson.Field
	ThresholdAmount    apijson.Field
	CustomCreditTypeID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommit struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID string `json:"product_id,required"`
	// Which products the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Which tags the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags"`
	Description           string   `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string `json:"name"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitSpecifier `json:"specifiers"`
	JSON       v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitJSON        `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommit]
type v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitJSON struct {
	ProductID             apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Description           apijson.Field
	Name                  apijson.Field
	Specifiers            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                                                                       `json:"product_tags"`
	JSON        v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitSpecifierJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitSpecifierJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitSpecifier]
type v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationCommitSpecifierJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType `json:"tax_type"`
	JSON    v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigJSON    `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfig]
type v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigJSON struct {
	PaymentGateType        apijson.Field
	PrecalculatedTaxConfig apijson.Field
	StripeConfig           apijson.Field
	TaxType                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigJSON) RawJSON() string {
	return r.raw
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType string

const (
	V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone     V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "NONE"
	V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe   V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "STRIPE"
	V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone, V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe, V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName string                                                                                                                 `json:"tax_name"`
	JSON    v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig]
type v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON struct {
	TaxAmount   apijson.Field
	TaxName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON) RawJSON() string {
	return r.raw
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string                                                                                            `json:"invoice_metadata"`
	JSON            v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig]
type v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON struct {
	PaymentType     apijson.Field
	InvoiceMetadata apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON) RawJSON() string {
	return r.raw
}

// If left blank, will default to INVOICE
type V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType string

const (
	V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice       V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "INVOICE"
	V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice, V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType string

const (
	V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone          V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "NONE"
	V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe        V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "STRIPE"
	V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok         V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "ANROK"
	V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "PRECALCULATED"
)

func (r V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone, V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe, V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok, V2ContractGetEditHistoryResponseDataAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated:
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
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfiguration `json:"hierarchy_configuration"`
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
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	RecurrenceFrequency V2ContractGetEditHistoryResponseDataAddRecurringCommitsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []V2ContractGetEditHistoryResponseDataAddRecurringCommitsSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfig `json:"subscription_config"`
	JSON               v2ContractGetEditHistoryResponseDataAddRecurringCommitJSON                `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataAddRecurringCommit]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitJSON struct {
	ID                     apijson.Field
	AccessAmount           apijson.Field
	CommitDuration         apijson.Field
	Priority               apijson.Field
	Product                apijson.Field
	RateType               apijson.Field
	StartingAt             apijson.Field
	ApplicableProductIDs   apijson.Field
	ApplicableProductTags  apijson.Field
	Contract               apijson.Field
	Description            apijson.Field
	EndingBefore           apijson.Field
	HierarchyConfiguration apijson.Field
	InvoiceAmount          apijson.Field
	Name                   apijson.Field
	NetsuiteSalesOrderID   apijson.Field
	Proration              apijson.Field
	RecurrenceFrequency    apijson.Field
	RolloverFraction       apijson.Field
	Specifiers             apijson.Field
	SubscriptionConfig     apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
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
	UnitPrice    float64                                                                 `json:"unit_price,required"`
	Quantity     float64                                                                 `json:"quantity"`
	JSON         v2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmountJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmountJSON contains
// the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmount]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	UnitPrice    apijson.Field
	Quantity     apijson.Field
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

// Optional configuration for recurring credit hierarchy access control
type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfiguration struct {
	ChildAccess V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationJSON        `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfiguration]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccess struct {
	Type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                                                  `json:"contract_ids"`
	JSON        v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessJSON `json:"-"`
	union       V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessUnion
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccess]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccess) AsUnion() V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
// or
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs{}),
		},
	)
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs []string                                                                                                                          `json:"contract_ids,required" format:"uuid"`
	Type        V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType `json:"type,required"`
	JSON        v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessType string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessTypeAll         V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessTypeNone        V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessTypeContractIDs V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessTypeAll, V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessTypeNone, V2ContractGetEditHistoryResponseDataAddRecurringCommitsHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
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
// be created aligned with the recurring commit's starting_at rather than the usage
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

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                                             `json:"product_tags"`
	JSON        v2ContractGetEditHistoryResponseDataAddRecurringCommitsSpecifierJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsSpecifierJSON contains
// the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsSpecifier]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsSpecifierJSON) RawJSON() string {
	return r.raw
}

// Attach a subscription to the recurring commit/credit.
type V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfig struct {
	Allocation              V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigAllocation              `json:"allocation,required"`
	ApplySeatIncreaseConfig V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,required"`
	SubscriptionID          string                                                                                           `json:"subscription_id,required" format:"uuid"`
	JSON                    v2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigJSON                    `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfig]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigJSON struct {
	Allocation              apijson.Field
	ApplySeatIncreaseConfig apijson.Field
	SubscriptionID          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigAllocation string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigAllocationIndividual V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigAllocation = "INDIVIDUAL"
	V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigAllocationPooled     V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigAllocation = "POOLED"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigAllocation) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigAllocationIndividual, V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigAllocationPooled:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool                                                                                                 `json:"is_prorated,required"`
	JSON       v2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig]
type v2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON struct {
	IsProrated  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCommitsSubscriptionConfigApplySeatIncreaseConfigJSON) RawJSON() string {
	return r.raw
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
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfiguration `json:"hierarchy_configuration"`
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
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	RecurrenceFrequency V2ContractGetEditHistoryResponseDataAddRecurringCreditsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64 `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []V2ContractGetEditHistoryResponseDataAddRecurringCreditsSpecifier `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfig `json:"subscription_config"`
	JSON               v2ContractGetEditHistoryResponseDataAddRecurringCreditJSON                `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataAddRecurringCredit]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditJSON struct {
	ID                     apijson.Field
	AccessAmount           apijson.Field
	CommitDuration         apijson.Field
	Priority               apijson.Field
	Product                apijson.Field
	RateType               apijson.Field
	StartingAt             apijson.Field
	ApplicableProductIDs   apijson.Field
	ApplicableProductTags  apijson.Field
	Contract               apijson.Field
	Description            apijson.Field
	EndingBefore           apijson.Field
	HierarchyConfiguration apijson.Field
	Name                   apijson.Field
	NetsuiteSalesOrderID   apijson.Field
	Proration              apijson.Field
	RecurrenceFrequency    apijson.Field
	RolloverFraction       apijson.Field
	Specifiers             apijson.Field
	SubscriptionConfig     apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
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
	UnitPrice    float64                                                                 `json:"unit_price,required"`
	Quantity     float64                                                                 `json:"quantity"`
	JSON         v2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmountJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmountJSON contains
// the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmount]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	UnitPrice    apijson.Field
	Quantity     apijson.Field
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

// Optional configuration for recurring credit hierarchy access control
type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfiguration struct {
	ChildAccess V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationJSON        `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfiguration]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccess struct {
	Type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                                                  `json:"contract_ids"`
	JSON        v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessJSON `json:"-"`
	union       V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessUnion
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccess]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccess) AsUnion() V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
// or
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs{}),
		},
	)
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs []string                                                                                                                          `json:"contract_ids,required" format:"uuid"`
	Type        V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType `json:"type,required"`
	JSON        v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessType string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessTypeAll         V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessTypeNone        V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessTypeContractIDs V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessTypeAll, V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessTypeNone, V2ContractGetEditHistoryResponseDataAddRecurringCreditsHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
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
// be created aligned with the recurring commit's starting_at rather than the usage
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

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                                             `json:"product_tags"`
	JSON        v2ContractGetEditHistoryResponseDataAddRecurringCreditsSpecifierJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsSpecifierJSON contains
// the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsSpecifier]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsSpecifierJSON) RawJSON() string {
	return r.raw
}

// Attach a subscription to the recurring commit/credit.
type V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfig struct {
	Allocation              V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigAllocation              `json:"allocation,required"`
	ApplySeatIncreaseConfig V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig `json:"apply_seat_increase_config,required"`
	SubscriptionID          string                                                                                           `json:"subscription_id,required" format:"uuid"`
	JSON                    v2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigJSON                    `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfig]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigJSON struct {
	Allocation              apijson.Field
	ApplySeatIncreaseConfig apijson.Field
	SubscriptionID          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigAllocation string

const (
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigAllocationIndividual V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigAllocation = "INDIVIDUAL"
	V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigAllocationPooled     V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigAllocation = "POOLED"
)

func (r V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigAllocation) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigAllocationIndividual, V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigAllocationPooled:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated bool                                                                                                 `json:"is_prorated,required"`
	JSON       v2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig]
type v2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON struct {
	IsProrated  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddRecurringCreditsSubscriptionConfigApplySeatIncreaseConfigJSON) RawJSON() string {
	return r.raw
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

type V2ContractGetEditHistoryResponseDataAddSpendThresholdConfiguration struct {
	Commit V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationCommit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                                                                                `json:"is_enabled,required"`
	PaymentGateConfig V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfig `json:"payment_gate_config,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64                                                                `json:"threshold_amount,required"`
	JSON            v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationJSON contains
// the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddSpendThresholdConfiguration]
type v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationJSON struct {
	Commit            apijson.Field
	IsEnabled         apijson.Field
	PaymentGateConfig apijson.Field
	ThresholdAmount   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddSpendThresholdConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationCommit struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID   string `json:"product_id,required"`
	Description string `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string                                                                       `json:"name"`
	JSON v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationCommitJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationCommitJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationCommit]
type v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationCommitJSON struct {
	ProductID   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfig `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxType `json:"tax_type"`
	JSON    v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigJSON    `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfig]
type v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigJSON struct {
	PaymentGateType        apijson.Field
	PrecalculatedTaxConfig apijson.Field
	StripeConfig           apijson.Field
	TaxType                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigJSON) RawJSON() string {
	return r.raw
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType string

const (
	V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone     V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "NONE"
	V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe   V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "STRIPE"
	V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone, V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe, V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName string                                                                                                        `json:"tax_name"`
	JSON    v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig]
type v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON struct {
	TaxAmount   apijson.Field
	TaxName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON) RawJSON() string {
	return r.raw
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string                                                                                   `json:"invoice_metadata"`
	JSON            v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfig]
type v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON struct {
	PaymentType     apijson.Field
	InvoiceMetadata apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON) RawJSON() string {
	return r.raw
}

// If left blank, will default to INVOICE
type V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType string

const (
	V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice       V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "INVOICE"
	V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice, V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxType string

const (
	V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxTypeNone          V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxType = "NONE"
	V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe        V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxType = "STRIPE"
	V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok         V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxType = "ANROK"
	V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxType = "PRECALCULATED"
)

func (r V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxTypeNone, V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe, V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok, V2ContractGetEditHistoryResponseDataAddSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddSubscription struct {
	CollectionSchedule V2ContractGetEditHistoryResponseDataAddSubscriptionsCollectionSchedule `json:"collection_schedule,required"`
	Proration          V2ContractGetEditHistoryResponseDataAddSubscriptionsProration          `json:"proration,required"`
	// List of quantity schedule items for the subscription. Only includes the current
	// quantity and future quantity changes.
	QuantitySchedule []V2ContractGetEditHistoryResponseDataAddSubscriptionsQuantitySchedule `json:"quantity_schedule,required"`
	StartingAt       time.Time                                                              `json:"starting_at,required" format:"date-time"`
	SubscriptionRate V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRate   `json:"subscription_rate,required"`
	ID               string                                                                 `json:"id" format:"uuid"`
	CustomFields     map[string]string                                                      `json:"custom_fields"`
	Description      string                                                                 `json:"description"`
	EndingBefore     time.Time                                                              `json:"ending_before" format:"date-time"`
	FiatCreditTypeID string                                                                 `json:"fiat_credit_type_id" format:"uuid"`
	Name             string                                                                 `json:"name"`
	JSON             v2ContractGetEditHistoryResponseDataAddSubscriptionJSON                `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddSubscriptionJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataAddSubscription]
type v2ContractGetEditHistoryResponseDataAddSubscriptionJSON struct {
	CollectionSchedule apijson.Field
	Proration          apijson.Field
	QuantitySchedule   apijson.Field
	StartingAt         apijson.Field
	SubscriptionRate   apijson.Field
	ID                 apijson.Field
	CustomFields       apijson.Field
	Description        apijson.Field
	EndingBefore       apijson.Field
	FiatCreditTypeID   apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddSubscriptionJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionsCollectionSchedule string

const (
	V2ContractGetEditHistoryResponseDataAddSubscriptionsCollectionScheduleAdvance V2ContractGetEditHistoryResponseDataAddSubscriptionsCollectionSchedule = "ADVANCE"
	V2ContractGetEditHistoryResponseDataAddSubscriptionsCollectionScheduleArrears V2ContractGetEditHistoryResponseDataAddSubscriptionsCollectionSchedule = "ARREARS"
)

func (r V2ContractGetEditHistoryResponseDataAddSubscriptionsCollectionSchedule) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddSubscriptionsCollectionScheduleAdvance, V2ContractGetEditHistoryResponseDataAddSubscriptionsCollectionScheduleArrears:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionsProration struct {
	InvoiceBehavior V2ContractGetEditHistoryResponseDataAddSubscriptionsProrationInvoiceBehavior `json:"invoice_behavior,required"`
	IsProrated      bool                                                                         `json:"is_prorated,required"`
	JSON            v2ContractGetEditHistoryResponseDataAddSubscriptionsProrationJSON            `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddSubscriptionsProrationJSON contains the
// JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddSubscriptionsProration]
type v2ContractGetEditHistoryResponseDataAddSubscriptionsProrationJSON struct {
	InvoiceBehavior apijson.Field
	IsProrated      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionsProration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddSubscriptionsProrationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionsProrationInvoiceBehavior string

const (
	V2ContractGetEditHistoryResponseDataAddSubscriptionsProrationInvoiceBehaviorBillImmediately          V2ContractGetEditHistoryResponseDataAddSubscriptionsProrationInvoiceBehavior = "BILL_IMMEDIATELY"
	V2ContractGetEditHistoryResponseDataAddSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate V2ContractGetEditHistoryResponseDataAddSubscriptionsProrationInvoiceBehavior = "BILL_ON_NEXT_COLLECTION_DATE"
)

func (r V2ContractGetEditHistoryResponseDataAddSubscriptionsProrationInvoiceBehavior) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddSubscriptionsProrationInvoiceBehaviorBillImmediately, V2ContractGetEditHistoryResponseDataAddSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionsQuantitySchedule struct {
	Quantity     float64                                                                  `json:"quantity,required"`
	StartingAt   time.Time                                                                `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time                                                                `json:"ending_before" format:"date-time"`
	JSON         v2ContractGetEditHistoryResponseDataAddSubscriptionsQuantityScheduleJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddSubscriptionsQuantityScheduleJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddSubscriptionsQuantitySchedule]
type v2ContractGetEditHistoryResponseDataAddSubscriptionsQuantityScheduleJSON struct {
	Quantity     apijson.Field
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionsQuantitySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddSubscriptionsQuantityScheduleJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRate struct {
	BillingFrequency V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequency `json:"billing_frequency,required"`
	Product          V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateProduct          `json:"product,required"`
	JSON             v2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateJSON             `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRate]
type v2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateJSON struct {
	BillingFrequency apijson.Field
	Product          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequency string

const (
	V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequencyMonthly   V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequency = "MONTHLY"
	V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequencyQuarterly V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequency = "QUARTERLY"
	V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequencyAnnual    V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequency = "ANNUAL"
	V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequencyWeekly    V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequency = "WEEKLY"
)

func (r V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequency) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequencyMonthly, V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequencyQuarterly, V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequencyAnnual, V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateBillingFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateProduct struct {
	ID   string                                                                          `json:"id,required" format:"uuid"`
	Name string                                                                          `json:"name,required"`
	JSON v2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateProductJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateProductJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateProduct]
type v2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataAddSubscriptionsSubscriptionRateProductJSON) RawJSON() string {
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

type V2ContractGetEditHistoryResponseDataArchiveCommit struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	JSON v2ContractGetEditHistoryResponseDataArchiveCommitJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataArchiveCommitJSON contains the JSON metadata
// for the struct [V2ContractGetEditHistoryResponseDataArchiveCommit]
type v2ContractGetEditHistoryResponseDataArchiveCommitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataArchiveCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataArchiveCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataArchiveCredit struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	JSON v2ContractGetEditHistoryResponseDataArchiveCreditJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataArchiveCreditJSON contains the JSON metadata
// for the struct [V2ContractGetEditHistoryResponseDataArchiveCredit]
type v2ContractGetEditHistoryResponseDataArchiveCreditJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataArchiveCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataArchiveCreditJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataArchiveScheduledCharge struct {
	ID   string                                                         `json:"id,required" format:"uuid"`
	JSON v2ContractGetEditHistoryResponseDataArchiveScheduledChargeJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataArchiveScheduledChargeJSON contains the JSON
// metadata for the struct
// [V2ContractGetEditHistoryResponseDataArchiveScheduledCharge]
type v2ContractGetEditHistoryResponseDataArchiveScheduledChargeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataArchiveScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataArchiveScheduledChargeJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataRemoveOverride struct {
	ID   string                                                 `json:"id,required" format:"uuid"`
	JSON v2ContractGetEditHistoryResponseDataRemoveOverrideJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataRemoveOverrideJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataRemoveOverride]
type v2ContractGetEditHistoryResponseDataRemoveOverrideJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataRemoveOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataRemoveOverrideJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCommit struct {
	ID             string                                                          `json:"id,required" format:"uuid"`
	AccessSchedule V2ContractGetEditHistoryResponseDataUpdateCommitsAccessSchedule `json:"access_schedule"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,nullable" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags []string `json:"applicable_product_tags,nullable"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfiguration `json:"hierarchy_configuration"`
	InvoiceSchedule        V2ContractGetEditHistoryResponseDataUpdateCommitsInvoiceSchedule        `json:"invoice_schedule"`
	Name                   string                                                                  `json:"name"`
	NetsuiteSalesOrderID   string                                                                  `json:"netsuite_sales_order_id,nullable"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority         float64 `json:"priority,nullable"`
	ProductID        string  `json:"product_id" format:"uuid"`
	RolloverFraction float64 `json:"rollover_fraction,nullable"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []V2ContractGetEditHistoryResponseDataUpdateCommitsSpecifier `json:"specifiers,nullable"`
	JSON       v2ContractGetEditHistoryResponseDataUpdateCommitJSON         `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitJSON contains the JSON metadata
// for the struct [V2ContractGetEditHistoryResponseDataUpdateCommit]
type v2ContractGetEditHistoryResponseDataUpdateCommitJSON struct {
	ID                     apijson.Field
	AccessSchedule         apijson.Field
	ApplicableProductIDs   apijson.Field
	ApplicableProductTags  apijson.Field
	HierarchyConfiguration apijson.Field
	InvoiceSchedule        apijson.Field
	Name                   apijson.Field
	NetsuiteSalesOrderID   apijson.Field
	Priority               apijson.Field
	ProductID              apijson.Field
	RolloverFraction       apijson.Field
	Specifiers             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
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

// Optional configuration for commit hierarchy access control
type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfiguration struct {
	ChildAccess V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationJSON        `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfiguration]
type v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccess struct {
	Type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                                            `json:"contract_ids"`
	JSON        v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessJSON `json:"-"`
	union       V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessUnion
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccess]
type v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
func (r V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccess) AsUnion() V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
// or
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs{}),
		},
	)
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll]
type v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
type v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs []string                                                                                                                    `json:"contract_ids,required" format:"uuid"`
	Type        V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType `json:"type,required"`
	JSON        v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs]
type v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessType string

const (
	V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessTypeAll         V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessTypeNone        V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessTypeContractIDs V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessTypeAll, V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessTypeNone, V2ContractGetEditHistoryResponseDataUpdateCommitsHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
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

type V2ContractGetEditHistoryResponseDataUpdateCommitsSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                                       `json:"product_tags"`
	JSON        v2ContractGetEditHistoryResponseDataUpdateCommitsSpecifierJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCommitsSpecifierJSON contains the JSON
// metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCommitsSpecifier]
type v2ContractGetEditHistoryResponseDataUpdateCommitsSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCommitsSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCommitsSpecifierJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCredit struct {
	ID             string                                                          `json:"id,required" format:"uuid"`
	AccessSchedule V2ContractGetEditHistoryResponseDataUpdateCreditsAccessSchedule `json:"access_schedule"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfiguration `json:"hierarchy_configuration"`
	Name                   string                                                                  `json:"name"`
	NetsuiteSalesOrderID   string                                                                  `json:"netsuite_sales_order_id,nullable"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority         float64                                              `json:"priority,nullable"`
	RolloverFraction float64                                              `json:"rollover_fraction,nullable"`
	JSON             v2ContractGetEditHistoryResponseDataUpdateCreditJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCreditJSON contains the JSON metadata
// for the struct [V2ContractGetEditHistoryResponseDataUpdateCredit]
type v2ContractGetEditHistoryResponseDataUpdateCreditJSON struct {
	ID                     apijson.Field
	AccessSchedule         apijson.Field
	HierarchyConfiguration apijson.Field
	Name                   apijson.Field
	NetsuiteSalesOrderID   apijson.Field
	Priority               apijson.Field
	RolloverFraction       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
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

// Optional configuration for credit hierarchy access control
type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfiguration struct {
	ChildAccess V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationJSON        `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfiguration]
type v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccess struct {
	Type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                                            `json:"contract_ids"`
	JSON        v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessJSON `json:"-"`
	union       V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessUnion
}

// v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccess]
type v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
func (r V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccess) AsUnion() V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
// or
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs].
type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs{}),
		},
	)
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll]
type v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType `json:"type,required"`
	JSON v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone]
type v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs []string                                                                                                                    `json:"contract_ids,required" format:"uuid"`
	Type        V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType `json:"type,required"`
	JSON        v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs]
type v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsJSON) RawJSON() string {
	return r.raw
}

func (r V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccess() {
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessType string

const (
	V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessTypeAll         V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessTypeNone        V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessTypeContractIDs V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessTypeAll, V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessTypeNone, V2ContractGetEditHistoryResponseDataUpdateCreditsHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
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
	// Defaults to USD (cents) if not passed.
	CreditTypeID string `json:"credit_type_id" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice bool `json:"do_not_invoice"`
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
	DoNotInvoice      apijson.Field
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

type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfiguration struct {
	Commit V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommit `json:"commit"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID string `json:"custom_credit_type_id,nullable" format:"uuid"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                                                                                            `json:"is_enabled"`
	PaymentGateConfig V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfig `json:"payment_gate_config"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount float64 `json:"recharge_to_amount"`
	// Specify the threshold amount for the contract. Each time the contract's balance
	// lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount float64                                                                            `json:"threshold_amount"`
	JSON            v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfiguration]
type v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationJSON struct {
	Commit             apijson.Field
	CustomCreditTypeID apijson.Field
	IsEnabled          apijson.Field
	PaymentGateConfig  apijson.Field
	RechargeToAmount   apijson.Field
	ThresholdAmount    apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommit struct {
	// Which products the threshold commit applies to. If both applicable_product_ids
	// and applicable_product_tags are not provided, the commit applies to all
	// products.
	ApplicableProductIDs []string `json:"applicable_product_ids,nullable" format:"uuid"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags,nullable"`
	Description           string   `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string `json:"name"`
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID string `json:"product_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers []V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitSpecifier `json:"specifiers,nullable"`
	JSON       v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitJSON        `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommit]
type v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitJSON struct {
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Description           apijson.Field
	Name                  apijson.Field
	ProductID             apijson.Field
	Specifiers            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                                                                          `json:"product_tags"`
	JSON        v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitSpecifierJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitSpecifierJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitSpecifier]
type v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationCommitSpecifierJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType `json:"tax_type"`
	JSON    v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigJSON    `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfig]
type v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigJSON struct {
	PaymentGateType        apijson.Field
	PrecalculatedTaxConfig apijson.Field
	StripeConfig           apijson.Field
	TaxType                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigJSON) RawJSON() string {
	return r.raw
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType string

const (
	V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone     V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "NONE"
	V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe   V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "STRIPE"
	V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone, V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe, V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName string                                                                                                                    `json:"tax_name"`
	JSON    v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig]
type v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON struct {
	TaxAmount   apijson.Field
	TaxName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON) RawJSON() string {
	return r.raw
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string                                                                                               `json:"invoice_metadata"`
	JSON            v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig]
type v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON struct {
	PaymentType     apijson.Field
	InvoiceMetadata apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigJSON) RawJSON() string {
	return r.raw
}

// If left blank, will default to INVOICE
type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType string

const (
	V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice       V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "INVOICE"
	V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice, V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType string

const (
	V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone          V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "NONE"
	V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe        V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "STRIPE"
	V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok         V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "ANROK"
	V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "PRECALCULATED"
)

func (r V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone, V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe, V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok, V2ContractGetEditHistoryResponseDataUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCommit struct {
	ID            string                                                                  `json:"id,required" format:"uuid"`
	AccessAmount  V2ContractGetEditHistoryResponseDataUpdateRecurringCommitsAccessAmount  `json:"access_amount"`
	EndingBefore  time.Time                                                               `json:"ending_before" format:"date-time"`
	InvoiceAmount V2ContractGetEditHistoryResponseDataUpdateRecurringCommitsInvoiceAmount `json:"invoice_amount"`
	JSON          v2ContractGetEditHistoryResponseDataUpdateRecurringCommitJSON           `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateRecurringCommitJSON contains the JSON
// metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateRecurringCommit]
type v2ContractGetEditHistoryResponseDataUpdateRecurringCommitJSON struct {
	ID            apijson.Field
	AccessAmount  apijson.Field
	EndingBefore  apijson.Field
	InvoiceAmount apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateRecurringCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCommitsAccessAmount struct {
	Quantity  float64                                                                    `json:"quantity"`
	UnitPrice float64                                                                    `json:"unit_price"`
	JSON      v2ContractGetEditHistoryResponseDataUpdateRecurringCommitsAccessAmountJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateRecurringCommitsAccessAmountJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateRecurringCommitsAccessAmount]
type v2ContractGetEditHistoryResponseDataUpdateRecurringCommitsAccessAmountJSON struct {
	Quantity    apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCommitsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateRecurringCommitsAccessAmountJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCommitsInvoiceAmount struct {
	Quantity  float64                                                                     `json:"quantity"`
	UnitPrice float64                                                                     `json:"unit_price"`
	JSON      v2ContractGetEditHistoryResponseDataUpdateRecurringCommitsInvoiceAmountJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateRecurringCommitsInvoiceAmountJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateRecurringCommitsInvoiceAmount]
type v2ContractGetEditHistoryResponseDataUpdateRecurringCommitsInvoiceAmountJSON struct {
	Quantity    apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCommitsInvoiceAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateRecurringCommitsInvoiceAmountJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCredit struct {
	ID           string                                                                 `json:"id,required" format:"uuid"`
	AccessAmount V2ContractGetEditHistoryResponseDataUpdateRecurringCreditsAccessAmount `json:"access_amount"`
	EndingBefore time.Time                                                              `json:"ending_before" format:"date-time"`
	JSON         v2ContractGetEditHistoryResponseDataUpdateRecurringCreditJSON          `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateRecurringCreditJSON contains the JSON
// metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateRecurringCredit]
type v2ContractGetEditHistoryResponseDataUpdateRecurringCreditJSON struct {
	ID           apijson.Field
	AccessAmount apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateRecurringCreditJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateRecurringCreditsAccessAmount struct {
	Quantity  float64                                                                    `json:"quantity"`
	UnitPrice float64                                                                    `json:"unit_price"`
	JSON      v2ContractGetEditHistoryResponseDataUpdateRecurringCreditsAccessAmountJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateRecurringCreditsAccessAmountJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateRecurringCreditsAccessAmount]
type v2ContractGetEditHistoryResponseDataUpdateRecurringCreditsAccessAmountJSON struct {
	Quantity    apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateRecurringCreditsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateRecurringCreditsAccessAmountJSON) RawJSON() string {
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

type V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfiguration struct {
	Commit V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationCommit `json:"commit"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         bool                                                                                   `json:"is_enabled"`
	PaymentGateConfig V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfig `json:"payment_gate_config"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64                                                                   `json:"threshold_amount"`
	JSON            v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfiguration]
type v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationJSON struct {
	Commit            apijson.Field
	IsEnabled         apijson.Field
	PaymentGateConfig apijson.Field
	ThresholdAmount   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationCommit struct {
	Description string `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string `json:"name"`
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID string                                                                          `json:"product_id"`
	JSON      v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationCommitJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationCommitJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationCommit]
type v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationCommitJSON struct {
	Description apijson.Field
	Name        apijson.Field
	ProductID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationCommitJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfig `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxType `json:"tax_type"`
	JSON    v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigJSON    `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfig]
type v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigJSON struct {
	PaymentGateType        apijson.Field
	PrecalculatedTaxConfig apijson.Field
	StripeConfig           apijson.Field
	TaxType                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigJSON) RawJSON() string {
	return r.raw
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType string

const (
	V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone     V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "NONE"
	V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe   V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "STRIPE"
	V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone, V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe, V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount float64 `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName string                                                                                                           `json:"tax_name"`
	JSON    v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig]
type v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON struct {
	TaxAmount   apijson.Field
	TaxName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigJSON) RawJSON() string {
	return r.raw
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata map[string]string                                                                                      `json:"invoice_metadata"`
	JSON            v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfig]
type v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON struct {
	PaymentType     apijson.Field
	InvoiceMetadata apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigJSON) RawJSON() string {
	return r.raw
}

// If left blank, will default to INVOICE
type V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType string

const (
	V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice       V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "INVOICE"
	V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice, V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxType string

const (
	V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeNone          V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxType = "NONE"
	V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe        V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxType = "STRIPE"
	V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok         V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxType = "ANROK"
	V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxType = "PRECALCULATED"
)

func (r V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeNone, V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe, V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok, V2ContractGetEditHistoryResponseDataUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated:
		return true
	}
	return false
}

type V2ContractGetEditHistoryResponseDataUpdateSubscription struct {
	ID              string                                                                  `json:"id,required" format:"uuid"`
	EndingBefore    time.Time                                                               `json:"ending_before" format:"date-time"`
	QuantityUpdates []V2ContractGetEditHistoryResponseDataUpdateSubscriptionsQuantityUpdate `json:"quantity_updates"`
	JSON            v2ContractGetEditHistoryResponseDataUpdateSubscriptionJSON              `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateSubscriptionJSON contains the JSON
// metadata for the struct [V2ContractGetEditHistoryResponseDataUpdateSubscription]
type v2ContractGetEditHistoryResponseDataUpdateSubscriptionJSON struct {
	ID              apijson.Field
	EndingBefore    apijson.Field
	QuantityUpdates apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateSubscriptionJSON) RawJSON() string {
	return r.raw
}

type V2ContractGetEditHistoryResponseDataUpdateSubscriptionsQuantityUpdate struct {
	StartingAt    time.Time                                                                 `json:"starting_at,required" format:"date-time"`
	Quantity      float64                                                                   `json:"quantity"`
	QuantityDelta float64                                                                   `json:"quantity_delta"`
	JSON          v2ContractGetEditHistoryResponseDataUpdateSubscriptionsQuantityUpdateJSON `json:"-"`
}

// v2ContractGetEditHistoryResponseDataUpdateSubscriptionsQuantityUpdateJSON
// contains the JSON metadata for the struct
// [V2ContractGetEditHistoryResponseDataUpdateSubscriptionsQuantityUpdate]
type v2ContractGetEditHistoryResponseDataUpdateSubscriptionsQuantityUpdateJSON struct {
	StartingAt    apijson.Field
	Quantity      apijson.Field
	QuantityDelta apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2ContractGetEditHistoryResponseDataUpdateSubscriptionsQuantityUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ContractGetEditHistoryResponseDataUpdateSubscriptionsQuantityUpdateJSON) RawJSON() string {
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
	CustomerID                              param.Field[string]                                                      `json:"customer_id,required" format:"uuid"`
	AddCommits                              param.Field[[]V2ContractEditParamsAddCommit]                             `json:"add_commits"`
	AddCredits                              param.Field[[]V2ContractEditParamsAddCredit]                             `json:"add_credits"`
	AddDiscounts                            param.Field[[]V2ContractEditParamsAddDiscount]                           `json:"add_discounts"`
	AddOverrides                            param.Field[[]V2ContractEditParamsAddOverride]                           `json:"add_overrides"`
	AddPrepaidBalanceThresholdConfiguration param.Field[V2ContractEditParamsAddPrepaidBalanceThresholdConfiguration] `json:"add_prepaid_balance_threshold_configuration"`
	// This field's availability is dependent on your client's configuration.
	AddProfessionalServices        param.Field[[]V2ContractEditParamsAddProfessionalService]       `json:"add_professional_services"`
	AddRecurringCommits            param.Field[[]V2ContractEditParamsAddRecurringCommit]           `json:"add_recurring_commits"`
	AddRecurringCredits            param.Field[[]V2ContractEditParamsAddRecurringCredit]           `json:"add_recurring_credits"`
	AddResellerRoyalties           param.Field[[]V2ContractEditParamsAddResellerRoyalty]           `json:"add_reseller_royalties"`
	AddScheduledCharges            param.Field[[]V2ContractEditParamsAddScheduledCharge]           `json:"add_scheduled_charges"`
	AddSpendThresholdConfiguration param.Field[V2ContractEditParamsAddSpendThresholdConfiguration] `json:"add_spend_threshold_configuration"`
	// Optional list of
	// [subscriptions](https://docs.metronome.com/manage-product-access/create-subscription/)
	// to add to the contract.
	AddSubscriptions param.Field[[]V2ContractEditParamsAddSubscription] `json:"add_subscriptions"`
	// If true, allows setting the contract end date earlier than the end_timestamp of
	// existing finalized invoices. Finalized invoices will be unchanged; if you want
	// to incorporate the new end date, you can void and regenerate finalized usage
	// invoices. Defaults to true.
	AllowContractEndingBeforeFinalizedInvoice param.Field[bool] `json:"allow_contract_ending_before_finalized_invoice"`
	// IDs of commits to archive
	ArchiveCommits param.Field[[]V2ContractEditParamsArchiveCommit] `json:"archive_commits"`
	// IDs of credits to archive
	ArchiveCredits param.Field[[]V2ContractEditParamsArchiveCredit] `json:"archive_credits"`
	// IDs of scheduled charges to archive
	ArchiveScheduledCharges param.Field[[]V2ContractEditParamsArchiveScheduledCharge] `json:"archive_scheduled_charges"`
	// IDs of overrides to remove
	RemoveOverrides param.Field[[]V2ContractEditParamsRemoveOverride] `json:"remove_overrides"`
	// Optional uniqueness key to prevent duplicate contract edits.
	UniquenessKey param.Field[string]                             `json:"uniqueness_key"`
	UpdateCommits param.Field[[]V2ContractEditParamsUpdateCommit] `json:"update_commits"`
	// RFC 3339 timestamp indicating when the contract will end (exclusive).
	UpdateContractEndDate param.Field[time.Time] `json:"update_contract_end_date" format:"date-time"`
	// Value to update the contract name to. If not provided, the contract name will
	// remain unchanged.
	UpdateContractName                         param.Field[string]                                                         `json:"update_contract_name"`
	UpdateCredits                              param.Field[[]V2ContractEditParamsUpdateCredit]                             `json:"update_credits"`
	UpdatePrepaidBalanceThresholdConfiguration param.Field[V2ContractEditParamsUpdatePrepaidBalanceThresholdConfiguration] `json:"update_prepaid_balance_threshold_configuration"`
	// Edits to these recurring commits will only affect commits whose access schedules
	// has not started. Expired commits, and commits with an active access schedule
	// will remain unchanged.
	UpdateRecurringCommits param.Field[[]V2ContractEditParamsUpdateRecurringCommit] `json:"update_recurring_commits"`
	// Edits to these recurring credits will only affect credits whose access schedules
	// has not started. Expired credits, and credits with an active access schedule
	// will remain unchanged.
	UpdateRecurringCredits            param.Field[[]V2ContractEditParamsUpdateRecurringCredit]           `json:"update_recurring_credits"`
	UpdateScheduledCharges            param.Field[[]V2ContractEditParamsUpdateScheduledCharge]           `json:"update_scheduled_charges"`
	UpdateSpendThresholdConfiguration param.Field[V2ContractEditParamsUpdateSpendThresholdConfiguration] `json:"update_spend_threshold_configuration"`
	// Optional list of subscriptions to update.
	UpdateSubscriptions param.Field[[]V2ContractEditParamsUpdateSubscription] `json:"update_subscriptions"`
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
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags param.Field[[]string]          `json:"applicable_product_tags"`
	CustomFields          param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration param.Field[V2ContractEditParamsAddCommitsHierarchyConfiguration] `json:"hierarchy_configuration"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this
	// time and only one schedule item is allowed; the total must match access_schedule
	// amount. Optional for "PREPAID" commits: if not provided, this will be a
	// "complimentary" commit with no invoice.
	InvoiceSchedule param.Field[V2ContractEditParamsAddCommitsInvoiceSchedule] `json:"invoice_schedule"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// optionally payment gate this commit
	PaymentGateConfig param.Field[V2ContractEditParamsAddCommitsPaymentGateConfig] `json:"payment_gate_config"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Field[float64]                                `json:"priority"`
	RateType param.Field[V2ContractEditParamsAddCommitsRateType] `json:"rate_type"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers param.Field[[]V2ContractEditParamsAddCommitsSpecifier] `json:"specifiers"`
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

// Optional configuration for commit hierarchy access control
type V2ContractEditParamsAddCommitsHierarchyConfiguration struct {
	ChildAccess param.Field[V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessUnion] `json:"child_access,required"`
}

func (r V2ContractEditParamsAddCommitsHierarchyConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccess struct {
	Type        param.Field[V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessType] `json:"type,required"`
	ContractIDs param.Field[interface{}]                                                         `json:"contract_ids"`
}

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccess) implementsV2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessUnion() {
}

// Satisfied by
// [V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs],
// [V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccess].
type V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessUnion()
}

type V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type param.Field[V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType] `json:"type,required"`
}

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type param.Field[V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType] `json:"type,required"`
}

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs param.Field[[]string]                                                                                                 `json:"contract_ids,required" format:"uuid"`
	Type        param.Field[V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType] `json:"type,required"`
}

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessType string

const (
	V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessTypeAll         V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessTypeNone        V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessTypeContractIDs V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessTypeAll, V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessTypeNone, V2ContractEditParamsAddCommitsHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match access_schedule
// amount. Optional for "PREPAID" commits: if not provided, this will be a
// "complimentary" commit with no invoice.
type V2ContractEditParamsAddCommitsInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Field[bool] `json:"do_not_invoice"`
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

// optionally payment gate this commit
type V2ContractEditParamsAddCommitsPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType param.Field[V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateType] `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig param.Field[V2ContractEditParamsAddCommitsPaymentGateConfigPrecalculatedTaxConfig] `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig param.Field[V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfig] `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType param.Field[V2ContractEditParamsAddCommitsPaymentGateConfigTaxType] `json:"tax_type"`
}

func (r V2ContractEditParamsAddCommitsPaymentGateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateType string

const (
	V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateTypeNone     V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateType = "NONE"
	V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateTypeStripe   V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateType = "STRIPE"
	V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateTypeExternal V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateTypeNone, V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateTypeStripe, V2ContractEditParamsAddCommitsPaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractEditParamsAddCommitsPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount param.Field[float64] `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Field[string] `json:"tax_name"`
}

func (r V2ContractEditParamsAddCommitsPaymentGateConfigPrecalculatedTaxConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType param.Field[V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfigPaymentType] `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata param.Field[map[string]string] `json:"invoice_metadata"`
	// If true, the payment will be made assuming the customer is present (i.e. on
	// session).
	//
	// If false, the payment will be made assuming the customer is not present (i.e.
	// off session). For cardholders from a country with an e-mandate requirement (e.g.
	// India), the payment may be declined.
	//
	// If left blank, will default to false.
	OnSessionPayment param.Field[bool] `json:"on_session_payment"`
}

func (r V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If left blank, will default to INVOICE
type V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfigPaymentType string

const (
	V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfigPaymentTypeInvoice       V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfigPaymentType = "INVOICE"
	V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfigPaymentTypePaymentIntent V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfigPaymentTypeInvoice, V2ContractEditParamsAddCommitsPaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type V2ContractEditParamsAddCommitsPaymentGateConfigTaxType string

const (
	V2ContractEditParamsAddCommitsPaymentGateConfigTaxTypeNone          V2ContractEditParamsAddCommitsPaymentGateConfigTaxType = "NONE"
	V2ContractEditParamsAddCommitsPaymentGateConfigTaxTypeStripe        V2ContractEditParamsAddCommitsPaymentGateConfigTaxType = "STRIPE"
	V2ContractEditParamsAddCommitsPaymentGateConfigTaxTypeAnrok         V2ContractEditParamsAddCommitsPaymentGateConfigTaxType = "ANROK"
	V2ContractEditParamsAddCommitsPaymentGateConfigTaxTypePrecalculated V2ContractEditParamsAddCommitsPaymentGateConfigTaxType = "PRECALCULATED"
)

func (r V2ContractEditParamsAddCommitsPaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCommitsPaymentGateConfigTaxTypeNone, V2ContractEditParamsAddCommitsPaymentGateConfigTaxTypeStripe, V2ContractEditParamsAddCommitsPaymentGateConfigTaxTypeAnrok, V2ContractEditParamsAddCommitsPaymentGateConfigTaxTypePrecalculated:
		return true
	}
	return false
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

type V2ContractEditParamsAddCommitsSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V2ContractEditParamsAddCommitsSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration param.Field[V2ContractEditParamsAddCreditsHierarchyConfiguration] `json:"hierarchy_configuration"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Field[float64]                                `json:"priority"`
	RateType param.Field[V2ContractEditParamsAddCreditsRateType] `json:"rate_type"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers param.Field[[]V2ContractEditParamsAddCreditsSpecifier] `json:"specifiers"`
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

// Optional configuration for credit hierarchy access control
type V2ContractEditParamsAddCreditsHierarchyConfiguration struct {
	ChildAccess param.Field[V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessUnion] `json:"child_access,required"`
}

func (r V2ContractEditParamsAddCreditsHierarchyConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccess struct {
	Type        param.Field[V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessType] `json:"type,required"`
	ContractIDs param.Field[interface{}]                                                         `json:"contract_ids"`
}

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccess) implementsV2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessUnion() {
}

// Satisfied by
// [V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs],
// [V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccess].
type V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessUnion()
}

type V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type param.Field[V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType] `json:"type,required"`
}

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type param.Field[V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType] `json:"type,required"`
}

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs param.Field[[]string]                                                                                                 `json:"contract_ids,required" format:"uuid"`
	Type        param.Field[V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType] `json:"type,required"`
}

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessType string

const (
	V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessTypeAll         V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessTypeNone        V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessTypeContractIDs V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessTypeAll, V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessTypeNone, V2ContractEditParamsAddCreditsHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
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

type V2ContractEditParamsAddCreditsSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V2ContractEditParamsAddCreditsSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Field[bool] `json:"do_not_invoice"`
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
	BillingFrequency param.Field[V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequency] `json:"billing_frequency"`
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

type V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequency string

const (
	V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequencyMonthly   V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequency = "MONTHLY"
	V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequencyQuarterly V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequency = "QUARTERLY"
	V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequencyAnnual    V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequency = "ANNUAL"
	V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequencyWeekly    V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequency = "WEEKLY"
)

func (r V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequency) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequencyMonthly, V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequencyQuarterly, V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequencyAnnual, V2ContractEditParamsAddOverridesOverrideSpecifiersBillingFrequencyWeekly:
		return true
	}
	return false
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

type V2ContractEditParamsAddPrepaidBalanceThresholdConfiguration struct {
	Commit param.Field[V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationCommit] `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         param.Field[bool]                                                                         `json:"is_enabled,required"`
	PaymentGateConfig param.Field[V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfig] `json:"payment_gate_config,required"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount param.Field[float64] `json:"recharge_to_amount,required"`
	// Specify the threshold amount for the contract. Each time the contract's balance
	// lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount param.Field[float64] `json:"threshold_amount,required"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID param.Field[string] `json:"custom_credit_type_id" format:"uuid"`
}

func (r V2ContractEditParamsAddPrepaidBalanceThresholdConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationCommit struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID param.Field[string] `json:"product_id,required"`
	// Which products the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the threshold commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	Description           param.Field[string]   `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name param.Field[string] `json:"name"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers param.Field[[]V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationCommitSpecifier] `json:"specifiers"`
}

func (r V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationCommitSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationCommitSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType param.Field[V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType] `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig param.Field[V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig] `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig param.Field[V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig] `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType param.Field[V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType] `json:"tax_type"`
}

func (r V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType string

const (
	V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone     V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "NONE"
	V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe   V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "STRIPE"
	V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone, V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe, V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount param.Field[float64] `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Field[string] `json:"tax_name"`
}

func (r V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType param.Field[V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType] `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata param.Field[map[string]string] `json:"invoice_metadata"`
}

func (r V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If left blank, will default to INVOICE
type V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType string

const (
	V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice       V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "INVOICE"
	V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice, V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType string

const (
	V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone          V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "NONE"
	V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe        V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "STRIPE"
	V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok         V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "ANROK"
	V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "PRECALCULATED"
)

func (r V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone, V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe, V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok, V2ContractEditParamsAddPrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated:
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
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
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
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration param.Field[V2ContractEditParamsAddRecurringCommitsHierarchyConfiguration] `json:"hierarchy_configuration"`
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
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	RecurrenceFrequency param.Field[V2ContractEditParamsAddRecurringCommitsRecurrenceFrequency] `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers param.Field[[]V2ContractEditParamsAddRecurringCommitsSpecifier] `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig param.Field[V2ContractEditParamsAddRecurringCommitsSubscriptionConfig] `json:"subscription_config"`
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
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
	// This field is required unless a subscription is attached via
	// `subscription_config`.
	Quantity param.Field[float64] `json:"quantity"`
}

func (r V2ContractEditParamsAddRecurringCommitsAccessAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines the length of the access schedule for each created commit/credit. The
// value represents the number of units. Unit defaults to "PERIODS", where the
// length of a period is determined by the recurrence_frequency.
type V2ContractEditParamsAddRecurringCommitsCommitDuration struct {
	Value param.Field[float64]                                                   `json:"value,required"`
	Unit  param.Field[V2ContractEditParamsAddRecurringCommitsCommitDurationUnit] `json:"unit"`
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

// Optional configuration for recurring credit hierarchy access control
type V2ContractEditParamsAddRecurringCommitsHierarchyConfiguration struct {
	ChildAccess param.Field[V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessUnion] `json:"child_access,required"`
}

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccess struct {
	Type        param.Field[V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessType] `json:"type,required"`
	ContractIDs param.Field[interface{}]                                                                  `json:"contract_ids"`
}

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccess) implementsV2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessUnion() {
}

// Satisfied by
// [V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs],
// [V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccess].
type V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessUnion()
}

type V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type param.Field[V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType] `json:"type,required"`
}

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type param.Field[V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType] `json:"type,required"`
}

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs param.Field[[]string]                                                                                                          `json:"contract_ids,required" format:"uuid"`
	Type        param.Field[V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType] `json:"type,required"`
}

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessType string

const (
	V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessTypeAll         V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessTypeNone        V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessTypeContractIDs V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessTypeAll, V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessTypeNone, V2ContractEditParamsAddRecurringCommitsHierarchyConfigurationChildAccessTypeContractIDs:
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
// be created aligned with the recurring commit's starting_at rather than the usage
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

type V2ContractEditParamsAddRecurringCommitsSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V2ContractEditParamsAddRecurringCommitsSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attach a subscription to the recurring commit/credit.
type V2ContractEditParamsAddRecurringCommitsSubscriptionConfig struct {
	ApplySeatIncreaseConfig param.Field[V2ContractEditParamsAddRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig] `json:"apply_seat_increase_config,required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID param.Field[string] `json:"subscription_id,required"`
	// If set to POOLED, allocation added per seat is pooled across the account.
	Allocation param.Field[V2ContractEditParamsAddRecurringCommitsSubscriptionConfigAllocation] `json:"allocation"`
}

func (r V2ContractEditParamsAddRecurringCommitsSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated param.Field[bool] `json:"is_prorated,required"`
}

func (r V2ContractEditParamsAddRecurringCommitsSubscriptionConfigApplySeatIncreaseConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set to POOLED, allocation added per seat is pooled across the account.
type V2ContractEditParamsAddRecurringCommitsSubscriptionConfigAllocation string

const (
	V2ContractEditParamsAddRecurringCommitsSubscriptionConfigAllocationPooled V2ContractEditParamsAddRecurringCommitsSubscriptionConfigAllocation = "POOLED"
)

func (r V2ContractEditParamsAddRecurringCommitsSubscriptionConfigAllocation) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCommitsSubscriptionConfigAllocationPooled:
		return true
	}
	return false
}

type V2ContractEditParamsAddRecurringCredit struct {
	// The amount of commit to grant.
	AccessAmount param.Field[V2ContractEditParamsAddRecurringCreditsAccessAmount] `json:"access_amount,required"`
	// Defines the length of the access schedule for each created commit/credit. The
	// value represents the number of units. Unit defaults to "PERIODS", where the
	// length of a period is determined by the recurrence_frequency.
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
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration param.Field[V2ContractEditParamsAddRecurringCreditsHierarchyConfiguration] `json:"hierarchy_configuration"`
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
	// be created aligned with the recurring commit's starting_at rather than the usage
	// invoice dates.
	RecurrenceFrequency param.Field[V2ContractEditParamsAddRecurringCreditsRecurrenceFrequency] `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers param.Field[[]V2ContractEditParamsAddRecurringCreditsSpecifier] `json:"specifiers"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig param.Field[V2ContractEditParamsAddRecurringCreditsSubscriptionConfig] `json:"subscription_config"`
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
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
	// This field is required unless a subscription is attached via
	// `subscription_config`.
	Quantity param.Field[float64] `json:"quantity"`
}

func (r V2ContractEditParamsAddRecurringCreditsAccessAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines the length of the access schedule for each created commit/credit. The
// value represents the number of units. Unit defaults to "PERIODS", where the
// length of a period is determined by the recurrence_frequency.
type V2ContractEditParamsAddRecurringCreditsCommitDuration struct {
	Value param.Field[float64]                                                   `json:"value,required"`
	Unit  param.Field[V2ContractEditParamsAddRecurringCreditsCommitDurationUnit] `json:"unit"`
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

// Optional configuration for recurring credit hierarchy access control
type V2ContractEditParamsAddRecurringCreditsHierarchyConfiguration struct {
	ChildAccess param.Field[V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessUnion] `json:"child_access,required"`
}

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccess struct {
	Type        param.Field[V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessType] `json:"type,required"`
	ContractIDs param.Field[interface{}]                                                                  `json:"contract_ids"`
}

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccess) implementsV2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessUnion() {
}

// Satisfied by
// [V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs],
// [V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccess].
type V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessUnion()
}

type V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type param.Field[V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType] `json:"type,required"`
}

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type param.Field[V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType] `json:"type,required"`
}

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs param.Field[[]string]                                                                                                          `json:"contract_ids,required" format:"uuid"`
	Type        param.Field[V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType] `json:"type,required"`
}

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessType string

const (
	V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessTypeAll         V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessTypeNone        V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessTypeContractIDs V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessTypeAll, V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessTypeNone, V2ContractEditParamsAddRecurringCreditsHierarchyConfigurationChildAccessTypeContractIDs:
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
// be created aligned with the recurring commit's starting_at rather than the usage
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

type V2ContractEditParamsAddRecurringCreditsSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V2ContractEditParamsAddRecurringCreditsSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attach a subscription to the recurring commit/credit.
type V2ContractEditParamsAddRecurringCreditsSubscriptionConfig struct {
	ApplySeatIncreaseConfig param.Field[V2ContractEditParamsAddRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig] `json:"apply_seat_increase_config,required"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID param.Field[string] `json:"subscription_id,required"`
	// If set to POOLED, allocation added per seat is pooled across the account.
	Allocation param.Field[V2ContractEditParamsAddRecurringCreditsSubscriptionConfigAllocation] `json:"allocation"`
}

func (r V2ContractEditParamsAddRecurringCreditsSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated param.Field[bool] `json:"is_prorated,required"`
}

func (r V2ContractEditParamsAddRecurringCreditsSubscriptionConfigApplySeatIncreaseConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set to POOLED, allocation added per seat is pooled across the account.
type V2ContractEditParamsAddRecurringCreditsSubscriptionConfigAllocation string

const (
	V2ContractEditParamsAddRecurringCreditsSubscriptionConfigAllocationPooled V2ContractEditParamsAddRecurringCreditsSubscriptionConfigAllocation = "POOLED"
)

func (r V2ContractEditParamsAddRecurringCreditsSubscriptionConfigAllocation) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddRecurringCreditsSubscriptionConfigAllocationPooled:
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
	Schedule     param.Field[V2ContractEditParamsAddScheduledChargesSchedule] `json:"schedule,required"`
	CustomFields param.Field[map[string]string]                               `json:"custom_fields"`
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
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// This field is only applicable to commit invoice schedules. If true, this
	// schedule will not generate an invoice.
	DoNotInvoice param.Field[bool] `json:"do_not_invoice"`
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

type V2ContractEditParamsAddSpendThresholdConfiguration struct {
	Commit param.Field[V2ContractEditParamsAddSpendThresholdConfigurationCommit] `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         param.Field[bool]                                                                `json:"is_enabled,required"`
	PaymentGateConfig param.Field[V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfig] `json:"payment_gate_config,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount param.Field[float64] `json:"threshold_amount,required"`
}

func (r V2ContractEditParamsAddSpendThresholdConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddSpendThresholdConfigurationCommit struct {
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID   param.Field[string] `json:"product_id,required"`
	Description param.Field[string] `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name param.Field[string] `json:"name"`
}

func (r V2ContractEditParamsAddSpendThresholdConfigurationCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType param.Field[V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType] `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig param.Field[V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig] `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig param.Field[V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfig] `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType param.Field[V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxType] `json:"tax_type"`
}

func (r V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType string

const (
	V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone     V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "NONE"
	V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe   V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "STRIPE"
	V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone, V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe, V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount param.Field[float64] `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Field[string] `json:"tax_name"`
}

func (r V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType param.Field[V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType] `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata param.Field[map[string]string] `json:"invoice_metadata"`
}

func (r V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If left blank, will default to INVOICE
type V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType string

const (
	V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice       V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "INVOICE"
	V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice, V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxType string

const (
	V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxTypeNone          V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxType = "NONE"
	V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe        V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxType = "STRIPE"
	V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok         V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxType = "ANROK"
	V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxType = "PRECALCULATED"
)

func (r V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxTypeNone, V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe, V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok, V2ContractEditParamsAddSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated:
		return true
	}
	return false
}

type V2ContractEditParamsAddSubscription struct {
	CollectionSchedule param.Field[V2ContractEditParamsAddSubscriptionsCollectionSchedule] `json:"collection_schedule,required"`
	InitialQuantity    param.Field[float64]                                                `json:"initial_quantity,required"`
	Proration          param.Field[V2ContractEditParamsAddSubscriptionsProration]          `json:"proration,required"`
	SubscriptionRate   param.Field[V2ContractEditParamsAddSubscriptionsSubscriptionRate]   `json:"subscription_rate,required"`
	CustomFields       param.Field[map[string]string]                                      `json:"custom_fields"`
	Description        param.Field[string]                                                 `json:"description"`
	// Exclusive end time for the subscription. If not provided, subscription inherits
	// contract end date.
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Name         param.Field[string]    `json:"name"`
	// Inclusive start time for the subscription. If not provided, defaults to contract
	// start date
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
	// A temporary ID used to reference the subscription in recurring commit/credit
	// subscription configs created within the same payload.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r V2ContractEditParamsAddSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsAddSubscriptionsCollectionSchedule string

const (
	V2ContractEditParamsAddSubscriptionsCollectionScheduleAdvance V2ContractEditParamsAddSubscriptionsCollectionSchedule = "ADVANCE"
	V2ContractEditParamsAddSubscriptionsCollectionScheduleArrears V2ContractEditParamsAddSubscriptionsCollectionSchedule = "ARREARS"
)

func (r V2ContractEditParamsAddSubscriptionsCollectionSchedule) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddSubscriptionsCollectionScheduleAdvance, V2ContractEditParamsAddSubscriptionsCollectionScheduleArrears:
		return true
	}
	return false
}

type V2ContractEditParamsAddSubscriptionsProration struct {
	// Indicates how mid-period quantity adjustments are invoiced.
	// **BILL_IMMEDIATELY**: Only available when collection schedule is `ADVANCE`. The
	// quantity increase will be billed immediately on the scheduled date.
	// **BILL_ON_NEXT_COLLECTION_DATE**: The quantity increase will be billed for
	// in-arrears at the end of the period.
	InvoiceBehavior param.Field[V2ContractEditParamsAddSubscriptionsProrationInvoiceBehavior] `json:"invoice_behavior"`
	// Indicates if the partial period will be prorated or charged a full amount.
	IsProrated param.Field[bool] `json:"is_prorated"`
}

func (r V2ContractEditParamsAddSubscriptionsProration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates how mid-period quantity adjustments are invoiced.
// **BILL_IMMEDIATELY**: Only available when collection schedule is `ADVANCE`. The
// quantity increase will be billed immediately on the scheduled date.
// **BILL_ON_NEXT_COLLECTION_DATE**: The quantity increase will be billed for
// in-arrears at the end of the period.
type V2ContractEditParamsAddSubscriptionsProrationInvoiceBehavior string

const (
	V2ContractEditParamsAddSubscriptionsProrationInvoiceBehaviorBillImmediately          V2ContractEditParamsAddSubscriptionsProrationInvoiceBehavior = "BILL_IMMEDIATELY"
	V2ContractEditParamsAddSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate V2ContractEditParamsAddSubscriptionsProrationInvoiceBehavior = "BILL_ON_NEXT_COLLECTION_DATE"
)

func (r V2ContractEditParamsAddSubscriptionsProrationInvoiceBehavior) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddSubscriptionsProrationInvoiceBehaviorBillImmediately, V2ContractEditParamsAddSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate:
		return true
	}
	return false
}

type V2ContractEditParamsAddSubscriptionsSubscriptionRate struct {
	// Frequency to bill subscription with. Together with product_id, must match
	// existing rate on the rate card.
	BillingFrequency param.Field[V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequency] `json:"billing_frequency,required"`
	// Must be subscription type product
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
}

func (r V2ContractEditParamsAddSubscriptionsSubscriptionRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Frequency to bill subscription with. Together with product_id, must match
// existing rate on the rate card.
type V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequency string

const (
	V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequencyMonthly   V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequency = "MONTHLY"
	V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequencyQuarterly V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequency = "QUARTERLY"
	V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequencyAnnual    V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequency = "ANNUAL"
	V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequencyWeekly    V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequency = "WEEKLY"
)

func (r V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequency) IsKnown() bool {
	switch r {
	case V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequencyMonthly, V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequencyQuarterly, V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequencyAnnual, V2ContractEditParamsAddSubscriptionsSubscriptionRateBillingFrequencyWeekly:
		return true
	}
	return false
}

type V2ContractEditParamsArchiveCommit struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V2ContractEditParamsArchiveCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsArchiveCredit struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V2ContractEditParamsArchiveCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsArchiveScheduledCharge struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V2ContractEditParamsArchiveScheduledCharge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsRemoveOverride struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r V2ContractEditParamsRemoveOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCommit struct {
	CommitID       param.Field[string]                                          `json:"commit_id,required" format:"uuid"`
	AccessSchedule param.Field[V2ContractEditParamsUpdateCommitsAccessSchedule] `json:"access_schedule"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration param.Field[V2ContractEditParamsUpdateCommitsHierarchyConfiguration] `json:"hierarchy_configuration"`
	InvoiceSchedule        param.Field[V2ContractEditParamsUpdateCommitsInvoiceSchedule]        `json:"invoice_schedule"`
	NetsuiteSalesOrderID   param.Field[string]                                                  `json:"netsuite_sales_order_id"`
	Priority               param.Field[float64]                                                 `json:"priority"`
	ProductID              param.Field[string]                                                  `json:"product_id" format:"uuid"`
	RolloverFraction       param.Field[float64]                                                 `json:"rollover_fraction"`
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

// Optional configuration for commit hierarchy access control
type V2ContractEditParamsUpdateCommitsHierarchyConfiguration struct {
	ChildAccess param.Field[V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessUnion] `json:"child_access,required"`
}

func (r V2ContractEditParamsUpdateCommitsHierarchyConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccess struct {
	Type        param.Field[V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessType] `json:"type,required"`
	ContractIDs param.Field[interface{}]                                                            `json:"contract_ids"`
}

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccess) implementsV2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessUnion() {
}

// Satisfied by
// [V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs],
// [V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccess].
type V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessUnion()
}

type V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type param.Field[V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType] `json:"type,required"`
}

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type param.Field[V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType] `json:"type,required"`
}

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs param.Field[[]string]                                                                                                    `json:"contract_ids,required" format:"uuid"`
	Type        param.Field[V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType] `json:"type,required"`
}

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessType string

const (
	V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessTypeAll         V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessTypeNone        V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessTypeContractIDs V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessTypeAll, V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessTypeNone, V2ContractEditParamsUpdateCommitsHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
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
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration param.Field[V2ContractEditParamsUpdateCreditsHierarchyConfiguration] `json:"hierarchy_configuration"`
	NetsuiteSalesOrderID   param.Field[string]                                                  `json:"netsuite_sales_order_id"`
	Priority               param.Field[float64]                                                 `json:"priority"`
	ProductID              param.Field[string]                                                  `json:"product_id" format:"uuid"`
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

// Optional configuration for commit hierarchy access control
type V2ContractEditParamsUpdateCreditsHierarchyConfiguration struct {
	ChildAccess param.Field[V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessUnion] `json:"child_access,required"`
}

func (r V2ContractEditParamsUpdateCreditsHierarchyConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccess struct {
	Type        param.Field[V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessType] `json:"type,required"`
	ContractIDs param.Field[interface{}]                                                            `json:"contract_ids"`
}

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccess) implementsV2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessUnion() {
}

// Satisfied by
// [V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll],
// [V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone],
// [V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs],
// [V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccess].
type V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessUnion interface {
	implementsV2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessUnion()
}

type V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll struct {
	Type param.Field[V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType] `json:"type,required"`
}

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAll) implementsV2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType string

const (
	V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType = "ALL"
)

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessAllTypeAll:
		return true
	}
	return false
}

type V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone struct {
	Type param.Field[V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType] `json:"type,required"`
}

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNone) implementsV2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType string

const (
	V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType = "NONE"
)

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessNoneTypeNone:
		return true
	}
	return false
}

type V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs struct {
	ContractIDs param.Field[[]string]                                                                                                    `json:"contract_ids,required" format:"uuid"`
	Type        param.Field[V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType] `json:"type,required"`
}

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDs) implementsV2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessUnion() {
}

type V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType string

const (
	V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessCommitHierarchyChildAccessContractIDsTypeContractIDs:
		return true
	}
	return false
}

type V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessType string

const (
	V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessTypeAll         V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessType = "ALL"
	V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessTypeNone        V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessType = "NONE"
	V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessTypeContractIDs V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessType = "CONTRACT_IDS"
)

func (r V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessTypeAll, V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessTypeNone, V2ContractEditParamsUpdateCreditsHierarchyConfigurationChildAccessTypeContractIDs:
		return true
	}
	return false
}

type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfiguration struct {
	Commit param.Field[V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommit] `json:"commit"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold
	// commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID param.Field[string] `json:"custom_credit_type_id" format:"uuid"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         param.Field[bool]                                                                            `json:"is_enabled"`
	PaymentGateConfig param.Field[V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfig] `json:"payment_gate_config"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount param.Field[float64] `json:"recharge_to_amount"`
	// Specify the threshold amount for the contract. Each time the contract's balance
	// lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount param.Field[float64] `json:"threshold_amount"`
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommit struct {
	// Which products the threshold commit applies to. If both applicable_product_ids
	// and applicable_product_tags are not provided, the commit applies to all
	// products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	Description           param.Field[string]   `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name param.Field[string] `json:"name"`
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID param.Field[string] `json:"product_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers param.Field[[]V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommitSpecifier] `json:"specifiers"`
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommitSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationCommitSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType param.Field[V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType] `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig param.Field[V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig] `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig param.Field[V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig] `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType param.Field[V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType] `json:"tax_type"`
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType string

const (
	V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone     V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "NONE"
	V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe   V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "STRIPE"
	V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeNone, V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe, V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount param.Field[float64] `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Field[string] `json:"tax_name"`
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType param.Field[V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType] `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata param.Field[map[string]string] `json:"invoice_metadata"`
}

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If left blank, will default to INVOICE
type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType string

const (
	V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice       V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "INVOICE"
	V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice, V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType string

const (
	V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone          V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "NONE"
	V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe        V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "STRIPE"
	V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok         V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "ANROK"
	V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType = "PRECALCULATED"
)

func (r V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeNone, V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeStripe, V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypeAnrok, V2ContractEditParamsUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigTaxTypePrecalculated:
		return true
	}
	return false
}

type V2ContractEditParamsUpdateRecurringCommit struct {
	RecurringCommitID param.Field[string]                                                  `json:"recurring_commit_id,required" format:"uuid"`
	AccessAmount      param.Field[V2ContractEditParamsUpdateRecurringCommitsAccessAmount]  `json:"access_amount"`
	EndingBefore      param.Field[time.Time]                                               `json:"ending_before" format:"date-time"`
	InvoiceAmount     param.Field[V2ContractEditParamsUpdateRecurringCommitsInvoiceAmount] `json:"invoice_amount"`
}

func (r V2ContractEditParamsUpdateRecurringCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateRecurringCommitsAccessAmount struct {
	Quantity  param.Field[float64] `json:"quantity"`
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V2ContractEditParamsUpdateRecurringCommitsAccessAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateRecurringCommitsInvoiceAmount struct {
	Quantity  param.Field[float64] `json:"quantity"`
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V2ContractEditParamsUpdateRecurringCommitsInvoiceAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateRecurringCredit struct {
	RecurringCreditID param.Field[string]                                                 `json:"recurring_credit_id,required" format:"uuid"`
	AccessAmount      param.Field[V2ContractEditParamsUpdateRecurringCreditsAccessAmount] `json:"access_amount"`
	EndingBefore      param.Field[time.Time]                                              `json:"ending_before" format:"date-time"`
}

func (r V2ContractEditParamsUpdateRecurringCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateRecurringCreditsAccessAmount struct {
	Quantity  param.Field[float64] `json:"quantity"`
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V2ContractEditParamsUpdateRecurringCreditsAccessAmount) MarshalJSON() (data []byte, err error) {
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

type V2ContractEditParamsUpdateSpendThresholdConfiguration struct {
	Commit param.Field[V2ContractEditParamsUpdateSpendThresholdConfigurationCommit] `json:"commit"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state.
	IsEnabled         param.Field[bool]                                                                   `json:"is_enabled"`
	PaymentGateConfig param.Field[V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfig] `json:"payment_gate_config"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount param.Field[float64] `json:"threshold_amount"`
}

func (r V2ContractEditParamsUpdateSpendThresholdConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateSpendThresholdConfigurationCommit struct {
	Description param.Field[string] `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name param.Field[string] `json:"name"`
	// The commit product that will be used to generate the line item for commit
	// payment.
	ProductID param.Field[string] `json:"product_id"`
}

func (r V2ContractEditParamsUpdateSpendThresholdConfigurationCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfig struct {
	// Gate access to the commit balance based on successful collection of payment.
	// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
	// facilitate payment using your own payment integration. Select NONE if you do not
	// wish to payment gate the commit balance.
	PaymentGateType param.Field[V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType] `json:"payment_gate_type,required"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig param.Field[V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig] `json:"precalculated_tax_config"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig param.Field[V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfig] `json:"stripe_config"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
	// not wish Metronome to calculate tax on your behalf. Leaving this field blank
	// will default to NONE.
	TaxType param.Field[V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxType] `json:"tax_type"`
}

func (r V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Gate access to the commit balance based on successful collection of payment.
// Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to
// facilitate payment using your own payment integration. Select NONE if you do not
// wish to payment gate the commit balance.
type V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType string

const (
	V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone     V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "NONE"
	V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe   V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "STRIPE"
	V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType = "EXTERNAL"
)

func (r V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeNone, V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeStripe, V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPaymentGateTypeExternal:
		return true
	}
	return false
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and
	// denomination as the commit's invoice schedule
	TaxAmount param.Field[float64] `json:"tax_amount,required"`
	// Name of the tax to be applied. This may be used in an invoice line item
	// description.
	TaxName param.Field[string] `json:"tax_name"`
}

func (r V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfig struct {
	// If left blank, will default to INVOICE
	PaymentType param.Field[V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType] `json:"payment_type,required"`
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as
	// your payment type.
	InvoiceMetadata param.Field[map[string]string] `json:"invoice_metadata"`
}

func (r V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If left blank, will default to INVOICE
type V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType string

const (
	V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice       V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "INVOICE"
	V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType = "PAYMENT_INTENT"
)

func (r V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypeInvoice, V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigPaymentTypePaymentIntent:
		return true
	}
	return false
}

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do
// not wish Metronome to calculate tax on your behalf. Leaving this field blank
// will default to NONE.
type V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxType string

const (
	V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeNone          V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxType = "NONE"
	V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe        V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxType = "STRIPE"
	V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok         V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxType = "ANROK"
	V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxType = "PRECALCULATED"
)

func (r V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxType) IsKnown() bool {
	switch r {
	case V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeNone, V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeStripe, V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypeAnrok, V2ContractEditParamsUpdateSpendThresholdConfigurationPaymentGateConfigTaxTypePrecalculated:
		return true
	}
	return false
}

type V2ContractEditParamsUpdateSubscription struct {
	SubscriptionID param.Field[string]    `json:"subscription_id,required" format:"uuid"`
	EndingBefore   param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// Quantity changes are applied on the effective date based on the order which they
	// are sent. For example, if I scheduled the quantity to be 12 on May 21 and then
	// scheduled a quantity delta change of -1, the result from that day would be 11.
	QuantityUpdates param.Field[[]V2ContractEditParamsUpdateSubscriptionsQuantityUpdate] `json:"quantity_updates"`
}

func (r V2ContractEditParamsUpdateSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditParamsUpdateSubscriptionsQuantityUpdate struct {
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// The new quantity for the subscription. Must be provided if quantity_delta is not
	// provided. Must be non-negative.
	Quantity param.Field[float64] `json:"quantity"`
	// The delta to add to the subscription's quantity. Must be provided if quantity is
	// not provided. Can't be zero. It also can't result in a negative quantity on the
	// subscription.
	QuantityDelta param.Field[float64] `json:"quantity_delta"`
}

func (r V2ContractEditParamsUpdateSubscriptionsQuantityUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractEditCommitParams struct {
	// ID of the commit to edit
	CommitID param.Field[string] `json:"commit_id,required" format:"uuid"`
	// ID of the customer whose commit is being edited
	CustomerID     param.Field[string]                                   `json:"customer_id,required" format:"uuid"`
	AccessSchedule param.Field[V2ContractEditCommitParamsAccessSchedule] `json:"access_schedule"`
	// Which products the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If applicable_product_ids,
	// applicable_product_tags or specifiers are not provided, the commit applies to
	// all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// ID of contract to use for invoicing
	InvoiceContractID param.Field[string]                                    `json:"invoice_contract_id" format:"uuid"`
	InvoiceSchedule   param.Field[V2ContractEditCommitParamsInvoiceSchedule] `json:"invoice_schedule"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority  param.Field[float64] `json:"priority"`
	ProductID param.Field[string]  `json:"product_id" format:"uuid"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers param.Field[[]V2ContractEditCommitParamsSpecifier] `json:"specifiers"`
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

type V2ContractEditCommitParamsSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V2ContractEditCommitParamsSpecifier) MarshalJSON() (data []byte, err error) {
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
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority  param.Field[float64] `json:"priority"`
	ProductID param.Field[string]  `json:"product_id" format:"uuid"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	// Instead, to target usage by product or product tag, pass those values in the
	// body of `specifiers`.
	Specifiers param.Field[[]V2ContractEditCreditParamsSpecifier] `json:"specifiers"`
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

type V2ContractEditCreditParamsSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V2ContractEditCreditParamsSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2ContractGetEditHistoryParams struct {
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r V2ContractGetEditHistoryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
