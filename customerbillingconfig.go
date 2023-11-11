// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/metronome/metronome-go/internal/apijson"
	"github.com/metronome/metronome-go/internal/param"
	"github.com/metronome/metronome-go/internal/requestconfig"
	"github.com/metronome/metronome-go/option"
)

// CustomerBillingConfigService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomerBillingConfigService]
// method instead.
type CustomerBillingConfigService struct {
	Options []option.RequestOption
}

// NewCustomerBillingConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomerBillingConfigService(opts ...option.RequestOption) (r *CustomerBillingConfigService) {
	r = &CustomerBillingConfigService{}
	r.Options = opts
	return
}

// Set the billing configuration for a given customer.
func (r *CustomerBillingConfigService) New(ctx context.Context, customerID string, billingProviderType CustomerBillingConfigNewParamsBillingProviderType, body CustomerBillingConfigNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("customers/%s/billing-config/%v", customerID, billingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Fetch the billing configuration for the given customer.
func (r *CustomerBillingConfigService) Get(ctx context.Context, customerID string, billingProviderType CustomerBillingConfigGetParamsBillingProviderType, opts ...option.RequestOption) (res *CustomerBillingConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s/billing-config/%v", customerID, billingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the billing configuration for a given customer. Note: this is unsupported
// for Azure and AWS Marketplace customers.
func (r *CustomerBillingConfigService) Delete(ctx context.Context, customerID string, billingProviderType CustomerBillingConfigDeleteParamsBillingProviderType, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("customers/%s/billing-config/%v", customerID, billingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CustomerBillingConfigGetResponse struct {
	Data CustomerBillingConfigGetResponseData `json:"data,required"`
	JSON customerBillingConfigGetResponseJSON
}

// customerBillingConfigGetResponseJSON contains the JSON metadata for the struct
// [CustomerBillingConfigGetResponse]
type customerBillingConfigGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerBillingConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerBillingConfigGetResponseData struct {
	// Contract expiration date for the customer. The expected format is RFC 3339 and
	// can be retrieved from AWS's GetEntitlements API. (See
	// https://docs.aws.amazon.com/marketplaceentitlement/latest/APIReference/API_GetEntitlements.html.)
	AwsExpirationDate time.Time                                     `json:"aws_expiration_date" format:"date-time"`
	AwsProductCode    string                                        `json:"aws_product_code"`
	AwsRegion         CustomerBillingConfigGetResponseDataAwsRegion `json:"aws_region"`
	// Subscription term start/end date for the customer. The expected format is RFC
	// 3339 and can be retrieved from Azure's Get Subscription API. (See
	// https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-subscription.)
	AzureExpirationDate time.Time `json:"azure_expiration_date" format:"date-time"`
	AzurePlanID         string    `json:"azure_plan_id" format:"uuid"`
	// Subscription term start/end date for the customer. The expected format is RFC
	// 3339 and can be retrieved from Azure's Get Subscription API. (See
	// https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-subscription.)
	AzureStartDate            time.Time                                                   `json:"azure_start_date" format:"date-time"`
	AzureSubscriptionStatus   CustomerBillingConfigGetResponseDataAzureSubscriptionStatus `json:"azure_subscription_status"`
	BillingProviderCustomerID string                                                      `json:"billing_provider_customer_id"`
	StripeCollectionMethod    CustomerBillingConfigGetResponseDataStripeCollectionMethod  `json:"stripe_collection_method"`
	JSON                      customerBillingConfigGetResponseDataJSON
}

// customerBillingConfigGetResponseDataJSON contains the JSON metadata for the
// struct [CustomerBillingConfigGetResponseData]
type customerBillingConfigGetResponseDataJSON struct {
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

func (r *CustomerBillingConfigGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerBillingConfigGetResponseDataAwsRegion string

const (
	CustomerBillingConfigGetResponseDataAwsRegionAfSouth1     CustomerBillingConfigGetResponseDataAwsRegion = "af-south-1"
	CustomerBillingConfigGetResponseDataAwsRegionApEast1      CustomerBillingConfigGetResponseDataAwsRegion = "ap-east-1"
	CustomerBillingConfigGetResponseDataAwsRegionApNortheast1 CustomerBillingConfigGetResponseDataAwsRegion = "ap-northeast-1"
	CustomerBillingConfigGetResponseDataAwsRegionApNortheast2 CustomerBillingConfigGetResponseDataAwsRegion = "ap-northeast-2"
	CustomerBillingConfigGetResponseDataAwsRegionApNortheast3 CustomerBillingConfigGetResponseDataAwsRegion = "ap-northeast-3"
	CustomerBillingConfigGetResponseDataAwsRegionApSouth1     CustomerBillingConfigGetResponseDataAwsRegion = "ap-south-1"
	CustomerBillingConfigGetResponseDataAwsRegionApSoutheast1 CustomerBillingConfigGetResponseDataAwsRegion = "ap-southeast-1"
	CustomerBillingConfigGetResponseDataAwsRegionApSoutheast2 CustomerBillingConfigGetResponseDataAwsRegion = "ap-southeast-2"
	CustomerBillingConfigGetResponseDataAwsRegionCaCentral1   CustomerBillingConfigGetResponseDataAwsRegion = "ca-central-1"
	CustomerBillingConfigGetResponseDataAwsRegionCnNorth1     CustomerBillingConfigGetResponseDataAwsRegion = "cn-north-1"
	CustomerBillingConfigGetResponseDataAwsRegionCnNorthwest1 CustomerBillingConfigGetResponseDataAwsRegion = "cn-northwest-1"
	CustomerBillingConfigGetResponseDataAwsRegionEuCentral1   CustomerBillingConfigGetResponseDataAwsRegion = "eu-central-1"
	CustomerBillingConfigGetResponseDataAwsRegionEuNorth1     CustomerBillingConfigGetResponseDataAwsRegion = "eu-north-1"
	CustomerBillingConfigGetResponseDataAwsRegionEuSouth1     CustomerBillingConfigGetResponseDataAwsRegion = "eu-south-1"
	CustomerBillingConfigGetResponseDataAwsRegionEuWest1      CustomerBillingConfigGetResponseDataAwsRegion = "eu-west-1"
	CustomerBillingConfigGetResponseDataAwsRegionEuWest2      CustomerBillingConfigGetResponseDataAwsRegion = "eu-west-2"
	CustomerBillingConfigGetResponseDataAwsRegionEuWest3      CustomerBillingConfigGetResponseDataAwsRegion = "eu-west-3"
	CustomerBillingConfigGetResponseDataAwsRegionMeSouth1     CustomerBillingConfigGetResponseDataAwsRegion = "me-south-1"
	CustomerBillingConfigGetResponseDataAwsRegionSaEast1      CustomerBillingConfigGetResponseDataAwsRegion = "sa-east-1"
	CustomerBillingConfigGetResponseDataAwsRegionUsEast1      CustomerBillingConfigGetResponseDataAwsRegion = "us-east-1"
	CustomerBillingConfigGetResponseDataAwsRegionUsEast2      CustomerBillingConfigGetResponseDataAwsRegion = "us-east-2"
	CustomerBillingConfigGetResponseDataAwsRegionUsGovEast1   CustomerBillingConfigGetResponseDataAwsRegion = "us-gov-east-1"
	CustomerBillingConfigGetResponseDataAwsRegionUsGovWest1   CustomerBillingConfigGetResponseDataAwsRegion = "us-gov-west-1"
	CustomerBillingConfigGetResponseDataAwsRegionUsWest1      CustomerBillingConfigGetResponseDataAwsRegion = "us-west-1"
	CustomerBillingConfigGetResponseDataAwsRegionUsWest2      CustomerBillingConfigGetResponseDataAwsRegion = "us-west-2"
)

type CustomerBillingConfigGetResponseDataAzureSubscriptionStatus string

const (
	CustomerBillingConfigGetResponseDataAzureSubscriptionStatusSubscribed              CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "Subscribed"
	CustomerBillingConfigGetResponseDataAzureSubscriptionStatusUnsubscribed            CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "Unsubscribed"
	CustomerBillingConfigGetResponseDataAzureSubscriptionStatusSuspended               CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "Suspended"
	CustomerBillingConfigGetResponseDataAzureSubscriptionStatusPendingFulfillmentStart CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "PendingFulfillmentStart"
)

type CustomerBillingConfigGetResponseDataStripeCollectionMethod string

const (
	CustomerBillingConfigGetResponseDataStripeCollectionMethodChargeAutomatically CustomerBillingConfigGetResponseDataStripeCollectionMethod = "charge_automatically"
	CustomerBillingConfigGetResponseDataStripeCollectionMethodSendInvoice         CustomerBillingConfigGetResponseDataStripeCollectionMethod = "send_invoice"
)

type CustomerBillingConfigNewParams struct {
	// The customer ID in the billing provider's system. For Azure, this is the
	// subscription ID.
	BillingProviderCustomerID param.Field[string]                                               `json:"billing_provider_customer_id,required"`
	AwsProductCode            param.Field[string]                                               `json:"aws_product_code"`
	AwsRegion                 param.Field[CustomerBillingConfigNewParamsAwsRegion]              `json:"aws_region"`
	StripeCollectionMethod    param.Field[CustomerBillingConfigNewParamsStripeCollectionMethod] `json:"stripe_collection_method"`
}

func (r CustomerBillingConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerBillingConfigNewParamsBillingProviderType string

const (
	CustomerBillingConfigNewParamsBillingProviderTypeAwsMarketplace   CustomerBillingConfigNewParamsBillingProviderType = "aws_marketplace"
	CustomerBillingConfigNewParamsBillingProviderTypeStripe           CustomerBillingConfigNewParamsBillingProviderType = "stripe"
	CustomerBillingConfigNewParamsBillingProviderTypeNetsuite         CustomerBillingConfigNewParamsBillingProviderType = "netsuite"
	CustomerBillingConfigNewParamsBillingProviderTypeCustom           CustomerBillingConfigNewParamsBillingProviderType = "custom"
	CustomerBillingConfigNewParamsBillingProviderTypeAzureMarketplace CustomerBillingConfigNewParamsBillingProviderType = "azure_marketplace"
	CustomerBillingConfigNewParamsBillingProviderTypeQuickbooksOnline CustomerBillingConfigNewParamsBillingProviderType = "quickbooks_online"
)

type CustomerBillingConfigNewParamsAwsRegion string

const (
	CustomerBillingConfigNewParamsAwsRegionAfSouth1     CustomerBillingConfigNewParamsAwsRegion = "af-south-1"
	CustomerBillingConfigNewParamsAwsRegionApEast1      CustomerBillingConfigNewParamsAwsRegion = "ap-east-1"
	CustomerBillingConfigNewParamsAwsRegionApNortheast1 CustomerBillingConfigNewParamsAwsRegion = "ap-northeast-1"
	CustomerBillingConfigNewParamsAwsRegionApNortheast2 CustomerBillingConfigNewParamsAwsRegion = "ap-northeast-2"
	CustomerBillingConfigNewParamsAwsRegionApNortheast3 CustomerBillingConfigNewParamsAwsRegion = "ap-northeast-3"
	CustomerBillingConfigNewParamsAwsRegionApSouth1     CustomerBillingConfigNewParamsAwsRegion = "ap-south-1"
	CustomerBillingConfigNewParamsAwsRegionApSoutheast1 CustomerBillingConfigNewParamsAwsRegion = "ap-southeast-1"
	CustomerBillingConfigNewParamsAwsRegionApSoutheast2 CustomerBillingConfigNewParamsAwsRegion = "ap-southeast-2"
	CustomerBillingConfigNewParamsAwsRegionCaCentral1   CustomerBillingConfigNewParamsAwsRegion = "ca-central-1"
	CustomerBillingConfigNewParamsAwsRegionCnNorth1     CustomerBillingConfigNewParamsAwsRegion = "cn-north-1"
	CustomerBillingConfigNewParamsAwsRegionCnNorthwest1 CustomerBillingConfigNewParamsAwsRegion = "cn-northwest-1"
	CustomerBillingConfigNewParamsAwsRegionEuCentral1   CustomerBillingConfigNewParamsAwsRegion = "eu-central-1"
	CustomerBillingConfigNewParamsAwsRegionEuNorth1     CustomerBillingConfigNewParamsAwsRegion = "eu-north-1"
	CustomerBillingConfigNewParamsAwsRegionEuSouth1     CustomerBillingConfigNewParamsAwsRegion = "eu-south-1"
	CustomerBillingConfigNewParamsAwsRegionEuWest1      CustomerBillingConfigNewParamsAwsRegion = "eu-west-1"
	CustomerBillingConfigNewParamsAwsRegionEuWest2      CustomerBillingConfigNewParamsAwsRegion = "eu-west-2"
	CustomerBillingConfigNewParamsAwsRegionEuWest3      CustomerBillingConfigNewParamsAwsRegion = "eu-west-3"
	CustomerBillingConfigNewParamsAwsRegionMeSouth1     CustomerBillingConfigNewParamsAwsRegion = "me-south-1"
	CustomerBillingConfigNewParamsAwsRegionSaEast1      CustomerBillingConfigNewParamsAwsRegion = "sa-east-1"
	CustomerBillingConfigNewParamsAwsRegionUsEast1      CustomerBillingConfigNewParamsAwsRegion = "us-east-1"
	CustomerBillingConfigNewParamsAwsRegionUsEast2      CustomerBillingConfigNewParamsAwsRegion = "us-east-2"
	CustomerBillingConfigNewParamsAwsRegionUsGovEast1   CustomerBillingConfigNewParamsAwsRegion = "us-gov-east-1"
	CustomerBillingConfigNewParamsAwsRegionUsGovWest1   CustomerBillingConfigNewParamsAwsRegion = "us-gov-west-1"
	CustomerBillingConfigNewParamsAwsRegionUsWest1      CustomerBillingConfigNewParamsAwsRegion = "us-west-1"
	CustomerBillingConfigNewParamsAwsRegionUsWest2      CustomerBillingConfigNewParamsAwsRegion = "us-west-2"
)

type CustomerBillingConfigNewParamsStripeCollectionMethod string

const (
	CustomerBillingConfigNewParamsStripeCollectionMethodChargeAutomatically CustomerBillingConfigNewParamsStripeCollectionMethod = "charge_automatically"
	CustomerBillingConfigNewParamsStripeCollectionMethodSendInvoice         CustomerBillingConfigNewParamsStripeCollectionMethod = "send_invoice"
)

type CustomerBillingConfigGetParamsBillingProviderType string

const (
	CustomerBillingConfigGetParamsBillingProviderTypeAwsMarketplace   CustomerBillingConfigGetParamsBillingProviderType = "aws_marketplace"
	CustomerBillingConfigGetParamsBillingProviderTypeStripe           CustomerBillingConfigGetParamsBillingProviderType = "stripe"
	CustomerBillingConfigGetParamsBillingProviderTypeNetsuite         CustomerBillingConfigGetParamsBillingProviderType = "netsuite"
	CustomerBillingConfigGetParamsBillingProviderTypeCustom           CustomerBillingConfigGetParamsBillingProviderType = "custom"
	CustomerBillingConfigGetParamsBillingProviderTypeAzureMarketplace CustomerBillingConfigGetParamsBillingProviderType = "azure_marketplace"
	CustomerBillingConfigGetParamsBillingProviderTypeQuickbooksOnline CustomerBillingConfigGetParamsBillingProviderType = "quickbooks_online"
)

type CustomerBillingConfigDeleteParamsBillingProviderType string

const (
	CustomerBillingConfigDeleteParamsBillingProviderTypeAwsMarketplace   CustomerBillingConfigDeleteParamsBillingProviderType = "aws_marketplace"
	CustomerBillingConfigDeleteParamsBillingProviderTypeStripe           CustomerBillingConfigDeleteParamsBillingProviderType = "stripe"
	CustomerBillingConfigDeleteParamsBillingProviderTypeNetsuite         CustomerBillingConfigDeleteParamsBillingProviderType = "netsuite"
	CustomerBillingConfigDeleteParamsBillingProviderTypeCustom           CustomerBillingConfigDeleteParamsBillingProviderType = "custom"
	CustomerBillingConfigDeleteParamsBillingProviderTypeAzureMarketplace CustomerBillingConfigDeleteParamsBillingProviderType = "azure_marketplace"
	CustomerBillingConfigDeleteParamsBillingProviderTypeQuickbooksOnline CustomerBillingConfigDeleteParamsBillingProviderType = "quickbooks_online"
)
