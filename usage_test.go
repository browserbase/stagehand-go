// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stagehand_test

import (
	"context"
	"os"
	"testing"

	"github.com/browserbase/stagehand-go"
	"github.com/browserbase/stagehand-go/internal/testutil"
	"github.com/browserbase/stagehand-go/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Sessions.Act(
		context.TODO(),
		"00000000-your-session-id-000000000000",
		stagehand.SessionActParams{
			Input: stagehand.SessionActParamsInputUnion{
				OfString: stagehand.String("click the first link on the page"),
			},
		},
	)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Actions)
}
