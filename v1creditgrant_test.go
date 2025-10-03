// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Metronome-Industries/metronome-go/v2"
	"github.com/Metronome-Industries/metronome-go/v2/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/v2/option"
)

func TestV1CreditGrantNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.CreditGrants.New(context.TODO(), metronome.V1CreditGrantNewParams{
		CustomerID: "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		ExpiresAt:  time.Now(),
		GrantAmount: metronome.V1CreditGrantNewParamsGrantAmount{
			Amount:       1000,
			CreditTypeID: "5ae401dc-a648-4b49-9ac3-391bb5bc4d7b",
		},
		Name: "Acme Corp Promotional Credit Grant",
		PaidAmount: metronome.V1CreditGrantNewParamsPaidAmount{
			Amount:       5000,
			CreditTypeID: "2714e483-4ff1-48e4-9e25-ac732e8f24f2",
		},
		Priority:        0.5,
		CreditGrantType: metronome.String("trial"),
		CustomFields: map[string]string{
			"foo": "string",
		},
		EffectiveAt: metronome.Time(time.Now()),
		InvoiceDate: metronome.Time(time.Now()),
		ProductIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		Reason:      metronome.String("Incentivize new customer"),
		RolloverSettings: metronome.V1CreditGrantNewParamsRolloverSettings{
			ExpiresAt: time.Now(),
			Priority:  0,
			RolloverAmount: metronome.V1CreditGrantNewParamsRolloverSettingsRolloverAmountUnion{
				OfRolloverAmountMaxPercentage: &metronome.RolloverAmountMaxPercentageParam{
					Type:  metronome.RolloverAmountMaxPercentageTypeMaxPercentage,
					Value: 0,
				},
			},
		},
		UniquenessKey: metronome.String("x"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CreditGrantListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.CreditGrants.List(context.TODO(), metronome.V1CreditGrantListParams{
		Limit:             metronome.Int(1),
		NextPage:          metronome.String("next_page"),
		CreditGrantIDs:    []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		CreditTypeIDs:     []string{"2714e483-4ff1-48e4-9e25-ac732e8f24f2"},
		CustomerIDs:       []string{"d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc", "0e5b8609-d901-4992-b394-c3c2e3f37b1c"},
		EffectiveBefore:   metronome.Time(time.Now()),
		NotExpiringBefore: metronome.Time(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CreditGrantEditWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.CreditGrants.Edit(context.TODO(), metronome.V1CreditGrantEditParams{
		ID:              "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		CreditGrantType: metronome.String("credit_grant_type"),
		ExpiresAt:       metronome.Time(time.Now()),
		Name:            metronome.String("Acme Corp Promotional Credit Grant"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CreditGrantListEntriesWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.CreditGrants.ListEntries(context.TODO(), metronome.V1CreditGrantListEntriesParams{
		NextPage:      metronome.String("next_page"),
		Sort:          metronome.V1CreditGrantListEntriesParamsSortAsc,
		CreditTypeIDs: []string{"2714e483-4ff1-48e4-9e25-ac732e8f24f2"},
		CustomerIDs:   []string{"6a37bb88-8538-48c5-b37b-a41c836328bd"},
		EndingBefore:  metronome.Time(time.Now()),
		StartingOn:    metronome.Time(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CreditGrantVoidWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.CreditGrants.Void(context.TODO(), metronome.V1CreditGrantVoidParams{
		ID:                        "9b85c1c1-5238-4f2a-a409-61412905e1e1",
		ReleaseUniquenessKey:      metronome.Bool(true),
		VoidCreditPurchaseInvoice: metronome.Bool(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
