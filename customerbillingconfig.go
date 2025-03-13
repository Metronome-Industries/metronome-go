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

// CustomerBillingConfigService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerBillingConfigService] method instead.
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
func (r *CustomerBillingConfigService) New(ctx context.Context, params CustomerBillingConfigNewParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("customers/%s/billing-config/%v", params.CustomerID, params.BillingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Fetch the billing configuration for the given customer.
func (r *CustomerBillingConfigService) Get(ctx context.Context, query CustomerBillingConfigGetParams, opts ...option.RequestOption) (res *CustomerBillingConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.CustomerID.Value == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if !query.BillingProviderType.Present {
		err = errors.New("missing required billing_provider_type parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/billing-config/%v", query.CustomerID, query.BillingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the billing configuration for a given customer. Note: this is unsupported
// for Azure and AWS Marketplace customers.
func (r *CustomerBillingConfigService) Delete(ctx context.Context, body CustomerBillingConfigDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("customers/%s/billing-config/%v", body.CustomerID, body.BillingProviderType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CustomerBillingConfigGetResponse struct {
	Data CustomerBillingConfigGetResponseData `json:"data,required"`
	JSON customerBillingConfigGetResponseJSON `json:"-"`
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

func (r customerBillingConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerBillingConfigGetResponseData struct {
	// Contract expiration date for the customer. The expected format is RFC 3339 and
	// can be retrieved from
	// [AWS's GetEntitlements API](https://docs.aws.amazon.com/marketplaceentitlement/latest/APIReference/API_GetEntitlements.html).
	AwsExpirationDate time.Time                                     `json:"aws_expiration_date" format:"date-time"`
	AwsProductCode    string                                        `json:"aws_product_code"`
	AwsRegion         CustomerBillingConfigGetResponseDataAwsRegion `json:"aws_region"`
	// Subscription term start/end date for the customer. The expected format is RFC
	// 3339 and can be retrieved from
	// [Azure's Get Subscription API](https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-subscription).
	AzureExpirationDate time.Time `json:"azure_expiration_date" format:"date-time"`
	AzurePlanID         string    `json:"azure_plan_id" format:"uuid"`
	// Subscription term start/end date for the customer. The expected format is RFC
	// 3339 and can be retrieved from
	// [Azure's Get Subscription API](https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-subscription).
	AzureStartDate            time.Time                                                   `json:"azure_start_date" format:"date-time"`
	AzureSubscriptionStatus   CustomerBillingConfigGetResponseDataAzureSubscriptionStatus `json:"azure_subscription_status"`
	BillingProviderCustomerID string                                                      `json:"billing_provider_customer_id"`
	StripeCollectionMethod    CustomerBillingConfigGetResponseDataStripeCollectionMethod  `json:"stripe_collection_method"`
	JSON                      customerBillingConfigGetResponseDataJSON                    `json:"-"`
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

func (r customerBillingConfigGetResponseDataJSON) RawJSON() string {
	return r.raw
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

func (r CustomerBillingConfigGetResponseDataAwsRegion) IsKnown() bool {
	switch r {
	case CustomerBillingConfigGetResponseDataAwsRegionAfSouth1, CustomerBillingConfigGetResponseDataAwsRegionApEast1, CustomerBillingConfigGetResponseDataAwsRegionApNortheast1, CustomerBillingConfigGetResponseDataAwsRegionApNortheast2, CustomerBillingConfigGetResponseDataAwsRegionApNortheast3, CustomerBillingConfigGetResponseDataAwsRegionApSouth1, CustomerBillingConfigGetResponseDataAwsRegionApSoutheast1, CustomerBillingConfigGetResponseDataAwsRegionApSoutheast2, CustomerBillingConfigGetResponseDataAwsRegionCaCentral1, CustomerBillingConfigGetResponseDataAwsRegionCnNorth1, CustomerBillingConfigGetResponseDataAwsRegionCnNorthwest1, CustomerBillingConfigGetResponseDataAwsRegionEuCentral1, CustomerBillingConfigGetResponseDataAwsRegionEuNorth1, CustomerBillingConfigGetResponseDataAwsRegionEuSouth1, CustomerBillingConfigGetResponseDataAwsRegionEuWest1, CustomerBillingConfigGetResponseDataAwsRegionEuWest2, CustomerBillingConfigGetResponseDataAwsRegionEuWest3, CustomerBillingConfigGetResponseDataAwsRegionMeSouth1, CustomerBillingConfigGetResponseDataAwsRegionSaEast1, CustomerBillingConfigGetResponseDataAwsRegionUsEast1, CustomerBillingConfigGetResponseDataAwsRegionUsEast2, CustomerBillingConfigGetResponseDataAwsRegionUsGovEast1, CustomerBillingConfigGetResponseDataAwsRegionUsGovWest1, CustomerBillingConfigGetResponseDataAwsRegionUsWest1, CustomerBillingConfigGetResponseDataAwsRegionUsWest2:
		return true
	}
	return false
}

type CustomerBillingConfigGetResponseDataAzureSubscriptionStatus string

const (
	CustomerBillingConfigGetResponseDataAzureSubscriptionStatusSubscribed              CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "Subscribed"
	CustomerBillingConfigGetResponseDataAzureSubscriptionStatusUnsubscribed            CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "Unsubscribed"
	CustomerBillingConfigGetResponseDataAzureSubscriptionStatusSuspended               CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "Suspended"
	CustomerBillingConfigGetResponseDataAzureSubscriptionStatusPendingFulfillmentStart CustomerBillingConfigGetResponseDataAzureSubscriptionStatus = "PendingFulfillmentStart"
)

func (r CustomerBillingConfigGetResponseDataAzureSubscriptionStatus) IsKnown() bool {
	switch r {
	case CustomerBillingConfigGetResponseDataAzureSubscriptionStatusSubscribed, CustomerBillingConfigGetResponseDataAzureSubscriptionStatusUnsubscribed, CustomerBillingConfigGetResponseDataAzureSubscriptionStatusSuspended, CustomerBillingConfigGetResponseDataAzureSubscriptionStatusPendingFulfillmentStart:
		return true
	}
	return false
}

type CustomerBillingConfigGetResponseDataStripeCollectionMethod string

const (
	CustomerBillingConfigGetResponseDataStripeCollectionMethodChargeAutomatically CustomerBillingConfigGetResponseDataStripeCollectionMethod = "charge_automatically"
	CustomerBillingConfigGetResponseDataStripeCollectionMethodSendInvoice         CustomerBillingConfigGetResponseDataStripeCollectionMethod = "send_invoice"
)

func (r CustomerBillingConfigGetResponseDataStripeCollectionMethod) IsKnown() bool {
	switch r {
	case CustomerBillingConfigGetResponseDataStripeCollectionMethodChargeAutomatically, CustomerBillingConfigGetResponseDataStripeCollectionMethodSendInvoice:
		return true
	}
	return false
}

type CustomerBillingConfigNewParams struct {
	CustomerID          param.Field[string]                                            `path:"customer_id,required" format:"uuid"`
	BillingProviderType param.Field[CustomerBillingConfigNewParamsBillingProviderType] `path:"billing_provider_type,required"`
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
	CustomerBillingConfigNewParamsBillingProviderTypeWorkday          CustomerBillingConfigNewParamsBillingProviderType = "workday"
	CustomerBillingConfigNewParamsBillingProviderTypeGcpMarketplace   CustomerBillingConfigNewParamsBillingProviderType = "gcp_marketplace"
)

func (r CustomerBillingConfigNewParamsBillingProviderType) IsKnown() bool {
	switch r {
	case CustomerBillingConfigNewParamsBillingProviderTypeAwsMarketplace, CustomerBillingConfigNewParamsBillingProviderTypeStripe, CustomerBillingConfigNewParamsBillingProviderTypeNetsuite, CustomerBillingConfigNewParamsBillingProviderTypeCustom, CustomerBillingConfigNewParamsBillingProviderTypeAzureMarketplace, CustomerBillingConfigNewParamsBillingProviderTypeQuickbooksOnline, CustomerBillingConfigNewParamsBillingProviderTypeWorkday, CustomerBillingConfigNewParamsBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

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

func (r CustomerBillingConfigNewParamsAwsRegion) IsKnown() bool {
	switch r {
	case CustomerBillingConfigNewParamsAwsRegionAfSouth1, CustomerBillingConfigNewParamsAwsRegionApEast1, CustomerBillingConfigNewParamsAwsRegionApNortheast1, CustomerBillingConfigNewParamsAwsRegionApNortheast2, CustomerBillingConfigNewParamsAwsRegionApNortheast3, CustomerBillingConfigNewParamsAwsRegionApSouth1, CustomerBillingConfigNewParamsAwsRegionApSoutheast1, CustomerBillingConfigNewParamsAwsRegionApSoutheast2, CustomerBillingConfigNewParamsAwsRegionCaCentral1, CustomerBillingConfigNewParamsAwsRegionCnNorth1, CustomerBillingConfigNewParamsAwsRegionCnNorthwest1, CustomerBillingConfigNewParamsAwsRegionEuCentral1, CustomerBillingConfigNewParamsAwsRegionEuNorth1, CustomerBillingConfigNewParamsAwsRegionEuSouth1, CustomerBillingConfigNewParamsAwsRegionEuWest1, CustomerBillingConfigNewParamsAwsRegionEuWest2, CustomerBillingConfigNewParamsAwsRegionEuWest3, CustomerBillingConfigNewParamsAwsRegionMeSouth1, CustomerBillingConfigNewParamsAwsRegionSaEast1, CustomerBillingConfigNewParamsAwsRegionUsEast1, CustomerBillingConfigNewParamsAwsRegionUsEast2, CustomerBillingConfigNewParamsAwsRegionUsGovEast1, CustomerBillingConfigNewParamsAwsRegionUsGovWest1, CustomerBillingConfigNewParamsAwsRegionUsWest1, CustomerBillingConfigNewParamsAwsRegionUsWest2:
		return true
	}
	return false
}

type CustomerBillingConfigNewParamsStripeCollectionMethod string

const (
	CustomerBillingConfigNewParamsStripeCollectionMethodChargeAutomatically CustomerBillingConfigNewParamsStripeCollectionMethod = "charge_automatically"
	CustomerBillingConfigNewParamsStripeCollectionMethodSendInvoice         CustomerBillingConfigNewParamsStripeCollectionMethod = "send_invoice"
)

func (r CustomerBillingConfigNewParamsStripeCollectionMethod) IsKnown() bool {
	switch r {
	case CustomerBillingConfigNewParamsStripeCollectionMethodChargeAutomatically, CustomerBillingConfigNewParamsStripeCollectionMethodSendInvoice:
		return true
	}
	return false
}

type CustomerBillingConfigGetParams struct {
	CustomerID          param.Field[string]                                            `path:"customer_id,required" format:"uuid"`
	BillingProviderType param.Field[CustomerBillingConfigGetParamsBillingProviderType] `path:"billing_provider_type,required"`
}

type CustomerBillingConfigGetParamsBillingProviderType string

const (
	CustomerBillingConfigGetParamsBillingProviderTypeAwsMarketplace   CustomerBillingConfigGetParamsBillingProviderType = "aws_marketplace"
	CustomerBillingConfigGetParamsBillingProviderTypeStripe           CustomerBillingConfigGetParamsBillingProviderType = "stripe"
	CustomerBillingConfigGetParamsBillingProviderTypeNetsuite         CustomerBillingConfigGetParamsBillingProviderType = "netsuite"
	CustomerBillingConfigGetParamsBillingProviderTypeCustom           CustomerBillingConfigGetParamsBillingProviderType = "custom"
	CustomerBillingConfigGetParamsBillingProviderTypeAzureMarketplace CustomerBillingConfigGetParamsBillingProviderType = "azure_marketplace"
	CustomerBillingConfigGetParamsBillingProviderTypeQuickbooksOnline CustomerBillingConfigGetParamsBillingProviderType = "quickbooks_online"
	CustomerBillingConfigGetParamsBillingProviderTypeWorkday          CustomerBillingConfigGetParamsBillingProviderType = "workday"
	CustomerBillingConfigGetParamsBillingProviderTypeGcpMarketplace   CustomerBillingConfigGetParamsBillingProviderType = "gcp_marketplace"
)

func (r CustomerBillingConfigGetParamsBillingProviderType) IsKnown() bool {
	switch r {
	case CustomerBillingConfigGetParamsBillingProviderTypeAwsMarketplace, CustomerBillingConfigGetParamsBillingProviderTypeStripe, CustomerBillingConfigGetParamsBillingProviderTypeNetsuite, CustomerBillingConfigGetParamsBillingProviderTypeCustom, CustomerBillingConfigGetParamsBillingProviderTypeAzureMarketplace, CustomerBillingConfigGetParamsBillingProviderTypeQuickbooksOnline, CustomerBillingConfigGetParamsBillingProviderTypeWorkday, CustomerBillingConfigGetParamsBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type CustomerBillingConfigDeleteParams struct {
	CustomerID          param.Field[string]                                               `path:"customer_id,required" format:"uuid"`
	BillingProviderType param.Field[CustomerBillingConfigDeleteParamsBillingProviderType] `path:"billing_provider_type,required"`
}

type CustomerBillingConfigDeleteParamsBillingProviderType string

const (
	CustomerBillingConfigDeleteParamsBillingProviderTypeAwsMarketplace   CustomerBillingConfigDeleteParamsBillingProviderType = "aws_marketplace"
	CustomerBillingConfigDeleteParamsBillingProviderTypeStripe           CustomerBillingConfigDeleteParamsBillingProviderType = "stripe"
	CustomerBillingConfigDeleteParamsBillingProviderTypeNetsuite         CustomerBillingConfigDeleteParamsBillingProviderType = "netsuite"
	CustomerBillingConfigDeleteParamsBillingProviderTypeCustom           CustomerBillingConfigDeleteParamsBillingProviderType = "custom"
	CustomerBillingConfigDeleteParamsBillingProviderTypeAzureMarketplace CustomerBillingConfigDeleteParamsBillingProviderType = "azure_marketplace"
	CustomerBillingConfigDeleteParamsBillingProviderTypeQuickbooksOnline CustomerBillingConfigDeleteParamsBillingProviderType = "quickbooks_online"
	CustomerBillingConfigDeleteParamsBillingProviderTypeWorkday          CustomerBillingConfigDeleteParamsBillingProviderType = "workday"
	CustomerBillingConfigDeleteParamsBillingProviderTypeGcpMarketplace   CustomerBillingConfigDeleteParamsBillingProviderType = "gcp_marketplace"
)

func (r CustomerBillingConfigDeleteParamsBillingProviderType) IsKnown() bool {
	switch r {
	case CustomerBillingConfigDeleteParamsBillingProviderTypeAwsMarketplace, CustomerBillingConfigDeleteParamsBillingProviderTypeStripe, CustomerBillingConfigDeleteParamsBillingProviderTypeNetsuite, CustomerBillingConfigDeleteParamsBillingProviderTypeCustom, CustomerBillingConfigDeleteParamsBillingProviderTypeAzureMarketplace, CustomerBillingConfigDeleteParamsBillingProviderTypeQuickbooksOnline, CustomerBillingConfigDeleteParamsBillingProviderTypeWorkday, CustomerBillingConfigDeleteParamsBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}
