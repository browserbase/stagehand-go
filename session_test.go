// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stagehand_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

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
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehand.SessionActParams{
			Input: stagehand.SessionActParamsInputUnion{
				OfString: stagehand.String("Click the login button"),
			},
			FrameID: stagehand.String("frameId"),
			Options: stagehand.SessionActParamsOptions{
				Model: stagehand.ModelConfigUnionParam{
					OfString: stagehand.String("openai/gpt-5-nano"),
				},
				Timeout: stagehand.Float(30000),
				Variables: map[string]string{
					"username": "john_doe",
				},
			},
			XLanguage:       stagehand.SessionActParamsXLanguageTypescript,
			XSDKVersion:     stagehand.String("3.0.6"),
			XSentAt:         stagehand.Time(time.Now()),
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
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehand.SessionEndParams{
			XLanguage:       stagehand.SessionEndParamsXLanguageTypescript,
			XSDKVersion:     stagehand.String("3.0.6"),
			XSentAt:         stagehand.Time(time.Now()),
			XStreamResponse: stagehand.SessionEndParamsXStreamResponseTrue,
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

func TestSessionExecuteWithOptionalParams(t *testing.T) {
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
	_, err := client.Sessions.Execute(
		context.TODO(),
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehand.SessionExecuteParams{
			AgentConfig: stagehand.SessionExecuteParamsAgentConfig{
				Cua: stagehand.Bool(true),
				Model: stagehand.ModelConfigUnionParam{
					OfString: stagehand.String("openai/gpt-5-nano"),
				},
				Provider:     "openai",
				SystemPrompt: stagehand.String("systemPrompt"),
			},
			ExecuteOptions: stagehand.SessionExecuteParamsExecuteOptions{
				Instruction:     "Log in with username 'demo' and password 'test123', then navigate to settings",
				HighlightCursor: stagehand.Bool(true),
				MaxSteps:        stagehand.Float(20),
			},
			FrameID:         stagehand.String("frameId"),
			XLanguage:       stagehand.SessionExecuteParamsXLanguageTypescript,
			XSDKVersion:     stagehand.String("3.0.6"),
			XSentAt:         stagehand.Time(time.Now()),
			XStreamResponse: stagehand.SessionExecuteParamsXStreamResponseTrue,
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
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehand.SessionExtractParams{
			FrameID:     stagehand.String("frameId"),
			Instruction: stagehand.String("Extract all product names and prices from the page"),
			Options: stagehand.SessionExtractParamsOptions{
				Model: stagehand.ModelConfigUnionParam{
					OfString: stagehand.String("openai/gpt-5-nano"),
				},
				Selector: stagehand.String("#main-content"),
				Timeout:  stagehand.Float(30000),
			},
			Schema: map[string]any{
				"foo": "bar",
			},
			XLanguage:       stagehand.SessionExtractParamsXLanguageTypescript,
			XSDKVersion:     stagehand.String("3.0.6"),
			XSentAt:         stagehand.Time(time.Now()),
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
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Navigate(
		context.TODO(),
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehand.SessionNavigateParams{
			URL:     "https://example.com",
			FrameID: stagehand.String("frameId"),
			Options: stagehand.SessionNavigateParamsOptions{
				Referer:   stagehand.String("referer"),
				Timeout:   stagehand.Float(30000),
				WaitUntil: "networkidle",
			},
			StreamResponse:  stagehand.Bool(true),
			XLanguage:       stagehand.SessionNavigateParamsXLanguageTypescript,
			XSDKVersion:     stagehand.String("3.0.6"),
			XSentAt:         stagehand.Time(time.Now()),
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
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Observe(
		context.TODO(),
		"c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		stagehand.SessionObserveParams{
			FrameID:     stagehand.String("frameId"),
			Instruction: stagehand.String("Find all clickable navigation links"),
			Options: stagehand.SessionObserveParamsOptions{
				Model: stagehand.ModelConfigUnionParam{
					OfString: stagehand.String("openai/gpt-5-nano"),
				},
				Selector: stagehand.String("nav"),
				Timeout:  stagehand.Float(30000),
			},
			XLanguage:       stagehand.SessionObserveParamsXLanguageTypescript,
			XSDKVersion:     stagehand.String("3.0.6"),
			XSentAt:         stagehand.Time(time.Now()),
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
		option.WithBrowserbaseAPIKey("My Browserbase API Key"),
		option.WithBrowserbaseProjectID("My Browserbase Project ID"),
		option.WithModelAPIKey("My Model API Key"),
	)
	_, err := client.Sessions.Start(context.TODO(), stagehand.SessionStartParams{
		ModelName:    "gpt-4o",
		ActTimeoutMs: stagehand.Float(0),
		Browser: stagehand.SessionStartParamsBrowser{
			CdpURL: stagehand.String("ws://localhost:9222"),
			LaunchOptions: stagehand.SessionStartParamsBrowserLaunchOptions{
				AcceptDownloads:   stagehand.Bool(true),
				Args:              []string{"string"},
				CdpURL:            stagehand.String("cdpUrl"),
				ChromiumSandbox:   stagehand.Bool(true),
				ConnectTimeoutMs:  stagehand.Float(0),
				DeviceScaleFactor: stagehand.Float(0),
				Devtools:          stagehand.Bool(true),
				DownloadsPath:     stagehand.String("downloadsPath"),
				ExecutablePath:    stagehand.String("executablePath"),
				HasTouch:          stagehand.Bool(true),
				Headless:          stagehand.Bool(true),
				IgnoreDefaultArgs: stagehand.SessionStartParamsBrowserLaunchOptionsIgnoreDefaultArgsUnion{
					OfBool: stagehand.Bool(true),
				},
				IgnoreHTTPSErrors:   stagehand.Bool(true),
				Locale:              stagehand.String("locale"),
				PreserveUserDataDir: stagehand.Bool(true),
				Proxy: stagehand.SessionStartParamsBrowserLaunchOptionsProxy{
					Server:   "server",
					Bypass:   stagehand.String("bypass"),
					Password: stagehand.String("password"),
					Username: stagehand.String("username"),
				},
				UserDataDir: stagehand.String("userDataDir"),
				Viewport: stagehand.SessionStartParamsBrowserLaunchOptionsViewport{
					Height: 0,
					Width:  0,
				},
			},
			Type: "local",
		},
		BrowserbaseSessionCreateParams: stagehand.SessionStartParamsBrowserbaseSessionCreateParams{
			BrowserSettings: stagehand.SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettings{
				AdvancedStealth: stagehand.Bool(true),
				BlockAds:        stagehand.Bool(true),
				Context: stagehand.SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsContext{
					ID:      "id",
					Persist: stagehand.Bool(true),
				},
				ExtensionID: stagehand.String("extensionId"),
				Fingerprint: stagehand.SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprint{
					Browsers:         []string{"chrome"},
					Devices:          []string{"desktop"},
					HTTPVersion:      "1",
					Locales:          []string{"string"},
					OperatingSystems: []string{"android"},
					Screen: stagehand.SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprintScreen{
						MaxHeight: stagehand.Float(0),
						MaxWidth:  stagehand.Float(0),
						MinHeight: stagehand.Float(0),
						MinWidth:  stagehand.Float(0),
					},
				},
				LogSession:    stagehand.Bool(true),
				RecordSession: stagehand.Bool(true),
				SolveCaptchas: stagehand.Bool(true),
				Viewport: stagehand.SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsViewport{
					Height: stagehand.Float(0),
					Width:  stagehand.Float(0),
				},
			},
			ExtensionID: stagehand.String("extensionId"),
			KeepAlive:   stagehand.Bool(true),
			ProjectID:   stagehand.String("projectId"),
			Proxies: stagehand.SessionStartParamsBrowserbaseSessionCreateParamsProxiesUnion{
				OfBool: stagehand.Bool(true),
			},
			Region:  "us-west-2",
			Timeout: stagehand.Float(0),
			UserMetadata: map[string]any{
				"foo": "bar",
			},
		},
		BrowserbaseSessionID: stagehand.String("browserbaseSessionID"),
		DomSettleTimeoutMs:   stagehand.Float(5000),
		Experimental:         stagehand.Bool(true),
		SelfHeal:             stagehand.Bool(true),
		SystemPrompt:         stagehand.String("systemPrompt"),
		Verbose:              stagehand.SessionStartParamsVerbose1,
		WaitForCaptchaSolves: stagehand.Bool(true),
		XLanguage:            stagehand.SessionStartParamsXLanguageTypescript,
		XSDKVersion:          stagehand.String("3.0.6"),
		XSentAt:              stagehand.Time(time.Now()),
		XStreamResponse:      stagehand.SessionStartParamsXStreamResponseTrue,
	})
	if err != nil {
		var apierr *stagehand.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
