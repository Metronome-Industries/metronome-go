// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// PaymentService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	Options []option.RequestOption
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r PaymentService) {
	r = PaymentService{}
	r.Options = opts
	return
}

// Fetch all payment attempts for the given invoice.
func (r *PaymentService) List(ctx context.Context, body PaymentListParams, opts ...option.RequestOption) (res *pagination.BodyCursorPage[PaymentListResponse], err error) {
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
func (r *PaymentService) ListAutoPaging(ctx context.Context, body PaymentListParams, opts ...option.RequestOption) *pagination.BodyCursorPageAutoPager[PaymentListResponse] {
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
func (r *PaymentService) AttemptPayment(ctx context.Context, body PaymentAttemptPaymentParams, opts ...option.RequestOption) (res *PaymentAttemptPaymentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/payments/attempt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PaymentListResponse struct {
	ID             string                            `json:"id,required" format:"uuid"`
	Amount         float64                           `json:"amount"`
	AmountPaid     float64                           `json:"amount_paid"`
	ContractID     string                            `json:"contract_id" format:"uuid"`
	CreatedAt      time.Time                         `json:"created_at" format:"date-time"`
	CustomerID     string                            `json:"customer_id" format:"uuid"`
	ErrorMessage   string                            `json:"error_message"`
	FiatCreditType shared.CreditTypeData             `json:"fiat_credit_type"`
	InvoiceID      string                            `json:"invoice_id" format:"uuid"`
	PaymentGateway PaymentListResponsePaymentGateway `json:"payment_gateway"`
	// Any of "pending", "requires_intervention", "paid", "canceled".
	Status    PaymentListResponseStatus `json:"status"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Amount         respjson.Field
		AmountPaid     respjson.Field
		ContractID     respjson.Field
		CreatedAt      respjson.Field
		CustomerID     respjson.Field
		ErrorMessage   respjson.Field
		FiatCreditType respjson.Field
		InvoiceID      respjson.Field
		PaymentGateway respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentListResponsePaymentGateway struct {
	Stripe PaymentListResponsePaymentGatewayStripe `json:"stripe,required"`
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
func (r PaymentListResponsePaymentGateway) RawJSON() string { return r.JSON.raw }
func (r *PaymentListResponsePaymentGateway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentListResponsePaymentGatewayStripe struct {
	PaymentIntentID string                                       `json:"payment_intent_id,required"`
	Error           PaymentListResponsePaymentGatewayStripeError `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentIntentID respjson.Field
		Error           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentListResponsePaymentGatewayStripe) RawJSON() string { return r.JSON.raw }
func (r *PaymentListResponsePaymentGatewayStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentListResponsePaymentGatewayStripeError struct {
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
func (r PaymentListResponsePaymentGatewayStripeError) RawJSON() string { return r.JSON.raw }
func (r *PaymentListResponsePaymentGatewayStripeError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentListResponseStatus string

const (
	PaymentListResponseStatusPending              PaymentListResponseStatus = "pending"
	PaymentListResponseStatusRequiresIntervention PaymentListResponseStatus = "requires_intervention"
	PaymentListResponseStatusPaid                 PaymentListResponseStatus = "paid"
	PaymentListResponseStatusCanceled             PaymentListResponseStatus = "canceled"
)

type PaymentAttemptPaymentResponse struct {
	Data PaymentAttemptPaymentResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentAttemptPaymentResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentAttemptPaymentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentAttemptPaymentResponseData struct {
	ID             string                                          `json:"id,required" format:"uuid"`
	Amount         float64                                         `json:"amount"`
	AmountPaid     float64                                         `json:"amount_paid"`
	ContractID     string                                          `json:"contract_id" format:"uuid"`
	CreatedAt      time.Time                                       `json:"created_at" format:"date-time"`
	CustomerID     string                                          `json:"customer_id" format:"uuid"`
	ErrorMessage   string                                          `json:"error_message"`
	FiatCreditType shared.CreditTypeData                           `json:"fiat_credit_type"`
	InvoiceID      string                                          `json:"invoice_id" format:"uuid"`
	PaymentGateway PaymentAttemptPaymentResponseDataPaymentGateway `json:"payment_gateway"`
	// Any of "pending", "requires_intervention", "paid", "canceled".
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Amount         respjson.Field
		AmountPaid     respjson.Field
		ContractID     respjson.Field
		CreatedAt      respjson.Field
		CustomerID     respjson.Field
		ErrorMessage   respjson.Field
		FiatCreditType respjson.Field
		InvoiceID      respjson.Field
		PaymentGateway respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentAttemptPaymentResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaymentAttemptPaymentResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentAttemptPaymentResponseDataPaymentGateway struct {
	Stripe PaymentAttemptPaymentResponseDataPaymentGatewayStripe `json:"stripe,required"`
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
func (r PaymentAttemptPaymentResponseDataPaymentGateway) RawJSON() string { return r.JSON.raw }
func (r *PaymentAttemptPaymentResponseDataPaymentGateway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentAttemptPaymentResponseDataPaymentGatewayStripe struct {
	PaymentIntentID string                                                     `json:"payment_intent_id,required"`
	Error           PaymentAttemptPaymentResponseDataPaymentGatewayStripeError `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentIntentID respjson.Field
		Error           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentAttemptPaymentResponseDataPaymentGatewayStripe) RawJSON() string { return r.JSON.raw }
func (r *PaymentAttemptPaymentResponseDataPaymentGatewayStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentAttemptPaymentResponseDataPaymentGatewayStripeError struct {
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
func (r PaymentAttemptPaymentResponseDataPaymentGatewayStripeError) RawJSON() string {
	return r.JSON.raw
}
func (r *PaymentAttemptPaymentResponseDataPaymentGatewayStripeError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentListParams struct {
	CustomerID string `json:"customer_id,required" format:"uuid"`
	InvoiceID  string `json:"invoice_id,required" format:"uuid"`
	// The maximum number of payments to return. Defaults to 25.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// The next page token from a previous response.
	NextPage param.Opt[string] `json:"next_page,omitzero"`
	// Any of "pending", "requires_intervention", "paid", "canceled".
	Statuses []string `json:"statuses,omitzero"`
	paramObj
}

func (r PaymentListParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentAttemptPaymentParams struct {
	CustomerID string `json:"customer_id,required" format:"uuid"`
	InvoiceID  string `json:"invoice_id,required" format:"uuid"`
	paramObj
}

func (r PaymentAttemptPaymentParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentAttemptPaymentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentAttemptPaymentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
