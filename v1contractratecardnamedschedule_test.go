// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/option"
)

func TestV1ContractRateCardNamedScheduleGetWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.RateCards.NamedSchedules.Get(context.TODO(), metronome.V1ContractRateCardNamedScheduleGetParams{
		ContractID:   "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:   "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		ScheduleName: "my-schedule",
		CoveringDate: metronome.Time(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContractRateCardNamedScheduleUpdateWithOptionalParams(t *testing.T) {
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
	err := client.V1.Contracts.RateCards.NamedSchedules.Update(context.TODO(), metronome.V1ContractRateCardNamedScheduleUpdateParams{
		ContractID:   "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		CustomerID:   "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		ScheduleName: "my-schedule",
		StartingAt:   time.Now(),
		Value: map[string]interface{}{
			"my_key": "my_value",
		},
		EndingBefore: metronome.Time(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
