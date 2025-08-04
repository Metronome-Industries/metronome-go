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

// V1CustomerCreditService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerCreditService] method instead.
type V1CustomerCreditService struct {
	Options []option.RequestOption
}

// NewV1CustomerCreditService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerCreditService(opts ...option.RequestOption) (r *V1CustomerCreditService) {
	r = &V1CustomerCreditService{}
	r.Options = opts
	return
}

// Create a new credit at the customer level.
func (r *V1CustomerCreditService) New(ctx context.Context, body V1CustomerCreditNewParams, opts ...option.RequestOption) (res *V1CustomerCreditNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCredits/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List credits.
func (r *V1CustomerCreditService) List(ctx context.Context, body V1CustomerCreditListParams, opts ...option.RequestOption) (res *V1CustomerCreditListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCredits/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Pull forward the end date of a credit. Use the "edit a credit" endpoint to
// extend the end date of a credit, or to make other edits to the credit.
func (r *V1CustomerCreditService) UpdateEndDate(ctx context.Context, body V1CustomerCreditUpdateEndDateParams, opts ...option.RequestOption) (res *V1CustomerCreditUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerCredits/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1CustomerCreditNewResponse struct {
	Data V1CustomerCreditNewResponseData `json:"data,required"`
	JSON v1CustomerCreditNewResponseJSON `json:"-"`
}

// v1CustomerCreditNewResponseJSON contains the JSON metadata for the struct
// [V1CustomerCreditNewResponse]
type v1CustomerCreditNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditNewResponseData struct {
	ID   string                              `json:"id,required" format:"uuid"`
	JSON v1CustomerCreditNewResponseDataJSON `json:"-"`
}

// v1CustomerCreditNewResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerCreditNewResponseData]
type v1CustomerCreditNewResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponse struct {
	Data     []V1CustomerCreditListResponseData `json:"data,required"`
	NextPage string                             `json:"next_page,required,nullable"`
	JSON     v1CustomerCreditListResponseJSON   `json:"-"`
}

// v1CustomerCreditListResponseJSON contains the JSON metadata for the struct
// [V1CustomerCreditListResponse]
type v1CustomerCreditListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseData struct {
	ID      string                                  `json:"id,required" format:"uuid"`
	Product V1CustomerCreditListResponseDataProduct `json:"product,required"`
	Type    V1CustomerCreditListResponseDataType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        V1CustomerCreditListResponseDataAccessSchedule `json:"access_schedule"`
	ApplicableContractIDs []string                                       `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string                                       `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string                                       `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64                                  `json:"balance"`
	Contract     V1CustomerCreditListResponseDataContract `json:"contract"`
	CustomFields map[string]string                        `json:"custom_fields"`
	Description  string                                   `json:"description"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration V1CustomerCreditListResponseDataHierarchyConfiguration `json:"hierarchy_configuration"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []V1CustomerCreditListResponseDataLedger `json:"ledger"`
	Name   string                                   `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64                                  `json:"priority"`
	RateType V1CustomerCreditListResponseDataRateType `json:"rate_type"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []V1CustomerCreditListResponseDataSpecifier `json:"specifiers"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey string                               `json:"uniqueness_key"`
	JSON          v1CustomerCreditListResponseDataJSON `json:"-"`
}

// v1CustomerCreditListResponseDataJSON contains the JSON metadata for the struct
// [V1CustomerCreditListResponseData]
type v1CustomerCreditListResponseDataJSON struct {
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
	HierarchyConfiguration  apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	RateType                apijson.Field
	SalesforceOpportunityID apijson.Field
	Specifiers              apijson.Field
	UniquenessKey           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseDataProduct struct {
	ID   string                                      `json:"id,required" format:"uuid"`
	Name string                                      `json:"name,required"`
	JSON v1CustomerCreditListResponseDataProductJSON `json:"-"`
}

// v1CustomerCreditListResponseDataProductJSON contains the JSON metadata for the
// struct [V1CustomerCreditListResponseDataProduct]
type v1CustomerCreditListResponseDataProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseDataProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseDataProductJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseDataType string

const (
	V1CustomerCreditListResponseDataTypeCredit V1CustomerCreditListResponseDataType = "CREDIT"
)

func (r V1CustomerCreditListResponseDataType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseDataTypeCredit:
		return true
	}
	return false
}

// The schedule that the customer will gain access to the credits.
type V1CustomerCreditListResponseDataAccessSchedule struct {
	ScheduleItems []V1CustomerCreditListResponseDataAccessScheduleScheduleItem `json:"schedule_items,required"`
	CreditType    V1CustomerCreditListResponseDataAccessScheduleCreditType     `json:"credit_type"`
	JSON          v1CustomerCreditListResponseDataAccessScheduleJSON           `json:"-"`
}

// v1CustomerCreditListResponseDataAccessScheduleJSON contains the JSON metadata
// for the struct [V1CustomerCreditListResponseDataAccessSchedule]
type v1CustomerCreditListResponseDataAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseDataAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseDataAccessScheduleJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseDataAccessScheduleScheduleItem struct {
	ID           string                                                         `json:"id,required" format:"uuid"`
	Amount       float64                                                        `json:"amount,required"`
	EndingBefore time.Time                                                      `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                                                      `json:"starting_at,required" format:"date-time"`
	JSON         v1CustomerCreditListResponseDataAccessScheduleScheduleItemJSON `json:"-"`
}

// v1CustomerCreditListResponseDataAccessScheduleScheduleItemJSON contains the JSON
// metadata for the struct
// [V1CustomerCreditListResponseDataAccessScheduleScheduleItem]
type v1CustomerCreditListResponseDataAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseDataAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseDataAccessScheduleScheduleItemJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseDataAccessScheduleCreditType struct {
	ID   string                                                       `json:"id,required" format:"uuid"`
	Name string                                                       `json:"name,required"`
	JSON v1CustomerCreditListResponseDataAccessScheduleCreditTypeJSON `json:"-"`
}

// v1CustomerCreditListResponseDataAccessScheduleCreditTypeJSON contains the JSON
// metadata for the struct
// [V1CustomerCreditListResponseDataAccessScheduleCreditType]
type v1CustomerCreditListResponseDataAccessScheduleCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseDataAccessScheduleCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseDataAccessScheduleCreditTypeJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseDataContract struct {
	ID   string                                       `json:"id,required" format:"uuid"`
	JSON v1CustomerCreditListResponseDataContractJSON `json:"-"`
}

// v1CustomerCreditListResponseDataContractJSON contains the JSON metadata for the
// struct [V1CustomerCreditListResponseDataContract]
type v1CustomerCreditListResponseDataContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseDataContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseDataContractJSON) RawJSON() string {
	return r.raw
}

// Optional configuration for credit hierarchy access control
type V1CustomerCreditListResponseDataHierarchyConfiguration struct {
	ChildAccess V1CustomerCreditListResponseDataHierarchyConfigurationChildAccess `json:"child_access,required"`
	JSON        v1CustomerCreditListResponseDataHierarchyConfigurationJSON        `json:"-"`
}

// v1CustomerCreditListResponseDataHierarchyConfigurationJSON contains the JSON
// metadata for the struct [V1CustomerCreditListResponseDataHierarchyConfiguration]
type v1CustomerCreditListResponseDataHierarchyConfigurationJSON struct {
	ChildAccess apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseDataHierarchyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseDataHierarchyConfigurationJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditListResponseDataHierarchyConfigurationChildAccess struct {
	Type V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	ContractIDs interface{}                                                           `json:"contract_ids"`
	JSON        v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessJSON `json:"-"`
	union       V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessUnion
}

// v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessJSON contains
// the JSON metadata for the struct
// [V1CustomerCreditListResponseDataHierarchyConfigurationChildAccess]
type v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessJSON struct {
	Type        apijson.Field
	ContractIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessJSON) RawJSON() string {
	return r.raw
}

func (r *V1CustomerCreditListResponseDataHierarchyConfigurationChildAccess) UnmarshalJSON(data []byte) (err error) {
	*r = V1CustomerCreditListResponseDataHierarchyConfigurationChildAccess{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessType],
// [V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessType],
// [V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObject].
func (r V1CustomerCreditListResponseDataHierarchyConfigurationChildAccess) AsUnion() V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessUnion {
	return r.union
}

// Union satisfied by
// [V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessType],
// [V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessType] or
// [V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObject].
type V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessUnion interface {
	implementsV1CustomerCreditListResponseDataHierarchyConfigurationChildAccess()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObject{}),
		},
	)
}

type V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessType struct {
	Type V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessTypeType `json:"type,required"`
	JSON v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessTypeJSON `json:"-"`
}

// v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessTypeJSON
// contains the JSON metadata for the struct
// [V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessType]
type v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessTypeJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessType) implementsV1CustomerCreditListResponseDataHierarchyConfigurationChildAccess() {
}

type V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessTypeType string

const (
	V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessTypeTypeAll V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessTypeType = "ALL"
)

func (r V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessTypeType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessTypeTypeAll:
		return true
	}
	return false
}

type V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObject struct {
	ContractIDs []string                                                                    `json:"contract_ids,required" format:"uuid"`
	Type        V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObjectType `json:"type,required"`
	JSON        v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObjectJSON `json:"-"`
}

// v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObjectJSON
// contains the JSON metadata for the struct
// [V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObject]
type v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObjectJSON struct {
	ContractIDs apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObjectJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObject) implementsV1CustomerCreditListResponseDataHierarchyConfigurationChildAccess() {
}

type V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObjectType string

const (
	V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObjectTypeContractIDs V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObjectType = "CONTRACT_IDS"
)

func (r V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObjectType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseDataHierarchyConfigurationChildAccessObjectTypeContractIDs:
		return true
	}
	return false
}

type V1CustomerCreditListResponseDataLedger struct {
	Amount     float64                                    `json:"amount,required"`
	Timestamp  time.Time                                  `json:"timestamp,required" format:"date-time"`
	Type       V1CustomerCreditListResponseDataLedgerType `json:"type,required"`
	ContractID string                                     `json:"contract_id" format:"uuid"`
	InvoiceID  string                                     `json:"invoice_id" format:"uuid"`
	Reason     string                                     `json:"reason"`
	SegmentID  string                                     `json:"segment_id" format:"uuid"`
	JSON       v1CustomerCreditListResponseDataLedgerJSON `json:"-"`
	union      V1CustomerCreditListResponseDataLedgerUnion
}

// v1CustomerCreditListResponseDataLedgerJSON contains the JSON metadata for the
// struct [V1CustomerCreditListResponseDataLedger]
type v1CustomerCreditListResponseDataLedgerJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	ContractID  apijson.Field
	InvoiceID   apijson.Field
	Reason      apijson.Field
	SegmentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v1CustomerCreditListResponseDataLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *V1CustomerCreditListResponseDataLedger) UnmarshalJSON(data []byte) (err error) {
	*r = V1CustomerCreditListResponseDataLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V1CustomerCreditListResponseDataLedgerUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V1CustomerCreditListResponseDataLedgerObject],
// [V1CustomerCreditListResponseDataLedgerObject],
// [V1CustomerCreditListResponseDataLedgerObject],
// [V1CustomerCreditListResponseDataLedgerObject],
// [V1CustomerCreditListResponseDataLedgerObject],
// [V1CustomerCreditListResponseDataLedgerObject],
// [V1CustomerCreditListResponseDataLedgerObject].
func (r V1CustomerCreditListResponseDataLedger) AsUnion() V1CustomerCreditListResponseDataLedgerUnion {
	return r.union
}

// Union satisfied by [V1CustomerCreditListResponseDataLedgerObject],
// [V1CustomerCreditListResponseDataLedgerObject],
// [V1CustomerCreditListResponseDataLedgerObject],
// [V1CustomerCreditListResponseDataLedgerObject],
// [V1CustomerCreditListResponseDataLedgerObject],
// [V1CustomerCreditListResponseDataLedgerObject] or
// [V1CustomerCreditListResponseDataLedgerObject].
type V1CustomerCreditListResponseDataLedgerUnion interface {
	implementsV1CustomerCreditListResponseDataLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1CustomerCreditListResponseDataLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseDataLedgerObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1CustomerCreditListResponseDataLedgerObject{}),
		},
	)
}

type V1CustomerCreditListResponseDataLedgerObject struct {
	Amount    float64                                          `json:"amount,required"`
	SegmentID string                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                        `json:"timestamp,required" format:"date-time"`
	Type      V1CustomerCreditListResponseDataLedgerObjectType `json:"type,required"`
	JSON      v1CustomerCreditListResponseDataLedgerObjectJSON `json:"-"`
}

// v1CustomerCreditListResponseDataLedgerObjectJSON contains the JSON metadata for
// the struct [V1CustomerCreditListResponseDataLedgerObject]
type v1CustomerCreditListResponseDataLedgerObjectJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseDataLedgerObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseDataLedgerObjectJSON) RawJSON() string {
	return r.raw
}

func (r V1CustomerCreditListResponseDataLedgerObject) implementsV1CustomerCreditListResponseDataLedger() {
}

type V1CustomerCreditListResponseDataLedgerObjectType string

const (
	V1CustomerCreditListResponseDataLedgerObjectTypeCreditSegmentStart V1CustomerCreditListResponseDataLedgerObjectType = "CREDIT_SEGMENT_START"
)

func (r V1CustomerCreditListResponseDataLedgerObjectType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseDataLedgerObjectTypeCreditSegmentStart:
		return true
	}
	return false
}

type V1CustomerCreditListResponseDataLedgerType string

const (
	V1CustomerCreditListResponseDataLedgerTypeCreditSegmentStart              V1CustomerCreditListResponseDataLedgerType = "CREDIT_SEGMENT_START"
	V1CustomerCreditListResponseDataLedgerTypeCreditAutomatedInvoiceDeduction V1CustomerCreditListResponseDataLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	V1CustomerCreditListResponseDataLedgerTypeCreditExpiration                V1CustomerCreditListResponseDataLedgerType = "CREDIT_EXPIRATION"
	V1CustomerCreditListResponseDataLedgerTypeCreditCanceled                  V1CustomerCreditListResponseDataLedgerType = "CREDIT_CANCELED"
	V1CustomerCreditListResponseDataLedgerTypeCreditCredited                  V1CustomerCreditListResponseDataLedgerType = "CREDIT_CREDITED"
	V1CustomerCreditListResponseDataLedgerTypeCreditManual                    V1CustomerCreditListResponseDataLedgerType = "CREDIT_MANUAL"
	V1CustomerCreditListResponseDataLedgerTypeCreditSeatBasedAdjustment       V1CustomerCreditListResponseDataLedgerType = "CREDIT_SEAT_BASED_ADJUSTMENT"
)

func (r V1CustomerCreditListResponseDataLedgerType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseDataLedgerTypeCreditSegmentStart, V1CustomerCreditListResponseDataLedgerTypeCreditAutomatedInvoiceDeduction, V1CustomerCreditListResponseDataLedgerTypeCreditExpiration, V1CustomerCreditListResponseDataLedgerTypeCreditCanceled, V1CustomerCreditListResponseDataLedgerTypeCreditCredited, V1CustomerCreditListResponseDataLedgerTypeCreditManual, V1CustomerCreditListResponseDataLedgerTypeCreditSeatBasedAdjustment:
		return true
	}
	return false
}

type V1CustomerCreditListResponseDataRateType string

const (
	V1CustomerCreditListResponseDataRateTypeCommitRate V1CustomerCreditListResponseDataRateType = "COMMIT_RATE"
	V1CustomerCreditListResponseDataRateTypeListRate   V1CustomerCreditListResponseDataRateType = "LIST_RATE"
)

func (r V1CustomerCreditListResponseDataRateType) IsKnown() bool {
	switch r {
	case V1CustomerCreditListResponseDataRateTypeCommitRate, V1CustomerCreditListResponseDataRateTypeListRate:
		return true
	}
	return false
}

type V1CustomerCreditListResponseDataSpecifier struct {
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	PricingGroupValues      map[string]string `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags []string                                      `json:"product_tags"`
	JSON        v1CustomerCreditListResponseDataSpecifierJSON `json:"-"`
}

// v1CustomerCreditListResponseDataSpecifierJSON contains the JSON metadata for the
// struct [V1CustomerCreditListResponseDataSpecifier]
type v1CustomerCreditListResponseDataSpecifierJSON struct {
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1CustomerCreditListResponseDataSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditListResponseDataSpecifierJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditUpdateEndDateResponse struct {
	Data V1CustomerCreditUpdateEndDateResponseData `json:"data,required"`
	JSON v1CustomerCreditUpdateEndDateResponseJSON `json:"-"`
}

// v1CustomerCreditUpdateEndDateResponseJSON contains the JSON metadata for the
// struct [V1CustomerCreditUpdateEndDateResponse]
type v1CustomerCreditUpdateEndDateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditUpdateEndDateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditUpdateEndDateResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditUpdateEndDateResponseData struct {
	ID   string                                        `json:"id,required" format:"uuid"`
	JSON v1CustomerCreditUpdateEndDateResponseDataJSON `json:"-"`
}

// v1CustomerCreditUpdateEndDateResponseDataJSON contains the JSON metadata for the
// struct [V1CustomerCreditUpdateEndDateResponseData]
type v1CustomerCreditUpdateEndDateResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerCreditUpdateEndDateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerCreditUpdateEndDateResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerCreditNewParams struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule param.Field[V1CustomerCreditNewParamsAccessSchedule] `json:"access_schedule,required"`
	CustomerID     param.Field[string]                                  `json:"customer_id,required" format:"uuid"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority  param.Field[float64] `json:"priority,required"`
	ProductID param.Field[string]  `json:"product_id,required" format:"uuid"`
	// Which contract the credit applies to. If not provided, the credit applies to all
	// contracts.
	ApplicableContractIDs param.Field[[]string] `json:"applicable_contract_ids"`
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
	NetsuiteSalesOrderID param.Field[string]                            `json:"netsuite_sales_order_id"`
	RateType             param.Field[V1CustomerCreditNewParamsRateType] `json:"rate_type"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string] `json:"salesforce_opportunity_id"`
	// List of filters that determine what kind of customer usage draws down a commit
	// or credit. A customer's usage needs to meet the condition of at least one of the
	// specifiers to contribute to a commit's or credit's drawdown. This field cannot
	// be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers param.Field[[]V1CustomerCreditNewParamsSpecifier] `json:"specifiers"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey param.Field[string] `json:"uniqueness_key"`
}

func (r V1CustomerCreditNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the credit to the customer.
type V1CustomerCreditNewParamsAccessSchedule struct {
	ScheduleItems param.Field[[]V1CustomerCreditNewParamsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r V1CustomerCreditNewParamsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCreditNewParamsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V1CustomerCreditNewParamsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCreditNewParamsRateType string

const (
	V1CustomerCreditNewParamsRateTypeCommitRate V1CustomerCreditNewParamsRateType = "COMMIT_RATE"
	V1CustomerCreditNewParamsRateTypeListRate   V1CustomerCreditNewParamsRateType = "LIST_RATE"
)

func (r V1CustomerCreditNewParamsRateType) IsKnown() bool {
	switch r {
	case V1CustomerCreditNewParamsRateTypeCommitRate, V1CustomerCreditNewParamsRateTypeListRate:
		return true
	}
	return false
}

type V1CustomerCreditNewParamsSpecifier struct {
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the specifier will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V1CustomerCreditNewParamsSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCreditListParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Return only credits that have access schedules that "cover" the provided date
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	CreditID     param.Field[string]    `json:"credit_id" format:"uuid"`
	// Include only credits that have any access before the provided date (exclusive)
	EffectiveBefore param.Field[time.Time] `json:"effective_before" format:"date-time"`
	// Include archived credits and credits from archived contracts.
	IncludeArchived param.Field[bool] `json:"include_archived"`
	// Include the balance in the response. Setting this flag may cause the query to be
	// slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include credits on the contract level.
	IncludeContractCredits param.Field[bool] `json:"include_contract_credits"`
	// Include credit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
	// The maximum number of commits to return. Defaults to 25.
	Limit param.Field[int64] `json:"limit"`
	// The next page token from a previous response.
	NextPage param.Field[string] `json:"next_page"`
	// Include only credits that have any access on or after the provided date
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V1CustomerCreditListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1CustomerCreditUpdateEndDateParams struct {
	// RFC 3339 timestamp indicating when access to the credit will end and it will no
	// longer be possible to draw it down (exclusive).
	AccessEndingBefore param.Field[time.Time] `json:"access_ending_before,required" format:"date-time"`
	// ID of the commit to update
	CreditID param.Field[string] `json:"credit_id,required" format:"uuid"`
	// ID of the customer whose credit is to be updated
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r V1CustomerCreditUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
