// Example: Running Stagehand locally
//
// This example demonstrates how to use Stagehand with a local browser
// instead of connecting to the Browserbase cloud.
//
// Prerequisites:
//   - Set MODEL_API_KEY environment variable
//
// Run:
//   cd examples/local
//   go mod download
//   go run main.go
package main

import (
	"context"
	"fmt"
	"os"

	stagehand "github.com/browserbase/stagehand-go/v3"
	"github.com/browserbase/stagehand-go/v3/option"
)

func main() {
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
	extractResp, err := client.Sessions.Extract(ctx, sessionID, stagehand.SessionExtractParams{
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
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to extract: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Extracted: %+v\n", extractResp.Data.Result)

	// End the session
	fmt.Println("Ending session...")
	_, err = client.Sessions.End(ctx, sessionID, stagehand.SessionEndParams{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to end session: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Session ended successfully")
}
