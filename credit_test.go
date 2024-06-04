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

func TestCreditNewGrantWithOptionalParams(t *testing.T) {
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
	_, err := client.Credits.NewGrant(context.TODO(), metronome.CreditNewGrantParams{
		CustomerID: metronome.F("9b85c1c1-5238-4f2a-a409-61412905e1e1"),
		ExpiresAt:  metronome.F(time.Now()),
		GrantAmount: metronome.F(metronome.CreditNewGrantParamsGrantAmount{
			Amount:       metronome.F(1000.000000),
			CreditTypeID: metronome.F("5ae401dc-a648-4b49-9ac3-391bb5bc4d7b"),
		}),
		Name: metronome.F("Acme Corp Promotional Credit Grant"),
		PaidAmount: metronome.F(metronome.CreditNewGrantParamsPaidAmount{
			Amount:       metronome.F(5000.000000),
			CreditTypeID: metronome.F("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
		}),
		Priority:        metronome.F(0.500000),
		CreditGrantType: metronome.F("trial"),
		CustomFields: metronome.F(map[string]string{
			"foo": "string",
		}),
		EffectiveAt: metronome.F(time.Now()),
		InvoiceDate: metronome.F(time.Now()),
		ProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		Reason:      metronome.F("Incentivize new customer"),
		RolloverSettings: metronome.F(metronome.CreditNewGrantParamsRolloverSettings{
			ExpiresAt: metronome.F(time.Now()),
			Priority:  metronome.F(0.000000),
			RolloverAmount: metronome.F[metronome.CreditNewGrantParamsRolloverSettingsRolloverAmountUnion](metronome.CreditNewGrantParamsRolloverSettingsRolloverAmountObject{
				Type:  metronome.F(metronome.CreditNewGrantParamsRolloverSettingsRolloverAmountObjectTypeMaxPercentage),
				Value: metronome.F(0.000000),
			}),
		}),
		UniquenessKey: metronome.F("x"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreditEditGrantWithOptionalParams(t *testing.T) {
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
	_, err := client.Credits.EditGrant(context.TODO(), metronome.CreditEditGrantParams{
		ID:        metronome.F("9b85c1c1-5238-4f2a-a409-61412905e1e1"),
		ExpiresAt: metronome.F(time.Now()),
		Name:      metronome.F("Acme Corp Promotional Credit Grant"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreditListEntriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Credits.ListEntries(context.TODO(), metronome.CreditListEntriesParams{
		NextPage:      metronome.F("string"),
		CreditTypeIDs: metronome.F([]string{"2714e483-4ff1-48e4-9e25-ac732e8f24f2"}),
		CustomerIDs:   metronome.F([]string{"6a37bb88-8538-48c5-b37b-a41c836328bd"}),
		EndingBefore:  metronome.F(time.Now()),
		StartingOn:    metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreditListGrantsWithOptionalParams(t *testing.T) {
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
	_, err := client.Credits.ListGrants(context.TODO(), metronome.CreditListGrantsParams{
		NextPage:          metronome.F("string"),
		CreditGrantIDs:    metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		CreditTypeIDs:     metronome.F([]string{"2714e483-4ff1-48e4-9e25-ac732e8f24f2"}),
		CustomerIDs:       metronome.F([]string{"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc", "0e5b8609-d901-4992-b394-c3c2e3f37b1c"}),
		EffectiveBefore:   metronome.F(time.Now()),
		NotExpiringBefore: metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreditVoidGrantWithOptionalParams(t *testing.T) {
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
	_, err := client.Credits.VoidGrant(context.TODO(), metronome.CreditVoidGrantParams{
		ID:                        metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		VoidCreditPurchaseInvoice: metronome.F(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
