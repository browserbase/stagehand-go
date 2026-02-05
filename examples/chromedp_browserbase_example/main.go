// Example: chromedp + Browserbase (cloud mode).
//
// Prerequisites:
//   - Set BROWSERBASE_API_KEY
//   - Set BROWSERBASE_PROJECT_ID
//   - Set MODEL_API_KEY
//
// Run:
//   cd examples/chromedp_browserbase_example
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
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func main() {
	// Environment variables required (same as other examples):
	// - BROWSERBASE_API_KEY
	// - BROWSERBASE_PROJECT_ID
	// - MODEL_API_KEY
	client := stagehand.NewClient()

	// 1) Start a Stagehand session and get the Browserbase CDP URL.
	startResponse, err := client.Sessions.Start(context.TODO(), stagehand.SessionStartParams{
		ModelName: "openai/gpt-5-nano",
		Browser: stagehand.SessionStartParamsBrowser{
			Type: "browserbase",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	sessionID := startResponse.Data.SessionID
	cdpURL := startResponse.Data.CdpURL
	if cdpURL == "" {
		log.Fatalf("start response missing cdpUrl (sessionID=%s); cannot attach chromedp", sessionID)
	}
	fmt.Printf("Session started: %s\n", sessionID)
	fmt.Printf("CDP URL: %s\n", cdpURL)

	// Patch CDP URL to add port if missing (chromedp requires explicit port)
	cdpURL = ensurePort(cdpURL)
	fmt.Printf("CDP URL (patched): %s\n", cdpURL)

	// 2) Navigate with Stagehand so we can attach chromedp to the existing tab.
	_, err = client.Sessions.Navigate(
		context.TODO(),
		sessionID,
		stagehand.SessionNavigateParams{
			URL: "https://example.com",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// 3) Connect chromedp to the same browser over CDP.
	// Use NoModifyURL to skip the /json/version fetch that Browserbase doesn't support.
	allocatorCtx, cancelAllocator := chromedp.NewRemoteAllocator(context.Background(), cdpURL, chromedp.NoModifyURL)
	defer cancelAllocator()

	// Suppress CDP protocol unmarshal errors (version mismatch warnings)
	browserCtx, cancelBrowser := chromedp.NewContext(allocatorCtx,
		chromedp.WithErrorf(func(format string, args ...interface{}) {}),
	)
	defer cancelBrowser()

	targetID, err := waitForTarget(browserCtx, "example.com", 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	tabCtx, cancelTab := chromedp.NewContext(browserCtx, chromedp.WithTargetID(targetID))
	defer cancelTab()

	// 4) Use chromedp to click a link in the same tab.
	err = chromedp.Run(
		tabCtx,
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.Click("a", chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Chromedp click completed")

	// 5) Use Stagehand methods without a frameId (defaults to active tab).
	observeResponse, err := client.Sessions.Observe(
		context.TODO(),
		sessionID,
		stagehand.SessionObserveParams{
			Instruction: stagehand.String("Find the most relevant click target on this page"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Observed %d possible actions\n", len(observeResponse.Data.Result))

	// 6) Demonstrate using Stagehand to click something in the same tab/frame.
	_, err = client.Sessions.Act(
		context.TODO(),
		sessionID,
		stagehand.SessionActParams{
			Input: stagehand.SessionActParamsInputUnion{
				OfString: stagehand.String("Click on the 'Learn more' link"),
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Give navigation a moment to settle.
	time.Sleep(2 * time.Second)

	// 7) Extract from the same frame.
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string"},
			"url":   map[string]any{"type": "string"},
		},
	}
	extractResponse, err := client.Sessions.Extract(
		context.TODO(),
		sessionID,
		stagehand.SessionExtractParams{
			Instruction: stagehand.String("Extract the page title and current URL"),
			Schema:      schema,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Extracted: %+v\n", extractResponse.Data.Result)

	// 8) Run an autonomous agent using Stagehand Execute.
	fmt.Println("Running autonomous agent...")
	executeResponse, err := client.Sessions.Execute(
		context.TODO(),
		sessionID,
		stagehand.SessionExecuteParams{
			ExecuteOptions: stagehand.SessionExecuteParamsExecuteOptions{
				Instruction: "Find and click on any link related to 'RFC' or 'specifications' on this page. " +
					"If no such link exists, find any other interesting link to click.",
				MaxSteps: stagehand.Float(5),
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
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Agent completed: %s\n", executeResponse.Data.Result.Message)
	fmt.Printf("Agent success: %t\n", executeResponse.Data.Result.Success)
	fmt.Printf("Agent actions taken: %d\n", len(executeResponse.Data.Result.Actions))

	// 9) Use chromedp to take a screenshot after the agent finishes.
	fmt.Println("Taking screenshot with chromedp...")
	var screenshotBuf []byte
	err = chromedp.Run(
		tabCtx,
		chromedp.Sleep(1*time.Second), // Allow page to settle
		chromedp.FullScreenshot(&screenshotBuf, 90),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Save the screenshot to a file
	screenshotPath := "screenshot.png"
	if err := os.WriteFile(screenshotPath, screenshotBuf, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Screenshot saved to: %s\n", screenshotPath)

	// 10) Print session metrics including log of all actions taken and LLM tokens used for each.
	replayResponse, err := client.Sessions.Replay(
		context.TODO(),
		sessionID,
		stagehand.SessionReplayParams{},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", replayResponse.Data)

	// 11) End session.
	_, err = client.Sessions.End(context.TODO(), sessionID, stagehand.SessionEndParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Session ended")
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
