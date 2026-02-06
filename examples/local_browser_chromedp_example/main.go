// Example: chromedp + local browser (SSE streaming by default).
//
// Prerequisites:
//   - Set MODEL_API_KEY
//   - Chrome/Chromium installed locally
//
// Run:
//   cd examples/local_browser_chromedp_example
//   go mod download
//   go run main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/browserbase/stagehand-go/v3"
	"github.com/browserbase/stagehand-go/v3/option"
	"github.com/browserbase/stagehand-go/v3/packages/ssestream"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func main() {
	ctx := context.Background()

	// 1) Launch a local browser with chromedp on a fixed debugging port.
	debugPort := "9222"
	fmt.Println("Launching local browser with chromedp...")
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("remote-debugging-port", debugPort),
	)

	allocatorCtx, cancelAllocator := chromedp.NewExecAllocator(ctx, opts...)
	defer cancelAllocator()

	browserCtx, cancelBrowser := chromedp.NewContext(allocatorCtx,
		chromedp.WithErrorf(func(format string, args ...interface{}) {}),
	)
	defer cancelBrowser()

	fmt.Println("Navigating to example.com...")
	if err := chromedp.Run(
		browserCtx,
		chromedp.Navigate("https://example.com"),
		chromedp.WaitReady("body", chromedp.ByQuery),
	); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	wsURL, err := getBrowserWebSocketURL(debugPort)
	if err != nil {
		log.Fatalf("Failed to get browser websocket URL: %v", err)
	}
	fmt.Printf("Browser WebSocket URL: %s\n", wsURL)

	// 2) Create Stagehand client in local mode and connect to the existing browser.
	client := stagehand.NewClient(option.WithServer("local"))
	defer client.Close()

	startResponse, err := client.Sessions.Start(ctx, stagehand.SessionStartParams{
		ModelName: "openai/gpt-5-nano",
		Browser: stagehand.SessionStartParamsBrowser{
			Type: "local",
			LaunchOptions: stagehand.SessionStartParamsBrowserLaunchOptions{
				CdpURL: stagehand.String(wsURL),
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	sessionID := startResponse.Data.SessionID
	fmt.Printf("Session started: %s\n", sessionID)

	// 3) Navigate with Stagehand to ensure it's on the same page.
	_, err = client.Sessions.Navigate(
		ctx,
		sessionID,
		stagehand.SessionNavigateParams{
			URL: "https://example.com",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// 4) Use chromedp to click a link in the same tab.
	allocatorCtx2, cancelAllocator2 := chromedp.NewRemoteAllocator(context.Background(), wsURL, chromedp.NoModifyURL)
	defer cancelAllocator2()

	browserCtx2, cancelBrowser2 := chromedp.NewContext(allocatorCtx2,
		chromedp.WithErrorf(func(format string, args ...interface{}) {}),
	)
	defer cancelBrowser2()

	targetID, err := waitForTarget(browserCtx2, "example.com", 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	
	tabCtx, cancelTab := chromedp.NewContext(browserCtx2, chromedp.WithTargetID(targetID))
	defer cancelTab()

	err = chromedp.Run(
		tabCtx,
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.Click("a", chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Chromedp click completed")

	// 5) Observe with SSE streaming.
	fmt.Println("Observing with SSE...")
	observeStream := client.Sessions.ObserveStreaming(
		ctx,
		sessionID,
		stagehand.SessionObserveParams{
			Instruction: stagehand.String("Find the most relevant click target on this page"),
		},
	)
	observeResult, err := consumeStream("observe", observeStream)
	if err != nil {
		log.Fatal(err)
	}
	observeCount, err := countSlice(observeResult)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Observed %d possible actions\n", observeCount)

	// 6) Act with SSE streaming.
	fmt.Println("Acting with SSE...")
	actStream := client.Sessions.ActStreaming(
		ctx,
		sessionID,
		stagehand.SessionActParams{
			Input: stagehand.SessionActParamsInputUnion{
				OfString: stagehand.String("Click the 'Learn more' link"),
			},
		},
	)
	_, err = consumeStream("act", actStream)
	if err != nil {
		log.Fatal(err)
	}

	// 7) Extract with SSE streaming.
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string"},
			"url":   map[string]any{"type": "string"},
		},
	}
	fmt.Println("Extracting with SSE...")
	extractStream := client.Sessions.ExtractStreaming(
		ctx,
		sessionID,
		stagehand.SessionExtractParams{
			Instruction: stagehand.String("Extract the page title and current URL"),
			Schema:      schema,
		},
	)
	extractResult, err := consumeStream("extract", extractStream)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Extracted: %+v\n", extractResult)

	// 8) Execute with SSE streaming.
	fmt.Println("Running autonomous agent with SSE...")
	executeStream := client.Sessions.ExecuteStreaming(
		ctx,
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
	executeResult, err := consumeStream("execute", executeStream)
	if err != nil {
		log.Fatal(err)
	}
	executeSummary, err := parseExecuteResult(executeResult)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Agent completed: %s\n", executeSummary.Message)
	fmt.Printf("Agent success: %t\n", executeSummary.Success)
	fmt.Printf("Agent actions taken: %d\n", executeSummary.Actions)

	// 9) End session.
	_, err = client.Sessions.End(ctx, sessionID, stagehand.SessionEndParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Session ended")
}

type executeSummary struct {
	Message string
	Success bool
	Actions int
}

func consumeStream(label string, stream *ssestream.Stream[stagehand.StreamEvent]) (any, error) {
	var result any
	for stream.Next() {
		event := stream.Current()
		fmt.Printf("[%s][%s] %s\n", label, event.Type, event.Data.RawJSON())
		if event.Type == stagehand.StreamEventTypeSystem {
			system := event.Data.AsStreamEventDataStreamEventSystemDataOutput()
			if system.JSON.Result.Valid() {
				result = system.Result
			}
			if system.Status == "error" {
				if system.Error != "" {
					return result, fmt.Errorf("stream error: %s", system.Error)
				}
				return result, fmt.Errorf("stream error: unknown error")
			}
		}
	}
	if err := stream.Err(); err != nil {
		return result, err
	}
	if result == nil {
		return result, fmt.Errorf("stream finished without result")
	}
	return result, nil
}

func parseExecuteResult(result any) (executeSummary, error) {
	var summary executeSummary
	raw, err := json.Marshal(result)
	if err != nil {
		return summary, err
	}
	var payload struct {
		Message string `json:"message"`
		Success bool   `json:"success"`
		Actions []any  `json:"actions"`
	}
	if err := json.Unmarshal(raw, &payload); err != nil {
		return summary, err
	}
	summary.Message = payload.Message
	summary.Success = payload.Success
	summary.Actions = len(payload.Actions)
	return summary, nil
}

func countSlice(result any) (int, error) {
	raw, err := json.Marshal(result)
	if err != nil {
		return 0, err
	}
	var items []any
	if err := json.Unmarshal(raw, &items); err != nil {
		return 0, err
	}
	return len(items), nil
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

// getBrowserWebSocketURL fetches the browser's websocket URL from Chrome's debug endpoint
func getBrowserWebSocketURL(port string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/json/version", port))
	if err != nil {
		return "", fmt.Errorf("failed to fetch /json/version: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if result.WebSocketDebuggerURL == "" {
		return "", fmt.Errorf("webSocketDebuggerUrl not found in response")
	}

	return result.WebSocketDebuggerURL, nil
}
