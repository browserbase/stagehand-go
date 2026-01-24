// Example: Local mode + Browserbase region + chromedp.
//
// Prerequisites:
//   - Set BROWSERBASE_API_KEY
//   - Set BROWSERBASE_PROJECT_ID
//   - Set MODEL_API_KEY
//
// Run:
//   cd examples/chromedp_multiregion_example
//   go mod download
//   go run main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/browserbase/stagehand-go/v3"
	"github.com/browserbase/stagehand-go/v3/option"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func main() {
	requireEnv("BROWSERBASE_API_KEY", "BROWSERBASE_PROJECT_ID", "MODEL_API_KEY")

	// Run the Stagehand driver locally (required for Browserbase regions other than us-west-2)
	client := stagehand.NewClient(option.WithServer("local"))
	defer client.Close()

	ctx := context.Background()

	startResp, err := client.Sessions.Start(ctx, stagehand.SessionStartParams{
		ModelName: "openai/gpt-5-nano",
		Browser: stagehand.SessionStartParamsBrowser{
			Type: "browserbase",
		},
		BrowserbaseSessionCreateParams: stagehand.SessionStartParamsBrowserbaseSessionCreateParams{
			Region: "eu-central-1",
		},
	})
	if err != nil {
		log.Fatalf("Failed to start session: %v", err)
	}
	sessionID := startResp.Data.SessionID
	cdpURL := startResp.Data.CdpURL
	if cdpURL == "" {
		log.Fatalf("Start response missing cdpUrl (sessionID=%s)", sessionID)
	}

	_, err = client.Sessions.Navigate(ctx, sessionID, stagehand.SessionNavigateParams{
		URL: "https://example.com",
	})
	if err != nil {
		log.Fatalf("Failed to navigate via Stagehand: %v", err)
	}

	cdpURL = ensurePort(cdpURL)

	// Connect ChromeDP to the main tab in the remote browser
	allocatorCtx, cancelAllocator := chromedp.NewRemoteAllocator(context.Background(), cdpURL, chromedp.NoModifyURL)
	defer cancelAllocator()

	browserCtx, cancelBrowser := chromedp.NewContext(allocatorCtx, chromedp.WithErrorf(func(format string, args ...interface{}) {}))
	defer cancelBrowser()

	targetID, err := waitForTarget(browserCtx, "example.com", 10*time.Second)
	if err != nil {
		log.Fatalf("Failed to find existing tab: %v", err)
	}

	tabCtx, cancelTab := chromedp.NewContext(browserCtx, chromedp.WithTargetID(targetID))
	defer cancelTab()

	if err := takeScreenshot(tabCtx, "screenshot_multiregion_start.png"); err != nil {
		log.Fatalf("Failed to take start screenshot: %v", err)
	}

	observeResp, err := client.Sessions.Observe(ctx, sessionID, stagehand.SessionObserveParams{
		Instruction: stagehand.String("Find all clickable links on this page"),
	})
	if err != nil {
		log.Fatalf("Failed to observe: %v", err)
	}
	fmt.Printf("Observed %d possible actions\n", len(observeResp.Data.Result))

	extractResp, err := client.Sessions.Extract(ctx, sessionID, stagehand.SessionExtractParams{
		Instruction: stagehand.String("Extract the page title and current URL"),
		Schema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"title": map[string]any{"type": "string"},
				"url":   map[string]any{"type": "string"},
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to extract: %v", err)
	}
	fmt.Printf("Extracted: %+v\n", extractResp.Data.Result)

	executeResp, err := client.Sessions.Execute(ctx, sessionID, stagehand.SessionExecuteParams{
		ExecuteOptions: stagehand.SessionExecuteParamsExecuteOptions{
			Instruction: "Make sure the page is loaded, then find the first link on the page, labeled Learn more, click it and wait for navigation, then scroll down partway",
			MaxSteps:    stagehand.Float(3),
		},
		AgentConfig: stagehand.SessionExecuteParamsAgentConfig{
			Model: stagehand.SessionExecuteParamsAgentConfigModelUnion{
				OfModelConfig: &stagehand.ModelConfigParam{
					ModelName: "openai/gpt-5-nano",
					APIKey:    stagehand.String(os.Getenv("MODEL_API_KEY")),
				},
			},
			Cua: stagehand.Bool(false),
		},
	})
	if err != nil {
		log.Fatalf("Failed to execute: %v", err)
	}
	fmt.Printf("Agent completed: %s\n", executeResp.Data.Result.Message)

	if err := takeScreenshot(tabCtx, "screenshot_multiregion_end.png"); err != nil {
		log.Fatalf("Failed to take end screenshot: %v", err)
	}

	_, err = client.Sessions.End(ctx, sessionID, stagehand.SessionEndParams{})
	if err != nil {
		log.Fatalf("Failed to end session: %v", err)
	}
	fmt.Println("Session ended")
}

func requireEnv(names ...string) {
	for _, name := range names {
		if os.Getenv(name) == "" {
			log.Fatalf("Missing %s", name)
		}
	}
}

// ensurePort adds the default port to a WebSocket URL if missing.
// chromedp requires an explicit port in the URL.
func ensurePort(wsURL string) string {
	u, err := url.Parse(wsURL)
	if err != nil {
		return wsURL
	}
	if u.Port() == "" {
		switch u.Scheme {
		case "wss":
			u.Host = u.Hostname() + ":443"
		case "ws":
			u.Host = u.Hostname() + ":80"
		}
	}
	return u.String()
}

func takeScreenshot(ctx context.Context, path string) error {
	var screenshotBuf []byte
	if err := chromedp.Run(
		ctx,
		chromedp.Sleep(1*time.Second),
		chromedp.FullScreenshot(&screenshotBuf, 90),
	); err != nil {
		return err
	}
	return os.WriteFile(path, screenshotBuf, 0644)
}

func waitForTarget(ctx context.Context, urlSubstring string, timeout time.Duration) (target.ID, error) {
	deadline := time.Now().Add(timeout)
	for {
		targets, err := chromedp.Targets(ctx)
		if err != nil {
			return "", err
		}
		for _, t := range targets {
			if t.Type == "page" && strings.Contains(t.URL, urlSubstring) {
				return t.TargetID, nil
			}
		}
		if time.Now().After(deadline) {
			return "", fmt.Errorf("no page target with URL containing %q", urlSubstring)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
