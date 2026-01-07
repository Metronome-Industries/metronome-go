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
	"github.com/Metronome-Industries/metronome-go/v2/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/v2/packages/param"
	"github.com/Metronome-Industries/metronome-go/v2/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/v2/shared"
)

// V1PaymentService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PaymentService] method instead.
type V1PaymentService struct {
	Options []option.RequestOption
}

// NewV1PaymentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1PaymentService(opts ...option.RequestOption) (r V1PaymentService) {
	r = V1PaymentService{}
	r.Options = opts
	return
}

// Fetch all payment attempts for the given invoice.
func (r *V1PaymentService) List(ctx context.Context, body V1PaymentListParams, opts ...option.RequestOption) (res *pagination.BodyCursorPage[Payment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/payments/list"
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

// Fetch all payment attempts for the given invoice.
func (r *V1PaymentService) ListAutoPaging(ctx context.Context, body V1PaymentListParams, opts ...option.RequestOption) *pagination.BodyCursorPageAutoPager[Payment] {
	return pagination.NewBodyCursorPageAutoPager(r.List(ctx, body, opts...))
}

// Trigger a new attempt by canceling any existing attempts for this invoice and
// creating a new Payment. This will trigger another attempt to charge the
// Customer's configured Payment Gateway. Payment can only be attempted if all of
// the following are true:
//
//   - The Metronome Invoice is finalized
//   - PLG Invoicing is configured for the Customer
//   - You cannot attempt payments for invoices that have already been `paid` or
//     `voided`.
//
// Attempting to payment on an ineligible Invoice or Customer will result in a
// `400` response.
func (r *V1PaymentService) Attempt(ctx context.Context, body V1PaymentAttemptParams, opts ...option.RequestOption) (res *V1PaymentAttemptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/payments/attempt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Cancel an existing payment attempt for an invoice.
func (r *V1PaymentService) Cancel(ctx context.Context, body V1PaymentCancelParams, opts ...option.RequestOption) (res *V1PaymentCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/payments/cancel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Payment struct {
	ID                    string                        `json:"id,required" format:"uuid"`
	Amount                float64                       `json:"amount"`
	AmountPaid            float64                       `json:"amount_paid"`
	ContractID            string                        `json:"contract_id" format:"uuid"`
	CreatedAt             time.Time                     `json:"created_at" format:"date-time"`
	CustomerID            string                        `json:"customer_id" format:"uuid"`
	ErrorMessage          string                        `json:"error_message"`
	FiatCreditType        shared.CreditTypeData         `json:"fiat_credit_type"`
	InvoiceID             string                        `json:"invoice_id" format:"uuid"`
	PaymentGateway        PaymentPaymentGateway         `json:"payment_gateway"`
	RevenueSystemPayments []PaymentRevenueSystemPayment `json:"revenue_system_payments"`
	// Any of "pending", "requires_intervention", "paid", "canceled".
	Status    PaymentStatus `json:"status"`
	UpdatedAt time.Time     `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Amount                respjson.Field
		AmountPaid            respjson.Field
		ContractID            respjson.Field
		CreatedAt             respjson.Field
		CustomerID            respjson.Field
		ErrorMessage          respjson.Field
		FiatCreditType        respjson.Field
		InvoiceID             respjson.Field
		PaymentGateway        respjson.Field
		RevenueSystemPayments respjson.Field
		Status                respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Payment) RawJSON() string { return r.JSON.raw }
func (r *Payment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentPaymentGateway struct {
	Stripe PaymentPaymentGatewayStripe `json:"stripe,required"`
	// Any of "stripe".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Stripe      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentPaymentGateway) RawJSON() string { return r.JSON.raw }
func (r *PaymentPaymentGateway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentPaymentGatewayStripe struct {
	PaymentIntentID string                           `json:"payment_intent_id,required"`
	Error           PaymentPaymentGatewayStripeError `json:"error"`
	PaymentMethodID string                           `json:"payment_method_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentIntentID respjson.Field
		Error           respjson.Field
		PaymentMethodID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentPaymentGatewayStripe) RawJSON() string { return r.JSON.raw }
func (r *PaymentPaymentGatewayStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentPaymentGatewayStripeError struct {
	Code        string `json:"code"`
	DeclineCode string `json:"decline_code"`
	Type        string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		DeclineCode respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentPaymentGatewayStripeError) RawJSON() string { return r.JSON.raw }
func (r *PaymentPaymentGatewayStripeError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentRevenueSystemPayment struct {
	RevenueSystemProvider string `json:"revenue_system_provider,required"`
	SyncStatus            string `json:"sync_status,required"`
	// The error message from the revenue system, if available.
	ErrorMessage                   string `json:"error_message"`
	RevenueSystemExternalPaymentID string `json:"revenue_system_external_payment_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RevenueSystemProvider          respjson.Field
		SyncStatus                     respjson.Field
		ErrorMessage                   respjson.Field
		RevenueSystemExternalPaymentID respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentRevenueSystemPayment) RawJSON() string { return r.JSON.raw }
func (r *PaymentRevenueSystemPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentStatus string

const (
	PaymentStatusPending              PaymentStatus = "pending"
	PaymentStatusRequiresIntervention PaymentStatus = "requires_intervention"
	PaymentStatusPaid                 PaymentStatus = "paid"
	PaymentStatusCanceled             PaymentStatus = "canceled"
)

type V1PaymentAttemptResponse struct {
	Data Payment `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PaymentAttemptResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PaymentAttemptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PaymentCancelResponse struct {
	Data Payment `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PaymentCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PaymentCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PaymentListParams struct {
	CustomerID string `json:"customer_id,required" format:"uuid"`
	InvoiceID  string `json:"invoice_id,required" format:"uuid"`
	// The maximum number of payments to return. Defaults to 25.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// The next page token from a previous response.
	NextPage param.Opt[string] `json:"next_page,omitzero"`
	Statuses []PaymentStatus   `json:"statuses,omitzero"`
	paramObj
}

func (r V1PaymentListParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PaymentListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PaymentListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PaymentAttemptParams struct {
	CustomerID string `json:"customer_id,required" format:"uuid"`
	InvoiceID  string `json:"invoice_id,required" format:"uuid"`
	paramObj
}

func (r V1PaymentAttemptParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PaymentAttemptParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PaymentAttemptParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PaymentCancelParams struct {
	CustomerID string `json:"customer_id,required" format:"uuid"`
	InvoiceID  string `json:"invoice_id,required" format:"uuid"`
	paramObj
}

func (r V1PaymentCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PaymentCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PaymentCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
