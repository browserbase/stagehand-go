// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stagehandsdk_test

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
	client := stagehandsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Act(
		context.TODO(),
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehandsdk.SessionActParams{
			Input: stagehandsdk.SessionActParamsInputUnion{
				OfString: stagehandsdk.String("Click the login button"),
			},
			FrameID: stagehandsdk.String("frameId"),
			Options: stagehandsdk.SessionActParamsOptions{
				Model: stagehandsdk.ModelConfigUnionParam{
					OfString: stagehandsdk.String("openai/gpt-4o"),
				},
				Timeout: stagehandsdk.Float(30000),
				Variables: map[string]string{
					"username": "john_doe",
				},
			},
			XStreamResponse: stagehandsdk.SessionActParamsXStreamResponseTrue,
		},
	)
	if err != nil {
		var apierr *stagehandsdk.Error
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
	client := stagehandsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.End(
		context.TODO(),
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehandsdk.SessionEndParams{
			ForceBody:       map[string]any{},
			XStreamResponse: stagehandsdk.SessionEndParamsXStreamResponseTrue,
		},
	)
	if err != nil {
		var apierr *stagehandsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionExecuteWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := stagehandsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Execute(
		context.TODO(),
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehandsdk.SessionExecuteParams{
			AgentConfig: stagehandsdk.SessionExecuteParamsAgentConfig{
				Cua: stagehandsdk.Bool(true),
				Model: stagehandsdk.ModelConfigUnionParam{
					OfString: stagehandsdk.String("openai/gpt-4o"),
				},
				Provider:     "openai",
				SystemPrompt: stagehandsdk.String("systemPrompt"),
			},
			ExecuteOptions: stagehandsdk.SessionExecuteParamsExecuteOptions{
				Instruction:     "Log in with username 'demo' and password 'test123', then navigate to settings",
				HighlightCursor: stagehandsdk.Bool(true),
				MaxSteps:        stagehandsdk.Float(20),
			},
			FrameID:         stagehandsdk.String("frameId"),
			XStreamResponse: stagehandsdk.SessionExecuteParamsXStreamResponseTrue,
		},
	)
	if err != nil {
		var apierr *stagehandsdk.Error
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
	client := stagehandsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Extract(
		context.TODO(),
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehandsdk.SessionExtractParams{
			FrameID:     stagehandsdk.String("frameId"),
			Instruction: stagehandsdk.String("Extract all product names and prices from the page"),
			Options: stagehandsdk.SessionExtractParamsOptions{
				Model: stagehandsdk.ModelConfigUnionParam{
					OfString: stagehandsdk.String("openai/gpt-4o"),
				},
				Selector: stagehandsdk.String("#main-content"),
				Timeout:  stagehandsdk.Float(30000),
			},
			Schema: map[string]any{
				"foo": "bar",
			},
			XStreamResponse: stagehandsdk.SessionExtractParamsXStreamResponseTrue,
		},
	)
	if err != nil {
		var apierr *stagehandsdk.Error
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
	client := stagehandsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Navigate(
		context.TODO(),
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehandsdk.SessionNavigateParams{
			URL:     "https://example.com",
			FrameID: stagehandsdk.String("frameId"),
			Options: stagehandsdk.SessionNavigateParamsOptions{
				Referer:   stagehandsdk.String("referer"),
				Timeout:   stagehandsdk.Float(30000),
				WaitUntil: "networkidle",
			},
			StreamResponse:  stagehandsdk.Bool(true),
			XStreamResponse: stagehandsdk.SessionNavigateParamsXStreamResponseTrue,
		},
	)
	if err != nil {
		var apierr *stagehandsdk.Error
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
	client := stagehandsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Observe(
		context.TODO(),
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehandsdk.SessionObserveParams{
			FrameID:     stagehandsdk.String("frameId"),
			Instruction: stagehandsdk.String("Find all clickable navigation links"),
			Options: stagehandsdk.SessionObserveParamsOptions{
				Model: stagehandsdk.ModelConfigUnionParam{
					OfString: stagehandsdk.String("openai/gpt-4o"),
				},
				Selector: stagehandsdk.String("nav"),
				Timeout:  stagehandsdk.Float(30000),
			},
			XStreamResponse: stagehandsdk.SessionObserveParamsXStreamResponseTrue,
		},
	)
	if err != nil {
		var apierr *stagehandsdk.Error
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
	client := stagehandsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Start(context.TODO(), stagehandsdk.SessionStartParams{
		ModelName:    "openai/gpt-4o",
		ActTimeoutMs: stagehandsdk.Float(0),
		Browser: stagehandsdk.SessionStartParamsBrowser{
			CdpURL: stagehandsdk.String("ws://localhost:9222"),
			LaunchOptions: stagehandsdk.SessionStartParamsBrowserLaunchOptions{
				AcceptDownloads:   stagehandsdk.Bool(true),
				Args:              []string{"string"},
				CdpURL:            stagehandsdk.String("cdpUrl"),
				ChromiumSandbox:   stagehandsdk.Bool(true),
				ConnectTimeoutMs:  stagehandsdk.Float(0),
				DeviceScaleFactor: stagehandsdk.Float(0),
				Devtools:          stagehandsdk.Bool(true),
				DownloadsPath:     stagehandsdk.String("downloadsPath"),
				ExecutablePath:    stagehandsdk.String("executablePath"),
				HasTouch:          stagehandsdk.Bool(true),
				Headless:          stagehandsdk.Bool(true),
				IgnoreDefaultArgs: stagehandsdk.SessionStartParamsBrowserLaunchOptionsIgnoreDefaultArgsUnion{
					OfBool: stagehandsdk.Bool(true),
				},
				IgnoreHTTPSErrors:   stagehandsdk.Bool(true),
				Locale:              stagehandsdk.String("locale"),
				PreserveUserDataDir: stagehandsdk.Bool(true),
				Proxy: stagehandsdk.SessionStartParamsBrowserLaunchOptionsProxy{
					Server:   "server",
					Bypass:   stagehandsdk.String("bypass"),
					Password: stagehandsdk.String("password"),
					Username: stagehandsdk.String("username"),
				},
				UserDataDir: stagehandsdk.String("userDataDir"),
				Viewport: stagehandsdk.SessionStartParamsBrowserLaunchOptionsViewport{
					Height: 0,
					Width:  0,
				},
			},
			Type: "local",
		},
		BrowserbaseSessionCreateParams: stagehandsdk.SessionStartParamsBrowserbaseSessionCreateParams{
			BrowserSettings: stagehandsdk.SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettings{
				AdvancedStealth: stagehandsdk.Bool(true),
				BlockAds:        stagehandsdk.Bool(true),
				Context: stagehandsdk.SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsContext{
					ID:      "id",
					Persist: stagehandsdk.Bool(true),
				},
				ExtensionID: stagehandsdk.String("extensionId"),
				Fingerprint: stagehandsdk.SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprint{
					Browsers:         []string{"chrome"},
					Devices:          []string{"desktop"},
					HTTPVersion:      "1",
					Locales:          []string{"string"},
					OperatingSystems: []string{"android"},
					Screen: stagehandsdk.SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprintScreen{
						MaxHeight: stagehandsdk.Float(0),
						MaxWidth:  stagehandsdk.Float(0),
						MinHeight: stagehandsdk.Float(0),
						MinWidth:  stagehandsdk.Float(0),
					},
				},
				LogSession:    stagehandsdk.Bool(true),
				RecordSession: stagehandsdk.Bool(true),
				SolveCaptchas: stagehandsdk.Bool(true),
				Viewport: stagehandsdk.SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsViewport{
					Height: stagehandsdk.Float(0),
					Width:  stagehandsdk.Float(0),
				},
			},
			ExtensionID: stagehandsdk.String("extensionId"),
			KeepAlive:   stagehandsdk.Bool(true),
			ProjectID:   stagehandsdk.String("projectId"),
			Proxies: stagehandsdk.SessionStartParamsBrowserbaseSessionCreateParamsProxiesUnion{
				OfBool: stagehandsdk.Bool(true),
			},
			Region:  "us-west-2",
			Timeout: stagehandsdk.Float(0),
			UserMetadata: map[string]any{
				"foo": "bar",
			},
		},
		BrowserbaseSessionID: stagehandsdk.String("browserbaseSessionID"),
		DomSettleTimeoutMs:   stagehandsdk.Float(5000),
		Experimental:         stagehandsdk.Bool(true),
		SelfHeal:             stagehandsdk.Bool(true),
		SystemPrompt:         stagehandsdk.String("systemPrompt"),
		Verbose:              1,
		WaitForCaptchaSolves: stagehandsdk.Bool(true),
		XStreamResponse:      stagehandsdk.SessionStartParamsXStreamResponseTrue,
	})
	if err != nil {
		var apierr *stagehandsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
