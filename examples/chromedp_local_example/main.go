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
//   - Set MODEL_API_KEY or OPENAI_API_KEY environment variable
//   - Chrome/Chromium installed locally
//
// Run: go run examples/chromedp_local_example/main.go
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
	"github.com/chromedp/chromedp"
)

func main() {
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
		ModelName: "openai/gpt-5-nano",
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

	// 3) Use Stagehand AI to observe available actions
	fmt.Println("Using Stagehand to observe page...")
	observeResp, err := client.Sessions.Observe(ctx, sessionID, stagehand.SessionObserveParams{
		Instruction: stagehand.String("Find all clickable links on this page"),
	})
	if err != nil {
		log.Fatalf("Failed to observe: %v", err)
	}
	fmt.Printf("Found %d possible actions\n", len(observeResp.Data.Result))

	// 4) Use Stagehand AI to extract data
	fmt.Println("Using Stagehand to extract data...")
	extractResp, err := client.Sessions.Extract(ctx, sessionID, stagehand.SessionExtractParams{
		Instruction: stagehand.String("Extract the main heading and any links on this page"),
		Schema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"heading": map[string]any{"type": "string"},
				"links":   map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to extract: %v", err)
	}
	fmt.Printf("Extracted: %+v\n", extractResp.Data.Result)

	// 5) Run autonomous agent with Stagehand
	fmt.Println("Running Stagehand autonomous agent...")
	executeResp, err := client.Sessions.Execute(ctx, sessionID, stagehand.SessionExecuteParams{
		ExecuteOptions: stagehand.SessionExecuteParamsExecuteOptions{
			Instruction: "Click on the 'Learn more' link if available",
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
	fmt.Printf("Agent success: %t\n", executeResp.Data.Result.Success)

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
