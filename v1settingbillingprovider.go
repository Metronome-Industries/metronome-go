// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"

	"github.com/Metronome-Industries/metronome-go/v2/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v2/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/packages/param"
	"github.com/Metronome-Industries/metronome-go/v2/packages/respjson"
)

// V1SettingBillingProviderService contains methods and other services that help
// with interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1SettingBillingProviderService] method instead.
type V1SettingBillingProviderService struct {
	Options []option.RequestOption
}

// NewV1SettingBillingProviderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1SettingBillingProviderService(opts ...option.RequestOption) (r V1SettingBillingProviderService) {
	r = V1SettingBillingProviderService{}
	r.Options = opts
	return
}

// Set up account-level configuration for a billing provider. Once configured,
// individual contracts across customers can be mapped to this configuration using
// the returned delivery_method_id.
func (r *V1SettingBillingProviderService) New(ctx context.Context, body V1SettingBillingProviderNewParams, opts ...option.RequestOption) (res *V1SettingBillingProviderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/setUpBillingProvider"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all configured billing providers and their delivery method configurations
// for your account. Returns provider details, delivery method IDs, and
// configuration settings needed for mapping individual customer contracts to
// billing integrations.
func (r *V1SettingBillingProviderService) List(ctx context.Context, body V1SettingBillingProviderListParams, opts ...option.RequestOption) (res *V1SettingBillingProviderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/listConfiguredBillingProviders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1SettingBillingProviderNewResponse struct {
	Data V1SettingBillingProviderNewResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SettingBillingProviderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SettingBillingProviderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SettingBillingProviderNewResponseData struct {
	DeliveryMethodID string `json:"delivery_method_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeliveryMethodID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SettingBillingProviderNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SettingBillingProviderNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SettingBillingProviderListResponse struct {
	Data     []V1SettingBillingProviderListResponseData `json:"data,required"`
	NextPage string                                     `json:"next_page,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextPage    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SettingBillingProviderListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SettingBillingProviderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SettingBillingProviderListResponseData struct {
	// The billing provider set for this configuration.
	//
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace", "metronome".
	BillingProvider string `json:"billing_provider,required"`
	// The method to use for delivering invoices to this customer.
	//
	// Any of "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".
	DeliveryMethod string `json:"delivery_method,required"`
	// Configuration for the delivery method. The structure of this object is specific
	// to the delivery method. Some configuration may be omitted for security reasons.
	DeliveryMethodConfiguration map[string]any `json:"delivery_method_configuration,required"`
	// ID of the delivery method to use for this customer.
	DeliveryMethodID string `json:"delivery_method_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingProvider             respjson.Field
		DeliveryMethod              respjson.Field
		DeliveryMethodConfiguration respjson.Field
		DeliveryMethodID            respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SettingBillingProviderListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SettingBillingProviderListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SettingBillingProviderNewParams struct {
	// The billing provider set for this configuration.
	//
	// Any of "aws_marketplace", "azure_marketplace", "gcp_marketplace".
	BillingProvider V1SettingBillingProviderNewParamsBillingProvider `json:"billing_provider,omitzero,required"`
	// Account-level configuration for the billing provider. The structure of this
	// object is specific to the billing provider and delivery provider combination.
	// See examples below.
	Configuration map[string]any `json:"configuration,omitzero,required"`
	// The method to use for delivering invoices for this configuration.
	//
	// Any of "direct_to_billing_provider", "aws_sqs", "aws_sns".
	DeliveryMethod V1SettingBillingProviderNewParamsDeliveryMethod `json:"delivery_method,omitzero,required"`
	paramObj
}

func (r V1SettingBillingProviderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SettingBillingProviderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SettingBillingProviderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The billing provider set for this configuration.
type V1SettingBillingProviderNewParamsBillingProvider string

const (
	V1SettingBillingProviderNewParamsBillingProviderAwsMarketplace   V1SettingBillingProviderNewParamsBillingProvider = "aws_marketplace"
	V1SettingBillingProviderNewParamsBillingProviderAzureMarketplace V1SettingBillingProviderNewParamsBillingProvider = "azure_marketplace"
	V1SettingBillingProviderNewParamsBillingProviderGcpMarketplace   V1SettingBillingProviderNewParamsBillingProvider = "gcp_marketplace"
)

// The method to use for delivering invoices for this configuration.
type V1SettingBillingProviderNewParamsDeliveryMethod string

const (
	V1SettingBillingProviderNewParamsDeliveryMethodDirectToBillingProvider V1SettingBillingProviderNewParamsDeliveryMethod = "direct_to_billing_provider"
	V1SettingBillingProviderNewParamsDeliveryMethodAwsSqs                  V1SettingBillingProviderNewParamsDeliveryMethod = "aws_sqs"
	V1SettingBillingProviderNewParamsDeliveryMethodAwsSns                  V1SettingBillingProviderNewParamsDeliveryMethod = "aws_sns"
)

type V1SettingBillingProviderListParams struct {
	// The cursor to the next page of results
	NextPage param.Opt[string] `json:"next_page,omitzero" format:"uuid"`
	paramObj
}

func (r V1SettingBillingProviderListParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SettingBillingProviderListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SettingBillingProviderListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
