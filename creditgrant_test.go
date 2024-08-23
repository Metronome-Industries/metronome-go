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

func TestCreditGrantNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditGrants.New(context.TODO(), metronome.CreditGrantNewParams{
		CustomerID: metronome.F("9b85c1c1-5238-4f2a-a409-61412905e1e1"),
		ExpiresAt:  metronome.F(time.Now()),
		GrantAmount: metronome.F(metronome.CreditGrantNewParamsGrantAmount{
			Amount:       metronome.F(1000.000000),
			CreditTypeID: metronome.F("5ae401dc-a648-4b49-9ac3-391bb5bc4d7b"),
		}),
		Name: metronome.F("Acme Corp Promotional Credit Grant"),
		PaidAmount: metronome.F(metronome.CreditGrantNewParamsPaidAmount{
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
		RolloverSettings: metronome.F(metronome.CreditGrantNewParamsRolloverSettings{
			ExpiresAt: metronome.F(time.Now()),
			Priority:  metronome.F(0.000000),
			RolloverAmount: metronome.F[metronome.CreditGrantNewParamsRolloverSettingsRolloverAmountUnion](metronome.RolloverAmountMaxPercentageParam{
				Type:  metronome.F(metronome.RolloverAmountMaxPercentageTypeMaxPercentage),
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

func TestCreditGrantListWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditGrants.List(context.TODO(), metronome.CreditGrantListParams{
		Limit:             metronome.F(int64(1)),
		NextPage:          metronome.F("next_page"),
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

func TestCreditGrantEditWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditGrants.Edit(context.TODO(), metronome.CreditGrantEditParams{
		ID:              metronome.F("9b85c1c1-5238-4f2a-a409-61412905e1e1"),
		CreditGrantType: metronome.F("credit_grant_type"),
		ExpiresAt:       metronome.F(time.Now()),
		Name:            metronome.F("Acme Corp Promotional Credit Grant"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreditGrantListCreditTypesWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditGrants.ListCreditTypes(context.TODO(), metronome.CreditGrantListCreditTypesParams{
		Limit:    metronome.F(int64(1)),
		NextPage: metronome.F("next_page"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreditGrantListEntriesWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditGrants.ListEntries(context.TODO(), metronome.CreditGrantListEntriesParams{
		NextPage:      metronome.F("next_page"),
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

func TestCreditGrantVoidWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditGrants.Void(context.TODO(), metronome.CreditGrantVoidParams{
		ID:                        metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ReleaseUniquenessKey:      metronome.F(true),
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