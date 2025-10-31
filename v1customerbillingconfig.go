// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/Metronome-Industries/metronome-go/v2/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v2/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v2/option"
	"github.com/Metronome-Industries/metronome-go/v2/packages/param"
	"github.com/Metronome-Industries/metronome-go/v2/packages/respjson"
)

// V1CustomerBillingConfigService contains methods and other services that help
// with interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerBillingConfigService] method instead.
type V1CustomerBillingConfigService struct {
	Options []option.RequestOption
}

// NewV1CustomerBillingConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerBillingConfigService(opts ...option.RequestOption) (r V1CustomerBillingConfigService) {
	r = V1CustomerBillingConfigService{}
	r.Options = opts
	return
}

// Set the billing configuration for a given customer. This is a Plans (deprecated)
// endpoint. New clients should implement using Contracts.
func (r *V1CustomerBillingConfigService) New(ctx context.Context, params V1CustomerBillingConfigNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/billing-config/%v", params.CustomerID, params.BillingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Fetch the billing configuration for the given customer. This is a Plans
// (deprecated) endpoint. New clients should implement using Contracts.
func (r *V1CustomerBillingConfigService) Get(ctx context.Context, query V1CustomerBillingConfigGetParams, opts ...option.RequestOption) (res *V1CustomerBillingConfigGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/billing-config/%v", query.CustomerID, query.BillingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the billing configuration for a given customer. Note: this is unsupported
// for Azure and AWS Marketplace customers. This is a Plans (deprecated) endpoint.
// New clients should implement using Contracts.
func (r *V1CustomerBillingConfigService) Delete(ctx context.Context, body V1CustomerBillingConfigDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/billing-config/%v", body.CustomerID, body.BillingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type V1CustomerBillingConfigGetResponse struct {
	Data V1CustomerBillingConfigGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerBillingConfigGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerBillingConfigGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerBillingConfigGetResponseData struct {
	// Contract expiration date for the customer. The expected format is RFC 3339 and
	// can be retrieved from
	// [AWS's GetEntitlements API](https://docs.aws.amazon.com/marketplaceentitlement/latest/APIReference/API_GetEntitlements.html).
	AwsExpirationDate time.Time `json:"aws_expiration_date" format:"date-time"`
	// True if the aws_product_code is a SAAS subscription product, false otherwise.
	AwsIsSubscriptionProduct bool   `json:"aws_is_subscription_product"`
	AwsProductCode           string `json:"aws_product_code"`
	// Any of "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2",
	// "ap-northeast-3", "ap-south-1", "ap-southeast-1", "ap-southeast-2",
	// "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1",
	// "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "me-south-1", "sa-east-1",
	// "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1",
	// "us-west-2".
	AwsRegion string `json:"aws_region"`
	// Subscription term start/end date for the customer. The expected format is RFC
	// 3339 and can be retrieved from
	// [Azure's Get Subscription API](https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-subscription).
	AzureExpirationDate time.Time `json:"azure_expiration_date" format:"date-time"`
	AzurePlanID         string    `json:"azure_plan_id" format:"uuid"`
	// Subscription term start/end date for the customer. The expected format is RFC
	// 3339 and can be retrieved from
	// [Azure's Get Subscription API](https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-subscription).
	AzureStartDate time.Time `json:"azure_start_date" format:"date-time"`
	// Any of "Subscribed", "Unsubscribed", "Suspended", "PendingFulfillmentStart".
	AzureSubscriptionStatus   string `json:"azure_subscription_status"`
	BillingProviderCustomerID string `json:"billing_provider_customer_id"`
	// The collection method for the customer's invoices. NOTE:
	// `auto_charge_payment_intent` and `manually_charge_payment_intent` are in beta.
	//
	// Any of "charge_automatically", "send_invoice", "auto_charge_payment_intent",
	// "manually_charge_payment_intent".
	StripeCollectionMethod string `json:"stripe_collection_method"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwsExpirationDate         respjson.Field
		AwsIsSubscriptionProduct  respjson.Field
		AwsProductCode            respjson.Field
		AwsRegion                 respjson.Field
		AzureExpirationDate       respjson.Field
		AzurePlanID               respjson.Field
		AzureStartDate            respjson.Field
		AzureSubscriptionStatus   respjson.Field
		BillingProviderCustomerID respjson.Field
		StripeCollectionMethod    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerBillingConfigGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerBillingConfigGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerBillingConfigNewParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace", "metronome".
	BillingProviderType V1CustomerBillingConfigNewParamsBillingProviderType `path:"billing_provider_type,omitzero,required" json:"-"`
	// The customer ID in the billing provider's system. For Azure, this is the
	// subscription ID.
	BillingProviderCustomerID string            `json:"billing_provider_customer_id,required"`
	AwsProductCode            param.Opt[string] `json:"aws_product_code,omitzero"`
	// Any of "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2",
	// "ap-northeast-3", "ap-south-1", "ap-southeast-1", "ap-southeast-2",
	// "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1",
	// "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "me-south-1", "sa-east-1",
	// "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1",
	// "us-west-2".
	AwsRegion V1CustomerBillingConfigNewParamsAwsRegion `json:"aws_region,omitzero"`
	// The collection method for the customer's invoices. NOTE:
	// `auto_charge_payment_intent` and `manually_charge_payment_intent` are in beta.
	//
	// Any of "charge_automatically", "send_invoice", "auto_charge_payment_intent",
	// "manually_charge_payment_intent".
	StripeCollectionMethod V1CustomerBillingConfigNewParamsStripeCollectionMethod `json:"stripe_collection_method,omitzero"`
	paramObj
}

func (r V1CustomerBillingConfigNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerBillingConfigNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerBillingConfigNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerBillingConfigNewParamsBillingProviderType string

const (
	V1CustomerBillingConfigNewParamsBillingProviderTypeAwsMarketplace   V1CustomerBillingConfigNewParamsBillingProviderType = "aws_marketplace"
	V1CustomerBillingConfigNewParamsBillingProviderTypeStripe           V1CustomerBillingConfigNewParamsBillingProviderType = "stripe"
	V1CustomerBillingConfigNewParamsBillingProviderTypeNetsuite         V1CustomerBillingConfigNewParamsBillingProviderType = "netsuite"
	V1CustomerBillingConfigNewParamsBillingProviderTypeCustom           V1CustomerBillingConfigNewParamsBillingProviderType = "custom"
	V1CustomerBillingConfigNewParamsBillingProviderTypeAzureMarketplace V1CustomerBillingConfigNewParamsBillingProviderType = "azure_marketplace"
	V1CustomerBillingConfigNewParamsBillingProviderTypeQuickbooksOnline V1CustomerBillingConfigNewParamsBillingProviderType = "quickbooks_online"
	V1CustomerBillingConfigNewParamsBillingProviderTypeWorkday          V1CustomerBillingConfigNewParamsBillingProviderType = "workday"
	V1CustomerBillingConfigNewParamsBillingProviderTypeGcpMarketplace   V1CustomerBillingConfigNewParamsBillingProviderType = "gcp_marketplace"
	V1CustomerBillingConfigNewParamsBillingProviderTypeMetronome        V1CustomerBillingConfigNewParamsBillingProviderType = "metronome"
)

type V1CustomerBillingConfigNewParamsAwsRegion string

const (
	V1CustomerBillingConfigNewParamsAwsRegionAfSouth1     V1CustomerBillingConfigNewParamsAwsRegion = "af-south-1"
	V1CustomerBillingConfigNewParamsAwsRegionApEast1      V1CustomerBillingConfigNewParamsAwsRegion = "ap-east-1"
	V1CustomerBillingConfigNewParamsAwsRegionApNortheast1 V1CustomerBillingConfigNewParamsAwsRegion = "ap-northeast-1"
	V1CustomerBillingConfigNewParamsAwsRegionApNortheast2 V1CustomerBillingConfigNewParamsAwsRegion = "ap-northeast-2"
	V1CustomerBillingConfigNewParamsAwsRegionApNortheast3 V1CustomerBillingConfigNewParamsAwsRegion = "ap-northeast-3"
	V1CustomerBillingConfigNewParamsAwsRegionApSouth1     V1CustomerBillingConfigNewParamsAwsRegion = "ap-south-1"
	V1CustomerBillingConfigNewParamsAwsRegionApSoutheast1 V1CustomerBillingConfigNewParamsAwsRegion = "ap-southeast-1"
	V1CustomerBillingConfigNewParamsAwsRegionApSoutheast2 V1CustomerBillingConfigNewParamsAwsRegion = "ap-southeast-2"
	V1CustomerBillingConfigNewParamsAwsRegionCaCentral1   V1CustomerBillingConfigNewParamsAwsRegion = "ca-central-1"
	V1CustomerBillingConfigNewParamsAwsRegionCnNorth1     V1CustomerBillingConfigNewParamsAwsRegion = "cn-north-1"
	V1CustomerBillingConfigNewParamsAwsRegionCnNorthwest1 V1CustomerBillingConfigNewParamsAwsRegion = "cn-northwest-1"
	V1CustomerBillingConfigNewParamsAwsRegionEuCentral1   V1CustomerBillingConfigNewParamsAwsRegion = "eu-central-1"
	V1CustomerBillingConfigNewParamsAwsRegionEuNorth1     V1CustomerBillingConfigNewParamsAwsRegion = "eu-north-1"
	V1CustomerBillingConfigNewParamsAwsRegionEuSouth1     V1CustomerBillingConfigNewParamsAwsRegion = "eu-south-1"
	V1CustomerBillingConfigNewParamsAwsRegionEuWest1      V1CustomerBillingConfigNewParamsAwsRegion = "eu-west-1"
	V1CustomerBillingConfigNewParamsAwsRegionEuWest2      V1CustomerBillingConfigNewParamsAwsRegion = "eu-west-2"
	V1CustomerBillingConfigNewParamsAwsRegionEuWest3      V1CustomerBillingConfigNewParamsAwsRegion = "eu-west-3"
	V1CustomerBillingConfigNewParamsAwsRegionMeSouth1     V1CustomerBillingConfigNewParamsAwsRegion = "me-south-1"
	V1CustomerBillingConfigNewParamsAwsRegionSaEast1      V1CustomerBillingConfigNewParamsAwsRegion = "sa-east-1"
	V1CustomerBillingConfigNewParamsAwsRegionUsEast1      V1CustomerBillingConfigNewParamsAwsRegion = "us-east-1"
	V1CustomerBillingConfigNewParamsAwsRegionUsEast2      V1CustomerBillingConfigNewParamsAwsRegion = "us-east-2"
	V1CustomerBillingConfigNewParamsAwsRegionUsGovEast1   V1CustomerBillingConfigNewParamsAwsRegion = "us-gov-east-1"
	V1CustomerBillingConfigNewParamsAwsRegionUsGovWest1   V1CustomerBillingConfigNewParamsAwsRegion = "us-gov-west-1"
	V1CustomerBillingConfigNewParamsAwsRegionUsWest1      V1CustomerBillingConfigNewParamsAwsRegion = "us-west-1"
	V1CustomerBillingConfigNewParamsAwsRegionUsWest2      V1CustomerBillingConfigNewParamsAwsRegion = "us-west-2"
)

// The collection method for the customer's invoices. NOTE:
// `auto_charge_payment_intent` and `manually_charge_payment_intent` are in beta.
type V1CustomerBillingConfigNewParamsStripeCollectionMethod string

const (
	V1CustomerBillingConfigNewParamsStripeCollectionMethodChargeAutomatically         V1CustomerBillingConfigNewParamsStripeCollectionMethod = "charge_automatically"
	V1CustomerBillingConfigNewParamsStripeCollectionMethodSendInvoice                 V1CustomerBillingConfigNewParamsStripeCollectionMethod = "send_invoice"
	V1CustomerBillingConfigNewParamsStripeCollectionMethodAutoChargePaymentIntent     V1CustomerBillingConfigNewParamsStripeCollectionMethod = "auto_charge_payment_intent"
	V1CustomerBillingConfigNewParamsStripeCollectionMethodManuallyChargePaymentIntent V1CustomerBillingConfigNewParamsStripeCollectionMethod = "manually_charge_payment_intent"
)

type V1CustomerBillingConfigGetParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace", "metronome".
	BillingProviderType V1CustomerBillingConfigGetParamsBillingProviderType `path:"billing_provider_type,omitzero,required" json:"-"`
	paramObj
}

type V1CustomerBillingConfigGetParamsBillingProviderType string

const (
	V1CustomerBillingConfigGetParamsBillingProviderTypeAwsMarketplace   V1CustomerBillingConfigGetParamsBillingProviderType = "aws_marketplace"
	V1CustomerBillingConfigGetParamsBillingProviderTypeStripe           V1CustomerBillingConfigGetParamsBillingProviderType = "stripe"
	V1CustomerBillingConfigGetParamsBillingProviderTypeNetsuite         V1CustomerBillingConfigGetParamsBillingProviderType = "netsuite"
	V1CustomerBillingConfigGetParamsBillingProviderTypeCustom           V1CustomerBillingConfigGetParamsBillingProviderType = "custom"
	V1CustomerBillingConfigGetParamsBillingProviderTypeAzureMarketplace V1CustomerBillingConfigGetParamsBillingProviderType = "azure_marketplace"
	V1CustomerBillingConfigGetParamsBillingProviderTypeQuickbooksOnline V1CustomerBillingConfigGetParamsBillingProviderType = "quickbooks_online"
	V1CustomerBillingConfigGetParamsBillingProviderTypeWorkday          V1CustomerBillingConfigGetParamsBillingProviderType = "workday"
	V1CustomerBillingConfigGetParamsBillingProviderTypeGcpMarketplace   V1CustomerBillingConfigGetParamsBillingProviderType = "gcp_marketplace"
	V1CustomerBillingConfigGetParamsBillingProviderTypeMetronome        V1CustomerBillingConfigGetParamsBillingProviderType = "metronome"
)

type V1CustomerBillingConfigDeleteParams struct {
	CustomerID string `path:"customer_id,required" format:"uuid" json:"-"`
	// Any of "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace",
	// "quickbooks_online", "workday", "gcp_marketplace", "metronome".
	BillingProviderType V1CustomerBillingConfigDeleteParamsBillingProviderType `path:"billing_provider_type,omitzero,required" json:"-"`
	paramObj
}

type V1CustomerBillingConfigDeleteParamsBillingProviderType string

const (
	V1CustomerBillingConfigDeleteParamsBillingProviderTypeAwsMarketplace   V1CustomerBillingConfigDeleteParamsBillingProviderType = "aws_marketplace"
	V1CustomerBillingConfigDeleteParamsBillingProviderTypeStripe           V1CustomerBillingConfigDeleteParamsBillingProviderType = "stripe"
	V1CustomerBillingConfigDeleteParamsBillingProviderTypeNetsuite         V1CustomerBillingConfigDeleteParamsBillingProviderType = "netsuite"
	V1CustomerBillingConfigDeleteParamsBillingProviderTypeCustom           V1CustomerBillingConfigDeleteParamsBillingProviderType = "custom"
	V1CustomerBillingConfigDeleteParamsBillingProviderTypeAzureMarketplace V1CustomerBillingConfigDeleteParamsBillingProviderType = "azure_marketplace"
	V1CustomerBillingConfigDeleteParamsBillingProviderTypeQuickbooksOnline V1CustomerBillingConfigDeleteParamsBillingProviderType = "quickbooks_online"
	V1CustomerBillingConfigDeleteParamsBillingProviderTypeWorkday          V1CustomerBillingConfigDeleteParamsBillingProviderType = "workday"
	V1CustomerBillingConfigDeleteParamsBillingProviderTypeGcpMarketplace   V1CustomerBillingConfigDeleteParamsBillingProviderType = "gcp_marketplace"
	V1CustomerBillingConfigDeleteParamsBillingProviderTypeMetronome        V1CustomerBillingConfigDeleteParamsBillingProviderType = "metronome"
)
