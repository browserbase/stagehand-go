// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stagehand_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/stagehand-go"
	"github.com/stainless-sdks/stagehand-go/internal/testutil"
	"github.com/stainless-sdks/stagehand-go/option"
)

func TestSessionActWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.Act(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		stagehand.SessionActParams{
			Input: stagehand.SessionActParamsInputUnion{
				OfString: stagehand.String("click the sign in button"),
			},
			FrameID: stagehand.String("frameId"),
			Options: stagehand.SessionActParamsOptions{
				Model: stagehand.ModelConfigParam{
					APIKey:   stagehand.String("apiKey"),
					BaseURL:  stagehand.String("https://example.com"),
					Model:    stagehand.String("model"),
					Provider: stagehand.ModelConfigProviderOpenAI,
				},
				Timeout: stagehand.Int(0),
				Variables: map[string]string{
					"foo": "string",
				},
			},
			XStreamResponse: stagehand.SessionActParamsXStreamResponseTrue,
		},
	)
	if err != nil {
		var apierr *stagehand.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionEnd(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.End(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *stagehand.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionExecuteAgentWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.ExecuteAgent(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		stagehand.SessionExecuteAgentParams{
			AgentConfig: stagehand.SessionExecuteAgentParamsAgentConfig{
				Cua: stagehand.Bool(true),
				Model: stagehand.SessionExecuteAgentParamsAgentConfigModelUnion{
					OfString: stagehand.String("openai/gpt-4o"),
				},
				Provider:     "openai",
				SystemPrompt: stagehand.String("systemPrompt"),
			},
			ExecuteOptions: stagehand.SessionExecuteAgentParamsExecuteOptions{
				Instruction:     "Find and click the first product",
				HighlightCursor: stagehand.Bool(true),
				MaxSteps:        stagehand.Int(10),
			},
			FrameID:         stagehand.String("frameId"),
			XStreamResponse: stagehand.SessionExecuteAgentParamsXStreamResponseTrue,
		},
	)
	if err != nil {
		var apierr *stagehand.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionExtractWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.Extract(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		stagehand.SessionExtractParams{
			FrameID:     stagehand.String("frameId"),
			Instruction: stagehand.String("extract the page title"),
			Options: stagehand.SessionExtractParamsOptions{
				Model: stagehand.ModelConfigParam{
					APIKey:   stagehand.String("apiKey"),
					BaseURL:  stagehand.String("https://example.com"),
					Model:    stagehand.String("model"),
					Provider: stagehand.ModelConfigProviderOpenAI,
				},
				Selector: stagehand.String("selector"),
				Timeout:  stagehand.Int(0),
			},
			Schema: map[string]any{
				"foo": "bar",
			},
			XStreamResponse: stagehand.SessionExtractParamsXStreamResponseTrue,
		},
	)
	if err != nil {
		var apierr *stagehand.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionNavigateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.Navigate(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		stagehand.SessionNavigateParams{
			URL:     "https://example.com",
			FrameID: stagehand.String("frameId"),
			Options: stagehand.SessionNavigateParamsOptions{
				WaitUntil: "load",
			},
			XStreamResponse: stagehand.SessionNavigateParamsXStreamResponseTrue,
		},
	)
	if err != nil {
		var apierr *stagehand.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionObserveWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.Observe(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		stagehand.SessionObserveParams{
			FrameID:     stagehand.String("frameId"),
			Instruction: stagehand.String("instruction"),
			Options: stagehand.SessionObserveParamsOptions{
				Model: stagehand.ModelConfigParam{
					APIKey:   stagehand.String("apiKey"),
					BaseURL:  stagehand.String("https://example.com"),
					Model:    stagehand.String("model"),
					Provider: stagehand.ModelConfigProviderOpenAI,
				},
				Selector: stagehand.String("selector"),
				Timeout:  stagehand.Int(0),
			},
			XStreamResponse: stagehand.SessionObserveParamsXStreamResponseTrue,
		},
	)
	if err != nil {
		var apierr *stagehand.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.Start(context.TODO(), stagehand.SessionStartParams{
		Env:              stagehand.SessionStartParamsEnvLocal,
		APIKey:           stagehand.String("apiKey"),
		DomSettleTimeout: stagehand.Int(0),
		LocalBrowserLaunchOptions: stagehand.SessionStartParamsLocalBrowserLaunchOptions{
			Headless: stagehand.Bool(true),
		},
		Model:        stagehand.String("openai/gpt-4o"),
		ProjectID:    stagehand.String("projectId"),
		SelfHeal:     stagehand.Bool(true),
		SystemPrompt: stagehand.String("systemPrompt"),
		Verbose:      stagehand.Int(1),
	})
	if err != nil {
		var apierr *stagehand.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
