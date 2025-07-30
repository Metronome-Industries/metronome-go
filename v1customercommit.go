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
	"github.com/tidwall/gjson"
)

// V1CustomerCommitService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerCommitService] method instead.
type V1CustomerCommitService struct {
	Options []option.RequestOption
}

// NewV1CustomerCommitService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerCommitService(opts ...option.RequestOption) (r *V1CustomerCommitService) {
	r = &V1CustomerCommitService{}
	r.Options = opts
	return
}

// Create a new commit at the customer level.
func (r *V1CustomerCommitService) New(ctx context.Context, body V1CustomerCommitNewParams, opts ...option.RequestOption) (res *V1CustomerCommitNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCommits/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List commits.
func (r *V1CustomerCommitService) List(ctx context.Context, body V1CustomerCommitListParams, opts ...option.RequestOption) (res *V1CustomerCommitListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCommits/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Pull forward the end date of a prepaid commit. Use the "edit a commit" endpoint
// to extend the end date of a prepaid commit, or to make other edits to the
// commit.
func (r *V1CustomerCommitService) UpdateEndDate(ctx context.Context, body V1CustomerCommitUpdateEndDateParams, opts ...option.RequestOption) (res *V1CustomerCommitUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCommits/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1CustomerCommitNewResponse struct {
	Data V1CustomerCommitNewResponseData `json:"data,required"`
	JSON v1CustomerCommitNewResponseJSON `json:"-"`
}

// v1CustomerCommitNewResponseJSON contains the JSON metadata for the struct
// [V1CustomerCommitNewResponse]
type v1CustomerCommitNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitNewResponseData struct {
	ID   string                              `json:"id,required" format:"uuid"`
	JSON v1CustomerCommitNewResponseDataJSON `json:"-"`
}

// v1CustomerCommitNewResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerCommitNewResponseData]
type v1CustomerCommitNewResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponse struct {
	Data     []V1CustomerCommitListResponseData `json:"data,required"`
	NextPage string                             `json:"next_page,required,nullable"`
	JSON     v1CustomerCommitListResponseJSON   `json:"-"`
}

// v1CustomerCommitListResponseJSON contains the JSON metadata for the struct
// [V1CustomerCommitListResponse]
type v1CustomerCommitListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseData struct {
	ID      string                                  `json:"id,required" format:"uuid"`
	Product V1CustomerCommitListResponseDataProduct `json:"product,required"`
	Type    V1CustomerCommitListResponseDataType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule V1CustomerCommitListResponseDataAccessSchedule `json:"access_schedule"`
	// (DEPRECATED) Use access_schedule + invoice_schedule instead.
	Amount                float64  `json:"amount"`
	ApplicableContractIDs []string `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string `json:"applicable_product_tags"`
	// RFC 3339 timestamp indicating when the commit was archived. If not provided, the
	// commit is not archived.
	ArchivedAt time.Time `json:"archived_at" format:"date-time"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64                                  `json:"balance"`
	Contract     V1CustomerCommitListResponseDataContract `json:"contract"`
	CustomFields map[string]string                        `json:"custom_fields"`
	Description  string                                   `json:"description"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration V1CustomerCommitListResponseDataHierarchyConfiguration `json:"hierarchy_configuration"`
	// The contract that this commit will be billed on.
	InvoiceContract V1CustomerCommitListResponseDataInvoiceContract `json:"invoice_contract"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule V1CustomerCommitListResponseDataInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger []V1CustomerCommitListResponseDataLedger `json:"ledger"`
	Name   string                                   `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority         float64                                        `json:"priority"`
	RateType         V1CustomerCommitListResponseDataRateType       `json:"rate_type"`
	RolledOverFrom   V1CustomerCommitListResponseDataRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction float64                                        `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []V1CustomerCommitListResponseDataSpecifier `json:"specifiers"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey string                               `json:"uniqueness_key"`
	JSON          v1CustomerCommitListResponseDataJSON `json:"-"`
}

// v1CustomerCommitListResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerCommitListResponseData]
type v1CustomerCommitListResponseDataJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	Amount                  apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	ArchivedAt              apijson.Field
	Balance                 apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	HierarchyConfiguration  apijson.Field
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
	Specifiers              apijson.Field
	UniquenessKey           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseDataProduct struct {
	ID   string                                      `json:"id,required" format:"uuid"`
	Name string                                      `json:"name,required"`
	JSON v1CustomerCommitListResponseDataProductJSON `json:"-"`
}

// v1CustomerCommitListResponseDataProductJSON contains the JSON metadata for the
// struct [V1CustomerCommitListResponseDataProduct]
type v1CustomerCommitListResponseDataProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataProductJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseDataType string

const (
	V1CustomerCommitListResponseDataTypePrepaid  V1CustomerCommitListResponseDataType = "PREPAID"
	V1CustomerCommitListResponseDataTypePostpaid V1CustomerCommitListResponseDataType = "POSTPAID"
)

func (r V1CustomerCommitListResponseDataType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseDataTypePrepaid, V1CustomerCommitListResponseDataTypePostpaid:
		return true
	}
	return false
}

// The schedule that the customer will gain access to the credits purposed with
// this commit.
type V1CustomerCommitListResponseDataAccessSchedule struct {
	ScheduleItems []V1CustomerCommitListResponseDataAccessScheduleScheduleItem `json:"schedule_items,required"`
	CreditType    V1CustomerCommitListResponseDataAccessScheduleCreditType     `json:"credit_type"`
	JSON          v1CustomerCommitListResponseDataAccessScheduleJSON           `json:"-"`
}

// v1CustomerCommitListResponseDataAccessScheduleJSON contains the JSON metadata
// for the struct [V1CustomerCommitListResponseDataAccessSchedule]
type v1CustomerCommitListResponseDataAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseDataAccessScheduleScheduleItem struct {
	ID           string                                                         `json:"id,required" format:"uuid"`
	Amount       float64                                                        `json:"amount,required"`
	EndingBefore time.Time                                                      `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                                                      `json:"starting_at,required" format:"date-time"`
	JSON         v1CustomerCommitListResponseDataAccessScheduleScheduleItemJSON `json:"-"`
}

// v1CustomerCommitListResponseDataAccessScheduleScheduleItemJSON contains the JSON
// metadata for the struct
// [V1CustomerCommitListResponseDataAccessScheduleScheduleItem]
type v1CustomerCommitListResponseDataAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataAccessScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseDataAccessScheduleCreditType struct {
	ID   string                                                       `json:"id,required" format:"uuid"`
	Name string                                                       `json:"name,required"`
	JSON v1CustomerCommitListResponseDataAccessScheduleCreditTypeJSON `json:"-"`
}

// v1CustomerCommitListResponseDataAccessScheduleCreditTypeJSON contains the JSON
// metadata for the struct
// [V1CustomerCommitListResponseDataAccessScheduleCreditType]
type v1CustomerCommitListResponseDataAccessScheduleCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataAccessScheduleCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataAccessScheduleCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseDataContract struct {
	ID   string                                       `json:"id,required" format:"uuid"`
	JSON v1CustomerCommitListResponseDataContractJSON `json:"-"`
}

// v1CustomerCommitListResponseDataContractJSON contains the JSON metadata for the
// struct [V1CustomerCommitListResponseDataContract]
type v1CustomerCommitListResponseDataContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataContractJSON) RawJSON() string {
	return r.raw
}

// Optional configuration for commit hierarchy access control
type V1CustomerCommitListResponseDataHierarchyConfiguration struct {
	ChildAccess V1CustomerCommitListResponseDataHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        v1CustomerCommitListResponseDataHierarchyConfigurationJSON        `json:"-"`
}

// v1CustomerCommitListResponseDataHierarchyConfigurationJSON contains the JSON
// metadata for the struct [V1CustomerCommitListResponseDataHierarchyConfiguration]
type v1CustomerCommitListResponseDataHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseDataHierarchyConfigurationChildAccess struct {
	Type V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                           `json:"contract_ids"`
	JSON        v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessJSON `json:"-"`
	union       V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessUnion
}

// v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessJSON contains
// the JSON metadata for the struct
// [V1CustomerCommitListResponseDataHierarchyConfigurationChildAccess]
type v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *V1CustomerCommitListResponseDataHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = V1CustomerCommitListResponseDataHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessType],
// [V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessType],
// [V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObject].
func (r V1CustomerCommitListResponseDataHierarchyConfigurationChildAccess) AsUnion() V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessType],
// [V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessType] or
// [V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObject].
type V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessUnion interface {
	implementsV1CustomerCommitListResponseDataHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObject{}),
		},
	)
}

type V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessType struct {
	Type V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessTypeType `json:"type,required"`
	JSON v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessTypeJSON `json:"-"`
}

// v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessTypeJSON
// contains the JSON metadata for the struct
// [V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessType]
type v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessTypeJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessType) implementsV1CustomerCommitListResponseDataHierarchyConfigurationChildAccess() {
}

type V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessTypeType string

const (
	V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessTypeTypeAll V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessTypeType = "ALL"
)

func (r V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessTypeType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessTypeTypeAll:
		return true
	}
	return false
}

type V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObject struct {
	ContractIDs []string                                                                    `json:"contract_ids,required" format:"uuid"`
	Type        V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObjectType `json:"type,required"`
	JSON        v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObjectJSON `json:"-"`
}

// v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObjectJSON
// contains the JSON metadata for the struct
// [V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObject]
type v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObjectJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObjectJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObject) implementsV1CustomerCommitListResponseDataHierarchyConfigurationChildAccess() {
}

type V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObjectType string

const (
	V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObjectTypeContractIDs V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObjectType = "CONTRACT_IDS"
)

func (r V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObjectType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseDataHierarchyConfigurationChildAccessObjectTypeContractIDs:
		return true
	}
	return false
}

// The contract that this commit will be billed on.
type V1CustomerCommitListResponseDataInvoiceContract struct {
	ID   string                                              `json:"id,required" format:"uuid"`
	JSON v1CustomerCommitListResponseDataInvoiceContractJSON `json:"-"`
}

// v1CustomerCommitListResponseDataInvoiceContractJSON contains the JSON metadata
// for the struct [V1CustomerCommitListResponseDataInvoiceContract]
type v1CustomerCommitListResponseDataInvoiceContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataInvoiceContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataInvoiceContractJSON) RawJSON() string {
	return r.raw
}

// The schedule that the customer will be invoiced for this commit.
type V1CustomerCommitListResponseDataInvoiceSchedule struct {
	CreditType    V1CustomerCommitListResponseDataInvoiceScheduleCreditType     `json:"credit_type"`
	ScheduleItems []V1CustomerCommitListResponseDataInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          v1CustomerCommitListResponseDataInvoiceScheduleJSON           `json:"-"`
}

// v1CustomerCommitListResponseDataInvoiceScheduleJSON contains the JSON metadata
// for the struct [V1CustomerCommitListResponseDataInvoiceSchedule]
type v1CustomerCommitListResponseDataInvoiceScheduleJSON struct {
	CreditType    apijson.Field
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataInvoiceScheduleJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseDataInvoiceScheduleCreditType struct {
	ID   string                                                        `json:"id,required" format:"uuid"`
	Name string                                                        `json:"name,required"`
	JSON v1CustomerCommitListResponseDataInvoiceScheduleCreditTypeJSON `json:"-"`
}

// v1CustomerCommitListResponseDataInvoiceScheduleCreditTypeJSON contains the JSON
// metadata for the struct
// [V1CustomerCommitListResponseDataInvoiceScheduleCreditType]
type v1CustomerCommitListResponseDataInvoiceScheduleCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataInvoiceScheduleCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataInvoiceScheduleCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseDataInvoiceScheduleScheduleItem struct {
	ID        string                                                          `json:"id,required" format:"uuid"`
	Amount    float64                                                         `json:"amount,required"`
	Quantity  float64                                                         `json:"quantity,required"`
	Timestamp time.Time                                                       `json:"timestamp,required" format:"date-time"`
	UnitPrice float64                                                         `json:"unit_price,required"`
	InvoiceID string                                                          `json:"invoice_id,nullable" format:"uuid"`
	JSON      v1CustomerCommitListResponseDataInvoiceScheduleScheduleItemJSON `json:"-"`
}

// v1CustomerCommitListResponseDataInvoiceScheduleScheduleItemJSON contains the
// JSON metadata for the struct
// [V1CustomerCommitListResponseDataInvoiceScheduleScheduleItem]
type v1CustomerCommitListResponseDataInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	InvoiceID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataInvoiceScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseDataLedger struct {
	Amount        float64                                    `json:"amount,required"`
	Timestamp     time.Time                                  `json:"timestamp,required" format:"date-time"`
	Type          V1CustomerCommitListResponseDataLedgerType `json:"type,required"`
	ContractID    string                                     `json:"contract_id" format:"uuid"`
	InvoiceID     string                                     `json:"invoice_id" format:"uuid"`
	NewContractID string                                     `json:"new_contract_id" format:"uuid"`
	Reason        string                                     `json:"reason"`
	SegmentID     string                                     `json:"segment_id" format:"uuid"`
	JSON          v1CustomerCommitListResponseDataLedgerJSON `json:"-"`
	union         V1CustomerCommitListResponseDataLedgerUnion
}

// v1CustomerCommitListResponseDataLedgerJSON contains the JSON metadata for the
// struct [V1CustomerCommitListResponseDataLedger]
type v1CustomerCommitListResponseDataLedgerJSON struct {
	Amount        apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	ContractID    apijson.Field
	InvoiceID     apijson.Field
	NewContractID apijson.Field
	Reason        apijson.Field
	SegmentID     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r v1CustomerCommitListResponseDataLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *V1CustomerCommitListResponseDataLedger) UnmarshalJSON(data []byte) (err error) {
	*r = V1CustomerCommitListResponseDataLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V1CustomerCommitListResponseDataLedgerUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject].
func (r V1CustomerCommitListResponseDataLedger) AsUnion() V1CustomerCommitListResponseDataLedgerUnion {
	return r.union
}

// Union satisfied by [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject],
// [V1CustomerCommitListResponseDataLedgerObject] or
// [V1CustomerCommitListResponseDataLedgerObject].
type V1CustomerCommitListResponseDataLedgerUnion interface {
	implementsV1CustomerCommitListResponseDataLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1CustomerCommitListResponseDataLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCommitListResponseDataLedgerObject{}),
		},
	)
}

type V1CustomerCommitListResponseDataLedgerObject struct {
	Amount    float64                                          `json:"amount,required"`
	SegmentID string                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                        `json:"timestamp,required" format:"date-time"`
	Type      V1CustomerCommitListResponseDataLedgerObjectType `json:"type,required"`
	JSON      v1CustomerCommitListResponseDataLedgerObjectJSON `json:"-"`
}

// v1CustomerCommitListResponseDataLedgerObjectJSON contains the JSON metadata for
// the struct [V1CustomerCommitListResponseDataLedgerObject]
type v1CustomerCommitListResponseDataLedgerObjectJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataLedgerObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataLedgerObjectJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCommitListResponseDataLedgerObject) implementsV1CustomerCommitListResponseDataLedger() {
}

type V1CustomerCommitListResponseDataLedgerObjectType string

const (
	V1CustomerCommitListResponseDataLedgerObjectTypePrepaidCommitSegmentStart V1CustomerCommitListResponseDataLedgerObjectType = "PREPAID_COMMIT_SEGMENT_START"
)

func (r V1CustomerCommitListResponseDataLedgerObjectType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseDataLedgerObjectTypePrepaidCommitSegmentStart:
		return true
	}
	return false
}

type V1CustomerCommitListResponseDataLedgerType string

const (
	V1CustomerCommitListResponseDataLedgerTypePrepaidCommitSegmentStart               V1CustomerCommitListResponseDataLedgerType = "PREPAID_COMMIT_SEGMENT_START"
	V1CustomerCommitListResponseDataLedgerTypePrepaidCommitAutomatedInvoiceDeduction  V1CustomerCommitListResponseDataLedgerType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	V1CustomerCommitListResponseDataLedgerTypePrepaidCommitRollover                   V1CustomerCommitListResponseDataLedgerType = "PREPAID_COMMIT_ROLLOVER"
	V1CustomerCommitListResponseDataLedgerTypePrepaidCommitExpiration                 V1CustomerCommitListResponseDataLedgerType = "PREPAID_COMMIT_EXPIRATION"
	V1CustomerCommitListResponseDataLedgerTypePrepaidCommitCanceled                   V1CustomerCommitListResponseDataLedgerType = "PREPAID_COMMIT_CANCELED"
	V1CustomerCommitListResponseDataLedgerTypePrepaidCommitCredited                   V1CustomerCommitListResponseDataLedgerType = "PREPAID_COMMIT_CREDITED"
	V1CustomerCommitListResponseDataLedgerTypePrepaidCommitSeatBasedAdjustment        V1CustomerCommitListResponseDataLedgerType = "PREPAID_COMMIT_SEAT_BASED_ADJUSTMENT"
	V1CustomerCommitListResponseDataLedgerTypePostpaidCommitInitialBalance            V1CustomerCommitListResponseDataLedgerType = "POSTPAID_COMMIT_INITIAL_BALANCE"
	V1CustomerCommitListResponseDataLedgerTypePostpaidCommitAutomatedInvoiceDeduction V1CustomerCommitListResponseDataLedgerType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	V1CustomerCommitListResponseDataLedgerTypePostpaidCommitRollover                  V1CustomerCommitListResponseDataLedgerType = "POSTPAID_COMMIT_ROLLOVER"
	V1CustomerCommitListResponseDataLedgerTypePostpaidCommitTrueup                    V1CustomerCommitListResponseDataLedgerType = "POSTPAID_COMMIT_TRUEUP"
	V1CustomerCommitListResponseDataLedgerTypePrepaidCommitManual                     V1CustomerCommitListResponseDataLedgerType = "PREPAID_COMMIT_MANUAL"
	V1CustomerCommitListResponseDataLedgerTypePostpaidCommitManual                    V1CustomerCommitListResponseDataLedgerType = "POSTPAID_COMMIT_MANUAL"
	V1CustomerCommitListResponseDataLedgerTypePostpaidCommitExpiration                V1CustomerCommitListResponseDataLedgerType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r V1CustomerCommitListResponseDataLedgerType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseDataLedgerTypePrepaidCommitSegmentStart, V1CustomerCommitListResponseDataLedgerTypePrepaidCommitAutomatedInvoiceDeduction, V1CustomerCommitListResponseDataLedgerTypePrepaidCommitRollover, V1CustomerCommitListResponseDataLedgerTypePrepaidCommitExpiration, V1CustomerCommitListResponseDataLedgerTypePrepaidCommitCanceled, V1CustomerCommitListResponseDataLedgerTypePrepaidCommitCredited, V1CustomerCommitListResponseDataLedgerTypePrepaidCommitSeatBasedAdjustment, V1CustomerCommitListResponseDataLedgerTypePostpaidCommitInitialBalance, V1CustomerCommitListResponseDataLedgerTypePostpaidCommitAutomatedInvoiceDeduction, V1CustomerCommitListResponseDataLedgerTypePostpaidCommitRollover, V1CustomerCommitListResponseDataLedgerTypePostpaidCommitTrueup, V1CustomerCommitListResponseDataLedgerTypePrepaidCommitManual, V1CustomerCommitListResponseDataLedgerTypePostpaidCommitManual, V1CustomerCommitListResponseDataLedgerTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type V1CustomerCommitListResponseDataRateType string

const (
	V1CustomerCommitListResponseDataRateTypeCommitRate V1CustomerCommitListResponseDataRateType = "COMMIT_RATE"
	V1CustomerCommitListResponseDataRateTypeListRate   V1CustomerCommitListResponseDataRateType = "LIST_RATE"
)

func (r V1CustomerCommitListResponseDataRateType) IsKnown() bool {
	switch r {
	case V1CustomerCommitListResponseDataRateTypeCommitRate, V1CustomerCommitListResponseDataRateTypeListRate:
		return true
	}
	return false
}

type V1CustomerCommitListResponseDataRolledOverFrom struct {
	CommitID   string                                             `json:"commit_id,required" format:"uuid"`
	ContractID string                                             `json:"contract_id,required" format:"uuid"`
	JSON       v1CustomerCommitListResponseDataRolledOverFromJSON `json:"-"`
}

// v1CustomerCommitListResponseDataRolledOverFromJSON contains the JSON metadata
// for the struct [V1CustomerCommitListResponseDataRolledOverFrom]
type v1CustomerCommitListResponseDataRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataRolledOverFromJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitListResponseDataSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                      `json:"product_tags"`
	JSON        v1CustomerCommitListResponseDataSpecifierJSON `json:"-"`
}

// v1CustomerCommitListResponseDataSpecifierJSON contains the JSON metadata for the
// struct [V1CustomerCommitListResponseDataSpecifier]
type v1CustomerCommitListResponseDataSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerCommitListResponseDataSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitListResponseDataSpecifierJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitUpdateEndDateResponse struct {
	Data V1CustomerCommitUpdateEndDateResponseData `json:"data,required"`
	JSON v1CustomerCommitUpdateEndDateResponseJSON `json:"-"`
}

// v1CustomerCommitUpdateEndDateResponseJSON contains the JSON metadata for the
// struct [V1CustomerCommitUpdateEndDateResponse]
type v1CustomerCommitUpdateEndDateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitUpdateEndDateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitUpdateEndDateResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitUpdateEndDateResponseData struct {
	ID   string                                        `json:"id,required" format:"uuid"`
	JSON v1CustomerCommitUpdateEndDateResponseDataJSON `json:"-"`
}

// v1CustomerCommitUpdateEndDateResponseDataJSON contains the JSON metadata for the
// struct [V1CustomerCommitUpdateEndDateResponseData]
type v1CustomerCommitUpdateEndDateResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCommitUpdateEndDateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCommitUpdateEndDateResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCommitNewParams struct {
	// Schedule for distributing the commit to the customer. For "POSTPAID" commits
	// only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule param.Field[V1CustomerCommitNewParamsAccessSchedule] `json:"access_schedule,required"`
	CustomerID     param.Field[string]                                  `json:"customer_id,required" format:"uuid"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority param.Field[float64] `json:"priority,required"`
	// ID of the fixed product associated with the commit. This is required because
	// products are used to invoice the commit amount.
	ProductID param.Field[string]                        `json:"product_id,required" format:"uuid"`
	Type      param.Field[V1CustomerCommitNewParamsType] `json:"type,required"`
	// Which contract the commit applies to. If not provided, the commit applies to all
	// contracts.
	ApplicableContractIDs param.Field[[]string] `json:"applicable_contract_ids"`
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
	// The contract that this commit will be billed on. This is required for "POSTPAID"
	// commits and for "PREPAID" commits unless there is no invoice schedule above
	// (i.e., the commit is 'free').
	InvoiceContractID param.Field[string] `json:"invoice_contract_id" format:"uuid"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this
	// time and only one schedule item is allowed; the total must match
	// accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this
	// will be a "complimentary" commit with no invoice.
	InvoiceSchedule param.Field[V1CustomerCommitNewParamsInvoiceSchedule] `json:"invoice_schedule"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string]                            `json:"netsuite_sales_order_id"`
	RateType             param.Field[V1CustomerCommitNewParamsRateType] `json:"rate_type"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string] `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers param.Field[[]V1CustomerCommitNewParamsSpecifier] `json:"specifiers"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey param.Field[string] `json:"uniqueness_key"`
}

func (r V1CustomerCommitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the commit to the customer. For "POSTPAID" commits
// only one schedule item is allowed and amount must match invoice_schedule total.
type V1CustomerCommitNewParamsAccessSchedule struct {
	ScheduleItems param.Field[[]V1CustomerCommitNewParamsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r V1CustomerCommitNewParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitNewParamsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V1CustomerCommitNewParamsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitNewParamsType string

const (
	V1CustomerCommitNewParamsTypePrepaid  V1CustomerCommitNewParamsType = "PREPAID"
	V1CustomerCommitNewParamsTypePostpaid V1CustomerCommitNewParamsType = "POSTPAID"
)

func (r V1CustomerCommitNewParamsType) IsKnown() bool {
	switch r {
	case V1CustomerCommitNewParamsTypePrepaid, V1CustomerCommitNewParamsTypePostpaid:
		return true
	}
	return false
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match
// accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this
// will be a "complimentary" commit with no invoice.
type V1CustomerCommitNewParamsInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]V1CustomerCommitNewParamsInvoiceScheduleScheduleItem] `json:"schedule_items"`
}

func (r V1CustomerCommitNewParamsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                          `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r V1CustomerCommitNewParamsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDivided        V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionEach           V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDivided, V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency string

const (
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyMonthly    V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyQuarterly  V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencySemiAnnual V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyAnnual     V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyMonthly, V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyQuarterly, V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, V1CustomerCommitNewParamsInvoiceScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type V1CustomerCommitNewParamsInvoiceScheduleScheduleItem struct {
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

func (r V1CustomerCommitNewParamsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitNewParamsRateType string

const (
	V1CustomerCommitNewParamsRateTypeCommitRate V1CustomerCommitNewParamsRateType = "COMMIT_RATE"
	V1CustomerCommitNewParamsRateTypeListRate   V1CustomerCommitNewParamsRateType = "LIST_RATE"
)

func (r V1CustomerCommitNewParamsRateType) IsKnown() bool {
	switch r {
	case V1CustomerCommitNewParamsRateTypeCommitRate, V1CustomerCommitNewParamsRateTypeListRate:
		return true
	}
	return false
}

type V1CustomerCommitNewParamsSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V1CustomerCommitNewParamsSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitListParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	CommitID   param.Field[string] `json:"commit_id" format:"uuid"`
	// Include only commits that have access schedules that "cover" the provided date
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	// Include only commits that have any access before the provided date (exclusive)
	EffectiveBefore param.Field[time.Time] `json:"effective_before" format:"date-time"`
	// Include archived commits and commits from archived contracts.
	IncludeArchived param.Field[bool] `json:"include_archived"`
	// Include the balance in the response. Setting this flag may cause the query to be
	// slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include commits on the contract level.
	IncludeContractCommits param.Field[bool] `json:"include_contract_commits"`
	// Include commit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
	// The maximum number of commits to return. Defaults to 25.
	Limit param.Field[int64] `json:"limit"`
	// The next page token from a previous response.
	NextPage param.Field[string] `json:"next_page"`
	// Include only commits that have any access on or after the provided date
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V1CustomerCommitListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCommitUpdateEndDateParams struct {
	// ID of the commit to update. Only supports "PREPAID" commits.
	CommitID param.Field[string] `json:"commit_id,required" format:"uuid"`
	// ID of the customer whose commit is to be updated
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// RFC 3339 timestamp indicating when access to the commit will end and it will no
	// longer be possible to draw it down (exclusive). If not provided, the access will
	// not be updated.
	AccessEndingBefore param.Field[time.Time] `json:"access_ending_before" format:"date-time"`
	// RFC 3339 timestamp indicating when the commit will stop being invoiced
	// (exclusive). If not provided, the invoice schedule will not be updated.
	InvoicesEndingBefore param.Field[time.Time] `json:"invoices_ending_before" format:"date-time"`
}

func (r V1CustomerCommitUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
