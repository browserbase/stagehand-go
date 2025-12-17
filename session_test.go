// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stagehand_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/browserbase/stagehand-go"
	"github.com/browserbase/stagehand-go/internal/testutil"
	"github.com/browserbase/stagehand-go/option"
)

func TestSessionStartWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := stagehand.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Start(context.TODO(), stagehand.SessionStartParams{
		Body:            map[string]any{},
		XLanguage:       map[string]any{},
		XSDKVersion:     map[string]any{},
		XSentAt:         map[string]any{},
		XStreamResponse: map[string]any{},
	})
	if err != nil {
		var apierr *stagehand.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
