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
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Act(
		context.TODO(),
		map[string]any{},
		stagehand.SessionActParams{
			Body:            map[string]any{},
			XLanguage:       map[string]any{},
			XSDKVersion:     map[string]any{},
			XSentAt:         map[string]any{},
			XStreamResponse: map[string]any{},
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

func TestSessionEndWithOptionalParams(t *testing.T) {
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
	_, err := client.Sessions.End(
		context.TODO(),
		map[string]any{},
		stagehand.SessionEndParams{
			XLanguage:       map[string]any{},
			XSDKVersion:     map[string]any{},
			XSentAt:         map[string]any{},
			XStreamResponse: map[string]any{},
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
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.ExecuteAgent(
		context.TODO(),
		map[string]any{},
		stagehand.SessionExecuteAgentParams{
			Body:            map[string]any{},
			XLanguage:       map[string]any{},
			XSDKVersion:     map[string]any{},
			XSentAt:         map[string]any{},
			XStreamResponse: map[string]any{},
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
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Extract(
		context.TODO(),
		map[string]any{},
		stagehand.SessionExtractParams{
			Body:            map[string]any{},
			XLanguage:       map[string]any{},
			XSDKVersion:     map[string]any{},
			XSentAt:         map[string]any{},
			XStreamResponse: map[string]any{},
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
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Navigate(
		context.TODO(),
		map[string]any{},
		stagehand.SessionNavigateParams{
			Body:            map[string]any{},
			XLanguage:       map[string]any{},
			XSDKVersion:     map[string]any{},
			XSentAt:         map[string]any{},
			XStreamResponse: map[string]any{},
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
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Observe(
		context.TODO(),
		map[string]any{},
		stagehand.SessionObserveParams{
			Body:            map[string]any{},
			XLanguage:       map[string]any{},
			XSDKVersion:     map[string]any{},
			XSentAt:         map[string]any{},
			XStreamResponse: map[string]any{},
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
