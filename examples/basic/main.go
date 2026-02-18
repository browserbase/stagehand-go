// Example: Cloud mode with Browserbase.
//
// Prerequisites:
//   - Set BROWSERBASE_API_KEY
//   - Set BROWSERBASE_PROJECT_ID
//   - Set MODEL_API_KEY
//
// Run:
//
//	cd examples/basic
//	go mod download
//	go run main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/browserbase/stagehand-go/v3"
	"github.com/browserbase/stagehand-go/v3/packages/ssestream"
)

func main() {
	loadExampleEnv()
	client := stagehand.NewClient() // Uses env vars: BROWSERBASE_API_KEY, BROWSERBASE_PROJECT_ID, MODEL_API_KEY

	startResponse, err := client.Sessions.Start(context.TODO(), stagehand.SessionStartParams{
		ModelName: "anthropic/claude-sonnet-4-6",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Session started: %s\n", startResponse.Data.SessionID)

	sessionID := startResponse.Data.SessionID

	// Navigate to Hacker News
	_, err = client.Sessions.Navigate(
		context.TODO(),
		sessionID,
		stagehand.SessionNavigateParams{
			URL: "https://news.ycombinator.com",
		},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Navigated to Hacker News")

	// Observe to find possible actions (streaming)
	observeStream := client.Sessions.ObserveStreaming(
		context.TODO(),
		sessionID,
		stagehand.SessionObserveParams{
			Instruction: stagehand.String("find the link to view comments for the top post"),
		},
	)
	observeResult, err := consumeStream("observe", observeStream, true)
	if err != nil {
		panic(err.Error())
	}
	actions, err := parseActions(observeResult)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Found %d possible actions\n", len(actions))

	if len(actions) == 0 {
		fmt.Println("No actions found")
		return
	}

	// Use the first action
	action := actions[0]
	fmt.Printf("Acting on: %s\n", action.Description)

	// Pass the action to Act
	actStream := client.Sessions.ActStreaming(
		context.TODO(),
		sessionID,
		stagehand.SessionActParams{
			Input: stagehand.SessionActParamsInputUnion{
				OfAction: &stagehand.ActionParam{
					Description: action.Description,
					Selector:    action.Selector,
					Method:      stagehand.String(action.Method),
					Arguments:   action.Arguments,
				},
			},
		},
	)
	_, err = consumeStream("act", actStream, false)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Act completed")

	// Extract data from the page
	// We're now on the comments page, so extract the top comment text
	extractStream := client.Sessions.ExtractStreaming(
		context.TODO(),
		sessionID,
		stagehand.SessionExtractParams{
			Instruction: stagehand.String("extract the text of the top comment on this page"),
			Schema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"commentText": map[string]any{
						"type":        "string",
						"description": "The text content of the top comment",
					},
					"author": map[string]any{
						"type":        "string",
						"description": "The username of the comment author",
					},
				},
				"required": []string{"commentText"},
			},
		},
	)
	extractResult, err := consumeStream("extract", extractStream, true)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Extracted data: %+v\n", extractResult)

	// Get the author from the extracted data
	extractedData, err := parseMap(extractResult)
	if err != nil {
		panic(err.Error())
	}
	author, _ := extractedData["author"].(string)
	fmt.Printf("Looking up profile for author: %s\n", author)

	// Use the Agent to find the author's profile
	// Execute runs an autonomous agent that can navigate and interact with pages
	executeStream := client.Sessions.ExecuteStreaming(
		context.TODO(),
		sessionID,
		stagehand.SessionExecuteParams{
			ExecuteOptions: stagehand.SessionExecuteParamsExecuteOptions{
				Instruction: fmt.Sprintf(
					"Find any personal website, GitHub, LinkedIn, or other best profile URL for the Hacker News user '%s'. "+
						"Click on their username to go to their profile page and look for any links they have shared. "+
						"Use Google Search with their username or other details from their profile if you dont find any direct links.",
					author,
				),
				MaxSteps: stagehand.Float(15),
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
		},
	)
	executeResult, err := consumeStream("execute", executeStream, false)
	if err != nil {
		panic(err.Error())
	}
	if executeResult != nil {
		executeSummary, err := parseExecuteResult(executeResult)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Agent completed: %s\n", executeSummary.Message)
		fmt.Printf("Agent success: %t\n", executeSummary.Success)
		fmt.Printf("Agent actions taken: %d\n", executeSummary.Actions)
	}

	// End the session to clean up resources
	_, err = client.Sessions.End(
		context.TODO(),
		sessionID,
		stagehand.SessionEndParams{},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Session ended")
}

type executeSummary struct {
	Message string
	Success bool
	Actions int
}

type actionSummary struct {
	Description string   `json:"description"`
	Selector    string   `json:"selector"`
	Method      string   `json:"method"`
	Arguments   []string `json:"arguments"`
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

func consumeStream(
	label string,
	stream *ssestream.Stream[stagehand.StreamEvent],
	requireResult bool,
) (any, error) {
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
		if requireResult {
			return result, fmt.Errorf("stream finished without result")
		}
		return result, nil
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

func parseActions(result any) ([]actionSummary, error) {
	raw, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	var actions []actionSummary
	if err := json.Unmarshal(raw, &actions); err != nil {
		return nil, err
	}
	return actions, nil
}

func parseMap(result any) (map[string]any, error) {
	switch value := result.(type) {
	case string:
		var payload map[string]any
		if err := json.Unmarshal([]byte(value), &payload); err == nil {
			return payload, nil
		}
		// If the stream result was a truncated string, keep it as commentText.
		return map[string]any{
			"commentText": value,
			"author":      "unknown",
		}, nil
	}

	raw, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil, err
	}
	return payload, nil
}
