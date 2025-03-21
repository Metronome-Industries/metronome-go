// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
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
func NewV1CustomerBillingConfigService(opts ...option.RequestOption) (r *V1CustomerBillingConfigService) {
	r = &V1CustomerBillingConfigService{}
	r.Options = opts
	return
}

// Set the billing configuration for a given customer.
func (r *V1CustomerBillingConfigService) New(ctx context.Context, params V1CustomerBillingConfigNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if !params.BillingProviderType.Present {
		err = errors.New("missing required billing_provider_type parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/billing-config/%v", params.CustomerID, params.BillingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Fetch the billing configuration for the given customer.
func (r *V1CustomerBillingConfigService) Get(ctx context.Context, query V1CustomerBillingConfigGetParams, opts ...option.RequestOption) (res *V1CustomerBillingConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if !query.BillingProviderType.Present {
		err = errors.New("missing required billing_provider_type parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/billing-config/%v", query.CustomerID, query.BillingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the billing configuration for a given customer. Note: this is unsupported
// for Azure and AWS Marketplace customers.
func (r *V1CustomerBillingConfigService) Delete(ctx context.Context, body V1CustomerBillingConfigDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if !body.BillingProviderType.Present {
		err = errors.New("missing required billing_provider_type parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/billing-config/%v", body.CustomerID, body.BillingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type V1CustomerBillingConfigGetResponse struct {
	Data V1CustomerBillingConfigGetResponseData `json:"data,required"`
	JSON v1CustomerBillingConfigGetResponseJSON `json:"-"`
}

// v1CustomerBillingConfigGetResponseJSON contains the JSON metadata for the struct
// [V1CustomerBillingConfigGetResponse]
type v1CustomerBillingConfigGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1CustomerBillingConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerBillingConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1CustomerBillingConfigGetResponseData struct {
	// Contract expiration date for the customer. The expected format is RFC 3339 and
	// can be retrieved from
	// [AWS's GetEntitlements API](https://docs.aws.amazon.com/marketplaceentitlement/latest/APIReference/API_GetEntitlements.html).
	AwsExpirationDate time.Time                                       `json:"aws_expiration_date" format:"date-time"`
	AwsProductCode    string                                          `json:"aws_product_code"`
	AwsRegion         V1CustomerBillingConfigGetResponseDataAwsRegion `json:"aws_region"`
	// Subscription term start/end date for the customer. The expected format is RFC
	// 3339 and can be retrieved from
	// [Azure's Get Subscription API](https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-subscription).
	AzureExpirationDate time.Time `json:"azure_expiration_date" format:"date-time"`
	AzurePlanID         string    `json:"azure_plan_id" format:"uuid"`
	// Subscription term start/end date for the customer. The expected format is RFC
	// 3339 and can be retrieved from
	// [Azure's Get Subscription API](https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-subscription).
	AzureStartDate            time.Time                                                     `json:"azure_start_date" format:"date-time"`
	AzureSubscriptionStatus   V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatus `json:"azure_subscription_status"`
	BillingProviderCustomerID string                                                        `json:"billing_provider_customer_id"`
	StripeCollectionMethod    V1CustomerBillingConfigGetResponseDataStripeCollectionMethod  `json:"stripe_collection_method"`
	JSON                      v1CustomerBillingConfigGetResponseDataJSON                    `json:"-"`
}

// v1CustomerBillingConfigGetResponseDataJSON contains the JSON metadata for the
// struct [V1CustomerBillingConfigGetResponseData]
type v1CustomerBillingConfigGetResponseDataJSON struct {
	AwsExpirationDate         apijson.Field
	AwsProductCode            apijson.Field
	AwsRegion                 apijson.Field
	AzureExpirationDate       apijson.Field
	AzurePlanID               apijson.Field
	AzureStartDate            apijson.Field
	AzureSubscriptionStatus   apijson.Field
	BillingProviderCustomerID apijson.Field
	StripeCollectionMethod    apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *V1CustomerBillingConfigGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1CustomerBillingConfigGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1CustomerBillingConfigGetResponseDataAwsRegion string

const (
	V1CustomerBillingConfigGetResponseDataAwsRegionAfSouth1     V1CustomerBillingConfigGetResponseDataAwsRegion = "af-south-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionApEast1      V1CustomerBillingConfigGetResponseDataAwsRegion = "ap-east-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionApNortheast1 V1CustomerBillingConfigGetResponseDataAwsRegion = "ap-northeast-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionApNortheast2 V1CustomerBillingConfigGetResponseDataAwsRegion = "ap-northeast-2"
	V1CustomerBillingConfigGetResponseDataAwsRegionApNortheast3 V1CustomerBillingConfigGetResponseDataAwsRegion = "ap-northeast-3"
	V1CustomerBillingConfigGetResponseDataAwsRegionApSouth1     V1CustomerBillingConfigGetResponseDataAwsRegion = "ap-south-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionApSoutheast1 V1CustomerBillingConfigGetResponseDataAwsRegion = "ap-southeast-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionApSoutheast2 V1CustomerBillingConfigGetResponseDataAwsRegion = "ap-southeast-2"
	V1CustomerBillingConfigGetResponseDataAwsRegionCaCentral1   V1CustomerBillingConfigGetResponseDataAwsRegion = "ca-central-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionCnNorth1     V1CustomerBillingConfigGetResponseDataAwsRegion = "cn-north-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionCnNorthwest1 V1CustomerBillingConfigGetResponseDataAwsRegion = "cn-northwest-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionEuCentral1   V1CustomerBillingConfigGetResponseDataAwsRegion = "eu-central-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionEuNorth1     V1CustomerBillingConfigGetResponseDataAwsRegion = "eu-north-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionEuSouth1     V1CustomerBillingConfigGetResponseDataAwsRegion = "eu-south-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionEuWest1      V1CustomerBillingConfigGetResponseDataAwsRegion = "eu-west-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionEuWest2      V1CustomerBillingConfigGetResponseDataAwsRegion = "eu-west-2"
	V1CustomerBillingConfigGetResponseDataAwsRegionEuWest3      V1CustomerBillingConfigGetResponseDataAwsRegion = "eu-west-3"
	V1CustomerBillingConfigGetResponseDataAwsRegionMeSouth1     V1CustomerBillingConfigGetResponseDataAwsRegion = "me-south-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionSaEast1      V1CustomerBillingConfigGetResponseDataAwsRegion = "sa-east-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionUsEast1      V1CustomerBillingConfigGetResponseDataAwsRegion = "us-east-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionUsEast2      V1CustomerBillingConfigGetResponseDataAwsRegion = "us-east-2"
	V1CustomerBillingConfigGetResponseDataAwsRegionUsGovEast1   V1CustomerBillingConfigGetResponseDataAwsRegion = "us-gov-east-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionUsGovWest1   V1CustomerBillingConfigGetResponseDataAwsRegion = "us-gov-west-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionUsWest1      V1CustomerBillingConfigGetResponseDataAwsRegion = "us-west-1"
	V1CustomerBillingConfigGetResponseDataAwsRegionUsWest2      V1CustomerBillingConfigGetResponseDataAwsRegion = "us-west-2"
)

func (r V1CustomerBillingConfigGetResponseDataAwsRegion) IsKnown() bool {
	switch r {
	case V1CustomerBillingConfigGetResponseDataAwsRegionAfSouth1, V1CustomerBillingConfigGetResponseDataAwsRegionApEast1, V1CustomerBillingConfigGetResponseDataAwsRegionApNortheast1, V1CustomerBillingConfigGetResponseDataAwsRegionApNortheast2, V1CustomerBillingConfigGetResponseDataAwsRegionApNortheast3, V1CustomerBillingConfigGetResponseDataAwsRegionApSouth1, V1CustomerBillingConfigGetResponseDataAwsRegionApSoutheast1, V1CustomerBillingConfigGetResponseDataAwsRegionApSoutheast2, V1CustomerBillingConfigGetResponseDataAwsRegionCaCentral1, V1CustomerBillingConfigGetResponseDataAwsRegionCnNorth1, V1CustomerBillingConfigGetResponseDataAwsRegionCnNorthwest1, V1CustomerBillingConfigGetResponseDataAwsRegionEuCentral1, V1CustomerBillingConfigGetResponseDataAwsRegionEuNorth1, V1CustomerBillingConfigGetResponseDataAwsRegionEuSouth1, V1CustomerBillingConfigGetResponseDataAwsRegionEuWest1, V1CustomerBillingConfigGetResponseDataAwsRegionEuWest2, V1CustomerBillingConfigGetResponseDataAwsRegionEuWest3, V1CustomerBillingConfigGetResponseDataAwsRegionMeSouth1, V1CustomerBillingConfigGetResponseDataAwsRegionSaEast1, V1CustomerBillingConfigGetResponseDataAwsRegionUsEast1, V1CustomerBillingConfigGetResponseDataAwsRegionUsEast2, V1CustomerBillingConfigGetResponseDataAwsRegionUsGovEast1, V1CustomerBillingConfigGetResponseDataAwsRegionUsGovWest1, V1CustomerBillingConfigGetResponseDataAwsRegionUsWest1, V1CustomerBillingConfigGetResponseDataAwsRegionUsWest2:
		return true
	}
	return false
}

type V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatus string

const (
	V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatusSubscribed              V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "Subscribed"
	V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatusUnsubscribed            V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "Unsubscribed"
	V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatusSuspended               V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "Suspended"
	V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatusPendingFulfillmentStart V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "PendingFulfillmentStart"
)

func (r V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatus) IsKnown() bool {
	switch r {
	case V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatusSubscribed, V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatusUnsubscribed, V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatusSuspended, V1CustomerBillingConfigGetResponseDataAzureSubscriptionStatusPendingFulfillmentStart:
		return true
	}
	return false
}

type V1CustomerBillingConfigGetResponseDataStripeCollectionMethod string

const (
	V1CustomerBillingConfigGetResponseDataStripeCollectionMethodChargeAutomatically V1CustomerBillingConfigGetResponseDataStripeCollectionMethod = "charge_automatically"
	V1CustomerBillingConfigGetResponseDataStripeCollectionMethodSendInvoice         V1CustomerBillingConfigGetResponseDataStripeCollectionMethod = "send_invoice"
)

func (r V1CustomerBillingConfigGetResponseDataStripeCollectionMethod) IsKnown() bool {
	switch r {
	case V1CustomerBillingConfigGetResponseDataStripeCollectionMethodChargeAutomatically, V1CustomerBillingConfigGetResponseDataStripeCollectionMethodSendInvoice:
		return true
	}
	return false
}

type V1CustomerBillingConfigNewParams struct {
	CustomerID          param.Field[string]                                              `path:"customer_id,required" format:"uuid"`
	BillingProviderType param.Field[V1CustomerBillingConfigNewParamsBillingProviderType] `path:"billing_provider_type,required"`
	// The customer ID in the billing provider's system. For Azure, this is the
	// subscription ID.
	BillingProviderCustomerID param.Field[string]                                                 `json:"billing_provider_customer_id,required"`
	AwsProductCode            param.Field[string]                                                 `json:"aws_product_code"`
	AwsRegion                 param.Field[V1CustomerBillingConfigNewParamsAwsRegion]              `json:"aws_region"`
	StripeCollectionMethod    param.Field[V1CustomerBillingConfigNewParamsStripeCollectionMethod] `json:"stripe_collection_method"`
}

func (r V1CustomerBillingConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
)

func (r V1CustomerBillingConfigNewParamsBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerBillingConfigNewParamsBillingProviderTypeAwsMarketplace, V1CustomerBillingConfigNewParamsBillingProviderTypeStripe, V1CustomerBillingConfigNewParamsBillingProviderTypeNetsuite, V1CustomerBillingConfigNewParamsBillingProviderTypeCustom, V1CustomerBillingConfigNewParamsBillingProviderTypeAzureMarketplace, V1CustomerBillingConfigNewParamsBillingProviderTypeQuickbooksOnline, V1CustomerBillingConfigNewParamsBillingProviderTypeWorkday, V1CustomerBillingConfigNewParamsBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

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

func (r V1CustomerBillingConfigNewParamsAwsRegion) IsKnown() bool {
	switch r {
	case V1CustomerBillingConfigNewParamsAwsRegionAfSouth1, V1CustomerBillingConfigNewParamsAwsRegionApEast1, V1CustomerBillingConfigNewParamsAwsRegionApNortheast1, V1CustomerBillingConfigNewParamsAwsRegionApNortheast2, V1CustomerBillingConfigNewParamsAwsRegionApNortheast3, V1CustomerBillingConfigNewParamsAwsRegionApSouth1, V1CustomerBillingConfigNewParamsAwsRegionApSoutheast1, V1CustomerBillingConfigNewParamsAwsRegionApSoutheast2, V1CustomerBillingConfigNewParamsAwsRegionCaCentral1, V1CustomerBillingConfigNewParamsAwsRegionCnNorth1, V1CustomerBillingConfigNewParamsAwsRegionCnNorthwest1, V1CustomerBillingConfigNewParamsAwsRegionEuCentral1, V1CustomerBillingConfigNewParamsAwsRegionEuNorth1, V1CustomerBillingConfigNewParamsAwsRegionEuSouth1, V1CustomerBillingConfigNewParamsAwsRegionEuWest1, V1CustomerBillingConfigNewParamsAwsRegionEuWest2, V1CustomerBillingConfigNewParamsAwsRegionEuWest3, V1CustomerBillingConfigNewParamsAwsRegionMeSouth1, V1CustomerBillingConfigNewParamsAwsRegionSaEast1, V1CustomerBillingConfigNewParamsAwsRegionUsEast1, V1CustomerBillingConfigNewParamsAwsRegionUsEast2, V1CustomerBillingConfigNewParamsAwsRegionUsGovEast1, V1CustomerBillingConfigNewParamsAwsRegionUsGovWest1, V1CustomerBillingConfigNewParamsAwsRegionUsWest1, V1CustomerBillingConfigNewParamsAwsRegionUsWest2:
		return true
	}
	return false
}

type V1CustomerBillingConfigNewParamsStripeCollectionMethod string

const (
	V1CustomerBillingConfigNewParamsStripeCollectionMethodChargeAutomatically V1CustomerBillingConfigNewParamsStripeCollectionMethod = "charge_automatically"
	V1CustomerBillingConfigNewParamsStripeCollectionMethodSendInvoice         V1CustomerBillingConfigNewParamsStripeCollectionMethod = "send_invoice"
)

func (r V1CustomerBillingConfigNewParamsStripeCollectionMethod) IsKnown() bool {
	switch r {
	case V1CustomerBillingConfigNewParamsStripeCollectionMethodChargeAutomatically, V1CustomerBillingConfigNewParamsStripeCollectionMethodSendInvoice:
		return true
	}
	return false
}

type V1CustomerBillingConfigGetParams struct {
	CustomerID          param.Field[string]                                              `path:"customer_id,required" format:"uuid"`
	BillingProviderType param.Field[V1CustomerBillingConfigGetParamsBillingProviderType] `path:"billing_provider_type,required"`
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
)

func (r V1CustomerBillingConfigGetParamsBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerBillingConfigGetParamsBillingProviderTypeAwsMarketplace, V1CustomerBillingConfigGetParamsBillingProviderTypeStripe, V1CustomerBillingConfigGetParamsBillingProviderTypeNetsuite, V1CustomerBillingConfigGetParamsBillingProviderTypeCustom, V1CustomerBillingConfigGetParamsBillingProviderTypeAzureMarketplace, V1CustomerBillingConfigGetParamsBillingProviderTypeQuickbooksOnline, V1CustomerBillingConfigGetParamsBillingProviderTypeWorkday, V1CustomerBillingConfigGetParamsBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type V1CustomerBillingConfigDeleteParams struct {
	CustomerID          param.Field[string]                                                 `path:"customer_id,required" format:"uuid"`
	BillingProviderType param.Field[V1CustomerBillingConfigDeleteParamsBillingProviderType] `path:"billing_provider_type,required"`
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
)

func (r V1CustomerBillingConfigDeleteParamsBillingProviderType) IsKnown() bool {
	switch r {
	case V1CustomerBillingConfigDeleteParamsBillingProviderTypeAwsMarketplace, V1CustomerBillingConfigDeleteParamsBillingProviderTypeStripe, V1CustomerBillingConfigDeleteParamsBillingProviderTypeNetsuite, V1CustomerBillingConfigDeleteParamsBillingProviderTypeCustom, V1CustomerBillingConfigDeleteParamsBillingProviderTypeAzureMarketplace, V1CustomerBillingConfigDeleteParamsBillingProviderTypeQuickbooksOnline, V1CustomerBillingConfigDeleteParamsBillingProviderTypeWorkday, V1CustomerBillingConfigDeleteParamsBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}
