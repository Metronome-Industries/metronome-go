// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Metronome-Industries/metronome-go/v2"
	"github.com/Metronome-Industries/metronome-go/v2/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/v2/option"
)

func TestV1SettingUpsertAvalaraCredentials(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.V1.Settings.UpsertAvalaraCredentials(context.TODO(), metronome.V1SettingUpsertAvalaraCredentialsParams{
		AvalaraEnvironment: metronome.V1SettingUpsertAvalaraCredentialsParamsAvalaraEnvironmentProduction,
		AvalaraPassword:    "my_password_123",
		AvalaraUsername:    "test@metronome.com",
		DeliveryMethodIDs:  []string{"9a906ebb-fbc7-42e8-8e29-53bfd2db3aca"},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
