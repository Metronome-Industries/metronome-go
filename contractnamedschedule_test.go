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

func TestContractNamedScheduleGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.NamedSchedules.Get(context.TODO(), metronome.ContractNamedScheduleGetParams{
		RateCardID:   metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		ScheduleName: metronome.F("my-schedule"),
		CoveringDate: metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractNamedScheduleUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Contracts.NamedSchedules.Update(context.TODO(), metronome.ContractNamedScheduleUpdateParams{
		RateCardID:   metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		ScheduleName: metronome.F("my-schedule"),
		StartingAt:   metronome.F(time.Now()),
		Value: metronome.F[any](map[string]interface{}{
			"my_key": "my_value",
		}),
		EndingBefore: metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}