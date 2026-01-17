// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"slices"

	"github.com/Metronome-Industries/metronome-go/v3/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/v3/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"github.com/Metronome-Industries/metronome-go/v3/packages/param"
	"github.com/Metronome-Industries/metronome-go/v3/packages/respjson"
)

// V1SettingService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1SettingService] method instead.
type V1SettingService struct {
	Options          []option.RequestOption
	BillingProviders V1SettingBillingProviderService
}

// NewV1SettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1SettingService(opts ...option.RequestOption) (r V1SettingService) {
	r = V1SettingService{}
	r.Options = opts
	r.BillingProviders = NewV1SettingBillingProviderService(opts...)
	return
}

// Set the Avalara credentials for some specified `delivery_method_ids`, which can
// be found in the `/listConfiguredBillingProviders` response. This maps the
// Avalara credentials to the appropriate billing entity. These credentials are
// only used for PLG Invoicing today.
func (r *V1SettingService) UpsertAvalaraCredentials(ctx context.Context, body V1SettingUpsertAvalaraCredentialsParams, opts ...option.RequestOption) (res *V1SettingUpsertAvalaraCredentialsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/upsertAvalaraCredentials"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1SettingUpsertAvalaraCredentialsResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SettingUpsertAvalaraCredentialsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SettingUpsertAvalaraCredentialsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SettingUpsertAvalaraCredentialsParams struct {
	// The Avalara environment to use (SANDBOX or PRODUCTION).
	//
	// Any of "PRODUCTION", "SANDBOX".
	AvalaraEnvironment V1SettingUpsertAvalaraCredentialsParamsAvalaraEnvironment `json:"avalara_environment,omitzero,required"`
	// The password for the Avalara account.
	AvalaraPassword string `json:"avalara_password,required"`
	// The username for the Avalara account.
	AvalaraUsername string `json:"avalara_username,required"`
	// The delivery method IDs of the billing provider configurations to update, can be
	// found in the response of the `/listConfiguredBillingProviders` endpoint.
	DeliveryMethodIDs []string `json:"delivery_method_ids,omitzero,required"`
	// Commit transactions if you want Metronome tax calculations used for reporting
	// and tax filings.
	CommitTransactions param.Opt[bool] `json:"commit_transactions,omitzero"`
	paramObj
}

func (r V1SettingUpsertAvalaraCredentialsParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SettingUpsertAvalaraCredentialsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SettingUpsertAvalaraCredentialsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Avalara environment to use (SANDBOX or PRODUCTION).
type V1SettingUpsertAvalaraCredentialsParamsAvalaraEnvironment string

const (
	V1SettingUpsertAvalaraCredentialsParamsAvalaraEnvironmentProduction V1SettingUpsertAvalaraCredentialsParamsAvalaraEnvironment = "PRODUCTION"
	V1SettingUpsertAvalaraCredentialsParamsAvalaraEnvironmentSandbox    V1SettingUpsertAvalaraCredentialsParamsAvalaraEnvironment = "SANDBOX"
)
