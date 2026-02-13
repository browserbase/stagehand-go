// Example: Running Stagehand locally
//
// This example demonstrates how to use Stagehand with a local browser
// instead of connecting to the Browserbase cloud.
//
// Prerequisites:
//   - Set MODEL_API_KEY environment variable
//
// Run:
//
//	cd examples/local
//	go mod download
//	go run main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	stagehand "github.com/browserbase/stagehand-go/v3"
	"github.com/browserbase/stagehand-go/v3/option"
	"github.com/browserbase/stagehand-go/v3/packages/ssestream"
)

func main() {
	loadExampleEnv()
	// Create a client in local mode.
	client := stagehand.NewClient(option.WithServer("local"))
	defer client.Close()

	ctx := context.Background()

	// Start a new session with a local browser.
	fmt.Println("Starting local session...")
	startResp, err := client.Sessions.Start(ctx, stagehand.SessionStartParams{
		ModelName: "openai/gpt-5-nano",
		Browser: stagehand.SessionStartParamsBrowser{
			Type: "local",
			LaunchOptions: stagehand.SessionStartParamsBrowserLaunchOptions{
				Headless: stagehand.Bool(false),
			},
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start session: %v\n", err)
		os.Exit(1)
	}
	sessionID := startResp.Data.SessionID
	fmt.Printf("Session started: %s\n", sessionID)

	// Navigate to a page
	fmt.Println("Navigating to example.com...")
	_, err = client.Sessions.Navigate(ctx, sessionID, stagehand.SessionNavigateParams{
		URL: "https://example.com",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to navigate: %v\n", err)
		os.Exit(1)
	}

	// Extract the page title using natural language
	fmt.Println("Extracting page title...")
	extractStream := client.Sessions.ExtractStreaming(ctx, sessionID, stagehand.SessionExtractParams{
		Instruction: stagehand.String("extract the main heading from this page"),
		Schema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"title": map[string]any{
					"type":        "string",
					"description": "The main heading text",
				},
			},
			"required": []string{"title"},
		},
	})
	extractResult, err := consumeStream("extract", extractStream)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to extract: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Extracted: %+v\n", extractResult)

	// End the session
	fmt.Println("Ending session...")
	_, err = client.Sessions.End(ctx, sessionID, stagehand.SessionEndParams{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to end session: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Session ended successfully")
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
