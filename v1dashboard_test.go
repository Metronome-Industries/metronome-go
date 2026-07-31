// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/v3/option"
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
			GroupKeyName: "group_key_name",
			DisplayName:  metronome.String("display_name"),
			ValueDisplayNames: map[string]any{
				"foo": "bar",
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
			Key:   "invoice_status_filter",
			Value: "FINALIZED",
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
