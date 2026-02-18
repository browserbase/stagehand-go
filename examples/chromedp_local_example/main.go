// Example: Using chromedp with Stagehand local mode
//
// This example demonstrates how to combine chromedp (for low-level browser control)
// with Stagehand (for AI-powered actions) using a local browser.
//
// Architecture:
//   - chromedp launches and controls the browser
//   - Stagehand connects to the same browser via CDP URL
//   - Both can interact with the same page
//
// Prerequisites:
//   - Set MODEL_API_KEY environment variable
//   - Chrome/Chromium installed locally
//
// Run:
//
//	cd examples/chromedp_local_example
//	go mod download
//	go run main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/browserbase/stagehand-go/v3"
	"github.com/browserbase/stagehand-go/v3/option"
	"github.com/browserbase/stagehand-go/v3/packages/ssestream"
	"github.com/chromedp/chromedp"
)

func main() {
	loadExampleEnv()
	ctx := context.Background()

	// 1) Launch browser with chromedp on a specific debugging port
	fmt.Println("Launching browser with chromedp...")
	debugPort := "9222"

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("remote-debugging-port", debugPort),
	)

	allocatorCtx, cancelAllocator := chromedp.NewExecAllocator(ctx, opts...)
	defer cancelAllocator()

	browserCtx, cancelBrowser := chromedp.NewContext(allocatorCtx,
		chromedp.WithErrorf(func(format string, args ...interface{}) {}), // Suppress warnings
	)
	defer cancelBrowser()

	// Navigate to get the browser started
	fmt.Println("Navigating to example.com...")
	err := chromedp.Run(
		browserCtx,
		chromedp.Navigate("https://example.com"),
		chromedp.WaitReady("body", chromedp.ByQuery),
	)
	if err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	// Fetch the browser's websocket URL from Chrome's debug endpoint
	wsURL, err := getBrowserWebSocketURL(debugPort)
	if err != nil {
		log.Fatalf("Failed to get browser websocket URL: %v", err)
	}
	fmt.Printf("Browser WebSocket URL: %s\n", wsURL)

	// 2) Create Stagehand client in local mode and connect to existing browser
	fmt.Println("Connecting Stagehand to existing browser...")
	client := stagehand.NewClient(option.WithServer("local"))
	defer client.Close()

	// Start a session connected to the chromedp browser
	startResp, err := client.Sessions.Start(ctx, stagehand.SessionStartParams{
		ModelName: "anthropic/claude-sonnet-4-6",
		Browser: stagehand.SessionStartParamsBrowser{
			Type: "local",
			LaunchOptions: stagehand.SessionStartParamsBrowserLaunchOptions{
				CdpURL: stagehand.String(wsURL),
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to start session: %v", err)
	}
	sessionID := startResp.Data.SessionID
	fmt.Printf("Session started: %s\n", sessionID)

	// Navigate with Stagehand to ensure it's on the same page
	fmt.Println("Navigating with Stagehand to ensure same page...")
	_, err = client.Sessions.Navigate(ctx, sessionID, stagehand.SessionNavigateParams{
		URL: "https://example.com",
	})
	if err != nil {
		log.Fatalf("Failed to navigate with Stagehand: %v", err)
	}

	// 3) Use Stagehand AI to observe available actions (streaming)
	fmt.Println("Using Stagehand to observe page...")
	observeStream := client.Sessions.ObserveStreaming(ctx, sessionID, stagehand.SessionObserveParams{
		Instruction: stagehand.String("Find all clickable links on this page"),
	})
	observeResult, err := consumeStream("observe", observeStream)
	if err != nil {
		log.Fatalf("Failed to observe: %v", err)
	}
	observeCount, err := countSlice(observeResult)
	if err != nil {
		log.Fatalf("Failed to parse observe result: %v", err)
	}
	fmt.Printf("Found %d possible actions\n", observeCount)

	// 4) Use Stagehand AI to extract data (streaming)
	fmt.Println("Using Stagehand to extract data...")
	extractStream := client.Sessions.ExtractStreaming(ctx, sessionID, stagehand.SessionExtractParams{
		Instruction: stagehand.String("Extract the main heading and any links on this page"),
		Schema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"heading": map[string]any{"type": "string"},
				"links":   map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
			},
		},
	})
	extractResult, err := consumeStream("extract", extractStream)
	if err != nil {
		log.Fatalf("Failed to extract: %v", err)
	}
	fmt.Printf("Extracted: %+v\n", extractResult)

	// 5) Run autonomous agent with Stagehand (streaming)
	fmt.Println("Running Stagehand autonomous agent...")
	executeStream := client.Sessions.ExecuteStreaming(ctx, sessionID, stagehand.SessionExecuteParams{
		ExecuteOptions: stagehand.SessionExecuteParamsExecuteOptions{
			Instruction: "Click on the 'Learn more' link if available",
			MaxSteps:    stagehand.Float(3),
		},
		AgentConfig: stagehand.SessionExecuteParamsAgentConfig{
			Model: stagehand.SessionExecuteParamsAgentConfigModelUnion{
				OfModelConfig: &stagehand.ModelConfigParam{
					ModelName: "anthropic/claude-opus-4-6",
					APIKey:    stagehand.String(os.Getenv("MODEL_API_KEY")),
				},
			},
			Cua: stagehand.Bool(false),
		},
	})
	executeResult, err := consumeStream("execute", executeStream)
	if err != nil {
		log.Fatalf("Failed to execute: %v", err)
	}
	executeSummary, err := parseExecuteResult(executeResult)
	if err != nil {
		log.Fatalf("Failed to parse execute result: %v", err)
	}
	fmt.Printf("Agent completed: %s\n", executeSummary.Message)
	fmt.Printf("Agent success: %t\n", executeSummary.Success)
	fmt.Printf("Agent actions taken: %d\n", executeSummary.Actions)

	// 6) Take a screenshot with chromedp
	fmt.Println("Taking screenshot with chromedp...")
	var screenshotBuf []byte
	err = chromedp.Run(
		browserCtx,
		chromedp.Sleep(1*time.Second),
		chromedp.FullScreenshot(&screenshotBuf, 90),
	)
	if err != nil {
		log.Fatalf("Failed to take screenshot: %v", err)
	}

	screenshotPath := "screenshot_local.png"
	if err := os.WriteFile(screenshotPath, screenshotBuf, 0644); err != nil {
		log.Fatalf("Failed to save screenshot: %v", err)
	}
	fmt.Printf("Screenshot saved to: %s\n", screenshotPath)

	// 7) End the Stagehand session
	fmt.Println("Ending session...")
	_, err = client.Sessions.End(ctx, sessionID, stagehand.SessionEndParams{})
	if err != nil {
		log.Fatalf("Failed to end session: %v", err)
	}
	fmt.Println("Session ended successfully")
}

type executeSummary struct {
	Message string
	Success bool
	Actions int
}

func parseLogResult(raw string) (any, bool) {
	var payload struct {
		Message struct {
			Auxiliary struct {
				Result struct {
					Value string `json:"value"`
				} `json:"result"`
			} `json:"auxiliary"`
		} `json:"message"`
	}
	if err := json.Unmarshal([]byte(raw), &payload); err != nil {
		return nil, false
	}
	if payload.Message.Auxiliary.Result.Value == "" {
		return nil, false
	}
	var result any
	if err := json.Unmarshal([]byte(payload.Message.Auxiliary.Result.Value), &result); err != nil {
		return payload.Message.Auxiliary.Result.Value, true
	}
	return result, true
}

func parseObservationElements(raw string) (any, bool) {
	var payload struct {
		Message struct {
			Category  string `json:"category"`
			Message   string `json:"message"`
			Auxiliary struct {
				Elements struct {
					Value string `json:"value"`
				} `json:"elements"`
			} `json:"auxiliary"`
		} `json:"message"`
	}
	if err := json.Unmarshal([]byte(raw), &payload); err != nil {
		return nil, false
	}
	if payload.Message.Category != "observation" || payload.Message.Message != "found elements" {
		return nil, false
	}
	if payload.Message.Auxiliary.Elements.Value == "" {
		return nil, false
	}
	var elements any
	if err := json.Unmarshal([]byte(payload.Message.Auxiliary.Elements.Value), &elements); err != nil {
		return nil, false
	}
	return elements, true
}

func consumeStream(label string, stream *ssestream.Stream[stagehand.StreamEvent]) (any, error) {
	var result any
	for stream.Next() {
		event := stream.Current()
		fmt.Printf("[%s][%s] %s\n", label, event.Type, event.Data.RawJSON())
		if event.Type == stagehand.StreamEventTypeLog && result == nil {
			if elements, ok := parseObservationElements(event.Data.RawJSON()); ok {
				result = elements
			}
		}
		if result == nil {
			if parsed, ok := parseLogResult(event.Data.RawJSON()); ok {
				result = parsed
			}
		}
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
		if label == "act" {
			return result, nil
		}
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
