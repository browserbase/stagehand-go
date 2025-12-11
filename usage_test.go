// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stagehand_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/stagehand-go"
	"github.com/stainless-sdks/stagehand-go/internal/testutil"
	"github.com/stainless-sdks/stagehand-go/option"
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
		option.WithAPIKey("My API Key"),
	)
	response, err := client.Sessions.Start(context.TODO(), stagehand.SessionStartParams{
		Env: stagehand.SessionStartParamsEnvLocal,
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Available)
}
