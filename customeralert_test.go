// File generated from our OpenAPI spec by Stainless.

package metronome_test

import (
  "testing"
  "context"
  "net/http/httputil"
  "github.com/metronome/metronome-go"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/testutil"
)

func TestCustomerAlertGet(t *testing.T) {
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
  _, err := client.CustomerAlerts.Get(context.TODO(), metronome.CustomerAlertGetParams{
    AlertID: metronome.F("8deed800-1b7a-495d-a207-6c52bac54dc9"),
    CustomerID: metronome.F("9b85c1c1-5238-4f2a-a409-61412905e1e1"),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestCustomerAlertListWithOptionalParams(t *testing.T) {
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
  _, err := client.CustomerAlerts.List(context.TODO(), metronome.CustomerAlertListParams{
    CustomerID: metronome.F("9b85c1c1-5238-4f2a-a409-61412905e1e1"),
    NextPage: metronome.F("string"),
    AlertStatuses: metronome.F([]metronome.CustomerAlertListParamsAlertStatus{metronome.CustomerAlertListParamsAlertStatusEnabled}),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}