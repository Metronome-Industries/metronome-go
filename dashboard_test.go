// File generated from our OpenAPI spec by Stainless.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/option"
)

func TestDashboardGetEmbeddableURLWithOptionalParams(t *testing.T) {
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
	_, err := client.Dashboards.GetEmbeddableURL(context.TODO(), metronome.DashboardGetEmbeddableURLParams{
		CustomerID: metronome.F("4db51251-61de-4bfe-b9ce-495e244f3491"),
		Dashboard:  metronome.F(metronome.DashboardGetEmbeddableURLParamsDashboardInvoices),
		ColorOverrides: metronome.F([]metronome.DashboardGetEmbeddableURLParamsColorOverride{{
			Name:  metronome.F(metronome.DashboardGetEmbeddableURLParamsColorOverridesNameGrayDark),
			Value: metronome.F("#ff0000"),
		}}),
		DashboardOptions: metronome.F([]metronome.DashboardGetEmbeddableURLParamsDashboardOption{{
			Key:   metronome.F("string"),
			Value: metronome.F("string"),
		}, {
			Key:   metronome.F("string"),
			Value: metronome.F("string"),
		}, {
			Key:   metronome.F("string"),
			Value: metronome.F("string"),
		}}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
