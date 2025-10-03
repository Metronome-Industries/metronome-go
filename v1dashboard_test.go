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

func TestV1DashboardGetEmbeddableURLWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Dashboards.GetEmbeddableURL(context.TODO(), metronome.V1DashboardGetEmbeddableURLParams{
		CustomerID: "4db51251-61de-4bfe-b9ce-495e244f3491",
		Dashboard:  metronome.V1DashboardGetEmbeddableURLParamsDashboardInvoices,
		BmGroupKeyOverrides: []metronome.V1DashboardGetEmbeddableURLParamsBmGroupKeyOverride{{
			GroupKeyName: "tenant_id",
			DisplayName:  metronome.String("Org ID"),
			ValueDisplayNames: map[string]any{
				"48ecb18f358f": "bar",
				"e358f3ce242d": "bar",
			},
		}},
		ColorOverrides: []metronome.V1DashboardGetEmbeddableURLParamsColorOverride{{
			Name:  "Gray_dark",
			Value: metronome.String("#ff0000"),
		}},
		DashboardOptions: []metronome.V1DashboardGetEmbeddableURLParamsDashboardOption{{
			Key:   "show_zero_usage_line_items",
			Value: "false",
		}, {
			Key:   "hide_voided_invoices",
			Value: "true",
		}},
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
